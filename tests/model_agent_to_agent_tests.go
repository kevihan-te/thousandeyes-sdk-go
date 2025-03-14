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

// checks if the AgentToAgentTests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentToAgentTests{}

// AgentToAgentTests struct for AgentToAgentTests
type AgentToAgentTests struct {
	Tests []UnexpandedAgentToAgentTest `json:"tests,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewAgentToAgentTests instantiates a new AgentToAgentTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToAgentTests() *AgentToAgentTests {
	this := AgentToAgentTests{}
	return &this
}

// NewAgentToAgentTestsWithDefaults instantiates a new AgentToAgentTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToAgentTestsWithDefaults() *AgentToAgentTests {
	this := AgentToAgentTests{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *AgentToAgentTests) GetTests() []UnexpandedAgentToAgentTest {
	if o == nil || utils.IsNil(o.Tests) {
		var ret []UnexpandedAgentToAgentTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentTests) GetTestsOk() ([]UnexpandedAgentToAgentTest, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *AgentToAgentTests) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []UnexpandedAgentToAgentTest and assigns it to the Tests field.
func (o *AgentToAgentTests) SetTests(v []UnexpandedAgentToAgentTest) {
	o.Tests = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgentToAgentTests) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentTests) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgentToAgentTests) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *AgentToAgentTests) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o AgentToAgentTests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToAgentTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableAgentToAgentTests struct {
	value *AgentToAgentTests
	isSet bool
}

func (v NullableAgentToAgentTests) Get() *AgentToAgentTests {
	return v.value
}

func (v *NullableAgentToAgentTests) Set(val *AgentToAgentTests) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToAgentTests) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToAgentTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToAgentTests(val *AgentToAgentTests) *NullableAgentToAgentTests {
	return &NullableAgentToAgentTests{value: val, isSet: true}
}

func (v NullableAgentToAgentTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToAgentTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


