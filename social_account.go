package odoo

import (
	"fmt"
)

// SocialAccount represents social.account model.
type SocialAccount struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	Active                     *Bool      `xmlrpc:"active,omptempty"`
	Audience                   *Int       `xmlrpc:"audience,omptempty"`
	AudienceTrend              *Float     `xmlrpc:"audience_trend,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	Engagement                 *Int       `xmlrpc:"engagement,omptempty"`
	EngagementTrend            *Float     `xmlrpc:"engagement_trend,omptempty"`
	FacebookAccessToken        *String    `xmlrpc:"facebook_access_token,omptempty"`
	FacebookAccountId          *String    `xmlrpc:"facebook_account_id,omptempty"`
	FirebaseAdminKeyFile       *String    `xmlrpc:"firebase_admin_key_file,omptempty"`
	FirebaseProjectId          *String    `xmlrpc:"firebase_project_id,omptempty"`
	FirebasePushCertificateKey *String    `xmlrpc:"firebase_push_certificate_key,omptempty"`
	FirebaseSenderId           *String    `xmlrpc:"firebase_sender_id,omptempty"`
	FirebaseUseOwnAccount      *Bool      `xmlrpc:"firebase_use_own_account,omptempty"`
	FirebaseWebApiKey          *String    `xmlrpc:"firebase_web_api_key,omptempty"`
	HasAccountStats            *Bool      `xmlrpc:"has_account_stats,omptempty"`
	HasTrends                  *Bool      `xmlrpc:"has_trends,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	Image                      *String    `xmlrpc:"image,omptempty"`
	InstagramAccessToken       *String    `xmlrpc:"instagram_access_token,omptempty"`
	InstagramAccountId         *String    `xmlrpc:"instagram_account_id,omptempty"`
	InstagramFacebookAccountId *String    `xmlrpc:"instagram_facebook_account_id,omptempty"`
	IsMediaDisconnected        *Bool      `xmlrpc:"is_media_disconnected,omptempty"`
	LinkedinAccessToken        *String    `xmlrpc:"linkedin_access_token,omptempty"`
	LinkedinAccountId          *String    `xmlrpc:"linkedin_account_id,omptempty"`
	LinkedinAccountUrn         *String    `xmlrpc:"linkedin_account_urn,omptempty"`
	MediaId                    *Many2One  `xmlrpc:"media_id,omptempty"`
	MediaType                  *Selection `xmlrpc:"media_type,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	NotificationRequestBody    *String    `xmlrpc:"notification_request_body,omptempty"`
	NotificationRequestDelay   *Int       `xmlrpc:"notification_request_delay,omptempty"`
	NotificationRequestIcon    *String    `xmlrpc:"notification_request_icon,omptempty"`
	NotificationRequestTitle   *String    `xmlrpc:"notification_request_title,omptempty"`
	SocialAccountHandle        *String    `xmlrpc:"social_account_handle,omptempty"`
	StatsLink                  *String    `xmlrpc:"stats_link,omptempty"`
	Stories                    *Int       `xmlrpc:"stories,omptempty"`
	StoriesTrend               *Float     `xmlrpc:"stories_trend,omptempty"`
	TwitterOauthToken          *String    `xmlrpc:"twitter_oauth_token,omptempty"`
	TwitterOauthTokenSecret    *String    `xmlrpc:"twitter_oauth_token_secret,omptempty"`
	TwitterUserId              *String    `xmlrpc:"twitter_user_id,omptempty"`
	UtmMediumId                *Many2One  `xmlrpc:"utm_medium_id,omptempty"`
	WebsiteId                  *Many2One  `xmlrpc:"website_id,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
	YoutubeAccessToken         *String    `xmlrpc:"youtube_access_token,omptempty"`
	YoutubeChannelId           *String    `xmlrpc:"youtube_channel_id,omptempty"`
	YoutubeRefreshToken        *String    `xmlrpc:"youtube_refresh_token,omptempty"`
	YoutubeTokenExpirationDate *Time      `xmlrpc:"youtube_token_expiration_date,omptempty"`
	YoutubeUploadPlaylistId    *String    `xmlrpc:"youtube_upload_playlist_id,omptempty"`
}

// SocialAccounts represents array of social.account model.
type SocialAccounts []SocialAccount

// SocialAccountModel is the odoo model name.
const SocialAccountModel = "social.account"

// Many2One convert SocialAccount to *Many2One.
func (sa *SocialAccount) Many2One() *Many2One {
	return NewMany2One(sa.Id.Get(), "")
}

// CreateSocialAccount creates a new social.account model and returns its id.
func (c *Client) CreateSocialAccount(sa *SocialAccount) (int64, error) {
	ids, err := c.CreateSocialAccounts([]*SocialAccount{sa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialAccount creates a new social.account model and returns its id.
func (c *Client) CreateSocialAccounts(sas []*SocialAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range sas {
		vv = append(vv, v)
	}
	return c.Create(SocialAccountModel, vv)
}

// UpdateSocialAccount updates an existing social.account record.
func (c *Client) UpdateSocialAccount(sa *SocialAccount) error {
	return c.UpdateSocialAccounts([]int64{sa.Id.Get()}, sa)
}

// UpdateSocialAccounts updates existing social.account records.
// All records (represented by ids) will be updated by sa values.
func (c *Client) UpdateSocialAccounts(ids []int64, sa *SocialAccount) error {
	return c.Update(SocialAccountModel, ids, sa)
}

// DeleteSocialAccount deletes an existing social.account record.
func (c *Client) DeleteSocialAccount(id int64) error {
	return c.DeleteSocialAccounts([]int64{id})
}

// DeleteSocialAccounts deletes existing social.account records.
func (c *Client) DeleteSocialAccounts(ids []int64) error {
	return c.Delete(SocialAccountModel, ids)
}

// GetSocialAccount gets social.account existing record.
func (c *Client) GetSocialAccount(id int64) (*SocialAccount, error) {
	sas, err := c.GetSocialAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.account not found", id)
}

// GetSocialAccounts gets social.account existing records.
func (c *Client) GetSocialAccounts(ids []int64) (*SocialAccounts, error) {
	sas := &SocialAccounts{}
	if err := c.Read(SocialAccountModel, ids, nil, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSocialAccount finds social.account record by querying it with criteria.
func (c *Client) FindSocialAccount(criteria *Criteria) (*SocialAccount, error) {
	sas := &SocialAccounts{}
	if err := c.SearchRead(SocialAccountModel, criteria, NewOptions().Limit(1), sas); err != nil {
		return nil, err
	}
	if sas != nil && len(*sas) > 0 {
		return &((*sas)[0]), nil
	}
	return nil, fmt.Errorf("social.account was not found with criteria %v", criteria)
}

// FindSocialAccounts finds social.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialAccounts(criteria *Criteria, options *Options) (*SocialAccounts, error) {
	sas := &SocialAccounts{}
	if err := c.SearchRead(SocialAccountModel, criteria, options, sas); err != nil {
		return nil, err
	}
	return sas, nil
}

// FindSocialAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialAccountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialAccountId finds record id by querying it with criteria.
func (c *Client) FindSocialAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.account was not found with criteria %v and options %v", criteria, options)
}
