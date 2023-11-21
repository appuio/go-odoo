package odoo

import (
	"fmt"
)

// HrLeaveEmployeeTypeReport represents hr.leave.employee.type.report model.
type HrLeaveEmployeeTypeReport struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	ActiveEmployee *Bool      `xmlrpc:"active_employee,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	DateFrom       *Time      `xmlrpc:"date_from,omptempty"`
	DateTo         *Time      `xmlrpc:"date_to,omptempty"`
	DepartmentId   *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId     *Many2One  `xmlrpc:"employee_id,omptempty"`
	HolidayStatus  *Selection `xmlrpc:"holiday_status,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	LeaveType      *Many2One  `xmlrpc:"leave_type,omptempty"`
	NumberOfDays   *Float     `xmlrpc:"number_of_days,omptempty"`
	State          *Selection `xmlrpc:"state,omptempty"`
}

// HrLeaveEmployeeTypeReports represents array of hr.leave.employee.type.report model.
type HrLeaveEmployeeTypeReports []HrLeaveEmployeeTypeReport

// HrLeaveEmployeeTypeReportModel is the odoo model name.
const HrLeaveEmployeeTypeReportModel = "hr.leave.employee.type.report"

// Many2One convert HrLeaveEmployeeTypeReport to *Many2One.
func (hletr *HrLeaveEmployeeTypeReport) Many2One() *Many2One {
	return NewMany2One(hletr.Id.Get(), "")
}

// CreateHrLeaveEmployeeTypeReport creates a new hr.leave.employee.type.report model and returns its id.
func (c *Client) CreateHrLeaveEmployeeTypeReport(hletr *HrLeaveEmployeeTypeReport) (int64, error) {
	ids, err := c.CreateHrLeaveEmployeeTypeReports([]*HrLeaveEmployeeTypeReport{hletr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveEmployeeTypeReport creates a new hr.leave.employee.type.report model and returns its id.
func (c *Client) CreateHrLeaveEmployeeTypeReports(hletrs []*HrLeaveEmployeeTypeReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hletrs {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveEmployeeTypeReportModel, vv)
}

// UpdateHrLeaveEmployeeTypeReport updates an existing hr.leave.employee.type.report record.
func (c *Client) UpdateHrLeaveEmployeeTypeReport(hletr *HrLeaveEmployeeTypeReport) error {
	return c.UpdateHrLeaveEmployeeTypeReports([]int64{hletr.Id.Get()}, hletr)
}

// UpdateHrLeaveEmployeeTypeReports updates existing hr.leave.employee.type.report records.
// All records (represented by ids) will be updated by hletr values.
func (c *Client) UpdateHrLeaveEmployeeTypeReports(ids []int64, hletr *HrLeaveEmployeeTypeReport) error {
	return c.Update(HrLeaveEmployeeTypeReportModel, ids, hletr)
}

// DeleteHrLeaveEmployeeTypeReport deletes an existing hr.leave.employee.type.report record.
func (c *Client) DeleteHrLeaveEmployeeTypeReport(id int64) error {
	return c.DeleteHrLeaveEmployeeTypeReports([]int64{id})
}

// DeleteHrLeaveEmployeeTypeReports deletes existing hr.leave.employee.type.report records.
func (c *Client) DeleteHrLeaveEmployeeTypeReports(ids []int64) error {
	return c.Delete(HrLeaveEmployeeTypeReportModel, ids)
}

// GetHrLeaveEmployeeTypeReport gets hr.leave.employee.type.report existing record.
func (c *Client) GetHrLeaveEmployeeTypeReport(id int64) (*HrLeaveEmployeeTypeReport, error) {
	hletrs, err := c.GetHrLeaveEmployeeTypeReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hletrs != nil && len(*hletrs) > 0 {
		return &((*hletrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.employee.type.report not found", id)
}

// GetHrLeaveEmployeeTypeReports gets hr.leave.employee.type.report existing records.
func (c *Client) GetHrLeaveEmployeeTypeReports(ids []int64) (*HrLeaveEmployeeTypeReports, error) {
	hletrs := &HrLeaveEmployeeTypeReports{}
	if err := c.Read(HrLeaveEmployeeTypeReportModel, ids, nil, hletrs); err != nil {
		return nil, err
	}
	return hletrs, nil
}

// FindHrLeaveEmployeeTypeReport finds hr.leave.employee.type.report record by querying it with criteria.
func (c *Client) FindHrLeaveEmployeeTypeReport(criteria *Criteria) (*HrLeaveEmployeeTypeReport, error) {
	hletrs := &HrLeaveEmployeeTypeReports{}
	if err := c.SearchRead(HrLeaveEmployeeTypeReportModel, criteria, NewOptions().Limit(1), hletrs); err != nil {
		return nil, err
	}
	if hletrs != nil && len(*hletrs) > 0 {
		return &((*hletrs)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.employee.type.report was not found with criteria %v", criteria)
}

// FindHrLeaveEmployeeTypeReports finds hr.leave.employee.type.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveEmployeeTypeReports(criteria *Criteria, options *Options) (*HrLeaveEmployeeTypeReports, error) {
	hletrs := &HrLeaveEmployeeTypeReports{}
	if err := c.SearchRead(HrLeaveEmployeeTypeReportModel, criteria, options, hletrs); err != nil {
		return nil, err
	}
	return hletrs, nil
}

// FindHrLeaveEmployeeTypeReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveEmployeeTypeReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveEmployeeTypeReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveEmployeeTypeReportId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveEmployeeTypeReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveEmployeeTypeReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.employee.type.report was not found with criteria %v and options %v", criteria, options)
}
