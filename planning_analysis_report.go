package odoo

import (
	"fmt"
)

// PlanningAnalysisReport represents planning.analysis.report model.
type PlanningAnalysisReport struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	AllocatedHours            *Float     `xmlrpc:"allocated_hours,omptempty"`
	AllocatedHoursCost        *Float     `xmlrpc:"allocated_hours_cost,omptempty"`
	AllocatedPercentage       *Float     `xmlrpc:"allocated_percentage,omptempty"`
	BillableAllocatedHours    *Float     `xmlrpc:"billable_allocated_hours,omptempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omptempty"`
	DepartmentId              *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	EffectiveHours            *Float     `xmlrpc:"effective_hours,omptempty"`
	EffectiveHoursCost        *Float     `xmlrpc:"effective_hours_cost,omptempty"`
	EmployeeId                *Many2One  `xmlrpc:"employee_id,omptempty"`
	EmployeeSkillIds          *Relation  `xmlrpc:"employee_skill_ids,omptempty"`
	EndDatetime               *Time      `xmlrpc:"end_datetime,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	JobTitle                  *String    `xmlrpc:"job_title,omptempty"`
	ManagerId                 *Many2One  `xmlrpc:"manager_id,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	NonBillableAllocatedHours *Float     `xmlrpc:"non_billable_allocated_hours,omptempty"`
	PercentageHours           *Float     `xmlrpc:"percentage_hours,omptempty"`
	ProjectId                 *Many2One  `xmlrpc:"project_id,omptempty"`
	PublicationWarning        *Bool      `xmlrpc:"publication_warning,omptempty"`
	RecurrencyId              *Many2One  `xmlrpc:"recurrency_id,omptempty"`
	RemainingHours            *Float     `xmlrpc:"remaining_hours,omptempty"`
	ResourceId                *Many2One  `xmlrpc:"resource_id,omptempty"`
	ResourceType              *Selection `xmlrpc:"resource_type,omptempty"`
	RoleId                    *Many2One  `xmlrpc:"role_id,omptempty"`
	RoleProductIds            *Relation  `xmlrpc:"role_product_ids,omptempty"`
	SaleLineId                *Many2One  `xmlrpc:"sale_line_id,omptempty"`
	SaleOrderId               *Many2One  `xmlrpc:"sale_order_id,omptempty"`
	StartDatetime             *Time      `xmlrpc:"start_datetime,omptempty"`
	State                     *Selection `xmlrpc:"state,omptempty"`
	UserId                    *Many2One  `xmlrpc:"user_id,omptempty"`
	WorkingDaysCount          *Float     `xmlrpc:"working_days_count,omptempty"`
}

// PlanningAnalysisReports represents array of planning.analysis.report model.
type PlanningAnalysisReports []PlanningAnalysisReport

// PlanningAnalysisReportModel is the odoo model name.
const PlanningAnalysisReportModel = "planning.analysis.report"

// Many2One convert PlanningAnalysisReport to *Many2One.
func (par *PlanningAnalysisReport) Many2One() *Many2One {
	return NewMany2One(par.Id.Get(), "")
}

// CreatePlanningAnalysisReport creates a new planning.analysis.report model and returns its id.
func (c *Client) CreatePlanningAnalysisReport(par *PlanningAnalysisReport) (int64, error) {
	ids, err := c.CreatePlanningAnalysisReports([]*PlanningAnalysisReport{par})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningAnalysisReport creates a new planning.analysis.report model and returns its id.
func (c *Client) CreatePlanningAnalysisReports(pars []*PlanningAnalysisReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range pars {
		vv = append(vv, v)
	}
	return c.Create(PlanningAnalysisReportModel, vv)
}

// UpdatePlanningAnalysisReport updates an existing planning.analysis.report record.
func (c *Client) UpdatePlanningAnalysisReport(par *PlanningAnalysisReport) error {
	return c.UpdatePlanningAnalysisReports([]int64{par.Id.Get()}, par)
}

// UpdatePlanningAnalysisReports updates existing planning.analysis.report records.
// All records (represented by ids) will be updated by par values.
func (c *Client) UpdatePlanningAnalysisReports(ids []int64, par *PlanningAnalysisReport) error {
	return c.Update(PlanningAnalysisReportModel, ids, par)
}

// DeletePlanningAnalysisReport deletes an existing planning.analysis.report record.
func (c *Client) DeletePlanningAnalysisReport(id int64) error {
	return c.DeletePlanningAnalysisReports([]int64{id})
}

// DeletePlanningAnalysisReports deletes existing planning.analysis.report records.
func (c *Client) DeletePlanningAnalysisReports(ids []int64) error {
	return c.Delete(PlanningAnalysisReportModel, ids)
}

// GetPlanningAnalysisReport gets planning.analysis.report existing record.
func (c *Client) GetPlanningAnalysisReport(id int64) (*PlanningAnalysisReport, error) {
	pars, err := c.GetPlanningAnalysisReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if pars != nil && len(*pars) > 0 {
		return &((*pars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of planning.analysis.report not found", id)
}

// GetPlanningAnalysisReports gets planning.analysis.report existing records.
func (c *Client) GetPlanningAnalysisReports(ids []int64) (*PlanningAnalysisReports, error) {
	pars := &PlanningAnalysisReports{}
	if err := c.Read(PlanningAnalysisReportModel, ids, nil, pars); err != nil {
		return nil, err
	}
	return pars, nil
}

// FindPlanningAnalysisReport finds planning.analysis.report record by querying it with criteria.
func (c *Client) FindPlanningAnalysisReport(criteria *Criteria) (*PlanningAnalysisReport, error) {
	pars := &PlanningAnalysisReports{}
	if err := c.SearchRead(PlanningAnalysisReportModel, criteria, NewOptions().Limit(1), pars); err != nil {
		return nil, err
	}
	if pars != nil && len(*pars) > 0 {
		return &((*pars)[0]), nil
	}
	return nil, fmt.Errorf("planning.analysis.report was not found with criteria %v", criteria)
}

// FindPlanningAnalysisReports finds planning.analysis.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningAnalysisReports(criteria *Criteria, options *Options) (*PlanningAnalysisReports, error) {
	pars := &PlanningAnalysisReports{}
	if err := c.SearchRead(PlanningAnalysisReportModel, criteria, options, pars); err != nil {
		return nil, err
	}
	return pars, nil
}

// FindPlanningAnalysisReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningAnalysisReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PlanningAnalysisReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPlanningAnalysisReportId finds record id by querying it with criteria.
func (c *Client) FindPlanningAnalysisReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningAnalysisReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("planning.analysis.report was not found with criteria %v and options %v", criteria, options)
}
