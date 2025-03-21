/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AgentInterfaces type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentInterfaces{}

// AgentInterfaces struct for AgentInterfaces
type AgentInterfaces struct {
	// IP address of the agent interface.
	IpAddress *string `json:"ipAddress,omitempty"`
	// The agent ID of the enterprise agent for the test.
	AgentId *string `json:"agentId,omitempty"`
}

// NewAgentInterfaces instantiates a new AgentInterfaces object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentInterfaces() *AgentInterfaces {
	this := AgentInterfaces{}
	return &this
}

// NewAgentInterfacesWithDefaults instantiates a new AgentInterfaces object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentInterfacesWithDefaults() *AgentInterfaces {
	this := AgentInterfaces{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *AgentInterfaces) GetIpAddress() string {
	if o == nil || utils.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentInterfaces) GetIpAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *AgentInterfaces) HasIpAddress() bool {
	if o != nil && !utils.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *AgentInterfaces) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AgentInterfaces) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentInterfaces) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AgentInterfaces) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *AgentInterfaces) SetAgentId(v string) {
	o.AgentId = &v
}

func (o AgentInterfaces) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentInterfaces) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	return toSerialize, nil
}

type NullableAgentInterfaces struct {
	value *AgentInterfaces
	isSet bool
}

func (v NullableAgentInterfaces) Get() *AgentInterfaces {
	return v.value
}

func (v *NullableAgentInterfaces) Set(val *AgentInterfaces) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentInterfaces) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentInterfaces) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentInterfaces(val *AgentInterfaces) *NullableAgentInterfaces {
	return &NullableAgentInterfaces{value: val, isSet: true}
}

func (v NullableAgentInterfaces) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentInterfaces) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


