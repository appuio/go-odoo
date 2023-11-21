package odoo

import (
	"fmt"
)

// CxMessageFilterCondition represents cx.message.filter.condition model.
type CxMessageFilterCondition struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Condition   *Selection `xmlrpc:"condition,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	PartnerIds  *Relation  `xmlrpc:"partner_ids,omptempty"`
	RuleId      *Many2One  `xmlrpc:"rule_id,omptempty"`
	Trigger     *Selection `xmlrpc:"trigger,omptempty"`
	Value       *String    `xmlrpc:"value,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CxMessageFilterConditions represents array of cx.message.filter.condition model.
type CxMessageFilterConditions []CxMessageFilterCondition

// CxMessageFilterConditionModel is the odoo model name.
const CxMessageFilterConditionModel = "cx.message.filter.condition"

// Many2One convert CxMessageFilterCondition to *Many2One.
func (cmfc *CxMessageFilterCondition) Many2One() *Many2One {
	return NewMany2One(cmfc.Id.Get(), "")
}

// CreateCxMessageFilterCondition creates a new cx.message.filter.condition model and returns its id.
func (c *Client) CreateCxMessageFilterCondition(cmfc *CxMessageFilterCondition) (int64, error) {
	ids, err := c.CreateCxMessageFilterConditions([]*CxMessageFilterCondition{cmfc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxMessageFilterCondition creates a new cx.message.filter.condition model and returns its id.
func (c *Client) CreateCxMessageFilterConditions(cmfcs []*CxMessageFilterCondition) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmfcs {
		vv = append(vv, v)
	}
	return c.Create(CxMessageFilterConditionModel, vv)
}

// UpdateCxMessageFilterCondition updates an existing cx.message.filter.condition record.
func (c *Client) UpdateCxMessageFilterCondition(cmfc *CxMessageFilterCondition) error {
	return c.UpdateCxMessageFilterConditions([]int64{cmfc.Id.Get()}, cmfc)
}

// UpdateCxMessageFilterConditions updates existing cx.message.filter.condition records.
// All records (represented by ids) will be updated by cmfc values.
func (c *Client) UpdateCxMessageFilterConditions(ids []int64, cmfc *CxMessageFilterCondition) error {
	return c.Update(CxMessageFilterConditionModel, ids, cmfc)
}

// DeleteCxMessageFilterCondition deletes an existing cx.message.filter.condition record.
func (c *Client) DeleteCxMessageFilterCondition(id int64) error {
	return c.DeleteCxMessageFilterConditions([]int64{id})
}

// DeleteCxMessageFilterConditions deletes existing cx.message.filter.condition records.
func (c *Client) DeleteCxMessageFilterConditions(ids []int64) error {
	return c.Delete(CxMessageFilterConditionModel, ids)
}

// GetCxMessageFilterCondition gets cx.message.filter.condition existing record.
func (c *Client) GetCxMessageFilterCondition(id int64) (*CxMessageFilterCondition, error) {
	cmfcs, err := c.GetCxMessageFilterConditions([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmfcs != nil && len(*cmfcs) > 0 {
		return &((*cmfcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.message.filter.condition not found", id)
}

// GetCxMessageFilterConditions gets cx.message.filter.condition existing records.
func (c *Client) GetCxMessageFilterConditions(ids []int64) (*CxMessageFilterConditions, error) {
	cmfcs := &CxMessageFilterConditions{}
	if err := c.Read(CxMessageFilterConditionModel, ids, nil, cmfcs); err != nil {
		return nil, err
	}
	return cmfcs, nil
}

// FindCxMessageFilterCondition finds cx.message.filter.condition record by querying it with criteria.
func (c *Client) FindCxMessageFilterCondition(criteria *Criteria) (*CxMessageFilterCondition, error) {
	cmfcs := &CxMessageFilterConditions{}
	if err := c.SearchRead(CxMessageFilterConditionModel, criteria, NewOptions().Limit(1), cmfcs); err != nil {
		return nil, err
	}
	if cmfcs != nil && len(*cmfcs) > 0 {
		return &((*cmfcs)[0]), nil
	}
	return nil, fmt.Errorf("cx.message.filter.condition was not found with criteria %v", criteria)
}

// FindCxMessageFilterConditions finds cx.message.filter.condition records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterConditions(criteria *Criteria, options *Options) (*CxMessageFilterConditions, error) {
	cmfcs := &CxMessageFilterConditions{}
	if err := c.SearchRead(CxMessageFilterConditionModel, criteria, options, cmfcs); err != nil {
		return nil, err
	}
	return cmfcs, nil
}

// FindCxMessageFilterConditionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterConditionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxMessageFilterConditionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxMessageFilterConditionId finds record id by querying it with criteria.
func (c *Client) FindCxMessageFilterConditionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxMessageFilterConditionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.message.filter.condition was not found with criteria %v and options %v", criteria, options)
}
