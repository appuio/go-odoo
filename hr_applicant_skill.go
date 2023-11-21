package odoo

import (
	"fmt"
)

// HrApplicantSkill represents hr.applicant.skill model.
type HrApplicantSkill struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	ApplicantId   *Many2One `xmlrpc:"applicant_id,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LevelProgress *Int      `xmlrpc:"level_progress,omptempty"`
	SkillId       *Many2One `xmlrpc:"skill_id,omptempty"`
	SkillLevelId  *Many2One `xmlrpc:"skill_level_id,omptempty"`
	SkillTypeId   *Many2One `xmlrpc:"skill_type_id,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HrApplicantSkills represents array of hr.applicant.skill model.
type HrApplicantSkills []HrApplicantSkill

// HrApplicantSkillModel is the odoo model name.
const HrApplicantSkillModel = "hr.applicant.skill"

// Many2One convert HrApplicantSkill to *Many2One.
func (has *HrApplicantSkill) Many2One() *Many2One {
	return NewMany2One(has.Id.Get(), "")
}

// CreateHrApplicantSkill creates a new hr.applicant.skill model and returns its id.
func (c *Client) CreateHrApplicantSkill(has *HrApplicantSkill) (int64, error) {
	ids, err := c.CreateHrApplicantSkills([]*HrApplicantSkill{has})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrApplicantSkill creates a new hr.applicant.skill model and returns its id.
func (c *Client) CreateHrApplicantSkills(hass []*HrApplicantSkill) ([]int64, error) {
	var vv []interface{}
	for _, v := range hass {
		vv = append(vv, v)
	}
	return c.Create(HrApplicantSkillModel, vv)
}

// UpdateHrApplicantSkill updates an existing hr.applicant.skill record.
func (c *Client) UpdateHrApplicantSkill(has *HrApplicantSkill) error {
	return c.UpdateHrApplicantSkills([]int64{has.Id.Get()}, has)
}

// UpdateHrApplicantSkills updates existing hr.applicant.skill records.
// All records (represented by ids) will be updated by has values.
func (c *Client) UpdateHrApplicantSkills(ids []int64, has *HrApplicantSkill) error {
	return c.Update(HrApplicantSkillModel, ids, has)
}

// DeleteHrApplicantSkill deletes an existing hr.applicant.skill record.
func (c *Client) DeleteHrApplicantSkill(id int64) error {
	return c.DeleteHrApplicantSkills([]int64{id})
}

// DeleteHrApplicantSkills deletes existing hr.applicant.skill records.
func (c *Client) DeleteHrApplicantSkills(ids []int64) error {
	return c.Delete(HrApplicantSkillModel, ids)
}

// GetHrApplicantSkill gets hr.applicant.skill existing record.
func (c *Client) GetHrApplicantSkill(id int64) (*HrApplicantSkill, error) {
	hass, err := c.GetHrApplicantSkills([]int64{id})
	if err != nil {
		return nil, err
	}
	if hass != nil && len(*hass) > 0 {
		return &((*hass)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.applicant.skill not found", id)
}

// GetHrApplicantSkills gets hr.applicant.skill existing records.
func (c *Client) GetHrApplicantSkills(ids []int64) (*HrApplicantSkills, error) {
	hass := &HrApplicantSkills{}
	if err := c.Read(HrApplicantSkillModel, ids, nil, hass); err != nil {
		return nil, err
	}
	return hass, nil
}

// FindHrApplicantSkill finds hr.applicant.skill record by querying it with criteria.
func (c *Client) FindHrApplicantSkill(criteria *Criteria) (*HrApplicantSkill, error) {
	hass := &HrApplicantSkills{}
	if err := c.SearchRead(HrApplicantSkillModel, criteria, NewOptions().Limit(1), hass); err != nil {
		return nil, err
	}
	if hass != nil && len(*hass) > 0 {
		return &((*hass)[0]), nil
	}
	return nil, fmt.Errorf("hr.applicant.skill was not found with criteria %v", criteria)
}

// FindHrApplicantSkills finds hr.applicant.skill records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantSkills(criteria *Criteria, options *Options) (*HrApplicantSkills, error) {
	hass := &HrApplicantSkills{}
	if err := c.SearchRead(HrApplicantSkillModel, criteria, options, hass); err != nil {
		return nil, err
	}
	return hass, nil
}

// FindHrApplicantSkillIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrApplicantSkillIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrApplicantSkillModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrApplicantSkillId finds record id by querying it with criteria.
func (c *Client) FindHrApplicantSkillId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrApplicantSkillModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.applicant.skill was not found with criteria %v and options %v", criteria, options)
}
