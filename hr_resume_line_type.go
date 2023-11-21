package odoo

import (
	"fmt"
)

// HrResumeLineType represents hr.resume.line.type model.
type HrResumeLineType struct {
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

// HrResumeLineTypes represents array of hr.resume.line.type model.
type HrResumeLineTypes []HrResumeLineType

// HrResumeLineTypeModel is the odoo model name.
const HrResumeLineTypeModel = "hr.resume.line.type"

// Many2One convert HrResumeLineType to *Many2One.
func (hrlt *HrResumeLineType) Many2One() *Many2One {
	return NewMany2One(hrlt.Id.Get(), "")
}

// CreateHrResumeLineType creates a new hr.resume.line.type model and returns its id.
func (c *Client) CreateHrResumeLineType(hrlt *HrResumeLineType) (int64, error) {
	ids, err := c.CreateHrResumeLineTypes([]*HrResumeLineType{hrlt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrResumeLineType creates a new hr.resume.line.type model and returns its id.
func (c *Client) CreateHrResumeLineTypes(hrlts []*HrResumeLineType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrlts {
		vv = append(vv, v)
	}
	return c.Create(HrResumeLineTypeModel, vv)
}

// UpdateHrResumeLineType updates an existing hr.resume.line.type record.
func (c *Client) UpdateHrResumeLineType(hrlt *HrResumeLineType) error {
	return c.UpdateHrResumeLineTypes([]int64{hrlt.Id.Get()}, hrlt)
}

// UpdateHrResumeLineTypes updates existing hr.resume.line.type records.
// All records (represented by ids) will be updated by hrlt values.
func (c *Client) UpdateHrResumeLineTypes(ids []int64, hrlt *HrResumeLineType) error {
	return c.Update(HrResumeLineTypeModel, ids, hrlt)
}

// DeleteHrResumeLineType deletes an existing hr.resume.line.type record.
func (c *Client) DeleteHrResumeLineType(id int64) error {
	return c.DeleteHrResumeLineTypes([]int64{id})
}

// DeleteHrResumeLineTypes deletes existing hr.resume.line.type records.
func (c *Client) DeleteHrResumeLineTypes(ids []int64) error {
	return c.Delete(HrResumeLineTypeModel, ids)
}

// GetHrResumeLineType gets hr.resume.line.type existing record.
func (c *Client) GetHrResumeLineType(id int64) (*HrResumeLineType, error) {
	hrlts, err := c.GetHrResumeLineTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrlts != nil && len(*hrlts) > 0 {
		return &((*hrlts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.resume.line.type not found", id)
}

// GetHrResumeLineTypes gets hr.resume.line.type existing records.
func (c *Client) GetHrResumeLineTypes(ids []int64) (*HrResumeLineTypes, error) {
	hrlts := &HrResumeLineTypes{}
	if err := c.Read(HrResumeLineTypeModel, ids, nil, hrlts); err != nil {
		return nil, err
	}
	return hrlts, nil
}

// FindHrResumeLineType finds hr.resume.line.type record by querying it with criteria.
func (c *Client) FindHrResumeLineType(criteria *Criteria) (*HrResumeLineType, error) {
	hrlts := &HrResumeLineTypes{}
	if err := c.SearchRead(HrResumeLineTypeModel, criteria, NewOptions().Limit(1), hrlts); err != nil {
		return nil, err
	}
	if hrlts != nil && len(*hrlts) > 0 {
		return &((*hrlts)[0]), nil
	}
	return nil, fmt.Errorf("hr.resume.line.type was not found with criteria %v", criteria)
}

// FindHrResumeLineTypes finds hr.resume.line.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLineTypes(criteria *Criteria, options *Options) (*HrResumeLineTypes, error) {
	hrlts := &HrResumeLineTypes{}
	if err := c.SearchRead(HrResumeLineTypeModel, criteria, options, hrlts); err != nil {
		return nil, err
	}
	return hrlts, nil
}

// FindHrResumeLineTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLineTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrResumeLineTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrResumeLineTypeId finds record id by querying it with criteria.
func (c *Client) FindHrResumeLineTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrResumeLineTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.resume.line.type was not found with criteria %v and options %v", criteria, options)
}
