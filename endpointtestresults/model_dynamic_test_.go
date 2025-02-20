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
	"time"
)

// checks if the DynamicTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DynamicTest{}

// DynamicTest struct for DynamicTest
type DynamicTest struct {
	// A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint.
	Aid *string `json:"aid,omitempty"`
	Links *DynamicTestLinks `json:"_links,omitempty"`
	AgentSelectorConfig *EndpointAgentSelectorConfig `json:"agentSelectorConfig,omitempty"`
	// Which supported application to monitor, can be one of `webex`, `zoom`, `microsoft-teams`.
	Application *string `json:"application,omitempty"`
	// UTC created date (ISO date-time format).
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	Interval *TestInterval `json:"interval,omitempty"`
	// Indicates if test is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Enables \"in session\" path trace. When enabled, this option initiates a TCP session with the target server and sends path trace packets within the established TCP session.
	HasPathTraceInSession *bool `json:"hasPathTraceInSession,omitempty"`
	// Optional flag indicating if the test should run ping.
	HasPing *bool `json:"hasPing,omitempty"`
	// Optional flag indicating if the test should run traceroute.
	HasTraceroute *bool `json:"hasTraceroute,omitempty"`
	// UTC last modification date (ISO date-time format).
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Enable or disable network measurements. Set to true to enable or false to disable network measurements.
	NetworkMeasurements *bool `json:"networkMeasurements,omitempty"`
	Protocol *EndpointTestProtocol `json:"protocol,omitempty"`
	IpVersion *EndpointIpVersionTemplate `json:"ipVersion,omitempty"`
	TcpProbeMode *TestProbeModeResponse `json:"tcpProbeMode,omitempty"`
	// Each test is assigned a unique ID; this is used to access test information and results from other endpoints.
	TestId *string `json:"testId,omitempty"`
	// Name of the test.
	TestName *string `json:"testName,omitempty"`
	Labels []TestLabel `json:"labels,omitempty"`
}

// NewDynamicTest instantiates a new DynamicTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicTest() *DynamicTest {
	this := DynamicTest{}
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = &interval
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	var hasPing bool = true
	this.HasPing = &hasPing
	var hasTraceroute bool = true
	this.HasTraceroute = &hasTraceroute
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var protocol EndpointTestProtocol = ENDPOINTTESTPROTOCOL_ICMP
	this.Protocol = &protocol
	var tcpProbeMode TestProbeModeResponse = TESTPROBEMODERESPONSE_AUTO
	this.TcpProbeMode = &tcpProbeMode
	return &this
}

// NewDynamicTestWithDefaults instantiates a new DynamicTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicTestWithDefaults() *DynamicTest {
	this := DynamicTest{}
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = &interval
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	var hasPing bool = true
	this.HasPing = &hasPing
	var hasTraceroute bool = true
	this.HasTraceroute = &hasTraceroute
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var protocol EndpointTestProtocol = ENDPOINTTESTPROTOCOL_ICMP
	this.Protocol = &protocol
	var tcpProbeMode TestProbeModeResponse = TESTPROBEMODERESPONSE_AUTO
	this.TcpProbeMode = &tcpProbeMode
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *DynamicTest) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *DynamicTest) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *DynamicTest) SetAid(v string) {
	o.Aid = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DynamicTest) GetLinks() DynamicTestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret DynamicTestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetLinksOk() (*DynamicTestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DynamicTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DynamicTestLinks and assigns it to the Links field.
func (o *DynamicTest) SetLinks(v DynamicTestLinks) {
	o.Links = &v
}

// GetAgentSelectorConfig returns the AgentSelectorConfig field value if set, zero value otherwise.
func (o *DynamicTest) GetAgentSelectorConfig() EndpointAgentSelectorConfig {
	if o == nil || utils.IsNil(o.AgentSelectorConfig) {
		var ret EndpointAgentSelectorConfig
		return ret
	}
	return *o.AgentSelectorConfig
}

// GetAgentSelectorConfigOk returns a tuple with the AgentSelectorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetAgentSelectorConfigOk() (*EndpointAgentSelectorConfig, bool) {
	if o == nil || utils.IsNil(o.AgentSelectorConfig) {
		return nil, false
	}
	return o.AgentSelectorConfig, true
}

// HasAgentSelectorConfig returns a boolean if a field has been set.
func (o *DynamicTest) HasAgentSelectorConfig() bool {
	if o != nil && !utils.IsNil(o.AgentSelectorConfig) {
		return true
	}

	return false
}

