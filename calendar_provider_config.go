package odoo

import (
	"fmt"
)

// CalendarProviderConfig represents calendar.provider.config model.
type CalendarProviderConfig struct {
	LastUpdate                       *Time      `xmlrpc:"__last_update,omptempty"`
	CalClientId                      *String    `xmlrpc:"cal_client_id,omptempty"`
	CalClientSecret                  *String    `xmlrpc:"cal_client_secret,omptempty"`
	CreateDate                       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                        *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                      *String    `xmlrpc:"display_name,omptempty"`
	ExternalCalendarProvider         *Selection `xmlrpc:"external_calendar_provider,omptempty"`
	Id                               *Int       `xmlrpc:"id,omptempty"`
	MicrosoftOutlookClientIdentifier *String    `xmlrpc:"microsoft_outlook_client_identifier,omptempty"`
	MicrosoftOutlookClientSecret     *String    `xmlrpc:"microsoft_outlook_client_secret,omptempty"`
	WriteDate                        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// CalendarProviderConfigs represents array of calendar.provider.config model.
type CalendarProviderConfigs []CalendarProviderConfig

// CalendarProviderConfigModel is the odoo model name.
const CalendarProviderConfigModel = "calendar.provider.config"

// Many2One convert CalendarProviderConfig to *Many2One.
func (cpc *CalendarProviderConfig) Many2One() *Many2One {
	return NewMany2One(cpc.Id.Get(), "")
}

// CreateCalendarProviderConfig creates a new calendar.provider.config model and returns its id.
func (c *Client) CreateCalendarProviderConfig(cpc *CalendarProviderConfig) (int64, error) {
	ids, err := c.CreateCalendarProviderConfigs([]*CalendarProviderConfig{cpc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCalendarProviderConfig creates a new calendar.provider.config model and returns its id.
func (c *Client) CreateCalendarProviderConfigs(cpcs []*CalendarProviderConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpcs {
		vv = append(vv, v)
	}
	return c.Create(CalendarProviderConfigModel, vv)
}

// UpdateCalendarProviderConfig updates an existing calendar.provider.config record.
func (c *Client) UpdateCalendarProviderConfig(cpc *CalendarProviderConfig) error {
	return c.UpdateCalendarProviderConfigs([]int64{cpc.Id.Get()}, cpc)
}

// UpdateCalendarProviderConfigs updates existing calendar.provider.config records.
// All records (represented by ids) will be updated by cpc values.
func (c *Client) UpdateCalendarProviderConfigs(ids []int64, cpc *CalendarProviderConfig) error {
	return c.Update(CalendarProviderConfigModel, ids, cpc)
}

// DeleteCalendarProviderConfig deletes an existing calendar.provider.config record.
func (c *Client) DeleteCalendarProviderConfig(id int64) error {
	return c.DeleteCalendarProviderConfigs([]int64{id})
}

// DeleteCalendarProviderConfigs deletes existing calendar.provider.config records.
func (c *Client) DeleteCalendarProviderConfigs(ids []int64) error {
	return c.Delete(CalendarProviderConfigModel, ids)
}

// GetCalendarProviderConfig gets calendar.provider.config existing record.
func (c *Client) GetCalendarProviderConfig(id int64) (*CalendarProviderConfig, error) {
	cpcs, err := c.GetCalendarProviderConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if cpcs != nil && len(*cpcs) > 0 {
		return &((*cpcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of calendar.provider.config not found", id)
}

// GetCalendarProviderConfigs gets calendar.provider.config existing records.
func (c *Client) GetCalendarProviderConfigs(ids []int64) (*CalendarProviderConfigs, error) {
	cpcs := &CalendarProviderConfigs{}
	if err := c.Read(CalendarProviderConfigModel, ids, nil, cpcs); err != nil {
		return nil, err
	}
	return cpcs, nil
}

// FindCalendarProviderConfig finds calendar.provider.config record by querying it with criteria.
func (c *Client) FindCalendarProviderConfig(criteria *Criteria) (*CalendarProviderConfig, error) {
	cpcs := &CalendarProviderConfigs{}
	if err := c.SearchRead(CalendarProviderConfigModel, criteria, NewOptions().Limit(1), cpcs); err != nil {
		return nil, err
	}
	if cpcs != nil && len(*cpcs) > 0 {
		return &((*cpcs)[0]), nil
	}
	return nil, fmt.Errorf("calendar.provider.config was not found with criteria %v", criteria)
}

// FindCalendarProviderConfigs finds calendar.provider.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarProviderConfigs(criteria *Criteria, options *Options) (*CalendarProviderConfigs, error) {
	cpcs := &CalendarProviderConfigs{}
	if err := c.SearchRead(CalendarProviderConfigModel, criteria, options, cpcs); err != nil {
		return nil, err
	}
	return cpcs, nil
}

// FindCalendarProviderConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCalendarProviderConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CalendarProviderConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCalendarProviderConfigId finds record id by querying it with criteria.
func (c *Client) FindCalendarProviderConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CalendarProviderConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("calendar.provider.config was not found with criteria %v and options %v", criteria, options)
}
