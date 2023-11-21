package odoo

import (
	"fmt"
)

// XWithholdingTaxConf represents x_withholding_tax_conf model.
type XWithholdingTaxConf struct {
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
}

// XWithholdingTaxConfs represents array of x_withholding_tax_conf model.
type XWithholdingTaxConfs []XWithholdingTaxConf

// XWithholdingTaxConfModel is the odoo model name.
const XWithholdingTaxConfModel = "x_withholding_tax_conf"

// Many2One convert XWithholdingTaxConf to *Many2One.
func (x *XWithholdingTaxConf) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXWithholdingTaxConf creates a new x_withholding_tax_conf model and returns its id.
func (c *Client) CreateXWithholdingTaxConf(x *XWithholdingTaxConf) (int64, error) {
	ids, err := c.CreateXWithholdingTaxConfs([]*XWithholdingTaxConf{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXWithholdingTaxConf creates a new x_withholding_tax_conf model and returns its id.
func (c *Client) CreateXWithholdingTaxConfs(xs []*XWithholdingTaxConf) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XWithholdingTaxConfModel, vv)
}

// UpdateXWithholdingTaxConf updates an existing x_withholding_tax_conf record.
func (c *Client) UpdateXWithholdingTaxConf(x *XWithholdingTaxConf) error {
	return c.UpdateXWithholdingTaxConfs([]int64{x.Id.Get()}, x)
}

// UpdateXWithholdingTaxConfs updates existing x_withholding_tax_conf records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXWithholdingTaxConfs(ids []int64, x *XWithholdingTaxConf) error {
	return c.Update(XWithholdingTaxConfModel, ids, x)
}

// DeleteXWithholdingTaxConf deletes an existing x_withholding_tax_conf record.
func (c *Client) DeleteXWithholdingTaxConf(id int64) error {
	return c.DeleteXWithholdingTaxConfs([]int64{id})
}

// DeleteXWithholdingTaxConfs deletes existing x_withholding_tax_conf records.
func (c *Client) DeleteXWithholdingTaxConfs(ids []int64) error {
	return c.Delete(XWithholdingTaxConfModel, ids)
}

// GetXWithholdingTaxConf gets x_withholding_tax_conf existing record.
func (c *Client) GetXWithholdingTaxConf(id int64) (*XWithholdingTaxConf, error) {
	xs, err := c.GetXWithholdingTaxConfs([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_withholding_tax_conf not found", id)
}

// GetXWithholdingTaxConfs gets x_withholding_tax_conf existing records.
func (c *Client) GetXWithholdingTaxConfs(ids []int64) (*XWithholdingTaxConfs, error) {
	xs := &XWithholdingTaxConfs{}
	if err := c.Read(XWithholdingTaxConfModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTaxConf finds x_withholding_tax_conf record by querying it with criteria.
func (c *Client) FindXWithholdingTaxConf(criteria *Criteria) (*XWithholdingTaxConf, error) {
	xs := &XWithholdingTaxConfs{}
	if err := c.SearchRead(XWithholdingTaxConfModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_withholding_tax_conf was not found with criteria %v", criteria)
}

// FindXWithholdingTaxConfs finds x_withholding_tax_conf records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxConfs(criteria *Criteria, options *Options) (*XWithholdingTaxConfs, error) {
	xs := &XWithholdingTaxConfs{}
	if err := c.SearchRead(XWithholdingTaxConfModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTaxConfIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxConfIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XWithholdingTaxConfModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXWithholdingTaxConfId finds record id by querying it with criteria.
func (c *Client) FindXWithholdingTaxConfId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XWithholdingTaxConfModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_withholding_tax_conf was not found with criteria %v and options %v", criteria, options)
}
