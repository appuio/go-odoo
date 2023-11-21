package odoo

import (
	"fmt"
)

// HrApplicant represents hr.applicant model.
type HrApplicant struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
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
	ApplicantSkillIds           *Relation  `xmlrpc:"applicant_skill_ids,omptempty"`
	ApplicationCount            *Int       `xmlrpc:"application_count,omptempty"`
	ApplicationStatus           *Selection `xmlrpc:"application_status,omptempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AttachmentNumber            *Int       `xmlrpc:"attachment_number,omptempty"`
	Availability                *Time      `xmlrpc:"availability,omptempty"`
	CampaignId                  *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CategIds                    *Relation  `xmlrpc:"categ_ids,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateClosed                  *Time      `xmlrpc:"date_closed,omptempty"`
	DateLastStageUpdate         *Time      `xmlrpc:"date_last_stage_update,omptempty"`
	DateOpen                    *Time      `xmlrpc:"date_open,omptempty"`
	DayClose                    *Float     `xmlrpc:"day_close,omptempty"`
	DayOpen                     *Float     `xmlrpc:"day_open,omptempty"`
	DelayClose                  *Float     `xmlrpc:"delay_close,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omptempty"`
	EmailFrom                   *String    `xmlrpc:"email_from,omptempty"`
	EmpId                       *Many2One  `xmlrpc:"emp_id,omptempty"`
	EmployeeName                *String    `xmlrpc:"employee_name,omptempty"`
	ExtractCanShowResendButton  *Bool      `xmlrpc:"extract_can_show_resend_button,omptempty"`
	ExtractCanShowSendButton    *Bool      `xmlrpc:"extract_can_show_send_button,omptempty"`
	ExtractErrorMessage         *String    `xmlrpc:"extract_error_message,omptempty"`
	ExtractRemoteId             *Int       `xmlrpc:"extract_remote_id,omptempty"`
	ExtractState                *Selection `xmlrpc:"extract_state,omptempty"`
	ExtractStatusCode           *Int       `xmlrpc:"extract_status_code,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	InterviewerIds              *Relation  `xmlrpc:"interviewer_ids,omptempty"`
	IsFirstStage                *Bool      `xmlrpc:"is_first_stage,omptempty"`
	IsInterviewer               *Bool      `xmlrpc:"is_interviewer,omptempty"`
	JobId                       *Many2One  `xmlrpc:"job_id,omptempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omptempty"`
	LastStageId                 *Many2One  `xmlrpc:"last_stage_id,omptempty"`
	LegendBlocked               *String    `xmlrpc:"legend_blocked,omptempty"`
	LegendDone                  *String    `xmlrpc:"legend_done,omptempty"`
	LegendNormal                *String    `xmlrpc:"legend_normal,omptempty"`
	LinkedinProfile             *String    `xmlrpc:"linkedin_profile,omptempty"`
	MediumId                    *Many2One  `xmlrpc:"medium_id,omptempty"`
	MeetingDisplayDate          *Time      `xmlrpc:"meeting_display_date,omptempty"`
	MeetingDisplayText          *String    `xmlrpc:"meeting_display_text,omptempty"`
	MeetingIds                  *Relation  `xmlrpc:"meeting_ids,omptempty"`
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
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerMobile               *String    `xmlrpc:"partner_mobile,omptempty"`
	PartnerName                 *String    `xmlrpc:"partner_name,omptempty"`
	PartnerPhone                *String    `xmlrpc:"partner_phone,omptempty"`
	Priority                    *Selection `xmlrpc:"priority,omptempty"`
	Probability                 *Float     `xmlrpc:"probability,omptempty"`
	RefuseReasonId              *Many2One  `xmlrpc:"refuse_reason_id,omptempty"`
	ResponseId                  *Many2One  `xmlrpc:"response_id,omptempty"`
	ResponseState               *Selection `xmlrpc:"response_state,omptempty"`
	SalaryExpected              *Float     `xmlrpc:"salary_expected,omptempty"`
	SalaryExpectedExtra         *String    `xmlrpc:"salary_expected_extra,omptempty"`
	SalaryProposed              *Float     `xmlrpc:"salary_proposed,omptempty"`
	SalaryProposedExtra         *String    `xmlrpc:"salary_proposed_extra,omptempty"`
	SkillIds                    *Relation  `xmlrpc:"skill_ids,omptempty"`
	SourceId                    *Many2One  `xmlrpc:"source_id,omptempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omptempty"`
	StateProcessed              *Bool      `xmlrpc:"state_processed,omptempty"`
	SurveyId                    *Many2One  `xmlrpc:"survey_id,omptempty"`
	TypeId                      *Many2One  `xmlrpc:"type_id,omptempty"`
	UserEmail                   *String    `xmlrpc:"user_email,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrApplicants represents array of hr.applicant model.
