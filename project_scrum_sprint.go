package odoo

import (
	"fmt"
)

// ProjectScrumSprint represents project.scrum.sprint model.
type ProjectScrumSprint struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	AvailableForAllProjects  *Bool     `xmlrpc:"available_for_all_projects,omptempty"`
	AvailableProjectIds      *Relation `xmlrpc:"available_project_ids,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DateEnd                  *Time     `xmlrpc:"date_end,omptempty"`
	DateStart                *Time     `xmlrpc:"date_start,omptempty"`
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
	Sequence                 *Int      `xmlrpc:"sequence,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectScrumSprints represents array of project.scrum.sprint model.
type ProjectScrumSprints []ProjectScrumSprint

// ProjectScrumSprintModel is the odoo model name.
const ProjectScrumSprintModel = "project.scrum.sprint"

// Many2One convert ProjectScrumSprint to *Many2One.
func (pss *ProjectScrumSprint) Many2One() *Many2One {
	return NewMany2One(pss.Id.Get(), "")
}

// CreateProjectScrumSprint creates a new project.scrum.sprint model and returns its id.
func (c *Client) CreateProjectScrumSprint(pss *ProjectScrumSprint) (int64, error) {
	ids, err := c.CreateProjectScrumSprints([]*ProjectScrumSprint{pss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectScrumSprint creates a new project.scrum.sprint model and returns its id.
func (c *Client) CreateProjectScrumSprints(psss []*ProjectScrumSprint) ([]int64, error) {
	var vv []interface{}
	for _, v := range psss {
		vv = append(vv, v)
	}
	return c.Create(ProjectScrumSprintModel, vv)
}

// UpdateProjectScrumSprint updates an existing project.scrum.sprint record.
func (c *Client) UpdateProjectScrumSprint(pss *ProjectScrumSprint) error {
	return c.UpdateProjectScrumSprints([]int64{pss.Id.Get()}, pss)
}

// UpdateProjectScrumSprints updates existing project.scrum.sprint records.
// All records (represented by ids) will be updated by pss values.
func (c *Client) UpdateProjectScrumSprints(ids []int64, pss *ProjectScrumSprint) error {
	return c.Update(ProjectScrumSprintModel, ids, pss)
}

// DeleteProjectScrumSprint deletes an existing project.scrum.sprint record.
func (c *Client) DeleteProjectScrumSprint(id int64) error {
	return c.DeleteProjectScrumSprints([]int64{id})
}

// DeleteProjectScrumSprints deletes existing project.scrum.sprint records.
func (c *Client) DeleteProjectScrumSprints(ids []int64) error {
	return c.Delete(ProjectScrumSprintModel, ids)
}

// GetProjectScrumSprint gets project.scrum.sprint existing record.
func (c *Client) GetProjectScrumSprint(id int64) (*ProjectScrumSprint, error) {
	psss, err := c.GetProjectScrumSprints([]int64{id})
	if err != nil {
		return nil, err
	}
	if psss != nil && len(*psss) > 0 {
		return &((*psss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.scrum.sprint not found", id)
}

// GetProjectScrumSprints gets project.scrum.sprint existing records.
func (c *Client) GetProjectScrumSprints(ids []int64) (*ProjectScrumSprints, error) {
	psss := &ProjectScrumSprints{}
	if err := c.Read(ProjectScrumSprintModel, ids, nil, psss); err != nil {
		return nil, err
	}
	return psss, nil
}

// FindProjectScrumSprint finds project.scrum.sprint record by querying it with criteria.
func (c *Client) FindProjectScrumSprint(criteria *Criteria) (*ProjectScrumSprint, error) {
	psss := &ProjectScrumSprints{}
	if err := c.SearchRead(ProjectScrumSprintModel, criteria, NewOptions().Limit(1), psss); err != nil {
		return nil, err
	}
	if psss != nil && len(*psss) > 0 {
		return &((*psss)[0]), nil
	}
	return nil, fmt.Errorf("project.scrum.sprint was not found with criteria %v", criteria)
}

// FindProjectScrumSprints finds project.scrum.sprint records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumSprints(criteria *Criteria, options *Options) (*ProjectScrumSprints, error) {
	psss := &ProjectScrumSprints{}
	if err := c.SearchRead(ProjectScrumSprintModel, criteria, options, psss); err != nil {
		return nil, err
	}
	return psss, nil
}

// FindProjectScrumSprintIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectScrumSprintIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectScrumSprintModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectScrumSprintId finds record id by querying it with criteria.
func (c *Client) FindProjectScrumSprintId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectScrumSprintModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.scrum.sprint was not found with criteria %v and options %v", criteria, options)
}
