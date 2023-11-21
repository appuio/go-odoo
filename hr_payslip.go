package odoo

import (
	"fmt"
)

// HrPayslip represents hr.payslip model.
type HrPayslip struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttendanceCount             *Int       `xmlrpc:"attendance_count,omptempty"`
	BasicWage                   *Float     `xmlrpc:"basic_wage,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	ComputeDate                 *Time      `xmlrpc:"compute_date,omptempty"`
	ContractDomainIds           *Relation  `xmlrpc:"contract_domain_ids,omptempty"`
	ContractId                  *Many2One  `xmlrpc:"contract_id,omptempty"`
	CountryCode                 *String    `xmlrpc:"country_code,omptempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CreditNote                  *Bool      `xmlrpc:"credit_note,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                        *Time      `xmlrpc:"date,omptempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Edited                      *Bool      `xmlrpc:"edited,omptempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	HasNegativeNetToReport      *Bool      `xmlrpc:"has_negative_net_to_report,omptempty"`
	HasRefundSlip               *Bool      `xmlrpc:"has_refund_slip,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	InputLineIds                *Relation  `xmlrpc:"input_line_ids,omptempty"`
	IsRegular                   *Bool      `xmlrpc:"is_regular,omptempty"`
	IsSuperuser                 *Bool      `xmlrpc:"is_superuser,omptempty"`
	JobId                       *Many2One  `xmlrpc:"job_id,omptempty"`
	JournalId                   *Many2One  `xmlrpc:"journal_id,omptempty"`
	LineIds                     *Relation  `xmlrpc:"line_ids,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent              *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MoveId                      *Many2One  `xmlrpc:"move_id,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	NegativeNetToReportAmount   *Float     `xmlrpc:"negative_net_to_report_amount,omptempty"`
	NegativeNetToReportDisplay  *Bool      `xmlrpc:"negative_net_to_report_display,omptempty"`
	NegativeNetToReportMessage  *String    `xmlrpc:"negative_net_to_report_message,omptempty"`
	NetWage                     *Float     `xmlrpc:"net_wage,omptempty"`
	NormalWage                  *Int       `xmlrpc:"normal_wage,omptempty"`
	Note                        *String    `xmlrpc:"note,omptempty"`
	Number                      *String    `xmlrpc:"number,omptempty"`
	Paid                        *Bool      `xmlrpc:"paid,omptempty"`
	PayslipRunId                *Many2One  `xmlrpc:"payslip_run_id,omptempty"`
	PlanningSlotCount           *Int       `xmlrpc:"planning_slot_count,omptempty"`
	QueuedForPdf                *Bool      `xmlrpc:"queued_for_pdf,omptempty"`
	SalaryAttachmentCount       *Int       `xmlrpc:"salary_attachment_count,omptempty"`
	SalaryAttachmentIds         *Relation  `xmlrpc:"salary_attachment_ids,omptempty"`
	SepaExport                  *String    `xmlrpc:"sepa_export,omptempty"`
	SepaExportDate              *Time      `xmlrpc:"sepa_export_date,omptempty"`
	SepaExportFilename          *String    `xmlrpc:"sepa_export_filename,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	StructId                    *Many2One  `xmlrpc:"struct_id,omptempty"`
	StructTypeId                *Many2One  `xmlrpc:"struct_type_id,omptempty"`
	SumWorkedHours              *Float     `xmlrpc:"sum_worked_hours,omptempty"`
	WageType                    *Selection `xmlrpc:"wage_type,omptempty"`
	WarningMessage              *String    `xmlrpc:"warning_message,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WorkedDaysLineIds           *Relation  `xmlrpc:"worked_days_line_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
	XStudioMany2ManyFieldLTdHn  *Relation  `xmlrpc:"x_studio_many2many_field_LTdHn,omptempty"`
	XStudioMonetaryFieldKKD5Z   *Float     `xmlrpc:"x_studio_monetary_field_kKD5Z,omptempty"`
	XStudioOne2ManyFieldJ0Wi8   *Relation  `xmlrpc:"x_studio_one2many_field_j0Wi8,omptempty"`
	XStudioRelatedFieldOWkbK    *Float     `xmlrpc:"x_studio_related_field_OWkbK,omptempty"`
	XStudioRelatedFieldVx980    *String    `xmlrpc:"x_studio_related_field_Vx980,omptempty"`
	XStudioRelatedFieldKoycg    *String    `xmlrpc:"x_studio_related_field_koycg,omptempty"`
	XStudioRelatedFieldLZU07    *String    `xmlrpc:"x_studio_related_field_lZU07,omptempty"`
	XStudioRelatedFieldNQXtY    *Float     `xmlrpc:"x_studio_related_field_nQXtY,omptempty"`
}

