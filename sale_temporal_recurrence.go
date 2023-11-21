package odoo

import (
	"fmt"
)

// SaleTemporalRecurrence represents sale.temporal.recurrence model.
type SaleTemporalRecurrence struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	Active                  *Bool      `xmlrpc:"active,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	Duration                *Int       `xmlrpc:"duration,omptempty"`
	DurationDisplay         *String    `xmlrpc:"duration_display,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	SubscriptionUnitDisplay *String    `xmlrpc:"subscription_unit_display,omptempty"`
	Unit                    *Selection `xmlrpc:"unit,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SaleTemporalRecurrences represents array of sale.temporal.recurrence model.
type SaleTemporalRecurrences []SaleTemporalRecurrence

// SaleTemporalRecurrenceModel is the odoo model name.
const SaleTemporalRecurrenceModel = "sale.temporal.recurrence"

// Many2One convert SaleTemporalRecurrence to *Many2One.
func (str *SaleTemporalRecurrence) Many2One() *Many2One {
	return NewMany2One(str.Id.Get(), "")
}

// CreateSaleTemporalRecurrence creates a new sale.temporal.recurrence model and returns its id.
func (c *Client) CreateSaleTemporalRecurrence(str *SaleTemporalRecurrence) (int64, error) {
	ids, err := c.CreateSaleTemporalRecurrences([]*SaleTemporalRecurrence{str})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleTemporalRecurrence creates a new sale.temporal.recurrence model and returns its id.
func (c *Client) CreateSaleTemporalRecurrences(strs []*SaleTemporalRecurrence) ([]int64, error) {
	var vv []interface{}
	for _, v := range strs {
		vv = append(vv, v)
	}
	return c.Create(SaleTemporalRecurrenceModel, vv)
}

// UpdateSaleTemporalRecurrence updates an existing sale.temporal.recurrence record.
func (c *Client) UpdateSaleTemporalRecurrence(str *SaleTemporalRecurrence) error {
	return c.UpdateSaleTemporalRecurrences([]int64{str.Id.Get()}, str)
}

// UpdateSaleTemporalRecurrences updates existing sale.temporal.recurrence records.
// All records (represented by ids) will be updated by str values.
func (c *Client) UpdateSaleTemporalRecurrences(ids []int64, str *SaleTemporalRecurrence) error {
	return c.Update(SaleTemporalRecurrenceModel, ids, str)
}

// DeleteSaleTemporalRecurrence deletes an existing sale.temporal.recurrence record.
func (c *Client) DeleteSaleTemporalRecurrence(id int64) error {
	return c.DeleteSaleTemporalRecurrences([]int64{id})
}

// DeleteSaleTemporalRecurrences deletes existing sale.temporal.recurrence records.
func (c *Client) DeleteSaleTemporalRecurrences(ids []int64) error {
	return c.Delete(SaleTemporalRecurrenceModel, ids)
}

// GetSaleTemporalRecurrence gets sale.temporal.recurrence existing record.
func (c *Client) GetSaleTemporalRecurrence(id int64) (*SaleTemporalRecurrence, error) {
	strs, err := c.GetSaleTemporalRecurrences([]int64{id})
	if err != nil {
		return nil, err
	}
	if strs != nil && len(*strs) > 0 {
		return &((*strs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.temporal.recurrence not found", id)
}

// GetSaleTemporalRecurrences gets sale.temporal.recurrence existing records.
func (c *Client) GetSaleTemporalRecurrences(ids []int64) (*SaleTemporalRecurrences, error) {
	strs := &SaleTemporalRecurrences{}
	if err := c.Read(SaleTemporalRecurrenceModel, ids, nil, strs); err != nil {
		return nil, err
	}
	return strs, nil
}

// FindSaleTemporalRecurrence finds sale.temporal.recurrence record by querying it with criteria.
func (c *Client) FindSaleTemporalRecurrence(criteria *Criteria) (*SaleTemporalRecurrence, error) {
	strs := &SaleTemporalRecurrences{}
	if err := c.SearchRead(SaleTemporalRecurrenceModel, criteria, NewOptions().Limit(1), strs); err != nil {
		return nil, err
	}
	if strs != nil && len(*strs) > 0 {
		return &((*strs)[0]), nil
	}
	return nil, fmt.Errorf("sale.temporal.recurrence was not found with criteria %v", criteria)
}

// FindSaleTemporalRecurrences finds sale.temporal.recurrence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleTemporalRecurrences(criteria *Criteria, options *Options) (*SaleTemporalRecurrences, error) {
	strs := &SaleTemporalRecurrences{}
	if err := c.SearchRead(SaleTemporalRecurrenceModel, criteria, options, strs); err != nil {
		return nil, err
	}
	return strs, nil
}

// FindSaleTemporalRecurrenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleTemporalRecurrenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleTemporalRecurrenceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleTemporalRecurrenceId finds record id by querying it with criteria.
func (c *Client) FindSaleTemporalRecurrenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleTemporalRecurrenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.temporal.recurrence was not found with criteria %v and options %v", criteria, options)
}
