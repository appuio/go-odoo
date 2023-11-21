package odoo

import (
	"fmt"
)

// EventStage represents event.stage model.
type EventStage struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	Description   *String   `xmlrpc:"description,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Fold          *Bool     `xmlrpc:"fold,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LegendBlocked *String   `xmlrpc:"legend_blocked,omptempty"`
	LegendDone    *String   `xmlrpc:"legend_done,omptempty"`
	LegendNormal  *String   `xmlrpc:"legend_normal,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	PipeEnd       *Bool     `xmlrpc:"pipe_end,omptempty"`
	Sequence      *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// EventStages represents array of event.stage model.
type EventStages []EventStage

// EventStageModel is the odoo model name.
const EventStageModel = "event.stage"

// Many2One convert EventStage to *Many2One.
func (es *EventStage) Many2One() *Many2One {
	return NewMany2One(es.Id.Get(), "")
}

// CreateEventStage creates a new event.stage model and returns its id.
func (c *Client) CreateEventStage(es *EventStage) (int64, error) {
	ids, err := c.CreateEventStages([]*EventStage{es})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventStage creates a new event.stage model and returns its id.
func (c *Client) CreateEventStages(ess []*EventStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range ess {
		vv = append(vv, v)
	}
	return c.Create(EventStageModel, vv)
}

// UpdateEventStage updates an existing event.stage record.
func (c *Client) UpdateEventStage(es *EventStage) error {
	return c.UpdateEventStages([]int64{es.Id.Get()}, es)
}

// UpdateEventStages updates existing event.stage records.
// All records (represented by ids) will be updated by es values.
func (c *Client) UpdateEventStages(ids []int64, es *EventStage) error {
	return c.Update(EventStageModel, ids, es)
}

// DeleteEventStage deletes an existing event.stage record.
func (c *Client) DeleteEventStage(id int64) error {
	return c.DeleteEventStages([]int64{id})
}

// DeleteEventStages deletes existing event.stage records.
func (c *Client) DeleteEventStages(ids []int64) error {
	return c.Delete(EventStageModel, ids)
}

// GetEventStage gets event.stage existing record.
func (c *Client) GetEventStage(id int64) (*EventStage, error) {
	ess, err := c.GetEventStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if ess != nil && len(*ess) > 0 {
		return &((*ess)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.stage not found", id)
}

// GetEventStages gets event.stage existing records.
func (c *Client) GetEventStages(ids []int64) (*EventStages, error) {
	ess := &EventStages{}
	if err := c.Read(EventStageModel, ids, nil, ess); err != nil {
		return nil, err
	}
	return ess, nil
}

// FindEventStage finds event.stage record by querying it with criteria.
func (c *Client) FindEventStage(criteria *Criteria) (*EventStage, error) {
	ess := &EventStages{}
	if err := c.SearchRead(EventStageModel, criteria, NewOptions().Limit(1), ess); err != nil {
		return nil, err
	}
	if ess != nil && len(*ess) > 0 {
		return &((*ess)[0]), nil
	}
	return nil, fmt.Errorf("event.stage was not found with criteria %v", criteria)
}

// FindEventStages finds event.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventStages(criteria *Criteria, options *Options) (*EventStages, error) {
	ess := &EventStages{}
	if err := c.SearchRead(EventStageModel, criteria, options, ess); err != nil {
		return nil, err
	}
	return ess, nil
}

// FindEventStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventStageId finds record id by querying it with criteria.
func (c *Client) FindEventStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.stage was not found with criteria %v and options %v", criteria, options)
}
