package odoo

import (
	"fmt"
)

// ProductPublicCategory represents product.public.category model.
type ProductPublicCategory struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	ChildId                *Relation `xmlrpc:"child_id,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	Image1024              *String   `xmlrpc:"image_1024,omptempty"`
	Image128               *String   `xmlrpc:"image_128,omptempty"`
	Image1920              *String   `xmlrpc:"image_1920,omptempty"`
	Image256               *String   `xmlrpc:"image_256,omptempty"`
	Image512               *String   `xmlrpc:"image_512,omptempty"`
	IsSeoOptimized         *Bool     `xmlrpc:"is_seo_optimized,omptempty"`
	Name                   *String   `xmlrpc:"name,omptempty"`
	ParentId               *Many2One `xmlrpc:"parent_id,omptempty"`
	ParentPath             *String   `xmlrpc:"parent_path,omptempty"`
	ParentsAndSelf         *Relation `xmlrpc:"parents_and_self,omptempty"`
	ProductTmplIds         *Relation `xmlrpc:"product_tmpl_ids,omptempty"`
	SeoName                *String   `xmlrpc:"seo_name,omptempty"`
	Sequence               *Int      `xmlrpc:"sequence,omptempty"`
	WebsiteDescription     *String   `xmlrpc:"website_description,omptempty"`
	WebsiteId              *Many2One `xmlrpc:"website_id,omptempty"`
	WebsiteMetaDescription *String   `xmlrpc:"website_meta_description,omptempty"`
	WebsiteMetaKeywords    *String   `xmlrpc:"website_meta_keywords,omptempty"`
	WebsiteMetaOgImg       *String   `xmlrpc:"website_meta_og_img,omptempty"`
	WebsiteMetaTitle       *String   `xmlrpc:"website_meta_title,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductPublicCategorys represents array of product.public.category model.
type ProductPublicCategorys []ProductPublicCategory

// ProductPublicCategoryModel is the odoo model name.
const ProductPublicCategoryModel = "product.public.category"

// Many2One convert ProductPublicCategory to *Many2One.
func (ppc *ProductPublicCategory) Many2One() *Many2One {
	return NewMany2One(ppc.Id.Get(), "")
}

// CreateProductPublicCategory creates a new product.public.category model and returns its id.
func (c *Client) CreateProductPublicCategory(ppc *ProductPublicCategory) (int64, error) {
	ids, err := c.CreateProductPublicCategorys([]*ProductPublicCategory{ppc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductPublicCategory creates a new product.public.category model and returns its id.
func (c *Client) CreateProductPublicCategorys(ppcs []*ProductPublicCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppcs {
		vv = append(vv, v)
	}
	return c.Create(ProductPublicCategoryModel, vv)
}

// UpdateProductPublicCategory updates an existing product.public.category record.
func (c *Client) UpdateProductPublicCategory(ppc *ProductPublicCategory) error {
	return c.UpdateProductPublicCategorys([]int64{ppc.Id.Get()}, ppc)
}

// UpdateProductPublicCategorys updates existing product.public.category records.
// All records (represented by ids) will be updated by ppc values.
func (c *Client) UpdateProductPublicCategorys(ids []int64, ppc *ProductPublicCategory) error {
	return c.Update(ProductPublicCategoryModel, ids, ppc)
}

// DeleteProductPublicCategory deletes an existing product.public.category record.
func (c *Client) DeleteProductPublicCategory(id int64) error {
	return c.DeleteProductPublicCategorys([]int64{id})
}

// DeleteProductPublicCategorys deletes existing product.public.category records.
func (c *Client) DeleteProductPublicCategorys(ids []int64) error {
	return c.Delete(ProductPublicCategoryModel, ids)
}

// GetProductPublicCategory gets product.public.category existing record.
func (c *Client) GetProductPublicCategory(id int64) (*ProductPublicCategory, error) {
	ppcs, err := c.GetProductPublicCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if ppcs != nil && len(*ppcs) > 0 {
		return &((*ppcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.public.category not found", id)
}

// GetProductPublicCategorys gets product.public.category existing records.
func (c *Client) GetProductPublicCategorys(ids []int64) (*ProductPublicCategorys, error) {
	ppcs := &ProductPublicCategorys{}
	if err := c.Read(ProductPublicCategoryModel, ids, nil, ppcs); err != nil {
		return nil, err
	}
	return ppcs, nil
}

// FindProductPublicCategory finds product.public.category record by querying it with criteria.
func (c *Client) FindProductPublicCategory(criteria *Criteria) (*ProductPublicCategory, error) {
	ppcs := &ProductPublicCategorys{}
	if err := c.SearchRead(ProductPublicCategoryModel, criteria, NewOptions().Limit(1), ppcs); err != nil {
		return nil, err
	}
	if ppcs != nil && len(*ppcs) > 0 {
		return &((*ppcs)[0]), nil
	}
	return nil, fmt.Errorf("product.public.category was not found with criteria %v", criteria)
}

// FindProductPublicCategorys finds product.public.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPublicCategorys(criteria *Criteria, options *Options) (*ProductPublicCategorys, error) {
	ppcs := &ProductPublicCategorys{}
	if err := c.SearchRead(ProductPublicCategoryModel, criteria, options, ppcs); err != nil {
		return nil, err
	}
	return ppcs, nil
}

// FindProductPublicCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPublicCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductPublicCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductPublicCategoryId finds record id by querying it with criteria.
func (c *Client) FindProductPublicCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPublicCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.public.category was not found with criteria %v and options %v", criteria, options)
}
