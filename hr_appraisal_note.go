package odoo

import (
	"fmt"
)

// HrAppraisalNote represents hr.appraisal.note model.
type HrAppraisalNote struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrAppraisalNotes represents array of hr.appraisal.note model.
type HrAppraisalNotes []HrAppraisalNote

// HrAppraisalNoteModel is the odoo model name.
const HrAppraisalNoteModel = "hr.appraisal.note"

// Many2One convert HrAppraisalNote to *Many2One.
func (han *HrAppraisalNote) Many2One() *Many2One {
	return NewMany2One(han.Id.Get(), "")
}

// CreateHrAppraisalNote creates a new hr.appraisal.note model and returns its id.
func (c *Client) CreateHrAppraisalNote(han *HrAppraisalNote) (int64, error) {
	ids, err := c.CreateHrAppraisalNotes([]*HrAppraisalNote{han})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAppraisalNote creates a new hr.appraisal.note model and returns its id.
func (c *Client) CreateHrAppraisalNotes(hans []*HrAppraisalNote) ([]int64, error) {
	var vv []interface{}
	for _, v := range hans {
		vv = append(vv, v)
	}
	return c.Create(HrAppraisalNoteModel, vv)
}

// UpdateHrAppraisalNote updates an existing hr.appraisal.note record.
func (c *Client) UpdateHrAppraisalNote(han *HrAppraisalNote) error {
	return c.UpdateHrAppraisalNotes([]int64{han.Id.Get()}, han)
}

// UpdateHrAppraisalNotes updates existing hr.appraisal.note records.
// All records (represented by ids) will be updated by han values.
func (c *Client) UpdateHrAppraisalNotes(ids []int64, han *HrAppraisalNote) error {
	return c.Update(HrAppraisalNoteModel, ids, han)
}

// DeleteHrAppraisalNote deletes an existing hr.appraisal.note record.
func (c *Client) DeleteHrAppraisalNote(id int64) error {
	return c.DeleteHrAppraisalNotes([]int64{id})
}

// DeleteHrAppraisalNotes deletes existing hr.appraisal.note records.
func (c *Client) DeleteHrAppraisalNotes(ids []int64) error {
	return c.Delete(HrAppraisalNoteModel, ids)
}

// GetHrAppraisalNote gets hr.appraisal.note existing record.
func (c *Client) GetHrAppraisalNote(id int64) (*HrAppraisalNote, error) {
	hans, err := c.GetHrAppraisalNotes([]int64{id})
	if err != nil {
		return nil, err
	}
	if hans != nil && len(*hans) > 0 {
		return &((*hans)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.appraisal.note not found", id)
}

// GetHrAppraisalNotes gets hr.appraisal.note existing records.
func (c *Client) GetHrAppraisalNotes(ids []int64) (*HrAppraisalNotes, error) {
	hans := &HrAppraisalNotes{}
	if err := c.Read(HrAppraisalNoteModel, ids, nil, hans); err != nil {
		return nil, err
	}
	return hans, nil
}

// FindHrAppraisalNote finds hr.appraisal.note record by querying it with criteria.
func (c *Client) FindHrAppraisalNote(criteria *Criteria) (*HrAppraisalNote, error) {
	hans := &HrAppraisalNotes{}
	if err := c.SearchRead(HrAppraisalNoteModel, criteria, NewOptions().Limit(1), hans); err != nil {
		return nil, err
	}
	if hans != nil && len(*hans) > 0 {
		return &((*hans)[0]), nil
	}
	return nil, fmt.Errorf("hr.appraisal.note was not found with criteria %v", criteria)
}

// FindHrAppraisalNotes finds hr.appraisal.note records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalNotes(criteria *Criteria, options *Options) (*HrAppraisalNotes, error) {
	hans := &HrAppraisalNotes{}
	if err := c.SearchRead(HrAppraisalNoteModel, criteria, options, hans); err != nil {
		return nil, err
	}
	return hans, nil
}

// FindHrAppraisalNoteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalNoteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAppraisalNoteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAppraisalNoteId finds record id by querying it with criteria.
func (c *Client) FindHrAppraisalNoteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAppraisalNoteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.appraisal.note was not found with criteria %v and options %v", criteria, options)
}
