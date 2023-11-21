package odoo

import (
	"fmt"
)

// HrHolidaysCancelLeave represents hr.holidays.cancel.leave model.
type HrHolidaysCancelLeave struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LeaveId     *Many2One `xmlrpc:"leave_id,omptempty"`
	Reason      *String   `xmlrpc:"reason,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrHolidaysCancelLeaves represents array of hr.holidays.cancel.leave model.
type HrHolidaysCancelLeaves []HrHolidaysCancelLeave

// HrHolidaysCancelLeaveModel is the odoo model name.
const HrHolidaysCancelLeaveModel = "hr.holidays.cancel.leave"

// Many2One convert HrHolidaysCancelLeave to *Many2One.
func (hhcl *HrHolidaysCancelLeave) Many2One() *Many2One {
	return NewMany2One(hhcl.Id.Get(), "")
}

// CreateHrHolidaysCancelLeave creates a new hr.holidays.cancel.leave model and returns its id.
func (c *Client) CreateHrHolidaysCancelLeave(hhcl *HrHolidaysCancelLeave) (int64, error) {
	ids, err := c.CreateHrHolidaysCancelLeaves([]*HrHolidaysCancelLeave{hhcl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrHolidaysCancelLeave creates a new hr.holidays.cancel.leave model and returns its id.
func (c *Client) CreateHrHolidaysCancelLeaves(hhcls []*HrHolidaysCancelLeave) ([]int64, error) {
	var vv []interface{}
	for _, v := range hhcls {
		vv = append(vv, v)
	}
	return c.Create(HrHolidaysCancelLeaveModel, vv)
}

// UpdateHrHolidaysCancelLeave updates an existing hr.holidays.cancel.leave record.
func (c *Client) UpdateHrHolidaysCancelLeave(hhcl *HrHolidaysCancelLeave) error {
	return c.UpdateHrHolidaysCancelLeaves([]int64{hhcl.Id.Get()}, hhcl)
}

// UpdateHrHolidaysCancelLeaves updates existing hr.holidays.cancel.leave records.
// All records (represented by ids) will be updated by hhcl values.
func (c *Client) UpdateHrHolidaysCancelLeaves(ids []int64, hhcl *HrHolidaysCancelLeave) error {
	return c.Update(HrHolidaysCancelLeaveModel, ids, hhcl)
}

// DeleteHrHolidaysCancelLeave deletes an existing hr.holidays.cancel.leave record.
func (c *Client) DeleteHrHolidaysCancelLeave(id int64) error {
	return c.DeleteHrHolidaysCancelLeaves([]int64{id})
}

// DeleteHrHolidaysCancelLeaves deletes existing hr.holidays.cancel.leave records.
func (c *Client) DeleteHrHolidaysCancelLeaves(ids []int64) error {
	return c.Delete(HrHolidaysCancelLeaveModel, ids)
}

// GetHrHolidaysCancelLeave gets hr.holidays.cancel.leave existing record.
func (c *Client) GetHrHolidaysCancelLeave(id int64) (*HrHolidaysCancelLeave, error) {
	hhcls, err := c.GetHrHolidaysCancelLeaves([]int64{id})
	if err != nil {
		return nil, err
	}
	if hhcls != nil && len(*hhcls) > 0 {
		return &((*hhcls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.holidays.cancel.leave not found", id)
}

// GetHrHolidaysCancelLeaves gets hr.holidays.cancel.leave existing records.
func (c *Client) GetHrHolidaysCancelLeaves(ids []int64) (*HrHolidaysCancelLeaves, error) {
	hhcls := &HrHolidaysCancelLeaves{}
	if err := c.Read(HrHolidaysCancelLeaveModel, ids, nil, hhcls); err != nil {
		return nil, err
	}
	return hhcls, nil
}

// FindHrHolidaysCancelLeave finds hr.holidays.cancel.leave record by querying it with criteria.
func (c *Client) FindHrHolidaysCancelLeave(criteria *Criteria) (*HrHolidaysCancelLeave, error) {
	hhcls := &HrHolidaysCancelLeaves{}
	if err := c.SearchRead(HrHolidaysCancelLeaveModel, criteria, NewOptions().Limit(1), hhcls); err != nil {
		return nil, err
	}
	if hhcls != nil && len(*hhcls) > 0 {
		return &((*hhcls)[0]), nil
	}
	return nil, fmt.Errorf("hr.holidays.cancel.leave was not found with criteria %v", criteria)
}

// FindHrHolidaysCancelLeaves finds hr.holidays.cancel.leave records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysCancelLeaves(criteria *Criteria, options *Options) (*HrHolidaysCancelLeaves, error) {
	hhcls := &HrHolidaysCancelLeaves{}
	if err := c.SearchRead(HrHolidaysCancelLeaveModel, criteria, options, hhcls); err != nil {
		return nil, err
	}
	return hhcls, nil
}

// FindHrHolidaysCancelLeaveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrHolidaysCancelLeaveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrHolidaysCancelLeaveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrHolidaysCancelLeaveId finds record id by querying it with criteria.
func (c *Client) FindHrHolidaysCancelLeaveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrHolidaysCancelLeaveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.holidays.cancel.leave was not found with criteria %v and options %v", criteria, options)
}
