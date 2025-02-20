/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the NetworkEndpointTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkEndpointTestResult{}

// NetworkEndpointTestResult struct for NetworkEndpointTestResult
type NetworkEndpointTestResult struct {
	// A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint.
	Aid *string `json:"aid,omitempty"`
	// Unique ID of endpoint test.
	TestId *string `json:"testId,omitempty"`
	// Unique ID of endpoint agent, from `/endpoint/agents` endpoint.
	AgentId *string `json:"agentId,omitempty"`
	// Epoch time (seconds) indicating the start time of the round.
	RoundId *int32 `json:"roundId,omitempty"`
	// IP address of target server.
	ServerIp *string `json:"serverIp,omitempty"`
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	SystemMetrics *SystemMetrics `json:"systemMetrics,omitempty"`
	OriginalTargetProfile *TargetProfile `json:"originalTargetProfile,omitempty"`
	VpnProfile *VpnProfile `json:"vpnProfile,omitempty"`
	// Average RTT for packets sent to destination.
	AvgLatency *float64 `json:"avgLatency,omitempty"`
	// Error details, if an error was encountered.
	ErrorDetails *string `json:"errorDetails,omitempty"`
	// Standard deviation of latency.
	Jitter *float64 `json:"jitter,omitempty"`
	// Set to `true` if network target is blocking ICMP echo (ping) queries.
	IsIcmpBlocked *bool `json:"isIcmpBlocked,omitempty"`
	// Percentage of packets not reaching destination.
	Loss *float64 `json:"loss,omitempty"`
	// Maximum RTT for packets sent to destination.
	MaxLatency *float64 `json:"maxLatency,omitempty"`
	// Minimum RTT for packets sent to destination.
	MinLatency *float64 `json:"minLatency,omitempty"`
}

// NewNetworkEndpointTestResult instantiates a new NetworkEndpointTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkEndpointTestResult() *NetworkEndpointTestResult {
	this := NetworkEndpointTestResult{}
	return &this
}

// NewNetworkEndpointTestResultWithDefaults instantiates a new NetworkEndpointTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkEndpointTestResultWithDefaults() *NetworkEndpointTestResult {
	this := NetworkEndpointTestResult{}
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *NetworkEndpointTestResult) SetAid(v string) {
	o.Aid = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *NetworkEndpointTestResult) SetTestId(v string) {
	o.TestId = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *NetworkEndpointTestResult) SetAgentId(v string) {
	o.AgentId = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *NetworkEndpointTestResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetServerIp() string {
	if o == nil || utils.IsNil(o.ServerIp) {
		var ret string
		return ret
	}
	return *o.ServerIp
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetServerIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerIp) {
		return nil, false
	}
	return o.ServerIp, true
}

// HasServerIp returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasServerIp() bool {
	if o != nil && !utils.IsNil(o.ServerIp) {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given string and assigns it to the ServerIp field.
func (o *NetworkEndpointTestResult) SetServerIp(v string) {
	o.ServerIp = &v
}

// GetNetworkProfile returns the NetworkProfile field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetNetworkProfile() NetworkProfile {
	if o == nil || utils.IsNil(o.NetworkProfile) {
		var ret NetworkProfile
		return ret
	}
	return *o.NetworkProfile
}

// GetNetworkProfileOk returns a tuple with the NetworkProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetNetworkProfileOk() (*NetworkProfile, bool) {
	if o == nil || utils.IsNil(o.NetworkProfile) {
		return nil, false
	}
	return o.NetworkProfile, true
}

// HasNetworkProfile returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasNetworkProfile() bool {
	if o != nil && !utils.IsNil(o.NetworkProfile) {
		return true
	}

	return false
}

// SetNetworkProfile gets a reference to the given NetworkProfile and assigns it to the NetworkProfile field.
func (o *NetworkEndpointTestResult) SetNetworkProfile(v NetworkProfile) {
	o.NetworkProfile = &v
}

// GetSystemMetrics returns the SystemMetrics field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetSystemMetrics() SystemMetrics {
	if o == nil || utils.IsNil(o.SystemMetrics) {
		var ret SystemMetrics
		return ret
	}
	return *o.SystemMetrics
}

// GetSystemMetricsOk returns a tuple with the SystemMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetSystemMetricsOk() (*SystemMetrics, bool) {
	if o == nil || utils.IsNil(o.SystemMetrics) {
		return nil, false
	}
	return o.SystemMetrics, true
}

