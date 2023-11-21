package odoo

import (
	"fmt"
)

// GamificationKarmaTracking represents gamification.karma.tracking model.
type GamificationKarmaTracking struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	Consolidated *Bool     `xmlrpc:"consolidated,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	NewValue     *Int      `xmlrpc:"new_value,omptempty"`
	OldValue     *Int      `xmlrpc:"old_value,omptempty"`
	TrackingDate *Time     `xmlrpc:"tracking_date,omptempty"`
	UserId       *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// GamificationKarmaTrackings represents array of gamification.karma.tracking model.
type GamificationKarmaTrackings []GamificationKarmaTracking

// GamificationKarmaTrackingModel is the odoo model name.
const GamificationKarmaTrackingModel = "gamification.karma.tracking"

// Many2One convert GamificationKarmaTracking to *Many2One.
func (gkt *GamificationKarmaTracking) Many2One() *Many2One {
	return NewMany2One(gkt.Id.Get(), "")
}

// CreateGamificationKarmaTracking creates a new gamification.karma.tracking model and returns its id.
func (c *Client) CreateGamificationKarmaTracking(gkt *GamificationKarmaTracking) (int64, error) {
	ids, err := c.CreateGamificationKarmaTrackings([]*GamificationKarmaTracking{gkt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGamificationKarmaTracking creates a new gamification.karma.tracking model and returns its id.
func (c *Client) CreateGamificationKarmaTrackings(gkts []*GamificationKarmaTracking) ([]int64, error) {
	var vv []interface{}
	for _, v := range gkts {
		vv = append(vv, v)
	}
	return c.Create(GamificationKarmaTrackingModel, vv)
}

// UpdateGamificationKarmaTracking updates an existing gamification.karma.tracking record.
func (c *Client) UpdateGamificationKarmaTracking(gkt *GamificationKarmaTracking) error {
	return c.UpdateGamificationKarmaTrackings([]int64{gkt.Id.Get()}, gkt)
}

// UpdateGamificationKarmaTrackings updates existing gamification.karma.tracking records.
// All records (represented by ids) will be updated by gkt values.
func (c *Client) UpdateGamificationKarmaTrackings(ids []int64, gkt *GamificationKarmaTracking) error {
	return c.Update(GamificationKarmaTrackingModel, ids, gkt)
}

// DeleteGamificationKarmaTracking deletes an existing gamification.karma.tracking record.
func (c *Client) DeleteGamificationKarmaTracking(id int64) error {
	return c.DeleteGamificationKarmaTrackings([]int64{id})
}

// DeleteGamificationKarmaTrackings deletes existing gamification.karma.tracking records.
func (c *Client) DeleteGamificationKarmaTrackings(ids []int64) error {
	return c.Delete(GamificationKarmaTrackingModel, ids)
}

// GetGamificationKarmaTracking gets gamification.karma.tracking existing record.
func (c *Client) GetGamificationKarmaTracking(id int64) (*GamificationKarmaTracking, error) {
	gkts, err := c.GetGamificationKarmaTrackings([]int64{id})
	if err != nil {
		return nil, err
	}
	if gkts != nil && len(*gkts) > 0 {
		return &((*gkts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of gamification.karma.tracking not found", id)
}

// GetGamificationKarmaTrackings gets gamification.karma.tracking existing records.
func (c *Client) GetGamificationKarmaTrackings(ids []int64) (*GamificationKarmaTrackings, error) {
	gkts := &GamificationKarmaTrackings{}
	if err := c.Read(GamificationKarmaTrackingModel, ids, nil, gkts); err != nil {
		return nil, err
	}
	return gkts, nil
}

// FindGamificationKarmaTracking finds gamification.karma.tracking record by querying it with criteria.
func (c *Client) FindGamificationKarmaTracking(criteria *Criteria) (*GamificationKarmaTracking, error) {
	gkts := &GamificationKarmaTrackings{}
	if err := c.SearchRead(GamificationKarmaTrackingModel, criteria, NewOptions().Limit(1), gkts); err != nil {
		return nil, err
	}
	if gkts != nil && len(*gkts) > 0 {
		return &((*gkts)[0]), nil
	}
	return nil, fmt.Errorf("gamification.karma.tracking was not found with criteria %v", criteria)
}

// FindGamificationKarmaTrackings finds gamification.karma.tracking records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationKarmaTrackings(criteria *Criteria, options *Options) (*GamificationKarmaTrackings, error) {
	gkts := &GamificationKarmaTrackings{}
	if err := c.SearchRead(GamificationKarmaTrackingModel, criteria, options, gkts); err != nil {
		return nil, err
	}
	return gkts, nil
}

// FindGamificationKarmaTrackingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGamificationKarmaTrackingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(GamificationKarmaTrackingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindGamificationKarmaTrackingId finds record id by querying it with criteria.
func (c *Client) FindGamificationKarmaTrackingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GamificationKarmaTrackingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("gamification.karma.tracking was not found with criteria %v and options %v", criteria, options)
}
