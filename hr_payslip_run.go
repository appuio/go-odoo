package odoo

import (
	"fmt"
)

// HrPayslipRun represents hr.payslip.run model.
type HrPayslipRun struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryCode                 *String    `xmlrpc:"country_code,omptempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omptempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	FailedMessageIds            *Relation  `xmlrpc:"failed_message_ids,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
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
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	PayslipCount                *Int       `xmlrpc:"payslip_count,omptempty"`
	SepaExport                  *String    `xmlrpc:"sepa_export,omptempty"`
	SepaExportDate              *Time      `xmlrpc:"sepa_export_date,omptempty"`
	SepaExportFilename          *String    `xmlrpc:"sepa_export_filename,omptempty"`
	SlipIds                     *Relation  `xmlrpc:"slip_ids,omptempty"`
	State                       *Selection `xmlrpc:"state,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// HrPayslipRuns represents array of hr.payslip.run model.
type HrPayslipRuns []HrPayslipRun

// HrPayslipRunModel is the odoo model name.
const HrPayslipRunModel = "hr.payslip.run"

// Many2One convert HrPayslipRun to *Many2One.
func (hpr *HrPayslipRun) Many2One() *Many2One {
	return NewMany2One(hpr.Id.Get(), "")
}

// CreateHrPayslipRun creates a new hr.payslip.run model and returns its id.
func (c *Client) CreateHrPayslipRun(hpr *HrPayslipRun) (int64, error) {
	ids, err := c.CreateHrPayslipRuns([]*HrPayslipRun{hpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayslipRun creates a new hr.payslip.run model and returns its id.
func (c *Client) CreateHrPayslipRuns(hprs []*HrPayslipRun) ([]int64, error) {
	var vv []interface{}
	for _, v := range hprs {
		vv = append(vv, v)
	}
	return c.Create(HrPayslipRunModel, vv)
}

// UpdateHrPayslipRun updates an existing hr.payslip.run record.
func (c *Client) UpdateHrPayslipRun(hpr *HrPayslipRun) error {
	return c.UpdateHrPayslipRuns([]int64{hpr.Id.Get()}, hpr)
}

// UpdateHrPayslipRuns updates existing hr.payslip.run records.
// All records (represented by ids) will be updated by hpr values.
func (c *Client) UpdateHrPayslipRuns(ids []int64, hpr *HrPayslipRun) error {
	return c.Update(HrPayslipRunModel, ids, hpr)
}

// DeleteHrPayslipRun deletes an existing hr.payslip.run record.
func (c *Client) DeleteHrPayslipRun(id int64) error {
	return c.DeleteHrPayslipRuns([]int64{id})
}

// DeleteHrPayslipRuns deletes existing hr.payslip.run records.
func (c *Client) DeleteHrPayslipRuns(ids []int64) error {
	return c.Delete(HrPayslipRunModel, ids)
}

// GetHrPayslipRun gets hr.payslip.run existing record.
func (c *Client) GetHrPayslipRun(id int64) (*HrPayslipRun, error) {
	hprs, err := c.GetHrPayslipRuns([]int64{id})
	if err != nil {
		return nil, err
	}
	if hprs != nil && len(*hprs) > 0 {
		return &((*hprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of hr.payslip.run not found", id)
}

// GetHrPayslipRuns gets hr.payslip.run existing records.
func (c *Client) GetHrPayslipRuns(ids []int64) (*HrPayslipRuns, error) {
	hprs := &HrPayslipRuns{}
	if err := c.Read(HrPayslipRunModel, ids, nil, hprs); err != nil {
		return nil, err
	}
	return hprs, nil
}

// FindHrPayslipRun finds hr.payslip.run record by querying it with criteria.
func (c *Client) FindHrPayslipRun(criteria *Criteria) (*HrPayslipRun, error) {
	hprs := &HrPayslipRuns{}
	if err := c.SearchRead(HrPayslipRunModel, criteria, NewOptions().Limit(1), hprs); err != nil {
		return nil, err
	}
	if hprs != nil && len(*hprs) > 0 {
		return &((*hprs)[0]), nil
	}
	return nil, fmt.Errorf("hr.payslip.run was not found with criteria %v", criteria)
}

// FindHrPayslipRuns finds hr.payslip.run records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipRuns(criteria *Criteria, options *Options) (*HrPayslipRuns, error) {
	hprs := &HrPayslipRuns{}
	if err := c.SearchRead(HrPayslipRunModel, criteria, options, hprs); err != nil {
		return nil, err
	}
	return hprs, nil
}

// FindHrPayslipRunIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayslipRunIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(HrPayslipRunModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindHrPayslipRunId finds record id by querying it with criteria.
func (c *Client) FindHrPayslipRunId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayslipRunModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("hr.payslip.run was not found with criteria %v and options %v", criteria, options)
}
