package odoo

import (
	"fmt"
)

// XHrEmployeeLine25Be5 represents x_hr_employee_line_25be5 model.
type XHrEmployeeLine25Be5 struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
	XHrEmployeeId   *Many2One `xmlrpc:"x_hr_employee_id,omptempty"`
	XName           *String   `xmlrpc:"x_name,omptempty"`
	XStudioSequence *Int      `xmlrpc:"x_studio_sequence,omptempty"`
}

// XHrEmployeeLine25Be5s represents array of x_hr_employee_line_25be5 model.
type XHrEmployeeLine25Be5s []XHrEmployeeLine25Be5

// XHrEmployeeLine25Be5Model is the odoo model name.
const XHrEmployeeLine25Be5Model = "x_hr_employee_line_25be5"

// Many2One convert XHrEmployeeLine25Be5 to *Many2One.
func (x *XHrEmployeeLine25Be5) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXHrEmployeeLine25Be5 creates a new x_hr_employee_line_25be5 model and returns its id.
func (c *Client) CreateXHrEmployeeLine25Be5(x *XHrEmployeeLine25Be5) (int64, error) {
	ids, err := c.CreateXHrEmployeeLine25Be5s([]*XHrEmployeeLine25Be5{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXHrEmployeeLine25Be5 creates a new x_hr_employee_line_25be5 model and returns its id.
func (c *Client) CreateXHrEmployeeLine25Be5s(xs []*XHrEmployeeLine25Be5) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XHrEmployeeLine25Be5Model, vv)
}

// UpdateXHrEmployeeLine25Be5 updates an existing x_hr_employee_line_25be5 record.
func (c *Client) UpdateXHrEmployeeLine25Be5(x *XHrEmployeeLine25Be5) error {
	return c.UpdateXHrEmployeeLine25Be5s([]int64{x.Id.Get()}, x)
}

// UpdateXHrEmployeeLine25Be5s updates existing x_hr_employee_line_25be5 records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXHrEmployeeLine25Be5s(ids []int64, x *XHrEmployeeLine25Be5) error {
	return c.Update(XHrEmployeeLine25Be5Model, ids, x)
}

// DeleteXHrEmployeeLine25Be5 deletes an existing x_hr_employee_line_25be5 record.
func (c *Client) DeleteXHrEmployeeLine25Be5(id int64) error {
	return c.DeleteXHrEmployeeLine25Be5s([]int64{id})
}

// DeleteXHrEmployeeLine25Be5s deletes existing x_hr_employee_line_25be5 records.
func (c *Client) DeleteXHrEmployeeLine25Be5s(ids []int64) error {
	return c.Delete(XHrEmployeeLine25Be5Model, ids)
}

// GetXHrEmployeeLine25Be5 gets x_hr_employee_line_25be5 existing record.
func (c *Client) GetXHrEmployeeLine25Be5(id int64) (*XHrEmployeeLine25Be5, error) {
	xs, err := c.GetXHrEmployeeLine25Be5s([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_hr_employee_line_25be5 not found", id)
}

// GetXHrEmployeeLine25Be5s gets x_hr_employee_line_25be5 existing records.
func (c *Client) GetXHrEmployeeLine25Be5s(ids []int64) (*XHrEmployeeLine25Be5s, error) {
	xs := &XHrEmployeeLine25Be5s{}
	if err := c.Read(XHrEmployeeLine25Be5Model, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrEmployeeLine25Be5 finds x_hr_employee_line_25be5 record by querying it with criteria.
func (c *Client) FindXHrEmployeeLine25Be5(criteria *Criteria) (*XHrEmployeeLine25Be5, error) {
	xs := &XHrEmployeeLine25Be5s{}
	if err := c.SearchRead(XHrEmployeeLine25Be5Model, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_hr_employee_line_25be5 was not found with criteria %v", criteria)
}

// FindXHrEmployeeLine25Be5s finds x_hr_employee_line_25be5 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrEmployeeLine25Be5s(criteria *Criteria, options *Options) (*XHrEmployeeLine25Be5s, error) {
	xs := &XHrEmployeeLine25Be5s{}
	if err := c.SearchRead(XHrEmployeeLine25Be5Model, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrEmployeeLine25Be5Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrEmployeeLine25Be5Ids(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XHrEmployeeLine25Be5Model, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXHrEmployeeLine25Be5Id finds record id by querying it with criteria.
func (c *Client) FindXHrEmployeeLine25Be5Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XHrEmployeeLine25Be5Model, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_hr_employee_line_25be5 was not found with criteria %v and options %v", criteria, options)
}
