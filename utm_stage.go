package odoo

import (
	"fmt"
)

// UtmStage represents utm.stage model.
type UtmStage struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// UtmStages represents array of utm.stage model.
type UtmStages []UtmStage

// UtmStageModel is the odoo model name.
const UtmStageModel = "utm.stage"

// Many2One convert UtmStage to *Many2One.
func (us *UtmStage) Many2One() *Many2One {
	return NewMany2One(us.Id.Get(), "")
}

// CreateUtmStage creates a new utm.stage model and returns its id.
func (c *Client) CreateUtmStage(us *UtmStage) (int64, error) {
	ids, err := c.CreateUtmStages([]*UtmStage{us})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateUtmStage creates a new utm.stage model and returns its id.
func (c *Client) CreateUtmStages(uss []*UtmStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range uss {
		vv = append(vv, v)
	}
	return c.Create(UtmStageModel, vv)
}

// UpdateUtmStage updates an existing utm.stage record.
func (c *Client) UpdateUtmStage(us *UtmStage) error {
	return c.UpdateUtmStages([]int64{us.Id.Get()}, us)
}

// UpdateUtmStages updates existing utm.stage records.
// All records (represented by ids) will be updated by us values.
func (c *Client) UpdateUtmStages(ids []int64, us *UtmStage) error {
	return c.Update(UtmStageModel, ids, us)
}

// DeleteUtmStage deletes an existing utm.stage record.
func (c *Client) DeleteUtmStage(id int64) error {
	return c.DeleteUtmStages([]int64{id})
}

// DeleteUtmStages deletes existing utm.stage records.
func (c *Client) DeleteUtmStages(ids []int64) error {
	return c.Delete(UtmStageModel, ids)
}

// GetUtmStage gets utm.stage existing record.
func (c *Client) GetUtmStage(id int64) (*UtmStage, error) {
	uss, err := c.GetUtmStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if uss != nil && len(*uss) > 0 {
		return &((*uss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of utm.stage not found", id)
}

// GetUtmStages gets utm.stage existing records.
func (c *Client) GetUtmStages(ids []int64) (*UtmStages, error) {
	uss := &UtmStages{}
	if err := c.Read(UtmStageModel, ids, nil, uss); err != nil {
		return nil, err
	}
	return uss, nil
}

// FindUtmStage finds utm.stage record by querying it with criteria.
func (c *Client) FindUtmStage(criteria *Criteria) (*UtmStage, error) {
	uss := &UtmStages{}
	if err := c.SearchRead(UtmStageModel, criteria, NewOptions().Limit(1), uss); err != nil {
		return nil, err
	}
	if uss != nil && len(*uss) > 0 {
		return &((*uss)[0]), nil
	}
	return nil, fmt.Errorf("utm.stage was not found with criteria %v", criteria)
}

// FindUtmStages finds utm.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmStages(criteria *Criteria, options *Options) (*UtmStages, error) {
	uss := &UtmStages{}
	if err := c.SearchRead(UtmStageModel, criteria, options, uss); err != nil {
		return nil, err
	}
	return uss, nil
}

// FindUtmStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindUtmStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(UtmStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindUtmStageId finds record id by querying it with criteria.
func (c *Client) FindUtmStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(UtmStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("utm.stage was not found with criteria %v and options %v", criteria, options)
}
