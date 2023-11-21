package odoo

import (
	"fmt"
)

// ProjectCreateSaleOrderLine represents project.create.sale.order.line model.
type ProjectCreateSaleOrderLine struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId  *Many2One `xmlrpc:"employee_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	PriceUnit   *Float    `xmlrpc:"price_unit,omptempty"`
	ProductId   *Many2One `xmlrpc:"product_id,omptempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectCreateSaleOrderLines represents array of project.create.sale.order.line model.
type ProjectCreateSaleOrderLines []ProjectCreateSaleOrderLine

// ProjectCreateSaleOrderLineModel is the odoo model name.
const ProjectCreateSaleOrderLineModel = "project.create.sale.order.line"

// Many2One convert ProjectCreateSaleOrderLine to *Many2One.
func (pcsol *ProjectCreateSaleOrderLine) Many2One() *Many2One {
	return NewMany2One(pcsol.Id.Get(), "")
}

// CreateProjectCreateSaleOrderLine creates a new project.create.sale.order.line model and returns its id.
func (c *Client) CreateProjectCreateSaleOrderLine(pcsol *ProjectCreateSaleOrderLine) (int64, error) {
	ids, err := c.CreateProjectCreateSaleOrderLines([]*ProjectCreateSaleOrderLine{pcsol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectCreateSaleOrderLine creates a new project.create.sale.order.line model and returns its id.
func (c *Client) CreateProjectCreateSaleOrderLines(pcsols []*ProjectCreateSaleOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcsols {
		vv = append(vv, v)
	}
	return c.Create(ProjectCreateSaleOrderLineModel, vv)
}

// UpdateProjectCreateSaleOrderLine updates an existing project.create.sale.order.line record.
func (c *Client) UpdateProjectCreateSaleOrderLine(pcsol *ProjectCreateSaleOrderLine) error {
	return c.UpdateProjectCreateSaleOrderLines([]int64{pcsol.Id.Get()}, pcsol)
}

// UpdateProjectCreateSaleOrderLines updates existing project.create.sale.order.line records.
// All records (represented by ids) will be updated by pcsol values.
func (c *Client) UpdateProjectCreateSaleOrderLines(ids []int64, pcsol *ProjectCreateSaleOrderLine) error {
	return c.Update(ProjectCreateSaleOrderLineModel, ids, pcsol)
}

// DeleteProjectCreateSaleOrderLine deletes an existing project.create.sale.order.line record.
func (c *Client) DeleteProjectCreateSaleOrderLine(id int64) error {
	return c.DeleteProjectCreateSaleOrderLines([]int64{id})
}

// DeleteProjectCreateSaleOrderLines deletes existing project.create.sale.order.line records.
func (c *Client) DeleteProjectCreateSaleOrderLines(ids []int64) error {
	return c.Delete(ProjectCreateSaleOrderLineModel, ids)
}

// GetProjectCreateSaleOrderLine gets project.create.sale.order.line existing record.
func (c *Client) GetProjectCreateSaleOrderLine(id int64) (*ProjectCreateSaleOrderLine, error) {
	pcsols, err := c.GetProjectCreateSaleOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcsols != nil && len(*pcsols) > 0 {
		return &((*pcsols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.create.sale.order.line not found", id)
}

// GetProjectCreateSaleOrderLines gets project.create.sale.order.line existing records.
func (c *Client) GetProjectCreateSaleOrderLines(ids []int64) (*ProjectCreateSaleOrderLines, error) {
	pcsols := &ProjectCreateSaleOrderLines{}
	if err := c.Read(ProjectCreateSaleOrderLineModel, ids, nil, pcsols); err != nil {
		return nil, err
	}
	return pcsols, nil
}

// FindProjectCreateSaleOrderLine finds project.create.sale.order.line record by querying it with criteria.
func (c *Client) FindProjectCreateSaleOrderLine(criteria *Criteria) (*ProjectCreateSaleOrderLine, error) {
	pcsols := &ProjectCreateSaleOrderLines{}
	if err := c.SearchRead(ProjectCreateSaleOrderLineModel, criteria, NewOptions().Limit(1), pcsols); err != nil {
		return nil, err
	}
	if pcsols != nil && len(*pcsols) > 0 {
		return &((*pcsols)[0]), nil
	}
	return nil, fmt.Errorf("project.create.sale.order.line was not found with criteria %v", criteria)
}

// FindProjectCreateSaleOrderLines finds project.create.sale.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateSaleOrderLines(criteria *Criteria, options *Options) (*ProjectCreateSaleOrderLines, error) {
	pcsols := &ProjectCreateSaleOrderLines{}
	if err := c.SearchRead(ProjectCreateSaleOrderLineModel, criteria, options, pcsols); err != nil {
		return nil, err
	}
	return pcsols, nil
}

// FindProjectCreateSaleOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateSaleOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectCreateSaleOrderLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectCreateSaleOrderLineId finds record id by querying it with criteria.
func (c *Client) FindProjectCreateSaleOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectCreateSaleOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.create.sale.order.line was not found with criteria %v and options %v", criteria, options)
}
