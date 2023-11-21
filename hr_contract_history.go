package odoo

import (
	"fmt"
)

// HrContractHistory represents hr.contract.history model.
type HrContractHistory struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	ActiveEmployee        *Bool      `xmlrpc:"active_employee,omptempty"`
	ActivityState         *Selection `xmlrpc:"activity_state,omptempty"`
	AnalyticAccountId     *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CompanyCountryId      *Many2One  `xmlrpc:"company_country_id,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	ContractCount         *Int       `xmlrpc:"contract_count,omptempty"`
	ContractId            *Many2One  `xmlrpc:"contract_id,omptempty"`
	ContractIds           *Relation  `xmlrpc:"contract_ids,omptempty"`
	ContractTypeId        *Many2One  `xmlrpc:"contract_type_id,omptempty"`
	CountryCode           *String    `xmlrpc:"country_code,omptempty"`
	CurrencyId            *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateEnd               *Time      `xmlrpc:"date_end,omptempty"`
	DateHired             *Time      `xmlrpc:"date_hired,omptempty"`
	DateStart             *Time      `xmlrpc:"date_start,omptempty"`
	DepartmentId          *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId            *Many2One  `xmlrpc:"employee_id,omptempty"`
	HrResponsibleId       *Many2One  `xmlrpc:"hr_responsible_id,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	IsUnderContract       *Bool      `xmlrpc:"is_under_contract,omptempty"`
	JobId                 *Many2One  `xmlrpc:"job_id,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	PayslipsCount         *Int       `xmlrpc:"payslips_count,omptempty"`
	ResourceCalendarId    *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	SalaryAttachmentCount *String    `xmlrpc:"salary_attachment_count,omptempty"`
	StandardCalendarId    *Many2One  `xmlrpc:"standard_calendar_id,omptempty"`
	State                 *Selection `xmlrpc:"state,omptempty"`
	StructureTypeId       *Many2One  `xmlrpc:"structure_type_id,omptempty"`
	TimeCredit            *Bool      `xmlrpc:"time_credit,omptempty"`
	UnderContractState    *Selection `xmlrpc:"under_contract_state,omptempty"`
	Wage                  *Float     `xmlrpc:"wage,omptempty"`
	WageType              *Selection `xmlrpc:"wage_type,omptempty"`
	WorkTimeRate          *Float     `xmlrpc:"work_time_rate,omptempty"`
}

// HrContractHistorys represents array of hr.contract.history model.
type HrContractHistorys []HrContractHistory

// HrContractHistoryModel is the odoo model name.
const HrContractHistoryModel = "hr.contract.history"

// Many2One convert HrContractHistory to *Many2One.
func (hch *HrContractHistory) Many2One() *Many2One {
	return NewMany2One(hch.Id.Get(), "")
}

// CreateHrContractHistory creates a new hr.contract.history model and returns its id.
func (c *Client) CreateHrContractHistory(hch *HrContractHistory) (int64, error) {
	ids, err := c.CreateHrContractHistorys([]*HrContractHistory{hch})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrContractHistory creates a new hr.contract.history model and returns its id.
func (c *Client) CreateHrContractHistorys(hchs []*HrContractHistory) ([]int64, error) {
	var vv []interface{}
	for _, v := range hchs {
		vv = append(vv, v)
	}
	return c.Create(HrContractHistoryModel, vv)
}

// UpdateHrContractHistory updates an existing hr.contract.history record.
func (c *Client) UpdateHrContractHistory(hch *HrContractHistory) error {
	return c.UpdateHrContractHistorys([]int64{hch.Id.Get()}, hch)
}

// UpdateHrContractHistorys updates existing hr.contract.history records.
// All records (represented by ids) will be updated by hch values.
func (c *Client) UpdateHrContractHistorys(ids []int64, hch *HrContractHistory) error {
	return c.Update(HrContractHistoryModel, ids, hch)
}

// DeleteHrContractHistory deletes an existing hr.contract.history record.
func (c *Client) DeleteHrContractHistory(id int64) error {
	return c.DeleteHrContractHistorys([]int64{id})
}

// DeleteHrContractHistorys deletes existing hr.contract.history records.
func (c *Client) DeleteHrContractHistorys(ids []int64) error {
	return c.Delete(HrContractHistoryModel, ids)
}

// GetHrContractHistory gets hr.contract.history existing record.
func (c *Client) GetHrContractHistory(id int64) (*HrContractHistory, error) {
	hchs, err := c.GetHrContractHistorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if hchs != nil && len(*hchs) > 0 {
		return &((*hchs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.contract.history not found", id)
}

// GetHrContractHistorys gets hr.contract.history existing records.
func (c *Client) GetHrContractHistorys(ids []int64) (*HrContractHistorys, error) {
	hchs := &HrContractHistorys{}
	if err := c.Read(HrContractHistoryModel, ids, nil, hchs); err != nil {
		return nil, err
	}
	return hchs, nil
}

// FindHrContractHistory finds hr.contract.history record by querying it with criteria.
func (c *Client) FindHrContractHistory(criteria *Criteria) (*HrContractHistory, error) {
	hchs := &HrContractHistorys{}
	if err := c.SearchRead(HrContractHistoryModel, criteria, NewOptions().Limit(1), hchs); err != nil {
		return nil, err
	}
	if hchs != nil && len(*hchs) > 0 {
		return &((*hchs)[0]), nil
	}
	return nil, fmt.Errorf("hr.contract.history was not found with criteria %v", criteria)
}

// FindHrContractHistorys finds hr.contract.history records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractHistorys(criteria *Criteria, options *Options) (*HrContractHistorys, error) {
	hchs := &HrContractHistorys{}
	if err := c.SearchRead(HrContractHistoryModel, criteria, options, hchs); err != nil {
		return nil, err
	}
	return hchs, nil
}

// FindHrContractHistoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractHistoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrContractHistoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrContractHistoryId finds record id by querying it with criteria.
func (c *Client) FindHrContractHistoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractHistoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.contract.history was not found with criteria %v and options %v", criteria, options)
}
