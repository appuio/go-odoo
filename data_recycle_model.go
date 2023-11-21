package odoo

import (
	"fmt"
)

// DataRecycleModel represents data_recycle.model model.
type DataRecycleModel struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	Active                *Bool      `xmlrpc:"active,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	Domain                *String    `xmlrpc:"domain,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	IncludeArchived       *Bool      `xmlrpc:"include_archived,omptempty"`
	LastNotification      *Time      `xmlrpc:"last_notification,omptempty"`
	Name                  *String    `xmlrpc:"name,omptempty"`
	NotifyFrequency       *Int       `xmlrpc:"notify_frequency,omptempty"`
	NotifyFrequencyPeriod *Selection `xmlrpc:"notify_frequency_period,omptempty"`
	NotifyUserIds         *Relation  `xmlrpc:"notify_user_ids,omptempty"`
	RecordsToRecycleCount *Int       `xmlrpc:"records_to_recycle_count,omptempty"`
	RecycleAction         *Selection `xmlrpc:"recycle_action,omptempty"`
	RecycleMode           *Selection `xmlrpc:"recycle_mode,omptempty"`
	RecycleRecordIds      *Relation  `xmlrpc:"recycle_record_ids,omptempty"`
	ResModelId            *Many2One  `xmlrpc:"res_model_id,omptempty"`
	ResModelName          *String    `xmlrpc:"res_model_name,omptempty"`
	TimeFieldDelta        *Int       `xmlrpc:"time_field_delta,omptempty"`
	TimeFieldDeltaUnit    *Selection `xmlrpc:"time_field_delta_unit,omptempty"`
	TimeFieldId           *Many2One  `xmlrpc:"time_field_id,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DataRecycleModels represents array of data_recycle.model model.
type DataRecycleModels []DataRecycleModel

// DataRecycleModelModel is the odoo model name.
const DataRecycleModelModel = "data_recycle.model"

// Many2One convert DataRecycleModel to *Many2One.
func (dm *DataRecycleModel) Many2One() *Many2One {
	return NewMany2One(dm.Id.Get(), "")
}

// CreateDataRecycleModel creates a new data_recycle.model model and returns its id.
func (c *Client) CreateDataRecycleModel(dm *DataRecycleModel) (int64, error) {
	ids, err := c.CreateDataRecycleModels([]*DataRecycleModel{dm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataRecycleModel creates a new data_recycle.model model and returns its id.
func (c *Client) CreateDataRecycleModels(dms []*DataRecycleModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range dms {
		vv = append(vv, v)
	}
	return c.Create(DataRecycleModelModel, vv)
}

// UpdateDataRecycleModel updates an existing data_recycle.model record.
func (c *Client) UpdateDataRecycleModel(dm *DataRecycleModel) error {
	return c.UpdateDataRecycleModels([]int64{dm.Id.Get()}, dm)
}

// UpdateDataRecycleModels updates existing data_recycle.model records.
// All records (represented by ids) will be updated by dm values.
func (c *Client) UpdateDataRecycleModels(ids []int64, dm *DataRecycleModel) error {
	return c.Update(DataRecycleModelModel, ids, dm)
}

// DeleteDataRecycleModel deletes an existing data_recycle.model record.
func (c *Client) DeleteDataRecycleModel(id int64) error {
	return c.DeleteDataRecycleModels([]int64{id})
}

// DeleteDataRecycleModels deletes existing data_recycle.model records.
func (c *Client) DeleteDataRecycleModels(ids []int64) error {
	return c.Delete(DataRecycleModelModel, ids)
}

// GetDataRecycleModel gets data_recycle.model existing record.
func (c *Client) GetDataRecycleModel(id int64) (*DataRecycleModel, error) {
	dms, err := c.GetDataRecycleModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if dms != nil && len(*dms) > 0 {
		return &((*dms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_recycle.model not found", id)
}

// GetDataRecycleModels gets data_recycle.model existing records.
func (c *Client) GetDataRecycleModels(ids []int64) (*DataRecycleModels, error) {
	dms := &DataRecycleModels{}
	if err := c.Read(DataRecycleModelModel, ids, nil, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDataRecycleModel finds data_recycle.model record by querying it with criteria.
func (c *Client) FindDataRecycleModel(criteria *Criteria) (*DataRecycleModel, error) {
	dms := &DataRecycleModels{}
	if err := c.SearchRead(DataRecycleModelModel, criteria, NewOptions().Limit(1), dms); err != nil {
		return nil, err
	}
	if dms != nil && len(*dms) > 0 {
		return &((*dms)[0]), nil
	}
	return nil, fmt.Errorf("data_recycle.model was not found with criteria %v", criteria)
}

// FindDataRecycleModels finds data_recycle.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataRecycleModels(criteria *Criteria, options *Options) (*DataRecycleModels, error) {
	dms := &DataRecycleModels{}
	if err := c.SearchRead(DataRecycleModelModel, criteria, options, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDataRecycleModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataRecycleModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataRecycleModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataRecycleModelId finds record id by querying it with criteria.
func (c *Client) FindDataRecycleModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataRecycleModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_recycle.model was not found with criteria %v and options %v", criteria, options)
}
