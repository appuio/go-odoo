package odoo

import (
	"fmt"
)

// CrossoveredBudgetLines represents crossovered.budget.lines model.
type CrossoveredBudgetLines struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	AnalyticAccountId      *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	AnalyticPlanId         *Many2One  `xmlrpc:"analytic_plan_id,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrossoveredBudgetId    *Many2One  `xmlrpc:"crossovered_budget_id,omptempty"`
	CrossoveredBudgetState *Selection `xmlrpc:"crossovered_budget_state,omptempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateFrom               *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                 *Time      `xmlrpc:"date_to,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	GeneralBudgetId        *Many2One  `xmlrpc:"general_budget_id,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	IsAboveBudget          *Bool      `xmlrpc:"is_above_budget,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	PaidDate               *Time      `xmlrpc:"paid_date,omptempty"`
	Percentage             *Float     `xmlrpc:"percentage,omptempty"`
	PlannedAmount          *Float     `xmlrpc:"planned_amount,omptempty"`
	PracticalAmount        *Float     `xmlrpc:"practical_amount,omptempty"`
	TheoriticalAmount      *Float     `xmlrpc:"theoritical_amount,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CrossoveredBudgetLiness represents array of crossovered.budget.lines model.
type CrossoveredBudgetLiness []CrossoveredBudgetLines

// CrossoveredBudgetLinesModel is the odoo model name.
const CrossoveredBudgetLinesModel = "crossovered.budget.lines"

// Many2One convert CrossoveredBudgetLines to *Many2One.
func (cbl *CrossoveredBudgetLines) Many2One() *Many2One {
	return NewMany2One(cbl.Id.Get(), "")
}

// CreateCrossoveredBudgetLines creates a new crossovered.budget.lines model and returns its id.
func (c *Client) CreateCrossoveredBudgetLines(cbl *CrossoveredBudgetLines) (int64, error) {
	ids, err := c.CreateCrossoveredBudgetLiness([]*CrossoveredBudgetLines{cbl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrossoveredBudgetLines creates a new crossovered.budget.lines model and returns its id.
func (c *Client) CreateCrossoveredBudgetLiness(cbls []*CrossoveredBudgetLines) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbls {
		vv = append(vv, v)
	}
	return c.Create(CrossoveredBudgetLinesModel, vv)
}

// UpdateCrossoveredBudgetLines updates an existing crossovered.budget.lines record.
func (c *Client) UpdateCrossoveredBudgetLines(cbl *CrossoveredBudgetLines) error {
	return c.UpdateCrossoveredBudgetLiness([]int64{cbl.Id.Get()}, cbl)
}

// UpdateCrossoveredBudgetLiness updates existing crossovered.budget.lines records.
// All records (represented by ids) will be updated by cbl values.
func (c *Client) UpdateCrossoveredBudgetLiness(ids []int64, cbl *CrossoveredBudgetLines) error {
	return c.Update(CrossoveredBudgetLinesModel, ids, cbl)
}

// DeleteCrossoveredBudgetLines deletes an existing crossovered.budget.lines record.
func (c *Client) DeleteCrossoveredBudgetLines(id int64) error {
	return c.DeleteCrossoveredBudgetLiness([]int64{id})
}

// DeleteCrossoveredBudgetLiness deletes existing crossovered.budget.lines records.
func (c *Client) DeleteCrossoveredBudgetLiness(ids []int64) error {
	return c.Delete(CrossoveredBudgetLinesModel, ids)
}

// GetCrossoveredBudgetLines gets crossovered.budget.lines existing record.
func (c *Client) GetCrossoveredBudgetLines(id int64) (*CrossoveredBudgetLines, error) {
	cbls, err := c.GetCrossoveredBudgetLiness([]int64{id})
	if err != nil {
		return nil, err
	}
	if cbls != nil && len(*cbls) > 0 {
		return &((*cbls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crossovered.budget.lines not found", id)
}

// GetCrossoveredBudgetLiness gets crossovered.budget.lines existing records.
func (c *Client) GetCrossoveredBudgetLiness(ids []int64) (*CrossoveredBudgetLiness, error) {
	cbls := &CrossoveredBudgetLiness{}
	if err := c.Read(CrossoveredBudgetLinesModel, ids, nil, cbls); err != nil {
		return nil, err
	}
	return cbls, nil
}

// FindCrossoveredBudgetLines finds crossovered.budget.lines record by querying it with criteria.
func (c *Client) FindCrossoveredBudgetLines(criteria *Criteria) (*CrossoveredBudgetLines, error) {
	cbls := &CrossoveredBudgetLiness{}
	if err := c.SearchRead(CrossoveredBudgetLinesModel, criteria, NewOptions().Limit(1), cbls); err != nil {
		return nil, err
	}
	if cbls != nil && len(*cbls) > 0 {
		return &((*cbls)[0]), nil
	}
	return nil, fmt.Errorf("crossovered.budget.lines was not found with criteria %v", criteria)
}

// FindCrossoveredBudgetLiness finds crossovered.budget.lines records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgetLiness(criteria *Criteria, options *Options) (*CrossoveredBudgetLiness, error) {
	cbls := &CrossoveredBudgetLiness{}
	if err := c.SearchRead(CrossoveredBudgetLinesModel, criteria, options, cbls); err != nil {
		return nil, err
	}
	return cbls, nil
}

// FindCrossoveredBudgetLinesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgetLinesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrossoveredBudgetLinesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrossoveredBudgetLinesId finds record id by querying it with criteria.
func (c *Client) FindCrossoveredBudgetLinesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrossoveredBudgetLinesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crossovered.budget.lines was not found with criteria %v and options %v", criteria, options)
}
