package odoo

import (
	"fmt"
)

// CalendarEvent represents calendar.event model.
type CalendarEvent struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken               *String    `xmlrpc:"access_token,omptempty"`
	Active                    *Bool      `xmlrpc:"active,omptempty"`
	ActivityIds               *Relation  `xmlrpc:"activity_ids,omptempty"`
	AlarmIds                  *Relation  `xmlrpc:"alarm_ids,omptempty"`
	Allday                    *Bool      `xmlrpc:"allday,omptempty"`
	ApplicantId               *Many2One  `xmlrpc:"applicant_id,omptempty"`
	AppointmentAnswerInputIds *Relation  `xmlrpc:"appointment_answer_input_ids,omptempty"`
	AppointmentInviteId       *Many2One  `xmlrpc:"appointment_invite_id,omptempty"`
	AppointmentTypeId         *Many2One  `xmlrpc:"appointment_type_id,omptempty"`
	AttendeeIds               *Relation  `xmlrpc:"attendee_ids,omptempty"`
	AttendeeStatus            *Selection `xmlrpc:"attendee_status,omptempty"`
	Byday                     *Selection `xmlrpc:"byday,omptempty"`
	CategIds                  *Relation  `xmlrpc:"categ_ids,omptempty"`
	Count                     *Int       `xmlrpc:"count,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	Day                       *Int       `xmlrpc:"day,omptempty"`
	Description               *String    `xmlrpc:"description,omptempty"`
	DisplayDescription        *Bool      `xmlrpc:"display_description,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	DisplayTime               *String    `xmlrpc:"display_time,omptempty"`
	Duration                  *Float     `xmlrpc:"duration,omptempty"`
	EndType                   *Selection `xmlrpc:"end_type,omptempty"`
	EventTz                   *Selection `xmlrpc:"event_tz,omptempty"`
	FailedMessageIds          *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FollowRecurrence          *Bool      `xmlrpc:"follow_recurrence,omptempty"`
	Fri                       *Bool      `xmlrpc:"fri,omptempty"`
	HasMessage                *Bool      `xmlrpc:"has_message,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	Interval                  *Int       `xmlrpc:"interval,omptempty"`
	InvalidEmailPartnerIds    *Relation  `xmlrpc:"invalid_email_partner_ids,omptempty"`
	IsHighlighted             *Bool      `xmlrpc:"is_highlighted,omptempty"`
	IsOrganizerAlone          *Bool      `xmlrpc:"is_organizer_alone,omptempty"`
	Location                  *String    `xmlrpc:"location,omptempty"`
	MessageAttachmentCount    *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent            *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds        *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError           *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter    *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError        *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower         *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId   *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction         *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter  *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds         *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Mon                       *Bool      `xmlrpc:"mon,omptempty"`
	MonthBy                   *Selection `xmlrpc:"month_by,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	OpportunityId             *Many2One  `xmlrpc:"opportunity_id,omptempty"`
	PartnerId                 *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerIds                *Relation  `xmlrpc:"partner_ids,omptempty"`
	Privacy                   *Selection `xmlrpc:"privacy,omptempty"`
	RecurrenceId              *Many2One  `xmlrpc:"recurrence_id,omptempty"`
	RecurrenceUpdate          *Selection `xmlrpc:"recurrence_update,omptempty"`
	Recurrency                *Bool      `xmlrpc:"recurrency,omptempty"`
	ResId                     *Many2One  `xmlrpc:"res_id,omptempty"`
	ResModel                  *String    `xmlrpc:"res_model,omptempty"`
	ResModelId                *Many2One  `xmlrpc:"res_model_id,omptempty"`
	Rrule                     *String    `xmlrpc:"rrule,omptempty"`
	RruleType                 *Selection `xmlrpc:"rrule_type,omptempty"`
	Sat                       *Bool      `xmlrpc:"sat,omptempty"`
	ShowAs                    *Selection `xmlrpc:"show_as,omptempty"`
	Start                     *Time      `xmlrpc:"start,omptempty"`
	StartDate                 *Time      `xmlrpc:"start_date,omptempty"`
	Stop                      *Time      `xmlrpc:"stop,omptempty"`
	StopDate                  *Time      `xmlrpc:"stop_date,omptempty"`
	Sun                       *Bool      `xmlrpc:"sun,omptempty"`
	Thu                       *Bool      `xmlrpc:"thu,omptempty"`
	Tue                       *Bool      `xmlrpc:"tue,omptempty"`
	Until                     *Time      `xmlrpc:"until,omptempty"`
	UserId                    *Many2One  `xmlrpc:"user_id,omptempty"`
	VideocallChannelId        *Many2One  `xmlrpc:"videocall_channel_id,omptempty"`
	VideocallLocation         *String    `xmlrpc:"videocall_location,omptempty"`
	VideocallSource           *Selection `xmlrpc:"videocall_source,omptempty"`
	WebsiteMessageIds         *Relation  `xmlrpc:"website_message_ids,omptempty"`
	Wed                       *Bool      `xmlrpc:"wed,omptempty"`
	Weekday                   *Selection `xmlrpc:"weekday,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CalendarEvents represents array of calendar.event model.
type CalendarEvents []CalendarEvent

// CalendarEventModel is the odoo model name.
const CalendarEventModel = "calendar.event"

// Many2One convert CalendarEvent to *Many2One.
func (ce *CalendarEvent) Many2One() *Many2One {
	return NewMany2One(ce.Id.Get(), "")
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvent(ce *CalendarEvent) (int64, error) {
	ids, err := c.CreateCalendarEvents([]*CalendarEvent{ce})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarEvent creates a new calendar.event model and returns its id.
func (c *Client) CreateCalendarEvents(ces []*CalendarEvent) ([]int64, error) {
	var vv []interface{}
	for _, v := range ces {
		vv = append(vv, v)
	}
	return c.Create(CalendarEventModel, vv)
}

// UpdateCalendarEvent updates an existing calendar.event record.
func (c *Client) UpdateCalendarEvent(ce *CalendarEvent) error {
	return c.UpdateCalendarEvents([]int64{ce.Id.Get()}, ce)
}

// UpdateCalendarEvents updates existing calendar.event records.
// All records (represented by ids) will be updated by ce values.
func (c *Client) UpdateCalendarEvents(ids []int64, ce *CalendarEvent) error {
	return c.Update(CalendarEventModel, ids, ce)
}

// DeleteCalendarEvent deletes an existing calendar.event record.
func (c *Client) DeleteCalendarEvent(id int64) error {
	return c.DeleteCalendarEvents([]int64{id})
}

// DeleteCalendarEvents deletes existing calendar.event records.
func (c *Client) DeleteCalendarEvents(ids []int64) error {
	return c.Delete(CalendarEventModel, ids)
}

// GetCalendarEvent gets calendar.event existing record.
func (c *Client) GetCalendarEvent(id int64) (*CalendarEvent, error) {
	ces, err := c.GetCalendarEvents([]int64{id})
	if err != nil {
		return nil, err
	}
	if ces != nil && len(*ces) > 0 {
		return &((*ces)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.event not found", id)
}

// GetCalendarEvents gets calendar.event existing records.
func (c *Client) GetCalendarEvents(ids []int64) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.Read(CalendarEventModel, ids, nil, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEvent finds calendar.event record by querying it with criteria.
func (c *Client) FindCalendarEvent(criteria *Criteria) (*CalendarEvent, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, NewOptions().Limit(1), ces); err != nil {
		return nil, err
	}
	if ces != nil && len(*ces) > 0 {
		return &((*ces)[0]), nil
	}
	return nil, fmt.Errorf("calendar.event was not found with criteria %v", criteria)
}

// FindCalendarEvents finds calendar.event records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEvents(criteria *Criteria, options *Options) (*CalendarEvents, error) {
	ces := &CalendarEvents{}
	if err := c.SearchRead(CalendarEventModel, criteria, options, ces); err != nil {
		return nil, err
	}
	return ces, nil
}

// FindCalendarEventIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarEventIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarEventModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarEventId finds record id by querying it with criteria.
func (c *Client) FindCalendarEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("calendar.event was not found with criteria %v and options %v", criteria, options)
}
