package odoo

import (
	"fmt"
)

// AccountFollowupManualReminder represents account_followup.manual_reminder model.
type AccountFollowupManualReminder struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	AttachmentIds        *Relation `xmlrpc:"attachment_ids,omptempty"`
	Body                 *String   `xmlrpc:"body,omptempty"`
	BodyHtml             *String   `xmlrpc:"body_html,omptempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	Email                *Bool     `xmlrpc:"email,omptempty"`
	EmailAddSignature    *Bool     `xmlrpc:"email_add_signature,omptempty"`
	EmailRecipientIds    *Relation `xmlrpc:"email_recipient_ids,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omptempty"`
	JoinInvoices         *Bool     `xmlrpc:"join_invoices,omptempty"`
	Lang                 *String   `xmlrpc:"lang,omptempty"`
	PartnerId            *Many2One `xmlrpc:"partner_id,omptempty"`
	Print                *Bool     `xmlrpc:"print,omptempty"`
	RenderModel          *String   `xmlrpc:"render_model,omptempty"`
	Sms                  *Bool     `xmlrpc:"sms,omptempty"`
	SmsBody              *String   `xmlrpc:"sms_body,omptempty"`
	SmsTemplateId        *Many2One `xmlrpc:"sms_template_id,omptempty"`
	Snailmail            *Bool     `xmlrpc:"snailmail,omptempty"`
	SnailmailCost        *Float    `xmlrpc:"snailmail_cost,omptempty"`
	Subject              *String   `xmlrpc:"subject,omptempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountFollowupManualReminders represents array of account_followup.manual_reminder model.
type AccountFollowupManualReminders []AccountFollowupManualReminder

// AccountFollowupManualReminderModel is the odoo model name.
const AccountFollowupManualReminderModel = "account_followup.manual_reminder"

// Many2One convert AccountFollowupManualReminder to *Many2One.
func (am *AccountFollowupManualReminder) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAccountFollowupManualReminder creates a new account_followup.manual_reminder model and returns its id.
func (c *Client) CreateAccountFollowupManualReminder(am *AccountFollowupManualReminder) (int64, error) {
	ids, err := c.CreateAccountFollowupManualReminders([]*AccountFollowupManualReminder{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFollowupManualReminder creates a new account_followup.manual_reminder model and returns its id.
func (c *Client) CreateAccountFollowupManualReminders(ams []*AccountFollowupManualReminder) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AccountFollowupManualReminderModel, vv)
}

// UpdateAccountFollowupManualReminder updates an existing account_followup.manual_reminder record.
func (c *Client) UpdateAccountFollowupManualReminder(am *AccountFollowupManualReminder) error {
	return c.UpdateAccountFollowupManualReminders([]int64{am.Id.Get()}, am)
}

// UpdateAccountFollowupManualReminders updates existing account_followup.manual_reminder records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAccountFollowupManualReminders(ids []int64, am *AccountFollowupManualReminder) error {
	return c.Update(AccountFollowupManualReminderModel, ids, am)
}

// DeleteAccountFollowupManualReminder deletes an existing account_followup.manual_reminder record.
func (c *Client) DeleteAccountFollowupManualReminder(id int64) error {
	return c.DeleteAccountFollowupManualReminders([]int64{id})
}

// DeleteAccountFollowupManualReminders deletes existing account_followup.manual_reminder records.
func (c *Client) DeleteAccountFollowupManualReminders(ids []int64) error {
	return c.Delete(AccountFollowupManualReminderModel, ids)
}

// GetAccountFollowupManualReminder gets account_followup.manual_reminder existing record.
func (c *Client) GetAccountFollowupManualReminder(id int64) (*AccountFollowupManualReminder, error) {
	ams, err := c.GetAccountFollowupManualReminders([]int64{id})
	if err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account_followup.manual_reminder not found", id)
}

// GetAccountFollowupManualReminders gets account_followup.manual_reminder existing records.
func (c *Client) GetAccountFollowupManualReminders(ids []int64) (*AccountFollowupManualReminders, error) {
	ams := &AccountFollowupManualReminders{}
	if err := c.Read(AccountFollowupManualReminderModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountFollowupManualReminder finds account_followup.manual_reminder record by querying it with criteria.
func (c *Client) FindAccountFollowupManualReminder(criteria *Criteria) (*AccountFollowupManualReminder, error) {
	ams := &AccountFollowupManualReminders{}
	if err := c.SearchRead(AccountFollowupManualReminderModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("account_followup.manual_reminder was not found with criteria %v", criteria)
}

// FindAccountFollowupManualReminders finds account_followup.manual_reminder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupManualReminders(criteria *Criteria, options *Options) (*AccountFollowupManualReminders, error) {
	ams := &AccountFollowupManualReminders{}
	if err := c.SearchRead(AccountFollowupManualReminderModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAccountFollowupManualReminderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupManualReminderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFollowupManualReminderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFollowupManualReminderId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupManualReminderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupManualReminderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account_followup.manual_reminder was not found with criteria %v and options %v", criteria, options)
}
