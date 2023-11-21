package odoo

import (
	"fmt"
)

// ChatbotScriptAnswer represents chatbot.script.answer model.
type ChatbotScriptAnswer struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	ChatbotScriptId *Many2One `xmlrpc:"chatbot_script_id,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	RedirectLink    *String   `xmlrpc:"redirect_link,omptempty"`
	ScriptStepId    *Many2One `xmlrpc:"script_step_id,omptempty"`
	Sequence        *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ChatbotScriptAnswers represents array of chatbot.script.answer model.
type ChatbotScriptAnswers []ChatbotScriptAnswer

// ChatbotScriptAnswerModel is the odoo model name.
const ChatbotScriptAnswerModel = "chatbot.script.answer"

// Many2One convert ChatbotScriptAnswer to *Many2One.
func (csa *ChatbotScriptAnswer) Many2One() *Many2One {
	return NewMany2One(csa.Id.Get(), "")
}

// CreateChatbotScriptAnswer creates a new chatbot.script.answer model and returns its id.
func (c *Client) CreateChatbotScriptAnswer(csa *ChatbotScriptAnswer) (int64, error) {
	ids, err := c.CreateChatbotScriptAnswers([]*ChatbotScriptAnswer{csa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChatbotScriptAnswer creates a new chatbot.script.answer model and returns its id.
func (c *Client) CreateChatbotScriptAnswers(csas []*ChatbotScriptAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range csas {
		vv = append(vv, v)
	}
	return c.Create(ChatbotScriptAnswerModel, vv)
}

// UpdateChatbotScriptAnswer updates an existing chatbot.script.answer record.
func (c *Client) UpdateChatbotScriptAnswer(csa *ChatbotScriptAnswer) error {
	return c.UpdateChatbotScriptAnswers([]int64{csa.Id.Get()}, csa)
}

// UpdateChatbotScriptAnswers updates existing chatbot.script.answer records.
// All records (represented by ids) will be updated by csa values.
func (c *Client) UpdateChatbotScriptAnswers(ids []int64, csa *ChatbotScriptAnswer) error {
	return c.Update(ChatbotScriptAnswerModel, ids, csa)
}

// DeleteChatbotScriptAnswer deletes an existing chatbot.script.answer record.
func (c *Client) DeleteChatbotScriptAnswer(id int64) error {
	return c.DeleteChatbotScriptAnswers([]int64{id})
}

// DeleteChatbotScriptAnswers deletes existing chatbot.script.answer records.
func (c *Client) DeleteChatbotScriptAnswers(ids []int64) error {
	return c.Delete(ChatbotScriptAnswerModel, ids)
}

// GetChatbotScriptAnswer gets chatbot.script.answer existing record.
func (c *Client) GetChatbotScriptAnswer(id int64) (*ChatbotScriptAnswer, error) {
	csas, err := c.GetChatbotScriptAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	if csas != nil && len(*csas) > 0 {
		return &((*csas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of chatbot.script.answer not found", id)
}

// GetChatbotScriptAnswers gets chatbot.script.answer existing records.
func (c *Client) GetChatbotScriptAnswers(ids []int64) (*ChatbotScriptAnswers, error) {
	csas := &ChatbotScriptAnswers{}
	if err := c.Read(ChatbotScriptAnswerModel, ids, nil, csas); err != nil {
		return nil, err
	}
	return csas, nil
}

// FindChatbotScriptAnswer finds chatbot.script.answer record by querying it with criteria.
func (c *Client) FindChatbotScriptAnswer(criteria *Criteria) (*ChatbotScriptAnswer, error) {
	csas := &ChatbotScriptAnswers{}
	if err := c.SearchRead(ChatbotScriptAnswerModel, criteria, NewOptions().Limit(1), csas); err != nil {
		return nil, err
	}
	if csas != nil && len(*csas) > 0 {
		return &((*csas)[0]), nil
	}
	return nil, fmt.Errorf("chatbot.script.answer was not found with criteria %v", criteria)
}

// FindChatbotScriptAnswers finds chatbot.script.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScriptAnswers(criteria *Criteria, options *Options) (*ChatbotScriptAnswers, error) {
	csas := &ChatbotScriptAnswers{}
	if err := c.SearchRead(ChatbotScriptAnswerModel, criteria, options, csas); err != nil {
		return nil, err
	}
	return csas, nil
}

// FindChatbotScriptAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScriptAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ChatbotScriptAnswerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindChatbotScriptAnswerId finds record id by querying it with criteria.
func (c *Client) FindChatbotScriptAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChatbotScriptAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("chatbot.script.answer was not found with criteria %v and options %v", criteria, options)
}
