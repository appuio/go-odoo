package odoo

import (
	"fmt"
)

// SocialAccountRevokeYoutube represents social.account.revoke.youtube model.
type SocialAccountRevokeYoutube struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	AccountId   *Many2One `xmlrpc:"account_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SocialAccountRevokeYoutubes represents array of social.account.revoke.youtube model.
type SocialAccountRevokeYoutubes []SocialAccountRevokeYoutube

// SocialAccountRevokeYoutubeModel is the odoo model name.
const SocialAccountRevokeYoutubeModel = "social.account.revoke.youtube"

// Many2One convert SocialAccountRevokeYoutube to *Many2One.
func (sary *SocialAccountRevokeYoutube) Many2One() *Many2One {
	return NewMany2One(sary.Id.Get(), "")
}

// CreateSocialAccountRevokeYoutube creates a new social.account.revoke.youtube model and returns its id.
func (c *Client) CreateSocialAccountRevokeYoutube(sary *SocialAccountRevokeYoutube) (int64, error) {
	ids, err := c.CreateSocialAccountRevokeYoutubes([]*SocialAccountRevokeYoutube{sary})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialAccountRevokeYoutube creates a new social.account.revoke.youtube model and returns its id.
func (c *Client) CreateSocialAccountRevokeYoutubes(sarys []*SocialAccountRevokeYoutube) ([]int64, error) {
	var vv []interface{}
	for _, v := range sarys {
		vv = append(vv, v)
	}
	return c.Create(SocialAccountRevokeYoutubeModel, vv)
}

// UpdateSocialAccountRevokeYoutube updates an existing social.account.revoke.youtube record.
func (c *Client) UpdateSocialAccountRevokeYoutube(sary *SocialAccountRevokeYoutube) error {
	return c.UpdateSocialAccountRevokeYoutubes([]int64{sary.Id.Get()}, sary)
}

// UpdateSocialAccountRevokeYoutubes updates existing social.account.revoke.youtube records.
// All records (represented by ids) will be updated by sary values.
func (c *Client) UpdateSocialAccountRevokeYoutubes(ids []int64, sary *SocialAccountRevokeYoutube) error {
	return c.Update(SocialAccountRevokeYoutubeModel, ids, sary)
}

// DeleteSocialAccountRevokeYoutube deletes an existing social.account.revoke.youtube record.
func (c *Client) DeleteSocialAccountRevokeYoutube(id int64) error {
	return c.DeleteSocialAccountRevokeYoutubes([]int64{id})
}

// DeleteSocialAccountRevokeYoutubes deletes existing social.account.revoke.youtube records.
func (c *Client) DeleteSocialAccountRevokeYoutubes(ids []int64) error {
	return c.Delete(SocialAccountRevokeYoutubeModel, ids)
}

// GetSocialAccountRevokeYoutube gets social.account.revoke.youtube existing record.
func (c *Client) GetSocialAccountRevokeYoutube(id int64) (*SocialAccountRevokeYoutube, error) {
	sarys, err := c.GetSocialAccountRevokeYoutubes([]int64{id})
	if err != nil {
		return nil, err
	}
	if sarys != nil && len(*sarys) > 0 {
		return &((*sarys)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.account.revoke.youtube not found", id)
}

// GetSocialAccountRevokeYoutubes gets social.account.revoke.youtube existing records.
func (c *Client) GetSocialAccountRevokeYoutubes(ids []int64) (*SocialAccountRevokeYoutubes, error) {
	sarys := &SocialAccountRevokeYoutubes{}
	if err := c.Read(SocialAccountRevokeYoutubeModel, ids, nil, sarys); err != nil {
		return nil, err
	}
	return sarys, nil
}

// FindSocialAccountRevokeYoutube finds social.account.revoke.youtube record by querying it with criteria.
func (c *Client) FindSocialAccountRevokeYoutube(criteria *Criteria) (*SocialAccountRevokeYoutube, error) {
	sarys := &SocialAccountRevokeYoutubes{}
	if err := c.SearchRead(SocialAccountRevokeYoutubeModel, criteria, NewOptions().Limit(1), sarys); err != nil {
		return nil, err
	}
	if sarys != nil && len(*sarys) > 0 {
		return &((*sarys)[0]), nil
	}
	return nil, fmt.Errorf("social.account.revoke.youtube was not found with criteria %v", criteria)
}

// FindSocialAccountRevokeYoutubes finds social.account.revoke.youtube records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialAccountRevokeYoutubes(criteria *Criteria, options *Options) (*SocialAccountRevokeYoutubes, error) {
	sarys := &SocialAccountRevokeYoutubes{}
	if err := c.SearchRead(SocialAccountRevokeYoutubeModel, criteria, options, sarys); err != nil {
		return nil, err
	}
	return sarys, nil
}

// FindSocialAccountRevokeYoutubeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialAccountRevokeYoutubeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialAccountRevokeYoutubeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialAccountRevokeYoutubeId finds record id by querying it with criteria.
func (c *Client) FindSocialAccountRevokeYoutubeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialAccountRevokeYoutubeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.account.revoke.youtube was not found with criteria %v and options %v", criteria, options)
}
