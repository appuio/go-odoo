package odoo

import (
	"fmt"
)

// OdooDataFetcherDataFetcher represents odoo_data_fetcher.data_fetcher model.
type OdooDataFetcherDataFetcher struct {
	LastUpdate              *Time     `xmlrpc:"__last_update,omptempty"`
	ConnectionError         *Bool     `xmlrpc:"connection_error,omptempty"`
	ConnectionErrorMessage  *String   `xmlrpc:"connection_error_message,omptempty"`
	ConnectionId            *Many2One `xmlrpc:"connection_id,omptempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String   `xmlrpc:"display_name,omptempty"`
	ErrorMessage            *String   `xmlrpc:"error_message,omptempty"`
	Id                      *Int      `xmlrpc:"id,omptempty"`
	LastUpdateDate          *Time     `xmlrpc:"last_update_date,omptempty"`
	ModelName               *String   `xmlrpc:"model_name,omptempty"`
	Name                    *String   `xmlrpc:"name,omptempty"`
	SharedModelError        *Bool     `xmlrpc:"shared_model_error,omptempty"`
	SharedModelErrorMessage *String   `xmlrpc:"shared_model_error_message,omptempty"`
	Status                  *String   `xmlrpc:"status,omptempty"`
	TaskId                  *Int      `xmlrpc:"task_id,omptempty"`
	TaskName                *String   `xmlrpc:"task_name,omptempty"`
	UrlError                *Bool     `xmlrpc:"url_error,omptempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OdooDataFetcherDataFetchers represents array of odoo_data_fetcher.data_fetcher model.
type OdooDataFetcherDataFetchers []OdooDataFetcherDataFetcher

// OdooDataFetcherDataFetcherModel is the odoo model name.
const OdooDataFetcherDataFetcherModel = "odoo_data_fetcher.data_fetcher"

// Many2One convert OdooDataFetcherDataFetcher to *Many2One.
func (od *OdooDataFetcherDataFetcher) Many2One() *Many2One {
	return NewMany2One(od.Id.Get(), "")
}

// CreateOdooDataFetcherDataFetcher creates a new odoo_data_fetcher.data_fetcher model and returns its id.
func (c *Client) CreateOdooDataFetcherDataFetcher(od *OdooDataFetcherDataFetcher) (int64, error) {
	ids, err := c.CreateOdooDataFetcherDataFetchers([]*OdooDataFetcherDataFetcher{od})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOdooDataFetcherDataFetcher creates a new odoo_data_fetcher.data_fetcher model and returns its id.
func (c *Client) CreateOdooDataFetcherDataFetchers(ods []*OdooDataFetcherDataFetcher) ([]int64, error) {
	var vv []interface{}
	for _, v := range ods {
		vv = append(vv, v)
	}
	return c.Create(OdooDataFetcherDataFetcherModel, vv)
}

// UpdateOdooDataFetcherDataFetcher updates an existing odoo_data_fetcher.data_fetcher record.
func (c *Client) UpdateOdooDataFetcherDataFetcher(od *OdooDataFetcherDataFetcher) error {
	return c.UpdateOdooDataFetcherDataFetchers([]int64{od.Id.Get()}, od)
}

// UpdateOdooDataFetcherDataFetchers updates existing odoo_data_fetcher.data_fetcher records.
// All records (represented by ids) will be updated by od values.
func (c *Client) UpdateOdooDataFetcherDataFetchers(ids []int64, od *OdooDataFetcherDataFetcher) error {
	return c.Update(OdooDataFetcherDataFetcherModel, ids, od)
}

// DeleteOdooDataFetcherDataFetcher deletes an existing odoo_data_fetcher.data_fetcher record.
func (c *Client) DeleteOdooDataFetcherDataFetcher(id int64) error {
	return c.DeleteOdooDataFetcherDataFetchers([]int64{id})
}

// DeleteOdooDataFetcherDataFetchers deletes existing odoo_data_fetcher.data_fetcher records.
func (c *Client) DeleteOdooDataFetcherDataFetchers(ids []int64) error {
	return c.Delete(OdooDataFetcherDataFetcherModel, ids)
}

// GetOdooDataFetcherDataFetcher gets odoo_data_fetcher.data_fetcher existing record.
func (c *Client) GetOdooDataFetcherDataFetcher(id int64) (*OdooDataFetcherDataFetcher, error) {
	ods, err := c.GetOdooDataFetcherDataFetchers([]int64{id})
	if err != nil {
		return nil, err
	}
	if ods != nil && len(*ods) > 0 {
		return &((*ods)[0]), nil
	}
	return nil, fmt.Errorf("id %v of odoo_data_fetcher.data_fetcher not found", id)
}

// GetOdooDataFetcherDataFetchers gets odoo_data_fetcher.data_fetcher existing records.
func (c *Client) GetOdooDataFetcherDataFetchers(ids []int64) (*OdooDataFetcherDataFetchers, error) {
	ods := &OdooDataFetcherDataFetchers{}
	if err := c.Read(OdooDataFetcherDataFetcherModel, ids, nil, ods); err != nil {
		return nil, err
	}
	return ods, nil
}

// FindOdooDataFetcherDataFetcher finds odoo_data_fetcher.data_fetcher record by querying it with criteria.
func (c *Client) FindOdooDataFetcherDataFetcher(criteria *Criteria) (*OdooDataFetcherDataFetcher, error) {
	ods := &OdooDataFetcherDataFetchers{}
	if err := c.SearchRead(OdooDataFetcherDataFetcherModel, criteria, NewOptions().Limit(1), ods); err != nil {
		return nil, err
	}
	if ods != nil && len(*ods) > 0 {
		return &((*ods)[0]), nil
	}
	return nil, fmt.Errorf("odoo_data_fetcher.data_fetcher was not found with criteria %v", criteria)
}

// FindOdooDataFetcherDataFetchers finds odoo_data_fetcher.data_fetcher records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherDataFetchers(criteria *Criteria, options *Options) (*OdooDataFetcherDataFetchers, error) {
	ods := &OdooDataFetcherDataFetchers{}
	if err := c.SearchRead(OdooDataFetcherDataFetcherModel, criteria, options, ods); err != nil {
		return nil, err
	}
	return ods, nil
}

// FindOdooDataFetcherDataFetcherIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherDataFetcherIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OdooDataFetcherDataFetcherModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOdooDataFetcherDataFetcherId finds record id by querying it with criteria.
func (c *Client) FindOdooDataFetcherDataFetcherId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OdooDataFetcherDataFetcherModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("odoo_data_fetcher.data_fetcher was not found with criteria %v and options %v", criteria, options)
}
