package odoo

import (
	"fmt"
)

// OdooDataFetcherSharedModel represents odoo_data_fetcher.shared_model model.
type OdooDataFetcherSharedModel struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	Active             *Bool     `xmlrpc:"active,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	EmailTechnical     *String   `xmlrpc:"email_technical,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	Name               *String   `xmlrpc:"name,omptempty"`
	SharedModelError   *Bool     `xmlrpc:"shared_model_error,omptempty"`
	SharedModelMessage *String   `xmlrpc:"shared_model_message,omptempty"`
	TaskName           *String   `xmlrpc:"task_name,omptempty"`
	TaskStatusId       *String   `xmlrpc:"task_status_id,omptempty"`
	TaskStatusModel    *String   `xmlrpc:"task_status_model,omptempty"`
	TaskStatusName     *String   `xmlrpc:"task_status_name,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// OdooDataFetcherSharedModels represents array of odoo_data_fetcher.shared_model model.
type OdooDataFetcherSharedModels []OdooDataFetcherSharedModel

// OdooDataFetcherSharedModelModel is the odoo model name.
const OdooDataFetcherSharedModelModel = "odoo_data_fetcher.shared_model"

// Many2One convert OdooDataFetcherSharedModel to *Many2One.
func (os *OdooDataFetcherSharedModel) Many2One() *Many2One {
	return NewMany2One(os.Id.Get(), "")
}

// CreateOdooDataFetcherSharedModel creates a new odoo_data_fetcher.shared_model model and returns its id.
func (c *Client) CreateOdooDataFetcherSharedModel(os *OdooDataFetcherSharedModel) (int64, error) {
	ids, err := c.CreateOdooDataFetcherSharedModels([]*OdooDataFetcherSharedModel{os})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOdooDataFetcherSharedModel creates a new odoo_data_fetcher.shared_model model and returns its id.
func (c *Client) CreateOdooDataFetcherSharedModels(oss []*OdooDataFetcherSharedModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range oss {
		vv = append(vv, v)
	}
	return c.Create(OdooDataFetcherSharedModelModel, vv)
}

// UpdateOdooDataFetcherSharedModel updates an existing odoo_data_fetcher.shared_model record.
func (c *Client) UpdateOdooDataFetcherSharedModel(os *OdooDataFetcherSharedModel) error {
	return c.UpdateOdooDataFetcherSharedModels([]int64{os.Id.Get()}, os)
}

// UpdateOdooDataFetcherSharedModels updates existing odoo_data_fetcher.shared_model records.
// All records (represented by ids) will be updated by os values.
func (c *Client) UpdateOdooDataFetcherSharedModels(ids []int64, os *OdooDataFetcherSharedModel) error {
	return c.Update(OdooDataFetcherSharedModelModel, ids, os)
}

// DeleteOdooDataFetcherSharedModel deletes an existing odoo_data_fetcher.shared_model record.
func (c *Client) DeleteOdooDataFetcherSharedModel(id int64) error {
	return c.DeleteOdooDataFetcherSharedModels([]int64{id})
}

// DeleteOdooDataFetcherSharedModels deletes existing odoo_data_fetcher.shared_model records.
func (c *Client) DeleteOdooDataFetcherSharedModels(ids []int64) error {
	return c.Delete(OdooDataFetcherSharedModelModel, ids)
}

// GetOdooDataFetcherSharedModel gets odoo_data_fetcher.shared_model existing record.
func (c *Client) GetOdooDataFetcherSharedModel(id int64) (*OdooDataFetcherSharedModel, error) {
	oss, err := c.GetOdooDataFetcherSharedModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if oss != nil && len(*oss) > 0 {
		return &((*oss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of odoo_data_fetcher.shared_model not found", id)
}

// GetOdooDataFetcherSharedModels gets odoo_data_fetcher.shared_model existing records.
func (c *Client) GetOdooDataFetcherSharedModels(ids []int64) (*OdooDataFetcherSharedModels, error) {
	oss := &OdooDataFetcherSharedModels{}
	if err := c.Read(OdooDataFetcherSharedModelModel, ids, nil, oss); err != nil {
		return nil, err
	}
	return oss, nil
}

// FindOdooDataFetcherSharedModel finds odoo_data_fetcher.shared_model record by querying it with criteria.
func (c *Client) FindOdooDataFetcherSharedModel(criteria *Criteria) (*OdooDataFetcherSharedModel, error) {
	oss := &OdooDataFetcherSharedModels{}
	if err := c.SearchRead(OdooDataFetcherSharedModelModel, criteria, NewOptions().Limit(1), oss); err != nil {
		return nil, err
	}
	if oss != nil && len(*oss) > 0 {
		return &((*oss)[0]), nil
	}
	return nil, fmt.Errorf("odoo_data_fetcher.shared_model was not found with criteria %v", criteria)
}

// FindOdooDataFetcherSharedModels finds odoo_data_fetcher.shared_model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherSharedModels(criteria *Criteria, options *Options) (*OdooDataFetcherSharedModels, error) {
	oss := &OdooDataFetcherSharedModels{}
	if err := c.SearchRead(OdooDataFetcherSharedModelModel, criteria, options, oss); err != nil {
		return nil, err
	}
	return oss, nil
}

// FindOdooDataFetcherSharedModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOdooDataFetcherSharedModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OdooDataFetcherSharedModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOdooDataFetcherSharedModelId finds record id by querying it with criteria.
func (c *Client) FindOdooDataFetcherSharedModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OdooDataFetcherSharedModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("odoo_data_fetcher.shared_model was not found with criteria %v and options %v", criteria, options)
}