// SetAgentSelectorConfig gets a reference to the given EndpointAgentSelectorConfig and assigns it to the AgentSelectorConfig field.
func (o *DynamicTest) SetAgentSelectorConfig(v EndpointAgentSelectorConfig) {
	o.AgentSelectorConfig = &v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *DynamicTest) GetApplication() string {
	if o == nil || utils.IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetApplicationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *DynamicTest) HasApplication() bool {
	if o != nil && !utils.IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *DynamicTest) SetApplication(v string) {
	o.Application = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *DynamicTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *DynamicTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *DynamicTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *DynamicTest) GetInterval() TestInterval {
	if o == nil || utils.IsNil(o.Interval) {
		var ret TestInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetIntervalOk() (*TestInterval, bool) {
	if o == nil || utils.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *DynamicTest) HasInterval() bool {
	if o != nil && !utils.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given TestInterval and assigns it to the Interval field.
func (o *DynamicTest) SetInterval(v TestInterval) {
	o.Interval = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *DynamicTest) GetIsEnabled() bool {
	if o == nil || utils.IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *DynamicTest) HasIsEnabled() bool {
	if o != nil && !utils.IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *DynamicTest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetHasPathTraceInSession returns the HasPathTraceInSession field value if set, zero value otherwise.
func (o *DynamicTest) GetHasPathTraceInSession() bool {
	if o == nil || utils.IsNil(o.HasPathTraceInSession) {
		var ret bool
		return ret
	}
	return *o.HasPathTraceInSession
}

// GetHasPathTraceInSessionOk returns a tuple with the HasPathTraceInSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetHasPathTraceInSessionOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasPathTraceInSession) {
		return nil, false
	}
	return o.HasPathTraceInSession, true
}

// HasHasPathTraceInSession returns a boolean if a field has been set.
func (o *DynamicTest) HasHasPathTraceInSession() bool {
	if o != nil && !utils.IsNil(o.HasPathTraceInSession) {
		return true
	}

	return false
}

// SetHasPathTraceInSession gets a reference to the given bool and assigns it to the HasPathTraceInSession field.
func (o *DynamicTest) SetHasPathTraceInSession(v bool) {
	o.HasPathTraceInSession = &v
}

// GetHasPing returns the HasPing field value if set, zero value otherwise.
func (o *DynamicTest) GetHasPing() bool {
	if o == nil || utils.IsNil(o.HasPing) {
		var ret bool
		return ret
	}
	return *o.HasPing
}

// GetHasPingOk returns a tuple with the HasPing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetHasPingOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasPing) {
		return nil, false
	}
	return o.HasPing, true
}

// HasHasPing returns a boolean if a field has been set.
func (o *DynamicTest) HasHasPing() bool {
	if o != nil && !utils.IsNil(o.HasPing) {
		return true
	}

	return false
}

// SetHasPing gets a reference to the given bool and assigns it to the HasPing field.
func (o *DynamicTest) SetHasPing(v bool) {
	o.HasPing = &v
}

// GetHasTraceroute returns the HasTraceroute field value if set, zero value otherwise.
func (o *DynamicTest) GetHasTraceroute() bool {
	if o == nil || utils.IsNil(o.HasTraceroute) {
		var ret bool
		return ret
	}
	return *o.HasTraceroute
}

// GetHasTracerouteOk returns a tuple with the HasTraceroute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetHasTracerouteOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasTraceroute) {
		return nil, false
	}
	return o.HasTraceroute, true
}

// HasHasTraceroute returns a boolean if a field has been set.
func (o *DynamicTest) HasHasTraceroute() bool {
	if o != nil && !utils.IsNil(o.HasTraceroute) {
		return true
	}

	return false
}

// SetHasTraceroute gets a reference to the given bool and assigns it to the HasTraceroute field.
func (o *DynamicTest) SetHasTraceroute(v bool) {
	o.HasTraceroute = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *DynamicTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *DynamicTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *DynamicTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetNetworkMeasurements returns the NetworkMeasurements field value if set, zero value otherwise.
func (o *DynamicTest) GetNetworkMeasurements() bool {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		var ret bool
		return ret
	}
	return *o.NetworkMeasurements
}

// GetNetworkMeasurementsOk returns a tuple with the NetworkMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetNetworkMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		return nil, false
	}
	return o.NetworkMeasurements, true
}

