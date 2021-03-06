// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type ConfigMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type VarsMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type FindingDifficulty string

const (
	FindingDifficultyZeroDifficulty     FindingDifficulty = "ZeroDifficulty"
	FindingDifficultyNoviceDifficulty   FindingDifficulty = "NoviceDifficulty"
	FindingDifficultyAdvancedDifficulty FindingDifficulty = "AdvancedDifficulty"
	FindingDifficultyExpertDifficulty   FindingDifficulty = "ExpertDifficulty"
	FindingDifficultyNullDifficulty     FindingDifficulty = "NullDifficulty"
)

var AllFindingDifficulty = []FindingDifficulty{
	FindingDifficultyZeroDifficulty,
	FindingDifficultyNoviceDifficulty,
	FindingDifficultyAdvancedDifficulty,
	FindingDifficultyExpertDifficulty,
	FindingDifficultyNullDifficulty,
}

func (e FindingDifficulty) IsValid() bool {
	switch e {
	case FindingDifficultyZeroDifficulty, FindingDifficultyNoviceDifficulty, FindingDifficultyAdvancedDifficulty, FindingDifficultyExpertDifficulty, FindingDifficultyNullDifficulty:
		return true
	}
	return false
}

func (e FindingDifficulty) String() string {
	return string(e)
}

func (e *FindingDifficulty) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FindingDifficulty(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FindingDifficulty", str)
	}
	return nil
}

func (e FindingDifficulty) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FindingSeverity string

const (
	FindingSeverityZeroSeverity     FindingSeverity = "ZeroSeverity"
	FindingSeverityLowSeverity      FindingSeverity = "LowSeverity"
	FindingSeverityMediumSeverity   FindingSeverity = "MediumSeverity"
	FindingSeverityHighSeverity     FindingSeverity = "HighSeverity"
	FindingSeverityCriticalSeverity FindingSeverity = "CriticalSeverity"
	FindingSeverityNullSeverity     FindingSeverity = "NullSeverity"
)

var AllFindingSeverity = []FindingSeverity{
	FindingSeverityZeroSeverity,
	FindingSeverityLowSeverity,
	FindingSeverityMediumSeverity,
	FindingSeverityHighSeverity,
	FindingSeverityCriticalSeverity,
	FindingSeverityNullSeverity,
}

func (e FindingSeverity) IsValid() bool {
	switch e {
	case FindingSeverityZeroSeverity, FindingSeverityLowSeverity, FindingSeverityMediumSeverity, FindingSeverityHighSeverity, FindingSeverityCriticalSeverity, FindingSeverityNullSeverity:
		return true
	}
	return false
}

func (e FindingSeverity) String() string {
	return string(e)
}

func (e *FindingSeverity) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FindingSeverity(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FindingSeverity", str)
	}
	return nil
}

func (e FindingSeverity) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProvisionStatus string

const (
	ProvisionStatusProvStatusUndefined  ProvisionStatus = "ProvStatusUndefined"
	ProvisionStatusProvStatusPlanning   ProvisionStatus = "ProvStatusPlanning"
	ProvisionStatusProvStatusAwaiting   ProvisionStatus = "ProvStatusAwaiting"
	ProvisionStatusProvStatusInProgress ProvisionStatus = "ProvStatusInProgress"
	ProvisionStatusProvStatusFailed     ProvisionStatus = "ProvStatusFailed"
	ProvisionStatusProvStatusComplete   ProvisionStatus = "ProvStatusComplete"
	ProvisionStatusProvStatusTainted    ProvisionStatus = "ProvStatusTainted"
)

var AllProvisionStatus = []ProvisionStatus{
	ProvisionStatusProvStatusUndefined,
	ProvisionStatusProvStatusPlanning,
	ProvisionStatusProvStatusAwaiting,
	ProvisionStatusProvStatusInProgress,
	ProvisionStatusProvStatusFailed,
	ProvisionStatusProvStatusComplete,
	ProvisionStatusProvStatusTainted,
}

func (e ProvisionStatus) IsValid() bool {
	switch e {
	case ProvisionStatusProvStatusUndefined, ProvisionStatusProvStatusPlanning, ProvisionStatusProvStatusAwaiting, ProvisionStatusProvStatusInProgress, ProvisionStatusProvStatusFailed, ProvisionStatusProvStatusComplete, ProvisionStatusProvStatusTainted:
		return true
	}
	return false
}

func (e ProvisionStatus) String() string {
	return string(e)
}

func (e *ProvisionStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProvisionStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProvisionStatus", str)
	}
	return nil
}

func (e ProvisionStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
