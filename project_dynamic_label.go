package odoo

import (
	"fmt"
)

// ProjectDynamicLabel represents project.dynamic.label model.
type ProjectDynamicLabel struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	AvailableProjectIds          *Relation  `xmlrpc:"available_project_ids,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	DynamicValueIds              *Relation  `xmlrpc:"dynamic_value_ids,omptempty"`
	FirstAttProjectIds           *Relation  `xmlrpc:"first_att_project_ids,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	LabelType                    *Selection `xmlrpc:"label_type,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	SecondAttProjectIds          *Relation  `xmlrpc:"second_att_project_ids,omptempty"`
	TransportFirstAttProjectIds  *Relation  `xmlrpc:"transport_first_att_project_ids,omptempty"`
	TransportSecondAttProjectIds *Relation  `xmlrpc:"transport_second_att_project_ids,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProjectDynamicLabels represents array of project.dynamic.label model.
type ProjectDynamicLabels []ProjectDynamicLabel

// ProjectDynamicLabelModel is the odoo model name.
const ProjectDynamicLabelModel = "project.dynamic.label"

// Many2One convert ProjectDynamicLabel to *Many2One.
func (pdl *ProjectDynamicLabel) Many2One() *Many2One {
	return NewMany2One(pdl.Id.Get(), "")
}

// CreateProjectDynamicLabel creates a new project.dynamic.label model and returns its id.
func (c *Client) CreateProjectDynamicLabel(pdl *ProjectDynamicLabel) (int64, error) {
	ids, err := c.CreateProjectDynamicLabels([]*ProjectDynamicLabel{pdl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectDynamicLabel creates a new project.dynamic.label model and returns its id.
func (c *Client) CreateProjectDynamicLabels(pdls []*ProjectDynamicLabel) ([]int64, error) {
	var vv []interface{}
	for _, v := range pdls {
		vv = append(vv, v)
	}
	return c.Create(ProjectDynamicLabelModel, vv)
}

// UpdateProjectDynamicLabel updates an existing project.dynamic.label record.
func (c *Client) UpdateProjectDynamicLabel(pdl *ProjectDynamicLabel) error {
	return c.UpdateProjectDynamicLabels([]int64{pdl.Id.Get()}, pdl)
}

// UpdateProjectDynamicLabels updates existing project.dynamic.label records.
// All records (represented by ids) will be updated by pdl values.
func (c *Client) UpdateProjectDynamicLabels(ids []int64, pdl *ProjectDynamicLabel) error {
	return c.Update(ProjectDynamicLabelModel, ids, pdl)
}

// DeleteProjectDynamicLabel deletes an existing project.dynamic.label record.
func (c *Client) DeleteProjectDynamicLabel(id int64) error {
	return c.DeleteProjectDynamicLabels([]int64{id})
}

// DeleteProjectDynamicLabels deletes existing project.dynamic.label records.
func (c *Client) DeleteProjectDynamicLabels(ids []int64) error {
	return c.Delete(ProjectDynamicLabelModel, ids)
}

// GetProjectDynamicLabel gets project.dynamic.label existing record.
func (c *Client) GetProjectDynamicLabel(id int64) (*ProjectDynamicLabel, error) {
	pdls, err := c.GetProjectDynamicLabels([]int64{id})
	if err != nil {
		return nil, err
	}
	if pdls != nil && len(*pdls) > 0 {
		return &((*pdls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.dynamic.label not found", id)
}

// GetProjectDynamicLabels gets project.dynamic.label existing records.
func (c *Client) GetProjectDynamicLabels(ids []int64) (*ProjectDynamicLabels, error) {
	pdls := &ProjectDynamicLabels{}
	if err := c.Read(ProjectDynamicLabelModel, ids, nil, pdls); err != nil {
		return nil, err
	}
	return pdls, nil
}

// FindProjectDynamicLabel finds project.dynamic.label record by querying it with criteria.
func (c *Client) FindProjectDynamicLabel(criteria *Criteria) (*ProjectDynamicLabel, error) {
	pdls := &ProjectDynamicLabels{}
	if err := c.SearchRead(ProjectDynamicLabelModel, criteria, NewOptions().Limit(1), pdls); err != nil {
		return nil, err
	}
	if pdls != nil && len(*pdls) > 0 {
		return &((*pdls)[0]), nil
	}
	return nil, fmt.Errorf("project.dynamic.label was not found with criteria %v", criteria)
}

// FindProjectDynamicLabels finds project.dynamic.label records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectDynamicLabels(criteria *Criteria, options *Options) (*ProjectDynamicLabels, error) {
	pdls := &ProjectDynamicLabels{}
	if err := c.SearchRead(ProjectDynamicLabelModel, criteria, options, pdls); err != nil {
		return nil, err
	}
	return pdls, nil
}

// FindProjectDynamicLabelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectDynamicLabelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectDynamicLabelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectDynamicLabelId finds record id by querying it with criteria.
func (c *Client) FindProjectDynamicLabelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectDynamicLabelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.dynamic.label was not found with criteria %v and options %v", criteria, options)
}
