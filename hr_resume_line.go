package odoo

import (
	"fmt"
)

// HrResumeLine represents hr.resume.line model.
type HrResumeLine struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	ChannelId   *Many2One  `xmlrpc:"channel_id,omptempty"`
	CourseUrl   *String    `xmlrpc:"course_url,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateEnd     *Time      `xmlrpc:"date_end,omptempty"`
	DateStart   *Time      `xmlrpc:"date_start,omptempty"`
	Description *String    `xmlrpc:"description,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	DisplayType *Selection `xmlrpc:"display_type,omptempty"`
	EmployeeId  *Many2One  `xmlrpc:"employee_id,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	LineTypeId  *Many2One  `xmlrpc:"line_type_id,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	SurveyId    *Many2One  `xmlrpc:"survey_id,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrResumeLines represents array of hr.resume.line model.
type HrResumeLines []HrResumeLine

// HrResumeLineModel is the odoo model name.
const HrResumeLineModel = "hr.resume.line"

// Many2One convert HrResumeLine to *Many2One.
func (hrl *HrResumeLine) Many2One() *Many2One {
	return NewMany2One(hrl.Id.Get(), "")
}

// CreateHrResumeLine creates a new hr.resume.line model and returns its id.
func (c *Client) CreateHrResumeLine(hrl *HrResumeLine) (int64, error) {
	ids, err := c.CreateHrResumeLines([]*HrResumeLine{hrl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrResumeLine creates a new hr.resume.line model and returns its id.
func (c *Client) CreateHrResumeLines(hrls []*HrResumeLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range hrls {
		vv = append(vv, v)
	}
	return c.Create(HrResumeLineModel, vv)
}

// UpdateHrResumeLine updates an existing hr.resume.line record.
func (c *Client) UpdateHrResumeLine(hrl *HrResumeLine) error {
	return c.UpdateHrResumeLines([]int64{hrl.Id.Get()}, hrl)
}

// UpdateHrResumeLines updates existing hr.resume.line records.
// All records (represented by ids) will be updated by hrl values.
func (c *Client) UpdateHrResumeLines(ids []int64, hrl *HrResumeLine) error {
	return c.Update(HrResumeLineModel, ids, hrl)
}

// DeleteHrResumeLine deletes an existing hr.resume.line record.
func (c *Client) DeleteHrResumeLine(id int64) error {
	return c.DeleteHrResumeLines([]int64{id})
}

// DeleteHrResumeLines deletes existing hr.resume.line records.
func (c *Client) DeleteHrResumeLines(ids []int64) error {
	return c.Delete(HrResumeLineModel, ids)
}

// GetHrResumeLine gets hr.resume.line existing record.
func (c *Client) GetHrResumeLine(id int64) (*HrResumeLine, error) {
	hrls, err := c.GetHrResumeLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if hrls != nil && len(*hrls) > 0 {
		return &((*hrls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.resume.line not found", id)
}

// GetHrResumeLines gets hr.resume.line existing records.
func (c *Client) GetHrResumeLines(ids []int64) (*HrResumeLines, error) {
	hrls := &HrResumeLines{}
	if err := c.Read(HrResumeLineModel, ids, nil, hrls); err != nil {
		return nil, err
	}
	return hrls, nil
}

// FindHrResumeLine finds hr.resume.line record by querying it with criteria.
func (c *Client) FindHrResumeLine(criteria *Criteria) (*HrResumeLine, error) {
	hrls := &HrResumeLines{}
	if err := c.SearchRead(HrResumeLineModel, criteria, NewOptions().Limit(1), hrls); err != nil {
		return nil, err
	}
	if hrls != nil && len(*hrls) > 0 {
		return &((*hrls)[0]), nil
	}
	return nil, fmt.Errorf("hr.resume.line was not found with criteria %v", criteria)
}

// FindHrResumeLines finds hr.resume.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLines(criteria *Criteria, options *Options) (*HrResumeLines, error) {
	hrls := &HrResumeLines{}
	if err := c.SearchRead(HrResumeLineModel, criteria, options, hrls); err != nil {
		return nil, err
	}
	return hrls, nil
}

// FindHrResumeLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrResumeLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrResumeLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrResumeLineId finds record id by querying it with criteria.
func (c *Client) FindHrResumeLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrResumeLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.resume.line was not found with criteria %v and options %v", criteria, options)
}
