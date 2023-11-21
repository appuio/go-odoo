package odoo

import (
	"fmt"
)

// CxMessageFilterRule represents cx.message.filter.rule model.
type CxMessageFilterRule struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	ConditionIds *Relation `xmlrpc:"condition_ids,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	FilterId     *Many2One `xmlrpc:"filter_id,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	MessageId    *Int      `xmlrpc:"message_id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CxMessageFilterRules represents array of cx.message.filter.rule model.
type CxMessageFilterRules []CxMessageFilterRule

// CxMessageFilterRuleModel is the odoo model name.
const CxMessageFilterRuleModel = "cx.message.filter.rule"

// Many2One convert CxMessageFilterRule to *Many2One.
func (cmfr *CxMessageFilterRule) Many2One() *Many2One {
	return NewMany2One(cmfr.Id.Get(), "")
}

// CreateCxMessageFilterRule creates a new cx.message.filter.rule model and returns its id.
func (c *Client) CreateCxMessageFilterRule(cmfr *CxMessageFilterRule) (int64, error) {
	ids, err := c.CreateCxMessageFilterRules([]*CxMessageFilterRule{cmfr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxMessageFilterRule creates a new cx.message.filter.rule model and returns its id.
func (c *Client) CreateCxMessageFilterRules(cmfrs []*CxMessageFilterRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmfrs {
		vv = append(vv, v)
	}
	return c.Create(CxMessageFilterRuleModel, vv)
}

// UpdateCxMessageFilterRule updates an existing cx.message.filter.rule record.
func (c *Client) UpdateCxMessageFilterRule(cmfr *CxMessageFilterRule) error {
	return c.UpdateCxMessageFilterRules([]int64{cmfr.Id.Get()}, cmfr)
}

// UpdateCxMessageFilterRules updates existing cx.message.filter.rule records.
// All records (represented by ids) will be updated by cmfr values.
func (c *Client) UpdateCxMessageFilterRules(ids []int64, cmfr *CxMessageFilterRule) error {
	return c.Update(CxMessageFilterRuleModel, ids, cmfr)
}

// DeleteCxMessageFilterRule deletes an existing cx.message.filter.rule record.
func (c *Client) DeleteCxMessageFilterRule(id int64) error {
	return c.DeleteCxMessageFilterRules([]int64{id})
}

// DeleteCxMessageFilterRules deletes existing cx.message.filter.rule records.
func (c *Client) DeleteCxMessageFilterRules(ids []int64) error {
	return c.Delete(CxMessageFilterRuleModel, ids)
}

// GetCxMessageFilterRule gets cx.message.filter.rule existing record.
func (c *Client) GetCxMessageFilterRule(id int64) (*CxMessageFilterRule, error) {
	cmfrs, err := c.GetCxMessageFilterRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmfrs != nil && len(*cmfrs) > 0 {
		return &((*cmfrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.message.filter.rule not found", id)
}

// GetCxMessageFilterRules gets cx.message.filter.rule existing records.
func (c *Client) GetCxMessageFilterRules(ids []int64) (*CxMessageFilterRules, error) {
	cmfrs := &CxMessageFilterRules{}
	if err := c.Read(CxMessageFilterRuleModel, ids, nil, cmfrs); err != nil {
		return nil, err
	}
	return cmfrs, nil
}

// FindCxMessageFilterRule finds cx.message.filter.rule record by querying it with criteria.
func (c *Client) FindCxMessageFilterRule(criteria *Criteria) (*CxMessageFilterRule, error) {
	cmfrs := &CxMessageFilterRules{}
	if err := c.SearchRead(CxMessageFilterRuleModel, criteria, NewOptions().Limit(1), cmfrs); err != nil {
		return nil, err
	}
	if cmfrs != nil && len(*cmfrs) > 0 {
		return &((*cmfrs)[0]), nil
	}
	return nil, fmt.Errorf("cx.message.filter.rule was not found with criteria %v", criteria)
}

// FindCxMessageFilterRules finds cx.message.filter.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterRules(criteria *Criteria, options *Options) (*CxMessageFilterRules, error) {
	cmfrs := &CxMessageFilterRules{}
	if err := c.SearchRead(CxMessageFilterRuleModel, criteria, options, cmfrs); err != nil {
		return nil, err
	}
	return cmfrs, nil
}

// FindCxMessageFilterRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxMessageFilterRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxMessageFilterRuleId finds record id by querying it with criteria.
func (c *Client) FindCxMessageFilterRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxMessageFilterRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.message.filter.rule was not found with criteria %v and options %v", criteria, options)
}
