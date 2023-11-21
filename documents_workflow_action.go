package odoo

import (
	"fmt"
)

// DocumentsWorkflowAction represents documents.workflow.action model.
type DocumentsWorkflowAction struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Action         *Selection `xmlrpc:"action,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	FacetId        *Many2One  `xmlrpc:"facet_id,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	TagId          *Many2One  `xmlrpc:"tag_id,omptempty"`
	WorkflowRuleId *Many2One  `xmlrpc:"workflow_rule_id,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DocumentsWorkflowActions represents array of documents.workflow.action model.
type DocumentsWorkflowActions []DocumentsWorkflowAction

// DocumentsWorkflowActionModel is the odoo model name.
const DocumentsWorkflowActionModel = "documents.workflow.action"

// Many2One convert DocumentsWorkflowAction to *Many2One.
func (dwa *DocumentsWorkflowAction) Many2One() *Many2One {
	return NewMany2One(dwa.Id.Get(), "")
}

// CreateDocumentsWorkflowAction creates a new documents.workflow.action model and returns its id.
func (c *Client) CreateDocumentsWorkflowAction(dwa *DocumentsWorkflowAction) (int64, error) {
	ids, err := c.CreateDocumentsWorkflowActions([]*DocumentsWorkflowAction{dwa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsWorkflowAction creates a new documents.workflow.action model and returns its id.
func (c *Client) CreateDocumentsWorkflowActions(dwas []*DocumentsWorkflowAction) ([]int64, error) {
	var vv []interface{}
	for _, v := range dwas {
		vv = append(vv, v)
	}
	return c.Create(DocumentsWorkflowActionModel, vv)
}

// UpdateDocumentsWorkflowAction updates an existing documents.workflow.action record.
func (c *Client) UpdateDocumentsWorkflowAction(dwa *DocumentsWorkflowAction) error {
	return c.UpdateDocumentsWorkflowActions([]int64{dwa.Id.Get()}, dwa)
}

// UpdateDocumentsWorkflowActions updates existing documents.workflow.action records.
// All records (represented by ids) will be updated by dwa values.
func (c *Client) UpdateDocumentsWorkflowActions(ids []int64, dwa *DocumentsWorkflowAction) error {
	return c.Update(DocumentsWorkflowActionModel, ids, dwa)
}

// DeleteDocumentsWorkflowAction deletes an existing documents.workflow.action record.
func (c *Client) DeleteDocumentsWorkflowAction(id int64) error {
	return c.DeleteDocumentsWorkflowActions([]int64{id})
}

// DeleteDocumentsWorkflowActions deletes existing documents.workflow.action records.
func (c *Client) DeleteDocumentsWorkflowActions(ids []int64) error {
	return c.Delete(DocumentsWorkflowActionModel, ids)
}

// GetDocumentsWorkflowAction gets documents.workflow.action existing record.
func (c *Client) GetDocumentsWorkflowAction(id int64) (*DocumentsWorkflowAction, error) {
	dwas, err := c.GetDocumentsWorkflowActions([]int64{id})
	if err != nil {
		return nil, err
	}
	if dwas != nil && len(*dwas) > 0 {
		return &((*dwas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.workflow.action not found", id)
}

// GetDocumentsWorkflowActions gets documents.workflow.action existing records.
func (c *Client) GetDocumentsWorkflowActions(ids []int64) (*DocumentsWorkflowActions, error) {
	dwas := &DocumentsWorkflowActions{}
	if err := c.Read(DocumentsWorkflowActionModel, ids, nil, dwas); err != nil {
		return nil, err
	}
	return dwas, nil
}

// FindDocumentsWorkflowAction finds documents.workflow.action record by querying it with criteria.
func (c *Client) FindDocumentsWorkflowAction(criteria *Criteria) (*DocumentsWorkflowAction, error) {
	dwas := &DocumentsWorkflowActions{}
	if err := c.SearchRead(DocumentsWorkflowActionModel, criteria, NewOptions().Limit(1), dwas); err != nil {
		return nil, err
	}
	if dwas != nil && len(*dwas) > 0 {
		return &((*dwas)[0]), nil
	}
	return nil, fmt.Errorf("documents.workflow.action was not found with criteria %v", criteria)
}

// FindDocumentsWorkflowActions finds documents.workflow.action records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsWorkflowActions(criteria *Criteria, options *Options) (*DocumentsWorkflowActions, error) {
	dwas := &DocumentsWorkflowActions{}
	if err := c.SearchRead(DocumentsWorkflowActionModel, criteria, options, dwas); err != nil {
		return nil, err
	}
	return dwas, nil
}

// FindDocumentsWorkflowActionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsWorkflowActionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsWorkflowActionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsWorkflowActionId finds record id by querying it with criteria.
func (c *Client) FindDocumentsWorkflowActionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsWorkflowActionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.workflow.action was not found with criteria %v and options %v", criteria, options)
}
