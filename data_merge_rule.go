package odoo

import (
	"fmt"
)

// DataMergeRule represents data_merge.rule model.
type DataMergeRule struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	FieldId     *Many2One  `xmlrpc:"field_id,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	MatchMode   *Selection `xmlrpc:"match_mode,omptempty"`
	ModelId     *Many2One  `xmlrpc:"model_id,omptempty"`
	ResModelId  *Many2One  `xmlrpc:"res_model_id,omptempty"`
	Sequence    *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DataMergeRules represents array of data_merge.rule model.
type DataMergeRules []DataMergeRule

// DataMergeRuleModel is the odoo model name.
const DataMergeRuleModel = "data_merge.rule"

// Many2One convert DataMergeRule to *Many2One.
func (dr *DataMergeRule) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDataMergeRule creates a new data_merge.rule model and returns its id.
func (c *Client) CreateDataMergeRule(dr *DataMergeRule) (int64, error) {
	ids, err := c.CreateDataMergeRules([]*DataMergeRule{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataMergeRule creates a new data_merge.rule model and returns its id.
func (c *Client) CreateDataMergeRules(drs []*DataMergeRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DataMergeRuleModel, vv)
}

// UpdateDataMergeRule updates an existing data_merge.rule record.
func (c *Client) UpdateDataMergeRule(dr *DataMergeRule) error {
	return c.UpdateDataMergeRules([]int64{dr.Id.Get()}, dr)
}

// UpdateDataMergeRules updates existing data_merge.rule records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDataMergeRules(ids []int64, dr *DataMergeRule) error {
	return c.Update(DataMergeRuleModel, ids, dr)
}

// DeleteDataMergeRule deletes an existing data_merge.rule record.
func (c *Client) DeleteDataMergeRule(id int64) error {
	return c.DeleteDataMergeRules([]int64{id})
}

// DeleteDataMergeRules deletes existing data_merge.rule records.
func (c *Client) DeleteDataMergeRules(ids []int64) error {
	return c.Delete(DataMergeRuleModel, ids)
}

// GetDataMergeRule gets data_merge.rule existing record.
func (c *Client) GetDataMergeRule(id int64) (*DataMergeRule, error) {
	drs, err := c.GetDataMergeRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_merge.rule not found", id)
}

// GetDataMergeRules gets data_merge.rule existing records.
func (c *Client) GetDataMergeRules(ids []int64) (*DataMergeRules, error) {
	drs := &DataMergeRules{}
	if err := c.Read(DataMergeRuleModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataMergeRule finds data_merge.rule record by querying it with criteria.
func (c *Client) FindDataMergeRule(criteria *Criteria) (*DataMergeRule, error) {
	drs := &DataMergeRules{}
	if err := c.SearchRead(DataMergeRuleModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("data_merge.rule was not found with criteria %v", criteria)
}

// FindDataMergeRules finds data_merge.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeRules(criteria *Criteria, options *Options) (*DataMergeRules, error) {
	drs := &DataMergeRules{}
	if err := c.SearchRead(DataMergeRuleModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataMergeRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataMergeRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataMergeRuleId finds record id by querying it with criteria.
func (c *Client) FindDataMergeRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataMergeRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_merge.rule was not found with criteria %v and options %v", criteria, options)
}
