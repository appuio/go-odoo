package odoo

import (
	"fmt"
)

// DataRecycleRecord represents data_recycle.record model.
type DataRecycleRecord struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	Active         *Bool     `xmlrpc:"active,omptempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	RecycleModelId *Many2One `xmlrpc:"recycle_model_id,omptempty"`
	ResId          *Int      `xmlrpc:"res_id,omptempty"`
	ResModelId     *Many2One `xmlrpc:"res_model_id,omptempty"`
	ResModelName   *String   `xmlrpc:"res_model_name,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DataRecycleRecords represents array of data_recycle.record model.
type DataRecycleRecords []DataRecycleRecord

// DataRecycleRecordModel is the odoo model name.
const DataRecycleRecordModel = "data_recycle.record"

// Many2One convert DataRecycleRecord to *Many2One.
func (dr *DataRecycleRecord) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDataRecycleRecord creates a new data_recycle.record model and returns its id.
func (c *Client) CreateDataRecycleRecord(dr *DataRecycleRecord) (int64, error) {
	ids, err := c.CreateDataRecycleRecords([]*DataRecycleRecord{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataRecycleRecord creates a new data_recycle.record model and returns its id.
func (c *Client) CreateDataRecycleRecords(drs []*DataRecycleRecord) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DataRecycleRecordModel, vv)
}

// UpdateDataRecycleRecord updates an existing data_recycle.record record.
func (c *Client) UpdateDataRecycleRecord(dr *DataRecycleRecord) error {
	return c.UpdateDataRecycleRecords([]int64{dr.Id.Get()}, dr)
}

// UpdateDataRecycleRecords updates existing data_recycle.record records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDataRecycleRecords(ids []int64, dr *DataRecycleRecord) error {
	return c.Update(DataRecycleRecordModel, ids, dr)
}

// DeleteDataRecycleRecord deletes an existing data_recycle.record record.
func (c *Client) DeleteDataRecycleRecord(id int64) error {
	return c.DeleteDataRecycleRecords([]int64{id})
}

// DeleteDataRecycleRecords deletes existing data_recycle.record records.
func (c *Client) DeleteDataRecycleRecords(ids []int64) error {
	return c.Delete(DataRecycleRecordModel, ids)
}

// GetDataRecycleRecord gets data_recycle.record existing record.
func (c *Client) GetDataRecycleRecord(id int64) (*DataRecycleRecord, error) {
	drs, err := c.GetDataRecycleRecords([]int64{id})
	if err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_recycle.record not found", id)
}

// GetDataRecycleRecords gets data_recycle.record existing records.
func (c *Client) GetDataRecycleRecords(ids []int64) (*DataRecycleRecords, error) {
	drs := &DataRecycleRecords{}
	if err := c.Read(DataRecycleRecordModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataRecycleRecord finds data_recycle.record record by querying it with criteria.
func (c *Client) FindDataRecycleRecord(criteria *Criteria) (*DataRecycleRecord, error) {
	drs := &DataRecycleRecords{}
	if err := c.SearchRead(DataRecycleRecordModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("data_recycle.record was not found with criteria %v", criteria)
}

// FindDataRecycleRecords finds data_recycle.record records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataRecycleRecords(criteria *Criteria, options *Options) (*DataRecycleRecords, error) {
	drs := &DataRecycleRecords{}
	if err := c.SearchRead(DataRecycleRecordModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDataRecycleRecordIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataRecycleRecordIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataRecycleRecordModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataRecycleRecordId finds record id by querying it with criteria.
func (c *Client) FindDataRecycleRecordId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataRecycleRecordModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_recycle.record was not found with criteria %v and options %v", criteria, options)
}
