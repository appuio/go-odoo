package odoo

import (
	"fmt"
)

// XQstUpload represents x_qst_upload model.
type XQstUpload struct {
	LastUpdate                      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                     *String   `xmlrpc:"display_name,omptempty"`
	Id                              *Int      `xmlrpc:"id,omptempty"`
	WriteDate                       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                        *Many2One `xmlrpc:"write_uid,omptempty"`
	XName                           *String   `xmlrpc:"x_name,omptempty"`
	XStudioCanton                   *String   `xmlrpc:"x_studio_canton,omptempty"`
	XStudioWage                     *Float    `xmlrpc:"x_studio_wage,omptempty"`
	XStudioWithholdingTaxCode       *String   `xmlrpc:"x_studio_withholding_tax_code,omptempty"`
	XStudioWithholdingTaxPercentage *Float    `xmlrpc:"x_studio_withholding_tax_percentage,omptempty"`
	XStudioYear                     *Int      `xmlrpc:"x_studio_year,omptempty"`
}

// XQstUploads represents array of x_qst_upload model.
type XQstUploads []XQstUpload

// XQstUploadModel is the odoo model name.
const XQstUploadModel = "x_qst_upload"

// Many2One convert XQstUpload to *Many2One.
func (x *XQstUpload) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXQstUpload creates a new x_qst_upload model and returns its id.
func (c *Client) CreateXQstUpload(x *XQstUpload) (int64, error) {
	ids, err := c.CreateXQstUploads([]*XQstUpload{x})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateXQstUpload creates a new x_qst_upload model and returns its id.
func (c *Client) CreateXQstUploads(xs []*XQstUpload) ([]int64, error) {
	var vv []interface{}
	for _, v := range xs {
		vv = append(vv, v)
	}
	return c.Create(XQstUploadModel, vv)
}

// UpdateXQstUpload updates an existing x_qst_upload record.
func (c *Client) UpdateXQstUpload(x *XQstUpload) error {
	return c.UpdateXQstUploads([]int64{x.Id.Get()}, x)
}

// UpdateXQstUploads updates existing x_qst_upload records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXQstUploads(ids []int64, x *XQstUpload) error {
	return c.Update(XQstUploadModel, ids, x)
}

// DeleteXQstUpload deletes an existing x_qst_upload record.
func (c *Client) DeleteXQstUpload(id int64) error {
	return c.DeleteXQstUploads([]int64{id})
}

// DeleteXQstUploads deletes existing x_qst_upload records.
func (c *Client) DeleteXQstUploads(ids []int64) error {
	return c.Delete(XQstUploadModel, ids)
}

// GetXQstUpload gets x_qst_upload existing record.
func (c *Client) GetXQstUpload(id int64) (*XQstUpload, error) {
	xs, err := c.GetXQstUploads([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_qst_upload not found", id)
}

// GetXQstUploads gets x_qst_upload existing records.
func (c *Client) GetXQstUploads(ids []int64) (*XQstUploads, error) {
	xs := &XQstUploads{}
	if err := c.Read(XQstUploadModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXQstUpload finds x_qst_upload record by querying it with criteria.
func (c *Client) FindXQstUpload(criteria *Criteria) (*XQstUpload, error) {
	xs := &XQstUploads{}
	if err := c.SearchRead(XQstUploadModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_qst_upload was not found with criteria %v", criteria)
}

// FindXQstUploads finds x_qst_upload records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXQstUploads(criteria *Criteria, options *Options) (*XQstUploads, error) {
	xs := &XQstUploads{}
	if err := c.SearchRead(XQstUploadModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXQstUploadIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXQstUploadIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XQstUploadModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXQstUploadId finds record id by querying it with criteria.
func (c *Client) FindXQstUploadId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XQstUploadModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_qst_upload was not found with criteria %v and options %v", criteria, options)
}
