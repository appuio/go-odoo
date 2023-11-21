package odoo

import (
	"fmt"
)

// CrmTeamMember represents crm.team.member model.
type CrmTeamMember struct {
	LastUpdate               *Time     `xmlrpc:"__last_update,omptempty"`
	Active                   *Bool     `xmlrpc:"active,omptempty"`
	AssignmentDomain         *String   `xmlrpc:"assignment_domain,omptempty"`
	AssignmentEnabled        *Bool     `xmlrpc:"assignment_enabled,omptempty"`
	AssignmentMax            *Int      `xmlrpc:"assignment_max,omptempty"`
	AssignmentOptout         *Bool     `xmlrpc:"assignment_optout,omptempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omptempty"`
	CrmTeamId                *Many2One `xmlrpc:"crm_team_id,omptempty"`
	DisplayName              *String   `xmlrpc:"display_name,omptempty"`
	Email                    *String   `xmlrpc:"email,omptempty"`
	FailedMessageIds         *Relation `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omptempty"`
	Id                       *Int      `xmlrpc:"id,omptempty"`
	Image128                 *String   `xmlrpc:"image_128,omptempty"`
	Image1920                *String   `xmlrpc:"image_1920,omptempty"`
	IsMembershipMulti        *Bool     `xmlrpc:"is_membership_multi,omptempty"`
	LeadMonthCount           *Int      `xmlrpc:"lead_month_count,omptempty"`
	MemberWarning            *String   `xmlrpc:"member_warning,omptempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent           *String   `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omptempty"`
	Mobile                   *String   `xmlrpc:"mobile,omptempty"`
	Name                     *String   `xmlrpc:"name,omptempty"`
	Phone                    *String   `xmlrpc:"phone,omptempty"`
	UserCompanyIds           *Relation `xmlrpc:"user_company_ids,omptempty"`
	UserId                   *Many2One `xmlrpc:"user_id,omptempty"`
	UserInTeamsIds           *Relation `xmlrpc:"user_in_teams_ids,omptempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omptempty"`
}

// CrmTeamMembers represents array of crm.team.member model.
type CrmTeamMembers []CrmTeamMember

// CrmTeamMemberModel is the odoo model name.
const CrmTeamMemberModel = "crm.team.member"

// Many2One convert CrmTeamMember to *Many2One.
func (ctm *CrmTeamMember) Many2One() *Many2One {
	return NewMany2One(ctm.Id.Get(), "")
}

// CreateCrmTeamMember creates a new crm.team.member model and returns its id.
func (c *Client) CreateCrmTeamMember(ctm *CrmTeamMember) (int64, error) {
	ids, err := c.CreateCrmTeamMembers([]*CrmTeamMember{ctm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmTeamMember creates a new crm.team.member model and returns its id.
func (c *Client) CreateCrmTeamMembers(ctms []*CrmTeamMember) ([]int64, error) {
	var vv []interface{}
	for _, v := range ctms {
		vv = append(vv, v)
	}
	return c.Create(CrmTeamMemberModel, vv)
}

// UpdateCrmTeamMember updates an existing crm.team.member record.
func (c *Client) UpdateCrmTeamMember(ctm *CrmTeamMember) error {
	return c.UpdateCrmTeamMembers([]int64{ctm.Id.Get()}, ctm)
}

// UpdateCrmTeamMembers updates existing crm.team.member records.
// All records (represented by ids) will be updated by ctm values.
func (c *Client) UpdateCrmTeamMembers(ids []int64, ctm *CrmTeamMember) error {
	return c.Update(CrmTeamMemberModel, ids, ctm)
}

// DeleteCrmTeamMember deletes an existing crm.team.member record.
func (c *Client) DeleteCrmTeamMember(id int64) error {
	return c.DeleteCrmTeamMembers([]int64{id})
}

// DeleteCrmTeamMembers deletes existing crm.team.member records.
func (c *Client) DeleteCrmTeamMembers(ids []int64) error {
	return c.Delete(CrmTeamMemberModel, ids)
}

// GetCrmTeamMember gets crm.team.member existing record.
func (c *Client) GetCrmTeamMember(id int64) (*CrmTeamMember, error) {
	ctms, err := c.GetCrmTeamMembers([]int64{id})
	if err != nil {
		return nil, err
	}
	if ctms != nil && len(*ctms) > 0 {
		return &((*ctms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.team.member not found", id)
}

// GetCrmTeamMembers gets crm.team.member existing records.
func (c *Client) GetCrmTeamMembers(ids []int64) (*CrmTeamMembers, error) {
	ctms := &CrmTeamMembers{}
	if err := c.Read(CrmTeamMemberModel, ids, nil, ctms); err != nil {
		return nil, err
	}
	return ctms, nil
}

// FindCrmTeamMember finds crm.team.member record by querying it with criteria.
func (c *Client) FindCrmTeamMember(criteria *Criteria) (*CrmTeamMember, error) {
	ctms := &CrmTeamMembers{}
	if err := c.SearchRead(CrmTeamMemberModel, criteria, NewOptions().Limit(1), ctms); err != nil {
		return nil, err
	}
	if ctms != nil && len(*ctms) > 0 {
		return &((*ctms)[0]), nil
	}
	return nil, fmt.Errorf("crm.team.member was not found with criteria %v", criteria)
}

// FindCrmTeamMembers finds crm.team.member records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTeamMembers(criteria *Criteria, options *Options) (*CrmTeamMembers, error) {
	ctms := &CrmTeamMembers{}
	if err := c.SearchRead(CrmTeamMemberModel, criteria, options, ctms); err != nil {
		return nil, err
	}
	return ctms, nil
}

// FindCrmTeamMemberIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTeamMemberIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(CrmTeamMemberModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmTeamMemberId finds record id by querying it with criteria.
func (c *Client) FindCrmTeamMemberId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmTeamMemberModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.team.member was not found with criteria %v and options %v", criteria, options)
}
