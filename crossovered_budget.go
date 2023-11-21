package odoo

import (
	"fmt"
)

// CrossoveredBudget represents crossovered.budget model.
type CrossoveredBudget struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrossoveredBudgetLine       *Relation  `xmlrpc:"crossovered_budget_line,omptempty"`
	DateFrom                    *Time      `xmlrpc:"date_from,omptempty"`
	DateTo                      *Time      `xmlrpc:"date_to,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent              *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CrossoveredBudgets represents array of crossovered.budget model.
type CrossoveredBudgets []CrossoveredBudget

// CrossoveredBudgetModel is the odoo model name.
const CrossoveredBudgetModel = "crossovered.budget"

// Many2One convert CrossoveredBudget to *Many2One.
func (cb *CrossoveredBudget) Many2One() *Many2One {
	return NewMany2One(cb.Id.Get(), "")
}

// CreateCrossoveredBudget creates a new crossovered.budget model and returns its id.
func (c *Client) CreateCrossoveredBudget(cb *CrossoveredBudget) (int64, error) {
	ids, err := c.CreateCrossoveredBudgets([]*CrossoveredBudget{cb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrossoveredBudget creates a new crossovered.budget model and returns its id.
func (c *Client) CreateCrossoveredBudgets(cbs []*CrossoveredBudget) ([]int64, error) {
	var vv []interface{}
	for _, v := range cbs {
		vv = append(vv, v)
	}
	return c.Create(CrossoveredBudgetModel, vv)
}

// UpdateCrossoveredBudget updates an existing crossovered.budget record.
func (c *Client) UpdateCrossoveredBudget(cb *CrossoveredBudget) error {
	return c.UpdateCrossoveredBudgets([]int64{cb.Id.Get()}, cb)
}

// UpdateCrossoveredBudgets updates existing crossovered.budget records.
// All records (represented by ids) will be updated by cb values.
func (c *Client) UpdateCrossoveredBudgets(ids []int64, cb *CrossoveredBudget) error {
	return c.Update(CrossoveredBudgetModel, ids, cb)
}

// DeleteCrossoveredBudget deletes an existing crossovered.budget record.
func (c *Client) DeleteCrossoveredBudget(id int64) error {
	return c.DeleteCrossoveredBudgets([]int64{id})
}

// DeleteCrossoveredBudgets deletes existing crossovered.budget records.
func (c *Client) DeleteCrossoveredBudgets(ids []int64) error {
	return c.Delete(CrossoveredBudgetModel, ids)
}

// GetCrossoveredBudget gets crossovered.budget existing record.
func (c *Client) GetCrossoveredBudget(id int64) (*CrossoveredBudget, error) {
	cbs, err := c.GetCrossoveredBudgets([]int64{id})
	if err != nil {
		return nil, err
	}
	if cbs != nil && len(*cbs) > 0 {
		return &((*cbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crossovered.budget not found", id)
}

// GetCrossoveredBudgets gets crossovered.budget existing records.
func (c *Client) GetCrossoveredBudgets(ids []int64) (*CrossoveredBudgets, error) {
	cbs := &CrossoveredBudgets{}
	if err := c.Read(CrossoveredBudgetModel, ids, nil, cbs); err != nil {
		return nil, err
	}
	return cbs, nil
}

// FindCrossoveredBudget finds crossovered.budget record by querying it with criteria.
func (c *Client) FindCrossoveredBudget(criteria *Criteria) (*CrossoveredBudget, error) {
	cbs := &CrossoveredBudgets{}
	if err := c.SearchRead(CrossoveredBudgetModel, criteria, NewOptions().Limit(1), cbs); err != nil {
		return nil, err
	}
	if cbs != nil && len(*cbs) > 0 {
		return &((*cbs)[0]), nil
	}
	return nil, fmt.Errorf("crossovered.budget was not found with criteria %v", criteria)
}

// FindCrossoveredBudgets finds crossovered.budget records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgets(criteria *Criteria, options *Options) (*CrossoveredBudgets, error) {
	cbs := &CrossoveredBudgets{}
	if err := c.SearchRead(CrossoveredBudgetModel, criteria, options, cbs); err != nil {
		return nil, err
	}
	return cbs, nil
}

// FindCrossoveredBudgetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrossoveredBudgetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrossoveredBudgetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrossoveredBudgetId finds record id by querying it with criteria.
func (c *Client) FindCrossoveredBudgetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrossoveredBudgetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crossovered.budget was not found with criteria %v and options %v", criteria, options)
}
