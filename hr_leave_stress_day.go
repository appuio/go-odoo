package odoo

import (
	"fmt"
)

// HrLeaveStressDay represents hr.leave.stress.day model.
type HrLeaveStressDay struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	Color              *Int      `xmlrpc:"color,omptempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DepartmentIds      *Relation `xmlrpc:"department_ids,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	EndDate            *Time     `xmlrpc:"end_date,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	Name               *String   `xmlrpc:"name,omptempty"`
	ResourceCalendarId *Many2One `xmlrpc:"resource_calendar_id,omptempty"`
	StartDate          *Time     `xmlrpc:"start_date,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrLeaveStressDays represents array of hr.leave.stress.day model.
type HrLeaveStressDays []HrLeaveStressDay

// HrLeaveStressDayModel is the odoo model name.
const HrLeaveStressDayModel = "hr.leave.stress.day"

// Many2One convert HrLeaveStressDay to *Many2One.
func (hlsd *HrLeaveStressDay) Many2One() *Many2One {
	return NewMany2One(hlsd.Id.Get(), "")
}

// CreateHrLeaveStressDay creates a new hr.leave.stress.day model and returns its id.
func (c *Client) CreateHrLeaveStressDay(hlsd *HrLeaveStressDay) (int64, error) {
	ids, err := c.CreateHrLeaveStressDays([]*HrLeaveStressDay{hlsd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveStressDay creates a new hr.leave.stress.day model and returns its id.
func (c *Client) CreateHrLeaveStressDays(hlsds []*HrLeaveStressDay) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlsds {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveStressDayModel, vv)
}

// UpdateHrLeaveStressDay updates an existing hr.leave.stress.day record.
func (c *Client) UpdateHrLeaveStressDay(hlsd *HrLeaveStressDay) error {
	return c.UpdateHrLeaveStressDays([]int64{hlsd.Id.Get()}, hlsd)
}

// UpdateHrLeaveStressDays updates existing hr.leave.stress.day records.
// All records (represented by ids) will be updated by hlsd values.
func (c *Client) UpdateHrLeaveStressDays(ids []int64, hlsd *HrLeaveStressDay) error {
	return c.Update(HrLeaveStressDayModel, ids, hlsd)
}

// DeleteHrLeaveStressDay deletes an existing hr.leave.stress.day record.
func (c *Client) DeleteHrLeaveStressDay(id int64) error {
	return c.DeleteHrLeaveStressDays([]int64{id})
}

// DeleteHrLeaveStressDays deletes existing hr.leave.stress.day records.
func (c *Client) DeleteHrLeaveStressDays(ids []int64) error {
	return c.Delete(HrLeaveStressDayModel, ids)
}

// GetHrLeaveStressDay gets hr.leave.stress.day existing record.
func (c *Client) GetHrLeaveStressDay(id int64) (*HrLeaveStressDay, error) {
	hlsds, err := c.GetHrLeaveStressDays([]int64{id})
	if err != nil {
		return nil, err
	}
	if hlsds != nil && len(*hlsds) > 0 {
		return &((*hlsds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.stress.day not found", id)
}

// GetHrLeaveStressDays gets hr.leave.stress.day existing records.
func (c *Client) GetHrLeaveStressDays(ids []int64) (*HrLeaveStressDays, error) {
	hlsds := &HrLeaveStressDays{}
	if err := c.Read(HrLeaveStressDayModel, ids, nil, hlsds); err != nil {
		return nil, err
	}
	return hlsds, nil
}

// FindHrLeaveStressDay finds hr.leave.stress.day record by querying it with criteria.
func (c *Client) FindHrLeaveStressDay(criteria *Criteria) (*HrLeaveStressDay, error) {
	hlsds := &HrLeaveStressDays{}
	if err := c.SearchRead(HrLeaveStressDayModel, criteria, NewOptions().Limit(1), hlsds); err != nil {
		return nil, err
	}
	if hlsds != nil && len(*hlsds) > 0 {
		return &((*hlsds)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.stress.day was not found with criteria %v", criteria)
}

// FindHrLeaveStressDays finds hr.leave.stress.day records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveStressDays(criteria *Criteria, options *Options) (*HrLeaveStressDays, error) {
	hlsds := &HrLeaveStressDays{}
	if err := c.SearchRead(HrLeaveStressDayModel, criteria, options, hlsds); err != nil {
		return nil, err
	}
	return hlsds, nil
}

// FindHrLeaveStressDayIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveStressDayIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveStressDayModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveStressDayId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveStressDayId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveStressDayModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.stress.day was not found with criteria %v and options %v", criteria, options)
}
