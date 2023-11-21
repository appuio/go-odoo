package odoo

import (
	"fmt"
)

// ProjectTaskBurndownChartReport represents project.task.burndown.chart.report model.
type ProjectTaskBurndownChartReport struct {
	Date                         *Time     `xmlrpc:"date,omptempty"`
	DateAssign                   *Time     `xmlrpc:"date_assign,omptempty"`
	DateDeadline                 *Time     `xmlrpc:"date_deadline,omptempty"`
	DisplayProjectId             *Many2One `xmlrpc:"display_project_id,omptempty"`
	HasLateAndUnreachedMilestone *Bool     `xmlrpc:"has_late_and_unreached_milestone,omptempty"`
	IsClosed                     *Bool     `xmlrpc:"is_closed,omptempty"`
	MilestoneId                  *Many2One `xmlrpc:"milestone_id,omptempty"`
	PartnerId                    *Many2One `xmlrpc:"partner_id,omptempty"`
	PlannedHours                 *Float    `xmlrpc:"planned_hours,omptempty"`
	ProjectId                    *Many2One `xmlrpc:"project_id,omptempty"`
	StageId                      *Many2One `xmlrpc:"stage_id,omptempty"`
	UserIds                      *Relation `xmlrpc:"user_ids,omptempty"`
}

// ProjectTaskBurndownChartReports represents array of project.task.burndown.chart.report model.
type ProjectTaskBurndownChartReports []ProjectTaskBurndownChartReport

// ProjectTaskBurndownChartReportModel is the odoo model name.
const ProjectTaskBurndownChartReportModel = "project.task.burndown.chart.report"

// Many2One convert ProjectTaskBurndownChartReport to *Many2One.
func (ptbcr *ProjectTaskBurndownChartReport) Many2One() *Many2One {
	return NewMany2One(ptbcr.Id.Get(), "")
}

// CreateProjectTaskBurndownChartReport creates a new project.task.burndown.chart.report model and returns its id.
func (c *Client) CreateProjectTaskBurndownChartReport(ptbcr *ProjectTaskBurndownChartReport) (int64, error) {
	ids, err := c.CreateProjectTaskBurndownChartReports([]*ProjectTaskBurndownChartReport{ptbcr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskBurndownChartReport creates a new project.task.burndown.chart.report model and returns its id.
func (c *Client) CreateProjectTaskBurndownChartReports(ptbcrs []*ProjectTaskBurndownChartReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptbcrs {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskBurndownChartReportModel, vv)
}

// UpdateProjectTaskBurndownChartReport updates an existing project.task.burndown.chart.report record.
func (c *Client) UpdateProjectTaskBurndownChartReport(ptbcr *ProjectTaskBurndownChartReport) error {
	return c.UpdateProjectTaskBurndownChartReports([]int64{ptbcr.Id.Get()}, ptbcr)
}

// UpdateProjectTaskBurndownChartReports updates existing project.task.burndown.chart.report records.
// All records (represented by ids) will be updated by ptbcr values.
func (c *Client) UpdateProjectTaskBurndownChartReports(ids []int64, ptbcr *ProjectTaskBurndownChartReport) error {
	return c.Update(ProjectTaskBurndownChartReportModel, ids, ptbcr)
}

// DeleteProjectTaskBurndownChartReport deletes an existing project.task.burndown.chart.report record.
func (c *Client) DeleteProjectTaskBurndownChartReport(id int64) error {
	return c.DeleteProjectTaskBurndownChartReports([]int64{id})
}

// DeleteProjectTaskBurndownChartReports deletes existing project.task.burndown.chart.report records.
func (c *Client) DeleteProjectTaskBurndownChartReports(ids []int64) error {
	return c.Delete(ProjectTaskBurndownChartReportModel, ids)
}

// GetProjectTaskBurndownChartReport gets project.task.burndown.chart.report existing record.
func (c *Client) GetProjectTaskBurndownChartReport(id int64) (*ProjectTaskBurndownChartReport, error) {
	ptbcrs, err := c.GetProjectTaskBurndownChartReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptbcrs != nil && len(*ptbcrs) > 0 {
		return &((*ptbcrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task.burndown.chart.report not found", id)
}

// GetProjectTaskBurndownChartReports gets project.task.burndown.chart.report existing records.
func (c *Client) GetProjectTaskBurndownChartReports(ids []int64) (*ProjectTaskBurndownChartReports, error) {
	ptbcrs := &ProjectTaskBurndownChartReports{}
	if err := c.Read(ProjectTaskBurndownChartReportModel, ids, nil, ptbcrs); err != nil {
		return nil, err
	}
	return ptbcrs, nil
}

// FindProjectTaskBurndownChartReport finds project.task.burndown.chart.report record by querying it with criteria.
func (c *Client) FindProjectTaskBurndownChartReport(criteria *Criteria) (*ProjectTaskBurndownChartReport, error) {
	ptbcrs := &ProjectTaskBurndownChartReports{}
	if err := c.SearchRead(ProjectTaskBurndownChartReportModel, criteria, NewOptions().Limit(1), ptbcrs); err != nil {
		return nil, err
	}
	if ptbcrs != nil && len(*ptbcrs) > 0 {
		return &((*ptbcrs)[0]), nil
	}
	return nil, fmt.Errorf("project.task.burndown.chart.report was not found with criteria %v", criteria)
}

// FindProjectTaskBurndownChartReports finds project.task.burndown.chart.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskBurndownChartReports(criteria *Criteria, options *Options) (*ProjectTaskBurndownChartReports, error) {
	ptbcrs := &ProjectTaskBurndownChartReports{}
	if err := c.SearchRead(ProjectTaskBurndownChartReportModel, criteria, options, ptbcrs); err != nil {
		return nil, err
	}
	return ptbcrs, nil
}

// FindProjectTaskBurndownChartReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskBurndownChartReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskBurndownChartReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskBurndownChartReportId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskBurndownChartReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskBurndownChartReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task.burndown.chart.report was not found with criteria %v and options %v", criteria, options)
}
