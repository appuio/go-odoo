package odoo

import (
	"fmt"
)

// AccountBatchPaymentRejection represents account.batch.payment.rejection model.
type AccountBatchPaymentRejection struct {
	LastUpdate            *Time     `xmlrpc:"__last_update,omptempty"`
	CancelActionTodo      *Bool     `xmlrpc:"cancel_action_todo,omptempty"`
	CreateDate            *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String   `xmlrpc:"display_name,omptempty"`
	Id                    *Int      `xmlrpc:"id,omptempty"`
	InReconcilePaymentIds *Relation `xmlrpc:"in_reconcile_payment_ids,omptempty"`
	NbBatchPaymentIds     *Int      `xmlrpc:"nb_batch_payment_ids,omptempty"`
	NbRejectedPaymentIds  *Int      `xmlrpc:"nb_rejected_payment_ids,omptempty"`
	NextActionTodo        *String   `xmlrpc:"next_action_todo,omptempty"`
	RejectedPaymentIds    *Relation `xmlrpc:"rejected_payment_ids,omptempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBatchPaymentRejections represents array of account.batch.payment.rejection model.
type AccountBatchPaymentRejections []AccountBatchPaymentRejection

// AccountBatchPaymentRejectionModel is the odoo model name.
const AccountBatchPaymentRejectionModel = "account.batch.payment.rejection"

// Many2One convert AccountBatchPaymentRejection to *Many2One.
func (abpr *AccountBatchPaymentRejection) Many2One() *Many2One {
	return NewMany2One(abpr.Id.Get(), "")
}

// CreateAccountBatchPaymentRejection creates a new account.batch.payment.rejection model and returns its id.
func (c *Client) CreateAccountBatchPaymentRejection(abpr *AccountBatchPaymentRejection) (int64, error) {
	ids, err := c.CreateAccountBatchPaymentRejections([]*AccountBatchPaymentRejection{abpr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountBatchPaymentRejection creates a new account.batch.payment.rejection model and returns its id.
func (c *Client) CreateAccountBatchPaymentRejections(abprs []*AccountBatchPaymentRejection) ([]int64, error) {
	var vv []interface{}
	for _, v := range abprs {
		vv = append(vv, v)
	}
	return c.Create(AccountBatchPaymentRejectionModel, vv)
}

// UpdateAccountBatchPaymentRejection updates an existing account.batch.payment.rejection record.
func (c *Client) UpdateAccountBatchPaymentRejection(abpr *AccountBatchPaymentRejection) error {
	return c.UpdateAccountBatchPaymentRejections([]int64{abpr.Id.Get()}, abpr)
}

// UpdateAccountBatchPaymentRejections updates existing account.batch.payment.rejection records.
// All records (represented by ids) will be updated by abpr values.
func (c *Client) UpdateAccountBatchPaymentRejections(ids []int64, abpr *AccountBatchPaymentRejection) error {
	return c.Update(AccountBatchPaymentRejectionModel, ids, abpr)
}

// DeleteAccountBatchPaymentRejection deletes an existing account.batch.payment.rejection record.
func (c *Client) DeleteAccountBatchPaymentRejection(id int64) error {
	return c.DeleteAccountBatchPaymentRejections([]int64{id})
}

// DeleteAccountBatchPaymentRejections deletes existing account.batch.payment.rejection records.
func (c *Client) DeleteAccountBatchPaymentRejections(ids []int64) error {
	return c.Delete(AccountBatchPaymentRejectionModel, ids)
}

// GetAccountBatchPaymentRejection gets account.batch.payment.rejection existing record.
func (c *Client) GetAccountBatchPaymentRejection(id int64) (*AccountBatchPaymentRejection, error) {
	abprs, err := c.GetAccountBatchPaymentRejections([]int64{id})
	if err != nil {
		return nil, err
	}
	if abprs != nil && len(*abprs) > 0 {
		return &((*abprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.batch.payment.rejection not found", id)
}

// GetAccountBatchPaymentRejections gets account.batch.payment.rejection existing records.
func (c *Client) GetAccountBatchPaymentRejections(ids []int64) (*AccountBatchPaymentRejections, error) {
	abprs := &AccountBatchPaymentRejections{}
	if err := c.Read(AccountBatchPaymentRejectionModel, ids, nil, abprs); err != nil {
		return nil, err
	}
	return abprs, nil
}

// FindAccountBatchPaymentRejection finds account.batch.payment.rejection record by querying it with criteria.
func (c *Client) FindAccountBatchPaymentRejection(criteria *Criteria) (*AccountBatchPaymentRejection, error) {
	abprs := &AccountBatchPaymentRejections{}
	if err := c.SearchRead(AccountBatchPaymentRejectionModel, criteria, NewOptions().Limit(1), abprs); err != nil {
		return nil, err
	}
	if abprs != nil && len(*abprs) > 0 {
		return &((*abprs)[0]), nil
	}
	return nil, fmt.Errorf("account.batch.payment.rejection was not found with criteria %v", criteria)
}

// FindAccountBatchPaymentRejections finds account.batch.payment.rejection records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPaymentRejections(criteria *Criteria, options *Options) (*AccountBatchPaymentRejections, error) {
	abprs := &AccountBatchPaymentRejections{}
	if err := c.SearchRead(AccountBatchPaymentRejectionModel, criteria, options, abprs); err != nil {
		return nil, err
	}
	return abprs, nil
}

// FindAccountBatchPaymentRejectionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBatchPaymentRejectionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBatchPaymentRejectionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBatchPaymentRejectionId finds record id by querying it with criteria.
func (c *Client) FindAccountBatchPaymentRejectionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBatchPaymentRejectionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.batch.payment.rejection was not found with criteria %v and options %v", criteria, options)
}
