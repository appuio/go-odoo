package odoo

import (
	"fmt"
)

// HrPayslipWorkedDays represents hr.payslip.worked_days model.
type HrPayslipWorkedDays struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Amount          *Float    `xmlrpc:"amount,omptempty"`
	Code            *String   `xmlrpc:"code,omptempty"`
	ContractId      *Many2One `xmlrpc:"contract_id,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	IsPaid          *Bool     `xmlrpc:"is_paid,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	NumberOfDays    *Float    `xmlrpc:"number_of_days,omptempty"`
	NumberOfHours   *Float    `xmlrpc:"number_of_hours,omptempty"`
	PayslipId       *Many2One `xmlrpc:"payslip_id,omptempty"`
	Sequence        *Int      `xmlrpc:"sequence,omptempty"`
	WorkEntryTypeId *Many2One `xmlrpc:"work_entry_type_id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipWorkedDayss represents array of hr.payslip.worked_days model.
type HrPayslipWorkedDayss []HrPayslipWorkedDays

// HrPayslipWorkedDaysModel is the odoo model name.
const HrPayslipWorkedDaysModel = "hr.payslip.worked_days"

// Many2One convert HrPayslipWorkedDays to *Many2One.
func (hpw *HrPayslipWorkedDays) Many2One() *Many2One {
	return NewMany2One(hpw.Id.Get(), "")
}

// CreateHrPayslipWorkedDays creates a new hr.payslip.worked_days model and returns its id.
func (c *Client) CreateHrPayslipWorkedDays(hpw *HrPayslipWorkedDays) (int64, error) {
	ids, err := c.CreateHrPayslipWorkedDayss([]*HrPayslipWorkedDays{hpw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipWorkedDays creates a new hr.payslip.worked_days model and returns its id.
func (c *Client) CreateHrPayslipWorkedDayss(hpws []*HrPayslipWorkedDays) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpws {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipWorkedDaysModel, vv)
}

// UpdateHrPayslipWorkedDays updates an existing hr.payslip.worked_days record.
func (c *Client) UpdateHrPayslipWorkedDays(hpw *HrPayslipWorkedDays) error {
	return c.UpdateHrPayslipWorkedDayss([]int64{hpw.Id.Get()}, hpw)
}

// UpdateHrPayslipWorkedDayss updates existing hr.payslip.worked_days records.
// All records (represented by ids) will be updated by hpw values.
func (c *Client) UpdateHrPayslipWorkedDayss(ids []int64, hpw *HrPayslipWorkedDays) error {
	return c.Update(HrPayslipWorkedDaysModel, ids, hpw)
}

// DeleteHrPayslipWorkedDays deletes an existing hr.payslip.worked_days record.
func (c *Client) DeleteHrPayslipWorkedDays(id int64) error {
	return c.DeleteHrPayslipWorkedDayss([]int64{id})
}

// DeleteHrPayslipWorkedDayss deletes existing hr.payslip.worked_days records.
func (c *Client) DeleteHrPayslipWorkedDayss(ids []int64) error {
	return c.Delete(HrPayslipWorkedDaysModel, ids)
}

// GetHrPayslipWorkedDays gets hr.payslip.worked_days existing record.
func (c *Client) GetHrPayslipWorkedDays(id int64) (*HrPayslipWorkedDays, error) {
	hpws, err := c.GetHrPayslipWorkedDayss([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpws != nil && len(*hpws) > 0 {
		return &((*hpws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.worked_days not found", id)
}

// GetHrPayslipWorkedDayss gets hr.payslip.worked_days existing records.
func (c *Client) GetHrPayslipWorkedDayss(ids []int64) (*HrPayslipWorkedDayss, error) {
	hpws := &HrPayslipWorkedDayss{}
	if err := c.Read(HrPayslipWorkedDaysModel, ids, nil, hpws); err != nil {
		return nil, err
	}
	return hpws, nil
}

// FindHrPayslipWorkedDays finds hr.payslip.worked_days record by querying it with criteria.
func (c *Client) FindHrPayslipWorkedDays(criteria *Criteria) (*HrPayslipWorkedDays, error) {
	hpws := &HrPayslipWorkedDayss{}
	if err := c.SearchRead(HrPayslipWorkedDaysModel, criteria, NewOptions().Limit(1), hpws); err != nil {
		return nil, err
	}
	if hpws != nil && len(*hpws) > 0 {
		return &((*hpws)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.worked_days was not found with criteria %v", criteria)
}

// FindHrPayslipWorkedDayss finds hr.payslip.worked_days records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipWorkedDayss(criteria *Criteria, options *Options) (*HrPayslipWorkedDayss, error) {
	hpws := &HrPayslipWorkedDayss{}
	if err := c.SearchRead(HrPayslipWorkedDaysModel, criteria, options, hpws); err != nil {
		return nil, err
	}
	return hpws, nil
}

// FindHrPayslipWorkedDaysIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipWorkedDaysIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipWorkedDaysModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipWorkedDaysId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipWorkedDaysId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipWorkedDaysModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.worked_days was not found with criteria %v and options %v", criteria, options)
}
