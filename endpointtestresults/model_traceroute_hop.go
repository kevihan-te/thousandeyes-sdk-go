/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the TracerouteHop type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TracerouteHop{}

// TracerouteHop struct for TracerouteHop
type TracerouteHop struct {
	// The hop index.
	Hop *int32 `json:"hop,omitempty"`
	// IP address of the hop.
	IpAddress *string `json:"ipAddress,omitempty"`
	// Prefix of IP address shown in CIDR.
	Prefix *string `json:"prefix,omitempty"`
	// Unique number assigned to an organization (also referred to as service provider).
	Asn *int32 `json:"asn,omitempty"`
	// Hop delay
	Delay *int32 `json:"delay,omitempty"`
	// Hop Multiprotocol Label Switching.
	Mpls []string `json:"mpls,omitempty"`
	// The hop name.
	Name *string `json:"name,omitempty"`
}

// NewTracerouteHop instantiates a new TracerouteHop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTracerouteHop() *TracerouteHop {
	this := TracerouteHop{}
	return &this
}

// NewTracerouteHopWithDefaults instantiates a new TracerouteHop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTracerouteHopWithDefaults() *TracerouteHop {
	this := TracerouteHop{}
	return &this
}

// GetHop returns the Hop field value if set, zero value otherwise.
func (o *TracerouteHop) GetHop() int32 {
	if o == nil || utils.IsNil(o.Hop) {
		var ret int32
		return ret
	}
	return *o.Hop
}

// GetHopOk returns a tuple with the Hop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TracerouteHop) GetHopOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Hop) {
		return nil, false
	}
	return o.Hop, true
}

// HasHop returns a boolean if a field has been set.
func (o *TracerouteHop) HasHop() bool {
	if o != nil && !utils.IsNil(o.Hop) {
		return true
	}

	return false
}

// SetHop gets a reference to the given int32 and assigns it to the Hop field.
func (o *TracerouteHop) SetHop(v int32) {
	o.Hop = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *TracerouteHop) GetIpAddress() string {
	if o == nil || utils.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TracerouteHop) GetIpAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *TracerouteHop) HasIpAddress() bool {
	if o != nil && !utils.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *TracerouteHop) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *TracerouteHop) GetPrefix() string {
	if o == nil || utils.IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TracerouteHop) GetPrefixOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *TracerouteHop) HasPrefix() bool {
	if o != nil && !utils.IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *TracerouteHop) SetPrefix(v string) {
	o.Prefix = &v
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *TracerouteHop) GetAsn() int32 {
	if o == nil || utils.IsNil(o.Asn) {
		var ret int32
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TracerouteHop) GetAsnOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *TracerouteHop) HasAsn() bool {
	if o != nil && !utils.IsNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int32 and assigns it to the Asn field.
func (o *TracerouteHop) SetAsn(v int32) {
	o.Asn = &v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *TracerouteHop) GetDelay() int32 {
	if o == nil || utils.IsNil(o.Delay) {
		var ret int32
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TracerouteHop) GetDelayOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *TracerouteHop) HasDelay() bool {
	if o != nil && !utils.IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int32 and assigns it to the Delay field.
func (o *TracerouteHop) SetDelay(v int32) {
	o.Delay = &v
}

// GetMpls returns the Mpls field value if set, zero value otherwise.
func (o *TracerouteHop) GetMpls() []string {
	if o == nil || utils.IsNil(o.Mpls) {
		var ret []string
		return ret
	}
	return o.Mpls
}

// GetMplsOk returns a tuple with the Mpls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TracerouteHop) GetMplsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Mpls) {
		return nil, false
	}
	return o.Mpls, true
}

// HasMpls returns a boolean if a field has been set.
func (o *TracerouteHop) HasMpls() bool {
	if o != nil && !utils.IsNil(o.Mpls) {
		return true
	}

	return false
}

// SetMpls gets a reference to the given []string and assigns it to the Mpls field.
func (o *TracerouteHop) SetMpls(v []string) {
	o.Mpls = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TracerouteHop) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TracerouteHop) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TracerouteHop) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TracerouteHop) SetName(v string) {
	o.Name = &v
}

func (o TracerouteHop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TracerouteHop) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Hop) {
		toSerialize["hop"] = o.Hop
	}
	if !utils.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !utils.IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !utils.IsNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !utils.IsNil(o.Delay) {
		toSerialize["delay"] = o.Delay
	}
	if !utils.IsNil(o.Mpls) {
		toSerialize["mpls"] = o.Mpls
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableTracerouteHop struct {
	value *TracerouteHop
	isSet bool
}

func (v NullableTracerouteHop) Get() *TracerouteHop {
	return v.value
}

func (v *NullableTracerouteHop) Set(val *TracerouteHop) {
	v.value = val
	v.isSet = true
}

func (v NullableTracerouteHop) IsSet() bool {
	return v.isSet
}

func (v *NullableTracerouteHop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTracerouteHop(val *TracerouteHop) *NullableTracerouteHop {
	return &NullableTracerouteHop{value: val, isSet: true}
}

func (v NullableTracerouteHop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTracerouteHop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


