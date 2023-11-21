package odoo

import (
	"fmt"
)

// WebsiteBaseUnit represents website.base.unit model.
type WebsiteBaseUnit struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// WebsiteBaseUnits represents array of website.base.unit model.
type WebsiteBaseUnits []WebsiteBaseUnit

// WebsiteBaseUnitModel is the odoo model name.
const WebsiteBaseUnitModel = "website.base.unit"

// Many2One convert WebsiteBaseUnit to *Many2One.
func (wbu *WebsiteBaseUnit) Many2One() *Many2One {
	return NewMany2One(wbu.Id.Get(), "")
}

// CreateWebsiteBaseUnit creates a new website.base.unit model and returns its id.
func (c *Client) CreateWebsiteBaseUnit(wbu *WebsiteBaseUnit) (int64, error) {
	ids, err := c.CreateWebsiteBaseUnits([]*WebsiteBaseUnit{wbu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteBaseUnit creates a new website.base.unit model and returns its id.
func (c *Client) CreateWebsiteBaseUnits(wbus []*WebsiteBaseUnit) ([]int64, error) {
	var vv []interface{}
	for _, v := range wbus {
		vv = append(vv, v)
	}
	return c.Create(WebsiteBaseUnitModel, vv)
}

// UpdateWebsiteBaseUnit updates an existing website.base.unit record.
func (c *Client) UpdateWebsiteBaseUnit(wbu *WebsiteBaseUnit) error {
	return c.UpdateWebsiteBaseUnits([]int64{wbu.Id.Get()}, wbu)
}

// UpdateWebsiteBaseUnits updates existing website.base.unit records.
// All records (represented by ids) will be updated by wbu values.
func (c *Client) UpdateWebsiteBaseUnits(ids []int64, wbu *WebsiteBaseUnit) error {
	return c.Update(WebsiteBaseUnitModel, ids, wbu)
}

// DeleteWebsiteBaseUnit deletes an existing website.base.unit record.
func (c *Client) DeleteWebsiteBaseUnit(id int64) error {
	return c.DeleteWebsiteBaseUnits([]int64{id})
}

// DeleteWebsiteBaseUnits deletes existing website.base.unit records.
func (c *Client) DeleteWebsiteBaseUnits(ids []int64) error {
	return c.Delete(WebsiteBaseUnitModel, ids)
}

// GetWebsiteBaseUnit gets website.base.unit existing record.
func (c *Client) GetWebsiteBaseUnit(id int64) (*WebsiteBaseUnit, error) {
	wbus, err := c.GetWebsiteBaseUnits([]int64{id})
	if err != nil {
		return nil, err
	}
	if wbus != nil && len(*wbus) > 0 {
		return &((*wbus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of website.base.unit not found", id)
}

// GetWebsiteBaseUnits gets website.base.unit existing records.
func (c *Client) GetWebsiteBaseUnits(ids []int64) (*WebsiteBaseUnits, error) {
	wbus := &WebsiteBaseUnits{}
	if err := c.Read(WebsiteBaseUnitModel, ids, nil, wbus); err != nil {
		return nil, err
	}
	return wbus, nil
}

// FindWebsiteBaseUnit finds website.base.unit record by querying it with criteria.
func (c *Client) FindWebsiteBaseUnit(criteria *Criteria) (*WebsiteBaseUnit, error) {
	wbus := &WebsiteBaseUnits{}
	if err := c.SearchRead(WebsiteBaseUnitModel, criteria, NewOptions().Limit(1), wbus); err != nil {
		return nil, err
	}
	if wbus != nil && len(*wbus) > 0 {
		return &((*wbus)[0]), nil
	}
	return nil, fmt.Errorf("website.base.unit was not found with criteria %v", criteria)
}

// FindWebsiteBaseUnits finds website.base.unit records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteBaseUnits(criteria *Criteria, options *Options) (*WebsiteBaseUnits, error) {
	wbus := &WebsiteBaseUnits{}
	if err := c.SearchRead(WebsiteBaseUnitModel, criteria, options, wbus); err != nil {
		return nil, err
	}
	return wbus, nil
}

// FindWebsiteBaseUnitIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteBaseUnitIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(WebsiteBaseUnitModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindWebsiteBaseUnitId finds record id by querying it with criteria.
func (c *Client) FindWebsiteBaseUnitId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteBaseUnitModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("website.base.unit was not found with criteria %v and options %v", criteria, options)
}
