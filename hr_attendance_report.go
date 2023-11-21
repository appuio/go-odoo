package odoo

import (
	"fmt"
)

// HrAttendanceReport represents hr.attendance.report model.
type HrAttendanceReport struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CheckIn       *Time     `xmlrpc:"check_in,omptempty"`
	CompanyId     *Many2One `xmlrpc:"company_id,omptempty"`
	DepartmentId  *Many2One `xmlrpc:"department_id,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	OvertimeHours *Float    `xmlrpc:"overtime_hours,omptempty"`
	WorkedHours   *Float    `xmlrpc:"worked_hours,omptempty"`
}

// HrAttendanceReports represents array of hr.attendance.report model.
type HrAttendanceReports []HrAttendanceReport

// HrAttendanceReportModel is the odoo model name.
const HrAttendanceReportModel = "hr.attendance.report"

// Many2One convert HrAttendanceReport to *Many2One.
func (har *HrAttendanceReport) Many2One() *Many2One {
	return NewMany2One(har.Id.Get(), "")
}

// CreateHrAttendanceReport creates a new hr.attendance.report model and returns its id.
func (c *Client) CreateHrAttendanceReport(har *HrAttendanceReport) (int64, error) {
	ids, err := c.CreateHrAttendanceReports([]*HrAttendanceReport{har})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAttendanceReport creates a new hr.attendance.report model and returns its id.
func (c *Client) CreateHrAttendanceReports(hars []*HrAttendanceReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hars {
		vv = append(vv, v)
	}
	return c.Create(HrAttendanceReportModel, vv)
}

// UpdateHrAttendanceReport updates an existing hr.attendance.report record.
func (c *Client) UpdateHrAttendanceReport(har *HrAttendanceReport) error {
	return c.UpdateHrAttendanceReports([]int64{har.Id.Get()}, har)
}

// UpdateHrAttendanceReports updates existing hr.attendance.report records.
// All records (represented by ids) will be updated by har values.
func (c *Client) UpdateHrAttendanceReports(ids []int64, har *HrAttendanceReport) error {
	return c.Update(HrAttendanceReportModel, ids, har)
}

// DeleteHrAttendanceReport deletes an existing hr.attendance.report record.
func (c *Client) DeleteHrAttendanceReport(id int64) error {
	return c.DeleteHrAttendanceReports([]int64{id})
}

// DeleteHrAttendanceReports deletes existing hr.attendance.report records.
func (c *Client) DeleteHrAttendanceReports(ids []int64) error {
	return c.Delete(HrAttendanceReportModel, ids)
}

// GetHrAttendanceReport gets hr.attendance.report existing record.
func (c *Client) GetHrAttendanceReport(id int64) (*HrAttendanceReport, error) {
	hars, err := c.GetHrAttendanceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hars != nil && len(*hars) > 0 {
		return &((*hars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.attendance.report not found", id)
}

// GetHrAttendanceReports gets hr.attendance.report existing records.
func (c *Client) GetHrAttendanceReports(ids []int64) (*HrAttendanceReports, error) {
	hars := &HrAttendanceReports{}
	if err := c.Read(HrAttendanceReportModel, ids, nil, hars); err != nil {
		return nil, err
	}
	return hars, nil
}

// FindHrAttendanceReport finds hr.attendance.report record by querying it with criteria.
func (c *Client) FindHrAttendanceReport(criteria *Criteria) (*HrAttendanceReport, error) {
	hars := &HrAttendanceReports{}
	if err := c.SearchRead(HrAttendanceReportModel, criteria, NewOptions().Limit(1), hars); err != nil {
		return nil, err
	}
	if hars != nil && len(*hars) > 0 {
		return &((*hars)[0]), nil
	}
	return nil, fmt.Errorf("hr.attendance.report was not found with criteria %v", criteria)
}

// FindHrAttendanceReports finds hr.attendance.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendanceReports(criteria *Criteria, options *Options) (*HrAttendanceReports, error) {
	hars := &HrAttendanceReports{}
	if err := c.SearchRead(HrAttendanceReportModel, criteria, options, hars); err != nil {
		return nil, err
	}
	return hars, nil
}

// FindHrAttendanceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAttendanceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAttendanceReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAttendanceReportId finds record id by querying it with criteria.
func (c *Client) FindHrAttendanceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAttendanceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.attendance.report was not found with criteria %v and options %v", criteria, options)
}
