package odoo

import (
	"fmt"
)

// XHrContractLine869Bf represents x_hr_contract_line_869bf model.
type XHrContractLine869Bf struct {
	LastUpdate                              *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                              *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                               *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                             *String   `xmlrpc:"display_name,omptempty"`
	Id                                      *Int      `xmlrpc:"id,omptempty"`
	WriteDate                               *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                                *Many2One `xmlrpc:"write_uid,omptempty"`
	XHrContractId                           *Many2One `xmlrpc:"x_hr_contract_id,omptempty"`
	XName                                   *String   `xmlrpc:"x_name,omptempty"`
	XStudioInternalStateCodeText            *String   `xmlrpc:"x_studio_internal_state_code_text,omptempty"`
	XStudioInternalWage                     *Float    `xmlrpc:"x_studio_internal_wage,omptempty"`
	XStudioInternalWithholdingTaxCode       *String   `xmlrpc:"x_studio_internal_withholding_tax_code,omptempty"`
	XStudioInternalWithholdingTaxPercentage *Float    `xmlrpc:"x_studio_internal_withholding_tax_percentage,omptempty"`
	XStudioSequence                         *Int      `xmlrpc:"x_studio_sequence,omptempty"`
}

// XHrContractLine869Bfs represents array of x_hr_contract_line_869bf model.
type XHrContractLine869Bfs []XHrContractLine869Bf

// XHrContractLine869BfModel is the odoo model name.
const XHrContractLine869BfModel = "x_hr_contract_line_869bf"

// Many2One convert XHrContractLine869Bf to *Many2One.
func (x *XHrContractLine869Bf) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXHrContractLine869Bf creates a new x_hr_contract_line_869bf model and returns its id.
func (c *Client) CreateXHrContractLine869Bf(x *XHrContractLine869Bf) (int64, error) {
	ids, err := c.CreateXHrContractLine869Bfs([]*XHrContractLine869Bf{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXHrContractLine869Bf creates a new x_hr_contract_line_869bf model and returns its id.
func (c *Client) CreateXHrContractLine869Bfs(xs []*XHrContractLine869Bf) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XHrContractLine869BfModel, vv)
}

// UpdateXHrContractLine869Bf updates an existing x_hr_contract_line_869bf record.
func (c *Client) UpdateXHrContractLine869Bf(x *XHrContractLine869Bf) error {
	return c.UpdateXHrContractLine869Bfs([]int64{x.Id.Get()}, x)
}

// UpdateXHrContractLine869Bfs updates existing x_hr_contract_line_869bf records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXHrContractLine869Bfs(ids []int64, x *XHrContractLine869Bf) error {
	return c.Update(XHrContractLine869BfModel, ids, x)
}

// DeleteXHrContractLine869Bf deletes an existing x_hr_contract_line_869bf record.
func (c *Client) DeleteXHrContractLine869Bf(id int64) error {
	return c.DeleteXHrContractLine869Bfs([]int64{id})
}

// DeleteXHrContractLine869Bfs deletes existing x_hr_contract_line_869bf records.
func (c *Client) DeleteXHrContractLine869Bfs(ids []int64) error {
	return c.Delete(XHrContractLine869BfModel, ids)
}

// GetXHrContractLine869Bf gets x_hr_contract_line_869bf existing record.
func (c *Client) GetXHrContractLine869Bf(id int64) (*XHrContractLine869Bf, error) {
	xs, err := c.GetXHrContractLine869Bfs([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_hr_contract_line_869bf not found", id)
}

// GetXHrContractLine869Bfs gets x_hr_contract_line_869bf existing records.
func (c *Client) GetXHrContractLine869Bfs(ids []int64) (*XHrContractLine869Bfs, error) {
	xs := &XHrContractLine869Bfs{}
	if err := c.Read(XHrContractLine869BfModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrContractLine869Bf finds x_hr_contract_line_869bf record by querying it with criteria.
func (c *Client) FindXHrContractLine869Bf(criteria *Criteria) (*XHrContractLine869Bf, error) {
	xs := &XHrContractLine869Bfs{}
	if err := c.SearchRead(XHrContractLine869BfModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_hr_contract_line_869bf was not found with criteria %v", criteria)
}

// FindXHrContractLine869Bfs finds x_hr_contract_line_869bf records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrContractLine869Bfs(criteria *Criteria, options *Options) (*XHrContractLine869Bfs, error) {
	xs := &XHrContractLine869Bfs{}
	if err := c.SearchRead(XHrContractLine869BfModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXHrContractLine869BfIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXHrContractLine869BfIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XHrContractLine869BfModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXHrContractLine869BfId finds record id by querying it with criteria.
func (c *Client) FindXHrContractLine869BfId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XHrContractLine869BfModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_hr_contract_line_869bf was not found with criteria %v and options %v", criteria, options)
}
