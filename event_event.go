package odoo

import (
	"fmt"
)

// EventEvent represents event.event model.
type EventEvent struct {
	LastUpdate                      *Time      `xmlrpc:"__last_update,omptempty"`
	Active                          *Bool      `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId         *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline            *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration     *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon           *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                     *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                   *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                 *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                  *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                  *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AddressId                       *Many2One  `xmlrpc:"address_id,omptempty"`
	AddressInline                   *String    `xmlrpc:"address_inline,omptempty"`
	AddressSearch                   *Many2One  `xmlrpc:"address_search,omptempty"`
	AutoConfirm                     *Bool      `xmlrpc:"auto_confirm,omptempty"`
	CanPublish                      *Bool      `xmlrpc:"can_publish,omptempty"`
	CommunityMenu                   *Bool      `xmlrpc:"community_menu,omptempty"`
	CommunityMenuIds                *Relation  `xmlrpc:"community_menu_ids,omptempty"`
	CompanyId                       *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId                       *Many2One  `xmlrpc:"country_id,omptempty"`
	CoverProperties                 *String    `xmlrpc:"cover_properties,omptempty"`
	CreateDate                      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateBegin                       *Time      `xmlrpc:"date_begin,omptempty"`
	DateBeginLocated                *String    `xmlrpc:"date_begin_located,omptempty"`
	DateEnd                         *Time      `xmlrpc:"date_end,omptempty"`
	DateEndLocated                  *String    `xmlrpc:"date_end_located,omptempty"`
	DateTz                          *Selection `xmlrpc:"date_tz,omptempty"`
	Description                     *String    `xmlrpc:"description,omptempty"`
	DisplayName                     *String    `xmlrpc:"display_name,omptempty"`
	EventMailIds                    *Relation  `xmlrpc:"event_mail_ids,omptempty"`
	EventRegistrationsOpen          *Bool      `xmlrpc:"event_registrations_open,omptempty"`
	EventRegistrationsSoldOut       *Bool      `xmlrpc:"event_registrations_sold_out,omptempty"`
	EventRegistrationsStarted       *Bool      `xmlrpc:"event_registrations_started,omptempty"`
	EventTicketIds                  *Relation  `xmlrpc:"event_ticket_ids,omptempty"`
	EventTypeId                     *Many2One  `xmlrpc:"event_type_id,omptempty"`
	FailedMessageIds                *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FirebaseEnablePushNotifications *Bool      `xmlrpc:"firebase_enable_push_notifications,omptempty"`
	HasMessage                      *Bool      `xmlrpc:"has_message,omptempty"`
	Id                              *Int       `xmlrpc:"id,omptempty"`
	IntroductionMenu                *Bool      `xmlrpc:"introduction_menu,omptempty"`
	IntroductionMenuIds             *Relation  `xmlrpc:"introduction_menu_ids,omptempty"`
	IsDone                          *Bool      `xmlrpc:"is_done,omptempty"`
	IsFinished                      *Bool      `xmlrpc:"is_finished,omptempty"`
	IsOneDay                        *Bool      `xmlrpc:"is_one_day,omptempty"`
	IsOngoing                       *Bool      `xmlrpc:"is_ongoing,omptempty"`
	IsParticipating                 *Bool      `xmlrpc:"is_participating,omptempty"`
	IsPublished                     *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized                  *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	KanbanState                     *Selection `xmlrpc:"kanban_state,omptempty"`
	KanbanStateLabel                *String    `xmlrpc:"kanban_state_label,omptempty"`
	LeadCount                       *Int       `xmlrpc:"lead_count,omptempty"`
	LeadIds                         *Relation  `xmlrpc:"lead_ids,omptempty"`
	LegendBlocked                   *String    `xmlrpc:"legend_blocked,omptempty"`
	LegendDone                      *String    `xmlrpc:"legend_done,omptempty"`
	LegendNormal                    *String    `xmlrpc:"legend_normal,omptempty"`
	LocationMenu                    *Bool      `xmlrpc:"location_menu,omptempty"`
	LocationMenuIds                 *Relation  `xmlrpc:"location_menu_ids,omptempty"`
	MenuId                          *Many2One  `xmlrpc:"menu_id,omptempty"`
	MenuRegisterCta                 *Bool      `xmlrpc:"menu_register_cta,omptempty"`
	MessageAttachmentCount          *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                  *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds              *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                 *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter          *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError              *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                      *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower               *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId         *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction               *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter        *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds               *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline          *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                            *String    `xmlrpc:"name,omptempty"`
	Note                            *String    `xmlrpc:"note,omptempty"`
	OrganizerId                     *Many2One  `xmlrpc:"organizer_id,omptempty"`
	RegisterMenu                    *Bool      `xmlrpc:"register_menu,omptempty"`
	RegisterMenuIds                 *Relation  `xmlrpc:"register_menu_ids,omptempty"`
	RegistrationIds                 *Relation  `xmlrpc:"registration_ids,omptempty"`
	SeatsAvailable                  *Int       `xmlrpc:"seats_available,omptempty"`
	SeatsExpected                   *Int       `xmlrpc:"seats_expected,omptempty"`
	SeatsLimited                    *Bool      `xmlrpc:"seats_limited,omptempty"`
	SeatsMax                        *Int       `xmlrpc:"seats_max,omptempty"`
	SeatsReserved                   *Int       `xmlrpc:"seats_reserved,omptempty"`
	SeatsUnconfirmed                *Int       `xmlrpc:"seats_unconfirmed,omptempty"`
	SeatsUsed                       *Int       `xmlrpc:"seats_used,omptempty"`
	SeoName                         *String    `xmlrpc:"seo_name,omptempty"`
	StageId                         *Many2One  `xmlrpc:"stage_id,omptempty"`
	StartRemaining                  *Int       `xmlrpc:"start_remaining,omptempty"`
	StartSaleDatetime               *Time      `xmlrpc:"start_sale_datetime,omptempty"`
	StartToday                      *Bool      `xmlrpc:"start_today,omptempty"`
	Subtitle                        *String    `xmlrpc:"subtitle,omptempty"`
	TagIds                          *Relation  `xmlrpc:"tag_ids,omptempty"`
	TicketInstructions              *String    `xmlrpc:"ticket_instructions,omptempty"`
	UserId                          *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteId                       *Many2One  `xmlrpc:"website_id,omptempty"`
	WebsiteMenu                     *Bool      `xmlrpc:"website_menu,omptempty"`
	WebsiteMessageIds               *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription          *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords             *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg                *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle                *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished                *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                      *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventEvents represents array of event.event model.
type EventEvents []EventEvent

// EventEventModel is the odoo model name.
const EventEventModel = "event.event"

// Many2One convert EventEvent to *Many2One.
func (ee *EventEvent) Many2One() *Many2One {
	return NewMany2One(ee.Id.Get(), "")
}

// CreateEventEvent creates a new event.event model and returns its id.
func (c *Client) CreateEventEvent(ee *EventEvent) (int64, error) {
	ids, err := c.CreateEventEvents([]*EventEvent{ee})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventEvent creates a new event.event model and returns its id.
func (c *Client) CreateEventEvents(ees []*EventEvent) ([]int64, error) {
	var vv []interface{}
	for _, v := range ees {
		vv = append(vv, v)
	}
	return c.Create(EventEventModel, vv)
}

// UpdateEventEvent updates an existing event.event record.
func (c *Client) UpdateEventEvent(ee *EventEvent) error {
	return c.UpdateEventEvents([]int64{ee.Id.Get()}, ee)
}

// UpdateEventEvents updates existing event.event records.
// All records (represented by ids) will be updated by ee values.
func (c *Client) UpdateEventEvents(ids []int64, ee *EventEvent) error {
	return c.Update(EventEventModel, ids, ee)
}

// DeleteEventEvent deletes an existing event.event record.
func (c *Client) DeleteEventEvent(id int64) error {
	return c.DeleteEventEvents([]int64{id})
}

// DeleteEventEvents deletes existing event.event records.
func (c *Client) DeleteEventEvents(ids []int64) error {
	return c.Delete(EventEventModel, ids)
}

// GetEventEvent gets event.event existing record.
func (c *Client) GetEventEvent(id int64) (*EventEvent, error) {
	ees, err := c.GetEventEvents([]int64{id})
	if err != nil {
		return nil, err
	}
	if ees != nil && len(*ees) > 0 {
		return &((*ees)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.event not found", id)
}

// GetEventEvents gets event.event existing records.
func (c *Client) GetEventEvents(ids []int64) (*EventEvents, error) {
	ees := &EventEvents{}
	if err := c.Read(EventEventModel, ids, nil, ees); err != nil {
		return nil, err
	}
	return ees, nil
}

// FindEventEvent finds event.event record by querying it with criteria.
func (c *Client) FindEventEvent(criteria *Criteria) (*EventEvent, error) {
	ees := &EventEvents{}
	if err := c.SearchRead(EventEventModel, criteria, NewOptions().Limit(1), ees); err != nil {
		return nil, err
	}
	if ees != nil && len(*ees) > 0 {
		return &((*ees)[0]), nil
	}
	return nil, fmt.Errorf("event.event was not found with criteria %v", criteria)
}

// FindEventEvents finds event.event records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEvents(criteria *Criteria, options *Options) (*EventEvents, error) {
	ees := &EventEvents{}
	if err := c.SearchRead(EventEventModel, criteria, options, ees); err != nil {
		return nil, err
	}
	return ees, nil
}

// FindEventEventIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventEventModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventEventId finds record id by querying it with criteria.
func (c *Client) FindEventEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.event was not found with criteria %v and options %v", criteria, options)
}
