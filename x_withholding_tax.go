package odoo

import (
	"fmt"
)

// XWithholdingTax represents x_withholding_tax model.
type XWithholdingTax struct {
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
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
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
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
	XActive                     *Bool      `xmlrpc:"x_active,omptempty"`
	XName                       *String    `xmlrpc:"x_name,omptempty"`
	XStudioCanton               *Selection `xmlrpc:"x_studio_canton,omptempty"`
	XStudioCanton1              *Selection `xmlrpc:"x_studio_canton_1,omptempty"`
	XStudioPartnerId            *Many2One  `xmlrpc:"x_studio_partner_id,omptempty"`
	XStudioSequence             *Int       `xmlrpc:"x_studio_sequence,omptempty"`
}

// XWithholdingTaxs represents array of x_withholding_tax model.
type XWithholdingTaxs []XWithholdingTax

// XWithholdingTaxModel is the odoo model name.
const XWithholdingTaxModel = "x_withholding_tax"

// Many2One convert XWithholdingTax to *Many2One.
func (x *XWithholdingTax) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXWithholdingTax creates a new x_withholding_tax model and returns its id.
func (c *Client) CreateXWithholdingTax(x *XWithholdingTax) (int64, error) {
	ids, err := c.CreateXWithholdingTaxs([]*XWithholdingTax{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXWithholdingTax creates a new x_withholding_tax model and returns its id.
func (c *Client) CreateXWithholdingTaxs(xs []*XWithholdingTax) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XWithholdingTaxModel, vv)
}

// UpdateXWithholdingTax updates an existing x_withholding_tax record.
func (c *Client) UpdateXWithholdingTax(x *XWithholdingTax) error {
	return c.UpdateXWithholdingTaxs([]int64{x.Id.Get()}, x)
}

// UpdateXWithholdingTaxs updates existing x_withholding_tax records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXWithholdingTaxs(ids []int64, x *XWithholdingTax) error {
	return c.Update(XWithholdingTaxModel, ids, x)
}

// DeleteXWithholdingTax deletes an existing x_withholding_tax record.
func (c *Client) DeleteXWithholdingTax(id int64) error {
	return c.DeleteXWithholdingTaxs([]int64{id})
}

// DeleteXWithholdingTaxs deletes existing x_withholding_tax records.
func (c *Client) DeleteXWithholdingTaxs(ids []int64) error {
	return c.Delete(XWithholdingTaxModel, ids)
}

// GetXWithholdingTax gets x_withholding_tax existing record.
func (c *Client) GetXWithholdingTax(id int64) (*XWithholdingTax, error) {
	xs, err := c.GetXWithholdingTaxs([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_withholding_tax not found", id)
}

// GetXWithholdingTaxs gets x_withholding_tax existing records.
func (c *Client) GetXWithholdingTaxs(ids []int64) (*XWithholdingTaxs, error) {
	xs := &XWithholdingTaxs{}
	if err := c.Read(XWithholdingTaxModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTax finds x_withholding_tax record by querying it with criteria.
func (c *Client) FindXWithholdingTax(criteria *Criteria) (*XWithholdingTax, error) {
	xs := &XWithholdingTaxs{}
	if err := c.SearchRead(XWithholdingTaxModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_withholding_tax was not found with criteria %v", criteria)
}

// FindXWithholdingTaxs finds x_withholding_tax records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxs(criteria *Criteria, options *Options) (*XWithholdingTaxs, error) {
	xs := &XWithholdingTaxs{}
	if err := c.SearchRead(XWithholdingTaxModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTaxIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XWithholdingTaxModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXWithholdingTaxId finds record id by querying it with criteria.
func (c *Client) FindXWithholdingTaxId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XWithholdingTaxModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_withholding_tax was not found with criteria %v and options %v", criteria, options)
}
