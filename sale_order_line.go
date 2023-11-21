package odoo

import (
	"fmt"
)

// SaleOrderLine represents sale.order.line model.
type SaleOrderLine struct {
	LastUpdate                        *Time       `xmlrpc:"__last_update,omptempty"`
	AccountAnalyticId                 *Many2One   `xmlrpc:"account_analytic_id,omptempty"`
	AnalyticDistribution              interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticDistributionSearch        interface{} `xmlrpc:"analytic_distribution_search,omptempty"`
	AnalyticLineIds                   *Relation   `xmlrpc:"analytic_line_ids,omptempty"`
	AnalyticPrecision                 *Int        `xmlrpc:"analytic_precision,omptempty"`
	CompanyId                         *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate                        *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                         *Many2One   `xmlrpc:"create_uid,omptempty"`
	CurrencyId                        *Many2One   `xmlrpc:"currency_id,omptempty"`
	CustomerLead                      *Float      `xmlrpc:"customer_lead,omptempty"`
	DateOrder                         *Time       `xmlrpc:"date_order,omptempty"`
	Discount                          *Float      `xmlrpc:"discount,omptempty"`
	DisplayName                       *String     `xmlrpc:"display_name,omptempty"`
	DisplayQtyWidget                  *Bool       `xmlrpc:"display_qty_widget,omptempty"`
	DisplayType                       *Selection  `xmlrpc:"display_type,omptempty"`
	ForecastExpectedDate              *Time       `xmlrpc:"forecast_expected_date,omptempty"`
	FreeQtyToday                      *Float      `xmlrpc:"free_qty_today,omptempty"`
	HasDisplayedWarningUpsell         *Bool       `xmlrpc:"has_displayed_warning_upsell,omptempty"`
	Id                                *Int        `xmlrpc:"id,omptempty"`
	InstanceId                        *String     `xmlrpc:"instance_id,omptempty"`
	InstanceIdReadonly                *Bool       `xmlrpc:"instance_id_readonly,omptempty"`
	InvoiceLines                      *Relation   `xmlrpc:"invoice_lines,omptempty"`
	InvoiceStatus                     *Selection  `xmlrpc:"invoice_status,omptempty"`
	IsConfigurableProduct             *Bool       `xmlrpc:"is_configurable_product,omptempty"`
	IsDownpayment                     *Bool       `xmlrpc:"is_downpayment,omptempty"`
	IsExpense                         *Bool       `xmlrpc:"is_expense,omptempty"`
	IsMto                             *Bool       `xmlrpc:"is_mto,omptempty"`
	IsService                         *Bool       `xmlrpc:"is_service,omptempty"`
	LinkedLineId                      *Many2One   `xmlrpc:"linked_line_id,omptempty"`
	MoveIds                           *Relation   `xmlrpc:"move_ids,omptempty"`
	Name                              *String     `xmlrpc:"name,omptempty"`
	NameShort                         *String     `xmlrpc:"name_short,omptempty"`
	OptionLineIds                     *Relation   `xmlrpc:"option_line_ids,omptempty"`
	OrderId                           *Many2One   `xmlrpc:"order_id,omptempty"`
	OrderPartnerId                    *Many2One   `xmlrpc:"order_partner_id,omptempty"`
	ParentLineId                      *Many2One   `xmlrpc:"parent_line_id,omptempty"`
	PlanningHoursPlanned              *Float      `xmlrpc:"planning_hours_planned,omptempty"`
	PlanningHoursToPlan               *Float      `xmlrpc:"planning_hours_to_plan,omptempty"`
	PlanningSlotIds                   *Relation   `xmlrpc:"planning_slot_ids,omptempty"`
	PriceReduce                       *Float      `xmlrpc:"price_reduce,omptempty"`
	PriceReduceTaxexcl                *Float      `xmlrpc:"price_reduce_taxexcl,omptempty"`
	PriceReduceTaxinc                 *Float      `xmlrpc:"price_reduce_taxinc,omptempty"`
	PriceSubtotal                     *Float      `xmlrpc:"price_subtotal,omptempty"`
	PriceTax                          *Float      `xmlrpc:"price_tax,omptempty"`
	PriceTotal                        *Float      `xmlrpc:"price_total,omptempty"`
	PriceUnit                         *Float      `xmlrpc:"price_unit,omptempty"`
	PricelistItemId                   *Many2One   `xmlrpc:"pricelist_item_id,omptempty"`
	PricingId                         *Many2One   `xmlrpc:"pricing_id,omptempty"`
	ProductAddMode                    *Selection  `xmlrpc:"product_add_mode,omptempty"`
	ProductCustomAttributeValueIds    *Relation   `xmlrpc:"product_custom_attribute_value_ids,omptempty"`
	ProductId                         *Many2One   `xmlrpc:"product_id,omptempty"`
	ProductNoVariantAttributeValueIds *Relation   `xmlrpc:"product_no_variant_attribute_value_ids,omptempty"`
	ProductPackagingId                *Many2One   `xmlrpc:"product_packaging_id,omptempty"`
	ProductPackagingQty               *Float      `xmlrpc:"product_packaging_qty,omptempty"`
	ProductTemplateAttributeValueIds  *Relation   `xmlrpc:"product_template_attribute_value_ids,omptempty"`
	ProductTemplateId                 *Many2One   `xmlrpc:"product_template_id,omptempty"`
	ProductType                       *Selection  `xmlrpc:"product_type,omptempty"`
	ProductUom                        *Many2One   `xmlrpc:"product_uom,omptempty"`
	ProductUomCategoryId              *Many2One   `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomQty                     *Float      `xmlrpc:"product_uom_qty,omptempty"`
	ProductUomReadonly                *Bool       `xmlrpc:"product_uom_readonly,omptempty"`
	ProductUpdatable                  *Bool       `xmlrpc:"product_updatable,omptempty"`
	ProjectId                         *Many2One   `xmlrpc:"project_id,omptempty"`
	PurchaseLineCount                 *Int        `xmlrpc:"purchase_line_count,omptempty"`
	PurchaseLineIds                   *Relation   `xmlrpc:"purchase_line_ids,omptempty"`
	QtyAvailableToday                 *Float      `xmlrpc:"qty_available_today,omptempty"`
	QtyDelivered                      *Float      `xmlrpc:"qty_delivered,omptempty"`
	QtyDeliveredMethod                *Selection  `xmlrpc:"qty_delivered_method,omptempty"`
	QtyInvoiced                       *Float      `xmlrpc:"qty_invoiced,omptempty"`
	QtyToDeliver                      *Float      `xmlrpc:"qty_to_deliver,omptempty"`
	QtyToInvoice                      *Float      `xmlrpc:"qty_to_invoice,omptempty"`
	ReachedMilestonesIds              *Relation   `xmlrpc:"reached_milestones_ids,omptempty"`
	RecurringMonthly                  *Float      `xmlrpc:"recurring_monthly,omptempty"`
	RemainingHours                    *Float      `xmlrpc:"remaining_hours,omptempty"`
	RemainingHoursAvailable           *Bool       `xmlrpc:"remaining_hours_available,omptempty"`
	RouteId                           *Many2One   `xmlrpc:"route_id,omptempty"`
	SaleOrderOptionIds                *Relation   `xmlrpc:"sale_order_option_ids,omptempty"`
	SalesmanId                        *Many2One   `xmlrpc:"salesman_id,omptempty"`
	ScheduledDate                     *Time       `xmlrpc:"scheduled_date,omptempty"`
	Sequence                          *Int        `xmlrpc:"sequence,omptempty"`
	ShopWarning                       *String     `xmlrpc:"shop_warning,omptempty"`
	State                             *Selection  `xmlrpc:"state,omptempty"`
	TaskId                            *Many2One   `xmlrpc:"task_id,omptempty"`
	TaxId                             *Relation   `xmlrpc:"tax_id,omptempty"`
	TemporalType                      *Selection  `xmlrpc:"temporal_type,omptempty"`
	TimesheetIds                      *Relation   `xmlrpc:"timesheet_ids,omptempty"`
	UntaxedAmountInvoiced             *Float      `xmlrpc:"untaxed_amount_invoiced,omptempty"`
	UntaxedAmountToInvoice            *Float      `xmlrpc:"untaxed_amount_to_invoice,omptempty"`
	VirtualAvailableAtDate            *Float      `xmlrpc:"virtual_available_at_date,omptempty"`
	WarehouseId                       *Many2One   `xmlrpc:"warehouse_id,omptempty"`
	WebsiteDescription                *String     `xmlrpc:"website_description,omptempty"`
	WriteDate                         *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                          *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// SaleOrderLines represents array of sale.order.line model.
type SaleOrderLines []SaleOrderLine

// SaleOrderLineModel is the odoo model name.
const SaleOrderLineModel = "sale.order.line"

// Many2One convert SaleOrderLine to *Many2One.
func (sol *SaleOrderLine) Many2One() *Many2One {
	return NewMany2One(sol.Id.Get(), "")
}

// CreateSaleOrderLine creates a new sale.order.line model and returns its id.
func (c *Client) CreateSaleOrderLine(sol *SaleOrderLine) (int64, error) {
	ids, err := c.CreateSaleOrderLines([]*SaleOrderLine{sol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderLine creates a new sale.order.line model and returns its id.
func (c *Client) CreateSaleOrderLines(sols []*SaleOrderLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range sols {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderLineModel, vv)
}

// UpdateSaleOrderLine updates an existing sale.order.line record.
func (c *Client) UpdateSaleOrderLine(sol *SaleOrderLine) error {
	return c.UpdateSaleOrderLines([]int64{sol.Id.Get()}, sol)
}

// UpdateSaleOrderLines updates existing sale.order.line records.
// All records (represented by ids) will be updated by sol values.
func (c *Client) UpdateSaleOrderLines(ids []int64, sol *SaleOrderLine) error {
	return c.Update(SaleOrderLineModel, ids, sol)
}

// DeleteSaleOrderLine deletes an existing sale.order.line record.
func (c *Client) DeleteSaleOrderLine(id int64) error {
	return c.DeleteSaleOrderLines([]int64{id})
}

// DeleteSaleOrderLines deletes existing sale.order.line records.
func (c *Client) DeleteSaleOrderLines(ids []int64) error {
	return c.Delete(SaleOrderLineModel, ids)
}

// GetSaleOrderLine gets sale.order.line existing record.
func (c *Client) GetSaleOrderLine(id int64) (*SaleOrderLine, error) {
	sols, err := c.GetSaleOrderLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if sols != nil && len(*sols) > 0 {
		return &((*sols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.line not found", id)
}

// GetSaleOrderLines gets sale.order.line existing records.
func (c *Client) GetSaleOrderLines(ids []int64) (*SaleOrderLines, error) {
	sols := &SaleOrderLines{}
	if err := c.Read(SaleOrderLineModel, ids, nil, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLine finds sale.order.line record by querying it with criteria.
func (c *Client) FindSaleOrderLine(criteria *Criteria) (*SaleOrderLine, error) {
	sols := &SaleOrderLines{}
	if err := c.SearchRead(SaleOrderLineModel, criteria, NewOptions().Limit(1), sols); err != nil {
		return nil, err
	}
	if sols != nil && len(*sols) > 0 {
		return &((*sols)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.line was not found with criteria %v", criteria)
}

// FindSaleOrderLines finds sale.order.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLines(criteria *Criteria, options *Options) (*SaleOrderLines, error) {
	sols := &SaleOrderLines{}
	if err := c.SearchRead(SaleOrderLineModel, criteria, options, sols); err != nil {
		return nil, err
	}
	return sols, nil
}

// FindSaleOrderLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderLineId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.line was not found with criteria %v and options %v", criteria, options)
}
