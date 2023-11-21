package odoo

import (
	"fmt"
)

// OnboardingProgressStep represents onboarding.progress.step model.
type OnboardingProgressStep struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String    `xmlrpc:"display_name,omptempty"`
	Id           *Int       `xmlrpc:"id,omptempty"`
	OnboardingId *Many2One  `xmlrpc:"onboarding_id,omptempty"`
	ProgressId   *Many2One  `xmlrpc:"progress_id,omptempty"`
	StepId       *Many2One  `xmlrpc:"step_id,omptempty"`
	StepState    *Selection `xmlrpc:"step_state,omptempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// OnboardingProgressSteps represents array of onboarding.progress.step model.
type OnboardingProgressSteps []OnboardingProgressStep

// OnboardingProgressStepModel is the odoo model name.
const OnboardingProgressStepModel = "onboarding.progress.step"

// Many2One convert OnboardingProgressStep to *Many2One.
func (ops *OnboardingProgressStep) Many2One() *Many2One {
	return NewMany2One(ops.Id.Get(), "")
}

// CreateOnboardingProgressStep creates a new onboarding.progress.step model and returns its id.
func (c *Client) CreateOnboardingProgressStep(ops *OnboardingProgressStep) (int64, error) {
	ids, err := c.CreateOnboardingProgressSteps([]*OnboardingProgressStep{ops})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOnboardingProgressStep creates a new onboarding.progress.step model and returns its id.
func (c *Client) CreateOnboardingProgressSteps(opss []*OnboardingProgressStep) ([]int64, error) {
	var vv []interface{}
	for _, v := range opss {
		vv = append(vv, v)
	}
	return c.Create(OnboardingProgressStepModel, vv)
}

// UpdateOnboardingProgressStep updates an existing onboarding.progress.step record.
func (c *Client) UpdateOnboardingProgressStep(ops *OnboardingProgressStep) error {
	return c.UpdateOnboardingProgressSteps([]int64{ops.Id.Get()}, ops)
}

// UpdateOnboardingProgressSteps updates existing onboarding.progress.step records.
// All records (represented by ids) will be updated by ops values.
func (c *Client) UpdateOnboardingProgressSteps(ids []int64, ops *OnboardingProgressStep) error {
	return c.Update(OnboardingProgressStepModel, ids, ops)
}

// DeleteOnboardingProgressStep deletes an existing onboarding.progress.step record.
func (c *Client) DeleteOnboardingProgressStep(id int64) error {
	return c.DeleteOnboardingProgressSteps([]int64{id})
}

// DeleteOnboardingProgressSteps deletes existing onboarding.progress.step records.
func (c *Client) DeleteOnboardingProgressSteps(ids []int64) error {
	return c.Delete(OnboardingProgressStepModel, ids)
}

// GetOnboardingProgressStep gets onboarding.progress.step existing record.
func (c *Client) GetOnboardingProgressStep(id int64) (*OnboardingProgressStep, error) {
	opss, err := c.GetOnboardingProgressSteps([]int64{id})
	if err != nil {
		return nil, err
	}
	if opss != nil && len(*opss) > 0 {
		return &((*opss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of onboarding.progress.step not found", id)
}

// GetOnboardingProgressSteps gets onboarding.progress.step existing records.
func (c *Client) GetOnboardingProgressSteps(ids []int64) (*OnboardingProgressSteps, error) {
	opss := &OnboardingProgressSteps{}
	if err := c.Read(OnboardingProgressStepModel, ids, nil, opss); err != nil {
		return nil, err
	}
	return opss, nil
}

// FindOnboardingProgressStep finds onboarding.progress.step record by querying it with criteria.
func (c *Client) FindOnboardingProgressStep(criteria *Criteria) (*OnboardingProgressStep, error) {
	opss := &OnboardingProgressSteps{}
	if err := c.SearchRead(OnboardingProgressStepModel, criteria, NewOptions().Limit(1), opss); err != nil {
		return nil, err
	}
	if opss != nil && len(*opss) > 0 {
		return &((*opss)[0]), nil
	}
	return nil, fmt.Errorf("onboarding.progress.step was not found with criteria %v", criteria)
}

// FindOnboardingProgressSteps finds onboarding.progress.step records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingProgressSteps(criteria *Criteria, options *Options) (*OnboardingProgressSteps, error) {
	opss := &OnboardingProgressSteps{}
	if err := c.SearchRead(OnboardingProgressStepModel, criteria, options, opss); err != nil {
		return nil, err
	}
	return opss, nil
}

// FindOnboardingProgressStepIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingProgressStepIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OnboardingProgressStepModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOnboardingProgressStepId finds record id by querying it with criteria.
func (c *Client) FindOnboardingProgressStepId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OnboardingProgressStepModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("onboarding.progress.step was not found with criteria %v and options %v", criteria, options)
}
