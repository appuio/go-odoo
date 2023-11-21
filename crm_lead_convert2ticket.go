package odoo

import (
	"fmt"
)

// CrmLeadConvert2Ticket represents crm.lead.convert2ticket model.
type CrmLeadConvert2Ticket struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	LeadId       *Many2One `xmlrpc:"lead_id,omptempty"`
	PartnerId    *Many2One `xmlrpc:"partner_id,omptempty"`
	TeamId       *Many2One `xmlrpc:"team_id,omptempty"`
	TicketTypeId *Many2One `xmlrpc:"ticket_type_id,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CrmLeadConvert2Tickets represents array of crm.lead.convert2ticket model.
type CrmLeadConvert2Tickets []CrmLeadConvert2Ticket

// CrmLeadConvert2TicketModel is the odoo model name.
const CrmLeadConvert2TicketModel = "crm.lead.convert2ticket"

// Many2One convert CrmLeadConvert2Ticket to *Many2One.
func (clc *CrmLeadConvert2Ticket) Many2One() *Many2One {
	return NewMany2One(clc.Id.Get(), "")
}

// CreateCrmLeadConvert2Ticket creates a new crm.lead.convert2ticket model and returns its id.
func (c *Client) CreateCrmLeadConvert2Ticket(clc *CrmLeadConvert2Ticket) (int64, error) {
	ids, err := c.CreateCrmLeadConvert2Tickets([]*CrmLeadConvert2Ticket{clc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLeadConvert2Ticket creates a new crm.lead.convert2ticket model and returns its id.
func (c *Client) CreateCrmLeadConvert2Tickets(clcs []*CrmLeadConvert2Ticket) ([]int64, error) {
	var vv []interface{}
	for _, v := range clcs {
		vv = append(vv, v)
	}
	return c.Create(CrmLeadConvert2TicketModel, vv)
}

// UpdateCrmLeadConvert2Ticket updates an existing crm.lead.convert2ticket record.
func (c *Client) UpdateCrmLeadConvert2Ticket(clc *CrmLeadConvert2Ticket) error {
	return c.UpdateCrmLeadConvert2Tickets([]int64{clc.Id.Get()}, clc)
}

// UpdateCrmLeadConvert2Tickets updates existing crm.lead.convert2ticket records.
// All records (represented by ids) will be updated by clc values.
func (c *Client) UpdateCrmLeadConvert2Tickets(ids []int64, clc *CrmLeadConvert2Ticket) error {
	return c.Update(CrmLeadConvert2TicketModel, ids, clc)
}

// DeleteCrmLeadConvert2Ticket deletes an existing crm.lead.convert2ticket record.
func (c *Client) DeleteCrmLeadConvert2Ticket(id int64) error {
	return c.DeleteCrmLeadConvert2Tickets([]int64{id})
}

// DeleteCrmLeadConvert2Tickets deletes existing crm.lead.convert2ticket records.
func (c *Client) DeleteCrmLeadConvert2Tickets(ids []int64) error {
	return c.Delete(CrmLeadConvert2TicketModel, ids)
}

// GetCrmLeadConvert2Ticket gets crm.lead.convert2ticket existing record.
func (c *Client) GetCrmLeadConvert2Ticket(id int64) (*CrmLeadConvert2Ticket, error) {
	clcs, err := c.GetCrmLeadConvert2Tickets([]int64{id})
	if err != nil {
		return nil, err
	}
	if clcs != nil && len(*clcs) > 0 {
		return &((*clcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.lead.convert2ticket not found", id)
}

// GetCrmLeadConvert2Tickets gets crm.lead.convert2ticket existing records.
func (c *Client) GetCrmLeadConvert2Tickets(ids []int64) (*CrmLeadConvert2Tickets, error) {
	clcs := &CrmLeadConvert2Tickets{}
	if err := c.Read(CrmLeadConvert2TicketModel, ids, nil, clcs); err != nil {
		return nil, err
	}
	return clcs, nil
}

// FindCrmLeadConvert2Ticket finds crm.lead.convert2ticket record by querying it with criteria.
func (c *Client) FindCrmLeadConvert2Ticket(criteria *Criteria) (*CrmLeadConvert2Ticket, error) {
	clcs := &CrmLeadConvert2Tickets{}
	if err := c.SearchRead(CrmLeadConvert2TicketModel, criteria, NewOptions().Limit(1), clcs); err != nil {
		return nil, err
	}
	if clcs != nil && len(*clcs) > 0 {
		return &((*clcs)[0]), nil
	}
	return nil, fmt.Errorf("crm.lead.convert2ticket was not found with criteria %v", criteria)
}

// FindCrmLeadConvert2Tickets finds crm.lead.convert2ticket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadConvert2Tickets(criteria *Criteria, options *Options) (*CrmLeadConvert2Tickets, error) {
	clcs := &CrmLeadConvert2Tickets{}
	if err := c.SearchRead(CrmLeadConvert2TicketModel, criteria, options, clcs); err != nil {
		return nil, err
	}
	return clcs, nil
}

// FindCrmLeadConvert2TicketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadConvert2TicketIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmLeadConvert2TicketModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmLeadConvert2TicketId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadConvert2TicketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadConvert2TicketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.lead.convert2ticket was not found with criteria %v and options %v", criteria, options)
}
