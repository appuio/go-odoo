package odoo

import (
	"fmt"
)

// SlideSlide represents slide.slide model.
type SlideSlide struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	Active                    *Bool      `xmlrpc:"active,omptempty"`
	BinaryContent             *String    `xmlrpc:"binary_content,omptempty"`
	CanPublish                *Bool      `xmlrpc:"can_publish,omptempty"`
	CanSelfMarkCompleted      *Bool      `xmlrpc:"can_self_mark_completed,omptempty"`
	CanSelfMarkUncompleted    *Bool      `xmlrpc:"can_self_mark_uncompleted,omptempty"`
	CategoryId                *Many2One  `xmlrpc:"category_id,omptempty"`
	ChannelAllowComment       *Bool      `xmlrpc:"channel_allow_comment,omptempty"`
	ChannelId                 *Many2One  `xmlrpc:"channel_id,omptempty"`
	ChannelType               *Selection `xmlrpc:"channel_type,omptempty"`
	CommentsCount             *Int       `xmlrpc:"comments_count,omptempty"`
	CompletionTime            *Float     `xmlrpc:"completion_time,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	DatePublished             *Time      `xmlrpc:"date_published,omptempty"`
	Description               *String    `xmlrpc:"description,omptempty"`
	Dislikes                  *Int       `xmlrpc:"dislikes,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	DocumentBinaryContent     *String    `xmlrpc:"document_binary_content,omptempty"`
	DocumentGoogleUrl         *String    `xmlrpc:"document_google_url,omptempty"`
	EmbedCode                 *String    `xmlrpc:"embed_code,omptempty"`
	EmbedCodeExternal         *String    `xmlrpc:"embed_code_external,omptempty"`
	EmbedCount                *Int       `xmlrpc:"embed_count,omptempty"`
	EmbedIds                  *Relation  `xmlrpc:"embed_ids,omptempty"`
	FailedMessageIds          *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	GoogleDriveId             *String    `xmlrpc:"google_drive_id,omptempty"`
	HasMessage                *Bool      `xmlrpc:"has_message,omptempty"`
	HtmlContent               *String    `xmlrpc:"html_content,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	Image1024                 *String    `xmlrpc:"image_1024,omptempty"`
	Image128                  *String    `xmlrpc:"image_128,omptempty"`
	Image1920                 *String    `xmlrpc:"image_1920,omptempty"`
	Image256                  *String    `xmlrpc:"image_256,omptempty"`
	Image512                  *String    `xmlrpc:"image_512,omptempty"`
	ImageBinaryContent        *String    `xmlrpc:"image_binary_content,omptempty"`
	ImageGoogleUrl            *String    `xmlrpc:"image_google_url,omptempty"`
	IsCategory                *Bool      `xmlrpc:"is_category,omptempty"`
	IsNewSlide                *Bool      `xmlrpc:"is_new_slide,omptempty"`
	IsPreview                 *Bool      `xmlrpc:"is_preview,omptempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized            *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	Likes                     *Int       `xmlrpc:"likes,omptempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent            *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	NbrArticle                *Int       `xmlrpc:"nbr_article,omptempty"`
	NbrCertification          *Int       `xmlrpc:"nbr_certification,omptempty"`
	NbrDocument               *Int       `xmlrpc:"nbr_document,omptempty"`
	NbrInfographic            *Int       `xmlrpc:"nbr_infographic,omptempty"`
	NbrQuiz                   *Int       `xmlrpc:"nbr_quiz,omptempty"`
	NbrVideo                  *Int       `xmlrpc:"nbr_video,omptempty"`
	PartnerIds                *Relation  `xmlrpc:"partner_ids,omptempty"`
	PublicViews               *Int       `xmlrpc:"public_views,omptempty"`
	QuestionIds               *Relation  `xmlrpc:"question_ids,omptempty"`
	QuestionsCount            *Int       `xmlrpc:"questions_count,omptempty"`
	QuizFirstAttemptReward    *Int       `xmlrpc:"quiz_first_attempt_reward,omptempty"`
	QuizFourthAttemptReward   *Int       `xmlrpc:"quiz_fourth_attempt_reward,omptempty"`
	QuizSecondAttemptReward   *Int       `xmlrpc:"quiz_second_attempt_reward,omptempty"`
	QuizThirdAttemptReward    *Int       `xmlrpc:"quiz_third_attempt_reward,omptempty"`
	SeoName                   *String    `xmlrpc:"seo_name,omptempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omptempty"`
	SlideCategory             *Selection `xmlrpc:"slide_category,omptempty"`
	SlideIds                  *Relation  `xmlrpc:"slide_ids,omptempty"`
	SlidePartnerIds           *Relation  `xmlrpc:"slide_partner_ids,omptempty"`
	SlideResourceDownloadable *Bool      `xmlrpc:"slide_resource_downloadable,omptempty"`
	SlideResourceIds          *Relation  `xmlrpc:"slide_resource_ids,omptempty"`
	SlideType                 *Selection `xmlrpc:"slide_type,omptempty"`
	SlideViews                *Int       `xmlrpc:"slide_views,omptempty"`
	SourceType                *Selection `xmlrpc:"source_type,omptempty"`
	SurveyId                  *Many2One  `xmlrpc:"survey_id,omptempty"`
	TagIds                    *Relation  `xmlrpc:"tag_ids,omptempty"`
	TotalSlides               *Int       `xmlrpc:"total_slides,omptempty"`
	TotalViews                *Int       `xmlrpc:"total_views,omptempty"`
	Url                       *String    `xmlrpc:"url,omptempty"`
	UserHasCompleted          *Bool      `xmlrpc:"user_has_completed,omptempty"`
	UserId                    *Many2One  `xmlrpc:"user_id,omptempty"`
	UserMembershipId          *Many2One  `xmlrpc:"user_membership_id,omptempty"`
	UserVote                  *Int       `xmlrpc:"user_vote,omptempty"`
	VideoSourceType           *Selection `xmlrpc:"video_source_type,omptempty"`
	VideoUrl                  *String    `xmlrpc:"video_url,omptempty"`
	VimeoId                   *String    `xmlrpc:"vimeo_id,omptempty"`
	WebsiteId                 *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription    *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords       *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg          *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle          *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished          *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteShareUrl           *String    `xmlrpc:"website_share_url,omptempty"`
	WebsiteUrl                *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
	YoutubeId                 *String    `xmlrpc:"youtube_id,omptempty"`
}

