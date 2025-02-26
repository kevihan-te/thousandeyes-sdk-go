/*
Agents API

 ## Overview Manage Cloud and Enterprise Agents available to your account in ThousandEyes.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the InterfaceIpMapping type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &InterfaceIpMapping{}

// InterfaceIpMapping struct for InterfaceIpMapping
type InterfaceIpMapping struct {
	// Name of the mapping
	InterfaceName *string `json:"interfaceName,omitempty"`
	// Array of ipAddress entries
	IpAddresses []string `json:"ipAddresses,omitempty"`
}

// NewInterfaceIpMapping instantiates a new InterfaceIpMapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInterfaceIpMapping() *InterfaceIpMapping {
	this := InterfaceIpMapping{}
	return &this
}

// NewInterfaceIpMappingWithDefaults instantiates a new InterfaceIpMapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInterfaceIpMappingWithDefaults() *InterfaceIpMapping {
	this := InterfaceIpMapping{}
	return &this
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *InterfaceIpMapping) GetInterfaceName() string {
	if o == nil || utils.IsNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceIpMapping) GetInterfaceNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InterfaceName) {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *InterfaceIpMapping) HasInterfaceName() bool {
	if o != nil && !utils.IsNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *InterfaceIpMapping) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetIpAddresses returns the IpAddresses field value if set, zero value otherwise.
func (o *InterfaceIpMapping) GetIpAddresses() []string {
	if o == nil || utils.IsNil(o.IpAddresses) {
		var ret []string
		return ret
	}
	return o.IpAddresses
}

// GetIpAddressesOk returns a tuple with the IpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InterfaceIpMapping) GetIpAddressesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IpAddresses) {
		return nil, false
	}
	return o.IpAddresses, true
}

// HasIpAddresses returns a boolean if a field has been set.
func (o *InterfaceIpMapping) HasIpAddresses() bool {
	if o != nil && !utils.IsNil(o.IpAddresses) {
		return true
	}

	return false
}

// SetIpAddresses gets a reference to the given []string and assigns it to the IpAddresses field.
func (o *InterfaceIpMapping) SetIpAddresses(v []string) {
	o.IpAddresses = v
}

func (o InterfaceIpMapping) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InterfaceIpMapping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.InterfaceName) {
		toSerialize["interfaceName"] = o.InterfaceName
	}
	if !utils.IsNil(o.IpAddresses) {
		toSerialize["ipAddresses"] = o.IpAddresses
	}
	return toSerialize, nil
}

type NullableInterfaceIpMapping struct {
	value *InterfaceIpMapping
	isSet bool
}

func (v NullableInterfaceIpMapping) Get() *InterfaceIpMapping {
	return v.value
}

func (v *NullableInterfaceIpMapping) Set(val *InterfaceIpMapping) {
	v.value = val
	v.isSet = true
}

func (v NullableInterfaceIpMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableInterfaceIpMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInterfaceIpMapping(val *InterfaceIpMapping) *NullableInterfaceIpMapping {
	return &NullableInterfaceIpMapping{value: val, isSet: true}
}

func (v NullableInterfaceIpMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInterfaceIpMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


