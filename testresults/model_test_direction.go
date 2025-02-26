/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
	"fmt"
)

// TestDirection Direction of the test, which affects how results are shown.
type TestDirection string

// List of TestDirection
const (
	TESTDIRECTION_TO_TARGET TestDirection = "to-target"
	TESTDIRECTION_FROM_TARGET TestDirection = "from-target"
	TESTDIRECTION_BIDIRECTIONAL TestDirection = "bidirectional"
)

// All allowed values of TestDirection enum
var AllowedTestDirectionEnumValues = []TestDirection{
	"to-target",
	"from-target",
	"bidirectional",
}

func (v *TestDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestDirection(value)
	for _, existing := range AllowedTestDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestDirection", value)
}

// NewTestDirectionFromValue returns a pointer to a valid TestDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestDirectionFromValue(v string) (*TestDirection, error) {
	ev := TestDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestDirection: valid values are %v", v, AllowedTestDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestDirection) IsValid() bool {
	for _, existing := range AllowedTestDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestDirection value
func (v TestDirection) Ptr() *TestDirection {
	return &v
}

type NullableTestDirection struct {
	value *TestDirection
	isSet bool
}

func (v NullableTestDirection) Get() *TestDirection {
	return v.value
}

func (v *NullableTestDirection) Set(val *TestDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableTestDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableTestDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestDirection(val *TestDirection) *NullableTestDirection {
	return &NullableTestDirection{value: val, isSet: true}
}

func (v NullableTestDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

