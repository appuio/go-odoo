package odoo

import (
	"fmt"
)

// ProductImage represents product.image model.
type ProductImage struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	CanImage1024BeZoomed *Bool     `xmlrpc:"can_image_1024_be_zoomed,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	EmbedCode            *String   `xmlrpc:"embed_code,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	Image1024            *String   `xmlrpc:"image_1024,omptempty"`
	Image128             *String   `xmlrpc:"image_128,omptempty"`
	Image1920            *String   `xmlrpc:"image_1920,omptempty"`
	Image256             *String   `xmlrpc:"image_256,omptempty"`
	Image512             *String   `xmlrpc:"image_512,omptempty"`
	Name                 *String   `xmlrpc:"name,omptempty"`
	ProductTmplId        *Many2One `xmlrpc:"product_tmpl_id,omptempty"`
	ProductVariantId     *Many2One `xmlrpc:"product_variant_id,omptempty"`
	Sequence             *Int      `xmlrpc:"sequence,omptempty"`
	VideoUrl             *String   `xmlrpc:"video_url,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ProductImages represents array of product.image model.
type ProductImages []ProductImage

// ProductImageModel is the odoo model name.
const ProductImageModel = "product.image"

// Many2One convert ProductImage to *Many2One.
func (pi *ProductImage) Many2One() *Many2One {
	return NewMany2One(pi.Id.Get(), "")
}

// CreateProductImage creates a new product.image model and returns its id.
func (c *Client) CreateProductImage(pi *ProductImage) (int64, error) {
	ids, err := c.CreateProductImages([]*ProductImage{pi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductImage creates a new product.image model and returns its id.
func (c *Client) CreateProductImages(pis []*ProductImage) ([]int64, error) {
	var vv []interface{}
	for _, v := range pis {
		vv = append(vv, v)
	}
	return c.Create(ProductImageModel, vv)
}

// UpdateProductImage updates an existing product.image record.
func (c *Client) UpdateProductImage(pi *ProductImage) error {
	return c.UpdateProductImages([]int64{pi.Id.Get()}, pi)
}

// UpdateProductImages updates existing product.image records.
// All records (represented by ids) will be updated by pi values.
func (c *Client) UpdateProductImages(ids []int64, pi *ProductImage) error {
	return c.Update(ProductImageModel, ids, pi)
}

// DeleteProductImage deletes an existing product.image record.
func (c *Client) DeleteProductImage(id int64) error {
	return c.DeleteProductImages([]int64{id})
}

// DeleteProductImages deletes existing product.image records.
func (c *Client) DeleteProductImages(ids []int64) error {
	return c.Delete(ProductImageModel, ids)
}

// GetProductImage gets product.image existing record.
func (c *Client) GetProductImage(id int64) (*ProductImage, error) {
	pis, err := c.GetProductImages([]int64{id})
	if err != nil {
		return nil, err
	}
	if pis != nil && len(*pis) > 0 {
		return &((*pis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.image not found", id)
}

// GetProductImages gets product.image existing records.
func (c *Client) GetProductImages(ids []int64) (*ProductImages, error) {
	pis := &ProductImages{}
	if err := c.Read(ProductImageModel, ids, nil, pis); err != nil {
		return nil, err
	}
	return pis, nil
}

// FindProductImage finds product.image record by querying it with criteria.
func (c *Client) FindProductImage(criteria *Criteria) (*ProductImage, error) {
	pis := &ProductImages{}
	if err := c.SearchRead(ProductImageModel, criteria, NewOptions().Limit(1), pis); err != nil {
		return nil, err
	}
	if pis != nil && len(*pis) > 0 {
		return &((*pis)[0]), nil
	}
	return nil, fmt.Errorf("product.image was not found with criteria %v", criteria)
}

// FindProductImages finds product.image records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductImages(criteria *Criteria, options *Options) (*ProductImages, error) {
	pis := &ProductImages{}
	if err := c.SearchRead(ProductImageModel, criteria, options, pis); err != nil {
		return nil, err
	}
	return pis, nil
}

// FindProductImageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductImageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductImageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductImageId finds record id by querying it with criteria.
func (c *Client) FindProductImageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductImageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.image was not found with criteria %v and options %v", criteria, options)
}
