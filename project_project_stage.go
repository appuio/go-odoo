package odoo

import (
	"fmt"
)

// ProjectProjectStage represents project.project.stage model.
type ProjectProjectStage struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	Active         *Bool     `xmlrpc:"active,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Fold           *Bool     `xmlrpc:"fold,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	MailTemplateId *Many2One `xmlrpc:"mail_template_id,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	Sequence       *Int      `xmlrpc:"sequence,omptempty"`
	SmsTemplateId  *Many2One `xmlrpc:"sms_template_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectProjectStages represents array of project.project.stage model.
type ProjectProjectStages []ProjectProjectStage

// ProjectProjectStageModel is the odoo model name.
const ProjectProjectStageModel = "project.project.stage"

// Many2One convert ProjectProjectStage to *Many2One.
func (pps *ProjectProjectStage) Many2One() *Many2One {
	return NewMany2One(pps.Id.Get(), "")
}

// CreateProjectProjectStage creates a new project.project.stage model and returns its id.
func (c *Client) CreateProjectProjectStage(pps *ProjectProjectStage) (int64, error) {
	ids, err := c.CreateProjectProjectStages([]*ProjectProjectStage{pps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectProjectStage creates a new project.project.stage model and returns its id.
func (c *Client) CreateProjectProjectStages(ppss []*ProjectProjectStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppss {
		vv = append(vv, v)
	}
	return c.Create(ProjectProjectStageModel, vv)
}

// UpdateProjectProjectStage updates an existing project.project.stage record.
func (c *Client) UpdateProjectProjectStage(pps *ProjectProjectStage) error {
	return c.UpdateProjectProjectStages([]int64{pps.Id.Get()}, pps)
}

// UpdateProjectProjectStages updates existing project.project.stage records.
// All records (represented by ids) will be updated by pps values.
func (c *Client) UpdateProjectProjectStages(ids []int64, pps *ProjectProjectStage) error {
	return c.Update(ProjectProjectStageModel, ids, pps)
}

// DeleteProjectProjectStage deletes an existing project.project.stage record.
func (c *Client) DeleteProjectProjectStage(id int64) error {
	return c.DeleteProjectProjectStages([]int64{id})
}

// DeleteProjectProjectStages deletes existing project.project.stage records.
func (c *Client) DeleteProjectProjectStages(ids []int64) error {
	return c.Delete(ProjectProjectStageModel, ids)
}

// GetProjectProjectStage gets project.project.stage existing record.
func (c *Client) GetProjectProjectStage(id int64) (*ProjectProjectStage, error) {
	ppss, err := c.GetProjectProjectStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if ppss != nil && len(*ppss) > 0 {
		return &((*ppss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.project.stage not found", id)
}

// GetProjectProjectStages gets project.project.stage existing records.
func (c *Client) GetProjectProjectStages(ids []int64) (*ProjectProjectStages, error) {
	ppss := &ProjectProjectStages{}
	if err := c.Read(ProjectProjectStageModel, ids, nil, ppss); err != nil {
		return nil, err
	}
	return ppss, nil
}

// FindProjectProjectStage finds project.project.stage record by querying it with criteria.
func (c *Client) FindProjectProjectStage(criteria *Criteria) (*ProjectProjectStage, error) {
	ppss := &ProjectProjectStages{}
	if err := c.SearchRead(ProjectProjectStageModel, criteria, NewOptions().Limit(1), ppss); err != nil {
		return nil, err
	}
	if ppss != nil && len(*ppss) > 0 {
		return &((*ppss)[0]), nil
	}
	return nil, fmt.Errorf("project.project.stage was not found with criteria %v", criteria)
}

// FindProjectProjectStages finds project.project.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectStages(criteria *Criteria, options *Options) (*ProjectProjectStages, error) {
	ppss := &ProjectProjectStages{}
	if err := c.SearchRead(ProjectProjectStageModel, criteria, options, ppss); err != nil {
		return nil, err
	}
	return ppss, nil
}

// FindProjectProjectStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectProjectStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectProjectStageId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.project.stage was not found with criteria %v and options %v", criteria, options)
}
