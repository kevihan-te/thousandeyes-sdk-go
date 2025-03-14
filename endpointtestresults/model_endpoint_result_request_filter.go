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

// checks if the EndpointResultRequestFilter type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointResultRequestFilter{}

// EndpointResultRequestFilter struct for EndpointResultRequestFilter
type EndpointResultRequestFilter struct {
	// Location of the endpoint agent.
	Location []string `json:"location,omitempty"`
	Connection []InterfaceHardwareType `json:"connection,omitempty"`
	Platform []Platform `json:"platform,omitempty"`
	// Endpoint agent default gateway IP address.
	Gateway []string `json:"gateway,omitempty"`
	// Endpoint agent proxy IP address.
	ProxyTarget []string `json:"proxyTarget,omitempty"`
	// Endpoint agent VPN endpoint IP address.
	VpnTarget []string `json:"vpnTarget,omitempty"`
	// Endpoint agent ID.
	AgentId []string `json:"agentId,omitempty"`
	// Network ID.
	NetworkId []string `json:"networkId,omitempty"`
	// WiFi SSID.
	Ssid []string `json:"ssid,omitempty"`
	// WiFi BSSID.
	Bssid []string `json:"bssid,omitempty"`
}

// NewEndpointResultRequestFilter instantiates a new EndpointResultRequestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointResultRequestFilter() *EndpointResultRequestFilter {
	this := EndpointResultRequestFilter{}
	return &this
}

