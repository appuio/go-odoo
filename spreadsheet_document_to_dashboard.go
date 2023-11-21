package odoo

import (
	"fmt"
)

// SpreadsheetDocumentToDashboard represents spreadsheet.document.to.dashboard model.
type SpreadsheetDocumentToDashboard struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DashboardGroupId *Many2One `xmlrpc:"dashboard_group_id,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	DocumentId       *Many2One `xmlrpc:"document_id,omptempty"`
	GroupIds         *Relation `xmlrpc:"group_ids,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SpreadsheetDocumentToDashboards represents array of spreadsheet.document.to.dashboard model.
type SpreadsheetDocumentToDashboards []SpreadsheetDocumentToDashboard

// SpreadsheetDocumentToDashboardModel is the odoo model name.
const SpreadsheetDocumentToDashboardModel = "spreadsheet.document.to.dashboard"

// Many2One convert SpreadsheetDocumentToDashboard to *Many2One.
func (sdtd *SpreadsheetDocumentToDashboard) Many2One() *Many2One {
	return NewMany2One(sdtd.Id.Get(), "")
}

// CreateSpreadsheetDocumentToDashboard creates a new spreadsheet.document.to.dashboard model and returns its id.
func (c *Client) CreateSpreadsheetDocumentToDashboard(sdtd *SpreadsheetDocumentToDashboard) (int64, error) {
	ids, err := c.CreateSpreadsheetDocumentToDashboards([]*SpreadsheetDocumentToDashboard{sdtd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetDocumentToDashboard creates a new spreadsheet.document.to.dashboard model and returns its id.
func (c *Client) CreateSpreadsheetDocumentToDashboards(sdtds []*SpreadsheetDocumentToDashboard) ([]int64, error) {
	var vv []interface{}
	for _, v := range sdtds {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetDocumentToDashboardModel, vv)
}

// UpdateSpreadsheetDocumentToDashboard updates an existing spreadsheet.document.to.dashboard record.
func (c *Client) UpdateSpreadsheetDocumentToDashboard(sdtd *SpreadsheetDocumentToDashboard) error {
	return c.UpdateSpreadsheetDocumentToDashboards([]int64{sdtd.Id.Get()}, sdtd)
}

// UpdateSpreadsheetDocumentToDashboards updates existing spreadsheet.document.to.dashboard records.
// All records (represented by ids) will be updated by sdtd values.
func (c *Client) UpdateSpreadsheetDocumentToDashboards(ids []int64, sdtd *SpreadsheetDocumentToDashboard) error {
	return c.Update(SpreadsheetDocumentToDashboardModel, ids, sdtd)
}

// DeleteSpreadsheetDocumentToDashboard deletes an existing spreadsheet.document.to.dashboard record.
func (c *Client) DeleteSpreadsheetDocumentToDashboard(id int64) error {
	return c.DeleteSpreadsheetDocumentToDashboards([]int64{id})
}

// DeleteSpreadsheetDocumentToDashboards deletes existing spreadsheet.document.to.dashboard records.
func (c *Client) DeleteSpreadsheetDocumentToDashboards(ids []int64) error {
	return c.Delete(SpreadsheetDocumentToDashboardModel, ids)
}

// GetSpreadsheetDocumentToDashboard gets spreadsheet.document.to.dashboard existing record.
func (c *Client) GetSpreadsheetDocumentToDashboard(id int64) (*SpreadsheetDocumentToDashboard, error) {
	sdtds, err := c.GetSpreadsheetDocumentToDashboards([]int64{id})
	if err != nil {
		return nil, err
	}
	if sdtds != nil && len(*sdtds) > 0 {
		return &((*sdtds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of spreadsheet.document.to.dashboard not found", id)
}

// GetSpreadsheetDocumentToDashboards gets spreadsheet.document.to.dashboard existing records.
func (c *Client) GetSpreadsheetDocumentToDashboards(ids []int64) (*SpreadsheetDocumentToDashboards, error) {
	sdtds := &SpreadsheetDocumentToDashboards{}
	if err := c.Read(SpreadsheetDocumentToDashboardModel, ids, nil, sdtds); err != nil {
		return nil, err
	}
	return sdtds, nil
}

// FindSpreadsheetDocumentToDashboard finds spreadsheet.document.to.dashboard record by querying it with criteria.
func (c *Client) FindSpreadsheetDocumentToDashboard(criteria *Criteria) (*SpreadsheetDocumentToDashboard, error) {
	sdtds := &SpreadsheetDocumentToDashboards{}
	if err := c.SearchRead(SpreadsheetDocumentToDashboardModel, criteria, NewOptions().Limit(1), sdtds); err != nil {
		return nil, err
	}
	if sdtds != nil && len(*sdtds) > 0 {
		return &((*sdtds)[0]), nil
	}
	return nil, fmt.Errorf("spreadsheet.document.to.dashboard was not found with criteria %v", criteria)
}

// FindSpreadsheetDocumentToDashboards finds spreadsheet.document.to.dashboard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDocumentToDashboards(criteria *Criteria, options *Options) (*SpreadsheetDocumentToDashboards, error) {
	sdtds := &SpreadsheetDocumentToDashboards{}
	if err := c.SearchRead(SpreadsheetDocumentToDashboardModel, criteria, options, sdtds); err != nil {
		return nil, err
	}
	return sdtds, nil
}

// FindSpreadsheetDocumentToDashboardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDocumentToDashboardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SpreadsheetDocumentToDashboardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSpreadsheetDocumentToDashboardId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetDocumentToDashboardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetDocumentToDashboardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("spreadsheet.document.to.dashboard was not found with criteria %v and options %v", criteria, options)
}
