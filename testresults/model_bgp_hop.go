/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the BgpHop type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BgpHop{}

// BgpHop struct for BgpHop
type BgpHop struct {
	// ASN of transit autonomous system
	Asn *int32 `json:"asn,omitempty"`
	// Name of autonomous system.
	AsName *string `json:"asName,omitempty"`
}

// NewBgpHop instantiates a new BgpHop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpHop() *BgpHop {
	this := BgpHop{}
	return &this
}

// NewBgpHopWithDefaults instantiates a new BgpHop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpHopWithDefaults() *BgpHop {
	this := BgpHop{}
	return &this
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *BgpHop) GetAsn() int32 {
	if o == nil || utils.IsNil(o.Asn) {
		var ret int32
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpHop) GetAsnOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *BgpHop) HasAsn() bool {
	if o != nil && !utils.IsNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int32 and assigns it to the Asn field.
func (o *BgpHop) SetAsn(v int32) {
	o.Asn = &v
}

// GetAsName returns the AsName field value if set, zero value otherwise.
func (o *BgpHop) GetAsName() string {
	if o == nil || utils.IsNil(o.AsName) {
		var ret string
		return ret
	}
	return *o.AsName
}

// GetAsNameOk returns a tuple with the AsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpHop) GetAsNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AsName) {
		return nil, false
	}
	return o.AsName, true
}

// HasAsName returns a boolean if a field has been set.
func (o *BgpHop) HasAsName() bool {
	if o != nil && !utils.IsNil(o.AsName) {
		return true
	}

	return false
}

// SetAsName gets a reference to the given string and assigns it to the AsName field.
func (o *BgpHop) SetAsName(v string) {
	o.AsName = &v
}

func (o BgpHop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpHop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !utils.IsNil(o.AsName) {
		toSerialize["asName"] = o.AsName
	}
	return toSerialize, nil
}

type NullableBgpHop struct {
	value *BgpHop
	isSet bool
}

func (v NullableBgpHop) Get() *BgpHop {
	return v.value
}

func (v *NullableBgpHop) Set(val *BgpHop) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpHop) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpHop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpHop(val *BgpHop) *NullableBgpHop {
	return &NullableBgpHop{value: val, isSet: true}
}

func (v NullableBgpHop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpHop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


