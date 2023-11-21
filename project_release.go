package odoo

import (
	"fmt"
)

// ProjectRelease represents project.release model.
type ProjectRelease struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	ProjectIds       *Relation `xmlrpc:"project_ids,omptempty"`
	ReleaseDate      *Time     `xmlrpc:"release_date,omptempty"`
	TransportIds     *Relation `xmlrpc:"transport_ids,omptempty"`
	TransportIdsNber *Int      `xmlrpc:"transport_ids_nber,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectReleases represents array of project.release model.
type ProjectReleases []ProjectRelease

// ProjectReleaseModel is the odoo model name.
const ProjectReleaseModel = "project.release"

// Many2One convert ProjectRelease to *Many2One.
func (pr *ProjectRelease) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreateProjectRelease creates a new project.release model and returns its id.
func (c *Client) CreateProjectRelease(pr *ProjectRelease) (int64, error) {
	ids, err := c.CreateProjectReleases([]*ProjectRelease{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectRelease creates a new project.release model and returns its id.
func (c *Client) CreateProjectReleases(prs []*ProjectRelease) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(ProjectReleaseModel, vv)
}

// UpdateProjectRelease updates an existing project.release record.
func (c *Client) UpdateProjectRelease(pr *ProjectRelease) error {
	return c.UpdateProjectReleases([]int64{pr.Id.Get()}, pr)
}

// UpdateProjectReleases updates existing project.release records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProjectReleases(ids []int64, pr *ProjectRelease) error {
	return c.Update(ProjectReleaseModel, ids, pr)
}

// DeleteProjectRelease deletes an existing project.release record.
func (c *Client) DeleteProjectRelease(id int64) error {
	return c.DeleteProjectReleases([]int64{id})
}

// DeleteProjectReleases deletes existing project.release records.
func (c *Client) DeleteProjectReleases(ids []int64) error {
	return c.Delete(ProjectReleaseModel, ids)
}

// GetProjectRelease gets project.release existing record.
func (c *Client) GetProjectRelease(id int64) (*ProjectRelease, error) {
	prs, err := c.GetProjectReleases([]int64{id})
	if err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.release not found", id)
}

// GetProjectReleases gets project.release existing records.
func (c *Client) GetProjectReleases(ids []int64) (*ProjectReleases, error) {
	prs := &ProjectReleases{}
	if err := c.Read(ProjectReleaseModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProjectRelease finds project.release record by querying it with criteria.
func (c *Client) FindProjectRelease(criteria *Criteria) (*ProjectRelease, error) {
	prs := &ProjectReleases{}
	if err := c.SearchRead(ProjectReleaseModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("project.release was not found with criteria %v", criteria)
}

// FindProjectReleases finds project.release records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectReleases(criteria *Criteria, options *Options) (*ProjectReleases, error) {
	prs := &ProjectReleases{}
	if err := c.SearchRead(ProjectReleaseModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProjectReleaseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectReleaseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectReleaseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectReleaseId finds record id by querying it with criteria.
func (c *Client) FindProjectReleaseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectReleaseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.release was not found with criteria %v and options %v", criteria, options)
}
