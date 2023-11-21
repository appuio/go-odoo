package odoo

import (
	"fmt"
)

// MukRestOauth2 represents muk_rest.oauth2 model.
type MukRestOauth2 struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Active            *Bool      `xmlrpc:"active,omptempty"`
	CallbackIds       *Relation  `xmlrpc:"callback_ids,omptempty"`
	ClientId          *String    `xmlrpc:"client_id,omptempty"`
	ClientSecret      *String    `xmlrpc:"client_secret,omptempty"`
	Color             *Int       `xmlrpc:"color,omptempty"`
	Company           *String    `xmlrpc:"company,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultCallbackId *Many2One  `xmlrpc:"default_callback_id,omptempty"`
	Description       *String    `xmlrpc:"description,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Homepage          *String    `xmlrpc:"homepage,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	LogoUrl           *String    `xmlrpc:"logo_url,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	Oauth1Ids         *Relation  `xmlrpc:"oauth1_ids,omptempty"`
	Oauth2Ids         *Relation  `xmlrpc:"oauth2_ids,omptempty"`
	OauthId           *Many2One  `xmlrpc:"oauth_id,omptempty"`
	PrivacyPolicy     *String    `xmlrpc:"privacy_policy,omptempty"`
	RuleIds           *Relation  `xmlrpc:"rule_ids,omptempty"`
	Security          *Selection `xmlrpc:"security,omptempty"`
	ServiceTerms      *String    `xmlrpc:"service_terms,omptempty"`
	Sessions          *Int       `xmlrpc:"sessions,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MukRestOauth2s represents array of muk_rest.oauth2 model.
type MukRestOauth2s []MukRestOauth2

// MukRestOauth2Model is the odoo model name.
const MukRestOauth2Model = "muk_rest.oauth2"

// Many2One convert MukRestOauth2 to *Many2One.
func (mo *MukRestOauth2) Many2One() *Many2One {
	return NewMany2One(mo.Id.Get(), "")
}

// CreateMukRestOauth2 creates a new muk_rest.oauth2 model and returns its id.
func (c *Client) CreateMukRestOauth2(mo *MukRestOauth2) (int64, error) {
	ids, err := c.CreateMukRestOauth2s([]*MukRestOauth2{mo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestOauth2 creates a new muk_rest.oauth2 model and returns its id.
func (c *Client) CreateMukRestOauth2s(mos []*MukRestOauth2) ([]int64, error) {
	var vv []interface{}
	for _, v := range mos {
		vv = append(vv, v)
	}
	return c.Create(MukRestOauth2Model, vv)
}

// UpdateMukRestOauth2 updates an existing muk_rest.oauth2 record.
func (c *Client) UpdateMukRestOauth2(mo *MukRestOauth2) error {
	return c.UpdateMukRestOauth2s([]int64{mo.Id.Get()}, mo)
}

// UpdateMukRestOauth2s updates existing muk_rest.oauth2 records.
// All records (represented by ids) will be updated by mo values.
func (c *Client) UpdateMukRestOauth2s(ids []int64, mo *MukRestOauth2) error {
	return c.Update(MukRestOauth2Model, ids, mo)
}

// DeleteMukRestOauth2 deletes an existing muk_rest.oauth2 record.
func (c *Client) DeleteMukRestOauth2(id int64) error {
	return c.DeleteMukRestOauth2s([]int64{id})
}

// DeleteMukRestOauth2s deletes existing muk_rest.oauth2 records.
func (c *Client) DeleteMukRestOauth2s(ids []int64) error {
	return c.Delete(MukRestOauth2Model, ids)
}

// GetMukRestOauth2 gets muk_rest.oauth2 existing record.
func (c *Client) GetMukRestOauth2(id int64) (*MukRestOauth2, error) {
	mos, err := c.GetMukRestOauth2s([]int64{id})
	if err != nil {
		return nil, err
	}
	if mos != nil && len(*mos) > 0 {
		return &((*mos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.oauth2 not found", id)
}

// GetMukRestOauth2s gets muk_rest.oauth2 existing records.
func (c *Client) GetMukRestOauth2s(ids []int64) (*MukRestOauth2s, error) {
	mos := &MukRestOauth2s{}
	if err := c.Read(MukRestOauth2Model, ids, nil, mos); err != nil {
		return nil, err
	}
	return mos, nil
}

// FindMukRestOauth2 finds muk_rest.oauth2 record by querying it with criteria.
func (c *Client) FindMukRestOauth2(criteria *Criteria) (*MukRestOauth2, error) {
	mos := &MukRestOauth2s{}
	if err := c.SearchRead(MukRestOauth2Model, criteria, NewOptions().Limit(1), mos); err != nil {
		return nil, err
	}
	if mos != nil && len(*mos) > 0 {
		return &((*mos)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.oauth2 was not found with criteria %v", criteria)
}

// FindMukRestOauth2s finds muk_rest.oauth2 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestOauth2s(criteria *Criteria, options *Options) (*MukRestOauth2s, error) {
	mos := &MukRestOauth2s{}
	if err := c.SearchRead(MukRestOauth2Model, criteria, options, mos); err != nil {
		return nil, err
	}
	return mos, nil
}

// FindMukRestOauth2Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestOauth2Ids(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestOauth2Model, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestOauth2Id finds record id by querying it with criteria.
func (c *Client) FindMukRestOauth2Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestOauth2Model, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.oauth2 was not found with criteria %v and options %v", criteria, options)
}
