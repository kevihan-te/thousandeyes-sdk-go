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

// checks if the NetworkInterface type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkInterface{}

// NetworkInterface struct for NetworkInterface
type NetworkInterface struct {
	// Network IP address.
	IpAddress *string `json:"ipAddress,omitempty"`
	// Network subnet mask - only for IPv4.
	SubnetMask *string `json:"subnetMask,omitempty"`
	// Network public IP address.
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
	// Network local prefix.
	LocalPrefix *string `json:"localPrefix,omitempty"`
	// Network public IP range.
	PublicIpRange *string `json:"publicIpRange,omitempty"`
	// Network DNS servers.
	DnsServers []string `json:"dnsServers,omitempty"`
	HardwareType *InterfaceHardwareType `json:"hardwareType,omitempty"`
	// Network interface name.
	InterfaceName *string `json:"interfaceName,omitempty"`
}

// NewNetworkInterface instantiates a new NetworkInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkInterface() *NetworkInterface {
	this := NetworkInterface{}
	return &this
}

// NewNetworkInterfaceWithDefaults instantiates a new NetworkInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkInterfaceWithDefaults() *NetworkInterface {
	this := NetworkInterface{}
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *NetworkInterface) GetIpAddress() string {
	if o == nil || utils.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetIpAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *NetworkInterface) HasIpAddress() bool {
	if o != nil && !utils.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *NetworkInterface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetSubnetMask returns the SubnetMask field value if set, zero value otherwise.
func (o *NetworkInterface) GetSubnetMask() string {
	if o == nil || utils.IsNil(o.SubnetMask) {
		var ret string
		return ret
	}
	return *o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetSubnetMaskOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SubnetMask) {
		return nil, false
	}
	return o.SubnetMask, true
}

// HasSubnetMask returns a boolean if a field has been set.
func (o *NetworkInterface) HasSubnetMask() bool {
	if o != nil && !utils.IsNil(o.SubnetMask) {
		return true
	}

	return false
}

// SetSubnetMask gets a reference to the given string and assigns it to the SubnetMask field.
func (o *NetworkInterface) SetSubnetMask(v string) {
	o.SubnetMask = &v
}

// GetPublicIpAddress returns the PublicIpAddress field value if set, zero value otherwise.
func (o *NetworkInterface) GetPublicIpAddress() string {
	if o == nil || utils.IsNil(o.PublicIpAddress) {
		var ret string
		return ret
	}
	return *o.PublicIpAddress
}

// GetPublicIpAddressOk returns a tuple with the PublicIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetPublicIpAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PublicIpAddress) {
		return nil, false
	}
	return o.PublicIpAddress, true
}

// HasPublicIpAddress returns a boolean if a field has been set.
func (o *NetworkInterface) HasPublicIpAddress() bool {
	if o != nil && !utils.IsNil(o.PublicIpAddress) {
		return true
	}

	return false
}

// SetPublicIpAddress gets a reference to the given string and assigns it to the PublicIpAddress field.
func (o *NetworkInterface) SetPublicIpAddress(v string) {
	o.PublicIpAddress = &v
}

// GetLocalPrefix returns the LocalPrefix field value if set, zero value otherwise.
func (o *NetworkInterface) GetLocalPrefix() string {
	if o == nil || utils.IsNil(o.LocalPrefix) {
		var ret string
		return ret
	}
	return *o.LocalPrefix
}

// GetLocalPrefixOk returns a tuple with the LocalPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetLocalPrefixOk() (*string, bool) {
	if o == nil || utils.IsNil(o.LocalPrefix) {
		return nil, false
	}
	return o.LocalPrefix, true
}

// HasLocalPrefix returns a boolean if a field has been set.
func (o *NetworkInterface) HasLocalPrefix() bool {
	if o != nil && !utils.IsNil(o.LocalPrefix) {
		return true
	}

	return false
}

// SetLocalPrefix gets a reference to the given string and assigns it to the LocalPrefix field.
func (o *NetworkInterface) SetLocalPrefix(v string) {
	o.LocalPrefix = &v
}

