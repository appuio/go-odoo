package odoo

import (
	"fmt"
)

// SaleSubscriptionReport represents sale.subscription.report model.
type SaleSubscriptionReport struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	AnalyticAccountId   *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CategId             *Many2One  `xmlrpc:"categ_id,omptempty"`
	CloseReasonId       *Many2One  `xmlrpc:"close_reason_id,omptempty"`
	CommercialPartnerId *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omptempty"`
	DateOrder           *Time      `xmlrpc:"date_order,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	EndDate             *Time      `xmlrpc:"end_date,omptempty"`
	Health              *Selection `xmlrpc:"health,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	IndustryId          *Many2One  `xmlrpc:"industry_id,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	PricelistId         *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductTmplId       *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductUom          *Many2One  `xmlrpc:"product_uom,omptempty"`
	Quantity            *Float     `xmlrpc:"quantity,omptempty"`
	RecurringMonthly    *Float     `xmlrpc:"recurring_monthly,omptempty"`
	RecurringTotal      *Float     `xmlrpc:"recurring_total,omptempty"`
	RecurringYearly     *Float     `xmlrpc:"recurring_yearly,omptempty"`
	StageCategory       *Selection `xmlrpc:"stage_category,omptempty"`
	StageId             *Many2One  `xmlrpc:"stage_id,omptempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omptempty"`
	TemplateId          *Many2One  `xmlrpc:"template_id,omptempty"`
	ToRenew             *Bool      `xmlrpc:"to_renew,omptempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omptempty"`
}

// SaleSubscriptionReports represents array of sale.subscription.report model.
type SaleSubscriptionReports []SaleSubscriptionReport

// SaleSubscriptionReportModel is the odoo model name.
const SaleSubscriptionReportModel = "sale.subscription.report"

// Many2One convert SaleSubscriptionReport to *Many2One.
func (ssr *SaleSubscriptionReport) Many2One() *Many2One {
	return NewMany2One(ssr.Id.Get(), "")
}

// CreateSaleSubscriptionReport creates a new sale.subscription.report model and returns its id.
func (c *Client) CreateSaleSubscriptionReport(ssr *SaleSubscriptionReport) (int64, error) {
	ids, err := c.CreateSaleSubscriptionReports([]*SaleSubscriptionReport{ssr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleSubscriptionReport creates a new sale.subscription.report model and returns its id.
func (c *Client) CreateSaleSubscriptionReports(ssrs []*SaleSubscriptionReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssrs {
		vv = append(vv, v)
	}
	return c.Create(SaleSubscriptionReportModel, vv)
}

// UpdateSaleSubscriptionReport updates an existing sale.subscription.report record.
func (c *Client) UpdateSaleSubscriptionReport(ssr *SaleSubscriptionReport) error {
	return c.UpdateSaleSubscriptionReports([]int64{ssr.Id.Get()}, ssr)
}

// UpdateSaleSubscriptionReports updates existing sale.subscription.report records.
// All records (represented by ids) will be updated by ssr values.
func (c *Client) UpdateSaleSubscriptionReports(ids []int64, ssr *SaleSubscriptionReport) error {
	return c.Update(SaleSubscriptionReportModel, ids, ssr)
}

// DeleteSaleSubscriptionReport deletes an existing sale.subscription.report record.
func (c *Client) DeleteSaleSubscriptionReport(id int64) error {
	return c.DeleteSaleSubscriptionReports([]int64{id})
}

// DeleteSaleSubscriptionReports deletes existing sale.subscription.report records.
func (c *Client) DeleteSaleSubscriptionReports(ids []int64) error {
	return c.Delete(SaleSubscriptionReportModel, ids)
}

// GetSaleSubscriptionReport gets sale.subscription.report existing record.
func (c *Client) GetSaleSubscriptionReport(id int64) (*SaleSubscriptionReport, error) {
	ssrs, err := c.GetSaleSubscriptionReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssrs != nil && len(*ssrs) > 0 {
		return &((*ssrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.subscription.report not found", id)
}

// GetSaleSubscriptionReports gets sale.subscription.report existing records.
func (c *Client) GetSaleSubscriptionReports(ids []int64) (*SaleSubscriptionReports, error) {
	ssrs := &SaleSubscriptionReports{}
	if err := c.Read(SaleSubscriptionReportModel, ids, nil, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSaleSubscriptionReport finds sale.subscription.report record by querying it with criteria.
func (c *Client) FindSaleSubscriptionReport(criteria *Criteria) (*SaleSubscriptionReport, error) {
	ssrs := &SaleSubscriptionReports{}
	if err := c.SearchRead(SaleSubscriptionReportModel, criteria, NewOptions().Limit(1), ssrs); err != nil {
		return nil, err
	}
	if ssrs != nil && len(*ssrs) > 0 {
		return &((*ssrs)[0]), nil
	}
	return nil, fmt.Errorf("sale.subscription.report was not found with criteria %v", criteria)
}

// FindSaleSubscriptionReports finds sale.subscription.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionReports(criteria *Criteria, options *Options) (*SaleSubscriptionReports, error) {
	ssrs := &SaleSubscriptionReports{}
	if err := c.SearchRead(SaleSubscriptionReportModel, criteria, options, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSaleSubscriptionReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleSubscriptionReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleSubscriptionReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleSubscriptionReportId finds record id by querying it with criteria.
func (c *Client) FindSaleSubscriptionReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleSubscriptionReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.subscription.report was not found with criteria %v and options %v", criteria, options)
}
