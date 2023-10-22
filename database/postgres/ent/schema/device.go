package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.Time("born_at").Immutable(),
		field.UUID("id", uuid.UUID{}).Immutable().Unique(),
		field.String("hashed_passwd"),
		field.Time("dead_at").Optional().Nillable(),
		field.String("reason").Optional().Nillable(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return nil
}
