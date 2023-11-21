package odoo

import (
	"fmt"
)

// GlobalSearchFields represents global.search.fields model.
type GlobalSearchFields struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	FieldId        *Many2One  `xmlrpc:"field_id,omptempty"`
	FieldIds       *Relation  `xmlrpc:"field_ids,omptempty"`
	GlobalSearchId *Many2One  `xmlrpc:"global_search_id,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	ModelId        *Many2One  `xmlrpc:"model_id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	RelatedModelId *String    `xmlrpc:"related_model_id,omptempty"`
	Sequence       *Int       `xmlrpc:"sequence,omptempty"`
	Ttype          *Selection `xmlrpc:"ttype,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// GlobalSearchFieldss represents array of global.search.fields model.
type GlobalSearchFieldss []GlobalSearchFields

// GlobalSearchFieldsModel is the odoo model name.
const GlobalSearchFieldsModel = "global.search.fields"

// Many2One convert GlobalSearchFields to *Many2One.
func (gsf *GlobalSearchFields) Many2One() *Many2One {
	return NewMany2One(gsf.Id.Get(), "")
}

// CreateGlobalSearchFields creates a new global.search.fields model and returns its id.
func (c *Client) CreateGlobalSearchFields(gsf *GlobalSearchFields) (int64, error) {
	ids, err := c.CreateGlobalSearchFieldss([]*GlobalSearchFields{gsf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGlobalSearchFields creates a new global.search.fields model and returns its id.
func (c *Client) CreateGlobalSearchFieldss(gsfs []*GlobalSearchFields) ([]int64, error) {
	var vv []interface{}
	for _, v := range gsfs {
		vv = append(vv, v)
	}
	return c.Create(GlobalSearchFieldsModel, vv)
}

// UpdateGlobalSearchFields updates an existing global.search.fields record.
func (c *Client) UpdateGlobalSearchFields(gsf *GlobalSearchFields) error {
	return c.UpdateGlobalSearchFieldss([]int64{gsf.Id.Get()}, gsf)
}

// UpdateGlobalSearchFieldss updates existing global.search.fields records.
// All records (represented by ids) will be updated by gsf values.
func (c *Client) UpdateGlobalSearchFieldss(ids []int64, gsf *GlobalSearchFields) error {
	return c.Update(GlobalSearchFieldsModel, ids, gsf)
}

// DeleteGlobalSearchFields deletes an existing global.search.fields record.
func (c *Client) DeleteGlobalSearchFields(id int64) error {
	return c.DeleteGlobalSearchFieldss([]int64{id})
}

// DeleteGlobalSearchFieldss deletes existing global.search.fields records.
func (c *Client) DeleteGlobalSearchFieldss(ids []int64) error {
	return c.Delete(GlobalSearchFieldsModel, ids)
}

// GetGlobalSearchFields gets global.search.fields existing record.
func (c *Client) GetGlobalSearchFields(id int64) (*GlobalSearchFields, error) {
	gsfs, err := c.GetGlobalSearchFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	if gsfs != nil && len(*gsfs) > 0 {
		return &((*gsfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of global.search.fields not found", id)
}

// GetGlobalSearchFieldss gets global.search.fields existing records.
func (c *Client) GetGlobalSearchFieldss(ids []int64) (*GlobalSearchFieldss, error) {
	gsfs := &GlobalSearchFieldss{}
	if err := c.Read(GlobalSearchFieldsModel, ids, nil, gsfs); err != nil {
		return nil, err
	}
	return gsfs, nil
}

// FindGlobalSearchFields finds global.search.fields record by querying it with criteria.
func (c *Client) FindGlobalSearchFields(criteria *Criteria) (*GlobalSearchFields, error) {
	gsfs := &GlobalSearchFieldss{}
	if err := c.SearchRead(GlobalSearchFieldsModel, criteria, NewOptions().Limit(1), gsfs); err != nil {
		return nil, err
	}
	if gsfs != nil && len(*gsfs) > 0 {
		return &((*gsfs)[0]), nil
	}
	return nil, fmt.Errorf("global.search.fields was not found with criteria %v", criteria)
}

// FindGlobalSearchFieldss finds global.search.fields records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGlobalSearchFieldss(criteria *Criteria, options *Options) (*GlobalSearchFieldss, error) {
	gsfs := &GlobalSearchFieldss{}
	if err := c.SearchRead(GlobalSearchFieldsModel, criteria, options, gsfs); err != nil {
		return nil, err
	}
	return gsfs, nil
}

// FindGlobalSearchFieldsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGlobalSearchFieldsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GlobalSearchFieldsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGlobalSearchFieldsId finds record id by querying it with criteria.
func (c *Client) FindGlobalSearchFieldsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GlobalSearchFieldsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("global.search.fields was not found with criteria %v and options %v", criteria, options)
}
