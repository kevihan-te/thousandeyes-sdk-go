/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
	"fmt"
)

// EnterpriseAgentState State of the agent.
type EnterpriseAgentState string

// List of EnterpriseAgentState
const (
	ENTERPRISEAGENTSTATE_ONLINE EnterpriseAgentState = "online"
	ENTERPRISEAGENTSTATE_OFFLINE EnterpriseAgentState = "offline"
	ENTERPRISEAGENTSTATE_DISABLED EnterpriseAgentState = "disabled"
)

// All allowed values of EnterpriseAgentState enum
var AllowedEnterpriseAgentStateEnumValues = []EnterpriseAgentState{
	"online",
	"offline",
	"disabled",
}

func (v *EnterpriseAgentState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnterpriseAgentState(value)
	for _, existing := range AllowedEnterpriseAgentStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnterpriseAgentState", value)
}

// NewEnterpriseAgentStateFromValue returns a pointer to a valid EnterpriseAgentState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnterpriseAgentStateFromValue(v string) (*EnterpriseAgentState, error) {
	ev := EnterpriseAgentState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnterpriseAgentState: valid values are %v", v, AllowedEnterpriseAgentStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnterpriseAgentState) IsValid() bool {
	for _, existing := range AllowedEnterpriseAgentStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnterpriseAgentState value
func (v EnterpriseAgentState) Ptr() *EnterpriseAgentState {
	return &v
}

type NullableEnterpriseAgentState struct {
	value *EnterpriseAgentState
	isSet bool
}

func (v NullableEnterpriseAgentState) Get() *EnterpriseAgentState {
	return v.value
}

func (v *NullableEnterpriseAgentState) Set(val *EnterpriseAgentState) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAgentState) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAgentState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAgentState(val *EnterpriseAgentState) *NullableEnterpriseAgentState {
	return &NullableEnterpriseAgentState{value: val, isSet: true}
}

func (v NullableEnterpriseAgentState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAgentState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

