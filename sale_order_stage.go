package odoo

import (
	"fmt"
)

// SaleOrderStage represents sale.order.stage model.
type SaleOrderStage struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	Category         *Selection `xmlrpc:"category,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description      *String    `xmlrpc:"description,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	Fold             *Bool      `xmlrpc:"fold,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	RatingTemplateId *Many2One  `xmlrpc:"rating_template_id,omptempty"`
	Sequence         *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderStages represents array of sale.order.stage model.
type SaleOrderStages []SaleOrderStage

// SaleOrderStageModel is the odoo model name.
const SaleOrderStageModel = "sale.order.stage"

// Many2One convert SaleOrderStage to *Many2One.
func (sos *SaleOrderStage) Many2One() *Many2One {
	return NewMany2One(sos.Id.Get(), "")
}

// CreateSaleOrderStage creates a new sale.order.stage model and returns its id.
func (c *Client) CreateSaleOrderStage(sos *SaleOrderStage) (int64, error) {
	ids, err := c.CreateSaleOrderStages([]*SaleOrderStage{sos})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderStage creates a new sale.order.stage model and returns its id.
func (c *Client) CreateSaleOrderStages(soss []*SaleOrderStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range soss {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderStageModel, vv)
}

// UpdateSaleOrderStage updates an existing sale.order.stage record.
func (c *Client) UpdateSaleOrderStage(sos *SaleOrderStage) error {
	return c.UpdateSaleOrderStages([]int64{sos.Id.Get()}, sos)
}

// UpdateSaleOrderStages updates existing sale.order.stage records.
// All records (represented by ids) will be updated by sos values.
func (c *Client) UpdateSaleOrderStages(ids []int64, sos *SaleOrderStage) error {
	return c.Update(SaleOrderStageModel, ids, sos)
}

// DeleteSaleOrderStage deletes an existing sale.order.stage record.
func (c *Client) DeleteSaleOrderStage(id int64) error {
	return c.DeleteSaleOrderStages([]int64{id})
}

// DeleteSaleOrderStages deletes existing sale.order.stage records.
func (c *Client) DeleteSaleOrderStages(ids []int64) error {
	return c.Delete(SaleOrderStageModel, ids)
}

// GetSaleOrderStage gets sale.order.stage existing record.
func (c *Client) GetSaleOrderStage(id int64) (*SaleOrderStage, error) {
	soss, err := c.GetSaleOrderStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if soss != nil && len(*soss) > 0 {
		return &((*soss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.stage not found", id)
}

// GetSaleOrderStages gets sale.order.stage existing records.
func (c *Client) GetSaleOrderStages(ids []int64) (*SaleOrderStages, error) {
	soss := &SaleOrderStages{}
	if err := c.Read(SaleOrderStageModel, ids, nil, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindSaleOrderStage finds sale.order.stage record by querying it with criteria.
func (c *Client) FindSaleOrderStage(criteria *Criteria) (*SaleOrderStage, error) {
	soss := &SaleOrderStages{}
	if err := c.SearchRead(SaleOrderStageModel, criteria, NewOptions().Limit(1), soss); err != nil {
		return nil, err
	}
	if soss != nil && len(*soss) > 0 {
		return &((*soss)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.stage was not found with criteria %v", criteria)
}

// FindSaleOrderStages finds sale.order.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderStages(criteria *Criteria, options *Options) (*SaleOrderStages, error) {
	soss := &SaleOrderStages{}
	if err := c.SearchRead(SaleOrderStageModel, criteria, options, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindSaleOrderStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderStageId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.stage was not found with criteria %v and options %v", criteria, options)
}
