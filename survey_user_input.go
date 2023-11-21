package odoo

import (
	"fmt"
)

// SurveyUserInput represents survey.user_input model.
type SurveyUserInput struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omptempty"`
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
	ApplicantId                 *Relation  `xmlrpc:"applicant_id,omptempty"`
	AppraisalId                 *Many2One  `xmlrpc:"appraisal_id,omptempty"`
	AttemptsCount               *Int       `xmlrpc:"attempts_count,omptempty"`
	AttemptsLimit               *Int       `xmlrpc:"attempts_limit,omptempty"`
	AttemptsNumber              *Int       `xmlrpc:"attempts_number,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Deadline                    *Time      `xmlrpc:"deadline,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Email                       *String    `xmlrpc:"email,omptempty"`
	EndDatetime                 *Time      `xmlrpc:"end_datetime,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	InviteToken                 *String    `xmlrpc:"invite_token,omptempty"`
	IsAttemptsLimited           *Bool      `xmlrpc:"is_attempts_limited,omptempty"`
	IsSessionAnswer             *Bool      `xmlrpc:"is_session_answer,omptempty"`
	LastDisplayedPageId         *Many2One  `xmlrpc:"last_displayed_page_id,omptempty"`
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
	Nickname                    *String    `xmlrpc:"nickname,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PredefinedQuestionIds       *Relation  `xmlrpc:"predefined_question_ids,omptempty"`
	QuestionTimeLimitReached    *Bool      `xmlrpc:"question_time_limit_reached,omptempty"`
	ScoringPercentage           *Float     `xmlrpc:"scoring_percentage,omptempty"`
	ScoringSuccess              *Bool      `xmlrpc:"scoring_success,omptempty"`
	ScoringTotal                *Float     `xmlrpc:"scoring_total,omptempty"`
	ScoringType                 *Selection `xmlrpc:"scoring_type,omptempty"`
	SlideId                     *Many2One  `xmlrpc:"slide_id,omptempty"`
	SlidePartnerId              *Many2One  `xmlrpc:"slide_partner_id,omptempty"`
	StartDatetime               *Time      `xmlrpc:"start_datetime,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	SurveyId                    *Many2One  `xmlrpc:"survey_id,omptempty"`
	SurveyTimeLimitReached      *Bool      `xmlrpc:"survey_time_limit_reached,omptempty"`
	TestEntry                   *Bool      `xmlrpc:"test_entry,omptempty"`
	UserInputLineIds            *Relation  `xmlrpc:"user_input_line_ids,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveyUserInputs represents array of survey.user_input model.
type SurveyUserInputs []SurveyUserInput

// SurveyUserInputModel is the odoo model name.
const SurveyUserInputModel = "survey.user_input"

// Many2One convert SurveyUserInput to *Many2One.
func (su *SurveyUserInput) Many2One() *Many2One {
	return NewMany2One(su.Id.Get(), "")
}

// CreateSurveyUserInput creates a new survey.user_input model and returns its id.
func (c *Client) CreateSurveyUserInput(su *SurveyUserInput) (int64, error) {
	ids, err := c.CreateSurveyUserInputs([]*SurveyUserInput{su})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyUserInput creates a new survey.user_input model and returns its id.
func (c *Client) CreateSurveyUserInputs(sus []*SurveyUserInput) ([]int64, error) {
	var vv []interface{}
	for _, v := range sus {
		vv = append(vv, v)
	}
	return c.Create(SurveyUserInputModel, vv)
}

// UpdateSurveyUserInput updates an existing survey.user_input record.
func (c *Client) UpdateSurveyUserInput(su *SurveyUserInput) error {
	return c.UpdateSurveyUserInputs([]int64{su.Id.Get()}, su)
}

// UpdateSurveyUserInputs updates existing survey.user_input records.
// All records (represented by ids) will be updated by su values.
func (c *Client) UpdateSurveyUserInputs(ids []int64, su *SurveyUserInput) error {
	return c.Update(SurveyUserInputModel, ids, su)
}

// DeleteSurveyUserInput deletes an existing survey.user_input record.
func (c *Client) DeleteSurveyUserInput(id int64) error {
	return c.DeleteSurveyUserInputs([]int64{id})
}

// DeleteSurveyUserInputs deletes existing survey.user_input records.
func (c *Client) DeleteSurveyUserInputs(ids []int64) error {
	return c.Delete(SurveyUserInputModel, ids)
}

// GetSurveyUserInput gets survey.user_input existing record.
func (c *Client) GetSurveyUserInput(id int64) (*SurveyUserInput, error) {
	sus, err := c.GetSurveyUserInputs([]int64{id})
	if err != nil {
		return nil, err
	}
	if sus != nil && len(*sus) > 0 {
		return &((*sus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.user_input not found", id)
}

// GetSurveyUserInputs gets survey.user_input existing records.
func (c *Client) GetSurveyUserInputs(ids []int64) (*SurveyUserInputs, error) {
	sus := &SurveyUserInputs{}
	if err := c.Read(SurveyUserInputModel, ids, nil, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInput finds survey.user_input record by querying it with criteria.
func (c *Client) FindSurveyUserInput(criteria *Criteria) (*SurveyUserInput, error) {
	sus := &SurveyUserInputs{}
	if err := c.SearchRead(SurveyUserInputModel, criteria, NewOptions().Limit(1), sus); err != nil {
		return nil, err
	}
	if sus != nil && len(*sus) > 0 {
		return &((*sus)[0]), nil
	}
	return nil, fmt.Errorf("survey.user_input was not found with criteria %v", criteria)
}

// FindSurveyUserInputs finds survey.user_input records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputs(criteria *Criteria, options *Options) (*SurveyUserInputs, error) {
	sus := &SurveyUserInputs{}
	if err := c.SearchRead(SurveyUserInputModel, criteria, options, sus); err != nil {
		return nil, err
	}
	return sus, nil
}

// FindSurveyUserInputIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyUserInputModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyUserInputId finds record id by querying it with criteria.
func (c *Client) FindSurveyUserInputId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyUserInputModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.user_input was not found with criteria %v and options %v", criteria, options)
}
