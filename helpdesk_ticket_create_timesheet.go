package odoo

import (
	"fmt"
)

// HelpdeskTicketCreateTimesheet represents helpdesk.ticket.create.timesheet model.
type HelpdeskTicketCreateTimesheet struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Description *String   `xmlrpc:"description,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	TicketId    *Many2One `xmlrpc:"ticket_id,omptempty"`
	TimeSpent   *Float    `xmlrpc:"time_spent,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskTicketCreateTimesheets represents array of helpdesk.ticket.create.timesheet model.
type HelpdeskTicketCreateTimesheets []HelpdeskTicketCreateTimesheet

// HelpdeskTicketCreateTimesheetModel is the odoo model name.
const HelpdeskTicketCreateTimesheetModel = "helpdesk.ticket.create.timesheet"

// Many2One convert HelpdeskTicketCreateTimesheet to *Many2One.
func (htct *HelpdeskTicketCreateTimesheet) Many2One() *Many2One {
	return NewMany2One(htct.Id.Get(), "")
}

// CreateHelpdeskTicketCreateTimesheet creates a new helpdesk.ticket.create.timesheet model and returns its id.
func (c *Client) CreateHelpdeskTicketCreateTimesheet(htct *HelpdeskTicketCreateTimesheet) (int64, error) {
	ids, err := c.CreateHelpdeskTicketCreateTimesheets([]*HelpdeskTicketCreateTimesheet{htct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicketCreateTimesheet creates a new helpdesk.ticket.create.timesheet model and returns its id.
func (c *Client) CreateHelpdeskTicketCreateTimesheets(htcts []*HelpdeskTicketCreateTimesheet) ([]int64, error) {
	var vv []interface{}
	for _, v := range htcts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketCreateTimesheetModel, vv)
}

// UpdateHelpdeskTicketCreateTimesheet updates an existing helpdesk.ticket.create.timesheet record.
func (c *Client) UpdateHelpdeskTicketCreateTimesheet(htct *HelpdeskTicketCreateTimesheet) error {
	return c.UpdateHelpdeskTicketCreateTimesheets([]int64{htct.Id.Get()}, htct)
}

// UpdateHelpdeskTicketCreateTimesheets updates existing helpdesk.ticket.create.timesheet records.
// All records (represented by ids) will be updated by htct values.
func (c *Client) UpdateHelpdeskTicketCreateTimesheets(ids []int64, htct *HelpdeskTicketCreateTimesheet) error {
	return c.Update(HelpdeskTicketCreateTimesheetModel, ids, htct)
}

// DeleteHelpdeskTicketCreateTimesheet deletes an existing helpdesk.ticket.create.timesheet record.
func (c *Client) DeleteHelpdeskTicketCreateTimesheet(id int64) error {
	return c.DeleteHelpdeskTicketCreateTimesheets([]int64{id})
}

// DeleteHelpdeskTicketCreateTimesheets deletes existing helpdesk.ticket.create.timesheet records.
func (c *Client) DeleteHelpdeskTicketCreateTimesheets(ids []int64) error {
	return c.Delete(HelpdeskTicketCreateTimesheetModel, ids)
}

// GetHelpdeskTicketCreateTimesheet gets helpdesk.ticket.create.timesheet existing record.
func (c *Client) GetHelpdeskTicketCreateTimesheet(id int64) (*HelpdeskTicketCreateTimesheet, error) {
	htcts, err := c.GetHelpdeskTicketCreateTimesheets([]int64{id})
	if err != nil {
		return nil, err
	}
	if htcts != nil && len(*htcts) > 0 {
		return &((*htcts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.ticket.create.timesheet not found", id)
}

// GetHelpdeskTicketCreateTimesheets gets helpdesk.ticket.create.timesheet existing records.
func (c *Client) GetHelpdeskTicketCreateTimesheets(ids []int64) (*HelpdeskTicketCreateTimesheets, error) {
	htcts := &HelpdeskTicketCreateTimesheets{}
	if err := c.Read(HelpdeskTicketCreateTimesheetModel, ids, nil, htcts); err != nil {
		return nil, err
	}
	return htcts, nil
}

// FindHelpdeskTicketCreateTimesheet finds helpdesk.ticket.create.timesheet record by querying it with criteria.
func (c *Client) FindHelpdeskTicketCreateTimesheet(criteria *Criteria) (*HelpdeskTicketCreateTimesheet, error) {
	htcts := &HelpdeskTicketCreateTimesheets{}
	if err := c.SearchRead(HelpdeskTicketCreateTimesheetModel, criteria, NewOptions().Limit(1), htcts); err != nil {
		return nil, err
	}
	if htcts != nil && len(*htcts) > 0 {
		return &((*htcts)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.ticket.create.timesheet was not found with criteria %v", criteria)
}

// FindHelpdeskTicketCreateTimesheets finds helpdesk.ticket.create.timesheet records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketCreateTimesheets(criteria *Criteria, options *Options) (*HelpdeskTicketCreateTimesheets, error) {
	htcts := &HelpdeskTicketCreateTimesheets{}
	if err := c.SearchRead(HelpdeskTicketCreateTimesheetModel, criteria, options, htcts); err != nil {
		return nil, err
	}
	return htcts, nil
}

// FindHelpdeskTicketCreateTimesheetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketCreateTimesheetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTicketCreateTimesheetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTicketCreateTimesheetId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketCreateTimesheetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketCreateTimesheetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.ticket.create.timesheet was not found with criteria %v and options %v", criteria, options)
}
