package odoo

import (
	"fmt"
)

// HrPayslipLine represents hr.payslip.line model.
type HrPayslipLine struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	Amount           *Float     `xmlrpc:"amount,omptempty"`
	AmountFix        *Float     `xmlrpc:"amount_fix,omptempty"`
	AmountPercentage *Float     `xmlrpc:"amount_percentage,omptempty"`
	AmountSelect     *Selection `xmlrpc:"amount_select,omptempty"`
	AppearsOnPayslip *Bool      `xmlrpc:"appears_on_payslip,omptempty"`
	CategoryId       *Many2One  `xmlrpc:"category_id,omptempty"`
	Code             *String    `xmlrpc:"code,omptempty"`
	CompanyId        *Many2One  `xmlrpc:"company_id,omptempty"`
	ContractId       *Many2One  `xmlrpc:"contract_id,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId       *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateFrom         *Time      `xmlrpc:"date_from,omptempty"`
	DateTo           *Time      `xmlrpc:"date_to,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId       *Many2One  `xmlrpc:"employee_id,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	Note             *String    `xmlrpc:"note,omptempty"`
	PartnerId        *Many2One  `xmlrpc:"partner_id,omptempty"`
	Quantity         *Float     `xmlrpc:"quantity,omptempty"`
	Rate             *Float     `xmlrpc:"rate,omptempty"`
	SalaryRuleId     *Many2One  `xmlrpc:"salary_rule_id,omptempty"`
	Sequence         *Int       `xmlrpc:"sequence,omptempty"`
	SlipId           *Many2One  `xmlrpc:"slip_id,omptempty"`
	Total            *Float     `xmlrpc:"total,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipLines represents array of hr.payslip.line model.
type HrPayslipLines []HrPayslipLine

// HrPayslipLineModel is the odoo model name.
const HrPayslipLineModel = "hr.payslip.line"

// Many2One convert HrPayslipLine to *Many2One.
func (hpl *HrPayslipLine) Many2One() *Many2One {
	return NewMany2One(hpl.Id.Get(), "")
}

// CreateHrPayslipLine creates a new hr.payslip.line model and returns its id.
func (c *Client) CreateHrPayslipLine(hpl *HrPayslipLine) (int64, error) {
	ids, err := c.CreateHrPayslipLines([]*HrPayslipLine{hpl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipLine creates a new hr.payslip.line model and returns its id.
func (c *Client) CreateHrPayslipLines(hpls []*HrPayslipLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpls {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipLineModel, vv)
}

// UpdateHrPayslipLine updates an existing hr.payslip.line record.
func (c *Client) UpdateHrPayslipLine(hpl *HrPayslipLine) error {
	return c.UpdateHrPayslipLines([]int64{hpl.Id.Get()}, hpl)
}

// UpdateHrPayslipLines updates existing hr.payslip.line records.
// All records (represented by ids) will be updated by hpl values.
func (c *Client) UpdateHrPayslipLines(ids []int64, hpl *HrPayslipLine) error {
	return c.Update(HrPayslipLineModel, ids, hpl)
}

// DeleteHrPayslipLine deletes an existing hr.payslip.line record.
func (c *Client) DeleteHrPayslipLine(id int64) error {
	return c.DeleteHrPayslipLines([]int64{id})
}

// DeleteHrPayslipLines deletes existing hr.payslip.line records.
func (c *Client) DeleteHrPayslipLines(ids []int64) error {
	return c.Delete(HrPayslipLineModel, ids)
}

// GetHrPayslipLine gets hr.payslip.line existing record.
func (c *Client) GetHrPayslipLine(id int64) (*HrPayslipLine, error) {
	hpls, err := c.GetHrPayslipLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpls != nil && len(*hpls) > 0 {
		return &((*hpls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.line not found", id)
}

// GetHrPayslipLines gets hr.payslip.line existing records.
func (c *Client) GetHrPayslipLines(ids []int64) (*HrPayslipLines, error) {
	hpls := &HrPayslipLines{}
	if err := c.Read(HrPayslipLineModel, ids, nil, hpls); err != nil {
		return nil, err
	}
	return hpls, nil
}

// FindHrPayslipLine finds hr.payslip.line record by querying it with criteria.
func (c *Client) FindHrPayslipLine(criteria *Criteria) (*HrPayslipLine, error) {
	hpls := &HrPayslipLines{}
	if err := c.SearchRead(HrPayslipLineModel, criteria, NewOptions().Limit(1), hpls); err != nil {
		return nil, err
	}
	if hpls != nil && len(*hpls) > 0 {
		return &((*hpls)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.line was not found with criteria %v", criteria)
}

// FindHrPayslipLines finds hr.payslip.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipLines(criteria *Criteria, options *Options) (*HrPayslipLines, error) {
	hpls := &HrPayslipLines{}
	if err := c.SearchRead(HrPayslipLineModel, criteria, options, hpls); err != nil {
		return nil, err
	}
	return hpls, nil
}

// FindHrPayslipLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipLineId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.line was not found with criteria %v and options %v", criteria, options)
}
