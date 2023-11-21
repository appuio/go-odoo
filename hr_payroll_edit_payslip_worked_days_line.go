package odoo

import (
	"fmt"
)

// HrPayrollEditPayslipWorkedDaysLine represents hr.payroll.edit.payslip.worked.days.line model.
type HrPayrollEditPayslipWorkedDaysLine struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	Amount                   *Float    `xmlrpc:"amount,omptempty"`
	Code                     *String   `xmlrpc:"code,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	EditPayslipLinesWizardId *Many2One `xmlrpc:"edit_payslip_lines_wizard_id,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	NumberOfDays             *Float    `xmlrpc:"number_of_days,omptempty"`
	NumberOfHours            *Float    `xmlrpc:"number_of_hours,omptempty"`
	Sequence                 *Int      `xmlrpc:"sequence,omptempty"`
	SlipId                   *Many2One `xmlrpc:"slip_id,omptempty"`
	WorkEntryTypeId          *Many2One `xmlrpc:"work_entry_type_id,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayrollEditPayslipWorkedDaysLines represents array of hr.payroll.edit.payslip.worked.days.line model.
type HrPayrollEditPayslipWorkedDaysLines []HrPayrollEditPayslipWorkedDaysLine

// HrPayrollEditPayslipWorkedDaysLineModel is the odoo model name.
const HrPayrollEditPayslipWorkedDaysLineModel = "hr.payroll.edit.payslip.worked.days.line"

// Many2One convert HrPayrollEditPayslipWorkedDaysLine to *Many2One.
func (hpepwdl *HrPayrollEditPayslipWorkedDaysLine) Many2One() *Many2One {
	return NewMany2One(hpepwdl.Id.Get(), "")
}

// CreateHrPayrollEditPayslipWorkedDaysLine creates a new hr.payroll.edit.payslip.worked.days.line model and returns its id.
func (c *Client) CreateHrPayrollEditPayslipWorkedDaysLine(hpepwdl *HrPayrollEditPayslipWorkedDaysLine) (int64, error) {
	ids, err := c.CreateHrPayrollEditPayslipWorkedDaysLines([]*HrPayrollEditPayslipWorkedDaysLine{hpepwdl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayrollEditPayslipWorkedDaysLine creates a new hr.payroll.edit.payslip.worked.days.line model and returns its id.
func (c *Client) CreateHrPayrollEditPayslipWorkedDaysLines(hpepwdls []*HrPayrollEditPayslipWorkedDaysLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpepwdls {
		vv = append(vv, v)
	}
	return c.Create(HrPayrollEditPayslipWorkedDaysLineModel, vv)
}

// UpdateHrPayrollEditPayslipWorkedDaysLine updates an existing hr.payroll.edit.payslip.worked.days.line record.
func (c *Client) UpdateHrPayrollEditPayslipWorkedDaysLine(hpepwdl *HrPayrollEditPayslipWorkedDaysLine) error {
	return c.UpdateHrPayrollEditPayslipWorkedDaysLines([]int64{hpepwdl.Id.Get()}, hpepwdl)
}

// UpdateHrPayrollEditPayslipWorkedDaysLines updates existing hr.payroll.edit.payslip.worked.days.line records.
// All records (represented by ids) will be updated by hpepwdl values.
func (c *Client) UpdateHrPayrollEditPayslipWorkedDaysLines(ids []int64, hpepwdl *HrPayrollEditPayslipWorkedDaysLine) error {
	return c.Update(HrPayrollEditPayslipWorkedDaysLineModel, ids, hpepwdl)
}

// DeleteHrPayrollEditPayslipWorkedDaysLine deletes an existing hr.payroll.edit.payslip.worked.days.line record.
func (c *Client) DeleteHrPayrollEditPayslipWorkedDaysLine(id int64) error {
	return c.DeleteHrPayrollEditPayslipWorkedDaysLines([]int64{id})
}

// DeleteHrPayrollEditPayslipWorkedDaysLines deletes existing hr.payroll.edit.payslip.worked.days.line records.
func (c *Client) DeleteHrPayrollEditPayslipWorkedDaysLines(ids []int64) error {
	return c.Delete(HrPayrollEditPayslipWorkedDaysLineModel, ids)
}

// GetHrPayrollEditPayslipWorkedDaysLine gets hr.payroll.edit.payslip.worked.days.line existing record.
func (c *Client) GetHrPayrollEditPayslipWorkedDaysLine(id int64) (*HrPayrollEditPayslipWorkedDaysLine, error) {
	hpepwdls, err := c.GetHrPayrollEditPayslipWorkedDaysLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpepwdls != nil && len(*hpepwdls) > 0 {
		return &((*hpepwdls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payroll.edit.payslip.worked.days.line not found", id)
}

// GetHrPayrollEditPayslipWorkedDaysLines gets hr.payroll.edit.payslip.worked.days.line existing records.
func (c *Client) GetHrPayrollEditPayslipWorkedDaysLines(ids []int64) (*HrPayrollEditPayslipWorkedDaysLines, error) {
	hpepwdls := &HrPayrollEditPayslipWorkedDaysLines{}
	if err := c.Read(HrPayrollEditPayslipWorkedDaysLineModel, ids, nil, hpepwdls); err != nil {
		return nil, err
	}
	return hpepwdls, nil
}

// FindHrPayrollEditPayslipWorkedDaysLine finds hr.payroll.edit.payslip.worked.days.line record by querying it with criteria.
func (c *Client) FindHrPayrollEditPayslipWorkedDaysLine(criteria *Criteria) (*HrPayrollEditPayslipWorkedDaysLine, error) {
	hpepwdls := &HrPayrollEditPayslipWorkedDaysLines{}
	if err := c.SearchRead(HrPayrollEditPayslipWorkedDaysLineModel, criteria, NewOptions().Limit(1), hpepwdls); err != nil {
		return nil, err
	}
	if hpepwdls != nil && len(*hpepwdls) > 0 {
		return &((*hpepwdls)[0]), nil
	}
	return nil, fmt.Errorf("hr.payroll.edit.payslip.worked.days.line was not found with criteria %v", criteria)
}

// FindHrPayrollEditPayslipWorkedDaysLines finds hr.payroll.edit.payslip.worked.days.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollEditPayslipWorkedDaysLines(criteria *Criteria, options *Options) (*HrPayrollEditPayslipWorkedDaysLines, error) {
	hpepwdls := &HrPayrollEditPayslipWorkedDaysLines{}
	if err := c.SearchRead(HrPayrollEditPayslipWorkedDaysLineModel, criteria, options, hpepwdls); err != nil {
		return nil, err
	}
	return hpepwdls, nil
}

// FindHrPayrollEditPayslipWorkedDaysLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollEditPayslipWorkedDaysLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayrollEditPayslipWorkedDaysLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayrollEditPayslipWorkedDaysLineId finds record id by querying it with criteria.
func (c *Client) FindHrPayrollEditPayslipWorkedDaysLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayrollEditPayslipWorkedDaysLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payroll.edit.payslip.worked.days.line was not found with criteria %v and options %v", criteria, options)
}
