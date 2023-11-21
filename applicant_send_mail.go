package odoo

import (
	"fmt"
)

// ApplicantSendMail represents applicant.send.mail model.
type ApplicantSendMail struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	ApplicantIds         *Relation `xmlrpc:"applicant_ids,omptempty"`
	AuthorId             *Many2One `xmlrpc:"author_id,omptempty"`
	Body                 *String   `xmlrpc:"body,omptempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omptempty"`
	Lang                 *String   `xmlrpc:"lang,omptempty"`
	RenderModel          *String   `xmlrpc:"render_model,omptempty"`
	Subject              *String   `xmlrpc:"subject,omptempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ApplicantSendMails represents array of applicant.send.mail model.
type ApplicantSendMails []ApplicantSendMail

// ApplicantSendMailModel is the odoo model name.
const ApplicantSendMailModel = "applicant.send.mail"

// Many2One convert ApplicantSendMail to *Many2One.
func (asm *ApplicantSendMail) Many2One() *Many2One {
	return NewMany2One(asm.Id.Get(), "")
}

// CreateApplicantSendMail creates a new applicant.send.mail model and returns its id.
func (c *Client) CreateApplicantSendMail(asm *ApplicantSendMail) (int64, error) {
	ids, err := c.CreateApplicantSendMails([]*ApplicantSendMail{asm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateApplicantSendMail creates a new applicant.send.mail model and returns its id.
func (c *Client) CreateApplicantSendMails(asms []*ApplicantSendMail) ([]int64, error) {
	var vv []interface{}
	for _, v := range asms {
		vv = append(vv, v)
	}
	return c.Create(ApplicantSendMailModel, vv)
}

// UpdateApplicantSendMail updates an existing applicant.send.mail record.
func (c *Client) UpdateApplicantSendMail(asm *ApplicantSendMail) error {
	return c.UpdateApplicantSendMails([]int64{asm.Id.Get()}, asm)
}

// UpdateApplicantSendMails updates existing applicant.send.mail records.
// All records (represented by ids) will be updated by asm values.
func (c *Client) UpdateApplicantSendMails(ids []int64, asm *ApplicantSendMail) error {
	return c.Update(ApplicantSendMailModel, ids, asm)
}

// DeleteApplicantSendMail deletes an existing applicant.send.mail record.
func (c *Client) DeleteApplicantSendMail(id int64) error {
	return c.DeleteApplicantSendMails([]int64{id})
}

// DeleteApplicantSendMails deletes existing applicant.send.mail records.
func (c *Client) DeleteApplicantSendMails(ids []int64) error {
	return c.Delete(ApplicantSendMailModel, ids)
}

// GetApplicantSendMail gets applicant.send.mail existing record.
func (c *Client) GetApplicantSendMail(id int64) (*ApplicantSendMail, error) {
	asms, err := c.GetApplicantSendMails([]int64{id})
	if err != nil {
		return nil, err
	}
	if asms != nil && len(*asms) > 0 {
		return &((*asms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of applicant.send.mail not found", id)
}

// GetApplicantSendMails gets applicant.send.mail existing records.
func (c *Client) GetApplicantSendMails(ids []int64) (*ApplicantSendMails, error) {
	asms := &ApplicantSendMails{}
	if err := c.Read(ApplicantSendMailModel, ids, nil, asms); err != nil {
		return nil, err
	}
	return asms, nil
}

// FindApplicantSendMail finds applicant.send.mail record by querying it with criteria.
func (c *Client) FindApplicantSendMail(criteria *Criteria) (*ApplicantSendMail, error) {
	asms := &ApplicantSendMails{}
	if err := c.SearchRead(ApplicantSendMailModel, criteria, NewOptions().Limit(1), asms); err != nil {
		return nil, err
	}
	if asms != nil && len(*asms) > 0 {
		return &((*asms)[0]), nil
	}
	return nil, fmt.Errorf("applicant.send.mail was not found with criteria %v", criteria)
}

// FindApplicantSendMails finds applicant.send.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindApplicantSendMails(criteria *Criteria, options *Options) (*ApplicantSendMails, error) {
	asms := &ApplicantSendMails{}
	if err := c.SearchRead(ApplicantSendMailModel, criteria, options, asms); err != nil {
		return nil, err
	}
	return asms, nil
}

// FindApplicantSendMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindApplicantSendMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ApplicantSendMailModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindApplicantSendMailId finds record id by querying it with criteria.
func (c *Client) FindApplicantSendMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ApplicantSendMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("applicant.send.mail was not found with criteria %v and options %v", criteria, options)
}
