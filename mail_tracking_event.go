package odoo

import (
	"fmt"
)

// MailTrackingEvent represents mail.tracking.event model.
type MailTrackingEvent struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date             *Time      `xmlrpc:"date,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	ErrorDescription *String    `xmlrpc:"error_description,omptempty"`
	ErrorDetails     *String    `xmlrpc:"error_details,omptempty"`
	ErrorType        *String    `xmlrpc:"error_type,omptempty"`
	EventType        *Selection `xmlrpc:"event_type,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	Ip               *String    `xmlrpc:"ip,omptempty"`
	MailgunId        *String    `xmlrpc:"mailgun_id,omptempty"`
	MassMailingId    *Many2One  `xmlrpc:"mass_mailing_id,omptempty"`
	Mobile           *Bool      `xmlrpc:"mobile,omptempty"`
	OsFamily         *String    `xmlrpc:"os_family,omptempty"`
	Recipient        *String    `xmlrpc:"recipient,omptempty"`
	RecipientAddress *String    `xmlrpc:"recipient_address,omptempty"`
	SmtpServer       *String    `xmlrpc:"smtp_server,omptempty"`
	Time             *Time      `xmlrpc:"time,omptempty"`
	Timestamp        *Float     `xmlrpc:"timestamp,omptempty"`
	TrackingEmailId  *Many2One  `xmlrpc:"tracking_email_id,omptempty"`
	UaFamily         *String    `xmlrpc:"ua_family,omptempty"`
	UaType           *String    `xmlrpc:"ua_type,omptempty"`
	Url              *String    `xmlrpc:"url,omptempty"`
	UserAgent        *String    `xmlrpc:"user_agent,omptempty"`
	UserCountryId    *Many2One  `xmlrpc:"user_country_id,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailTrackingEvents represents array of mail.tracking.event model.
type MailTrackingEvents []MailTrackingEvent

// MailTrackingEventModel is the odoo model name.
const MailTrackingEventModel = "mail.tracking.event"

// Many2One convert MailTrackingEvent to *Many2One.
func (mte *MailTrackingEvent) Many2One() *Many2One {
	return NewMany2One(mte.Id.Get(), "")
}

// CreateMailTrackingEvent creates a new mail.tracking.event model and returns its id.
func (c *Client) CreateMailTrackingEvent(mte *MailTrackingEvent) (int64, error) {
	ids, err := c.CreateMailTrackingEvents([]*MailTrackingEvent{mte})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailTrackingEvent creates a new mail.tracking.event model and returns its id.
func (c *Client) CreateMailTrackingEvents(mtes []*MailTrackingEvent) ([]int64, error) {
	var vv []interface{}
	for _, v := range mtes {
		vv = append(vv, v)
	}
	return c.Create(MailTrackingEventModel, vv)
}

// UpdateMailTrackingEvent updates an existing mail.tracking.event record.
func (c *Client) UpdateMailTrackingEvent(mte *MailTrackingEvent) error {
	return c.UpdateMailTrackingEvents([]int64{mte.Id.Get()}, mte)
}

// UpdateMailTrackingEvents updates existing mail.tracking.event records.
// All records (represented by ids) will be updated by mte values.
func (c *Client) UpdateMailTrackingEvents(ids []int64, mte *MailTrackingEvent) error {
	return c.Update(MailTrackingEventModel, ids, mte)
}

// DeleteMailTrackingEvent deletes an existing mail.tracking.event record.
func (c *Client) DeleteMailTrackingEvent(id int64) error {
	return c.DeleteMailTrackingEvents([]int64{id})
}

// DeleteMailTrackingEvents deletes existing mail.tracking.event records.
func (c *Client) DeleteMailTrackingEvents(ids []int64) error {
	return c.Delete(MailTrackingEventModel, ids)
}

// GetMailTrackingEvent gets mail.tracking.event existing record.
func (c *Client) GetMailTrackingEvent(id int64) (*MailTrackingEvent, error) {
	mtes, err := c.GetMailTrackingEvents([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtes != nil && len(*mtes) > 0 {
		return &((*mtes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.tracking.event not found", id)
}

// GetMailTrackingEvents gets mail.tracking.event existing records.
func (c *Client) GetMailTrackingEvents(ids []int64) (*MailTrackingEvents, error) {
	mtes := &MailTrackingEvents{}
	if err := c.Read(MailTrackingEventModel, ids, nil, mtes); err != nil {
		return nil, err
	}
	return mtes, nil
}

// FindMailTrackingEvent finds mail.tracking.event record by querying it with criteria.
func (c *Client) FindMailTrackingEvent(criteria *Criteria) (*MailTrackingEvent, error) {
	mtes := &MailTrackingEvents{}
	if err := c.SearchRead(MailTrackingEventModel, criteria, NewOptions().Limit(1), mtes); err != nil {
		return nil, err
	}
	if mtes != nil && len(*mtes) > 0 {
		return &((*mtes)[0]), nil
	}
	return nil, fmt.Errorf("mail.tracking.event was not found with criteria %v", criteria)
}

// FindMailTrackingEvents finds mail.tracking.event records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingEvents(criteria *Criteria, options *Options) (*MailTrackingEvents, error) {
	mtes := &MailTrackingEvents{}
	if err := c.SearchRead(MailTrackingEventModel, criteria, options, mtes); err != nil {
		return nil, err
	}
	return mtes, nil
}

// FindMailTrackingEventIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingEventIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailTrackingEventModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTrackingEventId finds record id by querying it with criteria.
func (c *Client) FindMailTrackingEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTrackingEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.tracking.event was not found with criteria %v and options %v", criteria, options)
}
