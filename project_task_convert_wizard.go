package odoo

import (
	"fmt"
)

// ProjectTaskConvertWizard represents project.task.convert.wizard model.
type ProjectTaskConvertWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	StageId     *Many2One `xmlrpc:"stage_id,omptempty"`
	TeamId      *Many2One `xmlrpc:"team_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTaskConvertWizards represents array of project.task.convert.wizard model.
type ProjectTaskConvertWizards []ProjectTaskConvertWizard

// ProjectTaskConvertWizardModel is the odoo model name.
const ProjectTaskConvertWizardModel = "project.task.convert.wizard"

// Many2One convert ProjectTaskConvertWizard to *Many2One.
func (ptcw *ProjectTaskConvertWizard) Many2One() *Many2One {
	return NewMany2One(ptcw.Id.Get(), "")
}

// CreateProjectTaskConvertWizard creates a new project.task.convert.wizard model and returns its id.
func (c *Client) CreateProjectTaskConvertWizard(ptcw *ProjectTaskConvertWizard) (int64, error) {
	ids, err := c.CreateProjectTaskConvertWizards([]*ProjectTaskConvertWizard{ptcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskConvertWizard creates a new project.task.convert.wizard model and returns its id.
func (c *Client) CreateProjectTaskConvertWizards(ptcws []*ProjectTaskConvertWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptcws {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskConvertWizardModel, vv)
}

// UpdateProjectTaskConvertWizard updates an existing project.task.convert.wizard record.
func (c *Client) UpdateProjectTaskConvertWizard(ptcw *ProjectTaskConvertWizard) error {
	return c.UpdateProjectTaskConvertWizards([]int64{ptcw.Id.Get()}, ptcw)
}

// UpdateProjectTaskConvertWizards updates existing project.task.convert.wizard records.
// All records (represented by ids) will be updated by ptcw values.
func (c *Client) UpdateProjectTaskConvertWizards(ids []int64, ptcw *ProjectTaskConvertWizard) error {
	return c.Update(ProjectTaskConvertWizardModel, ids, ptcw)
}

// DeleteProjectTaskConvertWizard deletes an existing project.task.convert.wizard record.
func (c *Client) DeleteProjectTaskConvertWizard(id int64) error {
	return c.DeleteProjectTaskConvertWizards([]int64{id})
}

// DeleteProjectTaskConvertWizards deletes existing project.task.convert.wizard records.
func (c *Client) DeleteProjectTaskConvertWizards(ids []int64) error {
	return c.Delete(ProjectTaskConvertWizardModel, ids)
}

// GetProjectTaskConvertWizard gets project.task.convert.wizard existing record.
func (c *Client) GetProjectTaskConvertWizard(id int64) (*ProjectTaskConvertWizard, error) {
	ptcws, err := c.GetProjectTaskConvertWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptcws != nil && len(*ptcws) > 0 {
		return &((*ptcws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task.convert.wizard not found", id)
}

// GetProjectTaskConvertWizards gets project.task.convert.wizard existing records.
func (c *Client) GetProjectTaskConvertWizards(ids []int64) (*ProjectTaskConvertWizards, error) {
	ptcws := &ProjectTaskConvertWizards{}
	if err := c.Read(ProjectTaskConvertWizardModel, ids, nil, ptcws); err != nil {
		return nil, err
	}
	return ptcws, nil
}

// FindProjectTaskConvertWizard finds project.task.convert.wizard record by querying it with criteria.
func (c *Client) FindProjectTaskConvertWizard(criteria *Criteria) (*ProjectTaskConvertWizard, error) {
	ptcws := &ProjectTaskConvertWizards{}
	if err := c.SearchRead(ProjectTaskConvertWizardModel, criteria, NewOptions().Limit(1), ptcws); err != nil {
		return nil, err
	}
	if ptcws != nil && len(*ptcws) > 0 {
		return &((*ptcws)[0]), nil
	}
	return nil, fmt.Errorf("project.task.convert.wizard was not found with criteria %v", criteria)
}

// FindProjectTaskConvertWizards finds project.task.convert.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskConvertWizards(criteria *Criteria, options *Options) (*ProjectTaskConvertWizards, error) {
	ptcws := &ProjectTaskConvertWizards{}
	if err := c.SearchRead(ProjectTaskConvertWizardModel, criteria, options, ptcws); err != nil {
		return nil, err
	}
	return ptcws, nil
}

// FindProjectTaskConvertWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskConvertWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskConvertWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskConvertWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskConvertWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskConvertWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task.convert.wizard was not found with criteria %v and options %v", criteria, options)
}
