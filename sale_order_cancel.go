package odoo

import (
	"fmt"
)

// SaleOrderCancel represents sale.order.cancel model.
type SaleOrderCancel struct {
	LastUpdate                 *Time     `xmlrpc:"__last_update,omptempty"`
	AuthorId                   *Many2One `xmlrpc:"author_id,omptempty"`
	Body                       *String   `xmlrpc:"body,omptempty"`
	CanEditBody                *Bool     `xmlrpc:"can_edit_body,omptempty"`
	CreateDate                 *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayDeliveryAlert       *Bool     `xmlrpc:"display_delivery_alert,omptempty"`
	DisplayInvoiceAlert        *Bool     `xmlrpc:"display_invoice_alert,omptempty"`
	DisplayName                *String   `xmlrpc:"display_name,omptempty"`
	DisplayPurchaseOrdersAlert *Bool     `xmlrpc:"display_purchase_orders_alert,omptempty"`
	EmailFrom                  *String   `xmlrpc:"email_from,omptempty"`
	Id                         *Int      `xmlrpc:"id,omptempty"`
	IsMailTemplateEditor       *Bool     `xmlrpc:"is_mail_template_editor,omptempty"`
	Lang                       *String   `xmlrpc:"lang,omptempty"`
	OrderId                    *Many2One `xmlrpc:"order_id,omptempty"`
	RecipientIds               *Relation `xmlrpc:"recipient_ids,omptempty"`
	RenderModel                *String   `xmlrpc:"render_model,omptempty"`
	Subject                    *String   `xmlrpc:"subject,omptempty"`
	TemplateId                 *Many2One `xmlrpc:"template_id,omptempty"`
	WriteDate                  *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderCancels represents array of sale.order.cancel model.
type SaleOrderCancels []SaleOrderCancel

// SaleOrderCancelModel is the odoo model name.
const SaleOrderCancelModel = "sale.order.cancel"

// Many2One convert SaleOrderCancel to *Many2One.
func (soc *SaleOrderCancel) Many2One() *Many2One {
	return NewMany2One(soc.Id.Get(), "")
}

// CreateSaleOrderCancel creates a new sale.order.cancel model and returns its id.
func (c *Client) CreateSaleOrderCancel(soc *SaleOrderCancel) (int64, error) {
	ids, err := c.CreateSaleOrderCancels([]*SaleOrderCancel{soc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderCancel creates a new sale.order.cancel model and returns its id.
func (c *Client) CreateSaleOrderCancels(socs []*SaleOrderCancel) ([]int64, error) {
	var vv []interface{}
	for _, v := range socs {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderCancelModel, vv)
}

// UpdateSaleOrderCancel updates an existing sale.order.cancel record.
func (c *Client) UpdateSaleOrderCancel(soc *SaleOrderCancel) error {
	return c.UpdateSaleOrderCancels([]int64{soc.Id.Get()}, soc)
}

// UpdateSaleOrderCancels updates existing sale.order.cancel records.
// All records (represented by ids) will be updated by soc values.
func (c *Client) UpdateSaleOrderCancels(ids []int64, soc *SaleOrderCancel) error {
	return c.Update(SaleOrderCancelModel, ids, soc)
}

// DeleteSaleOrderCancel deletes an existing sale.order.cancel record.
func (c *Client) DeleteSaleOrderCancel(id int64) error {
	return c.DeleteSaleOrderCancels([]int64{id})
}

// DeleteSaleOrderCancels deletes existing sale.order.cancel records.
func (c *Client) DeleteSaleOrderCancels(ids []int64) error {
	return c.Delete(SaleOrderCancelModel, ids)
}

// GetSaleOrderCancel gets sale.order.cancel existing record.
func (c *Client) GetSaleOrderCancel(id int64) (*SaleOrderCancel, error) {
	socs, err := c.GetSaleOrderCancels([]int64{id})
	if err != nil {
		return nil, err
	}
	if socs != nil && len(*socs) > 0 {
		return &((*socs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.cancel not found", id)
}

// GetSaleOrderCancels gets sale.order.cancel existing records.
func (c *Client) GetSaleOrderCancels(ids []int64) (*SaleOrderCancels, error) {
	socs := &SaleOrderCancels{}
	if err := c.Read(SaleOrderCancelModel, ids, nil, socs); err != nil {
		return nil, err
	}
	return socs, nil
}

// FindSaleOrderCancel finds sale.order.cancel record by querying it with criteria.
func (c *Client) FindSaleOrderCancel(criteria *Criteria) (*SaleOrderCancel, error) {
	socs := &SaleOrderCancels{}
	if err := c.SearchRead(SaleOrderCancelModel, criteria, NewOptions().Limit(1), socs); err != nil {
		return nil, err
	}
	if socs != nil && len(*socs) > 0 {
		return &((*socs)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.cancel was not found with criteria %v", criteria)
}

// FindSaleOrderCancels finds sale.order.cancel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCancels(criteria *Criteria, options *Options) (*SaleOrderCancels, error) {
	socs := &SaleOrderCancels{}
	if err := c.SearchRead(SaleOrderCancelModel, criteria, options, socs); err != nil {
		return nil, err
	}
	return socs, nil
}

// FindSaleOrderCancelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderCancelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderCancelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderCancelId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderCancelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderCancelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.cancel was not found with criteria %v and options %v", criteria, options)
}
