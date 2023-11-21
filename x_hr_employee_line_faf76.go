package odoo

import (
	"fmt"
)

// XHrEmployeeLineFaf76 represents x_hr_employee_line_faf76 model.
type XHrEmployeeLineFaf76 struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
	XHrEmployeeId        *Many2One `xmlrpc:"x_hr_employee_id,omptempty"`
	XName                *String   `xmlrpc:"x_name,omptempty"`
	XStudioBirthday      *Time     `xmlrpc:"x_studio_birthday,omptempty"`
	XStudioChildBirthday *Time     `xmlrpc:"x_studio_child_birthday,omptempty"`
	XStudioDateField     *Time     `xmlrpc:"x_studio_date_field,omptempty"`
	XStudioSequence      *Int      `xmlrpc:"x_studio_sequence,omptempty"`
}

// XHrEmployeeLineFaf76s represents array of x_hr_employee_line_faf76 model.
type XHrEmployeeLineFaf76s []XHrEmployeeLineFaf76

// XHrEmployeeLineFaf76Model is the odoo model name.
const XHrEmployeeLineFaf76Model = "x_hr_employee_line_faf76"

// Many2One convert XHrEmployeeLineFaf76 to *Many2One.
func (x *XHrEmployeeLineFaf76) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXHrEmployeeLineFaf76 creates a new x_hr_employee_line_faf76 model and returns its id.
func (c *Client) CreateXHrEmployeeLineFaf76(x *XHrEmployeeLineFaf76) (int64, error) {
	ids, err := c.CreateXHrEmployeeLineFaf76s([]*XHrEmployeeLineFaf76{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXHrEmployeeLineFaf76 creates a new x_hr_employee_line_faf76 model and returns its id.
func (c *Client) CreateXHrEmployeeLineFaf76s(xs []*XHrEmployeeLineFaf76) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XHrEmployeeLineFaf76Model, vv)
}

// UpdateXHrEmployeeLineFaf76 updates an existing x_hr_employee_line_faf76 record.
func (c *Client) UpdateXHrEmployeeLineFaf76(x *XHrEmployeeLineFaf76) error {
	return c.UpdateXHrEmployeeLineFaf76s([]int64{x.Id.Get()}, x)
}

// UpdateXHrEmployeeLineFaf76s updates existing x_hr_employee_line_faf76 records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXHrEmployeeLineFaf76s(ids []int64, x *XHrEmployeeLineFaf76) error {
	return c.Update(XHrEmployeeLineFaf76Model, ids, x)
}

// DeleteXHrEmployeeLineFaf76 deletes an existing x_hr_employee_line_faf76 record.
func (c *Client) DeleteXHrEmployeeLineFaf76(id int64) error {
	return c.DeleteXHrEmployeeLineFaf76s([]int64{id})
}

// DeleteXHrEmployeeLineFaf76s deletes existing x_hr_employee_line_faf76 records.
func (c *Client) DeleteXHrEmployeeLineFaf76s(ids []int64) error {
	return c.Delete(XHrEmployeeLineFaf76Model, ids)
}

// GetXHrEmployeeLineFaf76 gets x_hr_employee_line_faf76 existing record.
func (c *Client) GetXHrEmployeeLineFaf76(id int64) (*XHrEmployeeLineFaf76, error) {
	xs, err := c.GetXHrEmployeeLineFaf76s([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_hr_employee_line_faf76 not found", id)
}

// GetXHrEmployeeLineFaf76s gets x_hr_employee_line_faf76 existing records.
func (c *Client) GetXHrEmployeeLineFaf76s(ids []int64) (*XHrEmployeeLineFaf76s, error) {
	xs := &XHrEmployeeLineFaf76s{}
	if err := c.Read(XHrEmployeeLineFaf76Model, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrEmployeeLineFaf76 finds x_hr_employee_line_faf76 record by querying it with criteria.
func (c *Client) FindXHrEmployeeLineFaf76(criteria *Criteria) (*XHrEmployeeLineFaf76, error) {
	xs := &XHrEmployeeLineFaf76s{}
	if err := c.SearchRead(XHrEmployeeLineFaf76Model, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_hr_employee_line_faf76 was not found with criteria %v", criteria)
}

// FindXHrEmployeeLineFaf76s finds x_hr_employee_line_faf76 records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrEmployeeLineFaf76s(criteria *Criteria, options *Options) (*XHrEmployeeLineFaf76s, error) {
	xs := &XHrEmployeeLineFaf76s{}
	if err := c.SearchRead(XHrEmployeeLineFaf76Model, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrEmployeeLineFaf76Ids finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrEmployeeLineFaf76Ids(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XHrEmployeeLineFaf76Model, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXHrEmployeeLineFaf76Id finds record id by querying it with criteria.
func (c *Client) FindXHrEmployeeLineFaf76Id(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XHrEmployeeLineFaf76Model, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_hr_employee_line_faf76 was not found with criteria %v and options %v", criteria, options)
}
