package odoo

import (
	"fmt"
)

// ProjectTaskConfirmScheduleWizard represents project.task.confirm.schedule.wizard model.
type ProjectTaskConfirmScheduleWizard struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	LineIds           *Relation `xmlrpc:"line_ids,omptempty"`
	SelectedLineCount *Int      `xmlrpc:"selected_line_count,omptempty"`
	ShowWarnings      *Bool     `xmlrpc:"show_warnings,omptempty"`
	UserId            *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTaskConfirmScheduleWizards represents array of project.task.confirm.schedule.wizard model.
type ProjectTaskConfirmScheduleWizards []ProjectTaskConfirmScheduleWizard

// ProjectTaskConfirmScheduleWizardModel is the odoo model name.
const ProjectTaskConfirmScheduleWizardModel = "project.task.confirm.schedule.wizard"

// Many2One convert ProjectTaskConfirmScheduleWizard to *Many2One.
func (ptcsw *ProjectTaskConfirmScheduleWizard) Many2One() *Many2One {
	return NewMany2One(ptcsw.Id.Get(), "")
}

// CreateProjectTaskConfirmScheduleWizard creates a new project.task.confirm.schedule.wizard model and returns its id.
func (c *Client) CreateProjectTaskConfirmScheduleWizard(ptcsw *ProjectTaskConfirmScheduleWizard) (int64, error) {
	ids, err := c.CreateProjectTaskConfirmScheduleWizards([]*ProjectTaskConfirmScheduleWizard{ptcsw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskConfirmScheduleWizard creates a new project.task.confirm.schedule.wizard model and returns its id.
func (c *Client) CreateProjectTaskConfirmScheduleWizards(ptcsws []*ProjectTaskConfirmScheduleWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptcsws {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskConfirmScheduleWizardModel, vv)
}

// UpdateProjectTaskConfirmScheduleWizard updates an existing project.task.confirm.schedule.wizard record.
func (c *Client) UpdateProjectTaskConfirmScheduleWizard(ptcsw *ProjectTaskConfirmScheduleWizard) error {
	return c.UpdateProjectTaskConfirmScheduleWizards([]int64{ptcsw.Id.Get()}, ptcsw)
}

// UpdateProjectTaskConfirmScheduleWizards updates existing project.task.confirm.schedule.wizard records.
// All records (represented by ids) will be updated by ptcsw values.
func (c *Client) UpdateProjectTaskConfirmScheduleWizards(ids []int64, ptcsw *ProjectTaskConfirmScheduleWizard) error {
	return c.Update(ProjectTaskConfirmScheduleWizardModel, ids, ptcsw)
}

// DeleteProjectTaskConfirmScheduleWizard deletes an existing project.task.confirm.schedule.wizard record.
func (c *Client) DeleteProjectTaskConfirmScheduleWizard(id int64) error {
	return c.DeleteProjectTaskConfirmScheduleWizards([]int64{id})
}

// DeleteProjectTaskConfirmScheduleWizards deletes existing project.task.confirm.schedule.wizard records.
func (c *Client) DeleteProjectTaskConfirmScheduleWizards(ids []int64) error {
	return c.Delete(ProjectTaskConfirmScheduleWizardModel, ids)
}

// GetProjectTaskConfirmScheduleWizard gets project.task.confirm.schedule.wizard existing record.
func (c *Client) GetProjectTaskConfirmScheduleWizard(id int64) (*ProjectTaskConfirmScheduleWizard, error) {
	ptcsws, err := c.GetProjectTaskConfirmScheduleWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptcsws != nil && len(*ptcsws) > 0 {
		return &((*ptcsws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task.confirm.schedule.wizard not found", id)
}

// GetProjectTaskConfirmScheduleWizards gets project.task.confirm.schedule.wizard existing records.
func (c *Client) GetProjectTaskConfirmScheduleWizards(ids []int64) (*ProjectTaskConfirmScheduleWizards, error) {
	ptcsws := &ProjectTaskConfirmScheduleWizards{}
	if err := c.Read(ProjectTaskConfirmScheduleWizardModel, ids, nil, ptcsws); err != nil {
		return nil, err
	}
	return ptcsws, nil
}

// FindProjectTaskConfirmScheduleWizard finds project.task.confirm.schedule.wizard record by querying it with criteria.
func (c *Client) FindProjectTaskConfirmScheduleWizard(criteria *Criteria) (*ProjectTaskConfirmScheduleWizard, error) {
	ptcsws := &ProjectTaskConfirmScheduleWizards{}
	if err := c.SearchRead(ProjectTaskConfirmScheduleWizardModel, criteria, NewOptions().Limit(1), ptcsws); err != nil {
		return nil, err
	}
	if ptcsws != nil && len(*ptcsws) > 0 {
		return &((*ptcsws)[0]), nil
	}
	return nil, fmt.Errorf("project.task.confirm.schedule.wizard was not found with criteria %v", criteria)
}

// FindProjectTaskConfirmScheduleWizards finds project.task.confirm.schedule.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskConfirmScheduleWizards(criteria *Criteria, options *Options) (*ProjectTaskConfirmScheduleWizards, error) {
	ptcsws := &ProjectTaskConfirmScheduleWizards{}
	if err := c.SearchRead(ProjectTaskConfirmScheduleWizardModel, criteria, options, ptcsws); err != nil {
		return nil, err
	}
	return ptcsws, nil
}

// FindProjectTaskConfirmScheduleWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskConfirmScheduleWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskConfirmScheduleWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskConfirmScheduleWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskConfirmScheduleWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskConfirmScheduleWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task.confirm.schedule.wizard was not found with criteria %v and options %v", criteria, options)
}
