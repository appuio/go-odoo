package odoo

import (
	"fmt"
)

// HrContract represents hr.contract model.
type HrContract struct {
	LastUpdate                        *Time      `xmlrpc:"__last_update,omptempty"`
	Active                            *Bool      `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId           *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline              *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration       *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon             *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                       *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                     *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                   *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                  *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                    *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                    *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AnalyticAccountId                 *Many2One  `xmlrpc:"analytic_account_id,omptempty"`
	CalendarChanged                   *Bool      `xmlrpc:"calendar_changed,omptempty"`
	CalendarMismatch                  *Bool      `xmlrpc:"calendar_mismatch,omptempty"`
	CompanyCountryId                  *Many2One  `xmlrpc:"company_country_id,omptempty"`
	CompanyId                         *Many2One  `xmlrpc:"company_id,omptempty"`
	ContractTypeId                    *Many2One  `xmlrpc:"contract_type_id,omptempty"`
	ContractWage                      *Float     `xmlrpc:"contract_wage,omptempty"`
	CountryCode                       *String    `xmlrpc:"country_code,omptempty"`
	CreateDate                        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                        *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateEnd                           *Time      `xmlrpc:"date_end,omptempty"`
	DateGeneratedFrom                 *Time      `xmlrpc:"date_generated_from,omptempty"`
	DateGeneratedTo                   *Time      `xmlrpc:"date_generated_to,omptempty"`
	DateStart                         *Time      `xmlrpc:"date_start,omptempty"`
	DepartmentId                      *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                       *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId                        *Many2One  `xmlrpc:"employee_id,omptempty"`
	FailedMessageIds                  *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FirstContractDate                 *Time      `xmlrpc:"first_contract_date,omptempty"`
	FullTimeRequiredHours             *Float     `xmlrpc:"full_time_required_hours,omptempty"`
	HasMessage                        *Bool      `xmlrpc:"has_message,omptempty"`
	HourlyWage                        *Float     `xmlrpc:"hourly_wage,omptempty"`
	HoursPerWeek                      *Float     `xmlrpc:"hours_per_week,omptempty"`
	HrResponsibleId                   *Many2One  `xmlrpc:"hr_responsible_id,omptempty"`
	Id                                *Int       `xmlrpc:"id,omptempty"`
	IsFulltime                        *Bool      `xmlrpc:"is_fulltime,omptempty"`
	JobId                             *Many2One  `xmlrpc:"job_id,omptempty"`
	KanbanState                       *Selection `xmlrpc:"kanban_state,omptempty"`
	LastGenerationDate                *Time      `xmlrpc:"last_generation_date,omptempty"`
	MessageAttachmentCount            *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                    *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                   *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter            *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                        *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                 *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId           *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                 *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter          *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                 *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline            *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                              *String    `xmlrpc:"name,omptempty"`
	Notes                             *String    `xmlrpc:"notes,omptempty"`
	PayslipsCount                     *Int       `xmlrpc:"payslips_count,omptempty"`
	PermitNo                          *String    `xmlrpc:"permit_no,omptempty"`
	ResourceCalendarId                *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	SchedulePay                       *Selection `xmlrpc:"schedule_pay,omptempty"`
	StandardCalendarId                *Many2One  `xmlrpc:"standard_calendar_id,omptempty"`
	State                             *Selection `xmlrpc:"state,omptempty"`
	StructureTypeId                   *Many2One  `xmlrpc:"structure_type_id,omptempty"`
	TimeCredit                        *Bool      `xmlrpc:"time_credit,omptempty"`
	TimeCreditTypeId                  *Many2One  `xmlrpc:"time_credit_type_id,omptempty"`
	TrialDateEnd                      *Time      `xmlrpc:"trial_date_end,omptempty"`
	VisaExpire                        *Time      `xmlrpc:"visa_expire,omptempty"`
	VisaNo                            *String    `xmlrpc:"visa_no,omptempty"`
	Wage                              *Float     `xmlrpc:"wage,omptempty"`
	WageType                          *Selection `xmlrpc:"wage_type,omptempty"`
	WebsiteMessageIds                 *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WorkEntrySource                   *Selection `xmlrpc:"work_entry_source,omptempty"`
	WorkTimeRate                      *Float     `xmlrpc:"work_time_rate,omptempty"`
	WriteDate                         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                          *Many2One  `xmlrpc:"write_uid,omptempty"`
	XAvatarImage                      *String    `xmlrpc:"x_avatar_image,omptempty"`
	XStudioCalcQst                    *Float     `xmlrpc:"x_studio_calc_qst,omptempty"`
	XStudioCalculatedWithholdingTax   *Float     `xmlrpc:"x_studio_calculated_withholding_tax,omptempty"`
	XStudioCharFieldM9OEv             *String    `xmlrpc:"x_studio_char_field_m9oEv,omptempty"`
	XStudioDateFieldQYZi3             *Time      `xmlrpc:"x_studio_date_field_qYZi3,omptempty"`
	XStudioEmployeeQstStateCode       *String    `xmlrpc:"x_studio_employee_qst_state_code,omptempty"`
	XStudioEmployeeWithholdingTaxCode *Selection `xmlrpc:"x_studio_employee_withholding_tax_code,omptempty"`
	XStudioEo                         *Float     `xmlrpc:"x_studio_eo,omptempty"`
	XStudioFloatFieldSs8Qe            *Float     `xmlrpc:"x_studio_float_field_Ss8Qe,omptempty"`
	XStudioFloatFieldLBI36            *Float     `xmlrpc:"x_studio_float_field_lBI36,omptempty"`
	XStudioFloatFieldTvyQz            *Float     `xmlrpc:"x_studio_float_field_tvyQz,omptempty"`
	XStudioFloatFieldXpZEb            *Float     `xmlrpc:"x_studio_float_field_xpZEb,omptempty"`
	XStudioKizu                       *Float     `xmlrpc:"x_studio_kizu,omptempty"`
	XStudioKtgUvg                     *Float     `xmlrpc:"x_studio_ktg_uvg,omptempty"`
	XStudioLookupVat                  *Float     `xmlrpc:"x_studio_lookup_vat,omptempty"`
	XStudioLookupVat1                 *Float     `xmlrpc:"x_studio_lookup_vat_1,omptempty"`
	XStudioMany2ManyFieldKEXWF        *Relation  `xmlrpc:"x_studio_many2many_field_KEXWF,omptempty"`
	XStudioMany2ManyFieldSd6Vz        *Relation  `xmlrpc:"x_studio_many2many_field_Sd6vz,omptempty"`
	XStudioMany2ManyFieldUGMPj        *Relation  `xmlrpc:"x_studio_many2many_field_UGMPj,omptempty"`
	XStudioMany2ManyFieldXPaJe        *Relation  `xmlrpc:"x_studio_many2many_field_XPaJe,omptempty"`
	XStudioMany2ManyFieldAqEY7        *Relation  `xmlrpc:"x_studio_many2many_field_aqEY7,omptempty"`
	XStudioMany2ManyFieldHUsl3        *Relation  `xmlrpc:"x_studio_many2many_field_hUsl3,omptempty"`
	XStudioMany2OneField396Q1         *Many2One  `xmlrpc:"x_studio_many2one_field_396q1,omptempty"`
	XStudioMany2OneFieldKQLEy         *Many2One  `xmlrpc:"x_studio_many2one_field_KQLEy,omptempty"`
	XStudioMany2OneFieldYVljh         *Many2One  `xmlrpc:"x_studio_many2one_field_YVljh,omptempty"`
	XStudioMany2OneFieldC7Zgf         *Many2One  `xmlrpc:"x_studio_many2one_field_c7Zgf,omptempty"`
	XStudioName                       *String    `xmlrpc:"x_studio_name,omptempty"`
	XStudioPensionFund                *Float     `xmlrpc:"x_studio_pension_fund,omptempty"`
	XStudioQst                        *Float     `xmlrpc:"x_studio_qst,omptempty"`
	XStudioQstChf                     *Float     `xmlrpc:"x_studio_qst_chf,omptempty"`
	XStudioQstName                    *String    `xmlrpc:"x_studio_qst_name,omptempty"`
	XStudioQstNameNew                 *String    `xmlrpc:"x_studio_qst_name_new,omptempty"`
	XStudioRelatedField0W4TZ          *Time      `xmlrpc:"x_studio_related_field_0W4tZ,omptempty"`
	XStudioRelatedField3XeJ5          *Many2One  `xmlrpc:"x_studio_related_field_3xeJ5,omptempty"`
	XStudioRelatedFieldC285H          *String    `xmlrpc:"x_studio_related_field_C285h,omptempty"`
	XStudioRelatedFieldGI8Ed          *String    `xmlrpc:"x_studio_related_field_GI8ed,omptempty"`
	XStudioRelatedFieldHvzAA          *Bool      `xmlrpc:"x_studio_related_field_HvzAA,omptempty"`
	XStudioRelatedFieldLTTSI          *String    `xmlrpc:"x_studio_related_field_LTTSI,omptempty"`
	XStudioRelatedFieldOBblW          *String    `xmlrpc:"x_studio_related_field_OBblW,omptempty"`
	XStudioRelatedFieldTcSNP          *String    `xmlrpc:"x_studio_related_field_TcSNP,omptempty"`
	XStudioRelatedFieldJ5L55          *String    `xmlrpc:"x_studio_related_field_j5L55,omptempty"`
	XStudioRelatedFieldJ7Whi          *String    `xmlrpc:"x_studio_related_field_j7Whi,omptempty"`
	XStudioRelatedFieldJLsUK          *Float     `xmlrpc:"x_studio_related_field_jLsUK,omptempty"`
	XStudioRelatedFieldNhYk0          *String    `xmlrpc:"x_studio_related_field_nhYk0,omptempty"`
	XStudioRelatedFieldRxFhN          *Relation  `xmlrpc:"x_studio_related_field_rxFhN,omptempty"`
	XStudioRelatedFieldTtwXd          *Float     `xmlrpc:"x_studio_related_field_ttwXd,omptempty"`
	XStudioWage                       *Float     `xmlrpc:"x_studio_wage,omptempty"`
}

// HrContracts represents array of hr.contract model.
type HrContracts []HrContract

// HrContractModel is the odoo model name.
const HrContractModel = "hr.contract"

// Many2One convert HrContract to *Many2One.
func (hc *HrContract) Many2One() *Many2One {
	return NewMany2One(hc.Id.Get(), "")
}

// CreateHrContract creates a new hr.contract model and returns its id.
func (c *Client) CreateHrContract(hc *HrContract) (int64, error) {
	ids, err := c.CreateHrContracts([]*HrContract{hc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrContract creates a new hr.contract model and returns its id.
func (c *Client) CreateHrContracts(hcs []*HrContract) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcs {
		vv = append(vv, v)
	}
	return c.Create(HrContractModel, vv)
}

// UpdateHrContract updates an existing hr.contract record.
func (c *Client) UpdateHrContract(hc *HrContract) error {
	return c.UpdateHrContracts([]int64{hc.Id.Get()}, hc)
}

// UpdateHrContracts updates existing hr.contract records.
// All records (represented by ids) will be updated by hc values.
func (c *Client) UpdateHrContracts(ids []int64, hc *HrContract) error {
	return c.Update(HrContractModel, ids, hc)
}

// DeleteHrContract deletes an existing hr.contract record.
func (c *Client) DeleteHrContract(id int64) error {
	return c.DeleteHrContracts([]int64{id})
}

// DeleteHrContracts deletes existing hr.contract records.
func (c *Client) DeleteHrContracts(ids []int64) error {
	return c.Delete(HrContractModel, ids)
}

// GetHrContract gets hr.contract existing record.
func (c *Client) GetHrContract(id int64) (*HrContract, error) {
	hcs, err := c.GetHrContracts([]int64{id})
	if err != nil {
		return nil, err
	}
	if hcs != nil && len(*hcs) > 0 {
		return &((*hcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.contract not found", id)
}

// GetHrContracts gets hr.contract existing records.
func (c *Client) GetHrContracts(ids []int64) (*HrContracts, error) {
	hcs := &HrContracts{}
	if err := c.Read(HrContractModel, ids, nil, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrContract finds hr.contract record by querying it with criteria.
func (c *Client) FindHrContract(criteria *Criteria) (*HrContract, error) {
	hcs := &HrContracts{}
	if err := c.SearchRead(HrContractModel, criteria, NewOptions().Limit(1), hcs); err != nil {
		return nil, err
	}
	if hcs != nil && len(*hcs) > 0 {
		return &((*hcs)[0]), nil
	}
	return nil, fmt.Errorf("hr.contract was not found with criteria %v", criteria)
}

// FindHrContracts finds hr.contract records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContracts(criteria *Criteria, options *Options) (*HrContracts, error) {
	hcs := &HrContracts{}
	if err := c.SearchRead(HrContractModel, criteria, options, hcs); err != nil {
		return nil, err
	}
	return hcs, nil
}

// FindHrContractIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrContractModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrContractId finds record id by querying it with criteria.
func (c *Client) FindHrContractId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.contract was not found with criteria %v and options %v", criteria, options)
}
