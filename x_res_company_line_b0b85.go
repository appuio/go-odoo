package odoo

import (
	"fmt"
)

// XResCompanyLineB0B85 represents x_res_company_line_b0b85 model.
type XResCompanyLineB0B85 struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
	XName           *String   `xmlrpc:"x_name,omptempty"`
	XResCompanyId   *Many2One `xmlrpc:"x_res_company_id,omptempty"`
	XStudioSequence *Int      `xmlrpc:"x_studio_sequence,omptempty"`
}

// XResCompanyLineB0B85s represents array of x_res_company_line_b0b85 model.
type XResCompanyLineB0B85s []XResCompanyLineB0B85

// XResCompanyLineB0B85Model is the odoo model name.
const XResCompanyLineB0B85Model = "x_res_company_line_b0b85"

// Many2One convert XResCompanyLineB0B85 to *Many2One.
func (x *XResCompanyLineB0B85) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXResCompanyLineB0B85 creates a new x_res_company_line_b0b85 model and returns its id.
func (c *Client) CreateXResCompanyLineB0B85(x *XResCompanyLineB0B85) (int64, error) {
	ids, err := c.CreateXResCompanyLineB0B85s([]*XResCompanyLineB0B85{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXResCompanyLineB0B85 creates a new x_res_company_line_b0b85 model and returns its id.
func (c *Client) CreateXResCompanyLineB0B85s(xs []*XResCompanyLineB0B85) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XResCompanyLineB0B85Model, vv)
}

// UpdateXResCompanyLineB0B85 updates an existing x_res_company_line_b0b85 record.
func (c *Client) UpdateXResCompanyLineB0B85(x *XResCompanyLineB0B85) error {
	return c.UpdateXResCompanyLineB0B85s([]int64{x.Id.Get()}, x)
}

// UpdateXResCompanyLineB0B85s updates existing x_res_company_line_b0b85 records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXResCompanyLineB0B85s(ids []int64, x *XResCompanyLineB0B85) error {
	return c.Update(XResCompanyLineB0B85Model, ids, x)
}

// DeleteXResCompanyLineB0B85 deletes an existing x_res_company_line_b0b85 record.
func (c *Client) DeleteXResCompanyLineB0B85(id int64) error {
	return c.DeleteXResCompanyLineB0B85s([]int64{id})
}

// DeleteXResCompanyLineB0B85s deletes existing x_res_company_line_b0b85 records.
func (c *Client) DeleteXResCompanyLineB0B85s(ids []int64) error {
	return c.Delete(XResCompanyLineB0B85Model, ids)
}

// GetXResCompanyLineB0B85 gets x_res_company_line_b0b85 existing record.
func (c *Client) GetXResCompanyLineB0B85(id int64) (*XResCompanyLineB0B85, error) {
	xs, err := c.GetXResCompanyLineB0B85s([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_res_company_line_b0b85 not found", id)
}

// GetXResCompanyLineB0B85s gets x_res_company_line_b0b85 existing records.
func (c *Client) GetXResCompanyLineB0B85s(ids []int64) (*XResCompanyLineB0B85s, error) {
	xs := &XResCompanyLineB0B85s{}
	if err := c.Read(XResCompanyLineB0B85Model, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXResCompanyLineB0B85 finds x_res_company_line_b0b85 record by querying it with criteria.
func (c *Client) FindXResCompanyLineB0B85(criteria *Criteria) (*XResCompanyLineB0B85, error) {
	xs := &XResCompanyLineB0B85s{}
	if err := c.SearchRead(XResCompanyLineB0B85Model, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_res_company_line_b0b85 was not found with criteria %v", criteria)
}

// FindXResCompanyLineB0B85s finds x_res_company_line_b0b85 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXResCompanyLineB0B85s(criteria *Criteria, options *Options) (*XResCompanyLineB0B85s, error) {
	xs := &XResCompanyLineB0B85s{}
	if err := c.SearchRead(XResCompanyLineB0B85Model, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXResCompanyLineB0B85Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXResCompanyLineB0B85Ids(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XResCompanyLineB0B85Model, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXResCompanyLineB0B85Id finds record id by querying it with criteria.
func (c *Client) FindXResCompanyLineB0B85Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XResCompanyLineB0B85Model, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_res_company_line_b0b85 was not found with criteria %v and options %v", criteria, options)
}
