package odoo

import (
	"fmt"
)

// DataMergeGroup represents data_merge.group model.
type DataMergeGroup struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Active          *Bool     `xmlrpc:"active,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	DivergentFields *String   `xmlrpc:"divergent_fields,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	ModelId         *Many2One `xmlrpc:"model_id,omptempty"`
	RecordIds       *Relation `xmlrpc:"record_ids,omptempty"`
	ResModelId      *Many2One `xmlrpc:"res_model_id,omptempty"`
	ResModelName    *String   `xmlrpc:"res_model_name,omptempty"`
	Similarity      *Float    `xmlrpc:"similarity,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DataMergeGroups represents array of data_merge.group model.
type DataMergeGroups []DataMergeGroup

// DataMergeGroupModel is the odoo model name.
const DataMergeGroupModel = "data_merge.group"

// Many2One convert DataMergeGroup to *Many2One.
func (dg *DataMergeGroup) Many2One() *Many2One {
	return NewMany2One(dg.Id.Get(), "")
}

// CreateDataMergeGroup creates a new data_merge.group model and returns its id.
func (c *Client) CreateDataMergeGroup(dg *DataMergeGroup) (int64, error) {
	ids, err := c.CreateDataMergeGroups([]*DataMergeGroup{dg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDataMergeGroup creates a new data_merge.group model and returns its id.
func (c *Client) CreateDataMergeGroups(dgs []*DataMergeGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range dgs {
		vv = append(vv, v)
	}
	return c.Create(DataMergeGroupModel, vv)
}

// UpdateDataMergeGroup updates an existing data_merge.group record.
func (c *Client) UpdateDataMergeGroup(dg *DataMergeGroup) error {
	return c.UpdateDataMergeGroups([]int64{dg.Id.Get()}, dg)
}

// UpdateDataMergeGroups updates existing data_merge.group records.
// All records (represented by ids) will be updated by dg values.
func (c *Client) UpdateDataMergeGroups(ids []int64, dg *DataMergeGroup) error {
	return c.Update(DataMergeGroupModel, ids, dg)
}

// DeleteDataMergeGroup deletes an existing data_merge.group record.
func (c *Client) DeleteDataMergeGroup(id int64) error {
	return c.DeleteDataMergeGroups([]int64{id})
}

// DeleteDataMergeGroups deletes existing data_merge.group records.
func (c *Client) DeleteDataMergeGroups(ids []int64) error {
	return c.Delete(DataMergeGroupModel, ids)
}

// GetDataMergeGroup gets data_merge.group existing record.
func (c *Client) GetDataMergeGroup(id int64) (*DataMergeGroup, error) {
	dgs, err := c.GetDataMergeGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	if dgs != nil && len(*dgs) > 0 {
		return &((*dgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of data_merge.group not found", id)
}

// GetDataMergeGroups gets data_merge.group existing records.
func (c *Client) GetDataMergeGroups(ids []int64) (*DataMergeGroups, error) {
	dgs := &DataMergeGroups{}
	if err := c.Read(DataMergeGroupModel, ids, nil, dgs); err != nil {
		return nil, err
	}
	return dgs, nil
}

// FindDataMergeGroup finds data_merge.group record by querying it with criteria.
func (c *Client) FindDataMergeGroup(criteria *Criteria) (*DataMergeGroup, error) {
	dgs := &DataMergeGroups{}
	if err := c.SearchRead(DataMergeGroupModel, criteria, NewOptions().Limit(1), dgs); err != nil {
		return nil, err
	}
	if dgs != nil && len(*dgs) > 0 {
		return &((*dgs)[0]), nil
	}
	return nil, fmt.Errorf("data_merge.group was not found with criteria %v", criteria)
}

// FindDataMergeGroups finds data_merge.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeGroups(criteria *Criteria, options *Options) (*DataMergeGroups, error) {
	dgs := &DataMergeGroups{}
	if err := c.SearchRead(DataMergeGroupModel, criteria, options, dgs); err != nil {
		return nil, err
	}
	return dgs, nil
}

// FindDataMergeGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDataMergeGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DataMergeGroupModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDataMergeGroupId finds record id by querying it with criteria.
func (c *Client) FindDataMergeGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DataMergeGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("data_merge.group was not found with criteria %v and options %v", criteria, options)
}
