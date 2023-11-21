package odoo

import (
	"fmt"
)

// HrAppraisalSkillReport represents hr.appraisal.skill.report model.
type HrAppraisalSkillReport struct {
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CurrentLevelProgress  *Float     `xmlrpc:"current_level_progress,omptempty"`
	CurrentSkillLevelId   *Many2One  `xmlrpc:"current_skill_level_id,omptempty"`
	DepartmentId          *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId            *Many2One  `xmlrpc:"employee_id,omptempty"`
	Evolution             *Selection `xmlrpc:"evolution,omptempty"`
	EvolutionSequence     *Int       `xmlrpc:"evolution_sequence,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	Justification         *String    `xmlrpc:"justification,omptempty"`
	PreviousLevelProgress *Float     `xmlrpc:"previous_level_progress,omptempty"`
	PreviousSkillLevelId  *Many2One  `xmlrpc:"previous_skill_level_id,omptempty"`
	SkillId               *Many2One  `xmlrpc:"skill_id,omptempty"`
	SkillTypeId           *Many2One  `xmlrpc:"skill_type_id,omptempty"`
}

// HrAppraisalSkillReports represents array of hr.appraisal.skill.report model.
type HrAppraisalSkillReports []HrAppraisalSkillReport

// HrAppraisalSkillReportModel is the odoo model name.
const HrAppraisalSkillReportModel = "hr.appraisal.skill.report"

// Many2One convert HrAppraisalSkillReport to *Many2One.
func (hasr *HrAppraisalSkillReport) Many2One() *Many2One {
	return NewMany2One(hasr.Id.Get(), "")
}

// CreateHrAppraisalSkillReport creates a new hr.appraisal.skill.report model and returns its id.
func (c *Client) CreateHrAppraisalSkillReport(hasr *HrAppraisalSkillReport) (int64, error) {
	ids, err := c.CreateHrAppraisalSkillReports([]*HrAppraisalSkillReport{hasr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAppraisalSkillReport creates a new hr.appraisal.skill.report model and returns its id.
func (c *Client) CreateHrAppraisalSkillReports(hasrs []*HrAppraisalSkillReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range hasrs {
		vv = append(vv, v)
	}
	return c.Create(HrAppraisalSkillReportModel, vv)
}

// UpdateHrAppraisalSkillReport updates an existing hr.appraisal.skill.report record.
func (c *Client) UpdateHrAppraisalSkillReport(hasr *HrAppraisalSkillReport) error {
	return c.UpdateHrAppraisalSkillReports([]int64{hasr.Id.Get()}, hasr)
}

// UpdateHrAppraisalSkillReports updates existing hr.appraisal.skill.report records.
// All records (represented by ids) will be updated by hasr values.
func (c *Client) UpdateHrAppraisalSkillReports(ids []int64, hasr *HrAppraisalSkillReport) error {
	return c.Update(HrAppraisalSkillReportModel, ids, hasr)
}

// DeleteHrAppraisalSkillReport deletes an existing hr.appraisal.skill.report record.
func (c *Client) DeleteHrAppraisalSkillReport(id int64) error {
	return c.DeleteHrAppraisalSkillReports([]int64{id})
}

// DeleteHrAppraisalSkillReports deletes existing hr.appraisal.skill.report records.
func (c *Client) DeleteHrAppraisalSkillReports(ids []int64) error {
	return c.Delete(HrAppraisalSkillReportModel, ids)
}

// GetHrAppraisalSkillReport gets hr.appraisal.skill.report existing record.
func (c *Client) GetHrAppraisalSkillReport(id int64) (*HrAppraisalSkillReport, error) {
	hasrs, err := c.GetHrAppraisalSkillReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if hasrs != nil && len(*hasrs) > 0 {
		return &((*hasrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.appraisal.skill.report not found", id)
}

// GetHrAppraisalSkillReports gets hr.appraisal.skill.report existing records.
func (c *Client) GetHrAppraisalSkillReports(ids []int64) (*HrAppraisalSkillReports, error) {
	hasrs := &HrAppraisalSkillReports{}
	if err := c.Read(HrAppraisalSkillReportModel, ids, nil, hasrs); err != nil {
		return nil, err
	}
	return hasrs, nil
}

// FindHrAppraisalSkillReport finds hr.appraisal.skill.report record by querying it with criteria.
func (c *Client) FindHrAppraisalSkillReport(criteria *Criteria) (*HrAppraisalSkillReport, error) {
	hasrs := &HrAppraisalSkillReports{}
	if err := c.SearchRead(HrAppraisalSkillReportModel, criteria, NewOptions().Limit(1), hasrs); err != nil {
		return nil, err
	}
	if hasrs != nil && len(*hasrs) > 0 {
		return &((*hasrs)[0]), nil
	}
	return nil, fmt.Errorf("hr.appraisal.skill.report was not found with criteria %v", criteria)
}

// FindHrAppraisalSkillReports finds hr.appraisal.skill.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalSkillReports(criteria *Criteria, options *Options) (*HrAppraisalSkillReports, error) {
	hasrs := &HrAppraisalSkillReports{}
	if err := c.SearchRead(HrAppraisalSkillReportModel, criteria, options, hasrs); err != nil {
		return nil, err
	}
	return hasrs, nil
}

// FindHrAppraisalSkillReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalSkillReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAppraisalSkillReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAppraisalSkillReportId finds record id by querying it with criteria.
func (c *Client) FindHrAppraisalSkillReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAppraisalSkillReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.appraisal.skill.report was not found with criteria %v and options %v", criteria, options)
}
