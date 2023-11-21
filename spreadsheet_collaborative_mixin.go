package odoo

import (
	"fmt"
)

// SpreadsheetCollaborativeMixin represents spreadsheet.collaborative.mixin model.
type SpreadsheetCollaborativeMixin struct {
	Raw                    *String   `xmlrpc:"raw,omptempty"`
	SpreadsheetRevisionIds *Relation `xmlrpc:"spreadsheet_revision_ids,omptempty"`
	SpreadsheetSnapshot    *String   `xmlrpc:"spreadsheet_snapshot,omptempty"`
}

// SpreadsheetCollaborativeMixins represents array of spreadsheet.collaborative.mixin model.
type SpreadsheetCollaborativeMixins []SpreadsheetCollaborativeMixin

// SpreadsheetCollaborativeMixinModel is the odoo model name.
const SpreadsheetCollaborativeMixinModel = "spreadsheet.collaborative.mixin"

// Many2One convert SpreadsheetCollaborativeMixin to *Many2One.
func (scm *SpreadsheetCollaborativeMixin) Many2One() *Many2One {
	return NewMany2One(scm.Id.Get(), "")
}

// CreateSpreadsheetCollaborativeMixin creates a new spreadsheet.collaborative.mixin model and returns its id.
func (c *Client) CreateSpreadsheetCollaborativeMixin(scm *SpreadsheetCollaborativeMixin) (int64, error) {
	ids, err := c.CreateSpreadsheetCollaborativeMixins([]*SpreadsheetCollaborativeMixin{scm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetCollaborativeMixin creates a new spreadsheet.collaborative.mixin model and returns its id.
func (c *Client) CreateSpreadsheetCollaborativeMixins(scms []*SpreadsheetCollaborativeMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range scms {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetCollaborativeMixinModel, vv)
}

// UpdateSpreadsheetCollaborativeMixin updates an existing spreadsheet.collaborative.mixin record.
func (c *Client) UpdateSpreadsheetCollaborativeMixin(scm *SpreadsheetCollaborativeMixin) error {
	return c.UpdateSpreadsheetCollaborativeMixins([]int64{scm.Id.Get()}, scm)
}

// UpdateSpreadsheetCollaborativeMixins updates existing spreadsheet.collaborative.mixin records.
// All records (represented by ids) will be updated by scm values.
func (c *Client) UpdateSpreadsheetCollaborativeMixins(ids []int64, scm *SpreadsheetCollaborativeMixin) error {
	return c.Update(SpreadsheetCollaborativeMixinModel, ids, scm)
}

// DeleteSpreadsheetCollaborativeMixin deletes an existing spreadsheet.collaborative.mixin record.
func (c *Client) DeleteSpreadsheetCollaborativeMixin(id int64) error {
	return c.DeleteSpreadsheetCollaborativeMixins([]int64{id})
}

// DeleteSpreadsheetCollaborativeMixins deletes existing spreadsheet.collaborative.mixin records.
func (c *Client) DeleteSpreadsheetCollaborativeMixins(ids []int64) error {
	return c.Delete(SpreadsheetCollaborativeMixinModel, ids)
}

// GetSpreadsheetCollaborativeMixin gets spreadsheet.collaborative.mixin existing record.
func (c *Client) GetSpreadsheetCollaborativeMixin(id int64) (*SpreadsheetCollaborativeMixin, error) {
	scms, err := c.GetSpreadsheetCollaborativeMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if scms != nil && len(*scms) > 0 {
		return &((*scms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of spreadsheet.collaborative.mixin not found", id)
}

// GetSpreadsheetCollaborativeMixins gets spreadsheet.collaborative.mixin existing records.
func (c *Client) GetSpreadsheetCollaborativeMixins(ids []int64) (*SpreadsheetCollaborativeMixins, error) {
	scms := &SpreadsheetCollaborativeMixins{}
	if err := c.Read(SpreadsheetCollaborativeMixinModel, ids, nil, scms); err != nil {
		return nil, err
	}
	return scms, nil
}

// FindSpreadsheetCollaborativeMixin finds spreadsheet.collaborative.mixin record by querying it with criteria.
func (c *Client) FindSpreadsheetCollaborativeMixin(criteria *Criteria) (*SpreadsheetCollaborativeMixin, error) {
	scms := &SpreadsheetCollaborativeMixins{}
	if err := c.SearchRead(SpreadsheetCollaborativeMixinModel, criteria, NewOptions().Limit(1), scms); err != nil {
		return nil, err
	}
	if scms != nil && len(*scms) > 0 {
		return &((*scms)[0]), nil
	}
	return nil, fmt.Errorf("spreadsheet.collaborative.mixin was not found with criteria %v", criteria)
}

// FindSpreadsheetCollaborativeMixins finds spreadsheet.collaborative.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetCollaborativeMixins(criteria *Criteria, options *Options) (*SpreadsheetCollaborativeMixins, error) {
	scms := &SpreadsheetCollaborativeMixins{}
	if err := c.SearchRead(SpreadsheetCollaborativeMixinModel, criteria, options, scms); err != nil {
		return nil, err
	}
	return scms, nil
}

// FindSpreadsheetCollaborativeMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetCollaborativeMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SpreadsheetCollaborativeMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSpreadsheetCollaborativeMixinId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetCollaborativeMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetCollaborativeMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("spreadsheet.collaborative.mixin was not found with criteria %v and options %v", criteria, options)
}
