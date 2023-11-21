package odoo

import (
	"fmt"
)

// MukRestAccessToken represents muk_rest.access_token model.
type MukRestAccessToken struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	OauthId     *Many2One `xmlrpc:"oauth_id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
}

// MukRestAccessTokens represents array of muk_rest.access_token model.
type MukRestAccessTokens []MukRestAccessToken

// MukRestAccessTokenModel is the odoo model name.
const MukRestAccessTokenModel = "muk_rest.access_token"

// Many2One convert MukRestAccessToken to *Many2One.
func (ma *MukRestAccessToken) Many2One() *Many2One {
	return NewMany2One(ma.Id.Get(), "")
}

// CreateMukRestAccessToken creates a new muk_rest.access_token model and returns its id.
func (c *Client) CreateMukRestAccessToken(ma *MukRestAccessToken) (int64, error) {
	ids, err := c.CreateMukRestAccessTokens([]*MukRestAccessToken{ma})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestAccessToken creates a new muk_rest.access_token model and returns its id.
func (c *Client) CreateMukRestAccessTokens(mas []*MukRestAccessToken) ([]int64, error) {
	var vv []interface{}
	for _, v := range mas {
		vv = append(vv, v)
	}
	return c.Create(MukRestAccessTokenModel, vv)
}

// UpdateMukRestAccessToken updates an existing muk_rest.access_token record.
func (c *Client) UpdateMukRestAccessToken(ma *MukRestAccessToken) error {
	return c.UpdateMukRestAccessTokens([]int64{ma.Id.Get()}, ma)
}

// UpdateMukRestAccessTokens updates existing muk_rest.access_token records.
// All records (represented by ids) will be updated by ma values.
func (c *Client) UpdateMukRestAccessTokens(ids []int64, ma *MukRestAccessToken) error {
	return c.Update(MukRestAccessTokenModel, ids, ma)
}

// DeleteMukRestAccessToken deletes an existing muk_rest.access_token record.
func (c *Client) DeleteMukRestAccessToken(id int64) error {
	return c.DeleteMukRestAccessTokens([]int64{id})
}

// DeleteMukRestAccessTokens deletes existing muk_rest.access_token records.
func (c *Client) DeleteMukRestAccessTokens(ids []int64) error {
	return c.Delete(MukRestAccessTokenModel, ids)
}

// GetMukRestAccessToken gets muk_rest.access_token existing record.
func (c *Client) GetMukRestAccessToken(id int64) (*MukRestAccessToken, error) {
	mas, err := c.GetMukRestAccessTokens([]int64{id})
	if err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.access_token not found", id)
}

// GetMukRestAccessTokens gets muk_rest.access_token existing records.
func (c *Client) GetMukRestAccessTokens(ids []int64) (*MukRestAccessTokens, error) {
	mas := &MukRestAccessTokens{}
	if err := c.Read(MukRestAccessTokenModel, ids, nil, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMukRestAccessToken finds muk_rest.access_token record by querying it with criteria.
func (c *Client) FindMukRestAccessToken(criteria *Criteria) (*MukRestAccessToken, error) {
	mas := &MukRestAccessTokens{}
	if err := c.SearchRead(MukRestAccessTokenModel, criteria, NewOptions().Limit(1), mas); err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.access_token was not found with criteria %v", criteria)
}

// FindMukRestAccessTokens finds muk_rest.access_token records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAccessTokens(criteria *Criteria, options *Options) (*MukRestAccessTokens, error) {
	mas := &MukRestAccessTokens{}
	if err := c.SearchRead(MukRestAccessTokenModel, criteria, options, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMukRestAccessTokenIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAccessTokenIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestAccessTokenModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestAccessTokenId finds record id by querying it with criteria.
func (c *Client) FindMukRestAccessTokenId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestAccessTokenModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.access_token was not found with criteria %v and options %v", criteria, options)
}
