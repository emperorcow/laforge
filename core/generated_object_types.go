// Code generated by renum (github.com/gen0cide/renum)
// DO NOT EDIT!
// renum v0.0.6-b3fac17

package core

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// ObjectType is a generated type alias for the ObjectType enum.
type ObjectType int

const (
	// ObjectTypeUndefinedEnumValue is an enum value for type ObjectType.
	// ObjectTypeUndefinedEnumValue is the default value for enum type ObjectType. It is meant to be a placeholder and default for unknown values.
	// This value is a default placeholder for any unknown type for the core.ObjectType enum.
	ObjectTypeUndefinedEnumValue ObjectType = iota

	// ObjectTypeBuild is an enum value for type ObjectType.
	// Build is a type of Laforge object that describes the current and goal state of an environment configuration as realized by a builder implementation.
	ObjectTypeBuild

	// ObjectTypeCompetition is an enum value for type ObjectType.
	// Competition is a type of Laforge object that describes the root configuration object for a Laforge configuration directory.
	ObjectTypeCompetition

	// ObjectTypeCommand is an enum value for type ObjectType.
	// Command is a type of Laforge object that describes a single command to be run during provisioning.
	ObjectTypeCommand

	// ObjectTypeDNSRecord is an enum value for type ObjectType.
	// DNSRecord is a type of Laforge object that describes a DNS record that needs to be created in the root DNS of the environment.
	ObjectTypeDNSRecord

	// ObjectTypeEnvironment is an enum value for type ObjectType.
	// Environment is a type of Laforge object that describes a single, atomic configuration with an isolated state.
	ObjectTypeEnvironment

	// ObjectTypeHost is an enum value for type ObjectType.
	// Host is a type of Laforge object that describes a buildable configuration of a machine.
	ObjectTypeHost

	// ObjectTypeIdentity is an enum value for type ObjectType.
	// Identity is a type of Laforge object that describes a user, employee, or account on the built network.
	ObjectTypeIdentity

	// ObjectTypeNetwork is an enum value for type ObjectType.
	// Network is a type of Laforge object that describes a buildable subnet within an environment, containing many Hosts.
	ObjectTypeNetwork

	// ObjectTypeRemoteFile is an enum value for type ObjectType.
	// RemoteFile is a type of Laforge object that describes a provisioning step where a file is uploaded to the target Host.
	ObjectTypeRemoteFile

	// ObjectTypeScript is an enum value for type ObjectType.
	// Script is a type of Laforge object that describes a provisioning step where a script is uploaded and executed on the target Host.
	ObjectTypeScript

	// ObjectTypeTeam is an enum value for type ObjectType.
	// Team is a type of Laforge object that defines an isolated pod that exists within an environment, by which all network and host configurations are cloned.
	ObjectTypeTeam

	// ObjectTypeUser is an enum value for type ObjectType.
	// User is a type of Laforge object that defines an operator of the Laforge framework.
	ObjectTypeUser

	// ObjectTypeAMI is an enum value for type ObjectType.
	// AMI is a type of Laforge object that defines a reference to a custom image within an infrastructure provider.
	ObjectTypeAMI

	// ObjectTypeProvisionedHost is an enum value for type ObjectType.
	// ProvisionedHost is a type of Laforge object that defines a deployed Host configuration, contained within an Environment's ecology.
	ObjectTypeProvisionedHost

	// ObjectTypeProvisionedNetwork is an enum value for type ObjectType.
	// ProvisionedNetwork is a type of Laforge object that defines a deployed Network configuration, contained within an Environment ecology.
	ObjectTypeProvisionedNetwork

	// ObjectTypeProvisioningStep is an enum value for type ObjectType.
	// ProvisioningStep is a type of Laforge object that defines an action Laforge is required to take to advance the a build towards the desired state.
	ObjectTypeProvisioningStep

	// ObjectTypeConnection is an enum value for type ObjectType.
	// Connection is a type of Laforge object that defines the parameters by which the Laforge provisioner can use to make a remote connection to a provisioned host.
	ObjectTypeConnection

	// ObjectTypeIncluded is an enum value for type ObjectType.
	// Included is a classification of Laforge objects that help the compiler understand if the what hosts and networks should be included in an environment.
	ObjectTypeIncluded

	_ObjectTypeNamespace = `github.com.gen0cide.laforge.core`
	_ObjectTypePkgName   = `core`
	_ObjectTypePkgPath   = `github.com/gen0cide/laforge/core`
)

