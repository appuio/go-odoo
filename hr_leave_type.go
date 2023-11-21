package odoo

import (
	"fmt"
)

// HrLeaveType represents hr.leave.type model.
type HrLeaveType struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	AccrualCount              *Float     `xmlrpc:"accrual_count,omptempty"`
	AccrualsIds               *Relation  `xmlrpc:"accruals_ids,omptempty"`
	Active                    *Bool      `xmlrpc:"active,omptempty"`
	AllocationCount           *Int       `xmlrpc:"allocation_count,omptempty"`
	AllocationNotifSubtypeId  *Many2One  `xmlrpc:"allocation_notif_subtype_id,omptempty"`
	AllocationValidationType  *Selection `xmlrpc:"allocation_validation_type,omptempty"`
	ClosestAllocationToExpire *Many2One  `xmlrpc:"closest_allocation_to_expire,omptempty"`
	Color                     *Int       `xmlrpc:"color,omptempty"`
	ColorName                 *Selection `xmlrpc:"color_name,omptempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateCalendarMeeting     *Bool      `xmlrpc:"create_calendar_meeting,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	EmployeeRequests          *Selection `xmlrpc:"employee_requests,omptempty"`
	GroupDaysLeave            *Float     `xmlrpc:"group_days_leave,omptempty"`
	HasValidAllocation        *Bool      `xmlrpc:"has_valid_allocation,omptempty"`
	HrAttendanceOvertime      *Bool      `xmlrpc:"hr_attendance_overtime,omptempty"`
	IconId                    *Many2One  `xmlrpc:"icon_id,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	LeaveNotifSubtypeId       *Many2One  `xmlrpc:"leave_notif_subtype_id,omptempty"`
	LeaveValidationType       *Selection `xmlrpc:"leave_validation_type,omptempty"`
	LeavesTaken               *Float     `xmlrpc:"leaves_taken,omptempty"`
	MaxLeaves                 *Float     `xmlrpc:"max_leaves,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	OvertimeDeductible        *Bool      `xmlrpc:"overtime_deductible,omptempty"`
	RemainingLeaves           *Float     `xmlrpc:"remaining_leaves,omptempty"`
	RequestUnit               *Selection `xmlrpc:"request_unit,omptempty"`
	RequiresAllocation        *Selection `xmlrpc:"requires_allocation,omptempty"`
	ResponsibleId             *Many2One  `xmlrpc:"responsible_id,omptempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omptempty"`
	SupportDocument           *Bool      `xmlrpc:"support_document,omptempty"`
	TimeType                  *Selection `xmlrpc:"time_type,omptempty"`
	TimesheetGenerate         *Bool      `xmlrpc:"timesheet_generate,omptempty"`
	TimesheetProjectId        *Many2One  `xmlrpc:"timesheet_project_id,omptempty"`
	TimesheetTaskId           *Many2One  `xmlrpc:"timesheet_task_id,omptempty"`
	Unpaid                    *Bool      `xmlrpc:"unpaid,omptempty"`
	VirtualLeavesTaken        *Float     `xmlrpc:"virtual_leaves_taken,omptempty"`
	VirtualRemainingLeaves    *Float     `xmlrpc:"virtual_remaining_leaves,omptempty"`
	WorkEntryTypeId           *Many2One  `xmlrpc:"work_entry_type_id,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrLeaveTypes represents array of hr.leave.type model.
type HrLeaveTypes []HrLeaveType

// HrLeaveTypeModel is the odoo model name.
const HrLeaveTypeModel = "hr.leave.type"

// Many2One convert HrLeaveType to *Many2One.
func (hlt *HrLeaveType) Many2One() *Many2One {
	return NewMany2One(hlt.Id.Get(), "")
}

// CreateHrLeaveType creates a new hr.leave.type model and returns its id.
func (c *Client) CreateHrLeaveType(hlt *HrLeaveType) (int64, error) {
	ids, err := c.CreateHrLeaveTypes([]*HrLeaveType{hlt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrLeaveType creates a new hr.leave.type model and returns its id.
func (c *Client) CreateHrLeaveTypes(hlts []*HrLeaveType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hlts {
		vv = append(vv, v)
	}
	return c.Create(HrLeaveTypeModel, vv)
}

// UpdateHrLeaveType updates an existing hr.leave.type record.
func (c *Client) UpdateHrLeaveType(hlt *HrLeaveType) error {
	return c.UpdateHrLeaveTypes([]int64{hlt.Id.Get()}, hlt)
}

// UpdateHrLeaveTypes updates existing hr.leave.type records.
// All records (represented by ids) will be updated by hlt values.
func (c *Client) UpdateHrLeaveTypes(ids []int64, hlt *HrLeaveType) error {
	return c.Update(HrLeaveTypeModel, ids, hlt)
}

// DeleteHrLeaveType deletes an existing hr.leave.type record.
func (c *Client) DeleteHrLeaveType(id int64) error {
	return c.DeleteHrLeaveTypes([]int64{id})
}

// DeleteHrLeaveTypes deletes existing hr.leave.type records.
func (c *Client) DeleteHrLeaveTypes(ids []int64) error {
	return c.Delete(HrLeaveTypeModel, ids)
}

// GetHrLeaveType gets hr.leave.type existing record.
func (c *Client) GetHrLeaveType(id int64) (*HrLeaveType, error) {
	hlts, err := c.GetHrLeaveTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if hlts != nil && len(*hlts) > 0 {
		return &((*hlts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.leave.type not found", id)
}

// GetHrLeaveTypes gets hr.leave.type existing records.
func (c *Client) GetHrLeaveTypes(ids []int64) (*HrLeaveTypes, error) {
	hlts := &HrLeaveTypes{}
	if err := c.Read(HrLeaveTypeModel, ids, nil, hlts); err != nil {
		return nil, err
	}
	return hlts, nil
}

// FindHrLeaveType finds hr.leave.type record by querying it with criteria.
func (c *Client) FindHrLeaveType(criteria *Criteria) (*HrLeaveType, error) {
	hlts := &HrLeaveTypes{}
	if err := c.SearchRead(HrLeaveTypeModel, criteria, NewOptions().Limit(1), hlts); err != nil {
		return nil, err
	}
	if hlts != nil && len(*hlts) > 0 {
		return &((*hlts)[0]), nil
	}
	return nil, fmt.Errorf("hr.leave.type was not found with criteria %v", criteria)
}

// FindHrLeaveTypes finds hr.leave.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveTypes(criteria *Criteria, options *Options) (*HrLeaveTypes, error) {
	hlts := &HrLeaveTypes{}
	if err := c.SearchRead(HrLeaveTypeModel, criteria, options, hlts); err != nil {
		return nil, err
	}
	return hlts, nil
}

// FindHrLeaveTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrLeaveTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrLeaveTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrLeaveTypeId finds record id by querying it with criteria.
func (c *Client) FindHrLeaveTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrLeaveTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.leave.type was not found with criteria %v and options %v", criteria, options)
}
