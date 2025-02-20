/*
Agents API

 ## Overview Manage Cloud and Enterprise Agents available to your account in ThousandEyes.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agents

import (
	"encoding/json"
	"fmt"
)

// CloudEnterpriseAgentType Type of the agent.
type CloudEnterpriseAgentType string

// List of CloudEnterpriseAgentType
const (
	CLOUDENTERPRISEAGENTTYPE_CLOUD CloudEnterpriseAgentType = "cloud"
	CLOUDENTERPRISEAGENTTYPE_ENTERPRISE_CLUSTER CloudEnterpriseAgentType = "enterprise-cluster"
	CLOUDENTERPRISEAGENTTYPE_ENTERPRISE CloudEnterpriseAgentType = "enterprise"
)

// All allowed values of CloudEnterpriseAgentType enum
var AllowedCloudEnterpriseAgentTypeEnumValues = []CloudEnterpriseAgentType{
	"cloud",
	"enterprise-cluster",
	"enterprise",
}

func (v *CloudEnterpriseAgentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloudEnterpriseAgentType(value)
	for _, existing := range AllowedCloudEnterpriseAgentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CloudEnterpriseAgentType", value)
}

// NewCloudEnterpriseAgentTypeFromValue returns a pointer to a valid CloudEnterpriseAgentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCloudEnterpriseAgentTypeFromValue(v string) (*CloudEnterpriseAgentType, error) {
	ev := CloudEnterpriseAgentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CloudEnterpriseAgentType: valid values are %v", v, AllowedCloudEnterpriseAgentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CloudEnterpriseAgentType) IsValid() bool {
	for _, existing := range AllowedCloudEnterpriseAgentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CloudEnterpriseAgentType value
func (v CloudEnterpriseAgentType) Ptr() *CloudEnterpriseAgentType {
	return &v
}

type NullableCloudEnterpriseAgentType struct {
	value *CloudEnterpriseAgentType
	isSet bool
}

func (v NullableCloudEnterpriseAgentType) Get() *CloudEnterpriseAgentType {
	return v.value
}

func (v *NullableCloudEnterpriseAgentType) Set(val *CloudEnterpriseAgentType) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudEnterpriseAgentType) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudEnterpriseAgentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudEnterpriseAgentType(val *CloudEnterpriseAgentType) *NullableCloudEnterpriseAgentType {
	return &NullableCloudEnterpriseAgentType{value: val, isSet: true}
}

func (v NullableCloudEnterpriseAgentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudEnterpriseAgentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