// SlideSlides represents array of slide.slide model.
type SlideSlides []SlideSlide

// SlideSlideModel is the odoo model name.
const SlideSlideModel = "slide.slide"

// Many2One convert SlideSlide to *Many2One.
func (ss *SlideSlide) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSlideSlide creates a new slide.slide model and returns its id.
func (c *Client) CreateSlideSlide(ss *SlideSlide) (int64, error) {
	ids, err := c.CreateSlideSlides([]*SlideSlide{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSlideSlide creates a new slide.slide model and returns its id.
func (c *Client) CreateSlideSlides(sss []*SlideSlide) ([]int64, error) {
	var vv []interface{}
	for _, v := range sss {
		vv = append(vv, v)
	}
	return c.Create(SlideSlideModel, vv)
}

// UpdateSlideSlide updates an existing slide.slide record.
func (c *Client) UpdateSlideSlide(ss *SlideSlide) error {
	return c.UpdateSlideSlides([]int64{ss.Id.Get()}, ss)
}

// UpdateSlideSlides updates existing slide.slide records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSlideSlides(ids []int64, ss *SlideSlide) error {
	return c.Update(SlideSlideModel, ids, ss)
}

// DeleteSlideSlide deletes an existing slide.slide record.
func (c *Client) DeleteSlideSlide(id int64) error {
	return c.DeleteSlideSlides([]int64{id})
}

// DeleteSlideSlides deletes existing slide.slide records.
func (c *Client) DeleteSlideSlides(ids []int64) error {
	return c.Delete(SlideSlideModel, ids)
}

// GetSlideSlide gets slide.slide existing record.
func (c *Client) GetSlideSlide(id int64) (*SlideSlide, error) {
	sss, err := c.GetSlideSlides([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide not found", id)
}

// GetSlideSlides gets slide.slide existing records.
func (c *Client) GetSlideSlides(ids []int64) (*SlideSlides, error) {
	sss := &SlideSlides{}
	if err := c.Read(SlideSlideModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlide finds slide.slide record by querying it with criteria.
func (c *Client) FindSlideSlide(criteria *Criteria) (*SlideSlide, error) {
	sss := &SlideSlides{}
	if err := c.SearchRead(SlideSlideModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide was not found with criteria %v", criteria)
}

// FindSlideSlides finds slide.slide records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlides(criteria *Criteria, options *Options) (*SlideSlides, error) {
	sss := &SlideSlides{}
	if err := c.SearchRead(SlideSlideModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSlideSlideIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideSlideId finds record id by querying it with criteria.
func (c *Client) FindSlideSlideId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide was not found with criteria %v and options %v", criteria, options)
}
