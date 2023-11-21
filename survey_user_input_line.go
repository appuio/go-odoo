package odoo

import (
	"fmt"
)

// SurveyUserInputLine represents survey.user_input.line model.
type SurveyUserInputLine struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	AnswerIsCorrect   *Bool      `xmlrpc:"answer_is_correct,omptempty"`
	AnswerScore       *Float     `xmlrpc:"answer_score,omptempty"`
	AnswerType        *Selection `xmlrpc:"answer_type,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	MatrixRowId       *Many2One  `xmlrpc:"matrix_row_id,omptempty"`
	PageId            *Many2One  `xmlrpc:"page_id,omptempty"`
	QuestionId        *Many2One  `xmlrpc:"question_id,omptempty"`
	QuestionSequence  *Int       `xmlrpc:"question_sequence,omptempty"`
	Skipped           *Bool      `xmlrpc:"skipped,omptempty"`
	SuggestedAnswerId *Many2One  `xmlrpc:"suggested_answer_id,omptempty"`
	SurveyId          *Many2One  `xmlrpc:"survey_id,omptempty"`
	UserInputId       *Many2One  `xmlrpc:"user_input_id,omptempty"`
	ValueCharBox      *String    `xmlrpc:"value_char_box,omptempty"`
	ValueDate         *Time      `xmlrpc:"value_date,omptempty"`
	ValueDatetime     *Time      `xmlrpc:"value_datetime,omptempty"`
	ValueNumericalBox *Float     `xmlrpc:"value_numerical_box,omptempty"`
	ValueTextBox      *String    `xmlrpc:"value_text_box,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveyUserInputLines represents array of survey.user_input.line model.
type SurveyUserInputLines []SurveyUserInputLine

// SurveyUserInputLineModel is the odoo model name.
const SurveyUserInputLineModel = "survey.user_input.line"

// Many2One convert SurveyUserInputLine to *Many2One.
func (sul *SurveyUserInputLine) Many2One() *Many2One {
	return NewMany2One(sul.Id.Get(), "")
}

// CreateSurveyUserInputLine creates a new survey.user_input.line model and returns its id.
func (c *Client) CreateSurveyUserInputLine(sul *SurveyUserInputLine) (int64, error) {
	ids, err := c.CreateSurveyUserInputLines([]*SurveyUserInputLine{sul})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyUserInputLine creates a new survey.user_input.line model and returns its id.
func (c *Client) CreateSurveyUserInputLines(suls []*SurveyUserInputLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range suls {
		vv = append(vv, v)
	}
	return c.Create(SurveyUserInputLineModel, vv)
}

// UpdateSurveyUserInputLine updates an existing survey.user_input.line record.
func (c *Client) UpdateSurveyUserInputLine(sul *SurveyUserInputLine) error {
	return c.UpdateSurveyUserInputLines([]int64{sul.Id.Get()}, sul)
}

// UpdateSurveyUserInputLines updates existing survey.user_input.line records.
// All records (represented by ids) will be updated by sul values.
func (c *Client) UpdateSurveyUserInputLines(ids []int64, sul *SurveyUserInputLine) error {
	return c.Update(SurveyUserInputLineModel, ids, sul)
}

// DeleteSurveyUserInputLine deletes an existing survey.user_input.line record.
func (c *Client) DeleteSurveyUserInputLine(id int64) error {
	return c.DeleteSurveyUserInputLines([]int64{id})
}

// DeleteSurveyUserInputLines deletes existing survey.user_input.line records.
func (c *Client) DeleteSurveyUserInputLines(ids []int64) error {
	return c.Delete(SurveyUserInputLineModel, ids)
}

// GetSurveyUserInputLine gets survey.user_input.line existing record.
func (c *Client) GetSurveyUserInputLine(id int64) (*SurveyUserInputLine, error) {
	suls, err := c.GetSurveyUserInputLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if suls != nil && len(*suls) > 0 {
		return &((*suls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.user_input.line not found", id)
}

// GetSurveyUserInputLines gets survey.user_input.line existing records.
func (c *Client) GetSurveyUserInputLines(ids []int64) (*SurveyUserInputLines, error) {
	suls := &SurveyUserInputLines{}
	if err := c.Read(SurveyUserInputLineModel, ids, nil, suls); err != nil {
		return nil, err
	}
	return suls, nil
}

// FindSurveyUserInputLine finds survey.user_input.line record by querying it with criteria.
func (c *Client) FindSurveyUserInputLine(criteria *Criteria) (*SurveyUserInputLine, error) {
	suls := &SurveyUserInputLines{}
	if err := c.SearchRead(SurveyUserInputLineModel, criteria, NewOptions().Limit(1), suls); err != nil {
		return nil, err
	}
	if suls != nil && len(*suls) > 0 {
		return &((*suls)[0]), nil
	}
	return nil, fmt.Errorf("survey.user_input.line was not found with criteria %v", criteria)
}

// FindSurveyUserInputLines finds survey.user_input.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputLines(criteria *Criteria, options *Options) (*SurveyUserInputLines, error) {
	suls := &SurveyUserInputLines{}
	if err := c.SearchRead(SurveyUserInputLineModel, criteria, options, suls); err != nil {
		return nil, err
	}
	return suls, nil
}

// FindSurveyUserInputLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyUserInputLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyUserInputLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyUserInputLineId finds record id by querying it with criteria.
func (c *Client) FindSurveyUserInputLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyUserInputLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.user_input.line was not found with criteria %v and options %v", criteria, options)
}
