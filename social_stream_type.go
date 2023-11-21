package odoo

import (
	"fmt"
)

// SocialStreamType represents social.stream.type model.
type SocialStreamType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MediaId     *Many2One `xmlrpc:"media_id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	StreamType  *String   `xmlrpc:"stream_type,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SocialStreamTypes represents array of social.stream.type model.
type SocialStreamTypes []SocialStreamType

// SocialStreamTypeModel is the odoo model name.
const SocialStreamTypeModel = "social.stream.type"

// Many2One convert SocialStreamType to *Many2One.
func (sst *SocialStreamType) Many2One() *Many2One {
	return NewMany2One(sst.Id.Get(), "")
}

// CreateSocialStreamType creates a new social.stream.type model and returns its id.
func (c *Client) CreateSocialStreamType(sst *SocialStreamType) (int64, error) {
	ids, err := c.CreateSocialStreamTypes([]*SocialStreamType{sst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialStreamType creates a new social.stream.type model and returns its id.
func (c *Client) CreateSocialStreamTypes(ssts []*SocialStreamType) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssts {
		vv = append(vv, v)
	}
	return c.Create(SocialStreamTypeModel, vv)
}

// UpdateSocialStreamType updates an existing social.stream.type record.
func (c *Client) UpdateSocialStreamType(sst *SocialStreamType) error {
	return c.UpdateSocialStreamTypes([]int64{sst.Id.Get()}, sst)
}

// UpdateSocialStreamTypes updates existing social.stream.type records.
// All records (represented by ids) will be updated by sst values.
func (c *Client) UpdateSocialStreamTypes(ids []int64, sst *SocialStreamType) error {
	return c.Update(SocialStreamTypeModel, ids, sst)
}

// DeleteSocialStreamType deletes an existing social.stream.type record.
func (c *Client) DeleteSocialStreamType(id int64) error {
	return c.DeleteSocialStreamTypes([]int64{id})
}

// DeleteSocialStreamTypes deletes existing social.stream.type records.
func (c *Client) DeleteSocialStreamTypes(ids []int64) error {
	return c.Delete(SocialStreamTypeModel, ids)
}

// GetSocialStreamType gets social.stream.type existing record.
func (c *Client) GetSocialStreamType(id int64) (*SocialStreamType, error) {
	ssts, err := c.GetSocialStreamTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssts != nil && len(*ssts) > 0 {
		return &((*ssts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.stream.type not found", id)
}

// GetSocialStreamTypes gets social.stream.type existing records.
func (c *Client) GetSocialStreamTypes(ids []int64) (*SocialStreamTypes, error) {
	ssts := &SocialStreamTypes{}
	if err := c.Read(SocialStreamTypeModel, ids, nil, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSocialStreamType finds social.stream.type record by querying it with criteria.
func (c *Client) FindSocialStreamType(criteria *Criteria) (*SocialStreamType, error) {
	ssts := &SocialStreamTypes{}
	if err := c.SearchRead(SocialStreamTypeModel, criteria, NewOptions().Limit(1), ssts); err != nil {
		return nil, err
	}
	if ssts != nil && len(*ssts) > 0 {
		return &((*ssts)[0]), nil
	}
	return nil, fmt.Errorf("social.stream.type was not found with criteria %v", criteria)
}

// FindSocialStreamTypes finds social.stream.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreamTypes(criteria *Criteria, options *Options) (*SocialStreamTypes, error) {
	ssts := &SocialStreamTypes{}
	if err := c.SearchRead(SocialStreamTypeModel, criteria, options, ssts); err != nil {
		return nil, err
	}
	return ssts, nil
}

// FindSocialStreamTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreamTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialStreamTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialStreamTypeId finds record id by querying it with criteria.
func (c *Client) FindSocialStreamTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialStreamTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.stream.type was not found with criteria %v and options %v", criteria, options)
}
