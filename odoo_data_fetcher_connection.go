package odoo

import (
	"fmt"
)

// OdooDataFetcherConnection represents odoo_data_fetcher.connection model.
type OdooDataFetcherConnection struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Active          *Bool     `xmlrpc:"active,omptempty"`
	ConnMessage     *String   `xmlrpc:"conn_message,omptempty"`
	ConnectionError *Bool     `xmlrpc:"connection_error,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	Database        *String   `xmlrpc:"database,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Domain          *Relation `xmlrpc:"domain,omptempty"`
	EmailFunctional *String   `xmlrpc:"email_functional,omptempty"`
	EmailTechnical  *String   `xmlrpc:"email_technical,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	LastUpdatedDate *Time     `xmlrpc:"last_updated_date,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	Password        *String   `xmlrpc:"password,omptempty"`
	Port            *Int      `xmlrpc:"port,omptempty"`
	User            *String   `xmlrpc:"user,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OdooDataFetcherConnections represents array of odoo_data_fetcher.connection model.
type OdooDataFetcherConnections []OdooDataFetcherConnection

// OdooDataFetcherConnectionModel is the odoo model name.
const OdooDataFetcherConnectionModel = "odoo_data_fetcher.connection"

// Many2One convert OdooDataFetcherConnection to *Many2One.
func (oc *OdooDataFetcherConnection) Many2One() *Many2One {
	return NewMany2One(oc.Id.Get(), "")
}

// CreateOdooDataFetcherConnection creates a new odoo_data_fetcher.connection model and returns its id.
func (c *Client) CreateOdooDataFetcherConnection(oc *OdooDataFetcherConnection) (int64, error) {
	ids, err := c.CreateOdooDataFetcherConnections([]*OdooDataFetcherConnection{oc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOdooDataFetcherConnection creates a new odoo_data_fetcher.connection model and returns its id.
func (c *Client) CreateOdooDataFetcherConnections(ocs []*OdooDataFetcherConnection) ([]int64, error) {
	var vv []interface{}
	for _, v := range ocs {
		vv = append(vv, v)
	}
	return c.Create(OdooDataFetcherConnectionModel, vv)
}

// UpdateOdooDataFetcherConnection updates an existing odoo_data_fetcher.connection record.
func (c *Client) UpdateOdooDataFetcherConnection(oc *OdooDataFetcherConnection) error {
	return c.UpdateOdooDataFetcherConnections([]int64{oc.Id.Get()}, oc)
}

// UpdateOdooDataFetcherConnections updates existing odoo_data_fetcher.connection records.
// All records (represented by ids) will be updated by oc values.
func (c *Client) UpdateOdooDataFetcherConnections(ids []int64, oc *OdooDataFetcherConnection) error {
	return c.Update(OdooDataFetcherConnectionModel, ids, oc)
}

// DeleteOdooDataFetcherConnection deletes an existing odoo_data_fetcher.connection record.
func (c *Client) DeleteOdooDataFetcherConnection(id int64) error {
	return c.DeleteOdooDataFetcherConnections([]int64{id})
}

// DeleteOdooDataFetcherConnections deletes existing odoo_data_fetcher.connection records.
func (c *Client) DeleteOdooDataFetcherConnections(ids []int64) error {
	return c.Delete(OdooDataFetcherConnectionModel, ids)
}

// GetOdooDataFetcherConnection gets odoo_data_fetcher.connection existing record.
func (c *Client) GetOdooDataFetcherConnection(id int64) (*OdooDataFetcherConnection, error) {
	ocs, err := c.GetOdooDataFetcherConnections([]int64{id})
	if err != nil {
		return nil, err
	}
	if ocs != nil && len(*ocs) > 0 {
		return &((*ocs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of odoo_data_fetcher.connection not found", id)
}

// GetOdooDataFetcherConnections gets odoo_data_fetcher.connection existing records.
func (c *Client) GetOdooDataFetcherConnections(ids []int64) (*OdooDataFetcherConnections, error) {
	ocs := &OdooDataFetcherConnections{}
	if err := c.Read(OdooDataFetcherConnectionModel, ids, nil, ocs); err != nil {
		return nil, err
	}
	return ocs, nil
}

// FindOdooDataFetcherConnection finds odoo_data_fetcher.connection record by querying it with criteria.
func (c *Client) FindOdooDataFetcherConnection(criteria *Criteria) (*OdooDataFetcherConnection, error) {
	ocs := &OdooDataFetcherConnections{}
	if err := c.SearchRead(OdooDataFetcherConnectionModel, criteria, NewOptions().Limit(1), ocs); err != nil {
		return nil, err
	}
	if ocs != nil && len(*ocs) > 0 {
		return &((*ocs)[0]), nil
	}
	return nil, fmt.Errorf("odoo_data_fetcher.connection was not found with criteria %v", criteria)
}

// FindOdooDataFetcherConnections finds odoo_data_fetcher.connection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherConnections(criteria *Criteria, options *Options) (*OdooDataFetcherConnections, error) {
	ocs := &OdooDataFetcherConnections{}
	if err := c.SearchRead(OdooDataFetcherConnectionModel, criteria, options, ocs); err != nil {
		return nil, err
	}
	return ocs, nil
}

// FindOdooDataFetcherConnectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherConnectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OdooDataFetcherConnectionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOdooDataFetcherConnectionId finds record id by querying it with criteria.
func (c *Client) FindOdooDataFetcherConnectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OdooDataFetcherConnectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("odoo_data_fetcher.connection was not found with criteria %v and options %v", criteria, options)
}
