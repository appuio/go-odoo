package odoo

import (
	"fmt"
)

// CalendarRecurrence represents calendar.recurrence model.
type CalendarRecurrence struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	BaseEventId      *Many2One  `xmlrpc:"base_event_id,omptempty"`
	Byday            *Selection `xmlrpc:"byday,omptempty"`
	CalendarEventIds *Relation  `xmlrpc:"calendar_event_ids,omptempty"`
	Count            *Int       `xmlrpc:"count,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	Day              *Int       `xmlrpc:"day,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	Dtstart          *Time      `xmlrpc:"dtstart,omptempty"`
	EndType          *Selection `xmlrpc:"end_type,omptempty"`
	EventTz          *Selection `xmlrpc:"event_tz,omptempty"`
	Fri              *Bool      `xmlrpc:"fri,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Interval         *Int       `xmlrpc:"interval,omptempty"`
	Mon              *Bool      `xmlrpc:"mon,omptempty"`
	MonthBy          *Selection `xmlrpc:"month_by,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	Rrule            *String    `xmlrpc:"rrule,omptempty"`
	RruleType        *Selection `xmlrpc:"rrule_type,omptempty"`
	Sat              *Bool      `xmlrpc:"sat,omptempty"`
	Sun              *Bool      `xmlrpc:"sun,omptempty"`
	Thu              *Bool      `xmlrpc:"thu,omptempty"`
	Tue              *Bool      `xmlrpc:"tue,omptempty"`
	Until            *Time      `xmlrpc:"until,omptempty"`
	Wed              *Bool      `xmlrpc:"wed,omptempty"`
	Weekday          *Selection `xmlrpc:"weekday,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CalendarRecurrences represents array of calendar.recurrence model.
type CalendarRecurrences []CalendarRecurrence

// CalendarRecurrenceModel is the odoo model name.
const CalendarRecurrenceModel = "calendar.recurrence"

// Many2One convert CalendarRecurrence to *Many2One.
func (cr *CalendarRecurrence) Many2One() *Many2One {
	return NewMany2One(cr.Id.Get(), "")
}

// CreateCalendarRecurrence creates a new calendar.recurrence model and returns its id.
func (c *Client) CreateCalendarRecurrence(cr *CalendarRecurrence) (int64, error) {
	ids, err := c.CreateCalendarRecurrences([]*CalendarRecurrence{cr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarRecurrence creates a new calendar.recurrence model and returns its id.
func (c *Client) CreateCalendarRecurrences(crs []*CalendarRecurrence) ([]int64, error) {
	var vv []interface{}
	for _, v := range crs {
		vv = append(vv, v)
	}
	return c.Create(CalendarRecurrenceModel, vv)
}

// UpdateCalendarRecurrence updates an existing calendar.recurrence record.
func (c *Client) UpdateCalendarRecurrence(cr *CalendarRecurrence) error {
	return c.UpdateCalendarRecurrences([]int64{cr.Id.Get()}, cr)
}

// UpdateCalendarRecurrences updates existing calendar.recurrence records.
// All records (represented by ids) will be updated by cr values.
func (c *Client) UpdateCalendarRecurrences(ids []int64, cr *CalendarRecurrence) error {
	return c.Update(CalendarRecurrenceModel, ids, cr)
}

// DeleteCalendarRecurrence deletes an existing calendar.recurrence record.
func (c *Client) DeleteCalendarRecurrence(id int64) error {
	return c.DeleteCalendarRecurrences([]int64{id})
}

// DeleteCalendarRecurrences deletes existing calendar.recurrence records.
func (c *Client) DeleteCalendarRecurrences(ids []int64) error {
	return c.Delete(CalendarRecurrenceModel, ids)
}

// GetCalendarRecurrence gets calendar.recurrence existing record.
func (c *Client) GetCalendarRecurrence(id int64) (*CalendarRecurrence, error) {
	crs, err := c.GetCalendarRecurrences([]int64{id})
	if err != nil {
		return nil, err
	}
	if crs != nil && len(*crs) > 0 {
		return &((*crs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.recurrence not found", id)
}

// GetCalendarRecurrences gets calendar.recurrence existing records.
func (c *Client) GetCalendarRecurrences(ids []int64) (*CalendarRecurrences, error) {
	crs := &CalendarRecurrences{}
	if err := c.Read(CalendarRecurrenceModel, ids, nil, crs); err != nil {
		return nil, err
	}
	return crs, nil
}

// FindCalendarRecurrence finds calendar.recurrence record by querying it with criteria.
func (c *Client) FindCalendarRecurrence(criteria *Criteria) (*CalendarRecurrence, error) {
	crs := &CalendarRecurrences{}
	if err := c.SearchRead(CalendarRecurrenceModel, criteria, NewOptions().Limit(1), crs); err != nil {
		return nil, err
	}
	if crs != nil && len(*crs) > 0 {
		return &((*crs)[0]), nil
	}
	return nil, fmt.Errorf("calendar.recurrence was not found with criteria %v", criteria)
}

// FindCalendarRecurrences finds calendar.recurrence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarRecurrences(criteria *Criteria, options *Options) (*CalendarRecurrences, error) {
	crs := &CalendarRecurrences{}
	if err := c.SearchRead(CalendarRecurrenceModel, criteria, options, crs); err != nil {
		return nil, err
	}
	return crs, nil
}

// FindCalendarRecurrenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarRecurrenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarRecurrenceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarRecurrenceId finds record id by querying it with criteria.
func (c *Client) FindCalendarRecurrenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarRecurrenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("calendar.recurrence was not found with criteria %v and options %v", criteria, options)
}
