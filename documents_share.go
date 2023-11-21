package odoo

import (
	"fmt"
)

// DocumentsShare represents documents.share model.
type DocumentsShare struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                   *String    `xmlrpc:"access_token,omptempty"`
	Action                        *Selection `xmlrpc:"action,omptempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omptempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omptempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omptempty"`
	ActivityOption                *Bool      `xmlrpc:"activity_option,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AliasBouncedContent           *String    `xmlrpc:"alias_bounced_content,omptempty"`
	AliasContact                  *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults                 *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain                   *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId            *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                       *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId                  *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName                     *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId            *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId           *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId                   *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	CanUpload                     *Bool      `xmlrpc:"can_upload,omptempty"`
	CheckSpam                     *Bool      `xmlrpc:"check_spam,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateDeadline                  *Time      `xmlrpc:"date_deadline,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	DocumentIds                   *Relation  `xmlrpc:"document_ids,omptempty"`
	Domain                        *String    `xmlrpc:"domain,omptempty"`
	EmailDrop                     *Bool      `xmlrpc:"email_drop,omptempty"`
	FailedMessageIds              *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FolderId                      *Many2One  `xmlrpc:"folder_id,omptempty"`
	FullUrl                       *String    `xmlrpc:"full_url,omptempty"`
	HasMessage                    *Bool      `xmlrpc:"has_message,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	IncludeSubFolders             *Bool      `xmlrpc:"include_sub_folders,omptempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError            *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	OwnerId                       *Many2One  `xmlrpc:"owner_id,omptempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omptempty"`
	State                         *Selection `xmlrpc:"state,omptempty"`
	TagIds                        *Relation  `xmlrpc:"tag_ids,omptempty"`
	Type                          *Selection `xmlrpc:"type,omptempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DocumentsShares represents array of documents.share model.
type DocumentsShares []DocumentsShare

// DocumentsShareModel is the odoo model name.
const DocumentsShareModel = "documents.share"

// Many2One convert DocumentsShare to *Many2One.
func (ds *DocumentsShare) Many2One() *Many2One {
	return NewMany2One(ds.Id.Get(), "")
}

// CreateDocumentsShare creates a new documents.share model and returns its id.
func (c *Client) CreateDocumentsShare(ds *DocumentsShare) (int64, error) {
	ids, err := c.CreateDocumentsShares([]*DocumentsShare{ds})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsShare creates a new documents.share model and returns its id.
func (c *Client) CreateDocumentsShares(dss []*DocumentsShare) ([]int64, error) {
	var vv []interface{}
	for _, v := range dss {
		vv = append(vv, v)
	}
	return c.Create(DocumentsShareModel, vv)
}

// UpdateDocumentsShare updates an existing documents.share record.
func (c *Client) UpdateDocumentsShare(ds *DocumentsShare) error {
	return c.UpdateDocumentsShares([]int64{ds.Id.Get()}, ds)
}

// UpdateDocumentsShares updates existing documents.share records.
// All records (represented by ids) will be updated by ds values.
func (c *Client) UpdateDocumentsShares(ids []int64, ds *DocumentsShare) error {
	return c.Update(DocumentsShareModel, ids, ds)
}

// DeleteDocumentsShare deletes an existing documents.share record.
func (c *Client) DeleteDocumentsShare(id int64) error {
	return c.DeleteDocumentsShares([]int64{id})
}

// DeleteDocumentsShares deletes existing documents.share records.
func (c *Client) DeleteDocumentsShares(ids []int64) error {
	return c.Delete(DocumentsShareModel, ids)
}

// GetDocumentsShare gets documents.share existing record.
func (c *Client) GetDocumentsShare(id int64) (*DocumentsShare, error) {
	dss, err := c.GetDocumentsShares([]int64{id})
	if err != nil {
		return nil, err
	}
	if dss != nil && len(*dss) > 0 {
		return &((*dss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.share not found", id)
}

// GetDocumentsShares gets documents.share existing records.
func (c *Client) GetDocumentsShares(ids []int64) (*DocumentsShares, error) {
	dss := &DocumentsShares{}
	if err := c.Read(DocumentsShareModel, ids, nil, dss); err != nil {
		return nil, err
	}
	return dss, nil
}

// FindDocumentsShare finds documents.share record by querying it with criteria.
func (c *Client) FindDocumentsShare(criteria *Criteria) (*DocumentsShare, error) {
	dss := &DocumentsShares{}
	if err := c.SearchRead(DocumentsShareModel, criteria, NewOptions().Limit(1), dss); err != nil {
		return nil, err
	}
	if dss != nil && len(*dss) > 0 {
		return &((*dss)[0]), nil
	}
	return nil, fmt.Errorf("documents.share was not found with criteria %v", criteria)
}

// FindDocumentsShares finds documents.share records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsShares(criteria *Criteria, options *Options) (*DocumentsShares, error) {
	dss := &DocumentsShares{}
	if err := c.SearchRead(DocumentsShareModel, criteria, options, dss); err != nil {
		return nil, err
	}
	return dss, nil
}

// FindDocumentsShareIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsShareIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsShareModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsShareId finds record id by querying it with criteria.
func (c *Client) FindDocumentsShareId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsShareModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.share was not found with criteria %v and options %v", criteria, options)
}
