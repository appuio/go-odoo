package odoo

import (
	"fmt"
)

// MukRestOauth represents muk_rest.oauth model.
type MukRestOauth struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	Active        *Bool      `xmlrpc:"active,omptempty"`
	CallbackIds   *Relation  `xmlrpc:"callback_ids,omptempty"`
	Color         *Int       `xmlrpc:"color,omptempty"`
	Company       *String    `xmlrpc:"company,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description   *String    `xmlrpc:"description,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Homepage      *String    `xmlrpc:"homepage,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	LogoUrl       *String    `xmlrpc:"logo_url,omptempty"`
	Name          *String    `xmlrpc:"name,omptempty"`
	Oauth1Ids     *Relation  `xmlrpc:"oauth1_ids,omptempty"`
	Oauth2Ids     *Relation  `xmlrpc:"oauth2_ids,omptempty"`
	PrivacyPolicy *String    `xmlrpc:"privacy_policy,omptempty"`
	RuleIds       *Relation  `xmlrpc:"rule_ids,omptempty"`
	Security      *Selection `xmlrpc:"security,omptempty"`
	ServiceTerms  *String    `xmlrpc:"service_terms,omptempty"`
	Sessions      *Int       `xmlrpc:"sessions,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MukRestOauths represents array of muk_rest.oauth model.
type MukRestOauths []MukRestOauth

// MukRestOauthModel is the odoo model name.
const MukRestOauthModel = "muk_rest.oauth"

// Many2One convert MukRestOauth to *Many2One.
func (mo *MukRestOauth) Many2One() *Many2One {
	return NewMany2One(mo.Id.Get(), "")
}

// CreateMukRestOauth creates a new muk_rest.oauth model and returns its id.
func (c *Client) CreateMukRestOauth(mo *MukRestOauth) (int64, error) {
	ids, err := c.CreateMukRestOauths([]*MukRestOauth{mo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestOauth creates a new muk_rest.oauth model and returns its id.
func (c *Client) CreateMukRestOauths(mos []*MukRestOauth) ([]int64, error) {
	var vv []interface{}
	for _, v := range mos {
		vv = append(vv, v)
	}
	return c.Create(MukRestOauthModel, vv)
}

// UpdateMukRestOauth updates an existing muk_rest.oauth record.
func (c *Client) UpdateMukRestOauth(mo *MukRestOauth) error {
	return c.UpdateMukRestOauths([]int64{mo.Id.Get()}, mo)
}

// UpdateMukRestOauths updates existing muk_rest.oauth records.
// All records (represented by ids) will be updated by mo values.
func (c *Client) UpdateMukRestOauths(ids []int64, mo *MukRestOauth) error {
	return c.Update(MukRestOauthModel, ids, mo)
}

// DeleteMukRestOauth deletes an existing muk_rest.oauth record.
func (c *Client) DeleteMukRestOauth(id int64) error {
	return c.DeleteMukRestOauths([]int64{id})
}

// DeleteMukRestOauths deletes existing muk_rest.oauth records.
func (c *Client) DeleteMukRestOauths(ids []int64) error {
	return c.Delete(MukRestOauthModel, ids)
}

// GetMukRestOauth gets muk_rest.oauth existing record.
func (c *Client) GetMukRestOauth(id int64) (*MukRestOauth, error) {
	mos, err := c.GetMukRestOauths([]int64{id})
	if err != nil {
		return nil, err
	}
	if mos != nil && len(*mos) > 0 {
		return &((*mos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.oauth not found", id)
}

// GetMukRestOauths gets muk_rest.oauth existing records.
func (c *Client) GetMukRestOauths(ids []int64) (*MukRestOauths, error) {
	mos := &MukRestOauths{}
	if err := c.Read(MukRestOauthModel, ids, nil, mos); err != nil {
		return nil, err
	}
	return mos, nil
}

// FindMukRestOauth finds muk_rest.oauth record by querying it with criteria.
func (c *Client) FindMukRestOauth(criteria *Criteria) (*MukRestOauth, error) {
	mos := &MukRestOauths{}
	if err := c.SearchRead(MukRestOauthModel, criteria, NewOptions().Limit(1), mos); err != nil {
		return nil, err
	}
	if mos != nil && len(*mos) > 0 {
		return &((*mos)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.oauth was not found with criteria %v", criteria)
}

// FindMukRestOauths finds muk_rest.oauth records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestOauths(criteria *Criteria, options *Options) (*MukRestOauths, error) {
	mos := &MukRestOauths{}
	if err := c.SearchRead(MukRestOauthModel, criteria, options, mos); err != nil {
		return nil, err
	}
	return mos, nil
}

// FindMukRestOauthIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestOauthIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestOauthModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestOauthId finds record id by querying it with criteria.
func (c *Client) FindMukRestOauthId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestOauthModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.oauth was not found with criteria %v and options %v", criteria, options)
}
