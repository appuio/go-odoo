package odoo

import (
	"fmt"
)

// SaveSpreadsheetTemplate represents save.spreadsheet.template model.
type SaveSpreadsheetTemplate struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	Data         *String   `xmlrpc:"data,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	TemplateName *String   `xmlrpc:"template_name,omptempty"`
	Thumbnail    *String   `xmlrpc:"thumbnail,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaveSpreadsheetTemplates represents array of save.spreadsheet.template model.
type SaveSpreadsheetTemplates []SaveSpreadsheetTemplate

// SaveSpreadsheetTemplateModel is the odoo model name.
const SaveSpreadsheetTemplateModel = "save.spreadsheet.template"

// Many2One convert SaveSpreadsheetTemplate to *Many2One.
func (sst *SaveSpreadsheetTemplate) Many2One() *Many2One {
	return NewMany2One(sst.Id.Get(), "")
}

// CreateSaveSpreadsheetTemplate creates a new save.spreadsheet.template model and returns its id.
func (c *Client) CreateSaveSpreadsheetTemplate(sst *SaveSpreadsheetTemplate) (int64, error) {
	ids, err := c.CreateSaveSpreadsheetTemplates([]*SaveSpreadsheetTemplate{sst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaveSpreadsheetTemplate creates a new save.spreadsheet.template model and returns its id.
func (c *Client) CreateSaveSpreadsheetTemplates(ssts []*SaveSpreadsheetTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssts {
		vv = append(vv, v)
	}
	return c.Create(SaveSpreadsheetTemplateModel, vv)
}

// UpdateSaveSpreadsheetTemplate updates an existing save.spreadsheet.template record.
func (c *Client) UpdateSaveSpreadsheetTemplate(sst *SaveSpreadsheetTemplate) error {
	return c.UpdateSaveSpreadsheetTemplates([]int64{sst.Id.Get()}, sst)
}

// UpdateSaveSpreadsheetTemplates updates existing save.spreadsheet.template records.
// All records (represented by ids) will be updated by sst values.
func (c *Client) UpdateSaveSpreadsheetTemplates(ids []int64, sst *SaveSpreadsheetTemplate) error {
	return c.Update(SaveSpreadsheetTemplateModel, ids, sst)
}

// DeleteSaveSpreadsheetTemplate deletes an existing save.spreadsheet.template record.
func (c *Client) DeleteSaveSpreadsheetTemplate(id int64) error {
	return c.DeleteSaveSpreadsheetTemplates([]int64{id})
}

// DeleteSaveSpreadsheetTemplates deletes existing save.spreadsheet.template records.
func (c *Client) DeleteSaveSpreadsheetTemplates(ids []int64) error {
	return c.Delete(SaveSpreadsheetTemplateModel, ids)
}

// GetSaveSpreadsheetTemplate gets save.spreadsheet.template existing record.
func (c *Client) GetSaveSpreadsheetTemplate(id int64) (*SaveSpreadsheetTemplate, error) {
	ssts, err := c.GetSaveSpreadsheetTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssts != nil && len(*ssts) > 0 {
		return &((*ssts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of save.spreadsheet.template not found", id)
}

// GetSaveSpreadsheetTemplates gets save.spreadsheet.template existing records.
func (c *Client) GetSaveSpreadsheetTemplates(ids []int64) (*SaveSpreadsheetTemplates, error) {
	ssts := &SaveSpreadsheetTemplates{}
	if err := c.Read(SaveSpreadsheetTemplateModel, ids, nil, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSaveSpreadsheetTemplate finds save.spreadsheet.template record by querying it with criteria.
func (c *Client) FindSaveSpreadsheetTemplate(criteria *Criteria) (*SaveSpreadsheetTemplate, error) {
	ssts := &SaveSpreadsheetTemplates{}
	if err := c.SearchRead(SaveSpreadsheetTemplateModel, criteria, NewOptions().Limit(1), ssts); err != nil {
		return nil, err
	}
	if ssts != nil && len(*ssts) > 0 {
		return &((*ssts)[0]), nil
	}
	return nil, fmt.Errorf("save.spreadsheet.template was not found with criteria %v", criteria)
}

// FindSaveSpreadsheetTemplates finds save.spreadsheet.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaveSpreadsheetTemplates(criteria *Criteria, options *Options) (*SaveSpreadsheetTemplates, error) {
	ssts := &SaveSpreadsheetTemplates{}
	if err := c.SearchRead(SaveSpreadsheetTemplateModel, criteria, options, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSaveSpreadsheetTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaveSpreadsheetTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaveSpreadsheetTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaveSpreadsheetTemplateId finds record id by querying it with criteria.
func (c *Client) FindSaveSpreadsheetTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaveSpreadsheetTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("save.spreadsheet.template was not found with criteria %v and options %v", criteria, options)
}
