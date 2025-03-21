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

// checks if the HttpEndpointTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpEndpointTestResult{}

// HttpEndpointTestResult struct for HttpEndpointTestResult
type HttpEndpointTestResult struct {
	// A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint.
	Aid *string `json:"aid,omitempty"`
	// Unique ID of endpoint test.
	TestId *string `json:"testId,omitempty"`
	// Unique ID of endpoint agent, from `/endpoint/agents` endpoint.
	AgentId *string `json:"agentId,omitempty"`
	// Epoch time (seconds) indicating the start time of the round.
	RoundId *int32 `json:"roundId,omitempty"`
	// IP address of destination server.
	ServerIp *string `json:"serverIp,omitempty"`
	NetworkProfile *NetworkProfile `json:"networkProfile,omitempty"`
	SystemMetrics *SystemMetrics `json:"systemMetrics,omitempty"`
	OriginalTargetProfile *TargetProfile `json:"originalTargetProfile,omitempty"`
	VpnProfile *VpnProfile `json:"vpnProfile,omitempty"`
	// Time required to establish a TCP connection to the server in milliseconds.
	ConnectTime *int32 `json:"connectTime,omitempty"`
	// Time required to resolve DNS in milliseconds.
	DnsTime *int32 `json:"dnsTime,omitempty"`
	ErrorType *HttpErrorType `json:"errorType,omitempty"`
	// Error details, if an error were encountered.
	ErrorDetails *string `json:"errorDetails,omitempty"`
	Headers *HttpEndpointTestResultHeaders `json:"headers,omitempty"`
	// Number of redirects.
	NumRedirects *int32 `json:"numRedirects,omitempty"`
	// Elapsed time between first and last byte of response in milliseconds.
	ReceiveTime *int32 `json:"receiveTime,omitempty"`
	// Cumulative redirect timing in milliseconds.
	RedirectTime *int32 `json:"redirectTime,omitempty"`
	// HTTP response code.
	ResponseCode *int32 `json:"responseCode,omitempty"`
	// Time to first byte in milliseconds.
	ResponseTime *int32 `json:"responseTime,omitempty"`
	// Time to negotiate SSL/TLS in milliseconds.
	SslTime *int32 `json:"sslTime,omitempty"`
	// Total time is the response time + receive time.
	TotalTime *int32 `json:"totalTime,omitempty"`
	// Time elapsed between completion of request and first byte of response in milliseconds.
	WaitTime *int32 `json:"waitTime,omitempty"`
	// Size of content in bytes.
	WireSize *int32 `json:"wireSize,omitempty"`
}

// NewHttpEndpointTestResult instantiates a new HttpEndpointTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpEndpointTestResult() *HttpEndpointTestResult {
	this := HttpEndpointTestResult{}
	var errorType HttpErrorType = HTTPERRORTYPE_CONNECT
	this.ErrorType = &errorType
	return &this
}

// NewHttpEndpointTestResultWithDefaults instantiates a new HttpEndpointTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpEndpointTestResultWithDefaults() *HttpEndpointTestResult {
	this := HttpEndpointTestResult{}
	var errorType HttpErrorType = HTTPERRORTYPE_CONNECT
	this.ErrorType = &errorType
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *HttpEndpointTestResult) SetAid(v string) {
	o.Aid = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *HttpEndpointTestResult) SetTestId(v string) {
	o.TestId = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *HttpEndpointTestResult) SetAgentId(v string) {
	o.AgentId = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *HttpEndpointTestResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetServerIp returns the ServerIp field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetServerIp() string {
	if o == nil || utils.IsNil(o.ServerIp) {
		var ret string
		return ret
	}
	return *o.ServerIp
}

// GetServerIpOk returns a tuple with the ServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetServerIpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ServerIp) {
		return nil, false
	}
	return o.ServerIp, true
}

// HasServerIp returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasServerIp() bool {
	if o != nil && !utils.IsNil(o.ServerIp) {
		return true
	}

	return false
}

// SetServerIp gets a reference to the given string and assigns it to the ServerIp field.
func (o *HttpEndpointTestResult) SetServerIp(v string) {
	o.ServerIp = &v
}

// GetNetworkProfile returns the NetworkProfile field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetNetworkProfile() NetworkProfile {
	if o == nil || utils.IsNil(o.NetworkProfile) {
		var ret NetworkProfile
		return ret
	}
	return *o.NetworkProfile
}

// GetNetworkProfileOk returns a tuple with the NetworkProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetNetworkProfileOk() (*NetworkProfile, bool) {
	if o == nil || utils.IsNil(o.NetworkProfile) {
		return nil, false
	}
	return o.NetworkProfile, true
}

