package odoo

import (
	"fmt"
)

// AccountOnlineLink represents account.online.link model.
type AccountOnlineLink struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                 *String    `xmlrpc:"access_token,omptempty"`
	AccountOnlineAccountIds     *Relation  `xmlrpc:"account_online_account_ids,omptempty"`
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
	AutoSync                    *Bool      `xmlrpc:"auto_sync,omptempty"`
	ClientId                    *String    `xmlrpc:"client_id,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	ExpiringSynchronizationDate *Time      `xmlrpc:"expiring_synchronization_date,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	LastRefresh                 *Time      `xmlrpc:"last_refresh,omptempty"`
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
	NextRefresh                 *Time      `xmlrpc:"next_refresh,omptempty"`
	ProviderData                *String    `xmlrpc:"provider_data,omptempty"`
	RefreshToken                *String    `xmlrpc:"refresh_token,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountOnlineLinks represents array of account.online.link model.
type AccountOnlineLinks []AccountOnlineLink

// AccountOnlineLinkModel is the odoo model name.
const AccountOnlineLinkModel = "account.online.link"

// Many2One convert AccountOnlineLink to *Many2One.
func (aol *AccountOnlineLink) Many2One() *Many2One {
	return NewMany2One(aol.Id.Get(), "")
}

// CreateAccountOnlineLink creates a new account.online.link model and returns its id.
func (c *Client) CreateAccountOnlineLink(aol *AccountOnlineLink) (int64, error) {
	ids, err := c.CreateAccountOnlineLinks([]*AccountOnlineLink{aol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountOnlineLink creates a new account.online.link model and returns its id.
func (c *Client) CreateAccountOnlineLinks(aols []*AccountOnlineLink) ([]int64, error) {
	var vv []interface{}
	for _, v := range aols {
		vv = append(vv, v)
	}
	return c.Create(AccountOnlineLinkModel, vv)
}

// UpdateAccountOnlineLink updates an existing account.online.link record.
func (c *Client) UpdateAccountOnlineLink(aol *AccountOnlineLink) error {
	return c.UpdateAccountOnlineLinks([]int64{aol.Id.Get()}, aol)
}

// UpdateAccountOnlineLinks updates existing account.online.link records.
// All records (represented by ids) will be updated by aol values.
func (c *Client) UpdateAccountOnlineLinks(ids []int64, aol *AccountOnlineLink) error {
	return c.Update(AccountOnlineLinkModel, ids, aol)
}

// DeleteAccountOnlineLink deletes an existing account.online.link record.
func (c *Client) DeleteAccountOnlineLink(id int64) error {
	return c.DeleteAccountOnlineLinks([]int64{id})
}

// DeleteAccountOnlineLinks deletes existing account.online.link records.
func (c *Client) DeleteAccountOnlineLinks(ids []int64) error {
	return c.Delete(AccountOnlineLinkModel, ids)
}

// GetAccountOnlineLink gets account.online.link existing record.
func (c *Client) GetAccountOnlineLink(id int64) (*AccountOnlineLink, error) {
	aols, err := c.GetAccountOnlineLinks([]int64{id})
	if err != nil {
		return nil, err
	}
	if aols != nil && len(*aols) > 0 {
		return &((*aols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.online.link not found", id)
}

// GetAccountOnlineLinks gets account.online.link existing records.
func (c *Client) GetAccountOnlineLinks(ids []int64) (*AccountOnlineLinks, error) {
	aols := &AccountOnlineLinks{}
	if err := c.Read(AccountOnlineLinkModel, ids, nil, aols); err != nil {
		return nil, err
	}
	return aols, nil
}

// FindAccountOnlineLink finds account.online.link record by querying it with criteria.
func (c *Client) FindAccountOnlineLink(criteria *Criteria) (*AccountOnlineLink, error) {
	aols := &AccountOnlineLinks{}
	if err := c.SearchRead(AccountOnlineLinkModel, criteria, NewOptions().Limit(1), aols); err != nil {
		return nil, err
	}
	if aols != nil && len(*aols) > 0 {
		return &((*aols)[0]), nil
	}
	return nil, fmt.Errorf("account.online.link was not found with criteria %v", criteria)
}

// FindAccountOnlineLinks finds account.online.link records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineLinks(criteria *Criteria, options *Options) (*AccountOnlineLinks, error) {
	aols := &AccountOnlineLinks{}
	if err := c.SearchRead(AccountOnlineLinkModel, criteria, options, aols); err != nil {
		return nil, err
	}
	return aols, nil
}

// FindAccountOnlineLinkIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineLinkIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountOnlineLinkModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountOnlineLinkId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineLinkId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineLinkModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.online.link was not found with criteria %v and options %v", criteria, options)
}
