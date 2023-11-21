package odoo

import (
	"fmt"
)

// CrmLead represents crm.lead model.
type CrmLead struct {
	LastUpdate                      *Time       `xmlrpc:"__last_update,omptempty"`
	Active                          *Bool       `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId         *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline            *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration     *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon           *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                     *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState                   *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                 *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                  *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                  *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	AutomatedProbability            *Float      `xmlrpc:"automated_probability,omptempty"`
	CalendarEventCount              *Int        `xmlrpc:"calendar_event_count,omptempty"`
	CalendarEventIds                *Relation   `xmlrpc:"calendar_event_ids,omptempty"`
	CampaignId                      *Many2One   `xmlrpc:"campaign_id,omptempty"`
	City                            *String     `xmlrpc:"city,omptempty"`
	Color                           *Int        `xmlrpc:"color,omptempty"`
	CompanyCurrency                 *Many2One   `xmlrpc:"company_currency,omptempty"`
	CompanyId                       *Many2One   `xmlrpc:"company_id,omptempty"`
	ContactName                     *String     `xmlrpc:"contact_name,omptempty"`
	CountryId                       *Many2One   `xmlrpc:"country_id,omptempty"`
	CreateDate                      *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                       *Many2One   `xmlrpc:"create_uid,omptempty"`
	DateActionLast                  *Time       `xmlrpc:"date_action_last,omptempty"`
	DateClosed                      *Time       `xmlrpc:"date_closed,omptempty"`
	DateConversion                  *Time       `xmlrpc:"date_conversion,omptempty"`
	DateDeadline                    *Time       `xmlrpc:"date_deadline,omptempty"`
	DateLastStageUpdate             *Time       `xmlrpc:"date_last_stage_update,omptempty"`
	DateOpen                        *Time       `xmlrpc:"date_open,omptempty"`
	DayClose                        *Float      `xmlrpc:"day_close,omptempty"`
	DayOpen                         *Float      `xmlrpc:"day_open,omptempty"`
	DaysExceedingClosing            *Float      `xmlrpc:"days_exceeding_closing,omptempty"`
	DaysToConvert                   *Float      `xmlrpc:"days_to_convert,omptempty"`
	Description                     *String     `xmlrpc:"description,omptempty"`
	DisplayName                     *String     `xmlrpc:"display_name,omptempty"`
	DuplicateLeadCount              *Int        `xmlrpc:"duplicate_lead_count,omptempty"`
	DuplicateLeadIds                *Relation   `xmlrpc:"duplicate_lead_ids,omptempty"`
	EmailCc                         *String     `xmlrpc:"email_cc,omptempty"`
	EmailFrom                       *String     `xmlrpc:"email_from,omptempty"`
	EmailNormalized                 *String     `xmlrpc:"email_normalized,omptempty"`
	EmailState                      *Selection  `xmlrpc:"email_state,omptempty"`
	EventId                         *Many2One   `xmlrpc:"event_id,omptempty"`
	EventLeadRuleId                 *Many2One   `xmlrpc:"event_lead_rule_id,omptempty"`
	ExpectedRevenue                 *Float      `xmlrpc:"expected_revenue,omptempty"`
	FailedMessageIds                *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	Function                        *String     `xmlrpc:"function,omptempty"`
	HasMessage                      *Bool       `xmlrpc:"has_message,omptempty"`
	Id                              *Int        `xmlrpc:"id,omptempty"`
	IsAutomatedProbability          *Bool       `xmlrpc:"is_automated_probability,omptempty"`
	IsBlacklisted                   *Bool       `xmlrpc:"is_blacklisted,omptempty"`
	IsPartnerVisible                *Bool       `xmlrpc:"is_partner_visible,omptempty"`
	KanbanState                     *Selection  `xmlrpc:"kanban_state,omptempty"`
	LangActiveCount                 *Int        `xmlrpc:"lang_active_count,omptempty"`
	LangCode                        *String     `xmlrpc:"lang_code,omptempty"`
	LangId                          *Many2One   `xmlrpc:"lang_id,omptempty"`
	LeadProperties                  interface{} `xmlrpc:"lead_properties,omptempty"`
	LostReasonId                    *Many2One   `xmlrpc:"lost_reason_id,omptempty"`
	MediumId                        *Many2One   `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount          *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageBounce                   *Int        `xmlrpc:"message_bounce,omptempty"`
	MessageContent                  *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds              *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                 *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter          *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError              *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                      *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower               *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId         *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction               *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter        *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds               *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	Mobile                          *String     `xmlrpc:"mobile,omptempty"`
	MobileBlacklisted               *Bool       `xmlrpc:"mobile_blacklisted,omptempty"`
	MyActivityDateDeadline          *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                            *String     `xmlrpc:"name,omptempty"`
	OrderIds                        *Relation   `xmlrpc:"order_ids,omptempty"`
	PartnerEmailUpdate              *Bool       `xmlrpc:"partner_email_update,omptempty"`
	PartnerId                       *Many2One   `xmlrpc:"partner_id,omptempty"`
	PartnerIsBlacklisted            *Bool       `xmlrpc:"partner_is_blacklisted,omptempty"`
	PartnerName                     *String     `xmlrpc:"partner_name,omptempty"`
	PartnerPhoneUpdate              *Bool       `xmlrpc:"partner_phone_update,omptempty"`
	Phone                           *String     `xmlrpc:"phone,omptempty"`
	PhoneBlacklisted                *Bool       `xmlrpc:"phone_blacklisted,omptempty"`
	PhoneMobileSearch               *String     `xmlrpc:"phone_mobile_search,omptempty"`
	PhoneSanitized                  *String     `xmlrpc:"phone_sanitized,omptempty"`
	PhoneSanitizedBlacklisted       *Bool       `xmlrpc:"phone_sanitized_blacklisted,omptempty"`
	PhoneState                      *Selection  `xmlrpc:"phone_state,omptempty"`
	Priority                        *Selection  `xmlrpc:"priority,omptempty"`
	Probability                     *Float      `xmlrpc:"probability,omptempty"`
	ProratedRevenue                 *Float      `xmlrpc:"prorated_revenue,omptempty"`
	QuotationCount                  *Int        `xmlrpc:"quotation_count,omptempty"`
	RecurringPlan                   *Many2One   `xmlrpc:"recurring_plan,omptempty"`
	RecurringRevenue                *Float      `xmlrpc:"recurring_revenue,omptempty"`
	RecurringRevenueMonthly         *Float      `xmlrpc:"recurring_revenue_monthly,omptempty"`
	RecurringRevenueMonthlyProrated *Float      `xmlrpc:"recurring_revenue_monthly_prorated,omptempty"`
	Referred                        *String     `xmlrpc:"referred,omptempty"`
	RegistrationCount               *Int        `xmlrpc:"registration_count,omptempty"`
	RegistrationIds                 *Relation   `xmlrpc:"registration_ids,omptempty"`
	RevealId                        *String     `xmlrpc:"reveal_id,omptempty"`
	SaleAmountTotal                 *Float      `xmlrpc:"sale_amount_total,omptempty"`
	SaleOrderCount                  *Int        `xmlrpc:"sale_order_count,omptempty"`
	SourceId                        *Many2One   `xmlrpc:"source_id,omptempty"`
	StageId                         *Many2One   `xmlrpc:"stage_id,omptempty"`
	StateId                         *Many2One   `xmlrpc:"state_id,omptempty"`
	Street                          *String     `xmlrpc:"street,omptempty"`
	Street2                         *String     `xmlrpc:"street2,omptempty"`
	TagIds                          *Relation   `xmlrpc:"tag_ids,omptempty"`
	TeamId                          *Many2One   `xmlrpc:"team_id,omptempty"`
	Title                           *Many2One   `xmlrpc:"title,omptempty"`
	Type                            *Selection  `xmlrpc:"type,omptempty"`
	UserCompanyIds                  *Relation   `xmlrpc:"user_company_ids,omptempty"`
	UserId                          *Many2One   `xmlrpc:"user_id,omptempty"`
	VisitorIds                      *Relation   `xmlrpc:"visitor_ids,omptempty"`
	VisitorPageCount                *Int        `xmlrpc:"visitor_page_count,omptempty"`
	VisitorSessionsCount            *Int        `xmlrpc:"visitor_sessions_count,omptempty"`
	Website                         *String     `xmlrpc:"website,omptempty"`
	WebsiteMessageIds               *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WonStatus                       *Selection  `xmlrpc:"won_status,omptempty"`
	WriteDate                       *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                        *Many2One   `xmlrpc:"write_uid,omptempty"`
	Zip                             *String     `xmlrpc:"zip,omptempty"`
}

