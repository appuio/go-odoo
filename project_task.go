package odoo

import (
	"fmt"
)

// ProjectTask represents project.task model.
type ProjectTask struct {
	LastUpdate                         *Time       `xmlrpc:"__last_update,omptempty"`
	AcceptanceCriteria                 *String     `xmlrpc:"acceptance_criteria,omptempty"`
	AccessToken                        *String     `xmlrpc:"access_token,omptempty"`
	AccessUrl                          *String     `xmlrpc:"access_url,omptempty"`
	AccessWarning                      *String     `xmlrpc:"access_warning,omptempty"`
	Active                             *Bool       `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId            *Many2One   `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline               *Time       `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration        *Selection  `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon              *String     `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                        *Relation   `xmlrpc:"activity_ids,omptempty"`
	ActivityState                      *Selection  `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                    *String     `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                   *String     `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                     *Many2One   `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                     *Many2One   `xmlrpc:"activity_user_id,omptempty"`
	AddUrl                             *String     `xmlrpc:"add_url,omptempty"`
	AllocatedHours                     *Float      `xmlrpc:"allocated_hours,omptempty"`
	AllocationType                     *Selection  `xmlrpc:"allocation_type,omptempty"`
	AllowBillable                      *Bool       `xmlrpc:"allow_billable,omptempty"`
	AllowMilestones                    *Bool       `xmlrpc:"allow_milestones,omptempty"`
	AllowRecurringTasks                *Bool       `xmlrpc:"allow_recurring_tasks,omptempty"`
	AllowSubtasks                      *Bool       `xmlrpc:"allow_subtasks,omptempty"`
	AllowTaskDependencies              *Bool       `xmlrpc:"allow_task_dependencies,omptempty"`
	AllowTimesheets                    *Bool       `xmlrpc:"allow_timesheets,omptempty"`
	AnalyticAccountActive              *Bool       `xmlrpc:"analytic_account_active,omptempty"`
	AnalyticAccountId                  *Many2One   `xmlrpc:"analytic_account_id,omptempty"`
	AncestorId                         *Many2One   `xmlrpc:"ancestor_id,omptempty"`
	AttachmentIds                      *Relation   `xmlrpc:"attachment_ids,omptempty"`
	AvailableProjectScrumEpicIds       *Relation   `xmlrpc:"available_project_scrum_epic_ids,omptempty"`
	AvailableProjectScrumSprintIds     *Relation   `xmlrpc:"available_project_scrum_sprint_ids,omptempty"`
	AvailableProjectScrumTopicIds      *Relation   `xmlrpc:"available_project_scrum_topic_ids,omptempty"`
	AvailableProjectScrumWorkstreamIds *Relation   `xmlrpc:"available_project_scrum_workstream_ids,omptempty"`
	AvailableTagsIds                   *Relation   `xmlrpc:"available_tags_ids,omptempty"`
	BusinessValue                      *Int        `xmlrpc:"business_value,omptempty"`
	BusinessValueSelect                *Selection  `xmlrpc:"business_value_select,omptempty"`
	ChildIds                           *Relation   `xmlrpc:"child_ids,omptempty"`
	ChildText                          *String     `xmlrpc:"child_text,omptempty"`
	Color                              *Int        `xmlrpc:"color,omptempty"`
	CommercialPartnerId                *Many2One   `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyId                          *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate                         *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid                          *Many2One   `xmlrpc:"create_uid,omptempty"`
	DataFetcherIds                     *Relation   `xmlrpc:"data_fetcher_ids,omptempty"`
	DateAssign                         *Time       `xmlrpc:"date_assign,omptempty"`
	DateDeadline                       *Time       `xmlrpc:"date_deadline,omptempty"`
	DateEnd                            *Time       `xmlrpc:"date_end,omptempty"`
	DateLastStageUpdate                *Time       `xmlrpc:"date_last_stage_update,omptempty"`
	DependOnIds                        *Relation   `xmlrpc:"depend_on_ids,omptempty"`
	DependentIds                       *Relation   `xmlrpc:"dependent_ids,omptempty"`
	DependentTasksCount                *Int        `xmlrpc:"dependent_tasks_count,omptempty"`
	Description                        *String     `xmlrpc:"description,omptempty"`
	DisplayName                        *String     `xmlrpc:"display_name,omptempty"`
	DisplayParentTaskButton            *Bool       `xmlrpc:"display_parent_task_button,omptempty"`
	DisplayProjectId                   *Many2One   `xmlrpc:"display_project_id,omptempty"`
	DisplaySaleOrderButton             *Bool       `xmlrpc:"display_sale_order_button,omptempty"`
	DisplayTimerPause                  *Bool       `xmlrpc:"display_timer_pause,omptempty"`
	DisplayTimerResume                 *Bool       `xmlrpc:"display_timer_resume,omptempty"`
	DisplayTimerStartPrimary           *Bool       `xmlrpc:"display_timer_start_primary,omptempty"`
	DisplayTimerStartSecondary         *Bool       `xmlrpc:"display_timer_start_secondary,omptempty"`
	DisplayTimerStop                   *Bool       `xmlrpc:"display_timer_stop,omptempty"`
	DisplayTimesheetTimer              *Bool       `xmlrpc:"display_timesheet_timer,omptempty"`
	DisplayWarningDependencyInGantt    *Bool       `xmlrpc:"display_warning_dependency_in_gantt,omptempty"`
	DisplayedImageId                   *Many2One   `xmlrpc:"displayed_image_id,omptempty"`
	DocumentCount                      *Int        `xmlrpc:"document_count,omptempty"`
	DocumentIds                        *Relation   `xmlrpc:"document_ids,omptempty"`
	DocumentsFolderId                  *Many2One   `xmlrpc:"documents_folder_id,omptempty"`
	Duration                           *Float      `xmlrpc:"duration,omptempty"`
	DynamicLabel1Id                    *Many2One   `xmlrpc:"dynamic_label_1_id,omptempty"`
	DynamicLabel2Id                    *Many2One   `xmlrpc:"dynamic_label_2_id,omptempty"`
	DynamicValue1Id                    *Many2One   `xmlrpc:"dynamic_value_1_id,omptempty"`
	DynamicValue2Id                    *Many2One   `xmlrpc:"dynamic_value_2_id,omptempty"`
	EffectiveHours                     *Float      `xmlrpc:"effective_hours,omptempty"`
	EmailCc                            *String     `xmlrpc:"email_cc,omptempty"`
	EmailFrom                          *String     `xmlrpc:"email_from,omptempty"`
	EncodeUomInDays                    *Bool       `xmlrpc:"encode_uom_in_days,omptempty"`
	EpicId                             *Many2One   `xmlrpc:"epic_id,omptempty"`
	FailedMessageIds                   *Relation   `xmlrpc:"failed_message_ids,omptempty"`
	Fri                                *Bool       `xmlrpc:"fri,omptempty"`
	HasLateAndUnreachedMilestone       *Bool       `xmlrpc:"has_late_and_unreached_milestone,omptempty"`
	HasMessage                         *Bool       `xmlrpc:"has_message,omptempty"`
	HasMultiSol                        *Bool       `xmlrpc:"has_multi_sol,omptempty"`
	Id                                 *Int        `xmlrpc:"id,omptempty"`
	IsAbsent                           *Bool       `xmlrpc:"is_absent,omptempty"`
	IsAnalyticAccountIdChanged         *Bool       `xmlrpc:"is_analytic_account_id_changed,omptempty"`
	IsBlocked                          *Bool       `xmlrpc:"is_blocked,omptempty"`
	IsClosed                           *Bool       `xmlrpc:"is_closed,omptempty"`
	IsPrivate                          *Bool       `xmlrpc:"is_private,omptempty"`
	IsProjectMapEmpty                  *Bool       `xmlrpc:"is_project_map_empty,omptempty"`
	IsTimeoffTask                      *Bool       `xmlrpc:"is_timeoff_task,omptempty"`
	IsTimerRunning                     *Bool       `xmlrpc:"is_timer_running,omptempty"`
	IsVisibleAcceptanceCriteria        *Bool       `xmlrpc:"is_visible_acceptance_criteria,omptempty"`
	IsVisibleBusinessValue             *Bool       `xmlrpc:"is_visible_business_value,omptempty"`
	IsVisibleDynamicFields             *Bool       `xmlrpc:"is_visible_dynamic_fields,omptempty"`
	IsVisibleEpic                      *Bool       `xmlrpc:"is_visible_epic,omptempty"`
	IsVisibleFieldOne                  *Bool       `xmlrpc:"is_visible_field_one,omptempty"`
	IsVisibleReleasement               *Bool       `xmlrpc:"is_visible_releasement,omptempty"`
	IsVisibleRemarksDevelopment        *Bool       `xmlrpc:"is_visible_remarks_development,omptempty"`
	IsVisibleSecondOne                 *Bool       `xmlrpc:"is_visible_second_one,omptempty"`
	IsVisibleSprint                    *Bool       `xmlrpc:"is_visible_sprint,omptempty"`
	IsVisibleStoryPoints               *Bool       `xmlrpc:"is_visible_story_points,omptempty"`
	IsVisibleTestCases                 *Bool       `xmlrpc:"is_visible_test_cases,omptempty"`
	IsVisibleTopic                     *Bool       `xmlrpc:"is_visible_topic,omptempty"`
	IsVisibleWorkstream                *Bool       `xmlrpc:"is_visible_workstream,omptempty"`
	KanbanState                        *Selection  `xmlrpc:"kanban_state,omptempty"`
	KanbanStateLabel                   *String     `xmlrpc:"kanban_state_label,omptempty"`
	LeaveTypesCount                    *Int        `xmlrpc:"leave_types_count,omptempty"`
	LeaveWarning                       *String     `xmlrpc:"leave_warning,omptempty"`
	LegendBlocked                      *String     `xmlrpc:"legend_blocked,omptempty"`
	LegendDone                         *String     `xmlrpc:"legend_done,omptempty"`
	LegendNormal                       *String     `xmlrpc:"legend_normal,omptempty"`
	ManagerId                          *Many2One   `xmlrpc:"manager_id,omptempty"`
	MessageAttachmentCount             *Int        `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                     *String     `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                 *Relation   `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                    *Bool       `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter             *Int        `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                 *Bool       `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                         *Relation   `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                  *Bool       `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId            *Many2One   `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                  *Bool       `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter           *Int        `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                  *Relation   `xmlrpc:"message_partner_ids,omptempty"`
	MilestoneId                        *Many2One   `xmlrpc:"milestone_id,omptempty"`
	Mon                                *Bool       `xmlrpc:"mon,omptempty"`
	MyActivityDateDeadline             *Time       `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                               *String     `xmlrpc:"name,omptempty"`
	OverlapWarning                     *String     `xmlrpc:"overlap_warning,omptempty"`
	Overtime                           *Float      `xmlrpc:"overtime,omptempty"`
	OwnerId                            *Many2One   `xmlrpc:"owner_id,omptempty"`
	ParentId                           *Many2One   `xmlrpc:"parent_id,omptempty"`
	PartnerCity                        *String     `xmlrpc:"partner_city,omptempty"`
	PartnerEmail                       *String     `xmlrpc:"partner_email,omptempty"`
	PartnerId                          *Many2One   `xmlrpc:"partner_id,omptempty"`
	PartnerIsCompany                   *Bool       `xmlrpc:"partner_is_company,omptempty"`
	PartnerMobile                      *String     `xmlrpc:"partner_mobile,omptempty"`
	PartnerPhone                       *String     `xmlrpc:"partner_phone,omptempty"`
	PartnerStreet                      *String     `xmlrpc:"partner_street,omptempty"`
	PartnerZip                         *String     `xmlrpc:"partner_zip,omptempty"`
	PersonalStageId                    *Many2One   `xmlrpc:"personal_stage_id,omptempty"`
	PersonalStageTypeId                *Many2One   `xmlrpc:"personal_stage_type_id,omptempty"`
	PersonalStageTypeIds               *Relation   `xmlrpc:"personal_stage_type_ids,omptempty"`
	PlannedDateBegin                   *Time       `xmlrpc:"planned_date_begin,omptempty"`
	PlannedDateEnd                     *Time       `xmlrpc:"planned_date_end,omptempty"`
	PlannedHours                       *Float      `xmlrpc:"planned_hours,omptempty"`
	PlanningOverlap                    *Int        `xmlrpc:"planning_overlap,omptempty"`
	PortalEffectiveHours               *Float      `xmlrpc:"portal_effective_hours,omptempty"`
	PortalProgress                     *Float      `xmlrpc:"portal_progress,omptempty"`
	PortalRemainingHours               *Float      `xmlrpc:"portal_remaining_hours,omptempty"`
	PortalSubtaskEffectiveHours        *Float      `xmlrpc:"portal_subtask_effective_hours,omptempty"`
	PortalTotalHoursSpent              *Float      `xmlrpc:"portal_total_hours_spent,omptempty"`
	PortalUserNames                    *String     `xmlrpc:"portal_user_names,omptempty"`
	PricingType                        *Selection  `xmlrpc:"pricing_type,omptempty"`
	Priority                           *Selection  `xmlrpc:"priority,omptempty"`
	Progress                           *Float      `xmlrpc:"progress,omptempty"`
	ProjectAnalyticAccountId           *Many2One   `xmlrpc:"project_analytic_account_id,omptempty"`
	ProjectColor                       *Int        `xmlrpc:"project_color,omptempty"`
	ProjectId                          *Many2One   `xmlrpc:"project_id,omptempty"`
	ProjectPrivacyVisibility           *Selection  `xmlrpc:"project_privacy_visibility,omptempty"`
	ProjectSaleOrderId                 *Many2One   `xmlrpc:"project_sale_order_id,omptempty"`
	ProjectUseDocuments                *Bool       `xmlrpc:"project_use_documents,omptempty"`
	RatingActive                       *Bool       `xmlrpc:"rating_active,omptempty"`
	RatingAvg                          *Float      `xmlrpc:"rating_avg,omptempty"`
	RatingAvgText                      *Selection  `xmlrpc:"rating_avg_text,omptempty"`
	RatingCount                        *Int        `xmlrpc:"rating_count,omptempty"`
	RatingIds                          *Relation   `xmlrpc:"rating_ids,omptempty"`
	RatingLastFeedback                 *String     `xmlrpc:"rating_last_feedback,omptempty"`
	RatingLastImage                    *String     `xmlrpc:"rating_last_image,omptempty"`
	RatingLastText                     *Selection  `xmlrpc:"rating_last_text,omptempty"`
	RatingLastValue                    *Float      `xmlrpc:"rating_last_value,omptempty"`
	RatingPercentageSatisfaction       *Float      `xmlrpc:"rating_percentage_satisfaction,omptempty"`
	RecurrenceId                       *Many2One   `xmlrpc:"recurrence_id,omptempty"`
	RecurrenceMessage                  *String     `xmlrpc:"recurrence_message,omptempty"`
	RecurrenceUpdate                   *Selection  `xmlrpc:"recurrence_update,omptempty"`
	RecurringCount                     *Int        `xmlrpc:"recurring_count,omptempty"`
	RecurringTask                      *Bool       `xmlrpc:"recurring_task,omptempty"`
	ReleaseDate                        *Time       `xmlrpc:"release_date,omptempty"`
	ReleaseId                          *Many2One   `xmlrpc:"release_id,omptempty"`
	ReleaseIds                         *Relation   `xmlrpc:"release_ids,omptempty"`
	RemainingHours                     *Float      `xmlrpc:"remaining_hours,omptempty"`
	RemainingHoursAvailable            *Bool       `xmlrpc:"remaining_hours_available,omptempty"`
	RemainingHoursPercentage           *Float      `xmlrpc:"remaining_hours_percentage,omptempty"`
	RemainingHoursSo                   *Float      `xmlrpc:"remaining_hours_so,omptempty"`
	RemarksDevelopment                 *String     `xmlrpc:"remarks_development,omptempty"`
	RemoteError                        *Bool       `xmlrpc:"remote_error,omptempty"`
	RepeatDay                          *Selection  `xmlrpc:"repeat_day,omptempty"`
	RepeatInterval                     *Int        `xmlrpc:"repeat_interval,omptempty"`
	RepeatMonth                        *Selection  `xmlrpc:"repeat_month,omptempty"`
	RepeatNumber                       *Int        `xmlrpc:"repeat_number,omptempty"`
	RepeatOnMonth                      *Selection  `xmlrpc:"repeat_on_month,omptempty"`
	RepeatOnYear                       *Selection  `xmlrpc:"repeat_on_year,omptempty"`
	RepeatShowDay                      *Bool       `xmlrpc:"repeat_show_day,omptempty"`
	RepeatShowDow                      *Bool       `xmlrpc:"repeat_show_dow,omptempty"`
	RepeatShowMonth                    *Bool       `xmlrpc:"repeat_show_month,omptempty"`
	RepeatShowWeek                     *Bool       `xmlrpc:"repeat_show_week,omptempty"`
	RepeatType                         *Selection  `xmlrpc:"repeat_type,omptempty"`
	RepeatUnit                         *Selection  `xmlrpc:"repeat_unit,omptempty"`
	RepeatUntil                        *Time       `xmlrpc:"repeat_until,omptempty"`
	RepeatWeek                         *Selection  `xmlrpc:"repeat_week,omptempty"`
	RepeatWeekday                      *Selection  `xmlrpc:"repeat_weekday,omptempty"`
	SaleLineId                         *Many2One   `xmlrpc:"sale_line_id,omptempty"`
	SaleOrderId                        *Many2One   `xmlrpc:"sale_order_id,omptempty"`
	Sat                                *Bool       `xmlrpc:"sat,omptempty"`
	Sequence                           *Int        `xmlrpc:"sequence,omptempty"`
	SharedDocumentIds                  *Relation   `xmlrpc:"shared_document_ids,omptempty"`
	SoAnalyticAccountId                *Many2One   `xmlrpc:"so_analytic_account_id,omptempty"`
	SprintId                           *Many2One   `xmlrpc:"sprint_id,omptempty"`
	StageId                            *Many2One   `xmlrpc:"stage_id,omptempty"`
	StakeholderIds                     *Relation   `xmlrpc:"stakeholder_ids,omptempty"`
	StatusChangedDate                  *Time       `xmlrpc:"status_changed_date,omptempty"`
	Statuses                           *String     `xmlrpc:"statuses,omptempty"`
	StoryPoints                        *Int        `xmlrpc:"story_points,omptempty"`
	StoryPointsSelect                  *Selection  `xmlrpc:"story_points_select,omptempty"`
	SubtaskCount                       *Int        `xmlrpc:"subtask_count,omptempty"`
	SubtaskEffectiveHours              *Float      `xmlrpc:"subtask_effective_hours,omptempty"`
	SubtaskPlannedHours                *Float      `xmlrpc:"subtask_planned_hours,omptempty"`
	Sun                                *Bool       `xmlrpc:"sun,omptempty"`
	TagIds                             *Relation   `xmlrpc:"tag_ids,omptempty"`
	TaskProperties                     interface{} `xmlrpc:"task_properties,omptempty"`
	TaskToInvoice                      *Bool       `xmlrpc:"task_to_invoice,omptempty"`
	TestCases                          *String     `xmlrpc:"test_cases,omptempty"`
	Thu                                *Bool       `xmlrpc:"thu,omptempty"`
	TimerPause                         *Time       `xmlrpc:"timer_pause,omptempty"`
	TimerStart                         *Time       `xmlrpc:"timer_start,omptempty"`
	TimesheetIds                       *Relation   `xmlrpc:"timesheet_ids,omptempty"`
	TimesheetProductId                 *Many2One   `xmlrpc:"timesheet_product_id,omptempty"`
	TopicId                            *Many2One   `xmlrpc:"topic_id,omptempty"`
	TotalHoursSpent                    *Float      `xmlrpc:"total_hours_spent,omptempty"`
	TransportIds                       *Relation   `xmlrpc:"transport_ids,omptempty"`
	Tue                                *Bool       `xmlrpc:"tue,omptempty"`
	TypeId                             *Many2One   `xmlrpc:"type_id,omptempty"`
	UserIds                            *Relation   `xmlrpc:"user_ids,omptempty"`
	UserNames                          *String     `xmlrpc:"user_names,omptempty"`
	UserTimerId                        *Relation   `xmlrpc:"user_timer_id,omptempty"`
	WebsiteMessageIds                  *Relation   `xmlrpc:"website_message_ids,omptempty"`
	Wed                                *Bool       `xmlrpc:"wed,omptempty"`
	WorkingDaysClose                   *Float      `xmlrpc:"working_days_close,omptempty"`
	WorkingDaysOpen                    *Float      `xmlrpc:"working_days_open,omptempty"`
	WorkingHoursClose                  *Float      `xmlrpc:"working_hours_close,omptempty"`
	WorkingHoursOpen                   *Float      `xmlrpc:"working_hours_open,omptempty"`
	WorkstreamId                       *Many2One   `xmlrpc:"workstream_id,omptempty"`
	WriteDate                          *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid                           *Many2One   `xmlrpc:"write_uid,omptempty"`
	XTaskIds                           *Relation   `xmlrpc:"x_task_ids,omptempty"`
}

