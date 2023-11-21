package odoo

import (
	"fmt"
)

// OdooDataFetcherDomain represents odoo_data_fetcher.domain model.
type OdooDataFetcherDomain struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	ConnectionId *Many2One `xmlrpc:"connection_id,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	Sequence     *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OdooDataFetcherDomains represents array of odoo_data_fetcher.domain model.
type OdooDataFetcherDomains []OdooDataFetcherDomain

// OdooDataFetcherDomainModel is the odoo model name.
const OdooDataFetcherDomainModel = "odoo_data_fetcher.domain"

// Many2One convert OdooDataFetcherDomain to *Many2One.
func (od *OdooDataFetcherDomain) Many2One() *Many2One {
	return NewMany2One(od.Id.Get(), "")
}

// CreateOdooDataFetcherDomain creates a new odoo_data_fetcher.domain model and returns its id.
func (c *Client) CreateOdooDataFetcherDomain(od *OdooDataFetcherDomain) (int64, error) {
	ids, err := c.CreateOdooDataFetcherDomains([]*OdooDataFetcherDomain{od})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOdooDataFetcherDomain creates a new odoo_data_fetcher.domain model and returns its id.
func (c *Client) CreateOdooDataFetcherDomains(ods []*OdooDataFetcherDomain) ([]int64, error) {
	var vv []interface{}
	for _, v := range ods {
		vv = append(vv, v)
	}
	return c.Create(OdooDataFetcherDomainModel, vv)
}

// UpdateOdooDataFetcherDomain updates an existing odoo_data_fetcher.domain record.
func (c *Client) UpdateOdooDataFetcherDomain(od *OdooDataFetcherDomain) error {
	return c.UpdateOdooDataFetcherDomains([]int64{od.Id.Get()}, od)
}

// UpdateOdooDataFetcherDomains updates existing odoo_data_fetcher.domain records.
// All records (represented by ids) will be updated by od values.
func (c *Client) UpdateOdooDataFetcherDomains(ids []int64, od *OdooDataFetcherDomain) error {
	return c.Update(OdooDataFetcherDomainModel, ids, od)
}

// DeleteOdooDataFetcherDomain deletes an existing odoo_data_fetcher.domain record.
func (c *Client) DeleteOdooDataFetcherDomain(id int64) error {
	return c.DeleteOdooDataFetcherDomains([]int64{id})
}

// DeleteOdooDataFetcherDomains deletes existing odoo_data_fetcher.domain records.
func (c *Client) DeleteOdooDataFetcherDomains(ids []int64) error {
	return c.Delete(OdooDataFetcherDomainModel, ids)
}

// GetOdooDataFetcherDomain gets odoo_data_fetcher.domain existing record.
func (c *Client) GetOdooDataFetcherDomain(id int64) (*OdooDataFetcherDomain, error) {
	ods, err := c.GetOdooDataFetcherDomains([]int64{id})
	if err != nil {
		return nil, err
	}
	if ods != nil && len(*ods) > 0 {
		return &((*ods)[0]), nil
	}
	return nil, fmt.Errorf("id %v of odoo_data_fetcher.domain not found", id)
}

// GetOdooDataFetcherDomains gets odoo_data_fetcher.domain existing records.
func (c *Client) GetOdooDataFetcherDomains(ids []int64) (*OdooDataFetcherDomains, error) {
	ods := &OdooDataFetcherDomains{}
	if err := c.Read(OdooDataFetcherDomainModel, ids, nil, ods); err != nil {
		return nil, err
	}
	return ods, nil
}

// FindOdooDataFetcherDomain finds odoo_data_fetcher.domain record by querying it with criteria.
func (c *Client) FindOdooDataFetcherDomain(criteria *Criteria) (*OdooDataFetcherDomain, error) {
	ods := &OdooDataFetcherDomains{}
	if err := c.SearchRead(OdooDataFetcherDomainModel, criteria, NewOptions().Limit(1), ods); err != nil {
		return nil, err
	}
	if ods != nil && len(*ods) > 0 {
		return &((*ods)[0]), nil
	}
	return nil, fmt.Errorf("odoo_data_fetcher.domain was not found with criteria %v", criteria)
}

// FindOdooDataFetcherDomains finds odoo_data_fetcher.domain records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherDomains(criteria *Criteria, options *Options) (*OdooDataFetcherDomains, error) {
	ods := &OdooDataFetcherDomains{}
	if err := c.SearchRead(OdooDataFetcherDomainModel, criteria, options, ods); err != nil {
		return nil, err
	}
	return ods, nil
}

// FindOdooDataFetcherDomainIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherDomainIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OdooDataFetcherDomainModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOdooDataFetcherDomainId finds record id by querying it with criteria.
func (c *Client) FindOdooDataFetcherDomainId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OdooDataFetcherDomainModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("odoo_data_fetcher.domain was not found with criteria %v and options %v", criteria, options)
}
