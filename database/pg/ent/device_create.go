// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pg/ent/device"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// DeviceCreate is the builder for creating a Device entity.
type DeviceCreate struct {
	config
	mutation *DeviceMutation
	hooks    []Hook
}

// SetBornAt sets the "born_at" field.
func (dc *DeviceCreate) SetBornAt(t time.Time) *DeviceCreate {
	dc.mutation.SetBornAt(t)
	return dc
}

// SetHashedPasswd sets the "hashed_passwd" field.
func (dc *DeviceCreate) SetHashedPasswd(s string) *DeviceCreate {
	dc.mutation.SetHashedPasswd(s)
	return dc
}

// SetDeadAt sets the "dead_at" field.
func (dc *DeviceCreate) SetDeadAt(t time.Time) *DeviceCreate {
	dc.mutation.SetDeadAt(t)
	return dc
}

// SetNillableDeadAt sets the "dead_at" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableDeadAt(t *time.Time) *DeviceCreate {
	if t != nil {
		dc.SetDeadAt(*t)
	}
	return dc
}

// SetReason sets the "reason" field.
func (dc *DeviceCreate) SetReason(s string) *DeviceCreate {
	dc.mutation.SetReason(s)
	return dc
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableReason(s *string) *DeviceCreate {
	if s != nil {
		dc.SetReason(*s)
	}
	return dc
}

// SetParent sets the "parent" field.
func (dc *DeviceCreate) SetParent(u uuid.UUID) *DeviceCreate {
	dc.mutation.SetParent(u)
	return dc
}

// SetNillableParent sets the "parent" field if the given value is not nil.
func (dc *DeviceCreate) SetNillableParent(u *uuid.UUID) *DeviceCreate {
	if u != nil {
		dc.SetParent(*u)
	}
	return dc
}

// SetID sets the "id" field.
func (dc *DeviceCreate) SetID(u uuid.UUID) *DeviceCreate {
	dc.mutation.SetID(u)
	return dc
}

// SetMotherID sets the "mother" edge to the Device entity by ID.
func (dc *DeviceCreate) SetMotherID(id uuid.UUID) *DeviceCreate {
	dc.mutation.SetMotherID(id)
	return dc
}

// SetNillableMotherID sets the "mother" edge to the Device entity by ID if the given value is not nil.
func (dc *DeviceCreate) SetNillableMotherID(id *uuid.UUID) *DeviceCreate {
	if id != nil {
		dc = dc.SetMotherID(*id)
	}
	return dc
}

// SetMother sets the "mother" edge to the Device entity.
func (dc *DeviceCreate) SetMother(d *Device) *DeviceCreate {
	return dc.SetMotherID(d.ID)
}

// AddChildIDs adds the "children" edge to the Device entity by IDs.
func (dc *DeviceCreate) AddChildIDs(ids ...uuid.UUID) *DeviceCreate {
	dc.mutation.AddChildIDs(ids...)
	return dc
}

// AddChildren adds the "children" edges to the Device entity.
func (dc *DeviceCreate) AddChildren(d ...*Device) *DeviceCreate {
	ids := make([]uuid.UUID, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddChildIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (dc *DeviceCreate) Mutation() *DeviceMutation {
	return dc.mutation
}

// Save creates the Device in the database.
func (dc *DeviceCreate) Save(ctx context.Context) (*Device, error) {
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeviceCreate) SaveX(ctx context.Context) *Device {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeviceCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeviceCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeviceCreate) check() error {
	if _, ok := dc.mutation.BornAt(); !ok {
		return &ValidationError{Name: "born_at", err: errors.New(`ent: missing required field "Device.born_at"`)}
	}
	if _, ok := dc.mutation.HashedPasswd(); !ok {
		return &ValidationError{Name: "hashed_passwd", err: errors.New(`ent: missing required field "Device.hashed_passwd"`)}
	}
	return nil
}

func (dc *DeviceCreate) sqlSave(ctx context.Context) (*Device, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeviceCreate) createSpec() (*Device, *sqlgraph.CreateSpec) {
	var (
		_node = &Device{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(device.Table, sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID))
	)
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := dc.mutation.BornAt(); ok {
		_spec.SetField(device.FieldBornAt, field.TypeTime, value)
		_node.BornAt = value
	}
	if value, ok := dc.mutation.HashedPasswd(); ok {
		_spec.SetField(device.FieldHashedPasswd, field.TypeString, value)
		_node.HashedPasswd = value
	}
	if value, ok := dc.mutation.DeadAt(); ok {
		_spec.SetField(device.FieldDeadAt, field.TypeTime, value)
		_node.DeadAt = &value
	}
	if value, ok := dc.mutation.Reason(); ok {
		_spec.SetField(device.FieldReason, field.TypeString, value)
		_node.Reason = value
	}
	if nodes := dc.mutation.MotherIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   device.MotherTable,
			Columns: []string{device.MotherColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.Parent = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := dc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.ChildrenTable,
			Columns: []string{device.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DeviceCreateBulk is the builder for creating many Device entities in bulk.
type DeviceCreateBulk struct {
	config
	err      error
	builders []*DeviceCreate
}

// Save creates the Device entities in the database.
func (dcb *DeviceCreateBulk) Save(ctx context.Context) ([]*Device, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Device, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeviceCreateBulk) SaveX(ctx context.Context) []*Device {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeviceCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeviceCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
