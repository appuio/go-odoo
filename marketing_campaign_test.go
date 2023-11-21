package odoo

import (
	"fmt"
)

// MarketingCampaignTest represents marketing.campaign.test model.
type MarketingCampaignTest struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CampaignId  *Many2One `xmlrpc:"campaign_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	ModelId     *Many2One `xmlrpc:"model_id,omptempty"`
	ModelName   *String   `xmlrpc:"model_name,omptempty"`
	ResId       *Int      `xmlrpc:"res_id,omptempty"`
	ResourceRef *String   `xmlrpc:"resource_ref,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MarketingCampaignTests represents array of marketing.campaign.test model.
type MarketingCampaignTests []MarketingCampaignTest

// MarketingCampaignTestModel is the odoo model name.
const MarketingCampaignTestModel = "marketing.campaign.test"

// Many2One convert MarketingCampaignTest to *Many2One.
func (mct *MarketingCampaignTest) Many2One() *Many2One {
	return NewMany2One(mct.Id.Get(), "")
}

// CreateMarketingCampaignTest creates a new marketing.campaign.test model and returns its id.
func (c *Client) CreateMarketingCampaignTest(mct *MarketingCampaignTest) (int64, error) {
	ids, err := c.CreateMarketingCampaignTests([]*MarketingCampaignTest{mct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingCampaignTest creates a new marketing.campaign.test model and returns its id.
func (c *Client) CreateMarketingCampaignTests(mcts []*MarketingCampaignTest) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcts {
		vv = append(vv, v)
	}
	return c.Create(MarketingCampaignTestModel, vv)
}

// UpdateMarketingCampaignTest updates an existing marketing.campaign.test record.
func (c *Client) UpdateMarketingCampaignTest(mct *MarketingCampaignTest) error {
	return c.UpdateMarketingCampaignTests([]int64{mct.Id.Get()}, mct)
}

// UpdateMarketingCampaignTests updates existing marketing.campaign.test records.
// All records (represented by ids) will be updated by mct values.
func (c *Client) UpdateMarketingCampaignTests(ids []int64, mct *MarketingCampaignTest) error {
	return c.Update(MarketingCampaignTestModel, ids, mct)
}

// DeleteMarketingCampaignTest deletes an existing marketing.campaign.test record.
func (c *Client) DeleteMarketingCampaignTest(id int64) error {
	return c.DeleteMarketingCampaignTests([]int64{id})
}

// DeleteMarketingCampaignTests deletes existing marketing.campaign.test records.
func (c *Client) DeleteMarketingCampaignTests(ids []int64) error {
	return c.Delete(MarketingCampaignTestModel, ids)
}

// GetMarketingCampaignTest gets marketing.campaign.test existing record.
func (c *Client) GetMarketingCampaignTest(id int64) (*MarketingCampaignTest, error) {
	mcts, err := c.GetMarketingCampaignTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcts != nil && len(*mcts) > 0 {
		return &((*mcts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of marketing.campaign.test not found", id)
}

// GetMarketingCampaignTests gets marketing.campaign.test existing records.
func (c *Client) GetMarketingCampaignTests(ids []int64) (*MarketingCampaignTests, error) {
	mcts := &MarketingCampaignTests{}
	if err := c.Read(MarketingCampaignTestModel, ids, nil, mcts); err != nil {
		return nil, err
	}
	return mcts, nil
}

// FindMarketingCampaignTest finds marketing.campaign.test record by querying it with criteria.
func (c *Client) FindMarketingCampaignTest(criteria *Criteria) (*MarketingCampaignTest, error) {
	mcts := &MarketingCampaignTests{}
	if err := c.SearchRead(MarketingCampaignTestModel, criteria, NewOptions().Limit(1), mcts); err != nil {
		return nil, err
	}
	if mcts != nil && len(*mcts) > 0 {
		return &((*mcts)[0]), nil
	}
	return nil, fmt.Errorf("marketing.campaign.test was not found with criteria %v", criteria)
}

// FindMarketingCampaignTests finds marketing.campaign.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaignTests(criteria *Criteria, options *Options) (*MarketingCampaignTests, error) {
	mcts := &MarketingCampaignTests{}
	if err := c.SearchRead(MarketingCampaignTestModel, criteria, options, mcts); err != nil {
		return nil, err
	}
	return mcts, nil
}

// FindMarketingCampaignTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaignTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MarketingCampaignTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMarketingCampaignTestId finds record id by querying it with criteria.
func (c *Client) FindMarketingCampaignTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingCampaignTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("marketing.campaign.test was not found with criteria %v and options %v", criteria, options)
}
