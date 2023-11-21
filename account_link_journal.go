package odoo

import (
	"fmt"
)

// AccountLinkJournal represents account.link.journal model.
type AccountLinkJournal struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	AccountIds   *Relation `xmlrpc:"account_ids,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	NumberAdded  *Int      `xmlrpc:"number_added,omptempty"`
	SyncDate     *Time     `xmlrpc:"sync_date,omptempty"`
	Transactions *String   `xmlrpc:"transactions,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountLinkJournals represents array of account.link.journal model.
type AccountLinkJournals []AccountLinkJournal

// AccountLinkJournalModel is the odoo model name.
const AccountLinkJournalModel = "account.link.journal"

// Many2One convert AccountLinkJournal to *Many2One.
func (alj *AccountLinkJournal) Many2One() *Many2One {
	return NewMany2One(alj.Id.Get(), "")
}

// CreateAccountLinkJournal creates a new account.link.journal model and returns its id.
func (c *Client) CreateAccountLinkJournal(alj *AccountLinkJournal) (int64, error) {
	ids, err := c.CreateAccountLinkJournals([]*AccountLinkJournal{alj})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountLinkJournal creates a new account.link.journal model and returns its id.
func (c *Client) CreateAccountLinkJournals(aljs []*AccountLinkJournal) ([]int64, error) {
	var vv []interface{}
	for _, v := range aljs {
		vv = append(vv, v)
	}
	return c.Create(AccountLinkJournalModel, vv)
}

// UpdateAccountLinkJournal updates an existing account.link.journal record.
func (c *Client) UpdateAccountLinkJournal(alj *AccountLinkJournal) error {
	return c.UpdateAccountLinkJournals([]int64{alj.Id.Get()}, alj)
}

// UpdateAccountLinkJournals updates existing account.link.journal records.
// All records (represented by ids) will be updated by alj values.
func (c *Client) UpdateAccountLinkJournals(ids []int64, alj *AccountLinkJournal) error {
	return c.Update(AccountLinkJournalModel, ids, alj)
}

// DeleteAccountLinkJournal deletes an existing account.link.journal record.
func (c *Client) DeleteAccountLinkJournal(id int64) error {
	return c.DeleteAccountLinkJournals([]int64{id})
}

// DeleteAccountLinkJournals deletes existing account.link.journal records.
func (c *Client) DeleteAccountLinkJournals(ids []int64) error {
	return c.Delete(AccountLinkJournalModel, ids)
}

// GetAccountLinkJournal gets account.link.journal existing record.
func (c *Client) GetAccountLinkJournal(id int64) (*AccountLinkJournal, error) {
	aljs, err := c.GetAccountLinkJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	if aljs != nil && len(*aljs) > 0 {
		return &((*aljs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.link.journal not found", id)
}

// GetAccountLinkJournals gets account.link.journal existing records.
func (c *Client) GetAccountLinkJournals(ids []int64) (*AccountLinkJournals, error) {
	aljs := &AccountLinkJournals{}
	if err := c.Read(AccountLinkJournalModel, ids, nil, aljs); err != nil {
		return nil, err
	}
	return aljs, nil
}

// FindAccountLinkJournal finds account.link.journal record by querying it with criteria.
func (c *Client) FindAccountLinkJournal(criteria *Criteria) (*AccountLinkJournal, error) {
	aljs := &AccountLinkJournals{}
	if err := c.SearchRead(AccountLinkJournalModel, criteria, NewOptions().Limit(1), aljs); err != nil {
		return nil, err
	}
	if aljs != nil && len(*aljs) > 0 {
		return &((*aljs)[0]), nil
	}
	return nil, fmt.Errorf("account.link.journal was not found with criteria %v", criteria)
}

// FindAccountLinkJournals finds account.link.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLinkJournals(criteria *Criteria, options *Options) (*AccountLinkJournals, error) {
	aljs := &AccountLinkJournals{}
	if err := c.SearchRead(AccountLinkJournalModel, criteria, options, aljs); err != nil {
		return nil, err
	}
	return aljs, nil
}

// FindAccountLinkJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLinkJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountLinkJournalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountLinkJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountLinkJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountLinkJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.link.journal was not found with criteria %v and options %v", criteria, options)
}
