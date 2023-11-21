package odoo

import (
	"fmt"
)

// HrExpenseSplit represents hr.expense.split model.
type HrExpenseSplit struct {
	LastUpdate                 *Time       `xmlrpc:"__last_update,omptempty"`
	AmountTax                  *Float      `xmlrpc:"amount_tax,omptempty"`
	AnalyticDistribution       interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticDistributionSearch interface{} `xmlrpc:"analytic_distribution_search,omptempty"`
	AnalyticPrecision          *Int        `xmlrpc:"analytic_precision,omptempty"`
	CanBeReinvoiced            *Bool       `xmlrpc:"can_be_reinvoiced,omptempty"`
	CompanyId                  *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate                 *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One   `xmlrpc:"create_uid,omptempty"`
	CurrencyId                 *Many2One   `xmlrpc:"currency_id,omptempty"`
	DisplayName                *String     `xmlrpc:"display_name,omptempty"`
	EmployeeId                 *Many2One   `xmlrpc:"employee_id,omptempty"`
	ExpenseId                  *Many2One   `xmlrpc:"expense_id,omptempty"`
	Id                         *Int        `xmlrpc:"id,omptempty"`
	Name                       *String     `xmlrpc:"name,omptempty"`
	ProductHasCost             *Bool       `xmlrpc:"product_has_cost,omptempty"`
	ProductHasTax              *Bool       `xmlrpc:"product_has_tax,omptempty"`
	ProductId                  *Many2One   `xmlrpc:"product_id,omptempty"`
	SaleOrderId                *Many2One   `xmlrpc:"sale_order_id,omptempty"`
	TaxIds                     *Relation   `xmlrpc:"tax_ids,omptempty"`
	TotalAmount                *Float      `xmlrpc:"total_amount,omptempty"`
	WizardId                   *Many2One   `xmlrpc:"wizard_id,omptempty"`
	WriteDate                  *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// HrExpenseSplits represents array of hr.expense.split model.
type HrExpenseSplits []HrExpenseSplit

// HrExpenseSplitModel is the odoo model name.
const HrExpenseSplitModel = "hr.expense.split"

// Many2One convert HrExpenseSplit to *Many2One.
func (hes *HrExpenseSplit) Many2One() *Many2One {
	return NewMany2One(hes.Id.Get(), "")
}

// CreateHrExpenseSplit creates a new hr.expense.split model and returns its id.
func (c *Client) CreateHrExpenseSplit(hes *HrExpenseSplit) (int64, error) {
	ids, err := c.CreateHrExpenseSplits([]*HrExpenseSplit{hes})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseSplit creates a new hr.expense.split model and returns its id.
func (c *Client) CreateHrExpenseSplits(hess []*HrExpenseSplit) ([]int64, error) {
	var vv []interface{}
	for _, v := range hess {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseSplitModel, vv)
}

// UpdateHrExpenseSplit updates an existing hr.expense.split record.
func (c *Client) UpdateHrExpenseSplit(hes *HrExpenseSplit) error {
	return c.UpdateHrExpenseSplits([]int64{hes.Id.Get()}, hes)
}

// UpdateHrExpenseSplits updates existing hr.expense.split records.
// All records (represented by ids) will be updated by hes values.
func (c *Client) UpdateHrExpenseSplits(ids []int64, hes *HrExpenseSplit) error {
	return c.Update(HrExpenseSplitModel, ids, hes)
}

// DeleteHrExpenseSplit deletes an existing hr.expense.split record.
func (c *Client) DeleteHrExpenseSplit(id int64) error {
	return c.DeleteHrExpenseSplits([]int64{id})
}

// DeleteHrExpenseSplits deletes existing hr.expense.split records.
func (c *Client) DeleteHrExpenseSplits(ids []int64) error {
	return c.Delete(HrExpenseSplitModel, ids)
}

// GetHrExpenseSplit gets hr.expense.split existing record.
func (c *Client) GetHrExpenseSplit(id int64) (*HrExpenseSplit, error) {
	hess, err := c.GetHrExpenseSplits([]int64{id})
	if err != nil {
		return nil, err
	}
	if hess != nil && len(*hess) > 0 {
		return &((*hess)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.expense.split not found", id)
}

// GetHrExpenseSplits gets hr.expense.split existing records.
func (c *Client) GetHrExpenseSplits(ids []int64) (*HrExpenseSplits, error) {
	hess := &HrExpenseSplits{}
	if err := c.Read(HrExpenseSplitModel, ids, nil, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSplit finds hr.expense.split record by querying it with criteria.
func (c *Client) FindHrExpenseSplit(criteria *Criteria) (*HrExpenseSplit, error) {
	hess := &HrExpenseSplits{}
	if err := c.SearchRead(HrExpenseSplitModel, criteria, NewOptions().Limit(1), hess); err != nil {
		return nil, err
	}
	if hess != nil && len(*hess) > 0 {
		return &((*hess)[0]), nil
	}
	return nil, fmt.Errorf("hr.expense.split was not found with criteria %v", criteria)
}

// FindHrExpenseSplits finds hr.expense.split records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplits(criteria *Criteria, options *Options) (*HrExpenseSplits, error) {
	hess := &HrExpenseSplits{}
	if err := c.SearchRead(HrExpenseSplitModel, criteria, options, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrExpenseSplitIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplitIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrExpenseSplitModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrExpenseSplitId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSplitId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSplitModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.expense.split was not found with criteria %v and options %v", criteria, options)
}
