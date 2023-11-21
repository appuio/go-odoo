package odoo

import (
	"fmt"
)

// HrLeave represents hr.leave model.
type HrLeave struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActiveEmployee              *Bool      `xmlrpc:"active_employee,omptempty"`
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
	AllEmployeeIds              *Relation  `xmlrpc:"all_employee_ids,omptempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omptempty"`
	CanApprove                  *Bool      `xmlrpc:"can_approve,omptempty"`
	CanCancel                   *Bool      `xmlrpc:"can_cancel,omptempty"`
	CanReset                    *Bool      `xmlrpc:"can_reset,omptempty"`
	CategoryId                  *Many2One  `xmlrpc:"category_id,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	DurationDisplay             *String    `xmlrpc:"duration_display,omptempty"`
	EmployeeCompanyId           *Many2One  `xmlrpc:"employee_company_id,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	EmployeeIds                 *Relation  `xmlrpc:"employee_ids,omptempty"`
	EmployeeOvertime            *Float     `xmlrpc:"employee_overtime,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FirstApproverId             *Many2One  `xmlrpc:"first_approver_id,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	HasStressDay                *Bool      `xmlrpc:"has_stress_day,omptempty"`
	HolidayAllocationId         *Many2One  `xmlrpc:"holiday_allocation_id,omptempty"`
	HolidayStatusId             *Many2One  `xmlrpc:"holiday_status_id,omptempty"`
	HolidayType                 *Selection `xmlrpc:"holiday_type,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IsHatched                   *Bool      `xmlrpc:"is_hatched,omptempty"`
	IsStriked                   *Bool      `xmlrpc:"is_striked,omptempty"`
	LeaveTypeRequestUnit        *Selection `xmlrpc:"leave_type_request_unit,omptempty"`
	LeaveTypeSupportDocument    *Bool      `xmlrpc:"leave_type_support_document,omptempty"`
	LinkedRequestIds            *Relation  `xmlrpc:"linked_request_ids,omptempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omptempty"`
	MeetingId                   *Many2One  `xmlrpc:"meeting_id,omptempty"`
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
	ModeCompanyId               *Many2One  `xmlrpc:"mode_company_id,omptempty"`
	MultiEmployee               *Bool      `xmlrpc:"multi_employee,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	Notes                       *String    `xmlrpc:"notes,omptempty"`
	NumberOfDays                *Float     `xmlrpc:"number_of_days,omptempty"`
	NumberOfDaysDisplay         *Float     `xmlrpc:"number_of_days_display,omptempty"`
	NumberOfHoursDisplay        *Float     `xmlrpc:"number_of_hours_display,omptempty"`
	NumberOfHoursText           *String    `xmlrpc:"number_of_hours_text,omptempty"`
	OvertimeDeductible          *Bool      `xmlrpc:"overtime_deductible,omptempty"`
	OvertimeId                  *Many2One  `xmlrpc:"overtime_id,omptempty"`
	ParentId                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	PayslipState                *Selection `xmlrpc:"payslip_state,omptempty"`
	PrivateName                 *String    `xmlrpc:"private_name,omptempty"`
	ReportNote                  *String    `xmlrpc:"report_note,omptempty"`
	RequestDateFrom             *Time      `xmlrpc:"request_date_from,omptempty"`
	RequestDateFromPeriod       *Selection `xmlrpc:"request_date_from_period,omptempty"`
	RequestDateTo               *Time      `xmlrpc:"request_date_to,omptempty"`
	RequestHourFrom             *Selection `xmlrpc:"request_hour_from,omptempty"`
	RequestHourTo               *Selection `xmlrpc:"request_hour_to,omptempty"`
	RequestUnitHalf             *Bool      `xmlrpc:"request_unit_half,omptempty"`
	RequestUnitHours            *Bool      `xmlrpc:"request_unit_hours,omptempty"`
	SecondApproverId            *Many2One  `xmlrpc:"second_approver_id,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	SupportedAttachmentIds      *Relation  `xmlrpc:"supported_attachment_ids,omptempty"`
	SupportedAttachmentIdsCount *Int       `xmlrpc:"supported_attachment_ids_count,omptempty"`
	TimesheetIds                *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	Tz                          *Selection `xmlrpc:"tz,omptempty"`
	TzMismatch                  *Bool      `xmlrpc:"tz_mismatch,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	ValidationType              *Selection `xmlrpc:"validation_type,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrLeaves represents array of hr.leave model.
type HrLeaves []HrLeave

// HrLeaveModel is the odoo model name.
const HrLeaveModel = "hr.leave"

// Many2One convert HrLeave to *Many2One.
func (hl *HrLeave) Many2One() *Many2One {
	return NewMany2One(hl.Id.Get(), "")
}

// CreateHrLeave creates a new hr.leave model and returns its id.
func (c *Client) CreateHrLeave(hl *HrLeave) (int64, error) {
	ids, err := c.CreateHrLeaves([]*HrLeave{hl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeave creates a new hr.leave model and returns its id.
func (c *Client) CreateHrLeaves(hls []*HrLeave) ([]int64, error) {
	var vv []interface{}
	for _, v := range hls {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveModel, vv)
}

// UpdateHrLeave updates an existing hr.leave record.
func (c *Client) UpdateHrLeave(hl *HrLeave) error {
	return c.UpdateHrLeaves([]int64{hl.Id.Get()}, hl)
}

// UpdateHrLeaves updates existing hr.leave records.
// All records (represented by ids) will be updated by hl values.
func (c *Client) UpdateHrLeaves(ids []int64, hl *HrLeave) error {
	return c.Update(HrLeaveModel, ids, hl)
}

// DeleteHrLeave deletes an existing hr.leave record.
func (c *Client) DeleteHrLeave(id int64) error {
	return c.DeleteHrLeaves([]int64{id})
}

// DeleteHrLeaves deletes existing hr.leave records.
func (c *Client) DeleteHrLeaves(ids []int64) error {
	return c.Delete(HrLeaveModel, ids)
}

// GetHrLeave gets hr.leave existing record.
func (c *Client) GetHrLeave(id int64) (*HrLeave, error) {
	hls, err := c.GetHrLeaves([]int64{id})
	if err != nil {
		return nil, err
	}
	if hls != nil && len(*hls) > 0 {
		return &((*hls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave not found", id)
}

// GetHrLeaves gets hr.leave existing records.
func (c *Client) GetHrLeaves(ids []int64) (*HrLeaves, error) {
	hls := &HrLeaves{}
	if err := c.Read(HrLeaveModel, ids, nil, hls); err != nil {
		return nil, err
	}
	return hls, nil
}

// FindHrLeave finds hr.leave record by querying it with criteria.
func (c *Client) FindHrLeave(criteria *Criteria) (*HrLeave, error) {
	hls := &HrLeaves{}
	if err := c.SearchRead(HrLeaveModel, criteria, NewOptions().Limit(1), hls); err != nil {
		return nil, err
	}
	if hls != nil && len(*hls) > 0 {
		return &((*hls)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave was not found with criteria %v", criteria)
}

// FindHrLeaves finds hr.leave records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaves(criteria *Criteria, options *Options) (*HrLeaves, error) {
	hls := &HrLeaves{}
	if err := c.SearchRead(HrLeaveModel, criteria, options, hls); err != nil {
		return nil, err
	}
	return hls, nil
}

// FindHrLeaveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave was not found with criteria %v and options %v", criteria, options)
}
