package odoo

import (
	"fmt"
)

// HrRecruitmentReport represents hr.recruitment.report model.
type HrRecruitmentReport struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	Count           *Int       `xmlrpc:"count,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateClosed      *Time      `xmlrpc:"date_closed,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Hired           *Int       `xmlrpc:"hired,omptempty"`
	HiringRatio     *Int       `xmlrpc:"hiring_ratio,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	JobId           *Many2One  `xmlrpc:"job_id,omptempty"`
	MediumId        *Many2One  `xmlrpc:"medium_id,omptempty"`
	MeetingsAmount  *Int       `xmlrpc:"meetings_amount,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	ProcessDuration *Int       `xmlrpc:"process_duration,omptempty"`
	RefuseReasonId  *Many2One  `xmlrpc:"refuse_reason_id,omptempty"`
	Refused         *Int       `xmlrpc:"refused,omptempty"`
	SourceId        *Many2One  `xmlrpc:"source_id,omptempty"`
	StageId         *Many2One  `xmlrpc:"stage_id,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
	UserId          *Many2One  `xmlrpc:"user_id,omptempty"`
}

// HrRecruitmentReports represents array of hr.recruitment.report model.
type HrRecruitmentReports []HrRecruitmentReport

// HrRecruitmentReportModel is the odoo model name.
const HrRecruitmentReportModel = "hr.recruitment.report"

// Many2One convert HrRecruitmentReport to *Many2One.
func (hrr *HrRecruitmentReport) Many2One() *Many2One {
	return NewMany2One(hrr.Id.Get(), "")
}

// CreateHrRecruitmentReport creates a new hr.recruitment.report model and returns its id.
func (c *Client) CreateHrRecruitmentReport(hrr *HrRecruitmentReport) (int64, error) {
	ids, err := c.CreateHrRecruitmentReports([]*HrRecruitmentReport{hrr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrRecruitmentReport creates a new hr.recruitment.report model and returns its id.
func (c *Client) CreateHrRecruitmentReports(hrrs []*HrRecruitmentReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrrs {
		vv = append(vv, v)
	}
	return c.Create(HrRecruitmentReportModel, vv)
}

// UpdateHrRecruitmentReport updates an existing hr.recruitment.report record.
func (c *Client) UpdateHrRecruitmentReport(hrr *HrRecruitmentReport) error {
	return c.UpdateHrRecruitmentReports([]int64{hrr.Id.Get()}, hrr)
}

// UpdateHrRecruitmentReports updates existing hr.recruitment.report records.
// All records (represented by ids) will be updated by hrr values.
func (c *Client) UpdateHrRecruitmentReports(ids []int64, hrr *HrRecruitmentReport) error {
	return c.Update(HrRecruitmentReportModel, ids, hrr)
}

// DeleteHrRecruitmentReport deletes an existing hr.recruitment.report record.
func (c *Client) DeleteHrRecruitmentReport(id int64) error {
	return c.DeleteHrRecruitmentReports([]int64{id})
}

// DeleteHrRecruitmentReports deletes existing hr.recruitment.report records.
func (c *Client) DeleteHrRecruitmentReports(ids []int64) error {
	return c.Delete(HrRecruitmentReportModel, ids)
}

// GetHrRecruitmentReport gets hr.recruitment.report existing record.
func (c *Client) GetHrRecruitmentReport(id int64) (*HrRecruitmentReport, error) {
	hrrs, err := c.GetHrRecruitmentReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrrs != nil && len(*hrrs) > 0 {
		return &((*hrrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.recruitment.report not found", id)
}

// GetHrRecruitmentReports gets hr.recruitment.report existing records.
func (c *Client) GetHrRecruitmentReports(ids []int64) (*HrRecruitmentReports, error) {
	hrrs := &HrRecruitmentReports{}
	if err := c.Read(HrRecruitmentReportModel, ids, nil, hrrs); err != nil {
		return nil, err
	}
	return hrrs, nil
}

// FindHrRecruitmentReport finds hr.recruitment.report record by querying it with criteria.
func (c *Client) FindHrRecruitmentReport(criteria *Criteria) (*HrRecruitmentReport, error) {
	hrrs := &HrRecruitmentReports{}
	if err := c.SearchRead(HrRecruitmentReportModel, criteria, NewOptions().Limit(1), hrrs); err != nil {
		return nil, err
	}
	if hrrs != nil && len(*hrrs) > 0 {
		return &((*hrrs)[0]), nil
	}
	return nil, fmt.Errorf("hr.recruitment.report was not found with criteria %v", criteria)
}

// FindHrRecruitmentReports finds hr.recruitment.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentReports(criteria *Criteria, options *Options) (*HrRecruitmentReports, error) {
	hrrs := &HrRecruitmentReports{}
	if err := c.SearchRead(HrRecruitmentReportModel, criteria, options, hrrs); err != nil {
		return nil, err
	}
	return hrrs, nil
}

// FindHrRecruitmentReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRecruitmentReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrRecruitmentReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrRecruitmentReportId finds record id by querying it with criteria.
func (c *Client) FindHrRecruitmentReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRecruitmentReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.recruitment.report was not found with criteria %v and options %v", criteria, options)
}
