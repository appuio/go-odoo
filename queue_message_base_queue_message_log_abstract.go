package odoo

import (
	"fmt"
)

// QueueMessageBaseQueueMessageLogAbstract represents queue_message_base.queue_message_log_abstract model.
type QueueMessageBaseQueueMessageLogAbstract struct {
	CreateDate *Time   `xmlrpc:"create_date,omptempty"`
	Message    *String `xmlrpc:"message,omptempty"`
}

// QueueMessageBaseQueueMessageLogAbstracts represents array of queue_message_base.queue_message_log_abstract model.
type QueueMessageBaseQueueMessageLogAbstracts []QueueMessageBaseQueueMessageLogAbstract

// QueueMessageBaseQueueMessageLogAbstractModel is the odoo model name.
const QueueMessageBaseQueueMessageLogAbstractModel = "queue_message_base.queue_message_log_abstract"

// Many2One convert QueueMessageBaseQueueMessageLogAbstract to *Many2One.
func (qq *QueueMessageBaseQueueMessageLogAbstract) Many2One() *Many2One {
	return NewMany2One(qq.Id.Get(), "")
}

// CreateQueueMessageBaseQueueMessageLogAbstract creates a new queue_message_base.queue_message_log_abstract model and returns its id.
func (c *Client) CreateQueueMessageBaseQueueMessageLogAbstract(qq *QueueMessageBaseQueueMessageLogAbstract) (int64, error) {
	ids, err := c.CreateQueueMessageBaseQueueMessageLogAbstracts([]*QueueMessageBaseQueueMessageLogAbstract{qq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQueueMessageBaseQueueMessageLogAbstract creates a new queue_message_base.queue_message_log_abstract model and returns its id.
func (c *Client) CreateQueueMessageBaseQueueMessageLogAbstracts(qqs []*QueueMessageBaseQueueMessageLogAbstract) ([]int64, error) {
	var vv []interface{}
	for _, v := range qqs {
		vv = append(vv, v)
	}
	return c.Create(QueueMessageBaseQueueMessageLogAbstractModel, vv)
}

// UpdateQueueMessageBaseQueueMessageLogAbstract updates an existing queue_message_base.queue_message_log_abstract record.
func (c *Client) UpdateQueueMessageBaseQueueMessageLogAbstract(qq *QueueMessageBaseQueueMessageLogAbstract) error {
	return c.UpdateQueueMessageBaseQueueMessageLogAbstracts([]int64{qq.Id.Get()}, qq)
}

// UpdateQueueMessageBaseQueueMessageLogAbstracts updates existing queue_message_base.queue_message_log_abstract records.
// All records (represented by ids) will be updated by qq values.
func (c *Client) UpdateQueueMessageBaseQueueMessageLogAbstracts(ids []int64, qq *QueueMessageBaseQueueMessageLogAbstract) error {
	return c.Update(QueueMessageBaseQueueMessageLogAbstractModel, ids, qq)
}

// DeleteQueueMessageBaseQueueMessageLogAbstract deletes an existing queue_message_base.queue_message_log_abstract record.
func (c *Client) DeleteQueueMessageBaseQueueMessageLogAbstract(id int64) error {
	return c.DeleteQueueMessageBaseQueueMessageLogAbstracts([]int64{id})
}

// DeleteQueueMessageBaseQueueMessageLogAbstracts deletes existing queue_message_base.queue_message_log_abstract records.
func (c *Client) DeleteQueueMessageBaseQueueMessageLogAbstracts(ids []int64) error {
	return c.Delete(QueueMessageBaseQueueMessageLogAbstractModel, ids)
}

// GetQueueMessageBaseQueueMessageLogAbstract gets queue_message_base.queue_message_log_abstract existing record.
func (c *Client) GetQueueMessageBaseQueueMessageLogAbstract(id int64) (*QueueMessageBaseQueueMessageLogAbstract, error) {
	qqs, err := c.GetQueueMessageBaseQueueMessageLogAbstracts([]int64{id})
	if err != nil {
		return nil, err
	}
	if qqs != nil && len(*qqs) > 0 {
		return &((*qqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue_message_base.queue_message_log_abstract not found", id)
}

// GetQueueMessageBaseQueueMessageLogAbstracts gets queue_message_base.queue_message_log_abstract existing records.
func (c *Client) GetQueueMessageBaseQueueMessageLogAbstracts(ids []int64) (*QueueMessageBaseQueueMessageLogAbstracts, error) {
	qqs := &QueueMessageBaseQueueMessageLogAbstracts{}
	if err := c.Read(QueueMessageBaseQueueMessageLogAbstractModel, ids, nil, qqs); err != nil {
		return nil, err
	}
	return qqs, nil
}

// FindQueueMessageBaseQueueMessageLogAbstract finds queue_message_base.queue_message_log_abstract record by querying it with criteria.
func (c *Client) FindQueueMessageBaseQueueMessageLogAbstract(criteria *Criteria) (*QueueMessageBaseQueueMessageLogAbstract, error) {
	qqs := &QueueMessageBaseQueueMessageLogAbstracts{}
	if err := c.SearchRead(QueueMessageBaseQueueMessageLogAbstractModel, criteria, NewOptions().Limit(1), qqs); err != nil {
		return nil, err
	}
	if qqs != nil && len(*qqs) > 0 {
		return &((*qqs)[0]), nil
	}
	return nil, fmt.Errorf("queue_message_base.queue_message_log_abstract was not found with criteria %v", criteria)
}

// FindQueueMessageBaseQueueMessageLogAbstracts finds queue_message_base.queue_message_log_abstract records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageBaseQueueMessageLogAbstracts(criteria *Criteria, options *Options) (*QueueMessageBaseQueueMessageLogAbstracts, error) {
	qqs := &QueueMessageBaseQueueMessageLogAbstracts{}
	if err := c.SearchRead(QueueMessageBaseQueueMessageLogAbstractModel, criteria, options, qqs); err != nil {
		return nil, err
	}
	return qqs, nil
}

// FindQueueMessageBaseQueueMessageLogAbstractIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageBaseQueueMessageLogAbstractIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueMessageBaseQueueMessageLogAbstractModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueMessageBaseQueueMessageLogAbstractId finds record id by querying it with criteria.
func (c *Client) FindQueueMessageBaseQueueMessageLogAbstractId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueMessageBaseQueueMessageLogAbstractModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue_message_base.queue_message_log_abstract was not found with criteria %v and options %v", criteria, options)
}
