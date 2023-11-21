package odoo

import (
	"fmt"
)

// ProjectShareWizard represents project.share.wizard model.
type ProjectShareWizard struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	AccessMode        *Selection `xmlrpc:"access_mode,omptempty"`
	AccessWarning     *String    `xmlrpc:"access_warning,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayAccessMode *Bool      `xmlrpc:"display_access_mode,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	Note              *String    `xmlrpc:"note,omptempty"`
	PartnerIds        *Relation  `xmlrpc:"partner_ids,omptempty"`
	ResId             *Int       `xmlrpc:"res_id,omptempty"`
	ResModel          *String    `xmlrpc:"res_model,omptempty"`
	ResourceRef       *String    `xmlrpc:"resource_ref,omptempty"`
	ShareLink         *String    `xmlrpc:"share_link,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProjectShareWizards represents array of project.share.wizard model.
type ProjectShareWizards []ProjectShareWizard

// ProjectShareWizardModel is the odoo model name.
const ProjectShareWizardModel = "project.share.wizard"

// Many2One convert ProjectShareWizard to *Many2One.
func (psw *ProjectShareWizard) Many2One() *Many2One {
	return NewMany2One(psw.Id.Get(), "")
}

// CreateProjectShareWizard creates a new project.share.wizard model and returns its id.
func (c *Client) CreateProjectShareWizard(psw *ProjectShareWizard) (int64, error) {
	ids, err := c.CreateProjectShareWizards([]*ProjectShareWizard{psw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectShareWizard creates a new project.share.wizard model and returns its id.
func (c *Client) CreateProjectShareWizards(psws []*ProjectShareWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range psws {
		vv = append(vv, v)
	}
	return c.Create(ProjectShareWizardModel, vv)
}

// UpdateProjectShareWizard updates an existing project.share.wizard record.
func (c *Client) UpdateProjectShareWizard(psw *ProjectShareWizard) error {
	return c.UpdateProjectShareWizards([]int64{psw.Id.Get()}, psw)
}

// UpdateProjectShareWizards updates existing project.share.wizard records.
// All records (represented by ids) will be updated by psw values.
func (c *Client) UpdateProjectShareWizards(ids []int64, psw *ProjectShareWizard) error {
	return c.Update(ProjectShareWizardModel, ids, psw)
}

// DeleteProjectShareWizard deletes an existing project.share.wizard record.
func (c *Client) DeleteProjectShareWizard(id int64) error {
	return c.DeleteProjectShareWizards([]int64{id})
}

// DeleteProjectShareWizards deletes existing project.share.wizard records.
func (c *Client) DeleteProjectShareWizards(ids []int64) error {
	return c.Delete(ProjectShareWizardModel, ids)
}

// GetProjectShareWizard gets project.share.wizard existing record.
func (c *Client) GetProjectShareWizard(id int64) (*ProjectShareWizard, error) {
	psws, err := c.GetProjectShareWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if psws != nil && len(*psws) > 0 {
		return &((*psws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of project.share.wizard not found", id)
}

// GetProjectShareWizards gets project.share.wizard existing records.
func (c *Client) GetProjectShareWizards(ids []int64) (*ProjectShareWizards, error) {
	psws := &ProjectShareWizards{}
	if err := c.Read(ProjectShareWizardModel, ids, nil, psws); err != nil {
		return nil, err
	}
	return psws, nil
}

// FindProjectShareWizard finds project.share.wizard record by querying it with criteria.
func (c *Client) FindProjectShareWizard(criteria *Criteria) (*ProjectShareWizard, error) {
	psws := &ProjectShareWizards{}
	if err := c.SearchRead(ProjectShareWizardModel, criteria, NewOptions().Limit(1), psws); err != nil {
		return nil, err
	}
	if psws != nil && len(*psws) > 0 {
		return &((*psws)[0]), nil
	}
	return nil, fmt.Errorf("project.share.wizard was not found with criteria %v", criteria)
}

// FindProjectShareWizards finds project.share.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectShareWizards(criteria *Criteria, options *Options) (*ProjectShareWizards, error) {
	psws := &ProjectShareWizards{}
	if err := c.SearchRead(ProjectShareWizardModel, criteria, options, psws); err != nil {
		return nil, err
	}
	return psws, nil
}

// FindProjectShareWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectShareWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProjectShareWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProjectShareWizardId finds record id by querying it with criteria.
func (c *Client) FindProjectShareWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectShareWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("project.share.wizard was not found with criteria %v and options %v", criteria, options)
}
