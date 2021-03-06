// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/gen0cide/laforge/ent/disk"
	"github.com/gen0cide/laforge/ent/host"
	"github.com/gen0cide/laforge/ent/tag"
	"github.com/gen0cide/laforge/ent/user"
)

// HostCreate is the builder for creating a Host entity.
type HostCreate struct {
	config
	mutation *HostMutation
	hooks    []Hook
}

// SetHostname sets the hostname field.
func (hc *HostCreate) SetHostname(s string) *HostCreate {
	hc.mutation.SetHostname(s)
	return hc
}

// SetDescription sets the description field.
func (hc *HostCreate) SetDescription(s string) *HostCreate {
	hc.mutation.SetDescription(s)
	return hc
}

// SetOS sets the OS field.
func (hc *HostCreate) SetOS(s string) *HostCreate {
	hc.mutation.SetOS(s)
	return hc
}

// SetLastOctet sets the last_octet field.
func (hc *HostCreate) SetLastOctet(i int) *HostCreate {
	hc.mutation.SetLastOctet(i)
	return hc
}

// SetAllowMACChanges sets the allow_mac_changes field.
func (hc *HostCreate) SetAllowMACChanges(b bool) *HostCreate {
	hc.mutation.SetAllowMACChanges(b)
	return hc
}

// SetExposedTCPPorts sets the exposed_tcp_ports field.
func (hc *HostCreate) SetExposedTCPPorts(s []string) *HostCreate {
	hc.mutation.SetExposedTCPPorts(s)
	return hc
}

// SetExposedUDPPorts sets the exposed_udp_ports field.
func (hc *HostCreate) SetExposedUDPPorts(s []string) *HostCreate {
	hc.mutation.SetExposedUDPPorts(s)
	return hc
}

// SetOverridePassword sets the override_password field.
func (hc *HostCreate) SetOverridePassword(s string) *HostCreate {
	hc.mutation.SetOverridePassword(s)
	return hc
}

// SetVars sets the vars field.
func (hc *HostCreate) SetVars(m map[string]string) *HostCreate {
	hc.mutation.SetVars(m)
	return hc
}

// SetUserGroups sets the user_groups field.
func (hc *HostCreate) SetUserGroups(s []string) *HostCreate {
	hc.mutation.SetUserGroups(s)
	return hc
}

// SetDependsOn sets the depends_on field.
func (hc *HostCreate) SetDependsOn(s []string) *HostCreate {
	hc.mutation.SetDependsOn(s)
	return hc
}

// SetScripts sets the scripts field.
func (hc *HostCreate) SetScripts(s []string) *HostCreate {
	hc.mutation.SetScripts(s)
	return hc
}

// SetCommands sets the commands field.
func (hc *HostCreate) SetCommands(s []string) *HostCreate {
	hc.mutation.SetCommands(s)
	return hc
}

// SetRemoteFiles sets the remote_files field.
func (hc *HostCreate) SetRemoteFiles(s []string) *HostCreate {
	hc.mutation.SetRemoteFiles(s)
	return hc
}

// SetDNSRecords sets the dns_records field.
func (hc *HostCreate) SetDNSRecords(s []string) *HostCreate {
	hc.mutation.SetDNSRecords(s)
	return hc
}

// AddDiskIDs adds the disk edge to Disk by ids.
func (hc *HostCreate) AddDiskIDs(ids ...int) *HostCreate {
	hc.mutation.AddDiskIDs(ids...)
	return hc
}

// AddDisk adds the disk edges to Disk.
func (hc *HostCreate) AddDisk(d ...*Disk) *HostCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return hc.AddDiskIDs(ids...)
}

// AddMaintainerIDs adds the maintainer edge to User by ids.
func (hc *HostCreate) AddMaintainerIDs(ids ...int) *HostCreate {
	hc.mutation.AddMaintainerIDs(ids...)
	return hc
}

// AddMaintainer adds the maintainer edges to User.
func (hc *HostCreate) AddMaintainer(u ...*User) *HostCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return hc.AddMaintainerIDs(ids...)
}

// AddTagIDs adds the tag edge to Tag by ids.
func (hc *HostCreate) AddTagIDs(ids ...int) *HostCreate {
	hc.mutation.AddTagIDs(ids...)
	return hc
}

// AddTag adds the tag edges to Tag.
func (hc *HostCreate) AddTag(t ...*Tag) *HostCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return hc.AddTagIDs(ids...)
}

// Mutation returns the HostMutation object of the builder.
func (hc *HostCreate) Mutation() *HostMutation {
	return hc.mutation
}

