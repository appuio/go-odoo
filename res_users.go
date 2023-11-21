package odoo

import (
	"fmt"
)

// ResUsers represents res.users model.
type ResUsers struct {
	LastUpdate                         *Time      `xmlrpc:"__last_update,omptempty"`
	AccessesCount                      *Int       `xmlrpc:"accesses_count,omptempty"`
	AccountManager                     *Many2One  `xmlrpc:"account_manager,omptempty"`
	AccountRepresentedCompanyIds       *Relation  `xmlrpc:"account_represented_company_ids,omptempty"`
	ActionId                           *Many2One  `xmlrpc:"action_id,omptempty"`
	Active                             *Bool      `xmlrpc:"active,omptempty"`
	ActiveLangCount                    *Int       `xmlrpc:"active_lang_count,omptempty"`
	ActivePartner                      *Bool      `xmlrpc:"active_partner,omptempty"`
	ActivityCalendarEventId            *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline               *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration        *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon              *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                        *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                      *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                    *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                   *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                     *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                     *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalInfo                     *String    `xmlrpc:"additional_info,omptempty"`
	AdditionalNote                     *String    `xmlrpc:"additional_note,omptempty"`
	AddressHomeId                      *Many2One  `xmlrpc:"address_home_id,omptempty"`
	AddressId                          *Many2One  `xmlrpc:"address_id,omptempty"`
	AllocationCount                    *Float     `xmlrpc:"allocation_count,omptempty"`
	AllocationDisplay                  *String    `xmlrpc:"allocation_display,omptempty"`
	AllocationRemainingDisplay         *String    `xmlrpc:"allocation_remaining_display,omptempty"`
	ApiKeyIds                          *Relation  `xmlrpc:"api_key_ids,omptempty"`
	AttendanceState                    *Selection `xmlrpc:"attendance_state,omptempty"`
	Avatar1024                         *String    `xmlrpc:"avatar_1024,omptempty"`
	Avatar128                          *String    `xmlrpc:"avatar_128,omptempty"`
	Avatar1920                         *String    `xmlrpc:"avatar_1920,omptempty"`
	Avatar256                          *String    `xmlrpc:"avatar_256,omptempty"`
	Avatar512                          *String    `xmlrpc:"avatar_512,omptempty"`
	BadgeIds                           *Relation  `xmlrpc:"badge_ids,omptempty"`
	BankAccountCount                   *Int       `xmlrpc:"bank_account_count,omptempty"`
	BankAccountId                      *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	BankIds                            *Relation  `xmlrpc:"bank_ids,omptempty"`
	Barcode                            *String    `xmlrpc:"barcode,omptempty"`
	Birthday                           *Time      `xmlrpc:"birthday,omptempty"`
	BronzeBadge                        *Int       `xmlrpc:"bronze_badge,omptempty"`
	CalendarLastNotifAck               *Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CanEdit                            *Bool      `xmlrpc:"can_edit,omptempty"`
	CanPublish                         *Bool      `xmlrpc:"can_publish,omptempty"`
	CategoryId                         *Relation  `xmlrpc:"category_id,omptempty"`
	CategoryIds                        *Relation  `xmlrpc:"category_ids,omptempty"`
	Certificate                        *Selection `xmlrpc:"certificate,omptempty"`
	CertificationsCompanyCount         *Int       `xmlrpc:"certifications_company_count,omptempty"`
	CertificationsCount                *Int       `xmlrpc:"certifications_count,omptempty"`
	ChannelIds                         *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                           *Relation  `xmlrpc:"child_ids,omptempty"`
	Children                           *Int       `xmlrpc:"children,omptempty"`
	City                               *String    `xmlrpc:"city,omptempty"`
	CoachId                            *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                              *Int       `xmlrpc:"color,omptempty"`
	Comment                            *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName              *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerId                *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompaniesCount                     *Int       `xmlrpc:"companies_count,omptempty"`
	CompanyId                          *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyIds                         *Relation  `xmlrpc:"company_ids,omptempty"`
	CompanyName                        *String    `xmlrpc:"company_name,omptempty"`
	CompanyRegistry                    *String    `xmlrpc:"company_registry,omptempty"`
	CompanyType                        *Selection `xmlrpc:"company_type,omptempty"`
	ContactAddress                     *String    `xmlrpc:"contact_address,omptempty"`
	ContactAddressComplete             *String    `xmlrpc:"contact_address_complete,omptempty"`
	ContractIds                        *Relation  `xmlrpc:"contract_ids,omptempty"`
	CountryCode                        *String    `xmlrpc:"country_code,omptempty"`
	CountryId                          *Many2One  `xmlrpc:"country_id,omptempty"`
	CountryOfBirth                     *Many2One  `xmlrpc:"country_of_birth,omptempty"`
	CreateDate                         *Time      `xmlrpc:"create_date,omptempty"`
	CreateEmployee                     *Bool      `xmlrpc:"create_employee,omptempty"`
	CreateEmployeeId                   *Many2One  `xmlrpc:"create_employee_id,omptempty"`
	CreateUid                          *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                             *Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                        *Float     `xmlrpc:"credit_limit,omptempty"`
	CrmTeamIds                         *Relation  `xmlrpc:"crm_team_ids,omptempty"`
	CrmTeamMemberIds                   *Relation  `xmlrpc:"crm_team_member_ids,omptempty"`
	CurrencyId                         *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrentLeaveState                  *Selection `xmlrpc:"current_leave_state,omptempty"`
	CustomerRank                       *Int       `xmlrpc:"customer_rank,omptempty"`
	CustomerSummary                    *String    `xmlrpc:"customer_summary,omptempty"`
	Date                               *Time      `xmlrpc:"date,omptempty"`
	Debit                              *Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                         *Float     `xmlrpc:"debit_limit,omptempty"`
	DepartmentId                       *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                        *String    `xmlrpc:"display_name,omptempty"`
	DocumentCount                      *Int       `xmlrpc:"document_count,omptempty"`
	DocumentIds                        *Relation  `xmlrpc:"document_ids,omptempty"`
	DuplicatedBankAccountPartnersCount *Int       `xmlrpc:"duplicated_bank_account_partners_count,omptempty"`
	Email                              *String    `xmlrpc:"email,omptempty"`
	EmailBounced                       *Bool      `xmlrpc:"email_bounced,omptempty"`
	EmailFormatted                     *String    `xmlrpc:"email_formatted,omptempty"`
	EmailNormalized                    *String    `xmlrpc:"email_normalized,omptempty"`
	EmailScore                         *Float     `xmlrpc:"email_score,omptempty"`
	EmergencyContact                   *String    `xmlrpc:"emergency_contact,omptempty"`
	EmergencyPhone                     *String    `xmlrpc:"emergency_phone,omptempty"`
	Employee                           *Bool      `xmlrpc:"employee,omptempty"`
	EmployeeBankAccountId              *Many2One  `xmlrpc:"employee_bank_account_id,omptempty"`
	EmployeeCount                      *Int       `xmlrpc:"employee_count,omptempty"`
	EmployeeCountryId                  *Many2One  `xmlrpc:"employee_country_id,omptempty"`
	EmployeeId                         *Many2One  `xmlrpc:"employee_id,omptempty"`
	EmployeeIds                        *Relation  `xmlrpc:"employee_ids,omptempty"`
	EmployeeParentId                   *Many2One  `xmlrpc:"employee_parent_id,omptempty"`
	EmployeePhone                      *String    `xmlrpc:"employee_phone,omptempty"`
	EmployeeResourceCalendarId         *Many2One  `xmlrpc:"employee_resource_calendar_id,omptempty"`
	EmployeeSkillIds                   *Relation  `xmlrpc:"employee_skill_ids,omptempty"`
	EmployeeType                       *Selection `xmlrpc:"employee_type,omptempty"`
	EmployeesCount                     *Int       `xmlrpc:"employees_count,omptempty"`
	EquipmentCount                     *Int       `xmlrpc:"equipment_count,omptempty"`
	EquipmentIds                       *Relation  `xmlrpc:"equipment_ids,omptempty"`
	EventCount                         *Int       `xmlrpc:"event_count,omptempty"`
	ExpenseManagerId                   *Many2One  `xmlrpc:"expense_manager_id,omptempty"`
	FailedMessageIds                   *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	Firstname                          *String    `xmlrpc:"firstname,omptempty"`
	FollowupLineId                     *Many2One  `xmlrpc:"followup_line_id,omptempty"`
	FollowupNextActionDate             *Time      `xmlrpc:"followup_next_action_date,omptempty"`
	FollowupReminderType               *Selection `xmlrpc:"followup_reminder_type,omptempty"`
	FollowupResponsibleId              *Many2One  `xmlrpc:"followup_responsible_id,omptempty"`
	FollowupStatus                     *Selection `xmlrpc:"followup_status,omptempty"`
	Function                           *String    `xmlrpc:"function,omptempty"`
	Gender                             *Selection `xmlrpc:"gender,omptempty"`
	GoalIds                            *Relation  `xmlrpc:"goal_ids,omptempty"`
	GoldBadge                          *Int       `xmlrpc:"gold_badge,omptempty"`
	GroupsCount                        *Int       `xmlrpc:"groups_count,omptempty"`
	GroupsId                           *Relation  `xmlrpc:"groups_id,omptempty"`
	HasMessage                         *Bool      `xmlrpc:"has_message,omptempty"`
	HasUnreconciledEntries             *Bool      `xmlrpc:"has_unreconciled_entries,omptempty"`
	HelpdeskTargetClosed               *Int       `xmlrpc:"helpdesk_target_closed,omptempty"`
	HelpdeskTargetRating               *Float     `xmlrpc:"helpdesk_target_rating,omptempty"`
	HelpdeskTargetSuccess              *Float     `xmlrpc:"helpdesk_target_success,omptempty"`
	HoursLastMonth                     *Float     `xmlrpc:"hours_last_month,omptempty"`
	HoursLastMonthDisplay              *String    `xmlrpc:"hours_last_month_display,omptempty"`
	HrIconDisplay                      *Selection `xmlrpc:"hr_icon_display,omptempty"`
	HrPresenceState                    *Selection `xmlrpc:"hr_presence_state,omptempty"`
	IapEnrichInfo                      *String    `xmlrpc:"iap_enrich_info,omptempty"`
	IapSearchDomain                    *String    `xmlrpc:"iap_search_domain,omptempty"`
	Id                                 *Int       `xmlrpc:"id,omptempty"`
	IdentificationId                   *String    `xmlrpc:"identification_id,omptempty"`
	ImStatus                           *String    `xmlrpc:"im_status,omptempty"`
	Image1024                          *String    `xmlrpc:"image_1024,omptempty"`
	Image128                           *String    `xmlrpc:"image_128,omptempty"`
	Image1920                          *String    `xmlrpc:"image_1920,omptempty"`
	Image256                           *String    `xmlrpc:"image_256,omptempty"`
	Image512                           *String    `xmlrpc:"image_512,omptempty"`
	ImageMedium                        *String    `xmlrpc:"image_medium,omptempty"`
	IndustryId                         *Many2One  `xmlrpc:"industry_id,omptempty"`
	InvoiceIds                         *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceWarn                        *Selection `xmlrpc:"invoice_warn,omptempty"`
	InvoiceWarnMsg                     *String    `xmlrpc:"invoice_warn_msg,omptempty"`
	IsAbsent                           *Bool      `xmlrpc:"is_absent,omptempty"`
	IsAddressHomeACompany              *Bool      `xmlrpc:"is_address_home_a_company,omptempty"`
	IsBlacklisted                      *Bool      `xmlrpc:"is_blacklisted,omptempty"`
	IsCompany                          *Bool      `xmlrpc:"is_company,omptempty"`
	IsPublic                           *Bool      `xmlrpc:"is_public,omptempty"`
	IsPublished                        *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized                     *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	IsSystem                           *Bool      `xmlrpc:"is_system,omptempty"`
	JobTitle                           *String    `xmlrpc:"job_title,omptempty"`
	JournalItemCount                   *Int       `xmlrpc:"journal_item_count,omptempty"`
	Karma                              *Int       `xmlrpc:"karma,omptempty"`
	KarmaTrackingIds                   *Relation  `xmlrpc:"karma_tracking_ids,omptempty"`
	KmHomeWork                         *Int       `xmlrpc:"km_home_work,omptempty"`
	Lang                               *Selection `xmlrpc:"lang,omptempty"`
	LastActivity                       *Time      `xmlrpc:"last_activity,omptempty"`
	LastActivityTime                   *String    `xmlrpc:"last_activity_time,omptempty"`
	LastAppraisalDate                  *Time      `xmlrpc:"last_appraisal_date,omptempty"`
	LastAppraisalId                    *Many2One  `xmlrpc:"last_appraisal_id,omptempty"`
	LastCheckIn                        *Time      `xmlrpc:"last_check_in,omptempty"`
	LastCheckOut                       *Time      `xmlrpc:"last_check_out,omptempty"`
	LastPasswordChange                 *Time      `xmlrpc:"last_password_change,omptempty"`
	LastTimeEntriesChecked             *Time      `xmlrpc:"last_time_entries_checked,omptempty"`
	LastWebsiteSoId                    *Many2One  `xmlrpc:"last_website_so_id,omptempty"`
	Lastname                           *String    `xmlrpc:"lastname,omptempty"`
	LeaveDateTo                        *Time      `xmlrpc:"leave_date_to,omptempty"`
	LeaveManagerId                     *Many2One  `xmlrpc:"leave_manager_id,omptempty"`
	LivechatUsername                   *String    `xmlrpc:"livechat_username,omptempty"`
	LogIds                             *Relation  `xmlrpc:"log_ids,omptempty"`
	Login                              *String    `xmlrpc:"login,omptempty"`
	LoginDate                          *Time      `xmlrpc:"login_date,omptempty"`
	Marital                            *Selection `xmlrpc:"marital,omptempty"`
	MassMailingContactIds              *Relation  `xmlrpc:"mass_mailing_contact_ids,omptempty"`
	MassMailingContactsCount           *Int       `xmlrpc:"mass_mailing_contacts_count,omptempty"`
	MassMailingStatsCount              *Int       `xmlrpc:"mass_mailing_stats_count,omptempty"`
	MassMailingStatsIds                *Relation  `xmlrpc:"mass_mailing_stats_ids,omptempty"`
	MeetingCount                       *Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                         *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MessageAttachmentCount             *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageBounce                      *Int       `xmlrpc:"message_bounce,omptempty"`
	MessageContent                     *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                 *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                    *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter             *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                 *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                         *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                  *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId            *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                  *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter           *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                  *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessagesFromCount                  *Int       `xmlrpc:"messages_from_count,omptempty"`
	MessagesToCount                    *Int       `xmlrpc:"messages_to_count,omptempty"`
	Mobile                             *String    `xmlrpc:"mobile,omptempty"`
	MobileBlacklisted                  *Bool      `xmlrpc:"mobile_blacklisted,omptempty"`
	MobilePhone                        *String    `xmlrpc:"mobile_phone,omptempty"`
	MyActivityDateDeadline             *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                               *String    `xmlrpc:"name,omptempty"`
	NewPassword                        *String    `xmlrpc:"new_password,omptempty"`
	NextAppraisalDate                  *Time      `xmlrpc:"next_appraisal_date,omptempty"`
	NextRankId                         *Many2One  `xmlrpc:"next_rank_id,omptempty"`
	NotificationType                   *Selection `xmlrpc:"notification_type,omptempty"`
	Oauth1SessionIds                   *Relation  `xmlrpc:"oauth1_session_ids,omptempty"`
	Oauth2SessionIds                   *Relation  `xmlrpc:"oauth2_session_ids,omptempty"`
	OauthAccessToken                   *String    `xmlrpc:"oauth_access_token,omptempty"`
	OauthProviderId                    *Many2One  `xmlrpc:"oauth_provider_id,omptempty"`
	OauthUid                           *String    `xmlrpc:"oauth_uid,omptempty"`
	OcnToken                           *String    `xmlrpc:"ocn_token,omptempty"`
	OdoobotFailed                      *Bool      `xmlrpc:"odoobot_failed,omptempty"`
	OdoobotState                       *Selection `xmlrpc:"odoobot_state,omptempty"`
	OnTimeRate                         *Float     `xmlrpc:"on_time_rate,omptempty"`
	OnlinePartnerInformation           *String    `xmlrpc:"online_partner_information,omptempty"`
	OpportunityCount                   *Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityIds                     *Relation  `xmlrpc:"opportunity_ids,omptempty"`
	ParentId                           *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                         *String    `xmlrpc:"parent_name,omptempty"`
	PartnerGid                         *Int       `xmlrpc:"partner_gid,omptempty"`
	PartnerId                          *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerLatitude                    *Float     `xmlrpc:"partner_latitude,omptempty"`
	PartnerLongitude                   *Float     `xmlrpc:"partner_longitude,omptempty"`
	PartnerShare                       *Bool      `xmlrpc:"partner_share,omptempty"`
	PassportId                         *String    `xmlrpc:"passport_id,omptempty"`
	Password                           *String    `xmlrpc:"password,omptempty"`
	PaymentTokenCount                  *Int       `xmlrpc:"payment_token_count,omptempty"`
	PaymentTokenIds                    *Relation  `xmlrpc:"payment_token_ids,omptempty"`
	PermitNo                           *String    `xmlrpc:"permit_no,omptempty"`
	Phone                              *String    `xmlrpc:"phone,omptempty"`
	PhoneBlacklisted                   *Bool      `xmlrpc:"phone_blacklisted,omptempty"`
	PhoneMobileSearch                  *String    `xmlrpc:"phone_mobile_search,omptempty"`
	PhoneSanitized                     *String    `xmlrpc:"phone_sanitized,omptempty"`
	PhoneSanitizedBlacklisted          *Bool      `xmlrpc:"phone_sanitized_blacklisted,omptempty"`
	PickingWarn                        *Selection `xmlrpc:"picking_warn,omptempty"`
	PickingWarnMsg                     *String    `xmlrpc:"picking_warn_msg,omptempty"`
	Pin                                *String    `xmlrpc:"pin,omptempty"`
	PlaceOfBirth                       *String    `xmlrpc:"place_of_birth,omptempty"`
	PrivateCity                        *String    `xmlrpc:"private_city,omptempty"`
	PrivateCountryId                   *Many2One  `xmlrpc:"private_country_id,omptempty"`
	PrivateEmail                       *String    `xmlrpc:"private_email,omptempty"`
	PrivateLang                        *Selection `xmlrpc:"private_lang,omptempty"`
	PrivateStateId                     *Many2One  `xmlrpc:"private_state_id,omptempty"`
	PrivateStreet                      *String    `xmlrpc:"private_street,omptempty"`
	PrivateStreet2                     *String    `xmlrpc:"private_street2,omptempty"`
	PrivateZip                         *String    `xmlrpc:"private_zip,omptempty"`
	PropertyAccountPayableId           *Many2One  `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountPositionId          *Many2One  `xmlrpc:"property_account_position_id,omptempty"`
	PropertyAccountReceivableId        *Many2One  `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyPaymentTermId              *Many2One  `xmlrpc:"property_payment_term_id,omptempty"`
	PropertyProductPricelist           *Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertyPurchaseCurrencyId         *Many2One  `xmlrpc:"property_purchase_currency_id,omptempty"`
	PropertyStockCustomer              *Many2One  `xmlrpc:"property_stock_customer,omptempty"`
	PropertyStockSupplier              *Many2One  `xmlrpc:"property_stock_supplier,omptempty"`
	PropertySupplierPaymentTermId      *Many2One  `xmlrpc:"property_supplier_payment_term_id,omptempty"`
	PropertyWarehouseId                *Many2One  `xmlrpc:"property_warehouse_id,omptempty"`
	PurchaseLineIds                    *Relation  `xmlrpc:"purchase_line_ids,omptempty"`
	PurchaseOrderCount                 *Int       `xmlrpc:"purchase_order_count,omptempty"`
	PurchaseWarn                       *Selection `xmlrpc:"purchase_warn,omptempty"`
	PurchaseWarnMsg                    *String    `xmlrpc:"purchase_warn_msg,omptempty"`
	RankId                             *Many2One  `xmlrpc:"rank_id,omptempty"`
	ReceiptReminderEmail               *Bool      `xmlrpc:"receipt_reminder_email,omptempty"`
	Ref                                *String    `xmlrpc:"ref,omptempty"`
	RefCompanyIds                      *Relation  `xmlrpc:"ref_company_ids,omptempty"`
	RegistrationIds                    *Relation  `xmlrpc:"registration_ids,omptempty"`
	ReminderDateBeforeReceipt          *Int       `xmlrpc:"reminder_date_before_receipt,omptempty"`
	RequestOvertime                    *Bool      `xmlrpc:"request_overtime,omptempty"`
	ResUsersSettingsId                 *Many2One  `xmlrpc:"res_users_settings_id,omptempty"`
	ResUsersSettingsIds                *Relation  `xmlrpc:"res_users_settings_ids,omptempty"`
	ResourceCalendarId                 *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceIds                        *Relation  `xmlrpc:"resource_ids,omptempty"`
	ResumeLineIds                      *Relation  `xmlrpc:"resume_line_ids,omptempty"`
	RoleIds                            *Relation  `xmlrpc:"role_ids,omptempty"`
	RoleLineIds                        *Relation  `xmlrpc:"role_line_ids,omptempty"`
	RulesCount                         *Int       `xmlrpc:"rules_count,omptempty"`
	SaleOrderCount                     *Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                       *Relation  `xmlrpc:"sale_order_ids,omptempty"`
	SaleTeamId                         *Many2One  `xmlrpc:"sale_team_id,omptempty"`
	SaleWarn                           *Selection `xmlrpc:"sale_warn,omptempty"`
	SaleWarnMsg                        *String    `xmlrpc:"sale_warn_msg,omptempty"`
	SameCompanyRegistryPartnerId       *Many2One  `xmlrpc:"same_company_registry_partner_id,omptempty"`
	SameVatPartnerId                   *Many2One  `xmlrpc:"same_vat_partner_id,omptempty"`
	SddCount                           *Int       `xmlrpc:"sdd_count,omptempty"`
	SddMandateIds                      *Relation  `xmlrpc:"sdd_mandate_ids,omptempty"`
	Self                               *Many2One  `xmlrpc:"self,omptempty"`
	SeoName                            *String    `xmlrpc:"seo_name,omptempty"`
	Share                              *Bool      `xmlrpc:"share,omptempty"`
	ShowCreditLimit                    *Bool      `xmlrpc:"show_credit_limit,omptempty"`
	ShowLeaves                         *Bool      `xmlrpc:"show_leaves,omptempty"`
	ShowTechnicalFeatures              *Bool      `xmlrpc:"show_technical_features,omptempty"`
	Signature                          *String    `xmlrpc:"signature,omptempty"`
	SignupExpiration                   *Time      `xmlrpc:"signup_expiration,omptempty"`
	SignupToken                        *String    `xmlrpc:"signup_token,omptempty"`
	SignupType                         *String    `xmlrpc:"signup_type,omptempty"`
	SignupUrl                          *String    `xmlrpc:"signup_url,omptempty"`
	SignupValid                        *Bool      `xmlrpc:"signup_valid,omptempty"`
	SilverBadge                        *Int       `xmlrpc:"silver_badge,omptempty"`
	SlaIds                             *Relation  `xmlrpc:"sla_ids,omptempty"`
	SlideChannelCompanyCount           *Int       `xmlrpc:"slide_channel_company_count,omptempty"`
	SlideChannelCompletedIds           *Relation  `xmlrpc:"slide_channel_completed_ids,omptempty"`
	SlideChannelCount                  *Int       `xmlrpc:"slide_channel_count,omptempty"`
	SlideChannelIds                    *Relation  `xmlrpc:"slide_channel_ids,omptempty"`
	Spammer                            *Bool      `xmlrpc:"spammer,omptempty"`
	SpouseBirthdate                    *Time      `xmlrpc:"spouse_birthdate,omptempty"`
	SpouseCompleteName                 *String    `xmlrpc:"spouse_complete_name,omptempty"`
	State                              *Selection `xmlrpc:"state,omptempty"`
	StateId                            *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                             *String    `xmlrpc:"street,omptempty"`
	Street2                            *String    `xmlrpc:"street2,omptempty"`
	StudyField                         *String    `xmlrpc:"study_field,omptempty"`
	StudySchool                        *String    `xmlrpc:"study_school,omptempty"`
	SubscriptionCount                  *Int       `xmlrpc:"subscription_count,omptempty"`
	SupplierInvoiceCount               *Int       `xmlrpc:"supplier_invoice_count,omptempty"`
	SupplierRank                       *Int       `xmlrpc:"supplier_rank,omptempty"`
	TargetSalesDone                    *Int       `xmlrpc:"target_sales_done,omptempty"`
	TargetSalesInvoiced                *Int       `xmlrpc:"target_sales_invoiced,omptempty"`
	TargetSalesWon                     *Int       `xmlrpc:"target_sales_won,omptempty"`
	TaskCount                          *Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                            *Relation  `xmlrpc:"task_ids,omptempty"`
	TeamId                             *Many2One  `xmlrpc:"team_id,omptempty"`
	TechnicalFeatures                  *Bool      `xmlrpc:"technical_features,omptempty"`
	TicketCount                        *Int       `xmlrpc:"ticket_count,omptempty"`
	TimesheetManagerId                 *Many2One  `xmlrpc:"timesheet_manager_id,omptempty"`
	Title                              *Many2One  `xmlrpc:"title,omptempty"`
	TotalDue                           *Float     `xmlrpc:"total_due,omptempty"`
	TotalInvoiced                      *Float     `xmlrpc:"total_invoiced,omptempty"`
	TotalOverdue                       *Float     `xmlrpc:"total_overdue,omptempty"`
	TotalOvertime                      *Float     `xmlrpc:"total_overtime,omptempty"`
	TotpEnabled                        *Bool      `xmlrpc:"totp_enabled,omptempty"`
	TotpSecret                         *String    `xmlrpc:"totp_secret,omptempty"`
	TotpTrustedDeviceIds               *Relation  `xmlrpc:"totp_trusted_device_ids,omptempty"`
	TrackingEmailsCount                *Int       `xmlrpc:"tracking_emails_count,omptempty"`
	Trust                              *Selection `xmlrpc:"trust,omptempty"`
	Type                               *Selection `xmlrpc:"type,omptempty"`
	Tz                                 *Selection `xmlrpc:"tz,omptempty"`
	TzOffset                           *String    `xmlrpc:"tz_offset,omptempty"`
	UnpaidInvoiceIds                   *Relation  `xmlrpc:"unpaid_invoice_ids,omptempty"`
	UnpaidInvoicesCount                *Int       `xmlrpc:"unpaid_invoices_count,omptempty"`
	UnreconciledAmlIds                 *Relation  `xmlrpc:"unreconciled_aml_ids,omptempty"`
	UsePartnerCreditLimit              *Bool      `xmlrpc:"use_partner_credit_limit,omptempty"`
	UserGroupWarning                   *String    `xmlrpc:"user_group_warning,omptempty"`
	UserId                             *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                            *Relation  `xmlrpc:"user_ids,omptempty"`
	UserLivechatUsername               *String    `xmlrpc:"user_livechat_username,omptempty"`
	Vat                                *String    `xmlrpc:"vat,omptempty"`
	Vehicle                            *String    `xmlrpc:"vehicle,omptempty"`
	VisaExpire                         *Time      `xmlrpc:"visa_expire,omptempty"`
	VisaNo                             *String    `xmlrpc:"visa_no,omptempty"`
	VisitorIds                         *Relation  `xmlrpc:"visitor_ids,omptempty"`
	Website                            *String    `xmlrpc:"website,omptempty"`
	WebsiteDescription                 *String    `xmlrpc:"website_description,omptempty"`
	WebsiteId                          *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMessageIds                  *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription             *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords                *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg                   *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle                   *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished                   *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteShortDescription            *String    `xmlrpc:"website_short_description,omptempty"`
	WebsiteUrl                         *String    `xmlrpc:"website_url,omptempty"`
	WizardType                         *Selection `xmlrpc:"wizard_type,omptempty"`
	WorkEmail                          *String    `xmlrpc:"work_email,omptempty"`
	WorkLocationId                     *Many2One  `xmlrpc:"work_location_id,omptempty"`
	WorkPhone                          *String    `xmlrpc:"work_phone,omptempty"`
	WriteDate                          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                           *Many2One  `xmlrpc:"write_uid,omptempty"`
	XClassification                    *Selection `xmlrpc:"x_classification,omptempty"`
	XHubspotId                         *String    `xmlrpc:"x_hubspot_id,omptempty"`
	XInvoiceContact                    *String    `xmlrpc:"x_invoice_contact,omptempty"`
	XOdoo8ID                           *String    `xmlrpc:"x_odoo_8_ID,omptempty"`
	XStudioMany2ManyField2VqZb         *Relation  `xmlrpc:"x_studio_many2many_field_2vqZb,omptempty"`
	XStudioMany2OneFieldP3Zm5          *Many2One  `xmlrpc:"x_studio_many2one_field_P3Zm5,omptempty"`
	XStudioMany2OneFieldGaElv          *Many2One  `xmlrpc:"x_studio_many2one_field_gaElv,omptempty"`
	XStudioRelatedFieldHJLOY           *Selection `xmlrpc:"x_studio_related_field_hJLOY,omptempty"`
	XStudioRelatedFieldSmqd9           *Many2One  `xmlrpc:"x_studio_related_field_smqd9,omptempty"`
	Zip                                *String    `xmlrpc:"zip,omptempty"`
}