// HasNetworkProfile returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasNetworkProfile() bool {
	if o != nil && !utils.IsNil(o.NetworkProfile) {
		return true
	}

	return false
}

// SetNetworkProfile gets a reference to the given NetworkProfile and assigns it to the NetworkProfile field.
func (o *HttpEndpointTestResult) SetNetworkProfile(v NetworkProfile) {
	o.NetworkProfile = &v
}

// GetSystemMetrics returns the SystemMetrics field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetSystemMetrics() SystemMetrics {
	if o == nil || utils.IsNil(o.SystemMetrics) {
		var ret SystemMetrics
		return ret
	}
	return *o.SystemMetrics
}

// GetSystemMetricsOk returns a tuple with the SystemMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetSystemMetricsOk() (*SystemMetrics, bool) {
	if o == nil || utils.IsNil(o.SystemMetrics) {
		return nil, false
	}
	return o.SystemMetrics, true
}

// HasSystemMetrics returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasSystemMetrics() bool {
	if o != nil && !utils.IsNil(o.SystemMetrics) {
		return true
	}

	return false
}

// SetSystemMetrics gets a reference to the given SystemMetrics and assigns it to the SystemMetrics field.
func (o *HttpEndpointTestResult) SetSystemMetrics(v SystemMetrics) {
	o.SystemMetrics = &v
}

// GetOriginalTargetProfile returns the OriginalTargetProfile field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetOriginalTargetProfile() TargetProfile {
	if o == nil || utils.IsNil(o.OriginalTargetProfile) {
		var ret TargetProfile
		return ret
	}
	return *o.OriginalTargetProfile
}

// GetOriginalTargetProfileOk returns a tuple with the OriginalTargetProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetOriginalTargetProfileOk() (*TargetProfile, bool) {
	if o == nil || utils.IsNil(o.OriginalTargetProfile) {
		return nil, false
	}
	return o.OriginalTargetProfile, true
}

// HasOriginalTargetProfile returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasOriginalTargetProfile() bool {
	if o != nil && !utils.IsNil(o.OriginalTargetProfile) {
		return true
	}

	return false
}

// SetOriginalTargetProfile gets a reference to the given TargetProfile and assigns it to the OriginalTargetProfile field.
func (o *HttpEndpointTestResult) SetOriginalTargetProfile(v TargetProfile) {
	o.OriginalTargetProfile = &v
}

// GetVpnProfile returns the VpnProfile field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetVpnProfile() VpnProfile {
	if o == nil || utils.IsNil(o.VpnProfile) {
		var ret VpnProfile
		return ret
	}
	return *o.VpnProfile
}

// GetVpnProfileOk returns a tuple with the VpnProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetVpnProfileOk() (*VpnProfile, bool) {
	if o == nil || utils.IsNil(o.VpnProfile) {
		return nil, false
	}
	return o.VpnProfile, true
}

// HasVpnProfile returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasVpnProfile() bool {
	if o != nil && !utils.IsNil(o.VpnProfile) {
		return true
	}

	return false
}

// SetVpnProfile gets a reference to the given VpnProfile and assigns it to the VpnProfile field.
func (o *HttpEndpointTestResult) SetVpnProfile(v VpnProfile) {
	o.VpnProfile = &v
}

// GetConnectTime returns the ConnectTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetConnectTime() int32 {
	if o == nil || utils.IsNil(o.ConnectTime) {
		var ret int32
		return ret
	}
	return *o.ConnectTime
}

// GetConnectTimeOk returns a tuple with the ConnectTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetConnectTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ConnectTime) {
		return nil, false
	}
	return o.ConnectTime, true
}

// HasConnectTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasConnectTime() bool {
	if o != nil && !utils.IsNil(o.ConnectTime) {
		return true
	}

	return false
}

// SetConnectTime gets a reference to the given int32 and assigns it to the ConnectTime field.
func (o *HttpEndpointTestResult) SetConnectTime(v int32) {
	o.ConnectTime = &v
}

// GetDnsTime returns the DnsTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetDnsTime() int32 {
	if o == nil || utils.IsNil(o.DnsTime) {
		var ret int32
		return ret
	}
	return *o.DnsTime
}

// GetDnsTimeOk returns a tuple with the DnsTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetDnsTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.DnsTime) {
		return nil, false
	}
	return o.DnsTime, true
}

// HasDnsTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasDnsTime() bool {
	if o != nil && !utils.IsNil(o.DnsTime) {
		return true
	}

	return false
}

