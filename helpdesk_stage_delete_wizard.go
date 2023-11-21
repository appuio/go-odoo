package odoo

import (
	"fmt"
)

// HelpdeskStageDeleteWizard represents helpdesk.stage.delete.wizard model.
type HelpdeskStageDeleteWizard struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	StageIds     *Relation `xmlrpc:"stage_ids,omptempty"`
	StagesActive *Bool     `xmlrpc:"stages_active,omptempty"`
	TeamIds      *Relation `xmlrpc:"team_ids,omptempty"`
	TicketCount  *Int      `xmlrpc:"ticket_count,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskStageDeleteWizards represents array of helpdesk.stage.delete.wizard model.
type HelpdeskStageDeleteWizards []HelpdeskStageDeleteWizard

// HelpdeskStageDeleteWizardModel is the odoo model name.
const HelpdeskStageDeleteWizardModel = "helpdesk.stage.delete.wizard"

// Many2One convert HelpdeskStageDeleteWizard to *Many2One.
func (hsdw *HelpdeskStageDeleteWizard) Many2One() *Many2One {
	return NewMany2One(hsdw.Id.Get(), "")
}

// CreateHelpdeskStageDeleteWizard creates a new helpdesk.stage.delete.wizard model and returns its id.
func (c *Client) CreateHelpdeskStageDeleteWizard(hsdw *HelpdeskStageDeleteWizard) (int64, error) {
	ids, err := c.CreateHelpdeskStageDeleteWizards([]*HelpdeskStageDeleteWizard{hsdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskStageDeleteWizard creates a new helpdesk.stage.delete.wizard model and returns its id.
func (c *Client) CreateHelpdeskStageDeleteWizards(hsdws []*HelpdeskStageDeleteWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsdws {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskStageDeleteWizardModel, vv)
}

// UpdateHelpdeskStageDeleteWizard updates an existing helpdesk.stage.delete.wizard record.
func (c *Client) UpdateHelpdeskStageDeleteWizard(hsdw *HelpdeskStageDeleteWizard) error {
	return c.UpdateHelpdeskStageDeleteWizards([]int64{hsdw.Id.Get()}, hsdw)
}

// UpdateHelpdeskStageDeleteWizards updates existing helpdesk.stage.delete.wizard records.
// All records (represented by ids) will be updated by hsdw values.
func (c *Client) UpdateHelpdeskStageDeleteWizards(ids []int64, hsdw *HelpdeskStageDeleteWizard) error {
	return c.Update(HelpdeskStageDeleteWizardModel, ids, hsdw)
}

// DeleteHelpdeskStageDeleteWizard deletes an existing helpdesk.stage.delete.wizard record.
func (c *Client) DeleteHelpdeskStageDeleteWizard(id int64) error {
	return c.DeleteHelpdeskStageDeleteWizards([]int64{id})
}

// DeleteHelpdeskStageDeleteWizards deletes existing helpdesk.stage.delete.wizard records.
func (c *Client) DeleteHelpdeskStageDeleteWizards(ids []int64) error {
	return c.Delete(HelpdeskStageDeleteWizardModel, ids)
}

// GetHelpdeskStageDeleteWizard gets helpdesk.stage.delete.wizard existing record.
func (c *Client) GetHelpdeskStageDeleteWizard(id int64) (*HelpdeskStageDeleteWizard, error) {
	hsdws, err := c.GetHelpdeskStageDeleteWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsdws != nil && len(*hsdws) > 0 {
		return &((*hsdws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.stage.delete.wizard not found", id)
}

// GetHelpdeskStageDeleteWizards gets helpdesk.stage.delete.wizard existing records.
func (c *Client) GetHelpdeskStageDeleteWizards(ids []int64) (*HelpdeskStageDeleteWizards, error) {
	hsdws := &HelpdeskStageDeleteWizards{}
	if err := c.Read(HelpdeskStageDeleteWizardModel, ids, nil, hsdws); err != nil {
		return nil, err
	}
	return hsdws, nil
}

// FindHelpdeskStageDeleteWizard finds helpdesk.stage.delete.wizard record by querying it with criteria.
func (c *Client) FindHelpdeskStageDeleteWizard(criteria *Criteria) (*HelpdeskStageDeleteWizard, error) {
	hsdws := &HelpdeskStageDeleteWizards{}
	if err := c.SearchRead(HelpdeskStageDeleteWizardModel, criteria, NewOptions().Limit(1), hsdws); err != nil {
		return nil, err
	}
	if hsdws != nil && len(*hsdws) > 0 {
		return &((*hsdws)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.stage.delete.wizard was not found with criteria %v", criteria)
}

// FindHelpdeskStageDeleteWizards finds helpdesk.stage.delete.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskStageDeleteWizards(criteria *Criteria, options *Options) (*HelpdeskStageDeleteWizards, error) {
	hsdws := &HelpdeskStageDeleteWizards{}
	if err := c.SearchRead(HelpdeskStageDeleteWizardModel, criteria, options, hsdws); err != nil {
		return nil, err
	}
	return hsdws, nil
}

// FindHelpdeskStageDeleteWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskStageDeleteWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskStageDeleteWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskStageDeleteWizardId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskStageDeleteWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskStageDeleteWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.stage.delete.wizard was not found with criteria %v and options %v", criteria, options)
}
