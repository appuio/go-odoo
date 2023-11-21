package odoo

import (
	"fmt"
)

// MaintenanceStage represents maintenance.stage model.
type MaintenanceStage struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Done        *Bool     `xmlrpc:"done,omptempty"`
	Fold        *Bool     `xmlrpc:"fold,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MaintenanceStages represents array of maintenance.stage model.
type MaintenanceStages []MaintenanceStage

// MaintenanceStageModel is the odoo model name.
const MaintenanceStageModel = "maintenance.stage"

// Many2One convert MaintenanceStage to *Many2One.
func (ms *MaintenanceStage) Many2One() *Many2One {
	return NewMany2One(ms.Id.Get(), "")
}

// CreateMaintenanceStage creates a new maintenance.stage model and returns its id.
func (c *Client) CreateMaintenanceStage(ms *MaintenanceStage) (int64, error) {
	ids, err := c.CreateMaintenanceStages([]*MaintenanceStage{ms})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMaintenanceStage creates a new maintenance.stage model and returns its id.
func (c *Client) CreateMaintenanceStages(mss []*MaintenanceStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range mss {
		vv = append(vv, v)
	}
	return c.Create(MaintenanceStageModel, vv)
}

// UpdateMaintenanceStage updates an existing maintenance.stage record.
func (c *Client) UpdateMaintenanceStage(ms *MaintenanceStage) error {
	return c.UpdateMaintenanceStages([]int64{ms.Id.Get()}, ms)
}

// UpdateMaintenanceStages updates existing maintenance.stage records.
// All records (represented by ids) will be updated by ms values.
func (c *Client) UpdateMaintenanceStages(ids []int64, ms *MaintenanceStage) error {
	return c.Update(MaintenanceStageModel, ids, ms)
}

// DeleteMaintenanceStage deletes an existing maintenance.stage record.
func (c *Client) DeleteMaintenanceStage(id int64) error {
	return c.DeleteMaintenanceStages([]int64{id})
}

// DeleteMaintenanceStages deletes existing maintenance.stage records.
func (c *Client) DeleteMaintenanceStages(ids []int64) error {
	return c.Delete(MaintenanceStageModel, ids)
}

// GetMaintenanceStage gets maintenance.stage existing record.
func (c *Client) GetMaintenanceStage(id int64) (*MaintenanceStage, error) {
	mss, err := c.GetMaintenanceStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if mss != nil && len(*mss) > 0 {
		return &((*mss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of maintenance.stage not found", id)
}

// GetMaintenanceStages gets maintenance.stage existing records.
func (c *Client) GetMaintenanceStages(ids []int64) (*MaintenanceStages, error) {
	mss := &MaintenanceStages{}
	if err := c.Read(MaintenanceStageModel, ids, nil, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMaintenanceStage finds maintenance.stage record by querying it with criteria.
func (c *Client) FindMaintenanceStage(criteria *Criteria) (*MaintenanceStage, error) {
	mss := &MaintenanceStages{}
	if err := c.SearchRead(MaintenanceStageModel, criteria, NewOptions().Limit(1), mss); err != nil {
		return nil, err
	}
	if mss != nil && len(*mss) > 0 {
		return &((*mss)[0]), nil
	}
	return nil, fmt.Errorf("maintenance.stage was not found with criteria %v", criteria)
}

// FindMaintenanceStages finds maintenance.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceStages(criteria *Criteria, options *Options) (*MaintenanceStages, error) {
	mss := &MaintenanceStages{}
	if err := c.SearchRead(MaintenanceStageModel, criteria, options, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMaintenanceStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MaintenanceStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMaintenanceStageId finds record id by querying it with criteria.
func (c *Client) FindMaintenanceStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MaintenanceStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("maintenance.stage was not found with criteria %v and options %v", criteria, options)
}
