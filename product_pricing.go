package odoo

import (
	"fmt"
)

// ProductPricing represents product.pricing model.
type ProductPricing struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId        *Many2One `xmlrpc:"currency_id,omptempty"`
	Description       *String   `xmlrpc:"description,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	Name              *String   `xmlrpc:"name,omptempty"`
	Price             *Float    `xmlrpc:"price,omptempty"`
	PricelistId       *Many2One `xmlrpc:"pricelist_id,omptempty"`
	ProductTemplateId *Many2One `xmlrpc:"product_template_id,omptempty"`
	ProductVariantIds *Relation `xmlrpc:"product_variant_ids,omptempty"`
	RecurrenceId      *Many2One `xmlrpc:"recurrence_id,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductPricings represents array of product.pricing model.
type ProductPricings []ProductPricing

// ProductPricingModel is the odoo model name.
const ProductPricingModel = "product.pricing"

// Many2One convert ProductPricing to *Many2One.
func (pp *ProductPricing) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductPricing creates a new product.pricing model and returns its id.
func (c *Client) CreateProductPricing(pp *ProductPricing) (int64, error) {
	ids, err := c.CreateProductPricings([]*ProductPricing{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductPricing creates a new product.pricing model and returns its id.
func (c *Client) CreateProductPricings(pps []*ProductPricing) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductPricingModel, vv)
}

// UpdateProductPricing updates an existing product.pricing record.
func (c *Client) UpdateProductPricing(pp *ProductPricing) error {
	return c.UpdateProductPricings([]int64{pp.Id.Get()}, pp)
}

// UpdateProductPricings updates existing product.pricing records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductPricings(ids []int64, pp *ProductPricing) error {
	return c.Update(ProductPricingModel, ids, pp)
}

// DeleteProductPricing deletes an existing product.pricing record.
func (c *Client) DeleteProductPricing(id int64) error {
	return c.DeleteProductPricings([]int64{id})
}

// DeleteProductPricings deletes existing product.pricing records.
func (c *Client) DeleteProductPricings(ids []int64) error {
	return c.Delete(ProductPricingModel, ids)
}

// GetProductPricing gets product.pricing existing record.
func (c *Client) GetProductPricing(id int64) (*ProductPricing, error) {
	pps, err := c.GetProductPricings([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.pricing not found", id)
}

// GetProductPricings gets product.pricing existing records.
func (c *Client) GetProductPricings(ids []int64) (*ProductPricings, error) {
	pps := &ProductPricings{}
	if err := c.Read(ProductPricingModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricing finds product.pricing record by querying it with criteria.
func (c *Client) FindProductPricing(criteria *Criteria) (*ProductPricing, error) {
	pps := &ProductPricings{}
	if err := c.SearchRead(ProductPricingModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.pricing was not found with criteria %v", criteria)
}

// FindProductPricings finds product.pricing records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricings(criteria *Criteria, options *Options) (*ProductPricings, error) {
	pps := &ProductPricings{}
	if err := c.SearchRead(ProductPricingModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductPricingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPricingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPricingId finds record id by querying it with criteria.
func (c *Client) FindProductPricingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPricingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.pricing was not found with criteria %v and options %v", criteria, options)
}
