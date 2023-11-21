package odoo

import (
	"fmt"
)

// MailingFilter represents mailing.filter model.
type MailingFilter struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	MailingDomain    *String   `xmlrpc:"mailing_domain,omptempty"`
	MailingModelId   *Many2One `xmlrpc:"mailing_model_id,omptempty"`
	MailingModelName *String   `xmlrpc:"mailing_model_name,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailingFilters represents array of mailing.filter model.
type MailingFilters []MailingFilter

// MailingFilterModel is the odoo model name.
const MailingFilterModel = "mailing.filter"

// Many2One convert MailingFilter to *Many2One.
func (mf *MailingFilter) Many2One() *Many2One {
	return NewMany2One(mf.Id.Get(), "")
}

// CreateMailingFilter creates a new mailing.filter model and returns its id.
func (c *Client) CreateMailingFilter(mf *MailingFilter) (int64, error) {
	ids, err := c.CreateMailingFilters([]*MailingFilter{mf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingFilter creates a new mailing.filter model and returns its id.
func (c *Client) CreateMailingFilters(mfs []*MailingFilter) ([]int64, error) {
	var vv []interface{}
	for _, v := range mfs {
		vv = append(vv, v)
	}
	return c.Create(MailingFilterModel, vv)
}

// UpdateMailingFilter updates an existing mailing.filter record.
func (c *Client) UpdateMailingFilter(mf *MailingFilter) error {
	return c.UpdateMailingFilters([]int64{mf.Id.Get()}, mf)
}

// UpdateMailingFilters updates existing mailing.filter records.
// All records (represented by ids) will be updated by mf values.
func (c *Client) UpdateMailingFilters(ids []int64, mf *MailingFilter) error {
	return c.Update(MailingFilterModel, ids, mf)
}

// DeleteMailingFilter deletes an existing mailing.filter record.
func (c *Client) DeleteMailingFilter(id int64) error {
	return c.DeleteMailingFilters([]int64{id})
}

// DeleteMailingFilters deletes existing mailing.filter records.
func (c *Client) DeleteMailingFilters(ids []int64) error {
	return c.Delete(MailingFilterModel, ids)
}

// GetMailingFilter gets mailing.filter existing record.
func (c *Client) GetMailingFilter(id int64) (*MailingFilter, error) {
	mfs, err := c.GetMailingFilters([]int64{id})
	if err != nil {
		return nil, err
	}
	if mfs != nil && len(*mfs) > 0 {
		return &((*mfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.filter not found", id)
}

// GetMailingFilters gets mailing.filter existing records.
func (c *Client) GetMailingFilters(ids []int64) (*MailingFilters, error) {
	mfs := &MailingFilters{}
	if err := c.Read(MailingFilterModel, ids, nil, mfs); err != nil {
		return nil, err
	}
	return mfs, nil
}

// FindMailingFilter finds mailing.filter record by querying it with criteria.
func (c *Client) FindMailingFilter(criteria *Criteria) (*MailingFilter, error) {
	mfs := &MailingFilters{}
	if err := c.SearchRead(MailingFilterModel, criteria, NewOptions().Limit(1), mfs); err != nil {
		return nil, err
	}
	if mfs != nil && len(*mfs) > 0 {
		return &((*mfs)[0]), nil
	}
	return nil, fmt.Errorf("mailing.filter was not found with criteria %v", criteria)
}

// FindMailingFilters finds mailing.filter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingFilters(criteria *Criteria, options *Options) (*MailingFilters, error) {
	mfs := &MailingFilters{}
	if err := c.SearchRead(MailingFilterModel, criteria, options, mfs); err != nil {
		return nil, err
	}
	return mfs, nil
}

// FindMailingFilterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingFilterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingFilterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingFilterId finds record id by querying it with criteria.
func (c *Client) FindMailingFilterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingFilterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.filter was not found with criteria %v and options %v", criteria, options)
}
