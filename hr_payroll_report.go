package odoo

import (
	"fmt"
)

// HrPayrollReport represents hr.payroll.report model.
type HrPayrollReport struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	BasicWage              *Float     `xmlrpc:"basic_wage,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	Count                  *Int       `xmlrpc:"count,omptempty"`
	CountLeave             *Int       `xmlrpc:"count_leave,omptempty"`
	CountLeaveUnpaid       *Int       `xmlrpc:"count_leave_unpaid,omptempty"`
	CountUnforeseenAbsence *Int       `xmlrpc:"count_unforeseen_absence,omptempty"`
	CountWork              *Int       `xmlrpc:"count_work,omptempty"`
	CountWorkHours         *Int       `xmlrpc:"count_work_hours,omptempty"`
	DateFrom               *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                 *Time      `xmlrpc:"date_to,omptempty"`
	DepartmentId           *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId             *Many2One  `xmlrpc:"employee_id,omptempty"`
	GrossWage              *Float     `xmlrpc:"gross_wage,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	JobId                  *Many2One  `xmlrpc:"job_id,omptempty"`
	LeaveBasicWage         *Float     `xmlrpc:"leave_basic_wage,omptempty"`
	MasterDepartmentId     *Many2One  `xmlrpc:"master_department_id,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	NetWage                *Float     `xmlrpc:"net_wage,omptempty"`
	NumberOfDays           *Float     `xmlrpc:"number_of_days,omptempty"`
	NumberOfHours          *Float     `xmlrpc:"number_of_hours,omptempty"`
	WorkCode               *Many2One  `xmlrpc:"work_code,omptempty"`
	WorkEntrySource        *Selection `xmlrpc:"work_entry_source,omptempty"`
	WorkType               *Selection `xmlrpc:"work_type,omptempty"`
	XL10NXxPaid            *Float     `xmlrpc:"x_l10n_xx_paid,omptempty"`
}

// HrPayrollReports represents array of hr.payroll.report model.
type HrPayrollReports []HrPayrollReport

// HrPayrollReportModel is the odoo model name.
const HrPayrollReportModel = "hr.payroll.report"

// Many2One convert HrPayrollReport to *Many2One.
func (hpr *HrPayrollReport) Many2One() *Many2One {
	return NewMany2One(hpr.Id.Get(), "")
}

// CreateHrPayrollReport creates a new hr.payroll.report model and returns its id.
func (c *Client) CreateHrPayrollReport(hpr *HrPayrollReport) (int64, error) {
	ids, err := c.CreateHrPayrollReports([]*HrPayrollReport{hpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayrollReport creates a new hr.payroll.report model and returns its id.
func (c *Client) CreateHrPayrollReports(hprs []*HrPayrollReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hprs {
		vv = append(vv, v)
	}
	return c.Create(HrPayrollReportModel, vv)
}

// UpdateHrPayrollReport updates an existing hr.payroll.report record.
func (c *Client) UpdateHrPayrollReport(hpr *HrPayrollReport) error {
	return c.UpdateHrPayrollReports([]int64{hpr.Id.Get()}, hpr)
}

// UpdateHrPayrollReports updates existing hr.payroll.report records.
// All records (represented by ids) will be updated by hpr values.
func (c *Client) UpdateHrPayrollReports(ids []int64, hpr *HrPayrollReport) error {
	return c.Update(HrPayrollReportModel, ids, hpr)
}

// DeleteHrPayrollReport deletes an existing hr.payroll.report record.
func (c *Client) DeleteHrPayrollReport(id int64) error {
	return c.DeleteHrPayrollReports([]int64{id})
}

// DeleteHrPayrollReports deletes existing hr.payroll.report records.
func (c *Client) DeleteHrPayrollReports(ids []int64) error {
	return c.Delete(HrPayrollReportModel, ids)
}

// GetHrPayrollReport gets hr.payroll.report existing record.
func (c *Client) GetHrPayrollReport(id int64) (*HrPayrollReport, error) {
	hprs, err := c.GetHrPayrollReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hprs != nil && len(*hprs) > 0 {
		return &((*hprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payroll.report not found", id)
}

// GetHrPayrollReports gets hr.payroll.report existing records.
func (c *Client) GetHrPayrollReports(ids []int64) (*HrPayrollReports, error) {
	hprs := &HrPayrollReports{}
	if err := c.Read(HrPayrollReportModel, ids, nil, hprs); err != nil {
		return nil, err
	}
	return hprs, nil
}

// FindHrPayrollReport finds hr.payroll.report record by querying it with criteria.
func (c *Client) FindHrPayrollReport(criteria *Criteria) (*HrPayrollReport, error) {
	hprs := &HrPayrollReports{}
	if err := c.SearchRead(HrPayrollReportModel, criteria, NewOptions().Limit(1), hprs); err != nil {
		return nil, err
	}
	if hprs != nil && len(*hprs) > 0 {
		return &((*hprs)[0]), nil
	}
	return nil, fmt.Errorf("hr.payroll.report was not found with criteria %v", criteria)
}

// FindHrPayrollReports finds hr.payroll.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollReports(criteria *Criteria, options *Options) (*HrPayrollReports, error) {
	hprs := &HrPayrollReports{}
	if err := c.SearchRead(HrPayrollReportModel, criteria, options, hprs); err != nil {
		return nil, err
	}
	return hprs, nil
}

// FindHrPayrollReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayrollReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayrollReportId finds record id by querying it with criteria.
func (c *Client) FindHrPayrollReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayrollReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payroll.report was not found with criteria %v and options %v", criteria, options)
}
