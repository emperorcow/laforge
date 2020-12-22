// Code generated by entc, DO NOT EDIT.

package filedownload

const (
	// Label holds the string label denoting the filedownload type in the database.
	Label = "file_download"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSourceType holds the string denoting the source_type field in the database.
	FieldSourceType = "source_type"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldDestination holds the string denoting the destination field in the database.
	FieldDestination = "destination"
	// FieldTemplate holds the string denoting the template field in the database.
	FieldTemplate = "template"
	// FieldMode holds the string denoting the mode field in the database.
	FieldMode = "mode"
	// FieldDisabled holds the string denoting the disabled field in the database.
	FieldDisabled = "disabled"
	// FieldMd5 holds the string denoting the md5 field in the database.
	FieldMd5 = "md5"
	// FieldAbsPath holds the string denoting the abs_path field in the database.
	FieldAbsPath = "abs_path"

	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"

	// Table holds the table name of the filedownload in the database.
	Table = "file_downloads"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "file_download_tag"
)

// Columns holds all SQL columns for filedownload fields.
var Columns = []string{
	FieldID,
	FieldSourceType,
	FieldSource,
	FieldDestination,
	FieldTemplate,
	FieldMode,
	FieldDisabled,
	FieldMd5,
	FieldAbsPath,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}