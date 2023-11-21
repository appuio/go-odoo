package odoo

import (
	"fmt"
)

// AccountLinkJournalLine represents account.link.journal.line model.
type AccountLinkJournalLine struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	AccountNumber         *String   `xmlrpc:"account_number,omptempty"`
	AccountOnlineWizardId *Many2One `xmlrpc:"account_online_wizard_id,omptempty"`
	Balance               *Float    `xmlrpc:"balance,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId            *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	JournalId             *Many2One `xmlrpc:"journal_id,omptempty"`
	Name                  *String   `xmlrpc:"name,omptempty"`
	OnlineAccountId       *Many2One `xmlrpc:"online_account_id,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountLinkJournalLines represents array of account.link.journal.line model.
type AccountLinkJournalLines []AccountLinkJournalLine

// AccountLinkJournalLineModel is the odoo model name.
const AccountLinkJournalLineModel = "account.link.journal.line"

// Many2One convert AccountLinkJournalLine to *Many2One.
func (aljl *AccountLinkJournalLine) Many2One() *Many2One {
	return NewMany2One(aljl.Id.Get(), "")
}

// CreateAccountLinkJournalLine creates a new account.link.journal.line model and returns its id.
func (c *Client) CreateAccountLinkJournalLine(aljl *AccountLinkJournalLine) (int64, error) {
	ids, err := c.CreateAccountLinkJournalLines([]*AccountLinkJournalLine{aljl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountLinkJournalLine creates a new account.link.journal.line model and returns its id.
func (c *Client) CreateAccountLinkJournalLines(aljls []*AccountLinkJournalLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range aljls {
		vv = append(vv, v)
	}
	return c.Create(AccountLinkJournalLineModel, vv)
}

// UpdateAccountLinkJournalLine updates an existing account.link.journal.line record.
func (c *Client) UpdateAccountLinkJournalLine(aljl *AccountLinkJournalLine) error {
	return c.UpdateAccountLinkJournalLines([]int64{aljl.Id.Get()}, aljl)
}

// UpdateAccountLinkJournalLines updates existing account.link.journal.line records.
// All records (represented by ids) will be updated by aljl values.
func (c *Client) UpdateAccountLinkJournalLines(ids []int64, aljl *AccountLinkJournalLine) error {
	return c.Update(AccountLinkJournalLineModel, ids, aljl)
}

// DeleteAccountLinkJournalLine deletes an existing account.link.journal.line record.
func (c *Client) DeleteAccountLinkJournalLine(id int64) error {
	return c.DeleteAccountLinkJournalLines([]int64{id})
}

// DeleteAccountLinkJournalLines deletes existing account.link.journal.line records.
func (c *Client) DeleteAccountLinkJournalLines(ids []int64) error {
	return c.Delete(AccountLinkJournalLineModel, ids)
}

// GetAccountLinkJournalLine gets account.link.journal.line existing record.
func (c *Client) GetAccountLinkJournalLine(id int64) (*AccountLinkJournalLine, error) {
	aljls, err := c.GetAccountLinkJournalLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if aljls != nil && len(*aljls) > 0 {
		return &((*aljls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.link.journal.line not found", id)
}

// GetAccountLinkJournalLines gets account.link.journal.line existing records.
func (c *Client) GetAccountLinkJournalLines(ids []int64) (*AccountLinkJournalLines, error) {
	aljls := &AccountLinkJournalLines{}
	if err := c.Read(AccountLinkJournalLineModel, ids, nil, aljls); err != nil {
		return nil, err
	}
	return aljls, nil
}

// FindAccountLinkJournalLine finds account.link.journal.line record by querying it with criteria.
func (c *Client) FindAccountLinkJournalLine(criteria *Criteria) (*AccountLinkJournalLine, error) {
	aljls := &AccountLinkJournalLines{}
	if err := c.SearchRead(AccountLinkJournalLineModel, criteria, NewOptions().Limit(1), aljls); err != nil {
		return nil, err
	}
	if aljls != nil && len(*aljls) > 0 {
		return &((*aljls)[0]), nil
	}
	return nil, fmt.Errorf("account.link.journal.line was not found with criteria %v", criteria)
}

// FindAccountLinkJournalLines finds account.link.journal.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLinkJournalLines(criteria *Criteria, options *Options) (*AccountLinkJournalLines, error) {
	aljls := &AccountLinkJournalLines{}
	if err := c.SearchRead(AccountLinkJournalLineModel, criteria, options, aljls); err != nil {
		return nil, err
	}
	return aljls, nil
}

// FindAccountLinkJournalLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountLinkJournalLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountLinkJournalLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountLinkJournalLineId finds record id by querying it with criteria.
func (c *Client) FindAccountLinkJournalLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountLinkJournalLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.link.journal.line was not found with criteria %v and options %v", criteria, options)
}
