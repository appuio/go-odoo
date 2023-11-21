package odoo

import (
	"fmt"
)

// BankRecWidgetLine represents bank.rec.widget.line model.
type BankRecWidgetLine struct {
	LastUpdate                   *Time       `xmlrpc:"__last_update,omptempty"`
	AccountId                    *Many2One   `xmlrpc:"account_id,omptempty"`
	AmountCurrency               *Float      `xmlrpc:"amount_currency,omptempty"`
	AnalyticDistribution         interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticDistributionSearch   interface{} `xmlrpc:"analytic_distribution_search,omptempty"`
	AnalyticPrecision            *Int        `xmlrpc:"analytic_precision,omptempty"`
	Balance                      *Float      `xmlrpc:"balance,omptempty"`
	CompanyCurrencyId            *Many2One   `xmlrpc:"company_currency_id,omptempty"`
	Credit                       *Float      `xmlrpc:"credit,omptempty"`
	CurrencyId                   *Many2One   `xmlrpc:"currency_id,omptempty"`
	Date                         *Time       `xmlrpc:"date,omptempty"`
	Debit                        *Float      `xmlrpc:"debit,omptempty"`
	DisplayName                  *String     `xmlrpc:"display_name,omptempty"`
	DisplayStrokedAmountCurrency *Bool       `xmlrpc:"display_stroked_amount_currency,omptempty"`
	DisplayStrokedBalance        *Bool       `xmlrpc:"display_stroked_balance,omptempty"`
	Flag                         *Selection  `xmlrpc:"flag,omptempty"`
	ForcePriceIncludedTaxes      *Bool       `xmlrpc:"force_price_included_taxes,omptempty"`
	GroupTaxId                   *Many2One   `xmlrpc:"group_tax_id,omptempty"`
	Id                           *Int        `xmlrpc:"id,omptempty"`
	Index                        *String     `xmlrpc:"index,omptempty"`
	ManuallyModified             *Bool       `xmlrpc:"manually_modified,omptempty"`
	Name                         *String     `xmlrpc:"name,omptempty"`
	PartnerId                    *Many2One   `xmlrpc:"partner_id,omptempty"`
	ReconcileModelId             *Many2One   `xmlrpc:"reconcile_model_id,omptempty"`
	SourceAmlId                  *Many2One   `xmlrpc:"source_aml_id,omptempty"`
	SourceAmlMoveId              *Many2One   `xmlrpc:"source_aml_move_id,omptempty"`
	SourceAmlMoveName            *String     `xmlrpc:"source_aml_move_name,omptempty"`
	SourceAmountCurrency         *Float      `xmlrpc:"source_amount_currency,omptempty"`
	SourceBalance                *Float      `xmlrpc:"source_balance,omptempty"`
	SourceBatchPaymentId         *Many2One   `xmlrpc:"source_batch_payment_id,omptempty"`
	SourceCredit                 *Float      `xmlrpc:"source_credit,omptempty"`
	SourceDebit                  *Float      `xmlrpc:"source_debit,omptempty"`
	TaxBaseAmountCurrency        *Float      `xmlrpc:"tax_base_amount_currency,omptempty"`
	TaxIds                       *Relation   `xmlrpc:"tax_ids,omptempty"`
	TaxRepartitionLineId         *Many2One   `xmlrpc:"tax_repartition_line_id,omptempty"`
	TaxTagIds                    *Relation   `xmlrpc:"tax_tag_ids,omptempty"`
	WizardId                     *Many2One   `xmlrpc:"wizard_id,omptempty"`
}

// BankRecWidgetLines represents array of bank.rec.widget.line model.
type BankRecWidgetLines []BankRecWidgetLine

// BankRecWidgetLineModel is the odoo model name.
const BankRecWidgetLineModel = "bank.rec.widget.line"

// Many2One convert BankRecWidgetLine to *Many2One.
func (brwl *BankRecWidgetLine) Many2One() *Many2One {
	return NewMany2One(brwl.Id.Get(), "")
}

// CreateBankRecWidgetLine creates a new bank.rec.widget.line model and returns its id.
func (c *Client) CreateBankRecWidgetLine(brwl *BankRecWidgetLine) (int64, error) {
	ids, err := c.CreateBankRecWidgetLines([]*BankRecWidgetLine{brwl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBankRecWidgetLine creates a new bank.rec.widget.line model and returns its id.
func (c *Client) CreateBankRecWidgetLines(brwls []*BankRecWidgetLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range brwls {
		vv = append(vv, v)
	}
	return c.Create(BankRecWidgetLineModel, vv)
}

// UpdateBankRecWidgetLine updates an existing bank.rec.widget.line record.
func (c *Client) UpdateBankRecWidgetLine(brwl *BankRecWidgetLine) error {
	return c.UpdateBankRecWidgetLines([]int64{brwl.Id.Get()}, brwl)
}

// UpdateBankRecWidgetLines updates existing bank.rec.widget.line records.
// All records (represented by ids) will be updated by brwl values.
func (c *Client) UpdateBankRecWidgetLines(ids []int64, brwl *BankRecWidgetLine) error {
	return c.Update(BankRecWidgetLineModel, ids, brwl)
}

// DeleteBankRecWidgetLine deletes an existing bank.rec.widget.line record.
func (c *Client) DeleteBankRecWidgetLine(id int64) error {
	return c.DeleteBankRecWidgetLines([]int64{id})
}

// DeleteBankRecWidgetLines deletes existing bank.rec.widget.line records.
func (c *Client) DeleteBankRecWidgetLines(ids []int64) error {
	return c.Delete(BankRecWidgetLineModel, ids)
}

// GetBankRecWidgetLine gets bank.rec.widget.line existing record.
func (c *Client) GetBankRecWidgetLine(id int64) (*BankRecWidgetLine, error) {
	brwls, err := c.GetBankRecWidgetLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if brwls != nil && len(*brwls) > 0 {
		return &((*brwls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of bank.rec.widget.line not found", id)
}

// GetBankRecWidgetLines gets bank.rec.widget.line existing records.
func (c *Client) GetBankRecWidgetLines(ids []int64) (*BankRecWidgetLines, error) {
	brwls := &BankRecWidgetLines{}
	if err := c.Read(BankRecWidgetLineModel, ids, nil, brwls); err != nil {
		return nil, err
	}
	return brwls, nil
}

// FindBankRecWidgetLine finds bank.rec.widget.line record by querying it with criteria.
func (c *Client) FindBankRecWidgetLine(criteria *Criteria) (*BankRecWidgetLine, error) {
	brwls := &BankRecWidgetLines{}
	if err := c.SearchRead(BankRecWidgetLineModel, criteria, NewOptions().Limit(1), brwls); err != nil {
		return nil, err
	}
	if brwls != nil && len(*brwls) > 0 {
		return &((*brwls)[0]), nil
	}
	return nil, fmt.Errorf("bank.rec.widget.line was not found with criteria %v", criteria)
}

// FindBankRecWidgetLines finds bank.rec.widget.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgetLines(criteria *Criteria, options *Options) (*BankRecWidgetLines, error) {
	brwls := &BankRecWidgetLines{}
	if err := c.SearchRead(BankRecWidgetLineModel, criteria, options, brwls); err != nil {
		return nil, err
	}
	return brwls, nil
}

// FindBankRecWidgetLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBankRecWidgetLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BankRecWidgetLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBankRecWidgetLineId finds record id by querying it with criteria.
func (c *Client) FindBankRecWidgetLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BankRecWidgetLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("bank.rec.widget.line was not found with criteria %v and options %v", criteria, options)
}
