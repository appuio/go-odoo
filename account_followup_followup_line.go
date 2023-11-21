package odoo

import (
	"fmt"
)

// AccountFollowupFollowupLine represents account_followup.followup.line model.
type AccountFollowupFollowupLine struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDefaultResponsibleType *Selection `xmlrpc:"activity_default_responsible_type,omptempty"`
	ActivityNote                   *String    `xmlrpc:"activity_note,omptempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	AutoExecute                    *Bool      `xmlrpc:"auto_execute,omptempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateActivity                 *Bool      `xmlrpc:"create_activity,omptempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omptempty"`
	Delay                          *Int       `xmlrpc:"delay,omptempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omptempty"`
	Id                             *Int       `xmlrpc:"id,omptempty"`
	JoinInvoices                   *Bool      `xmlrpc:"join_invoices,omptempty"`
	MailTemplateId                 *Many2One  `xmlrpc:"mail_template_id,omptempty"`
	Name                           *String    `xmlrpc:"name,omptempty"`
	SendEmail                      *Bool      `xmlrpc:"send_email,omptempty"`
	SendLetter                     *Bool      `xmlrpc:"send_letter,omptempty"`
	SendSms                        *Bool      `xmlrpc:"send_sms,omptempty"`
	SmsTemplateId                  *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountFollowupFollowupLines represents array of account_followup.followup.line model.
type AccountFollowupFollowupLines []AccountFollowupFollowupLine

// AccountFollowupFollowupLineModel is the odoo model name.
const AccountFollowupFollowupLineModel = "account_followup.followup.line"

// Many2One convert AccountFollowupFollowupLine to *Many2One.
func (afl *AccountFollowupFollowupLine) Many2One() *Many2One {
	return NewMany2One(afl.Id.Get(), "")
}

// CreateAccountFollowupFollowupLine creates a new account_followup.followup.line model and returns its id.
func (c *Client) CreateAccountFollowupFollowupLine(afl *AccountFollowupFollowupLine) (int64, error) {
	ids, err := c.CreateAccountFollowupFollowupLines([]*AccountFollowupFollowupLine{afl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountFollowupFollowupLine creates a new account_followup.followup.line model and returns its id.
func (c *Client) CreateAccountFollowupFollowupLines(afls []*AccountFollowupFollowupLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range afls {
		vv = append(vv, v)
	}
	return c.Create(AccountFollowupFollowupLineModel, vv)
}

// UpdateAccountFollowupFollowupLine updates an existing account_followup.followup.line record.
func (c *Client) UpdateAccountFollowupFollowupLine(afl *AccountFollowupFollowupLine) error {
	return c.UpdateAccountFollowupFollowupLines([]int64{afl.Id.Get()}, afl)
}

// UpdateAccountFollowupFollowupLines updates existing account_followup.followup.line records.
// All records (represented by ids) will be updated by afl values.
func (c *Client) UpdateAccountFollowupFollowupLines(ids []int64, afl *AccountFollowupFollowupLine) error {
	return c.Update(AccountFollowupFollowupLineModel, ids, afl)
}

// DeleteAccountFollowupFollowupLine deletes an existing account_followup.followup.line record.
func (c *Client) DeleteAccountFollowupFollowupLine(id int64) error {
	return c.DeleteAccountFollowupFollowupLines([]int64{id})
}

// DeleteAccountFollowupFollowupLines deletes existing account_followup.followup.line records.
func (c *Client) DeleteAccountFollowupFollowupLines(ids []int64) error {
	return c.Delete(AccountFollowupFollowupLineModel, ids)
}

// GetAccountFollowupFollowupLine gets account_followup.followup.line existing record.
func (c *Client) GetAccountFollowupFollowupLine(id int64) (*AccountFollowupFollowupLine, error) {
	afls, err := c.GetAccountFollowupFollowupLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if afls != nil && len(*afls) > 0 {
		return &((*afls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account_followup.followup.line not found", id)
}

// GetAccountFollowupFollowupLines gets account_followup.followup.line existing records.
func (c *Client) GetAccountFollowupFollowupLines(ids []int64) (*AccountFollowupFollowupLines, error) {
	afls := &AccountFollowupFollowupLines{}
	if err := c.Read(AccountFollowupFollowupLineModel, ids, nil, afls); err != nil {
		return nil, err
	}
	return afls, nil
}

// FindAccountFollowupFollowupLine finds account_followup.followup.line record by querying it with criteria.
func (c *Client) FindAccountFollowupFollowupLine(criteria *Criteria) (*AccountFollowupFollowupLine, error) {
	afls := &AccountFollowupFollowupLines{}
	if err := c.SearchRead(AccountFollowupFollowupLineModel, criteria, NewOptions().Limit(1), afls); err != nil {
		return nil, err
	}
	if afls != nil && len(*afls) > 0 {
		return &((*afls)[0]), nil
	}
	return nil, fmt.Errorf("account_followup.followup.line was not found with criteria %v", criteria)
}

// FindAccountFollowupFollowupLines finds account_followup.followup.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupFollowupLines(criteria *Criteria, options *Options) (*AccountFollowupFollowupLines, error) {
	afls := &AccountFollowupFollowupLines{}
	if err := c.SearchRead(AccountFollowupFollowupLineModel, criteria, options, afls); err != nil {
		return nil, err
	}
	return afls, nil
}

// FindAccountFollowupFollowupLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupFollowupLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFollowupFollowupLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFollowupFollowupLineId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupFollowupLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupFollowupLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account_followup.followup.line was not found with criteria %v and options %v", criteria, options)
}
