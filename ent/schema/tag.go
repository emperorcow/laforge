package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/google/uuid"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("uuid", uuid.UUID{}),
		field.String("name"),
		field.JSON("description", map[string]string{}),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
