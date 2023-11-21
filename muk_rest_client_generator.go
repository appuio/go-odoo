package odoo

import (
	"fmt"
)

// MukRestClientGenerator represents muk_rest.client_generator model.
type MukRestClientGenerator struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Language    *Selection `xmlrpc:"language,omptempty"`
	Options     *String    `xmlrpc:"options,omptempty"`
	SendOptions *Bool      `xmlrpc:"send_options,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MukRestClientGenerators represents array of muk_rest.client_generator model.
type MukRestClientGenerators []MukRestClientGenerator

// MukRestClientGeneratorModel is the odoo model name.
const MukRestClientGeneratorModel = "muk_rest.client_generator"

// Many2One convert MukRestClientGenerator to *Many2One.
func (mc *MukRestClientGenerator) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMukRestClientGenerator creates a new muk_rest.client_generator model and returns its id.
func (c *Client) CreateMukRestClientGenerator(mc *MukRestClientGenerator) (int64, error) {
	ids, err := c.CreateMukRestClientGenerators([]*MukRestClientGenerator{mc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestClientGenerator creates a new muk_rest.client_generator model and returns its id.
func (c *Client) CreateMukRestClientGenerators(mcs []*MukRestClientGenerator) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcs {
		vv = append(vv, v)
	}
	return c.Create(MukRestClientGeneratorModel, vv)
}

// UpdateMukRestClientGenerator updates an existing muk_rest.client_generator record.
func (c *Client) UpdateMukRestClientGenerator(mc *MukRestClientGenerator) error {
	return c.UpdateMukRestClientGenerators([]int64{mc.Id.Get()}, mc)
}

// UpdateMukRestClientGenerators updates existing muk_rest.client_generator records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMukRestClientGenerators(ids []int64, mc *MukRestClientGenerator) error {
	return c.Update(MukRestClientGeneratorModel, ids, mc)
}

// DeleteMukRestClientGenerator deletes an existing muk_rest.client_generator record.
func (c *Client) DeleteMukRestClientGenerator(id int64) error {
	return c.DeleteMukRestClientGenerators([]int64{id})
}

// DeleteMukRestClientGenerators deletes existing muk_rest.client_generator records.
func (c *Client) DeleteMukRestClientGenerators(ids []int64) error {
	return c.Delete(MukRestClientGeneratorModel, ids)
}

// GetMukRestClientGenerator gets muk_rest.client_generator existing record.
func (c *Client) GetMukRestClientGenerator(id int64) (*MukRestClientGenerator, error) {
	mcs, err := c.GetMukRestClientGenerators([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.client_generator not found", id)
}

// GetMukRestClientGenerators gets muk_rest.client_generator existing records.
func (c *Client) GetMukRestClientGenerators(ids []int64) (*MukRestClientGenerators, error) {
	mcs := &MukRestClientGenerators{}
	if err := c.Read(MukRestClientGeneratorModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMukRestClientGenerator finds muk_rest.client_generator record by querying it with criteria.
func (c *Client) FindMukRestClientGenerator(criteria *Criteria) (*MukRestClientGenerator, error) {
	mcs := &MukRestClientGenerators{}
	if err := c.SearchRead(MukRestClientGeneratorModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.client_generator was not found with criteria %v", criteria)
}

// FindMukRestClientGenerators finds muk_rest.client_generator records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestClientGenerators(criteria *Criteria, options *Options) (*MukRestClientGenerators, error) {
	mcs := &MukRestClientGenerators{}
	if err := c.SearchRead(MukRestClientGeneratorModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMukRestClientGeneratorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestClientGeneratorIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestClientGeneratorModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestClientGeneratorId finds record id by querying it with criteria.
func (c *Client) FindMukRestClientGeneratorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestClientGeneratorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.client_generator was not found with criteria %v and options %v", criteria, options)
}
