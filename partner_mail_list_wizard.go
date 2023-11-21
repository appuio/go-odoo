package odoo

import (
	"fmt"
)

// PartnerMailListWizard represents partner.mail.list.wizard model.
type PartnerMailListWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MailListId  *Many2One `xmlrpc:"mail_list_id,omptempty"`
	PartnerIds  *Relation `xmlrpc:"partner_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PartnerMailListWizards represents array of partner.mail.list.wizard model.
type PartnerMailListWizards []PartnerMailListWizard

// PartnerMailListWizardModel is the odoo model name.
const PartnerMailListWizardModel = "partner.mail.list.wizard"

// Many2One convert PartnerMailListWizard to *Many2One.
func (pmlw *PartnerMailListWizard) Many2One() *Many2One {
	return NewMany2One(pmlw.Id.Get(), "")
}

// CreatePartnerMailListWizard creates a new partner.mail.list.wizard model and returns its id.
func (c *Client) CreatePartnerMailListWizard(pmlw *PartnerMailListWizard) (int64, error) {
	ids, err := c.CreatePartnerMailListWizards([]*PartnerMailListWizard{pmlw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePartnerMailListWizard creates a new partner.mail.list.wizard model and returns its id.
func (c *Client) CreatePartnerMailListWizards(pmlws []*PartnerMailListWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pmlws {
		vv = append(vv, v)
	}
	return c.Create(PartnerMailListWizardModel, vv)
}

// UpdatePartnerMailListWizard updates an existing partner.mail.list.wizard record.
func (c *Client) UpdatePartnerMailListWizard(pmlw *PartnerMailListWizard) error {
	return c.UpdatePartnerMailListWizards([]int64{pmlw.Id.Get()}, pmlw)
}

// UpdatePartnerMailListWizards updates existing partner.mail.list.wizard records.
// All records (represented by ids) will be updated by pmlw values.
func (c *Client) UpdatePartnerMailListWizards(ids []int64, pmlw *PartnerMailListWizard) error {
	return c.Update(PartnerMailListWizardModel, ids, pmlw)
}

// DeletePartnerMailListWizard deletes an existing partner.mail.list.wizard record.
func (c *Client) DeletePartnerMailListWizard(id int64) error {
	return c.DeletePartnerMailListWizards([]int64{id})
}

// DeletePartnerMailListWizards deletes existing partner.mail.list.wizard records.
func (c *Client) DeletePartnerMailListWizards(ids []int64) error {
	return c.Delete(PartnerMailListWizardModel, ids)
}

// GetPartnerMailListWizard gets partner.mail.list.wizard existing record.
func (c *Client) GetPartnerMailListWizard(id int64) (*PartnerMailListWizard, error) {
	pmlws, err := c.GetPartnerMailListWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if pmlws != nil && len(*pmlws) > 0 {
		return &((*pmlws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of partner.mail.list.wizard not found", id)
}

// GetPartnerMailListWizards gets partner.mail.list.wizard existing records.
func (c *Client) GetPartnerMailListWizards(ids []int64) (*PartnerMailListWizards, error) {
	pmlws := &PartnerMailListWizards{}
	if err := c.Read(PartnerMailListWizardModel, ids, nil, pmlws); err != nil {
		return nil, err
	}
	return pmlws, nil
}

// FindPartnerMailListWizard finds partner.mail.list.wizard record by querying it with criteria.
func (c *Client) FindPartnerMailListWizard(criteria *Criteria) (*PartnerMailListWizard, error) {
	pmlws := &PartnerMailListWizards{}
	if err := c.SearchRead(PartnerMailListWizardModel, criteria, NewOptions().Limit(1), pmlws); err != nil {
		return nil, err
	}
	if pmlws != nil && len(*pmlws) > 0 {
		return &((*pmlws)[0]), nil
	}
	return nil, fmt.Errorf("partner.mail.list.wizard was not found with criteria %v", criteria)
}

// FindPartnerMailListWizards finds partner.mail.list.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPartnerMailListWizards(criteria *Criteria, options *Options) (*PartnerMailListWizards, error) {
	pmlws := &PartnerMailListWizards{}
	if err := c.SearchRead(PartnerMailListWizardModel, criteria, options, pmlws); err != nil {
		return nil, err
	}
	return pmlws, nil
}

// FindPartnerMailListWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPartnerMailListWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PartnerMailListWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPartnerMailListWizardId finds record id by querying it with criteria.
func (c *Client) FindPartnerMailListWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PartnerMailListWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("partner.mail.list.wizard was not found with criteria %v and options %v", criteria, options)
}
