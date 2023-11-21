package odoo

import (
	"fmt"
)

// TimerMixin represents timer.mixin model.
type TimerMixin struct {
	DisplayTimerPause        *Bool     `xmlrpc:"display_timer_pause,omptempty"`
	DisplayTimerResume       *Bool     `xmlrpc:"display_timer_resume,omptempty"`
	DisplayTimerStartPrimary *Bool     `xmlrpc:"display_timer_start_primary,omptempty"`
	DisplayTimerStop         *Bool     `xmlrpc:"display_timer_stop,omptempty"`
	IsTimerRunning           *Bool     `xmlrpc:"is_timer_running,omptempty"`
	TimerPause               *Time     `xmlrpc:"timer_pause,omptempty"`
	TimerStart               *Time     `xmlrpc:"timer_start,omptempty"`
	UserTimerId              *Relation `xmlrpc:"user_timer_id,omptempty"`
}

// TimerMixins represents array of timer.mixin model.
type TimerMixins []TimerMixin

// TimerMixinModel is the odoo model name.
const TimerMixinModel = "timer.mixin"

// Many2One convert TimerMixin to *Many2One.
func (tm *TimerMixin) Many2One() *Many2One {
	return NewMany2One(tm.Id.Get(), "")
}

// CreateTimerMixin creates a new timer.mixin model and returns its id.
func (c *Client) CreateTimerMixin(tm *TimerMixin) (int64, error) {
	ids, err := c.CreateTimerMixins([]*TimerMixin{tm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimerMixin creates a new timer.mixin model and returns its id.
func (c *Client) CreateTimerMixins(tms []*TimerMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range tms {
		vv = append(vv, v)
	}
	return c.Create(TimerMixinModel, vv)
}

// UpdateTimerMixin updates an existing timer.mixin record.
func (c *Client) UpdateTimerMixin(tm *TimerMixin) error {
	return c.UpdateTimerMixins([]int64{tm.Id.Get()}, tm)
}

// UpdateTimerMixins updates existing timer.mixin records.
// All records (represented by ids) will be updated by tm values.
func (c *Client) UpdateTimerMixins(ids []int64, tm *TimerMixin) error {
	return c.Update(TimerMixinModel, ids, tm)
}

// DeleteTimerMixin deletes an existing timer.mixin record.
func (c *Client) DeleteTimerMixin(id int64) error {
	return c.DeleteTimerMixins([]int64{id})
}

// DeleteTimerMixins deletes existing timer.mixin records.
func (c *Client) DeleteTimerMixins(ids []int64) error {
	return c.Delete(TimerMixinModel, ids)
}

// GetTimerMixin gets timer.mixin existing record.
func (c *Client) GetTimerMixin(id int64) (*TimerMixin, error) {
	tms, err := c.GetTimerMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if tms != nil && len(*tms) > 0 {
		return &((*tms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of timer.mixin not found", id)
}

// GetTimerMixins gets timer.mixin existing records.
func (c *Client) GetTimerMixins(ids []int64) (*TimerMixins, error) {
	tms := &TimerMixins{}
	if err := c.Read(TimerMixinModel, ids, nil, tms); err != nil {
		return nil, err
	}
	return tms, nil
}

// FindTimerMixin finds timer.mixin record by querying it with criteria.
func (c *Client) FindTimerMixin(criteria *Criteria) (*TimerMixin, error) {
	tms := &TimerMixins{}
	if err := c.SearchRead(TimerMixinModel, criteria, NewOptions().Limit(1), tms); err != nil {
		return nil, err
	}
	if tms != nil && len(*tms) > 0 {
		return &((*tms)[0]), nil
	}
	return nil, fmt.Errorf("timer.mixin was not found with criteria %v", criteria)
}

// FindTimerMixins finds timer.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimerMixins(criteria *Criteria, options *Options) (*TimerMixins, error) {
	tms := &TimerMixins{}
	if err := c.SearchRead(TimerMixinModel, criteria, options, tms); err != nil {
		return nil, err
	}
	return tms, nil
}

// FindTimerMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimerMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(TimerMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTimerMixinId finds record id by querying it with criteria.
func (c *Client) FindTimerMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimerMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("timer.mixin was not found with criteria %v and options %v", criteria, options)
}
