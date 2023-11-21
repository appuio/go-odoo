package odoo

import (
	"fmt"
)

// DataMergeModel represents data_merge.model model.
type DataMergeModel struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	Active                  *Bool      `xmlrpc:"active,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateThreshold         *Int       `xmlrpc:"create_threshold,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomMergeMethod       *Bool      `xmlrpc:"custom_merge_method,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	Domain                  *String    `xmlrpc:"domain,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	IsContextualMergeAction *Bool      `xmlrpc:"is_contextual_merge_action,omptempty"`
	LastNotification        *Time      `xmlrpc:"last_notification,omptempty"`
	MergeMode               *Selection `xmlrpc:"merge_mode,omptempty"`
	MergeThreshold          *Int       `xmlrpc:"merge_threshold,omptempty"`
	MixByCompany            *Bool      `xmlrpc:"mix_by_company,omptempty"`
	Name                    *String    `xmlrpc:"name,omptempty"`
	NotifyFrequency         *Int       `xmlrpc:"notify_frequency,omptempty"`
	NotifyFrequencyPeriod   *Selection `xmlrpc:"notify_frequency_period,omptempty"`
	NotifyUserIds           *Relation  `xmlrpc:"notify_user_ids,omptempty"`
	RecordsToMergeCount     *Int       `xmlrpc:"records_to_merge_count,omptempty"`
	RemovalMode             *Selection `xmlrpc:"removal_mode,omptempty"`
	ResModelId              *Many2One  `xmlrpc:"res_model_id,omptempty"`
	ResModelName            *String    `xmlrpc:"res_model_name,omptempty"`
	RuleIds                 *Relation  `xmlrpc:"rule_ids,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DataMergeModels represents array of data_merge.model model.
type DataMergeModels []DataMergeModel

// DataMergeModelModel is the odoo model name.
const DataMergeModelModel = "data_merge.model"

// Many2One convert DataMergeModel to *Many2One.
func (dm *DataMergeModel) Many2One() *Many2One {
	return NewMany2One(dm.Id.Get(), "")
}

// CreateDataMergeModel creates a new data_merge.model model and returns its id.
func (c *Client) CreateDataMergeModel(dm *DataMergeModel) (int64, error) {
	ids, err := c.CreateDataMergeModels([]*DataMergeModel{dm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataMergeModel creates a new data_merge.model model and returns its id.
func (c *Client) CreateDataMergeModels(dms []*DataMergeModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range dms {
		vv = append(vv, v)
	}
	return c.Create(DataMergeModelModel, vv)
}

// UpdateDataMergeModel updates an existing data_merge.model record.
func (c *Client) UpdateDataMergeModel(dm *DataMergeModel) error {
	return c.UpdateDataMergeModels([]int64{dm.Id.Get()}, dm)
}

// UpdateDataMergeModels updates existing data_merge.model records.
// All records (represented by ids) will be updated by dm values.
func (c *Client) UpdateDataMergeModels(ids []int64, dm *DataMergeModel) error {
	return c.Update(DataMergeModelModel, ids, dm)
}

// DeleteDataMergeModel deletes an existing data_merge.model record.
func (c *Client) DeleteDataMergeModel(id int64) error {
	return c.DeleteDataMergeModels([]int64{id})
}

// DeleteDataMergeModels deletes existing data_merge.model records.
func (c *Client) DeleteDataMergeModels(ids []int64) error {
	return c.Delete(DataMergeModelModel, ids)
}

// GetDataMergeModel gets data_merge.model existing record.
func (c *Client) GetDataMergeModel(id int64) (*DataMergeModel, error) {
	dms, err := c.GetDataMergeModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if dms != nil && len(*dms) > 0 {
		return &((*dms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_merge.model not found", id)
}

// GetDataMergeModels gets data_merge.model existing records.
func (c *Client) GetDataMergeModels(ids []int64) (*DataMergeModels, error) {
	dms := &DataMergeModels{}
	if err := c.Read(DataMergeModelModel, ids, nil, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDataMergeModel finds data_merge.model record by querying it with criteria.
func (c *Client) FindDataMergeModel(criteria *Criteria) (*DataMergeModel, error) {
	dms := &DataMergeModels{}
	if err := c.SearchRead(DataMergeModelModel, criteria, NewOptions().Limit(1), dms); err != nil {
		return nil, err
	}
	if dms != nil && len(*dms) > 0 {
		return &((*dms)[0]), nil
	}
	return nil, fmt.Errorf("data_merge.model was not found with criteria %v", criteria)
}

// FindDataMergeModels finds data_merge.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeModels(criteria *Criteria, options *Options) (*DataMergeModels, error) {
	dms := &DataMergeModels{}
	if err := c.SearchRead(DataMergeModelModel, criteria, options, dms); err != nil {
		return nil, err
	}
	return dms, nil
}

// FindDataMergeModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataMergeModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataMergeModelId finds record id by querying it with criteria.
func (c *Client) FindDataMergeModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataMergeModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_merge.model was not found with criteria %v and options %v", criteria, options)
}
