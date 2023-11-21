package odoo

import (
	"fmt"
)

// HrPayslipEmployees represents hr.payslip.employees model.
type HrPayslipEmployees struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DepartmentId *Many2One `xmlrpc:"department_id,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	EmployeeIds  *Relation `xmlrpc:"employee_ids,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	StructureId  *Many2One `xmlrpc:"structure_id,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipEmployeess represents array of hr.payslip.employees model.
type HrPayslipEmployeess []HrPayslipEmployees

// HrPayslipEmployeesModel is the odoo model name.
const HrPayslipEmployeesModel = "hr.payslip.employees"

// Many2One convert HrPayslipEmployees to *Many2One.
func (hpe *HrPayslipEmployees) Many2One() *Many2One {
	return NewMany2One(hpe.Id.Get(), "")
}

// CreateHrPayslipEmployees creates a new hr.payslip.employees model and returns its id.
func (c *Client) CreateHrPayslipEmployees(hpe *HrPayslipEmployees) (int64, error) {
	ids, err := c.CreateHrPayslipEmployeess([]*HrPayslipEmployees{hpe})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipEmployees creates a new hr.payslip.employees model and returns its id.
func (c *Client) CreateHrPayslipEmployeess(hpes []*HrPayslipEmployees) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpes {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipEmployeesModel, vv)
}

// UpdateHrPayslipEmployees updates an existing hr.payslip.employees record.
func (c *Client) UpdateHrPayslipEmployees(hpe *HrPayslipEmployees) error {
	return c.UpdateHrPayslipEmployeess([]int64{hpe.Id.Get()}, hpe)
}

// UpdateHrPayslipEmployeess updates existing hr.payslip.employees records.
// All records (represented by ids) will be updated by hpe values.
func (c *Client) UpdateHrPayslipEmployeess(ids []int64, hpe *HrPayslipEmployees) error {
	return c.Update(HrPayslipEmployeesModel, ids, hpe)
}

// DeleteHrPayslipEmployees deletes an existing hr.payslip.employees record.
func (c *Client) DeleteHrPayslipEmployees(id int64) error {
	return c.DeleteHrPayslipEmployeess([]int64{id})
}

// DeleteHrPayslipEmployeess deletes existing hr.payslip.employees records.
func (c *Client) DeleteHrPayslipEmployeess(ids []int64) error {
	return c.Delete(HrPayslipEmployeesModel, ids)
}

// GetHrPayslipEmployees gets hr.payslip.employees existing record.
func (c *Client) GetHrPayslipEmployees(id int64) (*HrPayslipEmployees, error) {
	hpes, err := c.GetHrPayslipEmployeess([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpes != nil && len(*hpes) > 0 {
		return &((*hpes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.employees not found", id)
}

// GetHrPayslipEmployeess gets hr.payslip.employees existing records.
func (c *Client) GetHrPayslipEmployeess(ids []int64) (*HrPayslipEmployeess, error) {
	hpes := &HrPayslipEmployeess{}
	if err := c.Read(HrPayslipEmployeesModel, ids, nil, hpes); err != nil {
		return nil, err
	}
	return hpes, nil
}

// FindHrPayslipEmployees finds hr.payslip.employees record by querying it with criteria.
func (c *Client) FindHrPayslipEmployees(criteria *Criteria) (*HrPayslipEmployees, error) {
	hpes := &HrPayslipEmployeess{}
	if err := c.SearchRead(HrPayslipEmployeesModel, criteria, NewOptions().Limit(1), hpes); err != nil {
		return nil, err
	}
	if hpes != nil && len(*hpes) > 0 {
		return &((*hpes)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.employees was not found with criteria %v", criteria)
}

// FindHrPayslipEmployeess finds hr.payslip.employees records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipEmployeess(criteria *Criteria, options *Options) (*HrPayslipEmployeess, error) {
	hpes := &HrPayslipEmployeess{}
	if err := c.SearchRead(HrPayslipEmployeesModel, criteria, options, hpes); err != nil {
		return nil, err
	}
	return hpes, nil
}

// FindHrPayslipEmployeesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipEmployeesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipEmployeesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipEmployeesId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipEmployeesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipEmployeesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.employees was not found with criteria %v and options %v", criteria, options)
}
