package odoo

import (
	"fmt"
)

// ProductProduct represents product.product model.
type ProductProduct struct {
	LastUpdate                             *Time      `xmlrpc:"__last_update,omptempty"`
	AccessoryProductIds                    *Relation  `xmlrpc:"accessory_product_ids,omptempty"`
	AccountTagIds                          *Relation  `xmlrpc:"account_tag_ids,omptempty"`
	Active                                 *Bool      `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId                *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline                   *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration            *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                  *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                            *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                          *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                        *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                       *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                         *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                         *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalProductTagIds                *Relation  `xmlrpc:"additional_product_tag_ids,omptempty"`
	AllProductTagIds                       *Relation  `xmlrpc:"all_product_tag_ids,omptempty"`
	AllowOutOfStockOrder                   *Bool      `xmlrpc:"allow_out_of_stock_order,omptempty"`
	AlternativeProductIds                  *Relation  `xmlrpc:"alternative_product_ids,omptempty"`
	AttachmentCount                        *Int       `xmlrpc:"attachment_count,omptempty"`
	AttributeLineIds                       *Relation  `xmlrpc:"attribute_line_ids,omptempty"`
	AvailableThreshold                     *Float     `xmlrpc:"available_threshold,omptempty"`
	AvgCost                                *Float     `xmlrpc:"avg_cost,omptempty"`
	Barcode                                *String    `xmlrpc:"barcode,omptempty"`
	BaseUnitCount                          *Float     `xmlrpc:"base_unit_count,omptempty"`
	BaseUnitId                             *Many2One  `xmlrpc:"base_unit_id,omptempty"`
	BaseUnitName                           *String    `xmlrpc:"base_unit_name,omptempty"`
	BaseUnitPrice                          *Float     `xmlrpc:"base_unit_price,omptempty"`
	CanBeExpensed                          *Bool      `xmlrpc:"can_be_expensed,omptempty"`
	CanImage1024BeZoomed                   *Bool      `xmlrpc:"can_image_1024_be_zoomed,omptempty"`
	CanImageVariant1024BeZoomed            *Bool      `xmlrpc:"can_image_variant_1024_be_zoomed,omptempty"`
	CanPublish                             *Bool      `xmlrpc:"can_publish,omptempty"`
	CategId                                *Many2One  `xmlrpc:"categ_id,omptempty"`
	Code                                   *String    `xmlrpc:"code,omptempty"`
	Color                                  *Int       `xmlrpc:"color,omptempty"`
	CombinationIndices                     *String    `xmlrpc:"combination_indices,omptempty"`
	CompanyCurrencyId                      *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                              *Many2One  `xmlrpc:"company_id,omptempty"`
	CompareListPrice                       *Float     `xmlrpc:"compare_list_price,omptempty"`
	CostCurrencyId                         *Many2One  `xmlrpc:"cost_currency_id,omptempty"`
	CostMethod                             *Selection `xmlrpc:"cost_method,omptempty"`
	CreateDate                             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                             *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultCode                            *String    `xmlrpc:"default_code,omptempty"`
	Description                            *String    `xmlrpc:"description,omptempty"`
	DescriptionPicking                     *String    `xmlrpc:"description_picking,omptempty"`
	DescriptionPickingin                   *String    `xmlrpc:"description_pickingin,omptempty"`
	DescriptionPickingout                  *String    `xmlrpc:"description_pickingout,omptempty"`
	DescriptionPurchase                    *String    `xmlrpc:"description_purchase,omptempty"`
	DescriptionSale                        *String    `xmlrpc:"description_sale,omptempty"`
	DetailedType                           *Selection `xmlrpc:"detailed_type,omptempty"`
	DisplayName                            *String    `xmlrpc:"display_name,omptempty"`
	DisplayPrice                           *String    `xmlrpc:"display_price,omptempty"`
	DocumentsAllowedCompanyId              *Many2One  `xmlrpc:"documents_allowed_company_id,omptempty"`
	EmailTemplateId                        *Many2One  `xmlrpc:"email_template_id,omptempty"`
	ExpensePolicy                          *Selection `xmlrpc:"expense_policy,omptempty"`
	ExpensePolicyTooltip                   *String    `xmlrpc:"expense_policy_tooltip,omptempty"`
	FailedMessageIds                       *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FreeQty                                *Float     `xmlrpc:"free_qty,omptempty"`
	HasAvailableRouteIds                   *Bool      `xmlrpc:"has_available_route_ids,omptempty"`
	HasConfigurableAttributes              *Bool      `xmlrpc:"has_configurable_attributes,omptempty"`
	HasMessage                             *Bool      `xmlrpc:"has_message,omptempty"`
	Id                                     *Int       `xmlrpc:"id,omptempty"`
	Image1024                              *String    `xmlrpc:"image_1024,omptempty"`
	Image128                               *String    `xmlrpc:"image_128,omptempty"`
	Image1920                              *String    `xmlrpc:"image_1920,omptempty"`
	Image256                               *String    `xmlrpc:"image_256,omptempty"`
	Image512                               *String    `xmlrpc:"image_512,omptempty"`
	ImageVariant1024                       *String    `xmlrpc:"image_variant_1024,omptempty"`
	ImageVariant128                        *String    `xmlrpc:"image_variant_128,omptempty"`
	ImageVariant1920                       *String    `xmlrpc:"image_variant_1920,omptempty"`
	ImageVariant256                        *String    `xmlrpc:"image_variant_256,omptempty"`
	ImageVariant512                        *String    `xmlrpc:"image_variant_512,omptempty"`
	IncomingQty                            *Float     `xmlrpc:"incoming_qty,omptempty"`
	InvoicePolicy                          *Selection `xmlrpc:"invoice_policy,omptempty"`
	IsProductVariant                       *Bool      `xmlrpc:"is_product_variant,omptempty"`
	IsPublished                            *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized                         *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	IsTemporal                             *Bool      `xmlrpc:"is_temporal,omptempty"`
	ListPrice                              *Float     `xmlrpc:"list_price,omptempty"`
	LocationId                             *Many2One  `xmlrpc:"location_id,omptempty"`
	LstPrice                               *Float     `xmlrpc:"lst_price,omptempty"`
	MessageAttachmentCount                 *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                         *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                     *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                        *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                 *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                     *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                             *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                      *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId                *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                      *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter               *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                      *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MeteredBillingId                       *String    `xmlrpc:"metered_billing_id,omptempty"`
	MyActivityDateDeadline                 *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                   *String    `xmlrpc:"name,omptempty"`
	NbrMovesIn                             *Int       `xmlrpc:"nbr_moves_in,omptempty"`
	NbrMovesOut                            *Int       `xmlrpc:"nbr_moves_out,omptempty"`
	NbrReorderingRules                     *Int       `xmlrpc:"nbr_reordering_rules,omptempty"`
	OptionalProductIds                     *Relation  `xmlrpc:"optional_product_ids,omptempty"`
	OrderpointIds                          *Relation  `xmlrpc:"orderpoint_ids,omptempty"`
	OutOfStockMessage                      *String    `xmlrpc:"out_of_stock_message,omptempty"`
	OutgoingQty                            *Float     `xmlrpc:"outgoing_qty,omptempty"`
	PackagingIds                           *Relation  `xmlrpc:"packaging_ids,omptempty"`
	PartnerRef                             *String    `xmlrpc:"partner_ref,omptempty"`
	PlanningEnabled                        *Bool      `xmlrpc:"planning_enabled,omptempty"`
	PlanningRoleId                         *Many2One  `xmlrpc:"planning_role_id,omptempty"`
	PriceExtra                             *Float     `xmlrpc:"price_extra,omptempty"`
	PricelistItemCount                     *Int       `xmlrpc:"pricelist_item_count,omptempty"`
	Priority                               *Selection `xmlrpc:"priority,omptempty"`
	ProductAddMode                         *Selection `xmlrpc:"product_add_mode,omptempty"`
	ProductPricingIds                      *Relation  `xmlrpc:"product_pricing_ids,omptempty"`
	ProductTagIds                          *Relation  `xmlrpc:"product_tag_ids,omptempty"`
	ProductTemplateAttributeValueIds       *Relation  `xmlrpc:"product_template_attribute_value_ids,omptempty"`
	ProductTemplateImageIds                *Relation  `xmlrpc:"product_template_image_ids,omptempty"`
	ProductTemplateVariantValueIds         *Relation  `xmlrpc:"product_template_variant_value_ids,omptempty"`
	ProductTmplId                          *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductTooltip                         *String    `xmlrpc:"product_tooltip,omptempty"`
	ProductVariantCount                    *Int       `xmlrpc:"product_variant_count,omptempty"`
	ProductVariantId                       *Many2One  `xmlrpc:"product_variant_id,omptempty"`
	ProductVariantIds                      *Relation  `xmlrpc:"product_variant_ids,omptempty"`
	ProductVariantImageIds                 *Relation  `xmlrpc:"product_variant_image_ids,omptempty"`
	ProjectId                              *Many2One  `xmlrpc:"project_id,omptempty"`
	ProjectTemplateId                      *Many2One  `xmlrpc:"project_template_id,omptempty"`
	ProjectTemplateUseDocuments            *Bool      `xmlrpc:"project_template_use_documents,omptempty"`
	PropertyAccountCreditorPriceDifference *Many2One  `xmlrpc:"property_account_creditor_price_difference,omptempty"`
	PropertyAccountExpenseId               *Many2One  `xmlrpc:"property_account_expense_id,omptempty"`
	PropertyAccountIncomeId                *Many2One  `xmlrpc:"property_account_income_id,omptempty"`
	PropertyStockInventory                 *Many2One  `xmlrpc:"property_stock_inventory,omptempty"`
	PropertyStockProduction                *Many2One  `xmlrpc:"property_stock_production,omptempty"`
	PublicCategIds                         *Relation  `xmlrpc:"public_categ_ids,omptempty"`
	PurchaseLineWarn                       *Selection `xmlrpc:"purchase_line_warn,omptempty"`
	PurchaseLineWarnMsg                    *String    `xmlrpc:"purchase_line_warn_msg,omptempty"`
	PurchaseMethod                         *Selection `xmlrpc:"purchase_method,omptempty"`
	PurchaseOk                             *Bool      `xmlrpc:"purchase_ok,omptempty"`
	PurchaseOrderLineIds                   *Relation  `xmlrpc:"purchase_order_line_ids,omptempty"`
	PurchasedProductQty                    *Float     `xmlrpc:"purchased_product_qty,omptempty"`
	PutawayRuleIds                         *Relation  `xmlrpc:"putaway_rule_ids,omptempty"`
	QtyAvailable                           *Float     `xmlrpc:"qty_available,omptempty"`
	QuantitySvl                            *Float     `xmlrpc:"quantity_svl,omptempty"`
	QuotationDescription                   *String    `xmlrpc:"quotation_description,omptempty"`
	QuotationOnlyDescription               *String    `xmlrpc:"quotation_only_description,omptempty"`
	RatingAvg                              *Float     `xmlrpc:"rating_avg,omptempty"`
	RatingAvgText                          *Selection `xmlrpc:"rating_avg_text,omptempty"`
	RatingCount                            *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                              *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback                     *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage                        *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastText                         *Selection `xmlrpc:"rating_last_text,omptempty"`
	RatingLastValue                        *Float     `xmlrpc:"rating_last_value,omptempty"`
	RatingPercentageSatisfaction           *Float     `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RecurringInvoice                       *Bool      `xmlrpc:"recurring_invoice,omptempty"`
	ReorderingMaxQty                       *Float     `xmlrpc:"reordering_max_qty,omptempty"`
	ReorderingMinQty                       *Float     `xmlrpc:"reordering_min_qty,omptempty"`
	ResponsibleId                          *Many2One  `xmlrpc:"responsible_id,omptempty"`
	RouteFromCategIds                      *Relation  `xmlrpc:"route_from_categ_ids,omptempty"`
	RouteIds                               *Relation  `xmlrpc:"route_ids,omptempty"`
	SaleDelay                              *Float     `xmlrpc:"sale_delay,omptempty"`
	SaleLineWarn                           *Selection `xmlrpc:"sale_line_warn,omptempty"`
	SaleLineWarnMsg                        *String    `xmlrpc:"sale_line_warn_msg,omptempty"`
	SaleLinesCount                         *Int       `xmlrpc:"sale_lines_count,omptempty"`
	SaleOk                                 *Bool      `xmlrpc:"sale_ok,omptempty"`
	SalesCount                             *Float     `xmlrpc:"sales_count,omptempty"`
	SellerIds                              *Relation  `xmlrpc:"seller_ids,omptempty"`
	SeoName                                *String    `xmlrpc:"seo_name,omptempty"`
	Sequence                               *Int       `xmlrpc:"sequence,omptempty"`
	ServicePolicy                          *Selection `xmlrpc:"service_policy,omptempty"`
	ServiceToPurchase                      *Bool      `xmlrpc:"service_to_purchase,omptempty"`
	ServiceTracking                        *Selection `xmlrpc:"service_tracking,omptempty"`
	ServiceType                            *Selection `xmlrpc:"service_type,omptempty"`
	ServiceUpsellThreshold                 *Float     `xmlrpc:"service_upsell_threshold,omptempty"`
	ServiceUpsellThresholdRatio            *String    `xmlrpc:"service_upsell_threshold_ratio,omptempty"`
	ShowAvailability                       *Bool      `xmlrpc:"show_availability,omptempty"`
	ShowForecastedQtyStatusButton          *Bool      `xmlrpc:"show_forecasted_qty_status_button,omptempty"`
	ShowOnHandQtyStatusButton              *Bool      `xmlrpc:"show_on_hand_qty_status_button,omptempty"`
	SlaId                                  *Many2One  `xmlrpc:"sla_id,omptempty"`
	StandardPrice                          *Float     `xmlrpc:"standard_price,omptempty"`
	StockMoveIds                           *Relation  `xmlrpc:"stock_move_ids,omptempty"`
	StockNotificationPartnerIds            *Relation  `xmlrpc:"stock_notification_partner_ids,omptempty"`
	StockQuantIds                          *Relation  `xmlrpc:"stock_quant_ids,omptempty"`
	StockValuationLayerIds                 *Relation  `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	StorageCategoryCapacityIds             *Relation  `xmlrpc:"storage_category_capacity_ids,omptempty"`
	SupplierTaxesId                        *Relation  `xmlrpc:"supplier_taxes_id,omptempty"`
	TaxString                              *String    `xmlrpc:"tax_string,omptempty"`
	TaxesId                                *Relation  `xmlrpc:"taxes_id,omptempty"`
	TemplateFolderId                       *Many2One  `xmlrpc:"template_folder_id,omptempty"`
	TotalValue                             *Float     `xmlrpc:"total_value,omptempty"`
	Tracking                               *Selection `xmlrpc:"tracking,omptempty"`
	Type                                   *Selection `xmlrpc:"type,omptempty"`
	UomId                                  *Many2One  `xmlrpc:"uom_id,omptempty"`
	UomName                                *String    `xmlrpc:"uom_name,omptempty"`
	UomPoId                                *Many2One  `xmlrpc:"uom_po_id,omptempty"`
	ValidEan                               *Bool      `xmlrpc:"valid_ean,omptempty"`
	ValidProductTemplateAttributeLineIds   *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omptempty"`
	Valuation                              *Selection `xmlrpc:"valuation,omptempty"`
	ValueSvl                               *Float     `xmlrpc:"value_svl,omptempty"`
	VariantSellerIds                       *Relation  `xmlrpc:"variant_seller_ids,omptempty"`
	VirtualAvailable                       *Float     `xmlrpc:"virtual_available,omptempty"`
	VisibleExpensePolicy                   *Bool      `xmlrpc:"visible_expense_policy,omptempty"`
	VisibleQtyConfigurator                 *Bool      `xmlrpc:"visible_qty_configurator,omptempty"`
	Volume                                 *Float     `xmlrpc:"volume,omptempty"`
	VolumeUomName                          *String    `xmlrpc:"volume_uom_name,omptempty"`
	WarehouseId                            *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteDescription                     *String    `xmlrpc:"website_description,omptempty"`
	WebsiteId                              *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds                      *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription                 *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords                    *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg                       *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle                       *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished                       *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteRibbonId                        *Many2One  `xmlrpc:"website_ribbon_id,omptempty"`
	WebsiteSequence                        *Int       `xmlrpc:"website_sequence,omptempty"`
	WebsiteSizeX                           *Int       `xmlrpc:"website_size_x,omptempty"`
	WebsiteSizeY                           *Int       `xmlrpc:"website_size_y,omptempty"`
	WebsiteUrl                             *String    `xmlrpc:"website_url,omptempty"`
	Weight                                 *Float     `xmlrpc:"weight,omptempty"`
	WeightUomName                          *String    `xmlrpc:"weight_uom_name,omptempty"`
	WriteDate                              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                               *Many2One  `xmlrpc:"write_uid,omptempty"`
	XAccountableTeam                       *Many2One  `xmlrpc:"x_accountable_team,omptempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	ids, err := c.CreateProductProducts([]*ProductProduct{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProducts(pps []*ProductProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductProductModel, vv)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.product not found", id)
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.product was not found with criteria %v", criteria)
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.product was not found with criteria %v and options %v", criteria, options)
}
