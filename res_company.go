package odoo

import (
	"fmt"
)

// ResCompany represents res.company model.
type ResCompany struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omptempty"`
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omptempty"`
	AccountDashboardOnboardingState             *Selection `xmlrpc:"account_dashboard_onboarding_state,omptempty"`
	AccountDefaultPosReceivableAccountId        *Many2One  `xmlrpc:"account_default_pos_receivable_account_id,omptempty"`
	AccountDisplayRepresentativeField           *Bool      `xmlrpc:"account_display_representative_field,omptempty"`
	AccountEnabledTaxCountryIds                 *Relation  `xmlrpc:"account_enabled_tax_country_ids,omptempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omptempty"`
	AccountFolder                               *Many2One  `xmlrpc:"account_folder,omptempty"`
	AccountInvoiceOnboardingState               *Selection `xmlrpc:"account_invoice_onboarding_state,omptempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omptempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omptempty"`
	AccountJournalPaymentCreditAccountId        *Many2One  `xmlrpc:"account_journal_payment_credit_account_id,omptempty"`
	AccountJournalPaymentDebitAccountId         *Many2One  `xmlrpc:"account_journal_payment_debit_account_id,omptempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omptempty"`
	AccountOnboardingCreateInvoiceState         *Selection `xmlrpc:"account_onboarding_create_invoice_state,omptempty"`
	AccountOnboardingCreateInvoiceStateFlag     *Bool      `xmlrpc:"account_onboarding_create_invoice_state_flag,omptempty"`
	AccountOnboardingInvoiceLayoutState         *Selection `xmlrpc:"account_onboarding_invoice_layout_state,omptempty"`
	AccountOnboardingSaleTaxState               *Selection `xmlrpc:"account_onboarding_sale_tax_state,omptempty"`
	AccountOpeningDate                          *Time      `xmlrpc:"account_opening_date,omptempty"`
	AccountOpeningJournalId                     *Many2One  `xmlrpc:"account_opening_journal_id,omptempty"`
	AccountOpeningMoveId                        *Many2One  `xmlrpc:"account_opening_move_id,omptempty"`
	AccountPurchaseTaxId                        *Many2One  `xmlrpc:"account_purchase_tax_id,omptempty"`
	AccountRepresentativeId                     *Many2One  `xmlrpc:"account_representative_id,omptempty"`
	AccountRevaluationExpenseProvisionAccountId *Many2One  `xmlrpc:"account_revaluation_expense_provision_account_id,omptempty"`
	AccountRevaluationIncomeProvisionAccountId  *Many2One  `xmlrpc:"account_revaluation_income_provision_account_id,omptempty"`
	AccountRevaluationJournalId                 *Many2One  `xmlrpc:"account_revaluation_journal_id,omptempty"`
	AccountSaleTaxId                            *Many2One  `xmlrpc:"account_sale_tax_id,omptempty"`
	AccountSetupBankDataState                   *Selection `xmlrpc:"account_setup_bank_data_state,omptempty"`
	AccountSetupBillState                       *Selection `xmlrpc:"account_setup_bill_state,omptempty"`
	AccountSetupCoaState                        *Selection `xmlrpc:"account_setup_coa_state,omptempty"`
	AccountSetupFyDataState                     *Selection `xmlrpc:"account_setup_fy_data_state,omptempty"`
	AccountSetupTaxesState                      *Selection `xmlrpc:"account_setup_taxes_state,omptempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omptempty"`
	AccountTaxPeriodicity                       *Selection `xmlrpc:"account_tax_periodicity,omptempty"`
	AccountTaxPeriodicityJournalId              *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omptempty"`
	AccountTaxPeriodicityReminderDay            *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omptempty"`
	AccountTaxUnitIds                           *Relation  `xmlrpc:"account_tax_unit_ids,omptempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omptempty"`
	Active                                      *Bool      `xmlrpc:"active,omptempty"`
	ActiveGlobalAnalyticAccount                 *Bool      `xmlrpc:"active_global_analytic_account,omptempty"`
	AnalyticPlanId                              *Many2One  `xmlrpc:"analytic_plan_id,omptempty"`
	AngloSaxonAccounting                        *Bool      `xmlrpc:"anglo_saxon_accounting,omptempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omptempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omptempty"`
	AppraisalConfirmMailTemplate                *Many2One  `xmlrpc:"appraisal_confirm_mail_template,omptempty"`
	AppraisalEmployeeFeedbackTemplate           *String    `xmlrpc:"appraisal_employee_feedback_template,omptempty"`
	AppraisalManagerFeedbackTemplate            *String    `xmlrpc:"appraisal_manager_feedback_template,omptempty"`
	AppraisalPlan                               *Bool      `xmlrpc:"appraisal_plan,omptempty"`
	AppraisalSurveyTemplateId                   *Many2One  `xmlrpc:"appraisal_survey_template_id,omptempty"`
	AssessmentNoteIds                           *Relation  `xmlrpc:"assessment_note_ids,omptempty"`
	AttendanceBarcodeSource                     *Selection `xmlrpc:"attendance_barcode_source,omptempty"`
	AttendanceKioskDelay                        *Int       `xmlrpc:"attendance_kiosk_delay,omptempty"`
	AttendanceKioskMode                         *Selection `xmlrpc:"attendance_kiosk_mode,omptempty"`
	AutoValidation                              *Bool      `xmlrpc:"auto_validation,omptempty"`
	AutomaticEntryDefaultJournalId              *Many2One  `xmlrpc:"automatic_entry_default_journal_id,omptempty"`
	BackgroundImage                             *String    `xmlrpc:"background_image,omptempty"`
	BankAccountCodePrefix                       *String    `xmlrpc:"bank_account_code_prefix,omptempty"`
	BankIds                                     *Relation  `xmlrpc:"bank_ids,omptempty"`
	BankJournalIds                              *Relation  `xmlrpc:"bank_journal_ids,omptempty"`
	BaseLayout                                  *Selection `xmlrpc:"base_layout,omptempty"`
	BaseOnboardingCompanyState                  *Selection `xmlrpc:"base_onboarding_company_state,omptempty"`
	CashAccountCodePrefix                       *String    `xmlrpc:"cash_account_code_prefix,omptempty"`
	CatchallEmail                               *String    `xmlrpc:"catchall_email,omptempty"`
	CatchallFormatted                           *String    `xmlrpc:"catchall_formatted,omptempty"`
	ChartTemplateId                             *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	ChildIds                                    *Relation  `xmlrpc:"child_ids,omptempty"`
	City                                        *String    `xmlrpc:"city,omptempty"`
	CompanyDetails                              *String    `xmlrpc:"company_details,omptempty"`
	CompanyExpenseJournalId                     *Many2One  `xmlrpc:"company_expense_journal_id,omptempty"`
	CompanyRegistry                             *String    `xmlrpc:"company_registry,omptempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omptempty"`
	CountryId                                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrencyIntervalUnit                        *Selection `xmlrpc:"currency_interval_unit,omptempty"`
	CurrencyNextExecutionDate                   *Time      `xmlrpc:"currency_next_execution_date,omptempty"`
	CurrencyProvider                            *Selection `xmlrpc:"currency_provider,omptempty"`
	DaysToPurchase                              *Float     `xmlrpc:"days_to_purchase,omptempty"`
	DefaultCashDifferenceExpenseAccountId       *Many2One  `xmlrpc:"default_cash_difference_expense_account_id,omptempty"`
	DefaultCashDifferenceIncomeAccountId        *Many2One  `xmlrpc:"default_cash_difference_income_account_id,omptempty"`
	DeferredTimeOffManager                      *Many2One  `xmlrpc:"deferred_time_off_manager,omptempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omptempty"`
	DocumentLayoutId                            *Many2One  `xmlrpc:"document_layout_id,omptempty"`
	DocumentsAccountSettings                    *Bool      `xmlrpc:"documents_account_settings,omptempty"`
	DocumentsHrContractsTags                    *Relation  `xmlrpc:"documents_hr_contracts_tags,omptempty"`
	DocumentsHrFolder                           *Many2One  `xmlrpc:"documents_hr_folder,omptempty"`
	DocumentsHrPayslipsTags                     *Relation  `xmlrpc:"documents_hr_payslips_tags,omptempty"`
	DocumentsHrSettings                         *Bool      `xmlrpc:"documents_hr_settings,omptempty"`
	DocumentsPayrollFolderId                    *Many2One  `xmlrpc:"documents_payroll_folder_id,omptempty"`
	DocumentsProductSettings                    *Bool      `xmlrpc:"documents_product_settings,omptempty"`
	DocumentsRecruitmentSettings                *Bool      `xmlrpc:"documents_recruitment_settings,omptempty"`
	DocumentsSpreadsheetFolderId                *Many2One  `xmlrpc:"documents_spreadsheet_folder_id,omptempty"`
	DurationAfterRecruitment                    *Int       `xmlrpc:"duration_after_recruitment,omptempty"`
	DurationFirstAppraisal                      *Int       `xmlrpc:"duration_first_appraisal,omptempty"`
	DurationNextAppraisal                       *Int       `xmlrpc:"duration_next_appraisal,omptempty"`
	EarlyPayDiscountComputation                 *Selection `xmlrpc:"early_pay_discount_computation,omptempty"`
	Email                                       *String    `xmlrpc:"email,omptempty"`
	EmailFormatted                              *String    `xmlrpc:"email_formatted,omptempty"`
	EnableMenuSearch                            *Bool      `xmlrpc:"enable_menu_search,omptempty"`
	ExpectsChartOfAccounts                      *Bool      `xmlrpc:"expects_chart_of_accounts,omptempty"`
	ExpenseAccrualAccountId                     *Many2One  `xmlrpc:"expense_accrual_account_id,omptempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	ExpenseExtractShowOcrOptionSelection        *Selection `xmlrpc:"expense_extract_show_ocr_option_selection,omptempty"`
	ExpenseJournalId                            *Many2One  `xmlrpc:"expense_journal_id,omptempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	ExtractInInvoiceDigitalizationMode          *Selection `xmlrpc:"extract_in_invoice_digitalization_mode,omptempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection `xmlrpc:"extract_out_invoice_digitalization_mode,omptempty"`
	ExtractSingleLinePerTax                     *Bool      `xmlrpc:"extract_single_line_per_tax,omptempty"`
	FailedMessageIds                            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	Favicon                                     *String    `xmlrpc:"favicon,omptempty"`
	FiscalPositionIds                           *Relation  `xmlrpc:"fiscal_position_ids,omptempty"`
	FiscalyearLastDay                           *Int       `xmlrpc:"fiscalyear_last_day,omptempty"`
	FiscalyearLastMonth                         *Selection `xmlrpc:"fiscalyear_last_month,omptempty"`
	FiscalyearLockDate                          *Time      `xmlrpc:"fiscalyear_lock_date,omptempty"`
	Font                                        *Selection `xmlrpc:"font,omptempty"`
	GainAccountId                               *Many2One  `xmlrpc:"gain_account_id,omptempty"`
	GlobalAnalyticAccountId                     *Many2One  `xmlrpc:"global_analytic_account_id,omptempty"`
	HasMessage                                  *Bool      `xmlrpc:"has_message,omptempty"`
	HasReceivedWarningStockSms                  *Bool      `xmlrpc:"has_received_warning_stock_sms,omptempty"`
	HrAttendanceOvertime                        *Bool      `xmlrpc:"hr_attendance_overtime,omptempty"`
	HrPresenceControlEmailAmount                *Int       `xmlrpc:"hr_presence_control_email_amount,omptempty"`
	HrPresenceControlIpList                     *String    `xmlrpc:"hr_presence_control_ip_list,omptempty"`
	IapEnrichAutoDone                           *Bool      `xmlrpc:"iap_enrich_auto_done,omptempty"`
	Id                                          *Int       `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	IntercompanyTransactionMessage              *String    `xmlrpc:"intercompany_transaction_message,omptempty"`
	IntercompanyUserId                          *Many2One  `xmlrpc:"intercompany_user_id,omptempty"`
	InternalProjectId                           *Many2One  `xmlrpc:"internal_project_id,omptempty"`
	InternalTransitLocationId                   *Many2One  `xmlrpc:"internal_transit_location_id,omptempty"`
	InvoiceIsEmail                              *Bool      `xmlrpc:"invoice_is_email,omptempty"`
	InvoiceIsPrint                              *Bool      `xmlrpc:"invoice_is_print,omptempty"`
	InvoiceIsSnailmail                          *Bool      `xmlrpc:"invoice_is_snailmail,omptempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omptempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omptempty"`
	InvoicingSwitchThreshold                    *Time      `xmlrpc:"invoicing_switch_threshold,omptempty"`
	IsCompanyDetailsEmpty                       *Bool      `xmlrpc:"is_company_details_empty,omptempty"`
	L10NChIsrPreprintedAccount                  *Bool      `xmlrpc:"l10n_ch_isr_preprinted_account,omptempty"`
	L10NChIsrPreprintedBank                     *Bool      `xmlrpc:"l10n_ch_isr_preprinted_bank,omptempty"`
	L10NChIsrPrintBankLocation                  *Bool      `xmlrpc:"l10n_ch_isr_print_bank_location,omptempty"`
	L10NChIsrScanLineLeft                       *Float     `xmlrpc:"l10n_ch_isr_scan_line_left,omptempty"`
	L10NChIsrScanLineTop                        *Float     `xmlrpc:"l10n_ch_isr_scan_line_top,omptempty"`
	LayoutBackground                            *Selection `xmlrpc:"layout_background,omptempty"`
	LayoutBackgroundImage                       *String    `xmlrpc:"layout_background_image,omptempty"`
	LeadDelete                                  *Bool      `xmlrpc:"lead_delete,omptempty"`
	LeaveTimesheetTaskId                        *Many2One  `xmlrpc:"leave_timesheet_task_id,omptempty"`
	Logo                                        *String    `xmlrpc:"logo,omptempty"`
	LogoWeb                                     *String    `xmlrpc:"logo_web,omptempty"`
	LossAccountId                               *Many2One  `xmlrpc:"loss_account_id,omptempty"`
	MessageAttachmentCount                      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                              *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId                     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter                    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Mobile                                      *String    `xmlrpc:"mobile,omptempty"`
	MultiVatForeignCountryIds                   *Relation  `xmlrpc:"multi_vat_foreign_country_ids,omptempty"`
	Name                                        *String    `xmlrpc:"name,omptempty"`
	NomenclatureId                              *Many2One  `xmlrpc:"nomenclature_id,omptempty"`
	OppDelete                                   *Bool      `xmlrpc:"opp_delete,omptempty"`
	OvertimeCompanyThreshold                    *Int       `xmlrpc:"overtime_company_threshold,omptempty"`
	OvertimeEmployeeThreshold                   *Int       `xmlrpc:"overtime_employee_threshold,omptempty"`
	OvertimeStartDate                           *Time      `xmlrpc:"overtime_start_date,omptempty"`
	PaperformatId                               *Many2One  `xmlrpc:"paperformat_id,omptempty"`
	ParentId                                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerGid                                  *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerId                                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentOnboardingPaymentMethod              *Selection `xmlrpc:"payment_onboarding_payment_method,omptempty"`
	PaymentProviderOnboardingState              *Selection `xmlrpc:"payment_provider_onboarding_state,omptempty"`
	PeriodLockDate                              *Time      `xmlrpc:"period_lock_date,omptempty"`
	Phone                                       *String    `xmlrpc:"phone,omptempty"`
	PlanningAllowSelfUnassign                   *Bool      `xmlrpc:"planning_allow_self_unassign,omptempty"`
	PlanningGenerationInterval                  *Int       `xmlrpc:"planning_generation_interval,omptempty"`
	PlanningSelfUnassignDaysBefore              *Int       `xmlrpc:"planning_self_unassign_days_before,omptempty"`
	PoDoubleValidation                          *Selection `xmlrpc:"po_double_validation,omptempty"`
	PoDoubleValidationAmount                    *Float     `xmlrpc:"po_double_validation_amount,omptempty"`
	PoLead                                      *Float     `xmlrpc:"po_lead,omptempty"`
	PoLock                                      *Selection `xmlrpc:"po_lock,omptempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omptempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omptempty"`
	PreventOldTimesheetsEncoding                *Bool      `xmlrpc:"prevent_old_timesheets_encoding,omptempty"`
	PrimaryColor                                *String    `xmlrpc:"primary_color,omptempty"`
	ProductFolder                               *Many2One  `xmlrpc:"product_folder,omptempty"`
	ProductTags                                 *Relation  `xmlrpc:"product_tags,omptempty"`
	ProjectTimeModeId                           *Many2One  `xmlrpc:"project_time_mode_id,omptempty"`
	PropertyStockAccountInputCategId            *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omptempty"`
	PropertyStockAccountOutputCategId           *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omptempty"`
	PropertyStockValuationAccountId             *Many2One  `xmlrpc:"property_stock_valuation_account_id,omptempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omptempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omptempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omptempty"`
	RecruitmentExtractShowOcrOptionSelection    *Selection `xmlrpc:"recruitment_extract_show_ocr_option_selection,omptempty"`
	RecruitmentFolderId                         *Many2One  `xmlrpc:"recruitment_folder_id,omptempty"`
	RecruitmentTagIds                           *Relation  `xmlrpc:"recruitment_tag_ids,omptempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omptempty"`
	ReportHeader                                *String    `xmlrpc:"report_header,omptempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceCalendarIds                         *Relation  `xmlrpc:"resource_calendar_ids,omptempty"`
	RevenueAccrualAccountId                     *Many2One  `xmlrpc:"revenue_accrual_account_id,omptempty"`
	RuleType                                    *Selection `xmlrpc:"rule_type,omptempty"`
	SaleAutoUpdatePrice                         *Bool      `xmlrpc:"sale_auto_update_price,omptempty"`
	SaleOnboardingOrderConfirmationState        *Selection `xmlrpc:"sale_onboarding_order_confirmation_state,omptempty"`
	SaleOnboardingPaymentMethod                 *Selection `xmlrpc:"sale_onboarding_payment_method,omptempty"`
	SaleOnboardingSampleQuotationState          *Selection `xmlrpc:"sale_onboarding_sample_quotation_state,omptempty"`
	SaleOrderTemplateId                         *Many2One  `xmlrpc:"sale_order_template_id,omptempty"`
	SaleQuotationOnboardingState                *Selection `xmlrpc:"sale_quotation_onboarding_state,omptempty"`
	SddCreditorIdentifier                       *String    `xmlrpc:"sdd_creditor_identifier,omptempty"`
	SecondaryColor                              *String    `xmlrpc:"secondary_color,omptempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omptempty"`
	SepaInitiatingPartyName                     *String    `xmlrpc:"sepa_initiating_party_name,omptempty"`
	SepaOrgidId                                 *String    `xmlrpc:"sepa_orgid_id,omptempty"`
	SepaOrgidIssr                               *String    `xmlrpc:"sepa_orgid_issr,omptempty"`
	Sequence                                    *Int       `xmlrpc:"sequence,omptempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omptempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omptempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omptempty"`
	SocialFacebook                              *String    `xmlrpc:"social_facebook,omptempty"`
	SocialGithub                                *String    `xmlrpc:"social_github,omptempty"`
	SocialInstagram                             *String    `xmlrpc:"social_instagram,omptempty"`
	SocialLinkedin                              *String    `xmlrpc:"social_linkedin,omptempty"`
	SocialTwitter                               *String    `xmlrpc:"social_twitter,omptempty"`
	SocialYoutube                               *String    `xmlrpc:"social_youtube,omptempty"`
	StartSearchAfterLetter                      *Int       `xmlrpc:"start_search_after_letter,omptempty"`
	StateId                                     *Many2One  `xmlrpc:"state_id,omptempty"`
	StockMailConfirmationTemplateId             *Many2One  `xmlrpc:"stock_mail_confirmation_template_id,omptempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omptempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omptempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omptempty"`
	Street                                      *String    `xmlrpc:"street,omptempty"`
	Street2                                     *String    `xmlrpc:"street2,omptempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TaxLockDate                                 *Time      `xmlrpc:"tax_lock_date,omptempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omptempty"`
	TimesheetEncodeUomId                        *Many2One  `xmlrpc:"timesheet_encode_uom_id,omptempty"`
	TimesheetMailEmployeeAllow                  *Bool      `xmlrpc:"timesheet_mail_employee_allow,omptempty"`
	TimesheetMailEmployeeDelay                  *Int       `xmlrpc:"timesheet_mail_employee_delay,omptempty"`
	TimesheetMailEmployeeInterval               *Selection `xmlrpc:"timesheet_mail_employee_interval,omptempty"`
	TimesheetMailEmployeeNextdate               *Time      `xmlrpc:"timesheet_mail_employee_nextdate,omptempty"`
	TimesheetMailManagerAllow                   *Bool      `xmlrpc:"timesheet_mail_manager_allow,omptempty"`
	TimesheetMailManagerDelay                   *Int       `xmlrpc:"timesheet_mail_manager_delay,omptempty"`
	TimesheetMailManagerInterval                *Selection `xmlrpc:"timesheet_mail_manager_interval,omptempty"`
	TimesheetMailManagerNextdate                *Time      `xmlrpc:"timesheet_mail_manager_nextdate,omptempty"`
	TotalsBelowSections                         *Bool      `xmlrpc:"totals_below_sections,omptempty"`
	TransferAccountCodePrefix                   *String    `xmlrpc:"transfer_account_code_prefix,omptempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omptempty"`
	UserIds                                     *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                                         *String    `xmlrpc:"vat,omptempty"`
	WarehouseId                                 *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	Website                                     *String    `xmlrpc:"website,omptempty"`
	WebsiteId                                   *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds                           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteSaleDashboardOnboardingState         *Selection `xmlrpc:"website_sale_dashboard_onboarding_state,omptempty"`
	WebsiteSaleOnboardingPaymentProviderState   *Selection `xmlrpc:"website_sale_onboarding_payment_provider_state,omptempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omptempty"`
	XStudioMany2ManyCompanyWithholdingTax       *Relation  `xmlrpc:"x_studio_many2many_company_withholding_tax,omptempty"`
	XStudioMany2ManyFieldP56A4                  *Relation  `xmlrpc:"x_studio_many2many_field_p56A4,omptempty"`
	XStudioMany2OneCompanyQst                   *Many2One  `xmlrpc:"x_studio_many2one_company_qst,omptempty"`
	XStudioOne2ManyWithholdingTaxList           *Relation  `xmlrpc:"x_studio_one2many_withholding_tax_list,omptempty"`
	Zip                                         *String    `xmlrpc:"zip,omptempty"`
}

