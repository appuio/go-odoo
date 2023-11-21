package odoo

import (
	"fmt"
)

// HrPayslipRunSepaWizard represents hr.payslip.run.sepa.wizard model.
type HrPayslipRunSepaWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FileName    *String   `xmlrpc:"file_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JournalId   *Many2One `xmlrpc:"journal_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipRunSepaWizards represents array of hr.payslip.run.sepa.wizard model.
type HrPayslipRunSepaWizards []HrPayslipRunSepaWizard

// HrPayslipRunSepaWizardModel is the odoo model name.
const HrPayslipRunSepaWizardModel = "hr.payslip.run.sepa.wizard"

// Many2One convert HrPayslipRunSepaWizard to *Many2One.
func (hprsw *HrPayslipRunSepaWizard) Many2One() *Many2One {
	return NewMany2One(hprsw.Id.Get(), "")
}

// CreateHrPayslipRunSepaWizard creates a new hr.payslip.run.sepa.wizard model and returns its id.
func (c *Client) CreateHrPayslipRunSepaWizard(hprsw *HrPayslipRunSepaWizard) (int64, error) {
	ids, err := c.CreateHrPayslipRunSepaWizards([]*HrPayslipRunSepaWizard{hprsw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipRunSepaWizard creates a new hr.payslip.run.sepa.wizard model and returns its id.
func (c *Client) CreateHrPayslipRunSepaWizards(hprsws []*HrPayslipRunSepaWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hprsws {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipRunSepaWizardModel, vv)
}

// UpdateHrPayslipRunSepaWizard updates an existing hr.payslip.run.sepa.wizard record.
func (c *Client) UpdateHrPayslipRunSepaWizard(hprsw *HrPayslipRunSepaWizard) error {
	return c.UpdateHrPayslipRunSepaWizards([]int64{hprsw.Id.Get()}, hprsw)
}

// UpdateHrPayslipRunSepaWizards updates existing hr.payslip.run.sepa.wizard records.
// All records (represented by ids) will be updated by hprsw values.
func (c *Client) UpdateHrPayslipRunSepaWizards(ids []int64, hprsw *HrPayslipRunSepaWizard) error {
	return c.Update(HrPayslipRunSepaWizardModel, ids, hprsw)
}

// DeleteHrPayslipRunSepaWizard deletes an existing hr.payslip.run.sepa.wizard record.
func (c *Client) DeleteHrPayslipRunSepaWizard(id int64) error {
	return c.DeleteHrPayslipRunSepaWizards([]int64{id})
}

// DeleteHrPayslipRunSepaWizards deletes existing hr.payslip.run.sepa.wizard records.
func (c *Client) DeleteHrPayslipRunSepaWizards(ids []int64) error {
	return c.Delete(HrPayslipRunSepaWizardModel, ids)
}

// GetHrPayslipRunSepaWizard gets hr.payslip.run.sepa.wizard existing record.
func (c *Client) GetHrPayslipRunSepaWizard(id int64) (*HrPayslipRunSepaWizard, error) {
	hprsws, err := c.GetHrPayslipRunSepaWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hprsws != nil && len(*hprsws) > 0 {
		return &((*hprsws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.run.sepa.wizard not found", id)
}

// GetHrPayslipRunSepaWizards gets hr.payslip.run.sepa.wizard existing records.
func (c *Client) GetHrPayslipRunSepaWizards(ids []int64) (*HrPayslipRunSepaWizards, error) {
	hprsws := &HrPayslipRunSepaWizards{}
	if err := c.Read(HrPayslipRunSepaWizardModel, ids, nil, hprsws); err != nil {
		return nil, err
	}
	return hprsws, nil
}

// FindHrPayslipRunSepaWizard finds hr.payslip.run.sepa.wizard record by querying it with criteria.
func (c *Client) FindHrPayslipRunSepaWizard(criteria *Criteria) (*HrPayslipRunSepaWizard, error) {
	hprsws := &HrPayslipRunSepaWizards{}
	if err := c.SearchRead(HrPayslipRunSepaWizardModel, criteria, NewOptions().Limit(1), hprsws); err != nil {
		return nil, err
	}
	if hprsws != nil && len(*hprsws) > 0 {
		return &((*hprsws)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.run.sepa.wizard was not found with criteria %v", criteria)
}

// FindHrPayslipRunSepaWizards finds hr.payslip.run.sepa.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipRunSepaWizards(criteria *Criteria, options *Options) (*HrPayslipRunSepaWizards, error) {
	hprsws := &HrPayslipRunSepaWizards{}
	if err := c.SearchRead(HrPayslipRunSepaWizardModel, criteria, options, hprsws); err != nil {
		return nil, err
	}
	return hprsws, nil
}

// FindHrPayslipRunSepaWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipRunSepaWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipRunSepaWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipRunSepaWizardId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipRunSepaWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipRunSepaWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.run.sepa.wizard was not found with criteria %v and options %v", criteria, options)
}
