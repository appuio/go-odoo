package odoo

import (
	"fmt"
)

// MukRestCallback represents muk_rest.callback model.
type MukRestCallback struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	OauthId     *Many2One `xmlrpc:"oauth_id,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	Url         *String   `xmlrpc:"url,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MukRestCallbacks represents array of muk_rest.callback model.
type MukRestCallbacks []MukRestCallback

// MukRestCallbackModel is the odoo model name.
const MukRestCallbackModel = "muk_rest.callback"

// Many2One convert MukRestCallback to *Many2One.
func (mc *MukRestCallback) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMukRestCallback creates a new muk_rest.callback model and returns its id.
func (c *Client) CreateMukRestCallback(mc *MukRestCallback) (int64, error) {
	ids, err := c.CreateMukRestCallbacks([]*MukRestCallback{mc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestCallback creates a new muk_rest.callback model and returns its id.
func (c *Client) CreateMukRestCallbacks(mcs []*MukRestCallback) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcs {
		vv = append(vv, v)
	}
	return c.Create(MukRestCallbackModel, vv)
}

// UpdateMukRestCallback updates an existing muk_rest.callback record.
func (c *Client) UpdateMukRestCallback(mc *MukRestCallback) error {
	return c.UpdateMukRestCallbacks([]int64{mc.Id.Get()}, mc)
}

// UpdateMukRestCallbacks updates existing muk_rest.callback records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMukRestCallbacks(ids []int64, mc *MukRestCallback) error {
	return c.Update(MukRestCallbackModel, ids, mc)
}

// DeleteMukRestCallback deletes an existing muk_rest.callback record.
func (c *Client) DeleteMukRestCallback(id int64) error {
	return c.DeleteMukRestCallbacks([]int64{id})
}

// DeleteMukRestCallbacks deletes existing muk_rest.callback records.
func (c *Client) DeleteMukRestCallbacks(ids []int64) error {
	return c.Delete(MukRestCallbackModel, ids)
}

// GetMukRestCallback gets muk_rest.callback existing record.
func (c *Client) GetMukRestCallback(id int64) (*MukRestCallback, error) {
	mcs, err := c.GetMukRestCallbacks([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.callback not found", id)
}

// GetMukRestCallbacks gets muk_rest.callback existing records.
func (c *Client) GetMukRestCallbacks(ids []int64) (*MukRestCallbacks, error) {
	mcs := &MukRestCallbacks{}
	if err := c.Read(MukRestCallbackModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMukRestCallback finds muk_rest.callback record by querying it with criteria.
func (c *Client) FindMukRestCallback(criteria *Criteria) (*MukRestCallback, error) {
	mcs := &MukRestCallbacks{}
	if err := c.SearchRead(MukRestCallbackModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.callback was not found with criteria %v", criteria)
}

// FindMukRestCallbacks finds muk_rest.callback records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestCallbacks(criteria *Criteria, options *Options) (*MukRestCallbacks, error) {
	mcs := &MukRestCallbacks{}
	if err := c.SearchRead(MukRestCallbackModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMukRestCallbackIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestCallbackIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestCallbackModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestCallbackId finds record id by querying it with criteria.
func (c *Client) FindMukRestCallbackId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestCallbackModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.callback was not found with criteria %v and options %v", criteria, options)
}
