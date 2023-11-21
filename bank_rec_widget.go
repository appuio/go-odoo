package odoo

import (
	"fmt"
)

// BankRecWidget represents bank.rec.widget model.
type BankRecWidget struct {
	LastUpdate                      *Time       `xmlrpc:"__last_update,omptempty"`
	AmlsWidget                      *String     `xmlrpc:"amls_widget,omptempty"`
	AnalyticDistribution            interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticDistributionSearch      interface{} `xmlrpc:"analytic_distribution_search,omptempty"`
	AnalyticPrecision               *Int        `xmlrpc:"analytic_precision,omptempty"`
	BatchPaymentsWidget             *String     `xmlrpc:"batch_payments_widget,omptempty"`
	CompanyCurrencyId               *Many2One   `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                       *Many2One   `xmlrpc:"company_id,omptempty"`
	DisplayName                     *String     `xmlrpc:"display_name,omptempty"`
	FormAccountId                   *Many2One   `xmlrpc:"form_account_id,omptempty"`
	FormAmountCurrency              *Float      `xmlrpc:"form_amount_currency,omptempty"`
	FormBalance                     *Float      `xmlrpc:"form_balance,omptempty"`
	FormCurrencyId                  *Many2One   `xmlrpc:"form_currency_id,omptempty"`
	FormDate                        *Time       `xmlrpc:"form_date,omptempty"`
	FormExtraText                   *String     `xmlrpc:"form_extra_text,omptempty"`
	FormFlag                        *String     `xmlrpc:"form_flag,omptempty"`
	FormForceNegativeSign           *Bool       `xmlrpc:"form_force_negative_sign,omptempty"`
	FormIndex                       *String     `xmlrpc:"form_index,omptempty"`
	FormName                        *String     `xmlrpc:"form_name,omptempty"`
	FormPartnerCurrencyId           *Many2One   `xmlrpc:"form_partner_currency_id,omptempty"`
	FormPartnerId                   *Many2One   `xmlrpc:"form_partner_id,omptempty"`
	FormPartnerPayableAccountId     *Many2One   `xmlrpc:"form_partner_payable_account_id,omptempty"`
	FormPartnerPayableAmount        *Float      `xmlrpc:"form_partner_payable_amount,omptempty"`
	FormPartnerReceivableAccountId  *Many2One   `xmlrpc:"form_partner_receivable_account_id,omptempty"`
	FormPartnerReceivableAmount     *Float      `xmlrpc:"form_partner_receivable_amount,omptempty"`
	FormSingleCurrencyMode          *Bool       `xmlrpc:"form_single_currency_mode,omptempty"`
	FormSuggestAmountCurrency       *Float      `xmlrpc:"form_suggest_amount_currency,omptempty"`
	FormSuggestBalance              *Float      `xmlrpc:"form_suggest_balance,omptempty"`
	FormTaxIds                      *Relation   `xmlrpc:"form_tax_ids,omptempty"`
	Id                              *Int        `xmlrpc:"id,omptempty"`
	JournalCurrencyId               *Many2One   `xmlrpc:"journal_currency_id,omptempty"`
	LineIds                         *Relation   `xmlrpc:"line_ids,omptempty"`
	LinesWidget                     *String     `xmlrpc:"lines_widget,omptempty"`
	MatchedSaleOrderIds             *Relation   `xmlrpc:"matched_sale_order_ids,omptempty"`
	MatchingRulesAllowAutoReconcile *Bool       `xmlrpc:"matching_rules_allow_auto_reconcile,omptempty"`
	MoveId                          *Many2One   `xmlrpc:"move_id,omptempty"`
	NextActionTodo                  *String     `xmlrpc:"next_action_todo,omptempty"`
	PartnerId                       *Many2One   `xmlrpc:"partner_id,omptempty"`
	RecoModelsWidget                *String     `xmlrpc:"reco_models_widget,omptempty"`
	SelectedAmlIds                  *Relation   `xmlrpc:"selected_aml_ids,omptempty"`
	SelectedBatchPaymentIds         *Relation   `xmlrpc:"selected_batch_payment_ids,omptempty"`
	StLineId                        *Many2One   `xmlrpc:"st_line_id,omptempty"`
	StLineIsReconciled              *Bool       `xmlrpc:"st_line_is_reconciled,omptempty"`
	State                           *Selection  `xmlrpc:"state,omptempty"`
	ToCheck                         *Bool       `xmlrpc:"to_check,omptempty"`
	TodoCommand                     *String     `xmlrpc:"todo_command,omptempty"`
	TransactionCurrencyId           *Many2One   `xmlrpc:"transaction_currency_id,omptempty"`
}

// BankRecWidgets represents array of bank.rec.widget model.
type BankRecWidgets []BankRecWidget

// BankRecWidgetModel is the odoo model name.
const BankRecWidgetModel = "bank.rec.widget"

// Many2One convert BankRecWidget to *Many2One.
func (brw *BankRecWidget) Many2One() *Many2One {
	return NewMany2One(brw.Id.Get(), "")
}

// CreateBankRecWidget creates a new bank.rec.widget model and returns its id.
func (c *Client) CreateBankRecWidget(brw *BankRecWidget) (int64, error) {
	ids, err := c.CreateBankRecWidgets([]*BankRecWidget{brw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBankRecWidget creates a new bank.rec.widget model and returns its id.
func (c *Client) CreateBankRecWidgets(brws []*BankRecWidget) ([]int64, error) {
	var vv []interface{}
	for _, v := range brws {
		vv = append(vv, v)
	}
	return c.Create(BankRecWidgetModel, vv)
}

// UpdateBankRecWidget updates an existing bank.rec.widget record.
func (c *Client) UpdateBankRecWidget(brw *BankRecWidget) error {
	return c.UpdateBankRecWidgets([]int64{brw.Id.Get()}, brw)
}

// UpdateBankRecWidgets updates existing bank.rec.widget records.
// All records (represented by ids) will be updated by brw values.
func (c *Client) UpdateBankRecWidgets(ids []int64, brw *BankRecWidget) error {
	return c.Update(BankRecWidgetModel, ids, brw)
}

// DeleteBankRecWidget deletes an existing bank.rec.widget record.
func (c *Client) DeleteBankRecWidget(id int64) error {
	return c.DeleteBankRecWidgets([]int64{id})
}

// DeleteBankRecWidgets deletes existing bank.rec.widget records.
func (c *Client) DeleteBankRecWidgets(ids []int64) error {
	return c.Delete(BankRecWidgetModel, ids)
}

// GetBankRecWidget gets bank.rec.widget existing record.
func (c *Client) GetBankRecWidget(id int64) (*BankRecWidget, error) {
	brws, err := c.GetBankRecWidgets([]int64{id})
	if err != nil {
		return nil, err
	}
	if brws != nil && len(*brws) > 0 {
		return &((*brws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of bank.rec.widget not found", id)
}

// GetBankRecWidgets gets bank.rec.widget existing records.
func (c *Client) GetBankRecWidgets(ids []int64) (*BankRecWidgets, error) {
	brws := &BankRecWidgets{}
	if err := c.Read(BankRecWidgetModel, ids, nil, brws); err != nil {
		return nil, err
	}
	return brws, nil
}

// FindBankRecWidget finds bank.rec.widget record by querying it with criteria.
func (c *Client) FindBankRecWidget(criteria *Criteria) (*BankRecWidget, error) {
	brws := &BankRecWidgets{}
	if err := c.SearchRead(BankRecWidgetModel, criteria, NewOptions().Limit(1), brws); err != nil {
		return nil, err
	}
	if brws != nil && len(*brws) > 0 {
		return &((*brws)[0]), nil
	}
	return nil, fmt.Errorf("bank.rec.widget was not found with criteria %v", criteria)
}

// FindBankRecWidgets finds bank.rec.widget records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgets(criteria *Criteria, options *Options) (*BankRecWidgets, error) {
	brws := &BankRecWidgets{}
	if err := c.SearchRead(BankRecWidgetModel, criteria, options, brws); err != nil {
		return nil, err
	}
	return brws, nil
}

// FindBankRecWidgetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BankRecWidgetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBankRecWidgetId finds record id by querying it with criteria.
func (c *Client) FindBankRecWidgetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BankRecWidgetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("bank.rec.widget was not found with criteria %v and options %v", criteria, options)
}
