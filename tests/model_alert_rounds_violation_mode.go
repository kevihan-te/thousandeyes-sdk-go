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

// AlertRoundsViolationMode `exact` requires the same agents to meet the threshold in consecutive rounds. `auto` is only enabled for CEA and Endpoint Scheduled test rules. The default is `any`.
type AlertRoundsViolationMode string

// List of AlertRoundsViolationMode
const (
	ALERTROUNDSVIOLATIONMODE_EXACT AlertRoundsViolationMode = "exact"
	ALERTROUNDSVIOLATIONMODE_ANY AlertRoundsViolationMode = "any"
	ALERTROUNDSVIOLATIONMODE_AUTO AlertRoundsViolationMode = "auto"
)

// All allowed values of AlertRoundsViolationMode enum
var AllowedAlertRoundsViolationModeEnumValues = []AlertRoundsViolationMode{
	"exact",
	"any",
	"auto",
}

func (v *AlertRoundsViolationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertRoundsViolationMode(value)
	for _, existing := range AllowedAlertRoundsViolationModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertRoundsViolationMode", value)
}

// NewAlertRoundsViolationModeFromValue returns a pointer to a valid AlertRoundsViolationMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertRoundsViolationModeFromValue(v string) (*AlertRoundsViolationMode, error) {
	ev := AlertRoundsViolationMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertRoundsViolationMode: valid values are %v", v, AllowedAlertRoundsViolationModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertRoundsViolationMode) IsValid() bool {
	for _, existing := range AllowedAlertRoundsViolationModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertRoundsViolationMode value
func (v AlertRoundsViolationMode) Ptr() *AlertRoundsViolationMode {
	return &v
}

type NullableAlertRoundsViolationMode struct {
	value *AlertRoundsViolationMode
	isSet bool
}

func (v NullableAlertRoundsViolationMode) Get() *AlertRoundsViolationMode {
	return v.value
}

func (v *NullableAlertRoundsViolationMode) Set(val *AlertRoundsViolationMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRoundsViolationMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRoundsViolationMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRoundsViolationMode(val *AlertRoundsViolationMode) *NullableAlertRoundsViolationMode {
	return &NullableAlertRoundsViolationMode{value: val, isSet: true}
}

func (v NullableAlertRoundsViolationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRoundsViolationMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

