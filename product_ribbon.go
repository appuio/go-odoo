package odoo

import (
	"fmt"
)

// ProductRibbon represents product.ribbon model.
type ProductRibbon struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	BgColor       *String   `xmlrpc:"bg_color,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Html          *String   `xmlrpc:"html,omptempty"`
	HtmlClass     *String   `xmlrpc:"html_class,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	ProductTagIds *Relation `xmlrpc:"product_tag_ids,omptempty"`
	TextColor     *String   `xmlrpc:"text_color,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductRibbons represents array of product.ribbon model.
type ProductRibbons []ProductRibbon

// ProductRibbonModel is the odoo model name.
const ProductRibbonModel = "product.ribbon"

// Many2One convert ProductRibbon to *Many2One.
func (pr *ProductRibbon) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreateProductRibbon creates a new product.ribbon model and returns its id.
func (c *Client) CreateProductRibbon(pr *ProductRibbon) (int64, error) {
	ids, err := c.CreateProductRibbons([]*ProductRibbon{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductRibbon creates a new product.ribbon model and returns its id.
func (c *Client) CreateProductRibbons(prs []*ProductRibbon) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(ProductRibbonModel, vv)
}

// UpdateProductRibbon updates an existing product.ribbon record.
func (c *Client) UpdateProductRibbon(pr *ProductRibbon) error {
	return c.UpdateProductRibbons([]int64{pr.Id.Get()}, pr)
}

// UpdateProductRibbons updates existing product.ribbon records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProductRibbons(ids []int64, pr *ProductRibbon) error {
	return c.Update(ProductRibbonModel, ids, pr)
}

// DeleteProductRibbon deletes an existing product.ribbon record.
func (c *Client) DeleteProductRibbon(id int64) error {
	return c.DeleteProductRibbons([]int64{id})
}

// DeleteProductRibbons deletes existing product.ribbon records.
func (c *Client) DeleteProductRibbons(ids []int64) error {
	return c.Delete(ProductRibbonModel, ids)
}

// GetProductRibbon gets product.ribbon existing record.
func (c *Client) GetProductRibbon(id int64) (*ProductRibbon, error) {
	prs, err := c.GetProductRibbons([]int64{id})
	if err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.ribbon not found", id)
}

// GetProductRibbons gets product.ribbon existing records.
func (c *Client) GetProductRibbons(ids []int64) (*ProductRibbons, error) {
	prs := &ProductRibbons{}
	if err := c.Read(ProductRibbonModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductRibbon finds product.ribbon record by querying it with criteria.
func (c *Client) FindProductRibbon(criteria *Criteria) (*ProductRibbon, error) {
	prs := &ProductRibbons{}
	if err := c.SearchRead(ProductRibbonModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("product.ribbon was not found with criteria %v", criteria)
}

// FindProductRibbons finds product.ribbon records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductRibbons(criteria *Criteria, options *Options) (*ProductRibbons, error) {
	prs := &ProductRibbons{}
	if err := c.SearchRead(ProductRibbonModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductRibbonIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductRibbonIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductRibbonModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductRibbonId finds record id by querying it with criteria.
func (c *Client) FindProductRibbonId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductRibbonModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.ribbon was not found with criteria %v and options %v", criteria, options)
}
