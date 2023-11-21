package odoo

import (
	"fmt"
)

// HrAppraisalReport represents hr.appraisal.report model.
type HrAppraisalReport struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Color          *Int       `xmlrpc:"color,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	Deadline       *Time      `xmlrpc:"deadline,omptempty"`
	DepartmentId   *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId     *Many2One  `xmlrpc:"employee_id,omptempty"`
	FinalInterview *Time      `xmlrpc:"final_interview,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	State          *Selection `xmlrpc:"state,omptempty"`
}

// HrAppraisalReports represents array of hr.appraisal.report model.
type HrAppraisalReports []HrAppraisalReport

// HrAppraisalReportModel is the odoo model name.
const HrAppraisalReportModel = "hr.appraisal.report"

// Many2One convert HrAppraisalReport to *Many2One.
func (har *HrAppraisalReport) Many2One() *Many2One {
	return NewMany2One(har.Id.Get(), "")
}

// CreateHrAppraisalReport creates a new hr.appraisal.report model and returns its id.
func (c *Client) CreateHrAppraisalReport(har *HrAppraisalReport) (int64, error) {
	ids, err := c.CreateHrAppraisalReports([]*HrAppraisalReport{har})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAppraisalReport creates a new hr.appraisal.report model and returns its id.
func (c *Client) CreateHrAppraisalReports(hars []*HrAppraisalReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hars {
		vv = append(vv, v)
	}
	return c.Create(HrAppraisalReportModel, vv)
}

// UpdateHrAppraisalReport updates an existing hr.appraisal.report record.
func (c *Client) UpdateHrAppraisalReport(har *HrAppraisalReport) error {
	return c.UpdateHrAppraisalReports([]int64{har.Id.Get()}, har)
}

// UpdateHrAppraisalReports updates existing hr.appraisal.report records.
// All records (represented by ids) will be updated by har values.
func (c *Client) UpdateHrAppraisalReports(ids []int64, har *HrAppraisalReport) error {
	return c.Update(HrAppraisalReportModel, ids, har)
}

// DeleteHrAppraisalReport deletes an existing hr.appraisal.report record.
func (c *Client) DeleteHrAppraisalReport(id int64) error {
	return c.DeleteHrAppraisalReports([]int64{id})
}

// DeleteHrAppraisalReports deletes existing hr.appraisal.report records.
func (c *Client) DeleteHrAppraisalReports(ids []int64) error {
	return c.Delete(HrAppraisalReportModel, ids)
}

// GetHrAppraisalReport gets hr.appraisal.report existing record.
func (c *Client) GetHrAppraisalReport(id int64) (*HrAppraisalReport, error) {
	hars, err := c.GetHrAppraisalReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hars != nil && len(*hars) > 0 {
		return &((*hars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.appraisal.report not found", id)
}

// GetHrAppraisalReports gets hr.appraisal.report existing records.
func (c *Client) GetHrAppraisalReports(ids []int64) (*HrAppraisalReports, error) {
	hars := &HrAppraisalReports{}
	if err := c.Read(HrAppraisalReportModel, ids, nil, hars); err != nil {
		return nil, err
	}
	return hars, nil
}

// FindHrAppraisalReport finds hr.appraisal.report record by querying it with criteria.
func (c *Client) FindHrAppraisalReport(criteria *Criteria) (*HrAppraisalReport, error) {
	hars := &HrAppraisalReports{}
	if err := c.SearchRead(HrAppraisalReportModel, criteria, NewOptions().Limit(1), hars); err != nil {
		return nil, err
	}
	if hars != nil && len(*hars) > 0 {
		return &((*hars)[0]), nil
	}
	return nil, fmt.Errorf("hr.appraisal.report was not found with criteria %v", criteria)
}

// FindHrAppraisalReports finds hr.appraisal.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalReports(criteria *Criteria, options *Options) (*HrAppraisalReports, error) {
	hars := &HrAppraisalReports{}
	if err := c.SearchRead(HrAppraisalReportModel, criteria, options, hars); err != nil {
		return nil, err
	}
	return hars, nil
}

// FindHrAppraisalReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAppraisalReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAppraisalReportId finds record id by querying it with criteria.
func (c *Client) FindHrAppraisalReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAppraisalReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.appraisal.report was not found with criteria %v and options %v", criteria, options)
}
