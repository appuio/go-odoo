package odoo

import (
	"fmt"
)

// MailTrackingEmail represents mail.tracking.email model.
type MailTrackingEmail struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	BounceDescription *String    `xmlrpc:"bounce_description,omptempty"`
	BounceType        *String    `xmlrpc:"bounce_type,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date              *Time      `xmlrpc:"date,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	ErrorDescription  *String    `xmlrpc:"error_description,omptempty"`
	ErrorSmtpServer   *String    `xmlrpc:"error_smtp_server,omptempty"`
	ErrorType         *String    `xmlrpc:"error_type,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	MailId            *Many2One  `xmlrpc:"mail_id,omptempty"`
	MailIdInt         *Int       `xmlrpc:"mail_id_int,omptempty"`
	MailMessageId     *Many2One  `xmlrpc:"mail_message_id,omptempty"`
	MailStatsId       *Many2One  `xmlrpc:"mail_stats_id,omptempty"`
	MassMailingId     *Many2One  `xmlrpc:"mass_mailing_id,omptempty"`
	MessageId         *String    `xmlrpc:"message_id,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omptempty"`
	Recipient         *String    `xmlrpc:"recipient,omptempty"`
	RecipientAddress  *String    `xmlrpc:"recipient_address,omptempty"`
	Sender            *String    `xmlrpc:"sender,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	Time              *Time      `xmlrpc:"time,omptempty"`
	Timestamp         *Float     `xmlrpc:"timestamp,omptempty"`
	Token             *String    `xmlrpc:"token,omptempty"`
	TrackingEventIds  *Relation  `xmlrpc:"tracking_event_ids,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailTrackingEmails represents array of mail.tracking.email model.
type MailTrackingEmails []MailTrackingEmail

// MailTrackingEmailModel is the odoo model name.
const MailTrackingEmailModel = "mail.tracking.email"

// Many2One convert MailTrackingEmail to *Many2One.
func (mte *MailTrackingEmail) Many2One() *Many2One {
	return NewMany2One(mte.Id.Get(), "")
}

// CreateMailTrackingEmail creates a new mail.tracking.email model and returns its id.
func (c *Client) CreateMailTrackingEmail(mte *MailTrackingEmail) (int64, error) {
	ids, err := c.CreateMailTrackingEmails([]*MailTrackingEmail{mte})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailTrackingEmail creates a new mail.tracking.email model and returns its id.
func (c *Client) CreateMailTrackingEmails(mtes []*MailTrackingEmail) ([]int64, error) {
	var vv []interface{}
	for _, v := range mtes {
		vv = append(vv, v)
	}
	return c.Create(MailTrackingEmailModel, vv)
}

// UpdateMailTrackingEmail updates an existing mail.tracking.email record.
func (c *Client) UpdateMailTrackingEmail(mte *MailTrackingEmail) error {
	return c.UpdateMailTrackingEmails([]int64{mte.Id.Get()}, mte)
}

// UpdateMailTrackingEmails updates existing mail.tracking.email records.
// All records (represented by ids) will be updated by mte values.
func (c *Client) UpdateMailTrackingEmails(ids []int64, mte *MailTrackingEmail) error {
	return c.Update(MailTrackingEmailModel, ids, mte)
}

// DeleteMailTrackingEmail deletes an existing mail.tracking.email record.
func (c *Client) DeleteMailTrackingEmail(id int64) error {
	return c.DeleteMailTrackingEmails([]int64{id})
}

// DeleteMailTrackingEmails deletes existing mail.tracking.email records.
func (c *Client) DeleteMailTrackingEmails(ids []int64) error {
	return c.Delete(MailTrackingEmailModel, ids)
}

// GetMailTrackingEmail gets mail.tracking.email existing record.
func (c *Client) GetMailTrackingEmail(id int64) (*MailTrackingEmail, error) {
	mtes, err := c.GetMailTrackingEmails([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtes != nil && len(*mtes) > 0 {
		return &((*mtes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.tracking.email not found", id)
}

// GetMailTrackingEmails gets mail.tracking.email existing records.
func (c *Client) GetMailTrackingEmails(ids []int64) (*MailTrackingEmails, error) {
	mtes := &MailTrackingEmails{}
	if err := c.Read(MailTrackingEmailModel, ids, nil, mtes); err != nil {
		return nil, err
	}
	return mtes, nil
}

// FindMailTrackingEmail finds mail.tracking.email record by querying it with criteria.
func (c *Client) FindMailTrackingEmail(criteria *Criteria) (*MailTrackingEmail, error) {
	mtes := &MailTrackingEmails{}
	if err := c.SearchRead(MailTrackingEmailModel, criteria, NewOptions().Limit(1), mtes); err != nil {
		return nil, err
	}
	if mtes != nil && len(*mtes) > 0 {
		return &((*mtes)[0]), nil
	}
	return nil, fmt.Errorf("mail.tracking.email was not found with criteria %v", criteria)
}

// FindMailTrackingEmails finds mail.tracking.email records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingEmails(criteria *Criteria, options *Options) (*MailTrackingEmails, error) {
	mtes := &MailTrackingEmails{}
	if err := c.SearchRead(MailTrackingEmailModel, criteria, options, mtes); err != nil {
		return nil, err
	}
	return mtes, nil
}

// FindMailTrackingEmailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTrackingEmailIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailTrackingEmailModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTrackingEmailId finds record id by querying it with criteria.
func (c *Client) FindMailTrackingEmailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTrackingEmailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.tracking.email was not found with criteria %v and options %v", criteria, options)
}
