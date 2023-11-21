package odoo

import (
	"fmt"
)

// AppointmentSlot represents appointment.slot model.
type AppointmentSlot struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Allday            *Bool      `xmlrpc:"allday,omptempty"`
	AppointmentTypeId *Many2One  `xmlrpc:"appointment_type_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Duration          *Float     `xmlrpc:"duration,omptempty"`
	EndDatetime       *Time      `xmlrpc:"end_datetime,omptempty"`
	EndHour           *Float     `xmlrpc:"end_hour,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	RestrictToUserIds *Relation  `xmlrpc:"restrict_to_user_ids,omptempty"`
	SlotType          *Selection `xmlrpc:"slot_type,omptempty"`
	StartDatetime     *Time      `xmlrpc:"start_datetime,omptempty"`
	StartHour         *Float     `xmlrpc:"start_hour,omptempty"`
	Weekday           *Selection `xmlrpc:"weekday,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AppointmentSlots represents array of appointment.slot model.
type AppointmentSlots []AppointmentSlot

// AppointmentSlotModel is the odoo model name.
const AppointmentSlotModel = "appointment.slot"

// Many2One convert AppointmentSlot to *Many2One.
func (as *AppointmentSlot) Many2One() *Many2One {
	return NewMany2One(as.Id.Get(), "")
}

// CreateAppointmentSlot creates a new appointment.slot model and returns its id.
func (c *Client) CreateAppointmentSlot(as *AppointmentSlot) (int64, error) {
	ids, err := c.CreateAppointmentSlots([]*AppointmentSlot{as})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentSlot creates a new appointment.slot model and returns its id.
func (c *Client) CreateAppointmentSlots(ass []*AppointmentSlot) ([]int64, error) {
	var vv []interface{}
	for _, v := range ass {
		vv = append(vv, v)
	}
	return c.Create(AppointmentSlotModel, vv)
}

// UpdateAppointmentSlot updates an existing appointment.slot record.
func (c *Client) UpdateAppointmentSlot(as *AppointmentSlot) error {
	return c.UpdateAppointmentSlots([]int64{as.Id.Get()}, as)
}

// UpdateAppointmentSlots updates existing appointment.slot records.
// All records (represented by ids) will be updated by as values.
func (c *Client) UpdateAppointmentSlots(ids []int64, as *AppointmentSlot) error {
	return c.Update(AppointmentSlotModel, ids, as)
}

// DeleteAppointmentSlot deletes an existing appointment.slot record.
func (c *Client) DeleteAppointmentSlot(id int64) error {
	return c.DeleteAppointmentSlots([]int64{id})
}

// DeleteAppointmentSlots deletes existing appointment.slot records.
func (c *Client) DeleteAppointmentSlots(ids []int64) error {
	return c.Delete(AppointmentSlotModel, ids)
}

// GetAppointmentSlot gets appointment.slot existing record.
func (c *Client) GetAppointmentSlot(id int64) (*AppointmentSlot, error) {
	ass, err := c.GetAppointmentSlots([]int64{id})
	if err != nil {
		return nil, err
	}
	if ass != nil && len(*ass) > 0 {
		return &((*ass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of appointment.slot not found", id)
}

// GetAppointmentSlots gets appointment.slot existing records.
func (c *Client) GetAppointmentSlots(ids []int64) (*AppointmentSlots, error) {
	ass := &AppointmentSlots{}
	if err := c.Read(AppointmentSlotModel, ids, nil, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindAppointmentSlot finds appointment.slot record by querying it with criteria.
func (c *Client) FindAppointmentSlot(criteria *Criteria) (*AppointmentSlot, error) {
	ass := &AppointmentSlots{}
	if err := c.SearchRead(AppointmentSlotModel, criteria, NewOptions().Limit(1), ass); err != nil {
		return nil, err
	}
	if ass != nil && len(*ass) > 0 {
		return &((*ass)[0]), nil
	}
	return nil, fmt.Errorf("appointment.slot was not found with criteria %v", criteria)
}

// FindAppointmentSlots finds appointment.slot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentSlots(criteria *Criteria, options *Options) (*AppointmentSlots, error) {
	ass := &AppointmentSlots{}
	if err := c.SearchRead(AppointmentSlotModel, criteria, options, ass); err != nil {
		return nil, err
	}
	return ass, nil
}

// FindAppointmentSlotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentSlotIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AppointmentSlotModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAppointmentSlotId finds record id by querying it with criteria.
func (c *Client) FindAppointmentSlotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentSlotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("appointment.slot was not found with criteria %v and options %v", criteria, options)
}
