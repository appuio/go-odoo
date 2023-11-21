package odoo

import (
	"fmt"
)

// ProjectSaleLineEmployeeMap represents project.sale.line.employee.map model.
type ProjectSaleLineEmployeeMap struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omptempty"`
	Cost           *Float    `xmlrpc:"cost,omptempty"`
	CostCurrencyId *Many2One `xmlrpc:"cost_currency_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId     *Many2One `xmlrpc:"currency_id,omptempty"`
	DisplayCost    *Float    `xmlrpc:"display_cost,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId     *Many2One `xmlrpc:"employee_id,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	IsCostChanged  *Bool     `xmlrpc:"is_cost_changed,omptempty"`
	PartnerId      *Many2One `xmlrpc:"partner_id,omptempty"`
	PriceUnit      *Float    `xmlrpc:"price_unit,omptempty"`
	ProjectId      *Many2One `xmlrpc:"project_id,omptempty"`
	SaleLineId     *Many2One `xmlrpc:"sale_line_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProjectSaleLineEmployeeMaps represents array of project.sale.line.employee.map model.
type ProjectSaleLineEmployeeMaps []ProjectSaleLineEmployeeMap

// ProjectSaleLineEmployeeMapModel is the odoo model name.
const ProjectSaleLineEmployeeMapModel = "project.sale.line.employee.map"

// Many2One convert ProjectSaleLineEmployeeMap to *Many2One.
func (pslem *ProjectSaleLineEmployeeMap) Many2One() *Many2One {
	return NewMany2One(pslem.Id.Get(), "")
}

// CreateProjectSaleLineEmployeeMap creates a new project.sale.line.employee.map model and returns its id.
func (c *Client) CreateProjectSaleLineEmployeeMap(pslem *ProjectSaleLineEmployeeMap) (int64, error) {
	ids, err := c.CreateProjectSaleLineEmployeeMaps([]*ProjectSaleLineEmployeeMap{pslem})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectSaleLineEmployeeMap creates a new project.sale.line.employee.map model and returns its id.
func (c *Client) CreateProjectSaleLineEmployeeMaps(pslems []*ProjectSaleLineEmployeeMap) ([]int64, error) {
	var vv []interface{}
	for _, v := range pslems {
		vv = append(vv, v)
	}
	return c.Create(ProjectSaleLineEmployeeMapModel, vv)
}

// UpdateProjectSaleLineEmployeeMap updates an existing project.sale.line.employee.map record.
func (c *Client) UpdateProjectSaleLineEmployeeMap(pslem *ProjectSaleLineEmployeeMap) error {
	return c.UpdateProjectSaleLineEmployeeMaps([]int64{pslem.Id.Get()}, pslem)
}

// UpdateProjectSaleLineEmployeeMaps updates existing project.sale.line.employee.map records.
// All records (represented by ids) will be updated by pslem values.
func (c *Client) UpdateProjectSaleLineEmployeeMaps(ids []int64, pslem *ProjectSaleLineEmployeeMap) error {
	return c.Update(ProjectSaleLineEmployeeMapModel, ids, pslem)
}

// DeleteProjectSaleLineEmployeeMap deletes an existing project.sale.line.employee.map record.
func (c *Client) DeleteProjectSaleLineEmployeeMap(id int64) error {
	return c.DeleteProjectSaleLineEmployeeMaps([]int64{id})
}

// DeleteProjectSaleLineEmployeeMaps deletes existing project.sale.line.employee.map records.
func (c *Client) DeleteProjectSaleLineEmployeeMaps(ids []int64) error {
	return c.Delete(ProjectSaleLineEmployeeMapModel, ids)
}

// GetProjectSaleLineEmployeeMap gets project.sale.line.employee.map existing record.
func (c *Client) GetProjectSaleLineEmployeeMap(id int64) (*ProjectSaleLineEmployeeMap, error) {
	pslems, err := c.GetProjectSaleLineEmployeeMaps([]int64{id})
	if err != nil {
		return nil, err
	}
	if pslems != nil && len(*pslems) > 0 {
		return &((*pslems)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.sale.line.employee.map not found", id)
}

// GetProjectSaleLineEmployeeMaps gets project.sale.line.employee.map existing records.
func (c *Client) GetProjectSaleLineEmployeeMaps(ids []int64) (*ProjectSaleLineEmployeeMaps, error) {
	pslems := &ProjectSaleLineEmployeeMaps{}
	if err := c.Read(ProjectSaleLineEmployeeMapModel, ids, nil, pslems); err != nil {
		return nil, err
	}
	return pslems, nil
}

// FindProjectSaleLineEmployeeMap finds project.sale.line.employee.map record by querying it with criteria.
func (c *Client) FindProjectSaleLineEmployeeMap(criteria *Criteria) (*ProjectSaleLineEmployeeMap, error) {
	pslems := &ProjectSaleLineEmployeeMaps{}
	if err := c.SearchRead(ProjectSaleLineEmployeeMapModel, criteria, NewOptions().Limit(1), pslems); err != nil {
		return nil, err
	}
	if pslems != nil && len(*pslems) > 0 {
		return &((*pslems)[0]), nil
	}
	return nil, fmt.Errorf("project.sale.line.employee.map was not found with criteria %v", criteria)
}

// FindProjectSaleLineEmployeeMaps finds project.sale.line.employee.map records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectSaleLineEmployeeMaps(criteria *Criteria, options *Options) (*ProjectSaleLineEmployeeMaps, error) {
	pslems := &ProjectSaleLineEmployeeMaps{}
	if err := c.SearchRead(ProjectSaleLineEmployeeMapModel, criteria, options, pslems); err != nil {
		return nil, err
	}
	return pslems, nil
}

// FindProjectSaleLineEmployeeMapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectSaleLineEmployeeMapIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectSaleLineEmployeeMapModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectSaleLineEmployeeMapId finds record id by querying it with criteria.
func (c *Client) FindProjectSaleLineEmployeeMapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectSaleLineEmployeeMapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.sale.line.employee.map was not found with criteria %v and options %v", criteria, options)
}
