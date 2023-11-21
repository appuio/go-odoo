package odoo

import (
	"fmt"
)

// AccountReportHorizontalGroup represents account.report.horizontal.group model.
type AccountReportHorizontalGroup struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ReportIds   *Relation `xmlrpc:"report_ids,omptempty"`
	RuleIds     *Relation `xmlrpc:"rule_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountReportHorizontalGroups represents array of account.report.horizontal.group model.
type AccountReportHorizontalGroups []AccountReportHorizontalGroup

// AccountReportHorizontalGroupModel is the odoo model name.
const AccountReportHorizontalGroupModel = "account.report.horizontal.group"

// Many2One convert AccountReportHorizontalGroup to *Many2One.
func (arhg *AccountReportHorizontalGroup) Many2One() *Many2One {
	return NewMany2One(arhg.Id.Get(), "")
}

// CreateAccountReportHorizontalGroup creates a new account.report.horizontal.group model and returns its id.
func (c *Client) CreateAccountReportHorizontalGroup(arhg *AccountReportHorizontalGroup) (int64, error) {
	ids, err := c.CreateAccountReportHorizontalGroups([]*AccountReportHorizontalGroup{arhg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportHorizontalGroup creates a new account.report.horizontal.group model and returns its id.
func (c *Client) CreateAccountReportHorizontalGroups(arhgs []*AccountReportHorizontalGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range arhgs {
		vv = append(vv, v)
	}
	return c.Create(AccountReportHorizontalGroupModel, vv)
}

// UpdateAccountReportHorizontalGroup updates an existing account.report.horizontal.group record.
func (c *Client) UpdateAccountReportHorizontalGroup(arhg *AccountReportHorizontalGroup) error {
	return c.UpdateAccountReportHorizontalGroups([]int64{arhg.Id.Get()}, arhg)
}

// UpdateAccountReportHorizontalGroups updates existing account.report.horizontal.group records.
// All records (represented by ids) will be updated by arhg values.
func (c *Client) UpdateAccountReportHorizontalGroups(ids []int64, arhg *AccountReportHorizontalGroup) error {
	return c.Update(AccountReportHorizontalGroupModel, ids, arhg)
}

// DeleteAccountReportHorizontalGroup deletes an existing account.report.horizontal.group record.
func (c *Client) DeleteAccountReportHorizontalGroup(id int64) error {
	return c.DeleteAccountReportHorizontalGroups([]int64{id})
}

// DeleteAccountReportHorizontalGroups deletes existing account.report.horizontal.group records.
func (c *Client) DeleteAccountReportHorizontalGroups(ids []int64) error {
	return c.Delete(AccountReportHorizontalGroupModel, ids)
}

// GetAccountReportHorizontalGroup gets account.report.horizontal.group existing record.
func (c *Client) GetAccountReportHorizontalGroup(id int64) (*AccountReportHorizontalGroup, error) {
	arhgs, err := c.GetAccountReportHorizontalGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if arhgs != nil && len(*arhgs) > 0 {
		return &((*arhgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.horizontal.group not found", id)
}

// GetAccountReportHorizontalGroups gets account.report.horizontal.group existing records.
func (c *Client) GetAccountReportHorizontalGroups(ids []int64) (*AccountReportHorizontalGroups, error) {
	arhgs := &AccountReportHorizontalGroups{}
	if err := c.Read(AccountReportHorizontalGroupModel, ids, nil, arhgs); err != nil {
		return nil, err
	}
	return arhgs, nil
}

// FindAccountReportHorizontalGroup finds account.report.horizontal.group record by querying it with criteria.
func (c *Client) FindAccountReportHorizontalGroup(criteria *Criteria) (*AccountReportHorizontalGroup, error) {
	arhgs := &AccountReportHorizontalGroups{}
	if err := c.SearchRead(AccountReportHorizontalGroupModel, criteria, NewOptions().Limit(1), arhgs); err != nil {
		return nil, err
	}
	if arhgs != nil && len(*arhgs) > 0 {
		return &((*arhgs)[0]), nil
	}
	return nil, fmt.Errorf("account.report.horizontal.group was not found with criteria %v", criteria)
}

// FindAccountReportHorizontalGroups finds account.report.horizontal.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportHorizontalGroups(criteria *Criteria, options *Options) (*AccountReportHorizontalGroups, error) {
	arhgs := &AccountReportHorizontalGroups{}
	if err := c.SearchRead(AccountReportHorizontalGroupModel, criteria, options, arhgs); err != nil {
		return nil, err
	}
	return arhgs, nil
}

// FindAccountReportHorizontalGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportHorizontalGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportHorizontalGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportHorizontalGroupId finds record id by querying it with criteria.
func (c *Client) FindAccountReportHorizontalGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportHorizontalGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.horizontal.group was not found with criteria %v and options %v", criteria, options)
}
