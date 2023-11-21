package odoo

import (
	"fmt"
)

// SpreadsheetContributor represents spreadsheet.contributor model.
type SpreadsheetContributor struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	DocumentId     *Many2One `xmlrpc:"document_id,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LastUpdateDate *Time     `xmlrpc:"last_update_date,omptempty"`
	UserId         *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SpreadsheetContributors represents array of spreadsheet.contributor model.
type SpreadsheetContributors []SpreadsheetContributor

// SpreadsheetContributorModel is the odoo model name.
const SpreadsheetContributorModel = "spreadsheet.contributor"

// Many2One convert SpreadsheetContributor to *Many2One.
func (sc *SpreadsheetContributor) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSpreadsheetContributor creates a new spreadsheet.contributor model and returns its id.
func (c *Client) CreateSpreadsheetContributor(sc *SpreadsheetContributor) (int64, error) {
	ids, err := c.CreateSpreadsheetContributors([]*SpreadsheetContributor{sc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetContributor creates a new spreadsheet.contributor model and returns its id.
func (c *Client) CreateSpreadsheetContributors(scs []*SpreadsheetContributor) ([]int64, error) {
	var vv []interface{}
	for _, v := range scs {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetContributorModel, vv)
}

// UpdateSpreadsheetContributor updates an existing spreadsheet.contributor record.
func (c *Client) UpdateSpreadsheetContributor(sc *SpreadsheetContributor) error {
	return c.UpdateSpreadsheetContributors([]int64{sc.Id.Get()}, sc)
}

// UpdateSpreadsheetContributors updates existing spreadsheet.contributor records.
// All records (represented by ids) will be updated by sc values.
func (c *Client) UpdateSpreadsheetContributors(ids []int64, sc *SpreadsheetContributor) error {
	return c.Update(SpreadsheetContributorModel, ids, sc)
}

// DeleteSpreadsheetContributor deletes an existing spreadsheet.contributor record.
func (c *Client) DeleteSpreadsheetContributor(id int64) error {
	return c.DeleteSpreadsheetContributors([]int64{id})
}

// DeleteSpreadsheetContributors deletes existing spreadsheet.contributor records.
func (c *Client) DeleteSpreadsheetContributors(ids []int64) error {
	return c.Delete(SpreadsheetContributorModel, ids)
}

// GetSpreadsheetContributor gets spreadsheet.contributor existing record.
func (c *Client) GetSpreadsheetContributor(id int64) (*SpreadsheetContributor, error) {
	scs, err := c.GetSpreadsheetContributors([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of spreadsheet.contributor not found", id)
}

// GetSpreadsheetContributors gets spreadsheet.contributor existing records.
func (c *Client) GetSpreadsheetContributors(ids []int64) (*SpreadsheetContributors, error) {
	scs := &SpreadsheetContributors{}
	if err := c.Read(SpreadsheetContributorModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSpreadsheetContributor finds spreadsheet.contributor record by querying it with criteria.
func (c *Client) FindSpreadsheetContributor(criteria *Criteria) (*SpreadsheetContributor, error) {
	scs := &SpreadsheetContributors{}
	if err := c.SearchRead(SpreadsheetContributorModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("spreadsheet.contributor was not found with criteria %v", criteria)
}

// FindSpreadsheetContributors finds spreadsheet.contributor records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetContributors(criteria *Criteria, options *Options) (*SpreadsheetContributors, error) {
	scs := &SpreadsheetContributors{}
	if err := c.SearchRead(SpreadsheetContributorModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSpreadsheetContributorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetContributorIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SpreadsheetContributorModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSpreadsheetContributorId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetContributorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetContributorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("spreadsheet.contributor was not found with criteria %v and options %v", criteria, options)
}
