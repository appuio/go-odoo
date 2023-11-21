package odoo

import (
	"fmt"
)

// ProjectTaskConfirmScheduleLineWizard represents project.task.confirm.schedule.line.wizard model.
type ProjectTaskConfirmScheduleLineWizard struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DateBegin    *Time     `xmlrpc:"date_begin,omptempty"`
	DateEnd      *Time     `xmlrpc:"date_end,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	ParentId     *Many2One `xmlrpc:"parent_id,omptempty"`
	ScheduleTask *Bool     `xmlrpc:"schedule_task,omptempty"`
	TaskId       *Many2One `xmlrpc:"task_id,omptempty"`
	Warning      *String   `xmlrpc:"warning,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTaskConfirmScheduleLineWizards represents array of project.task.confirm.schedule.line.wizard model.
type ProjectTaskConfirmScheduleLineWizards []ProjectTaskConfirmScheduleLineWizard

// ProjectTaskConfirmScheduleLineWizardModel is the odoo model name.
const ProjectTaskConfirmScheduleLineWizardModel = "project.task.confirm.schedule.line.wizard"

// Many2One convert ProjectTaskConfirmScheduleLineWizard to *Many2One.
func (ptcslw *ProjectTaskConfirmScheduleLineWizard) Many2One() *Many2One {
	return NewMany2One(ptcslw.Id.Get(), "")
}

// CreateProjectTaskConfirmScheduleLineWizard creates a new project.task.confirm.schedule.line.wizard model and returns its id.
func (c *Client) CreateProjectTaskConfirmScheduleLineWizard(ptcslw *ProjectTaskConfirmScheduleLineWizard) (int64, error) {
	ids, err := c.CreateProjectTaskConfirmScheduleLineWizards([]*ProjectTaskConfirmScheduleLineWizard{ptcslw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskConfirmScheduleLineWizard creates a new project.task.confirm.schedule.line.wizard model and returns its id.
func (c *Client) CreateProjectTaskConfirmScheduleLineWizards(ptcslws []*ProjectTaskConfirmScheduleLineWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptcslws {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskConfirmScheduleLineWizardModel, vv)
}

// UpdateProjectTaskConfirmScheduleLineWizard updates an existing project.task.confirm.schedule.line.wizard record.
func (c *Client) UpdateProjectTaskConfirmScheduleLineWizard(ptcslw *ProjectTaskConfirmScheduleLineWizard) error {
	return c.UpdateProjectTaskConfirmScheduleLineWizards([]int64{ptcslw.Id.Get()}, ptcslw)
}

// UpdateProjectTaskConfirmScheduleLineWizards updates existing project.task.confirm.schedule.line.wizard records.
// All records (represented by ids) will be updated by ptcslw values.
func (c *Client) UpdateProjectTaskConfirmScheduleLineWizards(ids []int64, ptcslw *ProjectTaskConfirmScheduleLineWizard) error {
	return c.Update(ProjectTaskConfirmScheduleLineWizardModel, ids, ptcslw)
}

// DeleteProjectTaskConfirmScheduleLineWizard deletes an existing project.task.confirm.schedule.line.wizard record.
func (c *Client) DeleteProjectTaskConfirmScheduleLineWizard(id int64) error {
	return c.DeleteProjectTaskConfirmScheduleLineWizards([]int64{id})
}

// DeleteProjectTaskConfirmScheduleLineWizards deletes existing project.task.confirm.schedule.line.wizard records.
func (c *Client) DeleteProjectTaskConfirmScheduleLineWizards(ids []int64) error {
	return c.Delete(ProjectTaskConfirmScheduleLineWizardModel, ids)
}

// GetProjectTaskConfirmScheduleLineWizard gets project.task.confirm.schedule.line.wizard existing record.
func (c *Client) GetProjectTaskConfirmScheduleLineWizard(id int64) (*ProjectTaskConfirmScheduleLineWizard, error) {
	ptcslws, err := c.GetProjectTaskConfirmScheduleLineWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptcslws != nil && len(*ptcslws) > 0 {
		return &((*ptcslws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task.confirm.schedule.line.wizard not found", id)
}

// GetProjectTaskConfirmScheduleLineWizards gets project.task.confirm.schedule.line.wizard existing records.
func (c *Client) GetProjectTaskConfirmScheduleLineWizards(ids []int64) (*ProjectTaskConfirmScheduleLineWizards, error) {
	ptcslws := &ProjectTaskConfirmScheduleLineWizards{}
	if err := c.Read(ProjectTaskConfirmScheduleLineWizardModel, ids, nil, ptcslws); err != nil {
		return nil, err
	}
	return ptcslws, nil
}

// FindProjectTaskConfirmScheduleLineWizard finds project.task.confirm.schedule.line.wizard record by querying it with criteria.
func (c *Client) FindProjectTaskConfirmScheduleLineWizard(criteria *Criteria) (*ProjectTaskConfirmScheduleLineWizard, error) {
	ptcslws := &ProjectTaskConfirmScheduleLineWizards{}
	if err := c.SearchRead(ProjectTaskConfirmScheduleLineWizardModel, criteria, NewOptions().Limit(1), ptcslws); err != nil {
		return nil, err
	}
	if ptcslws != nil && len(*ptcslws) > 0 {
		return &((*ptcslws)[0]), nil
	}
	return nil, fmt.Errorf("project.task.confirm.schedule.line.wizard was not found with criteria %v", criteria)
}

// FindProjectTaskConfirmScheduleLineWizards finds project.task.confirm.schedule.line.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskConfirmScheduleLineWizards(criteria *Criteria, options *Options) (*ProjectTaskConfirmScheduleLineWizards, error) {
	ptcslws := &ProjectTaskConfirmScheduleLineWizards{}
	if err := c.SearchRead(ProjectTaskConfirmScheduleLineWizardModel, criteria, options, ptcslws); err != nil {
		return nil, err
	}
	return ptcslws, nil
}

// FindProjectTaskConfirmScheduleLineWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskConfirmScheduleLineWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskConfirmScheduleLineWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskConfirmScheduleLineWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskConfirmScheduleLineWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskConfirmScheduleLineWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task.confirm.schedule.line.wizard was not found with criteria %v and options %v", criteria, options)
}
