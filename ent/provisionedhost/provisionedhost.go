// Code generated by entc, DO NOT EDIT.

package provisionedhost

const (
	// Label holds the string label denoting the provisionedhost type in the database.
	Label = "provisioned_host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSubnetIP holds the string denoting the subnet_ip field in the database.
	FieldSubnetIP = "subnet_ip"

	// EdgeStatus holds the string denoting the status edge name in mutations.
	EdgeStatus = "status"
	// EdgeProvisionedNetwork holds the string denoting the provisioned_network edge name in mutations.
	EdgeProvisionedNetwork = "provisioned_network"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// EdgeProvisionedSteps holds the string denoting the provisioned_steps edge name in mutations.
	EdgeProvisionedSteps = "provisioned_steps"

	// Table holds the table name of the provisionedhost in the database.
	Table = "provisioned_hosts"
	// StatusTable is the table the holds the status relation/edge.
	StatusTable = "status"
	// StatusInverseTable is the table name for the Status entity.
	// It exists in this package in order to avoid circular dependency with the "status" package.
	StatusInverseTable = "status"
	// StatusColumn is the table column denoting the status relation/edge.
	StatusColumn = "provisioned_host_status"
	// ProvisionedNetworkTable is the table the holds the provisioned_network relation/edge. The primary key declared below.
	ProvisionedNetworkTable = "provisioned_host_provisioned_network"
	// ProvisionedNetworkInverseTable is the table name for the ProvisionedNetwork entity.
	// It exists in this package in order to avoid circular dependency with the "provisionednetwork" package.
	ProvisionedNetworkInverseTable = "provisioned_networks"
	// HostTable is the table the holds the host relation/edge.
	HostTable = "hosts"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "provisioned_host_host"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "provisioned_host_tag"
	// ProvisionedStepsTable is the table the holds the provisioned_steps relation/edge. The primary key declared below.
	ProvisionedStepsTable = "provisioning_step_provisioned_host"
	// ProvisionedStepsInverseTable is the table name for the ProvisioningStep entity.
	// It exists in this package in order to avoid circular dependency with the "provisioningstep" package.
	ProvisionedStepsInverseTable = "provisioning_steps"
)

// Columns holds all SQL columns for provisionedhost fields.
var Columns = []string{
	FieldID,
	FieldSubnetIP,
}

var (
	// ProvisionedNetworkPrimaryKey and ProvisionedNetworkColumn2 are the table columns denoting the
	// primary key for the provisioned_network relation (M2M).
	ProvisionedNetworkPrimaryKey = []string{"provisioned_host_id", "provisioned_network_id"}
	// ProvisionedStepsPrimaryKey and ProvisionedStepsColumn2 are the table columns denoting the
	// primary key for the provisioned_steps relation (M2M).
	ProvisionedStepsPrimaryKey = []string{"provisioning_step_id", "provisioned_host_id"}
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
