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

// TestTableFilterKey the model 'TestTableFilterKey'
type TestTableFilterKey string

// List of TestTableFilterKey
const (
	TESTTABLEFILTERKEY_ANYTHING TestTableFilterKey = "Anything"
	TESTTABLEFILTERKEY_TEST_NAME TestTableFilterKey = "Test Name"
	TESTTABLEFILTERKEY_TARGET TestTableFilterKey = "Target"
	TESTTABLEFILTERKEY_TEST_ID TestTableFilterKey = "Test ID"
	TESTTABLEFILTERKEY_TEST_TYPE TestTableFilterKey = "Test type"
	TESTTABLEFILTERKEY_LABEL_ID TestTableFilterKey = "Label ID"
)

// All allowed values of TestTableFilterKey enum
var AllowedTestTableFilterKeyEnumValues = []TestTableFilterKey{
	"Anything",
	"Test Name",
	"Target",
	"Test ID",
	"Test type",
	"Label ID",
}

func (v *TestTableFilterKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestTableFilterKey(value)
	for _, existing := range AllowedTestTableFilterKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestTableFilterKey", value)
}

// NewTestTableFilterKeyFromValue returns a pointer to a valid TestTableFilterKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestTableFilterKeyFromValue(v string) (*TestTableFilterKey, error) {
	ev := TestTableFilterKey(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestTableFilterKey: valid values are %v", v, AllowedTestTableFilterKeyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestTableFilterKey) IsValid() bool {
	for _, existing := range AllowedTestTableFilterKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestTableFilterKey value
func (v TestTableFilterKey) Ptr() *TestTableFilterKey {
	return &v
}

type NullableTestTableFilterKey struct {
	value *TestTableFilterKey
	isSet bool
}

func (v NullableTestTableFilterKey) Get() *TestTableFilterKey {
	return v.value
}

func (v *NullableTestTableFilterKey) Set(val *TestTableFilterKey) {
	v.value = val
	v.isSet = true
}

func (v NullableTestTableFilterKey) IsSet() bool {
	return v.isSet
}

func (v *NullableTestTableFilterKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestTableFilterKey(val *TestTableFilterKey) *NullableTestTableFilterKey {
	return &NullableTestTableFilterKey{value: val, isSet: true}
}

func (v NullableTestTableFilterKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestTableFilterKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

