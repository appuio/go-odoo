package odoo

import (
	"fmt"
)

// OnboardingProgress represents onboarding.progress model.
type OnboardingProgress struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId          *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	IsOnboardingClosed *Bool      `xmlrpc:"is_onboarding_closed,omptempty"`
	OnboardingId       *Many2One  `xmlrpc:"onboarding_id,omptempty"`
	OnboardingState    *Selection `xmlrpc:"onboarding_state,omptempty"`
	ProgressStepIds    *Relation  `xmlrpc:"progress_step_ids,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// OnboardingProgresss represents array of onboarding.progress model.
type OnboardingProgresss []OnboardingProgress

// OnboardingProgressModel is the odoo model name.
const OnboardingProgressModel = "onboarding.progress"

// Many2One convert OnboardingProgress to *Many2One.
func (op *OnboardingProgress) Many2One() *Many2One {
	return NewMany2One(op.Id.Get(), "")
}

// CreateOnboardingProgress creates a new onboarding.progress model and returns its id.
func (c *Client) CreateOnboardingProgress(op *OnboardingProgress) (int64, error) {
	ids, err := c.CreateOnboardingProgresss([]*OnboardingProgress{op})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateOnboardingProgress creates a new onboarding.progress model and returns its id.
func (c *Client) CreateOnboardingProgresss(ops []*OnboardingProgress) ([]int64, error) {
	var vv []interface{}
	for _, v := range ops {
		vv = append(vv, v)
	}
	return c.Create(OnboardingProgressModel, vv)
}

// UpdateOnboardingProgress updates an existing onboarding.progress record.
func (c *Client) UpdateOnboardingProgress(op *OnboardingProgress) error {
	return c.UpdateOnboardingProgresss([]int64{op.Id.Get()}, op)
}

// UpdateOnboardingProgresss updates existing onboarding.progress records.
// All records (represented by ids) will be updated by op values.
func (c *Client) UpdateOnboardingProgresss(ids []int64, op *OnboardingProgress) error {
	return c.Update(OnboardingProgressModel, ids, op)
}

// DeleteOnboardingProgress deletes an existing onboarding.progress record.
func (c *Client) DeleteOnboardingProgress(id int64) error {
	return c.DeleteOnboardingProgresss([]int64{id})
}

// DeleteOnboardingProgresss deletes existing onboarding.progress records.
func (c *Client) DeleteOnboardingProgresss(ids []int64) error {
	return c.Delete(OnboardingProgressModel, ids)
}

// GetOnboardingProgress gets onboarding.progress existing record.
func (c *Client) GetOnboardingProgress(id int64) (*OnboardingProgress, error) {
	ops, err := c.GetOnboardingProgresss([]int64{id})
	if err != nil {
		return nil, err
	}
	if ops != nil && len(*ops) > 0 {
		return &((*ops)[0]), nil
	}
	return nil, fmt.Errorf("id %v of onboarding.progress not found", id)
}

// GetOnboardingProgresss gets onboarding.progress existing records.
func (c *Client) GetOnboardingProgresss(ids []int64) (*OnboardingProgresss, error) {
	ops := &OnboardingProgresss{}
	if err := c.Read(OnboardingProgressModel, ids, nil, ops); err != nil {
		return nil, err
	}
	return ops, nil
}

// FindOnboardingProgress finds onboarding.progress record by querying it with criteria.
func (c *Client) FindOnboardingProgress(criteria *Criteria) (*OnboardingProgress, error) {
	ops := &OnboardingProgresss{}
	if err := c.SearchRead(OnboardingProgressModel, criteria, NewOptions().Limit(1), ops); err != nil {
		return nil, err
	}
	if ops != nil && len(*ops) > 0 {
		return &((*ops)[0]), nil
	}
	return nil, fmt.Errorf("onboarding.progress was not found with criteria %v", criteria)
}

// FindOnboardingProgresss finds onboarding.progress records by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingProgresss(criteria *Criteria, options *Options) (*OnboardingProgresss, error) {
	ops := &OnboardingProgresss{}
	if err := c.SearchRead(OnboardingProgressModel, criteria, options, ops); err != nil {
		return nil, err
	}
	return ops, nil
}

// FindOnboardingProgressIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindOnboardingProgressIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(OnboardingProgressModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindOnboardingProgressId finds record id by querying it with criteria.
func (c *Client) FindOnboardingProgressId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(OnboardingProgressModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("onboarding.progress was not found with criteria %v and options %v", criteria, options)
}
