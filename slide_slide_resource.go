package odoo

import (
	"fmt"
)

// SlideSlideResource represents slide.slide.resource model.
type SlideSlideResource struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	Data         *String    `xmlrpc:"data,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	FileName     *String    `xmlrpc:"file_name,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	Link         *String    `xmlrpc:"link,omptempty"`
	Name         *String    `xmlrpc:"name,omptempty"`
	ResourceType *Selection `xmlrpc:"resource_type,omptempty"`
	SlideId      *Many2One  `xmlrpc:"slide_id,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// SlideSlideResources represents array of slide.slide.resource model.
type SlideSlideResources []SlideSlideResource

// SlideSlideResourceModel is the odoo model name.
const SlideSlideResourceModel = "slide.slide.resource"

// Many2One convert SlideSlideResource to *Many2One.
func (ssr *SlideSlideResource) Many2One() *Many2One {
	return NewMany2One(ssr.Id.Get(), "")
}

// CreateSlideSlideResource creates a new slide.slide.resource model and returns its id.
func (c *Client) CreateSlideSlideResource(ssr *SlideSlideResource) (int64, error) {
	ids, err := c.CreateSlideSlideResources([]*SlideSlideResource{ssr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSlideSlideResource creates a new slide.slide.resource model and returns its id.
func (c *Client) CreateSlideSlideResources(ssrs []*SlideSlideResource) ([]int64, error) {
	var vv []interface{}
	for _, v := range ssrs {
		vv = append(vv, v)
	}
	return c.Create(SlideSlideResourceModel, vv)
}

// UpdateSlideSlideResource updates an existing slide.slide.resource record.
func (c *Client) UpdateSlideSlideResource(ssr *SlideSlideResource) error {
	return c.UpdateSlideSlideResources([]int64{ssr.Id.Get()}, ssr)
}

// UpdateSlideSlideResources updates existing slide.slide.resource records.
// All records (represented by ids) will be updated by ssr values.
func (c *Client) UpdateSlideSlideResources(ids []int64, ssr *SlideSlideResource) error {
	return c.Update(SlideSlideResourceModel, ids, ssr)
}

// DeleteSlideSlideResource deletes an existing slide.slide.resource record.
func (c *Client) DeleteSlideSlideResource(id int64) error {
	return c.DeleteSlideSlideResources([]int64{id})
}

// DeleteSlideSlideResources deletes existing slide.slide.resource records.
func (c *Client) DeleteSlideSlideResources(ids []int64) error {
	return c.Delete(SlideSlideResourceModel, ids)
}

// GetSlideSlideResource gets slide.slide.resource existing record.
func (c *Client) GetSlideSlideResource(id int64) (*SlideSlideResource, error) {
	ssrs, err := c.GetSlideSlideResources([]int64{id})
	if err != nil {
		return nil, err
	}
	if ssrs != nil && len(*ssrs) > 0 {
		return &((*ssrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slide.slide.resource not found", id)
}

// GetSlideSlideResources gets slide.slide.resource existing records.
func (c *Client) GetSlideSlideResources(ids []int64) (*SlideSlideResources, error) {
	ssrs := &SlideSlideResources{}
	if err := c.Read(SlideSlideResourceModel, ids, nil, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSlideSlideResource finds slide.slide.resource record by querying it with criteria.
func (c *Client) FindSlideSlideResource(criteria *Criteria) (*SlideSlideResource, error) {
	ssrs := &SlideSlideResources{}
	if err := c.SearchRead(SlideSlideResourceModel, criteria, NewOptions().Limit(1), ssrs); err != nil {
		return nil, err
	}
	if ssrs != nil && len(*ssrs) > 0 {
		return &((*ssrs)[0]), nil
	}
	return nil, fmt.Errorf("slide.slide.resource was not found with criteria %v", criteria)
}

// FindSlideSlideResources finds slide.slide.resource records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideResources(criteria *Criteria, options *Options) (*SlideSlideResources, error) {
	ssrs := &SlideSlideResources{}
	if err := c.SearchRead(SlideSlideResourceModel, criteria, options, ssrs); err != nil {
		return nil, err
	}
	return ssrs, nil
}

// FindSlideSlideResourceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlideSlideResourceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlideSlideResourceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlideSlideResourceId finds record id by querying it with criteria.
func (c *Client) FindSlideSlideResourceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlideSlideResourceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slide.slide.resource was not found with criteria %v and options %v", criteria, options)
}