// ResUserss represents array of res.users model.
type ResUserss []ResUsers

// ResUsersModel is the odoo model name.
const ResUsersModel = "res.users"

// Many2One convert ResUsers to *Many2One.
func (ru *ResUsers) Many2One() *Many2One {
	return NewMany2One(ru.Id.Get(), "")
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUsers(ru *ResUsers) (int64, error) {
	ids, err := c.CreateResUserss([]*ResUsers{ru})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUserss(rus []*ResUsers) ([]int64, error) {
	var vv []interface{}
	for _, v := range rus {
		vv = append(vv, v)
	}
	return c.Create(ResUsersModel, vv)
}

// UpdateResUsers updates an existing res.users record.
func (c *Client) UpdateResUsers(ru *ResUsers) error {
	return c.UpdateResUserss([]int64{ru.Id.Get()}, ru)
}

// UpdateResUserss updates existing res.users records.
// All records (represented by ids) will be updated by ru values.
func (c *Client) UpdateResUserss(ids []int64, ru *ResUsers) error {
	return c.Update(ResUsersModel, ids, ru)
}

// DeleteResUsers deletes an existing res.users record.
func (c *Client) DeleteResUsers(id int64) error {
	return c.DeleteResUserss([]int64{id})
}

// DeleteResUserss deletes existing res.users records.
func (c *Client) DeleteResUserss(ids []int64) error {
	return c.Delete(ResUsersModel, ids)
}

// GetResUsers gets res.users existing record.
func (c *Client) GetResUsers(id int64) (*ResUsers, error) {
	rus, err := c.GetResUserss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users not found", id)
}

// GetResUserss gets res.users existing records.
func (c *Client) GetResUserss(ids []int64) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.Read(ResUsersModel, ids, nil, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsers finds res.users record by querying it with criteria.
func (c *Client) FindResUsers(criteria *Criteria) (*ResUsers, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, NewOptions().Limit(1), rus); err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("res.users was not found with criteria %v", criteria)
}

// FindResUserss finds res.users records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUserss(criteria *Criteria, options *Options) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, options, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersId finds record id by querying it with criteria.
func (c *Client) FindResUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users was not found with criteria %v and options %v", criteria, options)
}
