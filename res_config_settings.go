package odoo

import (
	"fmt"
)

// ResConfigSettings represents res.config.settings model.
type ResConfigSettings struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omptempty"`
	AccountCashBasisBaseAccountId               *Many2One  `xmlrpc:"account_cash_basis_base_account_id,omptempty"`
	AccountDefaultCreditLimit                   *Float     `xmlrpc:"account_default_credit_limit,omptempty"`
	AccountFiscalCountryId                      *Many2One  `xmlrpc:"account_fiscal_country_id,omptempty"`
	AccountFolder                               *Many2One  `xmlrpc:"account_folder,omptempty"`
	AccountJournalEarlyPayDiscountGainAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_gain_account_id,omptempty"`
	AccountJournalEarlyPayDiscountLossAccountId *Many2One  `xmlrpc:"account_journal_early_pay_discount_loss_account_id,omptempty"`
	AccountJournalPaymentCreditAccountId        *Many2One  `xmlrpc:"account_journal_payment_credit_account_id,omptempty"`
	AccountJournalPaymentDebitAccountId         *Many2One  `xmlrpc:"account_journal_payment_debit_account_id,omptempty"`
	AccountJournalSuspenseAccountId             *Many2One  `xmlrpc:"account_journal_suspense_account_id,omptempty"`
	AccountOnCheckout                           *Selection `xmlrpc:"account_on_checkout,omptempty"`
	AccountStorno                               *Bool      `xmlrpc:"account_storno,omptempty"`
	AccountTaxPeriodicity                       *Selection `xmlrpc:"account_tax_periodicity,omptempty"`
	AccountTaxPeriodicityJournalId              *Many2One  `xmlrpc:"account_tax_periodicity_journal_id,omptempty"`
	AccountTaxPeriodicityReminderDay            *Int       `xmlrpc:"account_tax_periodicity_reminder_day,omptempty"`
	AccountUseCreditLimit                       *Bool      `xmlrpc:"account_use_credit_limit,omptempty"`
	ActiveGlobalAnalyticAccount                 *Bool      `xmlrpc:"active_global_analytic_account,omptempty"`
	ActiveUserCount                             *Int       `xmlrpc:"active_user_count,omptempty"`
	AddToCartAction                             *Selection `xmlrpc:"add_to_cart_action,omptempty"`
	AliasDomain                                 *String    `xmlrpc:"alias_domain,omptempty"`
	AllowOutOfStockOrder                        *Bool      `xmlrpc:"allow_out_of_stock_order,omptempty"`
	AnalyticPlanId                              *Many2One  `xmlrpc:"analytic_plan_id,omptempty"`
	AnnualInventoryDay                          *Int       `xmlrpc:"annual_inventory_day,omptempty"`
	AnnualInventoryMonth                        *Selection `xmlrpc:"annual_inventory_month,omptempty"`
	AppraisalEmployeeFeedbackTemplate           *String    `xmlrpc:"appraisal_employee_feedback_template,omptempty"`
	AppraisalManagerFeedbackTemplate            *String    `xmlrpc:"appraisal_manager_feedback_template,omptempty"`
	AppraisalPlan                               *Bool      `xmlrpc:"appraisal_plan,omptempty"`
	AppraisalSurveyTemplateId                   *Many2One  `xmlrpc:"appraisal_survey_template_id,omptempty"`
	AssessmentNoteIds                           *Relation  `xmlrpc:"assessment_note_ids,omptempty"`
	AttendanceBarcodeSource                     *Selection `xmlrpc:"attendance_barcode_source,omptempty"`
	AttendanceKioskDelay                        *Int       `xmlrpc:"attendance_kiosk_delay,omptempty"`
	AttendanceKioskMode                         *Selection `xmlrpc:"attendance_kiosk_mode,omptempty"`
	AuthOauthGoogleClientId                     *String    `xmlrpc:"auth_oauth_google_client_id,omptempty"`
	AuthOauthGoogleEnabled                      *Bool      `xmlrpc:"auth_oauth_google_enabled,omptempty"`
	AuthSignupResetPassword                     *Bool      `xmlrpc:"auth_signup_reset_password,omptempty"`
	AuthSignupTemplateUserId                    *Many2One  `xmlrpc:"auth_signup_template_user_id,omptempty"`
	AuthSignupUninvited                         *Selection `xmlrpc:"auth_signup_uninvited,omptempty"`
	AutoValidation                              *Bool      `xmlrpc:"auto_validation,omptempty"`
	AutomaticInvoice                            *Bool      `xmlrpc:"automatic_invoice,omptempty"`
	AvailableThreshold                          *Float     `xmlrpc:"available_threshold,omptempty"`
	CartAbandonedDelay                          *Float     `xmlrpc:"cart_abandoned_delay,omptempty"`
	CartRecoveryMailTemplate                    *Many2One  `xmlrpc:"cart_recovery_mail_template,omptempty"`
	CdnActivated                                *Bool      `xmlrpc:"cdn_activated,omptempty"`
	CdnFilters                                  *String    `xmlrpc:"cdn_filters,omptempty"`
	CdnUrl                                      *String    `xmlrpc:"cdn_url,omptempty"`
	ChannelId                                   *Many2One  `xmlrpc:"channel_id,omptempty"`
	ChartTemplateId                             *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	CompanyCount                                *Int       `xmlrpc:"company_count,omptempty"`
	CompanyCurrencyId                           *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyExpenseJournalId                     *Many2One  `xmlrpc:"company_expense_journal_id,omptempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyInformations                         *String    `xmlrpc:"company_informations,omptempty"`
	CompanyName                                 *String    `xmlrpc:"company_name,omptempty"`
	CompanySoTemplateId                         *Many2One  `xmlrpc:"company_so_template_id,omptempty"`
	CountryCode                                 *String    `xmlrpc:"country_code,omptempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrmAutoAssignmentAction                     *Selection `xmlrpc:"crm_auto_assignment_action,omptempty"`
	CrmAutoAssignmentIntervalNumber             *Int       `xmlrpc:"crm_auto_assignment_interval_number,omptempty"`
	CrmAutoAssignmentIntervalType               *Selection `xmlrpc:"crm_auto_assignment_interval_type,omptempty"`
	CrmAutoAssignmentRunDatetime                *Time      `xmlrpc:"crm_auto_assignment_run_datetime,omptempty"`
	CrmUseAutoAssignment                        *Bool      `xmlrpc:"crm_use_auto_assignment,omptempty"`
	CurrencyExchangeJournalId                   *Many2One  `xmlrpc:"currency_exchange_journal_id,omptempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrencyIntervalUnit                        *Selection `xmlrpc:"currency_interval_unit,omptempty"`
	CurrencyNextExecutionDate                   *Time      `xmlrpc:"currency_next_execution_date,omptempty"`
	CurrencyProvider                            *Selection `xmlrpc:"currency_provider,omptempty"`
	DaysToPurchase                              *Float     `xmlrpc:"days_to_purchase,omptempty"`
	DefaultInvoicePolicy                        *Selection `xmlrpc:"default_invoice_policy,omptempty"`
	DefaultPickingPolicy                        *Selection `xmlrpc:"default_picking_policy,omptempty"`
	DefaultPurchaseMethod                       *Selection `xmlrpc:"default_purchase_method,omptempty"`
	DeferredTimeOffManager                      *Many2One  `xmlrpc:"deferred_time_off_manager,omptempty"`
	DeleteSpamAfter                             *Int       `xmlrpc:"delete_spam_after,omptempty"`
	DepositDefaultProductId                     *Many2One  `xmlrpc:"deposit_default_product_id,omptempty"`
	DigestEmails                                *Bool      `xmlrpc:"digest_emails,omptempty"`
	DigestId                                    *Many2One  `xmlrpc:"digest_id,omptempty"`
	DisableRedirectFirebaseDynamicLink          *Bool      `xmlrpc:"disable_redirect_firebase_dynamic_link,omptempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omptempty"`
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
	EnableMenuSearch                            *Bool      `xmlrpc:"enable_menu_search,omptempty"`
	EnableOcn                                   *Bool      `xmlrpc:"enable_ocn,omptempty"`
	EnabledBuyNowButton                         *Bool      `xmlrpc:"enabled_buy_now_button,omptempty"`
	EnabledExtraCheckoutStep                    *Bool      `xmlrpc:"enabled_extra_checkout_step,omptempty"`
	ExpenseAliasPrefix                          *String    `xmlrpc:"expense_alias_prefix,omptempty"`
	ExpenseCurrencyExchangeAccountId            *Many2One  `xmlrpc:"expense_currency_exchange_account_id,omptempty"`
	ExpenseExtractShowOcrOptionSelection        *Selection `xmlrpc:"expense_extract_show_ocr_option_selection,omptempty"`
	ExpenseJournalId                            *Many2One  `xmlrpc:"expense_journal_id,omptempty"`
	ExternalEmailServerDefault                  *Bool      `xmlrpc:"external_email_server_default,omptempty"`
	ExternalReportLayoutId                      *Many2One  `xmlrpc:"external_report_layout_id,omptempty"`
	ExtractInInvoiceDigitalizationMode          *Selection `xmlrpc:"extract_in_invoice_digitalization_mode,omptempty"`
	ExtractOutInvoiceDigitalizationMode         *Selection `xmlrpc:"extract_out_invoice_digitalization_mode,omptempty"`
	ExtractSingleLinePerTax                     *Bool      `xmlrpc:"extract_single_line_per_tax,omptempty"`
	FacebookAppId                               *String    `xmlrpc:"facebook_app_id,omptempty"`
	FacebookClientSecret                        *String    `xmlrpc:"facebook_client_secret,omptempty"`
	FacebookUseOwnAccount                       *Bool      `xmlrpc:"facebook_use_own_account,omptempty"`
	FailCounter                                 *Int       `xmlrpc:"fail_counter,omptempty"`
	FallbackModelId                             *Many2One  `xmlrpc:"fallback_model_id,omptempty"`
	Favicon                                     *String    `xmlrpc:"favicon,omptempty"`
	FirebaseAdminKeyFile                        *String    `xmlrpc:"firebase_admin_key_file,omptempty"`
	FirebaseEnablePushNotifications             *Bool      `xmlrpc:"firebase_enable_push_notifications,omptempty"`
	FirebaseProjectId                           *String    `xmlrpc:"firebase_project_id,omptempty"`
	FirebasePushCertificateKey                  *String    `xmlrpc:"firebase_push_certificate_key,omptempty"`
	FirebaseSenderId                            *String    `xmlrpc:"firebase_sender_id,omptempty"`
	FirebaseUseOwnAccount                       *Bool      `xmlrpc:"firebase_use_own_account,omptempty"`
	FirebaseWebApiKey                           *String    `xmlrpc:"firebase_web_api_key,omptempty"`
	FirstProviderLabel                          *String    `xmlrpc:"first_provider_label,omptempty"`
	FiscalyearLastDay                           *Int       `xmlrpc:"fiscalyear_last_day,omptempty"`
	FiscalyearLastMonth                         *Selection `xmlrpc:"fiscalyear_last_month,omptempty"`
	FiscalyearLockDate                          *Time      `xmlrpc:"fiscalyear_lock_date,omptempty"`
	GlobalAnalyticAccountId                     *Many2One  `xmlrpc:"global_analytic_account_id,omptempty"`
	GoogleAnalyticsKey                          *String    `xmlrpc:"google_analytics_key,omptempty"`
	GoogleSearchConsole                         *String    `xmlrpc:"google_search_console,omptempty"`
	GroupAnalyticAccounting                     *Bool      `xmlrpc:"group_analytic_accounting,omptempty"`
	GroupApplicantCvDisplay                     *Bool      `xmlrpc:"group_applicant_cv_display,omptempty"`
	GroupAttendanceUsePin                       *Bool      `xmlrpc:"group_attendance_use_pin,omptempty"`
	GroupAutoDoneSetting                        *Bool      `xmlrpc:"group_auto_done_setting,omptempty"`
	GroupCashRounding                           *Bool      `xmlrpc:"group_cash_rounding,omptempty"`
	GroupDeliveryInvoiceAddress                 *Bool      `xmlrpc:"group_delivery_invoice_address,omptempty"`
	GroupDiscountPerSoLine                      *Bool      `xmlrpc:"group_discount_per_so_line,omptempty"`
	GroupDisplayIncoterm                        *Bool      `xmlrpc:"group_display_incoterm,omptempty"`
	GroupFiscalYear                             *Bool      `xmlrpc:"group_fiscal_year,omptempty"`
	GroupLotOnDeliverySlip                      *Bool      `xmlrpc:"group_lot_on_delivery_slip,omptempty"`
	GroupLotOnInvoice                           *Bool      `xmlrpc:"group_lot_on_invoice,omptempty"`
	GroupMassMailingCampaign                    *Bool      `xmlrpc:"group_mass_mailing_campaign,omptempty"`
	GroupMultiCurrency                          *Bool      `xmlrpc:"group_multi_currency,omptempty"`
	GroupMultiWebsite                           *Bool      `xmlrpc:"group_multi_website,omptempty"`
	GroupProductPriceComparison                 *Bool      `xmlrpc:"group_product_price_comparison,omptempty"`
	GroupProductPricelist                       *Bool      `xmlrpc:"group_product_pricelist,omptempty"`
	GroupProductVariant                         *Bool      `xmlrpc:"group_product_variant,omptempty"`
	GroupProformaSales                          *Bool      `xmlrpc:"group_proforma_sales,omptempty"`
	GroupProjectMilestone                       *Bool      `xmlrpc:"group_project_milestone,omptempty"`
	GroupProjectRating                          *Bool      `xmlrpc:"group_project_rating,omptempty"`
	GroupProjectRecurringTasks                  *Bool      `xmlrpc:"group_project_recurring_tasks,omptempty"`
	GroupProjectStages                          *Bool      `xmlrpc:"group_project_stages,omptempty"`
	GroupProjectTaskDependencies                *Bool      `xmlrpc:"group_project_task_dependencies,omptempty"`
	GroupSaleDeliveryAddress                    *Bool      `xmlrpc:"group_sale_delivery_address,omptempty"`
	GroupSaleOrderTemplate                      *Bool      `xmlrpc:"group_sale_order_template,omptempty"`
	GroupSalePricelist                          *Bool      `xmlrpc:"group_sale_pricelist,omptempty"`
	GroupSendReminder                           *Bool      `xmlrpc:"group_send_reminder,omptempty"`
	GroupShowLineSubtotalsTaxExcluded           *Bool      `xmlrpc:"group_show_line_subtotals_tax_excluded,omptempty"`
	GroupShowLineSubtotalsTaxIncluded           *Bool      `xmlrpc:"group_show_line_subtotals_tax_included,omptempty"`
	GroupShowPurchaseReceipts                   *Bool      `xmlrpc:"group_show_purchase_receipts,omptempty"`
	GroupShowSaleReceipts                       *Bool      `xmlrpc:"group_show_sale_receipts,omptempty"`
	GroupShowUomPrice                           *Bool      `xmlrpc:"group_show_uom_price,omptempty"`
	GroupStockAdvLocation                       *Bool      `xmlrpc:"group_stock_adv_location,omptempty"`
	GroupStockLotPrintGs1                       *Bool      `xmlrpc:"group_stock_lot_print_gs1,omptempty"`
	GroupStockMultiLocations                    *Bool      `xmlrpc:"group_stock_multi_locations,omptempty"`
	GroupStockPackaging                         *Bool      `xmlrpc:"group_stock_packaging,omptempty"`
	GroupStockPickingWave                       *Bool      `xmlrpc:"group_stock_picking_wave,omptempty"`
	GroupStockProductionLot                     *Bool      `xmlrpc:"group_stock_production_lot,omptempty"`
	GroupStockReceptionReport                   *Bool      `xmlrpc:"group_stock_reception_report,omptempty"`
	GroupStockSignDelivery                      *Bool      `xmlrpc:"group_stock_sign_delivery,omptempty"`
	GroupStockStorageCategories                 *Bool      `xmlrpc:"group_stock_storage_categories,omptempty"`
	GroupStockTrackingLot                       *Bool      `xmlrpc:"group_stock_tracking_lot,omptempty"`
	GroupStockTrackingOwner                     *Bool      `xmlrpc:"group_stock_tracking_owner,omptempty"`
	GroupSubtaskProject                         *Bool      `xmlrpc:"group_subtask_project,omptempty"`
	GroupUom                                    *Bool      `xmlrpc:"group_uom,omptempty"`
	GroupUseLead                                *Bool      `xmlrpc:"group_use_lead,omptempty"`
	GroupUseRecurringRevenues                   *Bool      `xmlrpc:"group_use_recurring_revenues,omptempty"`
	GroupWarningAccount                         *Bool      `xmlrpc:"group_warning_account,omptempty"`
	GroupWarningPurchase                        *Bool      `xmlrpc:"group_warning_purchase,omptempty"`
	GroupWarningSale                            *Bool      `xmlrpc:"group_warning_sale,omptempty"`
	GroupWarningStock                           *Bool      `xmlrpc:"group_warning_stock,omptempty"`
	HasAccountingEntries                        *Bool      `xmlrpc:"has_accounting_entries,omptempty"`
	HasChartOfAccounts                          *Bool      `xmlrpc:"has_chart_of_accounts,omptempty"`
	HasDefaultShareImage                        *Bool      `xmlrpc:"has_default_share_image,omptempty"`
	HasGoogleAnalytics                          *Bool      `xmlrpc:"has_google_analytics,omptempty"`
	HasGoogleSearchConsole                      *Bool      `xmlrpc:"has_google_search_console,omptempty"`
	HasPlausibleSharedKey                       *Bool      `xmlrpc:"has_plausible_shared_key,omptempty"`
	HrAttendanceOvertime                        *Bool      `xmlrpc:"hr_attendance_overtime,omptempty"`
	HrEmployeeSelfEdit                          *Bool      `xmlrpc:"hr_employee_self_edit,omptempty"`
	HrPresenceControlEmail                      *Bool      `xmlrpc:"hr_presence_control_email,omptempty"`
	HrPresenceControlEmailAmount                *Int       `xmlrpc:"hr_presence_control_email_amount,omptempty"`
	HrPresenceControlIp                         *Bool      `xmlrpc:"hr_presence_control_ip,omptempty"`
	HrPresenceControlIpList                     *String    `xmlrpc:"hr_presence_control_ip_list,omptempty"`
	HrPresenceControlLogin                      *Bool      `xmlrpc:"hr_presence_control_login,omptempty"`
	Id                                          *Int       `xmlrpc:"id,omptempty"`
	IncomeCurrencyExchangeAccountId             *Many2One  `xmlrpc:"income_currency_exchange_account_id,omptempty"`
	IncotermId                                  *Many2One  `xmlrpc:"incoterm_id,omptempty"`
	InstagramAppId                              *String    `xmlrpc:"instagram_app_id,omptempty"`
	InstagramClientSecret                       *String    `xmlrpc:"instagram_client_secret,omptempty"`
	InstagramUseOwnAccount                      *Bool      `xmlrpc:"instagram_use_own_account,omptempty"`
	IntercompanyTransactionMessage              *String    `xmlrpc:"intercompany_transaction_message,omptempty"`
	IntercompanyUserId                          *Many2One  `xmlrpc:"intercompany_user_id,omptempty"`
	InternalProjectId                           *Many2One  `xmlrpc:"internal_project_id,omptempty"`
	InvoiceIsEmail                              *Bool      `xmlrpc:"invoice_is_email,omptempty"`
	InvoiceIsPrint                              *Bool      `xmlrpc:"invoice_is_print,omptempty"`
	InvoiceIsSnailmail                          *Bool      `xmlrpc:"invoice_is_snailmail,omptempty"`
	InvoiceMailTemplateId                       *Many2One  `xmlrpc:"invoice_mail_template_id,omptempty"`
	InvoicePolicy                               *Bool      `xmlrpc:"invoice_policy,omptempty"`
	InvoiceTerms                                *String    `xmlrpc:"invoice_terms,omptempty"`
	InvoiceTermsHtml                            *String    `xmlrpc:"invoice_terms_html,omptempty"`
	InvoicedTimesheet                           *Selection `xmlrpc:"invoiced_timesheet,omptempty"`
	InvoicingSwitchThreshold                    *Time      `xmlrpc:"invoicing_switch_threshold,omptempty"`
	IsEncodeUomDays                             *Bool      `xmlrpc:"is_encode_uom_days,omptempty"`
	IsInstalledSale                             *Bool      `xmlrpc:"is_installed_sale,omptempty"`
	IsMembershipMulti                           *Bool      `xmlrpc:"is_membership_multi,omptempty"`
	IsStripeSupportedCountry                    *Bool      `xmlrpc:"is_stripe_supported_country,omptempty"`
	L10NChIsrPreprintedAccount                  *Bool      `xmlrpc:"l10n_ch_isr_preprinted_account,omptempty"`
	L10NChIsrPreprintedBank                     *Bool      `xmlrpc:"l10n_ch_isr_preprinted_bank,omptempty"`
	L10NChIsrPrintBankLocation                  *Bool      `xmlrpc:"l10n_ch_isr_print_bank_location,omptempty"`
	L10NChIsrScanLineLeft                       *Float     `xmlrpc:"l10n_ch_isr_scan_line_left,omptempty"`
	L10NChIsrScanLineTop                        *Float     `xmlrpc:"l10n_ch_isr_scan_line_top,omptempty"`
	LanguageCount                               *Int       `xmlrpc:"language_count,omptempty"`
	LanguageIds                                 *Relation  `xmlrpc:"language_ids,omptempty"`
	LeadEnrichAuto                              *Selection `xmlrpc:"lead_enrich_auto,omptempty"`
	LeadMiningInPipeline                        *Bool      `xmlrpc:"lead_mining_in_pipeline,omptempty"`
	LeaveTimesheetTaskId                        *Many2One  `xmlrpc:"leave_timesheet_task_id,omptempty"`
	LinkedinAppId                               *String    `xmlrpc:"linkedin_app_id,omptempty"`
	LinkedinClientSecret                        *String    `xmlrpc:"linkedin_client_secret,omptempty"`
	LinkedinUseOwnAccount                       *Bool      `xmlrpc:"linkedin_use_own_account,omptempty"`
	LockConfirmedPo                             *Bool      `xmlrpc:"lock_confirmed_po,omptempty"`
	MailIncomingSmartNotify                     *Bool      `xmlrpc:"mail_incoming_smart_notify,omptempty"`
	MailTrackingMailgunApiKey                   *String    `xmlrpc:"mail_tracking_mailgun_api_key,omptempty"`
	MailTrackingMailgunApiUrl                   *String    `xmlrpc:"mail_tracking_mailgun_api_url,omptempty"`
	MailTrackingMailgunDomain                   *String    `xmlrpc:"mail_tracking_mailgun_domain,omptempty"`
	MailTrackingMailgunEnabled                  *Bool      `xmlrpc:"mail_tracking_mailgun_enabled,omptempty"`
	MailTrackingMailgunValidationKey            *String    `xmlrpc:"mail_tracking_mailgun_validation_key,omptempty"`
	MailTrackingMailgunWebhookSigningKey        *String    `xmlrpc:"mail_tracking_mailgun_webhook_signing_key,omptempty"`
	MailTrackingMailgunWebhooksDomain           *String    `xmlrpc:"mail_tracking_mailgun_webhooks_domain,omptempty"`
	MapBoxToken                                 *String    `xmlrpc:"map_box_token,omptempty"`
	MassMailingMailServerId                     *Many2One  `xmlrpc:"mass_mailing_mail_server_id,omptempty"`
	MassMailingOutgoingMailServer               *Bool      `xmlrpc:"mass_mailing_outgoing_mail_server,omptempty"`
	MassMailingReports                          *Bool      `xmlrpc:"mass_mailing_reports,omptempty"`
	MessageSignatureLocation                    *Selection `xmlrpc:"message_signature_location,omptempty"`
	MessagesEasyColorNote                       *String    `xmlrpc:"messages_easy_color_note,omptempty"`
	MessagesEasyEmptyTrash                      *Int       `xmlrpc:"messages_easy_empty_trash,omptempty"`
	MessagesEasyTextPreview                     *Int       `xmlrpc:"messages_easy_text_preview,omptempty"`
	ModuleAccount                               *Bool      `xmlrpc:"module_account,omptempty"`
	ModuleAccount3WayMatch                      *Bool      `xmlrpc:"module_account_3way_match,omptempty"`
	ModuleAccountAccountant                     *Bool      `xmlrpc:"module_account_accountant,omptempty"`
	ModuleAccountBankStatementImportCamt        *Bool      `xmlrpc:"module_account_bank_statement_import_camt,omptempty"`
	ModuleAccountBankStatementImportCsv         *Bool      `xmlrpc:"module_account_bank_statement_import_csv,omptempty"`
	ModuleAccountBankStatementImportOfx         *Bool      `xmlrpc:"module_account_bank_statement_import_ofx,omptempty"`
	ModuleAccountBankStatementImportQif         *Bool      `xmlrpc:"module_account_bank_statement_import_qif,omptempty"`
	ModuleAccountBatchPayment                   *Bool      `xmlrpc:"module_account_batch_payment,omptempty"`
	ModuleAccountBudget                         *Bool      `xmlrpc:"module_account_budget,omptempty"`
	ModuleAccountCheckPrinting                  *Bool      `xmlrpc:"module_account_check_printing,omptempty"`
	ModuleAccountInterCompanyRules              *Bool      `xmlrpc:"module_account_inter_company_rules,omptempty"`
	ModuleAccountIntrastat                      *Bool      `xmlrpc:"module_account_intrastat,omptempty"`
	ModuleAccountInvoiceExtract                 *Bool      `xmlrpc:"module_account_invoice_extract,omptempty"`
	ModuleAccountPayment                        *Bool      `xmlrpc:"module_account_payment,omptempty"`
	ModuleAccountPredictiveBills                *Bool      `xmlrpc:"module_account_predictive_bills,omptempty"`
	ModuleAccountReports                        *Bool      `xmlrpc:"module_account_reports,omptempty"`
	ModuleAccountSepa                           *Bool      `xmlrpc:"module_account_sepa,omptempty"`
	ModuleAccountSepaDirectDebit                *Bool      `xmlrpc:"module_account_sepa_direct_debit,omptempty"`
	ModuleAccountTaxcloud                       *Bool      `xmlrpc:"module_account_taxcloud,omptempty"`
	ModuleAuthLdap                              *Bool      `xmlrpc:"module_auth_ldap,omptempty"`
	ModuleAuthOauth                             *Bool      `xmlrpc:"module_auth_oauth,omptempty"`
	ModuleBaseGengo                             *Bool      `xmlrpc:"module_base_gengo,omptempty"`
	ModuleBaseGeolocalize                       *Bool      `xmlrpc:"module_base_geolocalize,omptempty"`
	ModuleBaseImport                            *Bool      `xmlrpc:"module_base_import,omptempty"`
	ModuleCrmIapEnrich                          *Bool      `xmlrpc:"module_crm_iap_enrich,omptempty"`
	ModuleCrmIapMine                            *Bool      `xmlrpc:"module_crm_iap_mine,omptempty"`
	ModuleCurrencyRateLive                      *Bool      `xmlrpc:"module_currency_rate_live,omptempty"`
	ModuleDelivery                              *Bool      `xmlrpc:"module_delivery,omptempty"`
	ModuleDeliveryBpost                         *Bool      `xmlrpc:"module_delivery_bpost,omptempty"`
	ModuleDeliveryDhl                           *Bool      `xmlrpc:"module_delivery_dhl,omptempty"`
	ModuleDeliveryEasypost                      *Bool      `xmlrpc:"module_delivery_easypost,omptempty"`
	ModuleDeliveryFedex                         *Bool      `xmlrpc:"module_delivery_fedex,omptempty"`
	ModuleDeliveryMondialrelay                  *Bool      `xmlrpc:"module_delivery_mondialrelay,omptempty"`
	ModuleDeliverySendcloud                     *Bool      `xmlrpc:"module_delivery_sendcloud,omptempty"`
	ModuleDeliveryUps                           *Bool      `xmlrpc:"module_delivery_ups,omptempty"`
	ModuleDeliveryUsps                          *Bool      `xmlrpc:"module_delivery_usps,omptempty"`
	ModuleEventBarcode                          *Bool      `xmlrpc:"module_event_barcode,omptempty"`
	ModuleEventBooth                            *Bool      `xmlrpc:"module_event_booth,omptempty"`
	ModuleEventSale                             *Bool      `xmlrpc:"module_event_sale,omptempty"`
	ModuleGoogleCalendar                        *Bool      `xmlrpc:"module_google_calendar,omptempty"`
	ModuleGoogleGmail                           *Bool      `xmlrpc:"module_google_gmail,omptempty"`
	ModuleGoogleRecaptcha                       *Bool      `xmlrpc:"module_google_recaptcha,omptempty"`
	ModuleHrAppraisalSurvey                     *Bool      `xmlrpc:"module_hr_appraisal_survey,omptempty"`
	ModuleHrAttendance                          *Bool      `xmlrpc:"module_hr_attendance,omptempty"`
	ModuleHrExpenseExtract                      *Bool      `xmlrpc:"module_hr_expense_extract,omptempty"`
	ModuleHrHomeworking                         *Bool      `xmlrpc:"module_hr_homeworking,omptempty"`
	ModuleHrPayrollAccount                      *Bool      `xmlrpc:"module_hr_payroll_account,omptempty"`
	ModuleHrPayrollAccountSepa                  *Bool      `xmlrpc:"module_hr_payroll_account_sepa,omptempty"`
	ModuleHrPayrollExpense                      *Bool      `xmlrpc:"module_hr_payroll_expense,omptempty"`
	ModuleHrPresence                            *Bool      `xmlrpc:"module_hr_presence,omptempty"`
	ModuleHrRecruitmentExtract                  *Bool      `xmlrpc:"module_hr_recruitment_extract,omptempty"`
	ModuleHrRecruitmentSurvey                   *Bool      `xmlrpc:"module_hr_recruitment_survey,omptempty"`
	ModuleHrSkills                              *Bool      `xmlrpc:"module_hr_skills,omptempty"`
	ModuleHrTimesheet                           *Bool      `xmlrpc:"module_hr_timesheet,omptempty"`
	ModuleL10NBeHrPayroll                       *Bool      `xmlrpc:"module_l10n_be_hr_payroll,omptempty"`
	ModuleL10NEuOss                             *Bool      `xmlrpc:"module_l10n_eu_oss,omptempty"`
	ModuleL10NFrHrPayroll                       *Bool      `xmlrpc:"module_l10n_fr_hr_payroll,omptempty"`
	ModuleL10NInHrPayroll                       *Bool      `xmlrpc:"module_l10n_in_hr_payroll,omptempty"`
	ModuleLoyalty                               *Bool      `xmlrpc:"module_loyalty,omptempty"`
	ModuleMailPlugin                            *Bool      `xmlrpc:"module_mail_plugin,omptempty"`
	ModuleMarketingAutomation                   *Bool      `xmlrpc:"module_marketing_automation,omptempty"`
	ModuleMassMailingSlides                     *Bool      `xmlrpc:"module_mass_mailing_slides,omptempty"`
	ModuleMicrosoftCalendar                     *Bool      `xmlrpc:"module_microsoft_calendar,omptempty"`
	ModuleMicrosoftOutlook                      *Bool      `xmlrpc:"module_microsoft_outlook,omptempty"`
	ModulePartnerAutocomplete                   *Bool      `xmlrpc:"module_partner_autocomplete,omptempty"`
	ModulePaymentPaypal                         *Bool      `xmlrpc:"module_payment_paypal,omptempty"`
	ModuleProductEmailTemplate                  *Bool      `xmlrpc:"module_product_email_template,omptempty"`
	ModuleProductExpiry                         *Bool      `xmlrpc:"module_product_expiry,omptempty"`
	ModuleProductImages                         *Bool      `xmlrpc:"module_product_images,omptempty"`
	ModuleProductMargin                         *Bool      `xmlrpc:"module_product_margin,omptempty"`
	ModuleProjectForecast                       *Bool      `xmlrpc:"module_project_forecast,omptempty"`
	ModuleProjectTimesheetHolidays              *Bool      `xmlrpc:"module_project_timesheet_holidays,omptempty"`
	ModuleProjectTimesheetSynchro               *Bool      `xmlrpc:"module_project_timesheet_synchro,omptempty"`
	ModulePurchaseProductMatrix                 *Bool      `xmlrpc:"module_purchase_product_matrix,omptempty"`
	ModulePurchaseRequisition                   *Bool      `xmlrpc:"module_purchase_requisition,omptempty"`
	ModuleQualityControl                        *Bool      `xmlrpc:"module_quality_control,omptempty"`
	ModuleQualityControlWorksheet               *Bool      `xmlrpc:"module_quality_control_worksheet,omptempty"`
	ModuleSaleAmazon                            *Bool      `xmlrpc:"module_sale_amazon,omptempty"`
	ModuleSaleLoyalty                           *Bool      `xmlrpc:"module_sale_loyalty,omptempty"`
	ModuleSaleMargin                            *Bool      `xmlrpc:"module_sale_margin,omptempty"`
	ModuleSaleProductMatrix                     *Bool      `xmlrpc:"module_sale_product_matrix,omptempty"`
	ModuleSaleQuotationBuilder                  *Bool      `xmlrpc:"module_sale_quotation_builder,omptempty"`
	ModuleSnailmailAccount                      *Bool      `xmlrpc:"module_snailmail_account,omptempty"`
	ModuleSocialDemo                            *Bool      `xmlrpc:"module_social_demo,omptempty"`
	ModuleStockBarcode                          *Bool      `xmlrpc:"module_stock_barcode,omptempty"`
	ModuleStockDropshipping                     *Bool      `xmlrpc:"module_stock_dropshipping,omptempty"`
	ModuleStockLandedCosts                      *Bool      `xmlrpc:"module_stock_landed_costs,omptempty"`
	ModuleStockPickingBatch                     *Bool      `xmlrpc:"module_stock_picking_batch,omptempty"`
	ModuleStockSms                              *Bool      `xmlrpc:"module_stock_sms,omptempty"`
	ModuleVoip                                  *Bool      `xmlrpc:"module_voip,omptempty"`
	ModuleWebUnsplash                           *Bool      `xmlrpc:"module_web_unsplash,omptempty"`
	ModuleWebsiteCrmIapReveal                   *Bool      `xmlrpc:"module_website_crm_iap_reveal,omptempty"`
	ModuleWebsiteEventExhibitor                 *Bool      `xmlrpc:"module_website_event_exhibitor,omptempty"`
	ModuleWebsiteEventMeet                      *Bool      `xmlrpc:"module_website_event_meet,omptempty"`
	ModuleWebsiteEventQuestions                 *Bool      `xmlrpc:"module_website_event_questions,omptempty"`
	ModuleWebsiteEventSale                      *Bool      `xmlrpc:"module_website_event_sale,omptempty"`
	ModuleWebsiteEventTrack                     *Bool      `xmlrpc:"module_website_event_track,omptempty"`
	ModuleWebsiteEventTrackLive                 *Bool      `xmlrpc:"module_website_event_track_live,omptempty"`
	ModuleWebsiteEventTrackQuiz                 *Bool      `xmlrpc:"module_website_event_track_quiz,omptempty"`
	ModuleWebsiteHrRecruitment                  *Bool      `xmlrpc:"module_website_hr_recruitment,omptempty"`
	ModuleWebsiteLivechat                       *Bool      `xmlrpc:"module_website_livechat,omptempty"`
	ModuleWebsiteSaleAutocomplete               *Bool      `xmlrpc:"module_website_sale_autocomplete,omptempty"`
	ModuleWebsiteSaleComparison                 *Bool      `xmlrpc:"module_website_sale_comparison,omptempty"`
	ModuleWebsiteSaleDelivery                   *Bool      `xmlrpc:"module_website_sale_delivery,omptempty"`
	ModuleWebsiteSaleDigital                    *Bool      `xmlrpc:"module_website_sale_digital,omptempty"`
	ModuleWebsiteSalePicking                    *Bool      `xmlrpc:"module_website_sale_picking,omptempty"`
	ModuleWebsiteSaleSlides                     *Bool      `xmlrpc:"module_website_sale_slides,omptempty"`
	ModuleWebsiteSaleWishlist                   *Bool      `xmlrpc:"module_website_sale_wishlist,omptempty"`
	ModuleWebsiteSlidesForum                    *Bool      `xmlrpc:"module_website_slides_forum,omptempty"`
	ModuleWebsiteSlidesSurvey                   *Bool      `xmlrpc:"module_website_slides_survey,omptempty"`
	NotificationRequestBody                     *String    `xmlrpc:"notification_request_body,omptempty"`
	NotificationRequestDelay                    *Int       `xmlrpc:"notification_request_delay,omptempty"`
	NotificationRequestIcon                     *String    `xmlrpc:"notification_request_icon,omptempty"`
	NotificationRequestTitle                    *String    `xmlrpc:"notification_request_title,omptempty"`
	OvertimeCompanyThreshold                    *Int       `xmlrpc:"overtime_company_threshold,omptempty"`
	OvertimeEmployeeThreshold                   *Int       `xmlrpc:"overtime_employee_threshold,omptempty"`
	OvertimeStartDate                           *Time      `xmlrpc:"overtime_start_date,omptempty"`
	PartnerAutocompleteInsufficientCredit       *Bool      `xmlrpc:"partner_autocomplete_insufficient_credit,omptempty"`
	PartnerNamesOrder                           *Selection `xmlrpc:"partner_names_order,omptempty"`
	PartnerNamesOrderChanged                    *Bool      `xmlrpc:"partner_names_order_changed,omptempty"`
	PayInvoicesOnline                           *Bool      `xmlrpc:"pay_invoices_online,omptempty"`
	PeriodLockDate                              *Time      `xmlrpc:"period_lock_date,omptempty"`
	PlanningAllowSelfUnassign                   *Bool      `xmlrpc:"planning_allow_self_unassign,omptempty"`
	PlanningGenerationInterval                  *Int       `xmlrpc:"planning_generation_interval,omptempty"`
	PlanningSelfUnassignDaysBefore              *Int       `xmlrpc:"planning_self_unassign_days_before,omptempty"`
	PlausibleSharedKey                          *String    `xmlrpc:"plausible_shared_key,omptempty"`
	PlausibleSite                               *String    `xmlrpc:"plausible_site,omptempty"`
	PoDoubleValidation                          *Selection `xmlrpc:"po_double_validation,omptempty"`
	PoDoubleValidationAmount                    *Float     `xmlrpc:"po_double_validation_amount,omptempty"`
	PoLead                                      *Float     `xmlrpc:"po_lead,omptempty"`
	PoLock                                      *Selection `xmlrpc:"po_lock,omptempty"`
	PoOrderApproval                             *Bool      `xmlrpc:"po_order_approval,omptempty"`
	PortalAllowApiKeys                          *Bool      `xmlrpc:"portal_allow_api_keys,omptempty"`
	PortalConfirmationPay                       *Bool      `xmlrpc:"portal_confirmation_pay,omptempty"`
	PortalConfirmationSign                      *Bool      `xmlrpc:"portal_confirmation_sign,omptempty"`
	PredictiveLeadScoringFieldLabels            *String    `xmlrpc:"predictive_lead_scoring_field_labels,omptempty"`
	PredictiveLeadScoringFields                 *Relation  `xmlrpc:"predictive_lead_scoring_fields,omptempty"`
	PredictiveLeadScoringFieldsStr              *String    `xmlrpc:"predictive_lead_scoring_fields_str,omptempty"`
	PredictiveLeadScoringStartDate              *Time      `xmlrpc:"predictive_lead_scoring_start_date,omptempty"`
	PredictiveLeadScoringStartDateStr           *String    `xmlrpc:"predictive_lead_scoring_start_date_str,omptempty"`
	PreventOldTimesheetsEncoding                *Bool      `xmlrpc:"prevent_old_timesheets_encoding,omptempty"`
	PreviewReady                                *Bool      `xmlrpc:"preview_ready,omptempty"`
	PrimaryColor                                *String    `xmlrpc:"primary_color,omptempty"`
	ProductFolder                               *Many2One  `xmlrpc:"product_folder,omptempty"`
	ProductPricelistSetting                     *Selection `xmlrpc:"product_pricelist_setting,omptempty"`
	ProductTags                                 *Relation  `xmlrpc:"product_tags,omptempty"`
	ProductVolumeVolumeInCubicFeet              *Selection `xmlrpc:"product_volume_volume_in_cubic_feet,omptempty"`
	ProductWeightInLbs                          *Selection `xmlrpc:"product_weight_in_lbs,omptempty"`
	ProfilingEnabledUntil                       *Time      `xmlrpc:"profiling_enabled_until,omptempty"`
	ProjectTimeModeId                           *Many2One  `xmlrpc:"project_time_mode_id,omptempty"`
	ProvidersState                              *Selection `xmlrpc:"providers_state,omptempty"`
	PurchaseTaxId                               *Many2One  `xmlrpc:"purchase_tax_id,omptempty"`
	QrCode                                      *Bool      `xmlrpc:"qr_code,omptempty"`
	QuickEditMode                               *Selection `xmlrpc:"quick_edit_mode,omptempty"`
	QuotationValidityDays                       *Int       `xmlrpc:"quotation_validity_days,omptempty"`
	RecaptchaMinScore                           *Float     `xmlrpc:"recaptcha_min_score,omptempty"`
	RecaptchaPrivateKey                         *String    `xmlrpc:"recaptcha_private_key,omptempty"`
	RecaptchaPublicKey                          *String    `xmlrpc:"recaptcha_public_key,omptempty"`
	RecruitmentExtractShowOcrOptionSelection    *Selection `xmlrpc:"recruitment_extract_show_ocr_option_selection,omptempty"`
	RecruitmentFolderId                         *Many2One  `xmlrpc:"recruitment_folder_id,omptempty"`
	RecruitmentTagIds                           *Relation  `xmlrpc:"recruitment_tag_ids,omptempty"`
	ReminderManagerAllow                        *Bool      `xmlrpc:"reminder_manager_allow,omptempty"`
	ReminderManagerDelay                        *Int       `xmlrpc:"reminder_manager_delay,omptempty"`
	ReminderManagerInterval                     *Selection `xmlrpc:"reminder_manager_interval,omptempty"`
	ReminderUserAllow                           *Bool      `xmlrpc:"reminder_user_allow,omptempty"`
	ReminderUserDelay                           *Int       `xmlrpc:"reminder_user_delay,omptempty"`
	ReminderUserInterval                        *Selection `xmlrpc:"reminder_user_interval,omptempty"`
	ReportFooter                                *String    `xmlrpc:"report_footer,omptempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	RestDocsSecurityGroupId                     *Many2One  `xmlrpc:"rest_docs_security_group_id,omptempty"`
	RestDocsSecurityGroupXmlid                  *String    `xmlrpc:"rest_docs_security_group_xmlid,omptempty"`
	RestOauth2BearerAutovacuumDays              *Int       `xmlrpc:"rest_oauth2_bearer_autovacuum_days,omptempty"`
	RestOauth2BearerExpiresInSeconds            *Int       `xmlrpc:"rest_oauth2_bearer_expires_in_seconds,omptempty"`
	RestrictTemplateRendering                   *Bool      `xmlrpc:"restrict_template_rendering,omptempty"`
	RuleType                                    *Selection `xmlrpc:"rule_type,omptempty"`
	RulesCompanyId                              *Many2One  `xmlrpc:"rules_company_id,omptempty"`
	SaleAutoUpdatePrice                         *Bool      `xmlrpc:"sale_auto_update_price,omptempty"`
	SaleDeliverySettings                        *Selection `xmlrpc:"sale_delivery_settings,omptempty"`
	SaleTaxId                                   *Many2One  `xmlrpc:"sale_tax_id,omptempty"`
	SalespersonId                               *Many2One  `xmlrpc:"salesperson_id,omptempty"`
	SalesteamId                                 *Many2One  `xmlrpc:"salesteam_id,omptempty"`
	ScrumEstimatesCheck                         *Bool      `xmlrpc:"scrum_estimates_check,omptempty"`
	ScrumEstimatesVals                          *String    `xmlrpc:"scrum_estimates_vals,omptempty"`
	SddCreditorIdentifier                       *String    `xmlrpc:"sdd_creditor_identifier,omptempty"`
	SecondaryColor                              *String    `xmlrpc:"secondary_color,omptempty"`
	SecurityLead                                *Float     `xmlrpc:"security_lead,omptempty"`
	SendAbandonedCartEmail                      *Bool      `xmlrpc:"send_abandoned_cart_email,omptempty"`
	SepaInitiatingPartyName                     *String    `xmlrpc:"sepa_initiating_party_name,omptempty"`
	SepaOrgidId                                 *String    `xmlrpc:"sepa_orgid_id,omptempty"`
	SepaOrgidIssr                               *String    `xmlrpc:"sepa_orgid_issr,omptempty"`
	ServerUriGoogle                             *String    `xmlrpc:"server_uri_google,omptempty"`
	SharedUserAccount                           *Bool      `xmlrpc:"shared_user_account,omptempty"`
	ShowAvailability                            *Bool      `xmlrpc:"show_availability,omptempty"`
	ShowBlacklistButtons                        *Bool      `xmlrpc:"show_blacklist_buttons,omptempty"`
	ShowEffect                                  *Bool      `xmlrpc:"show_effect,omptempty"`
	ShowLineSubtotalsTaxSelection               *Selection `xmlrpc:"show_line_subtotals_tax_selection,omptempty"`
	SnailmailColor                              *Bool      `xmlrpc:"snailmail_color,omptempty"`
	SnailmailCover                              *Bool      `xmlrpc:"snailmail_cover,omptempty"`
	SnailmailDuplex                             *Bool      `xmlrpc:"snailmail_duplex,omptempty"`
	SocialDefaultImage                          *String    `xmlrpc:"social_default_image,omptempty"`
	StartSearchAfterLetter                      *Int       `xmlrpc:"start_search_after_letter,omptempty"`
	StockMoveEmailValidation                    *Bool      `xmlrpc:"stock_move_email_validation,omptempty"`
	StockMoveSmsValidation                      *Bool      `xmlrpc:"stock_move_sms_validation,omptempty"`
	StockSmsConfirmationTemplateId              *Many2One  `xmlrpc:"stock_sms_confirmation_template_id,omptempty"`
	TaxCalculationRoundingMethod                *Selection `xmlrpc:"tax_calculation_rounding_method,omptempty"`
	TaxCashBasisJournalId                       *Many2One  `xmlrpc:"tax_cash_basis_journal_id,omptempty"`
	TaxExigibility                              *Bool      `xmlrpc:"tax_exigibility,omptempty"`
	TaxLockDate                                 *Time      `xmlrpc:"tax_lock_date,omptempty"`
	TermsType                                   *Selection `xmlrpc:"terms_type,omptempty"`
	TermsUrl                                    *String    `xmlrpc:"terms_url,omptempty"`
	TimesheetEncodeUomId                        *Many2One  `xmlrpc:"timesheet_encode_uom_id,omptempty"`
	TimesheetMinDuration                        *Int       `xmlrpc:"timesheet_min_duration,omptempty"`
	TimesheetRounding                           *Int       `xmlrpc:"timesheet_rounding,omptempty"`
	TotalsBelowSections                         *Bool      `xmlrpc:"totals_below_sections,omptempty"`
	TransferAccountId                           *Many2One  `xmlrpc:"transfer_account_id,omptempty"`
	TwilioAccountSid                            *String    `xmlrpc:"twilio_account_sid,omptempty"`
	TwilioAccountToken                          *String    `xmlrpc:"twilio_account_token,omptempty"`
	TwitterConsumerKey                          *String    `xmlrpc:"twitter_consumer_key,omptempty"`
	TwitterConsumerSecretKey                    *String    `xmlrpc:"twitter_consumer_secret_key,omptempty"`
	TwitterUseOwnAccount                        *Bool      `xmlrpc:"twitter_use_own_account,omptempty"`
	UnsplashAccessKey                           *String    `xmlrpc:"unsplash_access_key,omptempty"`
	UnsplashAppId                               *String    `xmlrpc:"unsplash_app_id,omptempty"`
	UseAngloSaxon                               *Bool      `xmlrpc:"use_anglo_saxon,omptempty"`
	UseInvoiceTerms                             *Bool      `xmlrpc:"use_invoice_terms,omptempty"`
	UseMailgateway                              *Bool      `xmlrpc:"use_mailgateway,omptempty"`
	UsePoLead                                   *Bool      `xmlrpc:"use_po_lead,omptempty"`
	UseQuotationValidityDays                    *Bool      `xmlrpc:"use_quotation_validity_days,omptempty"`
	UseSecurityLead                             *Bool      `xmlrpc:"use_security_lead,omptempty"`
	UseTwilioRtcServers                         *Bool      `xmlrpc:"use_twilio_rtc_servers,omptempty"`
	UserDefaultRights                           *Bool      `xmlrpc:"user_default_rights,omptempty"`
	WarehouseId                                 *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WebsiteCompanyId                            *Many2One  `xmlrpc:"website_company_id,omptempty"`
	WebsiteCookiesBar                           *Bool      `xmlrpc:"website_cookies_bar,omptempty"`
	WebsiteDefaultLangCode                      *String    `xmlrpc:"website_default_lang_code,omptempty"`
	WebsiteDefaultLangId                        *Many2One  `xmlrpc:"website_default_lang_id,omptempty"`
	WebsiteDomain                               *String    `xmlrpc:"website_domain,omptempty"`
	WebsiteHomepageUrl                          *String    `xmlrpc:"website_homepage_url,omptempty"`
	WebsiteId                                   *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteLanguageCount                        *Int       `xmlrpc:"website_language_count,omptempty"`
	WebsiteLogo                                 *String    `xmlrpc:"website_logo,omptempty"`
	WebsiteName                                 *String    `xmlrpc:"website_name,omptempty"`
	WebsiteSaleContactUsButtonUrl               *String    `xmlrpc:"website_sale_contact_us_button_url,omptempty"`
	WebsiteSaleEnabledPortalReorderButton       *Bool      `xmlrpc:"website_sale_enabled_portal_reorder_button,omptempty"`
	WebsiteSalePreventZeroPriceSale             *Bool      `xmlrpc:"website_sale_prevent_zero_price_sale,omptempty"`
	WebsiteSlideGoogleAppKey                    *String    `xmlrpc:"website_slide_google_app_key,omptempty"`
	WebsiteWarehouseId                          *Many2One  `xmlrpc:"website_warehouse_id,omptempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omptempty"`
	YoutubeOauthClientId                        *String    `xmlrpc:"youtube_oauth_client_id,omptempty"`
	YoutubeOauthClientSecret                    *String    `xmlrpc:"youtube_oauth_client_secret,omptempty"`
	YoutubeUseOwnAccount                        *Bool      `xmlrpc:"youtube_use_own_account,omptempty"`
}

