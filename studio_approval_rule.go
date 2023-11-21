package odoo

import (
	"fmt"
)

// StudioApprovalRule represents studio.approval.rule model.
type StudioApprovalRule struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	ActionId      *Many2One `xmlrpc:"action_id,omptempty"`
	Active        *Bool     `xmlrpc:"active,omptempty"`
	CanValidate   *Bool     `xmlrpc:"can_validate,omptempty"`
	Conditional   *Bool     `xmlrpc:"conditional,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Domain        *String   `xmlrpc:"domain,omptempty"`
	EntriesCount  *Int      `xmlrpc:"entries_count,omptempty"`
	EntryIds      *Relation `xmlrpc:"entry_ids,omptempty"`
	ExclusiveUser *Bool     `xmlrpc:"exclusive_user,omptempty"`
	GroupId       *Many2One `xmlrpc:"group_id,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	Message       *String   `xmlrpc:"message,omptempty"`
	Method        *String   `xmlrpc:"method,omptempty"`
	ModelId       *Many2One `xmlrpc:"model_id,omptempty"`
	ModelName     *String   `xmlrpc:"model_name,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	ResponsibleId *Many2One `xmlrpc:"responsible_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StudioApprovalRules represents array of studio.approval.rule model.
type StudioApprovalRules []StudioApprovalRule

// StudioApprovalRuleModel is the odoo model name.
const StudioApprovalRuleModel = "studio.approval.rule"

// Many2One convert StudioApprovalRule to *Many2One.
func (sar *StudioApprovalRule) Many2One() *Many2One {
	return NewMany2One(sar.Id.Get(), "")
}

// CreateStudioApprovalRule creates a new studio.approval.rule model and returns its id.
func (c *Client) CreateStudioApprovalRule(sar *StudioApprovalRule) (int64, error) {
	ids, err := c.CreateStudioApprovalRules([]*StudioApprovalRule{sar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioApprovalRule creates a new studio.approval.rule model and returns its id.
func (c *Client) CreateStudioApprovalRules(sars []*StudioApprovalRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range sars {
		vv = append(vv, v)
	}
	return c.Create(StudioApprovalRuleModel, vv)
}

// UpdateStudioApprovalRule updates an existing studio.approval.rule record.
func (c *Client) UpdateStudioApprovalRule(sar *StudioApprovalRule) error {
	return c.UpdateStudioApprovalRules([]int64{sar.Id.Get()}, sar)
}

// UpdateStudioApprovalRules updates existing studio.approval.rule records.
// All records (represented by ids) will be updated by sar values.
func (c *Client) UpdateStudioApprovalRules(ids []int64, sar *StudioApprovalRule) error {
	return c.Update(StudioApprovalRuleModel, ids, sar)
}

// DeleteStudioApprovalRule deletes an existing studio.approval.rule record.
func (c *Client) DeleteStudioApprovalRule(id int64) error {
	return c.DeleteStudioApprovalRules([]int64{id})
}

// DeleteStudioApprovalRules deletes existing studio.approval.rule records.
func (c *Client) DeleteStudioApprovalRules(ids []int64) error {
	return c.Delete(StudioApprovalRuleModel, ids)
}

// GetStudioApprovalRule gets studio.approval.rule existing record.
func (c *Client) GetStudioApprovalRule(id int64) (*StudioApprovalRule, error) {
	sars, err := c.GetStudioApprovalRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if sars != nil && len(*sars) > 0 {
		return &((*sars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of studio.approval.rule not found", id)
}

// GetStudioApprovalRules gets studio.approval.rule existing records.
func (c *Client) GetStudioApprovalRules(ids []int64) (*StudioApprovalRules, error) {
	sars := &StudioApprovalRules{}
	if err := c.Read(StudioApprovalRuleModel, ids, nil, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRule finds studio.approval.rule record by querying it with criteria.
func (c *Client) FindStudioApprovalRule(criteria *Criteria) (*StudioApprovalRule, error) {
	sars := &StudioApprovalRules{}
	if err := c.SearchRead(StudioApprovalRuleModel, criteria, NewOptions().Limit(1), sars); err != nil {
		return nil, err
	}
	if sars != nil && len(*sars) > 0 {
		return &((*sars)[0]), nil
	}
	return nil, fmt.Errorf("studio.approval.rule was not found with criteria %v", criteria)
}

// FindStudioApprovalRules finds studio.approval.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRules(criteria *Criteria, options *Options) (*StudioApprovalRules, error) {
	sars := &StudioApprovalRules{}
	if err := c.SearchRead(StudioApprovalRuleModel, criteria, options, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StudioApprovalRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStudioApprovalRuleId finds record id by querying it with criteria.
func (c *Client) FindStudioApprovalRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioApprovalRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("studio.approval.rule was not found with criteria %v and options %v", criteria, options)
}