// GetPublicIpRange returns the PublicIpRange field value if set, zero value otherwise.
func (o *NetworkInterface) GetPublicIpRange() string {
	if o == nil || utils.IsNil(o.PublicIpRange) {
		var ret string
		return ret
	}
	return *o.PublicIpRange
}

// GetPublicIpRangeOk returns a tuple with the PublicIpRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetPublicIpRangeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PublicIpRange) {
		return nil, false
	}
	return o.PublicIpRange, true
}

// HasPublicIpRange returns a boolean if a field has been set.
func (o *NetworkInterface) HasPublicIpRange() bool {
	if o != nil && !utils.IsNil(o.PublicIpRange) {
		return true
	}

	return false
}

// SetPublicIpRange gets a reference to the given string and assigns it to the PublicIpRange field.
func (o *NetworkInterface) SetPublicIpRange(v string) {
	o.PublicIpRange = &v
}

// GetDnsServers returns the DnsServers field value if set, zero value otherwise.
func (o *NetworkInterface) GetDnsServers() []string {
	if o == nil || utils.IsNil(o.DnsServers) {
		var ret []string
		return ret
	}
	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetDnsServersOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.DnsServers) {
		return nil, false
	}
	return o.DnsServers, true
}

// HasDnsServers returns a boolean if a field has been set.
func (o *NetworkInterface) HasDnsServers() bool {
	if o != nil && !utils.IsNil(o.DnsServers) {
		return true
	}

	return false
}

// SetDnsServers gets a reference to the given []string and assigns it to the DnsServers field.
func (o *NetworkInterface) SetDnsServers(v []string) {
	o.DnsServers = v
}

// GetHardwareType returns the HardwareType field value if set, zero value otherwise.
func (o *NetworkInterface) GetHardwareType() InterfaceHardwareType {
	if o == nil || utils.IsNil(o.HardwareType) {
		var ret InterfaceHardwareType
		return ret
	}
	return *o.HardwareType
}

// GetHardwareTypeOk returns a tuple with the HardwareType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetHardwareTypeOk() (*InterfaceHardwareType, bool) {
	if o == nil || utils.IsNil(o.HardwareType) {
		return nil, false
	}
	return o.HardwareType, true
}

// HasHardwareType returns a boolean if a field has been set.
func (o *NetworkInterface) HasHardwareType() bool {
	if o != nil && !utils.IsNil(o.HardwareType) {
		return true
	}

	return false
}

// SetHardwareType gets a reference to the given InterfaceHardwareType and assigns it to the HardwareType field.
func (o *NetworkInterface) SetHardwareType(v InterfaceHardwareType) {
	o.HardwareType = &v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *NetworkInterface) GetInterfaceName() string {
	if o == nil || utils.IsNil(o.InterfaceName) {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkInterface) GetInterfaceNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.InterfaceName) {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *NetworkInterface) HasInterfaceName() bool {
	if o != nil && !utils.IsNil(o.InterfaceName) {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *NetworkInterface) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

func (o NetworkInterface) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !utils.IsNil(o.SubnetMask) {
		toSerialize["subnetMask"] = o.SubnetMask
	}
	if !utils.IsNil(o.PublicIpAddress) {
		toSerialize["publicIpAddress"] = o.PublicIpAddress
	}
	if !utils.IsNil(o.LocalPrefix) {
		toSerialize["localPrefix"] = o.LocalPrefix
	}
	if !utils.IsNil(o.PublicIpRange) {
		toSerialize["publicIpRange"] = o.PublicIpRange
	}
	if !utils.IsNil(o.DnsServers) {
		toSerialize["dnsServers"] = o.DnsServers
	}
	if !utils.IsNil(o.HardwareType) {
		toSerialize["hardwareType"] = o.HardwareType
	}
	if !utils.IsNil(o.InterfaceName) {
		toSerialize["interfaceName"] = o.InterfaceName
	}
	return toSerialize, nil
}

type NullableNetworkInterface struct {
	value *NetworkInterface
	isSet bool
}

func (v NullableNetworkInterface) Get() *NetworkInterface {
	return v.value
}

func (v *NullableNetworkInterface) Set(val *NetworkInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkInterface(val *NetworkInterface) *NullableNetworkInterface {
	return &NullableNetworkInterface{value: val, isSet: true}
}

func (v NullableNetworkInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


