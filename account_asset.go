package odoo

import (
	"fmt"
)

// AccountAsset represents account.asset model.
type AccountAsset struct {
	LastUpdate                     *Time       `xmlrpc:"__last_update,omptempty"`
	AccountAssetId                 *Many2One   `xmlrpc:"account_asset_id,omptempty"`
	AccountDepreciationExpenseId   *Many2One   `xmlrpc:"account_depreciation_expense_id,omptempty"`
	AccountDepreciationId          *Many2One   `xmlrpc:"account_depreciation_id,omptempty"`
	AccountType                    *Selection  `xmlrpc:"account_type,omptempty"`
	AcquisitionDate                *Time       `xmlrpc:"acquisition_date,omptempty"`
	Active                         *Bool       `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId        *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline           *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration    *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon          *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                    *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState                  *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon               *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                 *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                 *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	AlreadyDepreciatedAmountImport *Float      `xmlrpc:"already_depreciated_amount_import,omptempty"`
	AnalyticDistribution           interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticDistributionSearch     interface{} `xmlrpc:"analytic_distribution_search,omptempty"`
	AnalyticPrecision              *Int        `xmlrpc:"analytic_precision,omptempty"`
	AssetLifetimeDays              *Int        `xmlrpc:"asset_lifetime_days,omptempty"`
	AssetPausedDays                *Float      `xmlrpc:"asset_paused_days,omptempty"`
	AssetType                      *Selection  `xmlrpc:"asset_type,omptempty"`
	BookValue                      *Float      `xmlrpc:"book_value,omptempty"`
	ChildrenIds                    *Relation   `xmlrpc:"children_ids,omptempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate                     *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One   `xmlrpc:"create_uid,omptempty"`
	CurrencyId                     *Many2One   `xmlrpc:"currency_id,omptempty"`
	DepreciationEntriesCount       *Int        `xmlrpc:"depreciation_entries_count,omptempty"`
	DepreciationMoveIds            *Relation   `xmlrpc:"depreciation_move_ids,omptempty"`
	DisplayAccountAssetId          *Bool       `xmlrpc:"display_account_asset_id,omptempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omptempty"`
	DisposalDate                   *Time       `xmlrpc:"disposal_date,omptempty"`
	FailedMessageIds               *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	GrossIncreaseCount             *Int        `xmlrpc:"gross_increase_count,omptempty"`
	GrossIncreaseValue             *Float      `xmlrpc:"gross_increase_value,omptempty"`
	HasMessage                     *Bool       `xmlrpc:"has_message,omptempty"`
	Id                             *Int        `xmlrpc:"id,omptempty"`
	JournalId                      *Many2One   `xmlrpc:"journal_id,omptempty"`
	MessageAttachmentCount         *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                 *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds             *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter         *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError             *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                     *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower              *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId        *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction              *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter       *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds              *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	Method                         *Selection  `xmlrpc:"method,omptempty"`
	MethodNumber                   *Int        `xmlrpc:"method_number,omptempty"`
	MethodPeriod                   *Selection  `xmlrpc:"method_period,omptempty"`
	MethodProgressFactor           *Float      `xmlrpc:"method_progress_factor,omptempty"`
	ModelId                        *Many2One   `xmlrpc:"model_id,omptempty"`
	MyActivityDateDeadline         *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                           *String     `xmlrpc:"name,omptempty"`
	NonDeductibleTaxValue          *Float      `xmlrpc:"non_deductible_tax_value,omptempty"`
	OriginalMoveLineIds            *Relation   `xmlrpc:"original_move_line_ids,omptempty"`
	OriginalValue                  *Float      `xmlrpc:"original_value,omptempty"`
	ParentId                       *Many2One   `xmlrpc:"parent_id,omptempty"`
	PausedProrataDate              *Time       `xmlrpc:"paused_prorata_date,omptempty"`
	ProrataComputationType         *Selection  `xmlrpc:"prorata_computation_type,omptempty"`
	ProrataDate                    *Time       `xmlrpc:"prorata_date,omptempty"`
	RelatedPurchaseValue           *Float      `xmlrpc:"related_purchase_value,omptempty"`
	SalvageValue                   *Float      `xmlrpc:"salvage_value,omptempty"`
	State                          *Selection  `xmlrpc:"state,omptempty"`
	TotalDepreciableValue          *Float      `xmlrpc:"total_depreciable_value,omptempty"`
	TotalDepreciationEntriesCount  *Int        `xmlrpc:"total_depreciation_entries_count,omptempty"`
	ValueResidual                  *Float      `xmlrpc:"value_residual,omptempty"`
	WebsiteMessageIds              *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                      *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One   `xmlrpc:"write_uid,omptempty"`
	XMethodPeriod                  *Selection  `xmlrpc:"x_method_period,omptempty"`
	XStudioAssetValue              *Float      `xmlrpc:"x_studio_asset_value,omptempty"`
	XStudioCalculateValueForDate   *Time       `xmlrpc:"x_studio_calculate_value_for_date,omptempty"`
	XStudioCalculatedValue         *Float      `xmlrpc:"x_studio_calculated_value,omptempty"`
	XStudioCalculatedValue1        *Float      `xmlrpc:"x_studio_calculated_value_1,omptempty"`
	XStudioCharFieldVqjsX          *String     `xmlrpc:"x_studio_char_field_vqjsX,omptempty"`
	XStudioDatetimeFieldRAm51      *Time       `xmlrpc:"x_studio_datetime_field_RAm51,omptempty"`
	XStudioDatetimeNow             *Time       `xmlrpc:"x_studio_datetime_now,omptempty"`
	XStudioDaysBetween             *Int        `xmlrpc:"x_studio_days_between,omptempty"`
	XStudioDaysInThreeYears        *Int        `xmlrpc:"x_studio_days_in_three_years,omptempty"`
	XStudioDaysSinceAcquisition    *Int        `xmlrpc:"x_studio_days_since_acquisition,omptempty"`
	XStudioLocation                *Selection  `xmlrpc:"x_studio_location,omptempty"`
	XStudioMany2OneField3VCYY      *Many2One   `xmlrpc:"x_studio_many2one_field_3vCYY,omptempty"`
	XStudioRelatedFieldRq4ST       *String     `xmlrpc:"x_studio_related_field_Rq4sT,omptempty"`
	XStudioRelatedFieldVsrjI       *String     `xmlrpc:"x_studio_related_field_VsrjI,omptempty"`
	XStudioRelatedFieldX9ZR5       *String     `xmlrpc:"x_studio_related_field_X9ZR5,omptempty"`
	XStudioSelectionFieldOcRSQ     *Selection  `xmlrpc:"x_studio_selection_field_OcRSQ,omptempty"`
	XStudioSerial                  *String     `xmlrpc:"x_studio_serial_,omptempty"`
	XStudioToday                   *Time       `xmlrpc:"x_studio_today,omptempty"`
	XStudioTodayDate1              *Time       `xmlrpc:"x_studio_today_date_1,omptempty"`
	XStudioVshnAssetNumber1        *Int        `xmlrpc:"x_studio_vshn_asset_number_1,omptempty"`
	XStudioWarranty                *Int        `xmlrpc:"x_studio_warranty,omptempty"`
	XStudioWarrantyAssets          *Int        `xmlrpc:"x_studio_warranty_assets,omptempty"`
	XStudioWarrantyMonths          *Int        `xmlrpc:"x_studio_warranty_months,omptempty"`
	XStudioWarrantyUntil           *Time       `xmlrpc:"x_studio_warranty_until,omptempty"`
}

