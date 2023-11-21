package odoo

import (
	"fmt"
)

// VshnMeteredUsage represents vshn.metered.usage model.
type VshnMeteredUsage struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	AccountAnalyticLineId *Many2One  `xmlrpc:"account_analytic_line_id,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	ErrorMsg              *String    `xmlrpc:"error_msg,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	InstanceId            *String    `xmlrpc:"instance_id,omptempty"`
	ItemDescription       *String    `xmlrpc:"item_description,omptempty"`
	ItemGroupDescription  *String    `xmlrpc:"item_group_description,omptempty"`
	OrderId               *Many2One  `xmlrpc:"order_id,omptempty"`
	PartnerId             *Many2One  `xmlrpc:"partner_id,omptempty"`
	ProductId             *Many2One  `xmlrpc:"product_id,omptempty"`
	Quantity              *Float     `xmlrpc:"quantity,omptempty"`
	State                 *Selection `xmlrpc:"state,omptempty"`
	TimerangeEnd          *Time      `xmlrpc:"timerange_end,omptempty"`
	TimerangeStart        *Time      `xmlrpc:"timerange_start,omptempty"`
	Timestamp             *Time      `xmlrpc:"timestamp,omptempty"`
	UomId                 *Many2One  `xmlrpc:"uom_id,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// VshnMeteredUsages represents array of vshn.metered.usage model.
type VshnMeteredUsages []VshnMeteredUsage

// VshnMeteredUsageModel is the odoo model name.
const VshnMeteredUsageModel = "vshn.metered.usage"

// Many2One convert VshnMeteredUsage to *Many2One.
func (vmu *VshnMeteredUsage) Many2One() *Many2One {
	return NewMany2One(vmu.Id.Get(), "")
}

// CreateVshnMeteredUsage creates a new vshn.metered.usage model and returns its id.
func (c *Client) CreateVshnMeteredUsage(vmu *VshnMeteredUsage) (int64, error) {
	ids, err := c.CreateVshnMeteredUsages([]*VshnMeteredUsage{vmu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateVshnMeteredUsage creates a new vshn.metered.usage model and returns its id.
func (c *Client) CreateVshnMeteredUsages(vmus []*VshnMeteredUsage) ([]int64, error) {
	var vv []interface{}
	for _, v := range vmus {
		vv = append(vv, v)
	}
	return c.Create(VshnMeteredUsageModel, vv)
}

// UpdateVshnMeteredUsage updates an existing vshn.metered.usage record.
func (c *Client) UpdateVshnMeteredUsage(vmu *VshnMeteredUsage) error {
	return c.UpdateVshnMeteredUsages([]int64{vmu.Id.Get()}, vmu)
}

// UpdateVshnMeteredUsages updates existing vshn.metered.usage records.
// All records (represented by ids) will be updated by vmu values.
func (c *Client) UpdateVshnMeteredUsages(ids []int64, vmu *VshnMeteredUsage) error {
	return c.Update(VshnMeteredUsageModel, ids, vmu)
}

// DeleteVshnMeteredUsage deletes an existing vshn.metered.usage record.
func (c *Client) DeleteVshnMeteredUsage(id int64) error {
	return c.DeleteVshnMeteredUsages([]int64{id})
}

// DeleteVshnMeteredUsages deletes existing vshn.metered.usage records.
func (c *Client) DeleteVshnMeteredUsages(ids []int64) error {
	return c.Delete(VshnMeteredUsageModel, ids)
}

// GetVshnMeteredUsage gets vshn.metered.usage existing record.
func (c *Client) GetVshnMeteredUsage(id int64) (*VshnMeteredUsage, error) {
	vmus, err := c.GetVshnMeteredUsages([]int64{id})
	if err != nil {
		return nil, err
	}
	if vmus != nil && len(*vmus) > 0 {
		return &((*vmus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of vshn.metered.usage not found", id)
}

// GetVshnMeteredUsages gets vshn.metered.usage existing records.
func (c *Client) GetVshnMeteredUsages(ids []int64) (*VshnMeteredUsages, error) {
	vmus := &VshnMeteredUsages{}
	if err := c.Read(VshnMeteredUsageModel, ids, nil, vmus); err != nil {
		return nil, err
	}
	return vmus, nil
}

// FindVshnMeteredUsage finds vshn.metered.usage record by querying it with criteria.
func (c *Client) FindVshnMeteredUsage(criteria *Criteria) (*VshnMeteredUsage, error) {
	vmus := &VshnMeteredUsages{}
	if err := c.SearchRead(VshnMeteredUsageModel, criteria, NewOptions().Limit(1), vmus); err != nil {
		return nil, err
	}
	if vmus != nil && len(*vmus) > 0 {
		return &((*vmus)[0]), nil
	}
	return nil, fmt.Errorf("vshn.metered.usage was not found with criteria %v", criteria)
}

// FindVshnMeteredUsages finds vshn.metered.usage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindVshnMeteredUsages(criteria *Criteria, options *Options) (*VshnMeteredUsages, error) {
	vmus := &VshnMeteredUsages{}
	if err := c.SearchRead(VshnMeteredUsageModel, criteria, options, vmus); err != nil {
		return nil, err
	}
	return vmus, nil
}

// FindVshnMeteredUsageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindVshnMeteredUsageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(VshnMeteredUsageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindVshnMeteredUsageId finds record id by querying it with criteria.
func (c *Client) FindVshnMeteredUsageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(VshnMeteredUsageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("vshn.metered.usage was not found with criteria %v and options %v", criteria, options)
}
