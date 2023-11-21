package odoo

import (
	"fmt"
)

// HrAppraisalGoal represents hr.appraisal.goal model.
type HrAppraisalGoal struct {
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
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Deadline                    *Time      `xmlrpc:"deadline,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmployeeAutocompleteIds     *Relation  `xmlrpc:"employee_autocomplete_ids,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	IsImplicitManager           *Bool      `xmlrpc:"is_implicit_manager,omptempty"`
	IsManager                   *Bool      `xmlrpc:"is_manager,omptempty"`
	ManagerId                   *Many2One  `xmlrpc:"manager_id,omptempty"`
	ManagerUserId               *Many2One  `xmlrpc:"manager_user_id,omptempty"`
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
	Progression                 *Selection `xmlrpc:"progression,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrAppraisalGoals represents array of hr.appraisal.goal model.
type HrAppraisalGoals []HrAppraisalGoal

// HrAppraisalGoalModel is the odoo model name.
const HrAppraisalGoalModel = "hr.appraisal.goal"

// Many2One convert HrAppraisalGoal to *Many2One.
func (hag *HrAppraisalGoal) Many2One() *Many2One {
	return NewMany2One(hag.Id.Get(), "")
}

// CreateHrAppraisalGoal creates a new hr.appraisal.goal model and returns its id.
func (c *Client) CreateHrAppraisalGoal(hag *HrAppraisalGoal) (int64, error) {
	ids, err := c.CreateHrAppraisalGoals([]*HrAppraisalGoal{hag})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAppraisalGoal creates a new hr.appraisal.goal model and returns its id.
func (c *Client) CreateHrAppraisalGoals(hags []*HrAppraisalGoal) ([]int64, error) {
	var vv []interface{}
	for _, v := range hags {
		vv = append(vv, v)
	}
	return c.Create(HrAppraisalGoalModel, vv)
}

// UpdateHrAppraisalGoal updates an existing hr.appraisal.goal record.
func (c *Client) UpdateHrAppraisalGoal(hag *HrAppraisalGoal) error {
	return c.UpdateHrAppraisalGoals([]int64{hag.Id.Get()}, hag)
}

// UpdateHrAppraisalGoals updates existing hr.appraisal.goal records.
// All records (represented by ids) will be updated by hag values.
func (c *Client) UpdateHrAppraisalGoals(ids []int64, hag *HrAppraisalGoal) error {
	return c.Update(HrAppraisalGoalModel, ids, hag)
}

// DeleteHrAppraisalGoal deletes an existing hr.appraisal.goal record.
func (c *Client) DeleteHrAppraisalGoal(id int64) error {
	return c.DeleteHrAppraisalGoals([]int64{id})
}

// DeleteHrAppraisalGoals deletes existing hr.appraisal.goal records.
func (c *Client) DeleteHrAppraisalGoals(ids []int64) error {
	return c.Delete(HrAppraisalGoalModel, ids)
}

// GetHrAppraisalGoal gets hr.appraisal.goal existing record.
func (c *Client) GetHrAppraisalGoal(id int64) (*HrAppraisalGoal, error) {
	hags, err := c.GetHrAppraisalGoals([]int64{id})
	if err != nil {
		return nil, err
	}
	if hags != nil && len(*hags) > 0 {
		return &((*hags)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.appraisal.goal not found", id)
}

// GetHrAppraisalGoals gets hr.appraisal.goal existing records.
func (c *Client) GetHrAppraisalGoals(ids []int64) (*HrAppraisalGoals, error) {
	hags := &HrAppraisalGoals{}
	if err := c.Read(HrAppraisalGoalModel, ids, nil, hags); err != nil {
		return nil, err
	}
	return hags, nil
}

// FindHrAppraisalGoal finds hr.appraisal.goal record by querying it with criteria.
func (c *Client) FindHrAppraisalGoal(criteria *Criteria) (*HrAppraisalGoal, error) {
	hags := &HrAppraisalGoals{}
	if err := c.SearchRead(HrAppraisalGoalModel, criteria, NewOptions().Limit(1), hags); err != nil {
		return nil, err
	}
	if hags != nil && len(*hags) > 0 {
		return &((*hags)[0]), nil
	}
	return nil, fmt.Errorf("hr.appraisal.goal was not found with criteria %v", criteria)
}

// FindHrAppraisalGoals finds hr.appraisal.goal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalGoals(criteria *Criteria, options *Options) (*HrAppraisalGoals, error) {
	hags := &HrAppraisalGoals{}
	if err := c.SearchRead(HrAppraisalGoalModel, criteria, options, hags); err != nil {
		return nil, err
	}
	return hags, nil
}

// FindHrAppraisalGoalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalGoalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAppraisalGoalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAppraisalGoalId finds record id by querying it with criteria.
func (c *Client) FindHrAppraisalGoalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAppraisalGoalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.appraisal.goal was not found with criteria %v and options %v", criteria, options)
}
