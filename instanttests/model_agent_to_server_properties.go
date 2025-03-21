/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the AgentToServerProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentToServerProperties{}

// AgentToServerProperties struct for AgentToServerProperties
type AgentToServerProperties struct {
	// Set to `true` to enable bandwidth measurements, only applies to Enterprise agents assigned to the test.
	BandwidthMeasurements *bool `json:"bandwidthMeasurements,omitempty"`
	// To enable continuous monitoring, set this parameter to `true` to.  When continuous monitoring is enabled, the following actions occur: * `fixedPacketRate` is enforced * `bandwidthMeasurements` are disabled * If the `protocol` is set to `tcp`, `probeMode` is set to `syn`. 
	ContinuousMode *bool `json:"continuousMode,omitempty"`
	// If continuousMode is `false`, set the fixedPacketRate to a value between 10-100. If `continuousMode` is `true`, set the `fixedPacketRate` to `1`
	FixedPacketRate *int32 `json:"fixedPacketRate,omitempty"`
	// Set `true` to measure MTU sizes on network from agents to the target.
	MtuMeasurements *bool `json:"mtuMeasurements,omitempty"`
	// Number of path traces executed by the agent.
	NumPathTraces *int32 `json:"numPathTraces,omitempty"`
	PathTraceMode *TestPathTraceMode `json:"pathTraceMode,omitempty"`
	// Target port.
	Port *int32 `json:"port,omitempty"`
	ProbeMode *TestProbeMode `json:"probeMode,omitempty"`
	Protocol *TestProtocol `json:"protocol,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	// Target name or IP address.
	Server string `json:"server"`
	// DSCP label.
	Dscp *string `json:"dscp,omitempty"`
	DscpId *TestDscpId `json:"dscpId,omitempty"`
	Ipv6Policy *TestIpv6Policy `json:"ipv6Policy,omitempty"`
	// Payload size (not total packet size) for the end-to-end metric's probes, ping payload size allows values from 0 to 1400 bytes. When set to null, payload sizes are 0 bytes for ICMP-based tests and 1 byte for TCP-based tests.
	PingPayloadSize *int32 `json:"pingPayloadSize,omitempty"`
	// View packet loss in 1-second intervals. This is only available for 1-minute interval tests. Set to `true` to enable network measurements.
	NetworkMeasurements *bool `json:"networkMeasurements,omitempty"`
	Type *string `json:"type,omitempty"`
}

type _AgentToServerProperties AgentToServerProperties

// NewAgentToServerProperties instantiates a new AgentToServerProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToServerProperties(server string) *AgentToServerProperties {
	this := AgentToServerProperties{}
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var port int32 = 49153
	this.Port = &port
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	this.Server = server
	var dscpId TestDscpId = TESTDSCPID__0
	this.DscpId = &dscpId
	var ipv6Policy TestIpv6Policy = TESTIPV6POLICY_USE_AGENT_POLICY
	this.Ipv6Policy = &ipv6Policy
	var networkMeasurements bool = false
	this.NetworkMeasurements = &networkMeasurements
	return &this
}

// NewAgentToServerPropertiesWithDefaults instantiates a new AgentToServerProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToServerPropertiesWithDefaults() *AgentToServerProperties {
	this := AgentToServerProperties{}
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var port int32 = 49153
	this.Port = &port
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	var dscpId TestDscpId = TESTDSCPID__0
	this.DscpId = &dscpId
	var ipv6Policy TestIpv6Policy = TESTIPV6POLICY_USE_AGENT_POLICY
	this.Ipv6Policy = &ipv6Policy
	var networkMeasurements bool = false
	this.NetworkMeasurements = &networkMeasurements
	return &this
}

// GetBandwidthMeasurements returns the BandwidthMeasurements field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetBandwidthMeasurements() bool {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		var ret bool
		return ret
	}
	return *o.BandwidthMeasurements
}

// GetBandwidthMeasurementsOk returns a tuple with the BandwidthMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetBandwidthMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		return nil, false
	}
	return o.BandwidthMeasurements, true
}

// HasBandwidthMeasurements returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasBandwidthMeasurements() bool {
	if o != nil && !utils.IsNil(o.BandwidthMeasurements) {
		return true
	}

	return false
}

// SetBandwidthMeasurements gets a reference to the given bool and assigns it to the BandwidthMeasurements field.
func (o *AgentToServerProperties) SetBandwidthMeasurements(v bool) {
	o.BandwidthMeasurements = &v
}

// GetContinuousMode returns the ContinuousMode field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetContinuousMode() bool {
	if o == nil || utils.IsNil(o.ContinuousMode) {
		var ret bool
		return ret
	}
	return *o.ContinuousMode
}

// GetContinuousModeOk returns a tuple with the ContinuousMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetContinuousModeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ContinuousMode) {
		return nil, false
	}
	return o.ContinuousMode, true
}

// HasContinuousMode returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasContinuousMode() bool {
	if o != nil && !utils.IsNil(o.ContinuousMode) {
		return true
	}

	return false
}

// SetContinuousMode gets a reference to the given bool and assigns it to the ContinuousMode field.
func (o *AgentToServerProperties) SetContinuousMode(v bool) {
	o.ContinuousMode = &v
}

// GetFixedPacketRate returns the FixedPacketRate field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetFixedPacketRate() int32 {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		var ret int32
		return ret
	}
	return *o.FixedPacketRate
}

// GetFixedPacketRateOk returns a tuple with the FixedPacketRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetFixedPacketRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		return nil, false
	}
	return o.FixedPacketRate, true
}

// HasFixedPacketRate returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasFixedPacketRate() bool {
	if o != nil && !utils.IsNil(o.FixedPacketRate) {
		return true
	}

	return false
}

// SetFixedPacketRate gets a reference to the given int32 and assigns it to the FixedPacketRate field.
func (o *AgentToServerProperties) SetFixedPacketRate(v int32) {
	o.FixedPacketRate = &v
}

// GetMtuMeasurements returns the MtuMeasurements field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetMtuMeasurements() bool {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		var ret bool
		return ret
	}
	return *o.MtuMeasurements
}

// GetMtuMeasurementsOk returns a tuple with the MtuMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetMtuMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		return nil, false
	}
	return o.MtuMeasurements, true
}

// HasMtuMeasurements returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasMtuMeasurements() bool {
	if o != nil && !utils.IsNil(o.MtuMeasurements) {
		return true
	}

	return false
}

// SetMtuMeasurements gets a reference to the given bool and assigns it to the MtuMeasurements field.
func (o *AgentToServerProperties) SetMtuMeasurements(v bool) {
	o.MtuMeasurements = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *AgentToServerProperties) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPathTraceMode returns the PathTraceMode field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetPathTraceMode() TestPathTraceMode {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		var ret TestPathTraceMode
		return ret
	}
	return *o.PathTraceMode
}

// GetPathTraceModeOk returns a tuple with the PathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetPathTraceModeOk() (*TestPathTraceMode, bool) {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		return nil, false
	}
	return o.PathTraceMode, true
}

// HasPathTraceMode returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.PathTraceMode) {
		return true
	}

	return false
}

// SetPathTraceMode gets a reference to the given TestPathTraceMode and assigns it to the PathTraceMode field.
func (o *AgentToServerProperties) SetPathTraceMode(v TestPathTraceMode) {
	o.PathTraceMode = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *AgentToServerProperties) SetPort(v int32) {
	o.Port = &v
}

// GetProbeMode returns the ProbeMode field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetProbeMode() TestProbeMode {
	if o == nil || utils.IsNil(o.ProbeMode) {
		var ret TestProbeMode
		return ret
	}
	return *o.ProbeMode
}

// GetProbeModeOk returns a tuple with the ProbeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetProbeModeOk() (*TestProbeMode, bool) {
	if o == nil || utils.IsNil(o.ProbeMode) {
		return nil, false
	}
	return o.ProbeMode, true
}

// HasProbeMode returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasProbeMode() bool {
	if o != nil && !utils.IsNil(o.ProbeMode) {
		return true
	}

	return false
}

// SetProbeMode gets a reference to the given TestProbeMode and assigns it to the ProbeMode field.
func (o *AgentToServerProperties) SetProbeMode(v TestProbeMode) {
	o.ProbeMode = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetProtocol() TestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret TestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetProtocolOk() (*TestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given TestProtocol and assigns it to the Protocol field.
func (o *AgentToServerProperties) SetProtocol(v TestProtocol) {
	o.Protocol = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *AgentToServerProperties) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetServer returns the Server field value
func (o *AgentToServerProperties) GetServer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Server
}

// GetServerOk returns a tuple with the Server field value
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Server, true
}

// SetServer sets field value
func (o *AgentToServerProperties) SetServer(v string) {
	o.Server = v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetDscp() string {
	if o == nil || utils.IsNil(o.Dscp) {
		var ret string
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetDscpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasDscp() bool {
	if o != nil && !utils.IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given string and assigns it to the Dscp field.
func (o *AgentToServerProperties) SetDscp(v string) {
	o.Dscp = &v
}

// GetDscpId returns the DscpId field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetDscpId() TestDscpId {
	if o == nil || utils.IsNil(o.DscpId) {
		var ret TestDscpId
		return ret
	}
	return *o.DscpId
}

// GetDscpIdOk returns a tuple with the DscpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetDscpIdOk() (*TestDscpId, bool) {
	if o == nil || utils.IsNil(o.DscpId) {
		return nil, false
	}
	return o.DscpId, true
}

// HasDscpId returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasDscpId() bool {
	if o != nil && !utils.IsNil(o.DscpId) {
		return true
	}

	return false
}

// SetDscpId gets a reference to the given TestDscpId and assigns it to the DscpId field.
func (o *AgentToServerProperties) SetDscpId(v TestDscpId) {
	o.DscpId = &v
}

// GetIpv6Policy returns the Ipv6Policy field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetIpv6Policy() TestIpv6Policy {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		var ret TestIpv6Policy
		return ret
	}
	return *o.Ipv6Policy
}

// GetIpv6PolicyOk returns a tuple with the Ipv6Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetIpv6PolicyOk() (*TestIpv6Policy, bool) {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		return nil, false
	}
	return o.Ipv6Policy, true
}

// HasIpv6Policy returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasIpv6Policy() bool {
	if o != nil && !utils.IsNil(o.Ipv6Policy) {
		return true
	}

	return false
}

// SetIpv6Policy gets a reference to the given TestIpv6Policy and assigns it to the Ipv6Policy field.
func (o *AgentToServerProperties) SetIpv6Policy(v TestIpv6Policy) {
	o.Ipv6Policy = &v
}

// GetPingPayloadSize returns the PingPayloadSize field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetPingPayloadSize() int32 {
	if o == nil || utils.IsNil(o.PingPayloadSize) {
		var ret int32
		return ret
	}
	return *o.PingPayloadSize
}

// GetPingPayloadSizeOk returns a tuple with the PingPayloadSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetPingPayloadSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.PingPayloadSize) {
		return nil, false
	}
	return o.PingPayloadSize, true
}

// HasPingPayloadSize returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasPingPayloadSize() bool {
	if o != nil && !utils.IsNil(o.PingPayloadSize) {
		return true
	}

	return false
}

// SetPingPayloadSize gets a reference to the given int32 and assigns it to the PingPayloadSize field.
func (o *AgentToServerProperties) SetPingPayloadSize(v int32) {
	o.PingPayloadSize = &v
}

// GetNetworkMeasurements returns the NetworkMeasurements field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetNetworkMeasurements() bool {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		var ret bool
		return ret
	}
	return *o.NetworkMeasurements
}

// GetNetworkMeasurementsOk returns a tuple with the NetworkMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetNetworkMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		return nil, false
	}
	return o.NetworkMeasurements, true
}

// HasNetworkMeasurements returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasNetworkMeasurements() bool {
	if o != nil && !utils.IsNil(o.NetworkMeasurements) {
		return true
	}

	return false
}

// SetNetworkMeasurements gets a reference to the given bool and assigns it to the NetworkMeasurements field.
func (o *AgentToServerProperties) SetNetworkMeasurements(v bool) {
	o.NetworkMeasurements = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AgentToServerProperties) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToServerProperties) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AgentToServerProperties) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AgentToServerProperties) SetType(v string) {
	o.Type = &v
}

func (o AgentToServerProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToServerProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BandwidthMeasurements) {
		toSerialize["bandwidthMeasurements"] = o.BandwidthMeasurements
	}
	if !utils.IsNil(o.ContinuousMode) {
		toSerialize["continuousMode"] = o.ContinuousMode
	}
	if !utils.IsNil(o.FixedPacketRate) {
		toSerialize["fixedPacketRate"] = o.FixedPacketRate
	}
	if !utils.IsNil(o.MtuMeasurements) {
		toSerialize["mtuMeasurements"] = o.MtuMeasurements
	}
	if !utils.IsNil(o.NumPathTraces) {
		toSerialize["numPathTraces"] = o.NumPathTraces
	}
	if !utils.IsNil(o.PathTraceMode) {
		toSerialize["pathTraceMode"] = o.PathTraceMode
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.ProbeMode) {
		toSerialize["probeMode"] = o.ProbeMode
	}
	if !utils.IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !utils.IsNil(o.RandomizedStartTime) {
		toSerialize["randomizedStartTime"] = o.RandomizedStartTime
	}
	toSerialize["server"] = o.Server
	if !utils.IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !utils.IsNil(o.DscpId) {
		toSerialize["dscpId"] = o.DscpId
	}
	if !utils.IsNil(o.Ipv6Policy) {
		toSerialize["ipv6Policy"] = o.Ipv6Policy
	}
	if !utils.IsNil(o.PingPayloadSize) {
		toSerialize["pingPayloadSize"] = o.PingPayloadSize
	}
	if !utils.IsNil(o.NetworkMeasurements) {
		toSerialize["networkMeasurements"] = o.NetworkMeasurements
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *AgentToServerProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"server",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAgentToServerProperties := _AgentToServerProperties{}

    err = json.Unmarshal(data, &varAgentToServerProperties)

	if err != nil {
		return err
	}

	*o = AgentToServerProperties(varAgentToServerProperties)

	return err
}

type NullableAgentToServerProperties struct {
	value *AgentToServerProperties
	isSet bool
}

func (v NullableAgentToServerProperties) Get() *AgentToServerProperties {
	return v.value
}

func (v *NullableAgentToServerProperties) Set(val *AgentToServerProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToServerProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToServerProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToServerProperties(val *AgentToServerProperties) *NullableAgentToServerProperties {
	return &NullableAgentToServerProperties{value: val, isSet: true}
}

func (v NullableAgentToServerProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToServerProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


