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

// checks if the EndpointPathTrace type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointPathTrace{}

// EndpointPathTrace struct for EndpointPathTrace
type EndpointPathTrace struct {
	// IP address of the hop destination.
	IpAddress *string `json:"ipAddress,omitempty"`
	// Number of hops for path trace to destination.
	NumberOfHops *int32 `json:"numberOfHops,omitempty"`
	Protocol *EndpointTestResultProtocol `json:"protocol,omitempty"`
	TcpPathTraceMode *TcpPathTraceModeResponse `json:"tcpPathTraceMode,omitempty"`
	UdpPathTraceMode *UdpPathTraceModeResponse `json:"udpPathTraceMode,omitempty"`
	// Unique ID of path trace.
	PathId *string `json:"pathId,omitempty"`
	// RTT of the path trace to the destination in milliseconds.
	ResponseTime *int32 `json:"responseTime,omitempty"`
}

// NewEndpointPathTrace instantiates a new EndpointPathTrace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointPathTrace() *EndpointPathTrace {
	this := EndpointPathTrace{}
	var protocol EndpointTestResultProtocol = ENDPOINTTESTRESULTPROTOCOL_UNKNOWN
	this.Protocol = &protocol
	var tcpPathTraceMode TcpPathTraceModeResponse = TCPPATHTRACEMODERESPONSE_AUTO
	this.TcpPathTraceMode = &tcpPathTraceMode
	var udpPathTraceMode UdpPathTraceModeResponse = UDPPATHTRACEMODERESPONSE_UNKNOWN
	this.UdpPathTraceMode = &udpPathTraceMode
	return &this
}

// NewEndpointPathTraceWithDefaults instantiates a new EndpointPathTrace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointPathTraceWithDefaults() *EndpointPathTrace {
	this := EndpointPathTrace{}
	var protocol EndpointTestResultProtocol = ENDPOINTTESTRESULTPROTOCOL_UNKNOWN
	this.Protocol = &protocol
	var tcpPathTraceMode TcpPathTraceModeResponse = TCPPATHTRACEMODERESPONSE_AUTO
	this.TcpPathTraceMode = &tcpPathTraceMode
	var udpPathTraceMode UdpPathTraceModeResponse = UDPPATHTRACEMODERESPONSE_UNKNOWN
	this.UdpPathTraceMode = &udpPathTraceMode
	return &this
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *EndpointPathTrace) GetIpAddress() string {
	if o == nil || utils.IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPathTrace) GetIpAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *EndpointPathTrace) HasIpAddress() bool {
	if o != nil && !utils.IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *EndpointPathTrace) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetNumberOfHops returns the NumberOfHops field value if set, zero value otherwise.
func (o *EndpointPathTrace) GetNumberOfHops() int32 {
	if o == nil || utils.IsNil(o.NumberOfHops) {
		var ret int32
		return ret
	}
	return *o.NumberOfHops
}

// GetNumberOfHopsOk returns a tuple with the NumberOfHops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPathTrace) GetNumberOfHopsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumberOfHops) {
		return nil, false
	}
	return o.NumberOfHops, true
}

// HasNumberOfHops returns a boolean if a field has been set.
func (o *EndpointPathTrace) HasNumberOfHops() bool {
	if o != nil && !utils.IsNil(o.NumberOfHops) {
		return true
	}

	return false
}

// SetNumberOfHops gets a reference to the given int32 and assigns it to the NumberOfHops field.
func (o *EndpointPathTrace) SetNumberOfHops(v int32) {
	o.NumberOfHops = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *EndpointPathTrace) GetProtocol() EndpointTestResultProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret EndpointTestResultProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPathTrace) GetProtocolOk() (*EndpointTestResultProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *EndpointPathTrace) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given EndpointTestResultProtocol and assigns it to the Protocol field.
func (o *EndpointPathTrace) SetProtocol(v EndpointTestResultProtocol) {
	o.Protocol = &v
}

// GetTcpPathTraceMode returns the TcpPathTraceMode field value if set, zero value otherwise.
func (o *EndpointPathTrace) GetTcpPathTraceMode() TcpPathTraceModeResponse {
	if o == nil || utils.IsNil(o.TcpPathTraceMode) {
		var ret TcpPathTraceModeResponse
		return ret
	}
	return *o.TcpPathTraceMode
}

