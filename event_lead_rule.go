package odoo

import (
	"fmt"
)

// EventLeadRule represents event.lead.rule model.
type EventLeadRule struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	Active                  *Bool      `xmlrpc:"active,omptempty"`
	CompanyId               *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	EventId                 *Many2One  `xmlrpc:"event_id,omptempty"`
	EventRegistrationFilter *String    `xmlrpc:"event_registration_filter,omptempty"`
	EventTypeIds            *Relation  `xmlrpc:"event_type_ids,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	LeadCreationBasis       *Selection `xmlrpc:"lead_creation_basis,omptempty"`
	LeadCreationTrigger     *Selection `xmlrpc:"lead_creation_trigger,omptempty"`
	LeadIds                 *Relation  `xmlrpc:"lead_ids,omptempty"`
	LeadSalesTeamId         *Many2One  `xmlrpc:"lead_sales_team_id,omptempty"`
	LeadTagIds              *Relation  `xmlrpc:"lead_tag_ids,omptempty"`
	LeadType                *Selection `xmlrpc:"lead_type,omptempty"`
	LeadUserId              *Many2One  `xmlrpc:"lead_user_id,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventLeadRules represents array of event.lead.rule model.
type EventLeadRules []EventLeadRule

// EventLeadRuleModel is the odoo model name.
const EventLeadRuleModel = "event.lead.rule"

// Many2One convert EventLeadRule to *Many2One.
func (elr *EventLeadRule) Many2One() *Many2One {
	return NewMany2One(elr.Id.Get(), "")
}

// CreateEventLeadRule creates a new event.lead.rule model and returns its id.
func (c *Client) CreateEventLeadRule(elr *EventLeadRule) (int64, error) {
	ids, err := c.CreateEventLeadRules([]*EventLeadRule{elr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventLeadRule creates a new event.lead.rule model and returns its id.
func (c *Client) CreateEventLeadRules(elrs []*EventLeadRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range elrs {
		vv = append(vv, v)
	}
	return c.Create(EventLeadRuleModel, vv)
}

// UpdateEventLeadRule updates an existing event.lead.rule record.
func (c *Client) UpdateEventLeadRule(elr *EventLeadRule) error {
	return c.UpdateEventLeadRules([]int64{elr.Id.Get()}, elr)
}

// UpdateEventLeadRules updates existing event.lead.rule records.
// All records (represented by ids) will be updated by elr values.
func (c *Client) UpdateEventLeadRules(ids []int64, elr *EventLeadRule) error {
	return c.Update(EventLeadRuleModel, ids, elr)
}

// DeleteEventLeadRule deletes an existing event.lead.rule record.
func (c *Client) DeleteEventLeadRule(id int64) error {
	return c.DeleteEventLeadRules([]int64{id})
}

// DeleteEventLeadRules deletes existing event.lead.rule records.
func (c *Client) DeleteEventLeadRules(ids []int64) error {
	return c.Delete(EventLeadRuleModel, ids)
}

// GetEventLeadRule gets event.lead.rule existing record.
func (c *Client) GetEventLeadRule(id int64) (*EventLeadRule, error) {
	elrs, err := c.GetEventLeadRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if elrs != nil && len(*elrs) > 0 {
		return &((*elrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.lead.rule not found", id)
}

// GetEventLeadRules gets event.lead.rule existing records.
func (c *Client) GetEventLeadRules(ids []int64) (*EventLeadRules, error) {
	elrs := &EventLeadRules{}
	if err := c.Read(EventLeadRuleModel, ids, nil, elrs); err != nil {
		return nil, err
	}
	return elrs, nil
}

// FindEventLeadRule finds event.lead.rule record by querying it with criteria.
func (c *Client) FindEventLeadRule(criteria *Criteria) (*EventLeadRule, error) {
	elrs := &EventLeadRules{}
	if err := c.SearchRead(EventLeadRuleModel, criteria, NewOptions().Limit(1), elrs); err != nil {
		return nil, err
	}
	if elrs != nil && len(*elrs) > 0 {
		return &((*elrs)[0]), nil
	}
	return nil, fmt.Errorf("event.lead.rule was not found with criteria %v", criteria)
}

// FindEventLeadRules finds event.lead.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventLeadRules(criteria *Criteria, options *Options) (*EventLeadRules, error) {
	elrs := &EventLeadRules{}
	if err := c.SearchRead(EventLeadRuleModel, criteria, options, elrs); err != nil {
		return nil, err
	}
	return elrs, nil
}

// FindEventLeadRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventLeadRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventLeadRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventLeadRuleId finds record id by querying it with criteria.
func (c *Client) FindEventLeadRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventLeadRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.lead.rule was not found with criteria %v and options %v", criteria, options)
}
