package odoo

import (
	"fmt"
)

// WebsiteSnippetFilter represents website.snippet.filter model.
type WebsiteSnippetFilter struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	ActionServerId      *Many2One `xmlrpc:"action_server_id,omptempty"`
	CanPublish          *Bool     `xmlrpc:"can_publish,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	FieldNames          *String   `xmlrpc:"field_names,omptempty"`
	FilterId            *Many2One `xmlrpc:"filter_id,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	IsPublished         *Bool     `xmlrpc:"is_published,omptempty"`
	Limit               *Int      `xmlrpc:"limit,omptempty"`
	ModelName           *String   `xmlrpc:"model_name,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	ProductCrossSelling *Bool     `xmlrpc:"product_cross_selling,omptempty"`
	WebsiteId           *Many2One `xmlrpc:"website_id,omptempty"`
	WebsitePublished    *Bool     `xmlrpc:"website_published,omptempty"`
	WebsiteUrl          *String   `xmlrpc:"website_url,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WebsiteSnippetFilters represents array of website.snippet.filter model.
type WebsiteSnippetFilters []WebsiteSnippetFilter

// WebsiteSnippetFilterModel is the odoo model name.
const WebsiteSnippetFilterModel = "website.snippet.filter"

// Many2One convert WebsiteSnippetFilter to *Many2One.
func (wsf *WebsiteSnippetFilter) Many2One() *Many2One {
	return NewMany2One(wsf.Id.Get(), "")
}

// CreateWebsiteSnippetFilter creates a new website.snippet.filter model and returns its id.
func (c *Client) CreateWebsiteSnippetFilter(wsf *WebsiteSnippetFilter) (int64, error) {
	ids, err := c.CreateWebsiteSnippetFilters([]*WebsiteSnippetFilter{wsf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteSnippetFilter creates a new website.snippet.filter model and returns its id.
func (c *Client) CreateWebsiteSnippetFilters(wsfs []*WebsiteSnippetFilter) ([]int64, error) {
	var vv []interface{}
	for _, v := range wsfs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteSnippetFilterModel, vv)
}

// UpdateWebsiteSnippetFilter updates an existing website.snippet.filter record.
func (c *Client) UpdateWebsiteSnippetFilter(wsf *WebsiteSnippetFilter) error {
	return c.UpdateWebsiteSnippetFilters([]int64{wsf.Id.Get()}, wsf)
}

// UpdateWebsiteSnippetFilters updates existing website.snippet.filter records.
// All records (represented by ids) will be updated by wsf values.
func (c *Client) UpdateWebsiteSnippetFilters(ids []int64, wsf *WebsiteSnippetFilter) error {
	return c.Update(WebsiteSnippetFilterModel, ids, wsf)
}

// DeleteWebsiteSnippetFilter deletes an existing website.snippet.filter record.
func (c *Client) DeleteWebsiteSnippetFilter(id int64) error {
	return c.DeleteWebsiteSnippetFilters([]int64{id})
}

// DeleteWebsiteSnippetFilters deletes existing website.snippet.filter records.
func (c *Client) DeleteWebsiteSnippetFilters(ids []int64) error {
	return c.Delete(WebsiteSnippetFilterModel, ids)
}

// GetWebsiteSnippetFilter gets website.snippet.filter existing record.
func (c *Client) GetWebsiteSnippetFilter(id int64) (*WebsiteSnippetFilter, error) {
	wsfs, err := c.GetWebsiteSnippetFilters([]int64{id})
	if err != nil {
		return nil, err
	}
	if wsfs != nil && len(*wsfs) > 0 {
		return &((*wsfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.snippet.filter not found", id)
}

// GetWebsiteSnippetFilters gets website.snippet.filter existing records.
func (c *Client) GetWebsiteSnippetFilters(ids []int64) (*WebsiteSnippetFilters, error) {
	wsfs := &WebsiteSnippetFilters{}
	if err := c.Read(WebsiteSnippetFilterModel, ids, nil, wsfs); err != nil {
		return nil, err
	}
	return wsfs, nil
}

// FindWebsiteSnippetFilter finds website.snippet.filter record by querying it with criteria.
func (c *Client) FindWebsiteSnippetFilter(criteria *Criteria) (*WebsiteSnippetFilter, error) {
	wsfs := &WebsiteSnippetFilters{}
	if err := c.SearchRead(WebsiteSnippetFilterModel, criteria, NewOptions().Limit(1), wsfs); err != nil {
		return nil, err
	}
	if wsfs != nil && len(*wsfs) > 0 {
		return &((*wsfs)[0]), nil
	}
	return nil, fmt.Errorf("website.snippet.filter was not found with criteria %v", criteria)
}

// FindWebsiteSnippetFilters finds website.snippet.filter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSnippetFilters(criteria *Criteria, options *Options) (*WebsiteSnippetFilters, error) {
	wsfs := &WebsiteSnippetFilters{}
	if err := c.SearchRead(WebsiteSnippetFilterModel, criteria, options, wsfs); err != nil {
		return nil, err
	}
	return wsfs, nil
}

// FindWebsiteSnippetFilterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSnippetFilterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteSnippetFilterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteSnippetFilterId finds record id by querying it with criteria.
func (c *Client) FindWebsiteSnippetFilterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteSnippetFilterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.snippet.filter was not found with criteria %v and options %v", criteria, options)
}
