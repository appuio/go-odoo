package odoo

import (
	"fmt"
)

// ProjectMilestone represents project.milestone model.
type ProjectMilestone struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	AllowBillable            *Bool     `xmlrpc:"allow_billable,omptempty"`
	CanBeMarkedAsDone        *Bool     `xmlrpc:"can_be_marked_as_done,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	Deadline                 *Time     `xmlrpc:"deadline,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	FailedMessageIds         *Relation `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsDeadlineExceeded       *Bool     `xmlrpc:"is_deadline_exceeded,omptempty"`
	IsDeadlineFuture         *Bool     `xmlrpc:"is_deadline_future,omptempty"`
	IsReached                *Bool     `xmlrpc:"is_reached,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent           *String   `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	ProjectId                *Many2One `xmlrpc:"project_id,omptempty"`
	ProjectPartnerId         *Many2One `xmlrpc:"project_partner_id,omptempty"`
	QuantityPercentage       *Float    `xmlrpc:"quantity_percentage,omptempty"`
	ReachedDate              *Time     `xmlrpc:"reached_date,omptempty"`
	SaleLineId               *Many2One `xmlrpc:"sale_line_id,omptempty"`
	SaleLineName             *String   `xmlrpc:"sale_line_name,omptempty"`
	TaskCount                *Int      `xmlrpc:"task_count,omptempty"`
	TaskIds                  *Relation `xmlrpc:"task_ids,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectMilestones represents array of project.milestone model.
type ProjectMilestones []ProjectMilestone

// ProjectMilestoneModel is the odoo model name.
const ProjectMilestoneModel = "project.milestone"

// Many2One convert ProjectMilestone to *Many2One.
func (pm *ProjectMilestone) Many2One() *Many2One {
	return NewMany2One(pm.Id.Get(), "")
}

// CreateProjectMilestone creates a new project.milestone model and returns its id.
func (c *Client) CreateProjectMilestone(pm *ProjectMilestone) (int64, error) {
	ids, err := c.CreateProjectMilestones([]*ProjectMilestone{pm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectMilestone creates a new project.milestone model and returns its id.
func (c *Client) CreateProjectMilestones(pms []*ProjectMilestone) ([]int64, error) {
	var vv []interface{}
	for _, v := range pms {
		vv = append(vv, v)
	}
	return c.Create(ProjectMilestoneModel, vv)
}

// UpdateProjectMilestone updates an existing project.milestone record.
func (c *Client) UpdateProjectMilestone(pm *ProjectMilestone) error {
	return c.UpdateProjectMilestones([]int64{pm.Id.Get()}, pm)
}

// UpdateProjectMilestones updates existing project.milestone records.
// All records (represented by ids) will be updated by pm values.
func (c *Client) UpdateProjectMilestones(ids []int64, pm *ProjectMilestone) error {
	return c.Update(ProjectMilestoneModel, ids, pm)
}

// DeleteProjectMilestone deletes an existing project.milestone record.
func (c *Client) DeleteProjectMilestone(id int64) error {
	return c.DeleteProjectMilestones([]int64{id})
}

// DeleteProjectMilestones deletes existing project.milestone records.
func (c *Client) DeleteProjectMilestones(ids []int64) error {
	return c.Delete(ProjectMilestoneModel, ids)
}

// GetProjectMilestone gets project.milestone existing record.
func (c *Client) GetProjectMilestone(id int64) (*ProjectMilestone, error) {
	pms, err := c.GetProjectMilestones([]int64{id})
	if err != nil {
		return nil, err
	}
	if pms != nil && len(*pms) > 0 {
		return &((*pms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.milestone not found", id)
}

// GetProjectMilestones gets project.milestone existing records.
func (c *Client) GetProjectMilestones(ids []int64) (*ProjectMilestones, error) {
	pms := &ProjectMilestones{}
	if err := c.Read(ProjectMilestoneModel, ids, nil, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindProjectMilestone finds project.milestone record by querying it with criteria.
func (c *Client) FindProjectMilestone(criteria *Criteria) (*ProjectMilestone, error) {
	pms := &ProjectMilestones{}
	if err := c.SearchRead(ProjectMilestoneModel, criteria, NewOptions().Limit(1), pms); err != nil {
		return nil, err
	}
	if pms != nil && len(*pms) > 0 {
		return &((*pms)[0]), nil
	}
	return nil, fmt.Errorf("project.milestone was not found with criteria %v", criteria)
}

// FindProjectMilestones finds project.milestone records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectMilestones(criteria *Criteria, options *Options) (*ProjectMilestones, error) {
	pms := &ProjectMilestones{}
	if err := c.SearchRead(ProjectMilestoneModel, criteria, options, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindProjectMilestoneIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectMilestoneIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectMilestoneModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectMilestoneId finds record id by querying it with criteria.
func (c *Client) FindProjectMilestoneId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectMilestoneModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.milestone was not found with criteria %v and options %v", criteria, options)
}
