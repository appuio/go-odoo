package odoo

import (
	"fmt"
)

// SocialPost represents social.post model.
type SocialPost struct {
	LastUpdate                        *Time      `xmlrpc:"__last_update,omptempty"`
	AccountAllowedIds                 *Relation  `xmlrpc:"account_allowed_ids,omptempty"`
	AccountIds                        *Relation  `xmlrpc:"account_ids,omptempty"`
	ActivityCalendarEventId           *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline              *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration       *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon             *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                       *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                     *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                   *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                  *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                    *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                    *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CalendarDate                      *Time      `xmlrpc:"calendar_date,omptempty"`
	ClickCount                        *Int       `xmlrpc:"click_count,omptempty"`
	CompanyId                         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayFacebookPreview            *Bool      `xmlrpc:"display_facebook_preview,omptempty"`
	DisplayInstagramPreview           *Bool      `xmlrpc:"display_instagram_preview,omptempty"`
	DisplayLinkedinPreview            *Bool      `xmlrpc:"display_linkedin_preview,omptempty"`
	DisplayName                       *String    `xmlrpc:"display_name,omptempty"`
	DisplayPushNotificationAttributes *Bool      `xmlrpc:"display_push_notification_attributes,omptempty"`
	DisplayPushNotificationsPreview   *Bool      `xmlrpc:"display_push_notifications_preview,omptempty"`
	Engagement                        *Int       `xmlrpc:"engagement,omptempty"`
	FacebookPreview                   *String    `xmlrpc:"facebook_preview,omptempty"`
	FailedMessageIds                  *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasActiveAccounts                 *Bool      `xmlrpc:"has_active_accounts,omptempty"`
	HasMessage                        *Bool      `xmlrpc:"has_message,omptempty"`
	HasPostErrors                     *Bool      `xmlrpc:"has_post_errors,omptempty"`
	HasTwitterAccounts                *Bool      `xmlrpc:"has_twitter_accounts,omptempty"`
	Id                                *Int       `xmlrpc:"id,omptempty"`
	ImageIds                          *Relation  `xmlrpc:"image_ids,omptempty"`
	ImageUrls                         *String    `xmlrpc:"image_urls,omptempty"`
	InstagramAccessToken              *String    `xmlrpc:"instagram_access_token,omptempty"`
	InstagramImageId                  *Many2One  `xmlrpc:"instagram_image_id,omptempty"`
	InstagramPreview                  *String    `xmlrpc:"instagram_preview,omptempty"`
	IsTwitterPostLimitExceed          *Bool      `xmlrpc:"is_twitter_post_limit_exceed,omptempty"`
	LeadsOpportunitiesCount           *Int       `xmlrpc:"leads_opportunities_count,omptempty"`
	LinkedinPreview                   *String    `xmlrpc:"linkedin_preview,omptempty"`
	LivePostIds                       *Relation  `xmlrpc:"live_post_ids,omptempty"`
	LivePostsByMedia                  *String    `xmlrpc:"live_posts_by_media,omptempty"`
	MediaIds                          *Relation  `xmlrpc:"media_ids,omptempty"`
	Message                           *String    `xmlrpc:"message,omptempty"`
	MessageAttachmentCount            *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                    *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                   *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter            *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                        *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                 *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLength                     *Int       `xmlrpc:"message_length,omptempty"`
	MessageMainAttachmentId           *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                 *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter          *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                 *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline            *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                              *String    `xmlrpc:"name,omptempty"`
	PostMethod                        *Selection `xmlrpc:"post_method,omptempty"`
	PublishedDate                     *Time      `xmlrpc:"published_date,omptempty"`
	PushNotificationImage             *String    `xmlrpc:"push_notification_image,omptempty"`
	PushNotificationTargetUrl         *String    `xmlrpc:"push_notification_target_url,omptempty"`
	PushNotificationTitle             *String    `xmlrpc:"push_notification_title,omptempty"`
	PushNotificationsPreview          *String    `xmlrpc:"push_notifications_preview,omptempty"`
	SaleInvoicedAmount                *Int       `xmlrpc:"sale_invoiced_amount,omptempty"`
	SaleQuotationCount                *Int       `xmlrpc:"sale_quotation_count,omptempty"`
	ScheduledDate                     *Time      `xmlrpc:"scheduled_date,omptempty"`
	SourceId                          *Many2One  `xmlrpc:"source_id,omptempty"`
	State                             *Selection `xmlrpc:"state,omptempty"`
	StreamPostsCount                  *Int       `xmlrpc:"stream_posts_count,omptempty"`
	TwitterPostLimitMessage           *String    `xmlrpc:"twitter_post_limit_message,omptempty"`
	TwitterPreview                    *String    `xmlrpc:"twitter_preview,omptempty"`
	UseLeads                          *Bool      `xmlrpc:"use_leads,omptempty"`
	UseVisitorTimezone                *Bool      `xmlrpc:"use_visitor_timezone,omptempty"`
	UtmCampaignId                     *Many2One  `xmlrpc:"utm_campaign_id,omptempty"`
	VisitorDomain                     *String    `xmlrpc:"visitor_domain,omptempty"`
	WebsiteMessageIds                 *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                          *Many2One  `xmlrpc:"write_uid,omptempty"`
	YoutubeAccessToken                *String    `xmlrpc:"youtube_access_token,omptempty"`
	YoutubeAccountsCount              *Int       `xmlrpc:"youtube_accounts_count,omptempty"`
	YoutubeAccountsOtherCount         *Int       `xmlrpc:"youtube_accounts_other_count,omptempty"`
	YoutubeDescription                *String    `xmlrpc:"youtube_description,omptempty"`
	YoutubePreview                    *String    `xmlrpc:"youtube_preview,omptempty"`
	YoutubeThumbnailUrl               *String    `xmlrpc:"youtube_thumbnail_url,omptempty"`
	YoutubeTitle                      *String    `xmlrpc:"youtube_title,omptempty"`
	YoutubeVideo                      *String    `xmlrpc:"youtube_video,omptempty"`
	YoutubeVideoCategoryId            *String    `xmlrpc:"youtube_video_category_id,omptempty"`
	YoutubeVideoId                    *String    `xmlrpc:"youtube_video_id,omptempty"`
	YoutubeVideoPrivacy               *Selection `xmlrpc:"youtube_video_privacy,omptempty"`
	YoutubeVideoUrl                   *String    `xmlrpc:"youtube_video_url,omptempty"`
}

