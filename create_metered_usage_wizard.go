package odoo

import (
	"fmt"
)

// CreateMeteredUsageWizard represents create.metered.usage.wizard model.
type CreateMeteredUsageWizard struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	AccountId             *Many2One `xmlrpc:"account_id,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	OrderId               *Many2One `xmlrpc:"order_id,omptempty"`
	OrderName             *String   `xmlrpc:"order_name,omptempty"`
	OrderPartnerInvoiceId *Many2One `xmlrpc:"order_partner_invoice_id,omptempty"`
	ProductUomCategoryId  *Many2One `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomId          *Many2One `xmlrpc:"product_uom_id,omptempty"`
	SaleOrderLineId       *Many2One `xmlrpc:"sale_order_line_id,omptempty"`
	UnitAmount            *Float    `xmlrpc:"unit_amount,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CreateMeteredUsageWizards represents array of create.metered.usage.wizard model.
type CreateMeteredUsageWizards []CreateMeteredUsageWizard

// CreateMeteredUsageWizardModel is the odoo model name.
const CreateMeteredUsageWizardModel = "create.metered.usage.wizard"

// Many2One convert CreateMeteredUsageWizard to *Many2One.
func (cmuw *CreateMeteredUsageWizard) Many2One() *Many2One {
	return NewMany2One(cmuw.Id.Get(), "")
}

// CreateCreateMeteredUsageWizard creates a new create.metered.usage.wizard model and returns its id.
func (c *Client) CreateCreateMeteredUsageWizard(cmuw *CreateMeteredUsageWizard) (int64, error) {
	ids, err := c.CreateCreateMeteredUsageWizards([]*CreateMeteredUsageWizard{cmuw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCreateMeteredUsageWizard creates a new create.metered.usage.wizard model and returns its id.
func (c *Client) CreateCreateMeteredUsageWizards(cmuws []*CreateMeteredUsageWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmuws {
		vv = append(vv, v)
	}
	return c.Create(CreateMeteredUsageWizardModel, vv)
}

// UpdateCreateMeteredUsageWizard updates an existing create.metered.usage.wizard record.
func (c *Client) UpdateCreateMeteredUsageWizard(cmuw *CreateMeteredUsageWizard) error {
	return c.UpdateCreateMeteredUsageWizards([]int64{cmuw.Id.Get()}, cmuw)
}

// UpdateCreateMeteredUsageWizards updates existing create.metered.usage.wizard records.
// All records (represented by ids) will be updated by cmuw values.
func (c *Client) UpdateCreateMeteredUsageWizards(ids []int64, cmuw *CreateMeteredUsageWizard) error {
	return c.Update(CreateMeteredUsageWizardModel, ids, cmuw)
}

// DeleteCreateMeteredUsageWizard deletes an existing create.metered.usage.wizard record.
func (c *Client) DeleteCreateMeteredUsageWizard(id int64) error {
	return c.DeleteCreateMeteredUsageWizards([]int64{id})
}

// DeleteCreateMeteredUsageWizards deletes existing create.metered.usage.wizard records.
func (c *Client) DeleteCreateMeteredUsageWizards(ids []int64) error {
	return c.Delete(CreateMeteredUsageWizardModel, ids)
}

// GetCreateMeteredUsageWizard gets create.metered.usage.wizard existing record.
func (c *Client) GetCreateMeteredUsageWizard(id int64) (*CreateMeteredUsageWizard, error) {
	cmuws, err := c.GetCreateMeteredUsageWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmuws != nil && len(*cmuws) > 0 {
		return &((*cmuws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of create.metered.usage.wizard not found", id)
}

// GetCreateMeteredUsageWizards gets create.metered.usage.wizard existing records.
func (c *Client) GetCreateMeteredUsageWizards(ids []int64) (*CreateMeteredUsageWizards, error) {
	cmuws := &CreateMeteredUsageWizards{}
	if err := c.Read(CreateMeteredUsageWizardModel, ids, nil, cmuws); err != nil {
		return nil, err
	}
	return cmuws, nil
}

// FindCreateMeteredUsageWizard finds create.metered.usage.wizard record by querying it with criteria.
func (c *Client) FindCreateMeteredUsageWizard(criteria *Criteria) (*CreateMeteredUsageWizard, error) {
	cmuws := &CreateMeteredUsageWizards{}
	if err := c.SearchRead(CreateMeteredUsageWizardModel, criteria, NewOptions().Limit(1), cmuws); err != nil {
		return nil, err
	}
	if cmuws != nil && len(*cmuws) > 0 {
		return &((*cmuws)[0]), nil
	}
	return nil, fmt.Errorf("create.metered.usage.wizard was not found with criteria %v", criteria)
}

// FindCreateMeteredUsageWizards finds create.metered.usage.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCreateMeteredUsageWizards(criteria *Criteria, options *Options) (*CreateMeteredUsageWizards, error) {
	cmuws := &CreateMeteredUsageWizards{}
	if err := c.SearchRead(CreateMeteredUsageWizardModel, criteria, options, cmuws); err != nil {
		return nil, err
	}
	return cmuws, nil
}

// FindCreateMeteredUsageWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCreateMeteredUsageWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CreateMeteredUsageWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCreateMeteredUsageWizardId finds record id by querying it with criteria.
func (c *Client) FindCreateMeteredUsageWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CreateMeteredUsageWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("create.metered.usage.wizard was not found with criteria %v and options %v", criteria, options)
}
