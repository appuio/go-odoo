package odoo

import (
	"fmt"
)

// SpreadsheetTemplate represents spreadsheet.template model.
type SpreadsheetTemplate struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Data        *String   `xmlrpc:"data,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	Thumbnail   *String   `xmlrpc:"thumbnail,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SpreadsheetTemplates represents array of spreadsheet.template model.
type SpreadsheetTemplates []SpreadsheetTemplate

// SpreadsheetTemplateModel is the odoo model name.
const SpreadsheetTemplateModel = "spreadsheet.template"

// Many2One convert SpreadsheetTemplate to *Many2One.
func (st *SpreadsheetTemplate) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSpreadsheetTemplate creates a new spreadsheet.template model and returns its id.
func (c *Client) CreateSpreadsheetTemplate(st *SpreadsheetTemplate) (int64, error) {
	ids, err := c.CreateSpreadsheetTemplates([]*SpreadsheetTemplate{st})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetTemplate creates a new spreadsheet.template model and returns its id.
func (c *Client) CreateSpreadsheetTemplates(sts []*SpreadsheetTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sts {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetTemplateModel, vv)
}

// UpdateSpreadsheetTemplate updates an existing spreadsheet.template record.
func (c *Client) UpdateSpreadsheetTemplate(st *SpreadsheetTemplate) error {
	return c.UpdateSpreadsheetTemplates([]int64{st.Id.Get()}, st)
}

// UpdateSpreadsheetTemplates updates existing spreadsheet.template records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSpreadsheetTemplates(ids []int64, st *SpreadsheetTemplate) error {
	return c.Update(SpreadsheetTemplateModel, ids, st)
}

// DeleteSpreadsheetTemplate deletes an existing spreadsheet.template record.
func (c *Client) DeleteSpreadsheetTemplate(id int64) error {
	return c.DeleteSpreadsheetTemplates([]int64{id})
}

// DeleteSpreadsheetTemplates deletes existing spreadsheet.template records.
func (c *Client) DeleteSpreadsheetTemplates(ids []int64) error {
	return c.Delete(SpreadsheetTemplateModel, ids)
}

// GetSpreadsheetTemplate gets spreadsheet.template existing record.
func (c *Client) GetSpreadsheetTemplate(id int64) (*SpreadsheetTemplate, error) {
	sts, err := c.GetSpreadsheetTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of spreadsheet.template not found", id)
}

// GetSpreadsheetTemplates gets spreadsheet.template existing records.
func (c *Client) GetSpreadsheetTemplates(ids []int64) (*SpreadsheetTemplates, error) {
	sts := &SpreadsheetTemplates{}
	if err := c.Read(SpreadsheetTemplateModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSpreadsheetTemplate finds spreadsheet.template record by querying it with criteria.
func (c *Client) FindSpreadsheetTemplate(criteria *Criteria) (*SpreadsheetTemplate, error) {
	sts := &SpreadsheetTemplates{}
	if err := c.SearchRead(SpreadsheetTemplateModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("spreadsheet.template was not found with criteria %v", criteria)
}

// FindSpreadsheetTemplates finds spreadsheet.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetTemplates(criteria *Criteria, options *Options) (*SpreadsheetTemplates, error) {
	sts := &SpreadsheetTemplates{}
	if err := c.SearchRead(SpreadsheetTemplateModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSpreadsheetTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SpreadsheetTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSpreadsheetTemplateId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("spreadsheet.template was not found with criteria %v and options %v", criteria, options)
}
