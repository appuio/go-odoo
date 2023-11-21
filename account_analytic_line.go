package odoo

import (
	"fmt"
)

// AccountAnalyticLine represents account.analytic.line model.
type AccountAnalyticLine struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId                *Many2One  `xmlrpc:"account_id,omptempty"`
	Amount                   *Float     `xmlrpc:"amount,omptempty"`
	AncestorTaskId           *Many2One  `xmlrpc:"ancestor_task_id,omptempty"`
	BillableMeteredUsage     *Bool      `xmlrpc:"billable_metered_usage,omptempty"`
	CanEdit                  *Bool      `xmlrpc:"can_edit,omptempty"`
	Category                 *Selection `xmlrpc:"category,omptempty"`
	Code                     *String    `xmlrpc:"code,omptempty"`
	CommercialPartnerId      *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                     *Time      `xmlrpc:"date,omptempty"`
	DepartmentId             *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	DisplaySol               *Bool      `xmlrpc:"display_sol,omptempty"`
	DisplayTask              *Bool      `xmlrpc:"display_task,omptempty"`
	DisplayTimer             *Bool      `xmlrpc:"display_timer,omptempty"`
	DisplayTimerPause        *Bool      `xmlrpc:"display_timer_pause,omptempty"`
	DisplayTimerResume       *Bool      `xmlrpc:"display_timer_resume,omptempty"`
	DisplayTimerStartPrimary *Bool      `xmlrpc:"display_timer_start_primary,omptempty"`
	DisplayTimerStop         *Bool      `xmlrpc:"display_timer_stop,omptempty"`
	DurationUnitAmount       *Float     `xmlrpc:"duration_unit_amount,omptempty"`
	EmployeeId               *Many2One  `xmlrpc:"employee_id,omptempty"`
	EncodingUomId            *Many2One  `xmlrpc:"encoding_uom_id,omptempty"`
	GeneralAccountId         *Many2One  `xmlrpc:"general_account_id,omptempty"`
	GlobalLeaveId            *Many2One  `xmlrpc:"global_leave_id,omptempty"`
	HasHelpdeskTeam          *Bool      `xmlrpc:"has_helpdesk_team,omptempty"`
	HasSoAccess              *Bool      `xmlrpc:"has_so_access,omptempty"`
	HelpdeskTicketId         *Many2One  `xmlrpc:"helpdesk_ticket_id,omptempty"`
	HolidayId                *Many2One  `xmlrpc:"holiday_id,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IsSoLineEdited           *Bool      `xmlrpc:"is_so_line_edited,omptempty"`
	IsTimerRunning           *Bool      `xmlrpc:"is_timer_running,omptempty"`
	IsTimesheet              *Bool      `xmlrpc:"is_timesheet,omptempty"`
	JobTitle                 *String    `xmlrpc:"job_title,omptempty"`
	JournalId                *Many2One  `xmlrpc:"journal_id,omptempty"`
	L10NDin5008DocumentTitle *String    `xmlrpc:"l10n_din5008_document_title,omptempty"`
	L10NDin5008TemplateData  *String    `xmlrpc:"l10n_din5008_template_data,omptempty"`
	ManagerId                *Many2One  `xmlrpc:"manager_id,omptempty"`
	MeteredUsageInvoiceId    *Many2One  `xmlrpc:"metered_usage_invoice_id,omptempty"`
	MoveLineId               *Many2One  `xmlrpc:"move_line_id,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	OrderId                  *Many2One  `xmlrpc:"order_id,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PlanId                   *Many2One  `xmlrpc:"plan_id,omptempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId     *Many2One  `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomId             *Many2One  `xmlrpc:"product_uom_id,omptempty"`
	ProjectId                *Many2One  `xmlrpc:"project_id,omptempty"`
	Ref                      *String    `xmlrpc:"ref,omptempty"`
	SlotId                   *Many2One  `xmlrpc:"slot_id,omptempty"`
	SoLine                   *Many2One  `xmlrpc:"so_line,omptempty"`
	TaskId                   *Many2One  `xmlrpc:"task_id,omptempty"`
	TimerPause               *Time      `xmlrpc:"timer_pause,omptempty"`
	TimerStart               *Time      `xmlrpc:"timer_start,omptempty"`
	TimesheetInvoiceId       *Many2One  `xmlrpc:"timesheet_invoice_id,omptempty"`
	TimesheetInvoiceType     *Selection `xmlrpc:"timesheet_invoice_type,omptempty"`
	UnitAmount               *Float     `xmlrpc:"unit_amount,omptempty"`
	UnitAmountValidate       *Float     `xmlrpc:"unit_amount_validate,omptempty"`
	UserCanValidate          *Bool      `xmlrpc:"user_can_validate,omptempty"`
	UserId                   *Many2One  `xmlrpc:"user_id,omptempty"`
	UserTimerId              *Relation  `xmlrpc:"user_timer_id,omptempty"`
	Validated                *Bool      `xmlrpc:"validated,omptempty"`
	ValidatedStatus          *Selection `xmlrpc:"validated_status,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAnalyticLines represents array of account.analytic.line model.
type AccountAnalyticLines []AccountAnalyticLine

// AccountAnalyticLineModel is the odoo model name.
const AccountAnalyticLineModel = "account.analytic.line"

// Many2One convert AccountAnalyticLine to *Many2One.
func (aal *AccountAnalyticLine) Many2One() *Many2One {
	return NewMany2One(aal.Id.Get(), "")
}

// CreateAccountAnalyticLine creates a new account.analytic.line model and returns its id.
func (c *Client) CreateAccountAnalyticLine(aal *AccountAnalyticLine) (int64, error) {
	ids, err := c.CreateAccountAnalyticLines([]*AccountAnalyticLine{aal})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAnalyticLine creates a new account.analytic.line model and returns its id.
func (c *Client) CreateAccountAnalyticLines(aals []*AccountAnalyticLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range aals {
		vv = append(vv, v)
	}
	return c.Create(AccountAnalyticLineModel, vv)
}

// UpdateAccountAnalyticLine updates an existing account.analytic.line record.
func (c *Client) UpdateAccountAnalyticLine(aal *AccountAnalyticLine) error {
	return c.UpdateAccountAnalyticLines([]int64{aal.Id.Get()}, aal)
}

// UpdateAccountAnalyticLines updates existing account.analytic.line records.
// All records (represented by ids) will be updated by aal values.
func (c *Client) UpdateAccountAnalyticLines(ids []int64, aal *AccountAnalyticLine) error {
	return c.Update(AccountAnalyticLineModel, ids, aal)
}

// DeleteAccountAnalyticLine deletes an existing account.analytic.line record.
func (c *Client) DeleteAccountAnalyticLine(id int64) error {
	return c.DeleteAccountAnalyticLines([]int64{id})
}

// DeleteAccountAnalyticLines deletes existing account.analytic.line records.
func (c *Client) DeleteAccountAnalyticLines(ids []int64) error {
	return c.Delete(AccountAnalyticLineModel, ids)
}

// GetAccountAnalyticLine gets account.analytic.line existing record.
func (c *Client) GetAccountAnalyticLine(id int64) (*AccountAnalyticLine, error) {
	aals, err := c.GetAccountAnalyticLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if aals != nil && len(*aals) > 0 {
		return &((*aals)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.line not found", id)
}

// GetAccountAnalyticLines gets account.analytic.line existing records.
func (c *Client) GetAccountAnalyticLines(ids []int64) (*AccountAnalyticLines, error) {
	aals := &AccountAnalyticLines{}
	if err := c.Read(AccountAnalyticLineModel, ids, nil, aals); err != nil {
		return nil, err
	}
	return aals, nil
}

// FindAccountAnalyticLine finds account.analytic.line record by querying it with criteria.
func (c *Client) FindAccountAnalyticLine(criteria *Criteria) (*AccountAnalyticLine, error) {
	aals := &AccountAnalyticLines{}
	if err := c.SearchRead(AccountAnalyticLineModel, criteria, NewOptions().Limit(1), aals); err != nil {
		return nil, err
	}
	if aals != nil && len(*aals) > 0 {
		return &((*aals)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.line was not found with criteria %v", criteria)
}

// FindAccountAnalyticLines finds account.analytic.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticLines(criteria *Criteria, options *Options) (*AccountAnalyticLines, error) {
	aals := &AccountAnalyticLines{}
	if err := c.SearchRead(AccountAnalyticLineModel, criteria, options, aals); err != nil {
		return nil, err
	}
	return aals, nil
}

// FindAccountAnalyticLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticLineId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.line was not found with criteria %v and options %v", criteria, options)
}
