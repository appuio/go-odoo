package odoo

import (
	"fmt"
)

// HelpdeskTeam represents helpdesk.team model.
type HelpdeskTeam struct {
	LastUpdate                    *Time       `xmlrpc:"__last_update,omptempty"`
	AccessInstructionMessage      *String     `xmlrpc:"access_instruction_message,omptempty"`
	Active                        *Bool       `xmlrpc:"active,omptempty"`
	AliasBouncedContent           *String     `xmlrpc:"alias_bounced_content,omptempty"`
	AliasContact                  *Selection  `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults                 *String     `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain                   *String     `xmlrpc:"alias_domain,omptempty"`
	AliasEmailFrom                *String     `xmlrpc:"alias_email_from,omptempty"`
	AliasForceThreadId            *Int        `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                       *Many2One   `xmlrpc:"alias_id,omptempty"`
	AliasModelId                  *Many2One   `xmlrpc:"alias_model_id,omptempty"`
	AliasName                     *String     `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId            *Many2One   `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId           *Int        `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId                   *Many2One   `xmlrpc:"alias_user_id,omptempty"`
	AllowPortalTicketClosing      *Bool       `xmlrpc:"allow_portal_ticket_closing,omptempty"`
	AssignMethod                  *Selection  `xmlrpc:"assign_method,omptempty"`
	AutoAssignment                *Bool       `xmlrpc:"auto_assignment,omptempty"`
	AutoCloseDay                  *Int        `xmlrpc:"auto_close_day,omptempty"`
	AutoCloseTicket               *Bool       `xmlrpc:"auto_close_ticket,omptempty"`
	CanPublish                    *Bool       `xmlrpc:"can_publish,omptempty"`
	CheckSpam                     *Bool       `xmlrpc:"check_spam,omptempty"`
	Color                         *Int        `xmlrpc:"color,omptempty"`
	CompanyId                     *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate                    *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One   `xmlrpc:"create_uid,omptempty"`
	Description                   *String     `xmlrpc:"description,omptempty"`
	DisplayAliasName              *String     `xmlrpc:"display_alias_name,omptempty"`
	DisplayName                   *String     `xmlrpc:"display_name,omptempty"`
	FailedMessageIds              *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	FeatureFormUrl                *String     `xmlrpc:"feature_form_url,omptempty"`
	FeatureLivechatChannelId      *Many2One   `xmlrpc:"feature_livechat_channel_id,omptempty"`
	FromStageIds                  *Relation   `xmlrpc:"from_stage_ids,omptempty"`
	HasExternalMailServer         *Bool       `xmlrpc:"has_external_mail_server,omptempty"`
	HasMessage                    *Bool       `xmlrpc:"has_message,omptempty"`
	Id                            *Int        `xmlrpc:"id,omptempty"`
	IsPublished                   *Bool       `xmlrpc:"is_published,omptempty"`
	MemberIds                     *Relation   `xmlrpc:"member_ids,omptempty"`
	MessageAttachmentCount        *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds            *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError               *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter        *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError            *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                    *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower             *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId       *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction             *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter      *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds             *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	Name                          *String     `xmlrpc:"name,omptempty"`
	OpenTicketCount               *Int        `xmlrpc:"open_ticket_count,omptempty"`
	PortalShowRating              *Bool       `xmlrpc:"portal_show_rating,omptempty"`
	PrivacyVisibility             *Selection  `xmlrpc:"privacy_visibility,omptempty"`
	PrivacyVisibilityWarning      *String     `xmlrpc:"privacy_visibility_warning,omptempty"`
	ProjectId                     *Many2One   `xmlrpc:"project_id,omptempty"`
	RatingAvg                     *Float      `xmlrpc:"rating_avg,omptempty"`
	RatingAvgPercentage           *Float      `xmlrpc:"rating_avg_percentage,omptempty"`
	RatingCount                   *Int        `xmlrpc:"rating_count,omptempty"`
	RatingIds                     *Relation   `xmlrpc:"rating_ids,omptempty"`
	RatingPercentageSatisfaction  *Int        `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	ResourceCalendarId            *Many2One   `xmlrpc:"resource_calendar_id,omptempty"`
	Sequence                      *Int        `xmlrpc:"sequence,omptempty"`
	ShowKnowledgeBase             *Bool       `xmlrpc:"show_knowledge_base,omptempty"`
	ShowKnowledgeBaseArticle      *Bool       `xmlrpc:"show_knowledge_base_article,omptempty"`
	ShowKnowledgeBaseSlideChannel *Bool       `xmlrpc:"show_knowledge_base_slide_channel,omptempty"`
	SlaFailed                     *Int        `xmlrpc:"sla_failed,omptempty"`
	SlaPolicyCount                *Int        `xmlrpc:"sla_policy_count,omptempty"`
	StageIds                      *Relation   `xmlrpc:"stage_ids,omptempty"`
	SuccessRate                   *Float      `xmlrpc:"success_rate,omptempty"`
	TicketClosed                  *Int        `xmlrpc:"ticket_closed,omptempty"`
	TicketIds                     *Relation   `xmlrpc:"ticket_ids,omptempty"`
	TicketProperties              interface{} `xmlrpc:"ticket_properties,omptempty"`
	TimesheetEncodeUomId          *Many2One   `xmlrpc:"timesheet_encode_uom_id,omptempty"`
	ToStageId                     *Many2One   `xmlrpc:"to_stage_id,omptempty"`
	TotalTimesheetTime            *Int        `xmlrpc:"total_timesheet_time,omptempty"`
	UnassignedTickets             *Int        `xmlrpc:"unassigned_tickets,omptempty"`
	UrgentTicket                  *Int        `xmlrpc:"urgent_ticket,omptempty"`
	UseAlias                      *Bool       `xmlrpc:"use_alias,omptempty"`
	UseCoupons                    *Bool       `xmlrpc:"use_coupons,omptempty"`
	UseCreditNotes                *Bool       `xmlrpc:"use_credit_notes,omptempty"`
	UseFsm                        *Bool       `xmlrpc:"use_fsm,omptempty"`
	UseHelpdeskSaleTimesheet      *Bool       `xmlrpc:"use_helpdesk_sale_timesheet,omptempty"`
	UseHelpdeskTimesheet          *Bool       `xmlrpc:"use_helpdesk_timesheet,omptempty"`
	UseProductRepairs             *Bool       `xmlrpc:"use_product_repairs,omptempty"`
	UseProductReturns             *Bool       `xmlrpc:"use_product_returns,omptempty"`
	UseRating                     *Bool       `xmlrpc:"use_rating,omptempty"`
	UseSla                        *Bool       `xmlrpc:"use_sla,omptempty"`
	UseTwitter                    *Bool       `xmlrpc:"use_twitter,omptempty"`
	UseWebsiteHelpdeskForm        *Bool       `xmlrpc:"use_website_helpdesk_form,omptempty"`
	UseWebsiteHelpdeskForum       *Bool       `xmlrpc:"use_website_helpdesk_forum,omptempty"`
	UseWebsiteHelpdeskKnowledge   *Bool       `xmlrpc:"use_website_helpdesk_knowledge,omptempty"`
	UseWebsiteHelpdeskLivechat    *Bool       `xmlrpc:"use_website_helpdesk_livechat,omptempty"`
	UseWebsiteHelpdeskSlides      *Bool       `xmlrpc:"use_website_helpdesk_slides,omptempty"`
	WebsiteArticleId              *Many2One   `xmlrpc:"website_article_id,omptempty"`
	WebsiteFormViewId             *Many2One   `xmlrpc:"website_form_view_id,omptempty"`
	WebsiteId                     *Many2One   `xmlrpc:"website_id,omptempty"`
	WebsiteMenuId                 *Many2One   `xmlrpc:"website_menu_id,omptempty"`
	WebsiteMessageIds             *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WebsitePublished              *Bool       `xmlrpc:"website_published,omptempty"`
	WebsiteSlideChannelIds        *Relation   `xmlrpc:"website_slide_channel_ids,omptempty"`
	WebsiteUrl                    *String     `xmlrpc:"website_url,omptempty"`
	WriteDate                     *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskTeams represents array of helpdesk.team model.
type HelpdeskTeams []HelpdeskTeam

// HelpdeskTeamModel is the odoo model name.
const HelpdeskTeamModel = "helpdesk.team"

// Many2One convert HelpdeskTeam to *Many2One.
func (ht *HelpdeskTeam) Many2One() *Many2One {
	return NewMany2One(ht.Id.Get(), "")
}

// CreateHelpdeskTeam creates a new helpdesk.team model and returns its id.
func (c *Client) CreateHelpdeskTeam(ht *HelpdeskTeam) (int64, error) {
	ids, err := c.CreateHelpdeskTeams([]*HelpdeskTeam{ht})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskTeam creates a new helpdesk.team model and returns its id.
func (c *Client) CreateHelpdeskTeams(hts []*HelpdeskTeam) ([]int64, error) {
	var vv []interface{}
	for _, v := range hts {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskTeamModel, vv)
}

// UpdateHelpdeskTeam updates an existing helpdesk.team record.
func (c *Client) UpdateHelpdeskTeam(ht *HelpdeskTeam) error {
	return c.UpdateHelpdeskTeams([]int64{ht.Id.Get()}, ht)
}

// UpdateHelpdeskTeams updates existing helpdesk.team records.
// All records (represented by ids) will be updated by ht values.
func (c *Client) UpdateHelpdeskTeams(ids []int64, ht *HelpdeskTeam) error {
	return c.Update(HelpdeskTeamModel, ids, ht)
}

// DeleteHelpdeskTeam deletes an existing helpdesk.team record.
func (c *Client) DeleteHelpdeskTeam(id int64) error {
	return c.DeleteHelpdeskTeams([]int64{id})
}

// DeleteHelpdeskTeams deletes existing helpdesk.team records.
func (c *Client) DeleteHelpdeskTeams(ids []int64) error {
	return c.Delete(HelpdeskTeamModel, ids)
}

// GetHelpdeskTeam gets helpdesk.team existing record.
func (c *Client) GetHelpdeskTeam(id int64) (*HelpdeskTeam, error) {
	hts, err := c.GetHelpdeskTeams([]int64{id})
	if err != nil {
		return nil, err
	}
	if hts != nil && len(*hts) > 0 {
		return &((*hts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.team not found", id)
}

// GetHelpdeskTeams gets helpdesk.team existing records.
func (c *Client) GetHelpdeskTeams(ids []int64) (*HelpdeskTeams, error) {
	hts := &HelpdeskTeams{}
	if err := c.Read(HelpdeskTeamModel, ids, nil, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTeam finds helpdesk.team record by querying it with criteria.
func (c *Client) FindHelpdeskTeam(criteria *Criteria) (*HelpdeskTeam, error) {
	hts := &HelpdeskTeams{}
	if err := c.SearchRead(HelpdeskTeamModel, criteria, NewOptions().Limit(1), hts); err != nil {
		return nil, err
	}
	if hts != nil && len(*hts) > 0 {
		return &((*hts)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.team was not found with criteria %v", criteria)
}

// FindHelpdeskTeams finds helpdesk.team records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTeams(criteria *Criteria, options *Options) (*HelpdeskTeams, error) {
	hts := &HelpdeskTeams{}
	if err := c.SearchRead(HelpdeskTeamModel, criteria, options, hts); err != nil {
		return nil, err
	}
	return hts, nil
}

// FindHelpdeskTeamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskTeamIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskTeamModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskTeamId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskTeamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskTeamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.team was not found with criteria %v and options %v", criteria, options)
}
