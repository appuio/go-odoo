package odoo

import (
	"fmt"
)

// CrmTag represents crm.tag model.
type CrmTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CrmTags represents array of crm.tag model.
type CrmTags []CrmTag

// CrmTagModel is the odoo model name.
const CrmTagModel = "crm.tag"

// Many2One convert CrmTag to *Many2One.
func (ct *CrmTag) Many2One() *Many2One {
	return NewMany2One(ct.Id.Get(), "")
}

// CreateCrmTag creates a new crm.tag model and returns its id.
func (c *Client) CreateCrmTag(ct *CrmTag) (int64, error) {
	ids, err := c.CreateCrmTags([]*CrmTag{ct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmTag creates a new crm.tag model and returns its id.
func (c *Client) CreateCrmTags(cts []*CrmTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range cts {
		vv = append(vv, v)
	}
	return c.Create(CrmTagModel, vv)
}

// UpdateCrmTag updates an existing crm.tag record.
func (c *Client) UpdateCrmTag(ct *CrmTag) error {
	return c.UpdateCrmTags([]int64{ct.Id.Get()}, ct)
}

// UpdateCrmTags updates existing crm.tag records.
// All records (represented by ids) will be updated by ct values.
func (c *Client) UpdateCrmTags(ids []int64, ct *CrmTag) error {
	return c.Update(CrmTagModel, ids, ct)
}

// DeleteCrmTag deletes an existing crm.tag record.
func (c *Client) DeleteCrmTag(id int64) error {
	return c.DeleteCrmTags([]int64{id})
}

// DeleteCrmTags deletes existing crm.tag records.
func (c *Client) DeleteCrmTags(ids []int64) error {
	return c.Delete(CrmTagModel, ids)
}

// GetCrmTag gets crm.tag existing record.
func (c *Client) GetCrmTag(id int64) (*CrmTag, error) {
	cts, err := c.GetCrmTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if cts != nil && len(*cts) > 0 {
		return &((*cts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.tag not found", id)
}

// GetCrmTags gets crm.tag existing records.
func (c *Client) GetCrmTags(ids []int64) (*CrmTags, error) {
	cts := &CrmTags{}
	if err := c.Read(CrmTagModel, ids, nil, cts); err != nil {
		return nil, err
	}
	return cts, nil
}

// FindCrmTag finds crm.tag record by querying it with criteria.
func (c *Client) FindCrmTag(criteria *Criteria) (*CrmTag, error) {
	cts := &CrmTags{}
	if err := c.SearchRead(CrmTagModel, criteria, NewOptions().Limit(1), cts); err != nil {
		return nil, err
	}
	if cts != nil && len(*cts) > 0 {
		return &((*cts)[0]), nil
	}
	return nil, fmt.Errorf("crm.tag was not found with criteria %v", criteria)
}

// FindCrmTags finds crm.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTags(criteria *Criteria, options *Options) (*CrmTags, error) {
	cts := &CrmTags{}
	if err := c.SearchRead(CrmTagModel, criteria, options, cts); err != nil {
		return nil, err
	}
	return cts, nil
}

// FindCrmTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmTagId finds record id by querying it with criteria.
func (c *Client) FindCrmTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.tag was not found with criteria %v and options %v", criteria, options)
}
