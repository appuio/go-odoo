package odoo

import (
	"fmt"
)

// HelpdeskTicketConvertWizard represents helpdesk.ticket.convert.wizard model.
type HelpdeskTicketConvertWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ProjectId   *Many2One `xmlrpc:"project_id,omptempty"`
	StageId     *Many2One `xmlrpc:"stage_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskTicketConvertWizards represents array of helpdesk.ticket.convert.wizard model.
type HelpdeskTicketConvertWizards []HelpdeskTicketConvertWizard

// HelpdeskTicketConvertWizardModel is the odoo model name.
const HelpdeskTicketConvertWizardModel = "helpdesk.ticket.convert.wizard"

// Many2One convert HelpdeskTicketConvertWizard to *Many2One.
func (htcw *HelpdeskTicketConvertWizard) Many2One() *Many2One {
	return NewMany2One(htcw.Id.Get(), "")
}

// CreateHelpdeskTicketConvertWizard creates a new helpdesk.ticket.convert.wizard model and returns its id.
func (c *Client) CreateHelpdeskTicketConvertWizard(htcw *HelpdeskTicketConvertWizard) (int64, error) {
	ids, err := c.CreateHelpdeskTicketConvertWizards([]*HelpdeskTicketConvertWizard{htcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicketConvertWizard creates a new helpdesk.ticket.convert.wizard model and returns its id.
func (c *Client) CreateHelpdeskTicketConvertWizards(htcws []*HelpdeskTicketConvertWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range htcws {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketConvertWizardModel, vv)
}

// UpdateHelpdeskTicketConvertWizard updates an existing helpdesk.ticket.convert.wizard record.
func (c *Client) UpdateHelpdeskTicketConvertWizard(htcw *HelpdeskTicketConvertWizard) error {
	return c.UpdateHelpdeskTicketConvertWizards([]int64{htcw.Id.Get()}, htcw)
}

// UpdateHelpdeskTicketConvertWizards updates existing helpdesk.ticket.convert.wizard records.
// All records (represented by ids) will be updated by htcw values.
func (c *Client) UpdateHelpdeskTicketConvertWizards(ids []int64, htcw *HelpdeskTicketConvertWizard) error {
	return c.Update(HelpdeskTicketConvertWizardModel, ids, htcw)
}

// DeleteHelpdeskTicketConvertWizard deletes an existing helpdesk.ticket.convert.wizard record.
func (c *Client) DeleteHelpdeskTicketConvertWizard(id int64) error {
	return c.DeleteHelpdeskTicketConvertWizards([]int64{id})
}

// DeleteHelpdeskTicketConvertWizards deletes existing helpdesk.ticket.convert.wizard records.
func (c *Client) DeleteHelpdeskTicketConvertWizards(ids []int64) error {
	return c.Delete(HelpdeskTicketConvertWizardModel, ids)
}

// GetHelpdeskTicketConvertWizard gets helpdesk.ticket.convert.wizard existing record.
func (c *Client) GetHelpdeskTicketConvertWizard(id int64) (*HelpdeskTicketConvertWizard, error) {
	htcws, err := c.GetHelpdeskTicketConvertWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if htcws != nil && len(*htcws) > 0 {
		return &((*htcws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.ticket.convert.wizard not found", id)
}

// GetHelpdeskTicketConvertWizards gets helpdesk.ticket.convert.wizard existing records.
func (c *Client) GetHelpdeskTicketConvertWizards(ids []int64) (*HelpdeskTicketConvertWizards, error) {
	htcws := &HelpdeskTicketConvertWizards{}
	if err := c.Read(HelpdeskTicketConvertWizardModel, ids, nil, htcws); err != nil {
		return nil, err
	}
	return htcws, nil
}

// FindHelpdeskTicketConvertWizard finds helpdesk.ticket.convert.wizard record by querying it with criteria.
func (c *Client) FindHelpdeskTicketConvertWizard(criteria *Criteria) (*HelpdeskTicketConvertWizard, error) {
	htcws := &HelpdeskTicketConvertWizards{}
	if err := c.SearchRead(HelpdeskTicketConvertWizardModel, criteria, NewOptions().Limit(1), htcws); err != nil {
		return nil, err
	}
	if htcws != nil && len(*htcws) > 0 {
		return &((*htcws)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.ticket.convert.wizard was not found with criteria %v", criteria)
}

// FindHelpdeskTicketConvertWizards finds helpdesk.ticket.convert.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketConvertWizards(criteria *Criteria, options *Options) (*HelpdeskTicketConvertWizards, error) {
	htcws := &HelpdeskTicketConvertWizards{}
	if err := c.SearchRead(HelpdeskTicketConvertWizardModel, criteria, options, htcws); err != nil {
		return nil, err
	}
	return htcws, nil
}

// FindHelpdeskTicketConvertWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketConvertWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTicketConvertWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTicketConvertWizardId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketConvertWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketConvertWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.ticket.convert.wizard was not found with criteria %v and options %v", criteria, options)
}
