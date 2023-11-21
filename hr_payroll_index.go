package odoo

import (
	"fmt"
)

// HrPayrollIndex represents hr.payroll.index model.
type HrPayrollIndex struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	ContractIds    *Relation `xmlrpc:"contract_ids,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	Description    *String   `xmlrpc:"description,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	DisplayWarning *Bool     `xmlrpc:"display_warning,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	Percentage     *Float    `xmlrpc:"percentage,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayrollIndexs represents array of hr.payroll.index model.
type HrPayrollIndexs []HrPayrollIndex

// HrPayrollIndexModel is the odoo model name.
const HrPayrollIndexModel = "hr.payroll.index"

// Many2One convert HrPayrollIndex to *Many2One.
func (hpi *HrPayrollIndex) Many2One() *Many2One {
	return NewMany2One(hpi.Id.Get(), "")
}

// CreateHrPayrollIndex creates a new hr.payroll.index model and returns its id.
func (c *Client) CreateHrPayrollIndex(hpi *HrPayrollIndex) (int64, error) {
	ids, err := c.CreateHrPayrollIndexs([]*HrPayrollIndex{hpi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayrollIndex creates a new hr.payroll.index model and returns its id.
func (c *Client) CreateHrPayrollIndexs(hpis []*HrPayrollIndex) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpis {
		vv = append(vv, v)
	}
	return c.Create(HrPayrollIndexModel, vv)
}

// UpdateHrPayrollIndex updates an existing hr.payroll.index record.
func (c *Client) UpdateHrPayrollIndex(hpi *HrPayrollIndex) error {
	return c.UpdateHrPayrollIndexs([]int64{hpi.Id.Get()}, hpi)
}

// UpdateHrPayrollIndexs updates existing hr.payroll.index records.
// All records (represented by ids) will be updated by hpi values.
func (c *Client) UpdateHrPayrollIndexs(ids []int64, hpi *HrPayrollIndex) error {
	return c.Update(HrPayrollIndexModel, ids, hpi)
}

// DeleteHrPayrollIndex deletes an existing hr.payroll.index record.
func (c *Client) DeleteHrPayrollIndex(id int64) error {
	return c.DeleteHrPayrollIndexs([]int64{id})
}

// DeleteHrPayrollIndexs deletes existing hr.payroll.index records.
func (c *Client) DeleteHrPayrollIndexs(ids []int64) error {
	return c.Delete(HrPayrollIndexModel, ids)
}

// GetHrPayrollIndex gets hr.payroll.index existing record.
func (c *Client) GetHrPayrollIndex(id int64) (*HrPayrollIndex, error) {
	hpis, err := c.GetHrPayrollIndexs([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpis != nil && len(*hpis) > 0 {
		return &((*hpis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payroll.index not found", id)
}

// GetHrPayrollIndexs gets hr.payroll.index existing records.
func (c *Client) GetHrPayrollIndexs(ids []int64) (*HrPayrollIndexs, error) {
	hpis := &HrPayrollIndexs{}
	if err := c.Read(HrPayrollIndexModel, ids, nil, hpis); err != nil {
		return nil, err
	}
	return hpis, nil
}

// FindHrPayrollIndex finds hr.payroll.index record by querying it with criteria.
func (c *Client) FindHrPayrollIndex(criteria *Criteria) (*HrPayrollIndex, error) {
	hpis := &HrPayrollIndexs{}
	if err := c.SearchRead(HrPayrollIndexModel, criteria, NewOptions().Limit(1), hpis); err != nil {
		return nil, err
	}
	if hpis != nil && len(*hpis) > 0 {
		return &((*hpis)[0]), nil
	}
	return nil, fmt.Errorf("hr.payroll.index was not found with criteria %v", criteria)
}

// FindHrPayrollIndexs finds hr.payroll.index records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollIndexs(criteria *Criteria, options *Options) (*HrPayrollIndexs, error) {
	hpis := &HrPayrollIndexs{}
	if err := c.SearchRead(HrPayrollIndexModel, criteria, options, hpis); err != nil {
		return nil, err
	}
	return hpis, nil
}

// FindHrPayrollIndexIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollIndexIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayrollIndexModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayrollIndexId finds record id by querying it with criteria.
func (c *Client) FindHrPayrollIndexId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayrollIndexModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payroll.index was not found with criteria %v and options %v", criteria, options)
}
