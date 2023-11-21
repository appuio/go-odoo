package odoo

import (
	"fmt"
)

// ResPartnerIap represents res.partner.iap model.
type ResPartnerIap struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	IapEnrichInfo   *String   `xmlrpc:"iap_enrich_info,omptempty"`
	IapSearchDomain *String   `xmlrpc:"iap_search_domain,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	PartnerId       *Many2One `xmlrpc:"partner_id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResPartnerIaps represents array of res.partner.iap model.
type ResPartnerIaps []ResPartnerIap

// ResPartnerIapModel is the odoo model name.
const ResPartnerIapModel = "res.partner.iap"

// Many2One convert ResPartnerIap to *Many2One.
func (rpi *ResPartnerIap) Many2One() *Many2One {
	return NewMany2One(rpi.Id.Get(), "")
}

// CreateResPartnerIap creates a new res.partner.iap model and returns its id.
func (c *Client) CreateResPartnerIap(rpi *ResPartnerIap) (int64, error) {
	ids, err := c.CreateResPartnerIaps([]*ResPartnerIap{rpi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResPartnerIap creates a new res.partner.iap model and returns its id.
func (c *Client) CreateResPartnerIaps(rpis []*ResPartnerIap) ([]int64, error) {
	var vv []interface{}
	for _, v := range rpis {
		vv = append(vv, v)
	}
	return c.Create(ResPartnerIapModel, vv)
}

// UpdateResPartnerIap updates an existing res.partner.iap record.
func (c *Client) UpdateResPartnerIap(rpi *ResPartnerIap) error {
	return c.UpdateResPartnerIaps([]int64{rpi.Id.Get()}, rpi)
}

// UpdateResPartnerIaps updates existing res.partner.iap records.
// All records (represented by ids) will be updated by rpi values.
func (c *Client) UpdateResPartnerIaps(ids []int64, rpi *ResPartnerIap) error {
	return c.Update(ResPartnerIapModel, ids, rpi)
}

// DeleteResPartnerIap deletes an existing res.partner.iap record.
func (c *Client) DeleteResPartnerIap(id int64) error {
	return c.DeleteResPartnerIaps([]int64{id})
}

// DeleteResPartnerIaps deletes existing res.partner.iap records.
func (c *Client) DeleteResPartnerIaps(ids []int64) error {
	return c.Delete(ResPartnerIapModel, ids)
}

// GetResPartnerIap gets res.partner.iap existing record.
func (c *Client) GetResPartnerIap(id int64) (*ResPartnerIap, error) {
	rpis, err := c.GetResPartnerIaps([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpis != nil && len(*rpis) > 0 {
		return &((*rpis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner.iap not found", id)
}

// GetResPartnerIaps gets res.partner.iap existing records.
func (c *Client) GetResPartnerIaps(ids []int64) (*ResPartnerIaps, error) {
	rpis := &ResPartnerIaps{}
	if err := c.Read(ResPartnerIapModel, ids, nil, rpis); err != nil {
		return nil, err
	}
	return rpis, nil
}

// FindResPartnerIap finds res.partner.iap record by querying it with criteria.
func (c *Client) FindResPartnerIap(criteria *Criteria) (*ResPartnerIap, error) {
	rpis := &ResPartnerIaps{}
	if err := c.SearchRead(ResPartnerIapModel, criteria, NewOptions().Limit(1), rpis); err != nil {
		return nil, err
	}
	if rpis != nil && len(*rpis) > 0 {
		return &((*rpis)[0]), nil
	}
	return nil, fmt.Errorf("res.partner.iap was not found with criteria %v", criteria)
}

// FindResPartnerIaps finds res.partner.iap records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIaps(criteria *Criteria, options *Options) (*ResPartnerIaps, error) {
	rpis := &ResPartnerIaps{}
	if err := c.SearchRead(ResPartnerIapModel, criteria, options, rpis); err != nil {
		return nil, err
	}
	return rpis, nil
}

// FindResPartnerIapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIapIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerIapModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerIapId finds record id by querying it with criteria.
func (c *Client) FindResPartnerIapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerIapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner.iap was not found with criteria %v and options %v", criteria, options)
}
