package odoo

import (
	"fmt"
)

// HrExpense represents hr.expense model.
type HrExpense struct {
	LastUpdate                  *Time       `xmlrpc:"__last_update,omptempty"`
	AccountId                   *Many2One   `xmlrpc:"account_id,omptempty"`
	AccountingDate              *Time       `xmlrpc:"accounting_date,omptempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	AmountResidual              *Float      `xmlrpc:"amount_residual,omptempty"`
	AmountTax                   *Float      `xmlrpc:"amount_tax,omptempty"`
	AmountTaxCompany            *Float      `xmlrpc:"amount_tax_company,omptempty"`
	AnalyticDistribution        interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticDistributionSearch  interface{} `xmlrpc:"analytic_distribution_search,omptempty"`
	AnalyticPrecision           *Int        `xmlrpc:"analytic_precision,omptempty"`
	ApprovedBy                  *Many2One   `xmlrpc:"approved_by,omptempty"`
	ApprovedOn                  *Time       `xmlrpc:"approved_on,omptempty"`
	AttachmentNumber            *Int        `xmlrpc:"attachment_number,omptempty"`
	CanBeReinvoiced             *Bool       `xmlrpc:"can_be_reinvoiced,omptempty"`
	CompanyCurrencyId           *Many2One   `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One   `xmlrpc:"currency_id,omptempty"`
	CurrencyRate                *Float      `xmlrpc:"currency_rate,omptempty"`
	Date                        *Time       `xmlrpc:"date,omptempty"`
	Description                 *String     `xmlrpc:"description,omptempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omptempty"`
	DuplicateExpenseIds         *Relation   `xmlrpc:"duplicate_expense_ids,omptempty"`
	EmployeeId                  *Many2One   `xmlrpc:"employee_id,omptempty"`
	ExtractCanShowResendButton  *Bool       `xmlrpc:"extract_can_show_resend_button,omptempty"`
	ExtractCanShowSendButton    *Bool       `xmlrpc:"extract_can_show_send_button,omptempty"`
	ExtractErrorMessage         *String     `xmlrpc:"extract_error_message,omptempty"`
	ExtractRemoteId             *Int        `xmlrpc:"extract_remote_id,omptempty"`
	ExtractState                *Selection  `xmlrpc:"extract_state,omptempty"`
	ExtractStatusCode           *Int        `xmlrpc:"extract_status_code,omptempty"`
	ExtractWordIds              *Relation   `xmlrpc:"extract_word_ids,omptempty"`
	FailedMessageIds            *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omptempty"`
	Id                          *Int        `xmlrpc:"id,omptempty"`
	IsEditable                  *Bool       `xmlrpc:"is_editable,omptempty"`
	IsRefEditable               *Bool       `xmlrpc:"is_ref_editable,omptempty"`
	IsRefused                   *Bool       `xmlrpc:"is_refused,omptempty"`
	LabelConvertRate            *String     `xmlrpc:"label_convert_rate,omptempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent              *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String     `xmlrpc:"name,omptempty"`
	PaymentMode                 *Selection  `xmlrpc:"payment_mode,omptempty"`
	PredictedCategory           *Selection  `xmlrpc:"predicted_category,omptempty"`
	ProductDescription          *String     `xmlrpc:"product_description,omptempty"`
	ProductHasCost              *Bool       `xmlrpc:"product_has_cost,omptempty"`
	ProductHasTax               *Bool       `xmlrpc:"product_has_tax,omptempty"`
	ProductId                   *Many2One   `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId        *Many2One   `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomId                *Many2One   `xmlrpc:"product_uom_id,omptempty"`
	Quantity                    *Float      `xmlrpc:"quantity,omptempty"`
	Reference                   *String     `xmlrpc:"reference,omptempty"`
	SaleOrderId                 *Many2One   `xmlrpc:"sale_order_id,omptempty"`
	SameCurrency                *Bool       `xmlrpc:"same_currency,omptempty"`
	Sample                      *Bool       `xmlrpc:"sample,omptempty"`
	SheetId                     *Many2One   `xmlrpc:"sheet_id,omptempty"`
	SheetIsEditable             *Bool       `xmlrpc:"sheet_is_editable,omptempty"`
	State                       *Selection  `xmlrpc:"state,omptempty"`
	StateProcessed              *Bool       `xmlrpc:"state_processed,omptempty"`
	TaxIds                      *Relation   `xmlrpc:"tax_ids,omptempty"`
	TotalAmount                 *Float      `xmlrpc:"total_amount,omptempty"`
	TotalAmountCompany          *Float      `xmlrpc:"total_amount_company,omptempty"`
	UnitAmount                  *Float      `xmlrpc:"unit_amount,omptempty"`
	UnitAmountDisplay           *Float      `xmlrpc:"unit_amount_display,omptempty"`
	UntaxedAmount               *Float      `xmlrpc:"untaxed_amount,omptempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// HrExpenses represents array of hr.expense model.
type HrExpenses []HrExpense

// HrExpenseModel is the odoo model name.
const HrExpenseModel = "hr.expense"

// Many2One convert HrExpense to *Many2One.
func (he *HrExpense) Many2One() *Many2One {
	return NewMany2One(he.Id.Get(), "")
}

// CreateHrExpense creates a new hr.expense model and returns its id.
func (c *Client) CreateHrExpense(he *HrExpense) (int64, error) {
	ids, err := c.CreateHrExpenses([]*HrExpense{he})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrExpense creates a new hr.expense model and returns its id.
func (c *Client) CreateHrExpenses(hes []*HrExpense) ([]int64, error) {
	var vv []interface{}
	for _, v := range hes {
		vv = append(vv, v)
	}
	return c.Create(HrExpenseModel, vv)
}

// UpdateHrExpense updates an existing hr.expense record.
func (c *Client) UpdateHrExpense(he *HrExpense) error {
	return c.UpdateHrExpenses([]int64{he.Id.Get()}, he)
}

// UpdateHrExpenses updates existing hr.expense records.
// All records (represented by ids) will be updated by he values.
func (c *Client) UpdateHrExpenses(ids []int64, he *HrExpense) error {
	return c.Update(HrExpenseModel, ids, he)
}

// DeleteHrExpense deletes an existing hr.expense record.
func (c *Client) DeleteHrExpense(id int64) error {
	return c.DeleteHrExpenses([]int64{id})
}

// DeleteHrExpenses deletes existing hr.expense records.
func (c *Client) DeleteHrExpenses(ids []int64) error {
	return c.Delete(HrExpenseModel, ids)
}

// GetHrExpense gets hr.expense existing record.
func (c *Client) GetHrExpense(id int64) (*HrExpense, error) {
	hes, err := c.GetHrExpenses([]int64{id})
	if err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.expense not found", id)
}

// GetHrExpenses gets hr.expense existing records.
func (c *Client) GetHrExpenses(ids []int64) (*HrExpenses, error) {
	hes := &HrExpenses{}
	if err := c.Read(HrExpenseModel, ids, nil, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrExpense finds hr.expense record by querying it with criteria.
func (c *Client) FindHrExpense(criteria *Criteria) (*HrExpense, error) {
	hes := &HrExpenses{}
	if err := c.SearchRead(HrExpenseModel, criteria, NewOptions().Limit(1), hes); err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("hr.expense was not found with criteria %v", criteria)
}

// FindHrExpenses finds hr.expense records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenses(criteria *Criteria, options *Options) (*HrExpenses, error) {
	hes := &HrExpenses{}
	if err := c.SearchRead(HrExpenseModel, criteria, options, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrExpenseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrExpenseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrExpenseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrExpenseId finds record id by querying it with criteria.
func (c *Client) FindHrExpenseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrExpenseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.expense was not found with criteria %v and options %v", criteria, options)
}