// HasSystemMetrics returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasSystemMetrics() bool {
	if o != nil && !utils.IsNil(o.SystemMetrics) {
		return true
	}

	return false
}

// SetSystemMetrics gets a reference to the given SystemMetrics and assigns it to the SystemMetrics field.
func (o *NetworkEndpointTestResult) SetSystemMetrics(v SystemMetrics) {
	o.SystemMetrics = &v
}

// GetOriginalTargetProfile returns the OriginalTargetProfile field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetOriginalTargetProfile() TargetProfile {
	if o == nil || utils.IsNil(o.OriginalTargetProfile) {
		var ret TargetProfile
		return ret
	}
	return *o.OriginalTargetProfile
}

// GetOriginalTargetProfileOk returns a tuple with the OriginalTargetProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetOriginalTargetProfileOk() (*TargetProfile, bool) {
	if o == nil || utils.IsNil(o.OriginalTargetProfile) {
		return nil, false
	}
	return o.OriginalTargetProfile, true
}

// HasOriginalTargetProfile returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasOriginalTargetProfile() bool {
	if o != nil && !utils.IsNil(o.OriginalTargetProfile) {
		return true
	}

	return false
}

// SetOriginalTargetProfile gets a reference to the given TargetProfile and assigns it to the OriginalTargetProfile field.
func (o *NetworkEndpointTestResult) SetOriginalTargetProfile(v TargetProfile) {
	o.OriginalTargetProfile = &v
}

// GetVpnProfile returns the VpnProfile field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetVpnProfile() VpnProfile {
	if o == nil || utils.IsNil(o.VpnProfile) {
		var ret VpnProfile
		return ret
	}
	return *o.VpnProfile
}

// GetVpnProfileOk returns a tuple with the VpnProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetVpnProfileOk() (*VpnProfile, bool) {
	if o == nil || utils.IsNil(o.VpnProfile) {
		return nil, false
	}
	return o.VpnProfile, true
}

// HasVpnProfile returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasVpnProfile() bool {
	if o != nil && !utils.IsNil(o.VpnProfile) {
		return true
	}

	return false
}

// SetVpnProfile gets a reference to the given VpnProfile and assigns it to the VpnProfile field.
func (o *NetworkEndpointTestResult) SetVpnProfile(v VpnProfile) {
	o.VpnProfile = &v
}

// GetAvgLatency returns the AvgLatency field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetAvgLatency() float64 {
	if o == nil || utils.IsNil(o.AvgLatency) {
		var ret float64
		return ret
	}
	return *o.AvgLatency
}

// GetAvgLatencyOk returns a tuple with the AvgLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetAvgLatencyOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.AvgLatency) {
		return nil, false
	}
	return o.AvgLatency, true
}

// HasAvgLatency returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasAvgLatency() bool {
	if o != nil && !utils.IsNil(o.AvgLatency) {
		return true
	}

	return false
}

