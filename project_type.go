package odoo

import (
	"fmt"
)

// ProjectType represents project.type model.
type ProjectType struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	ChildIds     *Relation `xmlrpc:"child_ids,omptempty"`
	Code         *String   `xmlrpc:"code,omptempty"`
	CompleteName *String   `xmlrpc:"complete_name,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	Description  *String   `xmlrpc:"description,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	ParentId     *Many2One `xmlrpc:"parent_id,omptempty"`
	ProjectOk    *Bool     `xmlrpc:"project_ok,omptempty"`
	TaskOk       *Bool     `xmlrpc:"task_ok,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTypes represents array of project.type model.
type ProjectTypes []ProjectType

// ProjectTypeModel is the odoo model name.
const ProjectTypeModel = "project.type"

// Many2One convert ProjectType to *Many2One.
func (pt *ProjectType) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProjectType creates a new project.type model and returns its id.
func (c *Client) CreateProjectType(pt *ProjectType) (int64, error) {
	ids, err := c.CreateProjectTypes([]*ProjectType{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectType creates a new project.type model and returns its id.
func (c *Client) CreateProjectTypes(pts []*ProjectType) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTypeModel, vv)
}

// UpdateProjectType updates an existing project.type record.
func (c *Client) UpdateProjectType(pt *ProjectType) error {
	return c.UpdateProjectTypes([]int64{pt.Id.Get()}, pt)
}

// UpdateProjectTypes updates existing project.type records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProjectTypes(ids []int64, pt *ProjectType) error {
	return c.Update(ProjectTypeModel, ids, pt)
}

// DeleteProjectType deletes an existing project.type record.
func (c *Client) DeleteProjectType(id int64) error {
	return c.DeleteProjectTypes([]int64{id})
}

// DeleteProjectTypes deletes existing project.type records.
func (c *Client) DeleteProjectTypes(ids []int64) error {
	return c.Delete(ProjectTypeModel, ids)
}

// GetProjectType gets project.type existing record.
func (c *Client) GetProjectType(id int64) (*ProjectType, error) {
	pts, err := c.GetProjectTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.type not found", id)
}

// GetProjectTypes gets project.type existing records.
func (c *Client) GetProjectTypes(ids []int64) (*ProjectTypes, error) {
	pts := &ProjectTypes{}
	if err := c.Read(ProjectTypeModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectType finds project.type record by querying it with criteria.
func (c *Client) FindProjectType(criteria *Criteria) (*ProjectType, error) {
	pts := &ProjectTypes{}
	if err := c.SearchRead(ProjectTypeModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("project.type was not found with criteria %v", criteria)
}

// FindProjectTypes finds project.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTypes(criteria *Criteria, options *Options) (*ProjectTypes, error) {
	pts := &ProjectTypes{}
	if err := c.SearchRead(ProjectTypeModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTypeId finds record id by querying it with criteria.
func (c *Client) FindProjectTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.type was not found with criteria %v and options %v", criteria, options)
}
