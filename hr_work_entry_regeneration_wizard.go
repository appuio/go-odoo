package odoo

import (
	"fmt"
)

// HrWorkEntryRegenerationWizard represents hr.work.entry.regeneration.wizard model.
type HrWorkEntryRegenerationWizard struct {
	LastUpdate                   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omptempty"`
	DateFrom                     *Time     `xmlrpc:"date_from,omptempty"`
	DateTo                       *Time     `xmlrpc:"date_to,omptempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omptempty"`
	EarliestAvailableDate        *Time     `xmlrpc:"earliest_available_date,omptempty"`
	EarliestAvailableDateMessage *String   `xmlrpc:"earliest_available_date_message,omptempty"`
	EmployeeIds                  *Relation `xmlrpc:"employee_ids,omptempty"`
	Id                           *Int      `xmlrpc:"id,omptempty"`
	LatestAvailableDate          *Time     `xmlrpc:"latest_available_date,omptempty"`
	LatestAvailableDateMessage   *String   `xmlrpc:"latest_available_date_message,omptempty"`
	SearchCriteriaCompleted      *Bool     `xmlrpc:"search_criteria_completed,omptempty"`
	Valid                        *Bool     `xmlrpc:"valid,omptempty"`
	ValidatedWorkEntryIds        *Relation `xmlrpc:"validated_work_entry_ids,omptempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrWorkEntryRegenerationWizards represents array of hr.work.entry.regeneration.wizard model.
type HrWorkEntryRegenerationWizards []HrWorkEntryRegenerationWizard

// HrWorkEntryRegenerationWizardModel is the odoo model name.
const HrWorkEntryRegenerationWizardModel = "hr.work.entry.regeneration.wizard"

// Many2One convert HrWorkEntryRegenerationWizard to *Many2One.
func (hwerw *HrWorkEntryRegenerationWizard) Many2One() *Many2One {
	return NewMany2One(hwerw.Id.Get(), "")
}

// CreateHrWorkEntryRegenerationWizard creates a new hr.work.entry.regeneration.wizard model and returns its id.
func (c *Client) CreateHrWorkEntryRegenerationWizard(hwerw *HrWorkEntryRegenerationWizard) (int64, error) {
	ids, err := c.CreateHrWorkEntryRegenerationWizards([]*HrWorkEntryRegenerationWizard{hwerw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrWorkEntryRegenerationWizard creates a new hr.work.entry.regeneration.wizard model and returns its id.
func (c *Client) CreateHrWorkEntryRegenerationWizards(hwerws []*HrWorkEntryRegenerationWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hwerws {
		vv = append(vv, v)
	}
	return c.Create(HrWorkEntryRegenerationWizardModel, vv)
}

// UpdateHrWorkEntryRegenerationWizard updates an existing hr.work.entry.regeneration.wizard record.
func (c *Client) UpdateHrWorkEntryRegenerationWizard(hwerw *HrWorkEntryRegenerationWizard) error {
	return c.UpdateHrWorkEntryRegenerationWizards([]int64{hwerw.Id.Get()}, hwerw)
}

// UpdateHrWorkEntryRegenerationWizards updates existing hr.work.entry.regeneration.wizard records.
// All records (represented by ids) will be updated by hwerw values.
func (c *Client) UpdateHrWorkEntryRegenerationWizards(ids []int64, hwerw *HrWorkEntryRegenerationWizard) error {
	return c.Update(HrWorkEntryRegenerationWizardModel, ids, hwerw)
}

// DeleteHrWorkEntryRegenerationWizard deletes an existing hr.work.entry.regeneration.wizard record.
func (c *Client) DeleteHrWorkEntryRegenerationWizard(id int64) error {
	return c.DeleteHrWorkEntryRegenerationWizards([]int64{id})
}

// DeleteHrWorkEntryRegenerationWizards deletes existing hr.work.entry.regeneration.wizard records.
func (c *Client) DeleteHrWorkEntryRegenerationWizards(ids []int64) error {
	return c.Delete(HrWorkEntryRegenerationWizardModel, ids)
}

// GetHrWorkEntryRegenerationWizard gets hr.work.entry.regeneration.wizard existing record.
func (c *Client) GetHrWorkEntryRegenerationWizard(id int64) (*HrWorkEntryRegenerationWizard, error) {
	hwerws, err := c.GetHrWorkEntryRegenerationWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hwerws != nil && len(*hwerws) > 0 {
		return &((*hwerws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.work.entry.regeneration.wizard not found", id)
}

// GetHrWorkEntryRegenerationWizards gets hr.work.entry.regeneration.wizard existing records.
func (c *Client) GetHrWorkEntryRegenerationWizards(ids []int64) (*HrWorkEntryRegenerationWizards, error) {
	hwerws := &HrWorkEntryRegenerationWizards{}
	if err := c.Read(HrWorkEntryRegenerationWizardModel, ids, nil, hwerws); err != nil {
		return nil, err
	}
	return hwerws, nil
}

// FindHrWorkEntryRegenerationWizard finds hr.work.entry.regeneration.wizard record by querying it with criteria.
func (c *Client) FindHrWorkEntryRegenerationWizard(criteria *Criteria) (*HrWorkEntryRegenerationWizard, error) {
	hwerws := &HrWorkEntryRegenerationWizards{}
	if err := c.SearchRead(HrWorkEntryRegenerationWizardModel, criteria, NewOptions().Limit(1), hwerws); err != nil {
		return nil, err
	}
	if hwerws != nil && len(*hwerws) > 0 {
		return &((*hwerws)[0]), nil
	}
	return nil, fmt.Errorf("hr.work.entry.regeneration.wizard was not found with criteria %v", criteria)
}

// FindHrWorkEntryRegenerationWizards finds hr.work.entry.regeneration.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntryRegenerationWizards(criteria *Criteria, options *Options) (*HrWorkEntryRegenerationWizards, error) {
	hwerws := &HrWorkEntryRegenerationWizards{}
	if err := c.SearchRead(HrWorkEntryRegenerationWizardModel, criteria, options, hwerws); err != nil {
		return nil, err
	}
	return hwerws, nil
}

// FindHrWorkEntryRegenerationWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkEntryRegenerationWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrWorkEntryRegenerationWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrWorkEntryRegenerationWizardId finds record id by querying it with criteria.
func (c *Client) FindHrWorkEntryRegenerationWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrWorkEntryRegenerationWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.work.entry.regeneration.wizard was not found with criteria %v and options %v", criteria, options)
}
