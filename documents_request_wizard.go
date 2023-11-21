package odoo

import (
	"fmt"
)

// DocumentsRequestWizard represents documents.request_wizard model.
type DocumentsRequestWizard struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadlineRange     *Int       `xmlrpc:"activity_date_deadline_range,omptempty"`
	ActivityDateDeadlineRangeType *Selection `xmlrpc:"activity_date_deadline_range_type,omptempty"`
	ActivityNote                  *String    `xmlrpc:"activity_note,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	FolderId                      *Many2One  `xmlrpc:"folder_id,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	OwnerId                       *Many2One  `xmlrpc:"owner_id,omptempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omptempty"`
	ResId                         *Int       `xmlrpc:"res_id,omptempty"`
	ResModel                      *String    `xmlrpc:"res_model,omptempty"`
	TagIds                        *Relation  `xmlrpc:"tag_ids,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// DocumentsRequestWizards represents array of documents.request_wizard model.
type DocumentsRequestWizards []DocumentsRequestWizard

// DocumentsRequestWizardModel is the odoo model name.
const DocumentsRequestWizardModel = "documents.request_wizard"

// Many2One convert DocumentsRequestWizard to *Many2One.
func (dr *DocumentsRequestWizard) Many2One() *Many2One {
	return NewMany2One(dr.Id.Get(), "")
}

// CreateDocumentsRequestWizard creates a new documents.request_wizard model and returns its id.
func (c *Client) CreateDocumentsRequestWizard(dr *DocumentsRequestWizard) (int64, error) {
	ids, err := c.CreateDocumentsRequestWizards([]*DocumentsRequestWizard{dr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsRequestWizard creates a new documents.request_wizard model and returns its id.
func (c *Client) CreateDocumentsRequestWizards(drs []*DocumentsRequestWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range drs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsRequestWizardModel, vv)
}

// UpdateDocumentsRequestWizard updates an existing documents.request_wizard record.
func (c *Client) UpdateDocumentsRequestWizard(dr *DocumentsRequestWizard) error {
	return c.UpdateDocumentsRequestWizards([]int64{dr.Id.Get()}, dr)
}

// UpdateDocumentsRequestWizards updates existing documents.request_wizard records.
// All records (represented by ids) will be updated by dr values.
func (c *Client) UpdateDocumentsRequestWizards(ids []int64, dr *DocumentsRequestWizard) error {
	return c.Update(DocumentsRequestWizardModel, ids, dr)
}

// DeleteDocumentsRequestWizard deletes an existing documents.request_wizard record.
func (c *Client) DeleteDocumentsRequestWizard(id int64) error {
	return c.DeleteDocumentsRequestWizards([]int64{id})
}

// DeleteDocumentsRequestWizards deletes existing documents.request_wizard records.
func (c *Client) DeleteDocumentsRequestWizards(ids []int64) error {
	return c.Delete(DocumentsRequestWizardModel, ids)
}

// GetDocumentsRequestWizard gets documents.request_wizard existing record.
func (c *Client) GetDocumentsRequestWizard(id int64) (*DocumentsRequestWizard, error) {
	drs, err := c.GetDocumentsRequestWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.request_wizard not found", id)
}

// GetDocumentsRequestWizards gets documents.request_wizard existing records.
func (c *Client) GetDocumentsRequestWizards(ids []int64) (*DocumentsRequestWizards, error) {
	drs := &DocumentsRequestWizards{}
	if err := c.Read(DocumentsRequestWizardModel, ids, nil, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDocumentsRequestWizard finds documents.request_wizard record by querying it with criteria.
func (c *Client) FindDocumentsRequestWizard(criteria *Criteria) (*DocumentsRequestWizard, error) {
	drs := &DocumentsRequestWizards{}
	if err := c.SearchRead(DocumentsRequestWizardModel, criteria, NewOptions().Limit(1), drs); err != nil {
		return nil, err
	}
	if drs != nil && len(*drs) > 0 {
		return &((*drs)[0]), nil
	}
	return nil, fmt.Errorf("documents.request_wizard was not found with criteria %v", criteria)
}

// FindDocumentsRequestWizards finds documents.request_wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsRequestWizards(criteria *Criteria, options *Options) (*DocumentsRequestWizards, error) {
	drs := &DocumentsRequestWizards{}
	if err := c.SearchRead(DocumentsRequestWizardModel, criteria, options, drs); err != nil {
		return nil, err
	}
	return drs, nil
}

// FindDocumentsRequestWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsRequestWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsRequestWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsRequestWizardId finds record id by querying it with criteria.
func (c *Client) FindDocumentsRequestWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsRequestWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.request_wizard was not found with criteria %v and options %v", criteria, options)
}