type HrApplicants []HrApplicant

// HrApplicantModel is the odoo model name.
const HrApplicantModel = "hr.applicant"

// Many2One convert HrApplicant to *Many2One.
func (ha *HrApplicant) Many2One() *Many2One {
	return NewMany2One(ha.Id.Get(), "")
}

// CreateHrApplicant creates a new hr.applicant model and returns its id.
func (c *Client) CreateHrApplicant(ha *HrApplicant) (int64, error) {
	ids, err := c.CreateHrApplicants([]*HrApplicant{ha})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrApplicant creates a new hr.applicant model and returns its id.
func (c *Client) CreateHrApplicants(has []*HrApplicant) ([]int64, error) {
	var vv []interface{}
	for _, v := range has {
		vv = append(vv, v)
	}
	return c.Create(HrApplicantModel, vv)
}

// UpdateHrApplicant updates an existing hr.applicant record.
func (c *Client) UpdateHrApplicant(ha *HrApplicant) error {
	return c.UpdateHrApplicants([]int64{ha.Id.Get()}, ha)
}

// UpdateHrApplicants updates existing hr.applicant records.
// All records (represented by ids) will be updated by ha values.
func (c *Client) UpdateHrApplicants(ids []int64, ha *HrApplicant) error {
	return c.Update(HrApplicantModel, ids, ha)
}

// DeleteHrApplicant deletes an existing hr.applicant record.
func (c *Client) DeleteHrApplicant(id int64) error {
	return c.DeleteHrApplicants([]int64{id})
}

// DeleteHrApplicants deletes existing hr.applicant records.
func (c *Client) DeleteHrApplicants(ids []int64) error {
	return c.Delete(HrApplicantModel, ids)
}

// GetHrApplicant gets hr.applicant existing record.
func (c *Client) GetHrApplicant(id int64) (*HrApplicant, error) {
	has, err := c.GetHrApplicants([]int64{id})
	if err != nil {
		return nil, err
	}
	if has != nil && len(*has) > 0 {
		return &((*has)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.applicant not found", id)
}

// GetHrApplicants gets hr.applicant existing records.
func (c *Client) GetHrApplicants(ids []int64) (*HrApplicants, error) {
	has := &HrApplicants{}
	if err := c.Read(HrApplicantModel, ids, nil, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrApplicant finds hr.applicant record by querying it with criteria.
func (c *Client) FindHrApplicant(criteria *Criteria) (*HrApplicant, error) {
	has := &HrApplicants{}
	if err := c.SearchRead(HrApplicantModel, criteria, NewOptions().Limit(1), has); err != nil {
		return nil, err
	}
	if has != nil && len(*has) > 0 {
		return &((*has)[0]), nil
	}
	return nil, fmt.Errorf("hr.applicant was not found with criteria %v", criteria)
}

// FindHrApplicants finds hr.applicant records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicants(criteria *Criteria, options *Options) (*HrApplicants, error) {
	has := &HrApplicants{}
	if err := c.SearchRead(HrApplicantModel, criteria, options, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrApplicantIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrApplicantModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrApplicantId finds record id by querying it with criteria.
func (c *Client) FindHrApplicantId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrApplicantModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.applicant was not found with criteria %v and options %v", criteria, options)
}
