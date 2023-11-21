package odoo

import (
	"fmt"
)

// HelpdeskStage represents helpdesk.stage model.
type HelpdeskStage struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	Active        *Bool     `xmlrpc:"active,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	Description   *String   `xmlrpc:"description,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Fold          *Bool     `xmlrpc:"fold,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	LegendBlocked *String   `xmlrpc:"legend_blocked,omptempty"`
	LegendDone    *String   `xmlrpc:"legend_done,omptempty"`
	LegendNormal  *String   `xmlrpc:"legend_normal,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	Sequence      *Int      `xmlrpc:"sequence,omptempty"`
	SmsTemplateId *Many2One `xmlrpc:"sms_template_id,omptempty"`
	TeamIds       *Relation `xmlrpc:"team_ids,omptempty"`
	TemplateId    *Many2One `xmlrpc:"template_id,omptempty"`
	TicketCount   *Int      `xmlrpc:"ticket_count,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// HelpdeskStages represents array of helpdesk.stage model.
type HelpdeskStages []HelpdeskStage

// HelpdeskStageModel is the odoo model name.
const HelpdeskStageModel = "helpdesk.stage"

// Many2One convert HelpdeskStage to *Many2One.
func (hs *HelpdeskStage) Many2One() *Many2One {
	return NewMany2One(hs.Id.Get(), "")
}

// CreateHelpdeskStage creates a new helpdesk.stage model and returns its id.
func (c *Client) CreateHelpdeskStage(hs *HelpdeskStage) (int64, error) {
	ids, err := c.CreateHelpdeskStages([]*HelpdeskStage{hs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHelpdeskStage creates a new helpdesk.stage model and returns its id.
func (c *Client) CreateHelpdeskStages(hss []*HelpdeskStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range hss {
		vv = append(vv, v)
	}
	return c.Create(HelpdeskStageModel, vv)
}

// UpdateHelpdeskStage updates an existing helpdesk.stage record.
func (c *Client) UpdateHelpdeskStage(hs *HelpdeskStage) error {
	return c.UpdateHelpdeskStages([]int64{hs.Id.Get()}, hs)
}

// UpdateHelpdeskStages updates existing helpdesk.stage records.
// All records (represented by ids) will be updated by hs values.
func (c *Client) UpdateHelpdeskStages(ids []int64, hs *HelpdeskStage) error {
	return c.Update(HelpdeskStageModel, ids, hs)
}

// DeleteHelpdeskStage deletes an existing helpdesk.stage record.
func (c *Client) DeleteHelpdeskStage(id int64) error {
	return c.DeleteHelpdeskStages([]int64{id})
}

// DeleteHelpdeskStages deletes existing helpdesk.stage records.
func (c *Client) DeleteHelpdeskStages(ids []int64) error {
	return c.Delete(HelpdeskStageModel, ids)
}

// GetHelpdeskStage gets helpdesk.stage existing record.
func (c *Client) GetHelpdeskStage(id int64) (*HelpdeskStage, error) {
	hss, err := c.GetHelpdeskStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if hss != nil && len(*hss) > 0 {
		return &((*hss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of helpdesk.stage not found", id)
}

// GetHelpdeskStages gets helpdesk.stage existing records.
func (c *Client) GetHelpdeskStages(ids []int64) (*HelpdeskStages, error) {
	hss := &HelpdeskStages{}
	if err := c.Read(HelpdeskStageModel, ids, nil, hss); err != nil {
		return nil, err
	}
	return hss, nil
}

// FindHelpdeskStage finds helpdesk.stage record by querying it with criteria.
func (c *Client) FindHelpdeskStage(criteria *Criteria) (*HelpdeskStage, error) {
	hss := &HelpdeskStages{}
	if err := c.SearchRead(HelpdeskStageModel, criteria, NewOptions().Limit(1), hss); err != nil {
		return nil, err
	}
	if hss != nil && len(*hss) > 0 {
		return &((*hss)[0]), nil
	}
	return nil, fmt.Errorf("helpdesk.stage was not found with criteria %v", criteria)
}

// FindHelpdeskStages finds helpdesk.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskStages(criteria *Criteria, options *Options) (*HelpdeskStages, error) {
	hss := &HelpdeskStages{}
	if err := c.SearchRead(HelpdeskStageModel, criteria, options, hss); err != nil {
		return nil, err
	}
	return hss, nil
}

// FindHelpdeskStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHelpdeskStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HelpdeskStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHelpdeskStageId finds record id by querying it with criteria.
func (c *Client) FindHelpdeskStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HelpdeskStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("helpdesk.stage was not found with criteria %v and options %v", criteria, options)
}
