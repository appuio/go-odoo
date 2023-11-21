package odoo

import (
	"fmt"
)

// HrTimesheetMergeWizard represents hr_timesheet.merge.wizard model.
type HrTimesheetMergeWizard struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	Date          *Time     `xmlrpc:"date,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omptempty"`
	EncodingUomId *Many2One `xmlrpc:"encoding_uom_id,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	ProjectId     *Many2One `xmlrpc:"project_id,omptempty"`
	TaskId        *Many2One `xmlrpc:"task_id,omptempty"`
	TimesheetIds  *Relation `xmlrpc:"timesheet_ids,omptempty"`
	UnitAmount    *Float    `xmlrpc:"unit_amount,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrTimesheetMergeWizards represents array of hr_timesheet.merge.wizard model.
type HrTimesheetMergeWizards []HrTimesheetMergeWizard

// HrTimesheetMergeWizardModel is the odoo model name.
const HrTimesheetMergeWizardModel = "hr_timesheet.merge.wizard"

// Many2One convert HrTimesheetMergeWizard to *Many2One.
func (hmw *HrTimesheetMergeWizard) Many2One() *Many2One {
	return NewMany2One(hmw.Id.Get(), "")
}

// CreateHrTimesheetMergeWizard creates a new hr_timesheet.merge.wizard model and returns its id.
func (c *Client) CreateHrTimesheetMergeWizard(hmw *HrTimesheetMergeWizard) (int64, error) {
	ids, err := c.CreateHrTimesheetMergeWizards([]*HrTimesheetMergeWizard{hmw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrTimesheetMergeWizard creates a new hr_timesheet.merge.wizard model and returns its id.
func (c *Client) CreateHrTimesheetMergeWizards(hmws []*HrTimesheetMergeWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hmws {
		vv = append(vv, v)
	}
	return c.Create(HrTimesheetMergeWizardModel, vv)
}

// UpdateHrTimesheetMergeWizard updates an existing hr_timesheet.merge.wizard record.
func (c *Client) UpdateHrTimesheetMergeWizard(hmw *HrTimesheetMergeWizard) error {
	return c.UpdateHrTimesheetMergeWizards([]int64{hmw.Id.Get()}, hmw)
}

// UpdateHrTimesheetMergeWizards updates existing hr_timesheet.merge.wizard records.
// All records (represented by ids) will be updated by hmw values.
func (c *Client) UpdateHrTimesheetMergeWizards(ids []int64, hmw *HrTimesheetMergeWizard) error {
	return c.Update(HrTimesheetMergeWizardModel, ids, hmw)
}

// DeleteHrTimesheetMergeWizard deletes an existing hr_timesheet.merge.wizard record.
func (c *Client) DeleteHrTimesheetMergeWizard(id int64) error {
	return c.DeleteHrTimesheetMergeWizards([]int64{id})
}

// DeleteHrTimesheetMergeWizards deletes existing hr_timesheet.merge.wizard records.
func (c *Client) DeleteHrTimesheetMergeWizards(ids []int64) error {
	return c.Delete(HrTimesheetMergeWizardModel, ids)
}

// GetHrTimesheetMergeWizard gets hr_timesheet.merge.wizard existing record.
func (c *Client) GetHrTimesheetMergeWizard(id int64) (*HrTimesheetMergeWizard, error) {
	hmws, err := c.GetHrTimesheetMergeWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hmws != nil && len(*hmws) > 0 {
		return &((*hmws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr_timesheet.merge.wizard not found", id)
}

// GetHrTimesheetMergeWizards gets hr_timesheet.merge.wizard existing records.
func (c *Client) GetHrTimesheetMergeWizards(ids []int64) (*HrTimesheetMergeWizards, error) {
	hmws := &HrTimesheetMergeWizards{}
	if err := c.Read(HrTimesheetMergeWizardModel, ids, nil, hmws); err != nil {
		return nil, err
	}
	return hmws, nil
}

// FindHrTimesheetMergeWizard finds hr_timesheet.merge.wizard record by querying it with criteria.
func (c *Client) FindHrTimesheetMergeWizard(criteria *Criteria) (*HrTimesheetMergeWizard, error) {
	hmws := &HrTimesheetMergeWizards{}
	if err := c.SearchRead(HrTimesheetMergeWizardModel, criteria, NewOptions().Limit(1), hmws); err != nil {
		return nil, err
	}
	if hmws != nil && len(*hmws) > 0 {
		return &((*hmws)[0]), nil
	}
	return nil, fmt.Errorf("hr_timesheet.merge.wizard was not found with criteria %v", criteria)
}

// FindHrTimesheetMergeWizards finds hr_timesheet.merge.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetMergeWizards(criteria *Criteria, options *Options) (*HrTimesheetMergeWizards, error) {
	hmws := &HrTimesheetMergeWizards{}
	if err := c.SearchRead(HrTimesheetMergeWizardModel, criteria, options, hmws); err != nil {
		return nil, err
	}
	return hmws, nil
}

// FindHrTimesheetMergeWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrTimesheetMergeWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrTimesheetMergeWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrTimesheetMergeWizardId finds record id by querying it with criteria.
func (c *Client) FindHrTimesheetMergeWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrTimesheetMergeWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr_timesheet.merge.wizard was not found with criteria %v and options %v", criteria, options)
}
