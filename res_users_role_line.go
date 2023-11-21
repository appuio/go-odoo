package odoo

import (
	"fmt"
)

// ResUsersRoleLine represents res.users.role.line model.
type ResUsersRoleLine struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	Active            *Bool     `xmlrpc:"active,omptempty"`
	AllowedCompanyIds *Relation `xmlrpc:"allowed_company_ids,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DateFrom          *Time     `xmlrpc:"date_from,omptempty"`
	DateTo            *Time     `xmlrpc:"date_to,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	IsEnabled         *Bool     `xmlrpc:"is_enabled,omptempty"`
	RoleId            *Many2One `xmlrpc:"role_id,omptempty"`
	UserId            *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResUsersRoleLines represents array of res.users.role.line model.
type ResUsersRoleLines []ResUsersRoleLine

// ResUsersRoleLineModel is the odoo model name.
const ResUsersRoleLineModel = "res.users.role.line"

// Many2One convert ResUsersRoleLine to *Many2One.
func (rurl *ResUsersRoleLine) Many2One() *Many2One {
	return NewMany2One(rurl.Id.Get(), "")
}

// CreateResUsersRoleLine creates a new res.users.role.line model and returns its id.
func (c *Client) CreateResUsersRoleLine(rurl *ResUsersRoleLine) (int64, error) {
	ids, err := c.CreateResUsersRoleLines([]*ResUsersRoleLine{rurl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersRoleLine creates a new res.users.role.line model and returns its id.
func (c *Client) CreateResUsersRoleLines(rurls []*ResUsersRoleLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range rurls {
		vv = append(vv, v)
	}
	return c.Create(ResUsersRoleLineModel, vv)
}

// UpdateResUsersRoleLine updates an existing res.users.role.line record.
func (c *Client) UpdateResUsersRoleLine(rurl *ResUsersRoleLine) error {
	return c.UpdateResUsersRoleLines([]int64{rurl.Id.Get()}, rurl)
}

// UpdateResUsersRoleLines updates existing res.users.role.line records.
// All records (represented by ids) will be updated by rurl values.
func (c *Client) UpdateResUsersRoleLines(ids []int64, rurl *ResUsersRoleLine) error {
	return c.Update(ResUsersRoleLineModel, ids, rurl)
}

// DeleteResUsersRoleLine deletes an existing res.users.role.line record.
func (c *Client) DeleteResUsersRoleLine(id int64) error {
	return c.DeleteResUsersRoleLines([]int64{id})
}

// DeleteResUsersRoleLines deletes existing res.users.role.line records.
func (c *Client) DeleteResUsersRoleLines(ids []int64) error {
	return c.Delete(ResUsersRoleLineModel, ids)
}

// GetResUsersRoleLine gets res.users.role.line existing record.
func (c *Client) GetResUsersRoleLine(id int64) (*ResUsersRoleLine, error) {
	rurls, err := c.GetResUsersRoleLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if rurls != nil && len(*rurls) > 0 {
		return &((*rurls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users.role.line not found", id)
}

// GetResUsersRoleLines gets res.users.role.line existing records.
func (c *Client) GetResUsersRoleLines(ids []int64) (*ResUsersRoleLines, error) {
	rurls := &ResUsersRoleLines{}
	if err := c.Read(ResUsersRoleLineModel, ids, nil, rurls); err != nil {
		return nil, err
	}
	return rurls, nil
}

// FindResUsersRoleLine finds res.users.role.line record by querying it with criteria.
func (c *Client) FindResUsersRoleLine(criteria *Criteria) (*ResUsersRoleLine, error) {
	rurls := &ResUsersRoleLines{}
	if err := c.SearchRead(ResUsersRoleLineModel, criteria, NewOptions().Limit(1), rurls); err != nil {
		return nil, err
	}
	if rurls != nil && len(*rurls) > 0 {
		return &((*rurls)[0]), nil
	}
	return nil, fmt.Errorf("res.users.role.line was not found with criteria %v", criteria)
}

// FindResUsersRoleLines finds res.users.role.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersRoleLines(criteria *Criteria, options *Options) (*ResUsersRoleLines, error) {
	rurls := &ResUsersRoleLines{}
	if err := c.SearchRead(ResUsersRoleLineModel, criteria, options, rurls); err != nil {
		return nil, err
	}
	return rurls, nil
}

// FindResUsersRoleLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersRoleLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersRoleLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersRoleLineId finds record id by querying it with criteria.
func (c *Client) FindResUsersRoleLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersRoleLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users.role.line was not found with criteria %v and options %v", criteria, options)
}
