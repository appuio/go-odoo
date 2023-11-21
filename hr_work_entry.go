package odoo

import (
	"fmt"
)

// HrWorkEntry represents hr.work.entry model.
type HrWorkEntry struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Active          *Bool      `xmlrpc:"active,omptempty"`
	AttendanceId    *Many2One  `xmlrpc:"attendance_id,omptempty"`
	Color           *Int       `xmlrpc:"color,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	Conflict        *Bool      `xmlrpc:"conflict,omptempty"`
	ContractId      *Many2One  `xmlrpc:"contract_id,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateStart       *Time      `xmlrpc:"date_start,omptempty"`
	DateStop        *Time      `xmlrpc:"date_stop,omptempty"`
	DepartmentId    *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Duration        *Float     `xmlrpc:"duration,omptempty"`
	EmployeeId      *Many2One  `xmlrpc:"employee_id,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	LeaveId         *Many2One  `xmlrpc:"leave_id,omptempty"`
	LeaveState      *Selection `xmlrpc:"leave_state,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	PlanningSlotId  *Many2One  `xmlrpc:"planning_slot_id,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
	WorkEntrySource *Selection `xmlrpc:"work_entry_source,omptempty"`
	WorkEntryTypeId *Many2One  `xmlrpc:"work_entry_type_id,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrWorkEntrys represents array of hr.work.entry model.
type HrWorkEntrys []HrWorkEntry

// HrWorkEntryModel is the odoo model name.
const HrWorkEntryModel = "hr.work.entry"

// Many2One convert HrWorkEntry to *Many2One.
func (hwe *HrWorkEntry) Many2One() *Many2One {
	return NewMany2One(hwe.Id.Get(), "")
}

// CreateHrWorkEntry creates a new hr.work.entry model and returns its id.
func (c *Client) CreateHrWorkEntry(hwe *HrWorkEntry) (int64, error) {
	ids, err := c.CreateHrWorkEntrys([]*HrWorkEntry{hwe})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrWorkEntry creates a new hr.work.entry model and returns its id.
func (c *Client) CreateHrWorkEntrys(hwes []*HrWorkEntry) ([]int64, error) {
	var vv []interface{}
	for _, v := range hwes {
		vv = append(vv, v)
	}
	return c.Create(HrWorkEntryModel, vv)
}

// UpdateHrWorkEntry updates an existing hr.work.entry record.
func (c *Client) UpdateHrWorkEntry(hwe *HrWorkEntry) error {
	return c.UpdateHrWorkEntrys([]int64{hwe.Id.Get()}, hwe)
}

// UpdateHrWorkEntrys updates existing hr.work.entry records.
// All records (represented by ids) will be updated by hwe values.
func (c *Client) UpdateHrWorkEntrys(ids []int64, hwe *HrWorkEntry) error {
	return c.Update(HrWorkEntryModel, ids, hwe)
}

// DeleteHrWorkEntry deletes an existing hr.work.entry record.
func (c *Client) DeleteHrWorkEntry(id int64) error {
	return c.DeleteHrWorkEntrys([]int64{id})
}

// DeleteHrWorkEntrys deletes existing hr.work.entry records.
func (c *Client) DeleteHrWorkEntrys(ids []int64) error {
	return c.Delete(HrWorkEntryModel, ids)
}

// GetHrWorkEntry gets hr.work.entry existing record.
func (c *Client) GetHrWorkEntry(id int64) (*HrWorkEntry, error) {
	hwes, err := c.GetHrWorkEntrys([]int64{id})
	if err != nil {
		return nil, err
	}
	if hwes != nil && len(*hwes) > 0 {
		return &((*hwes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.work.entry not found", id)
}

// GetHrWorkEntrys gets hr.work.entry existing records.
func (c *Client) GetHrWorkEntrys(ids []int64) (*HrWorkEntrys, error) {
	hwes := &HrWorkEntrys{}
	if err := c.Read(HrWorkEntryModel, ids, nil, hwes); err != nil {
		return nil, err
	}
	return hwes, nil
}

// FindHrWorkEntry finds hr.work.entry record by querying it with criteria.
func (c *Client) FindHrWorkEntry(criteria *Criteria) (*HrWorkEntry, error) {
	hwes := &HrWorkEntrys{}
	if err := c.SearchRead(HrWorkEntryModel, criteria, NewOptions().Limit(1), hwes); err != nil {
		return nil, err
	}
	if hwes != nil && len(*hwes) > 0 {
		return &((*hwes)[0]), nil
	}
	return nil, fmt.Errorf("hr.work.entry was not found with criteria %v", criteria)
}

// FindHrWorkEntrys finds hr.work.entry records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntrys(criteria *Criteria, options *Options) (*HrWorkEntrys, error) {
	hwes := &HrWorkEntrys{}
	if err := c.SearchRead(HrWorkEntryModel, criteria, options, hwes); err != nil {
		return nil, err
	}
	return hwes, nil
}

// FindHrWorkEntryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrWorkEntryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrWorkEntryId finds record id by querying it with criteria.
func (c *Client) FindHrWorkEntryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrWorkEntryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.work.entry was not found with criteria %v and options %v", criteria, options)
}