// GetTcpPathTraceModeOk returns a tuple with the TcpPathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPathTrace) GetTcpPathTraceModeOk() (*TcpPathTraceModeResponse, bool) {
	if o == nil || utils.IsNil(o.TcpPathTraceMode) {
		return nil, false
	}
	return o.TcpPathTraceMode, true
}

// HasTcpPathTraceMode returns a boolean if a field has been set.
func (o *EndpointPathTrace) HasTcpPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.TcpPathTraceMode) {
		return true
	}

	return false
}

// SetTcpPathTraceMode gets a reference to the given TcpPathTraceModeResponse and assigns it to the TcpPathTraceMode field.
func (o *EndpointPathTrace) SetTcpPathTraceMode(v TcpPathTraceModeResponse) {
	o.TcpPathTraceMode = &v
}

// GetUdpPathTraceMode returns the UdpPathTraceMode field value if set, zero value otherwise.
func (o *EndpointPathTrace) GetUdpPathTraceMode() UdpPathTraceModeResponse {
	if o == nil || utils.IsNil(o.UdpPathTraceMode) {
		var ret UdpPathTraceModeResponse
		return ret
	}
	return *o.UdpPathTraceMode
}

// GetUdpPathTraceModeOk returns a tuple with the UdpPathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPathTrace) GetUdpPathTraceModeOk() (*UdpPathTraceModeResponse, bool) {
	if o == nil || utils.IsNil(o.UdpPathTraceMode) {
		return nil, false
	}
	return o.UdpPathTraceMode, true
}

// HasUdpPathTraceMode returns a boolean if a field has been set.
func (o *EndpointPathTrace) HasUdpPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.UdpPathTraceMode) {
		return true
	}

	return false
}

// SetUdpPathTraceMode gets a reference to the given UdpPathTraceModeResponse and assigns it to the UdpPathTraceMode field.
func (o *EndpointPathTrace) SetUdpPathTraceMode(v UdpPathTraceModeResponse) {
	o.UdpPathTraceMode = &v
}

// GetPathId returns the PathId field value if set, zero value otherwise.
func (o *EndpointPathTrace) GetPathId() string {
	if o == nil || utils.IsNil(o.PathId) {
		var ret string
		return ret
	}
	return *o.PathId
}

// GetPathIdOk returns a tuple with the PathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPathTrace) GetPathIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PathId) {
		return nil, false
	}
	return o.PathId, true
}

// HasPathId returns a boolean if a field has been set.
func (o *EndpointPathTrace) HasPathId() bool {
	if o != nil && !utils.IsNil(o.PathId) {
		return true
	}

	return false
}

// SetPathId gets a reference to the given string and assigns it to the PathId field.
func (o *EndpointPathTrace) SetPathId(v string) {
	o.PathId = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *EndpointPathTrace) GetResponseTime() int32 {
	if o == nil || utils.IsNil(o.ResponseTime) {
		var ret int32
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointPathTrace) GetResponseTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *EndpointPathTrace) HasResponseTime() bool {
	if o != nil && !utils.IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given int32 and assigns it to the ResponseTime field.
func (o *EndpointPathTrace) SetResponseTime(v int32) {
	o.ResponseTime = &v
}

func (o EndpointPathTrace) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointPathTrace) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !utils.IsNil(o.NumberOfHops) {
		toSerialize["numberOfHops"] = o.NumberOfHops
	}
	if !utils.IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !utils.IsNil(o.TcpPathTraceMode) {
		toSerialize["tcpPathTraceMode"] = o.TcpPathTraceMode
	}
	if !utils.IsNil(o.UdpPathTraceMode) {
		toSerialize["udpPathTraceMode"] = o.UdpPathTraceMode
	}
	if !utils.IsNil(o.PathId) {
		toSerialize["pathId"] = o.PathId
	}
	if !utils.IsNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	return toSerialize, nil
}

type NullableEndpointPathTrace struct {
	value *EndpointPathTrace
	isSet bool
}

func (v NullableEndpointPathTrace) Get() *EndpointPathTrace {
	return v.value
}

func (v *NullableEndpointPathTrace) Set(val *EndpointPathTrace) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointPathTrace) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointPathTrace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointPathTrace(val *EndpointPathTrace) *NullableEndpointPathTrace {
	return &NullableEndpointPathTrace{value: val, isSet: true}
}

func (v NullableEndpointPathTrace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointPathTrace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


