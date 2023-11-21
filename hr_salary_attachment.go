package odoo

import (
	"fmt"
)

// HrSalaryAttachment represents hr.salary.attachment model.
type HrSalaryAttachment struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActiveAmount                *Float     `xmlrpc:"active_amount,omptempty"`
	Attachment                  *String    `xmlrpc:"attachment,omptempty"`
	AttachmentName              *String    `xmlrpc:"attachment_name,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omptempty"`
	DateEstimatedEnd            *Time      `xmlrpc:"date_estimated_end,omptempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omptempty"`
	DeductionType               *Selection `xmlrpc:"deduction_type,omptempty"`
	Description                 *String    `xmlrpc:"description,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	HasSimilarAttachment        *Bool      `xmlrpc:"has_similar_attachment,omptempty"`
	HasSimilarAttachmentWarning *String    `xmlrpc:"has_similar_attachment_warning,omptempty"`
	HasTotalAmount              *Bool      `xmlrpc:"has_total_amount,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageContent              *String    `xmlrpc:"message_content,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MonthlyAmount               *Float     `xmlrpc:"monthly_amount,omptempty"`
	PaidAmount                  *Float     `xmlrpc:"paid_amount,omptempty"`
	PayslipCount                *Int       `xmlrpc:"payslip_count,omptempty"`
	PayslipIds                  *Relation  `xmlrpc:"payslip_ids,omptempty"`
	RemainingAmount             *Float     `xmlrpc:"remaining_amount,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	TotalAmount                 *Float     `xmlrpc:"total_amount,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrSalaryAttachments represents array of hr.salary.attachment model.
type HrSalaryAttachments []HrSalaryAttachment

// HrSalaryAttachmentModel is the odoo model name.
const HrSalaryAttachmentModel = "hr.salary.attachment"

// Many2One convert HrSalaryAttachment to *Many2One.
func (hsa *HrSalaryAttachment) Many2One() *Many2One {
	return NewMany2One(hsa.Id.Get(), "")
}

// CreateHrSalaryAttachment creates a new hr.salary.attachment model and returns its id.
func (c *Client) CreateHrSalaryAttachment(hsa *HrSalaryAttachment) (int64, error) {
	ids, err := c.CreateHrSalaryAttachments([]*HrSalaryAttachment{hsa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrSalaryAttachment creates a new hr.salary.attachment model and returns its id.
func (c *Client) CreateHrSalaryAttachments(hsas []*HrSalaryAttachment) ([]int64, error) {
	var vv []interface{}
	for _, v := range hsas {
		vv = append(vv, v)
	}
	return c.Create(HrSalaryAttachmentModel, vv)
}

// UpdateHrSalaryAttachment updates an existing hr.salary.attachment record.
func (c *Client) UpdateHrSalaryAttachment(hsa *HrSalaryAttachment) error {
	return c.UpdateHrSalaryAttachments([]int64{hsa.Id.Get()}, hsa)
}

// UpdateHrSalaryAttachments updates existing hr.salary.attachment records.
// All records (represented by ids) will be updated by hsa values.
func (c *Client) UpdateHrSalaryAttachments(ids []int64, hsa *HrSalaryAttachment) error {
	return c.Update(HrSalaryAttachmentModel, ids, hsa)
}

// DeleteHrSalaryAttachment deletes an existing hr.salary.attachment record.
func (c *Client) DeleteHrSalaryAttachment(id int64) error {
	return c.DeleteHrSalaryAttachments([]int64{id})
}

// DeleteHrSalaryAttachments deletes existing hr.salary.attachment records.
func (c *Client) DeleteHrSalaryAttachments(ids []int64) error {
	return c.Delete(HrSalaryAttachmentModel, ids)
}

// GetHrSalaryAttachment gets hr.salary.attachment existing record.
func (c *Client) GetHrSalaryAttachment(id int64) (*HrSalaryAttachment, error) {
	hsas, err := c.GetHrSalaryAttachments([]int64{id})
	if err != nil {
		return nil, err
	}
	if hsas != nil && len(*hsas) > 0 {
		return &((*hsas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.salary.attachment not found", id)
}

// GetHrSalaryAttachments gets hr.salary.attachment existing records.
func (c *Client) GetHrSalaryAttachments(ids []int64) (*HrSalaryAttachments, error) {
	hsas := &HrSalaryAttachments{}
	if err := c.Read(HrSalaryAttachmentModel, ids, nil, hsas); err != nil {
		return nil, err
	}
	return hsas, nil
}

// FindHrSalaryAttachment finds hr.salary.attachment record by querying it with criteria.
func (c *Client) FindHrSalaryAttachment(criteria *Criteria) (*HrSalaryAttachment, error) {
	hsas := &HrSalaryAttachments{}
	if err := c.SearchRead(HrSalaryAttachmentModel, criteria, NewOptions().Limit(1), hsas); err != nil {
		return nil, err
	}
	if hsas != nil && len(*hsas) > 0 {
		return &((*hsas)[0]), nil
	}
	return nil, fmt.Errorf("hr.salary.attachment was not found with criteria %v", criteria)
}

// FindHrSalaryAttachments finds hr.salary.attachment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryAttachments(criteria *Criteria, options *Options) (*HrSalaryAttachments, error) {
	hsas := &HrSalaryAttachments{}
	if err := c.SearchRead(HrSalaryAttachmentModel, criteria, options, hsas); err != nil {
		return nil, err
	}
	return hsas, nil
}

// FindHrSalaryAttachmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrSalaryAttachmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrSalaryAttachmentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrSalaryAttachmentId finds record id by querying it with criteria.
func (c *Client) FindHrSalaryAttachmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrSalaryAttachmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.salary.attachment was not found with criteria %v and options %v", criteria, options)
}
