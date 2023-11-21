package odoo

import (
	"fmt"
)

// WebsiteVisitorPushSubscription represents website.visitor.push.subscription model.
type WebsiteVisitorPushSubscription struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	PushToken        *String   `xmlrpc:"push_token,omptempty"`
	WebsiteVisitorId *Many2One `xmlrpc:"website_visitor_id,omptempty"`
}

// WebsiteVisitorPushSubscriptions represents array of website.visitor.push.subscription model.
type WebsiteVisitorPushSubscriptions []WebsiteVisitorPushSubscription

// WebsiteVisitorPushSubscriptionModel is the odoo model name.
const WebsiteVisitorPushSubscriptionModel = "website.visitor.push.subscription"

// Many2One convert WebsiteVisitorPushSubscription to *Many2One.
func (wvps *WebsiteVisitorPushSubscription) Many2One() *Many2One {
	return NewMany2One(wvps.Id.Get(), "")
}

// CreateWebsiteVisitorPushSubscription creates a new website.visitor.push.subscription model and returns its id.
func (c *Client) CreateWebsiteVisitorPushSubscription(wvps *WebsiteVisitorPushSubscription) (int64, error) {
	ids, err := c.CreateWebsiteVisitorPushSubscriptions([]*WebsiteVisitorPushSubscription{wvps})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteVisitorPushSubscription creates a new website.visitor.push.subscription model and returns its id.
func (c *Client) CreateWebsiteVisitorPushSubscriptions(wvpss []*WebsiteVisitorPushSubscription) ([]int64, error) {
	var vv []interface{}
	for _, v := range wvpss {
		vv = append(vv, v)
	}
	return c.Create(WebsiteVisitorPushSubscriptionModel, vv)
}

// UpdateWebsiteVisitorPushSubscription updates an existing website.visitor.push.subscription record.
func (c *Client) UpdateWebsiteVisitorPushSubscription(wvps *WebsiteVisitorPushSubscription) error {
	return c.UpdateWebsiteVisitorPushSubscriptions([]int64{wvps.Id.Get()}, wvps)
}

// UpdateWebsiteVisitorPushSubscriptions updates existing website.visitor.push.subscription records.
// All records (represented by ids) will be updated by wvps values.
func (c *Client) UpdateWebsiteVisitorPushSubscriptions(ids []int64, wvps *WebsiteVisitorPushSubscription) error {
	return c.Update(WebsiteVisitorPushSubscriptionModel, ids, wvps)
}

// DeleteWebsiteVisitorPushSubscription deletes an existing website.visitor.push.subscription record.
func (c *Client) DeleteWebsiteVisitorPushSubscription(id int64) error {
	return c.DeleteWebsiteVisitorPushSubscriptions([]int64{id})
}

// DeleteWebsiteVisitorPushSubscriptions deletes existing website.visitor.push.subscription records.
func (c *Client) DeleteWebsiteVisitorPushSubscriptions(ids []int64) error {
	return c.Delete(WebsiteVisitorPushSubscriptionModel, ids)
}

// GetWebsiteVisitorPushSubscription gets website.visitor.push.subscription existing record.
func (c *Client) GetWebsiteVisitorPushSubscription(id int64) (*WebsiteVisitorPushSubscription, error) {
	wvpss, err := c.GetWebsiteVisitorPushSubscriptions([]int64{id})
	if err != nil {
		return nil, err
	}
	if wvpss != nil && len(*wvpss) > 0 {
		return &((*wvpss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.visitor.push.subscription not found", id)
}

// GetWebsiteVisitorPushSubscriptions gets website.visitor.push.subscription existing records.
func (c *Client) GetWebsiteVisitorPushSubscriptions(ids []int64) (*WebsiteVisitorPushSubscriptions, error) {
	wvpss := &WebsiteVisitorPushSubscriptions{}
	if err := c.Read(WebsiteVisitorPushSubscriptionModel, ids, nil, wvpss); err != nil {
		return nil, err
	}
	return wvpss, nil
}

// FindWebsiteVisitorPushSubscription finds website.visitor.push.subscription record by querying it with criteria.
func (c *Client) FindWebsiteVisitorPushSubscription(criteria *Criteria) (*WebsiteVisitorPushSubscription, error) {
	wvpss := &WebsiteVisitorPushSubscriptions{}
	if err := c.SearchRead(WebsiteVisitorPushSubscriptionModel, criteria, NewOptions().Limit(1), wvpss); err != nil {
		return nil, err
	}
	if wvpss != nil && len(*wvpss) > 0 {
		return &((*wvpss)[0]), nil
	}
	return nil, fmt.Errorf("website.visitor.push.subscription was not found with criteria %v", criteria)
}

// FindWebsiteVisitorPushSubscriptions finds website.visitor.push.subscription records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteVisitorPushSubscriptions(criteria *Criteria, options *Options) (*WebsiteVisitorPushSubscriptions, error) {
	wvpss := &WebsiteVisitorPushSubscriptions{}
	if err := c.SearchRead(WebsiteVisitorPushSubscriptionModel, criteria, options, wvpss); err != nil {
		return nil, err
	}
	return wvpss, nil
}

// FindWebsiteVisitorPushSubscriptionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteVisitorPushSubscriptionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteVisitorPushSubscriptionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteVisitorPushSubscriptionId finds record id by querying it with criteria.
func (c *Client) FindWebsiteVisitorPushSubscriptionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteVisitorPushSubscriptionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.visitor.push.subscription was not found with criteria %v and options %v", criteria, options)
}
