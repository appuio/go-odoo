package odoo

import (
	"fmt"
)

// HrPayslipInputType represents hr.payslip.input.type model.
type HrPayslipInputType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Code        *String   `xmlrpc:"code,omptempty"`
	CountryId   *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	StructIds   *Relation `xmlrpc:"struct_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipInputTypes represents array of hr.payslip.input.type model.
type HrPayslipInputTypes []HrPayslipInputType

// HrPayslipInputTypeModel is the odoo model name.
const HrPayslipInputTypeModel = "hr.payslip.input.type"

// Many2One convert HrPayslipInputType to *Many2One.
func (hpit *HrPayslipInputType) Many2One() *Many2One {
	return NewMany2One(hpit.Id.Get(), "")
}

// CreateHrPayslipInputType creates a new hr.payslip.input.type model and returns its id.
func (c *Client) CreateHrPayslipInputType(hpit *HrPayslipInputType) (int64, error) {
	ids, err := c.CreateHrPayslipInputTypes([]*HrPayslipInputType{hpit})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipInputType creates a new hr.payslip.input.type model and returns its id.
func (c *Client) CreateHrPayslipInputTypes(hpits []*HrPayslipInputType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpits {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipInputTypeModel, vv)
}

// UpdateHrPayslipInputType updates an existing hr.payslip.input.type record.
func (c *Client) UpdateHrPayslipInputType(hpit *HrPayslipInputType) error {
	return c.UpdateHrPayslipInputTypes([]int64{hpit.Id.Get()}, hpit)
}

// UpdateHrPayslipInputTypes updates existing hr.payslip.input.type records.
// All records (represented by ids) will be updated by hpit values.
func (c *Client) UpdateHrPayslipInputTypes(ids []int64, hpit *HrPayslipInputType) error {
	return c.Update(HrPayslipInputTypeModel, ids, hpit)
}

// DeleteHrPayslipInputType deletes an existing hr.payslip.input.type record.
func (c *Client) DeleteHrPayslipInputType(id int64) error {
	return c.DeleteHrPayslipInputTypes([]int64{id})
}

// DeleteHrPayslipInputTypes deletes existing hr.payslip.input.type records.
func (c *Client) DeleteHrPayslipInputTypes(ids []int64) error {
	return c.Delete(HrPayslipInputTypeModel, ids)
}

// GetHrPayslipInputType gets hr.payslip.input.type existing record.
func (c *Client) GetHrPayslipInputType(id int64) (*HrPayslipInputType, error) {
	hpits, err := c.GetHrPayslipInputTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpits != nil && len(*hpits) > 0 {
		return &((*hpits)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.input.type not found", id)
}

// GetHrPayslipInputTypes gets hr.payslip.input.type existing records.
func (c *Client) GetHrPayslipInputTypes(ids []int64) (*HrPayslipInputTypes, error) {
	hpits := &HrPayslipInputTypes{}
	if err := c.Read(HrPayslipInputTypeModel, ids, nil, hpits); err != nil {
		return nil, err
	}
	return hpits, nil
}

// FindHrPayslipInputType finds hr.payslip.input.type record by querying it with criteria.
func (c *Client) FindHrPayslipInputType(criteria *Criteria) (*HrPayslipInputType, error) {
	hpits := &HrPayslipInputTypes{}
	if err := c.SearchRead(HrPayslipInputTypeModel, criteria, NewOptions().Limit(1), hpits); err != nil {
		return nil, err
	}
	if hpits != nil && len(*hpits) > 0 {
		return &((*hpits)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.input.type was not found with criteria %v", criteria)
}

// FindHrPayslipInputTypes finds hr.payslip.input.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipInputTypes(criteria *Criteria, options *Options) (*HrPayslipInputTypes, error) {
	hpits := &HrPayslipInputTypes{}
	if err := c.SearchRead(HrPayslipInputTypeModel, criteria, options, hpits); err != nil {
		return nil, err
	}
	return hpits, nil
}

// FindHrPayslipInputTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipInputTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipInputTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipInputTypeId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipInputTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipInputTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.input.type was not found with criteria %v and options %v", criteria, options)
}