// HasNetworkMeasurements returns a boolean if a field has been set.
func (o *DynamicTest) HasNetworkMeasurements() bool {
	if o != nil && !utils.IsNil(o.NetworkMeasurements) {
		return true
	}

	return false
}

// SetNetworkMeasurements gets a reference to the given bool and assigns it to the NetworkMeasurements field.
func (o *DynamicTest) SetNetworkMeasurements(v bool) {
	o.NetworkMeasurements = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *DynamicTest) GetProtocol() EndpointTestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret EndpointTestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetProtocolOk() (*EndpointTestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *DynamicTest) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given EndpointTestProtocol and assigns it to the Protocol field.
func (o *DynamicTest) SetProtocol(v EndpointTestProtocol) {
	o.Protocol = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *DynamicTest) GetIpVersion() EndpointIpVersionTemplate {
	if o == nil || utils.IsNil(o.IpVersion) {
		var ret EndpointIpVersionTemplate
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetIpVersionOk() (*EndpointIpVersionTemplate, bool) {
	if o == nil || utils.IsNil(o.IpVersion) {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *DynamicTest) HasIpVersion() bool {
	if o != nil && !utils.IsNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given EndpointIpVersionTemplate and assigns it to the IpVersion field.
func (o *DynamicTest) SetIpVersion(v EndpointIpVersionTemplate) {
	o.IpVersion = &v
}

// GetTcpProbeMode returns the TcpProbeMode field value if set, zero value otherwise.
func (o *DynamicTest) GetTcpProbeMode() TestProbeModeResponse {
	if o == nil || utils.IsNil(o.TcpProbeMode) {
		var ret TestProbeModeResponse
		return ret
	}
	return *o.TcpProbeMode
}

// GetTcpProbeModeOk returns a tuple with the TcpProbeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetTcpProbeModeOk() (*TestProbeModeResponse, bool) {
	if o == nil || utils.IsNil(o.TcpProbeMode) {
		return nil, false
	}
	return o.TcpProbeMode, true
}

// HasTcpProbeMode returns a boolean if a field has been set.
func (o *DynamicTest) HasTcpProbeMode() bool {
	if o != nil && !utils.IsNil(o.TcpProbeMode) {
		return true
	}

	return false
}

// SetTcpProbeMode gets a reference to the given TestProbeModeResponse and assigns it to the TcpProbeMode field.
func (o *DynamicTest) SetTcpProbeMode(v TestProbeModeResponse) {
	o.TcpProbeMode = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *DynamicTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *DynamicTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *DynamicTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *DynamicTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *DynamicTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *DynamicTest) SetTestName(v string) {
	o.TestName = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DynamicTest) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicTest) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DynamicTest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *DynamicTest) SetLabels(v []TestLabel) {
	o.Labels = v
}

func (o DynamicTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !utils.IsNil(o.AgentSelectorConfig) {
		toSerialize["agentSelectorConfig"] = o.AgentSelectorConfig
	}
	if !utils.IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !utils.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !utils.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !utils.IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !utils.IsNil(o.HasPathTraceInSession) {
		toSerialize["hasPathTraceInSession"] = o.HasPathTraceInSession
	}
	if !utils.IsNil(o.HasPing) {
		toSerialize["hasPing"] = o.HasPing
	}
	if !utils.IsNil(o.HasTraceroute) {
		toSerialize["hasTraceroute"] = o.HasTraceroute
	}
	if !utils.IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !utils.IsNil(o.NetworkMeasurements) {
		toSerialize["networkMeasurements"] = o.NetworkMeasurements
	}
	if !utils.IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !utils.IsNil(o.IpVersion) {
		toSerialize["ipVersion"] = o.IpVersion
	}
	if !utils.IsNil(o.TcpProbeMode) {
		toSerialize["tcpProbeMode"] = o.TcpProbeMode
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.TestName) {
		toSerialize["testName"] = o.TestName
	}
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableDynamicTest struct {
	value *DynamicTest
	isSet bool
}

func (v NullableDynamicTest) Get() *DynamicTest {
	return v.value
}

func (v *NullableDynamicTest) Set(val *DynamicTest) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicTest) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicTest(val *DynamicTest) *NullableDynamicTest {
	return &NullableDynamicTest{value: val, isSet: true}
}

func (v NullableDynamicTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


