package odoo

import (
	"fmt"
)

// WebsiteEventMenu represents website.event.menu model.
type WebsiteEventMenu struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	EventId     *Many2One  `xmlrpc:"event_id,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	MenuId      *Many2One  `xmlrpc:"menu_id,omptempty"`
	MenuType    *Selection `xmlrpc:"menu_type,omptempty"`
	ViewId      *Many2One  `xmlrpc:"view_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// WebsiteEventMenus represents array of website.event.menu model.
type WebsiteEventMenus []WebsiteEventMenu

// WebsiteEventMenuModel is the odoo model name.
const WebsiteEventMenuModel = "website.event.menu"

// Many2One convert WebsiteEventMenu to *Many2One.
func (wem *WebsiteEventMenu) Many2One() *Many2One {
	return NewMany2One(wem.Id.Get(), "")
}

// CreateWebsiteEventMenu creates a new website.event.menu model and returns its id.
func (c *Client) CreateWebsiteEventMenu(wem *WebsiteEventMenu) (int64, error) {
	ids, err := c.CreateWebsiteEventMenus([]*WebsiteEventMenu{wem})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteEventMenu creates a new website.event.menu model and returns its id.
func (c *Client) CreateWebsiteEventMenus(wems []*WebsiteEventMenu) ([]int64, error) {
	var vv []interface{}
	for _, v := range wems {
		vv = append(vv, v)
	}
	return c.Create(WebsiteEventMenuModel, vv)
}

// UpdateWebsiteEventMenu updates an existing website.event.menu record.
func (c *Client) UpdateWebsiteEventMenu(wem *WebsiteEventMenu) error {
	return c.UpdateWebsiteEventMenus([]int64{wem.Id.Get()}, wem)
}

// UpdateWebsiteEventMenus updates existing website.event.menu records.
// All records (represented by ids) will be updated by wem values.
func (c *Client) UpdateWebsiteEventMenus(ids []int64, wem *WebsiteEventMenu) error {
	return c.Update(WebsiteEventMenuModel, ids, wem)
}

// DeleteWebsiteEventMenu deletes an existing website.event.menu record.
func (c *Client) DeleteWebsiteEventMenu(id int64) error {
	return c.DeleteWebsiteEventMenus([]int64{id})
}

// DeleteWebsiteEventMenus deletes existing website.event.menu records.
func (c *Client) DeleteWebsiteEventMenus(ids []int64) error {
	return c.Delete(WebsiteEventMenuModel, ids)
}

// GetWebsiteEventMenu gets website.event.menu existing record.
func (c *Client) GetWebsiteEventMenu(id int64) (*WebsiteEventMenu, error) {
	wems, err := c.GetWebsiteEventMenus([]int64{id})
	if err != nil {
		return nil, err
	}
	if wems != nil && len(*wems) > 0 {
		return &((*wems)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.event.menu not found", id)
}

// GetWebsiteEventMenus gets website.event.menu existing records.
func (c *Client) GetWebsiteEventMenus(ids []int64) (*WebsiteEventMenus, error) {
	wems := &WebsiteEventMenus{}
	if err := c.Read(WebsiteEventMenuModel, ids, nil, wems); err != nil {
		return nil, err
	}
	return wems, nil
}

// FindWebsiteEventMenu finds website.event.menu record by querying it with criteria.
func (c *Client) FindWebsiteEventMenu(criteria *Criteria) (*WebsiteEventMenu, error) {
	wems := &WebsiteEventMenus{}
	if err := c.SearchRead(WebsiteEventMenuModel, criteria, NewOptions().Limit(1), wems); err != nil {
		return nil, err
	}
	if wems != nil && len(*wems) > 0 {
		return &((*wems)[0]), nil
	}
	return nil, fmt.Errorf("website.event.menu was not found with criteria %v", criteria)
}

// FindWebsiteEventMenus finds website.event.menu records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteEventMenus(criteria *Criteria, options *Options) (*WebsiteEventMenus, error) {
	wems := &WebsiteEventMenus{}
	if err := c.SearchRead(WebsiteEventMenuModel, criteria, options, wems); err != nil {
		return nil, err
	}
	return wems, nil
}

// FindWebsiteEventMenuIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteEventMenuIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteEventMenuModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteEventMenuId finds record id by querying it with criteria.
func (c *Client) FindWebsiteEventMenuId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteEventMenuModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.event.menu was not found with criteria %v and options %v", criteria, options)
}
