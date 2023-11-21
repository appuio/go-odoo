package odoo

import (
	"fmt"
)

// StudioApprovalRequest represents studio.approval.request model.
type StudioApprovalRequest struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	MailActivityId *Many2One `xmlrpc:"mail_activity_id,omptempty"`
	ResId          *Many2One `xmlrpc:"res_id,omptempty"`
	RuleId         *Many2One `xmlrpc:"rule_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StudioApprovalRequests represents array of studio.approval.request model.
type StudioApprovalRequests []StudioApprovalRequest

// StudioApprovalRequestModel is the odoo model name.
const StudioApprovalRequestModel = "studio.approval.request"

// Many2One convert StudioApprovalRequest to *Many2One.
func (sar *StudioApprovalRequest) Many2One() *Many2One {
	return NewMany2One(sar.Id.Get(), "")
}

// CreateStudioApprovalRequest creates a new studio.approval.request model and returns its id.
func (c *Client) CreateStudioApprovalRequest(sar *StudioApprovalRequest) (int64, error) {
	ids, err := c.CreateStudioApprovalRequests([]*StudioApprovalRequest{sar})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStudioApprovalRequest creates a new studio.approval.request model and returns its id.
func (c *Client) CreateStudioApprovalRequests(sars []*StudioApprovalRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range sars {
		vv = append(vv, v)
	}
	return c.Create(StudioApprovalRequestModel, vv)
}

// UpdateStudioApprovalRequest updates an existing studio.approval.request record.
func (c *Client) UpdateStudioApprovalRequest(sar *StudioApprovalRequest) error {
	return c.UpdateStudioApprovalRequests([]int64{sar.Id.Get()}, sar)
}

// UpdateStudioApprovalRequests updates existing studio.approval.request records.
// All records (represented by ids) will be updated by sar values.
func (c *Client) UpdateStudioApprovalRequests(ids []int64, sar *StudioApprovalRequest) error {
	return c.Update(StudioApprovalRequestModel, ids, sar)
}

// DeleteStudioApprovalRequest deletes an existing studio.approval.request record.
func (c *Client) DeleteStudioApprovalRequest(id int64) error {
	return c.DeleteStudioApprovalRequests([]int64{id})
}

// DeleteStudioApprovalRequests deletes existing studio.approval.request records.
func (c *Client) DeleteStudioApprovalRequests(ids []int64) error {
	return c.Delete(StudioApprovalRequestModel, ids)
}

// GetStudioApprovalRequest gets studio.approval.request existing record.
func (c *Client) GetStudioApprovalRequest(id int64) (*StudioApprovalRequest, error) {
	sars, err := c.GetStudioApprovalRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	if sars != nil && len(*sars) > 0 {
		return &((*sars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of studio.approval.request not found", id)
}

// GetStudioApprovalRequests gets studio.approval.request existing records.
func (c *Client) GetStudioApprovalRequests(ids []int64) (*StudioApprovalRequests, error) {
	sars := &StudioApprovalRequests{}
	if err := c.Read(StudioApprovalRequestModel, ids, nil, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRequest finds studio.approval.request record by querying it with criteria.
func (c *Client) FindStudioApprovalRequest(criteria *Criteria) (*StudioApprovalRequest, error) {
	sars := &StudioApprovalRequests{}
	if err := c.SearchRead(StudioApprovalRequestModel, criteria, NewOptions().Limit(1), sars); err != nil {
		return nil, err
	}
	if sars != nil && len(*sars) > 0 {
		return &((*sars)[0]), nil
	}
	return nil, fmt.Errorf("studio.approval.request was not found with criteria %v", criteria)
}

// FindStudioApprovalRequests finds studio.approval.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRequests(criteria *Criteria, options *Options) (*StudioApprovalRequests, error) {
	sars := &StudioApprovalRequests{}
	if err := c.SearchRead(StudioApprovalRequestModel, criteria, options, sars); err != nil {
		return nil, err
	}
	return sars, nil
}

// FindStudioApprovalRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStudioApprovalRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StudioApprovalRequestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStudioApprovalRequestId finds record id by querying it with criteria.
func (c *Client) FindStudioApprovalRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StudioApprovalRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("studio.approval.request was not found with criteria %v and options %v", criteria, options)
}
