package odoo

import (
	"fmt"
)

// HrDepartureReason represents hr.departure.reason model.
type HrDepartureReason struct {
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

// HrDepartureReasons represents array of hr.departure.reason model.
type HrDepartureReasons []HrDepartureReason

// HrDepartureReasonModel is the odoo model name.
const HrDepartureReasonModel = "hr.departure.reason"

// Many2One convert HrDepartureReason to *Many2One.
func (hdr *HrDepartureReason) Many2One() *Many2One {
	return NewMany2One(hdr.Id.Get(), "")
}

// CreateHrDepartureReason creates a new hr.departure.reason model and returns its id.
func (c *Client) CreateHrDepartureReason(hdr *HrDepartureReason) (int64, error) {
	ids, err := c.CreateHrDepartureReasons([]*HrDepartureReason{hdr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrDepartureReason creates a new hr.departure.reason model and returns its id.
func (c *Client) CreateHrDepartureReasons(hdrs []*HrDepartureReason) ([]int64, error) {
	var vv []interface{}
	for _, v := range hdrs {
		vv = append(vv, v)
	}
	return c.Create(HrDepartureReasonModel, vv)
}

// UpdateHrDepartureReason updates an existing hr.departure.reason record.
func (c *Client) UpdateHrDepartureReason(hdr *HrDepartureReason) error {
	return c.UpdateHrDepartureReasons([]int64{hdr.Id.Get()}, hdr)
}

// UpdateHrDepartureReasons updates existing hr.departure.reason records.
// All records (represented by ids) will be updated by hdr values.
func (c *Client) UpdateHrDepartureReasons(ids []int64, hdr *HrDepartureReason) error {
	return c.Update(HrDepartureReasonModel, ids, hdr)
}

// DeleteHrDepartureReason deletes an existing hr.departure.reason record.
func (c *Client) DeleteHrDepartureReason(id int64) error {
	return c.DeleteHrDepartureReasons([]int64{id})
}

// DeleteHrDepartureReasons deletes existing hr.departure.reason records.
func (c *Client) DeleteHrDepartureReasons(ids []int64) error {
	return c.Delete(HrDepartureReasonModel, ids)
}

// GetHrDepartureReason gets hr.departure.reason existing record.
func (c *Client) GetHrDepartureReason(id int64) (*HrDepartureReason, error) {
	hdrs, err := c.GetHrDepartureReasons([]int64{id})
	if err != nil {
		return nil, err
	}
	if hdrs != nil && len(*hdrs) > 0 {
		return &((*hdrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.departure.reason not found", id)
}

// GetHrDepartureReasons gets hr.departure.reason existing records.
func (c *Client) GetHrDepartureReasons(ids []int64) (*HrDepartureReasons, error) {
	hdrs := &HrDepartureReasons{}
	if err := c.Read(HrDepartureReasonModel, ids, nil, hdrs); err != nil {
		return nil, err
	}
	return hdrs, nil
}

// FindHrDepartureReason finds hr.departure.reason record by querying it with criteria.
func (c *Client) FindHrDepartureReason(criteria *Criteria) (*HrDepartureReason, error) {
	hdrs := &HrDepartureReasons{}
	if err := c.SearchRead(HrDepartureReasonModel, criteria, NewOptions().Limit(1), hdrs); err != nil {
		return nil, err
	}
	if hdrs != nil && len(*hdrs) > 0 {
		return &((*hdrs)[0]), nil
	}
	return nil, fmt.Errorf("hr.departure.reason was not found with criteria %v", criteria)
}

// FindHrDepartureReasons finds hr.departure.reason records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartureReasons(criteria *Criteria, options *Options) (*HrDepartureReasons, error) {
	hdrs := &HrDepartureReasons{}
	if err := c.SearchRead(HrDepartureReasonModel, criteria, options, hdrs); err != nil {
		return nil, err
	}
	return hdrs, nil
}

// FindHrDepartureReasonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrDepartureReasonIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrDepartureReasonModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrDepartureReasonId finds record id by querying it with criteria.
func (c *Client) FindHrDepartureReasonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrDepartureReasonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.departure.reason was not found with criteria %v and options %v", criteria, options)
}
