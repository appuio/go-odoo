package odoo

import (
	"fmt"
)

// CxTemplatePattern represents cx.template.pattern model.
type CxTemplatePattern struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CxTemplatePatterns represents array of cx.template.pattern model.
type CxTemplatePatterns []CxTemplatePattern

// CxTemplatePatternModel is the odoo model name.
const CxTemplatePatternModel = "cx.template.pattern"

// Many2One convert CxTemplatePattern to *Many2One.
func (ctp *CxTemplatePattern) Many2One() *Many2One {
	return NewMany2One(ctp.Id.Get(), "")
}

// CreateCxTemplatePattern creates a new cx.template.pattern model and returns its id.
func (c *Client) CreateCxTemplatePattern(ctp *CxTemplatePattern) (int64, error) {
	ids, err := c.CreateCxTemplatePatterns([]*CxTemplatePattern{ctp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxTemplatePattern creates a new cx.template.pattern model and returns its id.
func (c *Client) CreateCxTemplatePatterns(ctps []*CxTemplatePattern) ([]int64, error) {
	var vv []interface{}
	for _, v := range ctps {
		vv = append(vv, v)
	}
	return c.Create(CxTemplatePatternModel, vv)
}

// UpdateCxTemplatePattern updates an existing cx.template.pattern record.
func (c *Client) UpdateCxTemplatePattern(ctp *CxTemplatePattern) error {
	return c.UpdateCxTemplatePatterns([]int64{ctp.Id.Get()}, ctp)
}

// UpdateCxTemplatePatterns updates existing cx.template.pattern records.
// All records (represented by ids) will be updated by ctp values.
func (c *Client) UpdateCxTemplatePatterns(ids []int64, ctp *CxTemplatePattern) error {
	return c.Update(CxTemplatePatternModel, ids, ctp)
}

// DeleteCxTemplatePattern deletes an existing cx.template.pattern record.
func (c *Client) DeleteCxTemplatePattern(id int64) error {
	return c.DeleteCxTemplatePatterns([]int64{id})
}

// DeleteCxTemplatePatterns deletes existing cx.template.pattern records.
func (c *Client) DeleteCxTemplatePatterns(ids []int64) error {
	return c.Delete(CxTemplatePatternModel, ids)
}

// GetCxTemplatePattern gets cx.template.pattern existing record.
func (c *Client) GetCxTemplatePattern(id int64) (*CxTemplatePattern, error) {
	ctps, err := c.GetCxTemplatePatterns([]int64{id})
	if err != nil {
		return nil, err
	}
	if ctps != nil && len(*ctps) > 0 {
		return &((*ctps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.template.pattern not found", id)
}

// GetCxTemplatePatterns gets cx.template.pattern existing records.
func (c *Client) GetCxTemplatePatterns(ids []int64) (*CxTemplatePatterns, error) {
	ctps := &CxTemplatePatterns{}
	if err := c.Read(CxTemplatePatternModel, ids, nil, ctps); err != nil {
		return nil, err
	}
	return ctps, nil
}

// FindCxTemplatePattern finds cx.template.pattern record by querying it with criteria.
func (c *Client) FindCxTemplatePattern(criteria *Criteria) (*CxTemplatePattern, error) {
	ctps := &CxTemplatePatterns{}
	if err := c.SearchRead(CxTemplatePatternModel, criteria, NewOptions().Limit(1), ctps); err != nil {
		return nil, err
	}
	if ctps != nil && len(*ctps) > 0 {
		return &((*ctps)[0]), nil
	}
	return nil, fmt.Errorf("cx.template.pattern was not found with criteria %v", criteria)
}

// FindCxTemplatePatterns finds cx.template.pattern records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxTemplatePatterns(criteria *Criteria, options *Options) (*CxTemplatePatterns, error) {
	ctps := &CxTemplatePatterns{}
	if err := c.SearchRead(CxTemplatePatternModel, criteria, options, ctps); err != nil {
		return nil, err
	}
	return ctps, nil
}

// FindCxTemplatePatternIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxTemplatePatternIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxTemplatePatternModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxTemplatePatternId finds record id by querying it with criteria.
func (c *Client) FindCxTemplatePatternId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxTemplatePatternModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.template.pattern was not found with criteria %v and options %v", criteria, options)
}