// NewEndpointResultRequestFilterWithDefaults instantiates a new EndpointResultRequestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointResultRequestFilterWithDefaults() *EndpointResultRequestFilter {
	this := EndpointResultRequestFilter{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetLocation() []string {
	if o == nil || utils.IsNil(o.Location) {
		var ret []string
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetLocationOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasLocation() bool {
	if o != nil && !utils.IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given []string and assigns it to the Location field.
func (o *EndpointResultRequestFilter) SetLocation(v []string) {
	o.Location = v
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetConnection() []InterfaceHardwareType {
	if o == nil || utils.IsNil(o.Connection) {
		var ret []InterfaceHardwareType
		return ret
	}
	return o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetConnectionOk() ([]InterfaceHardwareType, bool) {
	if o == nil || utils.IsNil(o.Connection) {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasConnection() bool {
	if o != nil && !utils.IsNil(o.Connection) {
		return true
	}

	return false
}

// SetConnection gets a reference to the given []InterfaceHardwareType and assigns it to the Connection field.
func (o *EndpointResultRequestFilter) SetConnection(v []InterfaceHardwareType) {
	o.Connection = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetPlatform() []Platform {
	if o == nil || utils.IsNil(o.Platform) {
		var ret []Platform
		return ret
	}
	return o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetPlatformOk() ([]Platform, bool) {
	if o == nil || utils.IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasPlatform() bool {
	if o != nil && !utils.IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given []Platform and assigns it to the Platform field.
func (o *EndpointResultRequestFilter) SetPlatform(v []Platform) {
	o.Platform = v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetGateway() []string {
	if o == nil || utils.IsNil(o.Gateway) {
		var ret []string
		return ret
	}
	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetGatewayOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasGateway() bool {
	if o != nil && !utils.IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given []string and assigns it to the Gateway field.
func (o *EndpointResultRequestFilter) SetGateway(v []string) {
	o.Gateway = v
}

// GetProxyTarget returns the ProxyTarget field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetProxyTarget() []string {
	if o == nil || utils.IsNil(o.ProxyTarget) {
		var ret []string
		return ret
	}
	return o.ProxyTarget
}

// GetProxyTargetOk returns a tuple with the ProxyTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetProxyTargetOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.ProxyTarget) {
		return nil, false
	}
	return o.ProxyTarget, true
}

// HasProxyTarget returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasProxyTarget() bool {
	if o != nil && !utils.IsNil(o.ProxyTarget) {
		return true
	}

	return false
}

// SetProxyTarget gets a reference to the given []string and assigns it to the ProxyTarget field.
func (o *EndpointResultRequestFilter) SetProxyTarget(v []string) {
	o.ProxyTarget = v
}

// GetVpnTarget returns the VpnTarget field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetVpnTarget() []string {
	if o == nil || utils.IsNil(o.VpnTarget) {
		var ret []string
		return ret
	}
	return o.VpnTarget
}

// GetVpnTargetOk returns a tuple with the VpnTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetVpnTargetOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.VpnTarget) {
		return nil, false
	}
	return o.VpnTarget, true
}

// HasVpnTarget returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasVpnTarget() bool {
	if o != nil && !utils.IsNil(o.VpnTarget) {
		return true
	}

	return false
}

// SetVpnTarget gets a reference to the given []string and assigns it to the VpnTarget field.
func (o *EndpointResultRequestFilter) SetVpnTarget(v []string) {
	o.VpnTarget = v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetAgentId() []string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret []string
		return ret
	}
	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetAgentIdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given []string and assigns it to the AgentId field.
func (o *EndpointResultRequestFilter) SetAgentId(v []string) {
	o.AgentId = v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetNetworkId() []string {
	if o == nil || utils.IsNil(o.NetworkId) {
		var ret []string
		return ret
	}
	return o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetNetworkIdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasNetworkId() bool {
	if o != nil && !utils.IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given []string and assigns it to the NetworkId field.
func (o *EndpointResultRequestFilter) SetNetworkId(v []string) {
	o.NetworkId = v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetSsid() []string {
	if o == nil || utils.IsNil(o.Ssid) {
		var ret []string
		return ret
	}
	return o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetSsidOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasSsid() bool {
	if o != nil && !utils.IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given []string and assigns it to the Ssid field.
func (o *EndpointResultRequestFilter) SetSsid(v []string) {
	o.Ssid = v
}

// GetBssid returns the Bssid field value if set, zero value otherwise.
func (o *EndpointResultRequestFilter) GetBssid() []string {
	if o == nil || utils.IsNil(o.Bssid) {
		var ret []string
		return ret
	}
	return o.Bssid
}

// GetBssidOk returns a tuple with the Bssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointResultRequestFilter) GetBssidOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Bssid) {
		return nil, false
	}
	return o.Bssid, true
}

// HasBssid returns a boolean if a field has been set.
func (o *EndpointResultRequestFilter) HasBssid() bool {
	if o != nil && !utils.IsNil(o.Bssid) {
		return true
	}

	return false
}

// SetBssid gets a reference to the given []string and assigns it to the Bssid field.
func (o *EndpointResultRequestFilter) SetBssid(v []string) {
	o.Bssid = v
}

func (o EndpointResultRequestFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointResultRequestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !utils.IsNil(o.Connection) {
		toSerialize["connection"] = o.Connection
	}
	if !utils.IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !utils.IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !utils.IsNil(o.ProxyTarget) {
		toSerialize["proxyTarget"] = o.ProxyTarget
	}
	if !utils.IsNil(o.VpnTarget) {
		toSerialize["vpnTarget"] = o.VpnTarget
	}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.NetworkId) {
		toSerialize["networkId"] = o.NetworkId
	}
	if !utils.IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !utils.IsNil(o.Bssid) {
		toSerialize["bssid"] = o.Bssid
	}
	return toSerialize, nil
}

type NullableEndpointResultRequestFilter struct {
	value *EndpointResultRequestFilter
	isSet bool
}

func (v NullableEndpointResultRequestFilter) Get() *EndpointResultRequestFilter {
	return v.value
}

func (v *NullableEndpointResultRequestFilter) Set(val *EndpointResultRequestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointResultRequestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointResultRequestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointResultRequestFilter(val *EndpointResultRequestFilter) *NullableEndpointResultRequestFilter {
	return &NullableEndpointResultRequestFilter{value: val, isSet: true}
}

func (v NullableEndpointResultRequestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointResultRequestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


