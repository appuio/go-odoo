package odoo

import (
	"fmt"
)

// AccountTaxUnit represents account.tax.unit model.
type AccountTaxUnit struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyIds    *Relation `xmlrpc:"company_ids,omptempty"`
	CountryId     *Many2One `xmlrpc:"country_id,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	MainCompanyId *Many2One `xmlrpc:"main_company_id,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	Vat           *String   `xmlrpc:"vat,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountTaxUnits represents array of account.tax.unit model.
type AccountTaxUnits []AccountTaxUnit

// AccountTaxUnitModel is the odoo model name.
const AccountTaxUnitModel = "account.tax.unit"

// Many2One convert AccountTaxUnit to *Many2One.
func (atu *AccountTaxUnit) Many2One() *Many2One {
	return NewMany2One(atu.Id.Get(), "")
}

// CreateAccountTaxUnit creates a new account.tax.unit model and returns its id.
func (c *Client) CreateAccountTaxUnit(atu *AccountTaxUnit) (int64, error) {
	ids, err := c.CreateAccountTaxUnits([]*AccountTaxUnit{atu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountTaxUnit creates a new account.tax.unit model and returns its id.
func (c *Client) CreateAccountTaxUnits(atus []*AccountTaxUnit) ([]int64, error) {
	var vv []interface{}
	for _, v := range atus {
		vv = append(vv, v)
	}
	return c.Create(AccountTaxUnitModel, vv)
}

// UpdateAccountTaxUnit updates an existing account.tax.unit record.
func (c *Client) UpdateAccountTaxUnit(atu *AccountTaxUnit) error {
	return c.UpdateAccountTaxUnits([]int64{atu.Id.Get()}, atu)
}

// UpdateAccountTaxUnits updates existing account.tax.unit records.
// All records (represented by ids) will be updated by atu values.
func (c *Client) UpdateAccountTaxUnits(ids []int64, atu *AccountTaxUnit) error {
	return c.Update(AccountTaxUnitModel, ids, atu)
}

// DeleteAccountTaxUnit deletes an existing account.tax.unit record.
func (c *Client) DeleteAccountTaxUnit(id int64) error {
	return c.DeleteAccountTaxUnits([]int64{id})
}

// DeleteAccountTaxUnits deletes existing account.tax.unit records.
func (c *Client) DeleteAccountTaxUnits(ids []int64) error {
	return c.Delete(AccountTaxUnitModel, ids)
}

// GetAccountTaxUnit gets account.tax.unit existing record.
func (c *Client) GetAccountTaxUnit(id int64) (*AccountTaxUnit, error) {
	atus, err := c.GetAccountTaxUnits([]int64{id})
	if err != nil {
		return nil, err
	}
	if atus != nil && len(*atus) > 0 {
		return &((*atus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tax.unit not found", id)
}

// GetAccountTaxUnits gets account.tax.unit existing records.
func (c *Client) GetAccountTaxUnits(ids []int64) (*AccountTaxUnits, error) {
	atus := &AccountTaxUnits{}
	if err := c.Read(AccountTaxUnitModel, ids, nil, atus); err != nil {
		return nil, err
	}
	return atus, nil
}

// FindAccountTaxUnit finds account.tax.unit record by querying it with criteria.
func (c *Client) FindAccountTaxUnit(criteria *Criteria) (*AccountTaxUnit, error) {
	atus := &AccountTaxUnits{}
	if err := c.SearchRead(AccountTaxUnitModel, criteria, NewOptions().Limit(1), atus); err != nil {
		return nil, err
	}
	if atus != nil && len(*atus) > 0 {
		return &((*atus)[0]), nil
	}
	return nil, fmt.Errorf("account.tax.unit was not found with criteria %v", criteria)
}

// FindAccountTaxUnits finds account.tax.unit records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxUnits(criteria *Criteria, options *Options) (*AccountTaxUnits, error) {
	atus := &AccountTaxUnits{}
	if err := c.SearchRead(AccountTaxUnitModel, criteria, options, atus); err != nil {
		return nil, err
	}
	return atus, nil
}

// FindAccountTaxUnitIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxUnitIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTaxUnitModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTaxUnitId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxUnitId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxUnitModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tax.unit was not found with criteria %v and options %v", criteria, options)
}
