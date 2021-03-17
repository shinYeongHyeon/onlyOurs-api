package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Users holds the schema definition for the Users entity.
type Users struct {
	ent.Schema
}

// Fields of the Users.
func (Users) Fields() []ent.Field {
	return []ent.Field {
		field.Int("id").StorageKey("uuid"),
		field.String("user_id").Default(""),
		field.String("name").Default(""),
		field.String("password").Default(""),
	}
}

// Edges of the Users.
func (Users) Edges() []ent.Edge {
	return nil
}
