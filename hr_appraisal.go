package odoo

import (
	"fmt"
)

// HrAppraisal represents hr.appraisal model.
type HrAppraisal struct {
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
	AppraisalPlanPosted         *Bool      `xmlrpc:"appraisal_plan_posted,omptempty"`
	AssessmentNote              *Many2One  `xmlrpc:"assessment_note,omptempty"`
	Avatar128                   *String    `xmlrpc:"avatar_128,omptempty"`
	Avatar1920                  *String    `xmlrpc:"avatar_1920,omptempty"`
	CanSeeEmployeePublish       *Bool      `xmlrpc:"can_see_employee_publish,omptempty"`
	CanSeeManagerPublish        *Bool      `xmlrpc:"can_see_manager_publish,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateClose                   *Time      `xmlrpc:"date_close,omptempty"`
	DateFinalInterview          *Time      `xmlrpc:"date_final_interview,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmployeeAppraisalCount      *Int       `xmlrpc:"employee_appraisal_count,omptempty"`
	EmployeeAutocompleteIds     *Relation  `xmlrpc:"employee_autocomplete_ids,omptempty"`
	EmployeeFeedback            *String    `xmlrpc:"employee_feedback,omptempty"`
	EmployeeFeedbackIds         *Relation  `xmlrpc:"employee_feedback_ids,omptempty"`
	EmployeeFeedbackPublished   *Bool      `xmlrpc:"employee_feedback_published,omptempty"`
	EmployeeFeedbackTemplate    *String    `xmlrpc:"employee_feedback_template,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	EmployeeUserId              *Many2One  `xmlrpc:"employee_user_id,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	Image128                    *String    `xmlrpc:"image_128,omptempty"`
	Image1920                   *String    `xmlrpc:"image_1920,omptempty"`
	IsAppraisalManager          *Bool      `xmlrpc:"is_appraisal_manager,omptempty"`
	IsImplicitManager           *Bool      `xmlrpc:"is_implicit_manager,omptempty"`
	LastAppraisalDate           *Time      `xmlrpc:"last_appraisal_date,omptempty"`
	LastAppraisalId             *Many2One  `xmlrpc:"last_appraisal_id,omptempty"`
	ManagerFeedback             *String    `xmlrpc:"manager_feedback,omptempty"`
	ManagerFeedbackPublished    *Bool      `xmlrpc:"manager_feedback_published,omptempty"`
	ManagerFeedbackTemplate     *String    `xmlrpc:"manager_feedback_template,omptempty"`
	ManagerIds                  *Relation  `xmlrpc:"manager_ids,omptempty"`
	ManagerUserIds              *Relation  `xmlrpc:"manager_user_ids,omptempty"`
	MeetingCountDisplay         *String    `xmlrpc:"meeting_count_display,omptempty"`
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
	Note                        *String    `xmlrpc:"note,omptempty"`
	ShowEmployeeFeedbackFull    *Bool      `xmlrpc:"show_employee_feedback_full,omptempty"`
	ShowManagerFeedbackFull     *Bool      `xmlrpc:"show_manager_feedback_full,omptempty"`
	SkillIds                    *Relation  `xmlrpc:"skill_ids,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	SurveyIds                   *Relation  `xmlrpc:"survey_ids,omptempty"`
	UncompleteGoalsCount        *Int       `xmlrpc:"uncomplete_goals_count,omptempty"`
	WaitingFeedback             *Bool      `xmlrpc:"waiting_feedback,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrAppraisals represents array of hr.appraisal model.
type HrAppraisals []HrAppraisal

// HrAppraisalModel is the odoo model name.
const HrAppraisalModel = "hr.appraisal"

// Many2One convert HrAppraisal to *Many2One.
func (ha *HrAppraisal) Many2One() *Many2One {
	return NewMany2One(ha.Id.Get(), "")
}

// CreateHrAppraisal creates a new hr.appraisal model and returns its id.
func (c *Client) CreateHrAppraisal(ha *HrAppraisal) (int64, error) {
	ids, err := c.CreateHrAppraisals([]*HrAppraisal{ha})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAppraisal creates a new hr.appraisal model and returns its id.
func (c *Client) CreateHrAppraisals(has []*HrAppraisal) ([]int64, error) {
	var vv []interface{}
	for _, v := range has {
		vv = append(vv, v)
	}
	return c.Create(HrAppraisalModel, vv)
}

// UpdateHrAppraisal updates an existing hr.appraisal record.
func (c *Client) UpdateHrAppraisal(ha *HrAppraisal) error {
	return c.UpdateHrAppraisals([]int64{ha.Id.Get()}, ha)
}

// UpdateHrAppraisals updates existing hr.appraisal records.
// All records (represented by ids) will be updated by ha values.
func (c *Client) UpdateHrAppraisals(ids []int64, ha *HrAppraisal) error {
	return c.Update(HrAppraisalModel, ids, ha)
}

// DeleteHrAppraisal deletes an existing hr.appraisal record.
func (c *Client) DeleteHrAppraisal(id int64) error {
	return c.DeleteHrAppraisals([]int64{id})
}

// DeleteHrAppraisals deletes existing hr.appraisal records.
func (c *Client) DeleteHrAppraisals(ids []int64) error {
	return c.Delete(HrAppraisalModel, ids)
}

// GetHrAppraisal gets hr.appraisal existing record.
func (c *Client) GetHrAppraisal(id int64) (*HrAppraisal, error) {
	has, err := c.GetHrAppraisals([]int64{id})
	if err != nil {
		return nil, err
	}
	if has != nil && len(*has) > 0 {
		return &((*has)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.appraisal not found", id)
}

// GetHrAppraisals gets hr.appraisal existing records.
func (c *Client) GetHrAppraisals(ids []int64) (*HrAppraisals, error) {
	has := &HrAppraisals{}
	if err := c.Read(HrAppraisalModel, ids, nil, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrAppraisal finds hr.appraisal record by querying it with criteria.
func (c *Client) FindHrAppraisal(criteria *Criteria) (*HrAppraisal, error) {
	has := &HrAppraisals{}
	if err := c.SearchRead(HrAppraisalModel, criteria, NewOptions().Limit(1), has); err != nil {
		return nil, err
	}
	if has != nil && len(*has) > 0 {
		return &((*has)[0]), nil
	}
	return nil, fmt.Errorf("hr.appraisal was not found with criteria %v", criteria)
}

// FindHrAppraisals finds hr.appraisal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisals(criteria *Criteria, options *Options) (*HrAppraisals, error) {
	has := &HrAppraisals{}
	if err := c.SearchRead(HrAppraisalModel, criteria, options, has); err != nil {
		return nil, err
	}
	return has, nil
}

// FindHrAppraisalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAppraisalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAppraisalId finds record id by querying it with criteria.
func (c *Client) FindHrAppraisalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAppraisalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.appraisal was not found with criteria %v and options %v", criteria, options)
}
