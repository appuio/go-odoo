package odoo

import (
	"fmt"
)

// TrgmIndex represents trgm.index model.
type TrgmIndex struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	FieldId     *Many2One  `xmlrpc:"field_id,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	IndexName   *String    `xmlrpc:"index_name,omptempty"`
	IndexType   *Selection `xmlrpc:"index_type,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// TrgmIndexs represents array of trgm.index model.
type TrgmIndexs []TrgmIndex

// TrgmIndexModel is the odoo model name.
const TrgmIndexModel = "trgm.index"

// Many2One convert TrgmIndex to *Many2One.
func (ti *TrgmIndex) Many2One() *Many2One {
	return NewMany2One(ti.Id.Get(), "")
}

// CreateTrgmIndex creates a new trgm.index model and returns its id.
func (c *Client) CreateTrgmIndex(ti *TrgmIndex) (int64, error) {
	ids, err := c.CreateTrgmIndexs([]*TrgmIndex{ti})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTrgmIndex creates a new trgm.index model and returns its id.
func (c *Client) CreateTrgmIndexs(tis []*TrgmIndex) ([]int64, error) {
	var vv []interface{}
	for _, v := range tis {
		vv = append(vv, v)
	}
	return c.Create(TrgmIndexModel, vv)
}

// UpdateTrgmIndex updates an existing trgm.index record.
func (c *Client) UpdateTrgmIndex(ti *TrgmIndex) error {
	return c.UpdateTrgmIndexs([]int64{ti.Id.Get()}, ti)
}

// UpdateTrgmIndexs updates existing trgm.index records.
// All records (represented by ids) will be updated by ti values.
func (c *Client) UpdateTrgmIndexs(ids []int64, ti *TrgmIndex) error {
	return c.Update(TrgmIndexModel, ids, ti)
}

// DeleteTrgmIndex deletes an existing trgm.index record.
func (c *Client) DeleteTrgmIndex(id int64) error {
	return c.DeleteTrgmIndexs([]int64{id})
}

// DeleteTrgmIndexs deletes existing trgm.index records.
func (c *Client) DeleteTrgmIndexs(ids []int64) error {
	return c.Delete(TrgmIndexModel, ids)
}

// GetTrgmIndex gets trgm.index existing record.
func (c *Client) GetTrgmIndex(id int64) (*TrgmIndex, error) {
	tis, err := c.GetTrgmIndexs([]int64{id})
	if err != nil {
		return nil, err
	}
	if tis != nil && len(*tis) > 0 {
		return &((*tis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of trgm.index not found", id)
}

// GetTrgmIndexs gets trgm.index existing records.
func (c *Client) GetTrgmIndexs(ids []int64) (*TrgmIndexs, error) {
	tis := &TrgmIndexs{}
	if err := c.Read(TrgmIndexModel, ids, nil, tis); err != nil {
		return nil, err
	}
	return tis, nil
}

// FindTrgmIndex finds trgm.index record by querying it with criteria.
func (c *Client) FindTrgmIndex(criteria *Criteria) (*TrgmIndex, error) {
	tis := &TrgmIndexs{}
	if err := c.SearchRead(TrgmIndexModel, criteria, NewOptions().Limit(1), tis); err != nil {
		return nil, err
	}
	if tis != nil && len(*tis) > 0 {
		return &((*tis)[0]), nil
	}
	return nil, fmt.Errorf("trgm.index was not found with criteria %v", criteria)
}

// FindTrgmIndexs finds trgm.index records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTrgmIndexs(criteria *Criteria, options *Options) (*TrgmIndexs, error) {
	tis := &TrgmIndexs{}
	if err := c.SearchRead(TrgmIndexModel, criteria, options, tis); err != nil {
		return nil, err
	}
	return tis, nil
}

// FindTrgmIndexIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTrgmIndexIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(TrgmIndexModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTrgmIndexId finds record id by querying it with criteria.
func (c *Client) FindTrgmIndexId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TrgmIndexModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("trgm.index was not found with criteria %v and options %v", criteria, options)
}
