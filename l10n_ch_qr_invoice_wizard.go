package odoo

import (
	"fmt"
)

// L10NChQrInvoiceWizard represents l10n_ch.qr_invoice.wizard model.
type L10NChQrInvoiceWizard struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	ClassicInvText *String   `xmlrpc:"classic_inv_text,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	IsrInvText     *String   `xmlrpc:"isr_inv_text,omptempty"`
	NbClassicInv   *Int      `xmlrpc:"nb_classic_inv,omptempty"`
	NbIsrInv       *Int      `xmlrpc:"nb_isr_inv,omptempty"`
	NbQrInv        *Int      `xmlrpc:"nb_qr_inv,omptempty"`
	QrInvText      *String   `xmlrpc:"qr_inv_text,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// L10NChQrInvoiceWizards represents array of l10n_ch.qr_invoice.wizard model.
type L10NChQrInvoiceWizards []L10NChQrInvoiceWizard

// L10NChQrInvoiceWizardModel is the odoo model name.
const L10NChQrInvoiceWizardModel = "l10n_ch.qr_invoice.wizard"

// Many2One convert L10NChQrInvoiceWizard to *Many2One.
func (lqw *L10NChQrInvoiceWizard) Many2One() *Many2One {
	return NewMany2One(lqw.Id.Get(), "")
}

// CreateL10NChQrInvoiceWizard creates a new l10n_ch.qr_invoice.wizard model and returns its id.
func (c *Client) CreateL10NChQrInvoiceWizard(lqw *L10NChQrInvoiceWizard) (int64, error) {
	ids, err := c.CreateL10NChQrInvoiceWizards([]*L10NChQrInvoiceWizard{lqw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NChQrInvoiceWizard creates a new l10n_ch.qr_invoice.wizard model and returns its id.
func (c *Client) CreateL10NChQrInvoiceWizards(lqws []*L10NChQrInvoiceWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lqws {
		vv = append(vv, v)
	}
	return c.Create(L10NChQrInvoiceWizardModel, vv)
}

// UpdateL10NChQrInvoiceWizard updates an existing l10n_ch.qr_invoice.wizard record.
func (c *Client) UpdateL10NChQrInvoiceWizard(lqw *L10NChQrInvoiceWizard) error {
	return c.UpdateL10NChQrInvoiceWizards([]int64{lqw.Id.Get()}, lqw)
}

// UpdateL10NChQrInvoiceWizards updates existing l10n_ch.qr_invoice.wizard records.
// All records (represented by ids) will be updated by lqw values.
func (c *Client) UpdateL10NChQrInvoiceWizards(ids []int64, lqw *L10NChQrInvoiceWizard) error {
	return c.Update(L10NChQrInvoiceWizardModel, ids, lqw)
}

// DeleteL10NChQrInvoiceWizard deletes an existing l10n_ch.qr_invoice.wizard record.
func (c *Client) DeleteL10NChQrInvoiceWizard(id int64) error {
	return c.DeleteL10NChQrInvoiceWizards([]int64{id})
}

// DeleteL10NChQrInvoiceWizards deletes existing l10n_ch.qr_invoice.wizard records.
func (c *Client) DeleteL10NChQrInvoiceWizards(ids []int64) error {
	return c.Delete(L10NChQrInvoiceWizardModel, ids)
}

// GetL10NChQrInvoiceWizard gets l10n_ch.qr_invoice.wizard existing record.
func (c *Client) GetL10NChQrInvoiceWizard(id int64) (*L10NChQrInvoiceWizard, error) {
	lqws, err := c.GetL10NChQrInvoiceWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if lqws != nil && len(*lqws) > 0 {
		return &((*lqws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of l10n_ch.qr_invoice.wizard not found", id)
}

// GetL10NChQrInvoiceWizards gets l10n_ch.qr_invoice.wizard existing records.
func (c *Client) GetL10NChQrInvoiceWizards(ids []int64) (*L10NChQrInvoiceWizards, error) {
	lqws := &L10NChQrInvoiceWizards{}
	if err := c.Read(L10NChQrInvoiceWizardModel, ids, nil, lqws); err != nil {
		return nil, err
	}
	return lqws, nil
}

// FindL10NChQrInvoiceWizard finds l10n_ch.qr_invoice.wizard record by querying it with criteria.
func (c *Client) FindL10NChQrInvoiceWizard(criteria *Criteria) (*L10NChQrInvoiceWizard, error) {
	lqws := &L10NChQrInvoiceWizards{}
	if err := c.SearchRead(L10NChQrInvoiceWizardModel, criteria, NewOptions().Limit(1), lqws); err != nil {
		return nil, err
	}
	if lqws != nil && len(*lqws) > 0 {
		return &((*lqws)[0]), nil
	}
	return nil, fmt.Errorf("l10n_ch.qr_invoice.wizard was not found with criteria %v", criteria)
}

// FindL10NChQrInvoiceWizards finds l10n_ch.qr_invoice.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NChQrInvoiceWizards(criteria *Criteria, options *Options) (*L10NChQrInvoiceWizards, error) {
	lqws := &L10NChQrInvoiceWizards{}
	if err := c.SearchRead(L10NChQrInvoiceWizardModel, criteria, options, lqws); err != nil {
		return nil, err
	}
	return lqws, nil
}

// FindL10NChQrInvoiceWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NChQrInvoiceWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(L10NChQrInvoiceWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindL10NChQrInvoiceWizardId finds record id by querying it with criteria.
func (c *Client) FindL10NChQrInvoiceWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NChQrInvoiceWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("l10n_ch.qr_invoice.wizard was not found with criteria %v and options %v", criteria, options)
}
