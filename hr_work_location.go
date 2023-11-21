package odoo

import (
	"fmt"
)

// HrWorkLocation represents hr.work.location model.
type HrWorkLocation struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	Active         *Bool     `xmlrpc:"active,omptempty"`
	AddressId      *Many2One `xmlrpc:"address_id,omptempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LocationNumber *String   `xmlrpc:"location_number,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrWorkLocations represents array of hr.work.location model.
type HrWorkLocations []HrWorkLocation

// HrWorkLocationModel is the odoo model name.
const HrWorkLocationModel = "hr.work.location"

// Many2One convert HrWorkLocation to *Many2One.
func (hwl *HrWorkLocation) Many2One() *Many2One {
	return NewMany2One(hwl.Id.Get(), "")
}

// CreateHrWorkLocation creates a new hr.work.location model and returns its id.
func (c *Client) CreateHrWorkLocation(hwl *HrWorkLocation) (int64, error) {
	ids, err := c.CreateHrWorkLocations([]*HrWorkLocation{hwl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrWorkLocation creates a new hr.work.location model and returns its id.
func (c *Client) CreateHrWorkLocations(hwls []*HrWorkLocation) ([]int64, error) {
	var vv []interface{}
	for _, v := range hwls {
		vv = append(vv, v)
	}
	return c.Create(HrWorkLocationModel, vv)
}

// UpdateHrWorkLocation updates an existing hr.work.location record.
func (c *Client) UpdateHrWorkLocation(hwl *HrWorkLocation) error {
	return c.UpdateHrWorkLocations([]int64{hwl.Id.Get()}, hwl)
}

// UpdateHrWorkLocations updates existing hr.work.location records.
// All records (represented by ids) will be updated by hwl values.
func (c *Client) UpdateHrWorkLocations(ids []int64, hwl *HrWorkLocation) error {
	return c.Update(HrWorkLocationModel, ids, hwl)
}

// DeleteHrWorkLocation deletes an existing hr.work.location record.
func (c *Client) DeleteHrWorkLocation(id int64) error {
	return c.DeleteHrWorkLocations([]int64{id})
}

// DeleteHrWorkLocations deletes existing hr.work.location records.
func (c *Client) DeleteHrWorkLocations(ids []int64) error {
	return c.Delete(HrWorkLocationModel, ids)
}

// GetHrWorkLocation gets hr.work.location existing record.
func (c *Client) GetHrWorkLocation(id int64) (*HrWorkLocation, error) {
	hwls, err := c.GetHrWorkLocations([]int64{id})
	if err != nil {
		return nil, err
	}
	if hwls != nil && len(*hwls) > 0 {
		return &((*hwls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.work.location not found", id)
}

// GetHrWorkLocations gets hr.work.location existing records.
func (c *Client) GetHrWorkLocations(ids []int64) (*HrWorkLocations, error) {
	hwls := &HrWorkLocations{}
	if err := c.Read(HrWorkLocationModel, ids, nil, hwls); err != nil {
		return nil, err
	}
	return hwls, nil
}

// FindHrWorkLocation finds hr.work.location record by querying it with criteria.
func (c *Client) FindHrWorkLocation(criteria *Criteria) (*HrWorkLocation, error) {
	hwls := &HrWorkLocations{}
	if err := c.SearchRead(HrWorkLocationModel, criteria, NewOptions().Limit(1), hwls); err != nil {
		return nil, err
	}
	if hwls != nil && len(*hwls) > 0 {
		return &((*hwls)[0]), nil
	}
	return nil, fmt.Errorf("hr.work.location was not found with criteria %v", criteria)
}

// FindHrWorkLocations finds hr.work.location records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkLocations(criteria *Criteria, options *Options) (*HrWorkLocations, error) {
	hwls := &HrWorkLocations{}
	if err := c.SearchRead(HrWorkLocationModel, criteria, options, hwls); err != nil {
		return nil, err
	}
	return hwls, nil
}

// FindHrWorkLocationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrWorkLocationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrWorkLocationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrWorkLocationId finds record id by querying it with criteria.
func (c *Client) FindHrWorkLocationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrWorkLocationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.work.location was not found with criteria %v and options %v", criteria, options)
}
