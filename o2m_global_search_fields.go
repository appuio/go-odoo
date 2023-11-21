package odoo

import (
	"fmt"
)

// O2MGlobalSearchFields represents o2m.global.search.fields model.
type O2MGlobalSearchFields struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	FieldId           *Many2One  `xmlrpc:"field_id,omptempty"`
	GlobalO2MSearchId *Many2One  `xmlrpc:"global_o2m_search_id,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	ModelId           *Many2One  `xmlrpc:"model_id,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	RelatedModelId    *String    `xmlrpc:"related_model_id,omptempty"`
	Sequence          *Int       `xmlrpc:"sequence,omptempty"`
	Ttype             *Selection `xmlrpc:"ttype,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// O2MGlobalSearchFieldss represents array of o2m.global.search.fields model.
type O2MGlobalSearchFieldss []O2MGlobalSearchFields

// O2MGlobalSearchFieldsModel is the odoo model name.
const O2MGlobalSearchFieldsModel = "o2m.global.search.fields"

// Many2One convert O2MGlobalSearchFields to *Many2One.
func (ogsf *O2MGlobalSearchFields) Many2One() *Many2One {
	return NewMany2One(ogsf.Id.Get(), "")
}

// CreateO2MGlobalSearchFields creates a new o2m.global.search.fields model and returns its id.
func (c *Client) CreateO2MGlobalSearchFields(ogsf *O2MGlobalSearchFields) (int64, error) {
	ids, err := c.CreateO2MGlobalSearchFieldss([]*O2MGlobalSearchFields{ogsf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateO2MGlobalSearchFields creates a new o2m.global.search.fields model and returns its id.
func (c *Client) CreateO2MGlobalSearchFieldss(ogsfs []*O2MGlobalSearchFields) ([]int64, error) {
	var vv []interface{}
	for _, v := range ogsfs {
		vv = append(vv, v)
	}
	return c.Create(O2MGlobalSearchFieldsModel, vv)
}

// UpdateO2MGlobalSearchFields updates an existing o2m.global.search.fields record.
func (c *Client) UpdateO2MGlobalSearchFields(ogsf *O2MGlobalSearchFields) error {
	return c.UpdateO2MGlobalSearchFieldss([]int64{ogsf.Id.Get()}, ogsf)
}

// UpdateO2MGlobalSearchFieldss updates existing o2m.global.search.fields records.
// All records (represented by ids) will be updated by ogsf values.
func (c *Client) UpdateO2MGlobalSearchFieldss(ids []int64, ogsf *O2MGlobalSearchFields) error {
	return c.Update(O2MGlobalSearchFieldsModel, ids, ogsf)
}

// DeleteO2MGlobalSearchFields deletes an existing o2m.global.search.fields record.
func (c *Client) DeleteO2MGlobalSearchFields(id int64) error {
	return c.DeleteO2MGlobalSearchFieldss([]int64{id})
}

// DeleteO2MGlobalSearchFieldss deletes existing o2m.global.search.fields records.
func (c *Client) DeleteO2MGlobalSearchFieldss(ids []int64) error {
	return c.Delete(O2MGlobalSearchFieldsModel, ids)
}

// GetO2MGlobalSearchFields gets o2m.global.search.fields existing record.
func (c *Client) GetO2MGlobalSearchFields(id int64) (*O2MGlobalSearchFields, error) {
	ogsfs, err := c.GetO2MGlobalSearchFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	if ogsfs != nil && len(*ogsfs) > 0 {
		return &((*ogsfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of o2m.global.search.fields not found", id)
}

// GetO2MGlobalSearchFieldss gets o2m.global.search.fields existing records.
func (c *Client) GetO2MGlobalSearchFieldss(ids []int64) (*O2MGlobalSearchFieldss, error) {
	ogsfs := &O2MGlobalSearchFieldss{}
	if err := c.Read(O2MGlobalSearchFieldsModel, ids, nil, ogsfs); err != nil {
		return nil, err
	}
	return ogsfs, nil
}

// FindO2MGlobalSearchFields finds o2m.global.search.fields record by querying it with criteria.
func (c *Client) FindO2MGlobalSearchFields(criteria *Criteria) (*O2MGlobalSearchFields, error) {
	ogsfs := &O2MGlobalSearchFieldss{}
	if err := c.SearchRead(O2MGlobalSearchFieldsModel, criteria, NewOptions().Limit(1), ogsfs); err != nil {
		return nil, err
	}
	if ogsfs != nil && len(*ogsfs) > 0 {
		return &((*ogsfs)[0]), nil
	}
	return nil, fmt.Errorf("o2m.global.search.fields was not found with criteria %v", criteria)
}

// FindO2MGlobalSearchFieldss finds o2m.global.search.fields records by querying it
// and filtering it with criteria and options.
func (c *Client) FindO2MGlobalSearchFieldss(criteria *Criteria, options *Options) (*O2MGlobalSearchFieldss, error) {
	ogsfs := &O2MGlobalSearchFieldss{}
	if err := c.SearchRead(O2MGlobalSearchFieldsModel, criteria, options, ogsfs); err != nil {
		return nil, err
	}
	return ogsfs, nil
}

// FindO2MGlobalSearchFieldsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindO2MGlobalSearchFieldsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(O2MGlobalSearchFieldsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindO2MGlobalSearchFieldsId finds record id by querying it with criteria.
func (c *Client) FindO2MGlobalSearchFieldsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(O2MGlobalSearchFieldsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("o2m.global.search.fields was not found with criteria %v and options %v", criteria, options)
}
