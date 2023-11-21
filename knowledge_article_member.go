package odoo

import (
	"fmt"
)

// KnowledgeArticleMember represents knowledge.article.member model.
type KnowledgeArticleMember struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	ArticleId         *Many2One  `xmlrpc:"article_id,omptempty"`
	ArticlePermission *Selection `xmlrpc:"article_permission,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omptempty"`
	Permission        *Selection `xmlrpc:"permission,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// KnowledgeArticleMembers represents array of knowledge.article.member model.
type KnowledgeArticleMembers []KnowledgeArticleMember

// KnowledgeArticleMemberModel is the odoo model name.
const KnowledgeArticleMemberModel = "knowledge.article.member"

// Many2One convert KnowledgeArticleMember to *Many2One.
func (kam *KnowledgeArticleMember) Many2One() *Many2One {
	return NewMany2One(kam.Id.Get(), "")
}

// CreateKnowledgeArticleMember creates a new knowledge.article.member model and returns its id.
func (c *Client) CreateKnowledgeArticleMember(kam *KnowledgeArticleMember) (int64, error) {
	ids, err := c.CreateKnowledgeArticleMembers([]*KnowledgeArticleMember{kam})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeArticleMember creates a new knowledge.article.member model and returns its id.
func (c *Client) CreateKnowledgeArticleMembers(kams []*KnowledgeArticleMember) ([]int64, error) {
	var vv []interface{}
	for _, v := range kams {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeArticleMemberModel, vv)
}

// UpdateKnowledgeArticleMember updates an existing knowledge.article.member record.
func (c *Client) UpdateKnowledgeArticleMember(kam *KnowledgeArticleMember) error {
	return c.UpdateKnowledgeArticleMembers([]int64{kam.Id.Get()}, kam)
}

// UpdateKnowledgeArticleMembers updates existing knowledge.article.member records.
// All records (represented by ids) will be updated by kam values.
func (c *Client) UpdateKnowledgeArticleMembers(ids []int64, kam *KnowledgeArticleMember) error {
	return c.Update(KnowledgeArticleMemberModel, ids, kam)
}

// DeleteKnowledgeArticleMember deletes an existing knowledge.article.member record.
func (c *Client) DeleteKnowledgeArticleMember(id int64) error {
	return c.DeleteKnowledgeArticleMembers([]int64{id})
}

// DeleteKnowledgeArticleMembers deletes existing knowledge.article.member records.
func (c *Client) DeleteKnowledgeArticleMembers(ids []int64) error {
	return c.Delete(KnowledgeArticleMemberModel, ids)
}

// GetKnowledgeArticleMember gets knowledge.article.member existing record.
func (c *Client) GetKnowledgeArticleMember(id int64) (*KnowledgeArticleMember, error) {
	kams, err := c.GetKnowledgeArticleMembers([]int64{id})
	if err != nil {
		return nil, err
	}
	if kams != nil && len(*kams) > 0 {
		return &((*kams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowledge.article.member not found", id)
}

// GetKnowledgeArticleMembers gets knowledge.article.member existing records.
func (c *Client) GetKnowledgeArticleMembers(ids []int64) (*KnowledgeArticleMembers, error) {
	kams := &KnowledgeArticleMembers{}
	if err := c.Read(KnowledgeArticleMemberModel, ids, nil, kams); err != nil {
		return nil, err
	}
	return kams, nil
}

// FindKnowledgeArticleMember finds knowledge.article.member record by querying it with criteria.
func (c *Client) FindKnowledgeArticleMember(criteria *Criteria) (*KnowledgeArticleMember, error) {
	kams := &KnowledgeArticleMembers{}
	if err := c.SearchRead(KnowledgeArticleMemberModel, criteria, NewOptions().Limit(1), kams); err != nil {
		return nil, err
	}
	if kams != nil && len(*kams) > 0 {
		return &((*kams)[0]), nil
	}
	return nil, fmt.Errorf("knowledge.article.member was not found with criteria %v", criteria)
}

// FindKnowledgeArticleMembers finds knowledge.article.member records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleMembers(criteria *Criteria, options *Options) (*KnowledgeArticleMembers, error) {
	kams := &KnowledgeArticleMembers{}
	if err := c.SearchRead(KnowledgeArticleMemberModel, criteria, options, kams); err != nil {
		return nil, err
	}
	return kams, nil
}

// FindKnowledgeArticleMemberIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeArticleMemberIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowledgeArticleMemberModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowledgeArticleMemberId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeArticleMemberId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeArticleMemberModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowledge.article.member was not found with criteria %v and options %v", criteria, options)
}
