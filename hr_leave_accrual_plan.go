package odoo

import (
	"fmt"
)

// HrLeaveAccrualPlan represents hr.leave.accrual.plan model.
type HrLeaveAccrualPlan struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	AllocationIds      *Relation  `xmlrpc:"allocation_ids,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	EmployeesCount     *Int       `xmlrpc:"employees_count,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	LevelCount         *Int       `xmlrpc:"level_count,omptempty"`
	LevelIds           *Relation  `xmlrpc:"level_ids,omptempty"`
	Name               *String    `xmlrpc:"name,omptempty"`
	ShowTransitionMode *Bool      `xmlrpc:"show_transition_mode,omptempty"`
	TimeOffTypeId      *Many2One  `xmlrpc:"time_off_type_id,omptempty"`
	TransitionMode     *Selection `xmlrpc:"transition_mode,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrLeaveAccrualPlans represents array of hr.leave.accrual.plan model.
type HrLeaveAccrualPlans []HrLeaveAccrualPlan

// HrLeaveAccrualPlanModel is the odoo model name.
const HrLeaveAccrualPlanModel = "hr.leave.accrual.plan"

// Many2One convert HrLeaveAccrualPlan to *Many2One.
func (hlap *HrLeaveAccrualPlan) Many2One() *Many2One {
	return NewMany2One(hlap.Id.Get(), "")
}

// CreateHrLeaveAccrualPlan creates a new hr.leave.accrual.plan model and returns its id.
func (c *Client) CreateHrLeaveAccrualPlan(hlap *HrLeaveAccrualPlan) (int64, error) {
	ids, err := c.CreateHrLeaveAccrualPlans([]*HrLeaveAccrualPlan{hlap})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveAccrualPlan creates a new hr.leave.accrual.plan model and returns its id.
func (c *Client) CreateHrLeaveAccrualPlans(hlaps []*HrLeaveAccrualPlan) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlaps {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveAccrualPlanModel, vv)
}

// UpdateHrLeaveAccrualPlan updates an existing hr.leave.accrual.plan record.
func (c *Client) UpdateHrLeaveAccrualPlan(hlap *HrLeaveAccrualPlan) error {
	return c.UpdateHrLeaveAccrualPlans([]int64{hlap.Id.Get()}, hlap)
}

// UpdateHrLeaveAccrualPlans updates existing hr.leave.accrual.plan records.
// All records (represented by ids) will be updated by hlap values.
func (c *Client) UpdateHrLeaveAccrualPlans(ids []int64, hlap *HrLeaveAccrualPlan) error {
	return c.Update(HrLeaveAccrualPlanModel, ids, hlap)
}

// DeleteHrLeaveAccrualPlan deletes an existing hr.leave.accrual.plan record.
func (c *Client) DeleteHrLeaveAccrualPlan(id int64) error {
	return c.DeleteHrLeaveAccrualPlans([]int64{id})
}

// DeleteHrLeaveAccrualPlans deletes existing hr.leave.accrual.plan records.
func (c *Client) DeleteHrLeaveAccrualPlans(ids []int64) error {
	return c.Delete(HrLeaveAccrualPlanModel, ids)
}

// GetHrLeaveAccrualPlan gets hr.leave.accrual.plan existing record.
func (c *Client) GetHrLeaveAccrualPlan(id int64) (*HrLeaveAccrualPlan, error) {
	hlaps, err := c.GetHrLeaveAccrualPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	if hlaps != nil && len(*hlaps) > 0 {
		return &((*hlaps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.accrual.plan not found", id)
}

// GetHrLeaveAccrualPlans gets hr.leave.accrual.plan existing records.
func (c *Client) GetHrLeaveAccrualPlans(ids []int64) (*HrLeaveAccrualPlans, error) {
	hlaps := &HrLeaveAccrualPlans{}
	if err := c.Read(HrLeaveAccrualPlanModel, ids, nil, hlaps); err != nil {
		return nil, err
	}
	return hlaps, nil
}

// FindHrLeaveAccrualPlan finds hr.leave.accrual.plan record by querying it with criteria.
func (c *Client) FindHrLeaveAccrualPlan(criteria *Criteria) (*HrLeaveAccrualPlan, error) {
	hlaps := &HrLeaveAccrualPlans{}
	if err := c.SearchRead(HrLeaveAccrualPlanModel, criteria, NewOptions().Limit(1), hlaps); err != nil {
		return nil, err
	}
	if hlaps != nil && len(*hlaps) > 0 {
		return &((*hlaps)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.accrual.plan was not found with criteria %v", criteria)
}

// FindHrLeaveAccrualPlans finds hr.leave.accrual.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualPlans(criteria *Criteria, options *Options) (*HrLeaveAccrualPlans, error) {
	hlaps := &HrLeaveAccrualPlans{}
	if err := c.SearchRead(HrLeaveAccrualPlanModel, criteria, options, hlaps); err != nil {
		return nil, err
	}
	return hlaps, nil
}

// FindHrLeaveAccrualPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveAccrualPlanModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveAccrualPlanId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveAccrualPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveAccrualPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.accrual.plan was not found with criteria %v and options %v", criteria, options)
}
