package odoo

import (
	"fmt"
)

// WebsiteSaleExtraField represents website.sale.extra.field model.
type WebsiteSaleExtraField struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FieldId     *Many2One `xmlrpc:"field_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Label       *String   `xmlrpc:"label,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WebsiteId   *Many2One `xmlrpc:"website_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WebsiteSaleExtraFields represents array of website.sale.extra.field model.
type WebsiteSaleExtraFields []WebsiteSaleExtraField

// WebsiteSaleExtraFieldModel is the odoo model name.
const WebsiteSaleExtraFieldModel = "website.sale.extra.field"

// Many2One convert WebsiteSaleExtraField to *Many2One.
func (wsef *WebsiteSaleExtraField) Many2One() *Many2One {
	return NewMany2One(wsef.Id.Get(), "")
}

// CreateWebsiteSaleExtraField creates a new website.sale.extra.field model and returns its id.
func (c *Client) CreateWebsiteSaleExtraField(wsef *WebsiteSaleExtraField) (int64, error) {
	ids, err := c.CreateWebsiteSaleExtraFields([]*WebsiteSaleExtraField{wsef})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteSaleExtraField creates a new website.sale.extra.field model and returns its id.
func (c *Client) CreateWebsiteSaleExtraFields(wsefs []*WebsiteSaleExtraField) ([]int64, error) {
	var vv []interface{}
	for _, v := range wsefs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteSaleExtraFieldModel, vv)
}

// UpdateWebsiteSaleExtraField updates an existing website.sale.extra.field record.
func (c *Client) UpdateWebsiteSaleExtraField(wsef *WebsiteSaleExtraField) error {
	return c.UpdateWebsiteSaleExtraFields([]int64{wsef.Id.Get()}, wsef)
}

// UpdateWebsiteSaleExtraFields updates existing website.sale.extra.field records.
// All records (represented by ids) will be updated by wsef values.
func (c *Client) UpdateWebsiteSaleExtraFields(ids []int64, wsef *WebsiteSaleExtraField) error {
	return c.Update(WebsiteSaleExtraFieldModel, ids, wsef)
}

// DeleteWebsiteSaleExtraField deletes an existing website.sale.extra.field record.
func (c *Client) DeleteWebsiteSaleExtraField(id int64) error {
	return c.DeleteWebsiteSaleExtraFields([]int64{id})
}

// DeleteWebsiteSaleExtraFields deletes existing website.sale.extra.field records.
func (c *Client) DeleteWebsiteSaleExtraFields(ids []int64) error {
	return c.Delete(WebsiteSaleExtraFieldModel, ids)
}

// GetWebsiteSaleExtraField gets website.sale.extra.field existing record.
func (c *Client) GetWebsiteSaleExtraField(id int64) (*WebsiteSaleExtraField, error) {
	wsefs, err := c.GetWebsiteSaleExtraFields([]int64{id})
	if err != nil {
		return nil, err
	}
	if wsefs != nil && len(*wsefs) > 0 {
		return &((*wsefs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.sale.extra.field not found", id)
}

// GetWebsiteSaleExtraFields gets website.sale.extra.field existing records.
func (c *Client) GetWebsiteSaleExtraFields(ids []int64) (*WebsiteSaleExtraFields, error) {
	wsefs := &WebsiteSaleExtraFields{}
	if err := c.Read(WebsiteSaleExtraFieldModel, ids, nil, wsefs); err != nil {
		return nil, err
	}
	return wsefs, nil
}

// FindWebsiteSaleExtraField finds website.sale.extra.field record by querying it with criteria.
func (c *Client) FindWebsiteSaleExtraField(criteria *Criteria) (*WebsiteSaleExtraField, error) {
	wsefs := &WebsiteSaleExtraFields{}
	if err := c.SearchRead(WebsiteSaleExtraFieldModel, criteria, NewOptions().Limit(1), wsefs); err != nil {
		return nil, err
	}
	if wsefs != nil && len(*wsefs) > 0 {
		return &((*wsefs)[0]), nil
	}
	return nil, fmt.Errorf("website.sale.extra.field was not found with criteria %v", criteria)
}

// FindWebsiteSaleExtraFields finds website.sale.extra.field records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSaleExtraFields(criteria *Criteria, options *Options) (*WebsiteSaleExtraFields, error) {
	wsefs := &WebsiteSaleExtraFields{}
	if err := c.SearchRead(WebsiteSaleExtraFieldModel, criteria, options, wsefs); err != nil {
		return nil, err
	}
	return wsefs, nil
}

// FindWebsiteSaleExtraFieldIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSaleExtraFieldIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteSaleExtraFieldModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteSaleExtraFieldId finds record id by querying it with criteria.
func (c *Client) FindWebsiteSaleExtraFieldId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteSaleExtraFieldModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.sale.extra.field was not found with criteria %v and options %v", criteria, options)
}
