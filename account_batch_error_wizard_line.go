package odoo

import (
	"fmt"
)

// AccountBatchErrorWizardLine represents account.batch.error.wizard.line model.
type AccountBatchErrorWizardLine struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	Description      *String   `xmlrpc:"description,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	ErrorWizardId    *Many2One `xmlrpc:"error_wizard_id,omptempty"`
	HelpMessage      *String   `xmlrpc:"help_message,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	PaymentIds       *Relation `xmlrpc:"payment_ids,omptempty"`
	ShowRemoveButton *Bool     `xmlrpc:"show_remove_button,omptempty"`
	WarningWizardId  *Many2One `xmlrpc:"warning_wizard_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBatchErrorWizardLines represents array of account.batch.error.wizard.line model.
type AccountBatchErrorWizardLines []AccountBatchErrorWizardLine

// AccountBatchErrorWizardLineModel is the odoo model name.
const AccountBatchErrorWizardLineModel = "account.batch.error.wizard.line"

// Many2One convert AccountBatchErrorWizardLine to *Many2One.
func (abewl *AccountBatchErrorWizardLine) Many2One() *Many2One {
	return NewMany2One(abewl.Id.Get(), "")
}

// CreateAccountBatchErrorWizardLine creates a new account.batch.error.wizard.line model and returns its id.
func (c *Client) CreateAccountBatchErrorWizardLine(abewl *AccountBatchErrorWizardLine) (int64, error) {
	ids, err := c.CreateAccountBatchErrorWizardLines([]*AccountBatchErrorWizardLine{abewl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchErrorWizardLine creates a new account.batch.error.wizard.line model and returns its id.
func (c *Client) CreateAccountBatchErrorWizardLines(abewls []*AccountBatchErrorWizardLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range abewls {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchErrorWizardLineModel, vv)
}

// UpdateAccountBatchErrorWizardLine updates an existing account.batch.error.wizard.line record.
func (c *Client) UpdateAccountBatchErrorWizardLine(abewl *AccountBatchErrorWizardLine) error {
	return c.UpdateAccountBatchErrorWizardLines([]int64{abewl.Id.Get()}, abewl)
}

// UpdateAccountBatchErrorWizardLines updates existing account.batch.error.wizard.line records.
// All records (represented by ids) will be updated by abewl values.
func (c *Client) UpdateAccountBatchErrorWizardLines(ids []int64, abewl *AccountBatchErrorWizardLine) error {
	return c.Update(AccountBatchErrorWizardLineModel, ids, abewl)
}

// DeleteAccountBatchErrorWizardLine deletes an existing account.batch.error.wizard.line record.
func (c *Client) DeleteAccountBatchErrorWizardLine(id int64) error {
	return c.DeleteAccountBatchErrorWizardLines([]int64{id})
}

// DeleteAccountBatchErrorWizardLines deletes existing account.batch.error.wizard.line records.
func (c *Client) DeleteAccountBatchErrorWizardLines(ids []int64) error {
	return c.Delete(AccountBatchErrorWizardLineModel, ids)
}

// GetAccountBatchErrorWizardLine gets account.batch.error.wizard.line existing record.
func (c *Client) GetAccountBatchErrorWizardLine(id int64) (*AccountBatchErrorWizardLine, error) {
	abewls, err := c.GetAccountBatchErrorWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if abewls != nil && len(*abewls) > 0 {
		return &((*abewls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.batch.error.wizard.line not found", id)
}

// GetAccountBatchErrorWizardLines gets account.batch.error.wizard.line existing records.
func (c *Client) GetAccountBatchErrorWizardLines(ids []int64) (*AccountBatchErrorWizardLines, error) {
	abewls := &AccountBatchErrorWizardLines{}
	if err := c.Read(AccountBatchErrorWizardLineModel, ids, nil, abewls); err != nil {
		return nil, err
	}
	return abewls, nil
}

// FindAccountBatchErrorWizardLine finds account.batch.error.wizard.line record by querying it with criteria.
func (c *Client) FindAccountBatchErrorWizardLine(criteria *Criteria) (*AccountBatchErrorWizardLine, error) {
	abewls := &AccountBatchErrorWizardLines{}
	if err := c.SearchRead(AccountBatchErrorWizardLineModel, criteria, NewOptions().Limit(1), abewls); err != nil {
		return nil, err
	}
	if abewls != nil && len(*abewls) > 0 {
		return &((*abewls)[0]), nil
	}
	return nil, fmt.Errorf("account.batch.error.wizard.line was not found with criteria %v", criteria)
}

// FindAccountBatchErrorWizardLines finds account.batch.error.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchErrorWizardLines(criteria *Criteria, options *Options) (*AccountBatchErrorWizardLines, error) {
	abewls := &AccountBatchErrorWizardLines{}
	if err := c.SearchRead(AccountBatchErrorWizardLineModel, criteria, options, abewls); err != nil {
		return nil, err
	}
	return abewls, nil
}

// FindAccountBatchErrorWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchErrorWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBatchErrorWizardLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBatchErrorWizardLineId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchErrorWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchErrorWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.batch.error.wizard.line was not found with criteria %v and options %v", criteria, options)
}