// ResCompanys represents array of res.company model.
type ResCompanys []ResCompany

// ResCompanyModel is the odoo model name.
const ResCompanyModel = "res.company"

// Many2One convert ResCompany to *Many2One.
func (rc *ResCompany) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompany(rc *ResCompany) (int64, error) {
	ids, err := c.CreateResCompanys([]*ResCompany{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResCompany creates a new res.company model and returns its id.
func (c *Client) CreateResCompanys(rcs []*ResCompany) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResCompanyModel, vv)
}

// UpdateResCompany updates an existing res.company record.
func (c *Client) UpdateResCompany(rc *ResCompany) error {
	return c.UpdateResCompanys([]int64{rc.Id.Get()}, rc)
}

// UpdateResCompanys updates existing res.company records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResCompanys(ids []int64, rc *ResCompany) error {
	return c.Update(ResCompanyModel, ids, rc)
}

// DeleteResCompany deletes an existing res.company record.
func (c *Client) DeleteResCompany(id int64) error {
	return c.DeleteResCompanys([]int64{id})
}

// DeleteResCompanys deletes existing res.company records.
func (c *Client) DeleteResCompanys(ids []int64) error {
	return c.Delete(ResCompanyModel, ids)
}

// GetResCompany gets res.company existing record.
func (c *Client) GetResCompany(id int64) (*ResCompany, error) {
	rcs, err := c.GetResCompanys([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.company not found", id)
}

// GetResCompanys gets res.company existing records.
func (c *Client) GetResCompanys(ids []int64) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.Read(ResCompanyModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompany finds res.company record by querying it with criteria.
func (c *Client) FindResCompany(criteria *Criteria) (*ResCompany, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	if rcs != nil && len(*rcs) > 0 {
		return &((*rcs)[0]), nil
	}
	return nil, fmt.Errorf("res.company was not found with criteria %v", criteria)
}

// FindResCompanys finds res.company records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanys(criteria *Criteria, options *Options) (*ResCompanys, error) {
	rcs := &ResCompanys{}
	if err := c.SearchRead(ResCompanyModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResCompanyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResCompanyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResCompanyId finds record id by querying it with criteria.
func (c *Client) FindResCompanyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResCompanyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.company was not found with criteria %v and options %v", criteria, options)
}
