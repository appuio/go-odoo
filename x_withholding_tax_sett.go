package odoo

import (
	"fmt"
)

// XWithholdingTaxSett represents x_withholding_tax_sett model.
type XWithholdingTaxSett struct {
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
	XStudioCanton               *String    `xmlrpc:"x_studio_canton,omptempty"`
	XStudioCurrency             *Many2One  `xmlrpc:"x_studio_currency,omptempty"`
	XStudioMany2ManyFieldX8RJx  *Relation  `xmlrpc:"x_studio_many2many_field_X8RJx,omptempty"`
	XStudioQstCompany           *Many2One  `xmlrpc:"x_studio_qst_company,omptempty"`
	XStudioSequence             *Int       `xmlrpc:"x_studio_sequence,omptempty"`
	XStudioWithholdingTaxList   *Relation  `xmlrpc:"x_studio_withholding_tax_list,omptempty"`
	XStudioYear                 *Int       `xmlrpc:"x_studio_year,omptempty"`
	XStudioYearWithoutSeparator *String    `xmlrpc:"x_studio_year_without_separator,omptempty"`
}

// XWithholdingTaxSetts represents array of x_withholding_tax_sett model.
type XWithholdingTaxSetts []XWithholdingTaxSett

// XWithholdingTaxSettModel is the odoo model name.
const XWithholdingTaxSettModel = "x_withholding_tax_sett"

// Many2One convert XWithholdingTaxSett to *Many2One.
func (x *XWithholdingTaxSett) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXWithholdingTaxSett creates a new x_withholding_tax_sett model and returns its id.
func (c *Client) CreateXWithholdingTaxSett(x *XWithholdingTaxSett) (int64, error) {
	ids, err := c.CreateXWithholdingTaxSetts([]*XWithholdingTaxSett{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXWithholdingTaxSett creates a new x_withholding_tax_sett model and returns its id.
func (c *Client) CreateXWithholdingTaxSetts(xs []*XWithholdingTaxSett) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XWithholdingTaxSettModel, vv)
}

// UpdateXWithholdingTaxSett updates an existing x_withholding_tax_sett record.
func (c *Client) UpdateXWithholdingTaxSett(x *XWithholdingTaxSett) error {
	return c.UpdateXWithholdingTaxSetts([]int64{x.Id.Get()}, x)
}

// UpdateXWithholdingTaxSetts updates existing x_withholding_tax_sett records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXWithholdingTaxSetts(ids []int64, x *XWithholdingTaxSett) error {
	return c.Update(XWithholdingTaxSettModel, ids, x)
}

// DeleteXWithholdingTaxSett deletes an existing x_withholding_tax_sett record.
func (c *Client) DeleteXWithholdingTaxSett(id int64) error {
	return c.DeleteXWithholdingTaxSetts([]int64{id})
}

// DeleteXWithholdingTaxSetts deletes existing x_withholding_tax_sett records.
func (c *Client) DeleteXWithholdingTaxSetts(ids []int64) error {
	return c.Delete(XWithholdingTaxSettModel, ids)
}

// GetXWithholdingTaxSett gets x_withholding_tax_sett existing record.
func (c *Client) GetXWithholdingTaxSett(id int64) (*XWithholdingTaxSett, error) {
	xs, err := c.GetXWithholdingTaxSetts([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_withholding_tax_sett not found", id)
}

// GetXWithholdingTaxSetts gets x_withholding_tax_sett existing records.
func (c *Client) GetXWithholdingTaxSetts(ids []int64) (*XWithholdingTaxSetts, error) {
	xs := &XWithholdingTaxSetts{}
	if err := c.Read(XWithholdingTaxSettModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTaxSett finds x_withholding_tax_sett record by querying it with criteria.
func (c *Client) FindXWithholdingTaxSett(criteria *Criteria) (*XWithholdingTaxSett, error) {
	xs := &XWithholdingTaxSetts{}
	if err := c.SearchRead(XWithholdingTaxSettModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_withholding_tax_sett was not found with criteria %v", criteria)
}

// FindXWithholdingTaxSetts finds x_withholding_tax_sett records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxSetts(criteria *Criteria, options *Options) (*XWithholdingTaxSetts, error) {
	xs := &XWithholdingTaxSetts{}
	if err := c.SearchRead(XWithholdingTaxSettModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTaxSettIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxSettIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XWithholdingTaxSettModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXWithholdingTaxSettId finds record id by querying it with criteria.
func (c *Client) FindXWithholdingTaxSettId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XWithholdingTaxSettModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_withholding_tax_sett was not found with criteria %v and options %v", criteria, options)
}
