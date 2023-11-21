package odoo

import (
	"fmt"
)

// HelpdeskTag represents helpdesk.tag model.
type HelpdeskTag struct {
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

// HelpdeskTags represents array of helpdesk.tag model.
type HelpdeskTags []HelpdeskTag

// HelpdeskTagModel is the odoo model name.
const HelpdeskTagModel = "helpdesk.tag"

// Many2One convert HelpdeskTag to *Many2One.
func (ht *HelpdeskTag) Many2One() *Many2One {
	return NewMany2One(ht.Id.Get(), "")
}

// CreateHelpdeskTag creates a new helpdesk.tag model and returns its id.
func (c *Client) CreateHelpdeskTag(ht *HelpdeskTag) (int64, error) {
	ids, err := c.CreateHelpdeskTags([]*HelpdeskTag{ht})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTag creates a new helpdesk.tag model and returns its id.
func (c *Client) CreateHelpdeskTags(hts []*HelpdeskTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range hts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTagModel, vv)
}

// UpdateHelpdeskTag updates an existing helpdesk.tag record.
func (c *Client) UpdateHelpdeskTag(ht *HelpdeskTag) error {
	return c.UpdateHelpdeskTags([]int64{ht.Id.Get()}, ht)
}

// UpdateHelpdeskTags updates existing helpdesk.tag records.
// All records (represented by ids) will be updated by ht values.
func (c *Client) UpdateHelpdeskTags(ids []int64, ht *HelpdeskTag) error {
	return c.Update(HelpdeskTagModel, ids, ht)
}

// DeleteHelpdeskTag deletes an existing helpdesk.tag record.
func (c *Client) DeleteHelpdeskTag(id int64) error {
	return c.DeleteHelpdeskTags([]int64{id})
}

// DeleteHelpdeskTags deletes existing helpdesk.tag records.
func (c *Client) DeleteHelpdeskTags(ids []int64) error {
	return c.Delete(HelpdeskTagModel, ids)
}

// GetHelpdeskTag gets helpdesk.tag existing record.
func (c *Client) GetHelpdeskTag(id int64) (*HelpdeskTag, error) {
	hts, err := c.GetHelpdeskTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if hts != nil && len(*hts) > 0 {
		return &((*hts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.tag not found", id)
}

// GetHelpdeskTags gets helpdesk.tag existing records.
func (c *Client) GetHelpdeskTags(ids []int64) (*HelpdeskTags, error) {
	hts := &HelpdeskTags{}
	if err := c.Read(HelpdeskTagModel, ids, nil, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTag finds helpdesk.tag record by querying it with criteria.
func (c *Client) FindHelpdeskTag(criteria *Criteria) (*HelpdeskTag, error) {
	hts := &HelpdeskTags{}
	if err := c.SearchRead(HelpdeskTagModel, criteria, NewOptions().Limit(1), hts); err != nil {
		return nil, err
	}
	if hts != nil && len(*hts) > 0 {
		return &((*hts)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.tag was not found with criteria %v", criteria)
}

// FindHelpdeskTags finds helpdesk.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTags(criteria *Criteria, options *Options) (*HelpdeskTags, error) {
	hts := &HelpdeskTags{}
	if err := c.SearchRead(HelpdeskTagModel, criteria, options, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTagId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.tag was not found with criteria %v and options %v", criteria, options)
}
