package odoo

import (
	"fmt"
)

// AssetModify represents asset.modify model.
type AssetModify struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	AccountAssetCounterpartId    *Many2One  `xmlrpc:"account_asset_counterpart_id,omptempty"`
	AccountAssetId               *Many2One  `xmlrpc:"account_asset_id,omptempty"`
	AccountDepreciationExpenseId *Many2One  `xmlrpc:"account_depreciation_expense_id,omptempty"`
	AccountDepreciationId        *Many2One  `xmlrpc:"account_depreciation_id,omptempty"`
	AssetId                      *Many2One  `xmlrpc:"asset_id,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                         *Time      `xmlrpc:"date,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	GainAccountId                *Many2One  `xmlrpc:"gain_account_id,omptempty"`
	GainOrLoss                   *Selection `xmlrpc:"gain_or_loss,omptempty"`
	GainValue                    *Bool      `xmlrpc:"gain_value,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	InformationalText            *String    `xmlrpc:"informational_text,omptempty"`
	InvoiceIds                   *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceLineIds               *Relation  `xmlrpc:"invoice_line_ids,omptempty"`
	LossAccountId                *Many2One  `xmlrpc:"loss_account_id,omptempty"`
	MethodNumber                 *Int       `xmlrpc:"method_number,omptempty"`
	MethodPeriod                 *Selection `xmlrpc:"method_period,omptempty"`
	ModifyAction                 *Selection `xmlrpc:"modify_action,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	SalvageValue                 *Float     `xmlrpc:"salvage_value,omptempty"`
	SelectInvoiceLineId          *Bool      `xmlrpc:"select_invoice_line_id,omptempty"`
	ValueResidual                *Float     `xmlrpc:"value_residual,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
	XMethodPeriod                *Selection `xmlrpc:"x_method_period,omptempty"`
}

// AssetModifys represents array of asset.modify model.
type AssetModifys []AssetModify

// AssetModifyModel is the odoo model name.
const AssetModifyModel = "asset.modify"

// Many2One convert AssetModify to *Many2One.
func (am *AssetModify) Many2One() *Many2One {
	return NewMany2One(am.Id.Get(), "")
}

// CreateAssetModify creates a new asset.modify model and returns its id.
func (c *Client) CreateAssetModify(am *AssetModify) (int64, error) {
	ids, err := c.CreateAssetModifys([]*AssetModify{am})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAssetModify creates a new asset.modify model and returns its id.
func (c *Client) CreateAssetModifys(ams []*AssetModify) ([]int64, error) {
	var vv []interface{}
	for _, v := range ams {
		vv = append(vv, v)
	}
	return c.Create(AssetModifyModel, vv)
}

// UpdateAssetModify updates an existing asset.modify record.
func (c *Client) UpdateAssetModify(am *AssetModify) error {
	return c.UpdateAssetModifys([]int64{am.Id.Get()}, am)
}

// UpdateAssetModifys updates existing asset.modify records.
// All records (represented by ids) will be updated by am values.
func (c *Client) UpdateAssetModifys(ids []int64, am *AssetModify) error {
	return c.Update(AssetModifyModel, ids, am)
}

// DeleteAssetModify deletes an existing asset.modify record.
func (c *Client) DeleteAssetModify(id int64) error {
	return c.DeleteAssetModifys([]int64{id})
}

// DeleteAssetModifys deletes existing asset.modify records.
func (c *Client) DeleteAssetModifys(ids []int64) error {
	return c.Delete(AssetModifyModel, ids)
}

// GetAssetModify gets asset.modify existing record.
func (c *Client) GetAssetModify(id int64) (*AssetModify, error) {
	ams, err := c.GetAssetModifys([]int64{id})
	if err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of asset.modify not found", id)
}

// GetAssetModifys gets asset.modify existing records.
func (c *Client) GetAssetModifys(ids []int64) (*AssetModifys, error) {
	ams := &AssetModifys{}
	if err := c.Read(AssetModifyModel, ids, nil, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAssetModify finds asset.modify record by querying it with criteria.
func (c *Client) FindAssetModify(criteria *Criteria) (*AssetModify, error) {
	ams := &AssetModifys{}
	if err := c.SearchRead(AssetModifyModel, criteria, NewOptions().Limit(1), ams); err != nil {
		return nil, err
	}
	if ams != nil && len(*ams) > 0 {
		return &((*ams)[0]), nil
	}
	return nil, fmt.Errorf("asset.modify was not found with criteria %v", criteria)
}

// FindAssetModifys finds asset.modify records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetModifys(criteria *Criteria, options *Options) (*AssetModifys, error) {
	ams := &AssetModifys{}
	if err := c.SearchRead(AssetModifyModel, criteria, options, ams); err != nil {
		return nil, err
	}
	return ams, nil
}

// FindAssetModifyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAssetModifyIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AssetModifyModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAssetModifyId finds record id by querying it with criteria.
func (c *Client) FindAssetModifyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AssetModifyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("asset.modify was not found with criteria %v and options %v", criteria, options)
}
