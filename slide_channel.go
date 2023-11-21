package odoo

import (
	"fmt"
)

// SlideChannel represents slide.channel model.
type SlideChannel struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                  *String    `xmlrpc:"access_token,omptempty"`
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId      *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline         *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration  *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon        *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                  *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary              *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon             *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId               *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId               *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AllowComment                 *Bool      `xmlrpc:"allow_comment,omptempty"`
	CanComment                   *Bool      `xmlrpc:"can_comment,omptempty"`
	CanPublish                   *Bool      `xmlrpc:"can_publish,omptempty"`
	CanReview                    *Bool      `xmlrpc:"can_review,omptempty"`
	CanUpload                    *Bool      `xmlrpc:"can_upload,omptempty"`
	CanVote                      *Bool      `xmlrpc:"can_vote,omptempty"`
	ChannelPartnerIds            *Relation  `xmlrpc:"channel_partner_ids,omptempty"`
	ChannelType                  *Selection `xmlrpc:"channel_type,omptempty"`
	Color                        *Int       `xmlrpc:"color,omptempty"`
	Completed                    *Bool      `xmlrpc:"completed,omptempty"`
	CompletedTemplateId          *Many2One  `xmlrpc:"completed_template_id,omptempty"`
	Completion                   *Int       `xmlrpc:"completion,omptempty"`
	CoverProperties              *String    `xmlrpc:"cover_properties,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description                  *String    `xmlrpc:"description,omptempty"`
	DescriptionHtml              *String    `xmlrpc:"description_html,omptempty"`
	DescriptionShort             *String    `xmlrpc:"description_short,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	Enroll                       *Selection `xmlrpc:"enroll,omptempty"`
	EnrollGroupIds               *Relation  `xmlrpc:"enroll_group_ids,omptempty"`
	EnrollMsg                    *String    `xmlrpc:"enroll_msg,omptempty"`
	FailedMessageIds             *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omptempty"`
	HasRequestedAccess           *Bool      `xmlrpc:"has_requested_access,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	Image1024                    *String    `xmlrpc:"image_1024,omptempty"`
	Image128                     *String    `xmlrpc:"image_128,omptempty"`
	Image1920                    *String    `xmlrpc:"image_1920,omptempty"`
	Image256                     *String    `xmlrpc:"image_256,omptempty"`
	Image512                     *String    `xmlrpc:"image_512,omptempty"`
	IsMember                     *Bool      `xmlrpc:"is_member,omptempty"`
	IsPublished                  *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized               *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	KarmaGenChannelFinish        *Int       `xmlrpc:"karma_gen_channel_finish,omptempty"`
	KarmaGenChannelRank          *Int       `xmlrpc:"karma_gen_channel_rank,omptempty"`
	KarmaGenSlideVote            *Int       `xmlrpc:"karma_gen_slide_vote,omptempty"`
	KarmaReview                  *Int       `xmlrpc:"karma_review,omptempty"`
	KarmaSlideComment            *Int       `xmlrpc:"karma_slide_comment,omptempty"`
	KarmaSlideVote               *Int       `xmlrpc:"karma_slide_vote,omptempty"`
	MembersCount                 *Int       `xmlrpc:"members_count,omptempty"`
	MembersDoneCount             *Int       `xmlrpc:"members_done_count,omptempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent               *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId      *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline       *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	NbrArticle                   *Int       `xmlrpc:"nbr_article,omptempty"`
	NbrCertification             *Int       `xmlrpc:"nbr_certification,omptempty"`
	NbrDocument                  *Int       `xmlrpc:"nbr_document,omptempty"`
	NbrInfographic               *Int       `xmlrpc:"nbr_infographic,omptempty"`
	NbrQuiz                      *Int       `xmlrpc:"nbr_quiz,omptempty"`
	NbrVideo                     *Int       `xmlrpc:"nbr_video,omptempty"`
	PartnerHasNewContent         *Bool      `xmlrpc:"partner_has_new_content,omptempty"`
	PartnerIds                   *Relation  `xmlrpc:"partner_ids,omptempty"`
	PromoteStrategy              *Selection `xmlrpc:"promote_strategy,omptempty"`
	PromotedSlideId              *Many2One  `xmlrpc:"promoted_slide_id,omptempty"`
	PublishTemplateId            *Many2One  `xmlrpc:"publish_template_id,omptempty"`
	RatingAvg                    *Float     `xmlrpc:"rating_avg,omptempty"`
	RatingAvgStars               *Float     `xmlrpc:"rating_avg_stars,omptempty"`
	RatingAvgText                *Selection `xmlrpc:"rating_avg_text,omptempty"`
	RatingCount                  *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback           *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage              *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastText               *Selection `xmlrpc:"rating_last_text,omptempty"`
	RatingLastValue              *Float     `xmlrpc:"rating_last_value,omptempty"`
	RatingPercentageSatisfaction *Float     `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	SeoName                      *String    `xmlrpc:"seo_name,omptempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omptempty"`
	ShareChannelTemplateId       *Many2One  `xmlrpc:"share_channel_template_id,omptempty"`
	ShareSlideTemplateId         *Many2One  `xmlrpc:"share_slide_template_id,omptempty"`
	SlideCategoryIds             *Relation  `xmlrpc:"slide_category_ids,omptempty"`
	SlideContentIds              *Relation  `xmlrpc:"slide_content_ids,omptempty"`
	SlideIds                     *Relation  `xmlrpc:"slide_ids,omptempty"`
	SlideLastUpdate              *Time      `xmlrpc:"slide_last_update,omptempty"`
	SlidePartnerIds              *Relation  `xmlrpc:"slide_partner_ids,omptempty"`
	TagIds                       *Relation  `xmlrpc:"tag_ids,omptempty"`
	TotalSlides                  *Int       `xmlrpc:"total_slides,omptempty"`
	TotalTime                    *Float     `xmlrpc:"total_time,omptempty"`
	TotalViews                   *Int       `xmlrpc:"total_views,omptempty"`
	TotalVotes                   *Int       `xmlrpc:"total_votes,omptempty"`
	UploadGroupIds               *Relation  `xmlrpc:"upload_group_ids,omptempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omptempty"`
	Visibility                   *Selection `xmlrpc:"visibility,omptempty"`
	WebsiteBackgroundImageUrl    *String    `xmlrpc:"website_background_image_url,omptempty"`
	WebsiteId                    *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription       *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords          *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg             *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle             *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished             *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                   *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SlideChannels represents array of slide.channel model.
type SlideChannels []SlideChannel

// SlideChannelModel is the odoo model name.
const SlideChannelModel = "slide.channel"

// Many2One convert SlideChannel to *Many2One.
func (sc *SlideChannel) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSlideChannel creates a new slide.channel model and returns its id.
func (c *Client) CreateSlideChannel(sc *SlideChannel) (int64, error) {
	ids, err := c.CreateSlideChannels([]*SlideChannel{sc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSlideChannel creates a new slide.channel model and returns its id.
func (c *Client) CreateSlideChannels(scs []*SlideChannel) ([]int64, error) {
	var vv []interface{}
	for _, v := range scs {
		vv = append(vv, v)
	}
	return c.Create(SlideChannelModel, vv)
}

// UpdateSlideChannel updates an existing slide.channel record.
func (c *Client) UpdateSlideChannel(sc *SlideChannel) error {
	return c.UpdateSlideChannels([]int64{sc.Id.Get()}, sc)
}

// UpdateSlideChannels updates existing slide.channel records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateSlideChannels(ids []int64, sc *SlideChannel) error {
	return c.Update(SlideChannelModel, ids, sc)
}

// DeleteSlideChannel deletes an existing slide.channel record.
func (c *Client) DeleteSlideChannel(id int64) error {
	return c.DeleteSlideChannels([]int64{id})
}

// DeleteSlideChannels deletes existing slide.channel records.
func (c *Client) DeleteSlideChannels(ids []int64) error {
	return c.Delete(SlideChannelModel, ids)
}

// GetSlideChannel gets slide.channel existing record.
func (c *Client) GetSlideChannel(id int64) (*SlideChannel, error) {
	scs, err := c.GetSlideChannels([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.channel not found", id)
}

// GetSlideChannels gets slide.channel existing records.
func (c *Client) GetSlideChannels(ids []int64) (*SlideChannels, error) {
	scs := &SlideChannels{}
	if err := c.Read(SlideChannelModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannel finds slide.channel record by querying it with criteria.
func (c *Client) FindSlideChannel(criteria *Criteria) (*SlideChannel, error) {
	scs := &SlideChannels{}
	if err := c.SearchRead(SlideChannelModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("slide.channel was not found with criteria %v", criteria)
}

// FindSlideChannels finds slide.channel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannels(criteria *Criteria, options *Options) (*SlideChannels, error) {
	scs := &SlideChannels{}
	if err := c.SearchRead(SlideChannelModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSlideChannelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideChannelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideChannelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideChannelId finds record id by querying it with criteria.
func (c *Client) FindSlideChannelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideChannelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.channel was not found with criteria %v and options %v", criteria, options)
}
