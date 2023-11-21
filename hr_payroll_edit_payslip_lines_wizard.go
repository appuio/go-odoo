package odoo

import (
	"fmt"
)

// HrPayrollEditPayslipLinesWizard represents hr.payroll.edit.payslip.lines.wizard model.
type HrPayrollEditPayslipLinesWizard struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	LineIds           *Relation `xmlrpc:"line_ids,omptempty"`
	PayslipId         *Many2One `xmlrpc:"payslip_id,omptempty"`
	WorkedDaysLineIds *Relation `xmlrpc:"worked_days_line_ids,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayrollEditPayslipLinesWizards represents array of hr.payroll.edit.payslip.lines.wizard model.
type HrPayrollEditPayslipLinesWizards []HrPayrollEditPayslipLinesWizard

// HrPayrollEditPayslipLinesWizardModel is the odoo model name.
const HrPayrollEditPayslipLinesWizardModel = "hr.payroll.edit.payslip.lines.wizard"

// Many2One convert HrPayrollEditPayslipLinesWizard to *Many2One.
func (hpeplw *HrPayrollEditPayslipLinesWizard) Many2One() *Many2One {
	return NewMany2One(hpeplw.Id.Get(), "")
}

// CreateHrPayrollEditPayslipLinesWizard creates a new hr.payroll.edit.payslip.lines.wizard model and returns its id.
func (c *Client) CreateHrPayrollEditPayslipLinesWizard(hpeplw *HrPayrollEditPayslipLinesWizard) (int64, error) {
	ids, err := c.CreateHrPayrollEditPayslipLinesWizards([]*HrPayrollEditPayslipLinesWizard{hpeplw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayrollEditPayslipLinesWizard creates a new hr.payroll.edit.payslip.lines.wizard model and returns its id.
func (c *Client) CreateHrPayrollEditPayslipLinesWizards(hpeplws []*HrPayrollEditPayslipLinesWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpeplws {
		vv = append(vv, v)
	}
	return c.Create(HrPayrollEditPayslipLinesWizardModel, vv)
}

// UpdateHrPayrollEditPayslipLinesWizard updates an existing hr.payroll.edit.payslip.lines.wizard record.
func (c *Client) UpdateHrPayrollEditPayslipLinesWizard(hpeplw *HrPayrollEditPayslipLinesWizard) error {
	return c.UpdateHrPayrollEditPayslipLinesWizards([]int64{hpeplw.Id.Get()}, hpeplw)
}

// UpdateHrPayrollEditPayslipLinesWizards updates existing hr.payroll.edit.payslip.lines.wizard records.
// All records (represented by ids) will be updated by hpeplw values.
func (c *Client) UpdateHrPayrollEditPayslipLinesWizards(ids []int64, hpeplw *HrPayrollEditPayslipLinesWizard) error {
	return c.Update(HrPayrollEditPayslipLinesWizardModel, ids, hpeplw)
}

// DeleteHrPayrollEditPayslipLinesWizard deletes an existing hr.payroll.edit.payslip.lines.wizard record.
func (c *Client) DeleteHrPayrollEditPayslipLinesWizard(id int64) error {
	return c.DeleteHrPayrollEditPayslipLinesWizards([]int64{id})
}

// DeleteHrPayrollEditPayslipLinesWizards deletes existing hr.payroll.edit.payslip.lines.wizard records.
func (c *Client) DeleteHrPayrollEditPayslipLinesWizards(ids []int64) error {
	return c.Delete(HrPayrollEditPayslipLinesWizardModel, ids)
}

// GetHrPayrollEditPayslipLinesWizard gets hr.payroll.edit.payslip.lines.wizard existing record.
func (c *Client) GetHrPayrollEditPayslipLinesWizard(id int64) (*HrPayrollEditPayslipLinesWizard, error) {
	hpeplws, err := c.GetHrPayrollEditPayslipLinesWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpeplws != nil && len(*hpeplws) > 0 {
		return &((*hpeplws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payroll.edit.payslip.lines.wizard not found", id)
}

// GetHrPayrollEditPayslipLinesWizards gets hr.payroll.edit.payslip.lines.wizard existing records.
func (c *Client) GetHrPayrollEditPayslipLinesWizards(ids []int64) (*HrPayrollEditPayslipLinesWizards, error) {
	hpeplws := &HrPayrollEditPayslipLinesWizards{}
	if err := c.Read(HrPayrollEditPayslipLinesWizardModel, ids, nil, hpeplws); err != nil {
		return nil, err
	}
	return hpeplws, nil
}

// FindHrPayrollEditPayslipLinesWizard finds hr.payroll.edit.payslip.lines.wizard record by querying it with criteria.
func (c *Client) FindHrPayrollEditPayslipLinesWizard(criteria *Criteria) (*HrPayrollEditPayslipLinesWizard, error) {
	hpeplws := &HrPayrollEditPayslipLinesWizards{}
	if err := c.SearchRead(HrPayrollEditPayslipLinesWizardModel, criteria, NewOptions().Limit(1), hpeplws); err != nil {
		return nil, err
	}
	if hpeplws != nil && len(*hpeplws) > 0 {
		return &((*hpeplws)[0]), nil
	}
	return nil, fmt.Errorf("hr.payroll.edit.payslip.lines.wizard was not found with criteria %v", criteria)
}

// FindHrPayrollEditPayslipLinesWizards finds hr.payroll.edit.payslip.lines.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollEditPayslipLinesWizards(criteria *Criteria, options *Options) (*HrPayrollEditPayslipLinesWizards, error) {
	hpeplws := &HrPayrollEditPayslipLinesWizards{}
	if err := c.SearchRead(HrPayrollEditPayslipLinesWizardModel, criteria, options, hpeplws); err != nil {
		return nil, err
	}
	return hpeplws, nil
}

// FindHrPayrollEditPayslipLinesWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollEditPayslipLinesWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayrollEditPayslipLinesWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayrollEditPayslipLinesWizardId finds record id by querying it with criteria.
func (c *Client) FindHrPayrollEditPayslipLinesWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayrollEditPayslipLinesWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payroll.edit.payslip.lines.wizard was not found with criteria %v and options %v", criteria, options)
}
