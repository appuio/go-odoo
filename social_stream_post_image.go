package odoo

import (
	"fmt"
)

// SocialStreamPostImage represents social.stream.post.image model.
type SocialStreamPostImage struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	ImageUrl     *String   `xmlrpc:"image_url,omptempty"`
	StreamPostId *Many2One `xmlrpc:"stream_post_id,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SocialStreamPostImages represents array of social.stream.post.image model.
type SocialStreamPostImages []SocialStreamPostImage

// SocialStreamPostImageModel is the odoo model name.
const SocialStreamPostImageModel = "social.stream.post.image"

// Many2One convert SocialStreamPostImage to *Many2One.
func (sspi *SocialStreamPostImage) Many2One() *Many2One {
	return NewMany2One(sspi.Id.Get(), "")
}

// CreateSocialStreamPostImage creates a new social.stream.post.image model and returns its id.
func (c *Client) CreateSocialStreamPostImage(sspi *SocialStreamPostImage) (int64, error) {
	ids, err := c.CreateSocialStreamPostImages([]*SocialStreamPostImage{sspi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialStreamPostImage creates a new social.stream.post.image model and returns its id.
func (c *Client) CreateSocialStreamPostImages(sspis []*SocialStreamPostImage) ([]int64, error) {
	var vv []interface{}
	for _, v := range sspis {
		vv = append(vv, v)
	}
	return c.Create(SocialStreamPostImageModel, vv)
}

// UpdateSocialStreamPostImage updates an existing social.stream.post.image record.
func (c *Client) UpdateSocialStreamPostImage(sspi *SocialStreamPostImage) error {
	return c.UpdateSocialStreamPostImages([]int64{sspi.Id.Get()}, sspi)
}

// UpdateSocialStreamPostImages updates existing social.stream.post.image records.
// All records (represented by ids) will be updated by sspi values.
func (c *Client) UpdateSocialStreamPostImages(ids []int64, sspi *SocialStreamPostImage) error {
	return c.Update(SocialStreamPostImageModel, ids, sspi)
}

// DeleteSocialStreamPostImage deletes an existing social.stream.post.image record.
func (c *Client) DeleteSocialStreamPostImage(id int64) error {
	return c.DeleteSocialStreamPostImages([]int64{id})
}

// DeleteSocialStreamPostImages deletes existing social.stream.post.image records.
func (c *Client) DeleteSocialStreamPostImages(ids []int64) error {
	return c.Delete(SocialStreamPostImageModel, ids)
}

// GetSocialStreamPostImage gets social.stream.post.image existing record.
func (c *Client) GetSocialStreamPostImage(id int64) (*SocialStreamPostImage, error) {
	sspis, err := c.GetSocialStreamPostImages([]int64{id})
	if err != nil {
		return nil, err
	}
	if sspis != nil && len(*sspis) > 0 {
		return &((*sspis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.stream.post.image not found", id)
}

// GetSocialStreamPostImages gets social.stream.post.image existing records.
func (c *Client) GetSocialStreamPostImages(ids []int64) (*SocialStreamPostImages, error) {
	sspis := &SocialStreamPostImages{}
	if err := c.Read(SocialStreamPostImageModel, ids, nil, sspis); err != nil {
		return nil, err
	}
	return sspis, nil
}

// FindSocialStreamPostImage finds social.stream.post.image record by querying it with criteria.
func (c *Client) FindSocialStreamPostImage(criteria *Criteria) (*SocialStreamPostImage, error) {
	sspis := &SocialStreamPostImages{}
	if err := c.SearchRead(SocialStreamPostImageModel, criteria, NewOptions().Limit(1), sspis); err != nil {
		return nil, err
	}
	if sspis != nil && len(*sspis) > 0 {
		return &((*sspis)[0]), nil
	}
	return nil, fmt.Errorf("social.stream.post.image was not found with criteria %v", criteria)
}

// FindSocialStreamPostImages finds social.stream.post.image records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreamPostImages(criteria *Criteria, options *Options) (*SocialStreamPostImages, error) {
	sspis := &SocialStreamPostImages{}
	if err := c.SearchRead(SocialStreamPostImageModel, criteria, options, sspis); err != nil {
		return nil, err
	}
	return sspis, nil
}

// FindSocialStreamPostImageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreamPostImageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialStreamPostImageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialStreamPostImageId finds record id by querying it with criteria.
func (c *Client) FindSocialStreamPostImageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialStreamPostImageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.stream.post.image was not found with criteria %v and options %v", criteria, options)
}
