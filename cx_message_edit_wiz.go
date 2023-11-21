package odoo

import (
	"fmt"
)

// CxMessageEditWiz represents cx.message.edit.wiz model.
type CxMessageEditWiz struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Body        *String   `xmlrpc:"body,omptempty"`
	CanEdit     *Bool     `xmlrpc:"can_edit,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MessageId   *Many2One `xmlrpc:"message_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CxMessageEditWizs represents array of cx.message.edit.wiz model.
type CxMessageEditWizs []CxMessageEditWiz

// CxMessageEditWizModel is the odoo model name.
const CxMessageEditWizModel = "cx.message.edit.wiz"

// Many2One convert CxMessageEditWiz to *Many2One.
func (cmew *CxMessageEditWiz) Many2One() *Many2One {
	return NewMany2One(cmew.Id.Get(), "")
}

// CreateCxMessageEditWiz creates a new cx.message.edit.wiz model and returns its id.
func (c *Client) CreateCxMessageEditWiz(cmew *CxMessageEditWiz) (int64, error) {
	ids, err := c.CreateCxMessageEditWizs([]*CxMessageEditWiz{cmew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxMessageEditWiz creates a new cx.message.edit.wiz model and returns its id.
func (c *Client) CreateCxMessageEditWizs(cmews []*CxMessageEditWiz) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmews {
		vv = append(vv, v)
	}
	return c.Create(CxMessageEditWizModel, vv)
}

// UpdateCxMessageEditWiz updates an existing cx.message.edit.wiz record.
func (c *Client) UpdateCxMessageEditWiz(cmew *CxMessageEditWiz) error {
	return c.UpdateCxMessageEditWizs([]int64{cmew.Id.Get()}, cmew)
}

// UpdateCxMessageEditWizs updates existing cx.message.edit.wiz records.
// All records (represented by ids) will be updated by cmew values.
func (c *Client) UpdateCxMessageEditWizs(ids []int64, cmew *CxMessageEditWiz) error {
	return c.Update(CxMessageEditWizModel, ids, cmew)
}

// DeleteCxMessageEditWiz deletes an existing cx.message.edit.wiz record.
func (c *Client) DeleteCxMessageEditWiz(id int64) error {
	return c.DeleteCxMessageEditWizs([]int64{id})
}

// DeleteCxMessageEditWizs deletes existing cx.message.edit.wiz records.
func (c *Client) DeleteCxMessageEditWizs(ids []int64) error {
	return c.Delete(CxMessageEditWizModel, ids)
}

// GetCxMessageEditWiz gets cx.message.edit.wiz existing record.
func (c *Client) GetCxMessageEditWiz(id int64) (*CxMessageEditWiz, error) {
	cmews, err := c.GetCxMessageEditWizs([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmews != nil && len(*cmews) > 0 {
		return &((*cmews)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.message.edit.wiz not found", id)
}

// GetCxMessageEditWizs gets cx.message.edit.wiz existing records.
func (c *Client) GetCxMessageEditWizs(ids []int64) (*CxMessageEditWizs, error) {
	cmews := &CxMessageEditWizs{}
	if err := c.Read(CxMessageEditWizModel, ids, nil, cmews); err != nil {
		return nil, err
	}
	return cmews, nil
}

// FindCxMessageEditWiz finds cx.message.edit.wiz record by querying it with criteria.
func (c *Client) FindCxMessageEditWiz(criteria *Criteria) (*CxMessageEditWiz, error) {
	cmews := &CxMessageEditWizs{}
	if err := c.SearchRead(CxMessageEditWizModel, criteria, NewOptions().Limit(1), cmews); err != nil {
		return nil, err
	}
	if cmews != nil && len(*cmews) > 0 {
		return &((*cmews)[0]), nil
	}
	return nil, fmt.Errorf("cx.message.edit.wiz was not found with criteria %v", criteria)
}

// FindCxMessageEditWizs finds cx.message.edit.wiz records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageEditWizs(criteria *Criteria, options *Options) (*CxMessageEditWizs, error) {
	cmews := &CxMessageEditWizs{}
	if err := c.SearchRead(CxMessageEditWizModel, criteria, options, cmews); err != nil {
		return nil, err
	}
	return cmews, nil
}

// FindCxMessageEditWizIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageEditWizIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxMessageEditWizModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxMessageEditWizId finds record id by querying it with criteria.
func (c *Client) FindCxMessageEditWizId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxMessageEditWizModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.message.edit.wiz was not found with criteria %v and options %v", criteria, options)
}
