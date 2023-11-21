package odoo

import (
	"fmt"
)

// SpreadsheetRevision represents spreadsheet.revision model.
type SpreadsheetRevision struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	Active           *Bool     `xmlrpc:"active,omptempty"`
	Commands         *String   `xmlrpc:"commands,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	ParentRevisionId *String   `xmlrpc:"parent_revision_id,omptempty"`
	ResId            *Many2One `xmlrpc:"res_id,omptempty"`
	ResModel         *String   `xmlrpc:"res_model,omptempty"`
	RevisionId       *String   `xmlrpc:"revision_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SpreadsheetRevisions represents array of spreadsheet.revision model.
type SpreadsheetRevisions []SpreadsheetRevision

// SpreadsheetRevisionModel is the odoo model name.
const SpreadsheetRevisionModel = "spreadsheet.revision"

// Many2One convert SpreadsheetRevision to *Many2One.
func (sr *SpreadsheetRevision) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateSpreadsheetRevision creates a new spreadsheet.revision model and returns its id.
func (c *Client) CreateSpreadsheetRevision(sr *SpreadsheetRevision) (int64, error) {
	ids, err := c.CreateSpreadsheetRevisions([]*SpreadsheetRevision{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetRevision creates a new spreadsheet.revision model and returns its id.
func (c *Client) CreateSpreadsheetRevisions(srs []*SpreadsheetRevision) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetRevisionModel, vv)
}

// UpdateSpreadsheetRevision updates an existing spreadsheet.revision record.
func (c *Client) UpdateSpreadsheetRevision(sr *SpreadsheetRevision) error {
	return c.UpdateSpreadsheetRevisions([]int64{sr.Id.Get()}, sr)
}

// UpdateSpreadsheetRevisions updates existing spreadsheet.revision records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateSpreadsheetRevisions(ids []int64, sr *SpreadsheetRevision) error {
	return c.Update(SpreadsheetRevisionModel, ids, sr)
}

// DeleteSpreadsheetRevision deletes an existing spreadsheet.revision record.
func (c *Client) DeleteSpreadsheetRevision(id int64) error {
	return c.DeleteSpreadsheetRevisions([]int64{id})
}

// DeleteSpreadsheetRevisions deletes existing spreadsheet.revision records.
func (c *Client) DeleteSpreadsheetRevisions(ids []int64) error {
	return c.Delete(SpreadsheetRevisionModel, ids)
}

// GetSpreadsheetRevision gets spreadsheet.revision existing record.
func (c *Client) GetSpreadsheetRevision(id int64) (*SpreadsheetRevision, error) {
	srs, err := c.GetSpreadsheetRevisions([]int64{id})
	if err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of spreadsheet.revision not found", id)
}

// GetSpreadsheetRevisions gets spreadsheet.revision existing records.
func (c *Client) GetSpreadsheetRevisions(ids []int64) (*SpreadsheetRevisions, error) {
	srs := &SpreadsheetRevisions{}
	if err := c.Read(SpreadsheetRevisionModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSpreadsheetRevision finds spreadsheet.revision record by querying it with criteria.
func (c *Client) FindSpreadsheetRevision(criteria *Criteria) (*SpreadsheetRevision, error) {
	srs := &SpreadsheetRevisions{}
	if err := c.SearchRead(SpreadsheetRevisionModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("spreadsheet.revision was not found with criteria %v", criteria)
}

// FindSpreadsheetRevisions finds spreadsheet.revision records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetRevisions(criteria *Criteria, options *Options) (*SpreadsheetRevisions, error) {
	srs := &SpreadsheetRevisions{}
	if err := c.SearchRead(SpreadsheetRevisionModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSpreadsheetRevisionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetRevisionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SpreadsheetRevisionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSpreadsheetRevisionId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetRevisionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetRevisionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("spreadsheet.revision was not found with criteria %v and options %v", criteria, options)
}
