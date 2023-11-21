package odoo

import (
	"fmt"
)

// CetmixConversation represents cetmix.conversation model.
type CetmixConversation struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	AuthorId                 *Many2One `xmlrpc:"author_id,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	FailedMessageIds         *Relation `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	IsParticipant            *Bool     `xmlrpc:"is_participant,omptempty"`
	LastMessageBy            *Many2One `xmlrpc:"last_message_by,omptempty"`
	LastMessagePost          *Time     `xmlrpc:"last_message_post,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent           *String   `xmlrpc:"message_content,omptempty"`
	MessageCount             *Int      `xmlrpc:"message_count,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCount   *Int      `xmlrpc:"message_needaction_count,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	PartnerIds               *Relation `xmlrpc:"partner_ids,omptempty"`
	SubjectDisplay           *String   `xmlrpc:"subject_display,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CetmixConversations represents array of cetmix.conversation model.
type CetmixConversations []CetmixConversation

// CetmixConversationModel is the odoo model name.
const CetmixConversationModel = "cetmix.conversation"

// Many2One convert CetmixConversation to *Many2One.
func (cc *CetmixConversation) Many2One() *Many2One {
	return NewMany2One(cc.Id.Get(), "")
}

// CreateCetmixConversation creates a new cetmix.conversation model and returns its id.
func (c *Client) CreateCetmixConversation(cc *CetmixConversation) (int64, error) {
	ids, err := c.CreateCetmixConversations([]*CetmixConversation{cc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCetmixConversation creates a new cetmix.conversation model and returns its id.
func (c *Client) CreateCetmixConversations(ccs []*CetmixConversation) ([]int64, error) {
	var vv []interface{}
	for _, v := range ccs {
		vv = append(vv, v)
	}
	return c.Create(CetmixConversationModel, vv)
}

// UpdateCetmixConversation updates an existing cetmix.conversation record.
func (c *Client) UpdateCetmixConversation(cc *CetmixConversation) error {
	return c.UpdateCetmixConversations([]int64{cc.Id.Get()}, cc)
}

// UpdateCetmixConversations updates existing cetmix.conversation records.
// All records (represented by ids) will be updated by cc values.
func (c *Client) UpdateCetmixConversations(ids []int64, cc *CetmixConversation) error {
	return c.Update(CetmixConversationModel, ids, cc)
}

// DeleteCetmixConversation deletes an existing cetmix.conversation record.
func (c *Client) DeleteCetmixConversation(id int64) error {
	return c.DeleteCetmixConversations([]int64{id})
}

// DeleteCetmixConversations deletes existing cetmix.conversation records.
func (c *Client) DeleteCetmixConversations(ids []int64) error {
	return c.Delete(CetmixConversationModel, ids)
}

// GetCetmixConversation gets cetmix.conversation existing record.
func (c *Client) GetCetmixConversation(id int64) (*CetmixConversation, error) {
	ccs, err := c.GetCetmixConversations([]int64{id})
	if err != nil {
		return nil, err
	}
	if ccs != nil && len(*ccs) > 0 {
		return &((*ccs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of cetmix.conversation not found", id)
}

// GetCetmixConversations gets cetmix.conversation existing records.
func (c *Client) GetCetmixConversations(ids []int64) (*CetmixConversations, error) {
	ccs := &CetmixConversations{}
	if err := c.Read(CetmixConversationModel, ids, nil, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCetmixConversation finds cetmix.conversation record by querying it with criteria.
func (c *Client) FindCetmixConversation(criteria *Criteria) (*CetmixConversation, error) {
	ccs := &CetmixConversations{}
	if err := c.SearchRead(CetmixConversationModel, criteria, NewOptions().Limit(1), ccs); err != nil {
		return nil, err
	}
	if ccs != nil && len(*ccs) > 0 {
		return &((*ccs)[0]), nil
	}
	return nil, fmt.Errorf("cetmix.conversation was not found with criteria %v", criteria)
}

// FindCetmixConversations finds cetmix.conversation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCetmixConversations(criteria *Criteria, options *Options) (*CetmixConversations, error) {
	ccs := &CetmixConversations{}
	if err := c.SearchRead(CetmixConversationModel, criteria, options, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCetmixConversationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCetmixConversationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CetmixConversationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCetmixConversationId finds record id by querying it with criteria.
func (c *Client) FindCetmixConversationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CetmixConversationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("cetmix.conversation was not found with criteria %v and options %v", criteria, options)
}
