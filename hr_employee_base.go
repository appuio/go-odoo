package odoo

import (
	"fmt"
)

// HrEmployeeBase represents hr.employee.base model.
type HrEmployeeBase struct {
	Active                     *Bool      `xmlrpc:"active,omptempty"`
	AddressId                  *Many2One  `xmlrpc:"address_id,omptempty"`
	AllocationCount            *Float     `xmlrpc:"allocation_count,omptempty"`
	AllocationDisplay          *String    `xmlrpc:"allocation_display,omptempty"`
	AllocationRemainingDisplay *String    `xmlrpc:"allocation_remaining_display,omptempty"`
	AllocationsCount           *Int       `xmlrpc:"allocations_count,omptempty"`
	BadgeIds                   *Relation  `xmlrpc:"badge_ids,omptempty"`
	ChildAllCount              *Int       `xmlrpc:"child_all_count,omptempty"`
	CoachId                    *Many2One  `xmlrpc:"coach_id,omptempty"`
	Color                      *Int       `xmlrpc:"color,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	CurrentLeaveState          *Selection `xmlrpc:"current_leave_state,omptempty"`
	DepartmentId               *Many2One  `xmlrpc:"department_id,omptempty"`
	DirectBadgeIds             *Relation  `xmlrpc:"direct_badge_ids,omptempty"`
	EmployeeType               *Selection `xmlrpc:"employee_type,omptempty"`
	GoalIds                    *Relation  `xmlrpc:"goal_ids,omptempty"`
	HasBadges                  *Bool      `xmlrpc:"has_badges,omptempty"`
	HrIconDisplay              *Selection `xmlrpc:"hr_icon_display,omptempty"`
	HrPresenceState            *Selection `xmlrpc:"hr_presence_state,omptempty"`
	IsAbsent                   *Bool      `xmlrpc:"is_absent,omptempty"`
	JobId                      *Many2One  `xmlrpc:"job_id,omptempty"`
	JobTitle                   *String    `xmlrpc:"job_title,omptempty"`
	LastActivity               *Time      `xmlrpc:"last_activity,omptempty"`
	LastActivityTime           *String    `xmlrpc:"last_activity_time,omptempty"`
	LastAppraisalId            *Many2One  `xmlrpc:"last_appraisal_id,omptempty"`
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
	ShowHrIconDisplay          *Bool      `xmlrpc:"show_hr_icon_display,omptempty"`
	ShowLeaves                 *Bool      `xmlrpc:"show_leaves,omptempty"`
	Tz                         *Selection `xmlrpc:"tz,omptempty"`
	UserId                     *Many2One  `xmlrpc:"user_id,omptempty"`
	WorkContactId              *Many2One  `xmlrpc:"work_contact_id,omptempty"`
	WorkEmail                  *String    `xmlrpc:"work_email,omptempty"`
	WorkLocationId             *Many2One  `xmlrpc:"work_location_id,omptempty"`
	WorkPhone                  *String    `xmlrpc:"work_phone,omptempty"`
}

// HrEmployeeBases represents array of hr.employee.base model.
type HrEmployeeBases []HrEmployeeBase

// HrEmployeeBaseModel is the odoo model name.
const HrEmployeeBaseModel = "hr.employee.base"

// Many2One convert HrEmployeeBase to *Many2One.
func (heb *HrEmployeeBase) Many2One() *Many2One {
	return NewMany2One(heb.Id.Get(), "")
}

// CreateHrEmployeeBase creates a new hr.employee.base model and returns its id.
func (c *Client) CreateHrEmployeeBase(heb *HrEmployeeBase) (int64, error) {
	ids, err := c.CreateHrEmployeeBases([]*HrEmployeeBase{heb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeBase creates a new hr.employee.base model and returns its id.
func (c *Client) CreateHrEmployeeBases(hebs []*HrEmployeeBase) ([]int64, error) {
	var vv []interface{}
	for _, v := range hebs {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeBaseModel, vv)
}

// UpdateHrEmployeeBase updates an existing hr.employee.base record.
func (c *Client) UpdateHrEmployeeBase(heb *HrEmployeeBase) error {
	return c.UpdateHrEmployeeBases([]int64{heb.Id.Get()}, heb)
}

// UpdateHrEmployeeBases updates existing hr.employee.base records.
// All records (represented by ids) will be updated by heb values.
func (c *Client) UpdateHrEmployeeBases(ids []int64, heb *HrEmployeeBase) error {
	return c.Update(HrEmployeeBaseModel, ids, heb)
}

// DeleteHrEmployeeBase deletes an existing hr.employee.base record.
func (c *Client) DeleteHrEmployeeBase(id int64) error {
	return c.DeleteHrEmployeeBases([]int64{id})
}

// DeleteHrEmployeeBases deletes existing hr.employee.base records.
func (c *Client) DeleteHrEmployeeBases(ids []int64) error {
	return c.Delete(HrEmployeeBaseModel, ids)
}

// GetHrEmployeeBase gets hr.employee.base existing record.
func (c *Client) GetHrEmployeeBase(id int64) (*HrEmployeeBase, error) {
	hebs, err := c.GetHrEmployeeBases([]int64{id})
	if err != nil {
		return nil, err
	}
	if hebs != nil && len(*hebs) > 0 {
		return &((*hebs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.employee.base not found", id)
}

// GetHrEmployeeBases gets hr.employee.base existing records.
func (c *Client) GetHrEmployeeBases(ids []int64) (*HrEmployeeBases, error) {
	hebs := &HrEmployeeBases{}
	if err := c.Read(HrEmployeeBaseModel, ids, nil, hebs); err != nil {
		return nil, err
	}
	return hebs, nil
}

// FindHrEmployeeBase finds hr.employee.base record by querying it with criteria.
func (c *Client) FindHrEmployeeBase(criteria *Criteria) (*HrEmployeeBase, error) {
	hebs := &HrEmployeeBases{}
	if err := c.SearchRead(HrEmployeeBaseModel, criteria, NewOptions().Limit(1), hebs); err != nil {
		return nil, err
	}
	if hebs != nil && len(*hebs) > 0 {
		return &((*hebs)[0]), nil
	}
	return nil, fmt.Errorf("hr.employee.base was not found with criteria %v", criteria)
}

// FindHrEmployeeBases finds hr.employee.base records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeBases(criteria *Criteria, options *Options) (*HrEmployeeBases, error) {
	hebs := &HrEmployeeBases{}
	if err := c.SearchRead(HrEmployeeBaseModel, criteria, options, hebs); err != nil {
		return nil, err
	}
	return hebs, nil
}

// FindHrEmployeeBaseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeBaseIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrEmployeeBaseModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrEmployeeBaseId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeBaseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeBaseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.employee.base was not found with criteria %v and options %v", criteria, options)
}
