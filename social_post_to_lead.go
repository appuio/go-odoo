package odoo

import (
	"fmt"
)

// SocialPostToLead represents social.post.to.lead model.
type SocialPostToLead struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	Action             *Selection `xmlrpc:"action,omptempty"`
	AuthorName         *String    `xmlrpc:"author_name,omptempty"`
	ConversionSource   *Selection `xmlrpc:"conversion_source,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	PartnerId          *Many2One  `xmlrpc:"partner_id,omptempty"`
	PostContent        *String    `xmlrpc:"post_content,omptempty"`
	PostDatetime       *Time      `xmlrpc:"post_datetime,omptempty"`
	PostLink           *String    `xmlrpc:"post_link,omptempty"`
	SocialAccountId    *Many2One  `xmlrpc:"social_account_id,omptempty"`
	SocialStreamPostId *Many2One  `xmlrpc:"social_stream_post_id,omptempty"`
	UtmCampaignId      *Many2One  `xmlrpc:"utm_campaign_id,omptempty"`
	UtmMediumId        *Many2One  `xmlrpc:"utm_medium_id,omptempty"`
	UtmSourceId        *Many2One  `xmlrpc:"utm_source_id,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SocialPostToLeads represents array of social.post.to.lead model.
type SocialPostToLeads []SocialPostToLead

// SocialPostToLeadModel is the odoo model name.
const SocialPostToLeadModel = "social.post.to.lead"

// Many2One convert SocialPostToLead to *Many2One.
func (sptl *SocialPostToLead) Many2One() *Many2One {
	return NewMany2One(sptl.Id.Get(), "")
}

// CreateSocialPostToLead creates a new social.post.to.lead model and returns its id.
func (c *Client) CreateSocialPostToLead(sptl *SocialPostToLead) (int64, error) {
	ids, err := c.CreateSocialPostToLeads([]*SocialPostToLead{sptl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialPostToLead creates a new social.post.to.lead model and returns its id.
func (c *Client) CreateSocialPostToLeads(sptls []*SocialPostToLead) ([]int64, error) {
	var vv []interface{}
	for _, v := range sptls {
		vv = append(vv, v)
	}
	return c.Create(SocialPostToLeadModel, vv)
}

// UpdateSocialPostToLead updates an existing social.post.to.lead record.
func (c *Client) UpdateSocialPostToLead(sptl *SocialPostToLead) error {
	return c.UpdateSocialPostToLeads([]int64{sptl.Id.Get()}, sptl)
}

// UpdateSocialPostToLeads updates existing social.post.to.lead records.
// All records (represented by ids) will be updated by sptl values.
func (c *Client) UpdateSocialPostToLeads(ids []int64, sptl *SocialPostToLead) error {
	return c.Update(SocialPostToLeadModel, ids, sptl)
}

// DeleteSocialPostToLead deletes an existing social.post.to.lead record.
func (c *Client) DeleteSocialPostToLead(id int64) error {
	return c.DeleteSocialPostToLeads([]int64{id})
}

// DeleteSocialPostToLeads deletes existing social.post.to.lead records.
func (c *Client) DeleteSocialPostToLeads(ids []int64) error {
	return c.Delete(SocialPostToLeadModel, ids)
}

// GetSocialPostToLead gets social.post.to.lead existing record.
func (c *Client) GetSocialPostToLead(id int64) (*SocialPostToLead, error) {
	sptls, err := c.GetSocialPostToLeads([]int64{id})
	if err != nil {
		return nil, err
	}
	if sptls != nil && len(*sptls) > 0 {
		return &((*sptls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.post.to.lead not found", id)
}

// GetSocialPostToLeads gets social.post.to.lead existing records.
func (c *Client) GetSocialPostToLeads(ids []int64) (*SocialPostToLeads, error) {
	sptls := &SocialPostToLeads{}
	if err := c.Read(SocialPostToLeadModel, ids, nil, sptls); err != nil {
		return nil, err
	}
	return sptls, nil
}

// FindSocialPostToLead finds social.post.to.lead record by querying it with criteria.
func (c *Client) FindSocialPostToLead(criteria *Criteria) (*SocialPostToLead, error) {
	sptls := &SocialPostToLeads{}
	if err := c.SearchRead(SocialPostToLeadModel, criteria, NewOptions().Limit(1), sptls); err != nil {
		return nil, err
	}
	if sptls != nil && len(*sptls) > 0 {
		return &((*sptls)[0]), nil
	}
	return nil, fmt.Errorf("social.post.to.lead was not found with criteria %v", criteria)
}

// FindSocialPostToLeads finds social.post.to.lead records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialPostToLeads(criteria *Criteria, options *Options) (*SocialPostToLeads, error) {
	sptls := &SocialPostToLeads{}
	if err := c.SearchRead(SocialPostToLeadModel, criteria, options, sptls); err != nil {
		return nil, err
	}
	return sptls, nil
}

// FindSocialPostToLeadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialPostToLeadIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialPostToLeadModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialPostToLeadId finds record id by querying it with criteria.
func (c *Client) FindSocialPostToLeadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialPostToLeadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.post.to.lead was not found with criteria %v and options %v", criteria, options)
}
