package odoo

import (
	"fmt"
)

// MukRestAuthorizationCode represents muk_rest.authorization_code model.
type MukRestAuthorizationCode struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Callback    *String   `xmlrpc:"callback,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	OauthId     *Many2One `xmlrpc:"oauth_id,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
}

// MukRestAuthorizationCodes represents array of muk_rest.authorization_code model.
type MukRestAuthorizationCodes []MukRestAuthorizationCode

// MukRestAuthorizationCodeModel is the odoo model name.
const MukRestAuthorizationCodeModel = "muk_rest.authorization_code"

// Many2One convert MukRestAuthorizationCode to *Many2One.
func (ma *MukRestAuthorizationCode) Many2One() *Many2One {
	return NewMany2One(ma.Id.Get(), "")
}

// CreateMukRestAuthorizationCode creates a new muk_rest.authorization_code model and returns its id.
func (c *Client) CreateMukRestAuthorizationCode(ma *MukRestAuthorizationCode) (int64, error) {
	ids, err := c.CreateMukRestAuthorizationCodes([]*MukRestAuthorizationCode{ma})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestAuthorizationCode creates a new muk_rest.authorization_code model and returns its id.
func (c *Client) CreateMukRestAuthorizationCodes(mas []*MukRestAuthorizationCode) ([]int64, error) {
	var vv []interface{}
	for _, v := range mas {
		vv = append(vv, v)
	}
	return c.Create(MukRestAuthorizationCodeModel, vv)
}

// UpdateMukRestAuthorizationCode updates an existing muk_rest.authorization_code record.
func (c *Client) UpdateMukRestAuthorizationCode(ma *MukRestAuthorizationCode) error {
	return c.UpdateMukRestAuthorizationCodes([]int64{ma.Id.Get()}, ma)
}

// UpdateMukRestAuthorizationCodes updates existing muk_rest.authorization_code records.
// All records (represented by ids) will be updated by ma values.
func (c *Client) UpdateMukRestAuthorizationCodes(ids []int64, ma *MukRestAuthorizationCode) error {
	return c.Update(MukRestAuthorizationCodeModel, ids, ma)
}

// DeleteMukRestAuthorizationCode deletes an existing muk_rest.authorization_code record.
func (c *Client) DeleteMukRestAuthorizationCode(id int64) error {
	return c.DeleteMukRestAuthorizationCodes([]int64{id})
}

// DeleteMukRestAuthorizationCodes deletes existing muk_rest.authorization_code records.
func (c *Client) DeleteMukRestAuthorizationCodes(ids []int64) error {
	return c.Delete(MukRestAuthorizationCodeModel, ids)
}

// GetMukRestAuthorizationCode gets muk_rest.authorization_code existing record.
func (c *Client) GetMukRestAuthorizationCode(id int64) (*MukRestAuthorizationCode, error) {
	mas, err := c.GetMukRestAuthorizationCodes([]int64{id})
	if err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.authorization_code not found", id)
}

// GetMukRestAuthorizationCodes gets muk_rest.authorization_code existing records.
func (c *Client) GetMukRestAuthorizationCodes(ids []int64) (*MukRestAuthorizationCodes, error) {
	mas := &MukRestAuthorizationCodes{}
	if err := c.Read(MukRestAuthorizationCodeModel, ids, nil, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMukRestAuthorizationCode finds muk_rest.authorization_code record by querying it with criteria.
func (c *Client) FindMukRestAuthorizationCode(criteria *Criteria) (*MukRestAuthorizationCode, error) {
	mas := &MukRestAuthorizationCodes{}
	if err := c.SearchRead(MukRestAuthorizationCodeModel, criteria, NewOptions().Limit(1), mas); err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.authorization_code was not found with criteria %v", criteria)
}

// FindMukRestAuthorizationCodes finds muk_rest.authorization_code records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAuthorizationCodes(criteria *Criteria, options *Options) (*MukRestAuthorizationCodes, error) {
	mas := &MukRestAuthorizationCodes{}
	if err := c.SearchRead(MukRestAuthorizationCodeModel, criteria, options, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMukRestAuthorizationCodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAuthorizationCodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestAuthorizationCodeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestAuthorizationCodeId finds record id by querying it with criteria.
func (c *Client) FindMukRestAuthorizationCodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestAuthorizationCodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.authorization_code was not found with criteria %v and options %v", criteria, options)
}
