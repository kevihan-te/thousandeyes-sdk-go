/*
Endpoint Agents API

Manage ThousandEyes Endpoint Agents using this API.   For more information about Endpoint Agents, see [Endpoint Agents](https://docs.thousandeyes.com/product-documentation/global-vantage-points/endpoint-agents).

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointagents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the EndpointAgentEthernetProfile type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointAgentEthernetProfile{}

// EndpointAgentEthernetProfile Information about the ethernet connectivity of this device. Only present if the hardware type is `ethernet`. 
type EndpointAgentEthernetProfile struct {
	// Link speed in Mbps.
	LinkSpeed *int32 `json:"linkSpeed,omitempty"`
}

// NewEndpointAgentEthernetProfile instantiates a new EndpointAgentEthernetProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAgentEthernetProfile() *EndpointAgentEthernetProfile {
	this := EndpointAgentEthernetProfile{}
	return &this
}

// NewEndpointAgentEthernetProfileWithDefaults instantiates a new EndpointAgentEthernetProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAgentEthernetProfileWithDefaults() *EndpointAgentEthernetProfile {
	this := EndpointAgentEthernetProfile{}
	return &this
}

// GetLinkSpeed returns the LinkSpeed field value if set, zero value otherwise.
func (o *EndpointAgentEthernetProfile) GetLinkSpeed() int32 {
	if o == nil || utils.IsNil(o.LinkSpeed) {
		var ret int32
		return ret
	}
	return *o.LinkSpeed
}

// GetLinkSpeedOk returns a tuple with the LinkSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentEthernetProfile) GetLinkSpeedOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.LinkSpeed) {
		return nil, false
	}
	return o.LinkSpeed, true
}

// HasLinkSpeed returns a boolean if a field has been set.
func (o *EndpointAgentEthernetProfile) HasLinkSpeed() bool {
	if o != nil && !utils.IsNil(o.LinkSpeed) {
		return true
	}

	return false
}

// SetLinkSpeed gets a reference to the given int32 and assigns it to the LinkSpeed field.
func (o *EndpointAgentEthernetProfile) SetLinkSpeed(v int32) {
	o.LinkSpeed = &v
}

func (o EndpointAgentEthernetProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAgentEthernetProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.LinkSpeed) {
		toSerialize["linkSpeed"] = o.LinkSpeed
	}
	return toSerialize, nil
}

type NullableEndpointAgentEthernetProfile struct {
	value *EndpointAgentEthernetProfile
	isSet bool
}

func (v NullableEndpointAgentEthernetProfile) Get() *EndpointAgentEthernetProfile {
	return v.value
}

func (v *NullableEndpointAgentEthernetProfile) Set(val *EndpointAgentEthernetProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAgentEthernetProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAgentEthernetProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAgentEthernetProfile(val *EndpointAgentEthernetProfile) *NullableEndpointAgentEthernetProfile {
	return &NullableEndpointAgentEthernetProfile{value: val, isSet: true}
}

func (v NullableEndpointAgentEthernetProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAgentEthernetProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


