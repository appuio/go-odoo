package odoo

import (
	"fmt"
)

// MukRestRequestToken represents muk_rest.request_token model.
type MukRestRequestToken struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Callback    *String   `xmlrpc:"callback,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	OauthId     *Many2One `xmlrpc:"oauth_id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
}

// MukRestRequestTokens represents array of muk_rest.request_token model.
type MukRestRequestTokens []MukRestRequestToken

// MukRestRequestTokenModel is the odoo model name.
const MukRestRequestTokenModel = "muk_rest.request_token"

// Many2One convert MukRestRequestToken to *Many2One.
func (mr *MukRestRequestToken) Many2One() *Many2One {
	return NewMany2One(mr.Id.Get(), "")
}

// CreateMukRestRequestToken creates a new muk_rest.request_token model and returns its id.
func (c *Client) CreateMukRestRequestToken(mr *MukRestRequestToken) (int64, error) {
	ids, err := c.CreateMukRestRequestTokens([]*MukRestRequestToken{mr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestRequestToken creates a new muk_rest.request_token model and returns its id.
func (c *Client) CreateMukRestRequestTokens(mrs []*MukRestRequestToken) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrs {
		vv = append(vv, v)
	}
	return c.Create(MukRestRequestTokenModel, vv)
}

// UpdateMukRestRequestToken updates an existing muk_rest.request_token record.
func (c *Client) UpdateMukRestRequestToken(mr *MukRestRequestToken) error {
	return c.UpdateMukRestRequestTokens([]int64{mr.Id.Get()}, mr)
}

// UpdateMukRestRequestTokens updates existing muk_rest.request_token records.
// All records (represented by ids) will be updated by mr values.
func (c *Client) UpdateMukRestRequestTokens(ids []int64, mr *MukRestRequestToken) error {
	return c.Update(MukRestRequestTokenModel, ids, mr)
}

// DeleteMukRestRequestToken deletes an existing muk_rest.request_token record.
func (c *Client) DeleteMukRestRequestToken(id int64) error {
	return c.DeleteMukRestRequestTokens([]int64{id})
}

// DeleteMukRestRequestTokens deletes existing muk_rest.request_token records.
func (c *Client) DeleteMukRestRequestTokens(ids []int64) error {
	return c.Delete(MukRestRequestTokenModel, ids)
}

// GetMukRestRequestToken gets muk_rest.request_token existing record.
func (c *Client) GetMukRestRequestToken(id int64) (*MukRestRequestToken, error) {
	mrs, err := c.GetMukRestRequestTokens([]int64{id})
	if err != nil {
		return nil, err
	}
	if mrs != nil && len(*mrs) > 0 {
		return &((*mrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.request_token not found", id)
}

// GetMukRestRequestTokens gets muk_rest.request_token existing records.
func (c *Client) GetMukRestRequestTokens(ids []int64) (*MukRestRequestTokens, error) {
	mrs := &MukRestRequestTokens{}
	if err := c.Read(MukRestRequestTokenModel, ids, nil, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMukRestRequestToken finds muk_rest.request_token record by querying it with criteria.
func (c *Client) FindMukRestRequestToken(criteria *Criteria) (*MukRestRequestToken, error) {
	mrs := &MukRestRequestTokens{}
	if err := c.SearchRead(MukRestRequestTokenModel, criteria, NewOptions().Limit(1), mrs); err != nil {
		return nil, err
	}
	if mrs != nil && len(*mrs) > 0 {
		return &((*mrs)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.request_token was not found with criteria %v", criteria)
}

// FindMukRestRequestTokens finds muk_rest.request_token records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestRequestTokens(criteria *Criteria, options *Options) (*MukRestRequestTokens, error) {
	mrs := &MukRestRequestTokens{}
	if err := c.SearchRead(MukRestRequestTokenModel, criteria, options, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMukRestRequestTokenIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestRequestTokenIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestRequestTokenModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestRequestTokenId finds record id by querying it with criteria.
func (c *Client) FindMukRestRequestTokenId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestRequestTokenModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.request_token was not found with criteria %v and options %v", criteria, options)
}
