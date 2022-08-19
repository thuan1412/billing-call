package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Call holds the schema definition for the Call entity.
type Call struct {
	ent.Schema
}

// Fields of the Call.
func (Call) Fields() []ent.Field {
	return []ent.Field{
		field.Int("duration"),
		field.Int("block_count"),
	}
}

// Edges of the Call.
func (Call) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("calls"),
	}
}
