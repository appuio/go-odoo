package odoo

import (
	"fmt"
)

// HrExpenseSplitWizard represents hr.expense.split.wizard model.
type HrExpenseSplitWizard struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId          *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	ExpenseId           *Many2One `xmlrpc:"expense_id,omptempty"`
	ExpenseSplitLineIds *Relation `xmlrpc:"expense_split_line_ids,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	SplitPossible       *Bool     `xmlrpc:"split_possible,omptempty"`
	TotalAmount         *Float    `xmlrpc:"total_amount,omptempty"`
	TotalAmountOriginal *Float    `xmlrpc:"total_amount_original,omptempty"`
	TotalAmountTaxes    *Float    `xmlrpc:"total_amount_taxes,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrExpenseSplitWizards represents array of hr.expense.split.wizard model.
type HrExpenseSplitWizards []HrExpenseSplitWizard

// HrExpenseSplitWizardModel is the odoo model name.
const HrExpenseSplitWizardModel = "hr.expense.split.wizard"

// Many2One convert HrExpenseSplitWizard to *Many2One.
func (hesw *HrExpenseSplitWizard) Many2One() *Many2One {
	return NewMany2One(hesw.Id.Get(), "")
}

// CreateHrExpenseSplitWizard creates a new hr.expense.split.wizard model and returns its id.
func (c *Client) CreateHrExpenseSplitWizard(hesw *HrExpenseSplitWizard) (int64, error) {
	ids, err := c.CreateHrExpenseSplitWizards([]*HrExpenseSplitWizard{hesw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseSplitWizard creates a new hr.expense.split.wizard model and returns its id.
func (c *Client) CreateHrExpenseSplitWizards(hesws []*HrExpenseSplitWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hesws {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseSplitWizardModel, vv)
}

// UpdateHrExpenseSplitWizard updates an existing hr.expense.split.wizard record.
func (c *Client) UpdateHrExpenseSplitWizard(hesw *HrExpenseSplitWizard) error {
	return c.UpdateHrExpenseSplitWizards([]int64{hesw.Id.Get()}, hesw)
}

// UpdateHrExpenseSplitWizards updates existing hr.expense.split.wizard records.
// All records (represented by ids) will be updated by hesw values.
func (c *Client) UpdateHrExpenseSplitWizards(ids []int64, hesw *HrExpenseSplitWizard) error {
	return c.Update(HrExpenseSplitWizardModel, ids, hesw)
}

// DeleteHrExpenseSplitWizard deletes an existing hr.expense.split.wizard record.
func (c *Client) DeleteHrExpenseSplitWizard(id int64) error {
	return c.DeleteHrExpenseSplitWizards([]int64{id})
}

// DeleteHrExpenseSplitWizards deletes existing hr.expense.split.wizard records.
func (c *Client) DeleteHrExpenseSplitWizards(ids []int64) error {
	return c.Delete(HrExpenseSplitWizardModel, ids)
}

// GetHrExpenseSplitWizard gets hr.expense.split.wizard existing record.
func (c *Client) GetHrExpenseSplitWizard(id int64) (*HrExpenseSplitWizard, error) {
	hesws, err := c.GetHrExpenseSplitWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hesws != nil && len(*hesws) > 0 {
		return &((*hesws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.expense.split.wizard not found", id)
}

// GetHrExpenseSplitWizards gets hr.expense.split.wizard existing records.
func (c *Client) GetHrExpenseSplitWizards(ids []int64) (*HrExpenseSplitWizards, error) {
	hesws := &HrExpenseSplitWizards{}
	if err := c.Read(HrExpenseSplitWizardModel, ids, nil, hesws); err != nil {
		return nil, err
	}
	return hesws, nil
}

// FindHrExpenseSplitWizard finds hr.expense.split.wizard record by querying it with criteria.
func (c *Client) FindHrExpenseSplitWizard(criteria *Criteria) (*HrExpenseSplitWizard, error) {
	hesws := &HrExpenseSplitWizards{}
	if err := c.SearchRead(HrExpenseSplitWizardModel, criteria, NewOptions().Limit(1), hesws); err != nil {
		return nil, err
	}
	if hesws != nil && len(*hesws) > 0 {
		return &((*hesws)[0]), nil
	}
	return nil, fmt.Errorf("hr.expense.split.wizard was not found with criteria %v", criteria)
}

// FindHrExpenseSplitWizards finds hr.expense.split.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplitWizards(criteria *Criteria, options *Options) (*HrExpenseSplitWizards, error) {
	hesws := &HrExpenseSplitWizards{}
	if err := c.SearchRead(HrExpenseSplitWizardModel, criteria, options, hesws); err != nil {
		return nil, err
	}
	return hesws, nil
}

// FindHrExpenseSplitWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseSplitWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrExpenseSplitWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrExpenseSplitWizardId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseSplitWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseSplitWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.expense.split.wizard was not found with criteria %v and options %v", criteria, options)
}
