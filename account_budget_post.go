package odoo

import (
	"fmt"
)

// AccountBudgetPost represents account.budget.post model.
type AccountBudgetPost struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	AccountIds  *Relation `xmlrpc:"account_ids,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBudgetPosts represents array of account.budget.post model.
type AccountBudgetPosts []AccountBudgetPost

// AccountBudgetPostModel is the odoo model name.
const AccountBudgetPostModel = "account.budget.post"

// Many2One convert AccountBudgetPost to *Many2One.
func (abp *AccountBudgetPost) Many2One() *Many2One {
	return NewMany2One(abp.Id.Get(), "")
}

// CreateAccountBudgetPost creates a new account.budget.post model and returns its id.
func (c *Client) CreateAccountBudgetPost(abp *AccountBudgetPost) (int64, error) {
	ids, err := c.CreateAccountBudgetPosts([]*AccountBudgetPost{abp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBudgetPost creates a new account.budget.post model and returns its id.
func (c *Client) CreateAccountBudgetPosts(abps []*AccountBudgetPost) ([]int64, error) {
	var vv []interface{}
	for _, v := range abps {
		vv = append(vv, v)
	}
	return c.Create(AccountBudgetPostModel, vv)
}

// UpdateAccountBudgetPost updates an existing account.budget.post record.
func (c *Client) UpdateAccountBudgetPost(abp *AccountBudgetPost) error {
	return c.UpdateAccountBudgetPosts([]int64{abp.Id.Get()}, abp)
}

// UpdateAccountBudgetPosts updates existing account.budget.post records.
// All records (represented by ids) will be updated by abp values.
func (c *Client) UpdateAccountBudgetPosts(ids []int64, abp *AccountBudgetPost) error {
	return c.Update(AccountBudgetPostModel, ids, abp)
}

// DeleteAccountBudgetPost deletes an existing account.budget.post record.
func (c *Client) DeleteAccountBudgetPost(id int64) error {
	return c.DeleteAccountBudgetPosts([]int64{id})
}

// DeleteAccountBudgetPosts deletes existing account.budget.post records.
func (c *Client) DeleteAccountBudgetPosts(ids []int64) error {
	return c.Delete(AccountBudgetPostModel, ids)
}

// GetAccountBudgetPost gets account.budget.post existing record.
func (c *Client) GetAccountBudgetPost(id int64) (*AccountBudgetPost, error) {
	abps, err := c.GetAccountBudgetPosts([]int64{id})
	if err != nil {
		return nil, err
	}
	if abps != nil && len(*abps) > 0 {
		return &((*abps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.budget.post not found", id)
}

// GetAccountBudgetPosts gets account.budget.post existing records.
func (c *Client) GetAccountBudgetPosts(ids []int64) (*AccountBudgetPosts, error) {
	abps := &AccountBudgetPosts{}
	if err := c.Read(AccountBudgetPostModel, ids, nil, abps); err != nil {
		return nil, err
	}
	return abps, nil
}

// FindAccountBudgetPost finds account.budget.post record by querying it with criteria.
func (c *Client) FindAccountBudgetPost(criteria *Criteria) (*AccountBudgetPost, error) {
	abps := &AccountBudgetPosts{}
	if err := c.SearchRead(AccountBudgetPostModel, criteria, NewOptions().Limit(1), abps); err != nil {
		return nil, err
	}
	if abps != nil && len(*abps) > 0 {
		return &((*abps)[0]), nil
	}
	return nil, fmt.Errorf("account.budget.post was not found with criteria %v", criteria)
}

// FindAccountBudgetPosts finds account.budget.post records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBudgetPosts(criteria *Criteria, options *Options) (*AccountBudgetPosts, error) {
	abps := &AccountBudgetPosts{}
	if err := c.SearchRead(AccountBudgetPostModel, criteria, options, abps); err != nil {
		return nil, err
	}
	return abps, nil
}

// FindAccountBudgetPostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBudgetPostIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBudgetPostModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBudgetPostId finds record id by querying it with criteria.
func (c *Client) FindAccountBudgetPostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBudgetPostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.budget.post was not found with criteria %v and options %v", criteria, options)
}
