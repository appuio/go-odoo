package odoo

import (
	"fmt"
)

// ExpenseSampleReceipt represents expense.sample.receipt model.
type ExpenseSampleReceipt struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ExpenseSampleReceipts represents array of expense.sample.receipt model.
type ExpenseSampleReceipts []ExpenseSampleReceipt

// ExpenseSampleReceiptModel is the odoo model name.
const ExpenseSampleReceiptModel = "expense.sample.receipt"

// Many2One convert ExpenseSampleReceipt to *Many2One.
func (esr *ExpenseSampleReceipt) Many2One() *Many2One {
	return NewMany2One(esr.Id.Get(), "")
}

// CreateExpenseSampleReceipt creates a new expense.sample.receipt model and returns its id.
func (c *Client) CreateExpenseSampleReceipt(esr *ExpenseSampleReceipt) (int64, error) {
	ids, err := c.CreateExpenseSampleReceipts([]*ExpenseSampleReceipt{esr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateExpenseSampleReceipt creates a new expense.sample.receipt model and returns its id.
func (c *Client) CreateExpenseSampleReceipts(esrs []*ExpenseSampleReceipt) ([]int64, error) {
	var vv []interface{}
	for _, v := range esrs {
		vv = append(vv, v)
	}
	return c.Create(ExpenseSampleReceiptModel, vv)
}

// UpdateExpenseSampleReceipt updates an existing expense.sample.receipt record.
func (c *Client) UpdateExpenseSampleReceipt(esr *ExpenseSampleReceipt) error {
	return c.UpdateExpenseSampleReceipts([]int64{esr.Id.Get()}, esr)
}

// UpdateExpenseSampleReceipts updates existing expense.sample.receipt records.
// All records (represented by ids) will be updated by esr values.
func (c *Client) UpdateExpenseSampleReceipts(ids []int64, esr *ExpenseSampleReceipt) error {
	return c.Update(ExpenseSampleReceiptModel, ids, esr)
}

// DeleteExpenseSampleReceipt deletes an existing expense.sample.receipt record.
func (c *Client) DeleteExpenseSampleReceipt(id int64) error {
	return c.DeleteExpenseSampleReceipts([]int64{id})
}

// DeleteExpenseSampleReceipts deletes existing expense.sample.receipt records.
func (c *Client) DeleteExpenseSampleReceipts(ids []int64) error {
	return c.Delete(ExpenseSampleReceiptModel, ids)
}

// GetExpenseSampleReceipt gets expense.sample.receipt existing record.
func (c *Client) GetExpenseSampleReceipt(id int64) (*ExpenseSampleReceipt, error) {
	esrs, err := c.GetExpenseSampleReceipts([]int64{id})
	if err != nil {
		return nil, err
	}
	if esrs != nil && len(*esrs) > 0 {
		return &((*esrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of expense.sample.receipt not found", id)
}

// GetExpenseSampleReceipts gets expense.sample.receipt existing records.
func (c *Client) GetExpenseSampleReceipts(ids []int64) (*ExpenseSampleReceipts, error) {
	esrs := &ExpenseSampleReceipts{}
	if err := c.Read(ExpenseSampleReceiptModel, ids, nil, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindExpenseSampleReceipt finds expense.sample.receipt record by querying it with criteria.
func (c *Client) FindExpenseSampleReceipt(criteria *Criteria) (*ExpenseSampleReceipt, error) {
	esrs := &ExpenseSampleReceipts{}
	if err := c.SearchRead(ExpenseSampleReceiptModel, criteria, NewOptions().Limit(1), esrs); err != nil {
		return nil, err
	}
	if esrs != nil && len(*esrs) > 0 {
		return &((*esrs)[0]), nil
	}
	return nil, fmt.Errorf("expense.sample.receipt was not found with criteria %v", criteria)
}

// FindExpenseSampleReceipts finds expense.sample.receipt records by querying it
// and filtering it with criteria and options.
func (c *Client) FindExpenseSampleReceipts(criteria *Criteria, options *Options) (*ExpenseSampleReceipts, error) {
	esrs := &ExpenseSampleReceipts{}
	if err := c.SearchRead(ExpenseSampleReceiptModel, criteria, options, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindExpenseSampleReceiptIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindExpenseSampleReceiptIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ExpenseSampleReceiptModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindExpenseSampleReceiptId finds record id by querying it with criteria.
func (c *Client) FindExpenseSampleReceiptId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ExpenseSampleReceiptModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("expense.sample.receipt was not found with criteria %v and options %v", criteria, options)
}
