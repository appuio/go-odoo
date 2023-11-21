package odoo

import (
	"fmt"
)

// HrRecruitmentStageReport represents hr.recruitment.stage.report model.
type HrRecruitmentStageReport struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	ApplicantId *Many2One  `xmlrpc:"applicant_id,omptempty"`
	CompanyId   *Many2One  `xmlrpc:"company_id,omptempty"`
	DateBegin   *Time      `xmlrpc:"date_begin,omptempty"`
	DateEnd     *Time      `xmlrpc:"date_end,omptempty"`
	DaysInStage *Float     `xmlrpc:"days_in_stage,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	JobId       *Many2One  `xmlrpc:"job_id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	StageId     *Many2One  `xmlrpc:"stage_id,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
}

// HrRecruitmentStageReports represents array of hr.recruitment.stage.report model.
type HrRecruitmentStageReports []HrRecruitmentStageReport

// HrRecruitmentStageReportModel is the odoo model name.
const HrRecruitmentStageReportModel = "hr.recruitment.stage.report"

// Many2One convert HrRecruitmentStageReport to *Many2One.
func (hrsr *HrRecruitmentStageReport) Many2One() *Many2One {
	return NewMany2One(hrsr.Id.Get(), "")
}

// CreateHrRecruitmentStageReport creates a new hr.recruitment.stage.report model and returns its id.
func (c *Client) CreateHrRecruitmentStageReport(hrsr *HrRecruitmentStageReport) (int64, error) {
	ids, err := c.CreateHrRecruitmentStageReports([]*HrRecruitmentStageReport{hrsr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrRecruitmentStageReport creates a new hr.recruitment.stage.report model and returns its id.
func (c *Client) CreateHrRecruitmentStageReports(hrsrs []*HrRecruitmentStageReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrsrs {
		vv = append(vv, v)
	}
	return c.Create(HrRecruitmentStageReportModel, vv)
}

// UpdateHrRecruitmentStageReport updates an existing hr.recruitment.stage.report record.
func (c *Client) UpdateHrRecruitmentStageReport(hrsr *HrRecruitmentStageReport) error {
	return c.UpdateHrRecruitmentStageReports([]int64{hrsr.Id.Get()}, hrsr)
}

// UpdateHrRecruitmentStageReports updates existing hr.recruitment.stage.report records.
// All records (represented by ids) will be updated by hrsr values.
func (c *Client) UpdateHrRecruitmentStageReports(ids []int64, hrsr *HrRecruitmentStageReport) error {
	return c.Update(HrRecruitmentStageReportModel, ids, hrsr)
}

// DeleteHrRecruitmentStageReport deletes an existing hr.recruitment.stage.report record.
func (c *Client) DeleteHrRecruitmentStageReport(id int64) error {
	return c.DeleteHrRecruitmentStageReports([]int64{id})
}

// DeleteHrRecruitmentStageReports deletes existing hr.recruitment.stage.report records.
func (c *Client) DeleteHrRecruitmentStageReports(ids []int64) error {
	return c.Delete(HrRecruitmentStageReportModel, ids)
}

// GetHrRecruitmentStageReport gets hr.recruitment.stage.report existing record.
func (c *Client) GetHrRecruitmentStageReport(id int64) (*HrRecruitmentStageReport, error) {
	hrsrs, err := c.GetHrRecruitmentStageReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrsrs != nil && len(*hrsrs) > 0 {
		return &((*hrsrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.recruitment.stage.report not found", id)
}

// GetHrRecruitmentStageReports gets hr.recruitment.stage.report existing records.
func (c *Client) GetHrRecruitmentStageReports(ids []int64) (*HrRecruitmentStageReports, error) {
	hrsrs := &HrRecruitmentStageReports{}
	if err := c.Read(HrRecruitmentStageReportModel, ids, nil, hrsrs); err != nil {
		return nil, err
	}
	return hrsrs, nil
}

// FindHrRecruitmentStageReport finds hr.recruitment.stage.report record by querying it with criteria.
func (c *Client) FindHrRecruitmentStageReport(criteria *Criteria) (*HrRecruitmentStageReport, error) {
	hrsrs := &HrRecruitmentStageReports{}
	if err := c.SearchRead(HrRecruitmentStageReportModel, criteria, NewOptions().Limit(1), hrsrs); err != nil {
		return nil, err
	}
	if hrsrs != nil && len(*hrsrs) > 0 {
		return &((*hrsrs)[0]), nil
	}
	return nil, fmt.Errorf("hr.recruitment.stage.report was not found with criteria %v", criteria)
}

// FindHrRecruitmentStageReports finds hr.recruitment.stage.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentStageReports(criteria *Criteria, options *Options) (*HrRecruitmentStageReports, error) {
	hrsrs := &HrRecruitmentStageReports{}
	if err := c.SearchRead(HrRecruitmentStageReportModel, criteria, options, hrsrs); err != nil {
		return nil, err
	}
	return hrsrs, nil
}

// FindHrRecruitmentStageReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentStageReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrRecruitmentStageReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrRecruitmentStageReportId finds record id by querying it with criteria.
func (c *Client) FindHrRecruitmentStageReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRecruitmentStageReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.recruitment.stage.report was not found with criteria %v and options %v", criteria, options)
}
