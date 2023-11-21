package odoo

import (
	"fmt"
)

// HelpdeskTicket represents helpdesk.ticket model.
type HelpdeskTicket struct {
	LastUpdate                          *Time       `xmlrpc:"__last_update,omptempty"`
	AccessToken                         *String     `xmlrpc:"access_token,omptempty"`
	AccessUrl                           *String     `xmlrpc:"access_url,omptempty"`
	AccessWarning                       *String     `xmlrpc:"access_warning,omptempty"`
	Active                              *Bool       `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId             *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline                *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration         *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon               *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                         *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState                       *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                     *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                    *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                      *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                      *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	AddUrl                              *String     `xmlrpc:"add_url,omptempty"`
	AnalyticAccountId                   *Many2One   `xmlrpc:"analytic_account_id,omptempty"`
	AnsweredCustomerMessageCount        *Int        `xmlrpc:"answered_customer_message_count,omptempty"`
	AssignDate                          *Time       `xmlrpc:"assign_date,omptempty"`
	AssignHours                         *Int        `xmlrpc:"assign_hours,omptempty"`
	AvgResponseHours                    *Float      `xmlrpc:"avg_response_hours,omptempty"`
	CampaignId                          *Many2One   `xmlrpc:"campaign_id,omptempty"`
	CloseDate                           *Time       `xmlrpc:"close_date,omptempty"`
	CloseHours                          *Int        `xmlrpc:"close_hours,omptempty"`
	ClosedByPartner                     *Bool       `xmlrpc:"closed_by_partner,omptempty"`
	Color                               *Int        `xmlrpc:"color,omptempty"`
	CommercialPartnerId                 *Many2One   `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                           *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate                          *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                           *Many2One   `xmlrpc:"create_uid,omptempty"`
	DataFetcherIds                      *Relation   `xmlrpc:"data_fetcher_ids,omptempty"`
	DateLastStageUpdate                 *Time       `xmlrpc:"date_last_stage_update,omptempty"`
	Description                         *String     `xmlrpc:"description,omptempty"`
	DisplayName                         *String     `xmlrpc:"display_name,omptempty"`
	DisplayTimer                        *Bool       `xmlrpc:"display_timer,omptempty"`
	DisplayTimerPause                   *Bool       `xmlrpc:"display_timer_pause,omptempty"`
	DisplayTimerResume                  *Bool       `xmlrpc:"display_timer_resume,omptempty"`
	DisplayTimerStartPrimary            *Bool       `xmlrpc:"display_timer_start_primary,omptempty"`
	DisplayTimerStartSecondary          *Bool       `xmlrpc:"display_timer_start_secondary,omptempty"`
	DisplayTimerStop                    *Bool       `xmlrpc:"display_timer_stop,omptempty"`
	DisplayTimesheetTimer               *Bool       `xmlrpc:"display_timesheet_timer,omptempty"`
	DomainUserIds                       *Relation   `xmlrpc:"domain_user_ids,omptempty"`
	Email                               *String     `xmlrpc:"email,omptempty"`
	EmailCc                             *String     `xmlrpc:"email_cc,omptempty"`
	EncodeUomInDays                     *Bool       `xmlrpc:"encode_uom_in_days,omptempty"`
	FailedMessageIds                    *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	FirstResponseHours                  *Float      `xmlrpc:"first_response_hours,omptempty"`
	HasMessage                          *Bool       `xmlrpc:"has_message,omptempty"`
	Id                                  *Int        `xmlrpc:"id,omptempty"`
	IsPartnerEmailUpdate                *Bool       `xmlrpc:"is_partner_email_update,omptempty"`
	IsPartnerPhoneUpdate                *Bool       `xmlrpc:"is_partner_phone_update,omptempty"`
	IsTimerRunning                      *Bool       `xmlrpc:"is_timer_running,omptempty"`
	KanbanState                         *Selection  `xmlrpc:"kanban_state,omptempty"`
	KanbanStateLabel                    *String     `xmlrpc:"kanban_state_label,omptempty"`
	LegendBlocked                       *String     `xmlrpc:"legend_blocked,omptempty"`
	LegendDone                          *String     `xmlrpc:"legend_done,omptempty"`
	LegendNormal                        *String     `xmlrpc:"legend_normal,omptempty"`
	MediumId                            *Many2One   `xmlrpc:"medium_id,omptempty"`
	MessageAttachmentCount              *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                      *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                  *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                     *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter              *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                  *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                          *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                   *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId             *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                   *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter            *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                   *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline              *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                *String     `xmlrpc:"name,omptempty"`
	OldestUnansweredCustomerMessageDate *Time       `xmlrpc:"oldest_unanswered_customer_message_date,omptempty"`
	OpenHours                           *Int        `xmlrpc:"open_hours,omptempty"`
	PartnerEmail                        *String     `xmlrpc:"partner_email,omptempty"`
	PartnerId                           *Many2One   `xmlrpc:"partner_id,omptempty"`
	PartnerName                         *String     `xmlrpc:"partner_name,omptempty"`
	PartnerOpenTicketCount              *Int        `xmlrpc:"partner_open_ticket_count,omptempty"`
	PartnerPhone                        *String     `xmlrpc:"partner_phone,omptempty"`
	PartnerTicketCount                  *Int        `xmlrpc:"partner_ticket_count,omptempty"`
	PartnerTicketIds                    *Relation   `xmlrpc:"partner_ticket_ids,omptempty"`
	Priority                            *Selection  `xmlrpc:"priority,omptempty"`
	ProjectId                           *Many2One   `xmlrpc:"project_id,omptempty"`
	ProjectSaleOrderId                  *Many2One   `xmlrpc:"project_sale_order_id,omptempty"`
	Properties                          interface{} `xmlrpc:"properties,omptempty"`
	RatingAvg                           *Float      `xmlrpc:"rating_avg,omptempty"`
	RatingAvgText                       *Selection  `xmlrpc:"rating_avg_text,omptempty"`
	RatingCount                         *Int        `xmlrpc:"rating_count,omptempty"`
	RatingIds                           *Relation   `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback                  *String     `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage                     *String     `xmlrpc:"rating_last_image,omptempty"`
	RatingLastText                      *Selection  `xmlrpc:"rating_last_text,omptempty"`
	RatingLastValue                     *Float      `xmlrpc:"rating_last_value,omptempty"`
	RatingPercentageSatisfaction        *Float      `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RemainingHoursAvailable             *Bool       `xmlrpc:"remaining_hours_available,omptempty"`
	RemainingHoursSo                    *Float      `xmlrpc:"remaining_hours_so,omptempty"`
	RemoteError                         *Bool       `xmlrpc:"remote_error,omptempty"`
	SaleLineId                          *Many2One   `xmlrpc:"sale_line_id,omptempty"`
	SaleOrderId                         *Many2One   `xmlrpc:"sale_order_id,omptempty"`
	SlaDeadline                         *Time       `xmlrpc:"sla_deadline,omptempty"`
	SlaDeadlineHours                    *Float      `xmlrpc:"sla_deadline_hours,omptempty"`
	SlaFail                             *Bool       `xmlrpc:"sla_fail,omptempty"`
	SlaIds                              *Relation   `xmlrpc:"sla_ids,omptempty"`
	SlaReached                          *Bool       `xmlrpc:"sla_reached,omptempty"`
	SlaReachedLate                      *Bool       `xmlrpc:"sla_reached_late,omptempty"`
	SlaStatusIds                        *Relation   `xmlrpc:"sla_status_ids,omptempty"`
	SlaSuccess                          *Bool       `xmlrpc:"sla_success,omptempty"`
	SourceId                            *Many2One   `xmlrpc:"source_id,omptempty"`
	StageId                             *Many2One   `xmlrpc:"stage_id,omptempty"`
	StatusChangedDate                   *Time       `xmlrpc:"status_changed_date,omptempty"`
	Statuses                            *String     `xmlrpc:"statuses,omptempty"`
	TagIds                              *Relation   `xmlrpc:"tag_ids,omptempty"`
	TeamId                              *Many2One   `xmlrpc:"team_id,omptempty"`
	TeamPrivacyVisibility               *Selection  `xmlrpc:"team_privacy_visibility,omptempty"`
	TicketRef                           *String     `xmlrpc:"ticket_ref,omptempty"`
	TicketTypeId                        *Many2One   `xmlrpc:"ticket_type_id,omptempty"`
	TimerPause                          *Time       `xmlrpc:"timer_pause,omptempty"`
	TimerStart                          *Time       `xmlrpc:"timer_start,omptempty"`
	TimesheetIds                        *Relation   `xmlrpc:"timesheet_ids,omptempty"`
	TotalHoursSpent                     *Float      `xmlrpc:"total_hours_spent,omptempty"`
	TotalResponseHours                  *Float      `xmlrpc:"total_response_hours,omptempty"`
	UseCoupons                          *Bool       `xmlrpc:"use_coupons,omptempty"`
	UseCreditNotes                      *Bool       `xmlrpc:"use_credit_notes,omptempty"`
	UseHelpdeskSaleTimesheet            *Bool       `xmlrpc:"use_helpdesk_sale_timesheet,omptempty"`
	UseHelpdeskTimesheet                *Bool       `xmlrpc:"use_helpdesk_timesheet,omptempty"`
	UseProductRepairs                   *Bool       `xmlrpc:"use_product_repairs,omptempty"`
	UseProductReturns                   *Bool       `xmlrpc:"use_product_returns,omptempty"`
	UseRating                           *Bool       `xmlrpc:"use_rating,omptempty"`
	UseSla                              *Bool       `xmlrpc:"use_sla,omptempty"`
	UserId                              *Many2One   `xmlrpc:"user_id,omptempty"`
	UserTimerId                         *Relation   `xmlrpc:"user_timer_id,omptempty"`
	WebsiteMessageIds                   *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                           *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                            *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskTickets represents array of helpdesk.ticket model.
type HelpdeskTickets []HelpdeskTicket

// HelpdeskTicketModel is the odoo model name.
const HelpdeskTicketModel = "helpdesk.ticket"

// Many2One convert HelpdeskTicket to *Many2One.
func (ht *HelpdeskTicket) Many2One() *Many2One {
	return NewMany2One(ht.Id.Get(), "")
}

// CreateHelpdeskTicket creates a new helpdesk.ticket model and returns its id.
func (c *Client) CreateHelpdeskTicket(ht *HelpdeskTicket) (int64, error) {
	ids, err := c.CreateHelpdeskTickets([]*HelpdeskTicket{ht})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTicket creates a new helpdesk.ticket model and returns its id.
func (c *Client) CreateHelpdeskTickets(hts []*HelpdeskTicket) ([]int64, error) {
	var vv []interface{}
	for _, v := range hts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTicketModel, vv)
}

// UpdateHelpdeskTicket updates an existing helpdesk.ticket record.
func (c *Client) UpdateHelpdeskTicket(ht *HelpdeskTicket) error {
	return c.UpdateHelpdeskTickets([]int64{ht.Id.Get()}, ht)
}

// UpdateHelpdeskTickets updates existing helpdesk.ticket records.
// All records (represented by ids) will be updated by ht values.
func (c *Client) UpdateHelpdeskTickets(ids []int64, ht *HelpdeskTicket) error {
	return c.Update(HelpdeskTicketModel, ids, ht)
}

// DeleteHelpdeskTicket deletes an existing helpdesk.ticket record.
func (c *Client) DeleteHelpdeskTicket(id int64) error {
	return c.DeleteHelpdeskTickets([]int64{id})
}

// DeleteHelpdeskTickets deletes existing helpdesk.ticket records.
func (c *Client) DeleteHelpdeskTickets(ids []int64) error {
	return c.Delete(HelpdeskTicketModel, ids)
}

// GetHelpdeskTicket gets helpdesk.ticket existing record.
func (c *Client) GetHelpdeskTicket(id int64) (*HelpdeskTicket, error) {
	hts, err := c.GetHelpdeskTickets([]int64{id})
	if err != nil {
		return nil, err
	}
	if hts != nil && len(*hts) > 0 {
		return &((*hts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.ticket not found", id)
}

// GetHelpdeskTickets gets helpdesk.ticket existing records.
func (c *Client) GetHelpdeskTickets(ids []int64) (*HelpdeskTickets, error) {
	hts := &HelpdeskTickets{}
	if err := c.Read(HelpdeskTicketModel, ids, nil, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTicket finds helpdesk.ticket record by querying it with criteria.
func (c *Client) FindHelpdeskTicket(criteria *Criteria) (*HelpdeskTicket, error) {
	hts := &HelpdeskTickets{}
	if err := c.SearchRead(HelpdeskTicketModel, criteria, NewOptions().Limit(1), hts); err != nil {
		return nil, err
	}
	if hts != nil && len(*hts) > 0 {
		return &((*hts)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.ticket was not found with criteria %v", criteria)
}

// FindHelpdeskTickets finds helpdesk.ticket records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTickets(criteria *Criteria, options *Options) (*HelpdeskTickets, error) {
	hts := &HelpdeskTickets{}
	if err := c.SearchRead(HelpdeskTicketModel, criteria, options, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTicketIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTicketIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTicketModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTicketId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTicketId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTicketModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.ticket was not found with criteria %v and options %v", criteria, options)
}