var (
	// ErrUndefinedObjectTypeEnumValue is thrown when ParseObjectType(s string) cannot locate a valid enum for the provided string.
	ErrUndefinedObjectTypeEnumValue = errors.New("cannot identify enum for provided value")
)

const _ObjectTypeName = "undefined_enum_valuebuildcompetitioncommanddns_recordenvironmenthostidentitynetworkremote_filescriptteamuseramiprovisioned_hostprovisioned_networkprovisioning_stepconnectionincluded"

const _ObjectTypePascalName = "UndefinedEnumValueBuildCompetitionCommandDNSRecordEnvironmentHostIdentityNetworkRemoteFileScriptTeamUserAMIProvisionedHostProvisionedNetworkProvisioningStepConnectionIncluded"

const _ObjectTypeCamelName = "undefinedEnumValuebuildcompetitioncommanddNSRecordenvironmenthostidentitynetworkremoteFilescriptteamuseramiprovisionedHostprovisionedNetworkprovisioningStepconnectionincluded"

const _ObjectTypeScreamingName = "UNDEFINED_ENUM_VALUEBUILDCOMPETITIONCOMMANDDNS_RECORDENVIRONMENTHOSTIDENTITYNETWORKREMOTE_FILESCRIPTTEAMUSERAMIPROVISIONED_HOSTPROVISIONED_NETWORKPROVISIONING_STEPCONNECTIONINCLUDED"

const _ObjectTypeCommandName = "undefined-enum-valuebuildcompetitioncommanddns-recordenvironmenthostidentitynetworkremote-filescriptteamuseramiprovisioned-hostprovisioned-networkprovisioning-stepconnectionincluded"

var _ObjectTypeNames = []string{
	_ObjectTypeName[0:20],
	_ObjectTypeName[20:25],
	_ObjectTypeName[25:36],
	_ObjectTypeName[36:43],
	_ObjectTypeName[43:53],
	_ObjectTypeName[53:64],
	_ObjectTypeName[64:68],
	_ObjectTypeName[68:76],
	_ObjectTypeName[76:83],
	_ObjectTypeName[83:94],
	_ObjectTypeName[94:100],
	_ObjectTypeName[100:104],
	_ObjectTypeName[104:108],
	_ObjectTypeName[108:111],
	_ObjectTypeName[111:127],
	_ObjectTypeName[127:146],
	_ObjectTypeName[146:163],
	_ObjectTypeName[163:173],
	_ObjectTypeName[173:181],
}

// ObjectTypeNames returns a list of possible string values of ObjectType.
func ObjectTypeNames() []string {
	tmp := make([]string, len(_ObjectTypeNames))
	copy(tmp, _ObjectTypeNames)
	return tmp
}

var _ObjectTypeValueSlice = []ObjectType{
	ObjectTypeUndefinedEnumValue,
	ObjectTypeBuild,
	ObjectTypeCompetition,
	ObjectTypeCommand,
	ObjectTypeDNSRecord,
	ObjectTypeEnvironment,
	ObjectTypeHost,
	ObjectTypeIdentity,
	ObjectTypeNetwork,
	ObjectTypeRemoteFile,
	ObjectTypeScript,
	ObjectTypeTeam,
	ObjectTypeUser,
	ObjectTypeAMI,
	ObjectTypeProvisionedHost,
	ObjectTypeProvisionedNetwork,
	ObjectTypeProvisioningStep,
	ObjectTypeConnection,
	ObjectTypeIncluded,
}

// ObjectTypeValues returns a list of possible enum values for the ObjectType type.
func ObjectTypeValues() []ObjectType {
	tmp := make([]ObjectType, len(_ObjectTypeValueSlice))
	copy(tmp, _ObjectTypeValueSlice)
	return tmp
}

var _ObjectTypeValue = map[string]ObjectType{
	_ObjectTypeName[0:20]:    0,
	_ObjectTypeName[20:25]:   1,
	_ObjectTypeName[25:36]:   2,
	_ObjectTypeName[36:43]:   3,
	_ObjectTypeName[43:53]:   4,
	_ObjectTypeName[53:64]:   5,
	_ObjectTypeName[64:68]:   6,
	_ObjectTypeName[68:76]:   7,
	_ObjectTypeName[76:83]:   8,
	_ObjectTypeName[83:94]:   9,
	_ObjectTypeName[94:100]:  10,
	_ObjectTypeName[100:104]: 11,
	_ObjectTypeName[104:108]: 12,
	_ObjectTypeName[108:111]: 13,
	_ObjectTypeName[111:127]: 14,
	_ObjectTypeName[127:146]: 15,
	_ObjectTypeName[146:163]: 16,
	_ObjectTypeName[163:173]: 17,
	_ObjectTypeName[173:181]: 18,
}

