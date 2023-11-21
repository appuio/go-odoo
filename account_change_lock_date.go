package odoo

import (
	"fmt"
)

// AccountChangeLockDate represents account.change.lock.date model.
type AccountChangeLockDate struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	FiscalyearLockDate *Time     `xmlrpc:"fiscalyear_lock_date,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	PeriodLockDate     *Time     `xmlrpc:"period_lock_date,omptempty"`
	TaxLockDate        *Time     `xmlrpc:"tax_lock_date,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountChangeLockDates represents array of account.change.lock.date model.
type AccountChangeLockDates []AccountChangeLockDate

// AccountChangeLockDateModel is the odoo model name.
const AccountChangeLockDateModel = "account.change.lock.date"

// Many2One convert AccountChangeLockDate to *Many2One.
func (acld *AccountChangeLockDate) Many2One() *Many2One {
	return NewMany2One(acld.Id.Get(), "")
}

// CreateAccountChangeLockDate creates a new account.change.lock.date model and returns its id.
func (c *Client) CreateAccountChangeLockDate(acld *AccountChangeLockDate) (int64, error) {
	ids, err := c.CreateAccountChangeLockDates([]*AccountChangeLockDate{acld})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountChangeLockDate creates a new account.change.lock.date model and returns its id.
func (c *Client) CreateAccountChangeLockDates(aclds []*AccountChangeLockDate) ([]int64, error) {
	var vv []interface{}
	for _, v := range aclds {
		vv = append(vv, v)
	}
	return c.Create(AccountChangeLockDateModel, vv)
}

// UpdateAccountChangeLockDate updates an existing account.change.lock.date record.
func (c *Client) UpdateAccountChangeLockDate(acld *AccountChangeLockDate) error {
	return c.UpdateAccountChangeLockDates([]int64{acld.Id.Get()}, acld)
}

// UpdateAccountChangeLockDates updates existing account.change.lock.date records.
// All records (represented by ids) will be updated by acld values.
func (c *Client) UpdateAccountChangeLockDates(ids []int64, acld *AccountChangeLockDate) error {
	return c.Update(AccountChangeLockDateModel, ids, acld)
}

// DeleteAccountChangeLockDate deletes an existing account.change.lock.date record.
func (c *Client) DeleteAccountChangeLockDate(id int64) error {
	return c.DeleteAccountChangeLockDates([]int64{id})
}

// DeleteAccountChangeLockDates deletes existing account.change.lock.date records.
func (c *Client) DeleteAccountChangeLockDates(ids []int64) error {
	return c.Delete(AccountChangeLockDateModel, ids)
}

// GetAccountChangeLockDate gets account.change.lock.date existing record.
func (c *Client) GetAccountChangeLockDate(id int64) (*AccountChangeLockDate, error) {
	aclds, err := c.GetAccountChangeLockDates([]int64{id})
	if err != nil {
		return nil, err
	}
	if aclds != nil && len(*aclds) > 0 {
		return &((*aclds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.change.lock.date not found", id)
}

// GetAccountChangeLockDates gets account.change.lock.date existing records.
func (c *Client) GetAccountChangeLockDates(ids []int64) (*AccountChangeLockDates, error) {
	aclds := &AccountChangeLockDates{}
	if err := c.Read(AccountChangeLockDateModel, ids, nil, aclds); err != nil {
		return nil, err
	}
	return aclds, nil
}

// FindAccountChangeLockDate finds account.change.lock.date record by querying it with criteria.
func (c *Client) FindAccountChangeLockDate(criteria *Criteria) (*AccountChangeLockDate, error) {
	aclds := &AccountChangeLockDates{}
	if err := c.SearchRead(AccountChangeLockDateModel, criteria, NewOptions().Limit(1), aclds); err != nil {
		return nil, err
	}
	if aclds != nil && len(*aclds) > 0 {
		return &((*aclds)[0]), nil
	}
	return nil, fmt.Errorf("account.change.lock.date was not found with criteria %v", criteria)
}

// FindAccountChangeLockDates finds account.change.lock.date records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChangeLockDates(criteria *Criteria, options *Options) (*AccountChangeLockDates, error) {
	aclds := &AccountChangeLockDates{}
	if err := c.SearchRead(AccountChangeLockDateModel, criteria, options, aclds); err != nil {
		return nil, err
	}
	return aclds, nil
}

// FindAccountChangeLockDateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountChangeLockDateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountChangeLockDateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountChangeLockDateId finds record id by querying it with criteria.
func (c *Client) FindAccountChangeLockDateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountChangeLockDateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.change.lock.date was not found with criteria %v and options %v", criteria, options)
}
