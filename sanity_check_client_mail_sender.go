package odoo

import (
	"fmt"
)

// SanityCheckClientMailSender represents sanity_check_client_mail_sender model.
type SanityCheckClientMailSender struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SanityCheckClientMailSenders represents array of sanity_check_client_mail_sender model.
type SanityCheckClientMailSenders []SanityCheckClientMailSender

// SanityCheckClientMailSenderModel is the odoo model name.
const SanityCheckClientMailSenderModel = "sanity_check_client_mail_sender"

// Many2One convert SanityCheckClientMailSender to *Many2One.
func (s *SanityCheckClientMailSender) Many2One() *Many2One {
	return NewMany2One(s.Id.Get(), "")
}

// CreateSanityCheckClientMailSender creates a new sanity_check_client_mail_sender model and returns its id.
func (c *Client) CreateSanityCheckClientMailSender(s *SanityCheckClientMailSender) (int64, error) {
	ids, err := c.CreateSanityCheckClientMailSenders([]*SanityCheckClientMailSender{s})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSanityCheckClientMailSender creates a new sanity_check_client_mail_sender model and returns its id.
func (c *Client) CreateSanityCheckClientMailSenders(ss []*SanityCheckClientMailSender) ([]int64, error) {
	var vv []interface{}
	for _, v := range ss {
		vv = append(vv, v)
	}
	return c.Create(SanityCheckClientMailSenderModel, vv)
}

// UpdateSanityCheckClientMailSender updates an existing sanity_check_client_mail_sender record.
func (c *Client) UpdateSanityCheckClientMailSender(s *SanityCheckClientMailSender) error {
	return c.UpdateSanityCheckClientMailSenders([]int64{s.Id.Get()}, s)
}

// UpdateSanityCheckClientMailSenders updates existing sanity_check_client_mail_sender records.
// All records (represented by ids) will be updated by s values.
func (c *Client) UpdateSanityCheckClientMailSenders(ids []int64, s *SanityCheckClientMailSender) error {
	return c.Update(SanityCheckClientMailSenderModel, ids, s)
}

// DeleteSanityCheckClientMailSender deletes an existing sanity_check_client_mail_sender record.
func (c *Client) DeleteSanityCheckClientMailSender(id int64) error {
	return c.DeleteSanityCheckClientMailSenders([]int64{id})
}

// DeleteSanityCheckClientMailSenders deletes existing sanity_check_client_mail_sender records.
func (c *Client) DeleteSanityCheckClientMailSenders(ids []int64) error {
	return c.Delete(SanityCheckClientMailSenderModel, ids)
}

// GetSanityCheckClientMailSender gets sanity_check_client_mail_sender existing record.
func (c *Client) GetSanityCheckClientMailSender(id int64) (*SanityCheckClientMailSender, error) {
	ss, err := c.GetSanityCheckClientMailSenders([]int64{id})
	if err != nil {
		return nil, err
	}
	if ss != nil && len(*ss) > 0 {
		return &((*ss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sanity_check_client_mail_sender not found", id)
}

// GetSanityCheckClientMailSenders gets sanity_check_client_mail_sender existing records.
func (c *Client) GetSanityCheckClientMailSenders(ids []int64) (*SanityCheckClientMailSenders, error) {
	ss := &SanityCheckClientMailSenders{}
	if err := c.Read(SanityCheckClientMailSenderModel, ids, nil, ss); err != nil {
		return nil, err
	}
	return ss, nil
}

// FindSanityCheckClientMailSender finds sanity_check_client_mail_sender record by querying it with criteria.
func (c *Client) FindSanityCheckClientMailSender(criteria *Criteria) (*SanityCheckClientMailSender, error) {
	ss := &SanityCheckClientMailSenders{}
	if err := c.SearchRead(SanityCheckClientMailSenderModel, criteria, NewOptions().Limit(1), ss); err != nil {
		return nil, err
	}
	if ss != nil && len(*ss) > 0 {
		return &((*ss)[0]), nil
	}
	return nil, fmt.Errorf("sanity_check_client_mail_sender was not found with criteria %v", criteria)
}

// FindSanityCheckClientMailSenders finds sanity_check_client_mail_sender records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSanityCheckClientMailSenders(criteria *Criteria, options *Options) (*SanityCheckClientMailSenders, error) {
	ss := &SanityCheckClientMailSenders{}
	if err := c.SearchRead(SanityCheckClientMailSenderModel, criteria, options, ss); err != nil {
		return nil, err
	}
	return ss, nil
}

// FindSanityCheckClientMailSenderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSanityCheckClientMailSenderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SanityCheckClientMailSenderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSanityCheckClientMailSenderId finds record id by querying it with criteria.
func (c *Client) FindSanityCheckClientMailSenderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SanityCheckClientMailSenderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sanity_check_client_mail_sender was not found with criteria %v and options %v", criteria, options)
}