var _ObjectTypePascalValue = map[string]ObjectType{
	_ObjectTypePascalName[0:18]:    0,
	_ObjectTypePascalName[18:23]:   1,
	_ObjectTypePascalName[23:34]:   2,
	_ObjectTypePascalName[34:41]:   3,
	_ObjectTypePascalName[41:50]:   4,
	_ObjectTypePascalName[50:61]:   5,
	_ObjectTypePascalName[61:65]:   6,
	_ObjectTypePascalName[65:73]:   7,
	_ObjectTypePascalName[73:80]:   8,
	_ObjectTypePascalName[80:90]:   9,
	_ObjectTypePascalName[90:96]:   10,
	_ObjectTypePascalName[96:100]:  11,
	_ObjectTypePascalName[100:104]: 12,
	_ObjectTypePascalName[104:107]: 13,
	_ObjectTypePascalName[107:122]: 14,
	_ObjectTypePascalName[122:140]: 15,
	_ObjectTypePascalName[140:156]: 16,
	_ObjectTypePascalName[156:166]: 17,
	_ObjectTypePascalName[166:174]: 18,
}

var _ObjectTypeCamelValue = map[string]ObjectType{
	_ObjectTypeCamelName[0:18]:    0,
	_ObjectTypeCamelName[18:23]:   1,
	_ObjectTypeCamelName[23:34]:   2,
	_ObjectTypeCamelName[34:41]:   3,
	_ObjectTypeCamelName[41:50]:   4,
	_ObjectTypeCamelName[50:61]:   5,
	_ObjectTypeCamelName[61:65]:   6,
	_ObjectTypeCamelName[65:73]:   7,
	_ObjectTypeCamelName[73:80]:   8,
	_ObjectTypeCamelName[80:90]:   9,
	_ObjectTypeCamelName[90:96]:   10,
	_ObjectTypeCamelName[96:100]:  11,
	_ObjectTypeCamelName[100:104]: 12,
	_ObjectTypeCamelName[104:107]: 13,
	_ObjectTypeCamelName[107:122]: 14,
	_ObjectTypeCamelName[122:140]: 15,
	_ObjectTypeCamelName[140:156]: 16,
	_ObjectTypeCamelName[156:166]: 17,
	_ObjectTypeCamelName[166:174]: 18,
}

var _ObjectTypeScreamingValue = map[string]ObjectType{
	_ObjectTypeScreamingName[0:20]:    0,
	_ObjectTypeScreamingName[20:25]:   1,
	_ObjectTypeScreamingName[25:36]:   2,
	_ObjectTypeScreamingName[36:43]:   3,
	_ObjectTypeScreamingName[43:53]:   4,
	_ObjectTypeScreamingName[53:64]:   5,
	_ObjectTypeScreamingName[64:68]:   6,
	_ObjectTypeScreamingName[68:76]:   7,
	_ObjectTypeScreamingName[76:83]:   8,
	_ObjectTypeScreamingName[83:94]:   9,
	_ObjectTypeScreamingName[94:100]:  10,
	_ObjectTypeScreamingName[100:104]: 11,
	_ObjectTypeScreamingName[104:108]: 12,
	_ObjectTypeScreamingName[108:111]: 13,
	_ObjectTypeScreamingName[111:127]: 14,
	_ObjectTypeScreamingName[127:146]: 15,
	_ObjectTypeScreamingName[146:163]: 16,
	_ObjectTypeScreamingName[163:173]: 17,
	_ObjectTypeScreamingName[173:181]: 18,
}

