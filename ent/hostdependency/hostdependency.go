// Code generated by entc, DO NOT EDIT.

package hostdependency

const (
	// Label holds the string label denoting the hostdependency type in the database.
	Label = "host_dependency"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHostID holds the string denoting the host_id field in the database.
	FieldHostID = "host_id"
	// FieldNetworkID holds the string denoting the network_id field in the database.
	FieldNetworkID = "network_id"

	// EdgeHostDependencyToHost holds the string denoting the hostdependencytohost edge name in mutations.
	EdgeHostDependencyToHost = "HostDependencyToHost"
	// EdgeHostDependencyToNetwork holds the string denoting the hostdependencytonetwork edge name in mutations.
	EdgeHostDependencyToNetwork = "HostDependencyToNetwork"

	// Table holds the table name of the hostdependency in the database.
	Table = "host_dependencies"
	// HostDependencyToHostTable is the table the holds the HostDependencyToHost relation/edge. The primary key declared below.
	HostDependencyToHostTable = "host_dependency_HostDependencyToHost"
	// HostDependencyToHostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostDependencyToHostInverseTable = "hosts"
	// HostDependencyToNetworkTable is the table the holds the HostDependencyToNetwork relation/edge. The primary key declared below.
	HostDependencyToNetworkTable = "host_dependency_HostDependencyToNetwork"
	// HostDependencyToNetworkInverseTable is the table name for the Network entity.
	// It exists in this package in order to avoid circular dependency with the "network" package.
	HostDependencyToNetworkInverseTable = "networks"
)

// Columns holds all SQL columns for hostdependency fields.
var Columns = []string{
	FieldID,
	FieldHostID,
	FieldNetworkID,
}

var (
	// HostDependencyToHostPrimaryKey and HostDependencyToHostColumn2 are the table columns denoting the
	// primary key for the HostDependencyToHost relation (M2M).
	HostDependencyToHostPrimaryKey = []string{"host_dependency_id", "host_id"}
	// HostDependencyToNetworkPrimaryKey and HostDependencyToNetworkColumn2 are the table columns denoting the
	// primary key for the HostDependencyToNetwork relation (M2M).
	HostDependencyToNetworkPrimaryKey = []string{"host_dependency_id", "network_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}