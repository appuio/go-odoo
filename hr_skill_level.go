package odoo

import (
	"fmt"
)

// HrSkillLevel represents hr.skill.level model.
type HrSkillLevel struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DefaultLevel  *Bool     `xmlrpc:"default_level,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LevelProgress *Int      `xmlrpc:"level_progress,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	SkillTypeId   *Many2One `xmlrpc:"skill_type_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrSkillLevels represents array of hr.skill.level model.
type HrSkillLevels []HrSkillLevel

// HrSkillLevelModel is the odoo model name.
const HrSkillLevelModel = "hr.skill.level"

// Many2One convert HrSkillLevel to *Many2One.
func (hsl *HrSkillLevel) Many2One() *Many2One {
	return NewMany2One(hsl.Id.Get(), "")
}

// CreateHrSkillLevel creates a new hr.skill.level model and returns its id.
func (c *Client) CreateHrSkillLevel(hsl *HrSkillLevel) (int64, error) {
	ids, err := c.CreateHrSkillLevels([]*HrSkillLevel{hsl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSkillLevel creates a new hr.skill.level model and returns its id.
func (c *Client) CreateHrSkillLevels(hsls []*HrSkillLevel) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsls {
		vv = append(vv, v)
	}
	return c.Create(HrSkillLevelModel, vv)
}

// UpdateHrSkillLevel updates an existing hr.skill.level record.
func (c *Client) UpdateHrSkillLevel(hsl *HrSkillLevel) error {
	return c.UpdateHrSkillLevels([]int64{hsl.Id.Get()}, hsl)
}

// UpdateHrSkillLevels updates existing hr.skill.level records.
// All records (represented by ids) will be updated by hsl values.
func (c *Client) UpdateHrSkillLevels(ids []int64, hsl *HrSkillLevel) error {
	return c.Update(HrSkillLevelModel, ids, hsl)
}

// DeleteHrSkillLevel deletes an existing hr.skill.level record.
func (c *Client) DeleteHrSkillLevel(id int64) error {
	return c.DeleteHrSkillLevels([]int64{id})
}

// DeleteHrSkillLevels deletes existing hr.skill.level records.
func (c *Client) DeleteHrSkillLevels(ids []int64) error {
	return c.Delete(HrSkillLevelModel, ids)
}

// GetHrSkillLevel gets hr.skill.level existing record.
func (c *Client) GetHrSkillLevel(id int64) (*HrSkillLevel, error) {
	hsls, err := c.GetHrSkillLevels([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsls != nil && len(*hsls) > 0 {
		return &((*hsls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.skill.level not found", id)
}

// GetHrSkillLevels gets hr.skill.level existing records.
func (c *Client) GetHrSkillLevels(ids []int64) (*HrSkillLevels, error) {
	hsls := &HrSkillLevels{}
	if err := c.Read(HrSkillLevelModel, ids, nil, hsls); err != nil {
		return nil, err
	}
	return hsls, nil
}

// FindHrSkillLevel finds hr.skill.level record by querying it with criteria.
func (c *Client) FindHrSkillLevel(criteria *Criteria) (*HrSkillLevel, error) {
	hsls := &HrSkillLevels{}
	if err := c.SearchRead(HrSkillLevelModel, criteria, NewOptions().Limit(1), hsls); err != nil {
		return nil, err
	}
	if hsls != nil && len(*hsls) > 0 {
		return &((*hsls)[0]), nil
	}
	return nil, fmt.Errorf("hr.skill.level was not found with criteria %v", criteria)
}

// FindHrSkillLevels finds hr.skill.level records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkillLevels(criteria *Criteria, options *Options) (*HrSkillLevels, error) {
	hsls := &HrSkillLevels{}
	if err := c.SearchRead(HrSkillLevelModel, criteria, options, hsls); err != nil {
		return nil, err
	}
	return hsls, nil
}

// FindHrSkillLevelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSkillLevelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrSkillLevelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrSkillLevelId finds record id by querying it with criteria.
func (c *Client) FindHrSkillLevelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSkillLevelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.skill.level was not found with criteria %v and options %v", criteria, options)
}
