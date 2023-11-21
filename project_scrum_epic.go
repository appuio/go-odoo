package odoo

import (
	"fmt"
)

// ProjectScrumEpic represents project.scrum.epic model.
type ProjectScrumEpic struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	AvailableForAllProjects  *Bool     `xmlrpc:"available_for_all_projects,omptempty"`
	AvailableProjectIds      *Relation `xmlrpc:"available_project_ids,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
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
	TopicIds                 *Relation `xmlrpc:"topic_ids,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectScrumEpics represents array of project.scrum.epic model.
type ProjectScrumEpics []ProjectScrumEpic

// ProjectScrumEpicModel is the odoo model name.
const ProjectScrumEpicModel = "project.scrum.epic"

// Many2One convert ProjectScrumEpic to *Many2One.
func (pse *ProjectScrumEpic) Many2One() *Many2One {
	return NewMany2One(pse.Id.Get(), "")
}

// CreateProjectScrumEpic creates a new project.scrum.epic model and returns its id.
func (c *Client) CreateProjectScrumEpic(pse *ProjectScrumEpic) (int64, error) {
	ids, err := c.CreateProjectScrumEpics([]*ProjectScrumEpic{pse})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectScrumEpic creates a new project.scrum.epic model and returns its id.
func (c *Client) CreateProjectScrumEpics(pses []*ProjectScrumEpic) ([]int64, error) {
	var vv []interface{}
	for _, v := range pses {
		vv = append(vv, v)
	}
	return c.Create(ProjectScrumEpicModel, vv)
}

// UpdateProjectScrumEpic updates an existing project.scrum.epic record.
func (c *Client) UpdateProjectScrumEpic(pse *ProjectScrumEpic) error {
	return c.UpdateProjectScrumEpics([]int64{pse.Id.Get()}, pse)
}

// UpdateProjectScrumEpics updates existing project.scrum.epic records.
// All records (represented by ids) will be updated by pse values.
func (c *Client) UpdateProjectScrumEpics(ids []int64, pse *ProjectScrumEpic) error {
	return c.Update(ProjectScrumEpicModel, ids, pse)
}

// DeleteProjectScrumEpic deletes an existing project.scrum.epic record.
func (c *Client) DeleteProjectScrumEpic(id int64) error {
	return c.DeleteProjectScrumEpics([]int64{id})
}

// DeleteProjectScrumEpics deletes existing project.scrum.epic records.
func (c *Client) DeleteProjectScrumEpics(ids []int64) error {
	return c.Delete(ProjectScrumEpicModel, ids)
}

// GetProjectScrumEpic gets project.scrum.epic existing record.
func (c *Client) GetProjectScrumEpic(id int64) (*ProjectScrumEpic, error) {
	pses, err := c.GetProjectScrumEpics([]int64{id})
	if err != nil {
		return nil, err
	}
	if pses != nil && len(*pses) > 0 {
		return &((*pses)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.scrum.epic not found", id)
}

// GetProjectScrumEpics gets project.scrum.epic existing records.
func (c *Client) GetProjectScrumEpics(ids []int64) (*ProjectScrumEpics, error) {
	pses := &ProjectScrumEpics{}
	if err := c.Read(ProjectScrumEpicModel, ids, nil, pses); err != nil {
		return nil, err
	}
	return pses, nil
}

// FindProjectScrumEpic finds project.scrum.epic record by querying it with criteria.
func (c *Client) FindProjectScrumEpic(criteria *Criteria) (*ProjectScrumEpic, error) {
	pses := &ProjectScrumEpics{}
	if err := c.SearchRead(ProjectScrumEpicModel, criteria, NewOptions().Limit(1), pses); err != nil {
		return nil, err
	}
	if pses != nil && len(*pses) > 0 {
		return &((*pses)[0]), nil
	}
	return nil, fmt.Errorf("project.scrum.epic was not found with criteria %v", criteria)
}

// FindProjectScrumEpics finds project.scrum.epic records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumEpics(criteria *Criteria, options *Options) (*ProjectScrumEpics, error) {
	pses := &ProjectScrumEpics{}
	if err := c.SearchRead(ProjectScrumEpicModel, criteria, options, pses); err != nil {
		return nil, err
	}
	return pses, nil
}

// FindProjectScrumEpicIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumEpicIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectScrumEpicModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectScrumEpicId finds record id by querying it with criteria.
func (c *Client) FindProjectScrumEpicId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectScrumEpicModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.scrum.epic was not found with criteria %v and options %v", criteria, options)
}