// ResConfigSettingss represents array of res.config.settings model.
type ResConfigSettingss []ResConfigSettings

// ResConfigSettingsModel is the odoo model name.
const ResConfigSettingsModel = "res.config.settings"

// Many2One convert ResConfigSettings to *Many2One.
func (rcs *ResConfigSettings) Many2One() *Many2One {
	return NewMany2One(rcs.Id.Get(), "")
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettings(rcs *ResConfigSettings) (int64, error) {
	ids, err := c.CreateResConfigSettingss([]*ResConfigSettings{rcs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigSettings creates a new res.config.settings model and returns its id.
func (c *Client) CreateResConfigSettingss(rcss []*ResConfigSettings) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcss {
		vv = append(vv, v)
	}
	return c.Create(ResConfigSettingsModel, vv)
}

// UpdateResConfigSettings updates an existing res.config.settings record.
func (c *Client) UpdateResConfigSettings(rcs *ResConfigSettings) error {
	return c.UpdateResConfigSettingss([]int64{rcs.Id.Get()}, rcs)
}

// UpdateResConfigSettingss updates existing res.config.settings records.
// All records (represented by ids) will be updated by rcs values.
func (c *Client) UpdateResConfigSettingss(ids []int64, rcs *ResConfigSettings) error {
	return c.Update(ResConfigSettingsModel, ids, rcs)
}

// DeleteResConfigSettings deletes an existing res.config.settings record.
func (c *Client) DeleteResConfigSettings(id int64) error {
	return c.DeleteResConfigSettingss([]int64{id})
}

// DeleteResConfigSettingss deletes existing res.config.settings records.
func (c *Client) DeleteResConfigSettingss(ids []int64) error {
	return c.Delete(ResConfigSettingsModel, ids)
}

// GetResConfigSettings gets res.config.settings existing record.
func (c *Client) GetResConfigSettings(id int64) (*ResConfigSettings, error) {
	rcss, err := c.GetResConfigSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.config.settings not found", id)
}

// GetResConfigSettingss gets res.config.settings existing records.
func (c *Client) GetResConfigSettingss(ids []int64) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.Read(ResConfigSettingsModel, ids, nil, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettings finds res.config.settings record by querying it with criteria.
func (c *Client) FindResConfigSettings(criteria *Criteria) (*ResConfigSettings, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, NewOptions().Limit(1), rcss); err != nil {
		return nil, err
	}
	if rcss != nil && len(*rcss) > 0 {
		return &((*rcss)[0]), nil
	}
	return nil, fmt.Errorf("res.config.settings was not found with criteria %v", criteria)
}

// FindResConfigSettingss finds res.config.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingss(criteria *Criteria, options *Options) (*ResConfigSettingss, error) {
	rcss := &ResConfigSettingss{}
	if err := c.SearchRead(ResConfigSettingsModel, criteria, options, rcss); err != nil {
		return nil, err
	}
	return rcss, nil
}

// FindResConfigSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResConfigSettingsId finds record id by querying it with criteria.
func (c *Client) FindResConfigSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.config.settings was not found with criteria %v and options %v", criteria, options)
}
