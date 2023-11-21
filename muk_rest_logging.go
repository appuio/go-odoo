package odoo

import (
	"fmt"
)

// MukRestLogging represents muk_rest.logging model.
type MukRestLogging struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	IpAddress   *String   `xmlrpc:"ip_address,omptempty"`
	Method      *String   `xmlrpc:"method,omptempty"`
	Request     *String   `xmlrpc:"request,omptempty"`
	Response    *String   `xmlrpc:"response,omptempty"`
	Status      *String   `xmlrpc:"status,omptempty"`
	Url         *String   `xmlrpc:"url,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MukRestLoggings represents array of muk_rest.logging model.
type MukRestLoggings []MukRestLogging

// MukRestLoggingModel is the odoo model name.
const MukRestLoggingModel = "muk_rest.logging"

// Many2One convert MukRestLogging to *Many2One.
func (ml *MukRestLogging) Many2One() *Many2One {
	return NewMany2One(ml.Id.Get(), "")
}

// CreateMukRestLogging creates a new muk_rest.logging model and returns its id.
func (c *Client) CreateMukRestLogging(ml *MukRestLogging) (int64, error) {
	ids, err := c.CreateMukRestLoggings([]*MukRestLogging{ml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestLogging creates a new muk_rest.logging model and returns its id.
func (c *Client) CreateMukRestLoggings(mls []*MukRestLogging) ([]int64, error) {
	var vv []interface{}
	for _, v := range mls {
		vv = append(vv, v)
	}
	return c.Create(MukRestLoggingModel, vv)
}

// UpdateMukRestLogging updates an existing muk_rest.logging record.
func (c *Client) UpdateMukRestLogging(ml *MukRestLogging) error {
	return c.UpdateMukRestLoggings([]int64{ml.Id.Get()}, ml)
}

// UpdateMukRestLoggings updates existing muk_rest.logging records.
// All records (represented by ids) will be updated by ml values.
func (c *Client) UpdateMukRestLoggings(ids []int64, ml *MukRestLogging) error {
	return c.Update(MukRestLoggingModel, ids, ml)
}

// DeleteMukRestLogging deletes an existing muk_rest.logging record.
func (c *Client) DeleteMukRestLogging(id int64) error {
	return c.DeleteMukRestLoggings([]int64{id})
}

// DeleteMukRestLoggings deletes existing muk_rest.logging records.
func (c *Client) DeleteMukRestLoggings(ids []int64) error {
	return c.Delete(MukRestLoggingModel, ids)
}

// GetMukRestLogging gets muk_rest.logging existing record.
func (c *Client) GetMukRestLogging(id int64) (*MukRestLogging, error) {
	mls, err := c.GetMukRestLoggings([]int64{id})
	if err != nil {
		return nil, err
	}
	if mls != nil && len(*mls) > 0 {
		return &((*mls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.logging not found", id)
}

// GetMukRestLoggings gets muk_rest.logging existing records.
func (c *Client) GetMukRestLoggings(ids []int64) (*MukRestLoggings, error) {
	mls := &MukRestLoggings{}
	if err := c.Read(MukRestLoggingModel, ids, nil, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMukRestLogging finds muk_rest.logging record by querying it with criteria.
func (c *Client) FindMukRestLogging(criteria *Criteria) (*MukRestLogging, error) {
	mls := &MukRestLoggings{}
	if err := c.SearchRead(MukRestLoggingModel, criteria, NewOptions().Limit(1), mls); err != nil {
		return nil, err
	}
	if mls != nil && len(*mls) > 0 {
		return &((*mls)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.logging was not found with criteria %v", criteria)
}

// FindMukRestLoggings finds muk_rest.logging records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestLoggings(criteria *Criteria, options *Options) (*MukRestLoggings, error) {
	mls := &MukRestLoggings{}
	if err := c.SearchRead(MukRestLoggingModel, criteria, options, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMukRestLoggingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestLoggingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestLoggingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestLoggingId finds record id by querying it with criteria.
func (c *Client) FindMukRestLoggingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestLoggingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.logging was not found with criteria %v and options %v", criteria, options)
}
