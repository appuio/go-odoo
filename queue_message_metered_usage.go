package odoo

import (
	"fmt"
)

// QueueMessageMeteredUsage represents queue_message.metered_usage model.
type QueueMessageMeteredUsage struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	JobsCount   *Int       `xmlrpc:"jobs_count,omptempty"`
	JobsIds     *Relation  `xmlrpc:"jobs_ids,omptempty"`
	LogIds      *Relation  `xmlrpc:"log_ids,omptempty"`
	Payload     *String    `xmlrpc:"payload,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	StatusCode  *String    `xmlrpc:"status_code,omptempty"`
	StatusInfo  *String    `xmlrpc:"status_info,omptempty"`
	Timestamp   *Time      `xmlrpc:"timestamp,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// QueueMessageMeteredUsages represents array of queue_message.metered_usage model.
type QueueMessageMeteredUsages []QueueMessageMeteredUsage

// QueueMessageMeteredUsageModel is the odoo model name.
const QueueMessageMeteredUsageModel = "queue_message.metered_usage"

// Many2One convert QueueMessageMeteredUsage to *Many2One.
func (qm *QueueMessageMeteredUsage) Many2One() *Many2One {
	return NewMany2One(qm.Id.Get(), "")
}

// CreateQueueMessageMeteredUsage creates a new queue_message.metered_usage model and returns its id.
func (c *Client) CreateQueueMessageMeteredUsage(qm *QueueMessageMeteredUsage) (int64, error) {
	ids, err := c.CreateQueueMessageMeteredUsages([]*QueueMessageMeteredUsage{qm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQueueMessageMeteredUsage creates a new queue_message.metered_usage model and returns its id.
func (c *Client) CreateQueueMessageMeteredUsages(qms []*QueueMessageMeteredUsage) ([]int64, error) {
	var vv []interface{}
	for _, v := range qms {
		vv = append(vv, v)
	}
	return c.Create(QueueMessageMeteredUsageModel, vv)
}

// UpdateQueueMessageMeteredUsage updates an existing queue_message.metered_usage record.
func (c *Client) UpdateQueueMessageMeteredUsage(qm *QueueMessageMeteredUsage) error {
	return c.UpdateQueueMessageMeteredUsages([]int64{qm.Id.Get()}, qm)
}

// UpdateQueueMessageMeteredUsages updates existing queue_message.metered_usage records.
// All records (represented by ids) will be updated by qm values.
func (c *Client) UpdateQueueMessageMeteredUsages(ids []int64, qm *QueueMessageMeteredUsage) error {
	return c.Update(QueueMessageMeteredUsageModel, ids, qm)
}

// DeleteQueueMessageMeteredUsage deletes an existing queue_message.metered_usage record.
func (c *Client) DeleteQueueMessageMeteredUsage(id int64) error {
	return c.DeleteQueueMessageMeteredUsages([]int64{id})
}

// DeleteQueueMessageMeteredUsages deletes existing queue_message.metered_usage records.
func (c *Client) DeleteQueueMessageMeteredUsages(ids []int64) error {
	return c.Delete(QueueMessageMeteredUsageModel, ids)
}

// GetQueueMessageMeteredUsage gets queue_message.metered_usage existing record.
func (c *Client) GetQueueMessageMeteredUsage(id int64) (*QueueMessageMeteredUsage, error) {
	qms, err := c.GetQueueMessageMeteredUsages([]int64{id})
	if err != nil {
		return nil, err
	}
	if qms != nil && len(*qms) > 0 {
		return &((*qms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue_message.metered_usage not found", id)
}

// GetQueueMessageMeteredUsages gets queue_message.metered_usage existing records.
func (c *Client) GetQueueMessageMeteredUsages(ids []int64) (*QueueMessageMeteredUsages, error) {
	qms := &QueueMessageMeteredUsages{}
	if err := c.Read(QueueMessageMeteredUsageModel, ids, nil, qms); err != nil {
		return nil, err
	}
	return qms, nil
}

// FindQueueMessageMeteredUsage finds queue_message.metered_usage record by querying it with criteria.
func (c *Client) FindQueueMessageMeteredUsage(criteria *Criteria) (*QueueMessageMeteredUsage, error) {
	qms := &QueueMessageMeteredUsages{}
	if err := c.SearchRead(QueueMessageMeteredUsageModel, criteria, NewOptions().Limit(1), qms); err != nil {
		return nil, err
	}
	if qms != nil && len(*qms) > 0 {
		return &((*qms)[0]), nil
	}
	return nil, fmt.Errorf("queue_message.metered_usage was not found with criteria %v", criteria)
}

// FindQueueMessageMeteredUsages finds queue_message.metered_usage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageMeteredUsages(criteria *Criteria, options *Options) (*QueueMessageMeteredUsages, error) {
	qms := &QueueMessageMeteredUsages{}
	if err := c.SearchRead(QueueMessageMeteredUsageModel, criteria, options, qms); err != nil {
		return nil, err
	}
	return qms, nil
}

// FindQueueMessageMeteredUsageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageMeteredUsageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueMessageMeteredUsageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueMessageMeteredUsageId finds record id by querying it with criteria.
func (c *Client) FindQueueMessageMeteredUsageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueMessageMeteredUsageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue_message.metered_usage was not found with criteria %v and options %v", criteria, options)
}
