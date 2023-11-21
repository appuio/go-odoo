package odoo

import (
	"fmt"
)

// PrtMailMessageDraft represents prt.mail.message.draft model.
type PrtMailMessageDraft struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AttachmentIds               *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorId                    *Many2One  `xmlrpc:"author_id,omptempty"`
	Body                        *String    `xmlrpc:"body,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	Label                       *String    `xmlrpc:"label,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent              *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Model                       *String    `xmlrpc:"model,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	ParentId                    *Int       `xmlrpc:"parent_id,omptempty"`
	PartnerIds                  *Relation  `xmlrpc:"partner_ids,omptempty"`
	RecordRef                   *String    `xmlrpc:"record_ref,omptempty"`
	RefPartnerIds               *Relation  `xmlrpc:"ref_partner_ids,omptempty"`
	ResId                       *Int       `xmlrpc:"res_id,omptempty"`
	SignatureLocation           *Selection `xmlrpc:"signature_location,omptempty"`
	Subject                     *String    `xmlrpc:"subject,omptempty"`
	SubjectDisplay              *String    `xmlrpc:"subject_display,omptempty"`
	SubtypeId                   *Many2One  `xmlrpc:"subtype_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WizardMode                  *String    `xmlrpc:"wizard_mode,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PrtMailMessageDrafts represents array of prt.mail.message.draft model.
type PrtMailMessageDrafts []PrtMailMessageDraft

// PrtMailMessageDraftModel is the odoo model name.
const PrtMailMessageDraftModel = "prt.mail.message.draft"

// Many2One convert PrtMailMessageDraft to *Many2One.
func (pmmd *PrtMailMessageDraft) Many2One() *Many2One {
	return NewMany2One(pmmd.Id.Get(), "")
}

// CreatePrtMailMessageDraft creates a new prt.mail.message.draft model and returns its id.
func (c *Client) CreatePrtMailMessageDraft(pmmd *PrtMailMessageDraft) (int64, error) {
	ids, err := c.CreatePrtMailMessageDrafts([]*PrtMailMessageDraft{pmmd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePrtMailMessageDraft creates a new prt.mail.message.draft model and returns its id.
func (c *Client) CreatePrtMailMessageDrafts(pmmds []*PrtMailMessageDraft) ([]int64, error) {
	var vv []interface{}
	for _, v := range pmmds {
		vv = append(vv, v)
	}
	return c.Create(PrtMailMessageDraftModel, vv)
}

// UpdatePrtMailMessageDraft updates an existing prt.mail.message.draft record.
func (c *Client) UpdatePrtMailMessageDraft(pmmd *PrtMailMessageDraft) error {
	return c.UpdatePrtMailMessageDrafts([]int64{pmmd.Id.Get()}, pmmd)
}

// UpdatePrtMailMessageDrafts updates existing prt.mail.message.draft records.
// All records (represented by ids) will be updated by pmmd values.
func (c *Client) UpdatePrtMailMessageDrafts(ids []int64, pmmd *PrtMailMessageDraft) error {
	return c.Update(PrtMailMessageDraftModel, ids, pmmd)
}

// DeletePrtMailMessageDraft deletes an existing prt.mail.message.draft record.
func (c *Client) DeletePrtMailMessageDraft(id int64) error {
	return c.DeletePrtMailMessageDrafts([]int64{id})
}

// DeletePrtMailMessageDrafts deletes existing prt.mail.message.draft records.
func (c *Client) DeletePrtMailMessageDrafts(ids []int64) error {
	return c.Delete(PrtMailMessageDraftModel, ids)
}

// GetPrtMailMessageDraft gets prt.mail.message.draft existing record.
func (c *Client) GetPrtMailMessageDraft(id int64) (*PrtMailMessageDraft, error) {
	pmmds, err := c.GetPrtMailMessageDrafts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pmmds != nil && len(*pmmds) > 0 {
		return &((*pmmds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of prt.mail.message.draft not found", id)
}

// GetPrtMailMessageDrafts gets prt.mail.message.draft existing records.
func (c *Client) GetPrtMailMessageDrafts(ids []int64) (*PrtMailMessageDrafts, error) {
	pmmds := &PrtMailMessageDrafts{}
	if err := c.Read(PrtMailMessageDraftModel, ids, nil, pmmds); err != nil {
		return nil, err
	}
	return pmmds, nil
}

// FindPrtMailMessageDraft finds prt.mail.message.draft record by querying it with criteria.
func (c *Client) FindPrtMailMessageDraft(criteria *Criteria) (*PrtMailMessageDraft, error) {
	pmmds := &PrtMailMessageDrafts{}
	if err := c.SearchRead(PrtMailMessageDraftModel, criteria, NewOptions().Limit(1), pmmds); err != nil {
		return nil, err
	}
	if pmmds != nil && len(*pmmds) > 0 {
		return &((*pmmds)[0]), nil
	}
	return nil, fmt.Errorf("prt.mail.message.draft was not found with criteria %v", criteria)
}

// FindPrtMailMessageDrafts finds prt.mail.message.draft records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMailMessageDrafts(criteria *Criteria, options *Options) (*PrtMailMessageDrafts, error) {
	pmmds := &PrtMailMessageDrafts{}
	if err := c.SearchRead(PrtMailMessageDraftModel, criteria, options, pmmds); err != nil {
		return nil, err
	}
	return pmmds, nil
}

// FindPrtMailMessageDraftIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrtMailMessageDraftIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PrtMailMessageDraftModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPrtMailMessageDraftId finds record id by querying it with criteria.
func (c *Client) FindPrtMailMessageDraftId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrtMailMessageDraftModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("prt.mail.message.draft was not found with criteria %v and options %v", criteria, options)
}
