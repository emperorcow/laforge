// Code generated by entc, DO NOT EDIT.

package dns

const (
	// Label holds the string label denoting the dns type in the database.
	Label = "dns"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldRootDomain holds the string denoting the root_domain field in the database.
	FieldRootDomain = "root_domain"
	// FieldDNSServers holds the string denoting the dns_servers field in the database.
	FieldDNSServers = "dns_servers"
	// FieldNtpServers holds the string denoting the ntp_servers field in the database.
	FieldNtpServers = "ntp_servers"
	// FieldConfig holds the string denoting the config field in the database.
	FieldConfig = "config"

	// Table holds the table name of the dns in the database.
	Table = "dn_ss"
)

// Columns holds all SQL columns for dns fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldRootDomain,
	FieldDNSServers,
	FieldNtpServers,
	FieldConfig,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the DNS type.
var ForeignKeys = []string{
	"competition_dns",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
