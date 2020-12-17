// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/network"
)

// Network is the model entity for the Network schema.
type Network struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Cidr holds the value of the "cidr" field.
	Cidr string `json:"cidr,omitempty"`
	// VdiVisible holds the value of the "vdi_visible" field.
	VdiVisible bool `json:"vdi_visible,omitempty"`
	// Vars holds the value of the "vars" field.
	Vars []string `json:"vars,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NetworkQuery when eager-loading is set.
	Edges                       NetworkEdges `json:"edges"`
	provisioned_network_network *int
}

// NetworkEdges holds the relations/edges for other nodes in the graph.
type NetworkEdges struct {
	// Tag holds the value of the tag edge.
	Tag []*Tag
	// NetworkToEnvironment holds the value of the NetworkToEnvironment edge.
	NetworkToEnvironment []*Environment
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TagOrErr returns the Tag value or an error if the edge
// was not loaded in eager-loading.
func (e NetworkEdges) TagOrErr() ([]*Tag, error) {
	if e.loadedTypes[0] {
		return e.Tag, nil
	}
	return nil, &NotLoadedError{edge: "tag"}
}

// NetworkToEnvironmentOrErr returns the NetworkToEnvironment value or an error if the edge
// was not loaded in eager-loading.
func (e NetworkEdges) NetworkToEnvironmentOrErr() ([]*Environment, error) {
	if e.loadedTypes[1] {
		return e.NetworkToEnvironment, nil
	}
	return nil, &NotLoadedError{edge: "NetworkToEnvironment"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Network) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // cidr
		&sql.NullBool{},   // vdi_visible
		&[]byte{},         // vars
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Network) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // provisioned_network_network
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Network fields.
func (n *Network) assignValues(values ...interface{}) error {
	if m, n := len(values), len(network.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	n.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		n.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field cidr", values[1])
	} else if value.Valid {
		n.Cidr = value.String
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field vdi_visible", values[2])
	} else if value.Valid {
		n.VdiVisible = value.Bool
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field vars", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &n.Vars); err != nil {
			return fmt.Errorf("unmarshal field vars: %v", err)
		}
	}
	values = values[4:]
	if len(values) == len(network.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field provisioned_network_network", value)
		} else if value.Valid {
			n.provisioned_network_network = new(int)
			*n.provisioned_network_network = int(value.Int64)
		}
	}
	return nil
}

// QueryTag queries the tag edge of the Network.
func (n *Network) QueryTag() *TagQuery {
	return (&NetworkClient{config: n.config}).QueryTag(n)
}

// QueryNetworkToEnvironment queries the NetworkToEnvironment edge of the Network.
func (n *Network) QueryNetworkToEnvironment() *EnvironmentQuery {
	return (&NetworkClient{config: n.config}).QueryNetworkToEnvironment(n)
}

// Update returns a builder for updating this Network.
// Note that, you need to call Network.Unwrap() before calling this method, if this Network
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Network) Update() *NetworkUpdateOne {
	return (&NetworkClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (n *Network) Unwrap() *Network {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Network is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Network) String() string {
	var builder strings.Builder
	builder.WriteString("Network(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", name=")
	builder.WriteString(n.Name)
	builder.WriteString(", cidr=")
	builder.WriteString(n.Cidr)
	builder.WriteString(", vdi_visible=")
	builder.WriteString(fmt.Sprintf("%v", n.VdiVisible))
	builder.WriteString(", vars=")
	builder.WriteString(fmt.Sprintf("%v", n.Vars))
	builder.WriteByte(')')
	return builder.String()
}

// Networks is a parsable slice of Network.
type Networks []*Network

func (n Networks) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
