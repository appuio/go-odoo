package odoo

import (
	"fmt"
)

// HrWorkEntryType represents hr.work.entry.type model.
type HrWorkEntryType struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	Active             *Bool      `xmlrpc:"active,omptempty"`
	Code               *String    `xmlrpc:"code,omptempty"`
	Color              *Int       `xmlrpc:"color,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	IsLeave            *Bool      `xmlrpc:"is_leave,omptempty"`
	IsUnforeseen       *Bool      `xmlrpc:"is_unforeseen,omptempty"`
	LeaveTypeIds       *Relation  `xmlrpc:"leave_type_ids,omptempty"`
	Name               *String    `xmlrpc:"name,omptempty"`
	RoundDays          *Selection `xmlrpc:"round_days,omptempty"`
	RoundDaysType      *Selection `xmlrpc:"round_days_type,omptempty"`
	Sequence           *Int       `xmlrpc:"sequence,omptempty"`
	UnpaidStructureIds *Relation  `xmlrpc:"unpaid_structure_ids,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrWorkEntryTypes represents array of hr.work.entry.type model.
type HrWorkEntryTypes []HrWorkEntryType

// HrWorkEntryTypeModel is the odoo model name.
const HrWorkEntryTypeModel = "hr.work.entry.type"

// Many2One convert HrWorkEntryType to *Many2One.
func (hwet *HrWorkEntryType) Many2One() *Many2One {
	return NewMany2One(hwet.Id.Get(), "")
}

// CreateHrWorkEntryType creates a new hr.work.entry.type model and returns its id.
func (c *Client) CreateHrWorkEntryType(hwet *HrWorkEntryType) (int64, error) {
	ids, err := c.CreateHrWorkEntryTypes([]*HrWorkEntryType{hwet})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrWorkEntryType creates a new hr.work.entry.type model and returns its id.
func (c *Client) CreateHrWorkEntryTypes(hwets []*HrWorkEntryType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hwets {
		vv = append(vv, v)
	}
	return c.Create(HrWorkEntryTypeModel, vv)
}

// UpdateHrWorkEntryType updates an existing hr.work.entry.type record.
func (c *Client) UpdateHrWorkEntryType(hwet *HrWorkEntryType) error {
	return c.UpdateHrWorkEntryTypes([]int64{hwet.Id.Get()}, hwet)
}

// UpdateHrWorkEntryTypes updates existing hr.work.entry.type records.
// All records (represented by ids) will be updated by hwet values.
func (c *Client) UpdateHrWorkEntryTypes(ids []int64, hwet *HrWorkEntryType) error {
	return c.Update(HrWorkEntryTypeModel, ids, hwet)
}

// DeleteHrWorkEntryType deletes an existing hr.work.entry.type record.
func (c *Client) DeleteHrWorkEntryType(id int64) error {
	return c.DeleteHrWorkEntryTypes([]int64{id})
}

// DeleteHrWorkEntryTypes deletes existing hr.work.entry.type records.
func (c *Client) DeleteHrWorkEntryTypes(ids []int64) error {
	return c.Delete(HrWorkEntryTypeModel, ids)
}

// GetHrWorkEntryType gets hr.work.entry.type existing record.
func (c *Client) GetHrWorkEntryType(id int64) (*HrWorkEntryType, error) {
	hwets, err := c.GetHrWorkEntryTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if hwets != nil && len(*hwets) > 0 {
		return &((*hwets)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.work.entry.type not found", id)
}

// GetHrWorkEntryTypes gets hr.work.entry.type existing records.
func (c *Client) GetHrWorkEntryTypes(ids []int64) (*HrWorkEntryTypes, error) {
	hwets := &HrWorkEntryTypes{}
	if err := c.Read(HrWorkEntryTypeModel, ids, nil, hwets); err != nil {
		return nil, err
	}
	return hwets, nil
}

// FindHrWorkEntryType finds hr.work.entry.type record by querying it with criteria.
func (c *Client) FindHrWorkEntryType(criteria *Criteria) (*HrWorkEntryType, error) {
	hwets := &HrWorkEntryTypes{}
	if err := c.SearchRead(HrWorkEntryTypeModel, criteria, NewOptions().Limit(1), hwets); err != nil {
		return nil, err
	}
	if hwets != nil && len(*hwets) > 0 {
		return &((*hwets)[0]), nil
	}
	return nil, fmt.Errorf("hr.work.entry.type was not found with criteria %v", criteria)
}

// FindHrWorkEntryTypes finds hr.work.entry.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntryTypes(criteria *Criteria, options *Options) (*HrWorkEntryTypes, error) {
	hwets := &HrWorkEntryTypes{}
	if err := c.SearchRead(HrWorkEntryTypeModel, criteria, options, hwets); err != nil {
		return nil, err
	}
	return hwets, nil
}

// FindHrWorkEntryTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntryTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrWorkEntryTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrWorkEntryTypeId finds record id by querying it with criteria.
func (c *Client) FindHrWorkEntryTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrWorkEntryTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.work.entry.type was not found with criteria %v and options %v", criteria, options)
}
