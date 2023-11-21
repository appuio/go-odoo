package odoo

import (
	"fmt"
)

// AccountOnlineAccount represents account.online.account model.
type AccountOnlineAccount struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	AccountData            *String   `xmlrpc:"account_data,omptempty"`
	AccountNumber          *String   `xmlrpc:"account_number,omptempty"`
	AccountOnlineLinkId    *Many2One `xmlrpc:"account_online_link_id,omptempty"`
	Balance                *Float    `xmlrpc:"balance,omptempty"`
	CompanyId              *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId             *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	InverseBalanceSign     *Bool     `xmlrpc:"inverse_balance_sign,omptempty"`
	InverseTransactionSign *Bool     `xmlrpc:"inverse_transaction_sign,omptempty"`
	JournalIds             *Relation `xmlrpc:"journal_ids,omptempty"`
	LastSync               *Time     `xmlrpc:"last_sync,omptempty"`
	Name                   *String   `xmlrpc:"name,omptempty"`
	OnlineIdentifier       *String   `xmlrpc:"online_identifier,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountOnlineAccounts represents array of account.online.account model.
type AccountOnlineAccounts []AccountOnlineAccount

// AccountOnlineAccountModel is the odoo model name.
const AccountOnlineAccountModel = "account.online.account"

// Many2One convert AccountOnlineAccount to *Many2One.
func (aoa *AccountOnlineAccount) Many2One() *Many2One {
	return NewMany2One(aoa.Id.Get(), "")
}

// CreateAccountOnlineAccount creates a new account.online.account model and returns its id.
func (c *Client) CreateAccountOnlineAccount(aoa *AccountOnlineAccount) (int64, error) {
	ids, err := c.CreateAccountOnlineAccounts([]*AccountOnlineAccount{aoa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOnlineAccount creates a new account.online.account model and returns its id.
func (c *Client) CreateAccountOnlineAccounts(aoas []*AccountOnlineAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range aoas {
		vv = append(vv, v)
	}
	return c.Create(AccountOnlineAccountModel, vv)
}

// UpdateAccountOnlineAccount updates an existing account.online.account record.
func (c *Client) UpdateAccountOnlineAccount(aoa *AccountOnlineAccount) error {
	return c.UpdateAccountOnlineAccounts([]int64{aoa.Id.Get()}, aoa)
}

// UpdateAccountOnlineAccounts updates existing account.online.account records.
// All records (represented by ids) will be updated by aoa values.
func (c *Client) UpdateAccountOnlineAccounts(ids []int64, aoa *AccountOnlineAccount) error {
	return c.Update(AccountOnlineAccountModel, ids, aoa)
}

// DeleteAccountOnlineAccount deletes an existing account.online.account record.
func (c *Client) DeleteAccountOnlineAccount(id int64) error {
	return c.DeleteAccountOnlineAccounts([]int64{id})
}

// DeleteAccountOnlineAccounts deletes existing account.online.account records.
func (c *Client) DeleteAccountOnlineAccounts(ids []int64) error {
	return c.Delete(AccountOnlineAccountModel, ids)
}

// GetAccountOnlineAccount gets account.online.account existing record.
func (c *Client) GetAccountOnlineAccount(id int64) (*AccountOnlineAccount, error) {
	aoas, err := c.GetAccountOnlineAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if aoas != nil && len(*aoas) > 0 {
		return &((*aoas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.online.account not found", id)
}

// GetAccountOnlineAccounts gets account.online.account existing records.
func (c *Client) GetAccountOnlineAccounts(ids []int64) (*AccountOnlineAccounts, error) {
	aoas := &AccountOnlineAccounts{}
	if err := c.Read(AccountOnlineAccountModel, ids, nil, aoas); err != nil {
		return nil, err
	}
	return aoas, nil
}

// FindAccountOnlineAccount finds account.online.account record by querying it with criteria.
func (c *Client) FindAccountOnlineAccount(criteria *Criteria) (*AccountOnlineAccount, error) {
	aoas := &AccountOnlineAccounts{}
	if err := c.SearchRead(AccountOnlineAccountModel, criteria, NewOptions().Limit(1), aoas); err != nil {
		return nil, err
	}
	if aoas != nil && len(*aoas) > 0 {
		return &((*aoas)[0]), nil
	}
	return nil, fmt.Errorf("account.online.account was not found with criteria %v", criteria)
}

// FindAccountOnlineAccounts finds account.online.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineAccounts(criteria *Criteria, options *Options) (*AccountOnlineAccounts, error) {
	aoas := &AccountOnlineAccounts{}
	if err := c.SearchRead(AccountOnlineAccountModel, criteria, options, aoas); err != nil {
		return nil, err
	}
	return aoas, nil
}

// FindAccountOnlineAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountOnlineAccountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountOnlineAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.online.account was not found with criteria %v and options %v", criteria, options)
}
