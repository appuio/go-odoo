package odoo

import (
	"fmt"
)

// EventRegistration represents event.registration model.
type EventRegistration struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateClosed                  *Time      `xmlrpc:"date_closed,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Email                       *String    `xmlrpc:"email,omptempty"`
	EventBeginDate              *Time      `xmlrpc:"event_begin_date,omptempty"`
	EventEndDate                *Time      `xmlrpc:"event_end_date,omptempty"`
	EventId                     *Many2One  `xmlrpc:"event_id,omptempty"`
	EventOrganizerId            *Many2One  `xmlrpc:"event_organizer_id,omptempty"`
	EventTicketId               *Many2One  `xmlrpc:"event_ticket_id,omptempty"`
	EventUserId                 *Many2One  `xmlrpc:"event_user_id,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	LeadCount                   *Int       `xmlrpc:"lead_count,omptempty"`
	LeadIds                     *Relation  `xmlrpc:"lead_ids,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent              *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Mobile                      *String    `xmlrpc:"mobile,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	Phone                       *String    `xmlrpc:"phone,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	UtmCampaignId               *Many2One  `xmlrpc:"utm_campaign_id,omptempty"`
	UtmMediumId                 *Many2One  `xmlrpc:"utm_medium_id,omptempty"`
	UtmSourceId                 *Many2One  `xmlrpc:"utm_source_id,omptempty"`
	VisitorId                   *Many2One  `xmlrpc:"visitor_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// EventRegistrations represents array of event.registration model.
type EventRegistrations []EventRegistration

// EventRegistrationModel is the odoo model name.
const EventRegistrationModel = "event.registration"

// Many2One convert EventRegistration to *Many2One.
func (er *EventRegistration) Many2One() *Many2One {
	return NewMany2One(er.Id.Get(), "")
}

// CreateEventRegistration creates a new event.registration model and returns its id.
func (c *Client) CreateEventRegistration(er *EventRegistration) (int64, error) {
	ids, err := c.CreateEventRegistrations([]*EventRegistration{er})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventRegistration creates a new event.registration model and returns its id.
func (c *Client) CreateEventRegistrations(ers []*EventRegistration) ([]int64, error) {
	var vv []interface{}
	for _, v := range ers {
		vv = append(vv, v)
	}
	return c.Create(EventRegistrationModel, vv)
}

// UpdateEventRegistration updates an existing event.registration record.
func (c *Client) UpdateEventRegistration(er *EventRegistration) error {
	return c.UpdateEventRegistrations([]int64{er.Id.Get()}, er)
}

// UpdateEventRegistrations updates existing event.registration records.
// All records (represented by ids) will be updated by er values.
func (c *Client) UpdateEventRegistrations(ids []int64, er *EventRegistration) error {
	return c.Update(EventRegistrationModel, ids, er)
}

// DeleteEventRegistration deletes an existing event.registration record.
func (c *Client) DeleteEventRegistration(id int64) error {
	return c.DeleteEventRegistrations([]int64{id})
}

// DeleteEventRegistrations deletes existing event.registration records.
func (c *Client) DeleteEventRegistrations(ids []int64) error {
	return c.Delete(EventRegistrationModel, ids)
}

// GetEventRegistration gets event.registration existing record.
func (c *Client) GetEventRegistration(id int64) (*EventRegistration, error) {
	ers, err := c.GetEventRegistrations([]int64{id})
	if err != nil {
		return nil, err
	}
	if ers != nil && len(*ers) > 0 {
		return &((*ers)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.registration not found", id)
}

// GetEventRegistrations gets event.registration existing records.
func (c *Client) GetEventRegistrations(ids []int64) (*EventRegistrations, error) {
	ers := &EventRegistrations{}
	if err := c.Read(EventRegistrationModel, ids, nil, ers); err != nil {
		return nil, err
	}
	return ers, nil
}

// FindEventRegistration finds event.registration record by querying it with criteria.
func (c *Client) FindEventRegistration(criteria *Criteria) (*EventRegistration, error) {
	ers := &EventRegistrations{}
	if err := c.SearchRead(EventRegistrationModel, criteria, NewOptions().Limit(1), ers); err != nil {
		return nil, err
	}
	if ers != nil && len(*ers) > 0 {
		return &((*ers)[0]), nil
	}
	return nil, fmt.Errorf("event.registration was not found with criteria %v", criteria)
}

// FindEventRegistrations finds event.registration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrations(criteria *Criteria, options *Options) (*EventRegistrations, error) {
	ers := &EventRegistrations{}
	if err := c.SearchRead(EventRegistrationModel, criteria, options, ers); err != nil {
		return nil, err
	}
	return ers, nil
}

// FindEventRegistrationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventRegistrationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventRegistrationId finds record id by querying it with criteria.
func (c *Client) FindEventRegistrationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventRegistrationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.registration was not found with criteria %v and options %v", criteria, options)
}
