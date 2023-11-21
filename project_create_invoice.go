package odoo

import (
	"fmt"
)

// ProjectCreateInvoice represents project.create.invoice model.
type ProjectCreateInvoice struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CandidateOrders *Relation `xmlrpc:"_candidate_orders,omptempty"`
	AmountToInvoice *Float    `xmlrpc:"amount_to_invoice,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	ProjectId       *Many2One `xmlrpc:"project_id,omptempty"`
	SaleOrderId     *Many2One `xmlrpc:"sale_order_id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectCreateInvoices represents array of project.create.invoice model.
type ProjectCreateInvoices []ProjectCreateInvoice

// ProjectCreateInvoiceModel is the odoo model name.
const ProjectCreateInvoiceModel = "project.create.invoice"

// Many2One convert ProjectCreateInvoice to *Many2One.
func (pci *ProjectCreateInvoice) Many2One() *Many2One {
	return NewMany2One(pci.Id.Get(), "")
}

// CreateProjectCreateInvoice creates a new project.create.invoice model and returns its id.
func (c *Client) CreateProjectCreateInvoice(pci *ProjectCreateInvoice) (int64, error) {
	ids, err := c.CreateProjectCreateInvoices([]*ProjectCreateInvoice{pci})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectCreateInvoice creates a new project.create.invoice model and returns its id.
func (c *Client) CreateProjectCreateInvoices(pcis []*ProjectCreateInvoice) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcis {
		vv = append(vv, v)
	}
	return c.Create(ProjectCreateInvoiceModel, vv)
}

// UpdateProjectCreateInvoice updates an existing project.create.invoice record.
func (c *Client) UpdateProjectCreateInvoice(pci *ProjectCreateInvoice) error {
	return c.UpdateProjectCreateInvoices([]int64{pci.Id.Get()}, pci)
}

// UpdateProjectCreateInvoices updates existing project.create.invoice records.
// All records (represented by ids) will be updated by pci values.
func (c *Client) UpdateProjectCreateInvoices(ids []int64, pci *ProjectCreateInvoice) error {
	return c.Update(ProjectCreateInvoiceModel, ids, pci)
}

// DeleteProjectCreateInvoice deletes an existing project.create.invoice record.
func (c *Client) DeleteProjectCreateInvoice(id int64) error {
	return c.DeleteProjectCreateInvoices([]int64{id})
}

// DeleteProjectCreateInvoices deletes existing project.create.invoice records.
func (c *Client) DeleteProjectCreateInvoices(ids []int64) error {
	return c.Delete(ProjectCreateInvoiceModel, ids)
}

// GetProjectCreateInvoice gets project.create.invoice existing record.
func (c *Client) GetProjectCreateInvoice(id int64) (*ProjectCreateInvoice, error) {
	pcis, err := c.GetProjectCreateInvoices([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcis != nil && len(*pcis) > 0 {
		return &((*pcis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.create.invoice not found", id)
}

// GetProjectCreateInvoices gets project.create.invoice existing records.
func (c *Client) GetProjectCreateInvoices(ids []int64) (*ProjectCreateInvoices, error) {
	pcis := &ProjectCreateInvoices{}
	if err := c.Read(ProjectCreateInvoiceModel, ids, nil, pcis); err != nil {
		return nil, err
	}
	return pcis, nil
}

// FindProjectCreateInvoice finds project.create.invoice record by querying it with criteria.
func (c *Client) FindProjectCreateInvoice(criteria *Criteria) (*ProjectCreateInvoice, error) {
	pcis := &ProjectCreateInvoices{}
	if err := c.SearchRead(ProjectCreateInvoiceModel, criteria, NewOptions().Limit(1), pcis); err != nil {
		return nil, err
	}
	if pcis != nil && len(*pcis) > 0 {
		return &((*pcis)[0]), nil
	}
	return nil, fmt.Errorf("project.create.invoice was not found with criteria %v", criteria)
}

// FindProjectCreateInvoices finds project.create.invoice records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateInvoices(criteria *Criteria, options *Options) (*ProjectCreateInvoices, error) {
	pcis := &ProjectCreateInvoices{}
	if err := c.SearchRead(ProjectCreateInvoiceModel, criteria, options, pcis); err != nil {
		return nil, err
	}
	return pcis, nil
}

// FindProjectCreateInvoiceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectCreateInvoiceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectCreateInvoiceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectCreateInvoiceId finds record id by querying it with criteria.
func (c *Client) FindProjectCreateInvoiceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectCreateInvoiceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.create.invoice was not found with criteria %v and options %v", criteria, options)
}
