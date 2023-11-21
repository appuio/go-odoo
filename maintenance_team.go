package odoo

import (
	"fmt"
)

// MaintenanceTeam represents maintenance.team model.
type MaintenanceTeam struct {
	LastUpdate                   *Time     `xmlrpc:"__last_update,omptempty"`
	Active                       *Bool     `xmlrpc:"active,omptempty"`
	Color                        *Int      `xmlrpc:"color,omptempty"`
	CompanyId                    *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate                   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                  *String   `xmlrpc:"display_name,omptempty"`
	EquipmentIds                 *Relation `xmlrpc:"equipment_ids,omptempty"`
	Id                           *Int      `xmlrpc:"id,omptempty"`
	MemberIds                    *Relation `xmlrpc:"member_ids,omptempty"`
	Name                         *String   `xmlrpc:"name,omptempty"`
	RequestIds                   *Relation `xmlrpc:"request_ids,omptempty"`
	TodoRequestCount             *Int      `xmlrpc:"todo_request_count,omptempty"`
	TodoRequestCountBlock        *Int      `xmlrpc:"todo_request_count_block,omptempty"`
	TodoRequestCountDate         *Int      `xmlrpc:"todo_request_count_date,omptempty"`
	TodoRequestCountHighPriority *Int      `xmlrpc:"todo_request_count_high_priority,omptempty"`
	TodoRequestCountUnscheduled  *Int      `xmlrpc:"todo_request_count_unscheduled,omptempty"`
	TodoRequestIds               *Relation `xmlrpc:"todo_request_ids,omptempty"`
	WriteDate                    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MaintenanceTeams represents array of maintenance.team model.
type MaintenanceTeams []MaintenanceTeam

// MaintenanceTeamModel is the odoo model name.
const MaintenanceTeamModel = "maintenance.team"

// Many2One convert MaintenanceTeam to *Many2One.
func (mt *MaintenanceTeam) Many2One() *Many2One {
	return NewMany2One(mt.Id.Get(), "")
}

// CreateMaintenanceTeam creates a new maintenance.team model and returns its id.
func (c *Client) CreateMaintenanceTeam(mt *MaintenanceTeam) (int64, error) {
	ids, err := c.CreateMaintenanceTeams([]*MaintenanceTeam{mt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMaintenanceTeam creates a new maintenance.team model and returns its id.
func (c *Client) CreateMaintenanceTeams(mts []*MaintenanceTeam) ([]int64, error) {
	var vv []interface{}
	for _, v := range mts {
		vv = append(vv, v)
	}
	return c.Create(MaintenanceTeamModel, vv)
}

// UpdateMaintenanceTeam updates an existing maintenance.team record.
func (c *Client) UpdateMaintenanceTeam(mt *MaintenanceTeam) error {
	return c.UpdateMaintenanceTeams([]int64{mt.Id.Get()}, mt)
}

// UpdateMaintenanceTeams updates existing maintenance.team records.
// All records (represented by ids) will be updated by mt values.
func (c *Client) UpdateMaintenanceTeams(ids []int64, mt *MaintenanceTeam) error {
	return c.Update(MaintenanceTeamModel, ids, mt)
}

// DeleteMaintenanceTeam deletes an existing maintenance.team record.
func (c *Client) DeleteMaintenanceTeam(id int64) error {
	return c.DeleteMaintenanceTeams([]int64{id})
}

// DeleteMaintenanceTeams deletes existing maintenance.team records.
func (c *Client) DeleteMaintenanceTeams(ids []int64) error {
	return c.Delete(MaintenanceTeamModel, ids)
}

// GetMaintenanceTeam gets maintenance.team existing record.
func (c *Client) GetMaintenanceTeam(id int64) (*MaintenanceTeam, error) {
	mts, err := c.GetMaintenanceTeams([]int64{id})
	if err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of maintenance.team not found", id)
}

// GetMaintenanceTeams gets maintenance.team existing records.
func (c *Client) GetMaintenanceTeams(ids []int64) (*MaintenanceTeams, error) {
	mts := &MaintenanceTeams{}
	if err := c.Read(MaintenanceTeamModel, ids, nil, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMaintenanceTeam finds maintenance.team record by querying it with criteria.
func (c *Client) FindMaintenanceTeam(criteria *Criteria) (*MaintenanceTeam, error) {
	mts := &MaintenanceTeams{}
	if err := c.SearchRead(MaintenanceTeamModel, criteria, NewOptions().Limit(1), mts); err != nil {
		return nil, err
	}
	if mts != nil && len(*mts) > 0 {
		return &((*mts)[0]), nil
	}
	return nil, fmt.Errorf("maintenance.team was not found with criteria %v", criteria)
}

// FindMaintenanceTeams finds maintenance.team records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceTeams(criteria *Criteria, options *Options) (*MaintenanceTeams, error) {
	mts := &MaintenanceTeams{}
	if err := c.SearchRead(MaintenanceTeamModel, criteria, options, mts); err != nil {
		return nil, err
	}
	return mts, nil
}

// FindMaintenanceTeamIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMaintenanceTeamIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MaintenanceTeamModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMaintenanceTeamId finds record id by querying it with criteria.
func (c *Client) FindMaintenanceTeamId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MaintenanceTeamModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("maintenance.team was not found with criteria %v and options %v", criteria, options)
}
