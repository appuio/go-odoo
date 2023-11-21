package odoo

import (
	"fmt"
)

// DocumentsFolderDeletionWizard represents documents.folder.deletion.wizard model.
type DocumentsFolderDeletionWizard struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	FolderId       *Many2One `xmlrpc:"folder_id,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	ParentFolderId *Many2One `xmlrpc:"parent_folder_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DocumentsFolderDeletionWizards represents array of documents.folder.deletion.wizard model.
type DocumentsFolderDeletionWizards []DocumentsFolderDeletionWizard

// DocumentsFolderDeletionWizardModel is the odoo model name.
const DocumentsFolderDeletionWizardModel = "documents.folder.deletion.wizard"

// Many2One convert DocumentsFolderDeletionWizard to *Many2One.
func (dfdw *DocumentsFolderDeletionWizard) Many2One() *Many2One {
	return NewMany2One(dfdw.Id.Get(), "")
}

// CreateDocumentsFolderDeletionWizard creates a new documents.folder.deletion.wizard model and returns its id.
func (c *Client) CreateDocumentsFolderDeletionWizard(dfdw *DocumentsFolderDeletionWizard) (int64, error) {
	ids, err := c.CreateDocumentsFolderDeletionWizards([]*DocumentsFolderDeletionWizard{dfdw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsFolderDeletionWizard creates a new documents.folder.deletion.wizard model and returns its id.
func (c *Client) CreateDocumentsFolderDeletionWizards(dfdws []*DocumentsFolderDeletionWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range dfdws {
		vv = append(vv, v)
	}
	return c.Create(DocumentsFolderDeletionWizardModel, vv)
}

// UpdateDocumentsFolderDeletionWizard updates an existing documents.folder.deletion.wizard record.
func (c *Client) UpdateDocumentsFolderDeletionWizard(dfdw *DocumentsFolderDeletionWizard) error {
	return c.UpdateDocumentsFolderDeletionWizards([]int64{dfdw.Id.Get()}, dfdw)
}

// UpdateDocumentsFolderDeletionWizards updates existing documents.folder.deletion.wizard records.
// All records (represented by ids) will be updated by dfdw values.
func (c *Client) UpdateDocumentsFolderDeletionWizards(ids []int64, dfdw *DocumentsFolderDeletionWizard) error {
	return c.Update(DocumentsFolderDeletionWizardModel, ids, dfdw)
}

// DeleteDocumentsFolderDeletionWizard deletes an existing documents.folder.deletion.wizard record.
func (c *Client) DeleteDocumentsFolderDeletionWizard(id int64) error {
	return c.DeleteDocumentsFolderDeletionWizards([]int64{id})
}

// DeleteDocumentsFolderDeletionWizards deletes existing documents.folder.deletion.wizard records.
func (c *Client) DeleteDocumentsFolderDeletionWizards(ids []int64) error {
	return c.Delete(DocumentsFolderDeletionWizardModel, ids)
}

// GetDocumentsFolderDeletionWizard gets documents.folder.deletion.wizard existing record.
func (c *Client) GetDocumentsFolderDeletionWizard(id int64) (*DocumentsFolderDeletionWizard, error) {
	dfdws, err := c.GetDocumentsFolderDeletionWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if dfdws != nil && len(*dfdws) > 0 {
		return &((*dfdws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.folder.deletion.wizard not found", id)
}

// GetDocumentsFolderDeletionWizards gets documents.folder.deletion.wizard existing records.
func (c *Client) GetDocumentsFolderDeletionWizards(ids []int64) (*DocumentsFolderDeletionWizards, error) {
	dfdws := &DocumentsFolderDeletionWizards{}
	if err := c.Read(DocumentsFolderDeletionWizardModel, ids, nil, dfdws); err != nil {
		return nil, err
	}
	return dfdws, nil
}

// FindDocumentsFolderDeletionWizard finds documents.folder.deletion.wizard record by querying it with criteria.
func (c *Client) FindDocumentsFolderDeletionWizard(criteria *Criteria) (*DocumentsFolderDeletionWizard, error) {
	dfdws := &DocumentsFolderDeletionWizards{}
	if err := c.SearchRead(DocumentsFolderDeletionWizardModel, criteria, NewOptions().Limit(1), dfdws); err != nil {
		return nil, err
	}
	if dfdws != nil && len(*dfdws) > 0 {
		return &((*dfdws)[0]), nil
	}
	return nil, fmt.Errorf("documents.folder.deletion.wizard was not found with criteria %v", criteria)
}

// FindDocumentsFolderDeletionWizards finds documents.folder.deletion.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFolderDeletionWizards(criteria *Criteria, options *Options) (*DocumentsFolderDeletionWizards, error) {
	dfdws := &DocumentsFolderDeletionWizards{}
	if err := c.SearchRead(DocumentsFolderDeletionWizardModel, criteria, options, dfdws); err != nil {
		return nil, err
	}
	return dfdws, nil
}

// FindDocumentsFolderDeletionWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFolderDeletionWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsFolderDeletionWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsFolderDeletionWizardId finds record id by querying it with criteria.
func (c *Client) FindDocumentsFolderDeletionWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsFolderDeletionWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.folder.deletion.wizard was not found with criteria %v and options %v", criteria, options)
}
