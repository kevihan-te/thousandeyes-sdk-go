/*
Usage API

 These usage endpoints define the following operations:  * **Usage**: Retrieve usage data for the specified time period (default is one month).          * Users must have the `View organization usage` permission to access this endpoint.     * This operation offers visibility across all account groups within the organization.     * Users with `View organization usage` permission in multiple organizations should query the operation with the `aid` query string parameter (see optional parameters) for each organization.  * **Quotas**: Obtain organization and account usage quotas. Additionally, users with the appropriate permissions can create, update, or delete these quotas.          * Users must have the necessary permissions to perform quota-related actions.  Refer to the Usage API operations for detailed usage instructions and optional parameters. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package usage

import (
	"encoding/json"
	"fmt"
)

// ExpandUsageOptions the model 'ExpandUsageOptions'
type ExpandUsageOptions string

// List of ExpandUsageOptions
const (
	EXPANDUSAGEOPTIONS_TEST ExpandUsageOptions = "test"
	EXPANDUSAGEOPTIONS_ENTERPRISE_AGENT ExpandUsageOptions = "enterprise-agent"
	EXPANDUSAGEOPTIONS_ENTERPRISE_AGENT_UNIT ExpandUsageOptions = "enterprise-agent-unit"
	EXPANDUSAGEOPTIONS_ENDPOINT_AGENT ExpandUsageOptions = "endpoint-agent"
	EXPANDUSAGEOPTIONS_ENDPOINT_AGENT_ESSENTIAL ExpandUsageOptions = "endpoint-agent-essential"
	EXPANDUSAGEOPTIONS_ENDPOINT_AGENT_EMBEDDED ExpandUsageOptions = "endpoint-agent-embedded"
)

// All allowed values of ExpandUsageOptions enum
var AllowedExpandUsageOptionsEnumValues = []ExpandUsageOptions{
	"test",
	"enterprise-agent",
	"enterprise-agent-unit",
	"endpoint-agent",
	"endpoint-agent-essential",
	"endpoint-agent-embedded",
}

func (v *ExpandUsageOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpandUsageOptions(value)
	for _, existing := range AllowedExpandUsageOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpandUsageOptions", value)
}

// NewExpandUsageOptionsFromValue returns a pointer to a valid ExpandUsageOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpandUsageOptionsFromValue(v string) (*ExpandUsageOptions, error) {
	ev := ExpandUsageOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpandUsageOptions: valid values are %v", v, AllowedExpandUsageOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpandUsageOptions) IsValid() bool {
	for _, existing := range AllowedExpandUsageOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpandUsageOptions value
func (v ExpandUsageOptions) Ptr() *ExpandUsageOptions {
	return &v
}

type NullableExpandUsageOptions struct {
	value *ExpandUsageOptions
	isSet bool
}

func (v NullableExpandUsageOptions) Get() *ExpandUsageOptions {
	return v.value
}

func (v *NullableExpandUsageOptions) Set(val *ExpandUsageOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandUsageOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandUsageOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandUsageOptions(val *ExpandUsageOptions) *NullableExpandUsageOptions {
	return &NullableExpandUsageOptions{value: val, isSet: true}
}

func (v NullableExpandUsageOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandUsageOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

