package odoo

import (
	"fmt"
)

// BaseAutomation represents base.automation model.
type BaseAutomation struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	ActionServerId                *Many2One  `xmlrpc:"action_server_id,omptempty"`
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omptempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omptempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserFieldName         *String    `xmlrpc:"activity_user_field_name,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	ActivityUserType              *Selection `xmlrpc:"activity_user_type,omptempty"`
	BindingModelId                *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType                   *Selection `xmlrpc:"binding_type,omptempty"`
	BindingViewTypes              *String    `xmlrpc:"binding_view_types,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	Code                          *String    `xmlrpc:"code,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CrudModelId                   *Many2One  `xmlrpc:"crud_model_id,omptempty"`
	CrudModelName                 *String    `xmlrpc:"crud_model_name,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	FieldsLines                   *Relation  `xmlrpc:"fields_lines,omptempty"`
	FilterDomain                  *String    `xmlrpc:"filter_domain,omptempty"`
	FilterPreDomain               *String    `xmlrpc:"filter_pre_domain,omptempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omptempty"`
	Help                          *String    `xmlrpc:"help,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	LastRun                       *Time      `xmlrpc:"last_run,omptempty"`
	LeastDelayMsg                 *String    `xmlrpc:"least_delay_msg,omptempty"`
	LinkFieldId                   *Many2One  `xmlrpc:"link_field_id,omptempty"`
	MailPostAutofollow            *Bool      `xmlrpc:"mail_post_autofollow,omptempty"`
	MailPostMethod                *Selection `xmlrpc:"mail_post_method,omptempty"`
	ModelId                       *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName                     *String    `xmlrpc:"model_name,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	OnChangeFieldIds              *Relation  `xmlrpc:"on_change_field_ids,omptempty"`
	PartnerIds                    *Relation  `xmlrpc:"partner_ids,omptempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omptempty"`
	SmsMethod                     *Selection `xmlrpc:"sms_method,omptempty"`
	SmsTemplateId                 *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	State                         *Selection `xmlrpc:"state,omptempty"`
	TemplateId                    *Many2One  `xmlrpc:"template_id,omptempty"`
	TrgDateCalendarId             *Many2One  `xmlrpc:"trg_date_calendar_id,omptempty"`
	TrgDateId                     *Many2One  `xmlrpc:"trg_date_id,omptempty"`
	TrgDateRange                  *Int       `xmlrpc:"trg_date_range,omptempty"`
	TrgDateRangeType              *Selection `xmlrpc:"trg_date_range_type,omptempty"`
	TrgDateResourceFieldId        *Many2One  `xmlrpc:"trg_date_resource_field_id,omptempty"`
	Trigger                       *Selection `xmlrpc:"trigger,omptempty"`
	TriggerFieldIds               *Relation  `xmlrpc:"trigger_field_ids,omptempty"`
	Type                          *String    `xmlrpc:"type,omptempty"`
	Usage                         *Selection `xmlrpc:"usage,omptempty"`
	WebsitePath                   *String    `xmlrpc:"website_path,omptempty"`
	WebsitePublished              *Bool      `xmlrpc:"website_published,omptempty"`
	WebsiteUrl                    *String    `xmlrpc:"website_url,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId                         *String    `xmlrpc:"xml_id,omptempty"`
}

// BaseAutomations represents array of base.automation model.
type BaseAutomations []BaseAutomation

// BaseAutomationModel is the odoo model name.
const BaseAutomationModel = "base.automation"

// Many2One convert BaseAutomation to *Many2One.
func (ba *BaseAutomation) Many2One() *Many2One {
	return NewMany2One(ba.Id.Get(), "")
}

// CreateBaseAutomation creates a new base.automation model and returns its id.
func (c *Client) CreateBaseAutomation(ba *BaseAutomation) (int64, error) {
	ids, err := c.CreateBaseAutomations([]*BaseAutomation{ba})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseAutomation creates a new base.automation model and returns its id.
func (c *Client) CreateBaseAutomations(bas []*BaseAutomation) ([]int64, error) {
	var vv []interface{}
	for _, v := range bas {
		vv = append(vv, v)
	}
	return c.Create(BaseAutomationModel, vv)
}

// UpdateBaseAutomation updates an existing base.automation record.
func (c *Client) UpdateBaseAutomation(ba *BaseAutomation) error {
	return c.UpdateBaseAutomations([]int64{ba.Id.Get()}, ba)
}

// UpdateBaseAutomations updates existing base.automation records.
// All records (represented by ids) will be updated by ba values.
func (c *Client) UpdateBaseAutomations(ids []int64, ba *BaseAutomation) error {
	return c.Update(BaseAutomationModel, ids, ba)
}

// DeleteBaseAutomation deletes an existing base.automation record.
func (c *Client) DeleteBaseAutomation(id int64) error {
	return c.DeleteBaseAutomations([]int64{id})
}

// DeleteBaseAutomations deletes existing base.automation records.
func (c *Client) DeleteBaseAutomations(ids []int64) error {
	return c.Delete(BaseAutomationModel, ids)
}

// GetBaseAutomation gets base.automation existing record.
func (c *Client) GetBaseAutomation(id int64) (*BaseAutomation, error) {
	bas, err := c.GetBaseAutomations([]int64{id})
	if err != nil {
		return nil, err
	}
	if bas != nil && len(*bas) > 0 {
		return &((*bas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.automation not found", id)
}

// GetBaseAutomations gets base.automation existing records.
func (c *Client) GetBaseAutomations(ids []int64) (*BaseAutomations, error) {
	bas := &BaseAutomations{}
	if err := c.Read(BaseAutomationModel, ids, nil, bas); err != nil {
		return nil, err
	}
	return bas, nil
}

// FindBaseAutomation finds base.automation record by querying it with criteria.
func (c *Client) FindBaseAutomation(criteria *Criteria) (*BaseAutomation, error) {
	bas := &BaseAutomations{}
	if err := c.SearchRead(BaseAutomationModel, criteria, NewOptions().Limit(1), bas); err != nil {
		return nil, err
	}
	if bas != nil && len(*bas) > 0 {
		return &((*bas)[0]), nil
	}
	return nil, fmt.Errorf("base.automation was not found with criteria %v", criteria)
}

// FindBaseAutomations finds base.automation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomations(criteria *Criteria, options *Options) (*BaseAutomations, error) {
	bas := &BaseAutomations{}
	if err := c.SearchRead(BaseAutomationModel, criteria, options, bas); err != nil {
		return nil, err
	}
	return bas, nil
}

// FindBaseAutomationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseAutomationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseAutomationId finds record id by querying it with criteria.
func (c *Client) FindBaseAutomationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseAutomationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.automation was not found with criteria %v and options %v", criteria, options)
}
