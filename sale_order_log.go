package odoo

import (
	"fmt"
)

// SaleOrderLog represents sale.order.log model.
type SaleOrderLog struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	AmountCompanyCurrency *Float     `xmlrpc:"amount_company_currency,omptempty"`
	AmountSigned          *Float     `xmlrpc:"amount_signed,omptempty"`
	Category              *Selection `xmlrpc:"category,omptempty"`
	CompanyCurrencyId     *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId            *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	EventDate             *Time      `xmlrpc:"event_date,omptempty"`
	EventType             *Selection `xmlrpc:"event_type,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	OrderId               *Many2One  `xmlrpc:"order_id,omptempty"`
	RecurringMonthly      *Float     `xmlrpc:"recurring_monthly,omptempty"`
	TeamId                *Many2One  `xmlrpc:"team_id,omptempty"`
	UserId                *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderLogs represents array of sale.order.log model.
type SaleOrderLogs []SaleOrderLog

// SaleOrderLogModel is the odoo model name.
const SaleOrderLogModel = "sale.order.log"

// Many2One convert SaleOrderLog to *Many2One.
func (sol *SaleOrderLog) Many2One() *Many2One {
	return NewMany2One(sol.Id.Get(), "")
}

// CreateSaleOrderLog creates a new sale.order.log model and returns its id.
func (c *Client) CreateSaleOrderLog(sol *SaleOrderLog) (int64, error) {
	ids, err := c.CreateSaleOrderLogs([]*SaleOrderLog{sol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderLog creates a new sale.order.log model and returns its id.
func (c *Client) CreateSaleOrderLogs(sols []*SaleOrderLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range sols {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderLogModel, vv)
}

// UpdateSaleOrderLog updates an existing sale.order.log record.
func (c *Client) UpdateSaleOrderLog(sol *SaleOrderLog) error {
	return c.UpdateSaleOrderLogs([]int64{sol.Id.Get()}, sol)
}

// UpdateSaleOrderLogs updates existing sale.order.log records.
// All records (represented by ids) will be updated by sol values.
func (c *Client) UpdateSaleOrderLogs(ids []int64, sol *SaleOrderLog) error {
	return c.Update(SaleOrderLogModel, ids, sol)
}

// DeleteSaleOrderLog deletes an existing sale.order.log record.
func (c *Client) DeleteSaleOrderLog(id int64) error {
	return c.DeleteSaleOrderLogs([]int64{id})
}

// DeleteSaleOrderLogs deletes existing sale.order.log records.
func (c *Client) DeleteSaleOrderLogs(ids []int64) error {
	return c.Delete(SaleOrderLogModel, ids)
}

// GetSaleOrderLog gets sale.order.log existing record.
func (c *Client) GetSaleOrderLog(id int64) (*SaleOrderLog, error) {
	sols, err := c.GetSaleOrderLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	if sols != nil && len(*sols) > 0 {
		return &((*sols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.log not found", id)
}

// GetSaleOrderLogs gets sale.order.log existing records.
func (c *Client) GetSaleOrderLogs(ids []int64) (*SaleOrderLogs, error) {
	sols := &SaleOrderLogs{}
	if err := c.Read(SaleOrderLogModel, ids, nil, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLog finds sale.order.log record by querying it with criteria.
func (c *Client) FindSaleOrderLog(criteria *Criteria) (*SaleOrderLog, error) {
	sols := &SaleOrderLogs{}
	if err := c.SearchRead(SaleOrderLogModel, criteria, NewOptions().Limit(1), sols); err != nil {
		return nil, err
	}
	if sols != nil && len(*sols) > 0 {
		return &((*sols)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.log was not found with criteria %v", criteria)
}

// FindSaleOrderLogs finds sale.order.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLogs(criteria *Criteria, options *Options) (*SaleOrderLogs, error) {
	sols := &SaleOrderLogs{}
	if err := c.SearchRead(SaleOrderLogModel, criteria, options, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderLogModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderLogId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.log was not found with criteria %v and options %v", criteria, options)
}
