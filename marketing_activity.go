package odoo

import (
	"fmt"
)

// MarketingActivity represents marketing.activity model.
type MarketingActivity struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDomain           *String    `xmlrpc:"activity_domain,omptempty"`
	ActivityType             *Selection `xmlrpc:"activity_type,omptempty"`
	AllowedParentIds         *Relation  `xmlrpc:"allowed_parent_ids,omptempty"`
	CampaignId               *Many2One  `xmlrpc:"campaign_id,omptempty"`
	ChildIds                 *Relation  `xmlrpc:"child_ids,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	Domain                   *String    `xmlrpc:"domain,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IntervalNumber           *Int       `xmlrpc:"interval_number,omptempty"`
	IntervalStandardized     *Int       `xmlrpc:"interval_standardized,omptempty"`
	IntervalType             *Selection `xmlrpc:"interval_type,omptempty"`
	MassMailingId            *Many2One  `xmlrpc:"mass_mailing_id,omptempty"`
	MassMailingIdMailingType *Selection `xmlrpc:"mass_mailing_id_mailing_type,omptempty"`
	ModelId                  *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName                *String    `xmlrpc:"model_name,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	ParentId                 *Many2One  `xmlrpc:"parent_id,omptempty"`
	Processed                *Int       `xmlrpc:"processed,omptempty"`
	Rejected                 *Int       `xmlrpc:"rejected,omptempty"`
	RequireSync              *Bool      `xmlrpc:"require_sync,omptempty"`
	ServerActionId           *Many2One  `xmlrpc:"server_action_id,omptempty"`
	SourceId                 *Many2One  `xmlrpc:"source_id,omptempty"`
	StatisticsGraphData      *String    `xmlrpc:"statistics_graph_data,omptempty"`
	TotalBounce              *Int       `xmlrpc:"total_bounce,omptempty"`
	TotalClick               *Int       `xmlrpc:"total_click,omptempty"`
	TotalOpen                *Int       `xmlrpc:"total_open,omptempty"`
	TotalReply               *Int       `xmlrpc:"total_reply,omptempty"`
	TotalSent                *Int       `xmlrpc:"total_sent,omptempty"`
	TraceIds                 *Relation  `xmlrpc:"trace_ids,omptempty"`
	TriggerCategory          *Selection `xmlrpc:"trigger_category,omptempty"`
	TriggerType              *Selection `xmlrpc:"trigger_type,omptempty"`
	UtmCampaignId            *Many2One  `xmlrpc:"utm_campaign_id,omptempty"`
	ValidityDuration         *Bool      `xmlrpc:"validity_duration,omptempty"`
	ValidityDurationNumber   *Int       `xmlrpc:"validity_duration_number,omptempty"`
	ValidityDurationType     *Selection `xmlrpc:"validity_duration_type,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MarketingActivitys represents array of marketing.activity model.
type MarketingActivitys []MarketingActivity

// MarketingActivityModel is the odoo model name.
const MarketingActivityModel = "marketing.activity"

// Many2One convert MarketingActivity to *Many2One.
func (ma *MarketingActivity) Many2One() *Many2One {
	return NewMany2One(ma.Id.Get(), "")
}

// CreateMarketingActivity creates a new marketing.activity model and returns its id.
func (c *Client) CreateMarketingActivity(ma *MarketingActivity) (int64, error) {
	ids, err := c.CreateMarketingActivitys([]*MarketingActivity{ma})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingActivity creates a new marketing.activity model and returns its id.
func (c *Client) CreateMarketingActivitys(mas []*MarketingActivity) ([]int64, error) {
	var vv []interface{}
	for _, v := range mas {
		vv = append(vv, v)
	}
	return c.Create(MarketingActivityModel, vv)
}

// UpdateMarketingActivity updates an existing marketing.activity record.
func (c *Client) UpdateMarketingActivity(ma *MarketingActivity) error {
	return c.UpdateMarketingActivitys([]int64{ma.Id.Get()}, ma)
}

// UpdateMarketingActivitys updates existing marketing.activity records.
// All records (represented by ids) will be updated by ma values.
func (c *Client) UpdateMarketingActivitys(ids []int64, ma *MarketingActivity) error {
	return c.Update(MarketingActivityModel, ids, ma)
}

// DeleteMarketingActivity deletes an existing marketing.activity record.
func (c *Client) DeleteMarketingActivity(id int64) error {
	return c.DeleteMarketingActivitys([]int64{id})
}

// DeleteMarketingActivitys deletes existing marketing.activity records.
func (c *Client) DeleteMarketingActivitys(ids []int64) error {
	return c.Delete(MarketingActivityModel, ids)
}

// GetMarketingActivity gets marketing.activity existing record.
func (c *Client) GetMarketingActivity(id int64) (*MarketingActivity, error) {
	mas, err := c.GetMarketingActivitys([]int64{id})
	if err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of marketing.activity not found", id)
}

// GetMarketingActivitys gets marketing.activity existing records.
func (c *Client) GetMarketingActivitys(ids []int64) (*MarketingActivitys, error) {
	mas := &MarketingActivitys{}
	if err := c.Read(MarketingActivityModel, ids, nil, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMarketingActivity finds marketing.activity record by querying it with criteria.
func (c *Client) FindMarketingActivity(criteria *Criteria) (*MarketingActivity, error) {
	mas := &MarketingActivitys{}
	if err := c.SearchRead(MarketingActivityModel, criteria, NewOptions().Limit(1), mas); err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("marketing.activity was not found with criteria %v", criteria)
}

// FindMarketingActivitys finds marketing.activity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingActivitys(criteria *Criteria, options *Options) (*MarketingActivitys, error) {
	mas := &MarketingActivitys{}
	if err := c.SearchRead(MarketingActivityModel, criteria, options, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMarketingActivityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingActivityIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MarketingActivityModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMarketingActivityId finds record id by querying it with criteria.
func (c *Client) FindMarketingActivityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingActivityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("marketing.activity was not found with criteria %v and options %v", criteria, options)
}
