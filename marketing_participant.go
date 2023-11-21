package odoo

import (
	"fmt"
)

// MarketingParticipant represents marketing.participant model.
type MarketingParticipant struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CampaignId  *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	IsTest      *Bool      `xmlrpc:"is_test,omptempty"`
	ModelId     *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName   *String    `xmlrpc:"model_name,omptempty"`
	ResId       *Int       `xmlrpc:"res_id,omptempty"`
	ResourceRef *String    `xmlrpc:"resource_ref,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	TraceIds    *Relation  `xmlrpc:"trace_ids,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MarketingParticipants represents array of marketing.participant model.
type MarketingParticipants []MarketingParticipant

// MarketingParticipantModel is the odoo model name.
const MarketingParticipantModel = "marketing.participant"

// Many2One convert MarketingParticipant to *Many2One.
func (mp *MarketingParticipant) Many2One() *Many2One {
	return NewMany2One(mp.Id.Get(), "")
}

// CreateMarketingParticipant creates a new marketing.participant model and returns its id.
func (c *Client) CreateMarketingParticipant(mp *MarketingParticipant) (int64, error) {
	ids, err := c.CreateMarketingParticipants([]*MarketingParticipant{mp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingParticipant creates a new marketing.participant model and returns its id.
func (c *Client) CreateMarketingParticipants(mps []*MarketingParticipant) ([]int64, error) {
	var vv []interface{}
	for _, v := range mps {
		vv = append(vv, v)
	}
	return c.Create(MarketingParticipantModel, vv)
}

// UpdateMarketingParticipant updates an existing marketing.participant record.
func (c *Client) UpdateMarketingParticipant(mp *MarketingParticipant) error {
	return c.UpdateMarketingParticipants([]int64{mp.Id.Get()}, mp)
}

// UpdateMarketingParticipants updates existing marketing.participant records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMarketingParticipants(ids []int64, mp *MarketingParticipant) error {
	return c.Update(MarketingParticipantModel, ids, mp)
}

// DeleteMarketingParticipant deletes an existing marketing.participant record.
func (c *Client) DeleteMarketingParticipant(id int64) error {
	return c.DeleteMarketingParticipants([]int64{id})
}

// DeleteMarketingParticipants deletes existing marketing.participant records.
func (c *Client) DeleteMarketingParticipants(ids []int64) error {
	return c.Delete(MarketingParticipantModel, ids)
}

// GetMarketingParticipant gets marketing.participant existing record.
func (c *Client) GetMarketingParticipant(id int64) (*MarketingParticipant, error) {
	mps, err := c.GetMarketingParticipants([]int64{id})
	if err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of marketing.participant not found", id)
}

// GetMarketingParticipants gets marketing.participant existing records.
func (c *Client) GetMarketingParticipants(ids []int64) (*MarketingParticipants, error) {
	mps := &MarketingParticipants{}
	if err := c.Read(MarketingParticipantModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMarketingParticipant finds marketing.participant record by querying it with criteria.
func (c *Client) FindMarketingParticipant(criteria *Criteria) (*MarketingParticipant, error) {
	mps := &MarketingParticipants{}
	if err := c.SearchRead(MarketingParticipantModel, criteria, NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("marketing.participant was not found with criteria %v", criteria)
}

// FindMarketingParticipants finds marketing.participant records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingParticipants(criteria *Criteria, options *Options) (*MarketingParticipants, error) {
	mps := &MarketingParticipants{}
	if err := c.SearchRead(MarketingParticipantModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMarketingParticipantIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingParticipantIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MarketingParticipantModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMarketingParticipantId finds record id by querying it with criteria.
func (c *Client) FindMarketingParticipantId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingParticipantModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("marketing.participant was not found with criteria %v and options %v", criteria, options)
}
