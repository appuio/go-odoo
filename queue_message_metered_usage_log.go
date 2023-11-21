package odoo

import (
	"fmt"
)

// QueueMessageMeteredUsageLog represents queue_message.metered_usage_log model.
type QueueMessageMeteredUsageLog struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	Message        *String   `xmlrpc:"message,omptempty"`
	QueueMessageId *Many2One `xmlrpc:"queue_message_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// QueueMessageMeteredUsageLogs represents array of queue_message.metered_usage_log model.
type QueueMessageMeteredUsageLogs []QueueMessageMeteredUsageLog

// QueueMessageMeteredUsageLogModel is the odoo model name.
const QueueMessageMeteredUsageLogModel = "queue_message.metered_usage_log"

// Many2One convert QueueMessageMeteredUsageLog to *Many2One.
func (qm *QueueMessageMeteredUsageLog) Many2One() *Many2One {
	return NewMany2One(qm.Id.Get(), "")
}

// CreateQueueMessageMeteredUsageLog creates a new queue_message.metered_usage_log model and returns its id.
func (c *Client) CreateQueueMessageMeteredUsageLog(qm *QueueMessageMeteredUsageLog) (int64, error) {
	ids, err := c.CreateQueueMessageMeteredUsageLogs([]*QueueMessageMeteredUsageLog{qm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQueueMessageMeteredUsageLog creates a new queue_message.metered_usage_log model and returns its id.
func (c *Client) CreateQueueMessageMeteredUsageLogs(qms []*QueueMessageMeteredUsageLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range qms {
		vv = append(vv, v)
	}
	return c.Create(QueueMessageMeteredUsageLogModel, vv)
}

// UpdateQueueMessageMeteredUsageLog updates an existing queue_message.metered_usage_log record.
func (c *Client) UpdateQueueMessageMeteredUsageLog(qm *QueueMessageMeteredUsageLog) error {
	return c.UpdateQueueMessageMeteredUsageLogs([]int64{qm.Id.Get()}, qm)
}

// UpdateQueueMessageMeteredUsageLogs updates existing queue_message.metered_usage_log records.
// All records (represented by ids) will be updated by qm values.
func (c *Client) UpdateQueueMessageMeteredUsageLogs(ids []int64, qm *QueueMessageMeteredUsageLog) error {
	return c.Update(QueueMessageMeteredUsageLogModel, ids, qm)
}

// DeleteQueueMessageMeteredUsageLog deletes an existing queue_message.metered_usage_log record.
func (c *Client) DeleteQueueMessageMeteredUsageLog(id int64) error {
	return c.DeleteQueueMessageMeteredUsageLogs([]int64{id})
}

// DeleteQueueMessageMeteredUsageLogs deletes existing queue_message.metered_usage_log records.
func (c *Client) DeleteQueueMessageMeteredUsageLogs(ids []int64) error {
	return c.Delete(QueueMessageMeteredUsageLogModel, ids)
}

// GetQueueMessageMeteredUsageLog gets queue_message.metered_usage_log existing record.
func (c *Client) GetQueueMessageMeteredUsageLog(id int64) (*QueueMessageMeteredUsageLog, error) {
	qms, err := c.GetQueueMessageMeteredUsageLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	if qms != nil && len(*qms) > 0 {
		return &((*qms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue_message.metered_usage_log not found", id)
}

// GetQueueMessageMeteredUsageLogs gets queue_message.metered_usage_log existing records.
func (c *Client) GetQueueMessageMeteredUsageLogs(ids []int64) (*QueueMessageMeteredUsageLogs, error) {
	qms := &QueueMessageMeteredUsageLogs{}
	if err := c.Read(QueueMessageMeteredUsageLogModel, ids, nil, qms); err != nil {
		return nil, err
	}
	return qms, nil
}

// FindQueueMessageMeteredUsageLog finds queue_message.metered_usage_log record by querying it with criteria.
func (c *Client) FindQueueMessageMeteredUsageLog(criteria *Criteria) (*QueueMessageMeteredUsageLog, error) {
	qms := &QueueMessageMeteredUsageLogs{}
	if err := c.SearchRead(QueueMessageMeteredUsageLogModel, criteria, NewOptions().Limit(1), qms); err != nil {
		return nil, err
	}
	if qms != nil && len(*qms) > 0 {
		return &((*qms)[0]), nil
	}
	return nil, fmt.Errorf("queue_message.metered_usage_log was not found with criteria %v", criteria)
}

// FindQueueMessageMeteredUsageLogs finds queue_message.metered_usage_log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageMeteredUsageLogs(criteria *Criteria, options *Options) (*QueueMessageMeteredUsageLogs, error) {
	qms := &QueueMessageMeteredUsageLogs{}
	if err := c.SearchRead(QueueMessageMeteredUsageLogModel, criteria, options, qms); err != nil {
		return nil, err
	}
	return qms, nil
}

// FindQueueMessageMeteredUsageLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageMeteredUsageLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueMessageMeteredUsageLogModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueMessageMeteredUsageLogId finds record id by querying it with criteria.
func (c *Client) FindQueueMessageMeteredUsageLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueMessageMeteredUsageLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue_message.metered_usage_log was not found with criteria %v and options %v", criteria, options)
}
