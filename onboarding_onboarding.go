package odoo

import (
	"fmt"
)

// OnboardingOnboarding represents onboarding.onboarding model.
type OnboardingOnboarding struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrentOnboardingState *Selection `xmlrpc:"current_onboarding_state,omptempty"`
	CurrentProgressId      *Many2One  `xmlrpc:"current_progress_id,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	IsOnboardingClosed     *Bool      `xmlrpc:"is_onboarding_closed,omptempty"`
	IsPerCompany           *Bool      `xmlrpc:"is_per_company,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	PanelBackgroundColor   *Selection `xmlrpc:"panel_background_color,omptempty"`
	PanelBackgroundImage   *String    `xmlrpc:"panel_background_image,omptempty"`
	PanelCloseActionName   *String    `xmlrpc:"panel_close_action_name,omptempty"`
	ProgressIds            *Relation  `xmlrpc:"progress_ids,omptempty"`
	RouteName              *String    `xmlrpc:"route_name,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	StepIds                *Relation  `xmlrpc:"step_ids,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// OnboardingOnboardings represents array of onboarding.onboarding model.
type OnboardingOnboardings []OnboardingOnboarding

// OnboardingOnboardingModel is the odoo model name.
const OnboardingOnboardingModel = "onboarding.onboarding"

// Many2One convert OnboardingOnboarding to *Many2One.
func (oo *OnboardingOnboarding) Many2One() *Many2One {
	return NewMany2One(oo.Id.Get(), "")
}

// CreateOnboardingOnboarding creates a new onboarding.onboarding model and returns its id.
func (c *Client) CreateOnboardingOnboarding(oo *OnboardingOnboarding) (int64, error) {
	ids, err := c.CreateOnboardingOnboardings([]*OnboardingOnboarding{oo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOnboardingOnboarding creates a new onboarding.onboarding model and returns its id.
func (c *Client) CreateOnboardingOnboardings(oos []*OnboardingOnboarding) ([]int64, error) {
	var vv []interface{}
	for _, v := range oos {
		vv = append(vv, v)
	}
	return c.Create(OnboardingOnboardingModel, vv)
}

// UpdateOnboardingOnboarding updates an existing onboarding.onboarding record.
func (c *Client) UpdateOnboardingOnboarding(oo *OnboardingOnboarding) error {
	return c.UpdateOnboardingOnboardings([]int64{oo.Id.Get()}, oo)
}

// UpdateOnboardingOnboardings updates existing onboarding.onboarding records.
// All records (represented by ids) will be updated by oo values.
func (c *Client) UpdateOnboardingOnboardings(ids []int64, oo *OnboardingOnboarding) error {
	return c.Update(OnboardingOnboardingModel, ids, oo)
}

// DeleteOnboardingOnboarding deletes an existing onboarding.onboarding record.
func (c *Client) DeleteOnboardingOnboarding(id int64) error {
	return c.DeleteOnboardingOnboardings([]int64{id})
}

// DeleteOnboardingOnboardings deletes existing onboarding.onboarding records.
func (c *Client) DeleteOnboardingOnboardings(ids []int64) error {
	return c.Delete(OnboardingOnboardingModel, ids)
}

// GetOnboardingOnboarding gets onboarding.onboarding existing record.
func (c *Client) GetOnboardingOnboarding(id int64) (*OnboardingOnboarding, error) {
	oos, err := c.GetOnboardingOnboardings([]int64{id})
	if err != nil {
		return nil, err
	}
	if oos != nil && len(*oos) > 0 {
		return &((*oos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of onboarding.onboarding not found", id)
}

// GetOnboardingOnboardings gets onboarding.onboarding existing records.
func (c *Client) GetOnboardingOnboardings(ids []int64) (*OnboardingOnboardings, error) {
	oos := &OnboardingOnboardings{}
	if err := c.Read(OnboardingOnboardingModel, ids, nil, oos); err != nil {
		return nil, err
	}
	return oos, nil
}

// FindOnboardingOnboarding finds onboarding.onboarding record by querying it with criteria.
func (c *Client) FindOnboardingOnboarding(criteria *Criteria) (*OnboardingOnboarding, error) {
	oos := &OnboardingOnboardings{}
	if err := c.SearchRead(OnboardingOnboardingModel, criteria, NewOptions().Limit(1), oos); err != nil {
		return nil, err
	}
	if oos != nil && len(*oos) > 0 {
		return &((*oos)[0]), nil
	}
	return nil, fmt.Errorf("onboarding.onboarding was not found with criteria %v", criteria)
}

// FindOnboardingOnboardings finds onboarding.onboarding records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingOnboardings(criteria *Criteria, options *Options) (*OnboardingOnboardings, error) {
	oos := &OnboardingOnboardings{}
	if err := c.SearchRead(OnboardingOnboardingModel, criteria, options, oos); err != nil {
		return nil, err
	}
	return oos, nil
}

// FindOnboardingOnboardingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingOnboardingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OnboardingOnboardingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOnboardingOnboardingId finds record id by querying it with criteria.
func (c *Client) FindOnboardingOnboardingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OnboardingOnboardingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("onboarding.onboarding was not found with criteria %v and options %v", criteria, options)
}
