package odoo

import (
	"fmt"
)

// MaintenanceEquipmentCategory represents maintenance.equipment.category model.
type MaintenanceEquipmentCategory struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AliasBouncedContent      *String    `xmlrpc:"alias_bounced_content,omptempty"`
	AliasContact             *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults            *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain              *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId       *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId                  *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId             *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName                *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId       *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId      *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId              *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	CheckSpam                *Bool      `xmlrpc:"check_spam,omptempty"`
	Color                    *Int       `xmlrpc:"color,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	EquipmentCount           *Int       `xmlrpc:"equipment_count,omptempty"`
	EquipmentIds             *Relation  `xmlrpc:"equipment_ids,omptempty"`
	FailedMessageIds         *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	Fold                     *Bool      `xmlrpc:"fold,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	MaintenanceCount         *Int       `xmlrpc:"maintenance_count,omptempty"`
	MaintenanceIds           *Relation  `xmlrpc:"maintenance_ids,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent           *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	Note                     *String    `xmlrpc:"note,omptempty"`
	TechnicianUserId         *Many2One  `xmlrpc:"technician_user_id,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MaintenanceEquipmentCategorys represents array of maintenance.equipment.category model.
type MaintenanceEquipmentCategorys []MaintenanceEquipmentCategory

// MaintenanceEquipmentCategoryModel is the odoo model name.
const MaintenanceEquipmentCategoryModel = "maintenance.equipment.category"

// Many2One convert MaintenanceEquipmentCategory to *Many2One.
func (mec *MaintenanceEquipmentCategory) Many2One() *Many2One {
	return NewMany2One(mec.Id.Get(), "")
}

// CreateMaintenanceEquipmentCategory creates a new maintenance.equipment.category model and returns its id.
func (c *Client) CreateMaintenanceEquipmentCategory(mec *MaintenanceEquipmentCategory) (int64, error) {
	ids, err := c.CreateMaintenanceEquipmentCategorys([]*MaintenanceEquipmentCategory{mec})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMaintenanceEquipmentCategory creates a new maintenance.equipment.category model and returns its id.
func (c *Client) CreateMaintenanceEquipmentCategorys(mecs []*MaintenanceEquipmentCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range mecs {
		vv = append(vv, v)
	}
	return c.Create(MaintenanceEquipmentCategoryModel, vv)
}

// UpdateMaintenanceEquipmentCategory updates an existing maintenance.equipment.category record.
func (c *Client) UpdateMaintenanceEquipmentCategory(mec *MaintenanceEquipmentCategory) error {
	return c.UpdateMaintenanceEquipmentCategorys([]int64{mec.Id.Get()}, mec)
}

// UpdateMaintenanceEquipmentCategorys updates existing maintenance.equipment.category records.
// All records (represented by ids) will be updated by mec values.
func (c *Client) UpdateMaintenanceEquipmentCategorys(ids []int64, mec *MaintenanceEquipmentCategory) error {
	return c.Update(MaintenanceEquipmentCategoryModel, ids, mec)
}

// DeleteMaintenanceEquipmentCategory deletes an existing maintenance.equipment.category record.
func (c *Client) DeleteMaintenanceEquipmentCategory(id int64) error {
	return c.DeleteMaintenanceEquipmentCategorys([]int64{id})
}

// DeleteMaintenanceEquipmentCategorys deletes existing maintenance.equipment.category records.
func (c *Client) DeleteMaintenanceEquipmentCategorys(ids []int64) error {
	return c.Delete(MaintenanceEquipmentCategoryModel, ids)
}

// GetMaintenanceEquipmentCategory gets maintenance.equipment.category existing record.
func (c *Client) GetMaintenanceEquipmentCategory(id int64) (*MaintenanceEquipmentCategory, error) {
	mecs, err := c.GetMaintenanceEquipmentCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if mecs != nil && len(*mecs) > 0 {
		return &((*mecs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of maintenance.equipment.category not found", id)
}

// GetMaintenanceEquipmentCategorys gets maintenance.equipment.category existing records.
func (c *Client) GetMaintenanceEquipmentCategorys(ids []int64) (*MaintenanceEquipmentCategorys, error) {
	mecs := &MaintenanceEquipmentCategorys{}
	if err := c.Read(MaintenanceEquipmentCategoryModel, ids, nil, mecs); err != nil {
		return nil, err
	}
	return mecs, nil
}

// FindMaintenanceEquipmentCategory finds maintenance.equipment.category record by querying it with criteria.
func (c *Client) FindMaintenanceEquipmentCategory(criteria *Criteria) (*MaintenanceEquipmentCategory, error) {
	mecs := &MaintenanceEquipmentCategorys{}
	if err := c.SearchRead(MaintenanceEquipmentCategoryModel, criteria, NewOptions().Limit(1), mecs); err != nil {
		return nil, err
	}
	if mecs != nil && len(*mecs) > 0 {
		return &((*mecs)[0]), nil
	}
	return nil, fmt.Errorf("maintenance.equipment.category was not found with criteria %v", criteria)
}

// FindMaintenanceEquipmentCategorys finds maintenance.equipment.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceEquipmentCategorys(criteria *Criteria, options *Options) (*MaintenanceEquipmentCategorys, error) {
	mecs := &MaintenanceEquipmentCategorys{}
	if err := c.SearchRead(MaintenanceEquipmentCategoryModel, criteria, options, mecs); err != nil {
		return nil, err
	}
	return mecs, nil
}

// FindMaintenanceEquipmentCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceEquipmentCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MaintenanceEquipmentCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMaintenanceEquipmentCategoryId finds record id by querying it with criteria.
func (c *Client) FindMaintenanceEquipmentCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MaintenanceEquipmentCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("maintenance.equipment.category was not found with criteria %v and options %v", criteria, options)
}
