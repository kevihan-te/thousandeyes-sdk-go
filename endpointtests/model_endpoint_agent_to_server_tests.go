/*
Endpoint Tests API

 Manage endpoint agent dynamic and scheduled tests using the Endpoint Tests API. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the EndpointAgentToServerTests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointAgentToServerTests{}

// EndpointAgentToServerTests struct for EndpointAgentToServerTests
type EndpointAgentToServerTests struct {
	Tests []EndpointAgentToServerTest `json:"tests,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewEndpointAgentToServerTests instantiates a new EndpointAgentToServerTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAgentToServerTests() *EndpointAgentToServerTests {
	this := EndpointAgentToServerTests{}
	return &this
}

// NewEndpointAgentToServerTestsWithDefaults instantiates a new EndpointAgentToServerTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAgentToServerTestsWithDefaults() *EndpointAgentToServerTests {
	this := EndpointAgentToServerTests{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *EndpointAgentToServerTests) GetTests() []EndpointAgentToServerTest {
	if o == nil || utils.IsNil(o.Tests) {
		var ret []EndpointAgentToServerTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTests) GetTestsOk() ([]EndpointAgentToServerTest, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *EndpointAgentToServerTests) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []EndpointAgentToServerTest and assigns it to the Tests field.
func (o *EndpointAgentToServerTests) SetTests(v []EndpointAgentToServerTest) {
	o.Tests = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EndpointAgentToServerTests) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTests) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EndpointAgentToServerTests) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *EndpointAgentToServerTests) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o EndpointAgentToServerTests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAgentToServerTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableEndpointAgentToServerTests struct {
	value *EndpointAgentToServerTests
	isSet bool
}

func (v NullableEndpointAgentToServerTests) Get() *EndpointAgentToServerTests {
	return v.value
}

func (v *NullableEndpointAgentToServerTests) Set(val *EndpointAgentToServerTests) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAgentToServerTests) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAgentToServerTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAgentToServerTests(val *EndpointAgentToServerTests) *NullableEndpointAgentToServerTests {
	return &NullableEndpointAgentToServerTests{value: val, isSet: true}
}

func (v NullableEndpointAgentToServerTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAgentToServerTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


