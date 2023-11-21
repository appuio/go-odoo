package odoo

import (
	"fmt"
)

// HrUserWorkEntryEmployee represents hr.user.work.entry.employee model.
type HrUserWorkEntryEmployee struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Active      *Bool     `xmlrpc:"active,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrUserWorkEntryEmployees represents array of hr.user.work.entry.employee model.
type HrUserWorkEntryEmployees []HrUserWorkEntryEmployee

// HrUserWorkEntryEmployeeModel is the odoo model name.
const HrUserWorkEntryEmployeeModel = "hr.user.work.entry.employee"

// Many2One convert HrUserWorkEntryEmployee to *Many2One.
func (huwee *HrUserWorkEntryEmployee) Many2One() *Many2One {
	return NewMany2One(huwee.Id.Get(), "")
}

// CreateHrUserWorkEntryEmployee creates a new hr.user.work.entry.employee model and returns its id.
func (c *Client) CreateHrUserWorkEntryEmployee(huwee *HrUserWorkEntryEmployee) (int64, error) {
	ids, err := c.CreateHrUserWorkEntryEmployees([]*HrUserWorkEntryEmployee{huwee})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrUserWorkEntryEmployee creates a new hr.user.work.entry.employee model and returns its id.
func (c *Client) CreateHrUserWorkEntryEmployees(huwees []*HrUserWorkEntryEmployee) ([]int64, error) {
	var vv []interface{}
	for _, v := range huwees {
		vv = append(vv, v)
	}
	return c.Create(HrUserWorkEntryEmployeeModel, vv)
}

// UpdateHrUserWorkEntryEmployee updates an existing hr.user.work.entry.employee record.
func (c *Client) UpdateHrUserWorkEntryEmployee(huwee *HrUserWorkEntryEmployee) error {
	return c.UpdateHrUserWorkEntryEmployees([]int64{huwee.Id.Get()}, huwee)
}

// UpdateHrUserWorkEntryEmployees updates existing hr.user.work.entry.employee records.
// All records (represented by ids) will be updated by huwee values.
func (c *Client) UpdateHrUserWorkEntryEmployees(ids []int64, huwee *HrUserWorkEntryEmployee) error {
	return c.Update(HrUserWorkEntryEmployeeModel, ids, huwee)
}

// DeleteHrUserWorkEntryEmployee deletes an existing hr.user.work.entry.employee record.
func (c *Client) DeleteHrUserWorkEntryEmployee(id int64) error {
	return c.DeleteHrUserWorkEntryEmployees([]int64{id})
}

// DeleteHrUserWorkEntryEmployees deletes existing hr.user.work.entry.employee records.
func (c *Client) DeleteHrUserWorkEntryEmployees(ids []int64) error {
	return c.Delete(HrUserWorkEntryEmployeeModel, ids)
}

// GetHrUserWorkEntryEmployee gets hr.user.work.entry.employee existing record.
func (c *Client) GetHrUserWorkEntryEmployee(id int64) (*HrUserWorkEntryEmployee, error) {
	huwees, err := c.GetHrUserWorkEntryEmployees([]int64{id})
	if err != nil {
		return nil, err
	}
	if huwees != nil && len(*huwees) > 0 {
		return &((*huwees)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.user.work.entry.employee not found", id)
}

// GetHrUserWorkEntryEmployees gets hr.user.work.entry.employee existing records.
func (c *Client) GetHrUserWorkEntryEmployees(ids []int64) (*HrUserWorkEntryEmployees, error) {
	huwees := &HrUserWorkEntryEmployees{}
	if err := c.Read(HrUserWorkEntryEmployeeModel, ids, nil, huwees); err != nil {
		return nil, err
	}
	return huwees, nil
}

// FindHrUserWorkEntryEmployee finds hr.user.work.entry.employee record by querying it with criteria.
func (c *Client) FindHrUserWorkEntryEmployee(criteria *Criteria) (*HrUserWorkEntryEmployee, error) {
	huwees := &HrUserWorkEntryEmployees{}
	if err := c.SearchRead(HrUserWorkEntryEmployeeModel, criteria, NewOptions().Limit(1), huwees); err != nil {
		return nil, err
	}
	if huwees != nil && len(*huwees) > 0 {
		return &((*huwees)[0]), nil
	}
	return nil, fmt.Errorf("hr.user.work.entry.employee was not found with criteria %v", criteria)
}

// FindHrUserWorkEntryEmployees finds hr.user.work.entry.employee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrUserWorkEntryEmployees(criteria *Criteria, options *Options) (*HrUserWorkEntryEmployees, error) {
	huwees := &HrUserWorkEntryEmployees{}
	if err := c.SearchRead(HrUserWorkEntryEmployeeModel, criteria, options, huwees); err != nil {
		return nil, err
	}
	return huwees, nil
}

// FindHrUserWorkEntryEmployeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrUserWorkEntryEmployeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrUserWorkEntryEmployeeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrUserWorkEntryEmployeeId finds record id by querying it with criteria.
func (c *Client) FindHrUserWorkEntryEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrUserWorkEntryEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.user.work.entry.employee was not found with criteria %v and options %v", criteria, options)
}
