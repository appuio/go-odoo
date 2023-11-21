package odoo

import (
	"fmt"
)

// HelpdeskSlaReportAnalysis represents helpdesk.sla.report.analysis model.
type HelpdeskSlaReportAnalysis struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	CloseDate              *Time      `xmlrpc:"close_date,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	DepartmentId           *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId             *Many2One  `xmlrpc:"employee_id,omptempty"`
	FailedSlaRate          *Float     `xmlrpc:"failed_sla_rate,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	ManagerId              *Many2One  `xmlrpc:"manager_id,omptempty"`
	OngoingSlaRate         *Float     `xmlrpc:"ongoing_sla_rate,omptempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omptempty"`
	Priority               *Selection `xmlrpc:"priority,omptempty"`
	SaleOrderId            *Many2One  `xmlrpc:"sale_order_id,omptempty"`
	SlaDeadline            *Time      `xmlrpc:"sla_deadline,omptempty"`
	SlaExceededHours       *Int       `xmlrpc:"sla_exceeded_hours,omptempty"`
	SlaId                  *Many2One  `xmlrpc:"sla_id,omptempty"`
	SlaReachedDatetime     *Time      `xmlrpc:"sla_reached_datetime,omptempty"`
	SlaStageId             *Many2One  `xmlrpc:"sla_stage_id,omptempty"`
	SlaStatus              *Selection `xmlrpc:"sla_status,omptempty"`
	SlaStatusFail          *Bool      `xmlrpc:"sla_status_fail,omptempty"`
	SlaStatusFailed        *Int       `xmlrpc:"sla_status_failed,omptempty"`
	SlaStatusOngoing       *Int       `xmlrpc:"sla_status_ongoing,omptempty"`
	SlaStatusSuccessful    *Int       `xmlrpc:"sla_status_successful,omptempty"`
	SuccessfulSlaRate      *Float     `xmlrpc:"successful_sla_rate,omptempty"`
	TeamId                 *Many2One  `xmlrpc:"team_id,omptempty"`
	TicketAssignationHours *Int       `xmlrpc:"ticket_assignation_hours,omptempty"`
	TicketCloseHours       *Int       `xmlrpc:"ticket_close_hours,omptempty"`
	TicketClosed           *Bool      `xmlrpc:"ticket_closed,omptempty"`
	TicketDeadline         *Time      `xmlrpc:"ticket_deadline,omptempty"`
	TicketId               *Many2One  `xmlrpc:"ticket_id,omptempty"`
	TicketStageId          *Many2One  `xmlrpc:"ticket_stage_id,omptempty"`
	TicketTypeId           *Many2One  `xmlrpc:"ticket_type_id,omptempty"`
	UserId                 *Many2One  `xmlrpc:"user_id,omptempty"`
}

// HelpdeskSlaReportAnalysiss represents array of helpdesk.sla.report.analysis model.
type HelpdeskSlaReportAnalysiss []HelpdeskSlaReportAnalysis

// HelpdeskSlaReportAnalysisModel is the odoo model name.
const HelpdeskSlaReportAnalysisModel = "helpdesk.sla.report.analysis"

// Many2One convert HelpdeskSlaReportAnalysis to *Many2One.
func (hsra *HelpdeskSlaReportAnalysis) Many2One() *Many2One {
	return NewMany2One(hsra.Id.Get(), "")
}

// CreateHelpdeskSlaReportAnalysis creates a new helpdesk.sla.report.analysis model and returns its id.
func (c *Client) CreateHelpdeskSlaReportAnalysis(hsra *HelpdeskSlaReportAnalysis) (int64, error) {
	ids, err := c.CreateHelpdeskSlaReportAnalysiss([]*HelpdeskSlaReportAnalysis{hsra})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskSlaReportAnalysis creates a new helpdesk.sla.report.analysis model and returns its id.
func (c *Client) CreateHelpdeskSlaReportAnalysiss(hsras []*HelpdeskSlaReportAnalysis) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsras {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskSlaReportAnalysisModel, vv)
}

// UpdateHelpdeskSlaReportAnalysis updates an existing helpdesk.sla.report.analysis record.
func (c *Client) UpdateHelpdeskSlaReportAnalysis(hsra *HelpdeskSlaReportAnalysis) error {
	return c.UpdateHelpdeskSlaReportAnalysiss([]int64{hsra.Id.Get()}, hsra)
}

// UpdateHelpdeskSlaReportAnalysiss updates existing helpdesk.sla.report.analysis records.
// All records (represented by ids) will be updated by hsra values.
func (c *Client) UpdateHelpdeskSlaReportAnalysiss(ids []int64, hsra *HelpdeskSlaReportAnalysis) error {
	return c.Update(HelpdeskSlaReportAnalysisModel, ids, hsra)
}

// DeleteHelpdeskSlaReportAnalysis deletes an existing helpdesk.sla.report.analysis record.
func (c *Client) DeleteHelpdeskSlaReportAnalysis(id int64) error {
	return c.DeleteHelpdeskSlaReportAnalysiss([]int64{id})
}

// DeleteHelpdeskSlaReportAnalysiss deletes existing helpdesk.sla.report.analysis records.
func (c *Client) DeleteHelpdeskSlaReportAnalysiss(ids []int64) error {
	return c.Delete(HelpdeskSlaReportAnalysisModel, ids)
}

// GetHelpdeskSlaReportAnalysis gets helpdesk.sla.report.analysis existing record.
func (c *Client) GetHelpdeskSlaReportAnalysis(id int64) (*HelpdeskSlaReportAnalysis, error) {
	hsras, err := c.GetHelpdeskSlaReportAnalysiss([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsras != nil && len(*hsras) > 0 {
		return &((*hsras)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.sla.report.analysis not found", id)
}

// GetHelpdeskSlaReportAnalysiss gets helpdesk.sla.report.analysis existing records.
func (c *Client) GetHelpdeskSlaReportAnalysiss(ids []int64) (*HelpdeskSlaReportAnalysiss, error) {
	hsras := &HelpdeskSlaReportAnalysiss{}
	if err := c.Read(HelpdeskSlaReportAnalysisModel, ids, nil, hsras); err != nil {
		return nil, err
	}
	return hsras, nil
}

// FindHelpdeskSlaReportAnalysis finds helpdesk.sla.report.analysis record by querying it with criteria.
func (c *Client) FindHelpdeskSlaReportAnalysis(criteria *Criteria) (*HelpdeskSlaReportAnalysis, error) {
	hsras := &HelpdeskSlaReportAnalysiss{}
	if err := c.SearchRead(HelpdeskSlaReportAnalysisModel, criteria, NewOptions().Limit(1), hsras); err != nil {
		return nil, err
	}
	if hsras != nil && len(*hsras) > 0 {
		return &((*hsras)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.sla.report.analysis was not found with criteria %v", criteria)
}

// FindHelpdeskSlaReportAnalysiss finds helpdesk.sla.report.analysis records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlaReportAnalysiss(criteria *Criteria, options *Options) (*HelpdeskSlaReportAnalysiss, error) {
	hsras := &HelpdeskSlaReportAnalysiss{}
	if err := c.SearchRead(HelpdeskSlaReportAnalysisModel, criteria, options, hsras); err != nil {
		return nil, err
	}
	return hsras, nil
}

// FindHelpdeskSlaReportAnalysisIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlaReportAnalysisIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskSlaReportAnalysisModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskSlaReportAnalysisId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskSlaReportAnalysisId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskSlaReportAnalysisModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.sla.report.analysis was not found with criteria %v and options %v", criteria, options)
}
