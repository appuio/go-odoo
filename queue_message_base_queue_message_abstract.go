package odoo

import (
	"fmt"
)

// QueueMessageBaseQueueMessageAbstract represents queue_message_base.queue_message_abstract model.
type QueueMessageBaseQueueMessageAbstract struct {
	CreateDate *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid  *Many2One  `xmlrpc:"create_uid,omptempty"`
	JobsCount  *Int       `xmlrpc:"jobs_count,omptempty"`
	JobsIds    *Relation  `xmlrpc:"jobs_ids,omptempty"`
	Payload    *String    `xmlrpc:"payload,omptempty"`
	State      *Selection `xmlrpc:"state,omptempty"`
	StatusCode *String    `xmlrpc:"status_code,omptempty"`
	StatusInfo *String    `xmlrpc:"status_info,omptempty"`
	Timestamp  *Time      `xmlrpc:"timestamp,omptempty"`
}

// QueueMessageBaseQueueMessageAbstracts represents array of queue_message_base.queue_message_abstract model.
type QueueMessageBaseQueueMessageAbstracts []QueueMessageBaseQueueMessageAbstract

// QueueMessageBaseQueueMessageAbstractModel is the odoo model name.
const QueueMessageBaseQueueMessageAbstractModel = "queue_message_base.queue_message_abstract"

// Many2One convert QueueMessageBaseQueueMessageAbstract to *Many2One.
func (qq *QueueMessageBaseQueueMessageAbstract) Many2One() *Many2One {
	return NewMany2One(qq.Id.Get(), "")
}

// CreateQueueMessageBaseQueueMessageAbstract creates a new queue_message_base.queue_message_abstract model and returns its id.
func (c *Client) CreateQueueMessageBaseQueueMessageAbstract(qq *QueueMessageBaseQueueMessageAbstract) (int64, error) {
	ids, err := c.CreateQueueMessageBaseQueueMessageAbstracts([]*QueueMessageBaseQueueMessageAbstract{qq})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateQueueMessageBaseQueueMessageAbstract creates a new queue_message_base.queue_message_abstract model and returns its id.
func (c *Client) CreateQueueMessageBaseQueueMessageAbstracts(qqs []*QueueMessageBaseQueueMessageAbstract) ([]int64, error) {
	var vv []interface{}
	for _, v := range qqs {
		vv = append(vv, v)
	}
	return c.Create(QueueMessageBaseQueueMessageAbstractModel, vv)
}

// UpdateQueueMessageBaseQueueMessageAbstract updates an existing queue_message_base.queue_message_abstract record.
func (c *Client) UpdateQueueMessageBaseQueueMessageAbstract(qq *QueueMessageBaseQueueMessageAbstract) error {
	return c.UpdateQueueMessageBaseQueueMessageAbstracts([]int64{qq.Id.Get()}, qq)
}

// UpdateQueueMessageBaseQueueMessageAbstracts updates existing queue_message_base.queue_message_abstract records.
// All records (represented by ids) will be updated by qq values.
func (c *Client) UpdateQueueMessageBaseQueueMessageAbstracts(ids []int64, qq *QueueMessageBaseQueueMessageAbstract) error {
	return c.Update(QueueMessageBaseQueueMessageAbstractModel, ids, qq)
}

// DeleteQueueMessageBaseQueueMessageAbstract deletes an existing queue_message_base.queue_message_abstract record.
func (c *Client) DeleteQueueMessageBaseQueueMessageAbstract(id int64) error {
	return c.DeleteQueueMessageBaseQueueMessageAbstracts([]int64{id})
}

// DeleteQueueMessageBaseQueueMessageAbstracts deletes existing queue_message_base.queue_message_abstract records.
func (c *Client) DeleteQueueMessageBaseQueueMessageAbstracts(ids []int64) error {
	return c.Delete(QueueMessageBaseQueueMessageAbstractModel, ids)
}

// GetQueueMessageBaseQueueMessageAbstract gets queue_message_base.queue_message_abstract existing record.
func (c *Client) GetQueueMessageBaseQueueMessageAbstract(id int64) (*QueueMessageBaseQueueMessageAbstract, error) {
	qqs, err := c.GetQueueMessageBaseQueueMessageAbstracts([]int64{id})
	if err != nil {
		return nil, err
	}
	if qqs != nil && len(*qqs) > 0 {
		return &((*qqs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of queue_message_base.queue_message_abstract not found", id)
}

// GetQueueMessageBaseQueueMessageAbstracts gets queue_message_base.queue_message_abstract existing records.
func (c *Client) GetQueueMessageBaseQueueMessageAbstracts(ids []int64) (*QueueMessageBaseQueueMessageAbstracts, error) {
	qqs := &QueueMessageBaseQueueMessageAbstracts{}
	if err := c.Read(QueueMessageBaseQueueMessageAbstractModel, ids, nil, qqs); err != nil {
		return nil, err
	}
	return qqs, nil
}

// FindQueueMessageBaseQueueMessageAbstract finds queue_message_base.queue_message_abstract record by querying it with criteria.
func (c *Client) FindQueueMessageBaseQueueMessageAbstract(criteria *Criteria) (*QueueMessageBaseQueueMessageAbstract, error) {
	qqs := &QueueMessageBaseQueueMessageAbstracts{}
	if err := c.SearchRead(QueueMessageBaseQueueMessageAbstractModel, criteria, NewOptions().Limit(1), qqs); err != nil {
		return nil, err
	}
	if qqs != nil && len(*qqs) > 0 {
		return &((*qqs)[0]), nil
	}
	return nil, fmt.Errorf("queue_message_base.queue_message_abstract was not found with criteria %v", criteria)
}

// FindQueueMessageBaseQueueMessageAbstracts finds queue_message_base.queue_message_abstract records by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageBaseQueueMessageAbstracts(criteria *Criteria, options *Options) (*QueueMessageBaseQueueMessageAbstracts, error) {
	qqs := &QueueMessageBaseQueueMessageAbstracts{}
	if err := c.SearchRead(QueueMessageBaseQueueMessageAbstractModel, criteria, options, qqs); err != nil {
		return nil, err
	}
	return qqs, nil
}

// FindQueueMessageBaseQueueMessageAbstractIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindQueueMessageBaseQueueMessageAbstractIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(QueueMessageBaseQueueMessageAbstractModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindQueueMessageBaseQueueMessageAbstractId finds record id by querying it with criteria.
func (c *Client) FindQueueMessageBaseQueueMessageAbstractId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(QueueMessageBaseQueueMessageAbstractModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("queue_message_base.queue_message_abstract was not found with criteria %v and options %v", criteria, options)
}
