package odoo

import (
	"fmt"
)

// ManageMeteredUsageWizard represents manage.metered.usage.wizard model.
type ManageMeteredUsageWizard struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	AnalyticLineMeteredIds *Relation `xmlrpc:"analytic_line_metered_ids,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	OrderId                *Many2One `xmlrpc:"order_id,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ManageMeteredUsageWizards represents array of manage.metered.usage.wizard model.
type ManageMeteredUsageWizards []ManageMeteredUsageWizard

// ManageMeteredUsageWizardModel is the odoo model name.
const ManageMeteredUsageWizardModel = "manage.metered.usage.wizard"

// Many2One convert ManageMeteredUsageWizard to *Many2One.
func (mmuw *ManageMeteredUsageWizard) Many2One() *Many2One {
	return NewMany2One(mmuw.Id.Get(), "")
}

// CreateManageMeteredUsageWizard creates a new manage.metered.usage.wizard model and returns its id.
func (c *Client) CreateManageMeteredUsageWizard(mmuw *ManageMeteredUsageWizard) (int64, error) {
	ids, err := c.CreateManageMeteredUsageWizards([]*ManageMeteredUsageWizard{mmuw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateManageMeteredUsageWizard creates a new manage.metered.usage.wizard model and returns its id.
func (c *Client) CreateManageMeteredUsageWizards(mmuws []*ManageMeteredUsageWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmuws {
		vv = append(vv, v)
	}
	return c.Create(ManageMeteredUsageWizardModel, vv)
}

// UpdateManageMeteredUsageWizard updates an existing manage.metered.usage.wizard record.
func (c *Client) UpdateManageMeteredUsageWizard(mmuw *ManageMeteredUsageWizard) error {
	return c.UpdateManageMeteredUsageWizards([]int64{mmuw.Id.Get()}, mmuw)
}

// UpdateManageMeteredUsageWizards updates existing manage.metered.usage.wizard records.
// All records (represented by ids) will be updated by mmuw values.
func (c *Client) UpdateManageMeteredUsageWizards(ids []int64, mmuw *ManageMeteredUsageWizard) error {
	return c.Update(ManageMeteredUsageWizardModel, ids, mmuw)
}

// DeleteManageMeteredUsageWizard deletes an existing manage.metered.usage.wizard record.
func (c *Client) DeleteManageMeteredUsageWizard(id int64) error {
	return c.DeleteManageMeteredUsageWizards([]int64{id})
}

// DeleteManageMeteredUsageWizards deletes existing manage.metered.usage.wizard records.
func (c *Client) DeleteManageMeteredUsageWizards(ids []int64) error {
	return c.Delete(ManageMeteredUsageWizardModel, ids)
}

// GetManageMeteredUsageWizard gets manage.metered.usage.wizard existing record.
func (c *Client) GetManageMeteredUsageWizard(id int64) (*ManageMeteredUsageWizard, error) {
	mmuws, err := c.GetManageMeteredUsageWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if mmuws != nil && len(*mmuws) > 0 {
		return &((*mmuws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of manage.metered.usage.wizard not found", id)
}

// GetManageMeteredUsageWizards gets manage.metered.usage.wizard existing records.
func (c *Client) GetManageMeteredUsageWizards(ids []int64) (*ManageMeteredUsageWizards, error) {
	mmuws := &ManageMeteredUsageWizards{}
	if err := c.Read(ManageMeteredUsageWizardModel, ids, nil, mmuws); err != nil {
		return nil, err
	}
	return mmuws, nil
}

// FindManageMeteredUsageWizard finds manage.metered.usage.wizard record by querying it with criteria.
func (c *Client) FindManageMeteredUsageWizard(criteria *Criteria) (*ManageMeteredUsageWizard, error) {
	mmuws := &ManageMeteredUsageWizards{}
	if err := c.SearchRead(ManageMeteredUsageWizardModel, criteria, NewOptions().Limit(1), mmuws); err != nil {
		return nil, err
	}
	if mmuws != nil && len(*mmuws) > 0 {
		return &((*mmuws)[0]), nil
	}
	return nil, fmt.Errorf("manage.metered.usage.wizard was not found with criteria %v", criteria)
}

// FindManageMeteredUsageWizards finds manage.metered.usage.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindManageMeteredUsageWizards(criteria *Criteria, options *Options) (*ManageMeteredUsageWizards, error) {
	mmuws := &ManageMeteredUsageWizards{}
	if err := c.SearchRead(ManageMeteredUsageWizardModel, criteria, options, mmuws); err != nil {
		return nil, err
	}
	return mmuws, nil
}

// FindManageMeteredUsageWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindManageMeteredUsageWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ManageMeteredUsageWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindManageMeteredUsageWizardId finds record id by querying it with criteria.
func (c *Client) FindManageMeteredUsageWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ManageMeteredUsageWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("manage.metered.usage.wizard was not found with criteria %v and options %v", criteria, options)
}
