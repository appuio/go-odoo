package odoo

import (
	"fmt"
)

// HelpdeskTicketReportAnalysis represents helpdesk.ticket.report.analysis model.
type HelpdeskTicketReportAnalysis struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	AssignDate             *Time      `xmlrpc:"assign_date,omptempty"`
	AvgResponseHours       *Float     `xmlrpc:"avg_response_hours,omptempty"`
	CloseDate              *Time      `xmlrpc:"close_date,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	DepartmentId           *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId             *Many2One  `xmlrpc:"employee_id,omptempty"`
	EmployeeParentId       *Many2One  `xmlrpc:"employee_parent_id,omptempty"`
	FirstResponseHours     *Float     `xmlrpc:"first_response_hours,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	KanbanState            *Selection `xmlrpc:"kanban_state,omptempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omptempty"`
	Priority               *Selection `xmlrpc:"priority,omptempty"`
	RatingLastValue        *Float     `xmlrpc:"rating_last_value,omptempty"`
	SaleOrderId            *Many2One  `xmlrpc:"sale_order_id,omptempty"`
	SlaFail                *Bool      `xmlrpc:"sla_fail,omptempty"`
	TeamId                 *Many2One  `xmlrpc:"team_id,omptempty"`
	TicketAssignationHours *Float     `xmlrpc:"ticket_assignation_hours,omptempty"`
	TicketCloseHours       *Float     `xmlrpc:"ticket_close_hours,omptempty"`
	TicketDeadline         *Time      `xmlrpc:"ticket_deadline,omptempty"`
	TicketDeadlineHours    *Float     `xmlrpc:"ticket_deadline_hours,omptempty"`
	TicketId               *Many2One  `xmlrpc:"ticket_id,omptempty"`
	TicketOpenHours        *Float     `xmlrpc:"ticket_open_hours,omptempty"`
	TicketStageId          *Many2One  `xmlrpc:"ticket_stage_id,omptempty"`
	TicketTypeId           *Many2One  `xmlrpc:"ticket_type_id,omptempty"`
	TotalHoursSpent        *Float     `xmlrpc:"total_hours_spent,omptempty"`
	UserId                 *Many2One  `xmlrpc:"user_id,omptempty"`
}

// HelpdeskTicketReportAnalysiss represents array of helpdesk.ticket.report.analysis model.
type HelpdeskTicketReportAnalysiss []HelpdeskTicketReportAnalysis

// HelpdeskTicketReportAnalysisModel is the odoo model name.
const HelpdeskTicketReportAnalysisModel = "helpdesk.ticket.report.analysis"

// Many2One convert HelpdeskTicketReportAnalysis to *Many2One.
func (htra *HelpdeskTicketReportAnalysis) Many2One() *Many2One {
	return NewMany2One(htra.Id.Get(), "")
}

// CreateHelpdeskTicketReportAnalysis creates a new helpdesk.ticket.report.analysis model and returns its id.
func (c *Client) CreateHelpdeskTicketReportAnalysis(htra *HelpdeskTicketReportAnalysis) (int64, error) {
	ids, err := c.CreateHelpdeskTicketReportAnalysiss([]*HelpdeskTicketReportAnalysis{htra})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicketReportAnalysis creates a new helpdesk.ticket.report.analysis model and returns its id.
func (c *Client) CreateHelpdeskTicketReportAnalysiss(htras []*HelpdeskTicketReportAnalysis) ([]int64, error) {
	var vv []interface{}
	for _, v := range htras {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketReportAnalysisModel, vv)
}

// UpdateHelpdeskTicketReportAnalysis updates an existing helpdesk.ticket.report.analysis record.
func (c *Client) UpdateHelpdeskTicketReportAnalysis(htra *HelpdeskTicketReportAnalysis) error {
	return c.UpdateHelpdeskTicketReportAnalysiss([]int64{htra.Id.Get()}, htra)
}

// UpdateHelpdeskTicketReportAnalysiss updates existing helpdesk.ticket.report.analysis records.
// All records (represented by ids) will be updated by htra values.
func (c *Client) UpdateHelpdeskTicketReportAnalysiss(ids []int64, htra *HelpdeskTicketReportAnalysis) error {
	return c.Update(HelpdeskTicketReportAnalysisModel, ids, htra)
}

// DeleteHelpdeskTicketReportAnalysis deletes an existing helpdesk.ticket.report.analysis record.
func (c *Client) DeleteHelpdeskTicketReportAnalysis(id int64) error {
	return c.DeleteHelpdeskTicketReportAnalysiss([]int64{id})
}

// DeleteHelpdeskTicketReportAnalysiss deletes existing helpdesk.ticket.report.analysis records.
func (c *Client) DeleteHelpdeskTicketReportAnalysiss(ids []int64) error {
	return c.Delete(HelpdeskTicketReportAnalysisModel, ids)
}

// GetHelpdeskTicketReportAnalysis gets helpdesk.ticket.report.analysis existing record.
func (c *Client) GetHelpdeskTicketReportAnalysis(id int64) (*HelpdeskTicketReportAnalysis, error) {
	htras, err := c.GetHelpdeskTicketReportAnalysiss([]int64{id})
	if err != nil {
		return nil, err
	}
	if htras != nil && len(*htras) > 0 {
		return &((*htras)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.ticket.report.analysis not found", id)
}

// GetHelpdeskTicketReportAnalysiss gets helpdesk.ticket.report.analysis existing records.
func (c *Client) GetHelpdeskTicketReportAnalysiss(ids []int64) (*HelpdeskTicketReportAnalysiss, error) {
	htras := &HelpdeskTicketReportAnalysiss{}
	if err := c.Read(HelpdeskTicketReportAnalysisModel, ids, nil, htras); err != nil {
		return nil, err
	}
	return htras, nil
}

// FindHelpdeskTicketReportAnalysis finds helpdesk.ticket.report.analysis record by querying it with criteria.
func (c *Client) FindHelpdeskTicketReportAnalysis(criteria *Criteria) (*HelpdeskTicketReportAnalysis, error) {
	htras := &HelpdeskTicketReportAnalysiss{}
	if err := c.SearchRead(HelpdeskTicketReportAnalysisModel, criteria, NewOptions().Limit(1), htras); err != nil {
		return nil, err
	}
	if htras != nil && len(*htras) > 0 {
		return &((*htras)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.ticket.report.analysis was not found with criteria %v", criteria)
}

// FindHelpdeskTicketReportAnalysiss finds helpdesk.ticket.report.analysis records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketReportAnalysiss(criteria *Criteria, options *Options) (*HelpdeskTicketReportAnalysiss, error) {
	htras := &HelpdeskTicketReportAnalysiss{}
	if err := c.SearchRead(HelpdeskTicketReportAnalysisModel, criteria, options, htras); err != nil {
		return nil, err
	}
	return htras, nil
}

// FindHelpdeskTicketReportAnalysisIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketReportAnalysisIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTicketReportAnalysisModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTicketReportAnalysisId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketReportAnalysisId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketReportAnalysisModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.ticket.report.analysis was not found with criteria %v and options %v", criteria, options)
}
