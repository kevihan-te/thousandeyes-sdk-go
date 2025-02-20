/*
Endpoint Agents API

Manage ThousandEyes Endpoint Agents using this API.   For more information about Endpoint Agents, see [Endpoint Agents](https://docs.thousandeyes.com/product-documentation/global-vantage-points/endpoint-agents).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointagents

import (
	"encoding/json"
	"fmt"
)

// ThresholdFilterOperator the model 'ThresholdFilterOperator'
type ThresholdFilterOperator string

// List of ThresholdFilterOperator
const (
	THRESHOLDFILTEROPERATOR_GTE ThresholdFilterOperator = "gte"
	THRESHOLDFILTEROPERATOR_LTE ThresholdFilterOperator = "lte"
)

// All allowed values of ThresholdFilterOperator enum
var AllowedThresholdFilterOperatorEnumValues = []ThresholdFilterOperator{
	"gte",
	"lte",
}

func (v *ThresholdFilterOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ThresholdFilterOperator(value)
	for _, existing := range AllowedThresholdFilterOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ThresholdFilterOperator", value)
}

// NewThresholdFilterOperatorFromValue returns a pointer to a valid ThresholdFilterOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewThresholdFilterOperatorFromValue(v string) (*ThresholdFilterOperator, error) {
	ev := ThresholdFilterOperator(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ThresholdFilterOperator: valid values are %v", v, AllowedThresholdFilterOperatorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ThresholdFilterOperator) IsValid() bool {
	for _, existing := range AllowedThresholdFilterOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ThresholdFilterOperator value
func (v ThresholdFilterOperator) Ptr() *ThresholdFilterOperator {
	return &v
}

type NullableThresholdFilterOperator struct {
	value *ThresholdFilterOperator
	isSet bool
}

func (v NullableThresholdFilterOperator) Get() *ThresholdFilterOperator {
	return v.value
}

func (v *NullableThresholdFilterOperator) Set(val *ThresholdFilterOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableThresholdFilterOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableThresholdFilterOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThresholdFilterOperator(val *ThresholdFilterOperator) *NullableThresholdFilterOperator {
	return &NullableThresholdFilterOperator{value: val, isSet: true}
}

func (v NullableThresholdFilterOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableThresholdFilterOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

