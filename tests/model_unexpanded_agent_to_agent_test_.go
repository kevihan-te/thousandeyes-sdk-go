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
	"time"
	"fmt"
)

// checks if the UnexpandedAgentToAgentTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UnexpandedAgentToAgentTest{}

// UnexpandedAgentToAgentTest struct for UnexpandedAgentToAgentTest
type UnexpandedAgentToAgentTest struct {
	Interval TestInterval `json:"interval"`
	// Indicates if alerts are enabled.
	AlertsEnabled *bool `json:"alertsEnabled,omitempty"`
	// Test is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// User that created the test.
	CreatedBy *string `json:"createdBy,omitempty"`
	// UTC created date (ISO date-time format).
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// A description of the test.
	Description *string `json:"description,omitempty"`
	// Indicates if the test is shared with the account group.
	LiveShare *bool `json:"liveShare,omitempty"`
	// User that modified the test.
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// UTC last modification date (ISO date-time format).
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Indicates if the test is a saved event.
	SavedEvent *bool `json:"savedEvent,omitempty"`
	// Each test is assigned an unique ID; this is used to access test information and results from other endpoints.
	TestId *string `json:"testId,omitempty"`
	// The name of the test. Test name must be unique.
	TestName *string `json:"testName,omitempty"`
	Type *string `json:"type,omitempty"`
	Links *TestLinks `json:"_links,omitempty"`
	Direction *TestDirection `json:"direction,omitempty"`
	// DSCP label.
	Dscp *string `json:"dscp,omitempty"`
	DscpId *TestDscpId `json:"dscpId,omitempty"`
	// Maximum segment size, in bytes.
	Mss *int32 `json:"mss,omitempty"`
	// Number of path traces executed by the agent.
	NumPathTraces *int32 `json:"numPathTraces,omitempty"`
	PathTraceMode *TestPathTraceMode `json:"pathTraceMode,omitempty"`
	// Target port.
	Port *int32 `json:"port,omitempty"`
	Protocol *AgentToAgentTestProtocol `json:"protocol,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	// `agentId` of the target agent for the test.
	TargetAgentId string `json:"targetAgentId"`
	// Enable or disable throughput measurements. Throughput measurements cannot be enabled when the source or target of the test is a cloud agent.
	ThroughputMeasurements *bool `json:"throughputMeasurements,omitempty"`
	// The throughput duration.
	ThroughputDuration *int32 `json:"throughputDuration,omitempty"`
	// The throughput rate, only applicable for UDP protocol.
	ThroughputRate *int32 `json:"throughputRate,omitempty"`
	// Sets packets rate sent to measure the network in packets per second.
	FixedPacketRate *int32 `json:"fixedPacketRate,omitempty"`
	// Set to `true` to enable bgp measurements.
	BgpMeasurements *bool `json:"bgpMeasurements,omitempty"`
	// Indicate if all available public BGP monitors should be used, when ommited defaults to `bgpMeasurements` value.
	UsePublicBgp *bool `json:"usePublicBgp,omitempty"`
}

type _UnexpandedAgentToAgentTest UnexpandedAgentToAgentTest

// NewUnexpandedAgentToAgentTest instantiates a new UnexpandedAgentToAgentTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnexpandedAgentToAgentTest(interval TestInterval, targetAgentId string) *UnexpandedAgentToAgentTest {
	this := UnexpandedAgentToAgentTest{}
	this.Interval = interval
	var enabled bool = true
	this.Enabled = &enabled
	var direction TestDirection = TESTDIRECTION_TO_TARGET
	this.Direction = &direction
	var dscpId TestDscpId = TESTDSCPID__0
	this.DscpId = &dscpId
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var port int32 = 49153
	this.Port = &port
	var protocol AgentToAgentTestProtocol = AGENTTOAGENTTESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	this.TargetAgentId = targetAgentId
	var throughputMeasurements bool = false
	this.ThroughputMeasurements = &throughputMeasurements
	var throughputDuration int32 = 10000
	this.ThroughputDuration = &throughputDuration
	var bgpMeasurements bool = true
	this.BgpMeasurements = &bgpMeasurements
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	return &this
}

// NewUnexpandedAgentToAgentTestWithDefaults instantiates a new UnexpandedAgentToAgentTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnexpandedAgentToAgentTestWithDefaults() *UnexpandedAgentToAgentTest {
	this := UnexpandedAgentToAgentTest{}
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = interval
	var enabled bool = true
	this.Enabled = &enabled
	var direction TestDirection = TESTDIRECTION_TO_TARGET
	this.Direction = &direction
	var dscpId TestDscpId = TESTDSCPID__0
	this.DscpId = &dscpId
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var port int32 = 49153
	this.Port = &port
	var protocol AgentToAgentTestProtocol = AGENTTOAGENTTESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	var throughputMeasurements bool = false
	this.ThroughputMeasurements = &throughputMeasurements
	var throughputDuration int32 = 10000
	this.ThroughputDuration = &throughputDuration
	var bgpMeasurements bool = true
	this.BgpMeasurements = &bgpMeasurements
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	return &this
}

// GetInterval returns the Interval field value
func (o *UnexpandedAgentToAgentTest) GetInterval() TestInterval {
	if o == nil {
		var ret TestInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetIntervalOk() (*TestInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *UnexpandedAgentToAgentTest) SetInterval(v TestInterval) {
	o.Interval = v
}

// GetAlertsEnabled returns the AlertsEnabled field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetAlertsEnabled() bool {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		var ret bool
		return ret
	}
	return *o.AlertsEnabled
}

// GetAlertsEnabledOk returns a tuple with the AlertsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetAlertsEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		return nil, false
	}
	return o.AlertsEnabled, true
}

// HasAlertsEnabled returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasAlertsEnabled() bool {
	if o != nil && !utils.IsNil(o.AlertsEnabled) {
		return true
	}

	return false
}

// SetAlertsEnabled gets a reference to the given bool and assigns it to the AlertsEnabled field.
func (o *UnexpandedAgentToAgentTest) SetAlertsEnabled(v bool) {
	o.AlertsEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UnexpandedAgentToAgentTest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *UnexpandedAgentToAgentTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *UnexpandedAgentToAgentTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnexpandedAgentToAgentTest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *UnexpandedAgentToAgentTest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *UnexpandedAgentToAgentTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *UnexpandedAgentToAgentTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *UnexpandedAgentToAgentTest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *UnexpandedAgentToAgentTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *UnexpandedAgentToAgentTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UnexpandedAgentToAgentTest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *UnexpandedAgentToAgentTest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetDirection() TestDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret TestDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetDirectionOk() (*TestDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given TestDirection and assigns it to the Direction field.
func (o *UnexpandedAgentToAgentTest) SetDirection(v TestDirection) {
	o.Direction = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetDscp() string {
	if o == nil || utils.IsNil(o.Dscp) {
		var ret string
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetDscpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasDscp() bool {
	if o != nil && !utils.IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given string and assigns it to the Dscp field.
func (o *UnexpandedAgentToAgentTest) SetDscp(v string) {
	o.Dscp = &v
}

// GetDscpId returns the DscpId field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetDscpId() TestDscpId {
	if o == nil || utils.IsNil(o.DscpId) {
		var ret TestDscpId
		return ret
	}
	return *o.DscpId
}

// GetDscpIdOk returns a tuple with the DscpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetDscpIdOk() (*TestDscpId, bool) {
	if o == nil || utils.IsNil(o.DscpId) {
		return nil, false
	}
	return o.DscpId, true
}

// HasDscpId returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasDscpId() bool {
	if o != nil && !utils.IsNil(o.DscpId) {
		return true
	}

	return false
}

// SetDscpId gets a reference to the given TestDscpId and assigns it to the DscpId field.
func (o *UnexpandedAgentToAgentTest) SetDscpId(v TestDscpId) {
	o.DscpId = &v
}

// GetMss returns the Mss field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetMss() int32 {
	if o == nil || utils.IsNil(o.Mss) {
		var ret int32
		return ret
	}
	return *o.Mss
}

// GetMssOk returns a tuple with the Mss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetMssOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Mss) {
		return nil, false
	}
	return o.Mss, true
}

// HasMss returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasMss() bool {
	if o != nil && !utils.IsNil(o.Mss) {
		return true
	}

	return false
}

// SetMss gets a reference to the given int32 and assigns it to the Mss field.
func (o *UnexpandedAgentToAgentTest) SetMss(v int32) {
	o.Mss = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *UnexpandedAgentToAgentTest) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPathTraceMode returns the PathTraceMode field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetPathTraceMode() TestPathTraceMode {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		var ret TestPathTraceMode
		return ret
	}
	return *o.PathTraceMode
}

// GetPathTraceModeOk returns a tuple with the PathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetPathTraceModeOk() (*TestPathTraceMode, bool) {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		return nil, false
	}
	return o.PathTraceMode, true
}

// HasPathTraceMode returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.PathTraceMode) {
		return true
	}

	return false
}

// SetPathTraceMode gets a reference to the given TestPathTraceMode and assigns it to the PathTraceMode field.
func (o *UnexpandedAgentToAgentTest) SetPathTraceMode(v TestPathTraceMode) {
	o.PathTraceMode = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UnexpandedAgentToAgentTest) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetProtocol() AgentToAgentTestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret AgentToAgentTestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetProtocolOk() (*AgentToAgentTestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given AgentToAgentTestProtocol and assigns it to the Protocol field.
func (o *UnexpandedAgentToAgentTest) SetProtocol(v AgentToAgentTestProtocol) {
	o.Protocol = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *UnexpandedAgentToAgentTest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetTargetAgentId returns the TargetAgentId field value
func (o *UnexpandedAgentToAgentTest) GetTargetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAgentId
}

// GetTargetAgentIdOk returns a tuple with the TargetAgentId field value
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetTargetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAgentId, true
}

// SetTargetAgentId sets field value
func (o *UnexpandedAgentToAgentTest) SetTargetAgentId(v string) {
	o.TargetAgentId = v
}

// GetThroughputMeasurements returns the ThroughputMeasurements field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetThroughputMeasurements() bool {
	if o == nil || utils.IsNil(o.ThroughputMeasurements) {
		var ret bool
		return ret
	}
	return *o.ThroughputMeasurements
}

// GetThroughputMeasurementsOk returns a tuple with the ThroughputMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetThroughputMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ThroughputMeasurements) {
		return nil, false
	}
	return o.ThroughputMeasurements, true
}

// HasThroughputMeasurements returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasThroughputMeasurements() bool {
	if o != nil && !utils.IsNil(o.ThroughputMeasurements) {
		return true
	}

	return false
}

// SetThroughputMeasurements gets a reference to the given bool and assigns it to the ThroughputMeasurements field.
func (o *UnexpandedAgentToAgentTest) SetThroughputMeasurements(v bool) {
	o.ThroughputMeasurements = &v
}

// GetThroughputDuration returns the ThroughputDuration field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetThroughputDuration() int32 {
	if o == nil || utils.IsNil(o.ThroughputDuration) {
		var ret int32
		return ret
	}
	return *o.ThroughputDuration
}

// GetThroughputDurationOk returns a tuple with the ThroughputDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetThroughputDurationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ThroughputDuration) {
		return nil, false
	}
	return o.ThroughputDuration, true
}

// HasThroughputDuration returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasThroughputDuration() bool {
	if o != nil && !utils.IsNil(o.ThroughputDuration) {
		return true
	}

	return false
}

// SetThroughputDuration gets a reference to the given int32 and assigns it to the ThroughputDuration field.
func (o *UnexpandedAgentToAgentTest) SetThroughputDuration(v int32) {
	o.ThroughputDuration = &v
}

// GetThroughputRate returns the ThroughputRate field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetThroughputRate() int32 {
	if o == nil || utils.IsNil(o.ThroughputRate) {
		var ret int32
		return ret
	}
	return *o.ThroughputRate
}

// GetThroughputRateOk returns a tuple with the ThroughputRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetThroughputRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ThroughputRate) {
		return nil, false
	}
	return o.ThroughputRate, true
}

// HasThroughputRate returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasThroughputRate() bool {
	if o != nil && !utils.IsNil(o.ThroughputRate) {
		return true
	}

	return false
}

// SetThroughputRate gets a reference to the given int32 and assigns it to the ThroughputRate field.
func (o *UnexpandedAgentToAgentTest) SetThroughputRate(v int32) {
	o.ThroughputRate = &v
}

// GetFixedPacketRate returns the FixedPacketRate field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetFixedPacketRate() int32 {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		var ret int32
		return ret
	}
	return *o.FixedPacketRate
}

// GetFixedPacketRateOk returns a tuple with the FixedPacketRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetFixedPacketRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		return nil, false
	}
	return o.FixedPacketRate, true
}

// HasFixedPacketRate returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasFixedPacketRate() bool {
	if o != nil && !utils.IsNil(o.FixedPacketRate) {
		return true
	}

	return false
}

// SetFixedPacketRate gets a reference to the given int32 and assigns it to the FixedPacketRate field.
func (o *UnexpandedAgentToAgentTest) SetFixedPacketRate(v int32) {
	o.FixedPacketRate = &v
}

// GetBgpMeasurements returns the BgpMeasurements field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetBgpMeasurements() bool {
	if o == nil || utils.IsNil(o.BgpMeasurements) {
		var ret bool
		return ret
	}
	return *o.BgpMeasurements
}

// GetBgpMeasurementsOk returns a tuple with the BgpMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetBgpMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BgpMeasurements) {
		return nil, false
	}
	return o.BgpMeasurements, true
}

// HasBgpMeasurements returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasBgpMeasurements() bool {
	if o != nil && !utils.IsNil(o.BgpMeasurements) {
		return true
	}

	return false
}

// SetBgpMeasurements gets a reference to the given bool and assigns it to the BgpMeasurements field.
func (o *UnexpandedAgentToAgentTest) SetBgpMeasurements(v bool) {
	o.BgpMeasurements = &v
}

// GetUsePublicBgp returns the UsePublicBgp field value if set, zero value otherwise.
func (o *UnexpandedAgentToAgentTest) GetUsePublicBgp() bool {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		var ret bool
		return ret
	}
	return *o.UsePublicBgp
}

// GetUsePublicBgpOk returns a tuple with the UsePublicBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedAgentToAgentTest) GetUsePublicBgpOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		return nil, false
	}
	return o.UsePublicBgp, true
}

// HasUsePublicBgp returns a boolean if a field has been set.
func (o *UnexpandedAgentToAgentTest) HasUsePublicBgp() bool {
	if o != nil && !utils.IsNil(o.UsePublicBgp) {
		return true
	}

	return false
}

// SetUsePublicBgp gets a reference to the given bool and assigns it to the UsePublicBgp field.
func (o *UnexpandedAgentToAgentTest) SetUsePublicBgp(v bool) {
	o.UsePublicBgp = &v
}

func (o UnexpandedAgentToAgentTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnexpandedAgentToAgentTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interval"] = o.Interval
	if !utils.IsNil(o.AlertsEnabled) {
		toSerialize["alertsEnabled"] = o.AlertsEnabled
	}
	if !utils.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !utils.IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !utils.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.LiveShare) {
		toSerialize["liveShare"] = o.LiveShare
	}
	if !utils.IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !utils.IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !utils.IsNil(o.SavedEvent) {
		toSerialize["savedEvent"] = o.SavedEvent
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.TestName) {
		toSerialize["testName"] = o.TestName
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !utils.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !utils.IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !utils.IsNil(o.DscpId) {
		toSerialize["dscpId"] = o.DscpId
	}
	if !utils.IsNil(o.Mss) {
		toSerialize["mss"] = o.Mss
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
	if !utils.IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !utils.IsNil(o.RandomizedStartTime) {
		toSerialize["randomizedStartTime"] = o.RandomizedStartTime
	}
	toSerialize["targetAgentId"] = o.TargetAgentId
	if !utils.IsNil(o.ThroughputMeasurements) {
		toSerialize["throughputMeasurements"] = o.ThroughputMeasurements
	}
	if !utils.IsNil(o.ThroughputDuration) {
		toSerialize["throughputDuration"] = o.ThroughputDuration
	}
	if !utils.IsNil(o.ThroughputRate) {
		toSerialize["throughputRate"] = o.ThroughputRate
	}
	if !utils.IsNil(o.FixedPacketRate) {
		toSerialize["fixedPacketRate"] = o.FixedPacketRate
	}
	if !utils.IsNil(o.BgpMeasurements) {
		toSerialize["bgpMeasurements"] = o.BgpMeasurements
	}
	if !utils.IsNil(o.UsePublicBgp) {
		toSerialize["usePublicBgp"] = o.UsePublicBgp
	}
	return toSerialize, nil
}

func (o *UnexpandedAgentToAgentTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
		"targetAgentId",
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

	varUnexpandedAgentToAgentTest := _UnexpandedAgentToAgentTest{}

    err = json.Unmarshal(data, &varUnexpandedAgentToAgentTest)

	if err != nil {
		return err
	}

	*o = UnexpandedAgentToAgentTest(varUnexpandedAgentToAgentTest)

	return err
}

type NullableUnexpandedAgentToAgentTest struct {
	value *UnexpandedAgentToAgentTest
	isSet bool
}

func (v NullableUnexpandedAgentToAgentTest) Get() *UnexpandedAgentToAgentTest {
	return v.value
}

func (v *NullableUnexpandedAgentToAgentTest) Set(val *UnexpandedAgentToAgentTest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnexpandedAgentToAgentTest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnexpandedAgentToAgentTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnexpandedAgentToAgentTest(val *UnexpandedAgentToAgentTest) *NullableUnexpandedAgentToAgentTest {
	return &NullableUnexpandedAgentToAgentTest{value: val, isSet: true}
}

func (v NullableUnexpandedAgentToAgentTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnexpandedAgentToAgentTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