// AccountAssets represents array of account.asset model.
type AccountAssets []AccountAsset

// AccountAssetModel is the odoo model name.
const AccountAssetModel = "account.asset"

// Many2One convert AccountAsset to *Many2One.
func (aa *AccountAsset) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAccountAsset creates a new account.asset model and returns its id.
func (c *Client) CreateAccountAsset(aa *AccountAsset) (int64, error) {
	ids, err := c.CreateAccountAssets([]*AccountAsset{aa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAsset creates a new account.asset model and returns its id.
func (c *Client) CreateAccountAssets(aas []*AccountAsset) ([]int64, error) {
	var vv []interface{}
	for _, v := range aas {
		vv = append(vv, v)
	}
	return c.Create(AccountAssetModel, vv)
}

// UpdateAccountAsset updates an existing account.asset record.
func (c *Client) UpdateAccountAsset(aa *AccountAsset) error {
	return c.UpdateAccountAssets([]int64{aa.Id.Get()}, aa)
}

// UpdateAccountAssets updates existing account.asset records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAccountAssets(ids []int64, aa *AccountAsset) error {
	return c.Update(AccountAssetModel, ids, aa)
}

// DeleteAccountAsset deletes an existing account.asset record.
func (c *Client) DeleteAccountAsset(id int64) error {
	return c.DeleteAccountAssets([]int64{id})
}

// DeleteAccountAssets deletes existing account.asset records.
func (c *Client) DeleteAccountAssets(ids []int64) error {
	return c.Delete(AccountAssetModel, ids)
}

// GetAccountAsset gets account.asset existing record.
func (c *Client) GetAccountAsset(id int64) (*AccountAsset, error) {
	aas, err := c.GetAccountAssets([]int64{id})
	if err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.asset not found", id)
}

// GetAccountAssets gets account.asset existing records.
func (c *Client) GetAccountAssets(ids []int64) (*AccountAssets, error) {
	aas := &AccountAssets{}
	if err := c.Read(AccountAssetModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAsset finds account.asset record by querying it with criteria.
func (c *Client) FindAccountAsset(criteria *Criteria) (*AccountAsset, error) {
	aas := &AccountAssets{}
	if err := c.SearchRead(AccountAssetModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("account.asset was not found with criteria %v", criteria)
}

// FindAccountAssets finds account.asset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssets(criteria *Criteria, options *Options) (*AccountAssets, error) {
	aas := &AccountAssets{}
	if err := c.SearchRead(AccountAssetModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAssetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAssetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAssetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAssetId finds record id by querying it with criteria.
func (c *Client) FindAccountAssetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAssetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.asset was not found with criteria %v and options %v", criteria, options)
}
