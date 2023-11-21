package odoo

import (
	"fmt"
)

// MukRestEndpoint represents muk_rest.endpoint model.
type MukRestEndpoint struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	ActionId               *Many2One  `xmlrpc:"action_id,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	Code                   *String    `xmlrpc:"code,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	DocsComponents         *String    `xmlrpc:"docs_components,omptempty"`
	DocsDefaultResponse200 *Bool      `xmlrpc:"docs_default_response_200,omptempty"`
	DocsDefaultResponse400 *Bool      `xmlrpc:"docs_default_response_400,omptempty"`
	DocsDefaultResponse401 *Bool      `xmlrpc:"docs_default_response_401,omptempty"`
	DocsDefaultResponse500 *Bool      `xmlrpc:"docs_default_response_500,omptempty"`
	DocsDescription        *String    `xmlrpc:"docs_description,omptempty"`
	DocsParameters         *String    `xmlrpc:"docs_parameters,omptempty"`
	DocsResponses          *String    `xmlrpc:"docs_responses,omptempty"`
	DocsSummary            *String    `xmlrpc:"docs_summary,omptempty"`
	Domain                 *String    `xmlrpc:"domain,omptempty"`
	DomainFieldIds         *Relation  `xmlrpc:"domain_field_ids,omptempty"`
	Endpoint               *String    `xmlrpc:"endpoint,omptempty"`
	EvalSudo               *Bool      `xmlrpc:"eval_sudo,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	Logging                *Bool      `xmlrpc:"logging,omptempty"`
	Method                 *Selection `xmlrpc:"method,omptempty"`
	ModelId                *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName              *String    `xmlrpc:"model_name,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	Protected              *Bool      `xmlrpc:"protected,omptempty"`
	Route                  *String    `xmlrpc:"route,omptempty"`
	ShowLogging            *Bool      `xmlrpc:"show_logging,omptempty"`
	State                  *Selection `xmlrpc:"state,omptempty"`
	Url                    *String    `xmlrpc:"url,omptempty"`
	WrapResponse           *Bool      `xmlrpc:"wrap_response,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MukRestEndpoints represents array of muk_rest.endpoint model.
type MukRestEndpoints []MukRestEndpoint

// MukRestEndpointModel is the odoo model name.
const MukRestEndpointModel = "muk_rest.endpoint"

// Many2One convert MukRestEndpoint to *Many2One.
func (me *MukRestEndpoint) Many2One() *Many2One {
	return NewMany2One(me.Id.Get(), "")
}

// CreateMukRestEndpoint creates a new muk_rest.endpoint model and returns its id.
func (c *Client) CreateMukRestEndpoint(me *MukRestEndpoint) (int64, error) {
	ids, err := c.CreateMukRestEndpoints([]*MukRestEndpoint{me})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestEndpoint creates a new muk_rest.endpoint model and returns its id.
func (c *Client) CreateMukRestEndpoints(mes []*MukRestEndpoint) ([]int64, error) {
	var vv []interface{}
	for _, v := range mes {
		vv = append(vv, v)
	}
	return c.Create(MukRestEndpointModel, vv)
}

// UpdateMukRestEndpoint updates an existing muk_rest.endpoint record.
func (c *Client) UpdateMukRestEndpoint(me *MukRestEndpoint) error {
	return c.UpdateMukRestEndpoints([]int64{me.Id.Get()}, me)
}

// UpdateMukRestEndpoints updates existing muk_rest.endpoint records.
// All records (represented by ids) will be updated by me values.
func (c *Client) UpdateMukRestEndpoints(ids []int64, me *MukRestEndpoint) error {
	return c.Update(MukRestEndpointModel, ids, me)
}

// DeleteMukRestEndpoint deletes an existing muk_rest.endpoint record.
func (c *Client) DeleteMukRestEndpoint(id int64) error {
	return c.DeleteMukRestEndpoints([]int64{id})
}

// DeleteMukRestEndpoints deletes existing muk_rest.endpoint records.
func (c *Client) DeleteMukRestEndpoints(ids []int64) error {
	return c.Delete(MukRestEndpointModel, ids)
}

// GetMukRestEndpoint gets muk_rest.endpoint existing record.
func (c *Client) GetMukRestEndpoint(id int64) (*MukRestEndpoint, error) {
	mes, err := c.GetMukRestEndpoints([]int64{id})
	if err != nil {
		return nil, err
	}
	if mes != nil && len(*mes) > 0 {
		return &((*mes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.endpoint not found", id)
}

// GetMukRestEndpoints gets muk_rest.endpoint existing records.
func (c *Client) GetMukRestEndpoints(ids []int64) (*MukRestEndpoints, error) {
	mes := &MukRestEndpoints{}
	if err := c.Read(MukRestEndpointModel, ids, nil, mes); err != nil {
		return nil, err
	}
	return mes, nil
}

// FindMukRestEndpoint finds muk_rest.endpoint record by querying it with criteria.
func (c *Client) FindMukRestEndpoint(criteria *Criteria) (*MukRestEndpoint, error) {
	mes := &MukRestEndpoints{}
	if err := c.SearchRead(MukRestEndpointModel, criteria, NewOptions().Limit(1), mes); err != nil {
		return nil, err
	}
	if mes != nil && len(*mes) > 0 {
		return &((*mes)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.endpoint was not found with criteria %v", criteria)
}

// FindMukRestEndpoints finds muk_rest.endpoint records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestEndpoints(criteria *Criteria, options *Options) (*MukRestEndpoints, error) {
	mes := &MukRestEndpoints{}
	if err := c.SearchRead(MukRestEndpointModel, criteria, options, mes); err != nil {
		return nil, err
	}
	return mes, nil
}

// FindMukRestEndpointIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestEndpointIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestEndpointModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestEndpointId finds record id by querying it with criteria.
func (c *Client) FindMukRestEndpointId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestEndpointModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.endpoint was not found with criteria %v and options %v", criteria, options)
}
