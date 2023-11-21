package odoo

import (
	"fmt"
)

// PlanningSlotTemplate represents planning.slot.template model.
type PlanningSlotTemplate struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	Active       *Bool     `xmlrpc:"active,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Duration     *Float    `xmlrpc:"duration,omptempty"`
	DurationDays *Int      `xmlrpc:"duration_days,omptempty"`
	EndTime      *Float    `xmlrpc:"end_time,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	ProjectId    *Many2One `xmlrpc:"project_id,omptempty"`
	RoleId       *Many2One `xmlrpc:"role_id,omptempty"`
	Sequence     *Int      `xmlrpc:"sequence,omptempty"`
	StartTime    *Float    `xmlrpc:"start_time,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PlanningSlotTemplates represents array of planning.slot.template model.
type PlanningSlotTemplates []PlanningSlotTemplate

// PlanningSlotTemplateModel is the odoo model name.
const PlanningSlotTemplateModel = "planning.slot.template"

// Many2One convert PlanningSlotTemplate to *Many2One.
func (pst *PlanningSlotTemplate) Many2One() *Many2One {
	return NewMany2One(pst.Id.Get(), "")
}

// CreatePlanningSlotTemplate creates a new planning.slot.template model and returns its id.
func (c *Client) CreatePlanningSlotTemplate(pst *PlanningSlotTemplate) (int64, error) {
	ids, err := c.CreatePlanningSlotTemplates([]*PlanningSlotTemplate{pst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningSlotTemplate creates a new planning.slot.template model and returns its id.
func (c *Client) CreatePlanningSlotTemplates(psts []*PlanningSlotTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range psts {
		vv = append(vv, v)
	}
	return c.Create(PlanningSlotTemplateModel, vv)
}

// UpdatePlanningSlotTemplate updates an existing planning.slot.template record.
func (c *Client) UpdatePlanningSlotTemplate(pst *PlanningSlotTemplate) error {
	return c.UpdatePlanningSlotTemplates([]int64{pst.Id.Get()}, pst)
}

// UpdatePlanningSlotTemplates updates existing planning.slot.template records.
// All records (represented by ids) will be updated by pst values.
func (c *Client) UpdatePlanningSlotTemplates(ids []int64, pst *PlanningSlotTemplate) error {
	return c.Update(PlanningSlotTemplateModel, ids, pst)
}

// DeletePlanningSlotTemplate deletes an existing planning.slot.template record.
func (c *Client) DeletePlanningSlotTemplate(id int64) error {
	return c.DeletePlanningSlotTemplates([]int64{id})
}

// DeletePlanningSlotTemplates deletes existing planning.slot.template records.
func (c *Client) DeletePlanningSlotTemplates(ids []int64) error {
	return c.Delete(PlanningSlotTemplateModel, ids)
}

// GetPlanningSlotTemplate gets planning.slot.template existing record.
func (c *Client) GetPlanningSlotTemplate(id int64) (*PlanningSlotTemplate, error) {
	psts, err := c.GetPlanningSlotTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if psts != nil && len(*psts) > 0 {
		return &((*psts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of planning.slot.template not found", id)
}

// GetPlanningSlotTemplates gets planning.slot.template existing records.
func (c *Client) GetPlanningSlotTemplates(ids []int64) (*PlanningSlotTemplates, error) {
	psts := &PlanningSlotTemplates{}
	if err := c.Read(PlanningSlotTemplateModel, ids, nil, psts); err != nil {
		return nil, err
	}
	return psts, nil
}

// FindPlanningSlotTemplate finds planning.slot.template record by querying it with criteria.
func (c *Client) FindPlanningSlotTemplate(criteria *Criteria) (*PlanningSlotTemplate, error) {
	psts := &PlanningSlotTemplates{}
	if err := c.SearchRead(PlanningSlotTemplateModel, criteria, NewOptions().Limit(1), psts); err != nil {
		return nil, err
	}
	if psts != nil && len(*psts) > 0 {
		return &((*psts)[0]), nil
	}
	return nil, fmt.Errorf("planning.slot.template was not found with criteria %v", criteria)
}

// FindPlanningSlotTemplates finds planning.slot.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotTemplates(criteria *Criteria, options *Options) (*PlanningSlotTemplates, error) {
	psts := &PlanningSlotTemplates{}
	if err := c.SearchRead(PlanningSlotTemplateModel, criteria, options, psts); err != nil {
		return nil, err
	}
	return psts, nil
}

// FindPlanningSlotTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSlotTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PlanningSlotTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPlanningSlotTemplateId finds record id by querying it with criteria.
func (c *Client) FindPlanningSlotTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningSlotTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("planning.slot.template was not found with criteria %v and options %v", criteria, options)
}
