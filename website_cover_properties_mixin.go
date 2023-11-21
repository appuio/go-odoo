package odoo

import (
	"fmt"
)

// WebsiteCoverPropertiesMixin represents website.cover_properties.mixin model.
type WebsiteCoverPropertiesMixin struct {
	CoverProperties *String `xmlrpc:"cover_properties,omptempty"`
}

// WebsiteCoverPropertiesMixins represents array of website.cover_properties.mixin model.
type WebsiteCoverPropertiesMixins []WebsiteCoverPropertiesMixin

// WebsiteCoverPropertiesMixinModel is the odoo model name.
const WebsiteCoverPropertiesMixinModel = "website.cover_properties.mixin"

// Many2One convert WebsiteCoverPropertiesMixin to *Many2One.
func (wcm *WebsiteCoverPropertiesMixin) Many2One() *Many2One {
	return NewMany2One(wcm.Id.Get(), "")
}

// CreateWebsiteCoverPropertiesMixin creates a new website.cover_properties.mixin model and returns its id.
func (c *Client) CreateWebsiteCoverPropertiesMixin(wcm *WebsiteCoverPropertiesMixin) (int64, error) {
	ids, err := c.CreateWebsiteCoverPropertiesMixins([]*WebsiteCoverPropertiesMixin{wcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteCoverPropertiesMixin creates a new website.cover_properties.mixin model and returns its id.
func (c *Client) CreateWebsiteCoverPropertiesMixins(wcms []*WebsiteCoverPropertiesMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range wcms {
		vv = append(vv, v)
	}
	return c.Create(WebsiteCoverPropertiesMixinModel, vv)
}

// UpdateWebsiteCoverPropertiesMixin updates an existing website.cover_properties.mixin record.
func (c *Client) UpdateWebsiteCoverPropertiesMixin(wcm *WebsiteCoverPropertiesMixin) error {
	return c.UpdateWebsiteCoverPropertiesMixins([]int64{wcm.Id.Get()}, wcm)
}

// UpdateWebsiteCoverPropertiesMixins updates existing website.cover_properties.mixin records.
// All records (represented by ids) will be updated by wcm values.
func (c *Client) UpdateWebsiteCoverPropertiesMixins(ids []int64, wcm *WebsiteCoverPropertiesMixin) error {
	return c.Update(WebsiteCoverPropertiesMixinModel, ids, wcm)
}

// DeleteWebsiteCoverPropertiesMixin deletes an existing website.cover_properties.mixin record.
func (c *Client) DeleteWebsiteCoverPropertiesMixin(id int64) error {
	return c.DeleteWebsiteCoverPropertiesMixins([]int64{id})
}

// DeleteWebsiteCoverPropertiesMixins deletes existing website.cover_properties.mixin records.
func (c *Client) DeleteWebsiteCoverPropertiesMixins(ids []int64) error {
	return c.Delete(WebsiteCoverPropertiesMixinModel, ids)
}

// GetWebsiteCoverPropertiesMixin gets website.cover_properties.mixin existing record.
func (c *Client) GetWebsiteCoverPropertiesMixin(id int64) (*WebsiteCoverPropertiesMixin, error) {
	wcms, err := c.GetWebsiteCoverPropertiesMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if wcms != nil && len(*wcms) > 0 {
		return &((*wcms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.cover_properties.mixin not found", id)
}

// GetWebsiteCoverPropertiesMixins gets website.cover_properties.mixin existing records.
func (c *Client) GetWebsiteCoverPropertiesMixins(ids []int64) (*WebsiteCoverPropertiesMixins, error) {
	wcms := &WebsiteCoverPropertiesMixins{}
	if err := c.Read(WebsiteCoverPropertiesMixinModel, ids, nil, wcms); err != nil {
		return nil, err
	}
	return wcms, nil
}

// FindWebsiteCoverPropertiesMixin finds website.cover_properties.mixin record by querying it with criteria.
func (c *Client) FindWebsiteCoverPropertiesMixin(criteria *Criteria) (*WebsiteCoverPropertiesMixin, error) {
	wcms := &WebsiteCoverPropertiesMixins{}
	if err := c.SearchRead(WebsiteCoverPropertiesMixinModel, criteria, NewOptions().Limit(1), wcms); err != nil {
		return nil, err
	}
	if wcms != nil && len(*wcms) > 0 {
		return &((*wcms)[0]), nil
	}
	return nil, fmt.Errorf("website.cover_properties.mixin was not found with criteria %v", criteria)
}

// FindWebsiteCoverPropertiesMixins finds website.cover_properties.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteCoverPropertiesMixins(criteria *Criteria, options *Options) (*WebsiteCoverPropertiesMixins, error) {
	wcms := &WebsiteCoverPropertiesMixins{}
	if err := c.SearchRead(WebsiteCoverPropertiesMixinModel, criteria, options, wcms); err != nil {
		return nil, err
	}
	return wcms, nil
}

// FindWebsiteCoverPropertiesMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteCoverPropertiesMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteCoverPropertiesMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteCoverPropertiesMixinId finds record id by querying it with criteria.
func (c *Client) FindWebsiteCoverPropertiesMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteCoverPropertiesMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.cover_properties.mixin was not found with criteria %v and options %v", criteria, options)
}
