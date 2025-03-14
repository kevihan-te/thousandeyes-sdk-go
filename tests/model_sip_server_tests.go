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

// checks if the SipServerTests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SipServerTests{}

// SipServerTests struct for SipServerTests
type SipServerTests struct {
	Tests []UnexpandedSipServerTest `json:"tests,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewSipServerTests instantiates a new SipServerTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSipServerTests() *SipServerTests {
	this := SipServerTests{}
	return &this
}

// NewSipServerTestsWithDefaults instantiates a new SipServerTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSipServerTestsWithDefaults() *SipServerTests {
	this := SipServerTests{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *SipServerTests) GetTests() []UnexpandedSipServerTest {
	if o == nil || utils.IsNil(o.Tests) {
		var ret []UnexpandedSipServerTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipServerTests) GetTestsOk() ([]UnexpandedSipServerTest, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *SipServerTests) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []UnexpandedSipServerTest and assigns it to the Tests field.
func (o *SipServerTests) SetTests(v []UnexpandedSipServerTest) {
	o.Tests = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SipServerTests) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SipServerTests) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SipServerTests) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *SipServerTests) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o SipServerTests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SipServerTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSipServerTests struct {
	value *SipServerTests
	isSet bool
}

func (v NullableSipServerTests) Get() *SipServerTests {
	return v.value
}

func (v *NullableSipServerTests) Set(val *SipServerTests) {
	v.value = val
	v.isSet = true
}

func (v NullableSipServerTests) IsSet() bool {
	return v.isSet
}

func (v *NullableSipServerTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipServerTests(val *SipServerTests) *NullableSipServerTests {
	return &NullableSipServerTests{value: val, isSet: true}
}

func (v NullableSipServerTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipServerTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


