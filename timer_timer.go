package odoo

import (
	"fmt"
)

// TimerTimer represents timer.timer model.
type TimerTimer struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	IsTimerRunning *Bool     `xmlrpc:"is_timer_running,omptempty"`
	ResId          *Int      `xmlrpc:"res_id,omptempty"`
	ResModel       *String   `xmlrpc:"res_model,omptempty"`
	TimerPause     *Time     `xmlrpc:"timer_pause,omptempty"`
	TimerStart     *Time     `xmlrpc:"timer_start,omptempty"`
	UserId         *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// TimerTimers represents array of timer.timer model.
type TimerTimers []TimerTimer

// TimerTimerModel is the odoo model name.
const TimerTimerModel = "timer.timer"

// Many2One convert TimerTimer to *Many2One.
func (tt *TimerTimer) Many2One() *Many2One {
	return NewMany2One(tt.Id.Get(), "")
}

// CreateTimerTimer creates a new timer.timer model and returns its id.
func (c *Client) CreateTimerTimer(tt *TimerTimer) (int64, error) {
	ids, err := c.CreateTimerTimers([]*TimerTimer{tt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTimerTimer creates a new timer.timer model and returns its id.
func (c *Client) CreateTimerTimers(tts []*TimerTimer) ([]int64, error) {
	var vv []interface{}
	for _, v := range tts {
		vv = append(vv, v)
	}
	return c.Create(TimerTimerModel, vv)
}

// UpdateTimerTimer updates an existing timer.timer record.
func (c *Client) UpdateTimerTimer(tt *TimerTimer) error {
	return c.UpdateTimerTimers([]int64{tt.Id.Get()}, tt)
}

// UpdateTimerTimers updates existing timer.timer records.
// All records (represented by ids) will be updated by tt values.
func (c *Client) UpdateTimerTimers(ids []int64, tt *TimerTimer) error {
	return c.Update(TimerTimerModel, ids, tt)
}

// DeleteTimerTimer deletes an existing timer.timer record.
func (c *Client) DeleteTimerTimer(id int64) error {
	return c.DeleteTimerTimers([]int64{id})
}

// DeleteTimerTimers deletes existing timer.timer records.
func (c *Client) DeleteTimerTimers(ids []int64) error {
	return c.Delete(TimerTimerModel, ids)
}

// GetTimerTimer gets timer.timer existing record.
func (c *Client) GetTimerTimer(id int64) (*TimerTimer, error) {
	tts, err := c.GetTimerTimers([]int64{id})
	if err != nil {
		return nil, err
	}
	if tts != nil && len(*tts) > 0 {
		return &((*tts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of timer.timer not found", id)
}

// GetTimerTimers gets timer.timer existing records.
func (c *Client) GetTimerTimers(ids []int64) (*TimerTimers, error) {
	tts := &TimerTimers{}
	if err := c.Read(TimerTimerModel, ids, nil, tts); err != nil {
		return nil, err
	}
	return tts, nil
}

// FindTimerTimer finds timer.timer record by querying it with criteria.
func (c *Client) FindTimerTimer(criteria *Criteria) (*TimerTimer, error) {
	tts := &TimerTimers{}
	if err := c.SearchRead(TimerTimerModel, criteria, NewOptions().Limit(1), tts); err != nil {
		return nil, err
	}
	if tts != nil && len(*tts) > 0 {
		return &((*tts)[0]), nil
	}
	return nil, fmt.Errorf("timer.timer was not found with criteria %v", criteria)
}

// FindTimerTimers finds timer.timer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimerTimers(criteria *Criteria, options *Options) (*TimerTimers, error) {
	tts := &TimerTimers{}
	if err := c.SearchRead(TimerTimerModel, criteria, options, tts); err != nil {
		return nil, err
	}
	return tts, nil
}

// FindTimerTimerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTimerTimerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(TimerTimerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindTimerTimerId finds record id by querying it with criteria.
func (c *Client) FindTimerTimerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TimerTimerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("timer.timer was not found with criteria %v and options %v", criteria, options)
}