// CrmLeads represents array of crm.lead model.
type CrmLeads []CrmLead

// CrmLeadModel is the odoo model name.
const CrmLeadModel = "crm.lead"

// Many2One convert CrmLead to *Many2One.
func (cl *CrmLead) Many2One() *Many2One {
	return NewMany2One(cl.Id.Get(), "")
}

// CreateCrmLead creates a new crm.lead model and returns its id.
func (c *Client) CreateCrmLead(cl *CrmLead) (int64, error) {
	ids, err := c.CreateCrmLeads([]*CrmLead{cl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLead creates a new crm.lead model and returns its id.
func (c *Client) CreateCrmLeads(cls []*CrmLead) ([]int64, error) {
	var vv []interface{}
	for _, v := range cls {
		vv = append(vv, v)
	}
	return c.Create(CrmLeadModel, vv)
}

// UpdateCrmLead updates an existing crm.lead record.
func (c *Client) UpdateCrmLead(cl *CrmLead) error {
	return c.UpdateCrmLeads([]int64{cl.Id.Get()}, cl)
}

// UpdateCrmLeads updates existing crm.lead records.
// All records (represented by ids) will be updated by cl values.
func (c *Client) UpdateCrmLeads(ids []int64, cl *CrmLead) error {
	return c.Update(CrmLeadModel, ids, cl)
}

// DeleteCrmLead deletes an existing crm.lead record.
func (c *Client) DeleteCrmLead(id int64) error {
	return c.DeleteCrmLeads([]int64{id})
}

// DeleteCrmLeads deletes existing crm.lead records.
func (c *Client) DeleteCrmLeads(ids []int64) error {
	return c.Delete(CrmLeadModel, ids)
}

// GetCrmLead gets crm.lead existing record.
func (c *Client) GetCrmLead(id int64) (*CrmLead, error) {
	cls, err := c.GetCrmLeads([]int64{id})
	if err != nil {
		return nil, err
	}
	if cls != nil && len(*cls) > 0 {
		return &((*cls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.lead not found", id)
}

// GetCrmLeads gets crm.lead existing records.
func (c *Client) GetCrmLeads(ids []int64) (*CrmLeads, error) {
	cls := &CrmLeads{}
	if err := c.Read(CrmLeadModel, ids, nil, cls); err != nil {
		return nil, err
	}
	return cls, nil
}

// FindCrmLead finds crm.lead record by querying it with criteria.
func (c *Client) FindCrmLead(criteria *Criteria) (*CrmLead, error) {
	cls := &CrmLeads{}
	if err := c.SearchRead(CrmLeadModel, criteria, NewOptions().Limit(1), cls); err != nil {
		return nil, err
	}
	if cls != nil && len(*cls) > 0 {
		return &((*cls)[0]), nil
	}
	return nil, fmt.Errorf("crm.lead was not found with criteria %v", criteria)
}

// FindCrmLeads finds crm.lead records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeads(criteria *Criteria, options *Options) (*CrmLeads, error) {
	cls := &CrmLeads{}
	if err := c.SearchRead(CrmLeadModel, criteria, options, cls); err != nil {
		return nil, err
	}
	return cls, nil
}

// FindCrmLeadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmLeadModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmLeadId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.lead was not found with criteria %v and options %v", criteria, options)
}
