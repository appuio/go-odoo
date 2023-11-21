package odoo

import (
	"fmt"
)

// AppointmentOnboardingLink represents appointment.onboarding.link model.
type AppointmentOnboardingLink struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	AppointmentTypeId *Many2One `xmlrpc:"appointment_type_id,omptempty"`
	BaseBookUrl       *String   `xmlrpc:"base_book_url,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	ShortCode         *String   `xmlrpc:"short_code,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AppointmentOnboardingLinks represents array of appointment.onboarding.link model.
type AppointmentOnboardingLinks []AppointmentOnboardingLink

// AppointmentOnboardingLinkModel is the odoo model name.
const AppointmentOnboardingLinkModel = "appointment.onboarding.link"

// Many2One convert AppointmentOnboardingLink to *Many2One.
func (aol *AppointmentOnboardingLink) Many2One() *Many2One {
	return NewMany2One(aol.Id.Get(), "")
}

// CreateAppointmentOnboardingLink creates a new appointment.onboarding.link model and returns its id.
func (c *Client) CreateAppointmentOnboardingLink(aol *AppointmentOnboardingLink) (int64, error) {
	ids, err := c.CreateAppointmentOnboardingLinks([]*AppointmentOnboardingLink{aol})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentOnboardingLink creates a new appointment.onboarding.link model and returns its id.
func (c *Client) CreateAppointmentOnboardingLinks(aols []*AppointmentOnboardingLink) ([]int64, error) {
	var vv []interface{}
	for _, v := range aols {
		vv = append(vv, v)
	}
	return c.Create(AppointmentOnboardingLinkModel, vv)
}

// UpdateAppointmentOnboardingLink updates an existing appointment.onboarding.link record.
func (c *Client) UpdateAppointmentOnboardingLink(aol *AppointmentOnboardingLink) error {
	return c.UpdateAppointmentOnboardingLinks([]int64{aol.Id.Get()}, aol)
}

// UpdateAppointmentOnboardingLinks updates existing appointment.onboarding.link records.
// All records (represented by ids) will be updated by aol values.
func (c *Client) UpdateAppointmentOnboardingLinks(ids []int64, aol *AppointmentOnboardingLink) error {
	return c.Update(AppointmentOnboardingLinkModel, ids, aol)
}

// DeleteAppointmentOnboardingLink deletes an existing appointment.onboarding.link record.
func (c *Client) DeleteAppointmentOnboardingLink(id int64) error {
	return c.DeleteAppointmentOnboardingLinks([]int64{id})
}

// DeleteAppointmentOnboardingLinks deletes existing appointment.onboarding.link records.
func (c *Client) DeleteAppointmentOnboardingLinks(ids []int64) error {
	return c.Delete(AppointmentOnboardingLinkModel, ids)
}

// GetAppointmentOnboardingLink gets appointment.onboarding.link existing record.
func (c *Client) GetAppointmentOnboardingLink(id int64) (*AppointmentOnboardingLink, error) {
	aols, err := c.GetAppointmentOnboardingLinks([]int64{id})
	if err != nil {
		return nil, err
	}
	if aols != nil && len(*aols) > 0 {
		return &((*aols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of appointment.onboarding.link not found", id)
}

// GetAppointmentOnboardingLinks gets appointment.onboarding.link existing records.
func (c *Client) GetAppointmentOnboardingLinks(ids []int64) (*AppointmentOnboardingLinks, error) {
	aols := &AppointmentOnboardingLinks{}
	if err := c.Read(AppointmentOnboardingLinkModel, ids, nil, aols); err != nil {
		return nil, err
	}
	return aols, nil
}

// FindAppointmentOnboardingLink finds appointment.onboarding.link record by querying it with criteria.
func (c *Client) FindAppointmentOnboardingLink(criteria *Criteria) (*AppointmentOnboardingLink, error) {
	aols := &AppointmentOnboardingLinks{}
	if err := c.SearchRead(AppointmentOnboardingLinkModel, criteria, NewOptions().Limit(1), aols); err != nil {
		return nil, err
	}
	if aols != nil && len(*aols) > 0 {
		return &((*aols)[0]), nil
	}
	return nil, fmt.Errorf("appointment.onboarding.link was not found with criteria %v", criteria)
}

// FindAppointmentOnboardingLinks finds appointment.onboarding.link records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentOnboardingLinks(criteria *Criteria, options *Options) (*AppointmentOnboardingLinks, error) {
	aols := &AppointmentOnboardingLinks{}
	if err := c.SearchRead(AppointmentOnboardingLinkModel, criteria, options, aols); err != nil {
		return nil, err
	}
	return aols, nil
}

// FindAppointmentOnboardingLinkIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentOnboardingLinkIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AppointmentOnboardingLinkModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAppointmentOnboardingLinkId finds record id by querying it with criteria.
func (c *Client) FindAppointmentOnboardingLinkId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentOnboardingLinkModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("appointment.onboarding.link was not found with criteria %v and options %v", criteria, options)
}
