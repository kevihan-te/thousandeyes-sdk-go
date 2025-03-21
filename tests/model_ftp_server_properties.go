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
	"fmt"
)

// checks if the FtpServerProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FtpServerProperties{}

// FtpServerProperties struct for FtpServerProperties
type FtpServerProperties struct {
	// Set to `true` to enable bandwidth measurements, only applies to Enterprise agents assigned to the test.
	BandwidthMeasurements *bool `json:"bandwidthMeasurements,omitempty"`
	// Specify maximum number of bytes to download from the target object.
	DownloadLimit *int32 `json:"downloadLimit,omitempty"`
	// Target time for operation completion; specified in milliseconds.
	FtpTargetTime *int32 `json:"ftpTargetTime,omitempty"`
	// Set the time limit for the test in seconds.
	FtpTimeLimit *int32 `json:"ftpTimeLimit,omitempty"`
	// Set `true` to measure MTU sizes on network from agents to the target.
	MtuMeasurements *bool `json:"mtuMeasurements,omitempty"`
	// Enable or disable network measurements. Set to true to enable or false to disable network measurements.
	NetworkMeasurements *bool `json:"networkMeasurements,omitempty"`
	// Number of path traces executed by the agent.
	NumPathTraces *int32 `json:"numPathTraces,omitempty"`
	// Password for Basic/NTLM authentication.
	Password string `json:"password"`
	PathTraceMode *TestPathTraceMode `json:"pathTraceMode,omitempty"`
	ProbeMode *TestProbeMode `json:"probeMode,omitempty"`
	Protocol *TestProtocol `json:"protocol,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	RequestType FtpServerRequestType `json:"requestType"`
	// Target for the test.
	Url string `json:"url"`
	// Explicitly set the flag to use active FTP.
	UseActiveFtp *bool `json:"useActiveFtp,omitempty"`
	// Use explicit FTPS (ftp over SSL). By default, tests will autodetect when it is appropriate to use FTPS.
	UseExplicitFtps *bool `json:"useExplicitFtps,omitempty"`
	// Username for Basic/NTLM authentication.
	Username string `json:"username"`
	// Sets packets rate sent to measure the network in packets per second.
	FixedPacketRate *int32 `json:"fixedPacketRate,omitempty"`
	Ipv6Policy *TestIpv6Policy `json:"ipv6Policy,omitempty"`
	Type *string `json:"type,omitempty"`
}

type _FtpServerProperties FtpServerProperties

// NewFtpServerProperties instantiates a new FtpServerProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFtpServerProperties(password string, requestType FtpServerRequestType, url string, username string) *FtpServerProperties {
	this := FtpServerProperties{}
	var ftpTimeLimit int32 = 10
	this.FtpTimeLimit = &ftpTimeLimit
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	this.Password = password
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	this.RequestType = requestType
	this.Url = url
	var useActiveFtp bool = false
	this.UseActiveFtp = &useActiveFtp
	this.Username = username
	var ipv6Policy TestIpv6Policy = TESTIPV6POLICY_USE_AGENT_POLICY
	this.Ipv6Policy = &ipv6Policy
	return &this
}

// NewFtpServerPropertiesWithDefaults instantiates a new FtpServerProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFtpServerPropertiesWithDefaults() *FtpServerProperties {
	this := FtpServerProperties{}
	var ftpTimeLimit int32 = 10
	this.FtpTimeLimit = &ftpTimeLimit
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	var useActiveFtp bool = false
	this.UseActiveFtp = &useActiveFtp
	var ipv6Policy TestIpv6Policy = TESTIPV6POLICY_USE_AGENT_POLICY
	this.Ipv6Policy = &ipv6Policy
	return &this
}

// GetBandwidthMeasurements returns the BandwidthMeasurements field value if set, zero value otherwise.
func (o *FtpServerProperties) GetBandwidthMeasurements() bool {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		var ret bool
		return ret
	}
	return *o.BandwidthMeasurements
}

// GetBandwidthMeasurementsOk returns a tuple with the BandwidthMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetBandwidthMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		return nil, false
	}
	return o.BandwidthMeasurements, true
}

// HasBandwidthMeasurements returns a boolean if a field has been set.
func (o *FtpServerProperties) HasBandwidthMeasurements() bool {
	if o != nil && !utils.IsNil(o.BandwidthMeasurements) {
		return true
	}

	return false
}

// SetBandwidthMeasurements gets a reference to the given bool and assigns it to the BandwidthMeasurements field.
func (o *FtpServerProperties) SetBandwidthMeasurements(v bool) {
	o.BandwidthMeasurements = &v
}

// GetDownloadLimit returns the DownloadLimit field value if set, zero value otherwise.
func (o *FtpServerProperties) GetDownloadLimit() int32 {
	if o == nil || utils.IsNil(o.DownloadLimit) {
		var ret int32
		return ret
	}
	return *o.DownloadLimit
}

// GetDownloadLimitOk returns a tuple with the DownloadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetDownloadLimitOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.DownloadLimit) {
		return nil, false
	}
	return o.DownloadLimit, true
}

// HasDownloadLimit returns a boolean if a field has been set.
func (o *FtpServerProperties) HasDownloadLimit() bool {
	if o != nil && !utils.IsNil(o.DownloadLimit) {
		return true
	}

	return false
}

// SetDownloadLimit gets a reference to the given int32 and assigns it to the DownloadLimit field.
func (o *FtpServerProperties) SetDownloadLimit(v int32) {
	o.DownloadLimit = &v
}

// GetFtpTargetTime returns the FtpTargetTime field value if set, zero value otherwise.
func (o *FtpServerProperties) GetFtpTargetTime() int32 {
	if o == nil || utils.IsNil(o.FtpTargetTime) {
		var ret int32
		return ret
	}
	return *o.FtpTargetTime
}

// GetFtpTargetTimeOk returns a tuple with the FtpTargetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetFtpTargetTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FtpTargetTime) {
		return nil, false
	}
	return o.FtpTargetTime, true
}

// HasFtpTargetTime returns a boolean if a field has been set.
func (o *FtpServerProperties) HasFtpTargetTime() bool {
	if o != nil && !utils.IsNil(o.FtpTargetTime) {
		return true
	}

	return false
}

// SetFtpTargetTime gets a reference to the given int32 and assigns it to the FtpTargetTime field.
func (o *FtpServerProperties) SetFtpTargetTime(v int32) {
	o.FtpTargetTime = &v
}

// GetFtpTimeLimit returns the FtpTimeLimit field value if set, zero value otherwise.
func (o *FtpServerProperties) GetFtpTimeLimit() int32 {
	if o == nil || utils.IsNil(o.FtpTimeLimit) {
		var ret int32
		return ret
	}
	return *o.FtpTimeLimit
}

// GetFtpTimeLimitOk returns a tuple with the FtpTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetFtpTimeLimitOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FtpTimeLimit) {
		return nil, false
	}
	return o.FtpTimeLimit, true
}

// HasFtpTimeLimit returns a boolean if a field has been set.
func (o *FtpServerProperties) HasFtpTimeLimit() bool {
	if o != nil && !utils.IsNil(o.FtpTimeLimit) {
		return true
	}

	return false
}

// SetFtpTimeLimit gets a reference to the given int32 and assigns it to the FtpTimeLimit field.
func (o *FtpServerProperties) SetFtpTimeLimit(v int32) {
	o.FtpTimeLimit = &v
}

// GetMtuMeasurements returns the MtuMeasurements field value if set, zero value otherwise.
func (o *FtpServerProperties) GetMtuMeasurements() bool {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		var ret bool
		return ret
	}
	return *o.MtuMeasurements
}

// GetMtuMeasurementsOk returns a tuple with the MtuMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetMtuMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		return nil, false
	}
	return o.MtuMeasurements, true
}

// HasMtuMeasurements returns a boolean if a field has been set.
func (o *FtpServerProperties) HasMtuMeasurements() bool {
	if o != nil && !utils.IsNil(o.MtuMeasurements) {
		return true
	}

	return false
}

// SetMtuMeasurements gets a reference to the given bool and assigns it to the MtuMeasurements field.
func (o *FtpServerProperties) SetMtuMeasurements(v bool) {
	o.MtuMeasurements = &v
}

// GetNetworkMeasurements returns the NetworkMeasurements field value if set, zero value otherwise.
func (o *FtpServerProperties) GetNetworkMeasurements() bool {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		var ret bool
		return ret
	}
	return *o.NetworkMeasurements
}

// GetNetworkMeasurementsOk returns a tuple with the NetworkMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetNetworkMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		return nil, false
	}
	return o.NetworkMeasurements, true
}

// HasNetworkMeasurements returns a boolean if a field has been set.
func (o *FtpServerProperties) HasNetworkMeasurements() bool {
	if o != nil && !utils.IsNil(o.NetworkMeasurements) {
		return true
	}

	return false
}

// SetNetworkMeasurements gets a reference to the given bool and assigns it to the NetworkMeasurements field.
func (o *FtpServerProperties) SetNetworkMeasurements(v bool) {
	o.NetworkMeasurements = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *FtpServerProperties) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *FtpServerProperties) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *FtpServerProperties) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPassword returns the Password field value
func (o *FtpServerProperties) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *FtpServerProperties) SetPassword(v string) {
	o.Password = v
}

// GetPathTraceMode returns the PathTraceMode field value if set, zero value otherwise.
func (o *FtpServerProperties) GetPathTraceMode() TestPathTraceMode {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		var ret TestPathTraceMode
		return ret
	}
	return *o.PathTraceMode
}

// GetPathTraceModeOk returns a tuple with the PathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetPathTraceModeOk() (*TestPathTraceMode, bool) {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		return nil, false
	}
	return o.PathTraceMode, true
}

// HasPathTraceMode returns a boolean if a field has been set.
func (o *FtpServerProperties) HasPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.PathTraceMode) {
		return true
	}

	return false
}

// SetPathTraceMode gets a reference to the given TestPathTraceMode and assigns it to the PathTraceMode field.
func (o *FtpServerProperties) SetPathTraceMode(v TestPathTraceMode) {
	o.PathTraceMode = &v
}

// GetProbeMode returns the ProbeMode field value if set, zero value otherwise.
func (o *FtpServerProperties) GetProbeMode() TestProbeMode {
	if o == nil || utils.IsNil(o.ProbeMode) {
		var ret TestProbeMode
		return ret
	}
	return *o.ProbeMode
}

// GetProbeModeOk returns a tuple with the ProbeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetProbeModeOk() (*TestProbeMode, bool) {
	if o == nil || utils.IsNil(o.ProbeMode) {
		return nil, false
	}
	return o.ProbeMode, true
}

// HasProbeMode returns a boolean if a field has been set.
func (o *FtpServerProperties) HasProbeMode() bool {
	if o != nil && !utils.IsNil(o.ProbeMode) {
		return true
	}

	return false
}

// SetProbeMode gets a reference to the given TestProbeMode and assigns it to the ProbeMode field.
func (o *FtpServerProperties) SetProbeMode(v TestProbeMode) {
	o.ProbeMode = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *FtpServerProperties) GetProtocol() TestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret TestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetProtocolOk() (*TestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FtpServerProperties) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given TestProtocol and assigns it to the Protocol field.
func (o *FtpServerProperties) SetProtocol(v TestProtocol) {
	o.Protocol = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *FtpServerProperties) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *FtpServerProperties) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *FtpServerProperties) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetRequestType returns the RequestType field value
func (o *FtpServerProperties) GetRequestType() FtpServerRequestType {
	if o == nil {
		var ret FtpServerRequestType
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetRequestTypeOk() (*FtpServerRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *FtpServerProperties) SetRequestType(v FtpServerRequestType) {
	o.RequestType = v
}

// GetUrl returns the Url field value
func (o *FtpServerProperties) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FtpServerProperties) SetUrl(v string) {
	o.Url = v
}

// GetUseActiveFtp returns the UseActiveFtp field value if set, zero value otherwise.
func (o *FtpServerProperties) GetUseActiveFtp() bool {
	if o == nil || utils.IsNil(o.UseActiveFtp) {
		var ret bool
		return ret
	}
	return *o.UseActiveFtp
}

// GetUseActiveFtpOk returns a tuple with the UseActiveFtp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetUseActiveFtpOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UseActiveFtp) {
		return nil, false
	}
	return o.UseActiveFtp, true
}

// HasUseActiveFtp returns a boolean if a field has been set.
func (o *FtpServerProperties) HasUseActiveFtp() bool {
	if o != nil && !utils.IsNil(o.UseActiveFtp) {
		return true
	}

	return false
}

// SetUseActiveFtp gets a reference to the given bool and assigns it to the UseActiveFtp field.
func (o *FtpServerProperties) SetUseActiveFtp(v bool) {
	o.UseActiveFtp = &v
}

// GetUseExplicitFtps returns the UseExplicitFtps field value if set, zero value otherwise.
func (o *FtpServerProperties) GetUseExplicitFtps() bool {
	if o == nil || utils.IsNil(o.UseExplicitFtps) {
		var ret bool
		return ret
	}
	return *o.UseExplicitFtps
}

// GetUseExplicitFtpsOk returns a tuple with the UseExplicitFtps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetUseExplicitFtpsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UseExplicitFtps) {
		return nil, false
	}
	return o.UseExplicitFtps, true
}

// HasUseExplicitFtps returns a boolean if a field has been set.
func (o *FtpServerProperties) HasUseExplicitFtps() bool {
	if o != nil && !utils.IsNil(o.UseExplicitFtps) {
		return true
	}

	return false
}

// SetUseExplicitFtps gets a reference to the given bool and assigns it to the UseExplicitFtps field.
func (o *FtpServerProperties) SetUseExplicitFtps(v bool) {
	o.UseExplicitFtps = &v
}

// GetUsername returns the Username field value
func (o *FtpServerProperties) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *FtpServerProperties) SetUsername(v string) {
	o.Username = v
}

// GetFixedPacketRate returns the FixedPacketRate field value if set, zero value otherwise.
func (o *FtpServerProperties) GetFixedPacketRate() int32 {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		var ret int32
		return ret
	}
	return *o.FixedPacketRate
}

// GetFixedPacketRateOk returns a tuple with the FixedPacketRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetFixedPacketRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		return nil, false
	}
	return o.FixedPacketRate, true
}

// HasFixedPacketRate returns a boolean if a field has been set.
func (o *FtpServerProperties) HasFixedPacketRate() bool {
	if o != nil && !utils.IsNil(o.FixedPacketRate) {
		return true
	}

	return false
}

// SetFixedPacketRate gets a reference to the given int32 and assigns it to the FixedPacketRate field.
func (o *FtpServerProperties) SetFixedPacketRate(v int32) {
	o.FixedPacketRate = &v
}

// GetIpv6Policy returns the Ipv6Policy field value if set, zero value otherwise.
func (o *FtpServerProperties) GetIpv6Policy() TestIpv6Policy {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		var ret TestIpv6Policy
		return ret
	}
	return *o.Ipv6Policy
}

// GetIpv6PolicyOk returns a tuple with the Ipv6Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetIpv6PolicyOk() (*TestIpv6Policy, bool) {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		return nil, false
	}
	return o.Ipv6Policy, true
}

// HasIpv6Policy returns a boolean if a field has been set.
func (o *FtpServerProperties) HasIpv6Policy() bool {
	if o != nil && !utils.IsNil(o.Ipv6Policy) {
		return true
	}

	return false
}

// SetIpv6Policy gets a reference to the given TestIpv6Policy and assigns it to the Ipv6Policy field.
func (o *FtpServerProperties) SetIpv6Policy(v TestIpv6Policy) {
	o.Ipv6Policy = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FtpServerProperties) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerProperties) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FtpServerProperties) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FtpServerProperties) SetType(v string) {
	o.Type = &v
}

func (o FtpServerProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FtpServerProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BandwidthMeasurements) {
		toSerialize["bandwidthMeasurements"] = o.BandwidthMeasurements
	}
	if !utils.IsNil(o.DownloadLimit) {
		toSerialize["downloadLimit"] = o.DownloadLimit
	}
	if !utils.IsNil(o.FtpTargetTime) {
		toSerialize["ftpTargetTime"] = o.FtpTargetTime
	}
	if !utils.IsNil(o.FtpTimeLimit) {
		toSerialize["ftpTimeLimit"] = o.FtpTimeLimit
	}
	if !utils.IsNil(o.MtuMeasurements) {
		toSerialize["mtuMeasurements"] = o.MtuMeasurements
	}
	if !utils.IsNil(o.NetworkMeasurements) {
		toSerialize["networkMeasurements"] = o.NetworkMeasurements
	}
	if !utils.IsNil(o.NumPathTraces) {
		toSerialize["numPathTraces"] = o.NumPathTraces
	}
	toSerialize["password"] = o.Password
	if !utils.IsNil(o.PathTraceMode) {
		toSerialize["pathTraceMode"] = o.PathTraceMode
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
	toSerialize["requestType"] = o.RequestType
	toSerialize["url"] = o.Url
	if !utils.IsNil(o.UseActiveFtp) {
		toSerialize["useActiveFtp"] = o.UseActiveFtp
	}
	if !utils.IsNil(o.UseExplicitFtps) {
		toSerialize["useExplicitFtps"] = o.UseExplicitFtps
	}
	toSerialize["username"] = o.Username
	if !utils.IsNil(o.FixedPacketRate) {
		toSerialize["fixedPacketRate"] = o.FixedPacketRate
	}
	if !utils.IsNil(o.Ipv6Policy) {
		toSerialize["ipv6Policy"] = o.Ipv6Policy
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *FtpServerProperties) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
		"requestType",
		"url",
		"username",
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

	varFtpServerProperties := _FtpServerProperties{}

    err = json.Unmarshal(data, &varFtpServerProperties)

	if err != nil {
		return err
	}

	*o = FtpServerProperties(varFtpServerProperties)

	return err
}

type NullableFtpServerProperties struct {
	value *FtpServerProperties
	isSet bool
}

func (v NullableFtpServerProperties) Get() *FtpServerProperties {
	return v.value
}

func (v *NullableFtpServerProperties) Set(val *FtpServerProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableFtpServerProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableFtpServerProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFtpServerProperties(val *FtpServerProperties) *NullableFtpServerProperties {
	return &NullableFtpServerProperties{value: val, isSet: true}
}

func (v NullableFtpServerProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFtpServerProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


