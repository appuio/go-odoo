package odoo

import (
	"fmt"
)

// WizardCreateRoleFromUser represents wizard.create.role.from.user model.
type WizardCreateRoleFromUser struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	AssignToUser *Bool     `xmlrpc:"assign_to_user,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	Name         *String   `xmlrpc:"name,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WizardCreateRoleFromUsers represents array of wizard.create.role.from.user model.
type WizardCreateRoleFromUsers []WizardCreateRoleFromUser

// WizardCreateRoleFromUserModel is the odoo model name.
const WizardCreateRoleFromUserModel = "wizard.create.role.from.user"

// Many2One convert WizardCreateRoleFromUser to *Many2One.
func (wcrfu *WizardCreateRoleFromUser) Many2One() *Many2One {
	return NewMany2One(wcrfu.Id.Get(), "")
}

// CreateWizardCreateRoleFromUser creates a new wizard.create.role.from.user model and returns its id.
func (c *Client) CreateWizardCreateRoleFromUser(wcrfu *WizardCreateRoleFromUser) (int64, error) {
	ids, err := c.CreateWizardCreateRoleFromUsers([]*WizardCreateRoleFromUser{wcrfu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWizardCreateRoleFromUser creates a new wizard.create.role.from.user model and returns its id.
func (c *Client) CreateWizardCreateRoleFromUsers(wcrfus []*WizardCreateRoleFromUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range wcrfus {
		vv = append(vv, v)
	}
	return c.Create(WizardCreateRoleFromUserModel, vv)
}

// UpdateWizardCreateRoleFromUser updates an existing wizard.create.role.from.user record.
func (c *Client) UpdateWizardCreateRoleFromUser(wcrfu *WizardCreateRoleFromUser) error {
	return c.UpdateWizardCreateRoleFromUsers([]int64{wcrfu.Id.Get()}, wcrfu)
}

// UpdateWizardCreateRoleFromUsers updates existing wizard.create.role.from.user records.
// All records (represented by ids) will be updated by wcrfu values.
func (c *Client) UpdateWizardCreateRoleFromUsers(ids []int64, wcrfu *WizardCreateRoleFromUser) error {
	return c.Update(WizardCreateRoleFromUserModel, ids, wcrfu)
}

// DeleteWizardCreateRoleFromUser deletes an existing wizard.create.role.from.user record.
func (c *Client) DeleteWizardCreateRoleFromUser(id int64) error {
	return c.DeleteWizardCreateRoleFromUsers([]int64{id})
}

// DeleteWizardCreateRoleFromUsers deletes existing wizard.create.role.from.user records.
func (c *Client) DeleteWizardCreateRoleFromUsers(ids []int64) error {
	return c.Delete(WizardCreateRoleFromUserModel, ids)
}

// GetWizardCreateRoleFromUser gets wizard.create.role.from.user existing record.
func (c *Client) GetWizardCreateRoleFromUser(id int64) (*WizardCreateRoleFromUser, error) {
	wcrfus, err := c.GetWizardCreateRoleFromUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	if wcrfus != nil && len(*wcrfus) > 0 {
		return &((*wcrfus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of wizard.create.role.from.user not found", id)
}

// GetWizardCreateRoleFromUsers gets wizard.create.role.from.user existing records.
func (c *Client) GetWizardCreateRoleFromUsers(ids []int64) (*WizardCreateRoleFromUsers, error) {
	wcrfus := &WizardCreateRoleFromUsers{}
	if err := c.Read(WizardCreateRoleFromUserModel, ids, nil, wcrfus); err != nil {
		return nil, err
	}
	return wcrfus, nil
}

// FindWizardCreateRoleFromUser finds wizard.create.role.from.user record by querying it with criteria.
func (c *Client) FindWizardCreateRoleFromUser(criteria *Criteria) (*WizardCreateRoleFromUser, error) {
	wcrfus := &WizardCreateRoleFromUsers{}
	if err := c.SearchRead(WizardCreateRoleFromUserModel, criteria, NewOptions().Limit(1), wcrfus); err != nil {
		return nil, err
	}
	if wcrfus != nil && len(*wcrfus) > 0 {
		return &((*wcrfus)[0]), nil
	}
	return nil, fmt.Errorf("wizard.create.role.from.user was not found with criteria %v", criteria)
}

// FindWizardCreateRoleFromUsers finds wizard.create.role.from.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWizardCreateRoleFromUsers(criteria *Criteria, options *Options) (*WizardCreateRoleFromUsers, error) {
	wcrfus := &WizardCreateRoleFromUsers{}
	if err := c.SearchRead(WizardCreateRoleFromUserModel, criteria, options, wcrfus); err != nil {
		return nil, err
	}
	return wcrfus, nil
}

// FindWizardCreateRoleFromUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWizardCreateRoleFromUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WizardCreateRoleFromUserModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWizardCreateRoleFromUserId finds record id by querying it with criteria.
func (c *Client) FindWizardCreateRoleFromUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WizardCreateRoleFromUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("wizard.create.role.from.user was not found with criteria %v and options %v", criteria, options)
}
