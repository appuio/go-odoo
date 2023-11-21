package odoo

import (
	"fmt"
)

// ResUsersRole represents res.users.role model.
type ResUsersRole struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CategoryId       *Many2One `xmlrpc:"category_id,omptempty"`
	Color            *Int      `xmlrpc:"color,omptempty"`
	Comment          *String   `xmlrpc:"comment,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	FullName         *String   `xmlrpc:"full_name,omptempty"`
	GroupCategoryId  *Many2One `xmlrpc:"group_category_id,omptempty"`
	GroupId          *Many2One `xmlrpc:"group_id,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	ImpliedIds       *Relation `xmlrpc:"implied_ids,omptempty"`
	LineIds          *Relation `xmlrpc:"line_ids,omptempty"`
	MenuAccess       *Relation `xmlrpc:"menu_access,omptempty"`
	ModelAccess      *Relation `xmlrpc:"model_access,omptempty"`
	ModelAccessCount *Int      `xmlrpc:"model_access_count,omptempty"`
	ModelAccessIds   *Relation `xmlrpc:"model_access_ids,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	ParentIds        *Relation `xmlrpc:"parent_ids,omptempty"`
	RoleCount        *Int      `xmlrpc:"role_count,omptempty"`
	RoleId           *Relation `xmlrpc:"role_id,omptempty"`
	RoleIds          *Relation `xmlrpc:"role_ids,omptempty"`
	RuleGroups       *Relation `xmlrpc:"rule_groups,omptempty"`
	RuleIds          *Relation `xmlrpc:"rule_ids,omptempty"`
	RulesCount       *Int      `xmlrpc:"rules_count,omptempty"`
	Share            *Bool     `xmlrpc:"share,omptempty"`
	TransImpliedIds  *Relation `xmlrpc:"trans_implied_ids,omptempty"`
	TransParentIds   *Relation `xmlrpc:"trans_parent_ids,omptempty"`
	UserIds          *Relation `xmlrpc:"user_ids,omptempty"`
	Users            *Relation `xmlrpc:"users,omptempty"`
	ViewAccess       *Relation `xmlrpc:"view_access,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResUsersRoles represents array of res.users.role model.
type ResUsersRoles []ResUsersRole

// ResUsersRoleModel is the odoo model name.
const ResUsersRoleModel = "res.users.role"

// Many2One convert ResUsersRole to *Many2One.
func (rur *ResUsersRole) Many2One() *Many2One {
	return NewMany2One(rur.Id.Get(), "")
}

// CreateResUsersRole creates a new res.users.role model and returns its id.
func (c *Client) CreateResUsersRole(rur *ResUsersRole) (int64, error) {
	ids, err := c.CreateResUsersRoles([]*ResUsersRole{rur})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersRole creates a new res.users.role model and returns its id.
func (c *Client) CreateResUsersRoles(rurs []*ResUsersRole) ([]int64, error) {
	var vv []interface{}
	for _, v := range rurs {
		vv = append(vv, v)
	}
	return c.Create(ResUsersRoleModel, vv)
}

// UpdateResUsersRole updates an existing res.users.role record.
func (c *Client) UpdateResUsersRole(rur *ResUsersRole) error {
	return c.UpdateResUsersRoles([]int64{rur.Id.Get()}, rur)
}

// UpdateResUsersRoles updates existing res.users.role records.
// All records (represented by ids) will be updated by rur values.
func (c *Client) UpdateResUsersRoles(ids []int64, rur *ResUsersRole) error {
	return c.Update(ResUsersRoleModel, ids, rur)
}

// DeleteResUsersRole deletes an existing res.users.role record.
func (c *Client) DeleteResUsersRole(id int64) error {
	return c.DeleteResUsersRoles([]int64{id})
}

// DeleteResUsersRoles deletes existing res.users.role records.
func (c *Client) DeleteResUsersRoles(ids []int64) error {
	return c.Delete(ResUsersRoleModel, ids)
}

// GetResUsersRole gets res.users.role existing record.
func (c *Client) GetResUsersRole(id int64) (*ResUsersRole, error) {
	rurs, err := c.GetResUsersRoles([]int64{id})
	if err != nil {
		return nil, err
	}
	if rurs != nil && len(*rurs) > 0 {
		return &((*rurs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users.role not found", id)
}

// GetResUsersRoles gets res.users.role existing records.
func (c *Client) GetResUsersRoles(ids []int64) (*ResUsersRoles, error) {
	rurs := &ResUsersRoles{}
	if err := c.Read(ResUsersRoleModel, ids, nil, rurs); err != nil {
		return nil, err
	}
	return rurs, nil
}

// FindResUsersRole finds res.users.role record by querying it with criteria.
func (c *Client) FindResUsersRole(criteria *Criteria) (*ResUsersRole, error) {
	rurs := &ResUsersRoles{}
	if err := c.SearchRead(ResUsersRoleModel, criteria, NewOptions().Limit(1), rurs); err != nil {
		return nil, err
	}
	if rurs != nil && len(*rurs) > 0 {
		return &((*rurs)[0]), nil
	}
	return nil, fmt.Errorf("res.users.role was not found with criteria %v", criteria)
}

// FindResUsersRoles finds res.users.role records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersRoles(criteria *Criteria, options *Options) (*ResUsersRoles, error) {
	rurs := &ResUsersRoles{}
	if err := c.SearchRead(ResUsersRoleModel, criteria, options, rurs); err != nil {
		return nil, err
	}
	return rurs, nil
}

// FindResUsersRoleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersRoleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersRoleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersRoleId finds record id by querying it with criteria.
func (c *Client) FindResUsersRoleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersRoleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users.role was not found with criteria %v and options %v", criteria, options)
}
