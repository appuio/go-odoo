package odoo

import (
	"fmt"
)

// DocumentsWorkflowRule represents documents.workflow.rule model.
type DocumentsWorkflowRule struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omptempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omptempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omptempty"`
	ActivityOption                *Bool      `xmlrpc:"activity_option,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	ConditionType                 *Selection `xmlrpc:"condition_type,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateModel                   *Selection `xmlrpc:"create_model,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CriteriaOwnerId               *Many2One  `xmlrpc:"criteria_owner_id,omptempty"`
	CriteriaPartnerId             *Many2One  `xmlrpc:"criteria_partner_id,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	Domain                        *String    `xmlrpc:"domain,omptempty"`
	DomainFolderId                *Many2One  `xmlrpc:"domain_folder_id,omptempty"`
	ExcludedTagIds                *Relation  `xmlrpc:"excluded_tag_ids,omptempty"`
	FolderId                      *Many2One  `xmlrpc:"folder_id,omptempty"`
	HasOwnerActivity              *Bool      `xmlrpc:"has_owner_activity,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	LimitedToSingleRecord         *Bool      `xmlrpc:"limited_to_single_record,omptempty"`
	LinkModel                     *Many2One  `xmlrpc:"link_model,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	Note                          *String    `xmlrpc:"note,omptempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omptempty"`
	RemoveActivities              *Bool      `xmlrpc:"remove_activities,omptempty"`
	RequiredTagIds                *Relation  `xmlrpc:"required_tag_ids,omptempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omptempty"`
	TagActionIds                  *Relation  `xmlrpc:"tag_action_ids,omptempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DocumentsWorkflowRules represents array of documents.workflow.rule model.
type DocumentsWorkflowRules []DocumentsWorkflowRule

// DocumentsWorkflowRuleModel is the odoo model name.
const DocumentsWorkflowRuleModel = "documents.workflow.rule"

// Many2One convert DocumentsWorkflowRule to *Many2One.
func (dwr *DocumentsWorkflowRule) Many2One() *Many2One {
	return NewMany2One(dwr.Id.Get(), "")
}

// CreateDocumentsWorkflowRule creates a new documents.workflow.rule model and returns its id.
func (c *Client) CreateDocumentsWorkflowRule(dwr *DocumentsWorkflowRule) (int64, error) {
	ids, err := c.CreateDocumentsWorkflowRules([]*DocumentsWorkflowRule{dwr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsWorkflowRule creates a new documents.workflow.rule model and returns its id.
func (c *Client) CreateDocumentsWorkflowRules(dwrs []*DocumentsWorkflowRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range dwrs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsWorkflowRuleModel, vv)
}

// UpdateDocumentsWorkflowRule updates an existing documents.workflow.rule record.
func (c *Client) UpdateDocumentsWorkflowRule(dwr *DocumentsWorkflowRule) error {
	return c.UpdateDocumentsWorkflowRules([]int64{dwr.Id.Get()}, dwr)
}

// UpdateDocumentsWorkflowRules updates existing documents.workflow.rule records.
// All records (represented by ids) will be updated by dwr values.
func (c *Client) UpdateDocumentsWorkflowRules(ids []int64, dwr *DocumentsWorkflowRule) error {
	return c.Update(DocumentsWorkflowRuleModel, ids, dwr)
}

// DeleteDocumentsWorkflowRule deletes an existing documents.workflow.rule record.
func (c *Client) DeleteDocumentsWorkflowRule(id int64) error {
	return c.DeleteDocumentsWorkflowRules([]int64{id})
}

// DeleteDocumentsWorkflowRules deletes existing documents.workflow.rule records.
func (c *Client) DeleteDocumentsWorkflowRules(ids []int64) error {
	return c.Delete(DocumentsWorkflowRuleModel, ids)
}

// GetDocumentsWorkflowRule gets documents.workflow.rule existing record.
func (c *Client) GetDocumentsWorkflowRule(id int64) (*DocumentsWorkflowRule, error) {
	dwrs, err := c.GetDocumentsWorkflowRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if dwrs != nil && len(*dwrs) > 0 {
		return &((*dwrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.workflow.rule not found", id)
}

// GetDocumentsWorkflowRules gets documents.workflow.rule existing records.
func (c *Client) GetDocumentsWorkflowRules(ids []int64) (*DocumentsWorkflowRules, error) {
	dwrs := &DocumentsWorkflowRules{}
	if err := c.Read(DocumentsWorkflowRuleModel, ids, nil, dwrs); err != nil {
		return nil, err
	}
	return dwrs, nil
}

// FindDocumentsWorkflowRule finds documents.workflow.rule record by querying it with criteria.
func (c *Client) FindDocumentsWorkflowRule(criteria *Criteria) (*DocumentsWorkflowRule, error) {
	dwrs := &DocumentsWorkflowRules{}
	if err := c.SearchRead(DocumentsWorkflowRuleModel, criteria, NewOptions().Limit(1), dwrs); err != nil {
		return nil, err
	}
	if dwrs != nil && len(*dwrs) > 0 {
		return &((*dwrs)[0]), nil
	}
	return nil, fmt.Errorf("documents.workflow.rule was not found with criteria %v", criteria)
}

// FindDocumentsWorkflowRules finds documents.workflow.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsWorkflowRules(criteria *Criteria, options *Options) (*DocumentsWorkflowRules, error) {
	dwrs := &DocumentsWorkflowRules{}
	if err := c.SearchRead(DocumentsWorkflowRuleModel, criteria, options, dwrs); err != nil {
		return nil, err
	}
	return dwrs, nil
}

// FindDocumentsWorkflowRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsWorkflowRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsWorkflowRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsWorkflowRuleId finds record id by querying it with criteria.
func (c *Client) FindDocumentsWorkflowRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsWorkflowRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.workflow.rule was not found with criteria %v and options %v", criteria, options)
}
