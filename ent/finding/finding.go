// Code generated by entc, DO NOT EDIT.

package finding

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the finding type in the database.
	Label = "finding"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSeverity holds the string denoting the severity field in the database.
	FieldSeverity = "severity"
	// FieldDifficulty holds the string denoting the difficulty field in the database.
	FieldDifficulty = "difficulty"

	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// EdgeHost holds the string denoting the host edge name in mutations.
	EdgeHost = "host"
	// EdgeScript holds the string denoting the script edge name in mutations.
	EdgeScript = "script"

	// Table holds the table name of the finding in the database.
	Table = "findings"
	// UserTable is the table the holds the user relation/edge.
	UserTable = "users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "finding_user"
	// TagTable is the table the holds the tag relation/edge.
	TagTable = "tags"
	// TagInverseTable is the table name for the Tag entity.
	// It exists in this package in order to avoid circular dependency with the "tag" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "finding_tag"
	// HostTable is the table the holds the host relation/edge.
	HostTable = "hosts"
	// HostInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	HostInverseTable = "hosts"
	// HostColumn is the table column denoting the host relation/edge.
	HostColumn = "finding_host"
	// ScriptTable is the table the holds the script relation/edge. The primary key declared below.
	ScriptTable = "finding_script"
	// ScriptInverseTable is the table name for the Script entity.
	// It exists in this package in order to avoid circular dependency with the "script" package.
	ScriptInverseTable = "scripts"
)

// Columns holds all SQL columns for finding fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldSeverity,
	FieldDifficulty,
}

var (
	// ScriptPrimaryKey and ScriptColumn2 are the table columns denoting the
	// primary key for the script relation (M2M).
	ScriptPrimaryKey = []string{"finding_id", "script_id"}
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

// Severity defines the type for the severity enum field.
type Severity string

// Severity values.
const (
	SeverityZeroSeverity     Severity = "ZeroSeverity"
	SeverityLowSeverity      Severity = "LowSeverity"
	SeverityMediumSeverity   Severity = "MediumSeverity"
	SeverityHighSeverity     Severity = "HighSeverity"
	SeverityCriticalSeverity Severity = "CriticalSeverity"
	SeverityNullSeverity     Severity = "NullSeverity"
)

func (s Severity) String() string {
	return string(s)
}

// SeverityValidator is a validator for the "severity" field enum values. It is called by the builders before save.
func SeverityValidator(s Severity) error {
	switch s {
	case SeverityZeroSeverity, SeverityLowSeverity, SeverityMediumSeverity, SeverityHighSeverity, SeverityCriticalSeverity, SeverityNullSeverity:
		return nil
	default:
		return fmt.Errorf("finding: invalid enum value for severity field: %q", s)
	}
}

// Difficulty defines the type for the difficulty enum field.
type Difficulty string

// Difficulty values.
const (
	DifficultyZeroDifficulty     Difficulty = "ZeroDifficulty"
	DifficultyNoviceDifficulty   Difficulty = "NoviceDifficulty"
	DifficultyAdvancedDifficulty Difficulty = "AdvancedDifficulty"
	DifficultyExpertDifficulty   Difficulty = "ExpertDifficulty"
	DifficultyNullDifficulty     Difficulty = "NullDifficulty"
)

func (d Difficulty) String() string {
	return string(d)
}

// DifficultyValidator is a validator for the "difficulty" field enum values. It is called by the builders before save.
func DifficultyValidator(d Difficulty) error {
	switch d {
	case DifficultyZeroDifficulty, DifficultyNoviceDifficulty, DifficultyAdvancedDifficulty, DifficultyExpertDifficulty, DifficultyNullDifficulty:
		return nil
	default:
		return fmt.Errorf("finding: invalid enum value for difficulty field: %q", d)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (s Severity) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(s.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (s *Severity) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*s = Severity(str)
	if err := SeverityValidator(*s); err != nil {
		return fmt.Errorf("%s is not a valid Severity", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (d Difficulty) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(d.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (d *Difficulty) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*d = Difficulty(str)
	if err := DifficultyValidator(*d); err != nil {
		return fmt.Errorf("%s is not a valid Difficulty", str)
	}
	return nil
}