var _ObjectTypeCommandValue = map[string]ObjectType{
	_ObjectTypeCommandName[0:20]:    0,
	_ObjectTypeCommandName[20:25]:   1,
	_ObjectTypeCommandName[25:36]:   2,
	_ObjectTypeCommandName[36:43]:   3,
	_ObjectTypeCommandName[43:53]:   4,
	_ObjectTypeCommandName[53:64]:   5,
	_ObjectTypeCommandName[64:68]:   6,
	_ObjectTypeCommandName[68:76]:   7,
	_ObjectTypeCommandName[76:83]:   8,
	_ObjectTypeCommandName[83:94]:   9,
	_ObjectTypeCommandName[94:100]:  10,
	_ObjectTypeCommandName[100:104]: 11,
	_ObjectTypeCommandName[104:108]: 12,
	_ObjectTypeCommandName[108:111]: 13,
	_ObjectTypeCommandName[111:127]: 14,
	_ObjectTypeCommandName[127:146]: 15,
	_ObjectTypeCommandName[146:163]: 16,
	_ObjectTypeCommandName[163:173]: 17,
	_ObjectTypeCommandName[173:181]: 18,
}

var _ObjectTypeMap = map[ObjectType]string{
	0:  _ObjectTypeName[0:20],
	1:  _ObjectTypeName[20:25],
	2:  _ObjectTypeName[25:36],
	3:  _ObjectTypeName[36:43],
	4:  _ObjectTypeName[43:53],
	5:  _ObjectTypeName[53:64],
	6:  _ObjectTypeName[64:68],
	7:  _ObjectTypeName[68:76],
	8:  _ObjectTypeName[76:83],
	9:  _ObjectTypeName[83:94],
	10: _ObjectTypeName[94:100],
	11: _ObjectTypeName[100:104],
	12: _ObjectTypeName[104:108],
	13: _ObjectTypeName[108:111],
	14: _ObjectTypeName[111:127],
	15: _ObjectTypeName[127:146],
	16: _ObjectTypeName[146:163],
	17: _ObjectTypeName[163:173],
	18: _ObjectTypeName[173:181],
}

var _ObjectTypePascalMap = map[ObjectType]string{
	0:  _ObjectTypePascalName[0:18],
	1:  _ObjectTypePascalName[18:23],
	2:  _ObjectTypePascalName[23:34],
	3:  _ObjectTypePascalName[34:41],
	4:  _ObjectTypePascalName[41:50],
	5:  _ObjectTypePascalName[50:61],
	6:  _ObjectTypePascalName[61:65],
	7:  _ObjectTypePascalName[65:73],
	8:  _ObjectTypePascalName[73:80],
	9:  _ObjectTypePascalName[80:90],
	10: _ObjectTypePascalName[90:96],
	11: _ObjectTypePascalName[96:100],
	12: _ObjectTypePascalName[100:104],
	13: _ObjectTypePascalName[104:107],
	14: _ObjectTypePascalName[107:122],
	15: _ObjectTypePascalName[122:140],
	16: _ObjectTypePascalName[140:156],
	17: _ObjectTypePascalName[156:166],
	18: _ObjectTypePascalName[166:174],
}

var _ObjectTypeCamelMap = map[ObjectType]string{
	0:  _ObjectTypeCamelName[0:18],
	1:  _ObjectTypeCamelName[18:23],
	2:  _ObjectTypeCamelName[23:34],
	3:  _ObjectTypeCamelName[34:41],
	4:  _ObjectTypeCamelName[41:50],
	5:  _ObjectTypeCamelName[50:61],
	6:  _ObjectTypeCamelName[61:65],
	7:  _ObjectTypeCamelName[65:73],
	8:  _ObjectTypeCamelName[73:80],
	9:  _ObjectTypeCamelName[80:90],
	10: _ObjectTypeCamelName[90:96],
	11: _ObjectTypeCamelName[96:100],
	12: _ObjectTypeCamelName[100:104],
	13: _ObjectTypeCamelName[104:107],
	14: _ObjectTypeCamelName[107:122],
	15: _ObjectTypeCamelName[122:140],
	16: _ObjectTypeCamelName[140:156],
	17: _ObjectTypeCamelName[156:166],
	18: _ObjectTypeCamelName[166:174],
}

