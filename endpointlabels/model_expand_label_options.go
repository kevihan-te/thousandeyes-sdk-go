/*
Endpoint Agent Labels API

Manage labels applied to endpoint agents using this API. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointlabels

import (
	"encoding/json"
	"fmt"
)

// ExpandLabelOptions the model 'ExpandLabelOptions'
type ExpandLabelOptions string

// List of ExpandLabelOptions
const (
	EXPANDLABELOPTIONS_FILTERS ExpandLabelOptions = "filters"
)

// All allowed values of ExpandLabelOptions enum
var AllowedExpandLabelOptionsEnumValues = []ExpandLabelOptions{
	"filters",
}

func (v *ExpandLabelOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpandLabelOptions(value)
	for _, existing := range AllowedExpandLabelOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpandLabelOptions", value)
}

// NewExpandLabelOptionsFromValue returns a pointer to a valid ExpandLabelOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpandLabelOptionsFromValue(v string) (*ExpandLabelOptions, error) {
	ev := ExpandLabelOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpandLabelOptions: valid values are %v", v, AllowedExpandLabelOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpandLabelOptions) IsValid() bool {
	for _, existing := range AllowedExpandLabelOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpandLabelOptions value
func (v ExpandLabelOptions) Ptr() *ExpandLabelOptions {
	return &v
}

type NullableExpandLabelOptions struct {
	value *ExpandLabelOptions
	isSet bool
}

func (v NullableExpandLabelOptions) Get() *ExpandLabelOptions {
	return v.value
}

func (v *NullableExpandLabelOptions) Set(val *ExpandLabelOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandLabelOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandLabelOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandLabelOptions(val *ExpandLabelOptions) *NullableExpandLabelOptions {
	return &NullableExpandLabelOptions{value: val, isSet: true}
}

func (v NullableExpandLabelOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandLabelOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

