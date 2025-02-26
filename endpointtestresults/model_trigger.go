/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
	"fmt"
)

// Trigger the model 'Trigger'
type Trigger string

// List of Trigger
const (
	TRIGGER_AUTO Trigger = "auto"
	TRIGGER_USER Trigger = "user"
)

// All allowed values of Trigger enum
var AllowedTriggerEnumValues = []Trigger{
	"auto",
	"user",
}

func (v *Trigger) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Trigger(value)
	for _, existing := range AllowedTriggerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Trigger", value)
}

// NewTriggerFromValue returns a pointer to a valid Trigger
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerFromValue(v string) (*Trigger, error) {
	ev := Trigger(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Trigger: valid values are %v", v, AllowedTriggerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Trigger) IsValid() bool {
	for _, existing := range AllowedTriggerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Trigger value
func (v Trigger) Ptr() *Trigger {
	return &v
}

type NullableTrigger struct {
	value *Trigger
	isSet bool
}

func (v NullableTrigger) Get() *Trigger {
	return v.value
}

func (v *NullableTrigger) Set(val *Trigger) {
	v.value = val
	v.isSet = true
}

func (v NullableTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrigger(val *Trigger) *NullableTrigger {
	return &NullableTrigger{value: val, isSet: true}
}

func (v NullableTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