var _ObjectTypeScreamingMap = map[ObjectType]string{
	0:  _ObjectTypeScreamingName[0:20],
	1:  _ObjectTypeScreamingName[20:25],
	2:  _ObjectTypeScreamingName[25:36],
	3:  _ObjectTypeScreamingName[36:43],
	4:  _ObjectTypeScreamingName[43:53],
	5:  _ObjectTypeScreamingName[53:64],
	6:  _ObjectTypeScreamingName[64:68],
	7:  _ObjectTypeScreamingName[68:76],
	8:  _ObjectTypeScreamingName[76:83],
	9:  _ObjectTypeScreamingName[83:94],
	10: _ObjectTypeScreamingName[94:100],
	11: _ObjectTypeScreamingName[100:104],
	12: _ObjectTypeScreamingName[104:108],
	13: _ObjectTypeScreamingName[108:111],
	14: _ObjectTypeScreamingName[111:127],
	15: _ObjectTypeScreamingName[127:146],
	16: _ObjectTypeScreamingName[146:163],
	17: _ObjectTypeScreamingName[163:173],
	18: _ObjectTypeScreamingName[173:181],
}

var _ObjectTypeCommandMap = map[ObjectType]string{
	0:  _ObjectTypeCommandName[0:20],
	1:  _ObjectTypeCommandName[20:25],
	2:  _ObjectTypeCommandName[25:36],
	3:  _ObjectTypeCommandName[36:43],
	4:  _ObjectTypeCommandName[43:53],
	5:  _ObjectTypeCommandName[53:64],
	6:  _ObjectTypeCommandName[64:68],
	7:  _ObjectTypeCommandName[68:76],
	8:  _ObjectTypeCommandName[76:83],
	9:  _ObjectTypeCommandName[83:94],
	10: _ObjectTypeCommandName[94:100],
	11: _ObjectTypeCommandName[100:104],
	12: _ObjectTypeCommandName[104:108],
	13: _ObjectTypeCommandName[108:111],
	14: _ObjectTypeCommandName[111:127],
	15: _ObjectTypeCommandName[127:146],
	16: _ObjectTypeCommandName[146:163],
	17: _ObjectTypeCommandName[163:173],
	18: _ObjectTypeCommandName[173:181],
}

// String implements the Stringer interface.
func (x ObjectType) String() string {
	if str, ok := _ObjectTypeMap[x]; ok {
		return str
	}

	return _ObjectTypeMap[ObjectType(0)]
}

// SnakeCase returns the enum as a snake_case string.
func (x ObjectType) SnakeCase() string {
	if str, ok := _ObjectTypeMap[x]; ok {
		return str
	}

	return _ObjectTypeMap[ObjectType(0)]
}

// PascalCase returns the enum as a PascalCase string.
func (x ObjectType) PascalCase() string {
	if str, ok := _ObjectTypePascalMap[x]; ok {
		return str
	}

	return _ObjectTypePascalMap[ObjectType(0)]
}

// CamelCase returns the enum as a cascalCase string.
func (x ObjectType) CamelCase() string {
	if str, ok := _ObjectTypeCamelMap[x]; ok {
		return str
	}

	return _ObjectTypeCamelMap[ObjectType(0)]
}

// ScreamingCase returns the enum as a SCREAMING_CASE string.
func (x ObjectType) ScreamingCase() string {
	if str, ok := _ObjectTypeScreamingMap[x]; ok {
		return str
	}

	return _ObjectTypeScreamingMap[ObjectType(0)]
}

// CommandCase returns the enum as a command-case string.
func (x ObjectType) CommandCase() string {
	if str, ok := _ObjectTypeCommandMap[x]; ok {
		return str
	}

	return _ObjectTypeCommandMap[ObjectType(0)]
}

var _ObjectTypeKinds = map[ObjectType]string{
	ObjectTypeUndefinedEnumValue: `core.ObjectTypeUndefinedEnumValue`,
	ObjectTypeBuild:              `core.ObjectTypeBuild`,
	ObjectTypeCompetition:        `core.ObjectTypeCompetition`,
	ObjectTypeCommand:            `core.ObjectTypeCommand`,
	ObjectTypeDNSRecord:          `core.ObjectTypeDNSRecord`,
	ObjectTypeEnvironment:        `core.ObjectTypeEnvironment`,
	ObjectTypeHost:               `core.ObjectTypeHost`,
	ObjectTypeIdentity:           `core.ObjectTypeIdentity`,
	ObjectTypeNetwork:            `core.ObjectTypeNetwork`,
	ObjectTypeRemoteFile:         `core.ObjectTypeRemoteFile`,
	ObjectTypeScript:             `core.ObjectTypeScript`,
	ObjectTypeTeam:               `core.ObjectTypeTeam`,
	ObjectTypeUser:               `core.ObjectTypeUser`,
	ObjectTypeAMI:                `core.ObjectTypeAMI`,
	ObjectTypeProvisionedHost:    `core.ObjectTypeProvisionedHost`,
	ObjectTypeProvisionedNetwork: `core.ObjectTypeProvisionedNetwork`,
	ObjectTypeProvisioningStep:   `core.ObjectTypeProvisioningStep`,
	ObjectTypeConnection:         `core.ObjectTypeConnection`,
	ObjectTypeIncluded:           `core.ObjectTypeIncluded`,
}

