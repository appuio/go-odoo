package odoo

import (
	"fmt"
)

// AppointmentType represents appointment.type model.
type AppointmentType struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	AppointmentCount             *Int       `xmlrpc:"appointment_count,omptempty"`
	AppointmentCountReport       *Int       `xmlrpc:"appointment_count_report,omptempty"`
	AppointmentDuration          *Float     `xmlrpc:"appointment_duration,omptempty"`
	AppointmentDurationFormatted *String    `xmlrpc:"appointment_duration_formatted,omptempty"`
	AppointmentInviteIds         *Relation  `xmlrpc:"appointment_invite_ids,omptempty"`
	AppointmentTz                *Selection `xmlrpc:"appointment_tz,omptempty"`
	AssignMethod                 *Selection `xmlrpc:"assign_method,omptempty"`
	AvatarsDisplay               *Selection `xmlrpc:"avatars_display,omptempty"`
	CanPublish                   *Bool      `xmlrpc:"can_publish,omptempty"`
	Category                     *Selection `xmlrpc:"category,omptempty"`
	CountryIds                   *Relation  `xmlrpc:"country_ids,omptempty"`
	CoverProperties              *String    `xmlrpc:"cover_properties,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	FailedMessageIds             *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	IsPublished                  *Bool      `xmlrpc:"is_published,omptempty"`
	IsSeoOptimized               *Bool      `xmlrpc:"is_seo_optimized,omptempty"`
	LeadCreate                   *Bool      `xmlrpc:"lead_create,omptempty"`
	Location                     *String    `xmlrpc:"location,omptempty"`
	LocationId                   *Many2One  `xmlrpc:"location_id,omptempty"`
	MaxScheduleDays              *Int       `xmlrpc:"max_schedule_days,omptempty"`
	MeetingIds                   *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageConfirmation          *String    `xmlrpc:"message_confirmation,omptempty"`
	MessageContent               *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIntro                 *String    `xmlrpc:"message_intro,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId      *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MinCancellationHours         *Float     `xmlrpc:"min_cancellation_hours,omptempty"`
	MinScheduleHours             *Float     `xmlrpc:"min_schedule_hours,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	OpportunityId                *Many2One  `xmlrpc:"opportunity_id,omptempty"`
	QuestionIds                  *Relation  `xmlrpc:"question_ids,omptempty"`
	ReminderIds                  *Relation  `xmlrpc:"reminder_ids,omptempty"`
	SeoName                      *String    `xmlrpc:"seo_name,omptempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omptempty"`
	SlotIds                      *Relation  `xmlrpc:"slot_ids,omptempty"`
	StaffUserCount               *Int       `xmlrpc:"staff_user_count,omptempty"`
	StaffUserIds                 *Relation  `xmlrpc:"staff_user_ids,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WebsiteMetaDescription       *String    `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords          *String    `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg             *String    `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle             *String    `xmlrpc:"website_meta_title,omptempty"`
	WebsitePublished             *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                   *String    `xmlrpc:"website_url,omptempty"`
	WorkHoursActivated           *Bool      `xmlrpc:"work_hours_activated,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AppointmentTypes represents array of appointment.type model.
type AppointmentTypes []AppointmentType

// AppointmentTypeModel is the odoo model name.
const AppointmentTypeModel = "appointment.type"

// Many2One convert AppointmentType to *Many2One.
func (at *AppointmentType) Many2One() *Many2One {
	return NewMany2One(at.Id.Get(), "")
}

// CreateAppointmentType creates a new appointment.type model and returns its id.
func (c *Client) CreateAppointmentType(at *AppointmentType) (int64, error) {
	ids, err := c.CreateAppointmentTypes([]*AppointmentType{at})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentType creates a new appointment.type model and returns its id.
func (c *Client) CreateAppointmentTypes(ats []*AppointmentType) ([]int64, error) {
	var vv []interface{}
	for _, v := range ats {
		vv = append(vv, v)
	}
	return c.Create(AppointmentTypeModel, vv)
}

// UpdateAppointmentType updates an existing appointment.type record.
func (c *Client) UpdateAppointmentType(at *AppointmentType) error {
	return c.UpdateAppointmentTypes([]int64{at.Id.Get()}, at)
}

// UpdateAppointmentTypes updates existing appointment.type records.
// All records (represented by ids) will be updated by at values.
func (c *Client) UpdateAppointmentTypes(ids []int64, at *AppointmentType) error {
	return c.Update(AppointmentTypeModel, ids, at)
}

// DeleteAppointmentType deletes an existing appointment.type record.
func (c *Client) DeleteAppointmentType(id int64) error {
	return c.DeleteAppointmentTypes([]int64{id})
}

// DeleteAppointmentTypes deletes existing appointment.type records.
func (c *Client) DeleteAppointmentTypes(ids []int64) error {
	return c.Delete(AppointmentTypeModel, ids)
}

// GetAppointmentType gets appointment.type existing record.
func (c *Client) GetAppointmentType(id int64) (*AppointmentType, error) {
	ats, err := c.GetAppointmentTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if ats != nil && len(*ats) > 0 {
		return &((*ats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of appointment.type not found", id)
}

// GetAppointmentTypes gets appointment.type existing records.
func (c *Client) GetAppointmentTypes(ids []int64) (*AppointmentTypes, error) {
	ats := &AppointmentTypes{}
	if err := c.Read(AppointmentTypeModel, ids, nil, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAppointmentType finds appointment.type record by querying it with criteria.
func (c *Client) FindAppointmentType(criteria *Criteria) (*AppointmentType, error) {
	ats := &AppointmentTypes{}
	if err := c.SearchRead(AppointmentTypeModel, criteria, NewOptions().Limit(1), ats); err != nil {
		return nil, err
	}
	if ats != nil && len(*ats) > 0 {
		return &((*ats)[0]), nil
	}
	return nil, fmt.Errorf("appointment.type was not found with criteria %v", criteria)
}

// FindAppointmentTypes finds appointment.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentTypes(criteria *Criteria, options *Options) (*AppointmentTypes, error) {
	ats := &AppointmentTypes{}
	if err := c.SearchRead(AppointmentTypeModel, criteria, options, ats); err != nil {
		return nil, err
	}
	return ats, nil
}

// FindAppointmentTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AppointmentTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAppointmentTypeId finds record id by querying it with criteria.
func (c *Client) FindAppointmentTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("appointment.type was not found with criteria %v and options %v", criteria, options)
}
