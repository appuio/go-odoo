package odoo

import (
	"fmt"
)

// ProjectTransport represents project.transport model.
type ProjectTransport struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	DynamicTransportLabel1Id *Many2One `xmlrpc:"dynamic_transport_label_1_id,omptempty"`
	DynamicTransportLabel2Id *Many2One `xmlrpc:"dynamic_transport_label_2_id,omptempty"`
	DynamicValue1Id          *Many2One `xmlrpc:"dynamic_value_1_id,omptempty"`
	DynamicValue2Id          *Many2One `xmlrpc:"dynamic_value_2_id,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsVisibleDynamicFields   *Bool     `xmlrpc:"is_visible_dynamic_fields,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	OwnerId                  *Many2One `xmlrpc:"owner_id,omptempty"`
	ProjectFromTaskId        *Many2One `xmlrpc:"project_from_task_id,omptempty"`
	ProjectId                *Many2One `xmlrpc:"project_id,omptempty"`
	ReleaseDate              *Time     `xmlrpc:"release_date,omptempty"`
	ReleaseId                *Many2One `xmlrpc:"release_id,omptempty"`
	ReleaseProjectIds        *Relation `xmlrpc:"release_project_ids,omptempty"`
	TaskIds                  *Relation `xmlrpc:"task_ids,omptempty"`
	TaskProjectIds           *Relation `xmlrpc:"task_project_ids,omptempty"`
	Today                    *Time     `xmlrpc:"today,omptempty"`
	TransportTitle           *String   `xmlrpc:"transport_title,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectTransports represents array of project.transport model.
type ProjectTransports []ProjectTransport

// ProjectTransportModel is the odoo model name.
const ProjectTransportModel = "project.transport"

// Many2One convert ProjectTransport to *Many2One.
func (pt *ProjectTransport) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProjectTransport creates a new project.transport model and returns its id.
func (c *Client) CreateProjectTransport(pt *ProjectTransport) (int64, error) {
	ids, err := c.CreateProjectTransports([]*ProjectTransport{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTransport creates a new project.transport model and returns its id.
func (c *Client) CreateProjectTransports(pts []*ProjectTransport) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTransportModel, vv)
}

// UpdateProjectTransport updates an existing project.transport record.
func (c *Client) UpdateProjectTransport(pt *ProjectTransport) error {
	return c.UpdateProjectTransports([]int64{pt.Id.Get()}, pt)
}

// UpdateProjectTransports updates existing project.transport records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProjectTransports(ids []int64, pt *ProjectTransport) error {
	return c.Update(ProjectTransportModel, ids, pt)
}

// DeleteProjectTransport deletes an existing project.transport record.
func (c *Client) DeleteProjectTransport(id int64) error {
	return c.DeleteProjectTransports([]int64{id})
}

// DeleteProjectTransports deletes existing project.transport records.
func (c *Client) DeleteProjectTransports(ids []int64) error {
	return c.Delete(ProjectTransportModel, ids)
}

// GetProjectTransport gets project.transport existing record.
func (c *Client) GetProjectTransport(id int64) (*ProjectTransport, error) {
	pts, err := c.GetProjectTransports([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.transport not found", id)
}

// GetProjectTransports gets project.transport existing records.
func (c *Client) GetProjectTransports(ids []int64) (*ProjectTransports, error) {
	pts := &ProjectTransports{}
	if err := c.Read(ProjectTransportModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTransport finds project.transport record by querying it with criteria.
func (c *Client) FindProjectTransport(criteria *Criteria) (*ProjectTransport, error) {
	pts := &ProjectTransports{}
	if err := c.SearchRead(ProjectTransportModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("project.transport was not found with criteria %v", criteria)
}

// FindProjectTransports finds project.transport records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTransports(criteria *Criteria, options *Options) (*ProjectTransports, error) {
	pts := &ProjectTransports{}
	if err := c.SearchRead(ProjectTransportModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTransportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTransportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTransportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTransportId finds record id by querying it with criteria.
func (c *Client) FindProjectTransportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTransportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.transport was not found with criteria %v and options %v", criteria, options)
}
