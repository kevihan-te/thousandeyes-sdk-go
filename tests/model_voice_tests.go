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
)

// checks if the VoiceTests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VoiceTests{}

// VoiceTests struct for VoiceTests
type VoiceTests struct {
	Tests []UnexpandedVoiceTest `json:"tests,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewVoiceTests instantiates a new VoiceTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceTests() *VoiceTests {
	this := VoiceTests{}
	return &this
}

// NewVoiceTestsWithDefaults instantiates a new VoiceTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceTestsWithDefaults() *VoiceTests {
	this := VoiceTests{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *VoiceTests) GetTests() []UnexpandedVoiceTest {
	if o == nil || utils.IsNil(o.Tests) {
		var ret []UnexpandedVoiceTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTests) GetTestsOk() ([]UnexpandedVoiceTest, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *VoiceTests) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []UnexpandedVoiceTest and assigns it to the Tests field.
func (o *VoiceTests) SetTests(v []UnexpandedVoiceTest) {
	o.Tests = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VoiceTests) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTests) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VoiceTests) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *VoiceTests) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o VoiceTests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableVoiceTests struct {
	value *VoiceTests
	isSet bool
}

func (v NullableVoiceTests) Get() *VoiceTests {
	return v.value
}

func (v *NullableVoiceTests) Set(val *VoiceTests) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceTests) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceTests(val *VoiceTests) *NullableVoiceTests {
	return &NullableVoiceTests{value: val, isSet: true}
}

func (v NullableVoiceTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


