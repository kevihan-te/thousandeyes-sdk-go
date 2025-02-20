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

// checks if the InterfaceGroups type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InterfaceGroups{}

// InterfaceGroups struct for InterfaceGroups
type InterfaceGroups struct {
	PathVisInterfaceGroups []InterfaceGroup `json:"pathVisInterfaceGroups,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewInterfaceGroups instantiates a new InterfaceGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceGroups() *InterfaceGroups {
	this := InterfaceGroups{}
	return &this
}

// NewInterfaceGroupsWithDefaults instantiates a new InterfaceGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceGroupsWithDefaults() *InterfaceGroups {
	this := InterfaceGroups{}
	return &this
}

// GetPathVisInterfaceGroups returns the PathVisInterfaceGroups field value if set, zero value otherwise.
func (o *InterfaceGroups) GetPathVisInterfaceGroups() []InterfaceGroup {
	if o == nil || utils.IsNil(o.PathVisInterfaceGroups) {
		var ret []InterfaceGroup
		return ret
	}
	return o.PathVisInterfaceGroups
}

// GetPathVisInterfaceGroupsOk returns a tuple with the PathVisInterfaceGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceGroups) GetPathVisInterfaceGroupsOk() ([]InterfaceGroup, bool) {
	if o == nil || utils.IsNil(o.PathVisInterfaceGroups) {
		return nil, false
	}
	return o.PathVisInterfaceGroups, true
}

// HasPathVisInterfaceGroups returns a boolean if a field has been set.
func (o *InterfaceGroups) HasPathVisInterfaceGroups() bool {
	if o != nil && !utils.IsNil(o.PathVisInterfaceGroups) {
		return true
	}

	return false
}

// SetPathVisInterfaceGroups gets a reference to the given []InterfaceGroup and assigns it to the PathVisInterfaceGroups field.
func (o *InterfaceGroups) SetPathVisInterfaceGroups(v []InterfaceGroup) {
	o.PathVisInterfaceGroups = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *InterfaceGroups) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceGroups) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *InterfaceGroups) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *InterfaceGroups) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o InterfaceGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PathVisInterfaceGroups) {
		toSerialize["pathVisInterfaceGroups"] = o.PathVisInterfaceGroups
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableInterfaceGroups struct {
	value *InterfaceGroups
	isSet bool
}

func (v NullableInterfaceGroups) Get() *InterfaceGroups {
	return v.value
}

func (v *NullableInterfaceGroups) Set(val *InterfaceGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceGroups(val *InterfaceGroups) *NullableInterfaceGroups {
	return &NullableInterfaceGroups{value: val, isSet: true}
}

func (v NullableInterfaceGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


