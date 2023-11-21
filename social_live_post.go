package odoo

import (
	"fmt"
)

// SocialLivePost represents social.live.post model.
type SocialLivePost struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId         *Many2One  `xmlrpc:"account_id,omptempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Engagement        *Int       `xmlrpc:"engagement,omptempty"`
	FacebookPostId    *String    `xmlrpc:"facebook_post_id,omptempty"`
	FailureReason     *String    `xmlrpc:"failure_reason,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	InstagramPostId   *String    `xmlrpc:"instagram_post_id,omptempty"`
	LinkedinPostId    *String    `xmlrpc:"linkedin_post_id,omptempty"`
	LivePostLink      *String    `xmlrpc:"live_post_link,omptempty"`
	Message           *String    `xmlrpc:"message,omptempty"`
	PostId            *Many2One  `xmlrpc:"post_id,omptempty"`
	ReachedVisitorIds *Relation  `xmlrpc:"reached_visitor_ids,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	TwitterTweetId    *String    `xmlrpc:"twitter_tweet_id,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
	YoutubeVideoId    *String    `xmlrpc:"youtube_video_id,omptempty"`
}

// SocialLivePosts represents array of social.live.post model.
type SocialLivePosts []SocialLivePost

// SocialLivePostModel is the odoo model name.
const SocialLivePostModel = "social.live.post"

// Many2One convert SocialLivePost to *Many2One.
func (slp *SocialLivePost) Many2One() *Many2One {
	return NewMany2One(slp.Id.Get(), "")
}

// CreateSocialLivePost creates a new social.live.post model and returns its id.
func (c *Client) CreateSocialLivePost(slp *SocialLivePost) (int64, error) {
	ids, err := c.CreateSocialLivePosts([]*SocialLivePost{slp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialLivePost creates a new social.live.post model and returns its id.
func (c *Client) CreateSocialLivePosts(slps []*SocialLivePost) ([]int64, error) {
	var vv []interface{}
	for _, v := range slps {
		vv = append(vv, v)
	}
	return c.Create(SocialLivePostModel, vv)
}

// UpdateSocialLivePost updates an existing social.live.post record.
func (c *Client) UpdateSocialLivePost(slp *SocialLivePost) error {
	return c.UpdateSocialLivePosts([]int64{slp.Id.Get()}, slp)
}

// UpdateSocialLivePosts updates existing social.live.post records.
// All records (represented by ids) will be updated by slp values.
func (c *Client) UpdateSocialLivePosts(ids []int64, slp *SocialLivePost) error {
	return c.Update(SocialLivePostModel, ids, slp)
}

// DeleteSocialLivePost deletes an existing social.live.post record.
func (c *Client) DeleteSocialLivePost(id int64) error {
	return c.DeleteSocialLivePosts([]int64{id})
}

// DeleteSocialLivePosts deletes existing social.live.post records.
func (c *Client) DeleteSocialLivePosts(ids []int64) error {
	return c.Delete(SocialLivePostModel, ids)
}

// GetSocialLivePost gets social.live.post existing record.
func (c *Client) GetSocialLivePost(id int64) (*SocialLivePost, error) {
	slps, err := c.GetSocialLivePosts([]int64{id})
	if err != nil {
		return nil, err
	}
	if slps != nil && len(*slps) > 0 {
		return &((*slps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.live.post not found", id)
}

// GetSocialLivePosts gets social.live.post existing records.
func (c *Client) GetSocialLivePosts(ids []int64) (*SocialLivePosts, error) {
	slps := &SocialLivePosts{}
	if err := c.Read(SocialLivePostModel, ids, nil, slps); err != nil {
		return nil, err
	}
	return slps, nil
}

// FindSocialLivePost finds social.live.post record by querying it with criteria.
func (c *Client) FindSocialLivePost(criteria *Criteria) (*SocialLivePost, error) {
	slps := &SocialLivePosts{}
	if err := c.SearchRead(SocialLivePostModel, criteria, NewOptions().Limit(1), slps); err != nil {
		return nil, err
	}
	if slps != nil && len(*slps) > 0 {
		return &((*slps)[0]), nil
	}
	return nil, fmt.Errorf("social.live.post was not found with criteria %v", criteria)
}

// FindSocialLivePosts finds social.live.post records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialLivePosts(criteria *Criteria, options *Options) (*SocialLivePosts, error) {
	slps := &SocialLivePosts{}
	if err := c.SearchRead(SocialLivePostModel, criteria, options, slps); err != nil {
		return nil, err
	}
	return slps, nil
}

// FindSocialLivePostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialLivePostIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialLivePostModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialLivePostId finds record id by querying it with criteria.
func (c *Client) FindSocialLivePostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialLivePostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.live.post was not found with criteria %v and options %v", criteria, options)
}
