package odoo

import (
	"fmt"
)

// HelpdeskTicketType represents helpdesk.ticket.type model.
type HelpdeskTicketType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskTicketTypes represents array of helpdesk.ticket.type model.
type HelpdeskTicketTypes []HelpdeskTicketType

// HelpdeskTicketTypeModel is the odoo model name.
const HelpdeskTicketTypeModel = "helpdesk.ticket.type"

// Many2One convert HelpdeskTicketType to *Many2One.
func (htt *HelpdeskTicketType) Many2One() *Many2One {
	return NewMany2One(htt.Id.Get(), "")
}

// CreateHelpdeskTicketType creates a new helpdesk.ticket.type model and returns its id.
func (c *Client) CreateHelpdeskTicketType(htt *HelpdeskTicketType) (int64, error) {
	ids, err := c.CreateHelpdeskTicketTypes([]*HelpdeskTicketType{htt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicketType creates a new helpdesk.ticket.type model and returns its id.
func (c *Client) CreateHelpdeskTicketTypes(htts []*HelpdeskTicketType) ([]int64, error) {
	var vv []interface{}
	for _, v := range htts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketTypeModel, vv)
}

// UpdateHelpdeskTicketType updates an existing helpdesk.ticket.type record.
func (c *Client) UpdateHelpdeskTicketType(htt *HelpdeskTicketType) error {
	return c.UpdateHelpdeskTicketTypes([]int64{htt.Id.Get()}, htt)
}

// UpdateHelpdeskTicketTypes updates existing helpdesk.ticket.type records.
// All records (represented by ids) will be updated by htt values.
func (c *Client) UpdateHelpdeskTicketTypes(ids []int64, htt *HelpdeskTicketType) error {
	return c.Update(HelpdeskTicketTypeModel, ids, htt)
}

// DeleteHelpdeskTicketType deletes an existing helpdesk.ticket.type record.
func (c *Client) DeleteHelpdeskTicketType(id int64) error {
	return c.DeleteHelpdeskTicketTypes([]int64{id})
}

// DeleteHelpdeskTicketTypes deletes existing helpdesk.ticket.type records.
func (c *Client) DeleteHelpdeskTicketTypes(ids []int64) error {
	return c.Delete(HelpdeskTicketTypeModel, ids)
}

// GetHelpdeskTicketType gets helpdesk.ticket.type existing record.
func (c *Client) GetHelpdeskTicketType(id int64) (*HelpdeskTicketType, error) {
	htts, err := c.GetHelpdeskTicketTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if htts != nil && len(*htts) > 0 {
		return &((*htts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.ticket.type not found", id)
}

// GetHelpdeskTicketTypes gets helpdesk.ticket.type existing records.
func (c *Client) GetHelpdeskTicketTypes(ids []int64) (*HelpdeskTicketTypes, error) {
	htts := &HelpdeskTicketTypes{}
	if err := c.Read(HelpdeskTicketTypeModel, ids, nil, htts); err != nil {
		return nil, err
	}
	return htts, nil
}

// FindHelpdeskTicketType finds helpdesk.ticket.type record by querying it with criteria.
func (c *Client) FindHelpdeskTicketType(criteria *Criteria) (*HelpdeskTicketType, error) {
	htts := &HelpdeskTicketTypes{}
	if err := c.SearchRead(HelpdeskTicketTypeModel, criteria, NewOptions().Limit(1), htts); err != nil {
		return nil, err
	}
	if htts != nil && len(*htts) > 0 {
		return &((*htts)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.ticket.type was not found with criteria %v", criteria)
}

// FindHelpdeskTicketTypes finds helpdesk.ticket.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketTypes(criteria *Criteria, options *Options) (*HelpdeskTicketTypes, error) {
	htts := &HelpdeskTicketTypes{}
	if err := c.SearchRead(HelpdeskTicketTypeModel, criteria, options, htts); err != nil {
		return nil, err
	}
	return htts, nil
}

// FindHelpdeskTicketTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTicketTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTicketTypeId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.ticket.type was not found with criteria %v and options %v", criteria, options)
}
