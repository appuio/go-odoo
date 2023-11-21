package odoo

import (
	"fmt"
)

// KnowledgeCover represents knowledge.cover model.
type KnowledgeCover struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	ArticleIds    *Relation `xmlrpc:"article_ids,omptempty"`
	AttachmentId  *Many2One `xmlrpc:"attachment_id,omptempty"`
	AttachmentUrl *String   `xmlrpc:"attachment_url,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// KnowledgeCovers represents array of knowledge.cover model.
type KnowledgeCovers []KnowledgeCover

// KnowledgeCoverModel is the odoo model name.
const KnowledgeCoverModel = "knowledge.cover"

// Many2One convert KnowledgeCover to *Many2One.
func (kc *KnowledgeCover) Many2One() *Many2One {
	return NewMany2One(kc.Id.Get(), "")
}

// CreateKnowledgeCover creates a new knowledge.cover model and returns its id.
func (c *Client) CreateKnowledgeCover(kc *KnowledgeCover) (int64, error) {
	ids, err := c.CreateKnowledgeCovers([]*KnowledgeCover{kc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKnowledgeCover creates a new knowledge.cover model and returns its id.
func (c *Client) CreateKnowledgeCovers(kcs []*KnowledgeCover) ([]int64, error) {
	var vv []interface{}
	for _, v := range kcs {
		vv = append(vv, v)
	}
	return c.Create(KnowledgeCoverModel, vv)
}

// UpdateKnowledgeCover updates an existing knowledge.cover record.
func (c *Client) UpdateKnowledgeCover(kc *KnowledgeCover) error {
	return c.UpdateKnowledgeCovers([]int64{kc.Id.Get()}, kc)
}

// UpdateKnowledgeCovers updates existing knowledge.cover records.
// All records (represented by ids) will be updated by kc values.
func (c *Client) UpdateKnowledgeCovers(ids []int64, kc *KnowledgeCover) error {
	return c.Update(KnowledgeCoverModel, ids, kc)
}

// DeleteKnowledgeCover deletes an existing knowledge.cover record.
func (c *Client) DeleteKnowledgeCover(id int64) error {
	return c.DeleteKnowledgeCovers([]int64{id})
}

// DeleteKnowledgeCovers deletes existing knowledge.cover records.
func (c *Client) DeleteKnowledgeCovers(ids []int64) error {
	return c.Delete(KnowledgeCoverModel, ids)
}

// GetKnowledgeCover gets knowledge.cover existing record.
func (c *Client) GetKnowledgeCover(id int64) (*KnowledgeCover, error) {
	kcs, err := c.GetKnowledgeCovers([]int64{id})
	if err != nil {
		return nil, err
	}
	if kcs != nil && len(*kcs) > 0 {
		return &((*kcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of knowledge.cover not found", id)
}

// GetKnowledgeCovers gets knowledge.cover existing records.
func (c *Client) GetKnowledgeCovers(ids []int64) (*KnowledgeCovers, error) {
	kcs := &KnowledgeCovers{}
	if err := c.Read(KnowledgeCoverModel, ids, nil, kcs); err != nil {
		return nil, err
	}
	return kcs, nil
}

// FindKnowledgeCover finds knowledge.cover record by querying it with criteria.
func (c *Client) FindKnowledgeCover(criteria *Criteria) (*KnowledgeCover, error) {
	kcs := &KnowledgeCovers{}
	if err := c.SearchRead(KnowledgeCoverModel, criteria, NewOptions().Limit(1), kcs); err != nil {
		return nil, err
	}
	if kcs != nil && len(*kcs) > 0 {
		return &((*kcs)[0]), nil
	}
	return nil, fmt.Errorf("knowledge.cover was not found with criteria %v", criteria)
}

// FindKnowledgeCovers finds knowledge.cover records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeCovers(criteria *Criteria, options *Options) (*KnowledgeCovers, error) {
	kcs := &KnowledgeCovers{}
	if err := c.SearchRead(KnowledgeCoverModel, criteria, options, kcs); err != nil {
		return nil, err
	}
	return kcs, nil
}

// FindKnowledgeCoverIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKnowledgeCoverIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(KnowledgeCoverModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindKnowledgeCoverId finds record id by querying it with criteria.
func (c *Client) FindKnowledgeCoverId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KnowledgeCoverModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("knowledge.cover was not found with criteria %v and options %v", criteria, options)
}
