package odoo

import (
	"fmt"
)

// DocumentsTag represents documents.tag model.
type DocumentsTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FacetId     *Many2One `xmlrpc:"facet_id,omptempty"`
	FolderId    *Many2One `xmlrpc:"folder_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DocumentsTags represents array of documents.tag model.
type DocumentsTags []DocumentsTag

// DocumentsTagModel is the odoo model name.
const DocumentsTagModel = "documents.tag"

// Many2One convert DocumentsTag to *Many2One.
func (dt *DocumentsTag) Many2One() *Many2One {
	return NewMany2One(dt.Id.Get(), "")
}

// CreateDocumentsTag creates a new documents.tag model and returns its id.
func (c *Client) CreateDocumentsTag(dt *DocumentsTag) (int64, error) {
	ids, err := c.CreateDocumentsTags([]*DocumentsTag{dt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsTag creates a new documents.tag model and returns its id.
func (c *Client) CreateDocumentsTags(dts []*DocumentsTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range dts {
		vv = append(vv, v)
	}
	return c.Create(DocumentsTagModel, vv)
}

// UpdateDocumentsTag updates an existing documents.tag record.
func (c *Client) UpdateDocumentsTag(dt *DocumentsTag) error {
	return c.UpdateDocumentsTags([]int64{dt.Id.Get()}, dt)
}

// UpdateDocumentsTags updates existing documents.tag records.
// All records (represented by ids) will be updated by dt values.
func (c *Client) UpdateDocumentsTags(ids []int64, dt *DocumentsTag) error {
	return c.Update(DocumentsTagModel, ids, dt)
}

// DeleteDocumentsTag deletes an existing documents.tag record.
func (c *Client) DeleteDocumentsTag(id int64) error {
	return c.DeleteDocumentsTags([]int64{id})
}

// DeleteDocumentsTags deletes existing documents.tag records.
func (c *Client) DeleteDocumentsTags(ids []int64) error {
	return c.Delete(DocumentsTagModel, ids)
}

// GetDocumentsTag gets documents.tag existing record.
func (c *Client) GetDocumentsTag(id int64) (*DocumentsTag, error) {
	dts, err := c.GetDocumentsTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if dts != nil && len(*dts) > 0 {
		return &((*dts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.tag not found", id)
}

// GetDocumentsTags gets documents.tag existing records.
func (c *Client) GetDocumentsTags(ids []int64) (*DocumentsTags, error) {
	dts := &DocumentsTags{}
	if err := c.Read(DocumentsTagModel, ids, nil, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDocumentsTag finds documents.tag record by querying it with criteria.
func (c *Client) FindDocumentsTag(criteria *Criteria) (*DocumentsTag, error) {
	dts := &DocumentsTags{}
	if err := c.SearchRead(DocumentsTagModel, criteria, NewOptions().Limit(1), dts); err != nil {
		return nil, err
	}
	if dts != nil && len(*dts) > 0 {
		return &((*dts)[0]), nil
	}
	return nil, fmt.Errorf("documents.tag was not found with criteria %v", criteria)
}

// FindDocumentsTags finds documents.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsTags(criteria *Criteria, options *Options) (*DocumentsTags, error) {
	dts := &DocumentsTags{}
	if err := c.SearchRead(DocumentsTagModel, criteria, options, dts); err != nil {
		return nil, err
	}
	return dts, nil
}

// FindDocumentsTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsTagId finds record id by querying it with criteria.
func (c *Client) FindDocumentsTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.tag was not found with criteria %v and options %v", criteria, options)
}
