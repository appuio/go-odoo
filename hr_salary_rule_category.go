package odoo

import (
	"fmt"
)

// HrSalaryRuleCategory represents hr.salary.rule.category model.
type HrSalaryRuleCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	ChildrenIds *Relation `xmlrpc:"children_ids,omptempty"`
	Code        *String   `xmlrpc:"code,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Note        *String   `xmlrpc:"note,omptempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrSalaryRuleCategorys represents array of hr.salary.rule.category model.
type HrSalaryRuleCategorys []HrSalaryRuleCategory

// HrSalaryRuleCategoryModel is the odoo model name.
const HrSalaryRuleCategoryModel = "hr.salary.rule.category"

// Many2One convert HrSalaryRuleCategory to *Many2One.
func (hsrc *HrSalaryRuleCategory) Many2One() *Many2One {
	return NewMany2One(hsrc.Id.Get(), "")
}

// CreateHrSalaryRuleCategory creates a new hr.salary.rule.category model and returns its id.
func (c *Client) CreateHrSalaryRuleCategory(hsrc *HrSalaryRuleCategory) (int64, error) {
	ids, err := c.CreateHrSalaryRuleCategorys([]*HrSalaryRuleCategory{hsrc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSalaryRuleCategory creates a new hr.salary.rule.category model and returns its id.
func (c *Client) CreateHrSalaryRuleCategorys(hsrcs []*HrSalaryRuleCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsrcs {
		vv = append(vv, v)
	}
	return c.Create(HrSalaryRuleCategoryModel, vv)
}

// UpdateHrSalaryRuleCategory updates an existing hr.salary.rule.category record.
func (c *Client) UpdateHrSalaryRuleCategory(hsrc *HrSalaryRuleCategory) error {
	return c.UpdateHrSalaryRuleCategorys([]int64{hsrc.Id.Get()}, hsrc)
}

// UpdateHrSalaryRuleCategorys updates existing hr.salary.rule.category records.
// All records (represented by ids) will be updated by hsrc values.
func (c *Client) UpdateHrSalaryRuleCategorys(ids []int64, hsrc *HrSalaryRuleCategory) error {
	return c.Update(HrSalaryRuleCategoryModel, ids, hsrc)
}

// DeleteHrSalaryRuleCategory deletes an existing hr.salary.rule.category record.
func (c *Client) DeleteHrSalaryRuleCategory(id int64) error {
	return c.DeleteHrSalaryRuleCategorys([]int64{id})
}

// DeleteHrSalaryRuleCategorys deletes existing hr.salary.rule.category records.
func (c *Client) DeleteHrSalaryRuleCategorys(ids []int64) error {
	return c.Delete(HrSalaryRuleCategoryModel, ids)
}

// GetHrSalaryRuleCategory gets hr.salary.rule.category existing record.
func (c *Client) GetHrSalaryRuleCategory(id int64) (*HrSalaryRuleCategory, error) {
	hsrcs, err := c.GetHrSalaryRuleCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsrcs != nil && len(*hsrcs) > 0 {
		return &((*hsrcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.salary.rule.category not found", id)
}

// GetHrSalaryRuleCategorys gets hr.salary.rule.category existing records.
func (c *Client) GetHrSalaryRuleCategorys(ids []int64) (*HrSalaryRuleCategorys, error) {
	hsrcs := &HrSalaryRuleCategorys{}
	if err := c.Read(HrSalaryRuleCategoryModel, ids, nil, hsrcs); err != nil {
		return nil, err
	}
	return hsrcs, nil
}

// FindHrSalaryRuleCategory finds hr.salary.rule.category record by querying it with criteria.
func (c *Client) FindHrSalaryRuleCategory(criteria *Criteria) (*HrSalaryRuleCategory, error) {
	hsrcs := &HrSalaryRuleCategorys{}
	if err := c.SearchRead(HrSalaryRuleCategoryModel, criteria, NewOptions().Limit(1), hsrcs); err != nil {
		return nil, err
	}
	if hsrcs != nil && len(*hsrcs) > 0 {
		return &((*hsrcs)[0]), nil
	}
	return nil, fmt.Errorf("hr.salary.rule.category was not found with criteria %v", criteria)
}

// FindHrSalaryRuleCategorys finds hr.salary.rule.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryRuleCategorys(criteria *Criteria, options *Options) (*HrSalaryRuleCategorys, error) {
	hsrcs := &HrSalaryRuleCategorys{}
	if err := c.SearchRead(HrSalaryRuleCategoryModel, criteria, options, hsrcs); err != nil {
		return nil, err
	}
	return hsrcs, nil
}

// FindHrSalaryRuleCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryRuleCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrSalaryRuleCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrSalaryRuleCategoryId finds record id by querying it with criteria.
func (c *Client) FindHrSalaryRuleCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSalaryRuleCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.salary.rule.category was not found with criteria %v and options %v", criteria, options)
}
