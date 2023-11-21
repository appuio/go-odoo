package odoo

import (
	"fmt"
)

// ChangePasswordOwn represents change.password.own model.
type ChangePasswordOwn struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	ConfirmPassword *String   `xmlrpc:"confirm_password,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	NewPassword     *String   `xmlrpc:"new_password,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ChangePasswordOwns represents array of change.password.own model.
type ChangePasswordOwns []ChangePasswordOwn

// ChangePasswordOwnModel is the odoo model name.
const ChangePasswordOwnModel = "change.password.own"

// Many2One convert ChangePasswordOwn to *Many2One.
func (cpo *ChangePasswordOwn) Many2One() *Many2One {
	return NewMany2One(cpo.Id.Get(), "")
}

// CreateChangePasswordOwn creates a new change.password.own model and returns its id.
func (c *Client) CreateChangePasswordOwn(cpo *ChangePasswordOwn) (int64, error) {
	ids, err := c.CreateChangePasswordOwns([]*ChangePasswordOwn{cpo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChangePasswordOwn creates a new change.password.own model and returns its id.
func (c *Client) CreateChangePasswordOwns(cpos []*ChangePasswordOwn) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpos {
		vv = append(vv, v)
	}
	return c.Create(ChangePasswordOwnModel, vv)
}

// UpdateChangePasswordOwn updates an existing change.password.own record.
func (c *Client) UpdateChangePasswordOwn(cpo *ChangePasswordOwn) error {
	return c.UpdateChangePasswordOwns([]int64{cpo.Id.Get()}, cpo)
}

// UpdateChangePasswordOwns updates existing change.password.own records.
// All records (represented by ids) will be updated by cpo values.
func (c *Client) UpdateChangePasswordOwns(ids []int64, cpo *ChangePasswordOwn) error {
	return c.Update(ChangePasswordOwnModel, ids, cpo)
}

// DeleteChangePasswordOwn deletes an existing change.password.own record.
func (c *Client) DeleteChangePasswordOwn(id int64) error {
	return c.DeleteChangePasswordOwns([]int64{id})
}

// DeleteChangePasswordOwns deletes existing change.password.own records.
func (c *Client) DeleteChangePasswordOwns(ids []int64) error {
	return c.Delete(ChangePasswordOwnModel, ids)
}

// GetChangePasswordOwn gets change.password.own existing record.
func (c *Client) GetChangePasswordOwn(id int64) (*ChangePasswordOwn, error) {
	cpos, err := c.GetChangePasswordOwns([]int64{id})
	if err != nil {
		return nil, err
	}
	if cpos != nil && len(*cpos) > 0 {
		return &((*cpos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of change.password.own not found", id)
}

// GetChangePasswordOwns gets change.password.own existing records.
func (c *Client) GetChangePasswordOwns(ids []int64) (*ChangePasswordOwns, error) {
	cpos := &ChangePasswordOwns{}
	if err := c.Read(ChangePasswordOwnModel, ids, nil, cpos); err != nil {
		return nil, err
	}
	return cpos, nil
}

// FindChangePasswordOwn finds change.password.own record by querying it with criteria.
func (c *Client) FindChangePasswordOwn(criteria *Criteria) (*ChangePasswordOwn, error) {
	cpos := &ChangePasswordOwns{}
	if err := c.SearchRead(ChangePasswordOwnModel, criteria, NewOptions().Limit(1), cpos); err != nil {
		return nil, err
	}
	if cpos != nil && len(*cpos) > 0 {
		return &((*cpos)[0]), nil
	}
	return nil, fmt.Errorf("change.password.own was not found with criteria %v", criteria)
}

// FindChangePasswordOwns finds change.password.own records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangePasswordOwns(criteria *Criteria, options *Options) (*ChangePasswordOwns, error) {
	cpos := &ChangePasswordOwns{}
	if err := c.SearchRead(ChangePasswordOwnModel, criteria, options, cpos); err != nil {
		return nil, err
	}
	return cpos, nil
}

// FindChangePasswordOwnIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangePasswordOwnIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ChangePasswordOwnModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindChangePasswordOwnId finds record id by querying it with criteria.
func (c *Client) FindChangePasswordOwnId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChangePasswordOwnModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("change.password.own was not found with criteria %v and options %v", criteria, options)
}
