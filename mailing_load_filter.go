package odoo

import (
	"fmt"
)

// MailingLoadFilter represents mailing.load.filter model.
type MailingLoadFilter struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FilterId    *Many2One `xmlrpc:"filter_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailingLoadFilters represents array of mailing.load.filter model.
type MailingLoadFilters []MailingLoadFilter

// MailingLoadFilterModel is the odoo model name.
const MailingLoadFilterModel = "mailing.load.filter"

// Many2One convert MailingLoadFilter to *Many2One.
func (mlf *MailingLoadFilter) Many2One() *Many2One {
	return NewMany2One(mlf.Id.Get(), "")
}

// CreateMailingLoadFilter creates a new mailing.load.filter model and returns its id.
func (c *Client) CreateMailingLoadFilter(mlf *MailingLoadFilter) (int64, error) {
	ids, err := c.CreateMailingLoadFilters([]*MailingLoadFilter{mlf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailingLoadFilter creates a new mailing.load.filter model and returns its id.
func (c *Client) CreateMailingLoadFilters(mlfs []*MailingLoadFilter) ([]int64, error) {
	var vv []interface{}
	for _, v := range mlfs {
		vv = append(vv, v)
	}
	return c.Create(MailingLoadFilterModel, vv)
}

// UpdateMailingLoadFilter updates an existing mailing.load.filter record.
func (c *Client) UpdateMailingLoadFilter(mlf *MailingLoadFilter) error {
	return c.UpdateMailingLoadFilters([]int64{mlf.Id.Get()}, mlf)
}

// UpdateMailingLoadFilters updates existing mailing.load.filter records.
// All records (represented by ids) will be updated by mlf values.
func (c *Client) UpdateMailingLoadFilters(ids []int64, mlf *MailingLoadFilter) error {
	return c.Update(MailingLoadFilterModel, ids, mlf)
}

// DeleteMailingLoadFilter deletes an existing mailing.load.filter record.
func (c *Client) DeleteMailingLoadFilter(id int64) error {
	return c.DeleteMailingLoadFilters([]int64{id})
}

// DeleteMailingLoadFilters deletes existing mailing.load.filter records.
func (c *Client) DeleteMailingLoadFilters(ids []int64) error {
	return c.Delete(MailingLoadFilterModel, ids)
}

// GetMailingLoadFilter gets mailing.load.filter existing record.
func (c *Client) GetMailingLoadFilter(id int64) (*MailingLoadFilter, error) {
	mlfs, err := c.GetMailingLoadFilters([]int64{id})
	if err != nil {
		return nil, err
	}
	if mlfs != nil && len(*mlfs) > 0 {
		return &((*mlfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mailing.load.filter not found", id)
}

// GetMailingLoadFilters gets mailing.load.filter existing records.
func (c *Client) GetMailingLoadFilters(ids []int64) (*MailingLoadFilters, error) {
	mlfs := &MailingLoadFilters{}
	if err := c.Read(MailingLoadFilterModel, ids, nil, mlfs); err != nil {
		return nil, err
	}
	return mlfs, nil
}

// FindMailingLoadFilter finds mailing.load.filter record by querying it with criteria.
func (c *Client) FindMailingLoadFilter(criteria *Criteria) (*MailingLoadFilter, error) {
	mlfs := &MailingLoadFilters{}
	if err := c.SearchRead(MailingLoadFilterModel, criteria, NewOptions().Limit(1), mlfs); err != nil {
		return nil, err
	}
	if mlfs != nil && len(*mlfs) > 0 {
		return &((*mlfs)[0]), nil
	}
	return nil, fmt.Errorf("mailing.load.filter was not found with criteria %v", criteria)
}

// FindMailingLoadFilters finds mailing.load.filter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingLoadFilters(criteria *Criteria, options *Options) (*MailingLoadFilters, error) {
	mlfs := &MailingLoadFilters{}
	if err := c.SearchRead(MailingLoadFilterModel, criteria, options, mlfs); err != nil {
		return nil, err
	}
	return mlfs, nil
}

// FindMailingLoadFilterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailingLoadFilterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailingLoadFilterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailingLoadFilterId finds record id by querying it with criteria.
func (c *Client) FindMailingLoadFilterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailingLoadFilterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mailing.load.filter was not found with criteria %v and options %v", criteria, options)
}
