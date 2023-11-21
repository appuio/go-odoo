package odoo

import (
	"fmt"
)

// EventTag represents event.tag model.
type EventTag struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CategoryId       *Many2One `xmlrpc:"category_id,omptempty"`
	CategorySequence *Int      `xmlrpc:"category_sequence,omptempty"`
	Color            *Int      `xmlrpc:"color,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	Sequence         *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// EventTags represents array of event.tag model.
type EventTags []EventTag

// EventTagModel is the odoo model name.
const EventTagModel = "event.tag"

// Many2One convert EventTag to *Many2One.
func (et *EventTag) Many2One() *Many2One {
	return NewMany2One(et.Id.Get(), "")
}

// CreateEventTag creates a new event.tag model and returns its id.
func (c *Client) CreateEventTag(et *EventTag) (int64, error) {
	ids, err := c.CreateEventTags([]*EventTag{et})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventTag creates a new event.tag model and returns its id.
func (c *Client) CreateEventTags(ets []*EventTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range ets {
		vv = append(vv, v)
	}
	return c.Create(EventTagModel, vv)
}

// UpdateEventTag updates an existing event.tag record.
func (c *Client) UpdateEventTag(et *EventTag) error {
	return c.UpdateEventTags([]int64{et.Id.Get()}, et)
}

// UpdateEventTags updates existing event.tag records.
// All records (represented by ids) will be updated by et values.
func (c *Client) UpdateEventTags(ids []int64, et *EventTag) error {
	return c.Update(EventTagModel, ids, et)
}

// DeleteEventTag deletes an existing event.tag record.
func (c *Client) DeleteEventTag(id int64) error {
	return c.DeleteEventTags([]int64{id})
}

// DeleteEventTags deletes existing event.tag records.
func (c *Client) DeleteEventTags(ids []int64) error {
	return c.Delete(EventTagModel, ids)
}

// GetEventTag gets event.tag existing record.
func (c *Client) GetEventTag(id int64) (*EventTag, error) {
	ets, err := c.GetEventTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if ets != nil && len(*ets) > 0 {
		return &((*ets)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.tag not found", id)
}

// GetEventTags gets event.tag existing records.
func (c *Client) GetEventTags(ids []int64) (*EventTags, error) {
	ets := &EventTags{}
	if err := c.Read(EventTagModel, ids, nil, ets); err != nil {
		return nil, err
	}
	return ets, nil
}

// FindEventTag finds event.tag record by querying it with criteria.
func (c *Client) FindEventTag(criteria *Criteria) (*EventTag, error) {
	ets := &EventTags{}
	if err := c.SearchRead(EventTagModel, criteria, NewOptions().Limit(1), ets); err != nil {
		return nil, err
	}
	if ets != nil && len(*ets) > 0 {
		return &((*ets)[0]), nil
	}
	return nil, fmt.Errorf("event.tag was not found with criteria %v", criteria)
}

// FindEventTags finds event.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTags(criteria *Criteria, options *Options) (*EventTags, error) {
	ets := &EventTags{}
	if err := c.SearchRead(EventTagModel, criteria, options, ets); err != nil {
		return nil, err
	}
	return ets, nil
}

// FindEventTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventTagId finds record id by querying it with criteria.
func (c *Client) FindEventTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.tag was not found with criteria %v and options %v", criteria, options)
}
