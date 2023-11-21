package odoo

import (
	"fmt"
)

// PurchaseBillUnion represents purchase.bill.union model.
type PurchaseBillUnion struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Amount          *Float    `xmlrpc:"amount,omptempty"`
	CompanyId       *Many2One `xmlrpc:"company_id,omptempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omptempty"`
	Date            *Time     `xmlrpc:"date,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	PartnerId       *Many2One `xmlrpc:"partner_id,omptempty"`
	PurchaseOrderId *Many2One `xmlrpc:"purchase_order_id,omptempty"`
	Reference       *String   `xmlrpc:"reference,omptempty"`
	VendorBillId    *Many2One `xmlrpc:"vendor_bill_id,omptempty"`
}

// PurchaseBillUnions represents array of purchase.bill.union model.
type PurchaseBillUnions []PurchaseBillUnion

// PurchaseBillUnionModel is the odoo model name.
const PurchaseBillUnionModel = "purchase.bill.union"

// Many2One convert PurchaseBillUnion to *Many2One.
func (pbu *PurchaseBillUnion) Many2One() *Many2One {
	return NewMany2One(pbu.Id.Get(), "")
}

// CreatePurchaseBillUnion creates a new purchase.bill.union model and returns its id.
func (c *Client) CreatePurchaseBillUnion(pbu *PurchaseBillUnion) (int64, error) {
	ids, err := c.CreatePurchaseBillUnions([]*PurchaseBillUnion{pbu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePurchaseBillUnion creates a new purchase.bill.union model and returns its id.
func (c *Client) CreatePurchaseBillUnions(pbus []*PurchaseBillUnion) ([]int64, error) {
	var vv []interface{}
	for _, v := range pbus {
		vv = append(vv, v)
	}
	return c.Create(PurchaseBillUnionModel, vv)
}

// UpdatePurchaseBillUnion updates an existing purchase.bill.union record.
func (c *Client) UpdatePurchaseBillUnion(pbu *PurchaseBillUnion) error {
	return c.UpdatePurchaseBillUnions([]int64{pbu.Id.Get()}, pbu)
}

// UpdatePurchaseBillUnions updates existing purchase.bill.union records.
// All records (represented by ids) will be updated by pbu values.
func (c *Client) UpdatePurchaseBillUnions(ids []int64, pbu *PurchaseBillUnion) error {
	return c.Update(PurchaseBillUnionModel, ids, pbu)
}

// DeletePurchaseBillUnion deletes an existing purchase.bill.union record.
func (c *Client) DeletePurchaseBillUnion(id int64) error {
	return c.DeletePurchaseBillUnions([]int64{id})
}

// DeletePurchaseBillUnions deletes existing purchase.bill.union records.
func (c *Client) DeletePurchaseBillUnions(ids []int64) error {
	return c.Delete(PurchaseBillUnionModel, ids)
}

// GetPurchaseBillUnion gets purchase.bill.union existing record.
func (c *Client) GetPurchaseBillUnion(id int64) (*PurchaseBillUnion, error) {
	pbus, err := c.GetPurchaseBillUnions([]int64{id})
	if err != nil {
		return nil, err
	}
	if pbus != nil && len(*pbus) > 0 {
		return &((*pbus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of purchase.bill.union not found", id)
}

// GetPurchaseBillUnions gets purchase.bill.union existing records.
func (c *Client) GetPurchaseBillUnions(ids []int64) (*PurchaseBillUnions, error) {
	pbus := &PurchaseBillUnions{}
	if err := c.Read(PurchaseBillUnionModel, ids, nil, pbus); err != nil {
		return nil, err
	}
	return pbus, nil
}

// FindPurchaseBillUnion finds purchase.bill.union record by querying it with criteria.
func (c *Client) FindPurchaseBillUnion(criteria *Criteria) (*PurchaseBillUnion, error) {
	pbus := &PurchaseBillUnions{}
	if err := c.SearchRead(PurchaseBillUnionModel, criteria, NewOptions().Limit(1), pbus); err != nil {
		return nil, err
	}
	if pbus != nil && len(*pbus) > 0 {
		return &((*pbus)[0]), nil
	}
	return nil, fmt.Errorf("purchase.bill.union was not found with criteria %v", criteria)
}

// FindPurchaseBillUnions finds purchase.bill.union records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseBillUnions(criteria *Criteria, options *Options) (*PurchaseBillUnions, error) {
	pbus := &PurchaseBillUnions{}
	if err := c.SearchRead(PurchaseBillUnionModel, criteria, options, pbus); err != nil {
		return nil, err
	}
	return pbus, nil
}

// FindPurchaseBillUnionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPurchaseBillUnionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PurchaseBillUnionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPurchaseBillUnionId finds record id by querying it with criteria.
func (c *Client) FindPurchaseBillUnionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PurchaseBillUnionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("purchase.bill.union was not found with criteria %v and options %v", criteria, options)
}
