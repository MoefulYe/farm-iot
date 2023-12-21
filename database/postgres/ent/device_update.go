// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/device"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/predicate"
)

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetParent sets the "parent" field.
func (du *DeviceUpdate) SetParent(s string) *DeviceUpdate {
	du.mutation.SetParent(s)
	return du
}

// SetHashedPasswd sets the "hashed_passwd" field.
func (du *DeviceUpdate) SetHashedPasswd(s string) *DeviceUpdate {
	du.mutation.SetHashedPasswd(s)
	return du
}

// SetDeadAt sets the "dead_at" field.
func (du *DeviceUpdate) SetDeadAt(t time.Time) *DeviceUpdate {
	du.mutation.SetDeadAt(t)
	return du
}

// SetNillableDeadAt sets the "dead_at" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableDeadAt(t *time.Time) *DeviceUpdate {
	if t != nil {
		du.SetDeadAt(*t)
	}
	return du
}

// ClearDeadAt clears the value of the "dead_at" field.
func (du *DeviceUpdate) ClearDeadAt() *DeviceUpdate {
	du.mutation.ClearDeadAt()
	return du
}

// SetReason sets the "reason" field.
func (du *DeviceUpdate) SetReason(s string) *DeviceUpdate {
	du.mutation.SetReason(s)
	return du
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (du *DeviceUpdate) SetNillableReason(s *string) *DeviceUpdate {
	if s != nil {
		du.SetReason(*s)
	}
	return du
}

// ClearReason clears the value of the "reason" field.
func (du *DeviceUpdate) ClearReason() *DeviceUpdate {
	du.mutation.ClearReason()
	return du
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Parent(); ok {
		_spec.SetField(device.FieldParent, field.TypeString, value)
	}
	if value, ok := du.mutation.HashedPasswd(); ok {
		_spec.SetField(device.FieldHashedPasswd, field.TypeString, value)
	}
	if value, ok := du.mutation.DeadAt(); ok {
		_spec.SetField(device.FieldDeadAt, field.TypeTime, value)
	}
	if du.mutation.DeadAtCleared() {
		_spec.ClearField(device.FieldDeadAt, field.TypeTime)
	}
	if value, ok := du.mutation.Reason(); ok {
		_spec.SetField(device.FieldReason, field.TypeString, value)
	}
	if du.mutation.ReasonCleared() {
		_spec.ClearField(device.FieldReason, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetParent sets the "parent" field.
func (duo *DeviceUpdateOne) SetParent(s string) *DeviceUpdateOne {
	duo.mutation.SetParent(s)
	return duo
}

// SetHashedPasswd sets the "hashed_passwd" field.
func (duo *DeviceUpdateOne) SetHashedPasswd(s string) *DeviceUpdateOne {
	duo.mutation.SetHashedPasswd(s)
	return duo
}

// SetDeadAt sets the "dead_at" field.
func (duo *DeviceUpdateOne) SetDeadAt(t time.Time) *DeviceUpdateOne {
	duo.mutation.SetDeadAt(t)
	return duo
}

// SetNillableDeadAt sets the "dead_at" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableDeadAt(t *time.Time) *DeviceUpdateOne {
	if t != nil {
		duo.SetDeadAt(*t)
	}
	return duo
}

// ClearDeadAt clears the value of the "dead_at" field.
func (duo *DeviceUpdateOne) ClearDeadAt() *DeviceUpdateOne {
	duo.mutation.ClearDeadAt()
	return duo
}

// SetReason sets the "reason" field.
func (duo *DeviceUpdateOne) SetReason(s string) *DeviceUpdateOne {
	duo.mutation.SetReason(s)
	return duo
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableReason(s *string) *DeviceUpdateOne {
	if s != nil {
		duo.SetReason(*s)
	}
	return duo
}

// ClearReason clears the value of the "reason" field.
func (duo *DeviceUpdateOne) ClearReason() *DeviceUpdateOne {
	duo.mutation.ClearReason()
	return duo
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (duo *DeviceUpdateOne) Where(ps ...predicate.Device) *DeviceUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	_spec := sqlgraph.NewUpdateSpec(device.Table, device.Columns, sqlgraph.NewFieldSpec(device.FieldID, field.TypeUUID))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Device.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Parent(); ok {
		_spec.SetField(device.FieldParent, field.TypeString, value)
	}
	if value, ok := duo.mutation.HashedPasswd(); ok {
		_spec.SetField(device.FieldHashedPasswd, field.TypeString, value)
	}
	if value, ok := duo.mutation.DeadAt(); ok {
		_spec.SetField(device.FieldDeadAt, field.TypeTime, value)
	}
	if duo.mutation.DeadAtCleared() {
		_spec.ClearField(device.FieldDeadAt, field.TypeTime)
	}
	if value, ok := duo.mutation.Reason(); ok {
		_spec.SetField(device.FieldReason, field.TypeString, value)
	}
	if duo.mutation.ReasonCleared() {
		_spec.ClearField(device.FieldReason, field.TypeString)
	}
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
