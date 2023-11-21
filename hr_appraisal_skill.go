package odoo

import (
	"fmt"
)

// HrAppraisalSkill represents hr.appraisal.skill model.
type HrAppraisalSkill struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	AppraisalId          *Many2One `xmlrpc:"appraisal_id,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	EmployeeId           *Many2One `xmlrpc:"employee_id,omptempty"`
	EmployeeSkillId      *Many2One `xmlrpc:"employee_skill_id,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	Justification        *String   `xmlrpc:"justification,omptempty"`
	LevelProgress        *Int      `xmlrpc:"level_progress,omptempty"`
	ManagerIds           *Relation `xmlrpc:"manager_ids,omptempty"`
	PreviousSkillLevelId *Many2One `xmlrpc:"previous_skill_level_id,omptempty"`
	SkillId              *Many2One `xmlrpc:"skill_id,omptempty"`
	SkillLevelId         *Many2One `xmlrpc:"skill_level_id,omptempty"`
	SkillTypeId          *Many2One `xmlrpc:"skill_type_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrAppraisalSkills represents array of hr.appraisal.skill model.
type HrAppraisalSkills []HrAppraisalSkill

// HrAppraisalSkillModel is the odoo model name.
const HrAppraisalSkillModel = "hr.appraisal.skill"

// Many2One convert HrAppraisalSkill to *Many2One.
func (has *HrAppraisalSkill) Many2One() *Many2One {
	return NewMany2One(has.Id.Get(), "")
}

// CreateHrAppraisalSkill creates a new hr.appraisal.skill model and returns its id.
func (c *Client) CreateHrAppraisalSkill(has *HrAppraisalSkill) (int64, error) {
	ids, err := c.CreateHrAppraisalSkills([]*HrAppraisalSkill{has})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrAppraisalSkill creates a new hr.appraisal.skill model and returns its id.
func (c *Client) CreateHrAppraisalSkills(hass []*HrAppraisalSkill) ([]int64, error) {
	var vv []interface{}
	for _, v := range hass {
		vv = append(vv, v)
	}
	return c.Create(HrAppraisalSkillModel, vv)
}

// UpdateHrAppraisalSkill updates an existing hr.appraisal.skill record.
func (c *Client) UpdateHrAppraisalSkill(has *HrAppraisalSkill) error {
	return c.UpdateHrAppraisalSkills([]int64{has.Id.Get()}, has)
}

// UpdateHrAppraisalSkills updates existing hr.appraisal.skill records.
// All records (represented by ids) will be updated by has values.
func (c *Client) UpdateHrAppraisalSkills(ids []int64, has *HrAppraisalSkill) error {
	return c.Update(HrAppraisalSkillModel, ids, has)
}

// DeleteHrAppraisalSkill deletes an existing hr.appraisal.skill record.
func (c *Client) DeleteHrAppraisalSkill(id int64) error {
	return c.DeleteHrAppraisalSkills([]int64{id})
}

// DeleteHrAppraisalSkills deletes existing hr.appraisal.skill records.
func (c *Client) DeleteHrAppraisalSkills(ids []int64) error {
	return c.Delete(HrAppraisalSkillModel, ids)
}

// GetHrAppraisalSkill gets hr.appraisal.skill existing record.
func (c *Client) GetHrAppraisalSkill(id int64) (*HrAppraisalSkill, error) {
	hass, err := c.GetHrAppraisalSkills([]int64{id})
	if err != nil {
		return nil, err
	}
	if hass != nil && len(*hass) > 0 {
		return &((*hass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.appraisal.skill not found", id)
}

// GetHrAppraisalSkills gets hr.appraisal.skill existing records.
func (c *Client) GetHrAppraisalSkills(ids []int64) (*HrAppraisalSkills, error) {
	hass := &HrAppraisalSkills{}
	if err := c.Read(HrAppraisalSkillModel, ids, nil, hass); err != nil {
		return nil, err
	}
	return hass, nil
}

// FindHrAppraisalSkill finds hr.appraisal.skill record by querying it with criteria.
func (c *Client) FindHrAppraisalSkill(criteria *Criteria) (*HrAppraisalSkill, error) {
	hass := &HrAppraisalSkills{}
	if err := c.SearchRead(HrAppraisalSkillModel, criteria, NewOptions().Limit(1), hass); err != nil {
		return nil, err
	}
	if hass != nil && len(*hass) > 0 {
		return &((*hass)[0]), nil
	}
	return nil, fmt.Errorf("hr.appraisal.skill was not found with criteria %v", criteria)
}

// FindHrAppraisalSkills finds hr.appraisal.skill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalSkills(criteria *Criteria, options *Options) (*HrAppraisalSkills, error) {
	hass := &HrAppraisalSkills{}
	if err := c.SearchRead(HrAppraisalSkillModel, criteria, options, hass); err != nil {
		return nil, err
	}
	return hass, nil
}

// FindHrAppraisalSkillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrAppraisalSkillIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrAppraisalSkillModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrAppraisalSkillId finds record id by querying it with criteria.
func (c *Client) FindHrAppraisalSkillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrAppraisalSkillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.appraisal.skill was not found with criteria %v and options %v", criteria, options)
}
