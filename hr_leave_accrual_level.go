package odoo

import (
	"fmt"
)

// HrLeaveAccrualLevel represents hr.leave.accrual.level model.
type HrLeaveAccrualLevel struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccrualPlanId            *Many2One  `xmlrpc:"accrual_plan_id,omptempty"`
	ActionWithUnusedAccruals *Selection `xmlrpc:"action_with_unused_accruals,omptempty"`
	AddedValue               *Float     `xmlrpc:"added_value,omptempty"`
	AddedValueType           *Selection `xmlrpc:"added_value_type,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FirstDay                 *Int       `xmlrpc:"first_day,omptempty"`
	FirstDayDisplay          *Selection `xmlrpc:"first_day_display,omptempty"`
	FirstMonth               *Selection `xmlrpc:"first_month,omptempty"`
	FirstMonthDay            *Int       `xmlrpc:"first_month_day,omptempty"`
	FirstMonthDayDisplay     *Selection `xmlrpc:"first_month_day_display,omptempty"`
	Frequency                *Selection `xmlrpc:"frequency,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IsBasedOnWorkedTime      *Bool      `xmlrpc:"is_based_on_worked_time,omptempty"`
	MaximumLeave             *Float     `xmlrpc:"maximum_leave,omptempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omptempty"`
	PostponeMaxDays          *Int       `xmlrpc:"postpone_max_days,omptempty"`
	SecondDay                *Int       `xmlrpc:"second_day,omptempty"`
	SecondDayDisplay         *Selection `xmlrpc:"second_day_display,omptempty"`
	SecondMonth              *Selection `xmlrpc:"second_month,omptempty"`
	SecondMonthDay           *Int       `xmlrpc:"second_month_day,omptempty"`
	SecondMonthDayDisplay    *Selection `xmlrpc:"second_month_day_display,omptempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omptempty"`
	StartCount               *Int       `xmlrpc:"start_count,omptempty"`
	StartType                *Selection `xmlrpc:"start_type,omptempty"`
	WeekDay                  *Selection `xmlrpc:"week_day,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
	YearlyDay                *Int       `xmlrpc:"yearly_day,omptempty"`
	YearlyDayDisplay         *Selection `xmlrpc:"yearly_day_display,omptempty"`
	YearlyMonth              *Selection `xmlrpc:"yearly_month,omptempty"`
}

// HrLeaveAccrualLevels represents array of hr.leave.accrual.level model.
type HrLeaveAccrualLevels []HrLeaveAccrualLevel

// HrLeaveAccrualLevelModel is the odoo model name.
const HrLeaveAccrualLevelModel = "hr.leave.accrual.level"

// Many2One convert HrLeaveAccrualLevel to *Many2One.
func (hlal *HrLeaveAccrualLevel) Many2One() *Many2One {
	return NewMany2One(hlal.Id.Get(), "")
}

// CreateHrLeaveAccrualLevel creates a new hr.leave.accrual.level model and returns its id.
func (c *Client) CreateHrLeaveAccrualLevel(hlal *HrLeaveAccrualLevel) (int64, error) {
	ids, err := c.CreateHrLeaveAccrualLevels([]*HrLeaveAccrualLevel{hlal})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveAccrualLevel creates a new hr.leave.accrual.level model and returns its id.
func (c *Client) CreateHrLeaveAccrualLevels(hlals []*HrLeaveAccrualLevel) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlals {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveAccrualLevelModel, vv)
}

// UpdateHrLeaveAccrualLevel updates an existing hr.leave.accrual.level record.
func (c *Client) UpdateHrLeaveAccrualLevel(hlal *HrLeaveAccrualLevel) error {
	return c.UpdateHrLeaveAccrualLevels([]int64{hlal.Id.Get()}, hlal)
}

// UpdateHrLeaveAccrualLevels updates existing hr.leave.accrual.level records.
// All records (represented by ids) will be updated by hlal values.
func (c *Client) UpdateHrLeaveAccrualLevels(ids []int64, hlal *HrLeaveAccrualLevel) error {
	return c.Update(HrLeaveAccrualLevelModel, ids, hlal)
}

// DeleteHrLeaveAccrualLevel deletes an existing hr.leave.accrual.level record.
func (c *Client) DeleteHrLeaveAccrualLevel(id int64) error {
	return c.DeleteHrLeaveAccrualLevels([]int64{id})
}

// DeleteHrLeaveAccrualLevels deletes existing hr.leave.accrual.level records.
func (c *Client) DeleteHrLeaveAccrualLevels(ids []int64) error {
	return c.Delete(HrLeaveAccrualLevelModel, ids)
}

// GetHrLeaveAccrualLevel gets hr.leave.accrual.level existing record.
func (c *Client) GetHrLeaveAccrualLevel(id int64) (*HrLeaveAccrualLevel, error) {
	hlals, err := c.GetHrLeaveAccrualLevels([]int64{id})
	if err != nil {
		return nil, err
	}
	if hlals != nil && len(*hlals) > 0 {
		return &((*hlals)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.accrual.level not found", id)
}

// GetHrLeaveAccrualLevels gets hr.leave.accrual.level existing records.
func (c *Client) GetHrLeaveAccrualLevels(ids []int64) (*HrLeaveAccrualLevels, error) {
	hlals := &HrLeaveAccrualLevels{}
	if err := c.Read(HrLeaveAccrualLevelModel, ids, nil, hlals); err != nil {
		return nil, err
	}
	return hlals, nil
}

// FindHrLeaveAccrualLevel finds hr.leave.accrual.level record by querying it with criteria.
func (c *Client) FindHrLeaveAccrualLevel(criteria *Criteria) (*HrLeaveAccrualLevel, error) {
	hlals := &HrLeaveAccrualLevels{}
	if err := c.SearchRead(HrLeaveAccrualLevelModel, criteria, NewOptions().Limit(1), hlals); err != nil {
		return nil, err
	}
	if hlals != nil && len(*hlals) > 0 {
		return &((*hlals)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.accrual.level was not found with criteria %v", criteria)
}

// FindHrLeaveAccrualLevels finds hr.leave.accrual.level records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualLevels(criteria *Criteria, options *Options) (*HrLeaveAccrualLevels, error) {
	hlals := &HrLeaveAccrualLevels{}
	if err := c.SearchRead(HrLeaveAccrualLevelModel, criteria, options, hlals); err != nil {
		return nil, err
	}
	return hlals, nil
}

// FindHrLeaveAccrualLevelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveAccrualLevelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveAccrualLevelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveAccrualLevelId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveAccrualLevelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveAccrualLevelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.accrual.level was not found with criteria %v and options %v", criteria, options)
}
