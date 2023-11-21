package odoo

import (
	"fmt"
)

// CxMessagePartnerAssignWiz represents cx.message.partner.assign.wiz model.
type CxMessagePartnerAssignWiz struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Email       *String   `xmlrpc:"email,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omptempty"`
	SameEmail   *Bool     `xmlrpc:"same_email,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CxMessagePartnerAssignWizs represents array of cx.message.partner.assign.wiz model.
type CxMessagePartnerAssignWizs []CxMessagePartnerAssignWiz

// CxMessagePartnerAssignWizModel is the odoo model name.
const CxMessagePartnerAssignWizModel = "cx.message.partner.assign.wiz"

// Many2One convert CxMessagePartnerAssignWiz to *Many2One.
func (cmpaw *CxMessagePartnerAssignWiz) Many2One() *Many2One {
	return NewMany2One(cmpaw.Id.Get(), "")
}

// CreateCxMessagePartnerAssignWiz creates a new cx.message.partner.assign.wiz model and returns its id.
func (c *Client) CreateCxMessagePartnerAssignWiz(cmpaw *CxMessagePartnerAssignWiz) (int64, error) {
	ids, err := c.CreateCxMessagePartnerAssignWizs([]*CxMessagePartnerAssignWiz{cmpaw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxMessagePartnerAssignWiz creates a new cx.message.partner.assign.wiz model and returns its id.
func (c *Client) CreateCxMessagePartnerAssignWizs(cmpaws []*CxMessagePartnerAssignWiz) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmpaws {
		vv = append(vv, v)
	}
	return c.Create(CxMessagePartnerAssignWizModel, vv)
}

// UpdateCxMessagePartnerAssignWiz updates an existing cx.message.partner.assign.wiz record.
func (c *Client) UpdateCxMessagePartnerAssignWiz(cmpaw *CxMessagePartnerAssignWiz) error {
	return c.UpdateCxMessagePartnerAssignWizs([]int64{cmpaw.Id.Get()}, cmpaw)
}

// UpdateCxMessagePartnerAssignWizs updates existing cx.message.partner.assign.wiz records.
// All records (represented by ids) will be updated by cmpaw values.
func (c *Client) UpdateCxMessagePartnerAssignWizs(ids []int64, cmpaw *CxMessagePartnerAssignWiz) error {
	return c.Update(CxMessagePartnerAssignWizModel, ids, cmpaw)
}

// DeleteCxMessagePartnerAssignWiz deletes an existing cx.message.partner.assign.wiz record.
func (c *Client) DeleteCxMessagePartnerAssignWiz(id int64) error {
	return c.DeleteCxMessagePartnerAssignWizs([]int64{id})
}

// DeleteCxMessagePartnerAssignWizs deletes existing cx.message.partner.assign.wiz records.
func (c *Client) DeleteCxMessagePartnerAssignWizs(ids []int64) error {
	return c.Delete(CxMessagePartnerAssignWizModel, ids)
}

// GetCxMessagePartnerAssignWiz gets cx.message.partner.assign.wiz existing record.
func (c *Client) GetCxMessagePartnerAssignWiz(id int64) (*CxMessagePartnerAssignWiz, error) {
	cmpaws, err := c.GetCxMessagePartnerAssignWizs([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmpaws != nil && len(*cmpaws) > 0 {
		return &((*cmpaws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.message.partner.assign.wiz not found", id)
}

// GetCxMessagePartnerAssignWizs gets cx.message.partner.assign.wiz existing records.
func (c *Client) GetCxMessagePartnerAssignWizs(ids []int64) (*CxMessagePartnerAssignWizs, error) {
	cmpaws := &CxMessagePartnerAssignWizs{}
	if err := c.Read(CxMessagePartnerAssignWizModel, ids, nil, cmpaws); err != nil {
		return nil, err
	}
	return cmpaws, nil
}

// FindCxMessagePartnerAssignWiz finds cx.message.partner.assign.wiz record by querying it with criteria.
func (c *Client) FindCxMessagePartnerAssignWiz(criteria *Criteria) (*CxMessagePartnerAssignWiz, error) {
	cmpaws := &CxMessagePartnerAssignWizs{}
	if err := c.SearchRead(CxMessagePartnerAssignWizModel, criteria, NewOptions().Limit(1), cmpaws); err != nil {
		return nil, err
	}
	if cmpaws != nil && len(*cmpaws) > 0 {
		return &((*cmpaws)[0]), nil
	}
	return nil, fmt.Errorf("cx.message.partner.assign.wiz was not found with criteria %v", criteria)
}

// FindCxMessagePartnerAssignWizs finds cx.message.partner.assign.wiz records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessagePartnerAssignWizs(criteria *Criteria, options *Options) (*CxMessagePartnerAssignWizs, error) {
	cmpaws := &CxMessagePartnerAssignWizs{}
	if err := c.SearchRead(CxMessagePartnerAssignWizModel, criteria, options, cmpaws); err != nil {
		return nil, err
	}
	return cmpaws, nil
}

// FindCxMessagePartnerAssignWizIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessagePartnerAssignWizIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxMessagePartnerAssignWizModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxMessagePartnerAssignWizId finds record id by querying it with criteria.
func (c *Client) FindCxMessagePartnerAssignWizId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxMessagePartnerAssignWizModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.message.partner.assign.wiz was not found with criteria %v and options %v", criteria, options)
}
