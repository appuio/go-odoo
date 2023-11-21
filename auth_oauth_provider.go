package odoo

import (
	"fmt"
)

// AuthOauthProvider represents auth.oauth.provider model.
type AuthOauthProvider struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	AuthEndpoint       *String    `xmlrpc:"auth_endpoint,omptempty"`
	Body               *String    `xmlrpc:"body,omptempty"`
	ClientId           *String    `xmlrpc:"client_id,omptempty"`
	ClientSecret       *String    `xmlrpc:"client_secret,omptempty"`
	CodeVerifier       *String    `xmlrpc:"code_verifier,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	CssClass           *String    `xmlrpc:"css_class,omptempty"`
	DataEndpoint       *String    `xmlrpc:"data_endpoint,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Enabled            *Bool      `xmlrpc:"enabled,omptempty"`
	Flow               *Selection `xmlrpc:"flow,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	JwksUri            *String    `xmlrpc:"jwks_uri,omptempty"`
	Name               *String    `xmlrpc:"name,omptempty"`
	Scope              *String    `xmlrpc:"scope,omptempty"`
	Sequence           *Int       `xmlrpc:"sequence,omptempty"`
	TokenEndpoint      *String    `xmlrpc:"token_endpoint,omptempty"`
	TokenMap           *String    `xmlrpc:"token_map,omptempty"`
	ValidationEndpoint *String    `xmlrpc:"validation_endpoint,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AuthOauthProviders represents array of auth.oauth.provider model.
type AuthOauthProviders []AuthOauthProvider

// AuthOauthProviderModel is the odoo model name.
const AuthOauthProviderModel = "auth.oauth.provider"

// Many2One convert AuthOauthProvider to *Many2One.
func (aop *AuthOauthProvider) Many2One() *Many2One {
	return NewMany2One(aop.Id.Get(), "")
}

// CreateAuthOauthProvider creates a new auth.oauth.provider model and returns its id.
func (c *Client) CreateAuthOauthProvider(aop *AuthOauthProvider) (int64, error) {
	ids, err := c.CreateAuthOauthProviders([]*AuthOauthProvider{aop})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuthOauthProvider creates a new auth.oauth.provider model and returns its id.
func (c *Client) CreateAuthOauthProviders(aops []*AuthOauthProvider) ([]int64, error) {
	var vv []interface{}
	for _, v := range aops {
		vv = append(vv, v)
	}
	return c.Create(AuthOauthProviderModel, vv)
}

// UpdateAuthOauthProvider updates an existing auth.oauth.provider record.
func (c *Client) UpdateAuthOauthProvider(aop *AuthOauthProvider) error {
	return c.UpdateAuthOauthProviders([]int64{aop.Id.Get()}, aop)
}

// UpdateAuthOauthProviders updates existing auth.oauth.provider records.
// All records (represented by ids) will be updated by aop values.
func (c *Client) UpdateAuthOauthProviders(ids []int64, aop *AuthOauthProvider) error {
	return c.Update(AuthOauthProviderModel, ids, aop)
}

// DeleteAuthOauthProvider deletes an existing auth.oauth.provider record.
func (c *Client) DeleteAuthOauthProvider(id int64) error {
	return c.DeleteAuthOauthProviders([]int64{id})
}

// DeleteAuthOauthProviders deletes existing auth.oauth.provider records.
func (c *Client) DeleteAuthOauthProviders(ids []int64) error {
	return c.Delete(AuthOauthProviderModel, ids)
}

// GetAuthOauthProvider gets auth.oauth.provider existing record.
func (c *Client) GetAuthOauthProvider(id int64) (*AuthOauthProvider, error) {
	aops, err := c.GetAuthOauthProviders([]int64{id})
	if err != nil {
		return nil, err
	}
	if aops != nil && len(*aops) > 0 {
		return &((*aops)[0]), nil
	}
	return nil, fmt.Errorf("id %v of auth.oauth.provider not found", id)
}

// GetAuthOauthProviders gets auth.oauth.provider existing records.
func (c *Client) GetAuthOauthProviders(ids []int64) (*AuthOauthProviders, error) {
	aops := &AuthOauthProviders{}
	if err := c.Read(AuthOauthProviderModel, ids, nil, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAuthOauthProvider finds auth.oauth.provider record by querying it with criteria.
func (c *Client) FindAuthOauthProvider(criteria *Criteria) (*AuthOauthProvider, error) {
	aops := &AuthOauthProviders{}
	if err := c.SearchRead(AuthOauthProviderModel, criteria, NewOptions().Limit(1), aops); err != nil {
		return nil, err
	}
	if aops != nil && len(*aops) > 0 {
		return &((*aops)[0]), nil
	}
	return nil, fmt.Errorf("auth.oauth.provider was not found with criteria %v", criteria)
}

// FindAuthOauthProviders finds auth.oauth.provider records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthOauthProviders(criteria *Criteria, options *Options) (*AuthOauthProviders, error) {
	aops := &AuthOauthProviders{}
	if err := c.SearchRead(AuthOauthProviderModel, criteria, options, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAuthOauthProviderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthOauthProviderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AuthOauthProviderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAuthOauthProviderId finds record id by querying it with criteria.
func (c *Client) FindAuthOauthProviderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthOauthProviderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("auth.oauth.provider was not found with criteria %v and options %v", criteria, options)
}
