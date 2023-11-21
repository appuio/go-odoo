package odoo

import (
	"fmt"
)

// SocialMedia represents social.media model.
type SocialMedia struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountIds               *Relation  `xmlrpc:"account_ids,omptempty"`
	AccountsCount            *Int       `xmlrpc:"accounts_count,omptempty"`
	CanLinkAccounts          *Bool      `xmlrpc:"can_link_accounts,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CsrfToken                *String    `xmlrpc:"csrf_token,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FailedMessageIds         *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	HasStreams               *Bool      `xmlrpc:"has_streams,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	Image                    *String    `xmlrpc:"image,omptempty"`
	MaxPostLength            *Int       `xmlrpc:"max_post_length,omptempty"`
	MediaDescription         *String    `xmlrpc:"media_description,omptempty"`
	MediaType                *Selection `xmlrpc:"media_type,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent           *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	StreamTypeIds            *Relation  `xmlrpc:"stream_type_ids,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SocialMedias represents array of social.media model.
type SocialMedias []SocialMedia

// SocialMediaModel is the odoo model name.
const SocialMediaModel = "social.media"

// Many2One convert SocialMedia to *Many2One.
func (sm *SocialMedia) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateSocialMedia creates a new social.media model and returns its id.
func (c *Client) CreateSocialMedia(sm *SocialMedia) (int64, error) {
	ids, err := c.CreateSocialMedias([]*SocialMedia{sm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialMedia creates a new social.media model and returns its id.
func (c *Client) CreateSocialMedias(sms []*SocialMedia) ([]int64, error) {
	var vv []interface{}
	for _, v := range sms {
		vv = append(vv, v)
	}
	return c.Create(SocialMediaModel, vv)
}

// UpdateSocialMedia updates an existing social.media record.
func (c *Client) UpdateSocialMedia(sm *SocialMedia) error {
	return c.UpdateSocialMedias([]int64{sm.Id.Get()}, sm)
}

// UpdateSocialMedias updates existing social.media records.
// All records (represented by ids) will be updated by sm values.
func (c *Client) UpdateSocialMedias(ids []int64, sm *SocialMedia) error {
	return c.Update(SocialMediaModel, ids, sm)
}

// DeleteSocialMedia deletes an existing social.media record.
func (c *Client) DeleteSocialMedia(id int64) error {
	return c.DeleteSocialMedias([]int64{id})
}

// DeleteSocialMedias deletes existing social.media records.
func (c *Client) DeleteSocialMedias(ids []int64) error {
	return c.Delete(SocialMediaModel, ids)
}

// GetSocialMedia gets social.media existing record.
func (c *Client) GetSocialMedia(id int64) (*SocialMedia, error) {
	sms, err := c.GetSocialMedias([]int64{id})
	if err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.media not found", id)
}

// GetSocialMedias gets social.media existing records.
func (c *Client) GetSocialMedias(ids []int64) (*SocialMedias, error) {
	sms := &SocialMedias{}
	if err := c.Read(SocialMediaModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSocialMedia finds social.media record by querying it with criteria.
func (c *Client) FindSocialMedia(criteria *Criteria) (*SocialMedia, error) {
	sms := &SocialMedias{}
	if err := c.SearchRead(SocialMediaModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("social.media was not found with criteria %v", criteria)
}

// FindSocialMedias finds social.media records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialMedias(criteria *Criteria, options *Options) (*SocialMedias, error) {
	sms := &SocialMedias{}
	if err := c.SearchRead(SocialMediaModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSocialMediaIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialMediaIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialMediaModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialMediaId finds record id by querying it with criteria.
func (c *Client) FindSocialMediaId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialMediaModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.media was not found with criteria %v and options %v", criteria, options)
}
