package odoo

import (
	"fmt"
)

// ProjectCreateSaleOrder represents project.create.sale.order model.
type ProjectCreateSaleOrder struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	CommercialPartnerId *Many2One `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId           *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	InfoInvoice         *String   `xmlrpc:"info_invoice,omptempty"`
	LineIds             *Relation `xmlrpc:"line_ids,omptempty"`
	PartnerId           *Many2One `xmlrpc:"partner_id,omptempty"`
	ProjectId           *Many2One `xmlrpc:"project_id,omptempty"`
	SaleOrderId         *Many2One `xmlrpc:"sale_order_id,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectCreateSaleOrders represents array of project.create.sale.order model.
type ProjectCreateSaleOrders []ProjectCreateSaleOrder

// ProjectCreateSaleOrderModel is the odoo model name.
const ProjectCreateSaleOrderModel = "project.create.sale.order"

// Many2One convert ProjectCreateSaleOrder to *Many2One.
func (pcso *ProjectCreateSaleOrder) Many2One() *Many2One {
	return NewMany2One(pcso.Id.Get(), "")
}

// CreateProjectCreateSaleOrder creates a new project.create.sale.order model and returns its id.
func (c *Client) CreateProjectCreateSaleOrder(pcso *ProjectCreateSaleOrder) (int64, error) {
	ids, err := c.CreateProjectCreateSaleOrders([]*ProjectCreateSaleOrder{pcso})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectCreateSaleOrder creates a new project.create.sale.order model and returns its id.
func (c *Client) CreateProjectCreateSaleOrders(pcsos []*ProjectCreateSaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcsos {
		vv = append(vv, v)
	}
	return c.Create(ProjectCreateSaleOrderModel, vv)
}

// UpdateProjectCreateSaleOrder updates an existing project.create.sale.order record.
func (c *Client) UpdateProjectCreateSaleOrder(pcso *ProjectCreateSaleOrder) error {
	return c.UpdateProjectCreateSaleOrders([]int64{pcso.Id.Get()}, pcso)
}

// UpdateProjectCreateSaleOrders updates existing project.create.sale.order records.
// All records (represented by ids) will be updated by pcso values.
func (c *Client) UpdateProjectCreateSaleOrders(ids []int64, pcso *ProjectCreateSaleOrder) error {
	return c.Update(ProjectCreateSaleOrderModel, ids, pcso)
}

// DeleteProjectCreateSaleOrder deletes an existing project.create.sale.order record.
func (c *Client) DeleteProjectCreateSaleOrder(id int64) error {
	return c.DeleteProjectCreateSaleOrders([]int64{id})
}

// DeleteProjectCreateSaleOrders deletes existing project.create.sale.order records.
func (c *Client) DeleteProjectCreateSaleOrders(ids []int64) error {
	return c.Delete(ProjectCreateSaleOrderModel, ids)
}

// GetProjectCreateSaleOrder gets project.create.sale.order existing record.
func (c *Client) GetProjectCreateSaleOrder(id int64) (*ProjectCreateSaleOrder, error) {
	pcsos, err := c.GetProjectCreateSaleOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcsos != nil && len(*pcsos) > 0 {
		return &((*pcsos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.create.sale.order not found", id)
}

// GetProjectCreateSaleOrders gets project.create.sale.order existing records.
func (c *Client) GetProjectCreateSaleOrders(ids []int64) (*ProjectCreateSaleOrders, error) {
	pcsos := &ProjectCreateSaleOrders{}
	if err := c.Read(ProjectCreateSaleOrderModel, ids, nil, pcsos); err != nil {
		return nil, err
	}
	return pcsos, nil
}

// FindProjectCreateSaleOrder finds project.create.sale.order record by querying it with criteria.
func (c *Client) FindProjectCreateSaleOrder(criteria *Criteria) (*ProjectCreateSaleOrder, error) {
	pcsos := &ProjectCreateSaleOrders{}
	if err := c.SearchRead(ProjectCreateSaleOrderModel, criteria, NewOptions().Limit(1), pcsos); err != nil {
		return nil, err
	}
	if pcsos != nil && len(*pcsos) > 0 {
		return &((*pcsos)[0]), nil
	}
	return nil, fmt.Errorf("project.create.sale.order was not found with criteria %v", criteria)
}

// FindProjectCreateSaleOrders finds project.create.sale.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateSaleOrders(criteria *Criteria, options *Options) (*ProjectCreateSaleOrders, error) {
	pcsos := &ProjectCreateSaleOrders{}
	if err := c.SearchRead(ProjectCreateSaleOrderModel, criteria, options, pcsos); err != nil {
		return nil, err
	}
	return pcsos, nil
}

// FindProjectCreateSaleOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateSaleOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectCreateSaleOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectCreateSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindProjectCreateSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectCreateSaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.create.sale.order was not found with criteria %v and options %v", criteria, options)
}
