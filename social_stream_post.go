package odoo

import (
	"fmt"
)

// SocialStreamPost represents social.stream.post model.
type SocialStreamPost struct {
	LastUpdate                        *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId                         *Many2One  `xmlrpc:"account_id,omptempty"`
	AuthorLink                        *String    `xmlrpc:"author_link,omptempty"`
	AuthorName                        *String    `xmlrpc:"author_name,omptempty"`
	CompanyId                         *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                       *String    `xmlrpc:"display_name,omptempty"`
	FacebookAuthorId                  *String    `xmlrpc:"facebook_author_id,omptempty"`
	FacebookCommentsCount             *Int       `xmlrpc:"facebook_comments_count,omptempty"`
	FacebookIsEventPost               *Bool      `xmlrpc:"facebook_is_event_post,omptempty"`
	FacebookLikesCount                *Int       `xmlrpc:"facebook_likes_count,omptempty"`
	FacebookPostId                    *String    `xmlrpc:"facebook_post_id,omptempty"`
	FacebookReach                     *Int       `xmlrpc:"facebook_reach,omptempty"`
	FacebookSharesCount               *Int       `xmlrpc:"facebook_shares_count,omptempty"`
	FacebookUserLikes                 *Bool      `xmlrpc:"facebook_user_likes,omptempty"`
	FormattedPublishedDate            *String    `xmlrpc:"formatted_published_date,omptempty"`
	Id                                *Int       `xmlrpc:"id,omptempty"`
	InstagramCommentsCount            *Int       `xmlrpc:"instagram_comments_count,omptempty"`
	InstagramFacebookAuthorId         *String    `xmlrpc:"instagram_facebook_author_id,omptempty"`
	InstagramLikesCount               *Int       `xmlrpc:"instagram_likes_count,omptempty"`
	InstagramPostId                   *String    `xmlrpc:"instagram_post_id,omptempty"`
	InstagramPostLink                 *String    `xmlrpc:"instagram_post_link,omptempty"`
	LinkDescription                   *String    `xmlrpc:"link_description,omptempty"`
	LinkImageUrl                      *String    `xmlrpc:"link_image_url,omptempty"`
	LinkTitle                         *String    `xmlrpc:"link_title,omptempty"`
	LinkUrl                           *String    `xmlrpc:"link_url,omptempty"`
	LinkedinAuthorId                  *String    `xmlrpc:"linkedin_author_id,omptempty"`
	LinkedinAuthorImageUrl            *String    `xmlrpc:"linkedin_author_image_url,omptempty"`
	LinkedinAuthorUrn                 *String    `xmlrpc:"linkedin_author_urn,omptempty"`
	LinkedinAuthorVanityName          *String    `xmlrpc:"linkedin_author_vanity_name,omptempty"`
	LinkedinCommentsCount             *Int       `xmlrpc:"linkedin_comments_count,omptempty"`
	LinkedinLikesCount                *Int       `xmlrpc:"linkedin_likes_count,omptempty"`
	LinkedinPostUrn                   *String    `xmlrpc:"linkedin_post_urn,omptempty"`
	MediaType                         *Selection `xmlrpc:"media_type,omptempty"`
	Message                           *String    `xmlrpc:"message,omptempty"`
	PostLink                          *String    `xmlrpc:"post_link,omptempty"`
	PublishedDate                     *Time      `xmlrpc:"published_date,omptempty"`
	StreamId                          *Many2One  `xmlrpc:"stream_id,omptempty"`
	StreamPostImageIds                *Relation  `xmlrpc:"stream_post_image_ids,omptempty"`
	StreamPostImageUrls               *String    `xmlrpc:"stream_post_image_urls,omptempty"`
	TwitterAuthorId                   *String    `xmlrpc:"twitter_author_id,omptempty"`
	TwitterCanRetweet                 *Bool      `xmlrpc:"twitter_can_retweet,omptempty"`
	TwitterCommentsCount              *Int       `xmlrpc:"twitter_comments_count,omptempty"`
	TwitterLikesCount                 *Int       `xmlrpc:"twitter_likes_count,omptempty"`
	TwitterProfileImageUrl            *String    `xmlrpc:"twitter_profile_image_url,omptempty"`
	TwitterQuotedTweetAuthorLink      *String    `xmlrpc:"twitter_quoted_tweet_author_link,omptempty"`
	TwitterQuotedTweetAuthorName      *String    `xmlrpc:"twitter_quoted_tweet_author_name,omptempty"`
	TwitterQuotedTweetIdStr           *String    `xmlrpc:"twitter_quoted_tweet_id_str,omptempty"`
	TwitterQuotedTweetMessage         *String    `xmlrpc:"twitter_quoted_tweet_message,omptempty"`
	TwitterQuotedTweetProfileImageUrl *String    `xmlrpc:"twitter_quoted_tweet_profile_image_url,omptempty"`
	TwitterRetweetCount               *Int       `xmlrpc:"twitter_retweet_count,omptempty"`
	TwitterRetweetedTweetIdStr        *String    `xmlrpc:"twitter_retweeted_tweet_id_str,omptempty"`
	TwitterScreenName                 *String    `xmlrpc:"twitter_screen_name,omptempty"`
	TwitterTweetId                    *String    `xmlrpc:"twitter_tweet_id,omptempty"`
	TwitterUserLikes                  *Bool      `xmlrpc:"twitter_user_likes,omptempty"`
	WriteDate                         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                          *Many2One  `xmlrpc:"write_uid,omptempty"`
	YoutubeCommentsCount              *Int       `xmlrpc:"youtube_comments_count,omptempty"`
	YoutubeDislikesCount              *Int       `xmlrpc:"youtube_dislikes_count,omptempty"`
	YoutubeLikesCount                 *Int       `xmlrpc:"youtube_likes_count,omptempty"`
	YoutubeLikesRatio                 *Int       `xmlrpc:"youtube_likes_ratio,omptempty"`
	YoutubeThumbnailUrl               *String    `xmlrpc:"youtube_thumbnail_url,omptempty"`
	YoutubeVideoId                    *String    `xmlrpc:"youtube_video_id,omptempty"`
	YoutubeViewsCount                 *Int       `xmlrpc:"youtube_views_count,omptempty"`
}

