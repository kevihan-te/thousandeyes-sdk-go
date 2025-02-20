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

// checks if the BgpTests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BgpTests{}

// BgpTests struct for BgpTests
type BgpTests struct {
	Tests []UnexpandedBgpTest `json:"tests,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewBgpTests instantiates a new BgpTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpTests() *BgpTests {
	this := BgpTests{}
	return &this
}

// NewBgpTestsWithDefaults instantiates a new BgpTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpTestsWithDefaults() *BgpTests {
	this := BgpTests{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *BgpTests) GetTests() []UnexpandedBgpTest {
	if o == nil || utils.IsNil(o.Tests) {
		var ret []UnexpandedBgpTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTests) GetTestsOk() ([]UnexpandedBgpTest, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *BgpTests) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []UnexpandedBgpTest and assigns it to the Tests field.
func (o *BgpTests) SetTests(v []UnexpandedBgpTest) {
	o.Tests = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BgpTests) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTests) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BgpTests) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *BgpTests) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o BgpTests) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpTests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableBgpTests struct {
	value *BgpTests
	isSet bool
}

func (v NullableBgpTests) Get() *BgpTests {
	return v.value
}

func (v *NullableBgpTests) Set(val *BgpTests) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpTests) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpTests(val *BgpTests) *NullableBgpTests {
	return &NullableBgpTests{value: val, isSet: true}
}

func (v NullableBgpTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


