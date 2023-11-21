package odoo

import (
	"fmt"
)

// HrEmployeePublic represents hr.employee.public model.
type HrEmployeePublic struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	Active                     *Bool      `xmlrpc:"active,omptempty"`
	AddressId                  *Many2One  `xmlrpc:"address_id,omptempty"`
	AllocationCount            *Float     `xmlrpc:"allocation_count,omptempty"`
	AllocationDisplay          *String    `xmlrpc:"allocation_display,omptempty"`
	AllocationRemainingDisplay *String    `xmlrpc:"allocation_remaining_display,omptempty"`
	AllocationsCount           *Int       `xmlrpc:"allocations_count,omptempty"`
	AttendanceState            *Selection `xmlrpc:"attendance_state,omptempty"`
	Avatar1024                 *String    `xmlrpc:"avatar_1024,omptempty"`
	Avatar128                  *String    `xmlrpc:"avatar_128,omptempty"`
	Avatar1920                 *String    `xmlrpc:"avatar_1920,omptempty"`
	Avatar256                  *String    `xmlrpc:"avatar_256,omptempty"`
	Avatar512                  *String    `xmlrpc:"avatar_512,omptempty"`
	BadgeIds                   *Relation  `xmlrpc:"badge_ids,omptempty"`
	ChildAllCount              *Int       `xmlrpc:"child_all_count,omptempty"`
	ChildIds                   *Relation  `xmlrpc:"child_ids,omptempty"`
	CoachId                    *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                      *Int       `xmlrpc:"color,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrentLeaveState          *Selection `xmlrpc:"current_leave_state,omptempty"`
	DepartmentId               *Many2One  `xmlrpc:"department_id,omptempty"`
	DirectBadgeIds             *Relation  `xmlrpc:"direct_badge_ids,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId                 *Many2One  `xmlrpc:"employee_id,omptempty"`
	EmployeeSkillIds           *Relation  `xmlrpc:"employee_skill_ids,omptempty"`
	EmployeeType               *Selection `xmlrpc:"employee_type,omptempty"`
	ExpenseManagerId           *Many2One  `xmlrpc:"expense_manager_id,omptempty"`
	FirstContractDate          *Time      `xmlrpc:"first_contract_date,omptempty"`
	GoalIds                    *Relation  `xmlrpc:"goal_ids,omptempty"`
	HasBadges                  *Bool      `xmlrpc:"has_badges,omptempty"`
	HoursToday                 *Float     `xmlrpc:"hours_today,omptempty"`
	HrIconDisplay              *Selection `xmlrpc:"hr_icon_display,omptempty"`
	HrPresenceState            *Selection `xmlrpc:"hr_presence_state,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	Image1024                  *String    `xmlrpc:"image_1024,omptempty"`
	Image128                   *String    `xmlrpc:"image_128,omptempty"`
	Image1920                  *String    `xmlrpc:"image_1920,omptempty"`
	Image256                   *String    `xmlrpc:"image_256,omptempty"`
	Image512                   *String    `xmlrpc:"image_512,omptempty"`
	IsAbsent                   *Bool      `xmlrpc:"is_absent,omptempty"`
	JobId                      *Many2One  `xmlrpc:"job_id,omptempty"`
	JobTitle                   *String    `xmlrpc:"job_title,omptempty"`
	LastActivity               *Time      `xmlrpc:"last_activity,omptempty"`
	LastActivityTime           *String    `xmlrpc:"last_activity_time,omptempty"`
	LastAppraisalId            *Many2One  `xmlrpc:"last_appraisal_id,omptempty"`
	LastAttendanceId           *Many2One  `xmlrpc:"last_attendance_id,omptempty"`
	LeaveDateFrom              *Time      `xmlrpc:"leave_date_from,omptempty"`
	LeaveDateTo                *Time      `xmlrpc:"leave_date_to,omptempty"`
	LeaveManagerId             *Many2One  `xmlrpc:"leave_manager_id,omptempty"`
	LeavesCount                *Float     `xmlrpc:"leaves_count,omptempty"`
	MemberOfDepartment         *Bool      `xmlrpc:"member_of_department,omptempty"`
	MobilePhone                *String    `xmlrpc:"mobile_phone,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	ParentId                   *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentUserId               *Many2One  `xmlrpc:"parent_user_id,omptempty"`
	RelatedContactIds          *Relation  `xmlrpc:"related_contact_ids,omptempty"`
	RelatedContactsCount       *Int       `xmlrpc:"related_contacts_count,omptempty"`
	RemainingLeaves            *Float     `xmlrpc:"remaining_leaves,omptempty"`
	ResourceCalendarId         *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceId                 *Many2One  `xmlrpc:"resource_id,omptempty"`
	ResumeLineIds              *Relation  `xmlrpc:"resume_line_ids,omptempty"`
	ShowHrIconDisplay          *Bool      `xmlrpc:"show_hr_icon_display,omptempty"`
	ShowLeaves                 *Bool      `xmlrpc:"show_leaves,omptempty"`
	SubordinateIds             *Relation  `xmlrpc:"subordinate_ids,omptempty"`
	TimesheetManagerId         *Many2One  `xmlrpc:"timesheet_manager_id,omptempty"`
	TotalOvertime              *Float     `xmlrpc:"total_overtime,omptempty"`
	Tz                         *Selection `xmlrpc:"tz,omptempty"`
	UserId                     *Many2One  `xmlrpc:"user_id,omptempty"`
	UserPartnerId              *Many2One  `xmlrpc:"user_partner_id,omptempty"`
	WorkContactId              *Many2One  `xmlrpc:"work_contact_id,omptempty"`
	WorkEmail                  *String    `xmlrpc:"work_email,omptempty"`
	WorkLocationId             *Many2One  `xmlrpc:"work_location_id,omptempty"`
	WorkPhone                  *String    `xmlrpc:"work_phone,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrEmployeePublics represents array of hr.employee.public model.
type HrEmployeePublics []HrEmployeePublic

// HrEmployeePublicModel is the odoo model name.
const HrEmployeePublicModel = "hr.employee.public"

// Many2One convert HrEmployeePublic to *Many2One.
func (hep *HrEmployeePublic) Many2One() *Many2One {
	return NewMany2One(hep.Id.Get(), "")
}

// CreateHrEmployeePublic creates a new hr.employee.public model and returns its id.
func (c *Client) CreateHrEmployeePublic(hep *HrEmployeePublic) (int64, error) {
	ids, err := c.CreateHrEmployeePublics([]*HrEmployeePublic{hep})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeePublic creates a new hr.employee.public model and returns its id.
func (c *Client) CreateHrEmployeePublics(heps []*HrEmployeePublic) ([]int64, error) {
	var vv []interface{}
	for _, v := range heps {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeePublicModel, vv)
}

// UpdateHrEmployeePublic updates an existing hr.employee.public record.
func (c *Client) UpdateHrEmployeePublic(hep *HrEmployeePublic) error {
	return c.UpdateHrEmployeePublics([]int64{hep.Id.Get()}, hep)
}

// UpdateHrEmployeePublics updates existing hr.employee.public records.
// All records (represented by ids) will be updated by hep values.
func (c *Client) UpdateHrEmployeePublics(ids []int64, hep *HrEmployeePublic) error {
	return c.Update(HrEmployeePublicModel, ids, hep)
}

// DeleteHrEmployeePublic deletes an existing hr.employee.public record.
func (c *Client) DeleteHrEmployeePublic(id int64) error {
	return c.DeleteHrEmployeePublics([]int64{id})
}

// DeleteHrEmployeePublics deletes existing hr.employee.public records.
func (c *Client) DeleteHrEmployeePublics(ids []int64) error {
	return c.Delete(HrEmployeePublicModel, ids)
}

// GetHrEmployeePublic gets hr.employee.public existing record.
func (c *Client) GetHrEmployeePublic(id int64) (*HrEmployeePublic, error) {
	heps, err := c.GetHrEmployeePublics([]int64{id})
	if err != nil {
		return nil, err
	}
	if heps != nil && len(*heps) > 0 {
		return &((*heps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.employee.public not found", id)
}

// GetHrEmployeePublics gets hr.employee.public existing records.
func (c *Client) GetHrEmployeePublics(ids []int64) (*HrEmployeePublics, error) {
	heps := &HrEmployeePublics{}
	if err := c.Read(HrEmployeePublicModel, ids, nil, heps); err != nil {
		return nil, err
	}
	return heps, nil
}

// FindHrEmployeePublic finds hr.employee.public record by querying it with criteria.
func (c *Client) FindHrEmployeePublic(criteria *Criteria) (*HrEmployeePublic, error) {
	heps := &HrEmployeePublics{}
	if err := c.SearchRead(HrEmployeePublicModel, criteria, NewOptions().Limit(1), heps); err != nil {
		return nil, err
	}
	if heps != nil && len(*heps) > 0 {
		return &((*heps)[0]), nil
	}
	return nil, fmt.Errorf("hr.employee.public was not found with criteria %v", criteria)
}

// FindHrEmployeePublics finds hr.employee.public records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeePublics(criteria *Criteria, options *Options) (*HrEmployeePublics, error) {
	heps := &HrEmployeePublics{}
	if err := c.SearchRead(HrEmployeePublicModel, criteria, options, heps); err != nil {
		return nil, err
	}
	return heps, nil
}

// FindHrEmployeePublicIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeePublicIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrEmployeePublicModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrEmployeePublicId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeePublicId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeePublicModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.employee.public was not found with criteria %v and options %v", criteria, options)
}
