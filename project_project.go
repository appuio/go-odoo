package odoo

import (
	"fmt"
)

// ProjectProject represents project.project model.
type ProjectProject struct {
	LastUpdate                       *Time       `xmlrpc:"__last_update,omptempty"`
	AccessInstructionMessage         *String     `xmlrpc:"access_instruction_message,omptempty"`
	AccessToken                      *String     `xmlrpc:"access_token,omptempty"`
	AccessUrl                        *String     `xmlrpc:"access_url,omptempty"`
	AccessWarning                    *String     `xmlrpc:"access_warning,omptempty"`
	Active                           *Bool       `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId          *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline             *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration      *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon            *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                      *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState                    *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                  *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                 *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                   *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                   *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	AliasBouncedContent              *String     `xmlrpc:"alias_bounced_content,omptempty"`
	AliasContact                     *Selection  `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults                    *String     `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain                      *String     `xmlrpc:"alias_domain,omptempty"`
	AliasEnabled                     *Bool       `xmlrpc:"alias_enabled,omptempty"`
	AliasForceThreadId               *Int        `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                          *Many2One   `xmlrpc:"alias_id,omptempty"`
	AliasModelId                     *Many2One   `xmlrpc:"alias_model_id,omptempty"`
	AliasName                        *String     `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId               *Many2One   `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId              *Int        `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId                      *Many2One   `xmlrpc:"alias_user_id,omptempty"`
	AliasValue                       *String     `xmlrpc:"alias_value,omptempty"`
	AllocatedHours                   *Float      `xmlrpc:"allocated_hours,omptempty"`
	AllowBillable                    *Bool       `xmlrpc:"allow_billable,omptempty"`
	AllowForecast                    *Bool       `xmlrpc:"allow_forecast,omptempty"`
	AllowMilestones                  *Bool       `xmlrpc:"allow_milestones,omptempty"`
	AllowRating                      *Bool       `xmlrpc:"allow_rating,omptempty"`
	AllowRecurringTasks              *Bool       `xmlrpc:"allow_recurring_tasks,omptempty"`
	AllowSubtasks                    *Bool       `xmlrpc:"allow_subtasks,omptempty"`
	AllowTaskDependencies            *Bool       `xmlrpc:"allow_task_dependencies,omptempty"`
	AllowTimesheets                  *Bool       `xmlrpc:"allow_timesheets,omptempty"`
	AnalyticAccountBalance           *Float      `xmlrpc:"analytic_account_balance,omptempty"`
	AnalyticAccountId                *Many2One   `xmlrpc:"analytic_account_id,omptempty"`
	AssetsCount                      *Int        `xmlrpc:"assets_count,omptempty"`
	BillablePercentage               *Int        `xmlrpc:"billable_percentage,omptempty"`
	Budget                           *Int        `xmlrpc:"budget,omptempty"`
	CheckSpam                        *Bool       `xmlrpc:"check_spam,omptempty"`
	CollaboratorCount                *Int        `xmlrpc:"collaborator_count,omptempty"`
	CollaboratorIds                  *Relation   `xmlrpc:"collaborator_ids,omptempty"`
	Color                            *Int        `xmlrpc:"color,omptempty"`
	CommercialPartnerId              *Many2One   `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                        *Many2One   `xmlrpc:"company_id,omptempty"`
	ContractsCount                   *Int        `xmlrpc:"contracts_count,omptempty"`
	CreateDate                       *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                        *Many2One   `xmlrpc:"create_uid,omptempty"`
	CurrencyId                       *Many2One   `xmlrpc:"currency_id,omptempty"`
	Date                             *Time       `xmlrpc:"date,omptempty"`
	DateStart                        *Time       `xmlrpc:"date_start,omptempty"`
	Description                      *String     `xmlrpc:"description,omptempty"`
	DisplayCreateOrder               *Bool       `xmlrpc:"display_create_order,omptempty"`
	DisplayName                      *String     `xmlrpc:"display_name,omptempty"`
	DisplayPlanningTimesheetAnalysis *Bool       `xmlrpc:"display_planning_timesheet_analysis,omptempty"`
	DocCount                         *Int        `xmlrpc:"doc_count,omptempty"`
	DocumentCount                    *Int        `xmlrpc:"document_count,omptempty"`
	DocumentsFolderId                *Many2One   `xmlrpc:"documents_folder_id,omptempty"`
	DocumentsTagIds                  *Relation   `xmlrpc:"documents_tag_ids,omptempty"`
	DynamicLabel1Id                  *Many2One   `xmlrpc:"dynamic_label_1_id,omptempty"`
	DynamicLabel2Id                  *Many2One   `xmlrpc:"dynamic_label_2_id,omptempty"`
	DynamicTransportLabel1Id         *Many2One   `xmlrpc:"dynamic_transport_label_1_id,omptempty"`
	DynamicTransportLabel2Id         *Many2One   `xmlrpc:"dynamic_transport_label_2_id,omptempty"`
	EncodeUomInDays                  *Bool       `xmlrpc:"encode_uom_in_days,omptempty"`
	ExpensesCount                    *Int        `xmlrpc:"expenses_count,omptempty"`
	FailedMessageIds                 *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	FavoriteUserIds                  *Relation   `xmlrpc:"favorite_user_ids,omptempty"`
	HasAnySoToInvoice                *Bool       `xmlrpc:"has_any_so_to_invoice,omptempty"`
	HasAnySoWithNothingToInvoice     *Bool       `xmlrpc:"has_any_so_with_nothing_to_invoice,omptempty"`
	HasHelpdeskTeam                  *Bool       `xmlrpc:"has_helpdesk_team,omptempty"`
	HasMessage                       *Bool       `xmlrpc:"has_message,omptempty"`
	HelpdeskTeam                     *Relation   `xmlrpc:"helpdesk_team,omptempty"`
	Id                               *Int        `xmlrpc:"id,omptempty"`
	InvoiceCount                     *Int        `xmlrpc:"invoice_count,omptempty"`
	IsFavorite                       *Bool       `xmlrpc:"is_favorite,omptempty"`
	IsInternalProject                *Bool       `xmlrpc:"is_internal_project,omptempty"`
	IsMilestoneExceeded              *Bool       `xmlrpc:"is_milestone_exceeded,omptempty"`
	IsProjectOvertime                *Bool       `xmlrpc:"is_project_overtime,omptempty"`
	IsVisibleAcceptanceCriteria      *Bool       `xmlrpc:"is_visible_acceptance_criteria,omptempty"`
	IsVisibleBusinessValue           *Bool       `xmlrpc:"is_visible_business_value,omptempty"`
	IsVisibleDynamicFields           *Bool       `xmlrpc:"is_visible_dynamic_fields,omptempty"`
	IsVisibleEpic                    *Bool       `xmlrpc:"is_visible_epic,omptempty"`
	IsVisibleFieldOne                *Bool       `xmlrpc:"is_visible_field_one,omptempty"`
	IsVisibleReleasement             *Bool       `xmlrpc:"is_visible_releasement,omptempty"`
	IsVisibleRemarksDevelopment      *Bool       `xmlrpc:"is_visible_remarks_development,omptempty"`
	IsVisibleSecondOne               *Bool       `xmlrpc:"is_visible_second_one,omptempty"`
	IsVisibleSprint                  *Bool       `xmlrpc:"is_visible_sprint,omptempty"`
	IsVisibleStoryPoints             *Bool       `xmlrpc:"is_visible_story_points,omptempty"`
	IsVisibleTestCases               *Bool       `xmlrpc:"is_visible_test_cases,omptempty"`
	IsVisibleTopic                   *Bool       `xmlrpc:"is_visible_topic,omptempty"`
	IsVisibleWorkstream              *Bool       `xmlrpc:"is_visible_workstream,omptempty"`
	LabelTasks                       *String     `xmlrpc:"label_tasks,omptempty"`
	LastUpdateColor                  *Int        `xmlrpc:"last_update_color,omptempty"`
	LastUpdateId                     *Many2One   `xmlrpc:"last_update_id,omptempty"`
	LastUpdateStatus                 *Selection  `xmlrpc:"last_update_status,omptempty"`
	MessageAttachmentCount           *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                   *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds               *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                  *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter           *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError               *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                       *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId          *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter         *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MilestoneCount                   *Int        `xmlrpc:"milestone_count,omptempty"`
	MilestoneCountReached            *Int        `xmlrpc:"milestone_count_reached,omptempty"`
	MilestoneIds                     *Relation   `xmlrpc:"milestone_ids,omptempty"`
	MyActivityDateDeadline           *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                             *String     `xmlrpc:"name,omptempty"`
	PartnerEmail                     *String     `xmlrpc:"partner_email,omptempty"`
	PartnerId                        *Many2One   `xmlrpc:"partner_id,omptempty"`
	PartnerPhone                     *String     `xmlrpc:"partner_phone,omptempty"`
	PricingType                      *Selection  `xmlrpc:"pricing_type,omptempty"`
	PrivacyVisibility                *Selection  `xmlrpc:"privacy_visibility,omptempty"`
	PrivacyVisibilityWarning         *String     `xmlrpc:"privacy_visibility_warning,omptempty"`
	PurchaseOrdersCount              *Int        `xmlrpc:"purchase_orders_count,omptempty"`
	RatingActive                     *Bool       `xmlrpc:"rating_active,omptempty"`
	RatingAvg                        *Float      `xmlrpc:"rating_avg,omptempty"`
	RatingAvgPercentage              *Float      `xmlrpc:"rating_avg_percentage,omptempty"`
	RatingCount                      *Int        `xmlrpc:"rating_count,omptempty"`
	RatingIds                        *Relation   `xmlrpc:"rating_ids,omptempty"`
	RatingPercentageSatisfaction     *Int        `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RatingRequestDeadline            *Time       `xmlrpc:"rating_request_deadline,omptempty"`
	RatingStatus                     *Selection  `xmlrpc:"rating_status,omptempty"`
	RatingStatusPeriod               *Selection  `xmlrpc:"rating_status_period,omptempty"`
	RemainingHours                   *Float      `xmlrpc:"remaining_hours,omptempty"`
	ResourceCalendarId               *Many2One   `xmlrpc:"resource_calendar_id,omptempty"`
	SaleLineEmployeeIds              *Relation   `xmlrpc:"sale_line_employee_ids,omptempty"`
	SaleLineId                       *Many2One   `xmlrpc:"sale_line_id,omptempty"`
	SaleOrderCount                   *Int        `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderId                      *Many2One   `xmlrpc:"sale_order_id,omptempty"`
	Sequence                         *Int        `xmlrpc:"sequence,omptempty"`
	SharedDocumentCount              *Int        `xmlrpc:"shared_document_count,omptempty"`
	SharedDocumentIds                *Relation   `xmlrpc:"shared_document_ids,omptempty"`
	StageId                          *Many2One   `xmlrpc:"stage_id,omptempty"`
	SubscriptionsCount               *Int        `xmlrpc:"subscriptions_count,omptempty"`
	TagIds                           *Relation   `xmlrpc:"tag_ids,omptempty"`
	TaskCount                        *Int        `xmlrpc:"task_count,omptempty"`
	TaskCountWithSubtasks            *Int        `xmlrpc:"task_count_with_subtasks,omptempty"`
	TaskIds                          *Relation   `xmlrpc:"task_ids,omptempty"`
	TaskPropertiesDefinition         interface{} `xmlrpc:"task_properties_definition,omptempty"`
	Tasks                            *Relation   `xmlrpc:"tasks,omptempty"`
	TicketCount                      *Int        `xmlrpc:"ticket_count,omptempty"`
	TicketIds                        *Relation   `xmlrpc:"ticket_ids,omptempty"`
	TimesheetCount                   *Int        `xmlrpc:"timesheet_count,omptempty"`
	TimesheetEncodeUomId             *Many2One   `xmlrpc:"timesheet_encode_uom_id,omptempty"`
	TimesheetIds                     *Relation   `xmlrpc:"timesheet_ids,omptempty"`
	TimesheetProductId               *Many2One   `xmlrpc:"timesheet_product_id,omptempty"`
	TotalForecastTime                *Int        `xmlrpc:"total_forecast_time,omptempty"`
	TotalPlannedAmount               *Float      `xmlrpc:"total_planned_amount,omptempty"`
	TotalTimesheetTime               *Int        `xmlrpc:"total_timesheet_time,omptempty"`
	TransportIds                     *Relation   `xmlrpc:"transport_ids,omptempty"`
	TypeId                           *Many2One   `xmlrpc:"type_id,omptempty"`
	TypeIds                          *Relation   `xmlrpc:"type_ids,omptempty"`
	UpdateIds                        *Relation   `xmlrpc:"update_ids,omptempty"`
	UseDocuments                     *Bool       `xmlrpc:"use_documents,omptempty"`
	UserId                           *Many2One   `xmlrpc:"user_id,omptempty"`
	VendorBillCount                  *Int        `xmlrpc:"vendor_bill_count,omptempty"`
	WarningEmployeeRate              *Bool       `xmlrpc:"warning_employee_rate,omptempty"`
	WebsiteMessageIds                *Relation   `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                        *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                         *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// ProjectProjects represents array of project.project model.
type ProjectProjects []ProjectProject

// ProjectProjectModel is the odoo model name.
const ProjectProjectModel = "project.project"

// Many2One convert ProjectProject to *Many2One.
func (pp *ProjectProject) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProjectProject creates a new project.project model and returns its id.
func (c *Client) CreateProjectProject(pp *ProjectProject) (int64, error) {
	ids, err := c.CreateProjectProjects([]*ProjectProject{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectProject creates a new project.project model and returns its id.
func (c *Client) CreateProjectProjects(pps []*ProjectProject) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProjectProjectModel, vv)
}

// UpdateProjectProject updates an existing project.project record.
func (c *Client) UpdateProjectProject(pp *ProjectProject) error {
	return c.UpdateProjectProjects([]int64{pp.Id.Get()}, pp)
}

// UpdateProjectProjects updates existing project.project records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProjectProjects(ids []int64, pp *ProjectProject) error {
	return c.Update(ProjectProjectModel, ids, pp)
}

// DeleteProjectProject deletes an existing project.project record.
func (c *Client) DeleteProjectProject(id int64) error {
	return c.DeleteProjectProjects([]int64{id})
}

// DeleteProjectProjects deletes existing project.project records.
func (c *Client) DeleteProjectProjects(ids []int64) error {
	return c.Delete(ProjectProjectModel, ids)
}

// GetProjectProject gets project.project existing record.
func (c *Client) GetProjectProject(id int64) (*ProjectProject, error) {
	pps, err := c.GetProjectProjects([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.project not found", id)
}

// GetProjectProjects gets project.project existing records.
func (c *Client) GetProjectProjects(ids []int64) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.Read(ProjectProjectModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProject finds project.project record by querying it with criteria.
func (c *Client) FindProjectProject(criteria *Criteria) (*ProjectProject, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("project.project was not found with criteria %v", criteria)
}

// FindProjectProjects finds project.project records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjects(criteria *Criteria, options *Options) (*ProjectProjects, error) {
	pps := &ProjectProjects{}
	if err := c.SearchRead(ProjectProjectModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProjectProjectIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectProjectIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectProjectId finds record id by querying it with criteria.
func (c *Client) FindProjectProjectId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectProjectModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.project was not found with criteria %v and options %v", criteria, options)
}