// SocialPosts represents array of social.post model.
type SocialPosts []SocialPost

// SocialPostModel is the odoo model name.
const SocialPostModel = "social.post"

// Many2One convert SocialPost to *Many2One.
func (sp *SocialPost) Many2One() *Many2One {
	return NewMany2One(sp.Id.Get(), "")
}

// CreateSocialPost creates a new social.post model and returns its id.
func (c *Client) CreateSocialPost(sp *SocialPost) (int64, error) {
	ids, err := c.CreateSocialPosts([]*SocialPost{sp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialPost creates a new social.post model and returns its id.
func (c *Client) CreateSocialPosts(sps []*SocialPost) ([]int64, error) {
	var vv []interface{}
	for _, v := range sps {
		vv = append(vv, v)
	}
	return c.Create(SocialPostModel, vv)
}

// UpdateSocialPost updates an existing social.post record.
func (c *Client) UpdateSocialPost(sp *SocialPost) error {
	return c.UpdateSocialPosts([]int64{sp.Id.Get()}, sp)
}

// UpdateSocialPosts updates existing social.post records.
// All records (represented by ids) will be updated by sp values.
func (c *Client) UpdateSocialPosts(ids []int64, sp *SocialPost) error {
	return c.Update(SocialPostModel, ids, sp)
}

// DeleteSocialPost deletes an existing social.post record.
func (c *Client) DeleteSocialPost(id int64) error {
	return c.DeleteSocialPosts([]int64{id})
}

// DeleteSocialPosts deletes existing social.post records.
func (c *Client) DeleteSocialPosts(ids []int64) error {
	return c.Delete(SocialPostModel, ids)
}

// GetSocialPost gets social.post existing record.
func (c *Client) GetSocialPost(id int64) (*SocialPost, error) {
	sps, err := c.GetSocialPosts([]int64{id})
	if err != nil {
		return nil, err
	}
	if sps != nil && len(*sps) > 0 {
		return &((*sps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.post not found", id)
}

// GetSocialPosts gets social.post existing records.
func (c *Client) GetSocialPosts(ids []int64) (*SocialPosts, error) {
	sps := &SocialPosts{}
	if err := c.Read(SocialPostModel, ids, nil, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindSocialPost finds social.post record by querying it with criteria.
func (c *Client) FindSocialPost(criteria *Criteria) (*SocialPost, error) {
	sps := &SocialPosts{}
	if err := c.SearchRead(SocialPostModel, criteria, NewOptions().Limit(1), sps); err != nil {
		return nil, err
	}
	if sps != nil && len(*sps) > 0 {
		return &((*sps)[0]), nil
	}
	return nil, fmt.Errorf("social.post was not found with criteria %v", criteria)
}

// FindSocialPosts finds social.post records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialPosts(criteria *Criteria, options *Options) (*SocialPosts, error) {
	sps := &SocialPosts{}
	if err := c.SearchRead(SocialPostModel, criteria, options, sps); err != nil {
		return nil, err
	}
	return sps, nil
}

// FindSocialPostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialPostIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialPostModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialPostId finds record id by querying it with criteria.
func (c *Client) FindSocialPostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialPostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.post was not found with criteria %v and options %v", criteria, options)
}
