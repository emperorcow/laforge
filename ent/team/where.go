// Code generated by entc, DO NOT EDIT.

package team

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/gen0cide/laforge/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TeamNumber applies equality check predicate on the "team_number" field. It's identical to TeamNumberEQ.
func TeamNumber(v int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeamNumber), v))
	})
}

// Revision applies equality check predicate on the "revision" field. It's identical to RevisionEQ.
func Revision(v int64) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// TeamNumberEQ applies the EQ predicate on the "team_number" field.
func TeamNumberEQ(v int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTeamNumber), v))
	})
}

// TeamNumberNEQ applies the NEQ predicate on the "team_number" field.
func TeamNumberNEQ(v int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTeamNumber), v))
	})
}

// TeamNumberIn applies the In predicate on the "team_number" field.
func TeamNumberIn(vs ...int) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTeamNumber), v...))
	})
}

// TeamNumberNotIn applies the NotIn predicate on the "team_number" field.
func TeamNumberNotIn(vs ...int) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTeamNumber), v...))
	})
}

// TeamNumberGT applies the GT predicate on the "team_number" field.
func TeamNumberGT(v int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTeamNumber), v))
	})
}

// TeamNumberGTE applies the GTE predicate on the "team_number" field.
func TeamNumberGTE(v int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTeamNumber), v))
	})
}

// TeamNumberLT applies the LT predicate on the "team_number" field.
func TeamNumberLT(v int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTeamNumber), v))
	})
}

// TeamNumberLTE applies the LTE predicate on the "team_number" field.
func TeamNumberLTE(v int) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTeamNumber), v))
	})
}

// RevisionEQ applies the EQ predicate on the "revision" field.
func RevisionEQ(v int64) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRevision), v))
	})
}

// RevisionNEQ applies the NEQ predicate on the "revision" field.
func RevisionNEQ(v int64) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRevision), v))
	})
}

// RevisionIn applies the In predicate on the "revision" field.
func RevisionIn(vs ...int64) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRevision), v...))
	})
}

// RevisionNotIn applies the NotIn predicate on the "revision" field.
func RevisionNotIn(vs ...int64) predicate.Team {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Team(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRevision), v...))
	})
}

// RevisionGT applies the GT predicate on the "revision" field.
func RevisionGT(v int64) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRevision), v))
	})
}

// RevisionGTE applies the GTE predicate on the "revision" field.
func RevisionGTE(v int64) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRevision), v))
	})
}

// RevisionLT applies the LT predicate on the "revision" field.
func RevisionLT(v int64) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRevision), v))
	})
}

// RevisionLTE applies the LTE predicate on the "revision" field.
func RevisionLTE(v int64) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRevision), v))
	})
}

// HasMaintainer applies the HasEdge predicate on the "maintainer" edge.
func HasMaintainer() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MaintainerTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MaintainerTable, MaintainerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMaintainerWith applies the HasEdge predicate on the "maintainer" edge with a given conditions (other predicates).
func HasMaintainerWith(preds ...predicate.User) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MaintainerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MaintainerTable, MaintainerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBuild applies the HasEdge predicate on the "build" edge.
func HasBuild() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, BuildTable, BuildPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBuildWith applies the HasEdge predicate on the "build" edge with a given conditions (other predicates).
func HasBuildWith(preds ...predicate.Build) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BuildInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, BuildTable, BuildPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeamToEnvironment applies the HasEdge predicate on the "TeamToEnvironment" edge.
func HasTeamToEnvironment() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeamToEnvironmentTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TeamToEnvironmentTable, TeamToEnvironmentPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamToEnvironmentWith applies the HasEdge predicate on the "TeamToEnvironment" edge with a given conditions (other predicates).
func HasTeamToEnvironmentWith(preds ...predicate.Environment) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TeamToEnvironmentInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TeamToEnvironmentTable, TeamToEnvironmentPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProvisionedNetworks applies the HasEdge predicate on the "provisioned_networks" edge.
func HasProvisionedNetworks() predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProvisionedNetworksTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProvisionedNetworksTable, ProvisionedNetworksPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProvisionedNetworksWith applies the HasEdge predicate on the "provisioned_networks" edge with a given conditions (other predicates).
func HasProvisionedNetworksWith(preds ...predicate.ProvisionedNetwork) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ProvisionedNetworksInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProvisionedNetworksTable, ProvisionedNetworksPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Team) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Team) predicate.Team {
	return predicate.Team(func(s *sql.Selector) {
		p(s.Not())
	})
}
