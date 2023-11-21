package odoo

import (
	"fmt"
)

// ProjectTaskCreateTimesheet represents project.task.create.timesheet model.
type ProjectTaskCreateTimesheet struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Description *String   `xmlrpc:"description,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	TaskId      *Many2One `xmlrpc:"task_id,omptempty"`
	TimeSpent   *Float    `xmlrpc:"time_spent,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTaskCreateTimesheets represents array of project.task.create.timesheet model.
type ProjectTaskCreateTimesheets []ProjectTaskCreateTimesheet

// ProjectTaskCreateTimesheetModel is the odoo model name.
const ProjectTaskCreateTimesheetModel = "project.task.create.timesheet"

// Many2One convert ProjectTaskCreateTimesheet to *Many2One.
func (ptct *ProjectTaskCreateTimesheet) Many2One() *Many2One {
	return NewMany2One(ptct.Id.Get(), "")
}

// CreateProjectTaskCreateTimesheet creates a new project.task.create.timesheet model and returns its id.
func (c *Client) CreateProjectTaskCreateTimesheet(ptct *ProjectTaskCreateTimesheet) (int64, error) {
	ids, err := c.CreateProjectTaskCreateTimesheets([]*ProjectTaskCreateTimesheet{ptct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskCreateTimesheet creates a new project.task.create.timesheet model and returns its id.
func (c *Client) CreateProjectTaskCreateTimesheets(ptcts []*ProjectTaskCreateTimesheet) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptcts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskCreateTimesheetModel, vv)
}

// UpdateProjectTaskCreateTimesheet updates an existing project.task.create.timesheet record.
func (c *Client) UpdateProjectTaskCreateTimesheet(ptct *ProjectTaskCreateTimesheet) error {
	return c.UpdateProjectTaskCreateTimesheets([]int64{ptct.Id.Get()}, ptct)
}

// UpdateProjectTaskCreateTimesheets updates existing project.task.create.timesheet records.
// All records (represented by ids) will be updated by ptct values.
func (c *Client) UpdateProjectTaskCreateTimesheets(ids []int64, ptct *ProjectTaskCreateTimesheet) error {
	return c.Update(ProjectTaskCreateTimesheetModel, ids, ptct)
}

// DeleteProjectTaskCreateTimesheet deletes an existing project.task.create.timesheet record.
func (c *Client) DeleteProjectTaskCreateTimesheet(id int64) error {
	return c.DeleteProjectTaskCreateTimesheets([]int64{id})
}

// DeleteProjectTaskCreateTimesheets deletes existing project.task.create.timesheet records.
func (c *Client) DeleteProjectTaskCreateTimesheets(ids []int64) error {
	return c.Delete(ProjectTaskCreateTimesheetModel, ids)
}

// GetProjectTaskCreateTimesheet gets project.task.create.timesheet existing record.
func (c *Client) GetProjectTaskCreateTimesheet(id int64) (*ProjectTaskCreateTimesheet, error) {
	ptcts, err := c.GetProjectTaskCreateTimesheets([]int64{id})
	if err != nil {
		return nil, err
	}
	if ptcts != nil && len(*ptcts) > 0 {
		return &((*ptcts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task.create.timesheet not found", id)
}

// GetProjectTaskCreateTimesheets gets project.task.create.timesheet existing records.
func (c *Client) GetProjectTaskCreateTimesheets(ids []int64) (*ProjectTaskCreateTimesheets, error) {
	ptcts := &ProjectTaskCreateTimesheets{}
	if err := c.Read(ProjectTaskCreateTimesheetModel, ids, nil, ptcts); err != nil {
		return nil, err
	}
	return ptcts, nil
}

// FindProjectTaskCreateTimesheet finds project.task.create.timesheet record by querying it with criteria.
func (c *Client) FindProjectTaskCreateTimesheet(criteria *Criteria) (*ProjectTaskCreateTimesheet, error) {
	ptcts := &ProjectTaskCreateTimesheets{}
	if err := c.SearchRead(ProjectTaskCreateTimesheetModel, criteria, NewOptions().Limit(1), ptcts); err != nil {
		return nil, err
	}
	if ptcts != nil && len(*ptcts) > 0 {
		return &((*ptcts)[0]), nil
	}
	return nil, fmt.Errorf("project.task.create.timesheet was not found with criteria %v", criteria)
}

// FindProjectTaskCreateTimesheets finds project.task.create.timesheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskCreateTimesheets(criteria *Criteria, options *Options) (*ProjectTaskCreateTimesheets, error) {
	ptcts := &ProjectTaskCreateTimesheets{}
	if err := c.SearchRead(ProjectTaskCreateTimesheetModel, criteria, options, ptcts); err != nil {
		return nil, err
	}
	return ptcts, nil
}

// FindProjectTaskCreateTimesheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskCreateTimesheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskCreateTimesheetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskCreateTimesheetId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskCreateTimesheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskCreateTimesheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task.create.timesheet was not found with criteria %v and options %v", criteria, options)
}
