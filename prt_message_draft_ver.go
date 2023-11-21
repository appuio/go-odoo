package odoo

import (
	"fmt"
)

// PrtMessageDraftVer represents prt.message.draft.ver model.
type PrtMessageDraftVer struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	Body           *String   `xmlrpc:"body,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrentDraftId *Many2One `xmlrpc:"current_draft_id,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	DraftId        *Many2One `xmlrpc:"draft_id,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	Model          *String   `xmlrpc:"model,omptempty"`
	ResId          *Int      `xmlrpc:"res_id,omptempty"`
	Subject        *String   `xmlrpc:"subject,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PrtMessageDraftVers represents array of prt.message.draft.ver model.
type PrtMessageDraftVers []PrtMessageDraftVer

// PrtMessageDraftVerModel is the odoo model name.
const PrtMessageDraftVerModel = "prt.message.draft.ver"

// Many2One convert PrtMessageDraftVer to *Many2One.
func (pmdv *PrtMessageDraftVer) Many2One() *Many2One {
	return NewMany2One(pmdv.Id.Get(), "")
}

// CreatePrtMessageDraftVer creates a new prt.message.draft.ver model and returns its id.
func (c *Client) CreatePrtMessageDraftVer(pmdv *PrtMessageDraftVer) (int64, error) {
	ids, err := c.CreatePrtMessageDraftVers([]*PrtMessageDraftVer{pmdv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePrtMessageDraftVer creates a new prt.message.draft.ver model and returns its id.
func (c *Client) CreatePrtMessageDraftVers(pmdvs []*PrtMessageDraftVer) ([]int64, error) {
	var vv []interface{}
	for _, v := range pmdvs {
		vv = append(vv, v)
	}
	return c.Create(PrtMessageDraftVerModel, vv)
}

// UpdatePrtMessageDraftVer updates an existing prt.message.draft.ver record.
func (c *Client) UpdatePrtMessageDraftVer(pmdv *PrtMessageDraftVer) error {
	return c.UpdatePrtMessageDraftVers([]int64{pmdv.Id.Get()}, pmdv)
}

// UpdatePrtMessageDraftVers updates existing prt.message.draft.ver records.
// All records (represented by ids) will be updated by pmdv values.
func (c *Client) UpdatePrtMessageDraftVers(ids []int64, pmdv *PrtMessageDraftVer) error {
	return c.Update(PrtMessageDraftVerModel, ids, pmdv)
}

// DeletePrtMessageDraftVer deletes an existing prt.message.draft.ver record.
func (c *Client) DeletePrtMessageDraftVer(id int64) error {
	return c.DeletePrtMessageDraftVers([]int64{id})
}

// DeletePrtMessageDraftVers deletes existing prt.message.draft.ver records.
func (c *Client) DeletePrtMessageDraftVers(ids []int64) error {
	return c.Delete(PrtMessageDraftVerModel, ids)
}

// GetPrtMessageDraftVer gets prt.message.draft.ver existing record.
func (c *Client) GetPrtMessageDraftVer(id int64) (*PrtMessageDraftVer, error) {
	pmdvs, err := c.GetPrtMessageDraftVers([]int64{id})
	if err != nil {
		return nil, err
	}
	if pmdvs != nil && len(*pmdvs) > 0 {
		return &((*pmdvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of prt.message.draft.ver not found", id)
}

// GetPrtMessageDraftVers gets prt.message.draft.ver existing records.
func (c *Client) GetPrtMessageDraftVers(ids []int64) (*PrtMessageDraftVers, error) {
	pmdvs := &PrtMessageDraftVers{}
	if err := c.Read(PrtMessageDraftVerModel, ids, nil, pmdvs); err != nil {
		return nil, err
	}
	return pmdvs, nil
}

// FindPrtMessageDraftVer finds prt.message.draft.ver record by querying it with criteria.
func (c *Client) FindPrtMessageDraftVer(criteria *Criteria) (*PrtMessageDraftVer, error) {
	pmdvs := &PrtMessageDraftVers{}
	if err := c.SearchRead(PrtMessageDraftVerModel, criteria, NewOptions().Limit(1), pmdvs); err != nil {
		return nil, err
	}
	if pmdvs != nil && len(*pmdvs) > 0 {
		return &((*pmdvs)[0]), nil
	}
	return nil, fmt.Errorf("prt.message.draft.ver was not found with criteria %v", criteria)
}

// FindPrtMessageDraftVers finds prt.message.draft.ver records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMessageDraftVers(criteria *Criteria, options *Options) (*PrtMessageDraftVers, error) {
	pmdvs := &PrtMessageDraftVers{}
	if err := c.SearchRead(PrtMessageDraftVerModel, criteria, options, pmdvs); err != nil {
		return nil, err
	}
	return pmdvs, nil
}

// FindPrtMessageDraftVerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMessageDraftVerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PrtMessageDraftVerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPrtMessageDraftVerId finds record id by querying it with criteria.
func (c *Client) FindPrtMessageDraftVerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrtMessageDraftVerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("prt.message.draft.ver was not found with criteria %v and options %v", criteria, options)
}
