package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("hashed_passwd").Sensitive(),
		field.Time("dead_at").Optional().Nillable(),
		field.String("reason").Optional(),
		field.UUID("parent", uuid.UUID{}).Optional(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Device.Type).From("mother").Field("parent").Unique(),
	}
}

func (Device) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("born_at"),
		index.Fields("dead_at"),
	}
}
