package odoo

import (
	"fmt"
)

// MarketingTrace represents marketing.trace model.
type MarketingTrace struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityId         *Many2One  `xmlrpc:"activity_id,omptempty"`
	ActivityType       *Selection `xmlrpc:"activity_type,omptempty"`
	ChildIds           *Relation  `xmlrpc:"child_ids,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	IsTest             *Bool      `xmlrpc:"is_test,omptempty"`
	LinksClickDatetime *Time      `xmlrpc:"links_click_datetime,omptempty"`
	MailingTraceIds    *Relation  `xmlrpc:"mailing_trace_ids,omptempty"`
	MailingTraceStatus *Selection `xmlrpc:"mailing_trace_status,omptempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParticipantId      *Many2One  `xmlrpc:"participant_id,omptempty"`
	ResId              *Int       `xmlrpc:"res_id,omptempty"`
	ScheduleDate       *Time      `xmlrpc:"schedule_date,omptempty"`
	State              *Selection `xmlrpc:"state,omptempty"`
	StateMsg           *String    `xmlrpc:"state_msg,omptempty"`
	TriggerType        *Selection `xmlrpc:"trigger_type,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MarketingTraces represents array of marketing.trace model.
type MarketingTraces []MarketingTrace

// MarketingTraceModel is the odoo model name.
const MarketingTraceModel = "marketing.trace"

// Many2One convert MarketingTrace to *Many2One.
func (mt *MarketingTrace) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMarketingTrace creates a new marketing.trace model and returns its id.
func (c *Client) CreateMarketingTrace(mt *MarketingTrace) (int64, error) {
	ids, err := c.CreateMarketingTraces([]*MarketingTrace{mt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingTrace creates a new marketing.trace model and returns its id.
func (c *Client) CreateMarketingTraces(mts []*MarketingTrace) ([]int64, error) {
	var vv []interface{}
	for _, v := range mts {
		vv = append(vv, v)
	}
	return c.Create(MarketingTraceModel, vv)
}

// UpdateMarketingTrace updates an existing marketing.trace record.
func (c *Client) UpdateMarketingTrace(mt *MarketingTrace) error {
	return c.UpdateMarketingTraces([]int64{mt.Id.Get()}, mt)
}

// UpdateMarketingTraces updates existing marketing.trace records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMarketingTraces(ids []int64, mt *MarketingTrace) error {
	return c.Update(MarketingTraceModel, ids, mt)
}

// DeleteMarketingTrace deletes an existing marketing.trace record.
func (c *Client) DeleteMarketingTrace(id int64) error {
	return c.DeleteMarketingTraces([]int64{id})
}

// DeleteMarketingTraces deletes existing marketing.trace records.
func (c *Client) DeleteMarketingTraces(ids []int64) error {
	return c.Delete(MarketingTraceModel, ids)
}

// GetMarketingTrace gets marketing.trace existing record.
func (c *Client) GetMarketingTrace(id int64) (*MarketingTrace, error) {
	mts, err := c.GetMarketingTraces([]int64{id})
	if err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of marketing.trace not found", id)
}

// GetMarketingTraces gets marketing.trace existing records.
func (c *Client) GetMarketingTraces(ids []int64) (*MarketingTraces, error) {
	mts := &MarketingTraces{}
	if err := c.Read(MarketingTraceModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMarketingTrace finds marketing.trace record by querying it with criteria.
func (c *Client) FindMarketingTrace(criteria *Criteria) (*MarketingTrace, error) {
	mts := &MarketingTraces{}
	if err := c.SearchRead(MarketingTraceModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("marketing.trace was not found with criteria %v", criteria)
}

// FindMarketingTraces finds marketing.trace records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingTraces(criteria *Criteria, options *Options) (*MarketingTraces, error) {
	mts := &MarketingTraces{}
	if err := c.SearchRead(MarketingTraceModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMarketingTraceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingTraceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MarketingTraceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMarketingTraceId finds record id by querying it with criteria.
func (c *Client) FindMarketingTraceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingTraceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("marketing.trace was not found with criteria %v and options %v", criteria, options)
}
