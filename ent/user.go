// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID string `json:"uuid,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges             UserEdges `json:"edges"`
	build_maintainer  *int
	command_user      *int
	environment_user  *int
	finding_user      *int
	host_maintainer   *int
	script_maintainer *int
	team_maintainer   *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Tag holds the value of the tag edge.
	Tag []*Tag
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) TagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // uuid
		&sql.NullString{}, // email
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*User) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // build_maintainer
		&sql.NullInt64{}, // command_user
		&sql.NullInt64{}, // environment_user
		&sql.NullInt64{}, // finding_user
		&sql.NullInt64{}, // host_maintainer
		&sql.NullInt64{}, // script_maintainer
		&sql.NullInt64{}, // team_maintainer
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field uuid", values[1])
	} else if value.Valid {
		u.UUID = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field email", values[2])
	} else if value.Valid {
		u.Email = value.String
	}
	values = values[3:]
	if len(values) == len(user.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field build_maintainer", value)
		} else if value.Valid {
			u.build_maintainer = new(int)
			*u.build_maintainer = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field command_user", value)
		} else if value.Valid {
			u.command_user = new(int)
			*u.command_user = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field environment_user", value)
		} else if value.Valid {
			u.environment_user = new(int)
			*u.environment_user = int(value.Int64)
		}
		if value, ok := values[3].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field finding_user", value)
		} else if value.Valid {
			u.finding_user = new(int)
			*u.finding_user = int(value.Int64)
		}
		if value, ok := values[4].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field host_maintainer", value)
		} else if value.Valid {
			u.host_maintainer = new(int)
			*u.host_maintainer = int(value.Int64)
		}
		if value, ok := values[5].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field script_maintainer", value)
		} else if value.Valid {
			u.script_maintainer = new(int)
			*u.script_maintainer = int(value.Int64)
		}
		if value, ok := values[6].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field team_maintainer", value)
		} else if value.Valid {
			u.team_maintainer = new(int)
			*u.team_maintainer = int(value.Int64)
		}
	}
	return nil
}

// QueryTag queries the tag edge of the User.
func (u *User) QueryTag() *TagQuery {
	return (&UserClient{config: u.config}).QueryTag(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", uuid=")
	builder.WriteString(u.UUID)
	builder.WriteString(", email=")
	builder.WriteString(u.Email)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
