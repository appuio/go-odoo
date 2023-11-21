package odoo

import (
	"fmt"
)

// ProjectScrumTopic represents project.scrum.topic model.
type ProjectScrumTopic struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	AvailableForAllProjects  *Bool     `xmlrpc:"available_for_all_projects,omptempty"`
	AvailableProjectIds      *Relation `xmlrpc:"available_project_ids,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	EpicId                   *Many2One `xmlrpc:"epic_id,omptempty"`
	FailedMessageIds         *Relation `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
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
	Sequence                 *String   `xmlrpc:"sequence,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectScrumTopics represents array of project.scrum.topic model.
type ProjectScrumTopics []ProjectScrumTopic

// ProjectScrumTopicModel is the odoo model name.
const ProjectScrumTopicModel = "project.scrum.topic"

// Many2One convert ProjectScrumTopic to *Many2One.
func (pst *ProjectScrumTopic) Many2One() *Many2One {
	return NewMany2One(pst.Id.Get(), "")
}

// CreateProjectScrumTopic creates a new project.scrum.topic model and returns its id.
func (c *Client) CreateProjectScrumTopic(pst *ProjectScrumTopic) (int64, error) {
	ids, err := c.CreateProjectScrumTopics([]*ProjectScrumTopic{pst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectScrumTopic creates a new project.scrum.topic model and returns its id.
func (c *Client) CreateProjectScrumTopics(psts []*ProjectScrumTopic) ([]int64, error) {
	var vv []interface{}
	for _, v := range psts {
		vv = append(vv, v)
	}
	return c.Create(ProjectScrumTopicModel, vv)
}

// UpdateProjectScrumTopic updates an existing project.scrum.topic record.
func (c *Client) UpdateProjectScrumTopic(pst *ProjectScrumTopic) error {
	return c.UpdateProjectScrumTopics([]int64{pst.Id.Get()}, pst)
}

// UpdateProjectScrumTopics updates existing project.scrum.topic records.
// All records (represented by ids) will be updated by pst values.
func (c *Client) UpdateProjectScrumTopics(ids []int64, pst *ProjectScrumTopic) error {
	return c.Update(ProjectScrumTopicModel, ids, pst)
}

// DeleteProjectScrumTopic deletes an existing project.scrum.topic record.
func (c *Client) DeleteProjectScrumTopic(id int64) error {
	return c.DeleteProjectScrumTopics([]int64{id})
}

// DeleteProjectScrumTopics deletes existing project.scrum.topic records.
func (c *Client) DeleteProjectScrumTopics(ids []int64) error {
	return c.Delete(ProjectScrumTopicModel, ids)
}

// GetProjectScrumTopic gets project.scrum.topic existing record.
func (c *Client) GetProjectScrumTopic(id int64) (*ProjectScrumTopic, error) {
	psts, err := c.GetProjectScrumTopics([]int64{id})
	if err != nil {
		return nil, err
	}
	if psts != nil && len(*psts) > 0 {
		return &((*psts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.scrum.topic not found", id)
}

// GetProjectScrumTopics gets project.scrum.topic existing records.
func (c *Client) GetProjectScrumTopics(ids []int64) (*ProjectScrumTopics, error) {
	psts := &ProjectScrumTopics{}
	if err := c.Read(ProjectScrumTopicModel, ids, nil, psts); err != nil {
		return nil, err
	}
	return psts, nil
}

// FindProjectScrumTopic finds project.scrum.topic record by querying it with criteria.
func (c *Client) FindProjectScrumTopic(criteria *Criteria) (*ProjectScrumTopic, error) {
	psts := &ProjectScrumTopics{}
	if err := c.SearchRead(ProjectScrumTopicModel, criteria, NewOptions().Limit(1), psts); err != nil {
		return nil, err
	}
	if psts != nil && len(*psts) > 0 {
		return &((*psts)[0]), nil
	}
	return nil, fmt.Errorf("project.scrum.topic was not found with criteria %v", criteria)
}

// FindProjectScrumTopics finds project.scrum.topic records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumTopics(criteria *Criteria, options *Options) (*ProjectScrumTopics, error) {
	psts := &ProjectScrumTopics{}
	if err := c.SearchRead(ProjectScrumTopicModel, criteria, options, psts); err != nil {
		return nil, err
	}
	return psts, nil
}

// FindProjectScrumTopicIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumTopicIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectScrumTopicModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectScrumTopicId finds record id by querying it with criteria.
func (c *Client) FindProjectScrumTopicId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectScrumTopicModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.scrum.topic was not found with criteria %v and options %v", criteria, options)
}
