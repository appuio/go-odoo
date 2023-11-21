package odoo

import (
	"fmt"
)

// SocialTwitterAccount represents social.twitter.account model.
type SocialTwitterAccount struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	Description         *String   `xmlrpc:"description,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	Image               *String   `xmlrpc:"image,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	TwitterId           *String   `xmlrpc:"twitter_id,omptempty"`
	TwitterSearchedById *Many2One `xmlrpc:"twitter_searched_by_id,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SocialTwitterAccounts represents array of social.twitter.account model.
type SocialTwitterAccounts []SocialTwitterAccount

// SocialTwitterAccountModel is the odoo model name.
const SocialTwitterAccountModel = "social.twitter.account"

// Many2One convert SocialTwitterAccount to *Many2One.
func (sta *SocialTwitterAccount) Many2One() *Many2One {
	return NewMany2One(sta.Id.Get(), "")
}

// CreateSocialTwitterAccount creates a new social.twitter.account model and returns its id.
func (c *Client) CreateSocialTwitterAccount(sta *SocialTwitterAccount) (int64, error) {
	ids, err := c.CreateSocialTwitterAccounts([]*SocialTwitterAccount{sta})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSocialTwitterAccount creates a new social.twitter.account model and returns its id.
func (c *Client) CreateSocialTwitterAccounts(stas []*SocialTwitterAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range stas {
		vv = append(vv, v)
	}
	return c.Create(SocialTwitterAccountModel, vv)
}

// UpdateSocialTwitterAccount updates an existing social.twitter.account record.
func (c *Client) UpdateSocialTwitterAccount(sta *SocialTwitterAccount) error {
	return c.UpdateSocialTwitterAccounts([]int64{sta.Id.Get()}, sta)
}

// UpdateSocialTwitterAccounts updates existing social.twitter.account records.
// All records (represented by ids) will be updated by sta values.
func (c *Client) UpdateSocialTwitterAccounts(ids []int64, sta *SocialTwitterAccount) error {
	return c.Update(SocialTwitterAccountModel, ids, sta)
}

// DeleteSocialTwitterAccount deletes an existing social.twitter.account record.
func (c *Client) DeleteSocialTwitterAccount(id int64) error {
	return c.DeleteSocialTwitterAccounts([]int64{id})
}

// DeleteSocialTwitterAccounts deletes existing social.twitter.account records.
func (c *Client) DeleteSocialTwitterAccounts(ids []int64) error {
	return c.Delete(SocialTwitterAccountModel, ids)
}

// GetSocialTwitterAccount gets social.twitter.account existing record.
func (c *Client) GetSocialTwitterAccount(id int64) (*SocialTwitterAccount, error) {
	stas, err := c.GetSocialTwitterAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if stas != nil && len(*stas) > 0 {
		return &((*stas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of social.twitter.account not found", id)
}

// GetSocialTwitterAccounts gets social.twitter.account existing records.
func (c *Client) GetSocialTwitterAccounts(ids []int64) (*SocialTwitterAccounts, error) {
	stas := &SocialTwitterAccounts{}
	if err := c.Read(SocialTwitterAccountModel, ids, nil, stas); err != nil {
		return nil, err
	}
	return stas, nil
}

// FindSocialTwitterAccount finds social.twitter.account record by querying it with criteria.
func (c *Client) FindSocialTwitterAccount(criteria *Criteria) (*SocialTwitterAccount, error) {
	stas := &SocialTwitterAccounts{}
	if err := c.SearchRead(SocialTwitterAccountModel, criteria, NewOptions().Limit(1), stas); err != nil {
		return nil, err
	}
	if stas != nil && len(*stas) > 0 {
		return &((*stas)[0]), nil
	}
	return nil, fmt.Errorf("social.twitter.account was not found with criteria %v", criteria)
}

// FindSocialTwitterAccounts finds social.twitter.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialTwitterAccounts(criteria *Criteria, options *Options) (*SocialTwitterAccounts, error) {
	stas := &SocialTwitterAccounts{}
	if err := c.SearchRead(SocialTwitterAccountModel, criteria, options, stas); err != nil {
		return nil, err
	}
	return stas, nil
}

// FindSocialTwitterAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSocialTwitterAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SocialTwitterAccountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSocialTwitterAccountId finds record id by querying it with criteria.
func (c *Client) FindSocialTwitterAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SocialTwitterAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("social.twitter.account was not found with criteria %v and options %v", criteria, options)
}
