package odoo

import (
	"fmt"
)

// HrPayrollEditPayslipLine represents hr.payroll.edit.payslip.line model.
type HrPayrollEditPayslipLine struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	Amount                   *Float    `xmlrpc:"amount,omptempty"`
	CategoryId               *Many2One `xmlrpc:"category_id,omptempty"`
	Code                     *String   `xmlrpc:"code,omptempty"`
	ContractId               *Many2One `xmlrpc:"contract_id,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	EditPayslipLinesWizardId *Many2One `xmlrpc:"edit_payslip_lines_wizard_id,omptempty"`
	EmployeeId               *Many2One `xmlrpc:"employee_id,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	Note                     *String   `xmlrpc:"note,omptempty"`
	Quantity                 *Float    `xmlrpc:"quantity,omptempty"`
	Rate                     *Float    `xmlrpc:"rate,omptempty"`
	SalaryRuleId             *Many2One `xmlrpc:"salary_rule_id,omptempty"`
	Sequence                 *Int      `xmlrpc:"sequence,omptempty"`
	SlipId                   *Many2One `xmlrpc:"slip_id,omptempty"`
	StructId                 *Many2One `xmlrpc:"struct_id,omptempty"`
	Total                    *Float    `xmlrpc:"total,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayrollEditPayslipLines represents array of hr.payroll.edit.payslip.line model.
type HrPayrollEditPayslipLines []HrPayrollEditPayslipLine

// HrPayrollEditPayslipLineModel is the odoo model name.
const HrPayrollEditPayslipLineModel = "hr.payroll.edit.payslip.line"

// Many2One convert HrPayrollEditPayslipLine to *Many2One.
func (hpepl *HrPayrollEditPayslipLine) Many2One() *Many2One {
	return NewMany2One(hpepl.Id.Get(), "")
}

// CreateHrPayrollEditPayslipLine creates a new hr.payroll.edit.payslip.line model and returns its id.
func (c *Client) CreateHrPayrollEditPayslipLine(hpepl *HrPayrollEditPayslipLine) (int64, error) {
	ids, err := c.CreateHrPayrollEditPayslipLines([]*HrPayrollEditPayslipLine{hpepl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayrollEditPayslipLine creates a new hr.payroll.edit.payslip.line model and returns its id.
func (c *Client) CreateHrPayrollEditPayslipLines(hpepls []*HrPayrollEditPayslipLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpepls {
		vv = append(vv, v)
	}
	return c.Create(HrPayrollEditPayslipLineModel, vv)
}

// UpdateHrPayrollEditPayslipLine updates an existing hr.payroll.edit.payslip.line record.
func (c *Client) UpdateHrPayrollEditPayslipLine(hpepl *HrPayrollEditPayslipLine) error {
	return c.UpdateHrPayrollEditPayslipLines([]int64{hpepl.Id.Get()}, hpepl)
}

// UpdateHrPayrollEditPayslipLines updates existing hr.payroll.edit.payslip.line records.
// All records (represented by ids) will be updated by hpepl values.
func (c *Client) UpdateHrPayrollEditPayslipLines(ids []int64, hpepl *HrPayrollEditPayslipLine) error {
	return c.Update(HrPayrollEditPayslipLineModel, ids, hpepl)
}

// DeleteHrPayrollEditPayslipLine deletes an existing hr.payroll.edit.payslip.line record.
func (c *Client) DeleteHrPayrollEditPayslipLine(id int64) error {
	return c.DeleteHrPayrollEditPayslipLines([]int64{id})
}

// DeleteHrPayrollEditPayslipLines deletes existing hr.payroll.edit.payslip.line records.
func (c *Client) DeleteHrPayrollEditPayslipLines(ids []int64) error {
	return c.Delete(HrPayrollEditPayslipLineModel, ids)
}

// GetHrPayrollEditPayslipLine gets hr.payroll.edit.payslip.line existing record.
func (c *Client) GetHrPayrollEditPayslipLine(id int64) (*HrPayrollEditPayslipLine, error) {
	hpepls, err := c.GetHrPayrollEditPayslipLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpepls != nil && len(*hpepls) > 0 {
		return &((*hpepls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payroll.edit.payslip.line not found", id)
}

// GetHrPayrollEditPayslipLines gets hr.payroll.edit.payslip.line existing records.
func (c *Client) GetHrPayrollEditPayslipLines(ids []int64) (*HrPayrollEditPayslipLines, error) {
	hpepls := &HrPayrollEditPayslipLines{}
	if err := c.Read(HrPayrollEditPayslipLineModel, ids, nil, hpepls); err != nil {
		return nil, err
	}
	return hpepls, nil
}

// FindHrPayrollEditPayslipLine finds hr.payroll.edit.payslip.line record by querying it with criteria.
func (c *Client) FindHrPayrollEditPayslipLine(criteria *Criteria) (*HrPayrollEditPayslipLine, error) {
	hpepls := &HrPayrollEditPayslipLines{}
	if err := c.SearchRead(HrPayrollEditPayslipLineModel, criteria, NewOptions().Limit(1), hpepls); err != nil {
		return nil, err
	}
	if hpepls != nil && len(*hpepls) > 0 {
		return &((*hpepls)[0]), nil
	}
	return nil, fmt.Errorf("hr.payroll.edit.payslip.line was not found with criteria %v", criteria)
}

// FindHrPayrollEditPayslipLines finds hr.payroll.edit.payslip.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollEditPayslipLines(criteria *Criteria, options *Options) (*HrPayrollEditPayslipLines, error) {
	hpepls := &HrPayrollEditPayslipLines{}
	if err := c.SearchRead(HrPayrollEditPayslipLineModel, criteria, options, hpepls); err != nil {
		return nil, err
	}
	return hpepls, nil
}

// FindHrPayrollEditPayslipLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollEditPayslipLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayrollEditPayslipLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayrollEditPayslipLineId finds record id by querying it with criteria.
func (c *Client) FindHrPayrollEditPayslipLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayrollEditPayslipLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payroll.edit.payslip.line was not found with criteria %v and options %v", criteria, options)
}
