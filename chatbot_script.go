package odoo

import (
	"fmt"
)

// ChatbotScript represents chatbot.script model.
type ChatbotScript struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	Active               *Bool      `xmlrpc:"active,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	FirstStepWarning     *Selection `xmlrpc:"first_step_warning,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	Image1024            *String    `xmlrpc:"image_1024,omptempty"`
	Image128             *String    `xmlrpc:"image_128,omptempty"`
	Image1920            *String    `xmlrpc:"image_1920,omptempty"`
	Image256             *String    `xmlrpc:"image_256,omptempty"`
	Image512             *String    `xmlrpc:"image_512,omptempty"`
	LeadCount            *Int       `xmlrpc:"lead_count,omptempty"`
	LivechatChannelCount *Int       `xmlrpc:"livechat_channel_count,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	OperatorPartnerId    *Many2One  `xmlrpc:"operator_partner_id,omptempty"`
	ScriptStepIds        *Relation  `xmlrpc:"script_step_ids,omptempty"`
	SourceId             *Many2One  `xmlrpc:"source_id,omptempty"`
	TicketCount          *Int       `xmlrpc:"ticket_count,omptempty"`
	Title                *String    `xmlrpc:"title,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ChatbotScripts represents array of chatbot.script model.
type ChatbotScripts []ChatbotScript

// ChatbotScriptModel is the odoo model name.
const ChatbotScriptModel = "chatbot.script"

// Many2One convert ChatbotScript to *Many2One.
func (cs *ChatbotScript) Many2One() *Many2One {
	return NewMany2One(cs.Id.Get(), "")
}

// CreateChatbotScript creates a new chatbot.script model and returns its id.
func (c *Client) CreateChatbotScript(cs *ChatbotScript) (int64, error) {
	ids, err := c.CreateChatbotScripts([]*ChatbotScript{cs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChatbotScript creates a new chatbot.script model and returns its id.
func (c *Client) CreateChatbotScripts(css []*ChatbotScript) ([]int64, error) {
	var vv []interface{}
	for _, v := range css {
		vv = append(vv, v)
	}
	return c.Create(ChatbotScriptModel, vv)
}

// UpdateChatbotScript updates an existing chatbot.script record.
func (c *Client) UpdateChatbotScript(cs *ChatbotScript) error {
	return c.UpdateChatbotScripts([]int64{cs.Id.Get()}, cs)
}

// UpdateChatbotScripts updates existing chatbot.script records.
// All records (represented by ids) will be updated by cs values.
func (c *Client) UpdateChatbotScripts(ids []int64, cs *ChatbotScript) error {
	return c.Update(ChatbotScriptModel, ids, cs)
}

// DeleteChatbotScript deletes an existing chatbot.script record.
func (c *Client) DeleteChatbotScript(id int64) error {
	return c.DeleteChatbotScripts([]int64{id})
}

// DeleteChatbotScripts deletes existing chatbot.script records.
func (c *Client) DeleteChatbotScripts(ids []int64) error {
	return c.Delete(ChatbotScriptModel, ids)
}

// GetChatbotScript gets chatbot.script existing record.
func (c *Client) GetChatbotScript(id int64) (*ChatbotScript, error) {
	css, err := c.GetChatbotScripts([]int64{id})
	if err != nil {
		return nil, err
	}
	if css != nil && len(*css) > 0 {
		return &((*css)[0]), nil
	}
	return nil, fmt.Errorf("id %v of chatbot.script not found", id)
}

// GetChatbotScripts gets chatbot.script existing records.
func (c *Client) GetChatbotScripts(ids []int64) (*ChatbotScripts, error) {
	css := &ChatbotScripts{}
	if err := c.Read(ChatbotScriptModel, ids, nil, css); err != nil {
		return nil, err
	}
	return css, nil
}

// FindChatbotScript finds chatbot.script record by querying it with criteria.
func (c *Client) FindChatbotScript(criteria *Criteria) (*ChatbotScript, error) {
	css := &ChatbotScripts{}
	if err := c.SearchRead(ChatbotScriptModel, criteria, NewOptions().Limit(1), css); err != nil {
		return nil, err
	}
	if css != nil && len(*css) > 0 {
		return &((*css)[0]), nil
	}
	return nil, fmt.Errorf("chatbot.script was not found with criteria %v", criteria)
}

// FindChatbotScripts finds chatbot.script records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScripts(criteria *Criteria, options *Options) (*ChatbotScripts, error) {
	css := &ChatbotScripts{}
	if err := c.SearchRead(ChatbotScriptModel, criteria, options, css); err != nil {
		return nil, err
	}
	return css, nil
}

// FindChatbotScriptIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScriptIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ChatbotScriptModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindChatbotScriptId finds record id by querying it with criteria.
func (c *Client) FindChatbotScriptId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChatbotScriptModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("chatbot.script was not found with criteria %v and options %v", criteria, options)
}
