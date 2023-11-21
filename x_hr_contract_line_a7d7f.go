package odoo

import (
	"fmt"
)

// XHrContractLineA7D7F represents x_hr_contract_line_a7d7f model.
type XHrContractLineA7D7F struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
	XHrContractId   *Many2One `xmlrpc:"x_hr_contract_id,omptempty"`
	XName           *String   `xmlrpc:"x_name,omptempty"`
	XStudioSequence *Int      `xmlrpc:"x_studio_sequence,omptempty"`
}

// XHrContractLineA7D7Fs represents array of x_hr_contract_line_a7d7f model.
type XHrContractLineA7D7Fs []XHrContractLineA7D7F

// XHrContractLineA7D7FModel is the odoo model name.
const XHrContractLineA7D7FModel = "x_hr_contract_line_a7d7f"

// Many2One convert XHrContractLineA7D7F to *Many2One.
func (x *XHrContractLineA7D7F) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXHrContractLineA7D7F creates a new x_hr_contract_line_a7d7f model and returns its id.
func (c *Client) CreateXHrContractLineA7D7F(x *XHrContractLineA7D7F) (int64, error) {
	ids, err := c.CreateXHrContractLineA7D7Fs([]*XHrContractLineA7D7F{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXHrContractLineA7D7F creates a new x_hr_contract_line_a7d7f model and returns its id.
func (c *Client) CreateXHrContractLineA7D7Fs(xs []*XHrContractLineA7D7F) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XHrContractLineA7D7FModel, vv)
}

// UpdateXHrContractLineA7D7F updates an existing x_hr_contract_line_a7d7f record.
func (c *Client) UpdateXHrContractLineA7D7F(x *XHrContractLineA7D7F) error {
	return c.UpdateXHrContractLineA7D7Fs([]int64{x.Id.Get()}, x)
}

// UpdateXHrContractLineA7D7Fs updates existing x_hr_contract_line_a7d7f records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXHrContractLineA7D7Fs(ids []int64, x *XHrContractLineA7D7F) error {
	return c.Update(XHrContractLineA7D7FModel, ids, x)
}

// DeleteXHrContractLineA7D7F deletes an existing x_hr_contract_line_a7d7f record.
func (c *Client) DeleteXHrContractLineA7D7F(id int64) error {
	return c.DeleteXHrContractLineA7D7Fs([]int64{id})
}

// DeleteXHrContractLineA7D7Fs deletes existing x_hr_contract_line_a7d7f records.
func (c *Client) DeleteXHrContractLineA7D7Fs(ids []int64) error {
	return c.Delete(XHrContractLineA7D7FModel, ids)
}

// GetXHrContractLineA7D7F gets x_hr_contract_line_a7d7f existing record.
func (c *Client) GetXHrContractLineA7D7F(id int64) (*XHrContractLineA7D7F, error) {
	xs, err := c.GetXHrContractLineA7D7Fs([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_hr_contract_line_a7d7f not found", id)
}

// GetXHrContractLineA7D7Fs gets x_hr_contract_line_a7d7f existing records.
func (c *Client) GetXHrContractLineA7D7Fs(ids []int64) (*XHrContractLineA7D7Fs, error) {
	xs := &XHrContractLineA7D7Fs{}
	if err := c.Read(XHrContractLineA7D7FModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrContractLineA7D7F finds x_hr_contract_line_a7d7f record by querying it with criteria.
func (c *Client) FindXHrContractLineA7D7F(criteria *Criteria) (*XHrContractLineA7D7F, error) {
	xs := &XHrContractLineA7D7Fs{}
	if err := c.SearchRead(XHrContractLineA7D7FModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_hr_contract_line_a7d7f was not found with criteria %v", criteria)
}

// FindXHrContractLineA7D7Fs finds x_hr_contract_line_a7d7f records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrContractLineA7D7Fs(criteria *Criteria, options *Options) (*XHrContractLineA7D7Fs, error) {
	xs := &XHrContractLineA7D7Fs{}
	if err := c.SearchRead(XHrContractLineA7D7FModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrContractLineA7D7FIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrContractLineA7D7FIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XHrContractLineA7D7FModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXHrContractLineA7D7FId finds record id by querying it with criteria.
func (c *Client) FindXHrContractLineA7D7FId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XHrContractLineA7D7FModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_hr_contract_line_a7d7f was not found with criteria %v and options %v", criteria, options)
}
