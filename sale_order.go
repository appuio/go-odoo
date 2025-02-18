package odoo

import (
	"fmt"
)

// SaleOrder represents sale.order model.
type SaleOrder struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                  *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                    *String    `xmlrpc:"access_url,omptempty"`
	AccessWarning                *String    `xmlrpc:"access_warning,omptempty"`
	ActivityCalendarEventId      *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline         *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration  *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon        *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                  *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary              *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon             *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId               *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId               *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AmountTax                    *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTotal                  *Float     `xmlrpc:"amount_total,omptempty"`
	AmountUndiscounted           *Float     `xmlrpc:"amount_undiscounted,omptempty"`
	AmountUntaxed                *Float     `xmlrpc:"amount_untaxed,omptempty"`
	AnalyticAccountId            *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	ArchivedProductCount         *Int       `xmlrpc:"archived_product_count,omptempty"`
	ArchivedProductIds           *Relation  `xmlrpc:"archived_product_ids,omptempty"`
	AuthorizedTransactionIds     *Relation  `xmlrpc:"authorized_transaction_ids,omptempty"`
	AutoGenerated                *Bool      `xmlrpc:"auto_generated,omptempty"`
	AutoPurchaseOrderId          *Many2One  `xmlrpc:"auto_purchase_order_id,omptempty"`
	CampaignId                   *Many2One  `xmlrpc:"campaign_id,omptempty"`
	CartQuantity                 *Int       `xmlrpc:"cart_quantity,omptempty"`
	CartRecoveryEmailSent        *Bool      `xmlrpc:"cart_recovery_email_sent,omptempty"`
	ClientOrderRef               *String    `xmlrpc:"client_order_ref,omptempty"`
	CloseReasonId                *Many2One  `xmlrpc:"close_reason_id,omptempty"`
	CommercialPartnerId          *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CommitmentDate               *Time      `xmlrpc:"commitment_date,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryCode                  *String    `xmlrpc:"country_code,omptempty"`
	CountryId                    *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrencyRate                 *Float     `xmlrpc:"currency_rate,omptempty"`
	DateOrder                    *Time      `xmlrpc:"date_order,omptempty"`
	DeliveryCount                *Int       `xmlrpc:"delivery_count,omptempty"`
	DeliveryStatus               *Selection `xmlrpc:"delivery_status,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	EffectiveDate                *Time      `xmlrpc:"effective_date,omptempty"`
	EndDate                      *Time      `xmlrpc:"end_date,omptempty"`
	ExpectedDate                 *Time      `xmlrpc:"expected_date,omptempty"`
	ExpenseCount                 *Int       `xmlrpc:"expense_count,omptempty"`
	ExpenseIds                   *Relation  `xmlrpc:"expense_ids,omptempty"`
	FailedMessageIds             *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FiscalPositionId             *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	Grid                         *String    `xmlrpc:"grid,omptempty"`
	GridProductTmplId            *Many2One  `xmlrpc:"grid_product_tmpl_id,omptempty"`
	GridUpdate                   *Bool      `xmlrpc:"grid_update,omptempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omptempty"`
	Health                       *Selection `xmlrpc:"health,omptempty"`
	HistoryCount                 *Int       `xmlrpc:"history_count,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	Incoterm                     *Many2One  `xmlrpc:"incoterm,omptempty"`
	IncotermLocation             *String    `xmlrpc:"incoterm_location,omptempty"`
	IndustryId                   *Many2One  `xmlrpc:"industry_id,omptempty"`
	InternalNote                 *String    `xmlrpc:"internal_note,omptempty"`
	InvoiceCount                 *Int       `xmlrpc:"invoice_count,omptempty"`
	InvoiceIds                   *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceStatus                *Selection `xmlrpc:"invoice_status,omptempty"`
	IsAbandonedCart              *Bool      `xmlrpc:"is_abandoned_cart,omptempty"`
	IsBatch                      *Bool      `xmlrpc:"is_batch,omptempty"`
	IsExpired                    *Bool      `xmlrpc:"is_expired,omptempty"`
	IsInvoiceCron                *Bool      `xmlrpc:"is_invoice_cron,omptempty"`
	IsProductMilestone           *Bool      `xmlrpc:"is_product_milestone,omptempty"`
	IsSubscription               *Bool      `xmlrpc:"is_subscription,omptempty"`
	IsUpselling                  *Bool      `xmlrpc:"is_upselling,omptempty"`
	JsonPopover                  *String    `xmlrpc:"json_popover,omptempty"`
	Kpi1MonthMrrDelta            *Float     `xmlrpc:"kpi_1month_mrr_delta,omptempty"`
	Kpi1MonthMrrPercentage       *Float     `xmlrpc:"kpi_1month_mrr_percentage,omptempty"`
	Kpi3MonthsMrrDelta           *Float     `xmlrpc:"kpi_3months_mrr_delta,omptempty"`
	Kpi3MonthsMrrPercentage      *Float     `xmlrpc:"kpi_3months_mrr_percentage,omptempty"`
	L10NDin5008Addresses         *String    `xmlrpc:"l10n_din5008_addresses,omptempty"`
	L10NDin5008DocumentTitle     *String    `xmlrpc:"l10n_din5008_document_title,omptempty"`
	L10NDin5008TemplateData      *String    `xmlrpc:"l10n_din5008_template_data,omptempty"`
	LastInvoiceDate              *Time      `xmlrpc:"last_invoice_date,omptempty"`
	ManageMeteredOrder           *Bool      `xmlrpc:"manage_metered_order,omptempty"`
	MediumId                     *Many2One  `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent               *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId      *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MeteredUsageCount            *Int       `xmlrpc:"metered_usage_count,omptempty"`
	MilestoneCount               *Int       `xmlrpc:"milestone_count,omptempty"`
	MyActivityDateDeadline       *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	NextInvoiceDate              *Time      `xmlrpc:"next_invoice_date,omptempty"`
	Note                         *String    `xmlrpc:"note,omptempty"`
	OnlyServices                 *Bool      `xmlrpc:"only_services,omptempty"`
	OpportunityId                *Many2One  `xmlrpc:"opportunity_id,omptempty"`
	OrderLine                    *Relation  `xmlrpc:"order_line,omptempty"`
	OrderLogIds                  *Relation  `xmlrpc:"order_log_ids,omptempty"`
	Origin                       *String    `xmlrpc:"origin,omptempty"`
	OriginOrderId                *Many2One  `xmlrpc:"origin_order_id,omptempty"`
	PartnerCreditWarning         *String    `xmlrpc:"partner_credit_warning,omptempty"`
	PartnerId                    *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerInvoiceId             *Many2One  `xmlrpc:"partner_invoice_id,omptempty"`
	PartnerShippingId            *Many2One  `xmlrpc:"partner_shipping_id,omptempty"`
	PaymentException             *Bool      `xmlrpc:"payment_exception,omptempty"`
	PaymentTermId                *Many2One  `xmlrpc:"payment_term_id,omptempty"`
	PaymentTokenId               *Many2One  `xmlrpc:"payment_token_id,omptempty"`
	PercentageSatisfaction       *Int       `xmlrpc:"percentage_satisfaction,omptempty"`
	PickingIds                   *Relation  `xmlrpc:"picking_ids,omptempty"`
	PickingPolicy                *Selection `xmlrpc:"picking_policy,omptempty"`
	PlanningFirstSaleLineId      *Many2One  `xmlrpc:"planning_first_sale_line_id,omptempty"`
	PlanningHoursPlanned         *Float     `xmlrpc:"planning_hours_planned,omptempty"`
	PlanningHoursToPlan          *Float     `xmlrpc:"planning_hours_to_plan,omptempty"`
	PlanningInitialDate          *Time      `xmlrpc:"planning_initial_date,omptempty"`
	PricelistId                  *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProcurementGroupId           *Many2One  `xmlrpc:"procurement_group_id,omptempty"`
	ProjectCount                 *Int       `xmlrpc:"project_count,omptempty"`
	ProjectId                    *Many2One  `xmlrpc:"project_id,omptempty"`
	ProjectIds                   *Relation  `xmlrpc:"project_ids,omptempty"`
	PurchaseOrderCount           *Int       `xmlrpc:"purchase_order_count,omptempty"`
	RatingAvg                    *Float     `xmlrpc:"rating_avg,omptempty"`
	RatingAvgText                *Selection `xmlrpc:"rating_avg_text,omptempty"`
	RatingCount                  *Int       `xmlrpc:"rating_count,omptempty"`
	RatingIds                    *Relation  `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback           *String    `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage              *String    `xmlrpc:"rating_last_image,omptempty"`
	RatingLastText               *Selection `xmlrpc:"rating_last_text,omptempty"`
	RatingLastValue              *Float     `xmlrpc:"rating_last_value,omptempty"`
	RatingPercentageSatisfaction *Float     `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RecurrenceId                 *Many2One  `xmlrpc:"recurrence_id,omptempty"`
	RecurringLive                *Bool      `xmlrpc:"recurring_live,omptempty"`
	RecurringMonthly             *Float     `xmlrpc:"recurring_monthly,omptempty"`
	Reference                    *String    `xmlrpc:"reference,omptempty"`
	RenewState                   *Selection `xmlrpc:"renew_state,omptempty"`
	ReportGrids                  *Bool      `xmlrpc:"report_grids,omptempty"`
	RequirePayment               *Bool      `xmlrpc:"require_payment,omptempty"`
	RequireSignature             *Bool      `xmlrpc:"require_signature,omptempty"`
	SaleOrderOptionIds           *Relation  `xmlrpc:"sale_order_option_ids,omptempty"`
	SaleOrderTemplateId          *Many2One  `xmlrpc:"sale_order_template_id,omptempty"`
	ShopWarning                  *String    `xmlrpc:"shop_warning,omptempty"`
	ShowJsonPopover              *Bool      `xmlrpc:"show_json_popover,omptempty"`
	ShowRecInvoiceButton         *Bool      `xmlrpc:"show_rec_invoice_button,omptempty"`
	ShowUpdateFpos               *Bool      `xmlrpc:"show_update_fpos,omptempty"`
	ShowUpdatePricelist          *Bool      `xmlrpc:"show_update_pricelist,omptempty"`
	Signature                    *String    `xmlrpc:"signature,omptempty"`
	SignedBy                     *String    `xmlrpc:"signed_by,omptempty"`
	SignedOn                     *Time      `xmlrpc:"signed_on,omptempty"`
	SourceId                     *Many2One  `xmlrpc:"source_id,omptempty"`
	StageCategory                *Selection `xmlrpc:"stage_category,omptempty"`
	StageId                      *Many2One  `xmlrpc:"stage_id,omptempty"`
	Starred                      *Bool      `xmlrpc:"starred,omptempty"`
	StarredUserIds               *Relation  `xmlrpc:"starred_user_ids,omptempty"`
	StartDate                    *Time      `xmlrpc:"start_date,omptempty"`
	State                        *Selection `xmlrpc:"state,omptempty"`
	SubscriptionChildIds         *Relation  `xmlrpc:"subscription_child_ids,omptempty"`
	SubscriptionId               *Many2One  `xmlrpc:"subscription_id,omptempty"`
	SubscriptionManagement       *Selection `xmlrpc:"subscription_management,omptempty"`
	TagIds                       *Relation  `xmlrpc:"tag_ids,omptempty"`
	TasksCount                   *Int       `xmlrpc:"tasks_count,omptempty"`
	TasksIds                     *Relation  `xmlrpc:"tasks_ids,omptempty"`
	TaxCountryId                 *Many2One  `xmlrpc:"tax_country_id,omptempty"`
	TaxTotals                    *String    `xmlrpc:"tax_totals,omptempty"`
	TeamId                       *Many2One  `xmlrpc:"team_id,omptempty"`
	TeamUserId                   *Many2One  `xmlrpc:"team_user_id,omptempty"`
	TermsType                    *Selection `xmlrpc:"terms_type,omptempty"`
	TicketCount                  *Int       `xmlrpc:"ticket_count,omptempty"`
	TimesheetCount               *Float     `xmlrpc:"timesheet_count,omptempty"`
	TimesheetEncodeUomId         *Many2One  `xmlrpc:"timesheet_encode_uom_id,omptempty"`
	TimesheetIds                 *Relation  `xmlrpc:"timesheet_ids,omptempty"`
	TimesheetTotalDuration       *Int       `xmlrpc:"timesheet_total_duration,omptempty"`
	ToRenew                      *Bool      `xmlrpc:"to_renew,omptempty"`
	TransactionIds               *Relation  `xmlrpc:"transaction_ids,omptempty"`
	TypeName                     *String    `xmlrpc:"type_name,omptempty"`
	UserId                       *Many2One  `xmlrpc:"user_id,omptempty"`
	ValidityDate                 *Time      `xmlrpc:"validity_date,omptempty"`
	VisibleProject               *Bool      `xmlrpc:"visible_project,omptempty"`
	WarehouseId                  *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteDescription           *String    `xmlrpc:"website_description,omptempty"`
	WebsiteId                    *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteOrderLine             *Relation  `xmlrpc:"website_order_line,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
	XStudioReference             *String    `xmlrpc:"x_studio_reference,omptempty"`
}