// Kind returns a string of the Go type for the given message.
func (x ObjectType) Kind() string {
	if str, ok := _ObjectTypeKinds[x]; ok {
		return str
	}

	return _ObjectTypeKinds[ObjectType(0)]
}

var _ObjectTypeSources = map[ObjectType]string{
	ObjectTypeUndefinedEnumValue: `github.com/gen0cide/laforge/core.ObjectTypeUndefinedEnumValue`,
	ObjectTypeBuild:              `github.com/gen0cide/laforge/core.ObjectTypeBuild`,
	ObjectTypeCompetition:        `github.com/gen0cide/laforge/core.ObjectTypeCompetition`,
	ObjectTypeCommand:            `github.com/gen0cide/laforge/core.ObjectTypeCommand`,
	ObjectTypeDNSRecord:          `github.com/gen0cide/laforge/core.ObjectTypeDNSRecord`,
	ObjectTypeEnvironment:        `github.com/gen0cide/laforge/core.ObjectTypeEnvironment`,
	ObjectTypeHost:               `github.com/gen0cide/laforge/core.ObjectTypeHost`,
	ObjectTypeIdentity:           `github.com/gen0cide/laforge/core.ObjectTypeIdentity`,
	ObjectTypeNetwork:            `github.com/gen0cide/laforge/core.ObjectTypeNetwork`,
	ObjectTypeRemoteFile:         `github.com/gen0cide/laforge/core.ObjectTypeRemoteFile`,
	ObjectTypeScript:             `github.com/gen0cide/laforge/core.ObjectTypeScript`,
	ObjectTypeTeam:               `github.com/gen0cide/laforge/core.ObjectTypeTeam`,
	ObjectTypeUser:               `github.com/gen0cide/laforge/core.ObjectTypeUser`,
	ObjectTypeAMI:                `github.com/gen0cide/laforge/core.ObjectTypeAMI`,
	ObjectTypeProvisionedHost:    `github.com/gen0cide/laforge/core.ObjectTypeProvisionedHost`,
	ObjectTypeProvisionedNetwork: `github.com/gen0cide/laforge/core.ObjectTypeProvisionedNetwork`,
	ObjectTypeProvisioningStep:   `github.com/gen0cide/laforge/core.ObjectTypeProvisioningStep`,
	ObjectTypeConnection:         `github.com/gen0cide/laforge/core.ObjectTypeConnection`,
	ObjectTypeIncluded:           `github.com/gen0cide/laforge/core.ObjectTypeIncluded`,
}

// Source returns an import path directly to the type.
func (x ObjectType) Source() string {
	if str, ok := _ObjectTypeSources[x]; ok {
		return str
	}

	return _ObjectTypeSources[ObjectType(0)]
}

var _ObjectTypePaths = map[ObjectType]string{
	ObjectTypeUndefinedEnumValue: `github.com.gen0cide.laforge.core.object_type_undefined_enum_value`,
	ObjectTypeBuild:              `github.com.gen0cide.laforge.core.object_type_build`,
	ObjectTypeCompetition:        `github.com.gen0cide.laforge.core.object_type_competition`,
	ObjectTypeCommand:            `github.com.gen0cide.laforge.core.object_type_command`,
	ObjectTypeDNSRecord:          `github.com.gen0cide.laforge.core.object_type_dns_record`,
	ObjectTypeEnvironment:        `github.com.gen0cide.laforge.core.object_type_environment`,
	ObjectTypeHost:               `github.com.gen0cide.laforge.core.object_type_host`,
	ObjectTypeIdentity:           `github.com.gen0cide.laforge.core.object_type_identity`,
	ObjectTypeNetwork:            `github.com.gen0cide.laforge.core.object_type_network`,
	ObjectTypeRemoteFile:         `github.com.gen0cide.laforge.core.object_type_remote_file`,
	ObjectTypeScript:             `github.com.gen0cide.laforge.core.object_type_script`,
	ObjectTypeTeam:               `github.com.gen0cide.laforge.core.object_type_team`,
	ObjectTypeUser:               `github.com.gen0cide.laforge.core.object_type_user`,
	ObjectTypeAMI:                `github.com.gen0cide.laforge.core.object_type_ami`,
	ObjectTypeProvisionedHost:    `github.com.gen0cide.laforge.core.object_type_provisioned_host`,
	ObjectTypeProvisionedNetwork: `github.com.gen0cide.laforge.core.object_type_provisioned_network`,
	ObjectTypeProvisioningStep:   `github.com.gen0cide.laforge.core.object_type_provisioning_step`,
	ObjectTypeConnection:         `github.com.gen0cide.laforge.core.object_type_connection`,
	ObjectTypeIncluded:           `github.com.gen0cide.laforge.core.object_type_included`,
}

