package odoo

import (
	"fmt"
)

// SocialPostTemplate represents social.post.template model.
type SocialPostTemplate struct {
	LastUpdate                        *Time     `xmlrpc:"__last_update,omptempty"`
	AccountIds                        *Relation `xmlrpc:"account_ids,omptempty"`
	CreateDate                        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayFacebookPreview            *Bool     `xmlrpc:"display_facebook_preview,omptempty"`
	DisplayInstagramPreview           *Bool     `xmlrpc:"display_instagram_preview,omptempty"`
	DisplayLinkedinPreview            *Bool     `xmlrpc:"display_linkedin_preview,omptempty"`
	DisplayName                       *String   `xmlrpc:"display_name,omptempty"`
	DisplayPushNotificationAttributes *Bool     `xmlrpc:"display_push_notification_attributes,omptempty"`
	DisplayPushNotificationsPreview   *Bool     `xmlrpc:"display_push_notifications_preview,omptempty"`
	FacebookPreview                   *String   `xmlrpc:"facebook_preview,omptempty"`
	HasActiveAccounts                 *Bool     `xmlrpc:"has_active_accounts,omptempty"`
	HasTwitterAccounts                *Bool     `xmlrpc:"has_twitter_accounts,omptempty"`
	Id                                *Int      `xmlrpc:"id,omptempty"`
	ImageIds                          *Relation `xmlrpc:"image_ids,omptempty"`
	ImageUrls                         *String   `xmlrpc:"image_urls,omptempty"`
	InstagramAccessToken              *String   `xmlrpc:"instagram_access_token,omptempty"`
	InstagramImageId                  *Many2One `xmlrpc:"instagram_image_id,omptempty"`
	InstagramPreview                  *String   `xmlrpc:"instagram_preview,omptempty"`
	IsTwitterPostLimitExceed          *Bool     `xmlrpc:"is_twitter_post_limit_exceed,omptempty"`
	LinkedinPreview                   *String   `xmlrpc:"linkedin_preview,omptempty"`
	Message                           *String   `xmlrpc:"message,omptempty"`
	MessageLength                     *Int      `xmlrpc:"message_length,omptempty"`
	PushNotificationImage             *String   `xmlrpc:"push_notification_image,omptempty"`
	PushNotificationTargetUrl         *String   `xmlrpc:"push_notification_target_url,omptempty"`
	PushNotificationTitle             *String   `xmlrpc:"push_notification_title,omptempty"`
	PushNotificationsPreview          *String   `xmlrpc:"push_notifications_preview,omptempty"`
	TwitterPostLimitMessage           *String   `xmlrpc:"twitter_post_limit_message,omptempty"`
	TwitterPreview                    *String   `xmlrpc:"twitter_preview,omptempty"`
	UseVisitorTimezone                *Bool     `xmlrpc:"use_visitor_timezone,omptempty"`
	VisitorDomain                     *String   `xmlrpc:"visitor_domain,omptempty"`
	WriteDate                         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SocialPostTemplates represents array of social.post.template model.
type SocialPostTemplates []SocialPostTemplate

// SocialPostTemplateModel is the odoo model name.
const SocialPostTemplateModel = "social.post.template"

// Many2One convert SocialPostTemplate to *Many2One.
func (spt *SocialPostTemplate) Many2One() *Many2One {
	return NewMany2One(spt.Id.Get(), "")
}

// CreateSocialPostTemplate creates a new social.post.template model and returns its id.
func (c *Client) CreateSocialPostTemplate(spt *SocialPostTemplate) (int64, error) {
	ids, err := c.CreateSocialPostTemplates([]*SocialPostTemplate{spt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialPostTemplate creates a new social.post.template model and returns its id.
func (c *Client) CreateSocialPostTemplates(spts []*SocialPostTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range spts {
		vv = append(vv, v)
	}
	return c.Create(SocialPostTemplateModel, vv)
}

// UpdateSocialPostTemplate updates an existing social.post.template record.
func (c *Client) UpdateSocialPostTemplate(spt *SocialPostTemplate) error {
	return c.UpdateSocialPostTemplates([]int64{spt.Id.Get()}, spt)
}

// UpdateSocialPostTemplates updates existing social.post.template records.
// All records (represented by ids) will be updated by spt values.
func (c *Client) UpdateSocialPostTemplates(ids []int64, spt *SocialPostTemplate) error {
	return c.Update(SocialPostTemplateModel, ids, spt)
}

// DeleteSocialPostTemplate deletes an existing social.post.template record.
func (c *Client) DeleteSocialPostTemplate(id int64) error {
	return c.DeleteSocialPostTemplates([]int64{id})
}

// DeleteSocialPostTemplates deletes existing social.post.template records.
func (c *Client) DeleteSocialPostTemplates(ids []int64) error {
	return c.Delete(SocialPostTemplateModel, ids)
}

// GetSocialPostTemplate gets social.post.template existing record.
func (c *Client) GetSocialPostTemplate(id int64) (*SocialPostTemplate, error) {
	spts, err := c.GetSocialPostTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if spts != nil && len(*spts) > 0 {
		return &((*spts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.post.template not found", id)
}

// GetSocialPostTemplates gets social.post.template existing records.
func (c *Client) GetSocialPostTemplates(ids []int64) (*SocialPostTemplates, error) {
	spts := &SocialPostTemplates{}
	if err := c.Read(SocialPostTemplateModel, ids, nil, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindSocialPostTemplate finds social.post.template record by querying it with criteria.
func (c *Client) FindSocialPostTemplate(criteria *Criteria) (*SocialPostTemplate, error) {
	spts := &SocialPostTemplates{}
	if err := c.SearchRead(SocialPostTemplateModel, criteria, NewOptions().Limit(1), spts); err != nil {
		return nil, err
	}
	if spts != nil && len(*spts) > 0 {
		return &((*spts)[0]), nil
	}
	return nil, fmt.Errorf("social.post.template was not found with criteria %v", criteria)
}

// FindSocialPostTemplates finds social.post.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialPostTemplates(criteria *Criteria, options *Options) (*SocialPostTemplates, error) {
	spts := &SocialPostTemplates{}
	if err := c.SearchRead(SocialPostTemplateModel, criteria, options, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindSocialPostTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialPostTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialPostTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialPostTemplateId finds record id by querying it with criteria.
func (c *Client) FindSocialPostTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialPostTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.post.template was not found with criteria %v and options %v", criteria, options)
}
