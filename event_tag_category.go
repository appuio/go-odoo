package odoo

import (
	"fmt"
)

// EventTagCategory represents event.tag.category model.
type EventTagCategory struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CanPublish       *Bool     `xmlrpc:"can_publish,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	IsPublished      *Bool     `xmlrpc:"is_published,omptempty"`
	Name             *String   `xmlrpc:"name,omptempty"`
	Sequence         *Int      `xmlrpc:"sequence,omptempty"`
	TagIds           *Relation `xmlrpc:"tag_ids,omptempty"`
	WebsitePublished *Bool     `xmlrpc:"website_published,omptempty"`
	WebsiteUrl       *String   `xmlrpc:"website_url,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// EventTagCategorys represents array of event.tag.category model.
type EventTagCategorys []EventTagCategory

// EventTagCategoryModel is the odoo model name.
const EventTagCategoryModel = "event.tag.category"

// Many2One convert EventTagCategory to *Many2One.
func (etc *EventTagCategory) Many2One() *Many2One {
	return NewMany2One(etc.Id.Get(), "")
}

// CreateEventTagCategory creates a new event.tag.category model and returns its id.
func (c *Client) CreateEventTagCategory(etc *EventTagCategory) (int64, error) {
	ids, err := c.CreateEventTagCategorys([]*EventTagCategory{etc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventTagCategory creates a new event.tag.category model and returns its id.
func (c *Client) CreateEventTagCategorys(etcs []*EventTagCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range etcs {
		vv = append(vv, v)
	}
	return c.Create(EventTagCategoryModel, vv)
}

// UpdateEventTagCategory updates an existing event.tag.category record.
func (c *Client) UpdateEventTagCategory(etc *EventTagCategory) error {
	return c.UpdateEventTagCategorys([]int64{etc.Id.Get()}, etc)
}

// UpdateEventTagCategorys updates existing event.tag.category records.
// All records (represented by ids) will be updated by etc values.
func (c *Client) UpdateEventTagCategorys(ids []int64, etc *EventTagCategory) error {
	return c.Update(EventTagCategoryModel, ids, etc)
}

// DeleteEventTagCategory deletes an existing event.tag.category record.
func (c *Client) DeleteEventTagCategory(id int64) error {
	return c.DeleteEventTagCategorys([]int64{id})
}

// DeleteEventTagCategorys deletes existing event.tag.category records.
func (c *Client) DeleteEventTagCategorys(ids []int64) error {
	return c.Delete(EventTagCategoryModel, ids)
}

// GetEventTagCategory gets event.tag.category existing record.
func (c *Client) GetEventTagCategory(id int64) (*EventTagCategory, error) {
	etcs, err := c.GetEventTagCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if etcs != nil && len(*etcs) > 0 {
		return &((*etcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of event.tag.category not found", id)
}

// GetEventTagCategorys gets event.tag.category existing records.
func (c *Client) GetEventTagCategorys(ids []int64) (*EventTagCategorys, error) {
	etcs := &EventTagCategorys{}
	if err := c.Read(EventTagCategoryModel, ids, nil, etcs); err != nil {
		return nil, err
	}
	return etcs, nil
}

// FindEventTagCategory finds event.tag.category record by querying it with criteria.
func (c *Client) FindEventTagCategory(criteria *Criteria) (*EventTagCategory, error) {
	etcs := &EventTagCategorys{}
	if err := c.SearchRead(EventTagCategoryModel, criteria, NewOptions().Limit(1), etcs); err != nil {
		return nil, err
	}
	if etcs != nil && len(*etcs) > 0 {
		return &((*etcs)[0]), nil
	}
	return nil, fmt.Errorf("event.tag.category was not found with criteria %v", criteria)
}

// FindEventTagCategorys finds event.tag.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTagCategorys(criteria *Criteria, options *Options) (*EventTagCategorys, error) {
	etcs := &EventTagCategorys{}
	if err := c.SearchRead(EventTagCategoryModel, criteria, options, etcs); err != nil {
		return nil, err
	}
	return etcs, nil
}

// FindEventTagCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventTagCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(EventTagCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindEventTagCategoryId finds record id by querying it with criteria.
func (c *Client) FindEventTagCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventTagCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("event.tag.category was not found with criteria %v and options %v", criteria, options)
}
