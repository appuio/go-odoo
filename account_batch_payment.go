package odoo

import (
	"fmt"
)

// AccountBatchPayment represents account.batch.payment model.
type AccountBatchPayment struct {
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
	Amount                      *Float     `xmlrpc:"amount,omptempty"`
	AmountResidual              *Float     `xmlrpc:"amount_residual,omptempty"`
	AmountResidualCurrency      *Float     `xmlrpc:"amount_residual_currency,omptempty"`
	AvailablePaymentMethodIds   *Relation  `xmlrpc:"available_payment_method_ids,omptempty"`
	BatchType                   *Selection `xmlrpc:"batch_type,omptempty"`
	CompanyCurrencyId           *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                        *Time      `xmlrpc:"date,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	ExportFile                  *String    `xmlrpc:"export_file,omptempty"`
	ExportFileCreateDate        *Time      `xmlrpc:"export_file_create_date,omptempty"`
	ExportFilename              *String    `xmlrpc:"export_filename,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FileGenerationEnabled       *Bool      `xmlrpc:"file_generation_enabled,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	JournalId                   *Many2One  `xmlrpc:"journal_id,omptempty"`
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
	PaymentIds                  *Relation  `xmlrpc:"payment_ids,omptempty"`
	PaymentMethodCode           *String    `xmlrpc:"payment_method_code,omptempty"`
	PaymentMethodId             *Many2One  `xmlrpc:"payment_method_id,omptempty"`
	SctBatchBooking             *Bool      `xmlrpc:"sct_batch_booking,omptempty"`
	SctGeneric                  *Bool      `xmlrpc:"sct_generic,omptempty"`
	SddBatchBooking             *Bool      `xmlrpc:"sdd_batch_booking,omptempty"`
	SddRequiredCollectionDate   *Time      `xmlrpc:"sdd_required_collection_date,omptempty"`
	SddScheme                   *Selection `xmlrpc:"sdd_scheme,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountBatchPayments represents array of account.batch.payment model.
type AccountBatchPayments []AccountBatchPayment

// AccountBatchPaymentModel is the odoo model name.
const AccountBatchPaymentModel = "account.batch.payment"

// Many2One convert AccountBatchPayment to *Many2One.
func (abp *AccountBatchPayment) Many2One() *Many2One {
	return NewMany2One(abp.Id.Get(), "")
}

// CreateAccountBatchPayment creates a new account.batch.payment model and returns its id.
func (c *Client) CreateAccountBatchPayment(abp *AccountBatchPayment) (int64, error) {
	ids, err := c.CreateAccountBatchPayments([]*AccountBatchPayment{abp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchPayment creates a new account.batch.payment model and returns its id.
func (c *Client) CreateAccountBatchPayments(abps []*AccountBatchPayment) ([]int64, error) {
	var vv []interface{}
	for _, v := range abps {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchPaymentModel, vv)
}

// UpdateAccountBatchPayment updates an existing account.batch.payment record.
func (c *Client) UpdateAccountBatchPayment(abp *AccountBatchPayment) error {
	return c.UpdateAccountBatchPayments([]int64{abp.Id.Get()}, abp)
}

// UpdateAccountBatchPayments updates existing account.batch.payment records.
// All records (represented by ids) will be updated by abp values.
func (c *Client) UpdateAccountBatchPayments(ids []int64, abp *AccountBatchPayment) error {
	return c.Update(AccountBatchPaymentModel, ids, abp)
}

// DeleteAccountBatchPayment deletes an existing account.batch.payment record.
func (c *Client) DeleteAccountBatchPayment(id int64) error {
	return c.DeleteAccountBatchPayments([]int64{id})
}

// DeleteAccountBatchPayments deletes existing account.batch.payment records.
func (c *Client) DeleteAccountBatchPayments(ids []int64) error {
	return c.Delete(AccountBatchPaymentModel, ids)
}

// GetAccountBatchPayment gets account.batch.payment existing record.
func (c *Client) GetAccountBatchPayment(id int64) (*AccountBatchPayment, error) {
	abps, err := c.GetAccountBatchPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	if abps != nil && len(*abps) > 0 {
		return &((*abps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.batch.payment not found", id)
}

// GetAccountBatchPayments gets account.batch.payment existing records.
func (c *Client) GetAccountBatchPayments(ids []int64) (*AccountBatchPayments, error) {
	abps := &AccountBatchPayments{}
	if err := c.Read(AccountBatchPaymentModel, ids, nil, abps); err != nil {
		return nil, err
	}
	return abps, nil
}

// FindAccountBatchPayment finds account.batch.payment record by querying it with criteria.
func (c *Client) FindAccountBatchPayment(criteria *Criteria) (*AccountBatchPayment, error) {
	abps := &AccountBatchPayments{}
	if err := c.SearchRead(AccountBatchPaymentModel, criteria, NewOptions().Limit(1), abps); err != nil {
		return nil, err
	}
	if abps != nil && len(*abps) > 0 {
		return &((*abps)[0]), nil
	}
	return nil, fmt.Errorf("account.batch.payment was not found with criteria %v", criteria)
}

// FindAccountBatchPayments finds account.batch.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPayments(criteria *Criteria, options *Options) (*AccountBatchPayments, error) {
	abps := &AccountBatchPayments{}
	if err := c.SearchRead(AccountBatchPaymentModel, criteria, options, abps); err != nil {
		return nil, err
	}
	return abps, nil
}

// FindAccountBatchPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBatchPaymentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBatchPaymentId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.batch.payment was not found with criteria %v and options %v", criteria, options)
}
