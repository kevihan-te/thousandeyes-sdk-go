/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the TestAgentRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestAgentRequest{}

// TestAgentRequest struct for TestAgentRequest
type TestAgentRequest struct {
	// The agent ID. Get `agentId` from `/agents` endpoint.
	AgentId string `json:"agentId"`
	// The IP address from the `ipAddresses` field in agent details, used for interface selection. Get `ipAddresses` from the `/agents` endpoint.
	SourceIpAddress *string `json:"sourceIpAddress,omitempty"`
}

type _TestAgentRequest TestAgentRequest

// NewTestAgentRequest instantiates a new TestAgentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestAgentRequest(agentId string) *TestAgentRequest {
	this := TestAgentRequest{}
	this.AgentId = agentId
	return &this
}

// NewTestAgentRequestWithDefaults instantiates a new TestAgentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestAgentRequestWithDefaults() *TestAgentRequest {
	this := TestAgentRequest{}
	return &this
}

// GetAgentId returns the AgentId field value
func (o *TestAgentRequest) GetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *TestAgentRequest) GetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *TestAgentRequest) SetAgentId(v string) {
	o.AgentId = v
}

// GetSourceIpAddress returns the SourceIpAddress field value if set, zero value otherwise.
func (o *TestAgentRequest) GetSourceIpAddress() string {
	if o == nil || utils.IsNil(o.SourceIpAddress) {
		var ret string
		return ret
	}
	return *o.SourceIpAddress
}

// GetSourceIpAddressOk returns a tuple with the SourceIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestAgentRequest) GetSourceIpAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceIpAddress) {
		return nil, false
	}
	return o.SourceIpAddress, true
}

// HasSourceIpAddress returns a boolean if a field has been set.
func (o *TestAgentRequest) HasSourceIpAddress() bool {
	if o != nil && !utils.IsNil(o.SourceIpAddress) {
		return true
	}

	return false
}

// SetSourceIpAddress gets a reference to the given string and assigns it to the SourceIpAddress field.
func (o *TestAgentRequest) SetSourceIpAddress(v string) {
	o.SourceIpAddress = &v
}

func (o TestAgentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestAgentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agentId"] = o.AgentId
	if !utils.IsNil(o.SourceIpAddress) {
		toSerialize["sourceIpAddress"] = o.SourceIpAddress
	}
	return toSerialize, nil
}

func (o *TestAgentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agentId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTestAgentRequest := _TestAgentRequest{}

    err = json.Unmarshal(data, &varTestAgentRequest)

	if err != nil {
		return err
	}

	*o = TestAgentRequest(varTestAgentRequest)

	return err
}

type NullableTestAgentRequest struct {
	value *TestAgentRequest
	isSet bool
}

func (v NullableTestAgentRequest) Get() *TestAgentRequest {
	return v.value
}

func (v *NullableTestAgentRequest) Set(val *TestAgentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTestAgentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTestAgentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestAgentRequest(val *TestAgentRequest) *NullableTestAgentRequest {
	return &NullableTestAgentRequest{value: val, isSet: true}
}

func (v NullableTestAgentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestAgentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


