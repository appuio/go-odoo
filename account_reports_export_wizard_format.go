package odoo

import (
	"fmt"
)

// AccountReportsExportWizardFormat represents account_reports.export.wizard.format model.
type AccountReportsExportWizardFormat struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	ExportWizardId *Many2One `xmlrpc:"export_wizard_id,omptempty"`
	FunParam       *String   `xmlrpc:"fun_param,omptempty"`
	FunToCall      *String   `xmlrpc:"fun_to_call,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountReportsExportWizardFormats represents array of account_reports.export.wizard.format model.
type AccountReportsExportWizardFormats []AccountReportsExportWizardFormat

// AccountReportsExportWizardFormatModel is the odoo model name.
const AccountReportsExportWizardFormatModel = "account_reports.export.wizard.format"

// Many2One convert AccountReportsExportWizardFormat to *Many2One.
func (aewf *AccountReportsExportWizardFormat) Many2One() *Many2One {
	return NewMany2One(aewf.Id.Get(), "")
}

// CreateAccountReportsExportWizardFormat creates a new account_reports.export.wizard.format model and returns its id.
func (c *Client) CreateAccountReportsExportWizardFormat(aewf *AccountReportsExportWizardFormat) (int64, error) {
	ids, err := c.CreateAccountReportsExportWizardFormats([]*AccountReportsExportWizardFormat{aewf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportsExportWizardFormat creates a new account_reports.export.wizard.format model and returns its id.
func (c *Client) CreateAccountReportsExportWizardFormats(aewfs []*AccountReportsExportWizardFormat) ([]int64, error) {
	var vv []interface{}
	for _, v := range aewfs {
		vv = append(vv, v)
	}
	return c.Create(AccountReportsExportWizardFormatModel, vv)
}

// UpdateAccountReportsExportWizardFormat updates an existing account_reports.export.wizard.format record.
func (c *Client) UpdateAccountReportsExportWizardFormat(aewf *AccountReportsExportWizardFormat) error {
	return c.UpdateAccountReportsExportWizardFormats([]int64{aewf.Id.Get()}, aewf)
}

// UpdateAccountReportsExportWizardFormats updates existing account_reports.export.wizard.format records.
// All records (represented by ids) will be updated by aewf values.
func (c *Client) UpdateAccountReportsExportWizardFormats(ids []int64, aewf *AccountReportsExportWizardFormat) error {
	return c.Update(AccountReportsExportWizardFormatModel, ids, aewf)
}

// DeleteAccountReportsExportWizardFormat deletes an existing account_reports.export.wizard.format record.
func (c *Client) DeleteAccountReportsExportWizardFormat(id int64) error {
	return c.DeleteAccountReportsExportWizardFormats([]int64{id})
}

// DeleteAccountReportsExportWizardFormats deletes existing account_reports.export.wizard.format records.
func (c *Client) DeleteAccountReportsExportWizardFormats(ids []int64) error {
	return c.Delete(AccountReportsExportWizardFormatModel, ids)
}

// GetAccountReportsExportWizardFormat gets account_reports.export.wizard.format existing record.
func (c *Client) GetAccountReportsExportWizardFormat(id int64) (*AccountReportsExportWizardFormat, error) {
	aewfs, err := c.GetAccountReportsExportWizardFormats([]int64{id})
	if err != nil {
		return nil, err
	}
	if aewfs != nil && len(*aewfs) > 0 {
		return &((*aewfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account_reports.export.wizard.format not found", id)
}

// GetAccountReportsExportWizardFormats gets account_reports.export.wizard.format existing records.
func (c *Client) GetAccountReportsExportWizardFormats(ids []int64) (*AccountReportsExportWizardFormats, error) {
	aewfs := &AccountReportsExportWizardFormats{}
	if err := c.Read(AccountReportsExportWizardFormatModel, ids, nil, aewfs); err != nil {
		return nil, err
	}
	return aewfs, nil
}

// FindAccountReportsExportWizardFormat finds account_reports.export.wizard.format record by querying it with criteria.
func (c *Client) FindAccountReportsExportWizardFormat(criteria *Criteria) (*AccountReportsExportWizardFormat, error) {
	aewfs := &AccountReportsExportWizardFormats{}
	if err := c.SearchRead(AccountReportsExportWizardFormatModel, criteria, NewOptions().Limit(1), aewfs); err != nil {
		return nil, err
	}
	if aewfs != nil && len(*aewfs) > 0 {
		return &((*aewfs)[0]), nil
	}
	return nil, fmt.Errorf("account_reports.export.wizard.format was not found with criteria %v", criteria)
}

// FindAccountReportsExportWizardFormats finds account_reports.export.wizard.format records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportsExportWizardFormats(criteria *Criteria, options *Options) (*AccountReportsExportWizardFormats, error) {
	aewfs := &AccountReportsExportWizardFormats{}
	if err := c.SearchRead(AccountReportsExportWizardFormatModel, criteria, options, aewfs); err != nil {
		return nil, err
	}
	return aewfs, nil
}

// FindAccountReportsExportWizardFormatIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportsExportWizardFormatIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportsExportWizardFormatModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportsExportWizardFormatId finds record id by querying it with criteria.
func (c *Client) FindAccountReportsExportWizardFormatId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportsExportWizardFormatModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account_reports.export.wizard.format was not found with criteria %v and options %v", criteria, options)
}
