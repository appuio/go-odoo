package odoo

import (
	"fmt"
)

// DataMergeRecord represents data_merge.record model.
type DataMergeRecord struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	Active           *Bool     `xmlrpc:"active,omptempty"`
	CompanyId        *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	Differences      *String   `xmlrpc:"differences,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	FieldValues      *String   `xmlrpc:"field_values,omptempty"`
	GroupId          *Many2One `xmlrpc:"group_id,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	IsDeleted        *Bool     `xmlrpc:"is_deleted,omptempty"`
	IsDiscarded      *Bool     `xmlrpc:"is_discarded,omptempty"`
	IsMaster         *Bool     `xmlrpc:"is_master,omptempty"`
	ModelId          *Many2One `xmlrpc:"model_id,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	RecordCreateDate *Time     `xmlrpc:"record_create_date,omptempty"`
	RecordCreateUid  *String   `xmlrpc:"record_create_uid,omptempty"`
	RecordWriteDate  *Time     `xmlrpc:"record_write_date,omptempty"`
	RecordWriteUid   *String   `xmlrpc:"record_write_uid,omptempty"`
	ResId            *Int      `xmlrpc:"res_id,omptempty"`
	ResModelId       *Many2One `xmlrpc:"res_model_id,omptempty"`
	ResModelName     *String   `xmlrpc:"res_model_name,omptempty"`
	UsedIn           *String   `xmlrpc:"used_in,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DataMergeRecords represents array of data_merge.record model.
type DataMergeRecords []DataMergeRecord

// DataMergeRecordModel is the odoo model name.
const DataMergeRecordModel = "data_merge.record"

// Many2One convert DataMergeRecord to *Many2One.
func (dr *DataMergeRecord) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDataMergeRecord creates a new data_merge.record model and returns its id.
func (c *Client) CreateDataMergeRecord(dr *DataMergeRecord) (int64, error) {
	ids, err := c.CreateDataMergeRecords([]*DataMergeRecord{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataMergeRecord creates a new data_merge.record model and returns its id.
func (c *Client) CreateDataMergeRecords(drs []*DataMergeRecord) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DataMergeRecordModel, vv)
}

// UpdateDataMergeRecord updates an existing data_merge.record record.
func (c *Client) UpdateDataMergeRecord(dr *DataMergeRecord) error {
	return c.UpdateDataMergeRecords([]int64{dr.Id.Get()}, dr)
}

// UpdateDataMergeRecords updates existing data_merge.record records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDataMergeRecords(ids []int64, dr *DataMergeRecord) error {
	return c.Update(DataMergeRecordModel, ids, dr)
}

// DeleteDataMergeRecord deletes an existing data_merge.record record.
func (c *Client) DeleteDataMergeRecord(id int64) error {
	return c.DeleteDataMergeRecords([]int64{id})
}

// DeleteDataMergeRecords deletes existing data_merge.record records.
func (c *Client) DeleteDataMergeRecords(ids []int64) error {
	return c.Delete(DataMergeRecordModel, ids)
}

// GetDataMergeRecord gets data_merge.record existing record.
func (c *Client) GetDataMergeRecord(id int64) (*DataMergeRecord, error) {
	drs, err := c.GetDataMergeRecords([]int64{id})
	if err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_merge.record not found", id)
}

// GetDataMergeRecords gets data_merge.record existing records.
func (c *Client) GetDataMergeRecords(ids []int64) (*DataMergeRecords, error) {
	drs := &DataMergeRecords{}
	if err := c.Read(DataMergeRecordModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataMergeRecord finds data_merge.record record by querying it with criteria.
func (c *Client) FindDataMergeRecord(criteria *Criteria) (*DataMergeRecord, error) {
	drs := &DataMergeRecords{}
	if err := c.SearchRead(DataMergeRecordModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("data_merge.record was not found with criteria %v", criteria)
}

// FindDataMergeRecords finds data_merge.record records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeRecords(criteria *Criteria, options *Options) (*DataMergeRecords, error) {
	drs := &DataMergeRecords{}
	if err := c.SearchRead(DataMergeRecordModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataMergeRecordIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeRecordIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataMergeRecordModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataMergeRecordId finds record id by querying it with criteria.
func (c *Client) FindDataMergeRecordId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataMergeRecordModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_merge.record was not found with criteria %v and options %v", criteria, options)
}
