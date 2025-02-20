/*
Endpoint Instant Scheduled Tests API

 You can create and execute a new endpoint instant scheduled test within ThousandEyes using this API. The test parameters are specified in the `POST` data.  The following applies to the Endpoint Instant Scheduled Tests API:  * To initiate the creation and execution of an instant scheduled test, the user must possess the `Edit endpoint tests` permission.  * Upon successful creation of an instant scheduled test, the API responds with an HTTP/201 CREATED status code and return the test definition. * It's important to note that the response does not include the results of the instant scheduled test. To retrieve test results, users can utilize the Endpoint Test Data endpoints. The URLs for these API test data endpoints are provided within the test definition output when an instant scheduled test is created. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointinstanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
	"fmt"
)

// checks if the EndpointTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointTest{}

// EndpointTest struct for EndpointTest
type EndpointTest struct {
	// A unique identifier associated with your account group. You can retrieve your `AccountGroupId` from the `/account-groups` endpoint.
	Aid *string `json:"aid,omitempty"`
	Links *EndpointTestLinks `json:"_links,omitempty"`
	AgentSelectorConfig *EndpointAgentSelectorConfig `json:"agentSelectorConfig,omitempty"`
	// UTC created date (ISO date-time format).
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	Interval *TestInterval `json:"interval,omitempty"`
	// Indicates if test is enabled.
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// Indicates if the test is a saved event.
	IsSavedEvent *bool `json:"isSavedEvent,omitempty"`
	// Enables \"in session\" path trace. When enabled, this option initiates a TCP session with the target server and sends path trace packets within the established TCP session.
	HasPathTraceInSession *bool `json:"hasPathTraceInSession,omitempty"`
	// UTC last modification date (ISO date-time format).
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Enable or disable network measurements. Set to true to enable or false to disable network measurements.
	NetworkMeasurements *bool `json:"networkMeasurements,omitempty"`
	Protocol *EndpointTestProtocol `json:"protocol,omitempty"`
	IpVersion *EndpointIpVersionTemplate `json:"ipVersion,omitempty"`
	// Target domain name or IP address.
	Server *string `json:"server,omitempty"`
	// Each test is assigned a unique ID to access test data from other endpoints.
	TestId *string `json:"testId,omitempty"`
	// Name of the test.
	TestName *string `json:"testName,omitempty"`
	Type EndpointScheduledTestType `json:"type"`
	TcpProbeMode *TestProbeModeResponse `json:"tcpProbeMode,omitempty"`
	// Port number.
	Port *int32 `json:"port,omitempty"`
}

type _EndpointTest EndpointTest

// NewEndpointTest instantiates a new EndpointTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTest(type_ EndpointScheduledTestType) *EndpointTest {
	this := EndpointTest{}
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = &interval
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var protocol EndpointTestProtocol = ENDPOINTTESTPROTOCOL_ICMP
	this.Protocol = &protocol
	this.Type = type_
	var tcpProbeMode TestProbeModeResponse = TESTPROBEMODERESPONSE_AUTO
	this.TcpProbeMode = &tcpProbeMode
	var port int32 = 443
	this.Port = &port
	return &this
}

// NewEndpointTestWithDefaults instantiates a new EndpointTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTestWithDefaults() *EndpointTest {
	this := EndpointTest{}
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = &interval
	var isEnabled bool = true
	this.IsEnabled = &isEnabled
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var protocol EndpointTestProtocol = ENDPOINTTESTPROTOCOL_ICMP
	this.Protocol = &protocol
	var tcpProbeMode TestProbeModeResponse = TESTPROBEMODERESPONSE_AUTO
	this.TcpProbeMode = &tcpProbeMode
	var port int32 = 443
	this.Port = &port
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *EndpointTest) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *EndpointTest) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *EndpointTest) SetAid(v string) {
	o.Aid = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EndpointTest) GetLinks() EndpointTestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret EndpointTestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetLinksOk() (*EndpointTestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EndpointTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given EndpointTestLinks and assigns it to the Links field.
func (o *EndpointTest) SetLinks(v EndpointTestLinks) {
	o.Links = &v
}

// GetAgentSelectorConfig returns the AgentSelectorConfig field value if set, zero value otherwise.
func (o *EndpointTest) GetAgentSelectorConfig() EndpointAgentSelectorConfig {
	if o == nil || utils.IsNil(o.AgentSelectorConfig) {
		var ret EndpointAgentSelectorConfig
		return ret
	}
	return *o.AgentSelectorConfig
}

// GetAgentSelectorConfigOk returns a tuple with the AgentSelectorConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetAgentSelectorConfigOk() (*EndpointAgentSelectorConfig, bool) {
	if o == nil || utils.IsNil(o.AgentSelectorConfig) {
		return nil, false
	}
	return o.AgentSelectorConfig, true
}

// HasAgentSelectorConfig returns a boolean if a field has been set.
func (o *EndpointTest) HasAgentSelectorConfig() bool {
	if o != nil && !utils.IsNil(o.AgentSelectorConfig) {
		return true
	}

	return false
}

// SetAgentSelectorConfig gets a reference to the given EndpointAgentSelectorConfig and assigns it to the AgentSelectorConfig field.
func (o *EndpointTest) SetAgentSelectorConfig(v EndpointAgentSelectorConfig) {
	o.AgentSelectorConfig = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *EndpointTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *EndpointTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *EndpointTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *EndpointTest) GetInterval() TestInterval {
	if o == nil || utils.IsNil(o.Interval) {
		var ret TestInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetIntervalOk() (*TestInterval, bool) {
	if o == nil || utils.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *EndpointTest) HasInterval() bool {
	if o != nil && !utils.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given TestInterval and assigns it to the Interval field.
func (o *EndpointTest) SetInterval(v TestInterval) {
	o.Interval = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *EndpointTest) GetIsEnabled() bool {
	if o == nil || utils.IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetIsEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *EndpointTest) HasIsEnabled() bool {
	if o != nil && !utils.IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *EndpointTest) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsSavedEvent returns the IsSavedEvent field value if set, zero value otherwise.
func (o *EndpointTest) GetIsSavedEvent() bool {
	if o == nil || utils.IsNil(o.IsSavedEvent) {
		var ret bool
		return ret
	}
	return *o.IsSavedEvent
}

// GetIsSavedEventOk returns a tuple with the IsSavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetIsSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsSavedEvent) {
		return nil, false
	}
	return o.IsSavedEvent, true
}

// HasIsSavedEvent returns a boolean if a field has been set.
func (o *EndpointTest) HasIsSavedEvent() bool {
	if o != nil && !utils.IsNil(o.IsSavedEvent) {
		return true
	}

	return false
}

// SetIsSavedEvent gets a reference to the given bool and assigns it to the IsSavedEvent field.
func (o *EndpointTest) SetIsSavedEvent(v bool) {
	o.IsSavedEvent = &v
}

// GetHasPathTraceInSession returns the HasPathTraceInSession field value if set, zero value otherwise.
func (o *EndpointTest) GetHasPathTraceInSession() bool {
	if o == nil || utils.IsNil(o.HasPathTraceInSession) {
		var ret bool
		return ret
	}
	return *o.HasPathTraceInSession
}

// GetHasPathTraceInSessionOk returns a tuple with the HasPathTraceInSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetHasPathTraceInSessionOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasPathTraceInSession) {
		return nil, false
	}
	return o.HasPathTraceInSession, true
}

// HasHasPathTraceInSession returns a boolean if a field has been set.
func (o *EndpointTest) HasHasPathTraceInSession() bool {
	if o != nil && !utils.IsNil(o.HasPathTraceInSession) {
		return true
	}

	return false
}

// SetHasPathTraceInSession gets a reference to the given bool and assigns it to the HasPathTraceInSession field.
func (o *EndpointTest) SetHasPathTraceInSession(v bool) {
	o.HasPathTraceInSession = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *EndpointTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *EndpointTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *EndpointTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetNetworkMeasurements returns the NetworkMeasurements field value if set, zero value otherwise.
func (o *EndpointTest) GetNetworkMeasurements() bool {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		var ret bool
		return ret
	}
	return *o.NetworkMeasurements
}

// GetNetworkMeasurementsOk returns a tuple with the NetworkMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetNetworkMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		return nil, false
	}
	return o.NetworkMeasurements, true
}

// HasNetworkMeasurements returns a boolean if a field has been set.
func (o *EndpointTest) HasNetworkMeasurements() bool {
	if o != nil && !utils.IsNil(o.NetworkMeasurements) {
		return true
	}

	return false
}

// SetNetworkMeasurements gets a reference to the given bool and assigns it to the NetworkMeasurements field.
func (o *EndpointTest) SetNetworkMeasurements(v bool) {
	o.NetworkMeasurements = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *EndpointTest) GetProtocol() EndpointTestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret EndpointTestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetProtocolOk() (*EndpointTestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *EndpointTest) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given EndpointTestProtocol and assigns it to the Protocol field.
func (o *EndpointTest) SetProtocol(v EndpointTestProtocol) {
	o.Protocol = &v
}

// GetIpVersion returns the IpVersion field value if set, zero value otherwise.
func (o *EndpointTest) GetIpVersion() EndpointIpVersionTemplate {
	if o == nil || utils.IsNil(o.IpVersion) {
		var ret EndpointIpVersionTemplate
		return ret
	}
	return *o.IpVersion
}

// GetIpVersionOk returns a tuple with the IpVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetIpVersionOk() (*EndpointIpVersionTemplate, bool) {
	if o == nil || utils.IsNil(o.IpVersion) {
		return nil, false
	}
	return o.IpVersion, true
}

// HasIpVersion returns a boolean if a field has been set.
func (o *EndpointTest) HasIpVersion() bool {
	if o != nil && !utils.IsNil(o.IpVersion) {
		return true
	}

	return false
}

// SetIpVersion gets a reference to the given EndpointIpVersionTemplate and assigns it to the IpVersion field.
func (o *EndpointTest) SetIpVersion(v EndpointIpVersionTemplate) {
	o.IpVersion = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *EndpointTest) GetServer() string {
	if o == nil || utils.IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetServerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *EndpointTest) HasServer() bool {
	if o != nil && !utils.IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *EndpointTest) SetServer(v string) {
	o.Server = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *EndpointTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *EndpointTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *EndpointTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *EndpointTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *EndpointTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *EndpointTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value
func (o *EndpointTest) GetType() EndpointScheduledTestType {
	if o == nil {
		var ret EndpointScheduledTestType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetTypeOk() (*EndpointScheduledTestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EndpointTest) SetType(v EndpointScheduledTestType) {
	o.Type = v
}

// GetTcpProbeMode returns the TcpProbeMode field value if set, zero value otherwise.
func (o *EndpointTest) GetTcpProbeMode() TestProbeModeResponse {
	if o == nil || utils.IsNil(o.TcpProbeMode) {
		var ret TestProbeModeResponse
		return ret
	}
	return *o.TcpProbeMode
}

// GetTcpProbeModeOk returns a tuple with the TcpProbeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetTcpProbeModeOk() (*TestProbeModeResponse, bool) {
	if o == nil || utils.IsNil(o.TcpProbeMode) {
		return nil, false
	}
	return o.TcpProbeMode, true
}

// HasTcpProbeMode returns a boolean if a field has been set.
func (o *EndpointTest) HasTcpProbeMode() bool {
	if o != nil && !utils.IsNil(o.TcpProbeMode) {
		return true
	}

	return false
}

// SetTcpProbeMode gets a reference to the given TestProbeModeResponse and assigns it to the TcpProbeMode field.
func (o *EndpointTest) SetTcpProbeMode(v TestProbeModeResponse) {
	o.TcpProbeMode = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EndpointTest) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTest) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EndpointTest) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *EndpointTest) SetPort(v int32) {
	o.Port = &v
}

func (o EndpointTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTest) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !utils.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !utils.IsNil(o.IsEnabled) {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if !utils.IsNil(o.IsSavedEvent) {
		toSerialize["isSavedEvent"] = o.IsSavedEvent
	}
	if !utils.IsNil(o.HasPathTraceInSession) {
		toSerialize["hasPathTraceInSession"] = o.HasPathTraceInSession
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
	if !utils.IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.TestName) {
		toSerialize["testName"] = o.TestName
	}
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.TcpProbeMode) {
		toSerialize["tcpProbeMode"] = o.TcpProbeMode
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	return toSerialize, nil
}

func (o *EndpointTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varEndpointTest := _EndpointTest{}

    err = json.Unmarshal(data, &varEndpointTest)

	if err != nil {
		return err
	}

	*o = EndpointTest(varEndpointTest)

	return err
}

type NullableEndpointTest struct {
	value *EndpointTest
	isSet bool
}

func (v NullableEndpointTest) Get() *EndpointTest {
	return v.value
}

func (v *NullableEndpointTest) Set(val *EndpointTest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTest(val *EndpointTest) *NullableEndpointTest {
	return &NullableEndpointTest{value: val, isSet: true}
}

func (v NullableEndpointTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


