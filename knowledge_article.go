package odoo

import (
	"fmt"
)

// KnowledgeArticle represents knowledge.article model.
type KnowledgeArticle struct {
	LastUpdate                  *Time       `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool       `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	ArticleMemberIds            *Relation   `xmlrpc:"article_member_ids,omptempty"`
	ArticleProperties           interface{} `xmlrpc:"article_properties,omptempty"`
	ArticlePropertiesDefinition interface{} `xmlrpc:"article_properties_definition,omptempty"`
	ArticlePropertiesIsEmpty    *Bool       `xmlrpc:"article_properties_is_empty,omptempty"`
	ArticleUrl                  *String     `xmlrpc:"article_url,omptempty"`
	Body                        *String     `xmlrpc:"body,omptempty"`
	CanPublish                  *Bool       `xmlrpc:"can_publish,omptempty"`
	Category                    *Selection  `xmlrpc:"category,omptempty"`
	ChildIds                    *Relation   `xmlrpc:"child_ids,omptempty"`
	CoverImageId                *Many2One   `xmlrpc:"cover_image_id,omptempty"`
	CoverImageUrl               *String     `xmlrpc:"cover_image_url,omptempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omptempty"`
	DeletionDate                *Time       `xmlrpc:"deletion_date,omptempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omptempty"`
	FailedMessageIds            *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	FavoriteCount               *Int        `xmlrpc:"favorite_count,omptempty"`
	FavoriteIds                 *Relation   `xmlrpc:"favorite_ids,omptempty"`
	FullWidth                   *Bool       `xmlrpc:"full_width,omptempty"`
	HasArticleChildren          *Bool       `xmlrpc:"has_article_children,omptempty"`
	HasItemChildren             *Bool       `xmlrpc:"has_item_children,omptempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omptempty"`
	Icon                        *String     `xmlrpc:"icon,omptempty"`
	Id                          *Int        `xmlrpc:"id,omptempty"`
	InheritedPermission         *Selection  `xmlrpc:"inherited_permission,omptempty"`
	InheritedPermissionParentId *Many2One   `xmlrpc:"inherited_permission_parent_id,omptempty"`
	InternalPermission          *Selection  `xmlrpc:"internal_permission,omptempty"`
	IsArticleItem               *Bool       `xmlrpc:"is_article_item,omptempty"`
	IsDesynchronized            *Bool       `xmlrpc:"is_desynchronized,omptempty"`
	IsLocked                    *Bool       `xmlrpc:"is_locked,omptempty"`
	IsPublished                 *Bool       `xmlrpc:"is_published,omptempty"`
	IsUserFavorite              *Bool       `xmlrpc:"is_user_favorite,omptempty"`
	LastEditionDate             *Time       `xmlrpc:"last_edition_date,omptempty"`
	LastEditionUid              *Many2One   `xmlrpc:"last_edition_uid,omptempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent              *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String     `xmlrpc:"name,omptempty"`
	ParentId                    *Many2One   `xmlrpc:"parent_id,omptempty"`
	ParentPath                  *String     `xmlrpc:"parent_path,omptempty"`
	RootArticleId               *Many2One   `xmlrpc:"root_article_id,omptempty"`
	Sequence                    *Int        `xmlrpc:"sequence,omptempty"`
	ToDelete                    *Bool       `xmlrpc:"to_delete,omptempty"`
	UserCanRead                 *Bool       `xmlrpc:"user_can_read,omptempty"`
	UserCanWrite                *Bool       `xmlrpc:"user_can_write,omptempty"`
	UserFavoriteSequence        *Int        `xmlrpc:"user_favorite_sequence,omptempty"`
	UserHasAccess               *Bool       `xmlrpc:"user_has_access,omptempty"`
	UserHasWriteAccess          *Bool       `xmlrpc:"user_has_write_access,omptempty"`
	UserPermission              *Selection  `xmlrpc:"user_permission,omptempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WebsitePublished            *Bool       `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                  *String     `xmlrpc:"website_url,omptempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// KnowledgeArticles represents array of knowledge.article model.
type KnowledgeArticles []KnowledgeArticle

// KnowledgeArticleModel is the odoo model name.
const KnowledgeArticleModel = "knowledge.article"

// Many2One convert KnowledgeArticle to *Many2One.
func (ka *KnowledgeArticle) Many2One() *Many2One {
	return NewMany2One(ka.Id.Get(), "")
}

// CreateKnowledgeArticle creates a new knowledge.article model and returns its id.
func (c *Client) CreateKnowledgeArticle(ka *KnowledgeArticle) (int64, error) {
	ids, err := c.CreateKnowledgeArticles([]*KnowledgeArticle{ka})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticle creates a new knowledge.article model and returns its id.
func (c *Client) CreateKnowledgeArticles(kas []*KnowledgeArticle) ([]int64, error) {
	var vv []interface{}
	for _, v := range kas {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleModel, vv)
}

// UpdateKnowledgeArticle updates an existing knowledge.article record.
func (c *Client) UpdateKnowledgeArticle(ka *KnowledgeArticle) error {
	return c.UpdateKnowledgeArticles([]int64{ka.Id.Get()}, ka)
}

// UpdateKnowledgeArticles updates existing knowledge.article records.
// All records (represented by ids) will be updated by ka values.
func (c *Client) UpdateKnowledgeArticles(ids []int64, ka *KnowledgeArticle) error {
	return c.Update(KnowledgeArticleModel, ids, ka)
}

// DeleteKnowledgeArticle deletes an existing knowledge.article record.
func (c *Client) DeleteKnowledgeArticle(id int64) error {
	return c.DeleteKnowledgeArticles([]int64{id})
}

// DeleteKnowledgeArticles deletes existing knowledge.article records.
func (c *Client) DeleteKnowledgeArticles(ids []int64) error {
	return c.Delete(KnowledgeArticleModel, ids)
}

// GetKnowledgeArticle gets knowledge.article existing record.
func (c *Client) GetKnowledgeArticle(id int64) (*KnowledgeArticle, error) {
	kas, err := c.GetKnowledgeArticles([]int64{id})
	if err != nil {
		return nil, err
	}
	if kas != nil && len(*kas) > 0 {
		return &((*kas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowledge.article not found", id)
}

// GetKnowledgeArticles gets knowledge.article existing records.
func (c *Client) GetKnowledgeArticles(ids []int64) (*KnowledgeArticles, error) {
	kas := &KnowledgeArticles{}
	if err := c.Read(KnowledgeArticleModel, ids, nil, kas); err != nil {
		return nil, err
	}
	return kas, nil
}

// FindKnowledgeArticle finds knowledge.article record by querying it with criteria.
func (c *Client) FindKnowledgeArticle(criteria *Criteria) (*KnowledgeArticle, error) {
	kas := &KnowledgeArticles{}
	if err := c.SearchRead(KnowledgeArticleModel, criteria, NewOptions().Limit(1), kas); err != nil {
		return nil, err
	}
	if kas != nil && len(*kas) > 0 {
		return &((*kas)[0]), nil
	}
	return nil, fmt.Errorf("knowledge.article was not found with criteria %v", criteria)
}

// FindKnowledgeArticles finds knowledge.article records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticles(criteria *Criteria, options *Options) (*KnowledgeArticles, error) {
	kas := &KnowledgeArticles{}
	if err := c.SearchRead(KnowledgeArticleModel, criteria, options, kas); err != nil {
		return nil, err
	}
	return kas, nil
}

// FindKnowledgeArticleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowledgeArticleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowledgeArticleId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowledge.article was not found with criteria %v and options %v", criteria, options)
}
