package odoo

import (
	"fmt"
)

// EventMail represents event.mail model.
type EventMail struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	EventId             *Many2One  `xmlrpc:"event_id,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	IntervalNbr         *Int       `xmlrpc:"interval_nbr,omptempty"`
	IntervalType        *Selection `xmlrpc:"interval_type,omptempty"`
	IntervalUnit        *Selection `xmlrpc:"interval_unit,omptempty"`
	MailCountDone       *Int       `xmlrpc:"mail_count_done,omptempty"`
	MailDone            *Bool      `xmlrpc:"mail_done,omptempty"`
	MailRegistrationIds *Relation  `xmlrpc:"mail_registration_ids,omptempty"`
	MailState           *Selection `xmlrpc:"mail_state,omptempty"`
	NotificationType    *Selection `xmlrpc:"notification_type,omptempty"`
	ScheduledDate       *Time      `xmlrpc:"scheduled_date,omptempty"`
	Sequence            *Int       `xmlrpc:"sequence,omptempty"`
	TemplateModelId     *Many2One  `xmlrpc:"template_model_id,omptempty"`
	TemplateRef         *String    `xmlrpc:"template_ref,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventMails represents array of event.mail model.
type EventMails []EventMail

// EventMailModel is the odoo model name.
const EventMailModel = "event.mail"

// Many2One convert EventMail to *Many2One.
func (em *EventMail) Many2One() *Many2One {
	return NewMany2One(em.Id.Get(), "")
}

// CreateEventMail creates a new event.mail model and returns its id.
func (c *Client) CreateEventMail(em *EventMail) (int64, error) {
	ids, err := c.CreateEventMails([]*EventMail{em})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventMail creates a new event.mail model and returns its id.
func (c *Client) CreateEventMails(ems []*EventMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range ems {
		vv = append(vv, v)
	}
	return c.Create(EventMailModel, vv)
}

// UpdateEventMail updates an existing event.mail record.
func (c *Client) UpdateEventMail(em *EventMail) error {
	return c.UpdateEventMails([]int64{em.Id.Get()}, em)
}

// UpdateEventMails updates existing event.mail records.
// All records (represented by ids) will be updated by em values.
func (c *Client) UpdateEventMails(ids []int64, em *EventMail) error {
	return c.Update(EventMailModel, ids, em)
}

// DeleteEventMail deletes an existing event.mail record.
func (c *Client) DeleteEventMail(id int64) error {
	return c.DeleteEventMails([]int64{id})
}

// DeleteEventMails deletes existing event.mail records.
func (c *Client) DeleteEventMails(ids []int64) error {
	return c.Delete(EventMailModel, ids)
}

// GetEventMail gets event.mail existing record.
func (c *Client) GetEventMail(id int64) (*EventMail, error) {
	ems, err := c.GetEventMails([]int64{id})
	if err != nil {
		return nil, err
	}
	if ems != nil && len(*ems) > 0 {
		return &((*ems)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.mail not found", id)
}

// GetEventMails gets event.mail existing records.
func (c *Client) GetEventMails(ids []int64) (*EventMails, error) {
	ems := &EventMails{}
	if err := c.Read(EventMailModel, ids, nil, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindEventMail finds event.mail record by querying it with criteria.
func (c *Client) FindEventMail(criteria *Criteria) (*EventMail, error) {
	ems := &EventMails{}
	if err := c.SearchRead(EventMailModel, criteria, NewOptions().Limit(1), ems); err != nil {
		return nil, err
	}
	if ems != nil && len(*ems) > 0 {
		return &((*ems)[0]), nil
	}
	return nil, fmt.Errorf("event.mail was not found with criteria %v", criteria)
}

// FindEventMails finds event.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMails(criteria *Criteria, options *Options) (*EventMails, error) {
	ems := &EventMails{}
	if err := c.SearchRead(EventMailModel, criteria, options, ems); err != nil {
		return nil, err
	}
	return ems, nil
}

// FindEventMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventMailModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventMailId finds record id by querying it with criteria.
func (c *Client) FindEventMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.mail was not found with criteria %v and options %v", criteria, options)
}
