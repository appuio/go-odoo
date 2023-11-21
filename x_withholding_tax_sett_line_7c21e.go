package odoo

import (
	"fmt"
)

// XWithholdingTaxSettLine7C21E represents x_withholding_tax_sett_line_7c21e model.
type XWithholdingTaxSettLine7C21E struct {
	LastUpdate                      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                     *String   `xmlrpc:"display_name,omptempty"`
	Id                              *Int      `xmlrpc:"id,omptempty"`
	WriteDate                       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                        *Many2One `xmlrpc:"write_uid,omptempty"`
	XName                           *String   `xmlrpc:"x_name,omptempty"`
	XStudioCanton                   *String   `xmlrpc:"x_studio_canton,omptempty"`
	XStudioSequence                 *Int      `xmlrpc:"x_studio_sequence,omptempty"`
	XStudioWage                     *Float    `xmlrpc:"x_studio_wage,omptempty"`
	XStudioWithholdingTaxCode       *String   `xmlrpc:"x_studio_withholding_tax_code,omptempty"`
	XStudioWithholdingTaxPercentage *Float    `xmlrpc:"x_studio_withholding_tax_percentage,omptempty"`
	XStudioYear                     *Int      `xmlrpc:"x_studio_year,omptempty"`
	XStudioYearNoSeperator          *String   `xmlrpc:"x_studio_year_no_seperator,omptempty"`
	XWithholdingTaxSettId           *Many2One `xmlrpc:"x_withholding_tax_sett_id,omptempty"`
}

// XWithholdingTaxSettLine7C21Es represents array of x_withholding_tax_sett_line_7c21e model.
type XWithholdingTaxSettLine7C21Es []XWithholdingTaxSettLine7C21E

// XWithholdingTaxSettLine7C21EModel is the odoo model name.
const XWithholdingTaxSettLine7C21EModel = "x_withholding_tax_sett_line_7c21e"

// Many2One convert XWithholdingTaxSettLine7C21E to *Many2One.
func (x *XWithholdingTaxSettLine7C21E) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXWithholdingTaxSettLine7C21E creates a new x_withholding_tax_sett_line_7c21e model and returns its id.
func (c *Client) CreateXWithholdingTaxSettLine7C21E(x *XWithholdingTaxSettLine7C21E) (int64, error) {
	ids, err := c.CreateXWithholdingTaxSettLine7C21Es([]*XWithholdingTaxSettLine7C21E{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXWithholdingTaxSettLine7C21E creates a new x_withholding_tax_sett_line_7c21e model and returns its id.
func (c *Client) CreateXWithholdingTaxSettLine7C21Es(xs []*XWithholdingTaxSettLine7C21E) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XWithholdingTaxSettLine7C21EModel, vv)
}

// UpdateXWithholdingTaxSettLine7C21E updates an existing x_withholding_tax_sett_line_7c21e record.
func (c *Client) UpdateXWithholdingTaxSettLine7C21E(x *XWithholdingTaxSettLine7C21E) error {
	return c.UpdateXWithholdingTaxSettLine7C21Es([]int64{x.Id.Get()}, x)
}

// UpdateXWithholdingTaxSettLine7C21Es updates existing x_withholding_tax_sett_line_7c21e records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXWithholdingTaxSettLine7C21Es(ids []int64, x *XWithholdingTaxSettLine7C21E) error {
	return c.Update(XWithholdingTaxSettLine7C21EModel, ids, x)
}

// DeleteXWithholdingTaxSettLine7C21E deletes an existing x_withholding_tax_sett_line_7c21e record.
func (c *Client) DeleteXWithholdingTaxSettLine7C21E(id int64) error {
	return c.DeleteXWithholdingTaxSettLine7C21Es([]int64{id})
}

// DeleteXWithholdingTaxSettLine7C21Es deletes existing x_withholding_tax_sett_line_7c21e records.
func (c *Client) DeleteXWithholdingTaxSettLine7C21Es(ids []int64) error {
	return c.Delete(XWithholdingTaxSettLine7C21EModel, ids)
}

// GetXWithholdingTaxSettLine7C21E gets x_withholding_tax_sett_line_7c21e existing record.
func (c *Client) GetXWithholdingTaxSettLine7C21E(id int64) (*XWithholdingTaxSettLine7C21E, error) {
	xs, err := c.GetXWithholdingTaxSettLine7C21Es([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_withholding_tax_sett_line_7c21e not found", id)
}

// GetXWithholdingTaxSettLine7C21Es gets x_withholding_tax_sett_line_7c21e existing records.
func (c *Client) GetXWithholdingTaxSettLine7C21Es(ids []int64) (*XWithholdingTaxSettLine7C21Es, error) {
	xs := &XWithholdingTaxSettLine7C21Es{}
	if err := c.Read(XWithholdingTaxSettLine7C21EModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTaxSettLine7C21E finds x_withholding_tax_sett_line_7c21e record by querying it with criteria.
func (c *Client) FindXWithholdingTaxSettLine7C21E(criteria *Criteria) (*XWithholdingTaxSettLine7C21E, error) {
	xs := &XWithholdingTaxSettLine7C21Es{}
	if err := c.SearchRead(XWithholdingTaxSettLine7C21EModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_withholding_tax_sett_line_7c21e was not found with criteria %v", criteria)
}

// FindXWithholdingTaxSettLine7C21Es finds x_withholding_tax_sett_line_7c21e records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxSettLine7C21Es(criteria *Criteria, options *Options) (*XWithholdingTaxSettLine7C21Es, error) {
	xs := &XWithholdingTaxSettLine7C21Es{}
	if err := c.SearchRead(XWithholdingTaxSettLine7C21EModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXWithholdingTaxSettLine7C21EIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXWithholdingTaxSettLine7C21EIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XWithholdingTaxSettLine7C21EModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXWithholdingTaxSettLine7C21EId finds record id by querying it with criteria.
func (c *Client) FindXWithholdingTaxSettLine7C21EId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XWithholdingTaxSettLine7C21EModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_withholding_tax_sett_line_7c21e was not found with criteria %v and options %v", criteria, options)
}
