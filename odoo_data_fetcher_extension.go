package odoo

import (
	"fmt"
)

// OdooDataFetcherExtension represents odoo_data_fetcher.extension model.
type OdooDataFetcherExtension struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	AddUrl            *String   `xmlrpc:"add_url,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DataFetcherIds    *Relation `xmlrpc:"data_fetcher_ids,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	RemoteError       *Bool     `xmlrpc:"remote_error,omptempty"`
	StatusChangedDate *Time     `xmlrpc:"status_changed_date,omptempty"`
	Statuses          *String   `xmlrpc:"statuses,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OdooDataFetcherExtensions represents array of odoo_data_fetcher.extension model.
type OdooDataFetcherExtensions []OdooDataFetcherExtension

// OdooDataFetcherExtensionModel is the odoo model name.
const OdooDataFetcherExtensionModel = "odoo_data_fetcher.extension"

// Many2One convert OdooDataFetcherExtension to *Many2One.
func (oe *OdooDataFetcherExtension) Many2One() *Many2One {
	return NewMany2One(oe.Id.Get(), "")
}

// CreateOdooDataFetcherExtension creates a new odoo_data_fetcher.extension model and returns its id.
func (c *Client) CreateOdooDataFetcherExtension(oe *OdooDataFetcherExtension) (int64, error) {
	ids, err := c.CreateOdooDataFetcherExtensions([]*OdooDataFetcherExtension{oe})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOdooDataFetcherExtension creates a new odoo_data_fetcher.extension model and returns its id.
func (c *Client) CreateOdooDataFetcherExtensions(oes []*OdooDataFetcherExtension) ([]int64, error) {
	var vv []interface{}
	for _, v := range oes {
		vv = append(vv, v)
	}
	return c.Create(OdooDataFetcherExtensionModel, vv)
}

// UpdateOdooDataFetcherExtension updates an existing odoo_data_fetcher.extension record.
func (c *Client) UpdateOdooDataFetcherExtension(oe *OdooDataFetcherExtension) error {
	return c.UpdateOdooDataFetcherExtensions([]int64{oe.Id.Get()}, oe)
}

// UpdateOdooDataFetcherExtensions updates existing odoo_data_fetcher.extension records.
// All records (represented by ids) will be updated by oe values.
func (c *Client) UpdateOdooDataFetcherExtensions(ids []int64, oe *OdooDataFetcherExtension) error {
	return c.Update(OdooDataFetcherExtensionModel, ids, oe)
}

// DeleteOdooDataFetcherExtension deletes an existing odoo_data_fetcher.extension record.
func (c *Client) DeleteOdooDataFetcherExtension(id int64) error {
	return c.DeleteOdooDataFetcherExtensions([]int64{id})
}

// DeleteOdooDataFetcherExtensions deletes existing odoo_data_fetcher.extension records.
func (c *Client) DeleteOdooDataFetcherExtensions(ids []int64) error {
	return c.Delete(OdooDataFetcherExtensionModel, ids)
}

// GetOdooDataFetcherExtension gets odoo_data_fetcher.extension existing record.
func (c *Client) GetOdooDataFetcherExtension(id int64) (*OdooDataFetcherExtension, error) {
	oes, err := c.GetOdooDataFetcherExtensions([]int64{id})
	if err != nil {
		return nil, err
	}
	if oes != nil && len(*oes) > 0 {
		return &((*oes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of odoo_data_fetcher.extension not found", id)
}

// GetOdooDataFetcherExtensions gets odoo_data_fetcher.extension existing records.
func (c *Client) GetOdooDataFetcherExtensions(ids []int64) (*OdooDataFetcherExtensions, error) {
	oes := &OdooDataFetcherExtensions{}
	if err := c.Read(OdooDataFetcherExtensionModel, ids, nil, oes); err != nil {
		return nil, err
	}
	return oes, nil
}

// FindOdooDataFetcherExtension finds odoo_data_fetcher.extension record by querying it with criteria.
func (c *Client) FindOdooDataFetcherExtension(criteria *Criteria) (*OdooDataFetcherExtension, error) {
	oes := &OdooDataFetcherExtensions{}
	if err := c.SearchRead(OdooDataFetcherExtensionModel, criteria, NewOptions().Limit(1), oes); err != nil {
		return nil, err
	}
	if oes != nil && len(*oes) > 0 {
		return &((*oes)[0]), nil
	}
	return nil, fmt.Errorf("odoo_data_fetcher.extension was not found with criteria %v", criteria)
}

// FindOdooDataFetcherExtensions finds odoo_data_fetcher.extension records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherExtensions(criteria *Criteria, options *Options) (*OdooDataFetcherExtensions, error) {
	oes := &OdooDataFetcherExtensions{}
	if err := c.SearchRead(OdooDataFetcherExtensionModel, criteria, options, oes); err != nil {
		return nil, err
	}
	return oes, nil
}

// FindOdooDataFetcherExtensionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherExtensionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OdooDataFetcherExtensionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOdooDataFetcherExtensionId finds record id by querying it with criteria.
func (c *Client) FindOdooDataFetcherExtensionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OdooDataFetcherExtensionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("odoo_data_fetcher.extension was not found with criteria %v and options %v", criteria, options)
}
