package odoo

import (
	"fmt"
)

// StockReport represents stock.report model.
type StockReport struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	CategId         *Many2One  `xmlrpc:"categ_id,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	CreationDate    *Time      `xmlrpc:"creation_date,omptempty"`
	CycleTime       *Float     `xmlrpc:"cycle_time,omptempty"`
	DateDone        *Time      `xmlrpc:"date_done,omptempty"`
	Delay           *Float     `xmlrpc:"delay,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	IsBackorder     *Bool      `xmlrpc:"is_backorder,omptempty"`
	IsLate          *Bool      `xmlrpc:"is_late,omptempty"`
	OperationType   *String    `xmlrpc:"operation_type,omptempty"`
	PartnerId       *Many2One  `xmlrpc:"partner_id,omptempty"`
	PickingId       *Many2One  `xmlrpc:"picking_id,omptempty"`
	PickingName     *String    `xmlrpc:"picking_name,omptempty"`
	PickingTypeCode *Selection `xmlrpc:"picking_type_code,omptempty"`
	ProductId       *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductQty      *Float     `xmlrpc:"product_qty,omptempty"`
	Reference       *String    `xmlrpc:"reference,omptempty"`
	ScheduledDate   *Time      `xmlrpc:"scheduled_date,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
	StockValue      *Float     `xmlrpc:"stock_value,omptempty"`
	Valuation       *Float     `xmlrpc:"valuation,omptempty"`
}

// StockReports represents array of stock.report model.
type StockReports []StockReport

// StockReportModel is the odoo model name.
const StockReportModel = "stock.report"

// Many2One convert StockReport to *Many2One.
func (sr *StockReport) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateStockReport creates a new stock.report model and returns its id.
func (c *Client) CreateStockReport(sr *StockReport) (int64, error) {
	ids, err := c.CreateStockReports([]*StockReport{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockReport creates a new stock.report model and returns its id.
func (c *Client) CreateStockReports(srs []*StockReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(StockReportModel, vv)
}

// UpdateStockReport updates an existing stock.report record.
func (c *Client) UpdateStockReport(sr *StockReport) error {
	return c.UpdateStockReports([]int64{sr.Id.Get()}, sr)
}

// UpdateStockReports updates existing stock.report records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateStockReports(ids []int64, sr *StockReport) error {
	return c.Update(StockReportModel, ids, sr)
}

// DeleteStockReport deletes an existing stock.report record.
func (c *Client) DeleteStockReport(id int64) error {
	return c.DeleteStockReports([]int64{id})
}

// DeleteStockReports deletes existing stock.report records.
func (c *Client) DeleteStockReports(ids []int64) error {
	return c.Delete(StockReportModel, ids)
}

// GetStockReport gets stock.report existing record.
func (c *Client) GetStockReport(id int64) (*StockReport, error) {
	srs, err := c.GetStockReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.report not found", id)
}

// GetStockReports gets stock.report existing records.
func (c *Client) GetStockReports(ids []int64) (*StockReports, error) {
	srs := &StockReports{}
	if err := c.Read(StockReportModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockReport finds stock.report record by querying it with criteria.
func (c *Client) FindStockReport(criteria *Criteria) (*StockReport, error) {
	srs := &StockReports{}
	if err := c.SearchRead(StockReportModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("stock.report was not found with criteria %v", criteria)
}

// FindStockReports finds stock.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReports(criteria *Criteria, options *Options) (*StockReports, error) {
	srs := &StockReports{}
	if err := c.SearchRead(StockReportModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockReportId finds record id by querying it with criteria.
func (c *Client) FindStockReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.report was not found with criteria %v and options %v", criteria, options)
}
