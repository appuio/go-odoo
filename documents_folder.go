package odoo

import (
	"fmt"
)

// DocumentsFolder represents documents.folder model.
type DocumentsFolder struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	ActionCount       *Int      `xmlrpc:"action_count,omptempty"`
	ChildrenFolderIds *Relation `xmlrpc:"children_folder_ids,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	Description       *String   `xmlrpc:"description,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	DocumentCount     *Int      `xmlrpc:"document_count,omptempty"`
	DocumentIds       *Relation `xmlrpc:"document_ids,omptempty"`
	FacetIds          *Relation `xmlrpc:"facet_ids,omptempty"`
	GroupIds          *Relation `xmlrpc:"group_ids,omptempty"`
	HasWriteAccess    *Bool     `xmlrpc:"has_write_access,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	IsShared          *Bool     `xmlrpc:"is_shared,omptempty"`
	Name              *String   `xmlrpc:"name,omptempty"`
	ParentFolderId    *Many2One `xmlrpc:"parent_folder_id,omptempty"`
	ParentPath        *String   `xmlrpc:"parent_path,omptempty"`
	ProjectIds        *Relation `xmlrpc:"project_ids,omptempty"`
	ReadGroupIds      *Relation `xmlrpc:"read_group_ids,omptempty"`
	Sequence          *Int      `xmlrpc:"sequence,omptempty"`
	ShareLinkIds      *Relation `xmlrpc:"share_link_ids,omptempty"`
	UserSpecific      *Bool     `xmlrpc:"user_specific,omptempty"`
	UserSpecificWrite *Bool     `xmlrpc:"user_specific_write,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DocumentsFolders represents array of documents.folder model.
type DocumentsFolders []DocumentsFolder

// DocumentsFolderModel is the odoo model name.
const DocumentsFolderModel = "documents.folder"

// Many2One convert DocumentsFolder to *Many2One.
func (df *DocumentsFolder) Many2One() *Many2One {
	return NewMany2One(df.Id.Get(), "")
}

// CreateDocumentsFolder creates a new documents.folder model and returns its id.
func (c *Client) CreateDocumentsFolder(df *DocumentsFolder) (int64, error) {
	ids, err := c.CreateDocumentsFolders([]*DocumentsFolder{df})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsFolder creates a new documents.folder model and returns its id.
func (c *Client) CreateDocumentsFolders(dfs []*DocumentsFolder) ([]int64, error) {
	var vv []interface{}
	for _, v := range dfs {
		vv = append(vv, v)
	}
	return c.Create(DocumentsFolderModel, vv)
}

// UpdateDocumentsFolder updates an existing documents.folder record.
func (c *Client) UpdateDocumentsFolder(df *DocumentsFolder) error {
	return c.UpdateDocumentsFolders([]int64{df.Id.Get()}, df)
}

// UpdateDocumentsFolders updates existing documents.folder records.
// All records (represented by ids) will be updated by df values.
func (c *Client) UpdateDocumentsFolders(ids []int64, df *DocumentsFolder) error {
	return c.Update(DocumentsFolderModel, ids, df)
}

// DeleteDocumentsFolder deletes an existing documents.folder record.
func (c *Client) DeleteDocumentsFolder(id int64) error {
	return c.DeleteDocumentsFolders([]int64{id})
}

// DeleteDocumentsFolders deletes existing documents.folder records.
func (c *Client) DeleteDocumentsFolders(ids []int64) error {
	return c.Delete(DocumentsFolderModel, ids)
}

// GetDocumentsFolder gets documents.folder existing record.
func (c *Client) GetDocumentsFolder(id int64) (*DocumentsFolder, error) {
	dfs, err := c.GetDocumentsFolders([]int64{id})
	if err != nil {
		return nil, err
	}
	if dfs != nil && len(*dfs) > 0 {
		return &((*dfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.folder not found", id)
}

// GetDocumentsFolders gets documents.folder existing records.
func (c *Client) GetDocumentsFolders(ids []int64) (*DocumentsFolders, error) {
	dfs := &DocumentsFolders{}
	if err := c.Read(DocumentsFolderModel, ids, nil, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDocumentsFolder finds documents.folder record by querying it with criteria.
func (c *Client) FindDocumentsFolder(criteria *Criteria) (*DocumentsFolder, error) {
	dfs := &DocumentsFolders{}
	if err := c.SearchRead(DocumentsFolderModel, criteria, NewOptions().Limit(1), dfs); err != nil {
		return nil, err
	}
	if dfs != nil && len(*dfs) > 0 {
		return &((*dfs)[0]), nil
	}
	return nil, fmt.Errorf("documents.folder was not found with criteria %v", criteria)
}

// FindDocumentsFolders finds documents.folder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFolders(criteria *Criteria, options *Options) (*DocumentsFolders, error) {
	dfs := &DocumentsFolders{}
	if err := c.SearchRead(DocumentsFolderModel, criteria, options, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDocumentsFolderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsFolderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsFolderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsFolderId finds record id by querying it with criteria.
func (c *Client) FindDocumentsFolderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsFolderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.folder was not found with criteria %v and options %v", criteria, options)
}
