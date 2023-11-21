package odoo

import (
	"fmt"
)

// HrEmployee represents hr.employee model.
type HrEmployee struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityCalendarEventId                     *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline                        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration                 *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AdditionalNote                              *String    `xmlrpc:"additional_note,omptempty"`
	AddressHomeId                               *Many2One  `xmlrpc:"address_home_id,omptempty"`
	AddressId                                   *Many2One  `xmlrpc:"address_id,omptempty"`
	AllocationCount                             *Float     `xmlrpc:"allocation_count,omptempty"`
	AllocationDisplay                           *String    `xmlrpc:"allocation_display,omptempty"`
	AllocationRemainingDisplay                  *String    `xmlrpc:"allocation_remaining_display,omptempty"`
	AllocationsCount                            *Int       `xmlrpc:"allocations_count,omptempty"`
	ApplicantId                                 *Relation  `xmlrpc:"applicant_id,omptempty"`
	AppraisalCount                              *Int       `xmlrpc:"appraisal_count,omptempty"`
	AppraisalIds                                *Relation  `xmlrpc:"appraisal_ids,omptempty"`
	AttendanceIds                               *Relation  `xmlrpc:"attendance_ids,omptempty"`
	AttendanceState                             *Selection `xmlrpc:"attendance_state,omptempty"`
	Avatar1024                                  *String    `xmlrpc:"avatar_1024,omptempty"`
	Avatar128                                   *String    `xmlrpc:"avatar_128,omptempty"`
	Avatar1920                                  *String    `xmlrpc:"avatar_1920,omptempty"`
	Avatar256                                   *String    `xmlrpc:"avatar_256,omptempty"`
	Avatar512                                   *String    `xmlrpc:"avatar_512,omptempty"`
	BadgeIds                                    *Relation  `xmlrpc:"badge_ids,omptempty"`
	BankAccountId                               *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	Barcode                                     *String    `xmlrpc:"barcode,omptempty"`
	Birthday                                    *Time      `xmlrpc:"birthday,omptempty"`
	CalendarMismatch                            *Bool      `xmlrpc:"calendar_mismatch,omptempty"`
	CategoryIds                                 *Relation  `xmlrpc:"category_ids,omptempty"`
	Certificate                                 *Selection `xmlrpc:"certificate,omptempty"`
	ChildAllCount                               *Int       `xmlrpc:"child_all_count,omptempty"`
	ChildIds                                    *Relation  `xmlrpc:"child_ids,omptempty"`
	Children                                    *Int       `xmlrpc:"children,omptempty"`
	CoachId                                     *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                                       *Int       `xmlrpc:"color,omptempty"`
	CompanyCountryCode                          *String    `xmlrpc:"company_country_code,omptempty"`
	CompanyCountryId                            *Many2One  `xmlrpc:"company_country_id,omptempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omptempty"`
	ContractId                                  *Many2One  `xmlrpc:"contract_id,omptempty"`
	ContractIds                                 *Relation  `xmlrpc:"contract_ids,omptempty"`
	ContractWarning                             *Bool      `xmlrpc:"contract_warning,omptempty"`
	ContractsCount                              *Int       `xmlrpc:"contracts_count,omptempty"`
	CountryId                                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CountryOfBirth                              *Many2One  `xmlrpc:"country_of_birth,omptempty"`
	CoursesCompletionText                       *String    `xmlrpc:"courses_completion_text,omptempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrentLeaveId                              *Many2One  `xmlrpc:"current_leave_id,omptempty"`
	CurrentLeaveState                           *Selection `xmlrpc:"current_leave_state,omptempty"`
	DefaultPlanningRoleId                       *Many2One  `xmlrpc:"default_planning_role_id,omptempty"`
	DepartmentId                                *Many2One  `xmlrpc:"department_id,omptempty"`
	DepartureDate                               *Time      `xmlrpc:"departure_date,omptempty"`
	DepartureDescription                        *String    `xmlrpc:"departure_description,omptempty"`
	DepartureReasonId                           *Many2One  `xmlrpc:"departure_reason_id,omptempty"`
	DirectBadgeIds                              *Relation  `xmlrpc:"direct_badge_ids,omptempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omptempty"`
	DocumentCount                               *Int       `xmlrpc:"document_count,omptempty"`
	DocumentsShareId                            *Many2One  `xmlrpc:"documents_share_id,omptempty"`
	DrivingLicense                              *String    `xmlrpc:"driving_license,omptempty"`
	EmergencyContact                            *String    `xmlrpc:"emergency_contact,omptempty"`
	EmergencyPhone                              *String    `xmlrpc:"emergency_phone,omptempty"`
	EmployeeSkillIds                            *Relation  `xmlrpc:"employee_skill_ids,omptempty"`
	EmployeeToken                               *String    `xmlrpc:"employee_token,omptempty"`
	EmployeeType                                *Selection `xmlrpc:"employee_type,omptempty"`
	EquipmentCount                              *Int       `xmlrpc:"equipment_count,omptempty"`
	EquipmentIds                                *Relation  `xmlrpc:"equipment_ids,omptempty"`
	ExpenseManagerId                            *Many2One  `xmlrpc:"expense_manager_id,omptempty"`
	FailedMessageIds                            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	FirstContractDate                           *Time      `xmlrpc:"first_contract_date,omptempty"`
	FlexibleHours                               *Bool      `xmlrpc:"flexible_hours,omptempty"`
	Gender                                      *Selection `xmlrpc:"gender,omptempty"`
	GoalIds                                     *Relation  `xmlrpc:"goal_ids,omptempty"`
	HasBadges                                   *Bool      `xmlrpc:"has_badges,omptempty"`
	HasMessage                                  *Bool      `xmlrpc:"has_message,omptempty"`
	HasWorkPermit                               *String    `xmlrpc:"has_work_permit,omptempty"`
	HourlyCost                                  *Float     `xmlrpc:"hourly_cost,omptempty"`
	HoursLastMonth                              *Float     `xmlrpc:"hours_last_month,omptempty"`
	HoursLastMonthDisplay                       *String    `xmlrpc:"hours_last_month_display,omptempty"`
	HoursToday                                  *Float     `xmlrpc:"hours_today,omptempty"`
	HrIconDisplay                               *Selection `xmlrpc:"hr_icon_display,omptempty"`
	HrPresenceState                             *Selection `xmlrpc:"hr_presence_state,omptempty"`
	Id                                          *Int       `xmlrpc:"id,omptempty"`
	IdCard                                      *String    `xmlrpc:"id_card,omptempty"`
	IdentificationId                            *String    `xmlrpc:"identification_id,omptempty"`
	Image1024                                   *String    `xmlrpc:"image_1024,omptempty"`
	Image128                                    *String    `xmlrpc:"image_128,omptempty"`
	Image1920                                   *String    `xmlrpc:"image_1920,omptempty"`
	Image256                                    *String    `xmlrpc:"image_256,omptempty"`
	Image512                                    *String    `xmlrpc:"image_512,omptempty"`
	InternetInvoice                             *String    `xmlrpc:"internet_invoice,omptempty"`
	IsAbsent                                    *Bool      `xmlrpc:"is_absent,omptempty"`
	IsAddressHomeACompany                       *Bool      `xmlrpc:"is_address_home_a_company,omptempty"`
	JobId                                       *Many2One  `xmlrpc:"job_id,omptempty"`
	JobTitle                                    *String    `xmlrpc:"job_title,omptempty"`
	KmHomeWork                                  *Int       `xmlrpc:"km_home_work,omptempty"`
	Lang                                        *Selection `xmlrpc:"lang,omptempty"`
	LastActivity                                *Time      `xmlrpc:"last_activity,omptempty"`
	LastActivityTime                            *String    `xmlrpc:"last_activity_time,omptempty"`
	LastAppraisalDate                           *Time      `xmlrpc:"last_appraisal_date,omptempty"`
	LastAppraisalId                             *Many2One  `xmlrpc:"last_appraisal_id,omptempty"`
	LastAttendanceId                            *Many2One  `xmlrpc:"last_attendance_id,omptempty"`
	LastCheckIn                                 *Time      `xmlrpc:"last_check_in,omptempty"`
	LastCheckOut                                *Time      `xmlrpc:"last_check_out,omptempty"`
	LastValidatedTimesheetDate                  *Time      `xmlrpc:"last_validated_timesheet_date,omptempty"`
	LeaveDateFrom                               *Time      `xmlrpc:"leave_date_from,omptempty"`
	LeaveDateTo                                 *Time      `xmlrpc:"leave_date_to,omptempty"`
	LeaveManagerId                              *Many2One  `xmlrpc:"leave_manager_id,omptempty"`
	LeavesCount                                 *Float     `xmlrpc:"leaves_count,omptempty"`
	Marital                                     *Selection `xmlrpc:"marital,omptempty"`
	MemberOfDepartment                          *Bool      `xmlrpc:"member_of_department,omptempty"`
	MessageAttachmentCount                      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent                              *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds                          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId                     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter                    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MobileInvoice                               *String    `xmlrpc:"mobile_invoice,omptempty"`
	MobilePhone                                 *String    `xmlrpc:"mobile_phone,omptempty"`
	MyActivityDateDeadline                      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                        *String    `xmlrpc:"name,omptempty"`
	NewlyHiredEmployee                          *Bool      `xmlrpc:"newly_hired_employee,omptempty"`
	NextAppraisalDate                           *Time      `xmlrpc:"next_appraisal_date,omptempty"`
	Notes                                       *String    `xmlrpc:"notes,omptempty"`
	OngoingAppraisalCount                       *Int       `xmlrpc:"ongoing_appraisal_count,omptempty"`
	OvertimeIds                                 *Relation  `xmlrpc:"overtime_ids,omptempty"`
	ParentId                                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentUserId                                *Many2One  `xmlrpc:"parent_user_id,omptempty"`
	PassportId                                  *String    `xmlrpc:"passport_id,omptempty"`
	PayslipCount                                *Int       `xmlrpc:"payslip_count,omptempty"`
	PermitNo                                    *String    `xmlrpc:"permit_no,omptempty"`
	Phone                                       *String    `xmlrpc:"phone,omptempty"`
	Pin                                         *String    `xmlrpc:"pin,omptempty"`
	PlaceOfBirth                                *String    `xmlrpc:"place_of_birth,omptempty"`
	PlanningRoleIds                             *Relation  `xmlrpc:"planning_role_ids,omptempty"`
	PrivateEmail                                *String    `xmlrpc:"private_email,omptempty"`
	RegistrationNumber                          *String    `xmlrpc:"registration_number,omptempty"`
	RelatedContactIds                           *Relation  `xmlrpc:"related_contact_ids,omptempty"`
	RelatedContactsCount                        *Int       `xmlrpc:"related_contacts_count,omptempty"`
	RelatedPartnerId                            *Many2One  `xmlrpc:"related_partner_id,omptempty"`
	RemainingLeaves                             *Float     `xmlrpc:"remaining_leaves,omptempty"`
	ResourceCalendarId                          *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceId                                  *Many2One  `xmlrpc:"resource_id,omptempty"`
	ResumeLineIds                               *Relation  `xmlrpc:"resume_line_ids,omptempty"`
	SalaryAttachmentCount                       *Int       `xmlrpc:"salary_attachment_count,omptempty"`
	SalaryAttachmentIds                         *Relation  `xmlrpc:"salary_attachment_ids,omptempty"`
	ShowHrIconDisplay                           *Bool      `xmlrpc:"show_hr_icon_display,omptempty"`
	ShowLeaves                                  *Bool      `xmlrpc:"show_leaves,omptempty"`
	SimCard                                     *String    `xmlrpc:"sim_card,omptempty"`
	Sinid                                       *String    `xmlrpc:"sinid,omptempty"`
	SkillIds                                    *Relation  `xmlrpc:"skill_ids,omptempty"`
	SlipIds                                     *Relation  `xmlrpc:"slip_ids,omptempty"`
	SpouseBirthdate                             *Time      `xmlrpc:"spouse_birthdate,omptempty"`
	SpouseCompleteName                          *String    `xmlrpc:"spouse_complete_name,omptempty"`
	Ssnid                                       *String    `xmlrpc:"ssnid,omptempty"`
	StudyField                                  *String    `xmlrpc:"study_field,omptempty"`
	StudySchool                                 *String    `xmlrpc:"study_school,omptempty"`
	SubordinateIds                              *Relation  `xmlrpc:"subordinate_ids,omptempty"`
	SubscribedCourses                           *Relation  `xmlrpc:"subscribed_courses,omptempty"`
	TimesheetManagerId                          *Many2One  `xmlrpc:"timesheet_manager_id,omptempty"`
	TotalOvertime                               *Float     `xmlrpc:"total_overtime,omptempty"`
	Tz                                          *Selection `xmlrpc:"tz,omptempty"`
	UncompleteGoalsCount                        *Int       `xmlrpc:"uncomplete_goals_count,omptempty"`
	UserId                                      *Many2One  `xmlrpc:"user_id,omptempty"`
	UserPartnerId                               *Many2One  `xmlrpc:"user_partner_id,omptempty"`
	Vehicle                                     *String    `xmlrpc:"vehicle,omptempty"`
	VisaExpire                                  *Time      `xmlrpc:"visa_expire,omptempty"`
	VisaNo                                      *String    `xmlrpc:"visa_no,omptempty"`
	WebsiteMessageIds                           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WorkContactId                               *Many2One  `xmlrpc:"work_contact_id,omptempty"`
	WorkEmail                                   *String    `xmlrpc:"work_email,omptempty"`
	WorkLocationId                              *Many2One  `xmlrpc:"work_location_id,omptempty"`
	WorkPermitExpirationDate                    *Time      `xmlrpc:"work_permit_expiration_date,omptempty"`
	WorkPermitName                              *String    `xmlrpc:"work_permit_name,omptempty"`
	WorkPermitScheduledActivity                 *Bool      `xmlrpc:"work_permit_scheduled_activity,omptempty"`
	WorkPhone                                   *String    `xmlrpc:"work_phone,omptempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omptempty"`
	XStudioBirthdayChildren                     *Relation  `xmlrpc:"x_studio_birthday_children,omptempty"`
	XStudioConvertHolidayToOvertimeAtEndOfYear  *Selection `xmlrpc:"x_studio_convert_holiday_to_overtime_at_end_of_year,omptempty"`
	XStudioEmployeeQstCode                      *Selection `xmlrpc:"x_studio_employee_qst_code,omptempty"`
	XStudioOne2ManyFieldA4LEu                   *Relation  `xmlrpc:"x_studio_one2many_field_a4lEu,omptempty"`
	XXStudioMany2OneField3VCYYAccountAssetCount *Int       `xmlrpc:"x_x_studio_many2one_field_3vCYY_account_asset_count,omptempty"`
}

// HrEmployees represents array of hr.employee model.
type HrEmployees []HrEmployee

// HrEmployeeModel is the odoo model name.
const HrEmployeeModel = "hr.employee"

// Many2One convert HrEmployee to *Many2One.
func (he *HrEmployee) Many2One() *Many2One {
	return NewMany2One(he.Id.Get(), "")
}

// CreateHrEmployee creates a new hr.employee model and returns its id.
func (c *Client) CreateHrEmployee(he *HrEmployee) (int64, error) {
	ids, err := c.CreateHrEmployees([]*HrEmployee{he})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployee creates a new hr.employee model and returns its id.
func (c *Client) CreateHrEmployees(hes []*HrEmployee) ([]int64, error) {
	var vv []interface{}
	for _, v := range hes {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeModel, vv)
}

// UpdateHrEmployee updates an existing hr.employee record.
func (c *Client) UpdateHrEmployee(he *HrEmployee) error {
	return c.UpdateHrEmployees([]int64{he.Id.Get()}, he)
}

// UpdateHrEmployees updates existing hr.employee records.
// All records (represented by ids) will be updated by he values.
func (c *Client) UpdateHrEmployees(ids []int64, he *HrEmployee) error {
	return c.Update(HrEmployeeModel, ids, he)
}

// DeleteHrEmployee deletes an existing hr.employee record.
func (c *Client) DeleteHrEmployee(id int64) error {
	return c.DeleteHrEmployees([]int64{id})
}

// DeleteHrEmployees deletes existing hr.employee records.
func (c *Client) DeleteHrEmployees(ids []int64) error {
	return c.Delete(HrEmployeeModel, ids)
}

// GetHrEmployee gets hr.employee existing record.
func (c *Client) GetHrEmployee(id int64) (*HrEmployee, error) {
	hes, err := c.GetHrEmployees([]int64{id})
	if err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.employee not found", id)
}

// GetHrEmployees gets hr.employee existing records.
func (c *Client) GetHrEmployees(ids []int64) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.Read(HrEmployeeModel, ids, nil, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployee finds hr.employee record by querying it with criteria.
func (c *Client) FindHrEmployee(criteria *Criteria) (*HrEmployee, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, NewOptions().Limit(1), hes); err != nil {
		return nil, err
	}
	if hes != nil && len(*hes) > 0 {
		return &((*hes)[0]), nil
	}
	return nil, fmt.Errorf("hr.employee was not found with criteria %v", criteria)
}

// FindHrEmployees finds hr.employee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployees(criteria *Criteria, options *Options) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, options, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrEmployeeId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.employee was not found with criteria %v and options %v", criteria, options)
}
