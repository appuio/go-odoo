package odoo

import (
	"fmt"
)

// ExpenseSampleRegister represents expense.sample.register model.
type ExpenseSampleRegister struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	Amount                        *Float     `xmlrpc:"amount,omptempty"`
	AvailablePaymentMethodLineIds *Relation  `xmlrpc:"available_payment_method_line_ids,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                          *Time      `xmlrpc:"date,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	HidePartial                   *Bool      `xmlrpc:"hide_partial,omptempty"`
	HidePaymentMethodLine         *Bool      `xmlrpc:"hide_payment_method_line,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	JournalId                     *Many2One  `xmlrpc:"journal_id,omptempty"`
	Memo                          *String    `xmlrpc:"memo,omptempty"`
	PartialMode                   *Selection `xmlrpc:"partial_mode,omptempty"`
	PaymentMethodLineId           *Many2One  `xmlrpc:"payment_method_line_id,omptempty"`
	SheetId                       *Many2One  `xmlrpc:"sheet_id,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ExpenseSampleRegisters represents array of expense.sample.register model.
type ExpenseSampleRegisters []ExpenseSampleRegister

// ExpenseSampleRegisterModel is the odoo model name.
const ExpenseSampleRegisterModel = "expense.sample.register"

// Many2One convert ExpenseSampleRegister to *Many2One.
func (esr *ExpenseSampleRegister) Many2One() *Many2One {
	return NewMany2One(esr.Id.Get(), "")
}

// CreateExpenseSampleRegister creates a new expense.sample.register model and returns its id.
func (c *Client) CreateExpenseSampleRegister(esr *ExpenseSampleRegister) (int64, error) {
	ids, err := c.CreateExpenseSampleRegisters([]*ExpenseSampleRegister{esr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateExpenseSampleRegister creates a new expense.sample.register model and returns its id.
func (c *Client) CreateExpenseSampleRegisters(esrs []*ExpenseSampleRegister) ([]int64, error) {
	var vv []interface{}
	for _, v := range esrs {
		vv = append(vv, v)
	}
	return c.Create(ExpenseSampleRegisterModel, vv)
}

// UpdateExpenseSampleRegister updates an existing expense.sample.register record.
func (c *Client) UpdateExpenseSampleRegister(esr *ExpenseSampleRegister) error {
	return c.UpdateExpenseSampleRegisters([]int64{esr.Id.Get()}, esr)
}

// UpdateExpenseSampleRegisters updates existing expense.sample.register records.
// All records (represented by ids) will be updated by esr values.
func (c *Client) UpdateExpenseSampleRegisters(ids []int64, esr *ExpenseSampleRegister) error {
	return c.Update(ExpenseSampleRegisterModel, ids, esr)
}

// DeleteExpenseSampleRegister deletes an existing expense.sample.register record.
func (c *Client) DeleteExpenseSampleRegister(id int64) error {
	return c.DeleteExpenseSampleRegisters([]int64{id})
}

// DeleteExpenseSampleRegisters deletes existing expense.sample.register records.
func (c *Client) DeleteExpenseSampleRegisters(ids []int64) error {
	return c.Delete(ExpenseSampleRegisterModel, ids)
}

// GetExpenseSampleRegister gets expense.sample.register existing record.
func (c *Client) GetExpenseSampleRegister(id int64) (*ExpenseSampleRegister, error) {
	esrs, err := c.GetExpenseSampleRegisters([]int64{id})
	if err != nil {
		return nil, err
	}
	if esrs != nil && len(*esrs) > 0 {
		return &((*esrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of expense.sample.register not found", id)
}

// GetExpenseSampleRegisters gets expense.sample.register existing records.
func (c *Client) GetExpenseSampleRegisters(ids []int64) (*ExpenseSampleRegisters, error) {
	esrs := &ExpenseSampleRegisters{}
	if err := c.Read(ExpenseSampleRegisterModel, ids, nil, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindExpenseSampleRegister finds expense.sample.register record by querying it with criteria.
func (c *Client) FindExpenseSampleRegister(criteria *Criteria) (*ExpenseSampleRegister, error) {
	esrs := &ExpenseSampleRegisters{}
	if err := c.SearchRead(ExpenseSampleRegisterModel, criteria, NewOptions().Limit(1), esrs); err != nil {
		return nil, err
	}
	if esrs != nil && len(*esrs) > 0 {
		return &((*esrs)[0]), nil
	}
	return nil, fmt.Errorf("expense.sample.register was not found with criteria %v", criteria)
}

// FindExpenseSampleRegisters finds expense.sample.register records by querying it
// and filtering it with criteria and options.
func (c *Client) FindExpenseSampleRegisters(criteria *Criteria, options *Options) (*ExpenseSampleRegisters, error) {
	esrs := &ExpenseSampleRegisters{}
	if err := c.SearchRead(ExpenseSampleRegisterModel, criteria, options, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindExpenseSampleRegisterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindExpenseSampleRegisterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ExpenseSampleRegisterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindExpenseSampleRegisterId finds record id by querying it with criteria.
func (c *Client) FindExpenseSampleRegisterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ExpenseSampleRegisterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("expense.sample.register was not found with criteria %v and options %v", criteria, options)
}
