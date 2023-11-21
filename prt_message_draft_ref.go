package odoo

import (
	"fmt"
)

// PrtMessageDraftRef represents prt.message.draft.ref model.
type PrtMessageDraftRef struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	DraftId     *Int      `xmlrpc:"draft_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	RecordRef   *String   `xmlrpc:"record_ref,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PrtMessageDraftRefs represents array of prt.message.draft.ref model.
type PrtMessageDraftRefs []PrtMessageDraftRef

// PrtMessageDraftRefModel is the odoo model name.
const PrtMessageDraftRefModel = "prt.message.draft.ref"

// Many2One convert PrtMessageDraftRef to *Many2One.
func (pmdr *PrtMessageDraftRef) Many2One() *Many2One {
	return NewMany2One(pmdr.Id.Get(), "")
}

// CreatePrtMessageDraftRef creates a new prt.message.draft.ref model and returns its id.
func (c *Client) CreatePrtMessageDraftRef(pmdr *PrtMessageDraftRef) (int64, error) {
	ids, err := c.CreatePrtMessageDraftRefs([]*PrtMessageDraftRef{pmdr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePrtMessageDraftRef creates a new prt.message.draft.ref model and returns its id.
func (c *Client) CreatePrtMessageDraftRefs(pmdrs []*PrtMessageDraftRef) ([]int64, error) {
	var vv []interface{}
	for _, v := range pmdrs {
		vv = append(vv, v)
	}
	return c.Create(PrtMessageDraftRefModel, vv)
}

// UpdatePrtMessageDraftRef updates an existing prt.message.draft.ref record.
func (c *Client) UpdatePrtMessageDraftRef(pmdr *PrtMessageDraftRef) error {
	return c.UpdatePrtMessageDraftRefs([]int64{pmdr.Id.Get()}, pmdr)
}

// UpdatePrtMessageDraftRefs updates existing prt.message.draft.ref records.
// All records (represented by ids) will be updated by pmdr values.
func (c *Client) UpdatePrtMessageDraftRefs(ids []int64, pmdr *PrtMessageDraftRef) error {
	return c.Update(PrtMessageDraftRefModel, ids, pmdr)
}

// DeletePrtMessageDraftRef deletes an existing prt.message.draft.ref record.
func (c *Client) DeletePrtMessageDraftRef(id int64) error {
	return c.DeletePrtMessageDraftRefs([]int64{id})
}

// DeletePrtMessageDraftRefs deletes existing prt.message.draft.ref records.
func (c *Client) DeletePrtMessageDraftRefs(ids []int64) error {
	return c.Delete(PrtMessageDraftRefModel, ids)
}

// GetPrtMessageDraftRef gets prt.message.draft.ref existing record.
func (c *Client) GetPrtMessageDraftRef(id int64) (*PrtMessageDraftRef, error) {
	pmdrs, err := c.GetPrtMessageDraftRefs([]int64{id})
	if err != nil {
		return nil, err
	}
	if pmdrs != nil && len(*pmdrs) > 0 {
		return &((*pmdrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of prt.message.draft.ref not found", id)
}

// GetPrtMessageDraftRefs gets prt.message.draft.ref existing records.
func (c *Client) GetPrtMessageDraftRefs(ids []int64) (*PrtMessageDraftRefs, error) {
	pmdrs := &PrtMessageDraftRefs{}
	if err := c.Read(PrtMessageDraftRefModel, ids, nil, pmdrs); err != nil {
		return nil, err
	}
	return pmdrs, nil
}

// FindPrtMessageDraftRef finds prt.message.draft.ref record by querying it with criteria.
func (c *Client) FindPrtMessageDraftRef(criteria *Criteria) (*PrtMessageDraftRef, error) {
	pmdrs := &PrtMessageDraftRefs{}
	if err := c.SearchRead(PrtMessageDraftRefModel, criteria, NewOptions().Limit(1), pmdrs); err != nil {
		return nil, err
	}
	if pmdrs != nil && len(*pmdrs) > 0 {
		return &((*pmdrs)[0]), nil
	}
	return nil, fmt.Errorf("prt.message.draft.ref was not found with criteria %v", criteria)
}

// FindPrtMessageDraftRefs finds prt.message.draft.ref records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMessageDraftRefs(criteria *Criteria, options *Options) (*PrtMessageDraftRefs, error) {
	pmdrs := &PrtMessageDraftRefs{}
	if err := c.SearchRead(PrtMessageDraftRefModel, criteria, options, pmdrs); err != nil {
		return nil, err
	}
	return pmdrs, nil
}

// FindPrtMessageDraftRefIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMessageDraftRefIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PrtMessageDraftRefModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPrtMessageDraftRefId finds record id by querying it with criteria.
func (c *Client) FindPrtMessageDraftRefId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrtMessageDraftRefModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("prt.message.draft.ref was not found with criteria %v and options %v", criteria, options)
}
