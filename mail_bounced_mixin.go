package odoo

import (
	"fmt"
)

// MailBouncedMixin represents mail.bounced.mixin model.
type MailBouncedMixin struct {
	EmailBounced *Bool `xmlrpc:"email_bounced,omptempty"`
}

// MailBouncedMixins represents array of mail.bounced.mixin model.
type MailBouncedMixins []MailBouncedMixin

// MailBouncedMixinModel is the odoo model name.
const MailBouncedMixinModel = "mail.bounced.mixin"

// Many2One convert MailBouncedMixin to *Many2One.
func (mbm *MailBouncedMixin) Many2One() *Many2One {
	return NewMany2One(mbm.Id.Get(), "")
}

// CreateMailBouncedMixin creates a new mail.bounced.mixin model and returns its id.
func (c *Client) CreateMailBouncedMixin(mbm *MailBouncedMixin) (int64, error) {
	ids, err := c.CreateMailBouncedMixins([]*MailBouncedMixin{mbm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailBouncedMixin creates a new mail.bounced.mixin model and returns its id.
func (c *Client) CreateMailBouncedMixins(mbms []*MailBouncedMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range mbms {
		vv = append(vv, v)
	}
	return c.Create(MailBouncedMixinModel, vv)
}

// UpdateMailBouncedMixin updates an existing mail.bounced.mixin record.
func (c *Client) UpdateMailBouncedMixin(mbm *MailBouncedMixin) error {
	return c.UpdateMailBouncedMixins([]int64{mbm.Id.Get()}, mbm)
}

// UpdateMailBouncedMixins updates existing mail.bounced.mixin records.
// All records (represented by ids) will be updated by mbm values.
func (c *Client) UpdateMailBouncedMixins(ids []int64, mbm *MailBouncedMixin) error {
	return c.Update(MailBouncedMixinModel, ids, mbm)
}

// DeleteMailBouncedMixin deletes an existing mail.bounced.mixin record.
func (c *Client) DeleteMailBouncedMixin(id int64) error {
	return c.DeleteMailBouncedMixins([]int64{id})
}

// DeleteMailBouncedMixins deletes existing mail.bounced.mixin records.
func (c *Client) DeleteMailBouncedMixins(ids []int64) error {
	return c.Delete(MailBouncedMixinModel, ids)
}

// GetMailBouncedMixin gets mail.bounced.mixin existing record.
func (c *Client) GetMailBouncedMixin(id int64) (*MailBouncedMixin, error) {
	mbms, err := c.GetMailBouncedMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if mbms != nil && len(*mbms) > 0 {
		return &((*mbms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.bounced.mixin not found", id)
}

// GetMailBouncedMixins gets mail.bounced.mixin existing records.
func (c *Client) GetMailBouncedMixins(ids []int64) (*MailBouncedMixins, error) {
	mbms := &MailBouncedMixins{}
	if err := c.Read(MailBouncedMixinModel, ids, nil, mbms); err != nil {
		return nil, err
	}
	return mbms, nil
}

// FindMailBouncedMixin finds mail.bounced.mixin record by querying it with criteria.
func (c *Client) FindMailBouncedMixin(criteria *Criteria) (*MailBouncedMixin, error) {
	mbms := &MailBouncedMixins{}
	if err := c.SearchRead(MailBouncedMixinModel, criteria, NewOptions().Limit(1), mbms); err != nil {
		return nil, err
	}
	if mbms != nil && len(*mbms) > 0 {
		return &((*mbms)[0]), nil
	}
	return nil, fmt.Errorf("mail.bounced.mixin was not found with criteria %v", criteria)
}

// FindMailBouncedMixins finds mail.bounced.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBouncedMixins(criteria *Criteria, options *Options) (*MailBouncedMixins, error) {
	mbms := &MailBouncedMixins{}
	if err := c.SearchRead(MailBouncedMixinModel, criteria, options, mbms); err != nil {
		return nil, err
	}
	return mbms, nil
}

// FindMailBouncedMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailBouncedMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailBouncedMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailBouncedMixinId finds record id by querying it with criteria.
func (c *Client) FindMailBouncedMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailBouncedMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.bounced.mixin was not found with criteria %v and options %v", criteria, options)
}
