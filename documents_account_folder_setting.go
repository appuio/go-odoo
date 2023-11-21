package odoo

import (
	"fmt"
)

// DocumentsAccountFolderSetting represents documents.account.folder.setting model.
type DocumentsAccountFolderSetting struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId   *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	FolderId    *Many2One `xmlrpc:"folder_id,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JournalId   *Many2One `xmlrpc:"journal_id,omptempty"`
	TagIds      *Relation `xmlrpc:"tag_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// DocumentsAccountFolderSettings represents array of documents.account.folder.setting model.
type DocumentsAccountFolderSettings []DocumentsAccountFolderSetting

// DocumentsAccountFolderSettingModel is the odoo model name.
const DocumentsAccountFolderSettingModel = "documents.account.folder.setting"

// Many2One convert DocumentsAccountFolderSetting to *Many2One.
func (dafs *DocumentsAccountFolderSetting) Many2One() *Many2One {
	return NewMany2One(dafs.Id.Get(), "")
}

// CreateDocumentsAccountFolderSetting creates a new documents.account.folder.setting model and returns its id.
func (c *Client) CreateDocumentsAccountFolderSetting(dafs *DocumentsAccountFolderSetting) (int64, error) {
	ids, err := c.CreateDocumentsAccountFolderSettings([]*DocumentsAccountFolderSetting{dafs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDocumentsAccountFolderSetting creates a new documents.account.folder.setting model and returns its id.
func (c *Client) CreateDocumentsAccountFolderSettings(dafss []*DocumentsAccountFolderSetting) ([]int64, error) {
	var vv []interface{}
	for _, v := range dafss {
		vv = append(vv, v)
	}
	return c.Create(DocumentsAccountFolderSettingModel, vv)
}

// UpdateDocumentsAccountFolderSetting updates an existing documents.account.folder.setting record.
func (c *Client) UpdateDocumentsAccountFolderSetting(dafs *DocumentsAccountFolderSetting) error {
	return c.UpdateDocumentsAccountFolderSettings([]int64{dafs.Id.Get()}, dafs)
}

// UpdateDocumentsAccountFolderSettings updates existing documents.account.folder.setting records.
// All records (represented by ids) will be updated by dafs values.
func (c *Client) UpdateDocumentsAccountFolderSettings(ids []int64, dafs *DocumentsAccountFolderSetting) error {
	return c.Update(DocumentsAccountFolderSettingModel, ids, dafs)
}

// DeleteDocumentsAccountFolderSetting deletes an existing documents.account.folder.setting record.
func (c *Client) DeleteDocumentsAccountFolderSetting(id int64) error {
	return c.DeleteDocumentsAccountFolderSettings([]int64{id})
}

// DeleteDocumentsAccountFolderSettings deletes existing documents.account.folder.setting records.
func (c *Client) DeleteDocumentsAccountFolderSettings(ids []int64) error {
	return c.Delete(DocumentsAccountFolderSettingModel, ids)
}

// GetDocumentsAccountFolderSetting gets documents.account.folder.setting existing record.
func (c *Client) GetDocumentsAccountFolderSetting(id int64) (*DocumentsAccountFolderSetting, error) {
	dafss, err := c.GetDocumentsAccountFolderSettings([]int64{id})
	if err != nil {
		return nil, err
	}
	if dafss != nil && len(*dafss) > 0 {
		return &((*dafss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of documents.account.folder.setting not found", id)
}

// GetDocumentsAccountFolderSettings gets documents.account.folder.setting existing records.
func (c *Client) GetDocumentsAccountFolderSettings(ids []int64) (*DocumentsAccountFolderSettings, error) {
	dafss := &DocumentsAccountFolderSettings{}
	if err := c.Read(DocumentsAccountFolderSettingModel, ids, nil, dafss); err != nil {
		return nil, err
	}
	return dafss, nil
}

// FindDocumentsAccountFolderSetting finds documents.account.folder.setting record by querying it with criteria.
func (c *Client) FindDocumentsAccountFolderSetting(criteria *Criteria) (*DocumentsAccountFolderSetting, error) {
	dafss := &DocumentsAccountFolderSettings{}
	if err := c.SearchRead(DocumentsAccountFolderSettingModel, criteria, NewOptions().Limit(1), dafss); err != nil {
		return nil, err
	}
	if dafss != nil && len(*dafss) > 0 {
		return &((*dafss)[0]), nil
	}
	return nil, fmt.Errorf("documents.account.folder.setting was not found with criteria %v", criteria)
}

// FindDocumentsAccountFolderSettings finds documents.account.folder.setting records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsAccountFolderSettings(criteria *Criteria, options *Options) (*DocumentsAccountFolderSettings, error) {
	dafss := &DocumentsAccountFolderSettings{}
	if err := c.SearchRead(DocumentsAccountFolderSettingModel, criteria, options, dafss); err != nil {
		return nil, err
	}
	return dafss, nil
}

// FindDocumentsAccountFolderSettingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDocumentsAccountFolderSettingIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(DocumentsAccountFolderSettingModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindDocumentsAccountFolderSettingId finds record id by querying it with criteria.
func (c *Client) FindDocumentsAccountFolderSettingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DocumentsAccountFolderSettingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("documents.account.folder.setting was not found with criteria %v and options %v", criteria, options)
}
