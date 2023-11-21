package odoo

import (
	"fmt"
)

// DataCleaningModel represents data_cleaning.model model.
type DataCleaningModel struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Active                *Bool      `xmlrpc:"active,omptempty"`
	CleaningMode          *Selection `xmlrpc:"cleaning_mode,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	LastNotification      *Time      `xmlrpc:"last_notification,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	NotifyFrequency       *Int       `xmlrpc:"notify_frequency,omptempty"`
	NotifyFrequencyPeriod *Selection `xmlrpc:"notify_frequency_period,omptempty"`
	NotifyUserIds         *Relation  `xmlrpc:"notify_user_ids,omptempty"`
	RecordsToCleanCount   *Int       `xmlrpc:"records_to_clean_count,omptempty"`
	ResModelId            *Many2One  `xmlrpc:"res_model_id,omptempty"`
	ResModelName          *String    `xmlrpc:"res_model_name,omptempty"`
	RuleIds               *Relation  `xmlrpc:"rule_ids,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DataCleaningModels represents array of data_cleaning.model model.
type DataCleaningModels []DataCleaningModel

// DataCleaningModelModel is the odoo model name.
const DataCleaningModelModel = "data_cleaning.model"

// Many2One convert DataCleaningModel to *Many2One.
func (dm *DataCleaningModel) Many2One() *Many2One {
	return NewMany2One(dm.Id.Get(), "")
}

// CreateDataCleaningModel creates a new data_cleaning.model model and returns its id.
func (c *Client) CreateDataCleaningModel(dm *DataCleaningModel) (int64, error) {
	ids, err := c.CreateDataCleaningModels([]*DataCleaningModel{dm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataCleaningModel creates a new data_cleaning.model model and returns its id.
func (c *Client) CreateDataCleaningModels(dms []*DataCleaningModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range dms {
		vv = append(vv, v)
	}
	return c.Create(DataCleaningModelModel, vv)
}

// UpdateDataCleaningModel updates an existing data_cleaning.model record.
func (c *Client) UpdateDataCleaningModel(dm *DataCleaningModel) error {
	return c.UpdateDataCleaningModels([]int64{dm.Id.Get()}, dm)
}

// UpdateDataCleaningModels updates existing data_cleaning.model records.
// All records (represented by ids) will be updated by dm values.
func (c *Client) UpdateDataCleaningModels(ids []int64, dm *DataCleaningModel) error {
	return c.Update(DataCleaningModelModel, ids, dm)
}

// DeleteDataCleaningModel deletes an existing data_cleaning.model record.
func (c *Client) DeleteDataCleaningModel(id int64) error {
	return c.DeleteDataCleaningModels([]int64{id})
}

// DeleteDataCleaningModels deletes existing data_cleaning.model records.
func (c *Client) DeleteDataCleaningModels(ids []int64) error {
	return c.Delete(DataCleaningModelModel, ids)
}

// GetDataCleaningModel gets data_cleaning.model existing record.
func (c *Client) GetDataCleaningModel(id int64) (*DataCleaningModel, error) {
	dms, err := c.GetDataCleaningModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if dms != nil && len(*dms) > 0 {
		return &((*dms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_cleaning.model not found", id)
}

// GetDataCleaningModels gets data_cleaning.model existing records.
func (c *Client) GetDataCleaningModels(ids []int64) (*DataCleaningModels, error) {
	dms := &DataCleaningModels{}
	if err := c.Read(DataCleaningModelModel, ids, nil, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDataCleaningModel finds data_cleaning.model record by querying it with criteria.
func (c *Client) FindDataCleaningModel(criteria *Criteria) (*DataCleaningModel, error) {
	dms := &DataCleaningModels{}
	if err := c.SearchRead(DataCleaningModelModel, criteria, NewOptions().Limit(1), dms); err != nil {
		return nil, err
	}
	if dms != nil && len(*dms) > 0 {
		return &((*dms)[0]), nil
	}
	return nil, fmt.Errorf("data_cleaning.model was not found with criteria %v", criteria)
}

// FindDataCleaningModels finds data_cleaning.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataCleaningModels(criteria *Criteria, options *Options) (*DataCleaningModels, error) {
	dms := &DataCleaningModels{}
	if err := c.SearchRead(DataCleaningModelModel, criteria, options, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDataCleaningModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataCleaningModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataCleaningModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataCleaningModelId finds record id by querying it with criteria.
func (c *Client) FindDataCleaningModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataCleaningModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_cleaning.model was not found with criteria %v and options %v", criteria, options)
}
