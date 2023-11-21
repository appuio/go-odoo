package odoo

import (
	"fmt"
)

// SaleOrderCloseReason represents sale.order.close.reason model.
type SaleOrderCloseReason struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderCloseReasons represents array of sale.order.close.reason model.
type SaleOrderCloseReasons []SaleOrderCloseReason

// SaleOrderCloseReasonModel is the odoo model name.
const SaleOrderCloseReasonModel = "sale.order.close.reason"

// Many2One convert SaleOrderCloseReason to *Many2One.
func (socr *SaleOrderCloseReason) Many2One() *Many2One {
	return NewMany2One(socr.Id.Get(), "")
}

// CreateSaleOrderCloseReason creates a new sale.order.close.reason model and returns its id.
func (c *Client) CreateSaleOrderCloseReason(socr *SaleOrderCloseReason) (int64, error) {
	ids, err := c.CreateSaleOrderCloseReasons([]*SaleOrderCloseReason{socr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderCloseReason creates a new sale.order.close.reason model and returns its id.
func (c *Client) CreateSaleOrderCloseReasons(socrs []*SaleOrderCloseReason) ([]int64, error) {
	var vv []interface{}
	for _, v := range socrs {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderCloseReasonModel, vv)
}

// UpdateSaleOrderCloseReason updates an existing sale.order.close.reason record.
func (c *Client) UpdateSaleOrderCloseReason(socr *SaleOrderCloseReason) error {
	return c.UpdateSaleOrderCloseReasons([]int64{socr.Id.Get()}, socr)
}

// UpdateSaleOrderCloseReasons updates existing sale.order.close.reason records.
// All records (represented by ids) will be updated by socr values.
func (c *Client) UpdateSaleOrderCloseReasons(ids []int64, socr *SaleOrderCloseReason) error {
	return c.Update(SaleOrderCloseReasonModel, ids, socr)
}

// DeleteSaleOrderCloseReason deletes an existing sale.order.close.reason record.
func (c *Client) DeleteSaleOrderCloseReason(id int64) error {
	return c.DeleteSaleOrderCloseReasons([]int64{id})
}

// DeleteSaleOrderCloseReasons deletes existing sale.order.close.reason records.
func (c *Client) DeleteSaleOrderCloseReasons(ids []int64) error {
	return c.Delete(SaleOrderCloseReasonModel, ids)
}

// GetSaleOrderCloseReason gets sale.order.close.reason existing record.
func (c *Client) GetSaleOrderCloseReason(id int64) (*SaleOrderCloseReason, error) {
	socrs, err := c.GetSaleOrderCloseReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	if socrs != nil && len(*socrs) > 0 {
		return &((*socrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.close.reason not found", id)
}

// GetSaleOrderCloseReasons gets sale.order.close.reason existing records.
func (c *Client) GetSaleOrderCloseReasons(ids []int64) (*SaleOrderCloseReasons, error) {
	socrs := &SaleOrderCloseReasons{}
	if err := c.Read(SaleOrderCloseReasonModel, ids, nil, socrs); err != nil {
		return nil, err
	}
	return socrs, nil
}

// FindSaleOrderCloseReason finds sale.order.close.reason record by querying it with criteria.
func (c *Client) FindSaleOrderCloseReason(criteria *Criteria) (*SaleOrderCloseReason, error) {
	socrs := &SaleOrderCloseReasons{}
	if err := c.SearchRead(SaleOrderCloseReasonModel, criteria, NewOptions().Limit(1), socrs); err != nil {
		return nil, err
	}
	if socrs != nil && len(*socrs) > 0 {
		return &((*socrs)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.close.reason was not found with criteria %v", criteria)
}

// FindSaleOrderCloseReasons finds sale.order.close.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCloseReasons(criteria *Criteria, options *Options) (*SaleOrderCloseReasons, error) {
	socrs := &SaleOrderCloseReasons{}
	if err := c.SearchRead(SaleOrderCloseReasonModel, criteria, options, socrs); err != nil {
		return nil, err
	}
	return socrs, nil
}

// FindSaleOrderCloseReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCloseReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderCloseReasonModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderCloseReasonId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderCloseReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderCloseReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.close.reason was not found with criteria %v and options %v", criteria, options)
}
