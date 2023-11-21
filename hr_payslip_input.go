package odoo

import (
	"fmt"
)

// HrPayslipInput represents hr.payslip.input model.
type HrPayslipInput struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	AllowedInputTypeIds *Relation `xmlrpc:"_allowed_input_type_ids,omptempty"`
	Amount              *Float    `xmlrpc:"amount,omptempty"`
	Code                *String   `xmlrpc:"code,omptempty"`
	ContractId          *Many2One `xmlrpc:"contract_id,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	InputTypeId         *Many2One `xmlrpc:"input_type_id,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	PayslipId           *Many2One `xmlrpc:"payslip_id,omptempty"`
	Sequence            *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipInputs represents array of hr.payslip.input model.
type HrPayslipInputs []HrPayslipInput

// HrPayslipInputModel is the odoo model name.
const HrPayslipInputModel = "hr.payslip.input"

// Many2One convert HrPayslipInput to *Many2One.
func (hpi *HrPayslipInput) Many2One() *Many2One {
	return NewMany2One(hpi.Id.Get(), "")
}

// CreateHrPayslipInput creates a new hr.payslip.input model and returns its id.
func (c *Client) CreateHrPayslipInput(hpi *HrPayslipInput) (int64, error) {
	ids, err := c.CreateHrPayslipInputs([]*HrPayslipInput{hpi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipInput creates a new hr.payslip.input model and returns its id.
func (c *Client) CreateHrPayslipInputs(hpis []*HrPayslipInput) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpis {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipInputModel, vv)
}

// UpdateHrPayslipInput updates an existing hr.payslip.input record.
func (c *Client) UpdateHrPayslipInput(hpi *HrPayslipInput) error {
	return c.UpdateHrPayslipInputs([]int64{hpi.Id.Get()}, hpi)
}

// UpdateHrPayslipInputs updates existing hr.payslip.input records.
// All records (represented by ids) will be updated by hpi values.
func (c *Client) UpdateHrPayslipInputs(ids []int64, hpi *HrPayslipInput) error {
	return c.Update(HrPayslipInputModel, ids, hpi)
}

// DeleteHrPayslipInput deletes an existing hr.payslip.input record.
func (c *Client) DeleteHrPayslipInput(id int64) error {
	return c.DeleteHrPayslipInputs([]int64{id})
}

// DeleteHrPayslipInputs deletes existing hr.payslip.input records.
func (c *Client) DeleteHrPayslipInputs(ids []int64) error {
	return c.Delete(HrPayslipInputModel, ids)
}

// GetHrPayslipInput gets hr.payslip.input existing record.
func (c *Client) GetHrPayslipInput(id int64) (*HrPayslipInput, error) {
	hpis, err := c.GetHrPayslipInputs([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpis != nil && len(*hpis) > 0 {
		return &((*hpis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.input not found", id)
}

// GetHrPayslipInputs gets hr.payslip.input existing records.
func (c *Client) GetHrPayslipInputs(ids []int64) (*HrPayslipInputs, error) {
	hpis := &HrPayslipInputs{}
	if err := c.Read(HrPayslipInputModel, ids, nil, hpis); err != nil {
		return nil, err
	}
	return hpis, nil
}

// FindHrPayslipInput finds hr.payslip.input record by querying it with criteria.
func (c *Client) FindHrPayslipInput(criteria *Criteria) (*HrPayslipInput, error) {
	hpis := &HrPayslipInputs{}
	if err := c.SearchRead(HrPayslipInputModel, criteria, NewOptions().Limit(1), hpis); err != nil {
		return nil, err
	}
	if hpis != nil && len(*hpis) > 0 {
		return &((*hpis)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.input was not found with criteria %v", criteria)
}

// FindHrPayslipInputs finds hr.payslip.input records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipInputs(criteria *Criteria, options *Options) (*HrPayslipInputs, error) {
	hpis := &HrPayslipInputs{}
	if err := c.SearchRead(HrPayslipInputModel, criteria, options, hpis); err != nil {
		return nil, err
	}
	return hpis, nil
}

// FindHrPayslipInputIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipInputIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipInputModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipInputId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipInputId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipInputModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.input was not found with criteria %v and options %v", criteria, options)
}
