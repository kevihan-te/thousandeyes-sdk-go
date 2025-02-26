/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
	"fmt"
)

// TestInterval Interval between test runs in seconds.
type TestInterval int32

// List of TestInterval
const (
	TESTINTERVAL__60 TestInterval = 60
	TESTINTERVAL__120 TestInterval = 120
	TESTINTERVAL__300 TestInterval = 300
	TESTINTERVAL__600 TestInterval = 600
	TESTINTERVAL__900 TestInterval = 900
	TESTINTERVAL__1800 TestInterval = 1800
	TESTINTERVAL__3600 TestInterval = 3600
)

// All allowed values of TestInterval enum
var AllowedTestIntervalEnumValues = []TestInterval{
	60,
	120,
	300,
	600,
	900,
	1800,
	3600,
}

func (v *TestInterval) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestInterval(value)
	for _, existing := range AllowedTestIntervalEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestInterval", value)
}

// NewTestIntervalFromValue returns a pointer to a valid TestInterval
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestIntervalFromValue(v int32) (*TestInterval, error) {
	ev := TestInterval(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestInterval: valid values are %v", v, AllowedTestIntervalEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestInterval) IsValid() bool {
	for _, existing := range AllowedTestIntervalEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestInterval value
func (v TestInterval) Ptr() *TestInterval {
	return &v
}

type NullableTestInterval struct {
	value *TestInterval
	isSet bool
}

func (v NullableTestInterval) Get() *TestInterval {
	return v.value
}

func (v *NullableTestInterval) Set(val *TestInterval) {
	v.value = val
	v.isSet = true
}

func (v NullableTestInterval) IsSet() bool {
	return v.isSet
}

func (v *NullableTestInterval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestInterval(val *TestInterval) *NullableTestInterval {
	return &NullableTestInterval{value: val, isSet: true}
}

func (v NullableTestInterval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestInterval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

