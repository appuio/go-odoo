package odoo

import (
	"fmt"
)

// AuthAttempt represents auth.attempt model.
type AuthAttempt struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Blocked     *Bool     `xmlrpc:"blocked,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Delay       *Int      `xmlrpc:"delay,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Login       *String   `xmlrpc:"login,omptempty"`
	NAttempts   *Int      `xmlrpc:"n_attempts,omptempty"`
	Remote      *String   `xmlrpc:"remote,omptempty"`
	Time        *Time     `xmlrpc:"time,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AuthAttempts represents array of auth.attempt model.
type AuthAttempts []AuthAttempt

// AuthAttemptModel is the odoo model name.
const AuthAttemptModel = "auth.attempt"

// Many2One convert AuthAttempt to *Many2One.
func (aa *AuthAttempt) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAuthAttempt creates a new auth.attempt model and returns its id.
func (c *Client) CreateAuthAttempt(aa *AuthAttempt) (int64, error) {
	ids, err := c.CreateAuthAttempts([]*AuthAttempt{aa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuthAttempt creates a new auth.attempt model and returns its id.
func (c *Client) CreateAuthAttempts(aas []*AuthAttempt) ([]int64, error) {
	var vv []interface{}
	for _, v := range aas {
		vv = append(vv, v)
	}
	return c.Create(AuthAttemptModel, vv)
}

// UpdateAuthAttempt updates an existing auth.attempt record.
func (c *Client) UpdateAuthAttempt(aa *AuthAttempt) error {
	return c.UpdateAuthAttempts([]int64{aa.Id.Get()}, aa)
}

// UpdateAuthAttempts updates existing auth.attempt records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAuthAttempts(ids []int64, aa *AuthAttempt) error {
	return c.Update(AuthAttemptModel, ids, aa)
}

// DeleteAuthAttempt deletes an existing auth.attempt record.
func (c *Client) DeleteAuthAttempt(id int64) error {
	return c.DeleteAuthAttempts([]int64{id})
}

// DeleteAuthAttempts deletes existing auth.attempt records.
func (c *Client) DeleteAuthAttempts(ids []int64) error {
	return c.Delete(AuthAttemptModel, ids)
}

// GetAuthAttempt gets auth.attempt existing record.
func (c *Client) GetAuthAttempt(id int64) (*AuthAttempt, error) {
	aas, err := c.GetAuthAttempts([]int64{id})
	if err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of auth.attempt not found", id)
}

// GetAuthAttempts gets auth.attempt existing records.
func (c *Client) GetAuthAttempts(ids []int64) (*AuthAttempts, error) {
	aas := &AuthAttempts{}
	if err := c.Read(AuthAttemptModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAuthAttempt finds auth.attempt record by querying it with criteria.
func (c *Client) FindAuthAttempt(criteria *Criteria) (*AuthAttempt, error) {
	aas := &AuthAttempts{}
	if err := c.SearchRead(AuthAttemptModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("auth.attempt was not found with criteria %v", criteria)
}

// FindAuthAttempts finds auth.attempt records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthAttempts(criteria *Criteria, options *Options) (*AuthAttempts, error) {
	aas := &AuthAttempts{}
	if err := c.SearchRead(AuthAttemptModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAuthAttemptIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthAttemptIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AuthAttemptModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAuthAttemptId finds record id by querying it with criteria.
func (c *Client) FindAuthAttemptId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthAttemptModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("auth.attempt was not found with criteria %v and options %v", criteria, options)
}
