package odoo

import (
	"fmt"
)

// GlobalSearch represents global.search model.
type GlobalSearch struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	FieldIds       *Relation `xmlrpc:"field_ids,omptempty"`
	GlobalFieldIds *Relation `xmlrpc:"global_field_ids,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	MainFieldId    *Many2One `xmlrpc:"main_field_id,omptempty"`
	ModelId        *Many2One `xmlrpc:"model_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// GlobalSearchs represents array of global.search model.
type GlobalSearchs []GlobalSearch

// GlobalSearchModel is the odoo model name.
const GlobalSearchModel = "global.search"

// Many2One convert GlobalSearch to *Many2One.
func (gs *GlobalSearch) Many2One() *Many2One {
	return NewMany2One(gs.Id.Get(), "")
}

// CreateGlobalSearch creates a new global.search model and returns its id.
func (c *Client) CreateGlobalSearch(gs *GlobalSearch) (int64, error) {
	ids, err := c.CreateGlobalSearchs([]*GlobalSearch{gs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGlobalSearch creates a new global.search model and returns its id.
func (c *Client) CreateGlobalSearchs(gss []*GlobalSearch) ([]int64, error) {
	var vv []interface{}
	for _, v := range gss {
		vv = append(vv, v)
	}
	return c.Create(GlobalSearchModel, vv)
}

// UpdateGlobalSearch updates an existing global.search record.
func (c *Client) UpdateGlobalSearch(gs *GlobalSearch) error {
	return c.UpdateGlobalSearchs([]int64{gs.Id.Get()}, gs)
}

// UpdateGlobalSearchs updates existing global.search records.
// All records (represented by ids) will be updated by gs values.
func (c *Client) UpdateGlobalSearchs(ids []int64, gs *GlobalSearch) error {
	return c.Update(GlobalSearchModel, ids, gs)
}

// DeleteGlobalSearch deletes an existing global.search record.
func (c *Client) DeleteGlobalSearch(id int64) error {
	return c.DeleteGlobalSearchs([]int64{id})
}

// DeleteGlobalSearchs deletes existing global.search records.
func (c *Client) DeleteGlobalSearchs(ids []int64) error {
	return c.Delete(GlobalSearchModel, ids)
}

// GetGlobalSearch gets global.search existing record.
func (c *Client) GetGlobalSearch(id int64) (*GlobalSearch, error) {
	gss, err := c.GetGlobalSearchs([]int64{id})
	if err != nil {
		return nil, err
	}
	if gss != nil && len(*gss) > 0 {
		return &((*gss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of global.search not found", id)
}

// GetGlobalSearchs gets global.search existing records.
func (c *Client) GetGlobalSearchs(ids []int64) (*GlobalSearchs, error) {
	gss := &GlobalSearchs{}
	if err := c.Read(GlobalSearchModel, ids, nil, gss); err != nil {
		return nil, err
	}
	return gss, nil
}

// FindGlobalSearch finds global.search record by querying it with criteria.
func (c *Client) FindGlobalSearch(criteria *Criteria) (*GlobalSearch, error) {
	gss := &GlobalSearchs{}
	if err := c.SearchRead(GlobalSearchModel, criteria, NewOptions().Limit(1), gss); err != nil {
		return nil, err
	}
	if gss != nil && len(*gss) > 0 {
		return &((*gss)[0]), nil
	}
	return nil, fmt.Errorf("global.search was not found with criteria %v", criteria)
}

// FindGlobalSearchs finds global.search records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGlobalSearchs(criteria *Criteria, options *Options) (*GlobalSearchs, error) {
	gss := &GlobalSearchs{}
	if err := c.SearchRead(GlobalSearchModel, criteria, options, gss); err != nil {
		return nil, err
	}
	return gss, nil
}

// FindGlobalSearchIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGlobalSearchIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GlobalSearchModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGlobalSearchId finds record id by querying it with criteria.
func (c *Client) FindGlobalSearchId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GlobalSearchModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("global.search was not found with criteria %v and options %v", criteria, options)
}
