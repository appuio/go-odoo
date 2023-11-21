package odoo

import (
	"fmt"
)

// EventTypeTicket represents event.type.ticket model.
type EventTypeTicket struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	Description  *String   `xmlrpc:"description,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	EventTypeId  *Many2One `xmlrpc:"event_type_id,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	SeatsLimited *Bool     `xmlrpc:"seats_limited,omptempty"`
	SeatsMax     *Int      `xmlrpc:"seats_max,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// EventTypeTickets represents array of event.type.ticket model.
type EventTypeTickets []EventTypeTicket

// EventTypeTicketModel is the odoo model name.
const EventTypeTicketModel = "event.type.ticket"

// Many2One convert EventTypeTicket to *Many2One.
func (ett *EventTypeTicket) Many2One() *Many2One {
	return NewMany2One(ett.Id.Get(), "")
}

// CreateEventTypeTicket creates a new event.type.ticket model and returns its id.
func (c *Client) CreateEventTypeTicket(ett *EventTypeTicket) (int64, error) {
	ids, err := c.CreateEventTypeTickets([]*EventTypeTicket{ett})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventTypeTicket creates a new event.type.ticket model and returns its id.
func (c *Client) CreateEventTypeTickets(etts []*EventTypeTicket) ([]int64, error) {
	var vv []interface{}
	for _, v := range etts {
		vv = append(vv, v)
	}
	return c.Create(EventTypeTicketModel, vv)
}

// UpdateEventTypeTicket updates an existing event.type.ticket record.
func (c *Client) UpdateEventTypeTicket(ett *EventTypeTicket) error {
	return c.UpdateEventTypeTickets([]int64{ett.Id.Get()}, ett)
}

// UpdateEventTypeTickets updates existing event.type.ticket records.
// All records (represented by ids) will be updated by ett values.
func (c *Client) UpdateEventTypeTickets(ids []int64, ett *EventTypeTicket) error {
	return c.Update(EventTypeTicketModel, ids, ett)
}

// DeleteEventTypeTicket deletes an existing event.type.ticket record.
func (c *Client) DeleteEventTypeTicket(id int64) error {
	return c.DeleteEventTypeTickets([]int64{id})
}

// DeleteEventTypeTickets deletes existing event.type.ticket records.
func (c *Client) DeleteEventTypeTickets(ids []int64) error {
	return c.Delete(EventTypeTicketModel, ids)
}

// GetEventTypeTicket gets event.type.ticket existing record.
func (c *Client) GetEventTypeTicket(id int64) (*EventTypeTicket, error) {
	etts, err := c.GetEventTypeTickets([]int64{id})
	if err != nil {
		return nil, err
	}
	if etts != nil && len(*etts) > 0 {
		return &((*etts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.type.ticket not found", id)
}

// GetEventTypeTickets gets event.type.ticket existing records.
func (c *Client) GetEventTypeTickets(ids []int64) (*EventTypeTickets, error) {
	etts := &EventTypeTickets{}
	if err := c.Read(EventTypeTicketModel, ids, nil, etts); err != nil {
		return nil, err
	}
	return etts, nil
}

// FindEventTypeTicket finds event.type.ticket record by querying it with criteria.
func (c *Client) FindEventTypeTicket(criteria *Criteria) (*EventTypeTicket, error) {
	etts := &EventTypeTickets{}
	if err := c.SearchRead(EventTypeTicketModel, criteria, NewOptions().Limit(1), etts); err != nil {
		return nil, err
	}
	if etts != nil && len(*etts) > 0 {
		return &((*etts)[0]), nil
	}
	return nil, fmt.Errorf("event.type.ticket was not found with criteria %v", criteria)
}

// FindEventTypeTickets finds event.type.ticket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypeTickets(criteria *Criteria, options *Options) (*EventTypeTickets, error) {
	etts := &EventTypeTickets{}
	if err := c.SearchRead(EventTypeTicketModel, criteria, options, etts); err != nil {
		return nil, err
	}
	return etts, nil
}

// FindEventTypeTicketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTypeTicketIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventTypeTicketModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventTypeTicketId finds record id by querying it with criteria.
func (c *Client) FindEventTypeTicketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventTypeTicketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.type.ticket was not found with criteria %v and options %v", criteria, options)
}
