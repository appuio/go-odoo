package odoo

import (
	"fmt"
)

// AppraisalAskFeedback represents appraisal.ask.feedback model.
type AppraisalAskFeedback struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	AppraisalId          *Many2One `xmlrpc:"appraisal_id,omptempty"`
	AttachmentIds        *Relation `xmlrpc:"attachment_ids,omptempty"`
	AuthorId             *Many2One `xmlrpc:"author_id,omptempty"`
	Body                 *String   `xmlrpc:"body,omptempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	Deadline             *Time     `xmlrpc:"deadline,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	EmailFrom            *String   `xmlrpc:"email_from,omptempty"`
	EmployeeId           *Many2One `xmlrpc:"employee_id,omptempty"`
	EmployeeIds          *Relation `xmlrpc:"employee_ids,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omptempty"`
	Lang                 *String   `xmlrpc:"lang,omptempty"`
	RenderModel          *String   `xmlrpc:"render_model,omptempty"`
	Subject              *String   `xmlrpc:"subject,omptempty"`
	SurveyTemplateId     *Many2One `xmlrpc:"survey_template_id,omptempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omptempty"`
	UserBody             *String   `xmlrpc:"user_body,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AppraisalAskFeedbacks represents array of appraisal.ask.feedback model.
type AppraisalAskFeedbacks []AppraisalAskFeedback

// AppraisalAskFeedbackModel is the odoo model name.
const AppraisalAskFeedbackModel = "appraisal.ask.feedback"

// Many2One convert AppraisalAskFeedback to *Many2One.
func (aaf *AppraisalAskFeedback) Many2One() *Many2One {
	return NewMany2One(aaf.Id.Get(), "")
}

// CreateAppraisalAskFeedback creates a new appraisal.ask.feedback model and returns its id.
func (c *Client) CreateAppraisalAskFeedback(aaf *AppraisalAskFeedback) (int64, error) {
	ids, err := c.CreateAppraisalAskFeedbacks([]*AppraisalAskFeedback{aaf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppraisalAskFeedback creates a new appraisal.ask.feedback model and returns its id.
func (c *Client) CreateAppraisalAskFeedbacks(aafs []*AppraisalAskFeedback) ([]int64, error) {
	var vv []interface{}
	for _, v := range aafs {
		vv = append(vv, v)
	}
	return c.Create(AppraisalAskFeedbackModel, vv)
}

// UpdateAppraisalAskFeedback updates an existing appraisal.ask.feedback record.
func (c *Client) UpdateAppraisalAskFeedback(aaf *AppraisalAskFeedback) error {
	return c.UpdateAppraisalAskFeedbacks([]int64{aaf.Id.Get()}, aaf)
}

// UpdateAppraisalAskFeedbacks updates existing appraisal.ask.feedback records.
// All records (represented by ids) will be updated by aaf values.
func (c *Client) UpdateAppraisalAskFeedbacks(ids []int64, aaf *AppraisalAskFeedback) error {
	return c.Update(AppraisalAskFeedbackModel, ids, aaf)
}

// DeleteAppraisalAskFeedback deletes an existing appraisal.ask.feedback record.
func (c *Client) DeleteAppraisalAskFeedback(id int64) error {
	return c.DeleteAppraisalAskFeedbacks([]int64{id})
}

// DeleteAppraisalAskFeedbacks deletes existing appraisal.ask.feedback records.
func (c *Client) DeleteAppraisalAskFeedbacks(ids []int64) error {
	return c.Delete(AppraisalAskFeedbackModel, ids)
}

// GetAppraisalAskFeedback gets appraisal.ask.feedback existing record.
func (c *Client) GetAppraisalAskFeedback(id int64) (*AppraisalAskFeedback, error) {
	aafs, err := c.GetAppraisalAskFeedbacks([]int64{id})
	if err != nil {
		return nil, err
	}
	if aafs != nil && len(*aafs) > 0 {
		return &((*aafs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of appraisal.ask.feedback not found", id)
}

// GetAppraisalAskFeedbacks gets appraisal.ask.feedback existing records.
func (c *Client) GetAppraisalAskFeedbacks(ids []int64) (*AppraisalAskFeedbacks, error) {
	aafs := &AppraisalAskFeedbacks{}
	if err := c.Read(AppraisalAskFeedbackModel, ids, nil, aafs); err != nil {
		return nil, err
	}
	return aafs, nil
}

// FindAppraisalAskFeedback finds appraisal.ask.feedback record by querying it with criteria.
func (c *Client) FindAppraisalAskFeedback(criteria *Criteria) (*AppraisalAskFeedback, error) {
	aafs := &AppraisalAskFeedbacks{}
	if err := c.SearchRead(AppraisalAskFeedbackModel, criteria, NewOptions().Limit(1), aafs); err != nil {
		return nil, err
	}
	if aafs != nil && len(*aafs) > 0 {
		return &((*aafs)[0]), nil
	}
	return nil, fmt.Errorf("appraisal.ask.feedback was not found with criteria %v", criteria)
}

// FindAppraisalAskFeedbacks finds appraisal.ask.feedback records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppraisalAskFeedbacks(criteria *Criteria, options *Options) (*AppraisalAskFeedbacks, error) {
	aafs := &AppraisalAskFeedbacks{}
	if err := c.SearchRead(AppraisalAskFeedbackModel, criteria, options, aafs); err != nil {
		return nil, err
	}
	return aafs, nil
}

// FindAppraisalAskFeedbackIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppraisalAskFeedbackIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AppraisalAskFeedbackModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAppraisalAskFeedbackId finds record id by querying it with criteria.
func (c *Client) FindAppraisalAskFeedbackId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppraisalAskFeedbackModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("appraisal.ask.feedback was not found with criteria %v and options %v", criteria, options)
}
