package odoo

import (
	"fmt"
)

// AccountBatchErrorWizard represents account.batch.error.wizard model.
type AccountBatchErrorWizard struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	BatchPaymentId    *Many2One `xmlrpc:"batch_payment_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	ErrorLineIds      *Relation `xmlrpc:"error_line_ids,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	ShowRemoveOptions *Bool     `xmlrpc:"show_remove_options,omptempty"`
	WarningLineIds    *Relation `xmlrpc:"warning_line_ids,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBatchErrorWizards represents array of account.batch.error.wizard model.
type AccountBatchErrorWizards []AccountBatchErrorWizard

// AccountBatchErrorWizardModel is the odoo model name.
const AccountBatchErrorWizardModel = "account.batch.error.wizard"

// Many2One convert AccountBatchErrorWizard to *Many2One.
func (abew *AccountBatchErrorWizard) Many2One() *Many2One {
	return NewMany2One(abew.Id.Get(), "")
}

// CreateAccountBatchErrorWizard creates a new account.batch.error.wizard model and returns its id.
func (c *Client) CreateAccountBatchErrorWizard(abew *AccountBatchErrorWizard) (int64, error) {
	ids, err := c.CreateAccountBatchErrorWizards([]*AccountBatchErrorWizard{abew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchErrorWizard creates a new account.batch.error.wizard model and returns its id.
func (c *Client) CreateAccountBatchErrorWizards(abews []*AccountBatchErrorWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range abews {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchErrorWizardModel, vv)
}

// UpdateAccountBatchErrorWizard updates an existing account.batch.error.wizard record.
func (c *Client) UpdateAccountBatchErrorWizard(abew *AccountBatchErrorWizard) error {
	return c.UpdateAccountBatchErrorWizards([]int64{abew.Id.Get()}, abew)
}

// UpdateAccountBatchErrorWizards updates existing account.batch.error.wizard records.
// All records (represented by ids) will be updated by abew values.
func (c *Client) UpdateAccountBatchErrorWizards(ids []int64, abew *AccountBatchErrorWizard) error {
	return c.Update(AccountBatchErrorWizardModel, ids, abew)
}

// DeleteAccountBatchErrorWizard deletes an existing account.batch.error.wizard record.
func (c *Client) DeleteAccountBatchErrorWizard(id int64) error {
	return c.DeleteAccountBatchErrorWizards([]int64{id})
}

// DeleteAccountBatchErrorWizards deletes existing account.batch.error.wizard records.
func (c *Client) DeleteAccountBatchErrorWizards(ids []int64) error {
	return c.Delete(AccountBatchErrorWizardModel, ids)
}

// GetAccountBatchErrorWizard gets account.batch.error.wizard existing record.
func (c *Client) GetAccountBatchErrorWizard(id int64) (*AccountBatchErrorWizard, error) {
	abews, err := c.GetAccountBatchErrorWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if abews != nil && len(*abews) > 0 {
		return &((*abews)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.batch.error.wizard not found", id)
}

// GetAccountBatchErrorWizards gets account.batch.error.wizard existing records.
func (c *Client) GetAccountBatchErrorWizards(ids []int64) (*AccountBatchErrorWizards, error) {
	abews := &AccountBatchErrorWizards{}
	if err := c.Read(AccountBatchErrorWizardModel, ids, nil, abews); err != nil {
		return nil, err
	}
	return abews, nil
}

// FindAccountBatchErrorWizard finds account.batch.error.wizard record by querying it with criteria.
func (c *Client) FindAccountBatchErrorWizard(criteria *Criteria) (*AccountBatchErrorWizard, error) {
	abews := &AccountBatchErrorWizards{}
	if err := c.SearchRead(AccountBatchErrorWizardModel, criteria, NewOptions().Limit(1), abews); err != nil {
		return nil, err
	}
	if abews != nil && len(*abews) > 0 {
		return &((*abews)[0]), nil
	}
	return nil, fmt.Errorf("account.batch.error.wizard was not found with criteria %v", criteria)
}

// FindAccountBatchErrorWizards finds account.batch.error.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchErrorWizards(criteria *Criteria, options *Options) (*AccountBatchErrorWizards, error) {
	abews := &AccountBatchErrorWizards{}
	if err := c.SearchRead(AccountBatchErrorWizardModel, criteria, options, abews); err != nil {
		return nil, err
	}
	return abews, nil
}

// FindAccountBatchErrorWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchErrorWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBatchErrorWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBatchErrorWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchErrorWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchErrorWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.batch.error.wizard was not found with criteria %v and options %v", criteria, options)
}
