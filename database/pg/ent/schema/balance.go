package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

// Balance holds the schema definition for the Balance entity.
type Balance struct {
	ent.Schema
}

// Fields of the Balance.
func (Balance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("when").Immutable().Default(time.Now),
		field.Float("in").Immutable(),
		field.Float("out").Immutable(),
	}
}

// Edges of the Balance.
func (Balance) Edges() []ent.Edge {
	return nil
}