// HrPayslips represents array of hr.payslip model.
type HrPayslips []HrPayslip

// HrPayslipModel is the odoo model name.
const HrPayslipModel = "hr.payslip"

// Many2One convert HrPayslip to *Many2One.
func (hp *HrPayslip) Many2One() *Many2One {
	return NewMany2One(hp.Id.Get(), "")
}

// CreateHrPayslip creates a new hr.payslip model and returns its id.
func (c *Client) CreateHrPayslip(hp *HrPayslip) (int64, error) {
	ids, err := c.CreateHrPayslips([]*HrPayslip{hp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslip creates a new hr.payslip model and returns its id.
func (c *Client) CreateHrPayslips(hps []*HrPayslip) ([]int64, error) {
	var vv []interface{}
	for _, v := range hps {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipModel, vv)
}

// UpdateHrPayslip updates an existing hr.payslip record.
func (c *Client) UpdateHrPayslip(hp *HrPayslip) error {
	return c.UpdateHrPayslips([]int64{hp.Id.Get()}, hp)
}

// UpdateHrPayslips updates existing hr.payslip records.
// All records (represented by ids) will be updated by hp values.
func (c *Client) UpdateHrPayslips(ids []int64, hp *HrPayslip) error {
	return c.Update(HrPayslipModel, ids, hp)
}

// DeleteHrPayslip deletes an existing hr.payslip record.
func (c *Client) DeleteHrPayslip(id int64) error {
	return c.DeleteHrPayslips([]int64{id})
}

// DeleteHrPayslips deletes existing hr.payslip records.
func (c *Client) DeleteHrPayslips(ids []int64) error {
	return c.Delete(HrPayslipModel, ids)
}

// GetHrPayslip gets hr.payslip existing record.
func (c *Client) GetHrPayslip(id int64) (*HrPayslip, error) {
	hps, err := c.GetHrPayslips([]int64{id})
	if err != nil {
		return nil, err
	}
	if hps != nil && len(*hps) > 0 {
		return &((*hps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip not found", id)
}

// GetHrPayslips gets hr.payslip existing records.
func (c *Client) GetHrPayslips(ids []int64) (*HrPayslips, error) {
	hps := &HrPayslips{}
	if err := c.Read(HrPayslipModel, ids, nil, hps); err != nil {
		return nil, err
	}
	return hps, nil
}

// FindHrPayslip finds hr.payslip record by querying it with criteria.
func (c *Client) FindHrPayslip(criteria *Criteria) (*HrPayslip, error) {
	hps := &HrPayslips{}
	if err := c.SearchRead(HrPayslipModel, criteria, NewOptions().Limit(1), hps); err != nil {
		return nil, err
	}
	if hps != nil && len(*hps) > 0 {
		return &((*hps)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip was not found with criteria %v", criteria)
}

// FindHrPayslips finds hr.payslip records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslips(criteria *Criteria, options *Options) (*HrPayslips, error) {
	hps := &HrPayslips{}
	if err := c.SearchRead(HrPayslipModel, criteria, options, hps); err != nil {
		return nil, err
	}
	return hps, nil
}

// FindHrPayslipIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip was not found with criteria %v and options %v", criteria, options)
}
