package odoo

import (
	"fmt"
)

// IrMetric represents ir.metric model.
type IrMetric struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	Active      *Bool      `xmlrpc:"active,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description *String    `xmlrpc:"description,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Domain      *String    `xmlrpc:"domain,omptempty"`
	Field       *String    `xmlrpc:"field,omptempty"`
	FieldId     *Many2One  `xmlrpc:"field_id,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Model       *String    `xmlrpc:"model,omptempty"`
	ModelId     *Many2One  `xmlrpc:"model_id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	Operation   *Selection `xmlrpc:"operation,omptempty"`
	Type        *Selection `xmlrpc:"type,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// IrMetrics represents array of ir.metric model.
type IrMetrics []IrMetric

// IrMetricModel is the odoo model name.
const IrMetricModel = "ir.metric"

// Many2One convert IrMetric to *Many2One.
func (im *IrMetric) Many2One() *Many2One {
	return NewMany2One(im.Id.Get(), "")
}

// CreateIrMetric creates a new ir.metric model and returns its id.
func (c *Client) CreateIrMetric(im *IrMetric) (int64, error) {
	ids, err := c.CreateIrMetrics([]*IrMetric{im})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrMetric creates a new ir.metric model and returns its id.
func (c *Client) CreateIrMetrics(ims []*IrMetric) ([]int64, error) {
	var vv []interface{}
	for _, v := range ims {
		vv = append(vv, v)
	}
	return c.Create(IrMetricModel, vv)
}

// UpdateIrMetric updates an existing ir.metric record.
func (c *Client) UpdateIrMetric(im *IrMetric) error {
	return c.UpdateIrMetrics([]int64{im.Id.Get()}, im)
}

// UpdateIrMetrics updates existing ir.metric records.
// All records (represented by ids) will be updated by im values.
func (c *Client) UpdateIrMetrics(ids []int64, im *IrMetric) error {
	return c.Update(IrMetricModel, ids, im)
}

// DeleteIrMetric deletes an existing ir.metric record.
func (c *Client) DeleteIrMetric(id int64) error {
	return c.DeleteIrMetrics([]int64{id})
}

// DeleteIrMetrics deletes existing ir.metric records.
func (c *Client) DeleteIrMetrics(ids []int64) error {
	return c.Delete(IrMetricModel, ids)
}

// GetIrMetric gets ir.metric existing record.
func (c *Client) GetIrMetric(id int64) (*IrMetric, error) {
	ims, err := c.GetIrMetrics([]int64{id})
	if err != nil {
		return nil, err
	}
	if ims != nil && len(*ims) > 0 {
		return &((*ims)[0]), nil
	}
	return nil, fmt.Errorf("id %v of ir.metric not found", id)
}

// GetIrMetrics gets ir.metric existing records.
func (c *Client) GetIrMetrics(ids []int64) (*IrMetrics, error) {
	ims := &IrMetrics{}
	if err := c.Read(IrMetricModel, ids, nil, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindIrMetric finds ir.metric record by querying it with criteria.
func (c *Client) FindIrMetric(criteria *Criteria) (*IrMetric, error) {
	ims := &IrMetrics{}
	if err := c.SearchRead(IrMetricModel, criteria, NewOptions().Limit(1), ims); err != nil {
		return nil, err
	}
	if ims != nil && len(*ims) > 0 {
		return &((*ims)[0]), nil
	}
	return nil, fmt.Errorf("ir.metric was not found with criteria %v", criteria)
}

// FindIrMetrics finds ir.metric records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrMetrics(criteria *Criteria, options *Options) (*IrMetrics, error) {
	ims := &IrMetrics{}
	if err := c.SearchRead(IrMetricModel, criteria, options, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindIrMetricIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrMetricIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(IrMetricModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindIrMetricId finds record id by querying it with criteria.
func (c *Client) FindIrMetricId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrMetricModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("ir.metric was not found with criteria %v and options %v", criteria, options)
}
