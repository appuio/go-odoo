package odoo

import (
	"fmt"
)

// ProjectDynamicValue represents project.dynamic.value model.
type ProjectDynamicValue struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	LabelId     *Many2One  `xmlrpc:"label_id,omptempty"`
	LabelType   *Selection `xmlrpc:"label_type,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProjectDynamicValues represents array of project.dynamic.value model.
type ProjectDynamicValues []ProjectDynamicValue

// ProjectDynamicValueModel is the odoo model name.
const ProjectDynamicValueModel = "project.dynamic.value"

// Many2One convert ProjectDynamicValue to *Many2One.
func (pdv *ProjectDynamicValue) Many2One() *Many2One {
	return NewMany2One(pdv.Id.Get(), "")
}

// CreateProjectDynamicValue creates a new project.dynamic.value model and returns its id.
func (c *Client) CreateProjectDynamicValue(pdv *ProjectDynamicValue) (int64, error) {
	ids, err := c.CreateProjectDynamicValues([]*ProjectDynamicValue{pdv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectDynamicValue creates a new project.dynamic.value model and returns its id.
func (c *Client) CreateProjectDynamicValues(pdvs []*ProjectDynamicValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range pdvs {
		vv = append(vv, v)
	}
	return c.Create(ProjectDynamicValueModel, vv)
}

// UpdateProjectDynamicValue updates an existing project.dynamic.value record.
func (c *Client) UpdateProjectDynamicValue(pdv *ProjectDynamicValue) error {
	return c.UpdateProjectDynamicValues([]int64{pdv.Id.Get()}, pdv)
}

// UpdateProjectDynamicValues updates existing project.dynamic.value records.
// All records (represented by ids) will be updated by pdv values.
func (c *Client) UpdateProjectDynamicValues(ids []int64, pdv *ProjectDynamicValue) error {
	return c.Update(ProjectDynamicValueModel, ids, pdv)
}

// DeleteProjectDynamicValue deletes an existing project.dynamic.value record.
func (c *Client) DeleteProjectDynamicValue(id int64) error {
	return c.DeleteProjectDynamicValues([]int64{id})
}

// DeleteProjectDynamicValues deletes existing project.dynamic.value records.
func (c *Client) DeleteProjectDynamicValues(ids []int64) error {
	return c.Delete(ProjectDynamicValueModel, ids)
}

// GetProjectDynamicValue gets project.dynamic.value existing record.
func (c *Client) GetProjectDynamicValue(id int64) (*ProjectDynamicValue, error) {
	pdvs, err := c.GetProjectDynamicValues([]int64{id})
	if err != nil {
		return nil, err
	}
	if pdvs != nil && len(*pdvs) > 0 {
		return &((*pdvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.dynamic.value not found", id)
}

// GetProjectDynamicValues gets project.dynamic.value existing records.
func (c *Client) GetProjectDynamicValues(ids []int64) (*ProjectDynamicValues, error) {
	pdvs := &ProjectDynamicValues{}
	if err := c.Read(ProjectDynamicValueModel, ids, nil, pdvs); err != nil {
		return nil, err
	}
	return pdvs, nil
}

// FindProjectDynamicValue finds project.dynamic.value record by querying it with criteria.
func (c *Client) FindProjectDynamicValue(criteria *Criteria) (*ProjectDynamicValue, error) {
	pdvs := &ProjectDynamicValues{}
	if err := c.SearchRead(ProjectDynamicValueModel, criteria, NewOptions().Limit(1), pdvs); err != nil {
		return nil, err
	}
	if pdvs != nil && len(*pdvs) > 0 {
		return &((*pdvs)[0]), nil
	}
	return nil, fmt.Errorf("project.dynamic.value was not found with criteria %v", criteria)
}

// FindProjectDynamicValues finds project.dynamic.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectDynamicValues(criteria *Criteria, options *Options) (*ProjectDynamicValues, error) {
	pdvs := &ProjectDynamicValues{}
	if err := c.SearchRead(ProjectDynamicValueModel, criteria, options, pdvs); err != nil {
		return nil, err
	}
	return pdvs, nil
}

// FindProjectDynamicValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectDynamicValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectDynamicValueModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectDynamicValueId finds record id by querying it with criteria.
func (c *Client) FindProjectDynamicValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectDynamicValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.dynamic.value was not found with criteria %v and options %v", criteria, options)
}
