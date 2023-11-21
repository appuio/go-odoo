package odoo

import (
	"fmt"
)

// PlanningPlanning represents planning.planning model.
type PlanningPlanning struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	AccessToken            *String   `xmlrpc:"access_token,omptempty"`
	AllowSelfUnassign      *Bool     `xmlrpc:"allow_self_unassign,omptempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DateEnd                *Time     `xmlrpc:"date_end,omptempty"`
	DateStart              *Time     `xmlrpc:"date_start,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	EndDatetime            *Time     `xmlrpc:"end_datetime,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	IncludeUnassigned      *Bool     `xmlrpc:"include_unassigned,omptempty"`
	SelfUnassignDaysBefore *Int      `xmlrpc:"self_unassign_days_before,omptempty"`
	SlotIds                *Relation `xmlrpc:"slot_ids,omptempty"`
	StartDatetime          *Time     `xmlrpc:"start_datetime,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PlanningPlannings represents array of planning.planning model.
type PlanningPlannings []PlanningPlanning

// PlanningPlanningModel is the odoo model name.
const PlanningPlanningModel = "planning.planning"

// Many2One convert PlanningPlanning to *Many2One.
func (pp *PlanningPlanning) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePlanningPlanning creates a new planning.planning model and returns its id.
func (c *Client) CreatePlanningPlanning(pp *PlanningPlanning) (int64, error) {
	ids, err := c.CreatePlanningPlannings([]*PlanningPlanning{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningPlanning creates a new planning.planning model and returns its id.
func (c *Client) CreatePlanningPlannings(pps []*PlanningPlanning) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(PlanningPlanningModel, vv)
}

// UpdatePlanningPlanning updates an existing planning.planning record.
func (c *Client) UpdatePlanningPlanning(pp *PlanningPlanning) error {
	return c.UpdatePlanningPlannings([]int64{pp.Id.Get()}, pp)
}

// UpdatePlanningPlannings updates existing planning.planning records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePlanningPlannings(ids []int64, pp *PlanningPlanning) error {
	return c.Update(PlanningPlanningModel, ids, pp)
}

// DeletePlanningPlanning deletes an existing planning.planning record.
func (c *Client) DeletePlanningPlanning(id int64) error {
	return c.DeletePlanningPlannings([]int64{id})
}

// DeletePlanningPlannings deletes existing planning.planning records.
func (c *Client) DeletePlanningPlannings(ids []int64) error {
	return c.Delete(PlanningPlanningModel, ids)
}

// GetPlanningPlanning gets planning.planning existing record.
func (c *Client) GetPlanningPlanning(id int64) (*PlanningPlanning, error) {
	pps, err := c.GetPlanningPlannings([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of planning.planning not found", id)
}

// GetPlanningPlannings gets planning.planning existing records.
func (c *Client) GetPlanningPlannings(ids []int64) (*PlanningPlannings, error) {
	pps := &PlanningPlannings{}
	if err := c.Read(PlanningPlanningModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPlanningPlanning finds planning.planning record by querying it with criteria.
func (c *Client) FindPlanningPlanning(criteria *Criteria) (*PlanningPlanning, error) {
	pps := &PlanningPlannings{}
	if err := c.SearchRead(PlanningPlanningModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("planning.planning was not found with criteria %v", criteria)
}

// FindPlanningPlannings finds planning.planning records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningPlannings(criteria *Criteria, options *Options) (*PlanningPlannings, error) {
	pps := &PlanningPlannings{}
	if err := c.SearchRead(PlanningPlanningModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPlanningPlanningIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningPlanningIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PlanningPlanningModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPlanningPlanningId finds record id by querying it with criteria.
func (c *Client) FindPlanningPlanningId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningPlanningModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("planning.planning was not found with criteria %v and options %v", criteria, options)
}
