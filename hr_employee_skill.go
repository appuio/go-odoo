package odoo

import (
	"fmt"
)

// HrEmployeeSkill represents hr.employee.skill model.
type HrEmployeeSkill struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId    *Many2One `xmlrpc:"employee_id,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LevelProgress *Int      `xmlrpc:"level_progress,omptempty"`
	SkillId       *Many2One `xmlrpc:"skill_id,omptempty"`
	SkillLevelId  *Many2One `xmlrpc:"skill_level_id,omptempty"`
	SkillTypeId   *Many2One `xmlrpc:"skill_type_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrEmployeeSkills represents array of hr.employee.skill model.
type HrEmployeeSkills []HrEmployeeSkill

// HrEmployeeSkillModel is the odoo model name.
const HrEmployeeSkillModel = "hr.employee.skill"

// Many2One convert HrEmployeeSkill to *Many2One.
func (hes *HrEmployeeSkill) Many2One() *Many2One {
	return NewMany2One(hes.Id.Get(), "")
}

// CreateHrEmployeeSkill creates a new hr.employee.skill model and returns its id.
func (c *Client) CreateHrEmployeeSkill(hes *HrEmployeeSkill) (int64, error) {
	ids, err := c.CreateHrEmployeeSkills([]*HrEmployeeSkill{hes})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeSkill creates a new hr.employee.skill model and returns its id.
func (c *Client) CreateHrEmployeeSkills(hess []*HrEmployeeSkill) ([]int64, error) {
	var vv []interface{}
	for _, v := range hess {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeSkillModel, vv)
}

// UpdateHrEmployeeSkill updates an existing hr.employee.skill record.
func (c *Client) UpdateHrEmployeeSkill(hes *HrEmployeeSkill) error {
	return c.UpdateHrEmployeeSkills([]int64{hes.Id.Get()}, hes)
}

// UpdateHrEmployeeSkills updates existing hr.employee.skill records.
// All records (represented by ids) will be updated by hes values.
func (c *Client) UpdateHrEmployeeSkills(ids []int64, hes *HrEmployeeSkill) error {
	return c.Update(HrEmployeeSkillModel, ids, hes)
}

// DeleteHrEmployeeSkill deletes an existing hr.employee.skill record.
func (c *Client) DeleteHrEmployeeSkill(id int64) error {
	return c.DeleteHrEmployeeSkills([]int64{id})
}

// DeleteHrEmployeeSkills deletes existing hr.employee.skill records.
func (c *Client) DeleteHrEmployeeSkills(ids []int64) error {
	return c.Delete(HrEmployeeSkillModel, ids)
}

// GetHrEmployeeSkill gets hr.employee.skill existing record.
func (c *Client) GetHrEmployeeSkill(id int64) (*HrEmployeeSkill, error) {
	hess, err := c.GetHrEmployeeSkills([]int64{id})
	if err != nil {
		return nil, err
	}
	if hess != nil && len(*hess) > 0 {
		return &((*hess)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.employee.skill not found", id)
}

// GetHrEmployeeSkills gets hr.employee.skill existing records.
func (c *Client) GetHrEmployeeSkills(ids []int64) (*HrEmployeeSkills, error) {
	hess := &HrEmployeeSkills{}
	if err := c.Read(HrEmployeeSkillModel, ids, nil, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrEmployeeSkill finds hr.employee.skill record by querying it with criteria.
func (c *Client) FindHrEmployeeSkill(criteria *Criteria) (*HrEmployeeSkill, error) {
	hess := &HrEmployeeSkills{}
	if err := c.SearchRead(HrEmployeeSkillModel, criteria, NewOptions().Limit(1), hess); err != nil {
		return nil, err
	}
	if hess != nil && len(*hess) > 0 {
		return &((*hess)[0]), nil
	}
	return nil, fmt.Errorf("hr.employee.skill was not found with criteria %v", criteria)
}

// FindHrEmployeeSkills finds hr.employee.skill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkills(criteria *Criteria, options *Options) (*HrEmployeeSkills, error) {
	hess := &HrEmployeeSkills{}
	if err := c.SearchRead(HrEmployeeSkillModel, criteria, options, hess); err != nil {
		return nil, err
	}
	return hess, nil
}

// FindHrEmployeeSkillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeSkillIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrEmployeeSkillModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrEmployeeSkillId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeSkillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeSkillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.employee.skill was not found with criteria %v and options %v", criteria, options)
}
