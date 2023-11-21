package odoo

import (
	"fmt"
)

// AppointmentInvite represents appointment.invite model.
type AppointmentInvite struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken             *String    `xmlrpc:"access_token,omptempty"`
	AppointmentTypeCount    *Int       `xmlrpc:"appointment_type_count,omptempty"`
	AppointmentTypeIds      *Relation  `xmlrpc:"appointment_type_ids,omptempty"`
	AppointmentTypeInfoMsg  *String    `xmlrpc:"appointment_type_info_msg,omptempty"`
	BaseBookUrl             *String    `xmlrpc:"base_book_url,omptempty"`
	BookUrl                 *String    `xmlrpc:"book_url,omptempty"`
	CalendarEventIds        *Relation  `xmlrpc:"calendar_event_ids,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	RedirectUrl             *String    `xmlrpc:"redirect_url,omptempty"`
	ShortCode               *String    `xmlrpc:"short_code,omptempty"`
	ShortCodeFormatWarning  *Bool      `xmlrpc:"short_code_format_warning,omptempty"`
	ShortCodeUniqueWarning  *Bool      `xmlrpc:"short_code_unique_warning,omptempty"`
	StaffUserIds            *Relation  `xmlrpc:"staff_user_ids,omptempty"`
	StaffUsersChoice        *Selection `xmlrpc:"staff_users_choice,omptempty"`
	SuggestedStaffUserCount *Int       `xmlrpc:"suggested_staff_user_count,omptempty"`
	SuggestedStaffUserIds   *Relation  `xmlrpc:"suggested_staff_user_ids,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AppointmentInvites represents array of appointment.invite model.
type AppointmentInvites []AppointmentInvite

// AppointmentInviteModel is the odoo model name.
const AppointmentInviteModel = "appointment.invite"

// Many2One convert AppointmentInvite to *Many2One.
func (ai *AppointmentInvite) Many2One() *Many2One {
	return NewMany2One(ai.Id.Get(), "")
}

// CreateAppointmentInvite creates a new appointment.invite model and returns its id.
func (c *Client) CreateAppointmentInvite(ai *AppointmentInvite) (int64, error) {
	ids, err := c.CreateAppointmentInvites([]*AppointmentInvite{ai})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentInvite creates a new appointment.invite model and returns its id.
func (c *Client) CreateAppointmentInvites(ais []*AppointmentInvite) ([]int64, error) {
	var vv []interface{}
	for _, v := range ais {
		vv = append(vv, v)
	}
	return c.Create(AppointmentInviteModel, vv)
}

// UpdateAppointmentInvite updates an existing appointment.invite record.
func (c *Client) UpdateAppointmentInvite(ai *AppointmentInvite) error {
	return c.UpdateAppointmentInvites([]int64{ai.Id.Get()}, ai)
}

// UpdateAppointmentInvites updates existing appointment.invite records.
// All records (represented by ids) will be updated by ai values.
func (c *Client) UpdateAppointmentInvites(ids []int64, ai *AppointmentInvite) error {
	return c.Update(AppointmentInviteModel, ids, ai)
}

// DeleteAppointmentInvite deletes an existing appointment.invite record.
func (c *Client) DeleteAppointmentInvite(id int64) error {
	return c.DeleteAppointmentInvites([]int64{id})
}

// DeleteAppointmentInvites deletes existing appointment.invite records.
func (c *Client) DeleteAppointmentInvites(ids []int64) error {
	return c.Delete(AppointmentInviteModel, ids)
}

// GetAppointmentInvite gets appointment.invite existing record.
func (c *Client) GetAppointmentInvite(id int64) (*AppointmentInvite, error) {
	ais, err := c.GetAppointmentInvites([]int64{id})
	if err != nil {
		return nil, err
	}
	if ais != nil && len(*ais) > 0 {
		return &((*ais)[0]), nil
	}
	return nil, fmt.Errorf("id %v of appointment.invite not found", id)
}

// GetAppointmentInvites gets appointment.invite existing records.
func (c *Client) GetAppointmentInvites(ids []int64) (*AppointmentInvites, error) {
	ais := &AppointmentInvites{}
	if err := c.Read(AppointmentInviteModel, ids, nil, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAppointmentInvite finds appointment.invite record by querying it with criteria.
func (c *Client) FindAppointmentInvite(criteria *Criteria) (*AppointmentInvite, error) {
	ais := &AppointmentInvites{}
	if err := c.SearchRead(AppointmentInviteModel, criteria, NewOptions().Limit(1), ais); err != nil {
		return nil, err
	}
	if ais != nil && len(*ais) > 0 {
		return &((*ais)[0]), nil
	}
	return nil, fmt.Errorf("appointment.invite was not found with criteria %v", criteria)
}

// FindAppointmentInvites finds appointment.invite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentInvites(criteria *Criteria, options *Options) (*AppointmentInvites, error) {
	ais := &AppointmentInvites{}
	if err := c.SearchRead(AppointmentInviteModel, criteria, options, ais); err != nil {
		return nil, err
	}
	return ais, nil
}

// FindAppointmentInviteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentInviteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AppointmentInviteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAppointmentInviteId finds record id by querying it with criteria.
func (c *Client) FindAppointmentInviteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentInviteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("appointment.invite was not found with criteria %v and options %v", criteria, options)
}
