package odoo

import (
	"fmt"
)

// HrContractType represents hr.contract.type model.
type HrContractType struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrContractTypes represents array of hr.contract.type model.
type HrContractTypes []HrContractType

// HrContractTypeModel is the odoo model name.
const HrContractTypeModel = "hr.contract.type"

// Many2One convert HrContractType to *Many2One.
func (hct *HrContractType) Many2One() *Many2One {
	return NewMany2One(hct.Id.Get(), "")
}

// CreateHrContractType creates a new hr.contract.type model and returns its id.
func (c *Client) CreateHrContractType(hct *HrContractType) (int64, error) {
	ids, err := c.CreateHrContractTypes([]*HrContractType{hct})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrContractType creates a new hr.contract.type model and returns its id.
func (c *Client) CreateHrContractTypes(hcts []*HrContractType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hcts {
		vv = append(vv, v)
	}
	return c.Create(HrContractTypeModel, vv)
}

// UpdateHrContractType updates an existing hr.contract.type record.
func (c *Client) UpdateHrContractType(hct *HrContractType) error {
	return c.UpdateHrContractTypes([]int64{hct.Id.Get()}, hct)
}

// UpdateHrContractTypes updates existing hr.contract.type records.
// All records (represented by ids) will be updated by hct values.
func (c *Client) UpdateHrContractTypes(ids []int64, hct *HrContractType) error {
	return c.Update(HrContractTypeModel, ids, hct)
}

// DeleteHrContractType deletes an existing hr.contract.type record.
func (c *Client) DeleteHrContractType(id int64) error {
	return c.DeleteHrContractTypes([]int64{id})
}

// DeleteHrContractTypes deletes existing hr.contract.type records.
func (c *Client) DeleteHrContractTypes(ids []int64) error {
	return c.Delete(HrContractTypeModel, ids)
}

// GetHrContractType gets hr.contract.type existing record.
func (c *Client) GetHrContractType(id int64) (*HrContractType, error) {
	hcts, err := c.GetHrContractTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if hcts != nil && len(*hcts) > 0 {
		return &((*hcts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.contract.type not found", id)
}

// GetHrContractTypes gets hr.contract.type existing records.
func (c *Client) GetHrContractTypes(ids []int64) (*HrContractTypes, error) {
	hcts := &HrContractTypes{}
	if err := c.Read(HrContractTypeModel, ids, nil, hcts); err != nil {
		return nil, err
	}
	return hcts, nil
}

// FindHrContractType finds hr.contract.type record by querying it with criteria.
func (c *Client) FindHrContractType(criteria *Criteria) (*HrContractType, error) {
	hcts := &HrContractTypes{}
	if err := c.SearchRead(HrContractTypeModel, criteria, NewOptions().Limit(1), hcts); err != nil {
		return nil, err
	}
	if hcts != nil && len(*hcts) > 0 {
		return &((*hcts)[0]), nil
	}
	return nil, fmt.Errorf("hr.contract.type was not found with criteria %v", criteria)
}

// FindHrContractTypes finds hr.contract.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractTypes(criteria *Criteria, options *Options) (*HrContractTypes, error) {
	hcts := &HrContractTypes{}
	if err := c.SearchRead(HrContractTypeModel, criteria, options, hcts); err != nil {
		return nil, err
	}
	return hcts, nil
}

// FindHrContractTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrContractTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrContractTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrContractTypeId finds record id by querying it with criteria.
func (c *Client) FindHrContractTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrContractTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.contract.type was not found with criteria %v and options %v", criteria, options)
}
