/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
	"fmt"
)

// SensitivityLevel Used when `roundsViolatingMode` is set to `auto`. The default is `medium`. Higher sensitivity increases the likelihood of triggering alerts.
type SensitivityLevel string

// List of SensitivityLevel
const (
	SENSITIVITYLEVEL_HIGH SensitivityLevel = "high"
	SENSITIVITYLEVEL_MEDIUM SensitivityLevel = "medium"
	SENSITIVITYLEVEL_LOW SensitivityLevel = "low"
)

// All allowed values of SensitivityLevel enum
var AllowedSensitivityLevelEnumValues = []SensitivityLevel{
	"high",
	"medium",
	"low",
}

func (v *SensitivityLevel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SensitivityLevel(value)
	for _, existing := range AllowedSensitivityLevelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SensitivityLevel", value)
}

// NewSensitivityLevelFromValue returns a pointer to a valid SensitivityLevel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSensitivityLevelFromValue(v string) (*SensitivityLevel, error) {
	ev := SensitivityLevel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SensitivityLevel: valid values are %v", v, AllowedSensitivityLevelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SensitivityLevel) IsValid() bool {
	for _, existing := range AllowedSensitivityLevelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SensitivityLevel value
func (v SensitivityLevel) Ptr() *SensitivityLevel {
	return &v
}

type NullableSensitivityLevel struct {
	value *SensitivityLevel
	isSet bool
}

func (v NullableSensitivityLevel) Get() *SensitivityLevel {
	return v.value
}

func (v *NullableSensitivityLevel) Set(val *SensitivityLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableSensitivityLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableSensitivityLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensitivityLevel(val *SensitivityLevel) *NullableSensitivityLevel {
	return &NullableSensitivityLevel{value: val, isSet: true}
}

func (v NullableSensitivityLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensitivityLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

