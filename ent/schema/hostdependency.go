package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// HostDependency holds the schema definition for the HostDependency entity.
type HostDependency struct {
	ent.Schema
}

// Fields of the HostDependency.
func (HostDependency) Fields() []ent.Field {
	return []ent.Field{
		field.String("host_id").
			StructTag(`hcl:"host,attr"`),
		field.String("network_id").
			StructTag(`hcl:"network,attr"`),
	}
}

// Edges of the HostDependency.
func (HostDependency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("HostDependencyToHost", Host.Type),
		edge.To("HostDependencyToNetwork", Network.Type),
	}
}