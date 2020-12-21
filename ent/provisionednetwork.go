// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gen0cide/laforge/ent/provisionednetwork"
)

// ProvisionedNetwork is the model entity for the ProvisionedNetwork schema.
type ProvisionedNetwork struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Cidr holds the value of the "cidr" field.
	Cidr string `json:"cidr,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProvisionedNetworkQuery when eager-loading is set.
	Edges ProvisionedNetworkEdges `json:"edges"`
}

// ProvisionedNetworkEdges holds the relations/edges for other nodes in the graph.
type ProvisionedNetworkEdges struct {
	// Status holds the value of the status edge.
	Status []*Status
	// Network holds the value of the network edge.
	Network []*Network
	// Build holds the value of the build edge.
	Build []*Build
	// ProvisionedNetworkToTeam holds the value of the ProvisionedNetworkToTeam edge.
	ProvisionedNetworkToTeam []*Team
	// ProvisionedHosts holds the value of the provisioned_hosts edge.
	ProvisionedHosts []*ProvisionedHost
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// StatusOrErr returns the Status value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedNetworkEdges) StatusOrErr() ([]*Status, error) {
	if e.loadedTypes[0] {
		return e.Status, nil
	}
	return nil, &NotLoadedError{edge: "status"}
}

// NetworkOrErr returns the Network value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedNetworkEdges) NetworkOrErr() ([]*Network, error) {
	if e.loadedTypes[1] {
		return e.Network, nil
	}
	return nil, &NotLoadedError{edge: "network"}
}

// BuildOrErr returns the Build value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedNetworkEdges) BuildOrErr() ([]*Build, error) {
	if e.loadedTypes[2] {
		return e.Build, nil
	}
	return nil, &NotLoadedError{edge: "build"}
}

// ProvisionedNetworkToTeamOrErr returns the ProvisionedNetworkToTeam value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedNetworkEdges) ProvisionedNetworkToTeamOrErr() ([]*Team, error) {
	if e.loadedTypes[3] {
		return e.ProvisionedNetworkToTeam, nil
	}
	return nil, &NotLoadedError{edge: "ProvisionedNetworkToTeam"}
}

// ProvisionedHostsOrErr returns the ProvisionedHosts value or an error if the edge
// was not loaded in eager-loading.
func (e ProvisionedNetworkEdges) ProvisionedHostsOrErr() ([]*ProvisionedHost, error) {
	if e.loadedTypes[4] {
		return e.ProvisionedHosts, nil
	}
	return nil, &NotLoadedError{edge: "provisioned_hosts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProvisionedNetwork) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // cidr
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProvisionedNetwork fields.
func (pn *ProvisionedNetwork) assignValues(values ...interface{}) error {
	if m, n := len(values), len(provisionednetwork.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pn.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		pn.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field cidr", values[1])
	} else if value.Valid {
		pn.Cidr = value.String
	}
	return nil
}

// QueryStatus queries the status edge of the ProvisionedNetwork.
func (pn *ProvisionedNetwork) QueryStatus() *StatusQuery {
	return (&ProvisionedNetworkClient{config: pn.config}).QueryStatus(pn)
}

// QueryNetwork queries the network edge of the ProvisionedNetwork.
func (pn *ProvisionedNetwork) QueryNetwork() *NetworkQuery {
	return (&ProvisionedNetworkClient{config: pn.config}).QueryNetwork(pn)
}

// QueryBuild queries the build edge of the ProvisionedNetwork.
func (pn *ProvisionedNetwork) QueryBuild() *BuildQuery {
	return (&ProvisionedNetworkClient{config: pn.config}).QueryBuild(pn)
}

// QueryProvisionedNetworkToTeam queries the ProvisionedNetworkToTeam edge of the ProvisionedNetwork.
func (pn *ProvisionedNetwork) QueryProvisionedNetworkToTeam() *TeamQuery {
	return (&ProvisionedNetworkClient{config: pn.config}).QueryProvisionedNetworkToTeam(pn)
}

// QueryProvisionedHosts queries the provisioned_hosts edge of the ProvisionedNetwork.
func (pn *ProvisionedNetwork) QueryProvisionedHosts() *ProvisionedHostQuery {
	return (&ProvisionedNetworkClient{config: pn.config}).QueryProvisionedHosts(pn)
}

// Update returns a builder for updating this ProvisionedNetwork.
// Note that, you need to call ProvisionedNetwork.Unwrap() before calling this method, if this ProvisionedNetwork
// was returned from a transaction, and the transaction was committed or rolled back.
func (pn *ProvisionedNetwork) Update() *ProvisionedNetworkUpdateOne {
	return (&ProvisionedNetworkClient{config: pn.config}).UpdateOne(pn)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pn *ProvisionedNetwork) Unwrap() *ProvisionedNetwork {
	tx, ok := pn.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProvisionedNetwork is not a transactional entity")
	}
	pn.config.driver = tx.drv
	return pn
}

// String implements the fmt.Stringer.
func (pn *ProvisionedNetwork) String() string {
	var builder strings.Builder
	builder.WriteString("ProvisionedNetwork(")
	builder.WriteString(fmt.Sprintf("id=%v", pn.ID))
	builder.WriteString(", name=")
	builder.WriteString(pn.Name)
	builder.WriteString(", cidr=")
	builder.WriteString(pn.Cidr)
	builder.WriteByte(')')
	return builder.String()
}

// ProvisionedNetworks is a parsable slice of ProvisionedNetwork.
type ProvisionedNetworks []*ProvisionedNetwork

func (pn ProvisionedNetworks) config(cfg config) {
	for _i := range pn {
		pn[_i].config = cfg
	}
}