// Save creates the Host in the database.
func (hc *HostCreate) Save(ctx context.Context) (*Host, error) {
	var (
		err  error
		node *Host
	)
	if len(hc.hooks) == 0 {
		if err = hc.check(); err != nil {
			return nil, err
		}
		node, err = hc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = hc.check(); err != nil {
				return nil, err
			}
			hc.mutation = mutation
			node, err = hc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(hc.hooks) - 1; i >= 0; i-- {
			mut = hc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HostCreate) SaveX(ctx context.Context) *Host {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (hc *HostCreate) check() error {
	if _, ok := hc.mutation.Hostname(); !ok {
		return &ValidationError{Name: "hostname", err: errors.New("ent: missing required field \"hostname\"")}
	}
	if _, ok := hc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New("ent: missing required field \"description\"")}
	}
	if _, ok := hc.mutation.OS(); !ok {
		return &ValidationError{Name: "OS", err: errors.New("ent: missing required field \"OS\"")}
	}
	if _, ok := hc.mutation.LastOctet(); !ok {
		return &ValidationError{Name: "last_octet", err: errors.New("ent: missing required field \"last_octet\"")}
	}
	if _, ok := hc.mutation.AllowMACChanges(); !ok {
		return &ValidationError{Name: "allow_mac_changes", err: errors.New("ent: missing required field \"allow_mac_changes\"")}
	}
	if _, ok := hc.mutation.ExposedTCPPorts(); !ok {
		return &ValidationError{Name: "exposed_tcp_ports", err: errors.New("ent: missing required field \"exposed_tcp_ports\"")}
	}
	if _, ok := hc.mutation.ExposedUDPPorts(); !ok {
		return &ValidationError{Name: "exposed_udp_ports", err: errors.New("ent: missing required field \"exposed_udp_ports\"")}
	}
	if _, ok := hc.mutation.OverridePassword(); !ok {
		return &ValidationError{Name: "override_password", err: errors.New("ent: missing required field \"override_password\"")}
	}
	if _, ok := hc.mutation.Vars(); !ok {
		return &ValidationError{Name: "vars", err: errors.New("ent: missing required field \"vars\"")}
	}
	if _, ok := hc.mutation.UserGroups(); !ok {
		return &ValidationError{Name: "user_groups", err: errors.New("ent: missing required field \"user_groups\"")}
	}
	return nil
}

func (hc *HostCreate) sqlSave(ctx context.Context) (*Host, error) {
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (hc *HostCreate) createSpec() (*Host, *sqlgraph.CreateSpec) {
	var (
		_node = &Host{config: hc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: host.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: host.FieldID,
			},
		}
	)
	if value, ok := hc.mutation.Hostname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldHostname,
		})
		_node.Hostname = value
	}
	if value, ok := hc.mutation.Description(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldDescription,
		})
		_node.Description = value
	}
	if value, ok := hc.mutation.OS(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldOS,
		})
		_node.OS = value
	}
	if value, ok := hc.mutation.LastOctet(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: host.FieldLastOctet,
		})
		_node.LastOctet = value
	}
	if value, ok := hc.mutation.AllowMACChanges(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: host.FieldAllowMACChanges,
		})
		_node.AllowMACChanges = value
	}
	if value, ok := hc.mutation.ExposedTCPPorts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldExposedTCPPorts,
		})
		_node.ExposedTCPPorts = value
	}
	if value, ok := hc.mutation.ExposedUDPPorts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldExposedUDPPorts,
		})
		_node.ExposedUDPPorts = value
	}
	if value, ok := hc.mutation.OverridePassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: host.FieldOverridePassword,
		})
		_node.OverridePassword = value
	}
	if value, ok := hc.mutation.Vars(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldVars,
		})
		_node.Vars = value
	}
	if value, ok := hc.mutation.UserGroups(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldUserGroups,
		})
		_node.UserGroups = value
	}
	if value, ok := hc.mutation.DependsOn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldDependsOn,
		})
		_node.DependsOn = value
	}
	if value, ok := hc.mutation.Scripts(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldScripts,
		})
		_node.Scripts = value
	}
	if value, ok := hc.mutation.Commands(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldCommands,
		})
		_node.Commands = value
	}
	if value, ok := hc.mutation.RemoteFiles(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldRemoteFiles,
		})
		_node.RemoteFiles = value
	}
	if value, ok := hc.mutation.DNSRecords(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: host.FieldDNSRecords,
		})
		_node.DNSRecords = value
	}
	if nodes := hc.mutation.DiskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.DiskTable,
			Columns: []string{host.DiskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: disk.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.MaintainerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.MaintainerTable,
			Columns: []string{host.MaintainerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := hc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   host.TagTable,
			Columns: []string{host.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HostCreateBulk is the builder for creating a bulk of Host entities.
type HostCreateBulk struct {
	config
	builders []*HostCreate
}

// Save creates the Host entities in the database.
func (hcb *HostCreateBulk) Save(ctx context.Context) ([]*Host, error) {
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Host, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (hcb *HostCreateBulk) SaveX(ctx context.Context) []*Host {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
