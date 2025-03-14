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

// checks if the InterfaceGroup type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InterfaceGroup{}

// InterfaceGroup struct for InterfaceGroup
type InterfaceGroup struct {
	// Account Group Id
	Aid *string `json:"aid,omitempty"`
	// Group ID
	GroupId *string `json:"groupId,omitempty"`
	// Name of the path visualization interface group
	GroupName *string `json:"groupName,omitempty"`
	// Array of IP addresses associated with the interface group
	IpAddresses []string `json:"ipAddresses,omitempty"`
	// Array of RDNS Regexes associated with the interface group
	RdnsRegexes []string `json:"rdnsRegexes,omitempty"`
}

// NewInterfaceGroup instantiates a new InterfaceGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceGroup() *InterfaceGroup {
	this := InterfaceGroup{}
	return &this
}

// NewInterfaceGroupWithDefaults instantiates a new InterfaceGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceGroupWithDefaults() *InterfaceGroup {
	this := InterfaceGroup{}
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *InterfaceGroup) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceGroup) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *InterfaceGroup) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *InterfaceGroup) SetAid(v string) {
	o.Aid = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *InterfaceGroup) GetGroupId() string {
	if o == nil || utils.IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceGroup) GetGroupIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *InterfaceGroup) HasGroupId() bool {
	if o != nil && !utils.IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *InterfaceGroup) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *InterfaceGroup) GetGroupName() string {
	if o == nil || utils.IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceGroup) GetGroupNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *InterfaceGroup) HasGroupName() bool {
	if o != nil && !utils.IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *InterfaceGroup) SetGroupName(v string) {
	o.GroupName = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *InterfaceGroup) GetIpAddresses() []string {
	if o == nil || utils.IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceGroup) GetIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *InterfaceGroup) HasIpAddresses() bool {
	if o != nil && !utils.IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *InterfaceGroup) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

// GetRdnsRegexes returns the RdnsRegexes field value if set, zero value otherwise.
func (o *InterfaceGroup) GetRdnsRegexes() []string {
	if o == nil || utils.IsNil(o.RdnsRegexes) {
		var ret []string
		return ret
	}
	return o.RdnsRegexes
}

// GetRdnsRegexesOk returns a tuple with the RdnsRegexes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceGroup) GetRdnsRegexesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.RdnsRegexes) {
		return nil, false
	}
	return o.RdnsRegexes, true
}

// HasRdnsRegexes returns a boolean if a field has been set.
func (o *InterfaceGroup) HasRdnsRegexes() bool {
	if o != nil && !utils.IsNil(o.RdnsRegexes) {
		return true
	}

	return false
}

// SetRdnsRegexes gets a reference to the given []string and assigns it to the RdnsRegexes field.
func (o *InterfaceGroup) SetRdnsRegexes(v []string) {
	o.RdnsRegexes = v
}

func (o InterfaceGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !utils.IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !utils.IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	if !utils.IsNil(o.RdnsRegexes) {
		toSerialize["rdnsRegexes"] = o.RdnsRegexes
	}
	return toSerialize, nil
}

type NullableInterfaceGroup struct {
	value *InterfaceGroup
	isSet bool
}

func (v NullableInterfaceGroup) Get() *InterfaceGroup {
	return v.value
}

func (v *NullableInterfaceGroup) Set(val *InterfaceGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceGroup(val *InterfaceGroup) *NullableInterfaceGroup {
	return &NullableInterfaceGroup{value: val, isSet: true}
}

func (v NullableInterfaceGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


