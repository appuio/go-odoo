package odoo

import (
	"fmt"
)

// HrExpenseExtractWords represents hr.expense.extract.words model.
type HrExpenseExtractWords struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	ExpenseId   *Many2One `xmlrpc:"expense_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WordPage    *Int      `xmlrpc:"word_page,omptempty"`
	WordText    *String   `xmlrpc:"word_text,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrExpenseExtractWordss represents array of hr.expense.extract.words model.
type HrExpenseExtractWordss []HrExpenseExtractWords

// HrExpenseExtractWordsModel is the odoo model name.
const HrExpenseExtractWordsModel = "hr.expense.extract.words"

// Many2One convert HrExpenseExtractWords to *Many2One.
func (heew *HrExpenseExtractWords) Many2One() *Many2One {
	return NewMany2One(heew.Id.Get(), "")
}

// CreateHrExpenseExtractWords creates a new hr.expense.extract.words model and returns its id.
func (c *Client) CreateHrExpenseExtractWords(heew *HrExpenseExtractWords) (int64, error) {
	ids, err := c.CreateHrExpenseExtractWordss([]*HrExpenseExtractWords{heew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpenseExtractWords creates a new hr.expense.extract.words model and returns its id.
func (c *Client) CreateHrExpenseExtractWordss(heews []*HrExpenseExtractWords) ([]int64, error) {
	var vv []interface{}
	for _, v := range heews {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseExtractWordsModel, vv)
}

// UpdateHrExpenseExtractWords updates an existing hr.expense.extract.words record.
func (c *Client) UpdateHrExpenseExtractWords(heew *HrExpenseExtractWords) error {
	return c.UpdateHrExpenseExtractWordss([]int64{heew.Id.Get()}, heew)
}

// UpdateHrExpenseExtractWordss updates existing hr.expense.extract.words records.
// All records (represented by ids) will be updated by heew values.
func (c *Client) UpdateHrExpenseExtractWordss(ids []int64, heew *HrExpenseExtractWords) error {
	return c.Update(HrExpenseExtractWordsModel, ids, heew)
}

// DeleteHrExpenseExtractWords deletes an existing hr.expense.extract.words record.
func (c *Client) DeleteHrExpenseExtractWords(id int64) error {
	return c.DeleteHrExpenseExtractWordss([]int64{id})
}

// DeleteHrExpenseExtractWordss deletes existing hr.expense.extract.words records.
func (c *Client) DeleteHrExpenseExtractWordss(ids []int64) error {
	return c.Delete(HrExpenseExtractWordsModel, ids)
}

// GetHrExpenseExtractWords gets hr.expense.extract.words existing record.
func (c *Client) GetHrExpenseExtractWords(id int64) (*HrExpenseExtractWords, error) {
	heews, err := c.GetHrExpenseExtractWordss([]int64{id})
	if err != nil {
		return nil, err
	}
	if heews != nil && len(*heews) > 0 {
		return &((*heews)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.expense.extract.words not found", id)
}

// GetHrExpenseExtractWordss gets hr.expense.extract.words existing records.
func (c *Client) GetHrExpenseExtractWordss(ids []int64) (*HrExpenseExtractWordss, error) {
	heews := &HrExpenseExtractWordss{}
	if err := c.Read(HrExpenseExtractWordsModel, ids, nil, heews); err != nil {
		return nil, err
	}
	return heews, nil
}

// FindHrExpenseExtractWords finds hr.expense.extract.words record by querying it with criteria.
func (c *Client) FindHrExpenseExtractWords(criteria *Criteria) (*HrExpenseExtractWords, error) {
	heews := &HrExpenseExtractWordss{}
	if err := c.SearchRead(HrExpenseExtractWordsModel, criteria, NewOptions().Limit(1), heews); err != nil {
		return nil, err
	}
	if heews != nil && len(*heews) > 0 {
		return &((*heews)[0]), nil
	}
	return nil, fmt.Errorf("hr.expense.extract.words was not found with criteria %v", criteria)
}

// FindHrExpenseExtractWordss finds hr.expense.extract.words records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseExtractWordss(criteria *Criteria, options *Options) (*HrExpenseExtractWordss, error) {
	heews := &HrExpenseExtractWordss{}
	if err := c.SearchRead(HrExpenseExtractWordsModel, criteria, options, heews); err != nil {
		return nil, err
	}
	return heews, nil
}

// FindHrExpenseExtractWordsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseExtractWordsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrExpenseExtractWordsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrExpenseExtractWordsId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseExtractWordsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseExtractWordsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.expense.extract.words was not found with criteria %v and options %v", criteria, options)
}