// SocialStreamPosts represents array of social.stream.post model.
type SocialStreamPosts []SocialStreamPost

// SocialStreamPostModel is the odoo model name.
const SocialStreamPostModel = "social.stream.post"

// Many2One convert SocialStreamPost to *Many2One.
func (ssp *SocialStreamPost) Many2One() *Many2One {
	return NewMany2One(ssp.Id.Get(), "")
}

// CreateSocialStreamPost creates a new social.stream.post model and returns its id.
func (c *Client) CreateSocialStreamPost(ssp *SocialStreamPost) (int64, error) {
	ids, err := c.CreateSocialStreamPosts([]*SocialStreamPost{ssp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialStreamPost creates a new social.stream.post model and returns its id.
func (c *Client) CreateSocialStreamPosts(ssps []*SocialStreamPost) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssps {
		vv = append(vv, v)
	}
	return c.Create(SocialStreamPostModel, vv)
}

// UpdateSocialStreamPost updates an existing social.stream.post record.
func (c *Client) UpdateSocialStreamPost(ssp *SocialStreamPost) error {
	return c.UpdateSocialStreamPosts([]int64{ssp.Id.Get()}, ssp)
}

// UpdateSocialStreamPosts updates existing social.stream.post records.
// All records (represented by ids) will be updated by ssp values.
func (c *Client) UpdateSocialStreamPosts(ids []int64, ssp *SocialStreamPost) error {
	return c.Update(SocialStreamPostModel, ids, ssp)
}

// DeleteSocialStreamPost deletes an existing social.stream.post record.
func (c *Client) DeleteSocialStreamPost(id int64) error {
	return c.DeleteSocialStreamPosts([]int64{id})
}

// DeleteSocialStreamPosts deletes existing social.stream.post records.
func (c *Client) DeleteSocialStreamPosts(ids []int64) error {
	return c.Delete(SocialStreamPostModel, ids)
}

// GetSocialStreamPost gets social.stream.post existing record.
func (c *Client) GetSocialStreamPost(id int64) (*SocialStreamPost, error) {
	ssps, err := c.GetSocialStreamPosts([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssps != nil && len(*ssps) > 0 {
		return &((*ssps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.stream.post not found", id)
}

// GetSocialStreamPosts gets social.stream.post existing records.
func (c *Client) GetSocialStreamPosts(ids []int64) (*SocialStreamPosts, error) {
	ssps := &SocialStreamPosts{}
	if err := c.Read(SocialStreamPostModel, ids, nil, ssps); err != nil {
		return nil, err
	}
	return ssps, nil
}

// FindSocialStreamPost finds social.stream.post record by querying it with criteria.
func (c *Client) FindSocialStreamPost(criteria *Criteria) (*SocialStreamPost, error) {
	ssps := &SocialStreamPosts{}
	if err := c.SearchRead(SocialStreamPostModel, criteria, NewOptions().Limit(1), ssps); err != nil {
		return nil, err
	}
	if ssps != nil && len(*ssps) > 0 {
		return &((*ssps)[0]), nil
	}
	return nil, fmt.Errorf("social.stream.post was not found with criteria %v", criteria)
}

// FindSocialStreamPosts finds social.stream.post records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreamPosts(criteria *Criteria, options *Options) (*SocialStreamPosts, error) {
	ssps := &SocialStreamPosts{}
	if err := c.SearchRead(SocialStreamPostModel, criteria, options, ssps); err != nil {
		return nil, err
	}
	return ssps, nil
}

// FindSocialStreamPostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreamPostIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialStreamPostModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialStreamPostId finds record id by querying it with criteria.
func (c *Client) FindSocialStreamPostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialStreamPostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.stream.post was not found with criteria %v and options %v", criteria, options)
}
