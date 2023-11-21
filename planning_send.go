package odoo

import (
	"fmt"
)

// PlanningSend represents planning.send model.
type PlanningSend struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	EmployeeIds       *Relation `xmlrpc:"employee_ids,omptempty"`
	EmployeesNoEmail  *Relation `xmlrpc:"employees_no_email,omptempty"`
	EndDatetime       *Time     `xmlrpc:"end_datetime,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	IncludeUnassigned *Bool     `xmlrpc:"include_unassigned,omptempty"`
	Note              *String   `xmlrpc:"note,omptempty"`
	SlotIds           *Relation `xmlrpc:"slot_ids,omptempty"`
	StartDatetime     *Time     `xmlrpc:"start_datetime,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PlanningSends represents array of planning.send model.
type PlanningSends []PlanningSend

// PlanningSendModel is the odoo model name.
const PlanningSendModel = "planning.send"

// Many2One convert PlanningSend to *Many2One.
func (ps *PlanningSend) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePlanningSend creates a new planning.send model and returns its id.
func (c *Client) CreatePlanningSend(ps *PlanningSend) (int64, error) {
	ids, err := c.CreatePlanningSends([]*PlanningSend{ps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePlanningSend creates a new planning.send model and returns its id.
func (c *Client) CreatePlanningSends(pss []*PlanningSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range pss {
		vv = append(vv, v)
	}
	return c.Create(PlanningSendModel, vv)
}

// UpdatePlanningSend updates an existing planning.send record.
func (c *Client) UpdatePlanningSend(ps *PlanningSend) error {
	return c.UpdatePlanningSends([]int64{ps.Id.Get()}, ps)
}

// UpdatePlanningSends updates existing planning.send records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePlanningSends(ids []int64, ps *PlanningSend) error {
	return c.Update(PlanningSendModel, ids, ps)
}

// DeletePlanningSend deletes an existing planning.send record.
func (c *Client) DeletePlanningSend(id int64) error {
	return c.DeletePlanningSends([]int64{id})
}

// DeletePlanningSends deletes existing planning.send records.
func (c *Client) DeletePlanningSends(ids []int64) error {
	return c.Delete(PlanningSendModel, ids)
}

// GetPlanningSend gets planning.send existing record.
func (c *Client) GetPlanningSend(id int64) (*PlanningSend, error) {
	pss, err := c.GetPlanningSends([]int64{id})
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of planning.send not found", id)
}

// GetPlanningSends gets planning.send existing records.
func (c *Client) GetPlanningSends(ids []int64) (*PlanningSends, error) {
	pss := &PlanningSends{}
	if err := c.Read(PlanningSendModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSend finds planning.send record by querying it with criteria.
func (c *Client) FindPlanningSend(criteria *Criteria) (*PlanningSend, error) {
	pss := &PlanningSends{}
	if err := c.SearchRead(PlanningSendModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("planning.send was not found with criteria %v", criteria)
}

// FindPlanningSends finds planning.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSends(criteria *Criteria, options *Options) (*PlanningSends, error) {
	pss := &PlanningSends{}
	if err := c.SearchRead(PlanningSendModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPlanningSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPlanningSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PlanningSendModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPlanningSendId finds record id by querying it with criteria.
func (c *Client) FindPlanningSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PlanningSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("planning.send was not found with criteria %v and options %v", criteria, options)
}
