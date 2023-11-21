package odoo

import (
	"fmt"
)

// XResCompanyLine5Cf9D represents x_res_company_line_5cf9d model.
type XResCompanyLine5Cf9D struct {
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

// XResCompanyLine5Cf9Ds represents array of x_res_company_line_5cf9d model.
type XResCompanyLine5Cf9Ds []XResCompanyLine5Cf9D

// XResCompanyLine5Cf9DModel is the odoo model name.
const XResCompanyLine5Cf9DModel = "x_res_company_line_5cf9d"

// Many2One convert XResCompanyLine5Cf9D to *Many2One.
func (x *XResCompanyLine5Cf9D) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXResCompanyLine5Cf9D creates a new x_res_company_line_5cf9d model and returns its id.
func (c *Client) CreateXResCompanyLine5Cf9D(x *XResCompanyLine5Cf9D) (int64, error) {
	ids, err := c.CreateXResCompanyLine5Cf9Ds([]*XResCompanyLine5Cf9D{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXResCompanyLine5Cf9D creates a new x_res_company_line_5cf9d model and returns its id.
func (c *Client) CreateXResCompanyLine5Cf9Ds(xs []*XResCompanyLine5Cf9D) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XResCompanyLine5Cf9DModel, vv)
}

// UpdateXResCompanyLine5Cf9D updates an existing x_res_company_line_5cf9d record.
func (c *Client) UpdateXResCompanyLine5Cf9D(x *XResCompanyLine5Cf9D) error {
	return c.UpdateXResCompanyLine5Cf9Ds([]int64{x.Id.Get()}, x)
}

// UpdateXResCompanyLine5Cf9Ds updates existing x_res_company_line_5cf9d records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXResCompanyLine5Cf9Ds(ids []int64, x *XResCompanyLine5Cf9D) error {
	return c.Update(XResCompanyLine5Cf9DModel, ids, x)
}

// DeleteXResCompanyLine5Cf9D deletes an existing x_res_company_line_5cf9d record.
func (c *Client) DeleteXResCompanyLine5Cf9D(id int64) error {
	return c.DeleteXResCompanyLine5Cf9Ds([]int64{id})
}

// DeleteXResCompanyLine5Cf9Ds deletes existing x_res_company_line_5cf9d records.
func (c *Client) DeleteXResCompanyLine5Cf9Ds(ids []int64) error {
	return c.Delete(XResCompanyLine5Cf9DModel, ids)
}

// GetXResCompanyLine5Cf9D gets x_res_company_line_5cf9d existing record.
func (c *Client) GetXResCompanyLine5Cf9D(id int64) (*XResCompanyLine5Cf9D, error) {
	xs, err := c.GetXResCompanyLine5Cf9Ds([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_res_company_line_5cf9d not found", id)
}

// GetXResCompanyLine5Cf9Ds gets x_res_company_line_5cf9d existing records.
func (c *Client) GetXResCompanyLine5Cf9Ds(ids []int64) (*XResCompanyLine5Cf9Ds, error) {
	xs := &XResCompanyLine5Cf9Ds{}
	if err := c.Read(XResCompanyLine5Cf9DModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXResCompanyLine5Cf9D finds x_res_company_line_5cf9d record by querying it with criteria.
func (c *Client) FindXResCompanyLine5Cf9D(criteria *Criteria) (*XResCompanyLine5Cf9D, error) {
	xs := &XResCompanyLine5Cf9Ds{}
	if err := c.SearchRead(XResCompanyLine5Cf9DModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_res_company_line_5cf9d was not found with criteria %v", criteria)
}

// FindXResCompanyLine5Cf9Ds finds x_res_company_line_5cf9d records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXResCompanyLine5Cf9Ds(criteria *Criteria, options *Options) (*XResCompanyLine5Cf9Ds, error) {
	xs := &XResCompanyLine5Cf9Ds{}
	if err := c.SearchRead(XResCompanyLine5Cf9DModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXResCompanyLine5Cf9DIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXResCompanyLine5Cf9DIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XResCompanyLine5Cf9DModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXResCompanyLine5Cf9DId finds record id by querying it with criteria.
func (c *Client) FindXResCompanyLine5Cf9DId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XResCompanyLine5Cf9DModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_res_company_line_5cf9d was not found with criteria %v and options %v", criteria, options)
}
