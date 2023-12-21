// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/device"
	"github.com/google/uuid"
)

// Device is the model entity for the Device schema.
type Device struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// BornAt holds the value of the "born_at" field.
	BornAt time.Time `json:"born_at,omitempty"`
	// HashedPasswd holds the value of the "hashed_passwd" field.
	HashedPasswd string `json:"hashed_passwd,omitempty"`
	// DeadAt holds the value of the "dead_at" field.
	DeadAt *time.Time `json:"dead_at,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason *string `json:"reason,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges           DeviceEdges `json:"edges"`
	device_children *uuid.UUID
	selectValues    sql.SelectValues
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Device `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Device `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) ParentOrErr() (*Device, error) {
	if e.loadedTypes[0] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: device.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) ChildrenOrErr() ([]*Device, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldHashedPasswd, device.FieldReason:
			values[i] = new(sql.NullString)
		case device.FieldBornAt, device.FieldDeadAt:
			values[i] = new(sql.NullTime)
		case device.FieldID:
			values[i] = new(uuid.UUID)
		case device.ForeignKeys[0]: // device_children
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case device.FieldBornAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field born_at", values[i])
			} else if value.Valid {
				d.BornAt = value.Time
			}
		case device.FieldHashedPasswd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hashed_passwd", values[i])
			} else if value.Valid {
				d.HashedPasswd = value.String
			}
		case device.FieldDeadAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dead_at", values[i])
			} else if value.Valid {
				d.DeadAt = new(time.Time)
				*d.DeadAt = value.Time
			}
		case device.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				d.Reason = new(string)
				*d.Reason = value.String
			}
		case device.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field device_children", values[i])
			} else if value.Valid {
				d.device_children = new(uuid.UUID)
				*d.device_children = *value.S.(*uuid.UUID)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Device.
// This includes values selected through modifiers, order, etc.
func (d *Device) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Device entity.
func (d *Device) QueryParent() *DeviceQuery {
	return NewDeviceClient(d.config).QueryParent(d)
}

// QueryChildren queries the "children" edge of the Device entity.
func (d *Device) QueryChildren() *DeviceQuery {
	return NewDeviceClient(d.config).QueryChildren(d)
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return NewDeviceClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("born_at=")
	builder.WriteString(d.BornAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("hashed_passwd=")
	builder.WriteString(d.HashedPasswd)
	builder.WriteString(", ")
	if v := d.DeadAt; v != nil {
		builder.WriteString("dead_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := d.Reason; v != nil {
		builder.WriteString("reason=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device
