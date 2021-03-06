// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/status"
)

// Status is the model entity for the Status schema.
type Status struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// State holds the value of the "state" field.
	State status.State `json:"state,omitempty"`
	// StartedAt holds the value of the "started_at" field.
	StartedAt time.Time `json:"started_at,omitempty"`
	// EndedAt holds the value of the "ended_at" field.
	EndedAt time.Time `json:"ended_at,omitempty"`
	// Failed holds the value of the "failed" field.
	Failed bool `json:"failed,omitempty"`
	// Completed holds the value of the "completed" field.
	Completed bool `json:"completed,omitempty"`
	// Error holds the value of the "error" field.
	Error string `json:"error,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StatusQuery when eager-loading is set.
	Edges                      StatusEdges `json:"edges"`
	provisioned_host_status    *int
	provisioned_network_status *int
	provisioning_step_status   *int
}

// StatusEdges holds the relations/edges for other nodes in the graph.
type StatusEdges struct {
	// Tag holds the value of the tag edge.
	Tag []*Tag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e StatusEdges) TagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Status) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // state
		&sql.NullTime{},   // started_at
		&sql.NullTime{},   // ended_at
		&sql.NullBool{},   // failed
		&sql.NullBool{},   // completed
		&sql.NullString{}, // error
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Status) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // provisioned_host_status
		&sql.NullInt64{}, // provisioned_network_status
		&sql.NullInt64{}, // provisioning_step_status
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Status fields.
func (s *Status) assignValues(values ...interface{}) error {
	if m, n := len(values), len(status.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field state", values[0])
	} else if value.Valid {
		s.State = status.State(value.String)
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field started_at", values[1])
	} else if value.Valid {
		s.StartedAt = value.Time
	}
	if value, ok := values[2].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field ended_at", values[2])
	} else if value.Valid {
		s.EndedAt = value.Time
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field failed", values[3])
	} else if value.Valid {
		s.Failed = value.Bool
	}
	if value, ok := values[4].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field completed", values[4])
	} else if value.Valid {
		s.Completed = value.Bool
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field error", values[5])
	} else if value.Valid {
		s.Error = value.String
	}
	values = values[6:]
	if len(values) == len(status.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field provisioned_host_status", value)
		} else if value.Valid {
			s.provisioned_host_status = new(int)
			*s.provisioned_host_status = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field provisioned_network_status", value)
		} else if value.Valid {
			s.provisioned_network_status = new(int)
			*s.provisioned_network_status = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field provisioning_step_status", value)
		} else if value.Valid {
			s.provisioning_step_status = new(int)
			*s.provisioning_step_status = int(value.Int64)
		}
	}
	return nil
}

// QueryTag queries the tag edge of the Status.
func (s *Status) QueryTag() *TagQuery {
	return (&StatusClient{config: s.config}).QueryTag(s)
}

// Update returns a builder for updating this Status.
// Note that, you need to call Status.Unwrap() before calling this method, if this Status
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Status) Update() *StatusUpdateOne {
	return (&StatusClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Status) Unwrap() *Status {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Status is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Status) String() string {
	var builder strings.Builder
	builder.WriteString("Status(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", s.State))
	builder.WriteString(", started_at=")
	builder.WriteString(s.StartedAt.Format(time.ANSIC))
	builder.WriteString(", ended_at=")
	builder.WriteString(s.EndedAt.Format(time.ANSIC))
	builder.WriteString(", failed=")
	builder.WriteString(fmt.Sprintf("%v", s.Failed))
	builder.WriteString(", completed=")
	builder.WriteString(fmt.Sprintf("%v", s.Completed))
	builder.WriteString(", error=")
	builder.WriteString(s.Error)
	builder.WriteByte(')')
	return builder.String()
}

// StatusSlice is a parsable slice of Status.
type StatusSlice []*Status

func (s StatusSlice) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
