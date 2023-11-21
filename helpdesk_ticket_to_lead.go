package odoo

import (
	"fmt"
)

// HelpdeskTicketToLead represents helpdesk.ticket.to.lead model.
type HelpdeskTicketToLead struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Action      *Selection `xmlrpc:"action,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	PartnerId   *Many2One  `xmlrpc:"partner_id,omptempty"`
	TeamId      *Many2One  `xmlrpc:"team_id,omptempty"`
	TicketId    *Many2One  `xmlrpc:"ticket_id,omptempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskTicketToLeads represents array of helpdesk.ticket.to.lead model.
type HelpdeskTicketToLeads []HelpdeskTicketToLead

// HelpdeskTicketToLeadModel is the odoo model name.
const HelpdeskTicketToLeadModel = "helpdesk.ticket.to.lead"

// Many2One convert HelpdeskTicketToLead to *Many2One.
func (httl *HelpdeskTicketToLead) Many2One() *Many2One {
	return NewMany2One(httl.Id.Get(), "")
}

// CreateHelpdeskTicketToLead creates a new helpdesk.ticket.to.lead model and returns its id.
func (c *Client) CreateHelpdeskTicketToLead(httl *HelpdeskTicketToLead) (int64, error) {
	ids, err := c.CreateHelpdeskTicketToLeads([]*HelpdeskTicketToLead{httl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicketToLead creates a new helpdesk.ticket.to.lead model and returns its id.
func (c *Client) CreateHelpdeskTicketToLeads(httls []*HelpdeskTicketToLead) ([]int64, error) {
	var vv []interface{}
	for _, v := range httls {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketToLeadModel, vv)
}

// UpdateHelpdeskTicketToLead updates an existing helpdesk.ticket.to.lead record.
func (c *Client) UpdateHelpdeskTicketToLead(httl *HelpdeskTicketToLead) error {
	return c.UpdateHelpdeskTicketToLeads([]int64{httl.Id.Get()}, httl)
}

// UpdateHelpdeskTicketToLeads updates existing helpdesk.ticket.to.lead records.
// All records (represented by ids) will be updated by httl values.
func (c *Client) UpdateHelpdeskTicketToLeads(ids []int64, httl *HelpdeskTicketToLead) error {
	return c.Update(HelpdeskTicketToLeadModel, ids, httl)
}

// DeleteHelpdeskTicketToLead deletes an existing helpdesk.ticket.to.lead record.
func (c *Client) DeleteHelpdeskTicketToLead(id int64) error {
	return c.DeleteHelpdeskTicketToLeads([]int64{id})
}

// DeleteHelpdeskTicketToLeads deletes existing helpdesk.ticket.to.lead records.
func (c *Client) DeleteHelpdeskTicketToLeads(ids []int64) error {
	return c.Delete(HelpdeskTicketToLeadModel, ids)
}

// GetHelpdeskTicketToLead gets helpdesk.ticket.to.lead existing record.
func (c *Client) GetHelpdeskTicketToLead(id int64) (*HelpdeskTicketToLead, error) {
	httls, err := c.GetHelpdeskTicketToLeads([]int64{id})
	if err != nil {
		return nil, err
	}
	if httls != nil && len(*httls) > 0 {
		return &((*httls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.ticket.to.lead not found", id)
}

// GetHelpdeskTicketToLeads gets helpdesk.ticket.to.lead existing records.
func (c *Client) GetHelpdeskTicketToLeads(ids []int64) (*HelpdeskTicketToLeads, error) {
	httls := &HelpdeskTicketToLeads{}
	if err := c.Read(HelpdeskTicketToLeadModel, ids, nil, httls); err != nil {
		return nil, err
	}
	return httls, nil
}

// FindHelpdeskTicketToLead finds helpdesk.ticket.to.lead record by querying it with criteria.
func (c *Client) FindHelpdeskTicketToLead(criteria *Criteria) (*HelpdeskTicketToLead, error) {
	httls := &HelpdeskTicketToLeads{}
	if err := c.SearchRead(HelpdeskTicketToLeadModel, criteria, NewOptions().Limit(1), httls); err != nil {
		return nil, err
	}
	if httls != nil && len(*httls) > 0 {
		return &((*httls)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.ticket.to.lead was not found with criteria %v", criteria)
}

// FindHelpdeskTicketToLeads finds helpdesk.ticket.to.lead records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketToLeads(criteria *Criteria, options *Options) (*HelpdeskTicketToLeads, error) {
	httls := &HelpdeskTicketToLeads{}
	if err := c.SearchRead(HelpdeskTicketToLeadModel, criteria, options, httls); err != nil {
		return nil, err
	}
	return httls, nil
}

// FindHelpdeskTicketToLeadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketToLeadIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTicketToLeadModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTicketToLeadId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketToLeadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketToLeadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.ticket.to.lead was not found with criteria %v and options %v", criteria, options)
}
