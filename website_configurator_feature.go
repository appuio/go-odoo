package odoo

import (
	"fmt"
)

// WebsiteConfiguratorFeature represents website.configurator.feature model.
type WebsiteConfiguratorFeature struct {
	LastUpdate                *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omptempty"`
	Description               *String   `xmlrpc:"description,omptempty"`
	DisplayName               *String   `xmlrpc:"display_name,omptempty"`
	FeatureUrl                *String   `xmlrpc:"feature_url,omptempty"`
	IapPageCode               *String   `xmlrpc:"iap_page_code,omptempty"`
	Icon                      *String   `xmlrpc:"icon,omptempty"`
	Id                        *Int      `xmlrpc:"id,omptempty"`
	MenuCompany               *Bool     `xmlrpc:"menu_company,omptempty"`
	MenuSequence              *Int      `xmlrpc:"menu_sequence,omptempty"`
	ModuleId                  *Many2One `xmlrpc:"module_id,omptempty"`
	Name                      *String   `xmlrpc:"name,omptempty"`
	PageViewId                *Many2One `xmlrpc:"page_view_id,omptempty"`
	Sequence                  *Int      `xmlrpc:"sequence,omptempty"`
	WebsiteConfigPreselection *String   `xmlrpc:"website_config_preselection,omptempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WebsiteConfiguratorFeatures represents array of website.configurator.feature model.
type WebsiteConfiguratorFeatures []WebsiteConfiguratorFeature

// WebsiteConfiguratorFeatureModel is the odoo model name.
const WebsiteConfiguratorFeatureModel = "website.configurator.feature"

// Many2One convert WebsiteConfiguratorFeature to *Many2One.
func (wcf *WebsiteConfiguratorFeature) Many2One() *Many2One {
	return NewMany2One(wcf.Id.Get(), "")
}

// CreateWebsiteConfiguratorFeature creates a new website.configurator.feature model and returns its id.
func (c *Client) CreateWebsiteConfiguratorFeature(wcf *WebsiteConfiguratorFeature) (int64, error) {
	ids, err := c.CreateWebsiteConfiguratorFeatures([]*WebsiteConfiguratorFeature{wcf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteConfiguratorFeature creates a new website.configurator.feature model and returns its id.
func (c *Client) CreateWebsiteConfiguratorFeatures(wcfs []*WebsiteConfiguratorFeature) ([]int64, error) {
	var vv []interface{}
	for _, v := range wcfs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteConfiguratorFeatureModel, vv)
}

// UpdateWebsiteConfiguratorFeature updates an existing website.configurator.feature record.
func (c *Client) UpdateWebsiteConfiguratorFeature(wcf *WebsiteConfiguratorFeature) error {
	return c.UpdateWebsiteConfiguratorFeatures([]int64{wcf.Id.Get()}, wcf)
}

// UpdateWebsiteConfiguratorFeatures updates existing website.configurator.feature records.
// All records (represented by ids) will be updated by wcf values.
func (c *Client) UpdateWebsiteConfiguratorFeatures(ids []int64, wcf *WebsiteConfiguratorFeature) error {
	return c.Update(WebsiteConfiguratorFeatureModel, ids, wcf)
}

// DeleteWebsiteConfiguratorFeature deletes an existing website.configurator.feature record.
func (c *Client) DeleteWebsiteConfiguratorFeature(id int64) error {
	return c.DeleteWebsiteConfiguratorFeatures([]int64{id})
}

// DeleteWebsiteConfiguratorFeatures deletes existing website.configurator.feature records.
func (c *Client) DeleteWebsiteConfiguratorFeatures(ids []int64) error {
	return c.Delete(WebsiteConfiguratorFeatureModel, ids)
}

// GetWebsiteConfiguratorFeature gets website.configurator.feature existing record.
func (c *Client) GetWebsiteConfiguratorFeature(id int64) (*WebsiteConfiguratorFeature, error) {
	wcfs, err := c.GetWebsiteConfiguratorFeatures([]int64{id})
	if err != nil {
		return nil, err
	}
	if wcfs != nil && len(*wcfs) > 0 {
		return &((*wcfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.configurator.feature not found", id)
}

// GetWebsiteConfiguratorFeatures gets website.configurator.feature existing records.
func (c *Client) GetWebsiteConfiguratorFeatures(ids []int64) (*WebsiteConfiguratorFeatures, error) {
	wcfs := &WebsiteConfiguratorFeatures{}
	if err := c.Read(WebsiteConfiguratorFeatureModel, ids, nil, wcfs); err != nil {
		return nil, err
	}
	return wcfs, nil
}

// FindWebsiteConfiguratorFeature finds website.configurator.feature record by querying it with criteria.
func (c *Client) FindWebsiteConfiguratorFeature(criteria *Criteria) (*WebsiteConfiguratorFeature, error) {
	wcfs := &WebsiteConfiguratorFeatures{}
	if err := c.SearchRead(WebsiteConfiguratorFeatureModel, criteria, NewOptions().Limit(1), wcfs); err != nil {
		return nil, err
	}
	if wcfs != nil && len(*wcfs) > 0 {
		return &((*wcfs)[0]), nil
	}
	return nil, fmt.Errorf("website.configurator.feature was not found with criteria %v", criteria)
}

// FindWebsiteConfiguratorFeatures finds website.configurator.feature records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteConfiguratorFeatures(criteria *Criteria, options *Options) (*WebsiteConfiguratorFeatures, error) {
	wcfs := &WebsiteConfiguratorFeatures{}
	if err := c.SearchRead(WebsiteConfiguratorFeatureModel, criteria, options, wcfs); err != nil {
		return nil, err
	}
	return wcfs, nil
}

// FindWebsiteConfiguratorFeatureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteConfiguratorFeatureIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteConfiguratorFeatureModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteConfiguratorFeatureId finds record id by querying it with criteria.
func (c *Client) FindWebsiteConfiguratorFeatureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteConfiguratorFeatureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.configurator.feature was not found with criteria %v and options %v", criteria, options)
}