// SetDnsTime gets a reference to the given int32 and assigns it to the DnsTime field.
func (o *HttpEndpointTestResult) SetDnsTime(v int32) {
	o.DnsTime = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetErrorType() HttpErrorType {
	if o == nil || utils.IsNil(o.ErrorType) {
		var ret HttpErrorType
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetErrorTypeOk() (*HttpErrorType, bool) {
	if o == nil || utils.IsNil(o.ErrorType) {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasErrorType() bool {
	if o != nil && !utils.IsNil(o.ErrorType) {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given HttpErrorType and assigns it to the ErrorType field.
func (o *HttpEndpointTestResult) SetErrorType(v HttpErrorType) {
	o.ErrorType = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetErrorDetails() string {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		var ret string
		return ret
	}
	return *o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetErrorDetailsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasErrorDetails() bool {
	if o != nil && !utils.IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given string and assigns it to the ErrorDetails field.
func (o *HttpEndpointTestResult) SetErrorDetails(v string) {
	o.ErrorDetails = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetHeaders() HttpEndpointTestResultHeaders {
	if o == nil || utils.IsNil(o.Headers) {
		var ret HttpEndpointTestResultHeaders
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetHeadersOk() (*HttpEndpointTestResultHeaders, bool) {
	if o == nil || utils.IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasHeaders() bool {
	if o != nil && !utils.IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given HttpEndpointTestResultHeaders and assigns it to the Headers field.
func (o *HttpEndpointTestResult) SetHeaders(v HttpEndpointTestResultHeaders) {
	o.Headers = &v
}

// GetNumRedirects returns the NumRedirects field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetNumRedirects() int32 {
	if o == nil || utils.IsNil(o.NumRedirects) {
		var ret int32
		return ret
	}
	return *o.NumRedirects
}

// GetNumRedirectsOk returns a tuple with the NumRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetNumRedirectsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumRedirects) {
		return nil, false
	}
	return o.NumRedirects, true
}

// HasNumRedirects returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasNumRedirects() bool {
	if o != nil && !utils.IsNil(o.NumRedirects) {
		return true
	}

	return false
}

// SetNumRedirects gets a reference to the given int32 and assigns it to the NumRedirects field.
func (o *HttpEndpointTestResult) SetNumRedirects(v int32) {
	o.NumRedirects = &v
}

// GetReceiveTime returns the ReceiveTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetReceiveTime() int32 {
	if o == nil || utils.IsNil(o.ReceiveTime) {
		var ret int32
		return ret
	}
	return *o.ReceiveTime
}

// GetReceiveTimeOk returns a tuple with the ReceiveTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetReceiveTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ReceiveTime) {
		return nil, false
	}
	return o.ReceiveTime, true
}

// HasReceiveTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasReceiveTime() bool {
	if o != nil && !utils.IsNil(o.ReceiveTime) {
		return true
	}

	return false
}

// SetReceiveTime gets a reference to the given int32 and assigns it to the ReceiveTime field.
func (o *HttpEndpointTestResult) SetReceiveTime(v int32) {
	o.ReceiveTime = &v
}

// GetRedirectTime returns the RedirectTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetRedirectTime() int32 {
	if o == nil || utils.IsNil(o.RedirectTime) {
		var ret int32
		return ret
	}
	return *o.RedirectTime
}

// GetRedirectTimeOk returns a tuple with the RedirectTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetRedirectTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RedirectTime) {
		return nil, false
	}
	return o.RedirectTime, true
}

// HasRedirectTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasRedirectTime() bool {
	if o != nil && !utils.IsNil(o.RedirectTime) {
		return true
	}

	return false
}

// SetRedirectTime gets a reference to the given int32 and assigns it to the RedirectTime field.
func (o *HttpEndpointTestResult) SetRedirectTime(v int32) {
	o.RedirectTime = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetResponseCode() int32 {
	if o == nil || utils.IsNil(o.ResponseCode) {
		var ret int32
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetResponseCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasResponseCode() bool {
	if o != nil && !utils.IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int32 and assigns it to the ResponseCode field.
func (o *HttpEndpointTestResult) SetResponseCode(v int32) {
	o.ResponseCode = &v
}

// GetResponseTime returns the ResponseTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetResponseTime() int32 {
	if o == nil || utils.IsNil(o.ResponseTime) {
		var ret int32
		return ret
	}
	return *o.ResponseTime
}

// GetResponseTimeOk returns a tuple with the ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetResponseTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ResponseTime) {
		return nil, false
	}
	return o.ResponseTime, true
}

// HasResponseTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasResponseTime() bool {
	if o != nil && !utils.IsNil(o.ResponseTime) {
		return true
	}

	return false
}

// SetResponseTime gets a reference to the given int32 and assigns it to the ResponseTime field.
func (o *HttpEndpointTestResult) SetResponseTime(v int32) {
	o.ResponseTime = &v
}

// GetSslTime returns the SslTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetSslTime() int32 {
	if o == nil || utils.IsNil(o.SslTime) {
		var ret int32
		return ret
	}
	return *o.SslTime
}

// GetSslTimeOk returns a tuple with the SslTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetSslTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.SslTime) {
		return nil, false
	}
	return o.SslTime, true
}

// HasSslTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasSslTime() bool {
	if o != nil && !utils.IsNil(o.SslTime) {
		return true
	}

	return false
}

// SetSslTime gets a reference to the given int32 and assigns it to the SslTime field.
func (o *HttpEndpointTestResult) SetSslTime(v int32) {
	o.SslTime = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetTotalTime() int32 {
	if o == nil || utils.IsNil(o.TotalTime) {
		var ret int32
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetTotalTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalTime) {
		return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasTotalTime() bool {
	if o != nil && !utils.IsNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int32 and assigns it to the TotalTime field.
func (o *HttpEndpointTestResult) SetTotalTime(v int32) {
	o.TotalTime = &v
}

// GetWaitTime returns the WaitTime field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetWaitTime() int32 {
	if o == nil || utils.IsNil(o.WaitTime) {
		var ret int32
		return ret
	}
	return *o.WaitTime
}

// GetWaitTimeOk returns a tuple with the WaitTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetWaitTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.WaitTime) {
		return nil, false
	}
	return o.WaitTime, true
}

// HasWaitTime returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasWaitTime() bool {
	if o != nil && !utils.IsNil(o.WaitTime) {
		return true
	}

	return false
}

// SetWaitTime gets a reference to the given int32 and assigns it to the WaitTime field.
func (o *HttpEndpointTestResult) SetWaitTime(v int32) {
	o.WaitTime = &v
}

// GetWireSize returns the WireSize field value if set, zero value otherwise.
func (o *HttpEndpointTestResult) GetWireSize() int32 {
	if o == nil || utils.IsNil(o.WireSize) {
		var ret int32
		return ret
	}
	return *o.WireSize
}

// GetWireSizeOk returns a tuple with the WireSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResult) GetWireSizeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.WireSize) {
		return nil, false
	}
	return o.WireSize, true
}

// HasWireSize returns a boolean if a field has been set.
func (o *HttpEndpointTestResult) HasWireSize() bool {
	if o != nil && !utils.IsNil(o.WireSize) {
		return true
	}

	return false
}

// SetWireSize gets a reference to the given int32 and assigns it to the WireSize field.
func (o *HttpEndpointTestResult) SetWireSize(v int32) {
	o.WireSize = &v
}

func (o HttpEndpointTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpEndpointTestResult) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.ConnectTime) {
		toSerialize["connectTime"] = o.ConnectTime
	}
	if !utils.IsNil(o.DnsTime) {
		toSerialize["dnsTime"] = o.DnsTime
	}
	if !utils.IsNil(o.ErrorType) {
		toSerialize["errorType"] = o.ErrorType
	}
	if !utils.IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !utils.IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !utils.IsNil(o.NumRedirects) {
		toSerialize["numRedirects"] = o.NumRedirects
	}
	if !utils.IsNil(o.ReceiveTime) {
		toSerialize["receiveTime"] = o.ReceiveTime
	}
	if !utils.IsNil(o.RedirectTime) {
		toSerialize["redirectTime"] = o.RedirectTime
	}
	if !utils.IsNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !utils.IsNil(o.ResponseTime) {
		toSerialize["responseTime"] = o.ResponseTime
	}
	if !utils.IsNil(o.SslTime) {
		toSerialize["sslTime"] = o.SslTime
	}
	if !utils.IsNil(o.TotalTime) {
		toSerialize["totalTime"] = o.TotalTime
	}
	if !utils.IsNil(o.WaitTime) {
		toSerialize["waitTime"] = o.WaitTime
	}
	if !utils.IsNil(o.WireSize) {
		toSerialize["wireSize"] = o.WireSize
	}
	return toSerialize, nil
}

type NullableHttpEndpointTestResult struct {
	value *HttpEndpointTestResult
	isSet bool
}

func (v NullableHttpEndpointTestResult) Get() *HttpEndpointTestResult {
	return v.value
}

func (v *NullableHttpEndpointTestResult) Set(val *HttpEndpointTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpEndpointTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpEndpointTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpEndpointTestResult(val *HttpEndpointTestResult) *NullableHttpEndpointTestResult {
	return &NullableHttpEndpointTestResult{value: val, isSet: true}
}

func (v NullableHttpEndpointTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpEndpointTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


