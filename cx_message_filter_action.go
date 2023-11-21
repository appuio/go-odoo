package odoo

import (
	"fmt"
)

// CxMessageFilterAction represents cx.message.filter.action model.
type CxMessageFilterAction struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Action            *Selection `xmlrpc:"action,omptempty"`
	AuthorReplaceId   *Many2One  `xmlrpc:"author_replace_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	FilterId          *Many2One  `xmlrpc:"filter_id,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	MessageField      *Selection `xmlrpc:"message_field,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	PartnerIds        *Relation  `xmlrpc:"partner_ids,omptempty"`
	PartnerReplaceIds *Relation  `xmlrpc:"partner_replace_ids,omptempty"`
	Sequence          *Int       `xmlrpc:"sequence,omptempty"`
	Value             *String    `xmlrpc:"value,omptempty"`
	ValueReplace      *String    `xmlrpc:"value_replace,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CxMessageFilterActions represents array of cx.message.filter.action model.
type CxMessageFilterActions []CxMessageFilterAction

// CxMessageFilterActionModel is the odoo model name.
const CxMessageFilterActionModel = "cx.message.filter.action"

// Many2One convert CxMessageFilterAction to *Many2One.
func (cmfa *CxMessageFilterAction) Many2One() *Many2One {
	return NewMany2One(cmfa.Id.Get(), "")
}

// CreateCxMessageFilterAction creates a new cx.message.filter.action model and returns its id.
func (c *Client) CreateCxMessageFilterAction(cmfa *CxMessageFilterAction) (int64, error) {
	ids, err := c.CreateCxMessageFilterActions([]*CxMessageFilterAction{cmfa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxMessageFilterAction creates a new cx.message.filter.action model and returns its id.
func (c *Client) CreateCxMessageFilterActions(cmfas []*CxMessageFilterAction) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmfas {
		vv = append(vv, v)
	}
	return c.Create(CxMessageFilterActionModel, vv)
}

// UpdateCxMessageFilterAction updates an existing cx.message.filter.action record.
func (c *Client) UpdateCxMessageFilterAction(cmfa *CxMessageFilterAction) error {
	return c.UpdateCxMessageFilterActions([]int64{cmfa.Id.Get()}, cmfa)
}

// UpdateCxMessageFilterActions updates existing cx.message.filter.action records.
// All records (represented by ids) will be updated by cmfa values.
func (c *Client) UpdateCxMessageFilterActions(ids []int64, cmfa *CxMessageFilterAction) error {
	return c.Update(CxMessageFilterActionModel, ids, cmfa)
}

// DeleteCxMessageFilterAction deletes an existing cx.message.filter.action record.
func (c *Client) DeleteCxMessageFilterAction(id int64) error {
	return c.DeleteCxMessageFilterActions([]int64{id})
}

// DeleteCxMessageFilterActions deletes existing cx.message.filter.action records.
func (c *Client) DeleteCxMessageFilterActions(ids []int64) error {
	return c.Delete(CxMessageFilterActionModel, ids)
}

// GetCxMessageFilterAction gets cx.message.filter.action existing record.
func (c *Client) GetCxMessageFilterAction(id int64) (*CxMessageFilterAction, error) {
	cmfas, err := c.GetCxMessageFilterActions([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmfas != nil && len(*cmfas) > 0 {
		return &((*cmfas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.message.filter.action not found", id)
}

// GetCxMessageFilterActions gets cx.message.filter.action existing records.
func (c *Client) GetCxMessageFilterActions(ids []int64) (*CxMessageFilterActions, error) {
	cmfas := &CxMessageFilterActions{}
	if err := c.Read(CxMessageFilterActionModel, ids, nil, cmfas); err != nil {
		return nil, err
	}
	return cmfas, nil
}

// FindCxMessageFilterAction finds cx.message.filter.action record by querying it with criteria.
func (c *Client) FindCxMessageFilterAction(criteria *Criteria) (*CxMessageFilterAction, error) {
	cmfas := &CxMessageFilterActions{}
	if err := c.SearchRead(CxMessageFilterActionModel, criteria, NewOptions().Limit(1), cmfas); err != nil {
		return nil, err
	}
	if cmfas != nil && len(*cmfas) > 0 {
		return &((*cmfas)[0]), nil
	}
	return nil, fmt.Errorf("cx.message.filter.action was not found with criteria %v", criteria)
}

// FindCxMessageFilterActions finds cx.message.filter.action records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterActions(criteria *Criteria, options *Options) (*CxMessageFilterActions, error) {
	cmfas := &CxMessageFilterActions{}
	if err := c.SearchRead(CxMessageFilterActionModel, criteria, options, cmfas); err != nil {
		return nil, err
	}
	return cmfas, nil
}

// FindCxMessageFilterActionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterActionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxMessageFilterActionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxMessageFilterActionId finds record id by querying it with criteria.
func (c *Client) FindCxMessageFilterActionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxMessageFilterActionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.message.filter.action was not found with criteria %v and options %v", criteria, options)
}
