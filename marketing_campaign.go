package odoo

import (
	"fmt"
)

// MarketingCampaign represents marketing.campaign model.
type MarketingCampaign struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	AbTestingCompleted           *Bool      `xmlrpc:"ab_testing_completed,omptempty"`
	AbTestingMailingsCount       *Int       `xmlrpc:"ab_testing_mailings_count,omptempty"`
	AbTestingScheduleDatetime    *Time      `xmlrpc:"ab_testing_schedule_datetime,omptempty"`
	AbTestingTotalPc             *Int       `xmlrpc:"ab_testing_total_pc,omptempty"`
	AbTestingWinnerSelection     *Selection `xmlrpc:"ab_testing_winner_selection,omptempty"`
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	BouncedRatio                 *Int       `xmlrpc:"bounced_ratio,omptempty"`
	ClickCount                   *Int       `xmlrpc:"click_count,omptempty"`
	Color                        *Int       `xmlrpc:"color,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CompletedParticipantCount    *Int       `xmlrpc:"completed_participant_count,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrmLeadCount                 *Int       `xmlrpc:"crm_lead_count,omptempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	Domain                       *String    `xmlrpc:"domain,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	InvoicedAmount               *Int       `xmlrpc:"invoiced_amount,omptempty"`
	IsAutoCampaign               *Bool      `xmlrpc:"is_auto_campaign,omptempty"`
	IsMailingCampaignActivated   *Bool      `xmlrpc:"is_mailing_campaign_activated,omptempty"`
	LastSyncDate                 *Time      `xmlrpc:"last_sync_date,omptempty"`
	LinkTrackerClickCount        *Int       `xmlrpc:"link_tracker_click_count,omptempty"`
	MailingFilterCount           *Int       `xmlrpc:"mailing_filter_count,omptempty"`
	MailingFilterDomain          *String    `xmlrpc:"mailing_filter_domain,omptempty"`
	MailingFilterId              *Many2One  `xmlrpc:"mailing_filter_id,omptempty"`
	MailingMailCount             *Int       `xmlrpc:"mailing_mail_count,omptempty"`
	MailingMailIds               *Relation  `xmlrpc:"mailing_mail_ids,omptempty"`
	MarketingActivityIds         *Relation  `xmlrpc:"marketing_activity_ids,omptempty"`
	MassMailingCount             *Int       `xmlrpc:"mass_mailing_count,omptempty"`
	ModelId                      *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName                    *String    `xmlrpc:"model_name,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	OpenedRatio                  *Int       `xmlrpc:"opened_ratio,omptempty"`
	ParticipantIds               *Relation  `xmlrpc:"participant_ids,omptempty"`
	QuotationCount               *Int       `xmlrpc:"quotation_count,omptempty"`
	ReceivedRatio                *Int       `xmlrpc:"received_ratio,omptempty"`
	RepliedRatio                 *Int       `xmlrpc:"replied_ratio,omptempty"`
	RequireSync                  *Bool      `xmlrpc:"require_sync,omptempty"`
	RunningParticipantCount      *Int       `xmlrpc:"running_participant_count,omptempty"`
	SocialEngagement             *Int       `xmlrpc:"social_engagement,omptempty"`
	SocialPostIds                *Relation  `xmlrpc:"social_post_ids,omptempty"`
	SocialPostsCount             *Int       `xmlrpc:"social_posts_count,omptempty"`
	SocialPushNotificationIds    *Relation  `xmlrpc:"social_push_notification_ids,omptempty"`
	SocialPushNotificationsCount *Int       `xmlrpc:"social_push_notifications_count,omptempty"`
	StageId                      *Many2One  `xmlrpc:"stage_id,omptempty"`
	State                        *Selection `xmlrpc:"state,omptempty"`
	TagIds                       *Relation  `xmlrpc:"tag_ids,omptempty"`
	TestParticipantCount         *Int       `xmlrpc:"test_participant_count,omptempty"`
	Title                        *String    `xmlrpc:"title,omptempty"`
	TotalParticipantCount        *Int       `xmlrpc:"total_participant_count,omptempty"`
	UniqueFieldId                *Many2One  `xmlrpc:"unique_field_id,omptempty"`
	UseLeads                     *Bool      `xmlrpc:"use_leads,omptempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omptempty"`
	UtmCampaignId                *Many2One  `xmlrpc:"utm_campaign_id,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MarketingCampaigns represents array of marketing.campaign model.
type MarketingCampaigns []MarketingCampaign

// MarketingCampaignModel is the odoo model name.
const MarketingCampaignModel = "marketing.campaign"

// Many2One convert MarketingCampaign to *Many2One.
func (mc *MarketingCampaign) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMarketingCampaign creates a new marketing.campaign model and returns its id.
func (c *Client) CreateMarketingCampaign(mc *MarketingCampaign) (int64, error) {
	ids, err := c.CreateMarketingCampaigns([]*MarketingCampaign{mc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMarketingCampaign creates a new marketing.campaign model and returns its id.
func (c *Client) CreateMarketingCampaigns(mcs []*MarketingCampaign) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcs {
		vv = append(vv, v)
	}
	return c.Create(MarketingCampaignModel, vv)
}

// UpdateMarketingCampaign updates an existing marketing.campaign record.
func (c *Client) UpdateMarketingCampaign(mc *MarketingCampaign) error {
	return c.UpdateMarketingCampaigns([]int64{mc.Id.Get()}, mc)
}

// UpdateMarketingCampaigns updates existing marketing.campaign records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMarketingCampaigns(ids []int64, mc *MarketingCampaign) error {
	return c.Update(MarketingCampaignModel, ids, mc)
}

// DeleteMarketingCampaign deletes an existing marketing.campaign record.
func (c *Client) DeleteMarketingCampaign(id int64) error {
	return c.DeleteMarketingCampaigns([]int64{id})
}

// DeleteMarketingCampaigns deletes existing marketing.campaign records.
func (c *Client) DeleteMarketingCampaigns(ids []int64) error {
	return c.Delete(MarketingCampaignModel, ids)
}

// GetMarketingCampaign gets marketing.campaign existing record.
func (c *Client) GetMarketingCampaign(id int64) (*MarketingCampaign, error) {
	mcs, err := c.GetMarketingCampaigns([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of marketing.campaign not found", id)
}

// GetMarketingCampaigns gets marketing.campaign existing records.
func (c *Client) GetMarketingCampaigns(ids []int64) (*MarketingCampaigns, error) {
	mcs := &MarketingCampaigns{}
	if err := c.Read(MarketingCampaignModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMarketingCampaign finds marketing.campaign record by querying it with criteria.
func (c *Client) FindMarketingCampaign(criteria *Criteria) (*MarketingCampaign, error) {
	mcs := &MarketingCampaigns{}
	if err := c.SearchRead(MarketingCampaignModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("marketing.campaign was not found with criteria %v", criteria)
}

// FindMarketingCampaigns finds marketing.campaign records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaigns(criteria *Criteria, options *Options) (*MarketingCampaigns, error) {
	mcs := &MarketingCampaigns{}
	if err := c.SearchRead(MarketingCampaignModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMarketingCampaignIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMarketingCampaignIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MarketingCampaignModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMarketingCampaignId finds record id by querying it with criteria.
func (c *Client) FindMarketingCampaignId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MarketingCampaignModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("marketing.campaign was not found with criteria %v and options %v", criteria, options)
}
