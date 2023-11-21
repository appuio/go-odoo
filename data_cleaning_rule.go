package odoo

import (
	"fmt"
)

// DataCleaningRule represents data_cleaning.rule model.
type DataCleaningRule struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Action          *Selection `xmlrpc:"action,omptempty"`
	ActionCase      *Selection `xmlrpc:"action_case,omptempty"`
	ActionDisplay   *String    `xmlrpc:"action_display,omptempty"`
	ActionTechnical *String    `xmlrpc:"action_technical,omptempty"`
	ActionTrim      *Selection `xmlrpc:"action_trim,omptempty"`
	CleaningModelId *Many2One  `xmlrpc:"cleaning_model_id,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	FieldId         *Many2One  `xmlrpc:"field_id,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	ResModelId      *Many2One  `xmlrpc:"res_model_id,omptempty"`
	ResModelName    *String    `xmlrpc:"res_model_name,omptempty"`
	Sequence        *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DataCleaningRules represents array of data_cleaning.rule model.
type DataCleaningRules []DataCleaningRule

// DataCleaningRuleModel is the odoo model name.
const DataCleaningRuleModel = "data_cleaning.rule"

// Many2One convert DataCleaningRule to *Many2One.
func (dr *DataCleaningRule) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDataCleaningRule creates a new data_cleaning.rule model and returns its id.
func (c *Client) CreateDataCleaningRule(dr *DataCleaningRule) (int64, error) {
	ids, err := c.CreateDataCleaningRules([]*DataCleaningRule{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataCleaningRule creates a new data_cleaning.rule model and returns its id.
func (c *Client) CreateDataCleaningRules(drs []*DataCleaningRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DataCleaningRuleModel, vv)
}

// UpdateDataCleaningRule updates an existing data_cleaning.rule record.
func (c *Client) UpdateDataCleaningRule(dr *DataCleaningRule) error {
	return c.UpdateDataCleaningRules([]int64{dr.Id.Get()}, dr)
}

// UpdateDataCleaningRules updates existing data_cleaning.rule records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDataCleaningRules(ids []int64, dr *DataCleaningRule) error {
	return c.Update(DataCleaningRuleModel, ids, dr)
}

// DeleteDataCleaningRule deletes an existing data_cleaning.rule record.
func (c *Client) DeleteDataCleaningRule(id int64) error {
	return c.DeleteDataCleaningRules([]int64{id})
}

// DeleteDataCleaningRules deletes existing data_cleaning.rule records.
func (c *Client) DeleteDataCleaningRules(ids []int64) error {
	return c.Delete(DataCleaningRuleModel, ids)
}

// GetDataCleaningRule gets data_cleaning.rule existing record.
func (c *Client) GetDataCleaningRule(id int64) (*DataCleaningRule, error) {
	drs, err := c.GetDataCleaningRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_cleaning.rule not found", id)
}

// GetDataCleaningRules gets data_cleaning.rule existing records.
func (c *Client) GetDataCleaningRules(ids []int64) (*DataCleaningRules, error) {
	drs := &DataCleaningRules{}
	if err := c.Read(DataCleaningRuleModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataCleaningRule finds data_cleaning.rule record by querying it with criteria.
func (c *Client) FindDataCleaningRule(criteria *Criteria) (*DataCleaningRule, error) {
	drs := &DataCleaningRules{}
	if err := c.SearchRead(DataCleaningRuleModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("data_cleaning.rule was not found with criteria %v", criteria)
}

// FindDataCleaningRules finds data_cleaning.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataCleaningRules(criteria *Criteria, options *Options) (*DataCleaningRules, error) {
	drs := &DataCleaningRules{}
	if err := c.SearchRead(DataCleaningRuleModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataCleaningRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataCleaningRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataCleaningRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataCleaningRuleId finds record id by querying it with criteria.
func (c *Client) FindDataCleaningRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataCleaningRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_cleaning.rule was not found with criteria %v and options %v", criteria, options)
}
