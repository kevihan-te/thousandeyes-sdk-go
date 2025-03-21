/*
Endpoint Agents API

Manage ThousandEyes Endpoint Agents using this API.   For more information about Endpoint Agents, see [Endpoint Agents](https://docs.thousandeyes.com/product-documentation/global-vantage-points/endpoint-agents).

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointagents

import (
	"encoding/json"
	"fmt"
)

// AgentLicenseType the model 'AgentLicenseType'
type AgentLicenseType string

// List of AgentLicenseType
const (
	AGENTLICENSETYPE_ESSENTIALS AgentLicenseType = "essentials"
	AGENTLICENSETYPE_ADVANTAGE AgentLicenseType = "advantage"
	AGENTLICENSETYPE_EMBEDDED AgentLicenseType = "embedded"
)

// All allowed values of AgentLicenseType enum
var AllowedAgentLicenseTypeEnumValues = []AgentLicenseType{
	"essentials",
	"advantage",
	"embedded",
}

func (v *AgentLicenseType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentLicenseType(value)
	for _, existing := range AllowedAgentLicenseTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentLicenseType", value)
}

// NewAgentLicenseTypeFromValue returns a pointer to a valid AgentLicenseType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentLicenseTypeFromValue(v string) (*AgentLicenseType, error) {
	ev := AgentLicenseType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentLicenseType: valid values are %v", v, AllowedAgentLicenseTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentLicenseType) IsValid() bool {
	for _, existing := range AllowedAgentLicenseTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentLicenseType value
func (v AgentLicenseType) Ptr() *AgentLicenseType {
	return &v
}

type NullableAgentLicenseType struct {
	value *AgentLicenseType
	isSet bool
}

func (v NullableAgentLicenseType) Get() *AgentLicenseType {
	return v.value
}

func (v *NullableAgentLicenseType) Set(val *AgentLicenseType) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentLicenseType) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentLicenseType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentLicenseType(val *AgentLicenseType) *NullableAgentLicenseType {
	return &NullableAgentLicenseType{value: val, isSet: true}
}

func (v NullableAgentLicenseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentLicenseType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

