package odoo

import (
	"fmt"
)

// SocialStream represents social.stream model.
type SocialStream struct {
	LastUpdate                   *Time     `xmlrpc:"__last_update,omptempty"`
	AccountId                    *Many2One `xmlrpc:"account_id,omptempty"`
	CompanyId                    *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omptempty"`
	Id                           *Int      `xmlrpc:"id,omptempty"`
	MediaId                      *Many2One `xmlrpc:"media_id,omptempty"`
	MediaImage                   *String   `xmlrpc:"media_image,omptempty"`
	Name                         *String   `xmlrpc:"name,omptempty"`
	Sequence                     *Int      `xmlrpc:"sequence,omptempty"`
	StreamPostIds                *Relation `xmlrpc:"stream_post_ids,omptempty"`
	StreamTypeId                 *Many2One `xmlrpc:"stream_type_id,omptempty"`
	StreamTypeType               *String   `xmlrpc:"stream_type_type,omptempty"`
	TwitterFollowedAccountId     *Many2One `xmlrpc:"twitter_followed_account_id,omptempty"`
	TwitterFollowedAccountSearch *String   `xmlrpc:"twitter_followed_account_search,omptempty"`
	TwitterSearchedKeyword       *String   `xmlrpc:"twitter_searched_keyword,omptempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SocialStreams represents array of social.stream model.
type SocialStreams []SocialStream

// SocialStreamModel is the odoo model name.
const SocialStreamModel = "social.stream"

// Many2One convert SocialStream to *Many2One.
func (ss *SocialStream) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSocialStream creates a new social.stream model and returns its id.
func (c *Client) CreateSocialStream(ss *SocialStream) (int64, error) {
	ids, err := c.CreateSocialStreams([]*SocialStream{ss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialStream creates a new social.stream model and returns its id.
func (c *Client) CreateSocialStreams(sss []*SocialStream) ([]int64, error) {
	var vv []interface{}
	for _, v := range sss {
		vv = append(vv, v)
	}
	return c.Create(SocialStreamModel, vv)
}

// UpdateSocialStream updates an existing social.stream record.
func (c *Client) UpdateSocialStream(ss *SocialStream) error {
	return c.UpdateSocialStreams([]int64{ss.Id.Get()}, ss)
}

// UpdateSocialStreams updates existing social.stream records.
// All records (represented by ids) will be updated by ss values.
func (c *Client) UpdateSocialStreams(ids []int64, ss *SocialStream) error {
	return c.Update(SocialStreamModel, ids, ss)
}

// DeleteSocialStream deletes an existing social.stream record.
func (c *Client) DeleteSocialStream(id int64) error {
	return c.DeleteSocialStreams([]int64{id})
}

// DeleteSocialStreams deletes existing social.stream records.
func (c *Client) DeleteSocialStreams(ids []int64) error {
	return c.Delete(SocialStreamModel, ids)
}

// GetSocialStream gets social.stream existing record.
func (c *Client) GetSocialStream(id int64) (*SocialStream, error) {
	sss, err := c.GetSocialStreams([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.stream not found", id)
}

// GetSocialStreams gets social.stream existing records.
func (c *Client) GetSocialStreams(ids []int64) (*SocialStreams, error) {
	sss := &SocialStreams{}
	if err := c.Read(SocialStreamModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSocialStream finds social.stream record by querying it with criteria.
func (c *Client) FindSocialStream(criteria *Criteria) (*SocialStream, error) {
	sss := &SocialStreams{}
	if err := c.SearchRead(SocialStreamModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("social.stream was not found with criteria %v", criteria)
}

// FindSocialStreams finds social.stream records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreams(criteria *Criteria, options *Options) (*SocialStreams, error) {
	sss := &SocialStreams{}
	if err := c.SearchRead(SocialStreamModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSocialStreamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialStreamIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialStreamModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialStreamId finds record id by querying it with criteria.
func (c *Client) FindSocialStreamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialStreamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.stream was not found with criteria %v and options %v", criteria, options)
}
