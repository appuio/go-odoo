package odoo

import (
	"fmt"
)

// ServerEnvTechnameMixin represents server.env.techname.mixin model.
type ServerEnvTechnameMixin struct {
	TechName *String `xmlrpc:"tech_name,omptempty"`
}

// ServerEnvTechnameMixins represents array of server.env.techname.mixin model.
type ServerEnvTechnameMixins []ServerEnvTechnameMixin

// ServerEnvTechnameMixinModel is the odoo model name.
const ServerEnvTechnameMixinModel = "server.env.techname.mixin"

// Many2One convert ServerEnvTechnameMixin to *Many2One.
func (setm *ServerEnvTechnameMixin) Many2One() *Many2One {
	return NewMany2One(setm.Id.Get(), "")
}

// CreateServerEnvTechnameMixin creates a new server.env.techname.mixin model and returns its id.
func (c *Client) CreateServerEnvTechnameMixin(setm *ServerEnvTechnameMixin) (int64, error) {
	ids, err := c.CreateServerEnvTechnameMixins([]*ServerEnvTechnameMixin{setm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateServerEnvTechnameMixin creates a new server.env.techname.mixin model and returns its id.
func (c *Client) CreateServerEnvTechnameMixins(setms []*ServerEnvTechnameMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range setms {
		vv = append(vv, v)
	}
	return c.Create(ServerEnvTechnameMixinModel, vv)
}

// UpdateServerEnvTechnameMixin updates an existing server.env.techname.mixin record.
func (c *Client) UpdateServerEnvTechnameMixin(setm *ServerEnvTechnameMixin) error {
	return c.UpdateServerEnvTechnameMixins([]int64{setm.Id.Get()}, setm)
}

// UpdateServerEnvTechnameMixins updates existing server.env.techname.mixin records.
// All records (represented by ids) will be updated by setm values.
func (c *Client) UpdateServerEnvTechnameMixins(ids []int64, setm *ServerEnvTechnameMixin) error {
	return c.Update(ServerEnvTechnameMixinModel, ids, setm)
}

// DeleteServerEnvTechnameMixin deletes an existing server.env.techname.mixin record.
func (c *Client) DeleteServerEnvTechnameMixin(id int64) error {
	return c.DeleteServerEnvTechnameMixins([]int64{id})
}

// DeleteServerEnvTechnameMixins deletes existing server.env.techname.mixin records.
func (c *Client) DeleteServerEnvTechnameMixins(ids []int64) error {
	return c.Delete(ServerEnvTechnameMixinModel, ids)
}

// GetServerEnvTechnameMixin gets server.env.techname.mixin existing record.
func (c *Client) GetServerEnvTechnameMixin(id int64) (*ServerEnvTechnameMixin, error) {
	setms, err := c.GetServerEnvTechnameMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if setms != nil && len(*setms) > 0 {
		return &((*setms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of server.env.techname.mixin not found", id)
}

// GetServerEnvTechnameMixins gets server.env.techname.mixin existing records.
func (c *Client) GetServerEnvTechnameMixins(ids []int64) (*ServerEnvTechnameMixins, error) {
	setms := &ServerEnvTechnameMixins{}
	if err := c.Read(ServerEnvTechnameMixinModel, ids, nil, setms); err != nil {
		return nil, err
	}
	return setms, nil
}

// FindServerEnvTechnameMixin finds server.env.techname.mixin record by querying it with criteria.
func (c *Client) FindServerEnvTechnameMixin(criteria *Criteria) (*ServerEnvTechnameMixin, error) {
	setms := &ServerEnvTechnameMixins{}
	if err := c.SearchRead(ServerEnvTechnameMixinModel, criteria, NewOptions().Limit(1), setms); err != nil {
		return nil, err
	}
	if setms != nil && len(*setms) > 0 {
		return &((*setms)[0]), nil
	}
	return nil, fmt.Errorf("server.env.techname.mixin was not found with criteria %v", criteria)
}

// FindServerEnvTechnameMixins finds server.env.techname.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerEnvTechnameMixins(criteria *Criteria, options *Options) (*ServerEnvTechnameMixins, error) {
	setms := &ServerEnvTechnameMixins{}
	if err := c.SearchRead(ServerEnvTechnameMixinModel, criteria, options, setms); err != nil {
		return nil, err
	}
	return setms, nil
}

// FindServerEnvTechnameMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerEnvTechnameMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ServerEnvTechnameMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindServerEnvTechnameMixinId finds record id by querying it with criteria.
func (c *Client) FindServerEnvTechnameMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ServerEnvTechnameMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("server.env.techname.mixin was not found with criteria %v and options %v", criteria, options)
}
