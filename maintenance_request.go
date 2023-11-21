package odoo

import (
	"fmt"
)

// MaintenanceRequest represents maintenance.request model.
type MaintenanceRequest struct {
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
	Archive                     *Bool      `xmlrpc:"archive,omptempty"`
	CategoryId                  *Many2One  `xmlrpc:"category_id,omptempty"`
	CloseDate                   *Time      `xmlrpc:"close_date,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	Done                        *Bool      `xmlrpc:"done,omptempty"`
	Duration                    *Float     `xmlrpc:"duration,omptempty"`
	EmailCc                     *String    `xmlrpc:"email_cc,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	EquipmentId                 *Many2One  `xmlrpc:"equipment_id,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	KanbanState                 *Selection `xmlrpc:"kanban_state,omptempty"`
	MaintenanceTeamId           *Many2One  `xmlrpc:"maintenance_team_id,omptempty"`
	MaintenanceType             *Selection `xmlrpc:"maintenance_type,omptempty"`
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
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	OwnerUserId                 *Many2One  `xmlrpc:"owner_user_id,omptempty"`
	Priority                    *Selection `xmlrpc:"priority,omptempty"`
	RequestDate                 *Time      `xmlrpc:"request_date,omptempty"`
	ScheduleDate                *Time      `xmlrpc:"schedule_date,omptempty"`
	StageId                     *Many2One  `xmlrpc:"stage_id,omptempty"`
	UserId                      *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MaintenanceRequests represents array of maintenance.request model.
type MaintenanceRequests []MaintenanceRequest

// MaintenanceRequestModel is the odoo model name.
const MaintenanceRequestModel = "maintenance.request"

// Many2One convert MaintenanceRequest to *Many2One.
func (mr *MaintenanceRequest) Many2One() *Many2One {
	return NewMany2One(mr.Id.Get(), "")
}

// CreateMaintenanceRequest creates a new maintenance.request model and returns its id.
func (c *Client) CreateMaintenanceRequest(mr *MaintenanceRequest) (int64, error) {
	ids, err := c.CreateMaintenanceRequests([]*MaintenanceRequest{mr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMaintenanceRequest creates a new maintenance.request model and returns its id.
func (c *Client) CreateMaintenanceRequests(mrs []*MaintenanceRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrs {
		vv = append(vv, v)
	}
	return c.Create(MaintenanceRequestModel, vv)
}

// UpdateMaintenanceRequest updates an existing maintenance.request record.
func (c *Client) UpdateMaintenanceRequest(mr *MaintenanceRequest) error {
	return c.UpdateMaintenanceRequests([]int64{mr.Id.Get()}, mr)
}

// UpdateMaintenanceRequests updates existing maintenance.request records.
// All records (represented by ids) will be updated by mr values.
func (c *Client) UpdateMaintenanceRequests(ids []int64, mr *MaintenanceRequest) error {
	return c.Update(MaintenanceRequestModel, ids, mr)
}

// DeleteMaintenanceRequest deletes an existing maintenance.request record.
func (c *Client) DeleteMaintenanceRequest(id int64) error {
	return c.DeleteMaintenanceRequests([]int64{id})
}

// DeleteMaintenanceRequests deletes existing maintenance.request records.
func (c *Client) DeleteMaintenanceRequests(ids []int64) error {
	return c.Delete(MaintenanceRequestModel, ids)
}

// GetMaintenanceRequest gets maintenance.request existing record.
func (c *Client) GetMaintenanceRequest(id int64) (*MaintenanceRequest, error) {
	mrs, err := c.GetMaintenanceRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	if mrs != nil && len(*mrs) > 0 {
		return &((*mrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of maintenance.request not found", id)
}

// GetMaintenanceRequests gets maintenance.request existing records.
func (c *Client) GetMaintenanceRequests(ids []int64) (*MaintenanceRequests, error) {
	mrs := &MaintenanceRequests{}
	if err := c.Read(MaintenanceRequestModel, ids, nil, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMaintenanceRequest finds maintenance.request record by querying it with criteria.
func (c *Client) FindMaintenanceRequest(criteria *Criteria) (*MaintenanceRequest, error) {
	mrs := &MaintenanceRequests{}
	if err := c.SearchRead(MaintenanceRequestModel, criteria, NewOptions().Limit(1), mrs); err != nil {
		return nil, err
	}
	if mrs != nil && len(*mrs) > 0 {
		return &((*mrs)[0]), nil
	}
	return nil, fmt.Errorf("maintenance.request was not found with criteria %v", criteria)
}

// FindMaintenanceRequests finds maintenance.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceRequests(criteria *Criteria, options *Options) (*MaintenanceRequests, error) {
	mrs := &MaintenanceRequests{}
	if err := c.SearchRead(MaintenanceRequestModel, criteria, options, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMaintenanceRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MaintenanceRequestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMaintenanceRequestId finds record id by querying it with criteria.
func (c *Client) FindMaintenanceRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MaintenanceRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("maintenance.request was not found with criteria %v and options %v", criteria, options)
}
