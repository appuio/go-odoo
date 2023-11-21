package odoo

import (
	"fmt"
)

// MukRestBearerToken represents muk_rest.bearer_token model.
type MukRestBearerToken struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	ExpirationDate *Time     `xmlrpc:"expiration_date,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	OauthId        *Many2One `xmlrpc:"oauth_id,omptempty"`
	UserId         *Many2One `xmlrpc:"user_id,omptempty"`
}

// MukRestBearerTokens represents array of muk_rest.bearer_token model.
type MukRestBearerTokens []MukRestBearerToken

// MukRestBearerTokenModel is the odoo model name.
const MukRestBearerTokenModel = "muk_rest.bearer_token"

// Many2One convert MukRestBearerToken to *Many2One.
func (mb *MukRestBearerToken) Many2One() *Many2One {
	return NewMany2One(mb.Id.Get(), "")
}

// CreateMukRestBearerToken creates a new muk_rest.bearer_token model and returns its id.
func (c *Client) CreateMukRestBearerToken(mb *MukRestBearerToken) (int64, error) {
	ids, err := c.CreateMukRestBearerTokens([]*MukRestBearerToken{mb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestBearerToken creates a new muk_rest.bearer_token model and returns its id.
func (c *Client) CreateMukRestBearerTokens(mbs []*MukRestBearerToken) ([]int64, error) {
	var vv []interface{}
	for _, v := range mbs {
		vv = append(vv, v)
	}
	return c.Create(MukRestBearerTokenModel, vv)
}

// UpdateMukRestBearerToken updates an existing muk_rest.bearer_token record.
func (c *Client) UpdateMukRestBearerToken(mb *MukRestBearerToken) error {
	return c.UpdateMukRestBearerTokens([]int64{mb.Id.Get()}, mb)
}

// UpdateMukRestBearerTokens updates existing muk_rest.bearer_token records.
// All records (represented by ids) will be updated by mb values.
func (c *Client) UpdateMukRestBearerTokens(ids []int64, mb *MukRestBearerToken) error {
	return c.Update(MukRestBearerTokenModel, ids, mb)
}

// DeleteMukRestBearerToken deletes an existing muk_rest.bearer_token record.
func (c *Client) DeleteMukRestBearerToken(id int64) error {
	return c.DeleteMukRestBearerTokens([]int64{id})
}

// DeleteMukRestBearerTokens deletes existing muk_rest.bearer_token records.
func (c *Client) DeleteMukRestBearerTokens(ids []int64) error {
	return c.Delete(MukRestBearerTokenModel, ids)
}

// GetMukRestBearerToken gets muk_rest.bearer_token existing record.
func (c *Client) GetMukRestBearerToken(id int64) (*MukRestBearerToken, error) {
	mbs, err := c.GetMukRestBearerTokens([]int64{id})
	if err != nil {
		return nil, err
	}
	if mbs != nil && len(*mbs) > 0 {
		return &((*mbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.bearer_token not found", id)
}

// GetMukRestBearerTokens gets muk_rest.bearer_token existing records.
func (c *Client) GetMukRestBearerTokens(ids []int64) (*MukRestBearerTokens, error) {
	mbs := &MukRestBearerTokens{}
	if err := c.Read(MukRestBearerTokenModel, ids, nil, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMukRestBearerToken finds muk_rest.bearer_token record by querying it with criteria.
func (c *Client) FindMukRestBearerToken(criteria *Criteria) (*MukRestBearerToken, error) {
	mbs := &MukRestBearerTokens{}
	if err := c.SearchRead(MukRestBearerTokenModel, criteria, NewOptions().Limit(1), mbs); err != nil {
		return nil, err
	}
	if mbs != nil && len(*mbs) > 0 {
		return &((*mbs)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.bearer_token was not found with criteria %v", criteria)
}

// FindMukRestBearerTokens finds muk_rest.bearer_token records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestBearerTokens(criteria *Criteria, options *Options) (*MukRestBearerTokens, error) {
	mbs := &MukRestBearerTokens{}
	if err := c.SearchRead(MukRestBearerTokenModel, criteria, options, mbs); err != nil {
		return nil, err
	}
	return mbs, nil
}

// FindMukRestBearerTokenIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestBearerTokenIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestBearerTokenModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestBearerTokenId finds record id by querying it with criteria.
func (c *Client) FindMukRestBearerTokenId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestBearerTokenModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.bearer_token was not found with criteria %v and options %v", criteria, options)
}
