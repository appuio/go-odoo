package odoo

import (
	"fmt"
)

// SlotPlanningSelectSend represents slot.planning.select.send model.
type SlotPlanningSelectSend struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	EmployeeIds *Relation `xmlrpc:"employee_ids,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	SlotId      *Many2One `xmlrpc:"slot_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SlotPlanningSelectSends represents array of slot.planning.select.send model.
type SlotPlanningSelectSends []SlotPlanningSelectSend

// SlotPlanningSelectSendModel is the odoo model name.
const SlotPlanningSelectSendModel = "slot.planning.select.send"

// Many2One convert SlotPlanningSelectSend to *Many2One.
func (spss *SlotPlanningSelectSend) Many2One() *Many2One {
	return NewMany2One(spss.Id.Get(), "")
}

// CreateSlotPlanningSelectSend creates a new slot.planning.select.send model and returns its id.
func (c *Client) CreateSlotPlanningSelectSend(spss *SlotPlanningSelectSend) (int64, error) {
	ids, err := c.CreateSlotPlanningSelectSends([]*SlotPlanningSelectSend{spss})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSlotPlanningSelectSend creates a new slot.planning.select.send model and returns its id.
func (c *Client) CreateSlotPlanningSelectSends(spsss []*SlotPlanningSelectSend) ([]int64, error) {
	var vv []interface{}
	for _, v := range spsss {
		vv = append(vv, v)
	}
	return c.Create(SlotPlanningSelectSendModel, vv)
}

// UpdateSlotPlanningSelectSend updates an existing slot.planning.select.send record.
func (c *Client) UpdateSlotPlanningSelectSend(spss *SlotPlanningSelectSend) error {
	return c.UpdateSlotPlanningSelectSends([]int64{spss.Id.Get()}, spss)
}

// UpdateSlotPlanningSelectSends updates existing slot.planning.select.send records.
// All records (represented by ids) will be updated by spss values.
func (c *Client) UpdateSlotPlanningSelectSends(ids []int64, spss *SlotPlanningSelectSend) error {
	return c.Update(SlotPlanningSelectSendModel, ids, spss)
}

// DeleteSlotPlanningSelectSend deletes an existing slot.planning.select.send record.
func (c *Client) DeleteSlotPlanningSelectSend(id int64) error {
	return c.DeleteSlotPlanningSelectSends([]int64{id})
}

// DeleteSlotPlanningSelectSends deletes existing slot.planning.select.send records.
func (c *Client) DeleteSlotPlanningSelectSends(ids []int64) error {
	return c.Delete(SlotPlanningSelectSendModel, ids)
}

// GetSlotPlanningSelectSend gets slot.planning.select.send existing record.
func (c *Client) GetSlotPlanningSelectSend(id int64) (*SlotPlanningSelectSend, error) {
	spsss, err := c.GetSlotPlanningSelectSends([]int64{id})
	if err != nil {
		return nil, err
	}
	if spsss != nil && len(*spsss) > 0 {
		return &((*spsss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of slot.planning.select.send not found", id)
}

// GetSlotPlanningSelectSends gets slot.planning.select.send existing records.
func (c *Client) GetSlotPlanningSelectSends(ids []int64) (*SlotPlanningSelectSends, error) {
	spsss := &SlotPlanningSelectSends{}
	if err := c.Read(SlotPlanningSelectSendModel, ids, nil, spsss); err != nil {
		return nil, err
	}
	return spsss, nil
}

// FindSlotPlanningSelectSend finds slot.planning.select.send record by querying it with criteria.
func (c *Client) FindSlotPlanningSelectSend(criteria *Criteria) (*SlotPlanningSelectSend, error) {
	spsss := &SlotPlanningSelectSends{}
	if err := c.SearchRead(SlotPlanningSelectSendModel, criteria, NewOptions().Limit(1), spsss); err != nil {
		return nil, err
	}
	if spsss != nil && len(*spsss) > 0 {
		return &((*spsss)[0]), nil
	}
	return nil, fmt.Errorf("slot.planning.select.send was not found with criteria %v", criteria)
}

// FindSlotPlanningSelectSends finds slot.planning.select.send records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlotPlanningSelectSends(criteria *Criteria, options *Options) (*SlotPlanningSelectSends, error) {
	spsss := &SlotPlanningSelectSends{}
	if err := c.SearchRead(SlotPlanningSelectSendModel, criteria, options, spsss); err != nil {
		return nil, err
	}
	return spsss, nil
}

// FindSlotPlanningSelectSendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSlotPlanningSelectSendIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SlotPlanningSelectSendModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSlotPlanningSelectSendId finds record id by querying it with criteria.
func (c *Client) FindSlotPlanningSelectSendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SlotPlanningSelectSendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("slot.planning.select.send was not found with criteria %v and options %v", criteria, options)
}
