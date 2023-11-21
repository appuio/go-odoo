package odoo

import (
	"fmt"
)

// MukRestRequestData represents muk_rest.request_data model.
type MukRestRequestData struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	ClientKey   *String   `xmlrpc:"client_key,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Nonce       *String   `xmlrpc:"nonce,omptempty"`
	Timestamp   *String   `xmlrpc:"timestamp,omptempty"`
	TokenHash   *String   `xmlrpc:"token_hash,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MukRestRequestDatas represents array of muk_rest.request_data model.
type MukRestRequestDatas []MukRestRequestData

// MukRestRequestDataModel is the odoo model name.
const MukRestRequestDataModel = "muk_rest.request_data"

// Many2One convert MukRestRequestData to *Many2One.
func (mr *MukRestRequestData) Many2One() *Many2One {
	return NewMany2One(mr.Id.Get(), "")
}

// CreateMukRestRequestData creates a new muk_rest.request_data model and returns its id.
func (c *Client) CreateMukRestRequestData(mr *MukRestRequestData) (int64, error) {
	ids, err := c.CreateMukRestRequestDatas([]*MukRestRequestData{mr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMukRestRequestData creates a new muk_rest.request_data model and returns its id.
func (c *Client) CreateMukRestRequestDatas(mrs []*MukRestRequestData) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrs {
		vv = append(vv, v)
	}
	return c.Create(MukRestRequestDataModel, vv)
}

// UpdateMukRestRequestData updates an existing muk_rest.request_data record.
func (c *Client) UpdateMukRestRequestData(mr *MukRestRequestData) error {
	return c.UpdateMukRestRequestDatas([]int64{mr.Id.Get()}, mr)
}

// UpdateMukRestRequestDatas updates existing muk_rest.request_data records.
// All records (represented by ids) will be updated by mr values.
func (c *Client) UpdateMukRestRequestDatas(ids []int64, mr *MukRestRequestData) error {
	return c.Update(MukRestRequestDataModel, ids, mr)
}

// DeleteMukRestRequestData deletes an existing muk_rest.request_data record.
func (c *Client) DeleteMukRestRequestData(id int64) error {
	return c.DeleteMukRestRequestDatas([]int64{id})
}

// DeleteMukRestRequestDatas deletes existing muk_rest.request_data records.
func (c *Client) DeleteMukRestRequestDatas(ids []int64) error {
	return c.Delete(MukRestRequestDataModel, ids)
}

// GetMukRestRequestData gets muk_rest.request_data existing record.
func (c *Client) GetMukRestRequestData(id int64) (*MukRestRequestData, error) {
	mrs, err := c.GetMukRestRequestDatas([]int64{id})
	if err != nil {
		return nil, err
	}
	if mrs != nil && len(*mrs) > 0 {
		return &((*mrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of muk_rest.request_data not found", id)
}

// GetMukRestRequestDatas gets muk_rest.request_data existing records.
func (c *Client) GetMukRestRequestDatas(ids []int64) (*MukRestRequestDatas, error) {
	mrs := &MukRestRequestDatas{}
	if err := c.Read(MukRestRequestDataModel, ids, nil, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMukRestRequestData finds muk_rest.request_data record by querying it with criteria.
func (c *Client) FindMukRestRequestData(criteria *Criteria) (*MukRestRequestData, error) {
	mrs := &MukRestRequestDatas{}
	if err := c.SearchRead(MukRestRequestDataModel, criteria, NewOptions().Limit(1), mrs); err != nil {
		return nil, err
	}
	if mrs != nil && len(*mrs) > 0 {
		return &((*mrs)[0]), nil
	}
	return nil, fmt.Errorf("muk_rest.request_data was not found with criteria %v", criteria)
}

// FindMukRestRequestDatas finds muk_rest.request_data records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestRequestDatas(criteria *Criteria, options *Options) (*MukRestRequestDatas, error) {
	mrs := &MukRestRequestDatas{}
	if err := c.SearchRead(MukRestRequestDataModel, criteria, options, mrs); err != nil {
		return nil, err
	}
	return mrs, nil
}

// FindMukRestRequestDataIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMukRestRequestDataIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MukRestRequestDataModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMukRestRequestDataId finds record id by querying it with criteria.
func (c *Client) FindMukRestRequestDataId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MukRestRequestDataModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("muk_rest.request_data was not found with criteria %v and options %v", criteria, options)
}
