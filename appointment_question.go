package odoo

import (
	"fmt"
)

// AppointmentQuestion represents appointment.question model.
type AppointmentQuestion struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	AnswerIds         *Relation  `xmlrpc:"answer_ids,omptempty"`
	AnswerInputIds    *Relation  `xmlrpc:"answer_input_ids,omptempty"`
	AppointmentTypeId *Many2One  `xmlrpc:"appointment_type_id,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	Placeholder       *String    `xmlrpc:"placeholder,omptempty"`
	QuestionRequired  *Bool      `xmlrpc:"question_required,omptempty"`
	QuestionType      *Selection `xmlrpc:"question_type,omptempty"`
	Sequence          *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AppointmentQuestions represents array of appointment.question model.
type AppointmentQuestions []AppointmentQuestion

// AppointmentQuestionModel is the odoo model name.
const AppointmentQuestionModel = "appointment.question"

// Many2One convert AppointmentQuestion to *Many2One.
func (aq *AppointmentQuestion) Many2One() *Many2One {
	return NewMany2One(aq.Id.Get(), "")
}

// CreateAppointmentQuestion creates a new appointment.question model and returns its id.
func (c *Client) CreateAppointmentQuestion(aq *AppointmentQuestion) (int64, error) {
	ids, err := c.CreateAppointmentQuestions([]*AppointmentQuestion{aq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAppointmentQuestion creates a new appointment.question model and returns its id.
func (c *Client) CreateAppointmentQuestions(aqs []*AppointmentQuestion) ([]int64, error) {
	var vv []interface{}
	for _, v := range aqs {
		vv = append(vv, v)
	}
	return c.Create(AppointmentQuestionModel, vv)
}

// UpdateAppointmentQuestion updates an existing appointment.question record.
func (c *Client) UpdateAppointmentQuestion(aq *AppointmentQuestion) error {
	return c.UpdateAppointmentQuestions([]int64{aq.Id.Get()}, aq)
}

// UpdateAppointmentQuestions updates existing appointment.question records.
// All records (represented by ids) will be updated by aq values.
func (c *Client) UpdateAppointmentQuestions(ids []int64, aq *AppointmentQuestion) error {
	return c.Update(AppointmentQuestionModel, ids, aq)
}

// DeleteAppointmentQuestion deletes an existing appointment.question record.
func (c *Client) DeleteAppointmentQuestion(id int64) error {
	return c.DeleteAppointmentQuestions([]int64{id})
}

// DeleteAppointmentQuestions deletes existing appointment.question records.
func (c *Client) DeleteAppointmentQuestions(ids []int64) error {
	return c.Delete(AppointmentQuestionModel, ids)
}

// GetAppointmentQuestion gets appointment.question existing record.
func (c *Client) GetAppointmentQuestion(id int64) (*AppointmentQuestion, error) {
	aqs, err := c.GetAppointmentQuestions([]int64{id})
	if err != nil {
		return nil, err
	}
	if aqs != nil && len(*aqs) > 0 {
		return &((*aqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of appointment.question not found", id)
}

// GetAppointmentQuestions gets appointment.question existing records.
func (c *Client) GetAppointmentQuestions(ids []int64) (*AppointmentQuestions, error) {
	aqs := &AppointmentQuestions{}
	if err := c.Read(AppointmentQuestionModel, ids, nil, aqs); err != nil {
		return nil, err
	}
	return aqs, nil
}

// FindAppointmentQuestion finds appointment.question record by querying it with criteria.
func (c *Client) FindAppointmentQuestion(criteria *Criteria) (*AppointmentQuestion, error) {
	aqs := &AppointmentQuestions{}
	if err := c.SearchRead(AppointmentQuestionModel, criteria, NewOptions().Limit(1), aqs); err != nil {
		return nil, err
	}
	if aqs != nil && len(*aqs) > 0 {
		return &((*aqs)[0]), nil
	}
	return nil, fmt.Errorf("appointment.question was not found with criteria %v", criteria)
}

// FindAppointmentQuestions finds appointment.question records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentQuestions(criteria *Criteria, options *Options) (*AppointmentQuestions, error) {
	aqs := &AppointmentQuestions{}
	if err := c.SearchRead(AppointmentQuestionModel, criteria, options, aqs); err != nil {
		return nil, err
	}
	return aqs, nil
}

// FindAppointmentQuestionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAppointmentQuestionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AppointmentQuestionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAppointmentQuestionId finds record id by querying it with criteria.
func (c *Client) FindAppointmentQuestionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AppointmentQuestionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("appointment.question was not found with criteria %v and options %v", criteria, options)
}
