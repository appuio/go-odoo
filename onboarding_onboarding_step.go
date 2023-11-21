package odoo

import (
	"fmt"
)

// OnboardingOnboardingStep represents onboarding.onboarding.step model.
type OnboardingOnboardingStep struct {
	LastUpdate              *Time      `xmlrpc:"__last_update,omptempty"`
	ButtonText              *String    `xmlrpc:"button_text,omptempty"`
	CreateDate              *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrentProgressStepId   *Many2One  `xmlrpc:"current_progress_step_id,omptempty"`
	CurrentStepState        *Selection `xmlrpc:"current_step_state,omptempty"`
	Description             *String    `xmlrpc:"description,omptempty"`
	DisplayName             *String    `xmlrpc:"display_name,omptempty"`
	DoneIcon                *String    `xmlrpc:"done_icon,omptempty"`
	DoneText                *String    `xmlrpc:"done_text,omptempty"`
	Id                      *Int       `xmlrpc:"id,omptempty"`
	OnboardingId            *Many2One  `xmlrpc:"onboarding_id,omptempty"`
	PanelStepOpenActionName *String    `xmlrpc:"panel_step_open_action_name,omptempty"`
	ProgressIds             *Relation  `xmlrpc:"progress_ids,omptempty"`
	Sequence                *Int       `xmlrpc:"sequence,omptempty"`
	Title                   *String    `xmlrpc:"title,omptempty"`
	WriteDate               *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// OnboardingOnboardingSteps represents array of onboarding.onboarding.step model.
type OnboardingOnboardingSteps []OnboardingOnboardingStep

// OnboardingOnboardingStepModel is the odoo model name.
const OnboardingOnboardingStepModel = "onboarding.onboarding.step"

// Many2One convert OnboardingOnboardingStep to *Many2One.
func (oos *OnboardingOnboardingStep) Many2One() *Many2One {
	return NewMany2One(oos.Id.Get(), "")
}

// CreateOnboardingOnboardingStep creates a new onboarding.onboarding.step model and returns its id.
func (c *Client) CreateOnboardingOnboardingStep(oos *OnboardingOnboardingStep) (int64, error) {
	ids, err := c.CreateOnboardingOnboardingSteps([]*OnboardingOnboardingStep{oos})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOnboardingOnboardingStep creates a new onboarding.onboarding.step model and returns its id.
func (c *Client) CreateOnboardingOnboardingSteps(ooss []*OnboardingOnboardingStep) ([]int64, error) {
	var vv []interface{}
	for _, v := range ooss {
		vv = append(vv, v)
	}
	return c.Create(OnboardingOnboardingStepModel, vv)
}

// UpdateOnboardingOnboardingStep updates an existing onboarding.onboarding.step record.
func (c *Client) UpdateOnboardingOnboardingStep(oos *OnboardingOnboardingStep) error {
	return c.UpdateOnboardingOnboardingSteps([]int64{oos.Id.Get()}, oos)
}

// UpdateOnboardingOnboardingSteps updates existing onboarding.onboarding.step records.
// All records (represented by ids) will be updated by oos values.
func (c *Client) UpdateOnboardingOnboardingSteps(ids []int64, oos *OnboardingOnboardingStep) error {
	return c.Update(OnboardingOnboardingStepModel, ids, oos)
}

// DeleteOnboardingOnboardingStep deletes an existing onboarding.onboarding.step record.
func (c *Client) DeleteOnboardingOnboardingStep(id int64) error {
	return c.DeleteOnboardingOnboardingSteps([]int64{id})
}

// DeleteOnboardingOnboardingSteps deletes existing onboarding.onboarding.step records.
func (c *Client) DeleteOnboardingOnboardingSteps(ids []int64) error {
	return c.Delete(OnboardingOnboardingStepModel, ids)
}

// GetOnboardingOnboardingStep gets onboarding.onboarding.step existing record.
func (c *Client) GetOnboardingOnboardingStep(id int64) (*OnboardingOnboardingStep, error) {
	ooss, err := c.GetOnboardingOnboardingSteps([]int64{id})
	if err != nil {
		return nil, err
	}
	if ooss != nil && len(*ooss) > 0 {
		return &((*ooss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of onboarding.onboarding.step not found", id)
}

// GetOnboardingOnboardingSteps gets onboarding.onboarding.step existing records.
func (c *Client) GetOnboardingOnboardingSteps(ids []int64) (*OnboardingOnboardingSteps, error) {
	ooss := &OnboardingOnboardingSteps{}
	if err := c.Read(OnboardingOnboardingStepModel, ids, nil, ooss); err != nil {
		return nil, err
	}
	return ooss, nil
}

// FindOnboardingOnboardingStep finds onboarding.onboarding.step record by querying it with criteria.
func (c *Client) FindOnboardingOnboardingStep(criteria *Criteria) (*OnboardingOnboardingStep, error) {
	ooss := &OnboardingOnboardingSteps{}
	if err := c.SearchRead(OnboardingOnboardingStepModel, criteria, NewOptions().Limit(1), ooss); err != nil {
		return nil, err
	}
	if ooss != nil && len(*ooss) > 0 {
		return &((*ooss)[0]), nil
	}
	return nil, fmt.Errorf("onboarding.onboarding.step was not found with criteria %v", criteria)
}

// FindOnboardingOnboardingSteps finds onboarding.onboarding.step records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingOnboardingSteps(criteria *Criteria, options *Options) (*OnboardingOnboardingSteps, error) {
	ooss := &OnboardingOnboardingSteps{}
	if err := c.SearchRead(OnboardingOnboardingStepModel, criteria, options, ooss); err != nil {
		return nil, err
	}
	return ooss, nil
}

// FindOnboardingOnboardingStepIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingOnboardingStepIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OnboardingOnboardingStepModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOnboardingOnboardingStepId finds record id by querying it with criteria.
func (c *Client) FindOnboardingOnboardingStepId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OnboardingOnboardingStepModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("onboarding.onboarding.step was not found with criteria %v and options %v", criteria, options)
}
