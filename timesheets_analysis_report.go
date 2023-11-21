package odoo

import (
	"fmt"
)

// TimesheetsAnalysisReport represents timesheets.analysis.report model.
type TimesheetsAnalysisReport struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	Amount               *Float     `xmlrpc:"amount,omptempty"`
	AncestorTaskId       *Many2One  `xmlrpc:"ancestor_task_id,omptempty"`
	BillableTime         *Float     `xmlrpc:"billable_time,omptempty"`
	CompanyId            *Many2One  `xmlrpc:"company_id,omptempty"`
	CurrencyId           *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                 *Time      `xmlrpc:"date,omptempty"`
	DepartmentId         *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId           *Many2One  `xmlrpc:"employee_id,omptempty"`
	HelpdeskTicketId     *Many2One  `xmlrpc:"helpdesk_ticket_id,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	IsTimerRunning       *Bool      `xmlrpc:"is_timer_running,omptempty"`
	IsTimesheet          *Bool      `xmlrpc:"is_timesheet,omptempty"`
	ManagerId            *Many2One  `xmlrpc:"manager_id,omptempty"`
	Margin               *Float     `xmlrpc:"margin,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	NonBillableTime      *Float     `xmlrpc:"non_billable_time,omptempty"`
	OrderId              *Many2One  `xmlrpc:"order_id,omptempty"`
	ProjectId            *Many2One  `xmlrpc:"project_id,omptempty"`
	SoLine               *Many2One  `xmlrpc:"so_line,omptempty"`
	TaskId               *Many2One  `xmlrpc:"task_id,omptempty"`
	TimesheetInvoiceId   *Many2One  `xmlrpc:"timesheet_invoice_id,omptempty"`
	TimesheetInvoiceType *Selection `xmlrpc:"timesheet_invoice_type,omptempty"`
	TimesheetRevenues    *Float     `xmlrpc:"timesheet_revenues,omptempty"`
	UnitAmount           *Float     `xmlrpc:"unit_amount,omptempty"`
	UserId               *Many2One  `xmlrpc:"user_id,omptempty"`
	Validated            *Bool      `xmlrpc:"validated,omptempty"`
}

// TimesheetsAnalysisReports represents array of timesheets.analysis.report model.
type TimesheetsAnalysisReports []TimesheetsAnalysisReport

// TimesheetsAnalysisReportModel is the odoo model name.
const TimesheetsAnalysisReportModel = "timesheets.analysis.report"

// Many2One convert TimesheetsAnalysisReport to *Many2One.
func (tar *TimesheetsAnalysisReport) Many2One() *Many2One {
	return NewMany2One(tar.Id.Get(), "")
}

// CreateTimesheetsAnalysisReport creates a new timesheets.analysis.report model and returns its id.
func (c *Client) CreateTimesheetsAnalysisReport(tar *TimesheetsAnalysisReport) (int64, error) {
	ids, err := c.CreateTimesheetsAnalysisReports([]*TimesheetsAnalysisReport{tar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimesheetsAnalysisReport creates a new timesheets.analysis.report model and returns its id.
func (c *Client) CreateTimesheetsAnalysisReports(tars []*TimesheetsAnalysisReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range tars {
		vv = append(vv, v)
	}
	return c.Create(TimesheetsAnalysisReportModel, vv)
}

// UpdateTimesheetsAnalysisReport updates an existing timesheets.analysis.report record.
func (c *Client) UpdateTimesheetsAnalysisReport(tar *TimesheetsAnalysisReport) error {
	return c.UpdateTimesheetsAnalysisReports([]int64{tar.Id.Get()}, tar)
}

// UpdateTimesheetsAnalysisReports updates existing timesheets.analysis.report records.
// All records (represented by ids) will be updated by tar values.
func (c *Client) UpdateTimesheetsAnalysisReports(ids []int64, tar *TimesheetsAnalysisReport) error {
	return c.Update(TimesheetsAnalysisReportModel, ids, tar)
}

// DeleteTimesheetsAnalysisReport deletes an existing timesheets.analysis.report record.
func (c *Client) DeleteTimesheetsAnalysisReport(id int64) error {
	return c.DeleteTimesheetsAnalysisReports([]int64{id})
}

// DeleteTimesheetsAnalysisReports deletes existing timesheets.analysis.report records.
func (c *Client) DeleteTimesheetsAnalysisReports(ids []int64) error {
	return c.Delete(TimesheetsAnalysisReportModel, ids)
}

// GetTimesheetsAnalysisReport gets timesheets.analysis.report existing record.
func (c *Client) GetTimesheetsAnalysisReport(id int64) (*TimesheetsAnalysisReport, error) {
	tars, err := c.GetTimesheetsAnalysisReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if tars != nil && len(*tars) > 0 {
		return &((*tars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of timesheets.analysis.report not found", id)
}

// GetTimesheetsAnalysisReports gets timesheets.analysis.report existing records.
func (c *Client) GetTimesheetsAnalysisReports(ids []int64) (*TimesheetsAnalysisReports, error) {
	tars := &TimesheetsAnalysisReports{}
	if err := c.Read(TimesheetsAnalysisReportModel, ids, nil, tars); err != nil {
		return nil, err
	}
	return tars, nil
}

// FindTimesheetsAnalysisReport finds timesheets.analysis.report record by querying it with criteria.
func (c *Client) FindTimesheetsAnalysisReport(criteria *Criteria) (*TimesheetsAnalysisReport, error) {
	tars := &TimesheetsAnalysisReports{}
	if err := c.SearchRead(TimesheetsAnalysisReportModel, criteria, NewOptions().Limit(1), tars); err != nil {
		return nil, err
	}
	if tars != nil && len(*tars) > 0 {
		return &((*tars)[0]), nil
	}
	return nil, fmt.Errorf("timesheets.analysis.report was not found with criteria %v", criteria)
}

// FindTimesheetsAnalysisReports finds timesheets.analysis.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetsAnalysisReports(criteria *Criteria, options *Options) (*TimesheetsAnalysisReports, error) {
	tars := &TimesheetsAnalysisReports{}
	if err := c.SearchRead(TimesheetsAnalysisReportModel, criteria, options, tars); err != nil {
		return nil, err
	}
	return tars, nil
}

// FindTimesheetsAnalysisReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimesheetsAnalysisReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(TimesheetsAnalysisReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTimesheetsAnalysisReportId finds record id by querying it with criteria.
func (c *Client) FindTimesheetsAnalysisReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimesheetsAnalysisReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("timesheets.analysis.report was not found with criteria %v and options %v", criteria, options)
}
