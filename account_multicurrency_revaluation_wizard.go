package odoo

import (
	"fmt"
)

// AccountMulticurrencyRevaluationWizard represents account.multicurrency.revaluation.wizard model.
type AccountMulticurrencyRevaluationWizard struct {
	LastUpdate                *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId                 *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omptempty"`
	Date                      *Time     `xmlrpc:"date,omptempty"`
	DisplayName               *String   `xmlrpc:"display_name,omptempty"`
	ExpenseProvisionAccountId *Many2One `xmlrpc:"expense_provision_account_id,omptempty"`
	Id                        *Int      `xmlrpc:"id,omptempty"`
	IncomeProvisionAccountId  *Many2One `xmlrpc:"income_provision_account_id,omptempty"`
	JournalId                 *Many2One `xmlrpc:"journal_id,omptempty"`
	PreviewData               *String   `xmlrpc:"preview_data,omptempty"`
	ReversalDate              *Time     `xmlrpc:"reversal_date,omptempty"`
	ShowWarningMoveId         *Many2One `xmlrpc:"show_warning_move_id,omptempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountMulticurrencyRevaluationWizards represents array of account.multicurrency.revaluation.wizard model.
type AccountMulticurrencyRevaluationWizards []AccountMulticurrencyRevaluationWizard

// AccountMulticurrencyRevaluationWizardModel is the odoo model name.
const AccountMulticurrencyRevaluationWizardModel = "account.multicurrency.revaluation.wizard"

// Many2One convert AccountMulticurrencyRevaluationWizard to *Many2One.
func (amrw *AccountMulticurrencyRevaluationWizard) Many2One() *Many2One {
	return NewMany2One(amrw.Id.Get(), "")
}

// CreateAccountMulticurrencyRevaluationWizard creates a new account.multicurrency.revaluation.wizard model and returns its id.
func (c *Client) CreateAccountMulticurrencyRevaluationWizard(amrw *AccountMulticurrencyRevaluationWizard) (int64, error) {
	ids, err := c.CreateAccountMulticurrencyRevaluationWizards([]*AccountMulticurrencyRevaluationWizard{amrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountMulticurrencyRevaluationWizard creates a new account.multicurrency.revaluation.wizard model and returns its id.
func (c *Client) CreateAccountMulticurrencyRevaluationWizards(amrws []*AccountMulticurrencyRevaluationWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range amrws {
		vv = append(vv, v)
	}
	return c.Create(AccountMulticurrencyRevaluationWizardModel, vv)
}

// UpdateAccountMulticurrencyRevaluationWizard updates an existing account.multicurrency.revaluation.wizard record.
func (c *Client) UpdateAccountMulticurrencyRevaluationWizard(amrw *AccountMulticurrencyRevaluationWizard) error {
	return c.UpdateAccountMulticurrencyRevaluationWizards([]int64{amrw.Id.Get()}, amrw)
}

// UpdateAccountMulticurrencyRevaluationWizards updates existing account.multicurrency.revaluation.wizard records.
// All records (represented by ids) will be updated by amrw values.
func (c *Client) UpdateAccountMulticurrencyRevaluationWizards(ids []int64, amrw *AccountMulticurrencyRevaluationWizard) error {
	return c.Update(AccountMulticurrencyRevaluationWizardModel, ids, amrw)
}

// DeleteAccountMulticurrencyRevaluationWizard deletes an existing account.multicurrency.revaluation.wizard record.
func (c *Client) DeleteAccountMulticurrencyRevaluationWizard(id int64) error {
	return c.DeleteAccountMulticurrencyRevaluationWizards([]int64{id})
}

// DeleteAccountMulticurrencyRevaluationWizards deletes existing account.multicurrency.revaluation.wizard records.
func (c *Client) DeleteAccountMulticurrencyRevaluationWizards(ids []int64) error {
	return c.Delete(AccountMulticurrencyRevaluationWizardModel, ids)
}

// GetAccountMulticurrencyRevaluationWizard gets account.multicurrency.revaluation.wizard existing record.
func (c *Client) GetAccountMulticurrencyRevaluationWizard(id int64) (*AccountMulticurrencyRevaluationWizard, error) {
	amrws, err := c.GetAccountMulticurrencyRevaluationWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if amrws != nil && len(*amrws) > 0 {
		return &((*amrws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.multicurrency.revaluation.wizard not found", id)
}

// GetAccountMulticurrencyRevaluationWizards gets account.multicurrency.revaluation.wizard existing records.
func (c *Client) GetAccountMulticurrencyRevaluationWizards(ids []int64) (*AccountMulticurrencyRevaluationWizards, error) {
	amrws := &AccountMulticurrencyRevaluationWizards{}
	if err := c.Read(AccountMulticurrencyRevaluationWizardModel, ids, nil, amrws); err != nil {
		return nil, err
	}
	return amrws, nil
}

// FindAccountMulticurrencyRevaluationWizard finds account.multicurrency.revaluation.wizard record by querying it with criteria.
func (c *Client) FindAccountMulticurrencyRevaluationWizard(criteria *Criteria) (*AccountMulticurrencyRevaluationWizard, error) {
	amrws := &AccountMulticurrencyRevaluationWizards{}
	if err := c.SearchRead(AccountMulticurrencyRevaluationWizardModel, criteria, NewOptions().Limit(1), amrws); err != nil {
		return nil, err
	}
	if amrws != nil && len(*amrws) > 0 {
		return &((*amrws)[0]), nil
	}
	return nil, fmt.Errorf("account.multicurrency.revaluation.wizard was not found with criteria %v", criteria)
}

// FindAccountMulticurrencyRevaluationWizards finds account.multicurrency.revaluation.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMulticurrencyRevaluationWizards(criteria *Criteria, options *Options) (*AccountMulticurrencyRevaluationWizards, error) {
	amrws := &AccountMulticurrencyRevaluationWizards{}
	if err := c.SearchRead(AccountMulticurrencyRevaluationWizardModel, criteria, options, amrws); err != nil {
		return nil, err
	}
	return amrws, nil
}

// FindAccountMulticurrencyRevaluationWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountMulticurrencyRevaluationWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountMulticurrencyRevaluationWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountMulticurrencyRevaluationWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountMulticurrencyRevaluationWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountMulticurrencyRevaluationWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.multicurrency.revaluation.wizard was not found with criteria %v and options %v", criteria, options)
}
