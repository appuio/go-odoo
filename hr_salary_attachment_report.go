package odoo

import (
	"fmt"
)

// HrSalaryAttachmentReport represents hr.salary.attachment.report model.
type HrSalaryAttachmentReport struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	AssignmentAmount   *Float    `xmlrpc:"assignment_amount,omptempty"`
	AttachmentAmount   *Float    `xmlrpc:"attachment_amount,omptempty"`
	ChildSupportAmount *Float    `xmlrpc:"child_support_amount,omptempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId         *Many2One `xmlrpc:"employee_id,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	PayslipEndDate     *Time     `xmlrpc:"payslip_end_date,omptempty"`
	PayslipId          *Many2One `xmlrpc:"payslip_id,omptempty"`
	PayslipStartDate   *Time     `xmlrpc:"payslip_start_date,omptempty"`
}

// HrSalaryAttachmentReports represents array of hr.salary.attachment.report model.
type HrSalaryAttachmentReports []HrSalaryAttachmentReport

// HrSalaryAttachmentReportModel is the odoo model name.
const HrSalaryAttachmentReportModel = "hr.salary.attachment.report"

// Many2One convert HrSalaryAttachmentReport to *Many2One.
func (hsar *HrSalaryAttachmentReport) Many2One() *Many2One {
	return NewMany2One(hsar.Id.Get(), "")
}

// CreateHrSalaryAttachmentReport creates a new hr.salary.attachment.report model and returns its id.
func (c *Client) CreateHrSalaryAttachmentReport(hsar *HrSalaryAttachmentReport) (int64, error) {
	ids, err := c.CreateHrSalaryAttachmentReports([]*HrSalaryAttachmentReport{hsar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSalaryAttachmentReport creates a new hr.salary.attachment.report model and returns its id.
func (c *Client) CreateHrSalaryAttachmentReports(hsars []*HrSalaryAttachmentReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsars {
		vv = append(vv, v)
	}
	return c.Create(HrSalaryAttachmentReportModel, vv)
}

// UpdateHrSalaryAttachmentReport updates an existing hr.salary.attachment.report record.
func (c *Client) UpdateHrSalaryAttachmentReport(hsar *HrSalaryAttachmentReport) error {
	return c.UpdateHrSalaryAttachmentReports([]int64{hsar.Id.Get()}, hsar)
}

// UpdateHrSalaryAttachmentReports updates existing hr.salary.attachment.report records.
// All records (represented by ids) will be updated by hsar values.
func (c *Client) UpdateHrSalaryAttachmentReports(ids []int64, hsar *HrSalaryAttachmentReport) error {
	return c.Update(HrSalaryAttachmentReportModel, ids, hsar)
}

// DeleteHrSalaryAttachmentReport deletes an existing hr.salary.attachment.report record.
func (c *Client) DeleteHrSalaryAttachmentReport(id int64) error {
	return c.DeleteHrSalaryAttachmentReports([]int64{id})
}

// DeleteHrSalaryAttachmentReports deletes existing hr.salary.attachment.report records.
func (c *Client) DeleteHrSalaryAttachmentReports(ids []int64) error {
	return c.Delete(HrSalaryAttachmentReportModel, ids)
}

// GetHrSalaryAttachmentReport gets hr.salary.attachment.report existing record.
func (c *Client) GetHrSalaryAttachmentReport(id int64) (*HrSalaryAttachmentReport, error) {
	hsars, err := c.GetHrSalaryAttachmentReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsars != nil && len(*hsars) > 0 {
		return &((*hsars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.salary.attachment.report not found", id)
}

// GetHrSalaryAttachmentReports gets hr.salary.attachment.report existing records.
func (c *Client) GetHrSalaryAttachmentReports(ids []int64) (*HrSalaryAttachmentReports, error) {
	hsars := &HrSalaryAttachmentReports{}
	if err := c.Read(HrSalaryAttachmentReportModel, ids, nil, hsars); err != nil {
		return nil, err
	}
	return hsars, nil
}

// FindHrSalaryAttachmentReport finds hr.salary.attachment.report record by querying it with criteria.
func (c *Client) FindHrSalaryAttachmentReport(criteria *Criteria) (*HrSalaryAttachmentReport, error) {
	hsars := &HrSalaryAttachmentReports{}
	if err := c.SearchRead(HrSalaryAttachmentReportModel, criteria, NewOptions().Limit(1), hsars); err != nil {
		return nil, err
	}
	if hsars != nil && len(*hsars) > 0 {
		return &((*hsars)[0]), nil
	}
	return nil, fmt.Errorf("hr.salary.attachment.report was not found with criteria %v", criteria)
}

// FindHrSalaryAttachmentReports finds hr.salary.attachment.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryAttachmentReports(criteria *Criteria, options *Options) (*HrSalaryAttachmentReports, error) {
	hsars := &HrSalaryAttachmentReports{}
	if err := c.SearchRead(HrSalaryAttachmentReportModel, criteria, options, hsars); err != nil {
		return nil, err
	}
	return hsars, nil
}

// FindHrSalaryAttachmentReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryAttachmentReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrSalaryAttachmentReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrSalaryAttachmentReportId finds record id by querying it with criteria.
func (c *Client) FindHrSalaryAttachmentReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSalaryAttachmentReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.salary.attachment.report was not found with criteria %v and options %v", criteria, options)
}
