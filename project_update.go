package odoo

import (
	"fmt"
)

// ProjectUpdate represents project.update model.
type ProjectUpdate struct {
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
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date                        *Time      `xmlrpc:"date,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
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
	NameCropped                 *String    `xmlrpc:"name_cropped,omptempty"`
	Progress                    *Int       `xmlrpc:"progress,omptempty"`
	ProgressPercentage          *Float     `xmlrpc:"progress_percentage,omptempty"`
	ProjectId                   *Many2One  `xmlrpc:"project_id,omptempty"`
	Status                      *Selection `xmlrpc:"status,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProjectUpdates represents array of project.update model.
type ProjectUpdates []ProjectUpdate

// ProjectUpdateModel is the odoo model name.
const ProjectUpdateModel = "project.update"

// Many2One convert ProjectUpdate to *Many2One.
func (pu *ProjectUpdate) Many2One() *Many2One {
	return NewMany2One(pu.Id.Get(), "")
}

// CreateProjectUpdate creates a new project.update model and returns its id.
func (c *Client) CreateProjectUpdate(pu *ProjectUpdate) (int64, error) {
	ids, err := c.CreateProjectUpdates([]*ProjectUpdate{pu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectUpdate creates a new project.update model and returns its id.
func (c *Client) CreateProjectUpdates(pus []*ProjectUpdate) ([]int64, error) {
	var vv []interface{}
	for _, v := range pus {
		vv = append(vv, v)
	}
	return c.Create(ProjectUpdateModel, vv)
}

// UpdateProjectUpdate updates an existing project.update record.
func (c *Client) UpdateProjectUpdate(pu *ProjectUpdate) error {
	return c.UpdateProjectUpdates([]int64{pu.Id.Get()}, pu)
}

// UpdateProjectUpdates updates existing project.update records.
// All records (represented by ids) will be updated by pu values.
func (c *Client) UpdateProjectUpdates(ids []int64, pu *ProjectUpdate) error {
	return c.Update(ProjectUpdateModel, ids, pu)
}

// DeleteProjectUpdate deletes an existing project.update record.
func (c *Client) DeleteProjectUpdate(id int64) error {
	return c.DeleteProjectUpdates([]int64{id})
}

// DeleteProjectUpdates deletes existing project.update records.
func (c *Client) DeleteProjectUpdates(ids []int64) error {
	return c.Delete(ProjectUpdateModel, ids)
}

// GetProjectUpdate gets project.update existing record.
func (c *Client) GetProjectUpdate(id int64) (*ProjectUpdate, error) {
	pus, err := c.GetProjectUpdates([]int64{id})
	if err != nil {
		return nil, err
	}
	if pus != nil && len(*pus) > 0 {
		return &((*pus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.update not found", id)
}

// GetProjectUpdates gets project.update existing records.
func (c *Client) GetProjectUpdates(ids []int64) (*ProjectUpdates, error) {
	pus := &ProjectUpdates{}
	if err := c.Read(ProjectUpdateModel, ids, nil, pus); err != nil {
		return nil, err
	}
	return pus, nil
}

// FindProjectUpdate finds project.update record by querying it with criteria.
func (c *Client) FindProjectUpdate(criteria *Criteria) (*ProjectUpdate, error) {
	pus := &ProjectUpdates{}
	if err := c.SearchRead(ProjectUpdateModel, criteria, NewOptions().Limit(1), pus); err != nil {
		return nil, err
	}
	if pus != nil && len(*pus) > 0 {
		return &((*pus)[0]), nil
	}
	return nil, fmt.Errorf("project.update was not found with criteria %v", criteria)
}

// FindProjectUpdates finds project.update records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectUpdates(criteria *Criteria, options *Options) (*ProjectUpdates, error) {
	pus := &ProjectUpdates{}
	if err := c.SearchRead(ProjectUpdateModel, criteria, options, pus); err != nil {
		return nil, err
	}
	return pus, nil
}

// FindProjectUpdateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectUpdateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectUpdateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectUpdateId finds record id by querying it with criteria.
func (c *Client) FindProjectUpdateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectUpdateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.update was not found with criteria %v and options %v", criteria, options)
}
