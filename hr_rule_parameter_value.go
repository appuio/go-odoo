package odoo

import (
	"fmt"
)

// HrRuleParameterValue represents hr.rule.parameter.value model.
type HrRuleParameterValue struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	Code              *String   `xmlrpc:"code,omptempty"`
	CountryId         *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DateFrom          *Time     `xmlrpc:"date_from,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	ParameterValue    *String   `xmlrpc:"parameter_value,omptempty"`
	RuleParameterId   *Many2One `xmlrpc:"rule_parameter_id,omptempty"`
	RuleParameterName *String   `xmlrpc:"rule_parameter_name,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrRuleParameterValues represents array of hr.rule.parameter.value model.
type HrRuleParameterValues []HrRuleParameterValue

// HrRuleParameterValueModel is the odoo model name.
const HrRuleParameterValueModel = "hr.rule.parameter.value"

// Many2One convert HrRuleParameterValue to *Many2One.
func (hrpv *HrRuleParameterValue) Many2One() *Many2One {
	return NewMany2One(hrpv.Id.Get(), "")
}

// CreateHrRuleParameterValue creates a new hr.rule.parameter.value model and returns its id.
func (c *Client) CreateHrRuleParameterValue(hrpv *HrRuleParameterValue) (int64, error) {
	ids, err := c.CreateHrRuleParameterValues([]*HrRuleParameterValue{hrpv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrRuleParameterValue creates a new hr.rule.parameter.value model and returns its id.
func (c *Client) CreateHrRuleParameterValues(hrpvs []*HrRuleParameterValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrpvs {
		vv = append(vv, v)
	}
	return c.Create(HrRuleParameterValueModel, vv)
}

// UpdateHrRuleParameterValue updates an existing hr.rule.parameter.value record.
func (c *Client) UpdateHrRuleParameterValue(hrpv *HrRuleParameterValue) error {
	return c.UpdateHrRuleParameterValues([]int64{hrpv.Id.Get()}, hrpv)
}

// UpdateHrRuleParameterValues updates existing hr.rule.parameter.value records.
// All records (represented by ids) will be updated by hrpv values.
func (c *Client) UpdateHrRuleParameterValues(ids []int64, hrpv *HrRuleParameterValue) error {
	return c.Update(HrRuleParameterValueModel, ids, hrpv)
}

// DeleteHrRuleParameterValue deletes an existing hr.rule.parameter.value record.
func (c *Client) DeleteHrRuleParameterValue(id int64) error {
	return c.DeleteHrRuleParameterValues([]int64{id})
}

// DeleteHrRuleParameterValues deletes existing hr.rule.parameter.value records.
func (c *Client) DeleteHrRuleParameterValues(ids []int64) error {
	return c.Delete(HrRuleParameterValueModel, ids)
}

// GetHrRuleParameterValue gets hr.rule.parameter.value existing record.
func (c *Client) GetHrRuleParameterValue(id int64) (*HrRuleParameterValue, error) {
	hrpvs, err := c.GetHrRuleParameterValues([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrpvs != nil && len(*hrpvs) > 0 {
		return &((*hrpvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.rule.parameter.value not found", id)
}

// GetHrRuleParameterValues gets hr.rule.parameter.value existing records.
func (c *Client) GetHrRuleParameterValues(ids []int64) (*HrRuleParameterValues, error) {
	hrpvs := &HrRuleParameterValues{}
	if err := c.Read(HrRuleParameterValueModel, ids, nil, hrpvs); err != nil {
		return nil, err
	}
	return hrpvs, nil
}

// FindHrRuleParameterValue finds hr.rule.parameter.value record by querying it with criteria.
func (c *Client) FindHrRuleParameterValue(criteria *Criteria) (*HrRuleParameterValue, error) {
	hrpvs := &HrRuleParameterValues{}
	if err := c.SearchRead(HrRuleParameterValueModel, criteria, NewOptions().Limit(1), hrpvs); err != nil {
		return nil, err
	}
	if hrpvs != nil && len(*hrpvs) > 0 {
		return &((*hrpvs)[0]), nil
	}
	return nil, fmt.Errorf("hr.rule.parameter.value was not found with criteria %v", criteria)
}

// FindHrRuleParameterValues finds hr.rule.parameter.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRuleParameterValues(criteria *Criteria, options *Options) (*HrRuleParameterValues, error) {
	hrpvs := &HrRuleParameterValues{}
	if err := c.SearchRead(HrRuleParameterValueModel, criteria, options, hrpvs); err != nil {
		return nil, err
	}
	return hrpvs, nil
}

// FindHrRuleParameterValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrRuleParameterValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrRuleParameterValueModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrRuleParameterValueId finds record id by querying it with criteria.
func (c *Client) FindHrRuleParameterValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrRuleParameterValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.rule.parameter.value was not found with criteria %v and options %v", criteria, options)
}
