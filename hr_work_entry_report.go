package odoo

import (
	"fmt"
)

// HrWorkEntryReport represents hr.work.entry.report model.
type HrWorkEntryReport struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	DateStart       *Time      `xmlrpc:"date_start,omptempty"`
	DepartmentId    *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId      *Many2One  `xmlrpc:"employee_id,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	NumberOfDays    *Float     `xmlrpc:"number_of_days,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
	WorkEntrySource *Selection `xmlrpc:"work_entry_source,omptempty"`
	WorkEntryTypeId *Many2One  `xmlrpc:"work_entry_type_id,omptempty"`
}

// HrWorkEntryReports represents array of hr.work.entry.report model.
type HrWorkEntryReports []HrWorkEntryReport

// HrWorkEntryReportModel is the odoo model name.
const HrWorkEntryReportModel = "hr.work.entry.report"

// Many2One convert HrWorkEntryReport to *Many2One.
func (hwer *HrWorkEntryReport) Many2One() *Many2One {
	return NewMany2One(hwer.Id.Get(), "")
}

// CreateHrWorkEntryReport creates a new hr.work.entry.report model and returns its id.
func (c *Client) CreateHrWorkEntryReport(hwer *HrWorkEntryReport) (int64, error) {
	ids, err := c.CreateHrWorkEntryReports([]*HrWorkEntryReport{hwer})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrWorkEntryReport creates a new hr.work.entry.report model and returns its id.
func (c *Client) CreateHrWorkEntryReports(hwers []*HrWorkEntryReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hwers {
		vv = append(vv, v)
	}
	return c.Create(HrWorkEntryReportModel, vv)
}

// UpdateHrWorkEntryReport updates an existing hr.work.entry.report record.
func (c *Client) UpdateHrWorkEntryReport(hwer *HrWorkEntryReport) error {
	return c.UpdateHrWorkEntryReports([]int64{hwer.Id.Get()}, hwer)
}

// UpdateHrWorkEntryReports updates existing hr.work.entry.report records.
// All records (represented by ids) will be updated by hwer values.
func (c *Client) UpdateHrWorkEntryReports(ids []int64, hwer *HrWorkEntryReport) error {
	return c.Update(HrWorkEntryReportModel, ids, hwer)
}

// DeleteHrWorkEntryReport deletes an existing hr.work.entry.report record.
func (c *Client) DeleteHrWorkEntryReport(id int64) error {
	return c.DeleteHrWorkEntryReports([]int64{id})
}

// DeleteHrWorkEntryReports deletes existing hr.work.entry.report records.
func (c *Client) DeleteHrWorkEntryReports(ids []int64) error {
	return c.Delete(HrWorkEntryReportModel, ids)
}

// GetHrWorkEntryReport gets hr.work.entry.report existing record.
func (c *Client) GetHrWorkEntryReport(id int64) (*HrWorkEntryReport, error) {
	hwers, err := c.GetHrWorkEntryReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hwers != nil && len(*hwers) > 0 {
		return &((*hwers)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.work.entry.report not found", id)
}

// GetHrWorkEntryReports gets hr.work.entry.report existing records.
func (c *Client) GetHrWorkEntryReports(ids []int64) (*HrWorkEntryReports, error) {
	hwers := &HrWorkEntryReports{}
	if err := c.Read(HrWorkEntryReportModel, ids, nil, hwers); err != nil {
		return nil, err
	}
	return hwers, nil
}

// FindHrWorkEntryReport finds hr.work.entry.report record by querying it with criteria.
func (c *Client) FindHrWorkEntryReport(criteria *Criteria) (*HrWorkEntryReport, error) {
	hwers := &HrWorkEntryReports{}
	if err := c.SearchRead(HrWorkEntryReportModel, criteria, NewOptions().Limit(1), hwers); err != nil {
		return nil, err
	}
	if hwers != nil && len(*hwers) > 0 {
		return &((*hwers)[0]), nil
	}
	return nil, fmt.Errorf("hr.work.entry.report was not found with criteria %v", criteria)
}

// FindHrWorkEntryReports finds hr.work.entry.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntryReports(criteria *Criteria, options *Options) (*HrWorkEntryReports, error) {
	hwers := &HrWorkEntryReports{}
	if err := c.SearchRead(HrWorkEntryReportModel, criteria, options, hwers); err != nil {
		return nil, err
	}
	return hwers, nil
}

// FindHrWorkEntryReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntryReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrWorkEntryReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrWorkEntryReportId finds record id by querying it with criteria.
func (c *Client) FindHrWorkEntryReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrWorkEntryReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.work.entry.report was not found with criteria %v and options %v", criteria, options)
}
