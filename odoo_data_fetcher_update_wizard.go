package odoo

import (
	"fmt"
)

// OdooDataFetcherUpdateWizard represents odoo_data_fetcher.update_wizard model.
type OdooDataFetcherUpdateWizard struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	EndDate             *Time     `xmlrpc:"end_date,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	StartDate           *Time     `xmlrpc:"start_date,omptempty"`
	TimePeriodActivated *Bool     `xmlrpc:"time_period_activated,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OdooDataFetcherUpdateWizards represents array of odoo_data_fetcher.update_wizard model.
type OdooDataFetcherUpdateWizards []OdooDataFetcherUpdateWizard

// OdooDataFetcherUpdateWizardModel is the odoo model name.
const OdooDataFetcherUpdateWizardModel = "odoo_data_fetcher.update_wizard"

// Many2One convert OdooDataFetcherUpdateWizard to *Many2One.
func (ou *OdooDataFetcherUpdateWizard) Many2One() *Many2One {
	return NewMany2One(ou.Id.Get(), "")
}

// CreateOdooDataFetcherUpdateWizard creates a new odoo_data_fetcher.update_wizard model and returns its id.
func (c *Client) CreateOdooDataFetcherUpdateWizard(ou *OdooDataFetcherUpdateWizard) (int64, error) {
	ids, err := c.CreateOdooDataFetcherUpdateWizards([]*OdooDataFetcherUpdateWizard{ou})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOdooDataFetcherUpdateWizard creates a new odoo_data_fetcher.update_wizard model and returns its id.
func (c *Client) CreateOdooDataFetcherUpdateWizards(ous []*OdooDataFetcherUpdateWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range ous {
		vv = append(vv, v)
	}
	return c.Create(OdooDataFetcherUpdateWizardModel, vv)
}

// UpdateOdooDataFetcherUpdateWizard updates an existing odoo_data_fetcher.update_wizard record.
func (c *Client) UpdateOdooDataFetcherUpdateWizard(ou *OdooDataFetcherUpdateWizard) error {
	return c.UpdateOdooDataFetcherUpdateWizards([]int64{ou.Id.Get()}, ou)
}

// UpdateOdooDataFetcherUpdateWizards updates existing odoo_data_fetcher.update_wizard records.
// All records (represented by ids) will be updated by ou values.
func (c *Client) UpdateOdooDataFetcherUpdateWizards(ids []int64, ou *OdooDataFetcherUpdateWizard) error {
	return c.Update(OdooDataFetcherUpdateWizardModel, ids, ou)
}

// DeleteOdooDataFetcherUpdateWizard deletes an existing odoo_data_fetcher.update_wizard record.
func (c *Client) DeleteOdooDataFetcherUpdateWizard(id int64) error {
	return c.DeleteOdooDataFetcherUpdateWizards([]int64{id})
}

// DeleteOdooDataFetcherUpdateWizards deletes existing odoo_data_fetcher.update_wizard records.
func (c *Client) DeleteOdooDataFetcherUpdateWizards(ids []int64) error {
	return c.Delete(OdooDataFetcherUpdateWizardModel, ids)
}

// GetOdooDataFetcherUpdateWizard gets odoo_data_fetcher.update_wizard existing record.
func (c *Client) GetOdooDataFetcherUpdateWizard(id int64) (*OdooDataFetcherUpdateWizard, error) {
	ous, err := c.GetOdooDataFetcherUpdateWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if ous != nil && len(*ous) > 0 {
		return &((*ous)[0]), nil
	}
	return nil, fmt.Errorf("id %v of odoo_data_fetcher.update_wizard not found", id)
}

// GetOdooDataFetcherUpdateWizards gets odoo_data_fetcher.update_wizard existing records.
func (c *Client) GetOdooDataFetcherUpdateWizards(ids []int64) (*OdooDataFetcherUpdateWizards, error) {
	ous := &OdooDataFetcherUpdateWizards{}
	if err := c.Read(OdooDataFetcherUpdateWizardModel, ids, nil, ous); err != nil {
		return nil, err
	}
	return ous, nil
}

// FindOdooDataFetcherUpdateWizard finds odoo_data_fetcher.update_wizard record by querying it with criteria.
func (c *Client) FindOdooDataFetcherUpdateWizard(criteria *Criteria) (*OdooDataFetcherUpdateWizard, error) {
	ous := &OdooDataFetcherUpdateWizards{}
	if err := c.SearchRead(OdooDataFetcherUpdateWizardModel, criteria, NewOptions().Limit(1), ous); err != nil {
		return nil, err
	}
	if ous != nil && len(*ous) > 0 {
		return &((*ous)[0]), nil
	}
	return nil, fmt.Errorf("odoo_data_fetcher.update_wizard was not found with criteria %v", criteria)
}

// FindOdooDataFetcherUpdateWizards finds odoo_data_fetcher.update_wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherUpdateWizards(criteria *Criteria, options *Options) (*OdooDataFetcherUpdateWizards, error) {
	ous := &OdooDataFetcherUpdateWizards{}
	if err := c.SearchRead(OdooDataFetcherUpdateWizardModel, criteria, options, ous); err != nil {
		return nil, err
	}
	return ous, nil
}

// FindOdooDataFetcherUpdateWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherUpdateWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OdooDataFetcherUpdateWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOdooDataFetcherUpdateWizardId finds record id by querying it with criteria.
func (c *Client) FindOdooDataFetcherUpdateWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OdooDataFetcherUpdateWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("odoo_data_fetcher.update_wizard was not found with criteria %v and options %v", criteria, options)
}
