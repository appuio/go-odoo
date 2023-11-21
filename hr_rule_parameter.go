package odoo

import (
	"fmt"
)

// HrRuleParameter represents hr.rule.parameter model.
type HrRuleParameter struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	Code                *String   `xmlrpc:"code,omptempty"`
	CountryId           *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	Description         *String   `xmlrpc:"description,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	ParameterVersionIds *Relation `xmlrpc:"parameter_version_ids,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrRuleParameters represents array of hr.rule.parameter model.
type HrRuleParameters []HrRuleParameter

// HrRuleParameterModel is the odoo model name.
const HrRuleParameterModel = "hr.rule.parameter"

// Many2One convert HrRuleParameter to *Many2One.
func (hrp *HrRuleParameter) Many2One() *Many2One {
	return NewMany2One(hrp.Id.Get(), "")
}

// CreateHrRuleParameter creates a new hr.rule.parameter model and returns its id.
func (c *Client) CreateHrRuleParameter(hrp *HrRuleParameter) (int64, error) {
	ids, err := c.CreateHrRuleParameters([]*HrRuleParameter{hrp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrRuleParameter creates a new hr.rule.parameter model and returns its id.
func (c *Client) CreateHrRuleParameters(hrps []*HrRuleParameter) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrps {
		vv = append(vv, v)
	}
	return c.Create(HrRuleParameterModel, vv)
}

// UpdateHrRuleParameter updates an existing hr.rule.parameter record.
func (c *Client) UpdateHrRuleParameter(hrp *HrRuleParameter) error {
	return c.UpdateHrRuleParameters([]int64{hrp.Id.Get()}, hrp)
}

// UpdateHrRuleParameters updates existing hr.rule.parameter records.
// All records (represented by ids) will be updated by hrp values.
func (c *Client) UpdateHrRuleParameters(ids []int64, hrp *HrRuleParameter) error {
	return c.Update(HrRuleParameterModel, ids, hrp)
}

// DeleteHrRuleParameter deletes an existing hr.rule.parameter record.
func (c *Client) DeleteHrRuleParameter(id int64) error {
	return c.DeleteHrRuleParameters([]int64{id})
}

// DeleteHrRuleParameters deletes existing hr.rule.parameter records.
func (c *Client) DeleteHrRuleParameters(ids []int64) error {
	return c.Delete(HrRuleParameterModel, ids)
}

// GetHrRuleParameter gets hr.rule.parameter existing record.
func (c *Client) GetHrRuleParameter(id int64) (*HrRuleParameter, error) {
	hrps, err := c.GetHrRuleParameters([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrps != nil && len(*hrps) > 0 {
		return &((*hrps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.rule.parameter not found", id)
}

// GetHrRuleParameters gets hr.rule.parameter existing records.
func (c *Client) GetHrRuleParameters(ids []int64) (*HrRuleParameters, error) {
	hrps := &HrRuleParameters{}
	if err := c.Read(HrRuleParameterModel, ids, nil, hrps); err != nil {
		return nil, err
	}
	return hrps, nil
}

// FindHrRuleParameter finds hr.rule.parameter record by querying it with criteria.
func (c *Client) FindHrRuleParameter(criteria *Criteria) (*HrRuleParameter, error) {
	hrps := &HrRuleParameters{}
	if err := c.SearchRead(HrRuleParameterModel, criteria, NewOptions().Limit(1), hrps); err != nil {
		return nil, err
	}
	if hrps != nil && len(*hrps) > 0 {
		return &((*hrps)[0]), nil
	}
	return nil, fmt.Errorf("hr.rule.parameter was not found with criteria %v", criteria)
}

// FindHrRuleParameters finds hr.rule.parameter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRuleParameters(criteria *Criteria, options *Options) (*HrRuleParameters, error) {
	hrps := &HrRuleParameters{}
	if err := c.SearchRead(HrRuleParameterModel, criteria, options, hrps); err != nil {
		return nil, err
	}
	return hrps, nil
}

// FindHrRuleParameterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRuleParameterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrRuleParameterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrRuleParameterId finds record id by querying it with criteria.
func (c *Client) FindHrRuleParameterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRuleParameterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.rule.parameter was not found with criteria %v and options %v", criteria, options)
}
