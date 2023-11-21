package odoo

import (
	"fmt"
)

// ProjectScrumWorkstream represents project.scrum.workstream model.
type ProjectScrumWorkstream struct {
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
	SprintIds                *Relation `xmlrpc:"sprint_ids,omptempty"`
	TaskIds                  *Relation `xmlrpc:"task_ids,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectScrumWorkstreams represents array of project.scrum.workstream model.
type ProjectScrumWorkstreams []ProjectScrumWorkstream

// ProjectScrumWorkstreamModel is the odoo model name.
const ProjectScrumWorkstreamModel = "project.scrum.workstream"

// Many2One convert ProjectScrumWorkstream to *Many2One.
func (psw *ProjectScrumWorkstream) Many2One() *Many2One {
	return NewMany2One(psw.Id.Get(), "")
}

// CreateProjectScrumWorkstream creates a new project.scrum.workstream model and returns its id.
func (c *Client) CreateProjectScrumWorkstream(psw *ProjectScrumWorkstream) (int64, error) {
	ids, err := c.CreateProjectScrumWorkstreams([]*ProjectScrumWorkstream{psw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectScrumWorkstream creates a new project.scrum.workstream model and returns its id.
func (c *Client) CreateProjectScrumWorkstreams(psws []*ProjectScrumWorkstream) ([]int64, error) {
	var vv []interface{}
	for _, v := range psws {
		vv = append(vv, v)
	}
	return c.Create(ProjectScrumWorkstreamModel, vv)
}

// UpdateProjectScrumWorkstream updates an existing project.scrum.workstream record.
func (c *Client) UpdateProjectScrumWorkstream(psw *ProjectScrumWorkstream) error {
	return c.UpdateProjectScrumWorkstreams([]int64{psw.Id.Get()}, psw)
}

// UpdateProjectScrumWorkstreams updates existing project.scrum.workstream records.
// All records (represented by ids) will be updated by psw values.
func (c *Client) UpdateProjectScrumWorkstreams(ids []int64, psw *ProjectScrumWorkstream) error {
	return c.Update(ProjectScrumWorkstreamModel, ids, psw)
}

// DeleteProjectScrumWorkstream deletes an existing project.scrum.workstream record.
func (c *Client) DeleteProjectScrumWorkstream(id int64) error {
	return c.DeleteProjectScrumWorkstreams([]int64{id})
}

// DeleteProjectScrumWorkstreams deletes existing project.scrum.workstream records.
func (c *Client) DeleteProjectScrumWorkstreams(ids []int64) error {
	return c.Delete(ProjectScrumWorkstreamModel, ids)
}

// GetProjectScrumWorkstream gets project.scrum.workstream existing record.
func (c *Client) GetProjectScrumWorkstream(id int64) (*ProjectScrumWorkstream, error) {
	psws, err := c.GetProjectScrumWorkstreams([]int64{id})
	if err != nil {
		return nil, err
	}
	if psws != nil && len(*psws) > 0 {
		return &((*psws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.scrum.workstream not found", id)
}

// GetProjectScrumWorkstreams gets project.scrum.workstream existing records.
func (c *Client) GetProjectScrumWorkstreams(ids []int64) (*ProjectScrumWorkstreams, error) {
	psws := &ProjectScrumWorkstreams{}
	if err := c.Read(ProjectScrumWorkstreamModel, ids, nil, psws); err != nil {
		return nil, err
	}
	return psws, nil
}

// FindProjectScrumWorkstream finds project.scrum.workstream record by querying it with criteria.
func (c *Client) FindProjectScrumWorkstream(criteria *Criteria) (*ProjectScrumWorkstream, error) {
	psws := &ProjectScrumWorkstreams{}
	if err := c.SearchRead(ProjectScrumWorkstreamModel, criteria, NewOptions().Limit(1), psws); err != nil {
		return nil, err
	}
	if psws != nil && len(*psws) > 0 {
		return &((*psws)[0]), nil
	}
	return nil, fmt.Errorf("project.scrum.workstream was not found with criteria %v", criteria)
}

// FindProjectScrumWorkstreams finds project.scrum.workstream records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumWorkstreams(criteria *Criteria, options *Options) (*ProjectScrumWorkstreams, error) {
	psws := &ProjectScrumWorkstreams{}
	if err := c.SearchRead(ProjectScrumWorkstreamModel, criteria, options, psws); err != nil {
		return nil, err
	}
	return psws, nil
}

// FindProjectScrumWorkstreamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumWorkstreamIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectScrumWorkstreamModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectScrumWorkstreamId finds record id by querying it with criteria.
func (c *Client) FindProjectScrumWorkstreamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectScrumWorkstreamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.scrum.workstream was not found with criteria %v and options %v", criteria, options)
}
