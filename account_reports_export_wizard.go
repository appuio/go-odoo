package odoo

import (
	"fmt"
)

// AccountReportsExportWizard represents account_reports.export.wizard model.
type AccountReportsExportWizard struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	DocName         *String   `xmlrpc:"doc_name,omptempty"`
	ExportFormatIds *Relation `xmlrpc:"export_format_ids,omptempty"`
	FolderId        *Many2One `xmlrpc:"folder_id,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	ReportId        *Many2One `xmlrpc:"report_id,omptempty"`
	TagIds          *Relation `xmlrpc:"tag_ids,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountReportsExportWizards represents array of account_reports.export.wizard model.
type AccountReportsExportWizards []AccountReportsExportWizard

// AccountReportsExportWizardModel is the odoo model name.
const AccountReportsExportWizardModel = "account_reports.export.wizard"

// Many2One convert AccountReportsExportWizard to *Many2One.
func (aew *AccountReportsExportWizard) Many2One() *Many2One {
	return NewMany2One(aew.Id.Get(), "")
}

// CreateAccountReportsExportWizard creates a new account_reports.export.wizard model and returns its id.
func (c *Client) CreateAccountReportsExportWizard(aew *AccountReportsExportWizard) (int64, error) {
	ids, err := c.CreateAccountReportsExportWizards([]*AccountReportsExportWizard{aew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportsExportWizard creates a new account_reports.export.wizard model and returns its id.
func (c *Client) CreateAccountReportsExportWizards(aews []*AccountReportsExportWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aews {
		vv = append(vv, v)
	}
	return c.Create(AccountReportsExportWizardModel, vv)
}

// UpdateAccountReportsExportWizard updates an existing account_reports.export.wizard record.
func (c *Client) UpdateAccountReportsExportWizard(aew *AccountReportsExportWizard) error {
	return c.UpdateAccountReportsExportWizards([]int64{aew.Id.Get()}, aew)
}

// UpdateAccountReportsExportWizards updates existing account_reports.export.wizard records.
// All records (represented by ids) will be updated by aew values.
func (c *Client) UpdateAccountReportsExportWizards(ids []int64, aew *AccountReportsExportWizard) error {
	return c.Update(AccountReportsExportWizardModel, ids, aew)
}

// DeleteAccountReportsExportWizard deletes an existing account_reports.export.wizard record.
func (c *Client) DeleteAccountReportsExportWizard(id int64) error {
	return c.DeleteAccountReportsExportWizards([]int64{id})
}

// DeleteAccountReportsExportWizards deletes existing account_reports.export.wizard records.
func (c *Client) DeleteAccountReportsExportWizards(ids []int64) error {
	return c.Delete(AccountReportsExportWizardModel, ids)
}

// GetAccountReportsExportWizard gets account_reports.export.wizard existing record.
func (c *Client) GetAccountReportsExportWizard(id int64) (*AccountReportsExportWizard, error) {
	aews, err := c.GetAccountReportsExportWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if aews != nil && len(*aews) > 0 {
		return &((*aews)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account_reports.export.wizard not found", id)
}

// GetAccountReportsExportWizards gets account_reports.export.wizard existing records.
func (c *Client) GetAccountReportsExportWizards(ids []int64) (*AccountReportsExportWizards, error) {
	aews := &AccountReportsExportWizards{}
	if err := c.Read(AccountReportsExportWizardModel, ids, nil, aews); err != nil {
		return nil, err
	}
	return aews, nil
}

// FindAccountReportsExportWizard finds account_reports.export.wizard record by querying it with criteria.
func (c *Client) FindAccountReportsExportWizard(criteria *Criteria) (*AccountReportsExportWizard, error) {
	aews := &AccountReportsExportWizards{}
	if err := c.SearchRead(AccountReportsExportWizardModel, criteria, NewOptions().Limit(1), aews); err != nil {
		return nil, err
	}
	if aews != nil && len(*aews) > 0 {
		return &((*aews)[0]), nil
	}
	return nil, fmt.Errorf("account_reports.export.wizard was not found with criteria %v", criteria)
}

// FindAccountReportsExportWizards finds account_reports.export.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportsExportWizards(criteria *Criteria, options *Options) (*AccountReportsExportWizards, error) {
	aews := &AccountReportsExportWizards{}
	if err := c.SearchRead(AccountReportsExportWizardModel, criteria, options, aews); err != nil {
		return nil, err
	}
	return aews, nil
}

// FindAccountReportsExportWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportsExportWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportsExportWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportsExportWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountReportsExportWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportsExportWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account_reports.export.wizard was not found with criteria %v and options %v", criteria, options)
}
