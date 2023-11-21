package odoo

import (
	"fmt"
)

// MukRestAccessRulesExpression represents muk_rest.access_rules.expression model.
type MukRestAccessRulesExpression struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Expression  *String    `xmlrpc:"expression,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	Operation   *Selection `xmlrpc:"operation,omptempty"`
	Param       *String    `xmlrpc:"param,omptempty"`
	RuleId      *Many2One  `xmlrpc:"rule_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MukRestAccessRulesExpressions represents array of muk_rest.access_rules.expression model.
type MukRestAccessRulesExpressions []MukRestAccessRulesExpression

// MukRestAccessRulesExpressionModel is the odoo model name.
const MukRestAccessRulesExpressionModel = "muk_rest.access_rules.expression"

// Many2One convert MukRestAccessRulesExpression to *Many2One.
func (mae *MukRestAccessRulesExpression) Many2One() *Many2One {
	return NewMany2One(mae.Id.Get(), "")
}

// CreateMukRestAccessRulesExpression creates a new muk_rest.access_rules.expression model and returns its id.
func (c *Client) CreateMukRestAccessRulesExpression(mae *MukRestAccessRulesExpression) (int64, error) {
	ids, err := c.CreateMukRestAccessRulesExpressions([]*MukRestAccessRulesExpression{mae})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestAccessRulesExpression creates a new muk_rest.access_rules.expression model and returns its id.
func (c *Client) CreateMukRestAccessRulesExpressions(maes []*MukRestAccessRulesExpression) ([]int64, error) {
	var vv []interface{}
	for _, v := range maes {
		vv = append(vv, v)
	}
	return c.Create(MukRestAccessRulesExpressionModel, vv)
}

// UpdateMukRestAccessRulesExpression updates an existing muk_rest.access_rules.expression record.
func (c *Client) UpdateMukRestAccessRulesExpression(mae *MukRestAccessRulesExpression) error {
	return c.UpdateMukRestAccessRulesExpressions([]int64{mae.Id.Get()}, mae)
}

// UpdateMukRestAccessRulesExpressions updates existing muk_rest.access_rules.expression records.
// All records (represented by ids) will be updated by mae values.
func (c *Client) UpdateMukRestAccessRulesExpressions(ids []int64, mae *MukRestAccessRulesExpression) error {
	return c.Update(MukRestAccessRulesExpressionModel, ids, mae)
}

// DeleteMukRestAccessRulesExpression deletes an existing muk_rest.access_rules.expression record.
func (c *Client) DeleteMukRestAccessRulesExpression(id int64) error {
	return c.DeleteMukRestAccessRulesExpressions([]int64{id})
}

// DeleteMukRestAccessRulesExpressions deletes existing muk_rest.access_rules.expression records.
func (c *Client) DeleteMukRestAccessRulesExpressions(ids []int64) error {
	return c.Delete(MukRestAccessRulesExpressionModel, ids)
}

// GetMukRestAccessRulesExpression gets muk_rest.access_rules.expression existing record.
func (c *Client) GetMukRestAccessRulesExpression(id int64) (*MukRestAccessRulesExpression, error) {
	maes, err := c.GetMukRestAccessRulesExpressions([]int64{id})
	if err != nil {
		return nil, err
	}
	if maes != nil && len(*maes) > 0 {
		return &((*maes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.access_rules.expression not found", id)
}

// GetMukRestAccessRulesExpressions gets muk_rest.access_rules.expression existing records.
func (c *Client) GetMukRestAccessRulesExpressions(ids []int64) (*MukRestAccessRulesExpressions, error) {
	maes := &MukRestAccessRulesExpressions{}
	if err := c.Read(MukRestAccessRulesExpressionModel, ids, nil, maes); err != nil {
		return nil, err
	}
	return maes, nil
}

// FindMukRestAccessRulesExpression finds muk_rest.access_rules.expression record by querying it with criteria.
func (c *Client) FindMukRestAccessRulesExpression(criteria *Criteria) (*MukRestAccessRulesExpression, error) {
	maes := &MukRestAccessRulesExpressions{}
	if err := c.SearchRead(MukRestAccessRulesExpressionModel, criteria, NewOptions().Limit(1), maes); err != nil {
		return nil, err
	}
	if maes != nil && len(*maes) > 0 {
		return &((*maes)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.access_rules.expression was not found with criteria %v", criteria)
}

// FindMukRestAccessRulesExpressions finds muk_rest.access_rules.expression records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAccessRulesExpressions(criteria *Criteria, options *Options) (*MukRestAccessRulesExpressions, error) {
	maes := &MukRestAccessRulesExpressions{}
	if err := c.SearchRead(MukRestAccessRulesExpressionModel, criteria, options, maes); err != nil {
		return nil, err
	}
	return maes, nil
}

// FindMukRestAccessRulesExpressionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAccessRulesExpressionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestAccessRulesExpressionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestAccessRulesExpressionId finds record id by querying it with criteria.
func (c *Client) FindMukRestAccessRulesExpressionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestAccessRulesExpressionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.access_rules.expression was not found with criteria %v and options %v", criteria, options)
}
