package odoo

import (
	"fmt"
)

// SddMandate represents sdd.mandate model.
type SddMandate struct {
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
	DebtorIdCode                *String    `xmlrpc:"debtor_id_code,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EndDate                     *Time      `xmlrpc:"end_date,omptempty"`
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
	OneOff                      *Bool      `xmlrpc:"one_off,omptempty"`
	PaidInvoiceIds              *Relation  `xmlrpc:"paid_invoice_ids,omptempty"`
	PaidInvoicesNber            *Int       `xmlrpc:"paid_invoices_nber,omptempty"`
	PartnerBankId               *Many2One  `xmlrpc:"partner_bank_id,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentIds                  *Relation  `xmlrpc:"payment_ids,omptempty"`
	PaymentJournalId            *Many2One  `xmlrpc:"payment_journal_id,omptempty"`
	PaymentsToCollectNber       *Int       `xmlrpc:"payments_to_collect_nber,omptempty"`
	PhoneNumber                 *String    `xmlrpc:"phone_number,omptempty"`
	SddScheme                   *Selection `xmlrpc:"sdd_scheme,omptempty"`
	Signature                   *String    `xmlrpc:"signature,omptempty"`
	SignedBy                    *String    `xmlrpc:"signed_by,omptempty"`
	SignedOn                    *Time      `xmlrpc:"signed_on,omptempty"`
	StartDate                   *Time      `xmlrpc:"start_date,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	SuitableJournalIds          *Relation  `xmlrpc:"suitable_journal_ids,omptempty"`
	VerificationCode            *String    `xmlrpc:"verification_code,omptempty"`
	Verified                    *Bool      `xmlrpc:"verified,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SddMandates represents array of sdd.mandate model.
type SddMandates []SddMandate

// SddMandateModel is the odoo model name.
const SddMandateModel = "sdd.mandate"

// Many2One convert SddMandate to *Many2One.
func (sm *SddMandate) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateSddMandate creates a new sdd.mandate model and returns its id.
func (c *Client) CreateSddMandate(sm *SddMandate) (int64, error) {
	ids, err := c.CreateSddMandates([]*SddMandate{sm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSddMandate creates a new sdd.mandate model and returns its id.
func (c *Client) CreateSddMandates(sms []*SddMandate) ([]int64, error) {
	var vv []interface{}
	for _, v := range sms {
		vv = append(vv, v)
	}
	return c.Create(SddMandateModel, vv)
}

// UpdateSddMandate updates an existing sdd.mandate record.
func (c *Client) UpdateSddMandate(sm *SddMandate) error {
	return c.UpdateSddMandates([]int64{sm.Id.Get()}, sm)
}

// UpdateSddMandates updates existing sdd.mandate records.
// All records (represented by ids) will be updated by sm values.
func (c *Client) UpdateSddMandates(ids []int64, sm *SddMandate) error {
	return c.Update(SddMandateModel, ids, sm)
}

// DeleteSddMandate deletes an existing sdd.mandate record.
func (c *Client) DeleteSddMandate(id int64) error {
	return c.DeleteSddMandates([]int64{id})
}

// DeleteSddMandates deletes existing sdd.mandate records.
func (c *Client) DeleteSddMandates(ids []int64) error {
	return c.Delete(SddMandateModel, ids)
}

// GetSddMandate gets sdd.mandate existing record.
func (c *Client) GetSddMandate(id int64) (*SddMandate, error) {
	sms, err := c.GetSddMandates([]int64{id})
	if err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sdd.mandate not found", id)
}

// GetSddMandates gets sdd.mandate existing records.
func (c *Client) GetSddMandates(ids []int64) (*SddMandates, error) {
	sms := &SddMandates{}
	if err := c.Read(SddMandateModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSddMandate finds sdd.mandate record by querying it with criteria.
func (c *Client) FindSddMandate(criteria *Criteria) (*SddMandate, error) {
	sms := &SddMandates{}
	if err := c.SearchRead(SddMandateModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("sdd.mandate was not found with criteria %v", criteria)
}

// FindSddMandates finds sdd.mandate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSddMandates(criteria *Criteria, options *Options) (*SddMandates, error) {
	sms := &SddMandates{}
	if err := c.SearchRead(SddMandateModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindSddMandateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSddMandateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SddMandateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSddMandateId finds record id by querying it with criteria.
func (c *Client) FindSddMandateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SddMandateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sdd.mandate was not found with criteria %v and options %v", criteria, options)
}
