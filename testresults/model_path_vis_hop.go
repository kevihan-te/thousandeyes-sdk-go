/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the PathVisHop type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PathVisHop{}

// PathVisHop struct for PathVisHop
type PathVisHop struct {
	// Hop index
	Hop *int32 `json:"hop,omitempty"`
	// IP address of the hop
	IpAddress *string `json:"ipAddress,omitempty"`
	// Prefix of IP address shown in CIDR
	Prefix *string `json:"prefix,omitempty"`
	// Reverse DNS entry of IP, if available
	Rdns *string `json:"rdns,omitempty"`
	// Autonomous System originating the prefix
	Network *string `json:"network,omitempty"`
	// RTT to the hop’s IP in milliseconds
	ResponseTime *int32 `json:"responseTime,omitempty"`
	// Location information for the hop
	Location *string `json:"location,omitempty"`
	// Multiprotocol Label Switching information, if available
	Mpls *string `json:"mpls,omitempty"`
}

// NewPathVisHop instantiates a new PathVisHop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathVisHop() *PathVisHop {
	this := PathVisHop{}
	return &this
}

// NewPathVisHopWithDefaults instantiates a new PathVisHop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathVisHopWithDefaults() *PathVisHop {
	this := PathVisHop{}
	return &this
}

// GetHop returns the Hop field value if set, zero value otherwise.
func (o *PathVisHop) GetHop() int32 {
	if o == nil || utils.IsNil(o.Hop) {
		var ret int32
		return ret
	}
	return *o.Hop
}

// GetHopOk returns a tuple with the Hop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetHopOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Hop) {
		return nil, false
	}
	return o.Hop, true
}

// HasHop returns a boolean if a field has been set.
func (o *PathVisHop) HasHop() bool {
	if o != nil && !utils.IsNil(o.Hop) {
		return true
	}

	return false
}

// SetHop gets a reference to the given int32 and assigns it to the Hop field.
func (o *PathVisHop) SetHop(v int32) {
	o.Hop = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *PathVisHop) GetIpAddress() string {
	if o == nil || utils.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetIpAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *PathVisHop) HasIpAddress() bool {
	if o != nil && !utils.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *PathVisHop) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *PathVisHop) GetPrefix() string {
	if o == nil || utils.IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetPrefixOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *PathVisHop) HasPrefix() bool {
	if o != nil && !utils.IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *PathVisHop) SetPrefix(v string) {
	o.Prefix = &v
}

// GetRdns returns the Rdns field value if set, zero value otherwise.
func (o *PathVisHop) GetRdns() string {
	if o == nil || utils.IsNil(o.Rdns) {
		var ret string
		return ret
	}
	return *o.Rdns
}

// GetRdnsOk returns a tuple with the Rdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetRdnsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Rdns) {
		return nil, false
	}
	return o.Rdns, true
}

// HasRdns returns a boolean if a field has been set.
func (o *PathVisHop) HasRdns() bool {
	if o != nil && !utils.IsNil(o.Rdns) {
		return true
	}

	return false
}

// SetRdns gets a reference to the given string and assigns it to the Rdns field.
func (o *PathVisHop) SetRdns(v string) {
	o.Rdns = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *PathVisHop) GetNetwork() string {
	if o == nil || utils.IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetNetworkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *PathVisHop) HasNetwork() bool {
	if o != nil && !utils.IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *PathVisHop) SetNetwork(v string) {
	o.Network = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *PathVisHop) GetResponseTime() int32 {
	if o == nil || utils.IsNil(o.ResponseTime) {
		var ret int32
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetResponseTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *PathVisHop) HasResponseTime() bool {
	if o != nil && !utils.IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given int32 and assigns it to the ResponseTime field.
func (o *PathVisHop) SetResponseTime(v int32) {
	o.ResponseTime = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PathVisHop) GetLocation() string {
	if o == nil || utils.IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetLocationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PathVisHop) HasLocation() bool {
	if o != nil && !utils.IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PathVisHop) SetLocation(v string) {
	o.Location = &v
}

// GetMpls returns the Mpls field value if set, zero value otherwise.
func (o *PathVisHop) GetMpls() string {
	if o == nil || utils.IsNil(o.Mpls) {
		var ret string
		return ret
	}
	return *o.Mpls
}

// GetMplsOk returns a tuple with the Mpls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisHop) GetMplsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Mpls) {
		return nil, false
	}
	return o.Mpls, true
}

// HasMpls returns a boolean if a field has been set.
func (o *PathVisHop) HasMpls() bool {
	if o != nil && !utils.IsNil(o.Mpls) {
		return true
	}

	return false
}

// SetMpls gets a reference to the given string and assigns it to the Mpls field.
func (o *PathVisHop) SetMpls(v string) {
	o.Mpls = &v
}

func (o PathVisHop) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathVisHop) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Rdns) {
		toSerialize["rdns"] = o.Rdns
	}
	if !utils.IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !utils.IsNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	if !utils.IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !utils.IsNil(o.Mpls) {
		toSerialize["mpls"] = o.Mpls
	}
	return toSerialize, nil
}

type NullablePathVisHop struct {
	value *PathVisHop
	isSet bool
}

func (v NullablePathVisHop) Get() *PathVisHop {
	return v.value
}

func (v *NullablePathVisHop) Set(val *PathVisHop) {
	v.value = val
	v.isSet = true
}

func (v NullablePathVisHop) IsSet() bool {
	return v.isSet
}

func (v *NullablePathVisHop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathVisHop(val *PathVisHop) *NullablePathVisHop {
	return &NullablePathVisHop{value: val, isSet: true}
}

func (v NullablePathVisHop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathVisHop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


