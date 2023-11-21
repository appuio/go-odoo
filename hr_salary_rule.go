package odoo

import (
	"fmt"
)

// HrSalaryRule represents hr.salary.rule model.
type HrSalaryRule struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omptempty"`
	AccountCredit                  *Many2One  `xmlrpc:"account_credit,omptempty"`
	AccountDebit                   *Many2One  `xmlrpc:"account_debit,omptempty"`
	Active                         *Bool      `xmlrpc:"active,omptempty"`
	AmountFix                      *Float     `xmlrpc:"amount_fix,omptempty"`
	AmountPercentage               *Float     `xmlrpc:"amount_percentage,omptempty"`
	AmountPercentageBase           *String    `xmlrpc:"amount_percentage_base,omptempty"`
	AmountPythonCompute            *String    `xmlrpc:"amount_python_compute,omptempty"`
	AmountSelect                   *Selection `xmlrpc:"amount_select,omptempty"`
	AnalyticAccountId              *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	AppearsOnEmployeeCostDashboard *Bool      `xmlrpc:"appears_on_employee_cost_dashboard,omptempty"`
	AppearsOnPayrollReport         *Bool      `xmlrpc:"appears_on_payroll_report,omptempty"`
	AppearsOnPayslip               *Bool      `xmlrpc:"appears_on_payslip,omptempty"`
	CategoryId                     *Many2One  `xmlrpc:"category_id,omptempty"`
	Code                           *String    `xmlrpc:"code,omptempty"`
	ConditionPython                *String    `xmlrpc:"condition_python,omptempty"`
	ConditionRange                 *String    `xmlrpc:"condition_range,omptempty"`
	ConditionRangeMax              *Float     `xmlrpc:"condition_range_max,omptempty"`
	ConditionRangeMin              *Float     `xmlrpc:"condition_range_min,omptempty"`
	ConditionSelect                *Selection `xmlrpc:"condition_select,omptempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omptempty"`
	Id                             *Int       `xmlrpc:"id,omptempty"`
	Name                           *String    `xmlrpc:"name,omptempty"`
	NotComputedInNet               *Bool      `xmlrpc:"not_computed_in_net,omptempty"`
	Note                           *String    `xmlrpc:"note,omptempty"`
	PartnerId                      *Many2One  `xmlrpc:"partner_id,omptempty"`
	Quantity                       *String    `xmlrpc:"quantity,omptempty"`
	Sequence                       *Int       `xmlrpc:"sequence,omptempty"`
	StructId                       *Many2One  `xmlrpc:"struct_id,omptempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrSalaryRules represents array of hr.salary.rule model.
type HrSalaryRules []HrSalaryRule

// HrSalaryRuleModel is the odoo model name.
const HrSalaryRuleModel = "hr.salary.rule"

// Many2One convert HrSalaryRule to *Many2One.
func (hsr *HrSalaryRule) Many2One() *Many2One {
	return NewMany2One(hsr.Id.Get(), "")
}

// CreateHrSalaryRule creates a new hr.salary.rule model and returns its id.
func (c *Client) CreateHrSalaryRule(hsr *HrSalaryRule) (int64, error) {
	ids, err := c.CreateHrSalaryRules([]*HrSalaryRule{hsr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSalaryRule creates a new hr.salary.rule model and returns its id.
func (c *Client) CreateHrSalaryRules(hsrs []*HrSalaryRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsrs {
		vv = append(vv, v)
	}
	return c.Create(HrSalaryRuleModel, vv)
}

// UpdateHrSalaryRule updates an existing hr.salary.rule record.
func (c *Client) UpdateHrSalaryRule(hsr *HrSalaryRule) error {
	return c.UpdateHrSalaryRules([]int64{hsr.Id.Get()}, hsr)
}

// UpdateHrSalaryRules updates existing hr.salary.rule records.
// All records (represented by ids) will be updated by hsr values.
func (c *Client) UpdateHrSalaryRules(ids []int64, hsr *HrSalaryRule) error {
	return c.Update(HrSalaryRuleModel, ids, hsr)
}

// DeleteHrSalaryRule deletes an existing hr.salary.rule record.
func (c *Client) DeleteHrSalaryRule(id int64) error {
	return c.DeleteHrSalaryRules([]int64{id})
}

// DeleteHrSalaryRules deletes existing hr.salary.rule records.
func (c *Client) DeleteHrSalaryRules(ids []int64) error {
	return c.Delete(HrSalaryRuleModel, ids)
}

// GetHrSalaryRule gets hr.salary.rule existing record.
func (c *Client) GetHrSalaryRule(id int64) (*HrSalaryRule, error) {
	hsrs, err := c.GetHrSalaryRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsrs != nil && len(*hsrs) > 0 {
		return &((*hsrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.salary.rule not found", id)
}

// GetHrSalaryRules gets hr.salary.rule existing records.
func (c *Client) GetHrSalaryRules(ids []int64) (*HrSalaryRules, error) {
	hsrs := &HrSalaryRules{}
	if err := c.Read(HrSalaryRuleModel, ids, nil, hsrs); err != nil {
		return nil, err
	}
	return hsrs, nil
}

// FindHrSalaryRule finds hr.salary.rule record by querying it with criteria.
func (c *Client) FindHrSalaryRule(criteria *Criteria) (*HrSalaryRule, error) {
	hsrs := &HrSalaryRules{}
	if err := c.SearchRead(HrSalaryRuleModel, criteria, NewOptions().Limit(1), hsrs); err != nil {
		return nil, err
	}
	if hsrs != nil && len(*hsrs) > 0 {
		return &((*hsrs)[0]), nil
	}
	return nil, fmt.Errorf("hr.salary.rule was not found with criteria %v", criteria)
}

// FindHrSalaryRules finds hr.salary.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryRules(criteria *Criteria, options *Options) (*HrSalaryRules, error) {
	hsrs := &HrSalaryRules{}
	if err := c.SearchRead(HrSalaryRuleModel, criteria, options, hsrs); err != nil {
		return nil, err
	}
	return hsrs, nil
}

// FindHrSalaryRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrSalaryRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrSalaryRuleId finds record id by querying it with criteria.
func (c *Client) FindHrSalaryRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSalaryRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.salary.rule was not found with criteria %v and options %v", criteria, options)
}