// ProjectTasks represents array of project.task model.
type ProjectTasks []ProjectTask

// ProjectTaskModel is the odoo model name.
const ProjectTaskModel = "project.task"

// Many2One convert ProjectTask to *Many2One.
func (pt *ProjectTask) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProjectTask creates a new project.task model and returns its id.
func (c *Client) CreateProjectTask(pt *ProjectTask) (int64, error) {
	ids, err := c.CreateProjectTasks([]*ProjectTask{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTask creates a new project.task model and returns its id.
func (c *Client) CreateProjectTasks(pts []*ProjectTask) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskModel, vv)
}

// UpdateProjectTask updates an existing project.task record.
func (c *Client) UpdateProjectTask(pt *ProjectTask) error {
	return c.UpdateProjectTasks([]int64{pt.Id.Get()}, pt)
}

// UpdateProjectTasks updates existing project.task records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProjectTasks(ids []int64, pt *ProjectTask) error {
	return c.Update(ProjectTaskModel, ids, pt)
}

// DeleteProjectTask deletes an existing project.task record.
func (c *Client) DeleteProjectTask(id int64) error {
	return c.DeleteProjectTasks([]int64{id})
}

// DeleteProjectTasks deletes existing project.task records.
func (c *Client) DeleteProjectTasks(ids []int64) error {
	return c.Delete(ProjectTaskModel, ids)
}

// GetProjectTask gets project.task existing record.
func (c *Client) GetProjectTask(id int64) (*ProjectTask, error) {
	pts, err := c.GetProjectTasks([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.task not found", id)
}

// GetProjectTasks gets project.task existing records.
func (c *Client) GetProjectTasks(ids []int64) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.Read(ProjectTaskModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTask finds project.task record by querying it with criteria.
func (c *Client) FindProjectTask(criteria *Criteria) (*ProjectTask, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("project.task was not found with criteria %v", criteria)
}

// FindProjectTasks finds project.task records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTasks(criteria *Criteria, options *Options) (*ProjectTasks, error) {
	pts := &ProjectTasks{}
	if err := c.SearchRead(ProjectTaskModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProjectTaskIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectTaskModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectTaskId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.task was not found with criteria %v and options %v", criteria, options)
}
