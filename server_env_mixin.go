package odoo

import (
	"fmt"
)

// ServerEnvMixin represents server.env.mixin model.
type ServerEnvMixin struct {
	ServerEnvDefaults interface{} `xmlrpc:"server_env_defaults,omptempty"`
}

// ServerEnvMixins represents array of server.env.mixin model.
type ServerEnvMixins []ServerEnvMixin

// ServerEnvMixinModel is the odoo model name.
const ServerEnvMixinModel = "server.env.mixin"

// Many2One convert ServerEnvMixin to *Many2One.
func (sem *ServerEnvMixin) Many2One() *Many2One {
	return NewMany2One(sem.Id.Get(), "")
}

// CreateServerEnvMixin creates a new server.env.mixin model and returns its id.
func (c *Client) CreateServerEnvMixin(sem *ServerEnvMixin) (int64, error) {
	ids, err := c.CreateServerEnvMixins([]*ServerEnvMixin{sem})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateServerEnvMixin creates a new server.env.mixin model and returns its id.
func (c *Client) CreateServerEnvMixins(sems []*ServerEnvMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range sems {
		vv = append(vv, v)
	}
	return c.Create(ServerEnvMixinModel, vv)
}

// UpdateServerEnvMixin updates an existing server.env.mixin record.
func (c *Client) UpdateServerEnvMixin(sem *ServerEnvMixin) error {
	return c.UpdateServerEnvMixins([]int64{sem.Id.Get()}, sem)
}

// UpdateServerEnvMixins updates existing server.env.mixin records.
// All records (represented by ids) will be updated by sem values.
func (c *Client) UpdateServerEnvMixins(ids []int64, sem *ServerEnvMixin) error {
	return c.Update(ServerEnvMixinModel, ids, sem)
}

// DeleteServerEnvMixin deletes an existing server.env.mixin record.
func (c *Client) DeleteServerEnvMixin(id int64) error {
	return c.DeleteServerEnvMixins([]int64{id})
}

// DeleteServerEnvMixins deletes existing server.env.mixin records.
func (c *Client) DeleteServerEnvMixins(ids []int64) error {
	return c.Delete(ServerEnvMixinModel, ids)
}

// GetServerEnvMixin gets server.env.mixin existing record.
func (c *Client) GetServerEnvMixin(id int64) (*ServerEnvMixin, error) {
	sems, err := c.GetServerEnvMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if sems != nil && len(*sems) > 0 {
		return &((*sems)[0]), nil
	}
	return nil, fmt.Errorf("id %v of server.env.mixin not found", id)
}

// GetServerEnvMixins gets server.env.mixin existing records.
func (c *Client) GetServerEnvMixins(ids []int64) (*ServerEnvMixins, error) {
	sems := &ServerEnvMixins{}
	if err := c.Read(ServerEnvMixinModel, ids, nil, sems); err != nil {
		return nil, err
	}
	return sems, nil
}

// FindServerEnvMixin finds server.env.mixin record by querying it with criteria.
func (c *Client) FindServerEnvMixin(criteria *Criteria) (*ServerEnvMixin, error) {
	sems := &ServerEnvMixins{}
	if err := c.SearchRead(ServerEnvMixinModel, criteria, NewOptions().Limit(1), sems); err != nil {
		return nil, err
	}
	if sems != nil && len(*sems) > 0 {
		return &((*sems)[0]), nil
	}
	return nil, fmt.Errorf("server.env.mixin was not found with criteria %v", criteria)
}

// FindServerEnvMixins finds server.env.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerEnvMixins(criteria *Criteria, options *Options) (*ServerEnvMixins, error) {
	sems := &ServerEnvMixins{}
	if err := c.SearchRead(ServerEnvMixinModel, criteria, options, sems); err != nil {
		return nil, err
	}
	return sems, nil
}

// FindServerEnvMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerEnvMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ServerEnvMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindServerEnvMixinId finds record id by querying it with criteria.
func (c *Client) FindServerEnvMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ServerEnvMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("server.env.mixin was not found with criteria %v and options %v", criteria, options)
}
