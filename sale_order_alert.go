package odoo

import (
	"fmt"
)

// SaleOrderAlert represents sale.order.alert model.
type SaleOrderAlert struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	Action                        *Selection `xmlrpc:"action,omptempty"`
	ActionServerId                *Many2One  `xmlrpc:"action_server_id,omptempty"`
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omptempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omptempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUser                  *Selection `xmlrpc:"activity_user,omptempty"`
	ActivityUserFieldName         *String    `xmlrpc:"activity_user_field_name,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	ActivityUserIds               *Relation  `xmlrpc:"activity_user_ids,omptempty"`
	ActivityUserType              *Selection `xmlrpc:"activity_user_type,omptempty"`
	AutomationId                  *Many2One  `xmlrpc:"automation_id,omptempty"`
	BindingModelId                *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType                   *Selection `xmlrpc:"binding_type,omptempty"`
	BindingViewTypes              *String    `xmlrpc:"binding_view_types,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	Code                          *String    `xmlrpc:"code,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CronNextcall                  *Time      `xmlrpc:"cron_nextcall,omptempty"`
	CrudModelId                   *Many2One  `xmlrpc:"crud_model_id,omptempty"`
	CrudModelName                 *String    `xmlrpc:"crud_model_name,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	CustomerIds                   *Relation  `xmlrpc:"customer_ids,omptempty"`
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
	MrrChangeAmount               *Float     `xmlrpc:"mrr_change_amount,omptempty"`
	MrrChangePeriod               *Selection `xmlrpc:"mrr_change_period,omptempty"`
	MrrChangeUnit                 *Selection `xmlrpc:"mrr_change_unit,omptempty"`
	MrrMax                        *Float     `xmlrpc:"mrr_max,omptempty"`
	MrrMin                        *Float     `xmlrpc:"mrr_min,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	OnChangeFieldIds              *Relation  `xmlrpc:"on_change_field_ids,omptempty"`
	PartnerIds                    *Relation  `xmlrpc:"partner_ids,omptempty"`
	ProductIds                    *Relation  `xmlrpc:"product_ids,omptempty"`
	RatingOperator                *Selection `xmlrpc:"rating_operator,omptempty"`
	RatingPercentage              *Int       `xmlrpc:"rating_percentage,omptempty"`
	Sequence                      *Int       `xmlrpc:"sequence,omptempty"`
	SmsMethod                     *Selection `xmlrpc:"sms_method,omptempty"`
	SmsTemplateId                 *Many2One  `xmlrpc:"sms_template_id,omptempty"`
	StageFromId                   *Many2One  `xmlrpc:"stage_from_id,omptempty"`
	StageId                       *Many2One  `xmlrpc:"stage_id,omptempty"`
	StageToId                     *Many2One  `xmlrpc:"stage_to_id,omptempty"`
	State                         *Selection `xmlrpc:"state,omptempty"`
	SubscriptionCount             *Int       `xmlrpc:"subscription_count,omptempty"`
	SubscriptionTemplateIds       *Relation  `xmlrpc:"subscription_template_ids,omptempty"`
	TeamIds                       *Relation  `xmlrpc:"team_ids,omptempty"`
	TemplateId                    *Many2One  `xmlrpc:"template_id,omptempty"`
	TrgDateCalendarId             *Many2One  `xmlrpc:"trg_date_calendar_id,omptempty"`
	TrgDateId                     *Many2One  `xmlrpc:"trg_date_id,omptempty"`
	TrgDateRange                  *Int       `xmlrpc:"trg_date_range,omptempty"`
	TrgDateRangeType              *Selection `xmlrpc:"trg_date_range_type,omptempty"`
	TrgDateResourceFieldId        *Many2One  `xmlrpc:"trg_date_resource_field_id,omptempty"`
	Trigger                       *Selection `xmlrpc:"trigger,omptempty"`
	TriggerCondition              *Selection `xmlrpc:"trigger_condition,omptempty"`
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

