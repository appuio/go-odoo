package odoo

import (
	"fmt"
)

// DocLayout represents doc.layout model.
type DocLayout struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Address           *Bool      `xmlrpc:"address,omptempty"`
	BaseColor         *String    `xmlrpc:"base_color,omptempty"`
	City              *Bool      `xmlrpc:"city,omptempty"`
	CompanyPosition   *Selection `xmlrpc:"company_position,omptempty"`
	CompanyTextColor  *String    `xmlrpc:"company_text_color,omptempty"`
	Country           *Bool      `xmlrpc:"country,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomerPosition  *Selection `xmlrpc:"customer_position,omptempty"`
	CustomerTextColor *String    `xmlrpc:"customer_text_color,omptempty"`
	Description       *Bool      `xmlrpc:"description,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	HeadingTextColor  *String    `xmlrpc:"heading_text_color,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	LogoPosition      *Selection `xmlrpc:"logo_position,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	Reference         *Bool      `xmlrpc:"reference,omptempty"`
	SalesPerson       *Bool      `xmlrpc:"sales_person,omptempty"`
	Source            *Bool      `xmlrpc:"source,omptempty"`
	TaglinePosition   *Selection `xmlrpc:"tagline_position,omptempty"`
	TaxValue          *Bool      `xmlrpc:"tax_value,omptempty"`
	TextColor         *String    `xmlrpc:"text_color,omptempty"`
	Vat               *Bool      `xmlrpc:"vat,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DocLayouts represents array of doc.layout model.
type DocLayouts []DocLayout

// DocLayoutModel is the odoo model name.
const DocLayoutModel = "doc.layout"

// Many2One convert DocLayout to *Many2One.
func (dl *DocLayout) Many2One() *Many2One {
	return NewMany2One(dl.Id.Get(), "")
}

// CreateDocLayout creates a new doc.layout model and returns its id.
func (c *Client) CreateDocLayout(dl *DocLayout) (int64, error) {
	ids, err := c.CreateDocLayouts([]*DocLayout{dl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocLayout creates a new doc.layout model and returns its id.
func (c *Client) CreateDocLayouts(dls []*DocLayout) ([]int64, error) {
	var vv []interface{}
	for _, v := range dls {
		vv = append(vv, v)
	}
	return c.Create(DocLayoutModel, vv)
}

// UpdateDocLayout updates an existing doc.layout record.
func (c *Client) UpdateDocLayout(dl *DocLayout) error {
	return c.UpdateDocLayouts([]int64{dl.Id.Get()}, dl)
}

// UpdateDocLayouts updates existing doc.layout records.
// All records (represented by ids) will be updated by dl values.
func (c *Client) UpdateDocLayouts(ids []int64, dl *DocLayout) error {
	return c.Update(DocLayoutModel, ids, dl)
}

// DeleteDocLayout deletes an existing doc.layout record.
func (c *Client) DeleteDocLayout(id int64) error {
	return c.DeleteDocLayouts([]int64{id})
}

// DeleteDocLayouts deletes existing doc.layout records.
func (c *Client) DeleteDocLayouts(ids []int64) error {
	return c.Delete(DocLayoutModel, ids)
}

// GetDocLayout gets doc.layout existing record.
func (c *Client) GetDocLayout(id int64) (*DocLayout, error) {
	dls, err := c.GetDocLayouts([]int64{id})
	if err != nil {
		return nil, err
	}
	if dls != nil && len(*dls) > 0 {
		return &((*dls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of doc.layout not found", id)
}

// GetDocLayouts gets doc.layout existing records.
func (c *Client) GetDocLayouts(ids []int64) (*DocLayouts, error) {
	dls := &DocLayouts{}
	if err := c.Read(DocLayoutModel, ids, nil, dls); err != nil {
		return nil, err
	}
	return dls, nil
}

// FindDocLayout finds doc.layout record by querying it with criteria.
func (c *Client) FindDocLayout(criteria *Criteria) (*DocLayout, error) {
	dls := &DocLayouts{}
	if err := c.SearchRead(DocLayoutModel, criteria, NewOptions().Limit(1), dls); err != nil {
		return nil, err
	}
	if dls != nil && len(*dls) > 0 {
		return &((*dls)[0]), nil
	}
	return nil, fmt.Errorf("doc.layout was not found with criteria %v", criteria)
}

// FindDocLayouts finds doc.layout records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocLayouts(criteria *Criteria, options *Options) (*DocLayouts, error) {
	dls := &DocLayouts{}
	if err := c.SearchRead(DocLayoutModel, criteria, options, dls); err != nil {
		return nil, err
	}
	return dls, nil
}

// FindDocLayoutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocLayoutIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocLayoutModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocLayoutId finds record id by querying it with criteria.
func (c *Client) FindDocLayoutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocLayoutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("doc.layout was not found with criteria %v and options %v", criteria, options)
}
