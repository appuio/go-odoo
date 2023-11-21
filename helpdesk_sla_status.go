package odoo

import (
	"fmt"
)

// HelpdeskSlaStatus represents helpdesk.sla.status model.
type HelpdeskSlaStatus struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Color           *Int       `xmlrpc:"color,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	Deadline        *Time      `xmlrpc:"deadline,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	ExceededHours   *Float     `xmlrpc:"exceeded_hours,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	ReachedDatetime *Time      `xmlrpc:"reached_datetime,omptempty"`
	SlaId           *Many2One  `xmlrpc:"sla_id,omptempty"`
	SlaStageId      *Many2One  `xmlrpc:"sla_stage_id,omptempty"`
	Status          *Selection `xmlrpc:"status,omptempty"`
	TicketId        *Many2One  `xmlrpc:"ticket_id,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskSlaStatuss represents array of helpdesk.sla.status model.
type HelpdeskSlaStatuss []HelpdeskSlaStatus

// HelpdeskSlaStatusModel is the odoo model name.
const HelpdeskSlaStatusModel = "helpdesk.sla.status"

// Many2One convert HelpdeskSlaStatus to *Many2One.
func (hss *HelpdeskSlaStatus) Many2One() *Many2One {
	return NewMany2One(hss.Id.Get(), "")
}

// CreateHelpdeskSlaStatus creates a new helpdesk.sla.status model and returns its id.
func (c *Client) CreateHelpdeskSlaStatus(hss *HelpdeskSlaStatus) (int64, error) {
	ids, err := c.CreateHelpdeskSlaStatuss([]*HelpdeskSlaStatus{hss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskSlaStatus creates a new helpdesk.sla.status model and returns its id.
func (c *Client) CreateHelpdeskSlaStatuss(hsss []*HelpdeskSlaStatus) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsss {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskSlaStatusModel, vv)
}

// UpdateHelpdeskSlaStatus updates an existing helpdesk.sla.status record.
func (c *Client) UpdateHelpdeskSlaStatus(hss *HelpdeskSlaStatus) error {
	return c.UpdateHelpdeskSlaStatuss([]int64{hss.Id.Get()}, hss)
}

// UpdateHelpdeskSlaStatuss updates existing helpdesk.sla.status records.
// All records (represented by ids) will be updated by hss values.
func (c *Client) UpdateHelpdeskSlaStatuss(ids []int64, hss *HelpdeskSlaStatus) error {
	return c.Update(HelpdeskSlaStatusModel, ids, hss)
}

// DeleteHelpdeskSlaStatus deletes an existing helpdesk.sla.status record.
func (c *Client) DeleteHelpdeskSlaStatus(id int64) error {
	return c.DeleteHelpdeskSlaStatuss([]int64{id})
}

// DeleteHelpdeskSlaStatuss deletes existing helpdesk.sla.status records.
func (c *Client) DeleteHelpdeskSlaStatuss(ids []int64) error {
	return c.Delete(HelpdeskSlaStatusModel, ids)
}

// GetHelpdeskSlaStatus gets helpdesk.sla.status existing record.
func (c *Client) GetHelpdeskSlaStatus(id int64) (*HelpdeskSlaStatus, error) {
	hsss, err := c.GetHelpdeskSlaStatuss([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsss != nil && len(*hsss) > 0 {
		return &((*hsss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.sla.status not found", id)
}

// GetHelpdeskSlaStatuss gets helpdesk.sla.status existing records.
func (c *Client) GetHelpdeskSlaStatuss(ids []int64) (*HelpdeskSlaStatuss, error) {
	hsss := &HelpdeskSlaStatuss{}
	if err := c.Read(HelpdeskSlaStatusModel, ids, nil, hsss); err != nil {
		return nil, err
	}
	return hsss, nil
}

// FindHelpdeskSlaStatus finds helpdesk.sla.status record by querying it with criteria.
func (c *Client) FindHelpdeskSlaStatus(criteria *Criteria) (*HelpdeskSlaStatus, error) {
	hsss := &HelpdeskSlaStatuss{}
	if err := c.SearchRead(HelpdeskSlaStatusModel, criteria, NewOptions().Limit(1), hsss); err != nil {
		return nil, err
	}
	if hsss != nil && len(*hsss) > 0 {
		return &((*hsss)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.sla.status was not found with criteria %v", criteria)
}

// FindHelpdeskSlaStatuss finds helpdesk.sla.status records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlaStatuss(criteria *Criteria, options *Options) (*HelpdeskSlaStatuss, error) {
	hsss := &HelpdeskSlaStatuss{}
	if err := c.SearchRead(HelpdeskSlaStatusModel, criteria, options, hsss); err != nil {
		return nil, err
	}
	return hsss, nil
}

// FindHelpdeskSlaStatusIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskSlaStatusIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskSlaStatusModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskSlaStatusId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskSlaStatusId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskSlaStatusModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.sla.status was not found with criteria %v and options %v", criteria, options)
}
