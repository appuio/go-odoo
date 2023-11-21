package odoo

import (
	"fmt"
)

// HrPayslipSepaWizard represents hr.payslip.sepa.wizard model.
type HrPayslipSepaWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JournalId   *Many2One `xmlrpc:"journal_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipSepaWizards represents array of hr.payslip.sepa.wizard model.
type HrPayslipSepaWizards []HrPayslipSepaWizard

// HrPayslipSepaWizardModel is the odoo model name.
const HrPayslipSepaWizardModel = "hr.payslip.sepa.wizard"

// Many2One convert HrPayslipSepaWizard to *Many2One.
func (hpsw *HrPayslipSepaWizard) Many2One() *Many2One {
	return NewMany2One(hpsw.Id.Get(), "")
}

// CreateHrPayslipSepaWizard creates a new hr.payslip.sepa.wizard model and returns its id.
func (c *Client) CreateHrPayslipSepaWizard(hpsw *HrPayslipSepaWizard) (int64, error) {
	ids, err := c.CreateHrPayslipSepaWizards([]*HrPayslipSepaWizard{hpsw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipSepaWizard creates a new hr.payslip.sepa.wizard model and returns its id.
func (c *Client) CreateHrPayslipSepaWizards(hpsws []*HrPayslipSepaWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpsws {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipSepaWizardModel, vv)
}

// UpdateHrPayslipSepaWizard updates an existing hr.payslip.sepa.wizard record.
func (c *Client) UpdateHrPayslipSepaWizard(hpsw *HrPayslipSepaWizard) error {
	return c.UpdateHrPayslipSepaWizards([]int64{hpsw.Id.Get()}, hpsw)
}

// UpdateHrPayslipSepaWizards updates existing hr.payslip.sepa.wizard records.
// All records (represented by ids) will be updated by hpsw values.
func (c *Client) UpdateHrPayslipSepaWizards(ids []int64, hpsw *HrPayslipSepaWizard) error {
	return c.Update(HrPayslipSepaWizardModel, ids, hpsw)
}

// DeleteHrPayslipSepaWizard deletes an existing hr.payslip.sepa.wizard record.
func (c *Client) DeleteHrPayslipSepaWizard(id int64) error {
	return c.DeleteHrPayslipSepaWizards([]int64{id})
}

// DeleteHrPayslipSepaWizards deletes existing hr.payslip.sepa.wizard records.
func (c *Client) DeleteHrPayslipSepaWizards(ids []int64) error {
	return c.Delete(HrPayslipSepaWizardModel, ids)
}

// GetHrPayslipSepaWizard gets hr.payslip.sepa.wizard existing record.
func (c *Client) GetHrPayslipSepaWizard(id int64) (*HrPayslipSepaWizard, error) {
	hpsws, err := c.GetHrPayslipSepaWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpsws != nil && len(*hpsws) > 0 {
		return &((*hpsws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.sepa.wizard not found", id)
}

// GetHrPayslipSepaWizards gets hr.payslip.sepa.wizard existing records.
func (c *Client) GetHrPayslipSepaWizards(ids []int64) (*HrPayslipSepaWizards, error) {
	hpsws := &HrPayslipSepaWizards{}
	if err := c.Read(HrPayslipSepaWizardModel, ids, nil, hpsws); err != nil {
		return nil, err
	}
	return hpsws, nil
}

// FindHrPayslipSepaWizard finds hr.payslip.sepa.wizard record by querying it with criteria.
func (c *Client) FindHrPayslipSepaWizard(criteria *Criteria) (*HrPayslipSepaWizard, error) {
	hpsws := &HrPayslipSepaWizards{}
	if err := c.SearchRead(HrPayslipSepaWizardModel, criteria, NewOptions().Limit(1), hpsws); err != nil {
		return nil, err
	}
	if hpsws != nil && len(*hpsws) > 0 {
		return &((*hpsws)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.sepa.wizard was not found with criteria %v", criteria)
}

// FindHrPayslipSepaWizards finds hr.payslip.sepa.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipSepaWizards(criteria *Criteria, options *Options) (*HrPayslipSepaWizards, error) {
	hpsws := &HrPayslipSepaWizards{}
	if err := c.SearchRead(HrPayslipSepaWizardModel, criteria, options, hpsws); err != nil {
		return nil, err
	}
	return hpsws, nil
}

// FindHrPayslipSepaWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipSepaWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipSepaWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipSepaWizardId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipSepaWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipSepaWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.sepa.wizard was not found with criteria %v and options %v", criteria, options)
}
