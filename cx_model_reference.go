package odoo

import (
	"fmt"
)

// CxModelReference represents cx.model.reference model.
type CxModelReference struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	CustomName  *String   `xmlrpc:"custom_name,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	IrModelId   *Many2One `xmlrpc:"ir_model_id,omptempty"`
	Model       *String   `xmlrpc:"model,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CxModelReferences represents array of cx.model.reference model.
type CxModelReferences []CxModelReference

// CxModelReferenceModel is the odoo model name.
const CxModelReferenceModel = "cx.model.reference"

// Many2One convert CxModelReference to *Many2One.
func (cmr *CxModelReference) Many2One() *Many2One {
	return NewMany2One(cmr.Id.Get(), "")
}

// CreateCxModelReference creates a new cx.model.reference model and returns its id.
func (c *Client) CreateCxModelReference(cmr *CxModelReference) (int64, error) {
	ids, err := c.CreateCxModelReferences([]*CxModelReference{cmr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCxModelReference creates a new cx.model.reference model and returns its id.
func (c *Client) CreateCxModelReferences(cmrs []*CxModelReference) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmrs {
		vv = append(vv, v)
	}
	return c.Create(CxModelReferenceModel, vv)
}

// UpdateCxModelReference updates an existing cx.model.reference record.
func (c *Client) UpdateCxModelReference(cmr *CxModelReference) error {
	return c.UpdateCxModelReferences([]int64{cmr.Id.Get()}, cmr)
}

// UpdateCxModelReferences updates existing cx.model.reference records.
// All records (represented by ids) will be updated by cmr values.
func (c *Client) UpdateCxModelReferences(ids []int64, cmr *CxModelReference) error {
	return c.Update(CxModelReferenceModel, ids, cmr)
}

// DeleteCxModelReference deletes an existing cx.model.reference record.
func (c *Client) DeleteCxModelReference(id int64) error {
	return c.DeleteCxModelReferences([]int64{id})
}

// DeleteCxModelReferences deletes existing cx.model.reference records.
func (c *Client) DeleteCxModelReferences(ids []int64) error {
	return c.Delete(CxModelReferenceModel, ids)
}

// GetCxModelReference gets cx.model.reference existing record.
func (c *Client) GetCxModelReference(id int64) (*CxModelReference, error) {
	cmrs, err := c.GetCxModelReferences([]int64{id})
	if err != nil {
		return nil, err
	}
	if cmrs != nil && len(*cmrs) > 0 {
		return &((*cmrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cx.model.reference not found", id)
}

// GetCxModelReferences gets cx.model.reference existing records.
func (c *Client) GetCxModelReferences(ids []int64) (*CxModelReferences, error) {
	cmrs := &CxModelReferences{}
	if err := c.Read(CxModelReferenceModel, ids, nil, cmrs); err != nil {
		return nil, err
	}
	return cmrs, nil
}

// FindCxModelReference finds cx.model.reference record by querying it with criteria.
func (c *Client) FindCxModelReference(criteria *Criteria) (*CxModelReference, error) {
	cmrs := &CxModelReferences{}
	if err := c.SearchRead(CxModelReferenceModel, criteria, NewOptions().Limit(1), cmrs); err != nil {
		return nil, err
	}
	if cmrs != nil && len(*cmrs) > 0 {
		return &((*cmrs)[0]), nil
	}
	return nil, fmt.Errorf("cx.model.reference was not found with criteria %v", criteria)
}

// FindCxModelReferences finds cx.model.reference records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxModelReferences(criteria *Criteria, options *Options) (*CxModelReferences, error) {
	cmrs := &CxModelReferences{}
	if err := c.SearchRead(CxModelReferenceModel, criteria, options, cmrs); err != nil {
		return nil, err
	}
	return cmrs, nil
}

// FindCxModelReferenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCxModelReferenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CxModelReferenceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCxModelReferenceId finds record id by querying it with criteria.
func (c *Client) FindCxModelReferenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CxModelReferenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cx.model.reference was not found with criteria %v and options %v", criteria, options)
}