// Path returns an import path directly to the type.
func (x ObjectType) Path() string {
	if str, ok := _ObjectTypePaths[x]; ok {
		return str
	}

	return _ObjectTypePaths[ObjectType(0)]
}

// PackageName returns the name of the parent package for this type.
func (x ObjectType) PackageName() string {
	return _ObjectTypePkgName
}

// ImportPath returns the full import path of the parent package
func (x ObjectType) ImportPath() string {
	return _ObjectTypePkgPath
}

// Namespace implements the emitter.Namespaced interface.
func (x ObjectType) Namespace() string {
	return _ObjectTypeNamespace
}

// ParseObjectType attempts to convert a string to a ObjectType
func ParseObjectType(name string) (ObjectType, error) {
	if len(name) < 1 {
		return ObjectType(0), ErrUndefinedObjectTypeEnumValue
	}

	first, _ := utf8.DecodeRuneInString(name)
	if first == utf8.RuneError {
		return ObjectType(0), ErrUndefinedObjectTypeEnumValue
	}

	switch {
	case unicode.IsLower(first):
		// test for snake_case
		if x, ok := _ObjectTypeValue[name]; ok {
			return x, nil
		}

		// test for command-case
		if x, ok := _ObjectTypeCommandValue[name]; ok {
			return x, nil
		}

		// test for camelCase
		if x, ok := _ObjectTypeCamelValue[name]; ok {
			return x, nil
		}
	case unicode.IsUpper(first):
		// test for PascalCase
		if x, ok := _ObjectTypePascalValue[name]; ok {
			return x, nil
		}

		// test for SCREAMING_CASE
		if x, ok := _ObjectTypeScreamingValue[name]; ok {
			return x, nil
		}
	default:
		return ObjectType(0), ErrUndefinedObjectTypeEnumValue
	}

	return ObjectType(0), ErrUndefinedObjectTypeEnumValue
}

// MarshalText implements the text marshaller method
func (x ObjectType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *ObjectType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseObjectType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Scan implements the Scanner interface.
func (x *ObjectType) Scan(value interface{}) error {
	var name string

	switch v := value.(type) {
	case string:
		name = v
	case []byte:
		name = string(v)
	case nil:
		*x = ObjectType(0)
		return nil
	}

	tmp, err := ParseObjectType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Value implements the driver Valuer interface.
func (x ObjectType) Value() (driver.Value, error) {
	return x.String(), nil
}

// LookupObjectType attempts to convert a int to it's equivelent ObjectType value.
func LookupObjectType(id int) (ObjectType, error) {
	if _, ok := _ObjectTypeMap[ObjectType(id)]; ok {
		return ObjectType(id), nil
	}
	return ObjectType(0), fmt.Errorf("%T(%v) is not a valid ObjectType, try [%s]", id, id, strings.Join(_ObjectTypeNames, ", "))
}

// Code implements the Coder interface.
func (x ObjectType) Code() int {
	return int(x)
}

// MarshalJSON implements the json.Marshaler interface.
func (x ObjectType) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (x *ObjectType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return fmt.Errorf("error unmarshaling JSON value: %v", err)
	}
	tmp, err := ParseObjectType(s)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// MarshalYAML implements the yaml.Marshaler interface.
func (x ObjectType) MarshalYAML() (interface{}, error) {
	return x.String(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (x *ObjectType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return fmt.Errorf("error unmarshaling YAML value: %v", err)
	}

	tmp, err := ParseObjectType(s)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
