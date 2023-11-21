package odoo

import (
	"fmt"
)

// LotLabelLayout represents lot.label.layout model.
type LotLabelLayout struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	LabelQuantity *Selection `xmlrpc:"label_quantity,omptempty"`
	PickingIds    *Relation  `xmlrpc:"picking_ids,omptempty"`
	PrintFormat   *Selection `xmlrpc:"print_format,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LotLabelLayouts represents array of lot.label.layout model.
type LotLabelLayouts []LotLabelLayout

// LotLabelLayoutModel is the odoo model name.
const LotLabelLayoutModel = "lot.label.layout"

// Many2One convert LotLabelLayout to *Many2One.
func (lll *LotLabelLayout) Many2One() *Many2One {
	return NewMany2One(lll.Id.Get(), "")
}

// CreateLotLabelLayout creates a new lot.label.layout model and returns its id.
func (c *Client) CreateLotLabelLayout(lll *LotLabelLayout) (int64, error) {
	ids, err := c.CreateLotLabelLayouts([]*LotLabelLayout{lll})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLotLabelLayout creates a new lot.label.layout model and returns its id.
func (c *Client) CreateLotLabelLayouts(llls []*LotLabelLayout) ([]int64, error) {
	var vv []interface{}
	for _, v := range llls {
		vv = append(vv, v)
	}
	return c.Create(LotLabelLayoutModel, vv)
}

// UpdateLotLabelLayout updates an existing lot.label.layout record.
func (c *Client) UpdateLotLabelLayout(lll *LotLabelLayout) error {
	return c.UpdateLotLabelLayouts([]int64{lll.Id.Get()}, lll)
}

// UpdateLotLabelLayouts updates existing lot.label.layout records.
// All records (represented by ids) will be updated by lll values.
func (c *Client) UpdateLotLabelLayouts(ids []int64, lll *LotLabelLayout) error {
	return c.Update(LotLabelLayoutModel, ids, lll)
}

// DeleteLotLabelLayout deletes an existing lot.label.layout record.
func (c *Client) DeleteLotLabelLayout(id int64) error {
	return c.DeleteLotLabelLayouts([]int64{id})
}

// DeleteLotLabelLayouts deletes existing lot.label.layout records.
func (c *Client) DeleteLotLabelLayouts(ids []int64) error {
	return c.Delete(LotLabelLayoutModel, ids)
}

// GetLotLabelLayout gets lot.label.layout existing record.
func (c *Client) GetLotLabelLayout(id int64) (*LotLabelLayout, error) {
	llls, err := c.GetLotLabelLayouts([]int64{id})
	if err != nil {
		return nil, err
	}
	if llls != nil && len(*llls) > 0 {
		return &((*llls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of lot.label.layout not found", id)
}

// GetLotLabelLayouts gets lot.label.layout existing records.
func (c *Client) GetLotLabelLayouts(ids []int64) (*LotLabelLayouts, error) {
	llls := &LotLabelLayouts{}
	if err := c.Read(LotLabelLayoutModel, ids, nil, llls); err != nil {
		return nil, err
	}
	return llls, nil
}

// FindLotLabelLayout finds lot.label.layout record by querying it with criteria.
func (c *Client) FindLotLabelLayout(criteria *Criteria) (*LotLabelLayout, error) {
	llls := &LotLabelLayouts{}
	if err := c.SearchRead(LotLabelLayoutModel, criteria, NewOptions().Limit(1), llls); err != nil {
		return nil, err
	}
	if llls != nil && len(*llls) > 0 {
		return &((*llls)[0]), nil
	}
	return nil, fmt.Errorf("lot.label.layout was not found with criteria %v", criteria)
}

// FindLotLabelLayouts finds lot.label.layout records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLotLabelLayouts(criteria *Criteria, options *Options) (*LotLabelLayouts, error) {
	llls := &LotLabelLayouts{}
	if err := c.SearchRead(LotLabelLayoutModel, criteria, options, llls); err != nil {
		return nil, err
	}
	return llls, nil
}

// FindLotLabelLayoutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLotLabelLayoutIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LotLabelLayoutModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLotLabelLayoutId finds record id by querying it with criteria.
func (c *Client) FindLotLabelLayoutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LotLabelLayoutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("lot.label.layout was not found with criteria %v and options %v", criteria, options)
}
