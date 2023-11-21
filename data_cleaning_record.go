package odoo

import (
	"fmt"
)

// DataCleaningRecord represents data_cleaning.record model.
type DataCleaningRecord struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	Action                *String   `xmlrpc:"action,omptempty"`
	Active                *Bool     `xmlrpc:"active,omptempty"`
	CleaningModelId       *Many2One `xmlrpc:"cleaning_model_id,omptempty"`
	CompanyId             *Many2One `xmlrpc:"company_id,omptempty"`
	CountryId             *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrentValue          *String   `xmlrpc:"current_value,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	FieldId               *Many2One `xmlrpc:"field_id,omptempty"`
	FieldName             *String   `xmlrpc:"field_name,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	Name                  *String   `xmlrpc:"name,omptempty"`
	ResId                 *Int      `xmlrpc:"res_id,omptempty"`
	ResModelId            *Many2One `xmlrpc:"res_model_id,omptempty"`
	ResModelName          *String   `xmlrpc:"res_model_name,omptempty"`
	RuleIds               *Relation `xmlrpc:"rule_ids,omptempty"`
	SuggestedValue        *String   `xmlrpc:"suggested_value,omptempty"`
	SuggestedValueDisplay *String   `xmlrpc:"suggested_value_display,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DataCleaningRecords represents array of data_cleaning.record model.
type DataCleaningRecords []DataCleaningRecord

// DataCleaningRecordModel is the odoo model name.
const DataCleaningRecordModel = "data_cleaning.record"

// Many2One convert DataCleaningRecord to *Many2One.
func (dr *DataCleaningRecord) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDataCleaningRecord creates a new data_cleaning.record model and returns its id.
func (c *Client) CreateDataCleaningRecord(dr *DataCleaningRecord) (int64, error) {
	ids, err := c.CreateDataCleaningRecords([]*DataCleaningRecord{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataCleaningRecord creates a new data_cleaning.record model and returns its id.
func (c *Client) CreateDataCleaningRecords(drs []*DataCleaningRecord) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DataCleaningRecordModel, vv)
}

// UpdateDataCleaningRecord updates an existing data_cleaning.record record.
func (c *Client) UpdateDataCleaningRecord(dr *DataCleaningRecord) error {
	return c.UpdateDataCleaningRecords([]int64{dr.Id.Get()}, dr)
}

// UpdateDataCleaningRecords updates existing data_cleaning.record records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDataCleaningRecords(ids []int64, dr *DataCleaningRecord) error {
	return c.Update(DataCleaningRecordModel, ids, dr)
}

// DeleteDataCleaningRecord deletes an existing data_cleaning.record record.
func (c *Client) DeleteDataCleaningRecord(id int64) error {
	return c.DeleteDataCleaningRecords([]int64{id})
}

// DeleteDataCleaningRecords deletes existing data_cleaning.record records.
func (c *Client) DeleteDataCleaningRecords(ids []int64) error {
	return c.Delete(DataCleaningRecordModel, ids)
}

// GetDataCleaningRecord gets data_cleaning.record existing record.
func (c *Client) GetDataCleaningRecord(id int64) (*DataCleaningRecord, error) {
	drs, err := c.GetDataCleaningRecords([]int64{id})
	if err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_cleaning.record not found", id)
}

// GetDataCleaningRecords gets data_cleaning.record existing records.
func (c *Client) GetDataCleaningRecords(ids []int64) (*DataCleaningRecords, error) {
	drs := &DataCleaningRecords{}
	if err := c.Read(DataCleaningRecordModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataCleaningRecord finds data_cleaning.record record by querying it with criteria.
func (c *Client) FindDataCleaningRecord(criteria *Criteria) (*DataCleaningRecord, error) {
	drs := &DataCleaningRecords{}
	if err := c.SearchRead(DataCleaningRecordModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("data_cleaning.record was not found with criteria %v", criteria)
}

// FindDataCleaningRecords finds data_cleaning.record records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataCleaningRecords(criteria *Criteria, options *Options) (*DataCleaningRecords, error) {
	drs := &DataCleaningRecords{}
	if err := c.SearchRead(DataCleaningRecordModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataCleaningRecordIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataCleaningRecordIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataCleaningRecordModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataCleaningRecordId finds record id by querying it with criteria.
func (c *Client) FindDataCleaningRecordId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataCleaningRecordModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_cleaning.record was not found with criteria %v and options %v", criteria, options)
}
