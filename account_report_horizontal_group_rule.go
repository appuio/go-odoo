package odoo

import (
	"fmt"
)

// AccountReportHorizontalGroupRule represents account.report.horizontal.group.rule model.
type AccountReportHorizontalGroupRule struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Domain            *String    `xmlrpc:"domain,omptempty"`
	FieldName         *Selection `xmlrpc:"field_name,omptempty"`
	HorizontalGroupId *Many2One  `xmlrpc:"horizontal_group_id,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	ResModelName      *String    `xmlrpc:"res_model_name,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReportHorizontalGroupRules represents array of account.report.horizontal.group.rule model.
type AccountReportHorizontalGroupRules []AccountReportHorizontalGroupRule

// AccountReportHorizontalGroupRuleModel is the odoo model name.
const AccountReportHorizontalGroupRuleModel = "account.report.horizontal.group.rule"

// Many2One convert AccountReportHorizontalGroupRule to *Many2One.
func (arhgr *AccountReportHorizontalGroupRule) Many2One() *Many2One {
	return NewMany2One(arhgr.Id.Get(), "")
}

// CreateAccountReportHorizontalGroupRule creates a new account.report.horizontal.group.rule model and returns its id.
func (c *Client) CreateAccountReportHorizontalGroupRule(arhgr *AccountReportHorizontalGroupRule) (int64, error) {
	ids, err := c.CreateAccountReportHorizontalGroupRules([]*AccountReportHorizontalGroupRule{arhgr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReportHorizontalGroupRule creates a new account.report.horizontal.group.rule model and returns its id.
func (c *Client) CreateAccountReportHorizontalGroupRules(arhgrs []*AccountReportHorizontalGroupRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range arhgrs {
		vv = append(vv, v)
	}
	return c.Create(AccountReportHorizontalGroupRuleModel, vv)
}

// UpdateAccountReportHorizontalGroupRule updates an existing account.report.horizontal.group.rule record.
func (c *Client) UpdateAccountReportHorizontalGroupRule(arhgr *AccountReportHorizontalGroupRule) error {
	return c.UpdateAccountReportHorizontalGroupRules([]int64{arhgr.Id.Get()}, arhgr)
}

// UpdateAccountReportHorizontalGroupRules updates existing account.report.horizontal.group.rule records.
// All records (represented by ids) will be updated by arhgr values.
func (c *Client) UpdateAccountReportHorizontalGroupRules(ids []int64, arhgr *AccountReportHorizontalGroupRule) error {
	return c.Update(AccountReportHorizontalGroupRuleModel, ids, arhgr)
}

// DeleteAccountReportHorizontalGroupRule deletes an existing account.report.horizontal.group.rule record.
func (c *Client) DeleteAccountReportHorizontalGroupRule(id int64) error {
	return c.DeleteAccountReportHorizontalGroupRules([]int64{id})
}

// DeleteAccountReportHorizontalGroupRules deletes existing account.report.horizontal.group.rule records.
func (c *Client) DeleteAccountReportHorizontalGroupRules(ids []int64) error {
	return c.Delete(AccountReportHorizontalGroupRuleModel, ids)
}

// GetAccountReportHorizontalGroupRule gets account.report.horizontal.group.rule existing record.
func (c *Client) GetAccountReportHorizontalGroupRule(id int64) (*AccountReportHorizontalGroupRule, error) {
	arhgrs, err := c.GetAccountReportHorizontalGroupRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if arhgrs != nil && len(*arhgrs) > 0 {
		return &((*arhgrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.horizontal.group.rule not found", id)
}

// GetAccountReportHorizontalGroupRules gets account.report.horizontal.group.rule existing records.
func (c *Client) GetAccountReportHorizontalGroupRules(ids []int64) (*AccountReportHorizontalGroupRules, error) {
	arhgrs := &AccountReportHorizontalGroupRules{}
	if err := c.Read(AccountReportHorizontalGroupRuleModel, ids, nil, arhgrs); err != nil {
		return nil, err
	}
	return arhgrs, nil
}

// FindAccountReportHorizontalGroupRule finds account.report.horizontal.group.rule record by querying it with criteria.
func (c *Client) FindAccountReportHorizontalGroupRule(criteria *Criteria) (*AccountReportHorizontalGroupRule, error) {
	arhgrs := &AccountReportHorizontalGroupRules{}
	if err := c.SearchRead(AccountReportHorizontalGroupRuleModel, criteria, NewOptions().Limit(1), arhgrs); err != nil {
		return nil, err
	}
	if arhgrs != nil && len(*arhgrs) > 0 {
		return &((*arhgrs)[0]), nil
	}
	return nil, fmt.Errorf("account.report.horizontal.group.rule was not found with criteria %v", criteria)
}

// FindAccountReportHorizontalGroupRules finds account.report.horizontal.group.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportHorizontalGroupRules(criteria *Criteria, options *Options) (*AccountReportHorizontalGroupRules, error) {
	arhgrs := &AccountReportHorizontalGroupRules{}
	if err := c.SearchRead(AccountReportHorizontalGroupRuleModel, criteria, options, arhgrs); err != nil {
		return nil, err
	}
	return arhgrs, nil
}

// FindAccountReportHorizontalGroupRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportHorizontalGroupRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportHorizontalGroupRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportHorizontalGroupRuleId finds record id by querying it with criteria.
func (c *Client) FindAccountReportHorizontalGroupRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportHorizontalGroupRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.horizontal.group.rule was not found with criteria %v and options %v", criteria, options)
}
