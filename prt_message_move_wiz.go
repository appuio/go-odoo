package odoo

import (
	"fmt"
)

// PrtMessageMoveWiz represents prt.message.move.wiz model.
type PrtMessageMoveWiz struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	IsConversation *Bool      `xmlrpc:"is_conversation,omptempty"`
	IsLead         *Bool      `xmlrpc:"is_lead,omptempty"`
	LeadDelete     *Bool      `xmlrpc:"lead_delete,omptempty"`
	ModelTo        *String    `xmlrpc:"model_to,omptempty"`
	Notify         *Selection `xmlrpc:"notify,omptempty"`
	OppDelete      *Bool      `xmlrpc:"opp_delete,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PrtMessageMoveWizs represents array of prt.message.move.wiz model.
type PrtMessageMoveWizs []PrtMessageMoveWiz

// PrtMessageMoveWizModel is the odoo model name.
const PrtMessageMoveWizModel = "prt.message.move.wiz"

// Many2One convert PrtMessageMoveWiz to *Many2One.
func (pmmw *PrtMessageMoveWiz) Many2One() *Many2One {
	return NewMany2One(pmmw.Id.Get(), "")
}

// CreatePrtMessageMoveWiz creates a new prt.message.move.wiz model and returns its id.
func (c *Client) CreatePrtMessageMoveWiz(pmmw *PrtMessageMoveWiz) (int64, error) {
	ids, err := c.CreatePrtMessageMoveWizs([]*PrtMessageMoveWiz{pmmw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePrtMessageMoveWiz creates a new prt.message.move.wiz model and returns its id.
func (c *Client) CreatePrtMessageMoveWizs(pmmws []*PrtMessageMoveWiz) ([]int64, error) {
	var vv []interface{}
	for _, v := range pmmws {
		vv = append(vv, v)
	}
	return c.Create(PrtMessageMoveWizModel, vv)
}

// UpdatePrtMessageMoveWiz updates an existing prt.message.move.wiz record.
func (c *Client) UpdatePrtMessageMoveWiz(pmmw *PrtMessageMoveWiz) error {
	return c.UpdatePrtMessageMoveWizs([]int64{pmmw.Id.Get()}, pmmw)
}

// UpdatePrtMessageMoveWizs updates existing prt.message.move.wiz records.
// All records (represented by ids) will be updated by pmmw values.
func (c *Client) UpdatePrtMessageMoveWizs(ids []int64, pmmw *PrtMessageMoveWiz) error {
	return c.Update(PrtMessageMoveWizModel, ids, pmmw)
}

// DeletePrtMessageMoveWiz deletes an existing prt.message.move.wiz record.
func (c *Client) DeletePrtMessageMoveWiz(id int64) error {
	return c.DeletePrtMessageMoveWizs([]int64{id})
}

// DeletePrtMessageMoveWizs deletes existing prt.message.move.wiz records.
func (c *Client) DeletePrtMessageMoveWizs(ids []int64) error {
	return c.Delete(PrtMessageMoveWizModel, ids)
}

// GetPrtMessageMoveWiz gets prt.message.move.wiz existing record.
func (c *Client) GetPrtMessageMoveWiz(id int64) (*PrtMessageMoveWiz, error) {
	pmmws, err := c.GetPrtMessageMoveWizs([]int64{id})
	if err != nil {
		return nil, err
	}
	if pmmws != nil && len(*pmmws) > 0 {
		return &((*pmmws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of prt.message.move.wiz not found", id)
}

// GetPrtMessageMoveWizs gets prt.message.move.wiz existing records.
func (c *Client) GetPrtMessageMoveWizs(ids []int64) (*PrtMessageMoveWizs, error) {
	pmmws := &PrtMessageMoveWizs{}
	if err := c.Read(PrtMessageMoveWizModel, ids, nil, pmmws); err != nil {
		return nil, err
	}
	return pmmws, nil
}

// FindPrtMessageMoveWiz finds prt.message.move.wiz record by querying it with criteria.
func (c *Client) FindPrtMessageMoveWiz(criteria *Criteria) (*PrtMessageMoveWiz, error) {
	pmmws := &PrtMessageMoveWizs{}
	if err := c.SearchRead(PrtMessageMoveWizModel, criteria, NewOptions().Limit(1), pmmws); err != nil {
		return nil, err
	}
	if pmmws != nil && len(*pmmws) > 0 {
		return &((*pmmws)[0]), nil
	}
	return nil, fmt.Errorf("prt.message.move.wiz was not found with criteria %v", criteria)
}

// FindPrtMessageMoveWizs finds prt.message.move.wiz records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMessageMoveWizs(criteria *Criteria, options *Options) (*PrtMessageMoveWizs, error) {
	pmmws := &PrtMessageMoveWizs{}
	if err := c.SearchRead(PrtMessageMoveWizModel, criteria, options, pmmws); err != nil {
		return nil, err
	}
	return pmmws, nil
}

// FindPrtMessageMoveWizIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMessageMoveWizIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PrtMessageMoveWizModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPrtMessageMoveWizId finds record id by querying it with criteria.
func (c *Client) FindPrtMessageMoveWizId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrtMessageMoveWizModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("prt.message.move.wiz was not found with criteria %v and options %v", criteria, options)
}
