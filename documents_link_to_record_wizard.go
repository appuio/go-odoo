package odoo

import (
	"fmt"
)

// DocumentsLinkToRecordWizard represents documents.link_to_record_wizard model.
type DocumentsLinkToRecordWizard struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	AccessibleModelIds *Relation `xmlrpc:"accessible_model_ids,omptempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	DocumentIds        *Relation `xmlrpc:"document_ids,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	IsReadonlyModel    *Bool     `xmlrpc:"is_readonly_model,omptempty"`
	ModelId            *Many2One `xmlrpc:"model_id,omptempty"`
	ResourceRef        *String   `xmlrpc:"resource_ref,omptempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DocumentsLinkToRecordWizards represents array of documents.link_to_record_wizard model.
type DocumentsLinkToRecordWizards []DocumentsLinkToRecordWizard

// DocumentsLinkToRecordWizardModel is the odoo model name.
const DocumentsLinkToRecordWizardModel = "documents.link_to_record_wizard"

// Many2One convert DocumentsLinkToRecordWizard to *Many2One.
func (dl *DocumentsLinkToRecordWizard) Many2One() *Many2One {
	return NewMany2One(dl.Id.Get(), "")
}

// CreateDocumentsLinkToRecordWizard creates a new documents.link_to_record_wizard model and returns its id.
func (c *Client) CreateDocumentsLinkToRecordWizard(dl *DocumentsLinkToRecordWizard) (int64, error) {
	ids, err := c.CreateDocumentsLinkToRecordWizards([]*DocumentsLinkToRecordWizard{dl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsLinkToRecordWizard creates a new documents.link_to_record_wizard model and returns its id.
func (c *Client) CreateDocumentsLinkToRecordWizards(dls []*DocumentsLinkToRecordWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range dls {
		vv = append(vv, v)
	}
	return c.Create(DocumentsLinkToRecordWizardModel, vv)
}

// UpdateDocumentsLinkToRecordWizard updates an existing documents.link_to_record_wizard record.
func (c *Client) UpdateDocumentsLinkToRecordWizard(dl *DocumentsLinkToRecordWizard) error {
	return c.UpdateDocumentsLinkToRecordWizards([]int64{dl.Id.Get()}, dl)
}

// UpdateDocumentsLinkToRecordWizards updates existing documents.link_to_record_wizard records.
// All records (represented by ids) will be updated by dl values.
func (c *Client) UpdateDocumentsLinkToRecordWizards(ids []int64, dl *DocumentsLinkToRecordWizard) error {
	return c.Update(DocumentsLinkToRecordWizardModel, ids, dl)
}

// DeleteDocumentsLinkToRecordWizard deletes an existing documents.link_to_record_wizard record.
func (c *Client) DeleteDocumentsLinkToRecordWizard(id int64) error {
	return c.DeleteDocumentsLinkToRecordWizards([]int64{id})
}

// DeleteDocumentsLinkToRecordWizards deletes existing documents.link_to_record_wizard records.
func (c *Client) DeleteDocumentsLinkToRecordWizards(ids []int64) error {
	return c.Delete(DocumentsLinkToRecordWizardModel, ids)
}

// GetDocumentsLinkToRecordWizard gets documents.link_to_record_wizard existing record.
func (c *Client) GetDocumentsLinkToRecordWizard(id int64) (*DocumentsLinkToRecordWizard, error) {
	dls, err := c.GetDocumentsLinkToRecordWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if dls != nil && len(*dls) > 0 {
		return &((*dls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.link_to_record_wizard not found", id)
}

// GetDocumentsLinkToRecordWizards gets documents.link_to_record_wizard existing records.
func (c *Client) GetDocumentsLinkToRecordWizards(ids []int64) (*DocumentsLinkToRecordWizards, error) {
	dls := &DocumentsLinkToRecordWizards{}
	if err := c.Read(DocumentsLinkToRecordWizardModel, ids, nil, dls); err != nil {
		return nil, err
	}
	return dls, nil
}

// FindDocumentsLinkToRecordWizard finds documents.link_to_record_wizard record by querying it with criteria.
func (c *Client) FindDocumentsLinkToRecordWizard(criteria *Criteria) (*DocumentsLinkToRecordWizard, error) {
	dls := &DocumentsLinkToRecordWizards{}
	if err := c.SearchRead(DocumentsLinkToRecordWizardModel, criteria, NewOptions().Limit(1), dls); err != nil {
		return nil, err
	}
	if dls != nil && len(*dls) > 0 {
		return &((*dls)[0]), nil
	}
	return nil, fmt.Errorf("documents.link_to_record_wizard was not found with criteria %v", criteria)
}

// FindDocumentsLinkToRecordWizards finds documents.link_to_record_wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsLinkToRecordWizards(criteria *Criteria, options *Options) (*DocumentsLinkToRecordWizards, error) {
	dls := &DocumentsLinkToRecordWizards{}
	if err := c.SearchRead(DocumentsLinkToRecordWizardModel, criteria, options, dls); err != nil {
		return nil, err
	}
	return dls, nil
}

// FindDocumentsLinkToRecordWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsLinkToRecordWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsLinkToRecordWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsLinkToRecordWizardId finds record id by querying it with criteria.
func (c *Client) FindDocumentsLinkToRecordWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsLinkToRecordWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.link_to_record_wizard was not found with criteria %v and options %v", criteria, options)
}
