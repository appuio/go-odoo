package odoo

import (
	"fmt"
)

// SurveyQuestion represents survey.question model.
type SurveyQuestion struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	AnswerDate              *Time      `xmlrpc:"answer_date,omptempty"`
	AnswerDatetime          *Time      `xmlrpc:"answer_datetime,omptempty"`
	AnswerNumericalBox      *Float     `xmlrpc:"answer_numerical_box,omptempty"`
	AnswerScore             *Float     `xmlrpc:"answer_score,omptempty"`
	BackgroundImage         *String    `xmlrpc:"background_image,omptempty"`
	BackgroundImageUrl      *String    `xmlrpc:"background_image_url,omptempty"`
	CommentCountAsAnswer    *Bool      `xmlrpc:"comment_count_as_answer,omptempty"`
	CommentsAllowed         *Bool      `xmlrpc:"comments_allowed,omptempty"`
	CommentsMessage         *String    `xmlrpc:"comments_message,omptempty"`
	ConstrErrorMsg          *String    `xmlrpc:"constr_error_msg,omptempty"`
	ConstrMandatory         *Bool      `xmlrpc:"constr_mandatory,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description             *String    `xmlrpc:"description,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	IsConditional           *Bool      `xmlrpc:"is_conditional,omptempty"`
	IsPage                  *Bool      `xmlrpc:"is_page,omptempty"`
	IsScoredQuestion        *Bool      `xmlrpc:"is_scored_question,omptempty"`
	IsTimeLimited           *Bool      `xmlrpc:"is_time_limited,omptempty"`
	MatrixRowIds            *Relation  `xmlrpc:"matrix_row_ids,omptempty"`
	MatrixSubtype           *Selection `xmlrpc:"matrix_subtype,omptempty"`
	PageId                  *Many2One  `xmlrpc:"page_id,omptempty"`
	QuestionIds             *Relation  `xmlrpc:"question_ids,omptempty"`
	QuestionPlaceholder     *String    `xmlrpc:"question_placeholder,omptempty"`
	QuestionType            *Selection `xmlrpc:"question_type,omptempty"`
	QuestionsSelection      *Selection `xmlrpc:"questions_selection,omptempty"`
	RandomQuestionsCount    *Int       `xmlrpc:"random_questions_count,omptempty"`
	SaveAsEmail             *Bool      `xmlrpc:"save_as_email,omptempty"`
	SaveAsNickname          *Bool      `xmlrpc:"save_as_nickname,omptempty"`
	ScoringType             *Selection `xmlrpc:"scoring_type,omptempty"`
	Sequence                *Int       `xmlrpc:"sequence,omptempty"`
	SuggestedAnswerIds      *Relation  `xmlrpc:"suggested_answer_ids,omptempty"`
	SurveyId                *Many2One  `xmlrpc:"survey_id,omptempty"`
	TimeLimit               *Int       `xmlrpc:"time_limit,omptempty"`
	Title                   *String    `xmlrpc:"title,omptempty"`
	TriggeringAnswerId      *Many2One  `xmlrpc:"triggering_answer_id,omptempty"`
	TriggeringQuestionId    *Many2One  `xmlrpc:"triggering_question_id,omptempty"`
	UserInputLineIds        *Relation  `xmlrpc:"user_input_line_ids,omptempty"`
	ValidationEmail         *Bool      `xmlrpc:"validation_email,omptempty"`
	ValidationErrorMsg      *String    `xmlrpc:"validation_error_msg,omptempty"`
	ValidationLengthMax     *Int       `xmlrpc:"validation_length_max,omptempty"`
	ValidationLengthMin     *Int       `xmlrpc:"validation_length_min,omptempty"`
	ValidationMaxDate       *Time      `xmlrpc:"validation_max_date,omptempty"`
	ValidationMaxDatetime   *Time      `xmlrpc:"validation_max_datetime,omptempty"`
	ValidationMaxFloatValue *Float     `xmlrpc:"validation_max_float_value,omptempty"`
	ValidationMinDate       *Time      `xmlrpc:"validation_min_date,omptempty"`
	ValidationMinDatetime   *Time      `xmlrpc:"validation_min_datetime,omptempty"`
	ValidationMinFloatValue *Float     `xmlrpc:"validation_min_float_value,omptempty"`
	ValidationRequired      *Bool      `xmlrpc:"validation_required,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SurveyQuestions represents array of survey.question model.
type SurveyQuestions []SurveyQuestion

// SurveyQuestionModel is the odoo model name.
const SurveyQuestionModel = "survey.question"

// Many2One convert SurveyQuestion to *Many2One.
func (sq *SurveyQuestion) Many2One() *Many2One {
	return NewMany2One(sq.Id.Get(), "")
}

// CreateSurveyQuestion creates a new survey.question model and returns its id.
func (c *Client) CreateSurveyQuestion(sq *SurveyQuestion) (int64, error) {
	ids, err := c.CreateSurveyQuestions([]*SurveyQuestion{sq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSurveyQuestion creates a new survey.question model and returns its id.
func (c *Client) CreateSurveyQuestions(sqs []*SurveyQuestion) ([]int64, error) {
	var vv []interface{}
	for _, v := range sqs {
		vv = append(vv, v)
	}
	return c.Create(SurveyQuestionModel, vv)
}

// UpdateSurveyQuestion updates an existing survey.question record.
func (c *Client) UpdateSurveyQuestion(sq *SurveyQuestion) error {
	return c.UpdateSurveyQuestions([]int64{sq.Id.Get()}, sq)
}

// UpdateSurveyQuestions updates existing survey.question records.
// All records (represented by ids) will be updated by sq values.
func (c *Client) UpdateSurveyQuestions(ids []int64, sq *SurveyQuestion) error {
	return c.Update(SurveyQuestionModel, ids, sq)
}

// DeleteSurveyQuestion deletes an existing survey.question record.
func (c *Client) DeleteSurveyQuestion(id int64) error {
	return c.DeleteSurveyQuestions([]int64{id})
}

// DeleteSurveyQuestions deletes existing survey.question records.
func (c *Client) DeleteSurveyQuestions(ids []int64) error {
	return c.Delete(SurveyQuestionModel, ids)
}

// GetSurveyQuestion gets survey.question existing record.
func (c *Client) GetSurveyQuestion(id int64) (*SurveyQuestion, error) {
	sqs, err := c.GetSurveyQuestions([]int64{id})
	if err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of survey.question not found", id)
}

// GetSurveyQuestions gets survey.question existing records.
func (c *Client) GetSurveyQuestions(ids []int64) (*SurveyQuestions, error) {
	sqs := &SurveyQuestions{}
	if err := c.Read(SurveyQuestionModel, ids, nil, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindSurveyQuestion finds survey.question record by querying it with criteria.
func (c *Client) FindSurveyQuestion(criteria *Criteria) (*SurveyQuestion, error) {
	sqs := &SurveyQuestions{}
	if err := c.SearchRead(SurveyQuestionModel, criteria, NewOptions().Limit(1), sqs); err != nil {
		return nil, err
	}
	if sqs != nil && len(*sqs) > 0 {
		return &((*sqs)[0]), nil
	}
	return nil, fmt.Errorf("survey.question was not found with criteria %v", criteria)
}

// FindSurveyQuestions finds survey.question records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestions(criteria *Criteria, options *Options) (*SurveyQuestions, error) {
	sqs := &SurveyQuestions{}
	if err := c.SearchRead(SurveyQuestionModel, criteria, options, sqs); err != nil {
		return nil, err
	}
	return sqs, nil
}

// FindSurveyQuestionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSurveyQuestionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SurveyQuestionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSurveyQuestionId finds record id by querying it with criteria.
func (c *Client) FindSurveyQuestionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SurveyQuestionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("survey.question was not found with criteria %v and options %v", criteria, options)
}
