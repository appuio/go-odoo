package odoo

import (
	"fmt"
)

// RequestAppraisal represents request.appraisal model.
type RequestAppraisal struct {
	LastUpdate           *Time     `xmlrpc:"__last_update,omptempty"`
	AppraisalId          *Many2One `xmlrpc:"appraisal_id,omptempty"`
	AuthorId             *Many2One `xmlrpc:"author_id,omptempty"`
	Body                 *String   `xmlrpc:"body,omptempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omptempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String   `xmlrpc:"display_name,omptempty"`
	EmailFrom            *String   `xmlrpc:"email_from,omptempty"`
	EmployeeId           *Many2One `xmlrpc:"employee_id,omptempty"`
	Id                   *Int      `xmlrpc:"id,omptempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omptempty"`
	Lang                 *String   `xmlrpc:"lang,omptempty"`
	RecipientIds         *Relation `xmlrpc:"recipient_ids,omptempty"`
	RenderModel          *String   `xmlrpc:"render_model,omptempty"`
	Subject              *String   `xmlrpc:"subject,omptempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omptempty"`
	UserBody             *String   `xmlrpc:"user_body,omptempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// RequestAppraisals represents array of request.appraisal model.
type RequestAppraisals []RequestAppraisal

// RequestAppraisalModel is the odoo model name.
const RequestAppraisalModel = "request.appraisal"

// Many2One convert RequestAppraisal to *Many2One.
func (ra *RequestAppraisal) Many2One() *Many2One {
	return NewMany2One(ra.Id.Get(), "")
}

// CreateRequestAppraisal creates a new request.appraisal model and returns its id.
func (c *Client) CreateRequestAppraisal(ra *RequestAppraisal) (int64, error) {
	ids, err := c.CreateRequestAppraisals([]*RequestAppraisal{ra})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateRequestAppraisal creates a new request.appraisal model and returns its id.
func (c *Client) CreateRequestAppraisals(ras []*RequestAppraisal) ([]int64, error) {
	var vv []interface{}
	for _, v := range ras {
		vv = append(vv, v)
	}
	return c.Create(RequestAppraisalModel, vv)
}

// UpdateRequestAppraisal updates an existing request.appraisal record.
func (c *Client) UpdateRequestAppraisal(ra *RequestAppraisal) error {
	return c.UpdateRequestAppraisals([]int64{ra.Id.Get()}, ra)
}

// UpdateRequestAppraisals updates existing request.appraisal records.
// All records (represented by ids) will be updated by ra values.
func (c *Client) UpdateRequestAppraisals(ids []int64, ra *RequestAppraisal) error {
	return c.Update(RequestAppraisalModel, ids, ra)
}

// DeleteRequestAppraisal deletes an existing request.appraisal record.
func (c *Client) DeleteRequestAppraisal(id int64) error {
	return c.DeleteRequestAppraisals([]int64{id})
}

// DeleteRequestAppraisals deletes existing request.appraisal records.
func (c *Client) DeleteRequestAppraisals(ids []int64) error {
	return c.Delete(RequestAppraisalModel, ids)
}

// GetRequestAppraisal gets request.appraisal existing record.
func (c *Client) GetRequestAppraisal(id int64) (*RequestAppraisal, error) {
	ras, err := c.GetRequestAppraisals([]int64{id})
	if err != nil {
		return nil, err
	}
	if ras != nil && len(*ras) > 0 {
		return &((*ras)[0]), nil
	}
	return nil, fmt.Errorf("id %v of request.appraisal not found", id)
}

// GetRequestAppraisals gets request.appraisal existing records.
func (c *Client) GetRequestAppraisals(ids []int64) (*RequestAppraisals, error) {
	ras := &RequestAppraisals{}
	if err := c.Read(RequestAppraisalModel, ids, nil, ras); err != nil {
		return nil, err
	}
	return ras, nil
}

// FindRequestAppraisal finds request.appraisal record by querying it with criteria.
func (c *Client) FindRequestAppraisal(criteria *Criteria) (*RequestAppraisal, error) {
	ras := &RequestAppraisals{}
	if err := c.SearchRead(RequestAppraisalModel, criteria, NewOptions().Limit(1), ras); err != nil {
		return nil, err
	}
	if ras != nil && len(*ras) > 0 {
		return &((*ras)[0]), nil
	}
	return nil, fmt.Errorf("request.appraisal was not found with criteria %v", criteria)
}

// FindRequestAppraisals finds request.appraisal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindRequestAppraisals(criteria *Criteria, options *Options) (*RequestAppraisals, error) {
	ras := &RequestAppraisals{}
	if err := c.SearchRead(RequestAppraisalModel, criteria, options, ras); err != nil {
		return nil, err
	}
	return ras, nil
}

// FindRequestAppraisalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindRequestAppraisalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(RequestAppraisalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindRequestAppraisalId finds record id by querying it with criteria.
func (c *Client) FindRequestAppraisalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(RequestAppraisalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("request.appraisal was not found with criteria %v and options %v", criteria, options)
}
