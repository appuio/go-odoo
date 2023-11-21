package odoo

import (
	"fmt"
)

// ProjectTimesheetForecastReportAnalysis represents project.timesheet.forecast.report.analysis model.
type ProjectTimesheetForecastReportAnalysis struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	Difference     *Float     `xmlrpc:"difference,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	EffectiveHours *Float     `xmlrpc:"effective_hours,omptempty"`
	EmployeeId     *Many2One  `xmlrpc:"employee_id,omptempty"`
	EntryDate      *Time      `xmlrpc:"entry_date,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	IsPublished    *Bool      `xmlrpc:"is_published,omptempty"`
	LineType       *Selection `xmlrpc:"line_type,omptempty"`
	PlannedHours   *Float     `xmlrpc:"planned_hours,omptempty"`
	ProjectId      *Many2One  `xmlrpc:"project_id,omptempty"`
	UserId         *Many2One  `xmlrpc:"user_id,omptempty"`
}

// ProjectTimesheetForecastReportAnalysiss represents array of project.timesheet.forecast.report.analysis model.
type ProjectTimesheetForecastReportAnalysiss []ProjectTimesheetForecastReportAnalysis

// ProjectTimesheetForecastReportAnalysisModel is the odoo model name.
const ProjectTimesheetForecastReportAnalysisModel = "project.timesheet.forecast.report.analysis"

// Many2One convert ProjectTimesheetForecastReportAnalysis to *Many2One.
func (ptfra *ProjectTimesheetForecastReportAnalysis) Many2One() *Many2One {
	return NewMany2One(ptfra.Id.Get(), "")
}

// CreateProjectTimesheetForecastReportAnalysis creates a new project.timesheet.forecast.report.analysis model and returns its id.
func (c *Client) CreateProjectTimesheetForecastReportAnalysis(ptfra *ProjectTimesheetForecastReportAnalysis) (int64, error) {
	ids, err := c.CreateProjectTimesheetForecastReportAnalysiss([]*ProjectTimesheetForecastReportAnalysis{ptfra})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTimesheetForecastReportAnalysis creates a new project.timesheet.forecast.report.analysis model and returns its id.
func (c *Client) CreateProjectTimesheetForecastReportAnalysiss(ptfras []*ProjectTimesheetForecastReportAnalysis) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptfras {
		vv = append(vv, v)
	}
	return c.Create(ProjectTimesheetForecastReportAnalysisModel, vv)
}

// UpdateProjectTimesheetForecastReportAnalysis updates an existing project.timesheet.forecast.report.analysis record.
func (c *Client) UpdateProjectTimesheetForecastReportAnalysis(ptfra *ProjectTimesheetForecastReportAnalysis) error {
	return c.UpdateProjectTimesheetForecastReportAnalysiss([]int64{ptfra.Id.Get()}, ptfra)
}

// UpdateProjectTimesheetForecastReportAnalysiss updates existing project.timesheet.forecast.report.analysis records.
// All records (represented by ids) will be updated by ptfra values.
func (c *Client) UpdateProjectTimesheetForecastReportAnalysiss(ids []int64, ptfra *ProjectTimesheetForecastReportAnalysis) error {
	return c.Update(ProjectTimesheetForecastReportAnalysisModel, ids, ptfra)
}

// DeleteProjectTimesheetForecastReportAnalysis deletes an existing project.timesheet.forecast.report.analysis record.
func (c *Client) DeleteProjectTimesheetForecastReportAnalysis(id int64) error {
	return c.DeleteProjectTimesheetForecastReportAnalysiss([]int64{id})
}

// DeleteProjectTimesheetForecastReportAnalysiss deletes existing project.timesheet.forecast.report.analysis records.
func (c *Client) DeleteProjectTimesheetForecastReportAnalysiss(ids []int64) error {
	return c.Delete(ProjectTimesheetForecastReportAnalysisModel, ids)
}

// GetProjectTimesheetForecastReportAnalysis gets project.timesheet.forecast.report.analysis existing record.
func (c *Client) GetProjectTimesheetForecastReportAnalysis(id int64) (*ProjectTimesheetForecastReportAnalysis, error) {
	ptfras, err := c.GetProjectTimesheetForecastReportAnalysiss([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptfras != nil && len(*ptfras) > 0 {
		return &((*ptfras)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.timesheet.forecast.report.analysis not found", id)
}

// GetProjectTimesheetForecastReportAnalysiss gets project.timesheet.forecast.report.analysis existing records.
func (c *Client) GetProjectTimesheetForecastReportAnalysiss(ids []int64) (*ProjectTimesheetForecastReportAnalysiss, error) {
	ptfras := &ProjectTimesheetForecastReportAnalysiss{}
	if err := c.Read(ProjectTimesheetForecastReportAnalysisModel, ids, nil, ptfras); err != nil {
		return nil, err
	}
	return ptfras, nil
}

// FindProjectTimesheetForecastReportAnalysis finds project.timesheet.forecast.report.analysis record by querying it with criteria.
func (c *Client) FindProjectTimesheetForecastReportAnalysis(criteria *Criteria) (*ProjectTimesheetForecastReportAnalysis, error) {
	ptfras := &ProjectTimesheetForecastReportAnalysiss{}
	if err := c.SearchRead(ProjectTimesheetForecastReportAnalysisModel, criteria, NewOptions().Limit(1), ptfras); err != nil {
		return nil, err
	}
	if ptfras != nil && len(*ptfras) > 0 {
		return &((*ptfras)[0]), nil
	}
	return nil, fmt.Errorf("project.timesheet.forecast.report.analysis was not found with criteria %v", criteria)
}

// FindProjectTimesheetForecastReportAnalysiss finds project.timesheet.forecast.report.analysis records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTimesheetForecastReportAnalysiss(criteria *Criteria, options *Options) (*ProjectTimesheetForecastReportAnalysiss, error) {
	ptfras := &ProjectTimesheetForecastReportAnalysiss{}
	if err := c.SearchRead(ProjectTimesheetForecastReportAnalysisModel, criteria, options, ptfras); err != nil {
		return nil, err
	}
	return ptfras, nil
}

// FindProjectTimesheetForecastReportAnalysisIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTimesheetForecastReportAnalysisIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTimesheetForecastReportAnalysisModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTimesheetForecastReportAnalysisId finds record id by querying it with criteria.
func (c *Client) FindProjectTimesheetForecastReportAnalysisId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTimesheetForecastReportAnalysisModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.timesheet.forecast.report.analysis was not found with criteria %v and options %v", criteria, options)
}
