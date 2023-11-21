package odoo

import (
	"fmt"
)

// MukRestOauth1 represents muk_rest.oauth1 model.
type MukRestOauth1 struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Active         *Bool      `xmlrpc:"active,omptempty"`
	CallbackIds    *Relation  `xmlrpc:"callback_ids,omptempty"`
	Color          *Int       `xmlrpc:"color,omptempty"`
	Company        *String    `xmlrpc:"company,omptempty"`
	ConsumerKey    *String    `xmlrpc:"consumer_key,omptempty"`
	ConsumerSecret *String    `xmlrpc:"consumer_secret,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description    *String    `xmlrpc:"description,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Homepage       *String    `xmlrpc:"homepage,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	LogoUrl        *String    `xmlrpc:"logo_url,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	Oauth1Ids      *Relation  `xmlrpc:"oauth1_ids,omptempty"`
	Oauth2Ids      *Relation  `xmlrpc:"oauth2_ids,omptempty"`
	OauthId        *Many2One  `xmlrpc:"oauth_id,omptempty"`
	PrivacyPolicy  *String    `xmlrpc:"privacy_policy,omptempty"`
	RuleIds        *Relation  `xmlrpc:"rule_ids,omptempty"`
	Security       *Selection `xmlrpc:"security,omptempty"`
	ServiceTerms   *String    `xmlrpc:"service_terms,omptempty"`
	Sessions       *Int       `xmlrpc:"sessions,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MukRestOauth1s represents array of muk_rest.oauth1 model.
type MukRestOauth1s []MukRestOauth1

// MukRestOauth1Model is the odoo model name.
const MukRestOauth1Model = "muk_rest.oauth1"

// Many2One convert MukRestOauth1 to *Many2One.
func (mo *MukRestOauth1) Many2One() *Many2One {
	return NewMany2One(mo.Id.Get(), "")
}

// CreateMukRestOauth1 creates a new muk_rest.oauth1 model and returns its id.
func (c *Client) CreateMukRestOauth1(mo *MukRestOauth1) (int64, error) {
	ids, err := c.CreateMukRestOauth1s([]*MukRestOauth1{mo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestOauth1 creates a new muk_rest.oauth1 model and returns its id.
func (c *Client) CreateMukRestOauth1s(mos []*MukRestOauth1) ([]int64, error) {
	var vv []interface{}
	for _, v := range mos {
		vv = append(vv, v)
	}
	return c.Create(MukRestOauth1Model, vv)
}

// UpdateMukRestOauth1 updates an existing muk_rest.oauth1 record.
func (c *Client) UpdateMukRestOauth1(mo *MukRestOauth1) error {
	return c.UpdateMukRestOauth1s([]int64{mo.Id.Get()}, mo)
}

// UpdateMukRestOauth1s updates existing muk_rest.oauth1 records.
// All records (represented by ids) will be updated by mo values.
func (c *Client) UpdateMukRestOauth1s(ids []int64, mo *MukRestOauth1) error {
	return c.Update(MukRestOauth1Model, ids, mo)
}

// DeleteMukRestOauth1 deletes an existing muk_rest.oauth1 record.
func (c *Client) DeleteMukRestOauth1(id int64) error {
	return c.DeleteMukRestOauth1s([]int64{id})
}

// DeleteMukRestOauth1s deletes existing muk_rest.oauth1 records.
func (c *Client) DeleteMukRestOauth1s(ids []int64) error {
	return c.Delete(MukRestOauth1Model, ids)
}

// GetMukRestOauth1 gets muk_rest.oauth1 existing record.
func (c *Client) GetMukRestOauth1(id int64) (*MukRestOauth1, error) {
	mos, err := c.GetMukRestOauth1s([]int64{id})
	if err != nil {
		return nil, err
	}
	if mos != nil && len(*mos) > 0 {
		return &((*mos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.oauth1 not found", id)
}

// GetMukRestOauth1s gets muk_rest.oauth1 existing records.
func (c *Client) GetMukRestOauth1s(ids []int64) (*MukRestOauth1s, error) {
	mos := &MukRestOauth1s{}
	if err := c.Read(MukRestOauth1Model, ids, nil, mos); err != nil {
		return nil, err
	}
	return mos, nil
}

// FindMukRestOauth1 finds muk_rest.oauth1 record by querying it with criteria.
func (c *Client) FindMukRestOauth1(criteria *Criteria) (*MukRestOauth1, error) {
	mos := &MukRestOauth1s{}
	if err := c.SearchRead(MukRestOauth1Model, criteria, NewOptions().Limit(1), mos); err != nil {
		return nil, err
	}
	if mos != nil && len(*mos) > 0 {
		return &((*mos)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.oauth1 was not found with criteria %v", criteria)
}

// FindMukRestOauth1s finds muk_rest.oauth1 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestOauth1s(criteria *Criteria, options *Options) (*MukRestOauth1s, error) {
	mos := &MukRestOauth1s{}
	if err := c.SearchRead(MukRestOauth1Model, criteria, options, mos); err != nil {
		return nil, err
	}
	return mos, nil
}

// FindMukRestOauth1Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestOauth1Ids(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestOauth1Model, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestOauth1Id finds record id by querying it with criteria.
func (c *Client) FindMukRestOauth1Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestOauth1Model, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.oauth1 was not found with criteria %v and options %v", criteria, options)
}