// SetAvgLatency gets a reference to the given float64 and assigns it to the AvgLatency field.
func (o *NetworkEndpointTestResult) SetAvgLatency(v float64) {
	o.AvgLatency = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetErrorDetails() string {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		var ret string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetErrorDetailsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasErrorDetails() bool {
	if o != nil && !utils.IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given string and assigns it to the ErrorDetails field.
func (o *NetworkEndpointTestResult) SetErrorDetails(v string) {
	o.ErrorDetails = &v
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetJitter() float64 {
	if o == nil || utils.IsNil(o.Jitter) {
		var ret float64
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetJitterOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Jitter) {
		return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasJitter() bool {
	if o != nil && !utils.IsNil(o.Jitter) {
		return true
	}

	return false
}

// SetJitter gets a reference to the given float64 and assigns it to the Jitter field.
func (o *NetworkEndpointTestResult) SetJitter(v float64) {
	o.Jitter = &v
}

// GetIsIcmpBlocked returns the IsIcmpBlocked field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetIsIcmpBlocked() bool {
	if o == nil || utils.IsNil(o.IsIcmpBlocked) {
		var ret bool
		return ret
	}
	return *o.IsIcmpBlocked
}

// GetIsIcmpBlockedOk returns a tuple with the IsIcmpBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetIsIcmpBlockedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsIcmpBlocked) {
		return nil, false
	}
	return o.IsIcmpBlocked, true
}

// HasIsIcmpBlocked returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasIsIcmpBlocked() bool {
	if o != nil && !utils.IsNil(o.IsIcmpBlocked) {
		return true
	}

	return false
}

// SetIsIcmpBlocked gets a reference to the given bool and assigns it to the IsIcmpBlocked field.
func (o *NetworkEndpointTestResult) SetIsIcmpBlocked(v bool) {
	o.IsIcmpBlocked = &v
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetLoss() float64 {
	if o == nil || utils.IsNil(o.Loss) {
		var ret float64
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetLossOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Loss) {
		return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasLoss() bool {
	if o != nil && !utils.IsNil(o.Loss) {
		return true
	}

	return false
}

// SetLoss gets a reference to the given float64 and assigns it to the Loss field.
func (o *NetworkEndpointTestResult) SetLoss(v float64) {
	o.Loss = &v
}

// GetMaxLatency returns the MaxLatency field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetMaxLatency() float64 {
	if o == nil || utils.IsNil(o.MaxLatency) {
		var ret float64
		return ret
	}
	return *o.MaxLatency
}

// GetMaxLatencyOk returns a tuple with the MaxLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetMaxLatencyOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.MaxLatency) {
		return nil, false
	}
	return o.MaxLatency, true
}

// HasMaxLatency returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasMaxLatency() bool {
	if o != nil && !utils.IsNil(o.MaxLatency) {
		return true
	}

	return false
}

// SetMaxLatency gets a reference to the given float64 and assigns it to the MaxLatency field.
func (o *NetworkEndpointTestResult) SetMaxLatency(v float64) {
	o.MaxLatency = &v
}

// GetMinLatency returns the MinLatency field value if set, zero value otherwise.
func (o *NetworkEndpointTestResult) GetMinLatency() float64 {
	if o == nil || utils.IsNil(o.MinLatency) {
		var ret float64
		return ret
	}
	return *o.MinLatency
}

// GetMinLatencyOk returns a tuple with the MinLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkEndpointTestResult) GetMinLatencyOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.MinLatency) {
		return nil, false
	}
	return o.MinLatency, true
}

// HasMinLatency returns a boolean if a field has been set.
func (o *NetworkEndpointTestResult) HasMinLatency() bool {
	if o != nil && !utils.IsNil(o.MinLatency) {
		return true
	}

	return false
}

// SetMinLatency gets a reference to the given float64 and assigns it to the MinLatency field.
func (o *NetworkEndpointTestResult) SetMinLatency(v float64) {
	o.MinLatency = &v
}

func (o NetworkEndpointTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkEndpointTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.RoundId) {
		toSerialize["roundId"] = o.RoundId
	}
	if !utils.IsNil(o.ServerIp) {
		toSerialize["serverIp"] = o.ServerIp
	}
	if !utils.IsNil(o.NetworkProfile) {
		toSerialize["networkProfile"] = o.NetworkProfile
	}
	if !utils.IsNil(o.SystemMetrics) {
		toSerialize["systemMetrics"] = o.SystemMetrics
	}
	if !utils.IsNil(o.OriginalTargetProfile) {
		toSerialize["originalTargetProfile"] = o.OriginalTargetProfile
	}
	if !utils.IsNil(o.VpnProfile) {
		toSerialize["vpnProfile"] = o.VpnProfile
	}
	if !utils.IsNil(o.AvgLatency) {
		toSerialize["avgLatency"] = o.AvgLatency
	}
	if !utils.IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !utils.IsNil(o.Jitter) {
		toSerialize["jitter"] = o.Jitter
	}
	if !utils.IsNil(o.IsIcmpBlocked) {
		toSerialize["isIcmpBlocked"] = o.IsIcmpBlocked
	}
	if !utils.IsNil(o.Loss) {
		toSerialize["loss"] = o.Loss
	}
	if !utils.IsNil(o.MaxLatency) {
		toSerialize["maxLatency"] = o.MaxLatency
	}
	if !utils.IsNil(o.MinLatency) {
		toSerialize["minLatency"] = o.MinLatency
	}
	return toSerialize, nil
}

type NullableNetworkEndpointTestResult struct {
	value *NetworkEndpointTestResult
	isSet bool
}

func (v NullableNetworkEndpointTestResult) Get() *NetworkEndpointTestResult {
	return v.value
}

func (v *NullableNetworkEndpointTestResult) Set(val *NetworkEndpointTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkEndpointTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkEndpointTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkEndpointTestResult(val *NetworkEndpointTestResult) *NullableNetworkEndpointTestResult {
	return &NullableNetworkEndpointTestResult{value: val, isSet: true}
}

func (v NullableNetworkEndpointTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkEndpointTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


