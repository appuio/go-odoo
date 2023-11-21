package odoo

import (
	"fmt"
)

// MaintenanceEquipment represents maintenance.equipment model.
type MaintenanceEquipment struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
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
	AssignDate                  *Time      `xmlrpc:"assign_date,omptempty"`
	CategoryId                  *Many2One  `xmlrpc:"category_id,omptempty"`
	Color                       *Int       `xmlrpc:"color,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	Cost                        *Float     `xmlrpc:"cost,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EffectiveDate               *Time      `xmlrpc:"effective_date,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	EquipmentAssignTo           *Selection `xmlrpc:"equipment_assign_to,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	Location                    *String    `xmlrpc:"location,omptempty"`
	MaintenanceCount            *Int       `xmlrpc:"maintenance_count,omptempty"`
	MaintenanceDuration         *Float     `xmlrpc:"maintenance_duration,omptempty"`
	MaintenanceIds              *Relation  `xmlrpc:"maintenance_ids,omptempty"`
	MaintenanceOpenCount        *Int       `xmlrpc:"maintenance_open_count,omptempty"`
	MaintenanceTeamId           *Many2One  `xmlrpc:"maintenance_team_id,omptempty"`
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
	Name                        *String    `xmlrpc:"name,omptempty"`
	NextActionDate              *Time      `xmlrpc:"next_action_date,omptempty"`
	Note                        *String    `xmlrpc:"note,omptempty"`
	OwnerUserId                 *Many2One  `xmlrpc:"owner_user_id,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerRef                  *String    `xmlrpc:"partner_ref,omptempty"`
	Period                      *Int       `xmlrpc:"period,omptempty"`
	ScrapDate                   *Time      `xmlrpc:"scrap_date,omptempty"`
	SerialNo                    *String    `xmlrpc:"serial_no,omptempty"`
	TechnicianUserId            *Many2One  `xmlrpc:"technician_user_id,omptempty"`
	WarrantyDate                *Time      `xmlrpc:"warranty_date,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MaintenanceEquipments represents array of maintenance.equipment model.
type MaintenanceEquipments []MaintenanceEquipment

// MaintenanceEquipmentModel is the odoo model name.
const MaintenanceEquipmentModel = "maintenance.equipment"

// Many2One convert MaintenanceEquipment to *Many2One.
func (me *MaintenanceEquipment) Many2One() *Many2One {
	return NewMany2One(me.Id.Get(), "")
}

// CreateMaintenanceEquipment creates a new maintenance.equipment model and returns its id.
func (c *Client) CreateMaintenanceEquipment(me *MaintenanceEquipment) (int64, error) {
	ids, err := c.CreateMaintenanceEquipments([]*MaintenanceEquipment{me})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMaintenanceEquipment creates a new maintenance.equipment model and returns its id.
func (c *Client) CreateMaintenanceEquipments(mes []*MaintenanceEquipment) ([]int64, error) {
	var vv []interface{}
	for _, v := range mes {
		vv = append(vv, v)
	}
	return c.Create(MaintenanceEquipmentModel, vv)
}

// UpdateMaintenanceEquipment updates an existing maintenance.equipment record.
func (c *Client) UpdateMaintenanceEquipment(me *MaintenanceEquipment) error {
	return c.UpdateMaintenanceEquipments([]int64{me.Id.Get()}, me)
}

// UpdateMaintenanceEquipments updates existing maintenance.equipment records.
// All records (represented by ids) will be updated by me values.
func (c *Client) UpdateMaintenanceEquipments(ids []int64, me *MaintenanceEquipment) error {
	return c.Update(MaintenanceEquipmentModel, ids, me)
}

// DeleteMaintenanceEquipment deletes an existing maintenance.equipment record.
func (c *Client) DeleteMaintenanceEquipment(id int64) error {
	return c.DeleteMaintenanceEquipments([]int64{id})
}

// DeleteMaintenanceEquipments deletes existing maintenance.equipment records.
func (c *Client) DeleteMaintenanceEquipments(ids []int64) error {
	return c.Delete(MaintenanceEquipmentModel, ids)
}

// GetMaintenanceEquipment gets maintenance.equipment existing record.
func (c *Client) GetMaintenanceEquipment(id int64) (*MaintenanceEquipment, error) {
	mes, err := c.GetMaintenanceEquipments([]int64{id})
	if err != nil {
		return nil, err
	}
	if mes != nil && len(*mes) > 0 {
		return &((*mes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of maintenance.equipment not found", id)
}

// GetMaintenanceEquipments gets maintenance.equipment existing records.
func (c *Client) GetMaintenanceEquipments(ids []int64) (*MaintenanceEquipments, error) {
	mes := &MaintenanceEquipments{}
	if err := c.Read(MaintenanceEquipmentModel, ids, nil, mes); err != nil {
		return nil, err
	}
	return mes, nil
}

// FindMaintenanceEquipment finds maintenance.equipment record by querying it with criteria.
func (c *Client) FindMaintenanceEquipment(criteria *Criteria) (*MaintenanceEquipment, error) {
	mes := &MaintenanceEquipments{}
	if err := c.SearchRead(MaintenanceEquipmentModel, criteria, NewOptions().Limit(1), mes); err != nil {
		return nil, err
	}
	if mes != nil && len(*mes) > 0 {
		return &((*mes)[0]), nil
	}
	return nil, fmt.Errorf("maintenance.equipment was not found with criteria %v", criteria)
}

// FindMaintenanceEquipments finds maintenance.equipment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceEquipments(criteria *Criteria, options *Options) (*MaintenanceEquipments, error) {
	mes := &MaintenanceEquipments{}
	if err := c.SearchRead(MaintenanceEquipmentModel, criteria, options, mes); err != nil {
		return nil, err
	}
	return mes, nil
}

// FindMaintenanceEquipmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceEquipmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MaintenanceEquipmentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMaintenanceEquipmentId finds record id by querying it with criteria.
func (c *Client) FindMaintenanceEquipmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MaintenanceEquipmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("maintenance.equipment was not found with criteria %v and options %v", criteria, options)
}
