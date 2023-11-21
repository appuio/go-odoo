package odoo

import (
	"fmt"
)

// HrPayrollStructure represents hr.payroll.structure model.
type HrPayrollStructure struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	Code                   *String    `xmlrpc:"code,omptempty"`
	CountryId              *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	InputLineTypeIds       *Relation  `xmlrpc:"input_line_type_ids,omptempty"`
	JournalId              *Many2One  `xmlrpc:"journal_id,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	Note                   *String    `xmlrpc:"note,omptempty"`
	PayslipName            *String    `xmlrpc:"payslip_name,omptempty"`
	ReportId               *Many2One  `xmlrpc:"report_id,omptempty"`
	RuleIds                *Relation  `xmlrpc:"rule_ids,omptempty"`
	SchedulePay            *Selection `xmlrpc:"schedule_pay,omptempty"`
	TypeId                 *Many2One  `xmlrpc:"type_id,omptempty"`
	UnpaidWorkEntryTypeIds *Relation  `xmlrpc:"unpaid_work_entry_type_ids,omptempty"`
	UseWorkedDayLines      *Bool      `xmlrpc:"use_worked_day_lines,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrPayrollStructures represents array of hr.payroll.structure model.
type HrPayrollStructures []HrPayrollStructure

// HrPayrollStructureModel is the odoo model name.
const HrPayrollStructureModel = "hr.payroll.structure"

// Many2One convert HrPayrollStructure to *Many2One.
func (hps *HrPayrollStructure) Many2One() *Many2One {
	return NewMany2One(hps.Id.Get(), "")
}

// CreateHrPayrollStructure creates a new hr.payroll.structure model and returns its id.
func (c *Client) CreateHrPayrollStructure(hps *HrPayrollStructure) (int64, error) {
	ids, err := c.CreateHrPayrollStructures([]*HrPayrollStructure{hps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayrollStructure creates a new hr.payroll.structure model and returns its id.
func (c *Client) CreateHrPayrollStructures(hpss []*HrPayrollStructure) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpss {
		vv = append(vv, v)
	}
	return c.Create(HrPayrollStructureModel, vv)
}

// UpdateHrPayrollStructure updates an existing hr.payroll.structure record.
func (c *Client) UpdateHrPayrollStructure(hps *HrPayrollStructure) error {
	return c.UpdateHrPayrollStructures([]int64{hps.Id.Get()}, hps)
}

// UpdateHrPayrollStructures updates existing hr.payroll.structure records.
// All records (represented by ids) will be updated by hps values.
func (c *Client) UpdateHrPayrollStructures(ids []int64, hps *HrPayrollStructure) error {
	return c.Update(HrPayrollStructureModel, ids, hps)
}

// DeleteHrPayrollStructure deletes an existing hr.payroll.structure record.
func (c *Client) DeleteHrPayrollStructure(id int64) error {
	return c.DeleteHrPayrollStructures([]int64{id})
}

// DeleteHrPayrollStructures deletes existing hr.payroll.structure records.
func (c *Client) DeleteHrPayrollStructures(ids []int64) error {
	return c.Delete(HrPayrollStructureModel, ids)
}

// GetHrPayrollStructure gets hr.payroll.structure existing record.
func (c *Client) GetHrPayrollStructure(id int64) (*HrPayrollStructure, error) {
	hpss, err := c.GetHrPayrollStructures([]int64{id})
	if err != nil {
		return nil, err
	}
	if hpss != nil && len(*hpss) > 0 {
		return &((*hpss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payroll.structure not found", id)
}

// GetHrPayrollStructures gets hr.payroll.structure existing records.
func (c *Client) GetHrPayrollStructures(ids []int64) (*HrPayrollStructures, error) {
	hpss := &HrPayrollStructures{}
	if err := c.Read(HrPayrollStructureModel, ids, nil, hpss); err != nil {
		return nil, err
	}
	return hpss, nil
}

// FindHrPayrollStructure finds hr.payroll.structure record by querying it with criteria.
func (c *Client) FindHrPayrollStructure(criteria *Criteria) (*HrPayrollStructure, error) {
	hpss := &HrPayrollStructures{}
	if err := c.SearchRead(HrPayrollStructureModel, criteria, NewOptions().Limit(1), hpss); err != nil {
		return nil, err
	}
	if hpss != nil && len(*hpss) > 0 {
		return &((*hpss)[0]), nil
	}
	return nil, fmt.Errorf("hr.payroll.structure was not found with criteria %v", criteria)
}

// FindHrPayrollStructures finds hr.payroll.structure records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollStructures(criteria *Criteria, options *Options) (*HrPayrollStructures, error) {
	hpss := &HrPayrollStructures{}
	if err := c.SearchRead(HrPayrollStructureModel, criteria, options, hpss); err != nil {
		return nil, err
	}
	return hpss, nil
}

// FindHrPayrollStructureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollStructureIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayrollStructureModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayrollStructureId finds record id by querying it with criteria.
func (c *Client) FindHrPayrollStructureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayrollStructureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payroll.structure was not found with criteria %v and options %v", criteria, options)
}
