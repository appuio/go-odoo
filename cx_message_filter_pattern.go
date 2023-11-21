package odoo

import (
	"fmt"
)

// CxMessageFilterPattern represents cx.message.filter.pattern model.
type CxMessageFilterPattern struct {
	LastUpdate                *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omptempty"`
	DestinationRelatedModelId *Many2One `xmlrpc:"destination_related_model_id,omptempty"`
	DisplayName               *String   `xmlrpc:"display_name,omptempty"`
	FilterId                  *Many2One `xmlrpc:"filter_id,omptempty"`
	Id                        *Int      `xmlrpc:"id,omptempty"`
	ModelFieldId              *Many2One `xmlrpc:"model_field_id,omptempty"`
	PatternIds                *Relation `xmlrpc:"pattern_ids,omptempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CxMessageFilterPatterns represents array of cx.message.filter.pattern model.
type CxMessageFilterPatterns []CxMessageFilterPattern

// CxMessageFilterPatternModel is the odoo model name.
const CxMessageFilterPatternModel = "cx.message.filter.pattern"

// Many2One convert CxMessageFilterPattern to *Many2One.
func (cmfp *CxMessageFilterPattern) Many2One() *Many2One {
	return NewMany2One(cmfp.Id.Get(), "")
}

// CreateCxMessageFilterPattern creates a new cx.message.filter.pattern model and returns its id.
func (c *Client) CreateCxMessageFilterPattern(cmfp *CxMessageFilterPattern) (int64, error) {
	ids, err := c.CreateCxMessageFilterPatterns([]*CxMessageFilterPattern{cmfp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxMessageFilterPattern creates a new cx.message.filter.pattern model and returns its id.
func (c *Client) CreateCxMessageFilterPatterns(cmfps []*CxMessageFilterPattern) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmfps {
		vv = append(vv, v)
	}
	return c.Create(CxMessageFilterPatternModel, vv)
}

// UpdateCxMessageFilterPattern updates an existing cx.message.filter.pattern record.
func (c *Client) UpdateCxMessageFilterPattern(cmfp *CxMessageFilterPattern) error {
	return c.UpdateCxMessageFilterPatterns([]int64{cmfp.Id.Get()}, cmfp)
}

// UpdateCxMessageFilterPatterns updates existing cx.message.filter.pattern records.
// All records (represented by ids) will be updated by cmfp values.
func (c *Client) UpdateCxMessageFilterPatterns(ids []int64, cmfp *CxMessageFilterPattern) error {
	return c.Update(CxMessageFilterPatternModel, ids, cmfp)
}

// DeleteCxMessageFilterPattern deletes an existing cx.message.filter.pattern record.
func (c *Client) DeleteCxMessageFilterPattern(id int64) error {
	return c.DeleteCxMessageFilterPatterns([]int64{id})
}

// DeleteCxMessageFilterPatterns deletes existing cx.message.filter.pattern records.
func (c *Client) DeleteCxMessageFilterPatterns(ids []int64) error {
	return c.Delete(CxMessageFilterPatternModel, ids)
}

// GetCxMessageFilterPattern gets cx.message.filter.pattern existing record.
func (c *Client) GetCxMessageFilterPattern(id int64) (*CxMessageFilterPattern, error) {
	cmfps, err := c.GetCxMessageFilterPatterns([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmfps != nil && len(*cmfps) > 0 {
		return &((*cmfps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.message.filter.pattern not found", id)
}

// GetCxMessageFilterPatterns gets cx.message.filter.pattern existing records.
func (c *Client) GetCxMessageFilterPatterns(ids []int64) (*CxMessageFilterPatterns, error) {
	cmfps := &CxMessageFilterPatterns{}
	if err := c.Read(CxMessageFilterPatternModel, ids, nil, cmfps); err != nil {
		return nil, err
	}
	return cmfps, nil
}

// FindCxMessageFilterPattern finds cx.message.filter.pattern record by querying it with criteria.
func (c *Client) FindCxMessageFilterPattern(criteria *Criteria) (*CxMessageFilterPattern, error) {
	cmfps := &CxMessageFilterPatterns{}
	if err := c.SearchRead(CxMessageFilterPatternModel, criteria, NewOptions().Limit(1), cmfps); err != nil {
		return nil, err
	}
	if cmfps != nil && len(*cmfps) > 0 {
		return &((*cmfps)[0]), nil
	}
	return nil, fmt.Errorf("cx.message.filter.pattern was not found with criteria %v", criteria)
}

// FindCxMessageFilterPatterns finds cx.message.filter.pattern records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterPatterns(criteria *Criteria, options *Options) (*CxMessageFilterPatterns, error) {
	cmfps := &CxMessageFilterPatterns{}
	if err := c.SearchRead(CxMessageFilterPatternModel, criteria, options, cmfps); err != nil {
		return nil, err
	}
	return cmfps, nil
}

// FindCxMessageFilterPatternIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterPatternIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxMessageFilterPatternModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxMessageFilterPatternId finds record id by querying it with criteria.
func (c *Client) FindCxMessageFilterPatternId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxMessageFilterPatternModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.message.filter.pattern was not found with criteria %v and options %v", criteria, options)
}
