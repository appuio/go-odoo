package odoo

import (
	"fmt"
)

// MailingList represents mailing.list model.
type MailingList struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	Active                  *Bool      `xmlrpc:"active,omptempty"`
	ContactCount            *Int       `xmlrpc:"contact_count,omptempty"`
	ContactCountBlacklisted *Int       `xmlrpc:"contact_count_blacklisted,omptempty"`
	ContactCountEmail       *Int       `xmlrpc:"contact_count_email,omptempty"`
	ContactCountOptOut      *Int       `xmlrpc:"contact_count_opt_out,omptempty"`
	ContactIds              *Relation  `xmlrpc:"contact_ids,omptempty"`
	ContactPctBlacklisted   *Float     `xmlrpc:"contact_pct_blacklisted,omptempty"`
	ContactPctBounce        *Float     `xmlrpc:"contact_pct_bounce,omptempty"`
	ContactPctOptOut        *Float     `xmlrpc:"contact_pct_opt_out,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	Dynamic                 *Bool      `xmlrpc:"dynamic,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	IsPublic                *Bool      `xmlrpc:"is_public,omptempty"`
	IsSynced                *Bool      `xmlrpc:"is_synced,omptempty"`
	MailingCount            *Int       `xmlrpc:"mailing_count,omptempty"`
	MailingIds              *Relation  `xmlrpc:"mailing_ids,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	PartnerCategory         *Many2One  `xmlrpc:"partner_category,omptempty"`
	PartnerMandatory        *Bool      `xmlrpc:"partner_mandatory,omptempty"`
	SubscriptionIds         *Relation  `xmlrpc:"subscription_ids,omptempty"`
	SyncDomain              *String    `xmlrpc:"sync_domain,omptempty"`
	SyncMethod              *Selection `xmlrpc:"sync_method,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailingLists represents array of mailing.list model.
type MailingLists []MailingList

// MailingListModel is the odoo model name.
const MailingListModel = "mailing.list"

// Many2One convert MailingList to *Many2One.
func (ml *MailingList) Many2One() *Many2One {
	return NewMany2One(ml.Id.Get(), "")
}

// CreateMailingList creates a new mailing.list model and returns its id.
func (c *Client) CreateMailingList(ml *MailingList) (int64, error) {
	ids, err := c.CreateMailingLists([]*MailingList{ml})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingList creates a new mailing.list model and returns its id.
func (c *Client) CreateMailingLists(mls []*MailingList) ([]int64, error) {
	var vv []interface{}
	for _, v := range mls {
		vv = append(vv, v)
	}
	return c.Create(MailingListModel, vv)
}

// UpdateMailingList updates an existing mailing.list record.
func (c *Client) UpdateMailingList(ml *MailingList) error {
	return c.UpdateMailingLists([]int64{ml.Id.Get()}, ml)
}

// UpdateMailingLists updates existing mailing.list records.
// All records (represented by ids) will be updated by ml values.
func (c *Client) UpdateMailingLists(ids []int64, ml *MailingList) error {
	return c.Update(MailingListModel, ids, ml)
}

// DeleteMailingList deletes an existing mailing.list record.
func (c *Client) DeleteMailingList(id int64) error {
	return c.DeleteMailingLists([]int64{id})
}

// DeleteMailingLists deletes existing mailing.list records.
func (c *Client) DeleteMailingLists(ids []int64) error {
	return c.Delete(MailingListModel, ids)
}

// GetMailingList gets mailing.list existing record.
func (c *Client) GetMailingList(id int64) (*MailingList, error) {
	mls, err := c.GetMailingLists([]int64{id})
	if err != nil {
		return nil, err
	}
	if mls != nil && len(*mls) > 0 {
		return &((*mls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.list not found", id)
}

// GetMailingLists gets mailing.list existing records.
func (c *Client) GetMailingLists(ids []int64) (*MailingLists, error) {
	mls := &MailingLists{}
	if err := c.Read(MailingListModel, ids, nil, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMailingList finds mailing.list record by querying it with criteria.
func (c *Client) FindMailingList(criteria *Criteria) (*MailingList, error) {
	mls := &MailingLists{}
	if err := c.SearchRead(MailingListModel, criteria, NewOptions().Limit(1), mls); err != nil {
		return nil, err
	}
	if mls != nil && len(*mls) > 0 {
		return &((*mls)[0]), nil
	}
	return nil, fmt.Errorf("mailing.list was not found with criteria %v", criteria)
}

// FindMailingLists finds mailing.list records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingLists(criteria *Criteria, options *Options) (*MailingLists, error) {
	mls := &MailingLists{}
	if err := c.SearchRead(MailingListModel, criteria, options, mls); err != nil {
		return nil, err
	}
	return mls, nil
}

// FindMailingListIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingListIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingListModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingListId finds record id by querying it with criteria.
func (c *Client) FindMailingListId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingListModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.list was not found with criteria %v and options %v", criteria, options)
}
