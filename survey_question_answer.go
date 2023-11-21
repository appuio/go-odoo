package odoo

import (
	"fmt"
)

// SurveyQuestionAnswer represents survey.question.answer model.
type SurveyQuestionAnswer struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	AnswerScore        *Float     `xmlrpc:"answer_score,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	IsCorrect          *Bool      `xmlrpc:"is_correct,omptempty"`
	MatrixQuestionId   *Many2One  `xmlrpc:"matrix_question_id,omptempty"`
	QuestionId         *Many2One  `xmlrpc:"question_id,omptempty"`
	QuestionType       *Selection `xmlrpc:"question_type,omptempty"`
	ScoringType        *Selection `xmlrpc:"scoring_type,omptempty"`
	Sequence           *Int       `xmlrpc:"sequence,omptempty"`
	SurveyId           *Many2One  `xmlrpc:"survey_id,omptempty"`
	Value              *String    `xmlrpc:"value,omptempty"`
	ValueImage         *String    `xmlrpc:"value_image,omptempty"`
	ValueImageFilename *String    `xmlrpc:"value_image_filename,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveyQuestionAnswers represents array of survey.question.answer model.
type SurveyQuestionAnswers []SurveyQuestionAnswer

// SurveyQuestionAnswerModel is the odoo model name.
const SurveyQuestionAnswerModel = "survey.question.answer"

// Many2One convert SurveyQuestionAnswer to *Many2One.
func (sqa *SurveyQuestionAnswer) Many2One() *Many2One {
	return NewMany2One(sqa.Id.Get(), "")
}

// CreateSurveyQuestionAnswer creates a new survey.question.answer model and returns its id.
func (c *Client) CreateSurveyQuestionAnswer(sqa *SurveyQuestionAnswer) (int64, error) {
	ids, err := c.CreateSurveyQuestionAnswers([]*SurveyQuestionAnswer{sqa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyQuestionAnswer creates a new survey.question.answer model and returns its id.
func (c *Client) CreateSurveyQuestionAnswers(sqas []*SurveyQuestionAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqas {
		vv = append(vv, v)
	}
	return c.Create(SurveyQuestionAnswerModel, vv)
}

// UpdateSurveyQuestionAnswer updates an existing survey.question.answer record.
func (c *Client) UpdateSurveyQuestionAnswer(sqa *SurveyQuestionAnswer) error {
	return c.UpdateSurveyQuestionAnswers([]int64{sqa.Id.Get()}, sqa)
}

// UpdateSurveyQuestionAnswers updates existing survey.question.answer records.
// All records (represented by ids) will be updated by sqa values.
func (c *Client) UpdateSurveyQuestionAnswers(ids []int64, sqa *SurveyQuestionAnswer) error {
	return c.Update(SurveyQuestionAnswerModel, ids, sqa)
}

// DeleteSurveyQuestionAnswer deletes an existing survey.question.answer record.
func (c *Client) DeleteSurveyQuestionAnswer(id int64) error {
	return c.DeleteSurveyQuestionAnswers([]int64{id})
}

// DeleteSurveyQuestionAnswers deletes existing survey.question.answer records.
func (c *Client) DeleteSurveyQuestionAnswers(ids []int64) error {
	return c.Delete(SurveyQuestionAnswerModel, ids)
}

// GetSurveyQuestionAnswer gets survey.question.answer existing record.
func (c *Client) GetSurveyQuestionAnswer(id int64) (*SurveyQuestionAnswer, error) {
	sqas, err := c.GetSurveyQuestionAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	if sqas != nil && len(*sqas) > 0 {
		return &((*sqas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.question.answer not found", id)
}

// GetSurveyQuestionAnswers gets survey.question.answer existing records.
func (c *Client) GetSurveyQuestionAnswers(ids []int64) (*SurveyQuestionAnswers, error) {
	sqas := &SurveyQuestionAnswers{}
	if err := c.Read(SurveyQuestionAnswerModel, ids, nil, sqas); err != nil {
		return nil, err
	}
	return sqas, nil
}

// FindSurveyQuestionAnswer finds survey.question.answer record by querying it with criteria.
func (c *Client) FindSurveyQuestionAnswer(criteria *Criteria) (*SurveyQuestionAnswer, error) {
	sqas := &SurveyQuestionAnswers{}
	if err := c.SearchRead(SurveyQuestionAnswerModel, criteria, NewOptions().Limit(1), sqas); err != nil {
		return nil, err
	}
	if sqas != nil && len(*sqas) > 0 {
		return &((*sqas)[0]), nil
	}
	return nil, fmt.Errorf("survey.question.answer was not found with criteria %v", criteria)
}

// FindSurveyQuestionAnswers finds survey.question.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestionAnswers(criteria *Criteria, options *Options) (*SurveyQuestionAnswers, error) {
	sqas := &SurveyQuestionAnswers{}
	if err := c.SearchRead(SurveyQuestionAnswerModel, criteria, options, sqas); err != nil {
		return nil, err
	}
	return sqas, nil
}

// FindSurveyQuestionAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestionAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyQuestionAnswerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyQuestionAnswerId finds record id by querying it with criteria.
func (c *Client) FindSurveyQuestionAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyQuestionAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.question.answer was not found with criteria %v and options %v", criteria, options)
}
