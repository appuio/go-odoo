package odoo

import (
	"fmt"
)

// MukRestAccessRules represents muk_rest.access_rules model.
type MukRestAccessRules struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	Applied        *Bool      `xmlrpc:"applied,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	ExpressionIds  *Relation  `xmlrpc:"expression_ids,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	OauthId        *Many2One  `xmlrpc:"oauth_id,omptempty"`
	Route          *String    `xmlrpc:"route,omptempty"`
	RouteSelection *Selection `xmlrpc:"route_selection,omptempty"`
	Rule           *String    `xmlrpc:"rule,omptempty"`
	Sequence       *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MukRestAccessRuless represents array of muk_rest.access_rules model.
type MukRestAccessRuless []MukRestAccessRules

// MukRestAccessRulesModel is the odoo model name.
const MukRestAccessRulesModel = "muk_rest.access_rules"

// Many2One convert MukRestAccessRules to *Many2One.
func (ma *MukRestAccessRules) Many2One() *Many2One {
	return NewMany2One(ma.Id.Get(), "")
}

// CreateMukRestAccessRules creates a new muk_rest.access_rules model and returns its id.
func (c *Client) CreateMukRestAccessRules(ma *MukRestAccessRules) (int64, error) {
	ids, err := c.CreateMukRestAccessRuless([]*MukRestAccessRules{ma})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestAccessRules creates a new muk_rest.access_rules model and returns its id.
func (c *Client) CreateMukRestAccessRuless(mas []*MukRestAccessRules) ([]int64, error) {
	var vv []interface{}
	for _, v := range mas {
		vv = append(vv, v)
	}
	return c.Create(MukRestAccessRulesModel, vv)
}

// UpdateMukRestAccessRules updates an existing muk_rest.access_rules record.
func (c *Client) UpdateMukRestAccessRules(ma *MukRestAccessRules) error {
	return c.UpdateMukRestAccessRuless([]int64{ma.Id.Get()}, ma)
}

// UpdateMukRestAccessRuless updates existing muk_rest.access_rules records.
// All records (represented by ids) will be updated by ma values.
func (c *Client) UpdateMukRestAccessRuless(ids []int64, ma *MukRestAccessRules) error {
	return c.Update(MukRestAccessRulesModel, ids, ma)
}

// DeleteMukRestAccessRules deletes an existing muk_rest.access_rules record.
func (c *Client) DeleteMukRestAccessRules(id int64) error {
	return c.DeleteMukRestAccessRuless([]int64{id})
}

// DeleteMukRestAccessRuless deletes existing muk_rest.access_rules records.
func (c *Client) DeleteMukRestAccessRuless(ids []int64) error {
	return c.Delete(MukRestAccessRulesModel, ids)
}

// GetMukRestAccessRules gets muk_rest.access_rules existing record.
func (c *Client) GetMukRestAccessRules(id int64) (*MukRestAccessRules, error) {
	mas, err := c.GetMukRestAccessRuless([]int64{id})
	if err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.access_rules not found", id)
}

// GetMukRestAccessRuless gets muk_rest.access_rules existing records.
func (c *Client) GetMukRestAccessRuless(ids []int64) (*MukRestAccessRuless, error) {
	mas := &MukRestAccessRuless{}
	if err := c.Read(MukRestAccessRulesModel, ids, nil, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMukRestAccessRules finds muk_rest.access_rules record by querying it with criteria.
func (c *Client) FindMukRestAccessRules(criteria *Criteria) (*MukRestAccessRules, error) {
	mas := &MukRestAccessRuless{}
	if err := c.SearchRead(MukRestAccessRulesModel, criteria, NewOptions().Limit(1), mas); err != nil {
		return nil, err
	}
	if mas != nil && len(*mas) > 0 {
		return &((*mas)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.access_rules was not found with criteria %v", criteria)
}

// FindMukRestAccessRuless finds muk_rest.access_rules records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAccessRuless(criteria *Criteria, options *Options) (*MukRestAccessRuless, error) {
	mas := &MukRestAccessRuless{}
	if err := c.SearchRead(MukRestAccessRulesModel, criteria, options, mas); err != nil {
		return nil, err
	}
	return mas, nil
}

// FindMukRestAccessRulesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestAccessRulesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestAccessRulesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestAccessRulesId finds record id by querying it with criteria.
func (c *Client) FindMukRestAccessRulesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestAccessRulesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.access_rules was not found with criteria %v and options %v", criteria, options)
}