// SaleOrders represents array of sale.order model.
type SaleOrders []SaleOrder

// SaleOrderModel is the odoo model name.
const SaleOrderModel = "sale.order"

// Many2One convert SaleOrder to *Many2One.
func (so *SaleOrder) Many2One() *Many2One {
	return NewMany2One(so.Id.Get(), "")
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrder(so *SaleOrder) (int64, error) {
	ids, err := c.CreateSaleOrders([]*SaleOrder{so})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrder creates a new sale.order model and returns its id.
func (c *Client) CreateSaleOrders(sos []*SaleOrder) ([]int64, error) {
	var vv []interface{}
	for _, v := range sos {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderModel, vv)
}

// UpdateSaleOrder updates an existing sale.order record.
func (c *Client) UpdateSaleOrder(so *SaleOrder) error {
	return c.UpdateSaleOrders([]int64{so.Id.Get()}, so)
}

// UpdateSaleOrders updates existing sale.order records.
// All records (represented by ids) will be updated by so values.
func (c *Client) UpdateSaleOrders(ids []int64, so *SaleOrder) error {
	return c.Update(SaleOrderModel, ids, so)
}

// DeleteSaleOrder deletes an existing sale.order record.
func (c *Client) DeleteSaleOrder(id int64) error {
	return c.DeleteSaleOrders([]int64{id})
}

// DeleteSaleOrders deletes existing sale.order records.
func (c *Client) DeleteSaleOrders(ids []int64) error {
	return c.Delete(SaleOrderModel, ids)
}

// GetSaleOrder gets sale.order existing record.
func (c *Client) GetSaleOrder(id int64) (*SaleOrder, error) {
	sos, err := c.GetSaleOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if sos != nil && len(*sos) > 0 {
		return &((*sos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order not found", id)
}

// GetSaleOrders gets sale.order existing records.
func (c *Client) GetSaleOrders(ids []int64) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.Read(SaleOrderModel, ids, nil, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrder finds sale.order record by querying it with criteria.
func (c *Client) FindSaleOrder(criteria *Criteria) (*SaleOrder, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, NewOptions().Limit(1), sos); err != nil {
		return nil, err
	}
	if sos != nil && len(*sos) > 0 {
		return &((*sos)[0]), nil
	}
	return nil, fmt.Errorf("sale.order was not found with criteria %v", criteria)
}

// FindSaleOrders finds sale.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrders(criteria *Criteria, options *Options) (*SaleOrders, error) {
	sos := &SaleOrders{}
	if err := c.SearchRead(SaleOrderModel, criteria, options, sos); err != nil {
		return nil, err
	}
	return sos, nil
}

// FindSaleOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order was not found with criteria %v and options %v", criteria, options)
}
