/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
	"fmt"
)

// LegacyDurationUnit Timespan unit.
type LegacyDurationUnit string

// List of LegacyDurationUnit
const (
	LEGACYDURATIONUNIT_MINUTES LegacyDurationUnit = "Minutes"
	LEGACYDURATIONUNIT_HOURS LegacyDurationUnit = "Hours"
	LEGACYDURATIONUNIT_DAYS LegacyDurationUnit = "Days"
)

// All allowed values of LegacyDurationUnit enum
var AllowedLegacyDurationUnitEnumValues = []LegacyDurationUnit{
	"Minutes",
	"Hours",
	"Days",
}

func (v *LegacyDurationUnit) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LegacyDurationUnit(value)
	for _, existing := range AllowedLegacyDurationUnitEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LegacyDurationUnit", value)
}

// NewLegacyDurationUnitFromValue returns a pointer to a valid LegacyDurationUnit
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLegacyDurationUnitFromValue(v string) (*LegacyDurationUnit, error) {
	ev := LegacyDurationUnit(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LegacyDurationUnit: valid values are %v", v, AllowedLegacyDurationUnitEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LegacyDurationUnit) IsValid() bool {
	for _, existing := range AllowedLegacyDurationUnitEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LegacyDurationUnit value
func (v LegacyDurationUnit) Ptr() *LegacyDurationUnit {
	return &v
}

type NullableLegacyDurationUnit struct {
	value *LegacyDurationUnit
	isSet bool
}

func (v NullableLegacyDurationUnit) Get() *LegacyDurationUnit {
	return v.value
}

func (v *NullableLegacyDurationUnit) Set(val *LegacyDurationUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyDurationUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyDurationUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyDurationUnit(val *LegacyDurationUnit) *NullableLegacyDurationUnit {
	return &NullableLegacyDurationUnit{value: val, isSet: true}
}

func (v NullableLegacyDurationUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyDurationUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

