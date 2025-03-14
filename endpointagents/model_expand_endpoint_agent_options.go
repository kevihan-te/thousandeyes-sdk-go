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

// ExpandEndpointAgentOptions the model 'ExpandEndpointAgentOptions'
type ExpandEndpointAgentOptions string

// List of ExpandEndpointAgentOptions
const (
	EXPANDENDPOINTAGENTOPTIONS_CLIENTS ExpandEndpointAgentOptions = "clients"
	EXPANDENDPOINTAGENTOPTIONS_VPN_PROFILES ExpandEndpointAgentOptions = "vpnProfiles"
	EXPANDENDPOINTAGENTOPTIONS_NETWORK_INTERFACE_PROFILES ExpandEndpointAgentOptions = "networkInterfaceProfiles"
)

// All allowed values of ExpandEndpointAgentOptions enum
var AllowedExpandEndpointAgentOptionsEnumValues = []ExpandEndpointAgentOptions{
	"clients",
	"vpnProfiles",
	"networkInterfaceProfiles",
}

func (v *ExpandEndpointAgentOptions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpandEndpointAgentOptions(value)
	for _, existing := range AllowedExpandEndpointAgentOptionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpandEndpointAgentOptions", value)
}

// NewExpandEndpointAgentOptionsFromValue returns a pointer to a valid ExpandEndpointAgentOptions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpandEndpointAgentOptionsFromValue(v string) (*ExpandEndpointAgentOptions, error) {
	ev := ExpandEndpointAgentOptions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpandEndpointAgentOptions: valid values are %v", v, AllowedExpandEndpointAgentOptionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpandEndpointAgentOptions) IsValid() bool {
	for _, existing := range AllowedExpandEndpointAgentOptionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpandEndpointAgentOptions value
func (v ExpandEndpointAgentOptions) Ptr() *ExpandEndpointAgentOptions {
	return &v
}

type NullableExpandEndpointAgentOptions struct {
	value *ExpandEndpointAgentOptions
	isSet bool
}

func (v NullableExpandEndpointAgentOptions) Get() *ExpandEndpointAgentOptions {
	return v.value
}

func (v *NullableExpandEndpointAgentOptions) Set(val *ExpandEndpointAgentOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandEndpointAgentOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandEndpointAgentOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandEndpointAgentOptions(val *ExpandEndpointAgentOptions) *NullableExpandEndpointAgentOptions {
	return &NullableExpandEndpointAgentOptions{value: val, isSet: true}
}

func (v NullableExpandEndpointAgentOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandEndpointAgentOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

