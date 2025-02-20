/*
Endpoint Tests API

 Manage endpoint agent dynamic and scheduled tests using the Endpoint Tests API. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtests

import (
	"encoding/json"
	"fmt"
)

// TestProbeMode Probe mode used by network test, only valid when the protocol is set to TCP.
type TestProbeMode string

// List of TestProbeMode
const (
	TESTPROBEMODE_AUTO TestProbeMode = "auto"
	TESTPROBEMODE_SACK TestProbeMode = "sack"
	TESTPROBEMODE_SYN TestProbeMode = "syn"
)

// All allowed values of TestProbeMode enum
var AllowedTestProbeModeEnumValues = []TestProbeMode{
	"auto",
	"sack",
	"syn",
}

func (v *TestProbeMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestProbeMode(value)
	for _, existing := range AllowedTestProbeModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestProbeMode", value)
}

// NewTestProbeModeFromValue returns a pointer to a valid TestProbeMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestProbeModeFromValue(v string) (*TestProbeMode, error) {
	ev := TestProbeMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestProbeMode: valid values are %v", v, AllowedTestProbeModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestProbeMode) IsValid() bool {
	for _, existing := range AllowedTestProbeModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestProbeMode value
func (v TestProbeMode) Ptr() *TestProbeMode {
	return &v
}

type NullableTestProbeMode struct {
	value *TestProbeMode
	isSet bool
}

func (v NullableTestProbeMode) Get() *TestProbeMode {
	return v.value
}

func (v *NullableTestProbeMode) Set(val *TestProbeMode) {
	v.value = val
	v.isSet = true
}

func (v NullableTestProbeMode) IsSet() bool {
	return v.isSet
}

func (v *NullableTestProbeMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestProbeMode(val *TestProbeMode) *NullableTestProbeMode {
	return &NullableTestProbeMode{value: val, isSet: true}
}

func (v NullableTestProbeMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestProbeMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

