package odoo

import (
	"fmt"
)

// AccountTransferModel represents account.transfer.model model.
type AccountTransferModel struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	AccountIds   *Relation  `xmlrpc:"account_ids,omptempty"`
	CompanyId    *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateStart    *Time      `xmlrpc:"date_start,omptempty"`
	DateStop     *Time      `xmlrpc:"date_stop,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	Frequency    *Selection `xmlrpc:"frequency,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	JournalId    *Many2One  `xmlrpc:"journal_id,omptempty"`
	LineIds      *Relation  `xmlrpc:"line_ids,omptempty"`
	MoveIds      *Relation  `xmlrpc:"move_ids,omptempty"`
	MoveIdsCount *Int       `xmlrpc:"move_ids_count,omptempty"`
	Name         *String    `xmlrpc:"name,omptempty"`
	State        *Selection `xmlrpc:"state,omptempty"`
	TotalPercent *Float     `xmlrpc:"total_percent,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountTransferModels represents array of account.transfer.model model.
type AccountTransferModels []AccountTransferModel

// AccountTransferModelModel is the odoo model name.
const AccountTransferModelModel = "account.transfer.model"

// Many2One convert AccountTransferModel to *Many2One.
func (atm *AccountTransferModel) Many2One() *Many2One {
	return NewMany2One(atm.Id.Get(), "")
}

// CreateAccountTransferModel creates a new account.transfer.model model and returns its id.
func (c *Client) CreateAccountTransferModel(atm *AccountTransferModel) (int64, error) {
	ids, err := c.CreateAccountTransferModels([]*AccountTransferModel{atm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTransferModel creates a new account.transfer.model model and returns its id.
func (c *Client) CreateAccountTransferModels(atms []*AccountTransferModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range atms {
		vv = append(vv, v)
	}
	return c.Create(AccountTransferModelModel, vv)
}

// UpdateAccountTransferModel updates an existing account.transfer.model record.
func (c *Client) UpdateAccountTransferModel(atm *AccountTransferModel) error {
	return c.UpdateAccountTransferModels([]int64{atm.Id.Get()}, atm)
}

// UpdateAccountTransferModels updates existing account.transfer.model records.
// All records (represented by ids) will be updated by atm values.
func (c *Client) UpdateAccountTransferModels(ids []int64, atm *AccountTransferModel) error {
	return c.Update(AccountTransferModelModel, ids, atm)
}

// DeleteAccountTransferModel deletes an existing account.transfer.model record.
func (c *Client) DeleteAccountTransferModel(id int64) error {
	return c.DeleteAccountTransferModels([]int64{id})
}

// DeleteAccountTransferModels deletes existing account.transfer.model records.
func (c *Client) DeleteAccountTransferModels(ids []int64) error {
	return c.Delete(AccountTransferModelModel, ids)
}

// GetAccountTransferModel gets account.transfer.model existing record.
func (c *Client) GetAccountTransferModel(id int64) (*AccountTransferModel, error) {
	atms, err := c.GetAccountTransferModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if atms != nil && len(*atms) > 0 {
		return &((*atms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.transfer.model not found", id)
}

// GetAccountTransferModels gets account.transfer.model existing records.
func (c *Client) GetAccountTransferModels(ids []int64) (*AccountTransferModels, error) {
	atms := &AccountTransferModels{}
	if err := c.Read(AccountTransferModelModel, ids, nil, atms); err != nil {
		return nil, err
	}
	return atms, nil
}

// FindAccountTransferModel finds account.transfer.model record by querying it with criteria.
func (c *Client) FindAccountTransferModel(criteria *Criteria) (*AccountTransferModel, error) {
	atms := &AccountTransferModels{}
	if err := c.SearchRead(AccountTransferModelModel, criteria, NewOptions().Limit(1), atms); err != nil {
		return nil, err
	}
	if atms != nil && len(*atms) > 0 {
		return &((*atms)[0]), nil
	}
	return nil, fmt.Errorf("account.transfer.model was not found with criteria %v", criteria)
}

// FindAccountTransferModels finds account.transfer.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTransferModels(criteria *Criteria, options *Options) (*AccountTransferModels, error) {
	atms := &AccountTransferModels{}
	if err := c.SearchRead(AccountTransferModelModel, criteria, options, atms); err != nil {
		return nil, err
	}
	return atms, nil
}

// FindAccountTransferModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTransferModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTransferModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTransferModelId finds record id by querying it with criteria.
func (c *Client) FindAccountTransferModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTransferModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.transfer.model was not found with criteria %v and options %v", criteria, options)
}
