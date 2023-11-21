package odoo

import (
	"fmt"
)

// SaleSubscriptionCloseReasonWizard represents sale.subscription.close.reason.wizard model.
type SaleSubscriptionCloseReasonWizard struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CloseReasonId *Many2One `xmlrpc:"close_reason_id,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleSubscriptionCloseReasonWizards represents array of sale.subscription.close.reason.wizard model.
type SaleSubscriptionCloseReasonWizards []SaleSubscriptionCloseReasonWizard

// SaleSubscriptionCloseReasonWizardModel is the odoo model name.
const SaleSubscriptionCloseReasonWizardModel = "sale.subscription.close.reason.wizard"

// Many2One convert SaleSubscriptionCloseReasonWizard to *Many2One.
func (sscrw *SaleSubscriptionCloseReasonWizard) Many2One() *Many2One {
	return NewMany2One(sscrw.Id.Get(), "")
}

// CreateSaleSubscriptionCloseReasonWizard creates a new sale.subscription.close.reason.wizard model and returns its id.
func (c *Client) CreateSaleSubscriptionCloseReasonWizard(sscrw *SaleSubscriptionCloseReasonWizard) (int64, error) {
	ids, err := c.CreateSaleSubscriptionCloseReasonWizards([]*SaleSubscriptionCloseReasonWizard{sscrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionCloseReasonWizard creates a new sale.subscription.close.reason.wizard model and returns its id.
func (c *Client) CreateSaleSubscriptionCloseReasonWizards(sscrws []*SaleSubscriptionCloseReasonWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range sscrws {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionCloseReasonWizardModel, vv)
}

// UpdateSaleSubscriptionCloseReasonWizard updates an existing sale.subscription.close.reason.wizard record.
func (c *Client) UpdateSaleSubscriptionCloseReasonWizard(sscrw *SaleSubscriptionCloseReasonWizard) error {
	return c.UpdateSaleSubscriptionCloseReasonWizards([]int64{sscrw.Id.Get()}, sscrw)
}

// UpdateSaleSubscriptionCloseReasonWizards updates existing sale.subscription.close.reason.wizard records.
// All records (represented by ids) will be updated by sscrw values.
func (c *Client) UpdateSaleSubscriptionCloseReasonWizards(ids []int64, sscrw *SaleSubscriptionCloseReasonWizard) error {
	return c.Update(SaleSubscriptionCloseReasonWizardModel, ids, sscrw)
}

// DeleteSaleSubscriptionCloseReasonWizard deletes an existing sale.subscription.close.reason.wizard record.
func (c *Client) DeleteSaleSubscriptionCloseReasonWizard(id int64) error {
	return c.DeleteSaleSubscriptionCloseReasonWizards([]int64{id})
}

// DeleteSaleSubscriptionCloseReasonWizards deletes existing sale.subscription.close.reason.wizard records.
func (c *Client) DeleteSaleSubscriptionCloseReasonWizards(ids []int64) error {
	return c.Delete(SaleSubscriptionCloseReasonWizardModel, ids)
}

// GetSaleSubscriptionCloseReasonWizard gets sale.subscription.close.reason.wizard existing record.
func (c *Client) GetSaleSubscriptionCloseReasonWizard(id int64) (*SaleSubscriptionCloseReasonWizard, error) {
	sscrws, err := c.GetSaleSubscriptionCloseReasonWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if sscrws != nil && len(*sscrws) > 0 {
		return &((*sscrws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.close.reason.wizard not found", id)
}

// GetSaleSubscriptionCloseReasonWizards gets sale.subscription.close.reason.wizard existing records.
func (c *Client) GetSaleSubscriptionCloseReasonWizards(ids []int64) (*SaleSubscriptionCloseReasonWizards, error) {
	sscrws := &SaleSubscriptionCloseReasonWizards{}
	if err := c.Read(SaleSubscriptionCloseReasonWizardModel, ids, nil, sscrws); err != nil {
		return nil, err
	}
	return sscrws, nil
}

// FindSaleSubscriptionCloseReasonWizard finds sale.subscription.close.reason.wizard record by querying it with criteria.
func (c *Client) FindSaleSubscriptionCloseReasonWizard(criteria *Criteria) (*SaleSubscriptionCloseReasonWizard, error) {
	sscrws := &SaleSubscriptionCloseReasonWizards{}
	if err := c.SearchRead(SaleSubscriptionCloseReasonWizardModel, criteria, NewOptions().Limit(1), sscrws); err != nil {
		return nil, err
	}
	if sscrws != nil && len(*sscrws) > 0 {
		return &((*sscrws)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.close.reason.wizard was not found with criteria %v", criteria)
}

// FindSaleSubscriptionCloseReasonWizards finds sale.subscription.close.reason.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionCloseReasonWizards(criteria *Criteria, options *Options) (*SaleSubscriptionCloseReasonWizards, error) {
	sscrws := &SaleSubscriptionCloseReasonWizards{}
	if err := c.SearchRead(SaleSubscriptionCloseReasonWizardModel, criteria, options, sscrws); err != nil {
		return nil, err
	}
	return sscrws, nil
}

// FindSaleSubscriptionCloseReasonWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionCloseReasonWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionCloseReasonWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionCloseReasonWizardId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionCloseReasonWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionCloseReasonWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.close.reason.wizard was not found with criteria %v and options %v", criteria, options)
}
