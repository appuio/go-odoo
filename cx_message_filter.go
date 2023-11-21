package odoo

import (
	"fmt"
)

// CxMessageFilter represents cx.message.filter model.
type CxMessageFilter struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Action                   *Selection `xmlrpc:"action,omptempty"`
	ActionIds                *Relation  `xmlrpc:"action_ids,omptempty"`
	Active                   *Bool      `xmlrpc:"active,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreatePartner            *Bool      `xmlrpc:"create_partner,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomValues             *String    `xmlrpc:"custom_values,omptempty"`
	DestinationModelId       *Many2One  `xmlrpc:"destination_model_id,omptempty"`
	DestinationRec           *String    `xmlrpc:"destination_rec,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FailedMessageIds         *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent           *String    `xmlrpc:"message_content,omptempty"`
	MessageCount             *Int       `xmlrpc:"message_count,omptempty"`
	MessageFilteredIds       *Relation  `xmlrpc:"message_filtered_ids,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	Order                    *Int       `xmlrpc:"order,omptempty"`
	RuleIds                  *Relation  `xmlrpc:"rule_ids,omptempty"`
	SearchBody               *Bool      `xmlrpc:"search_body,omptempty"`
	SearchFieldPattern       *Relation  `xmlrpc:"search_field_pattern,omptempty"`
	SearchRecord             *Bool      `xmlrpc:"search_record,omptempty"`
	SearchSubject            *Bool      `xmlrpc:"search_subject,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CxMessageFilters represents array of cx.message.filter model.
type CxMessageFilters []CxMessageFilter

// CxMessageFilterModel is the odoo model name.
const CxMessageFilterModel = "cx.message.filter"

// Many2One convert CxMessageFilter to *Many2One.
func (cmf *CxMessageFilter) Many2One() *Many2One {
	return NewMany2One(cmf.Id.Get(), "")
}

// CreateCxMessageFilter creates a new cx.message.filter model and returns its id.
func (c *Client) CreateCxMessageFilter(cmf *CxMessageFilter) (int64, error) {
	ids, err := c.CreateCxMessageFilters([]*CxMessageFilter{cmf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxMessageFilter creates a new cx.message.filter model and returns its id.
func (c *Client) CreateCxMessageFilters(cmfs []*CxMessageFilter) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmfs {
		vv = append(vv, v)
	}
	return c.Create(CxMessageFilterModel, vv)
}

// UpdateCxMessageFilter updates an existing cx.message.filter record.
func (c *Client) UpdateCxMessageFilter(cmf *CxMessageFilter) error {
	return c.UpdateCxMessageFilters([]int64{cmf.Id.Get()}, cmf)
}

// UpdateCxMessageFilters updates existing cx.message.filter records.
// All records (represented by ids) will be updated by cmf values.
func (c *Client) UpdateCxMessageFilters(ids []int64, cmf *CxMessageFilter) error {
	return c.Update(CxMessageFilterModel, ids, cmf)
}

// DeleteCxMessageFilter deletes an existing cx.message.filter record.
func (c *Client) DeleteCxMessageFilter(id int64) error {
	return c.DeleteCxMessageFilters([]int64{id})
}

// DeleteCxMessageFilters deletes existing cx.message.filter records.
func (c *Client) DeleteCxMessageFilters(ids []int64) error {
	return c.Delete(CxMessageFilterModel, ids)
}

// GetCxMessageFilter gets cx.message.filter existing record.
func (c *Client) GetCxMessageFilter(id int64) (*CxMessageFilter, error) {
	cmfs, err := c.GetCxMessageFilters([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmfs != nil && len(*cmfs) > 0 {
		return &((*cmfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.message.filter not found", id)
}

// GetCxMessageFilters gets cx.message.filter existing records.
func (c *Client) GetCxMessageFilters(ids []int64) (*CxMessageFilters, error) {
	cmfs := &CxMessageFilters{}
	if err := c.Read(CxMessageFilterModel, ids, nil, cmfs); err != nil {
		return nil, err
	}
	return cmfs, nil
}

// FindCxMessageFilter finds cx.message.filter record by querying it with criteria.
func (c *Client) FindCxMessageFilter(criteria *Criteria) (*CxMessageFilter, error) {
	cmfs := &CxMessageFilters{}
	if err := c.SearchRead(CxMessageFilterModel, criteria, NewOptions().Limit(1), cmfs); err != nil {
		return nil, err
	}
	if cmfs != nil && len(*cmfs) > 0 {
		return &((*cmfs)[0]), nil
	}
	return nil, fmt.Errorf("cx.message.filter was not found with criteria %v", criteria)
}

// FindCxMessageFilters finds cx.message.filter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilters(criteria *Criteria, options *Options) (*CxMessageFilters, error) {
	cmfs := &CxMessageFilters{}
	if err := c.SearchRead(CxMessageFilterModel, criteria, options, cmfs); err != nil {
		return nil, err
	}
	return cmfs, nil
}

// FindCxMessageFilterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxMessageFilterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxMessageFilterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxMessageFilterId finds record id by querying it with criteria.
func (c *Client) FindCxMessageFilterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxMessageFilterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.message.filter was not found with criteria %v and options %v", criteria, options)
}