// SaleOrderAlerts represents array of sale.order.alert model.
type SaleOrderAlerts []SaleOrderAlert

// SaleOrderAlertModel is the odoo model name.
const SaleOrderAlertModel = "sale.order.alert"

// Many2One convert SaleOrderAlert to *Many2One.
func (soa *SaleOrderAlert) Many2One() *Many2One {
	return NewMany2One(soa.Id.Get(), "")
}

// CreateSaleOrderAlert creates a new sale.order.alert model and returns its id.
func (c *Client) CreateSaleOrderAlert(soa *SaleOrderAlert) (int64, error) {
	ids, err := c.CreateSaleOrderAlerts([]*SaleOrderAlert{soa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSaleOrderAlert creates a new sale.order.alert model and returns its id.
func (c *Client) CreateSaleOrderAlerts(soas []*SaleOrderAlert) ([]int64, error) {
	var vv []interface{}
	for _, v := range soas {
		vv = append(vv, v)
	}
	return c.Create(SaleOrderAlertModel, vv)
}

// UpdateSaleOrderAlert updates an existing sale.order.alert record.
func (c *Client) UpdateSaleOrderAlert(soa *SaleOrderAlert) error {
	return c.UpdateSaleOrderAlerts([]int64{soa.Id.Get()}, soa)
}

// UpdateSaleOrderAlerts updates existing sale.order.alert records.
// All records (represented by ids) will be updated by soa values.
func (c *Client) UpdateSaleOrderAlerts(ids []int64, soa *SaleOrderAlert) error {
	return c.Update(SaleOrderAlertModel, ids, soa)
}

// DeleteSaleOrderAlert deletes an existing sale.order.alert record.
func (c *Client) DeleteSaleOrderAlert(id int64) error {
	return c.DeleteSaleOrderAlerts([]int64{id})
}

// DeleteSaleOrderAlerts deletes existing sale.order.alert records.
func (c *Client) DeleteSaleOrderAlerts(ids []int64) error {
	return c.Delete(SaleOrderAlertModel, ids)
}

// GetSaleOrderAlert gets sale.order.alert existing record.
func (c *Client) GetSaleOrderAlert(id int64) (*SaleOrderAlert, error) {
	soas, err := c.GetSaleOrderAlerts([]int64{id})
	if err != nil {
		return nil, err
	}
	if soas != nil && len(*soas) > 0 {
		return &((*soas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sale.order.alert not found", id)
}

// GetSaleOrderAlerts gets sale.order.alert existing records.
func (c *Client) GetSaleOrderAlerts(ids []int64) (*SaleOrderAlerts, error) {
	soas := &SaleOrderAlerts{}
	if err := c.Read(SaleOrderAlertModel, ids, nil, soas); err != nil {
		return nil, err
	}
	return soas, nil
}

// FindSaleOrderAlert finds sale.order.alert record by querying it with criteria.
func (c *Client) FindSaleOrderAlert(criteria *Criteria) (*SaleOrderAlert, error) {
	soas := &SaleOrderAlerts{}
	if err := c.SearchRead(SaleOrderAlertModel, criteria, NewOptions().Limit(1), soas); err != nil {
		return nil, err
	}
	if soas != nil && len(*soas) > 0 {
		return &((*soas)[0]), nil
	}
	return nil, fmt.Errorf("sale.order.alert was not found with criteria %v", criteria)
}

// FindSaleOrderAlerts finds sale.order.alert records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderAlerts(criteria *Criteria, options *Options) (*SaleOrderAlerts, error) {
	soas := &SaleOrderAlerts{}
	if err := c.SearchRead(SaleOrderAlertModel, criteria, options, soas); err != nil {
		return nil, err
	}
	return soas, nil
}

// FindSaleOrderAlertIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSaleOrderAlertIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SaleOrderAlertModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSaleOrderAlertId finds record id by querying it with criteria.
func (c *Client) FindSaleOrderAlertId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SaleOrderAlertModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sale.order.alert was not found with criteria %v and options %v", criteria, options)
}
