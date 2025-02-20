/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
	"fmt"
)

// checks if the AgentToAgentInstantTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentToAgentInstantTest{}

// AgentToAgentInstantTest struct for AgentToAgentInstantTest
type AgentToAgentInstantTest struct {
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
	Labels []TestLabel `json:"labels,omitempty"`
	SharedWithAccounts []SharedWithAccount `json:"sharedWithAccounts,omitempty"`
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
}

type _AgentToAgentInstantTest AgentToAgentInstantTest

// NewAgentToAgentInstantTest instantiates a new AgentToAgentInstantTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentToAgentInstantTest(targetAgentId string) *AgentToAgentInstantTest {
	this := AgentToAgentInstantTest{}
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
	return &this
}

// NewAgentToAgentInstantTestWithDefaults instantiates a new AgentToAgentInstantTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentToAgentInstantTestWithDefaults() *AgentToAgentInstantTest {
	this := AgentToAgentInstantTest{}
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
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *AgentToAgentInstantTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *AgentToAgentInstantTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AgentToAgentInstantTest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *AgentToAgentInstantTest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *AgentToAgentInstantTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *AgentToAgentInstantTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *AgentToAgentInstantTest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *AgentToAgentInstantTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *AgentToAgentInstantTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AgentToAgentInstantTest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *AgentToAgentInstantTest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *AgentToAgentInstantTest) SetLabels(v []TestLabel) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetSharedWithAccounts() []SharedWithAccount {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []SharedWithAccount
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetSharedWithAccountsOk() ([]SharedWithAccount, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []SharedWithAccount and assigns it to the SharedWithAccounts field.
func (o *AgentToAgentInstantTest) SetSharedWithAccounts(v []SharedWithAccount) {
	o.SharedWithAccounts = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetDirection() TestDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret TestDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetDirectionOk() (*TestDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given TestDirection and assigns it to the Direction field.
func (o *AgentToAgentInstantTest) SetDirection(v TestDirection) {
	o.Direction = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetDscp() string {
	if o == nil || utils.IsNil(o.Dscp) {
		var ret string
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetDscpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasDscp() bool {
	if o != nil && !utils.IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given string and assigns it to the Dscp field.
func (o *AgentToAgentInstantTest) SetDscp(v string) {
	o.Dscp = &v
}

// GetDscpId returns the DscpId field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetDscpId() TestDscpId {
	if o == nil || utils.IsNil(o.DscpId) {
		var ret TestDscpId
		return ret
	}
	return *o.DscpId
}

// GetDscpIdOk returns a tuple with the DscpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetDscpIdOk() (*TestDscpId, bool) {
	if o == nil || utils.IsNil(o.DscpId) {
		return nil, false
	}
	return o.DscpId, true
}

// HasDscpId returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasDscpId() bool {
	if o != nil && !utils.IsNil(o.DscpId) {
		return true
	}

	return false
}

// SetDscpId gets a reference to the given TestDscpId and assigns it to the DscpId field.
func (o *AgentToAgentInstantTest) SetDscpId(v TestDscpId) {
	o.DscpId = &v
}

// GetMss returns the Mss field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetMss() int32 {
	if o == nil || utils.IsNil(o.Mss) {
		var ret int32
		return ret
	}
	return *o.Mss
}

// GetMssOk returns a tuple with the Mss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetMssOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Mss) {
		return nil, false
	}
	return o.Mss, true
}

// HasMss returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasMss() bool {
	if o != nil && !utils.IsNil(o.Mss) {
		return true
	}

	return false
}

// SetMss gets a reference to the given int32 and assigns it to the Mss field.
func (o *AgentToAgentInstantTest) SetMss(v int32) {
	o.Mss = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *AgentToAgentInstantTest) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPathTraceMode returns the PathTraceMode field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetPathTraceMode() TestPathTraceMode {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		var ret TestPathTraceMode
		return ret
	}
	return *o.PathTraceMode
}

// GetPathTraceModeOk returns a tuple with the PathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetPathTraceModeOk() (*TestPathTraceMode, bool) {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		return nil, false
	}
	return o.PathTraceMode, true
}

// HasPathTraceMode returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.PathTraceMode) {
		return true
	}

	return false
}

// SetPathTraceMode gets a reference to the given TestPathTraceMode and assigns it to the PathTraceMode field.
func (o *AgentToAgentInstantTest) SetPathTraceMode(v TestPathTraceMode) {
	o.PathTraceMode = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *AgentToAgentInstantTest) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetProtocol() AgentToAgentTestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret AgentToAgentTestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetProtocolOk() (*AgentToAgentTestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given AgentToAgentTestProtocol and assigns it to the Protocol field.
func (o *AgentToAgentInstantTest) SetProtocol(v AgentToAgentTestProtocol) {
	o.Protocol = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *AgentToAgentInstantTest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetTargetAgentId returns the TargetAgentId field value
func (o *AgentToAgentInstantTest) GetTargetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAgentId
}

// GetTargetAgentIdOk returns a tuple with the TargetAgentId field value
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetTargetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAgentId, true
}

// SetTargetAgentId sets field value
func (o *AgentToAgentInstantTest) SetTargetAgentId(v string) {
	o.TargetAgentId = v
}

// GetThroughputMeasurements returns the ThroughputMeasurements field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetThroughputMeasurements() bool {
	if o == nil || utils.IsNil(o.ThroughputMeasurements) {
		var ret bool
		return ret
	}
	return *o.ThroughputMeasurements
}

// GetThroughputMeasurementsOk returns a tuple with the ThroughputMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetThroughputMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ThroughputMeasurements) {
		return nil, false
	}
	return o.ThroughputMeasurements, true
}

// HasThroughputMeasurements returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasThroughputMeasurements() bool {
	if o != nil && !utils.IsNil(o.ThroughputMeasurements) {
		return true
	}

	return false
}

// SetThroughputMeasurements gets a reference to the given bool and assigns it to the ThroughputMeasurements field.
func (o *AgentToAgentInstantTest) SetThroughputMeasurements(v bool) {
	o.ThroughputMeasurements = &v
}

// GetThroughputDuration returns the ThroughputDuration field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetThroughputDuration() int32 {
	if o == nil || utils.IsNil(o.ThroughputDuration) {
		var ret int32
		return ret
	}
	return *o.ThroughputDuration
}

// GetThroughputDurationOk returns a tuple with the ThroughputDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetThroughputDurationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ThroughputDuration) {
		return nil, false
	}
	return o.ThroughputDuration, true
}

// HasThroughputDuration returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasThroughputDuration() bool {
	if o != nil && !utils.IsNil(o.ThroughputDuration) {
		return true
	}

	return false
}

// SetThroughputDuration gets a reference to the given int32 and assigns it to the ThroughputDuration field.
func (o *AgentToAgentInstantTest) SetThroughputDuration(v int32) {
	o.ThroughputDuration = &v
}

// GetThroughputRate returns the ThroughputRate field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetThroughputRate() int32 {
	if o == nil || utils.IsNil(o.ThroughputRate) {
		var ret int32
		return ret
	}
	return *o.ThroughputRate
}

// GetThroughputRateOk returns a tuple with the ThroughputRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetThroughputRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ThroughputRate) {
		return nil, false
	}
	return o.ThroughputRate, true
}

// HasThroughputRate returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasThroughputRate() bool {
	if o != nil && !utils.IsNil(o.ThroughputRate) {
		return true
	}

	return false
}

// SetThroughputRate gets a reference to the given int32 and assigns it to the ThroughputRate field.
func (o *AgentToAgentInstantTest) SetThroughputRate(v int32) {
	o.ThroughputRate = &v
}

// GetFixedPacketRate returns the FixedPacketRate field value if set, zero value otherwise.
func (o *AgentToAgentInstantTest) GetFixedPacketRate() int32 {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		var ret int32
		return ret
	}
	return *o.FixedPacketRate
}

// GetFixedPacketRateOk returns a tuple with the FixedPacketRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentToAgentInstantTest) GetFixedPacketRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		return nil, false
	}
	return o.FixedPacketRate, true
}

// HasFixedPacketRate returns a boolean if a field has been set.
func (o *AgentToAgentInstantTest) HasFixedPacketRate() bool {
	if o != nil && !utils.IsNil(o.FixedPacketRate) {
		return true
	}

	return false
}

// SetFixedPacketRate gets a reference to the given int32 and assigns it to the FixedPacketRate field.
func (o *AgentToAgentInstantTest) SetFixedPacketRate(v int32) {
	o.FixedPacketRate = &v
}

func (o AgentToAgentInstantTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentToAgentInstantTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !utils.IsNil(o.SharedWithAccounts) {
		toSerialize["sharedWithAccounts"] = o.SharedWithAccounts
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
	return toSerialize, nil
}

func (o *AgentToAgentInstantTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varAgentToAgentInstantTest := _AgentToAgentInstantTest{}

    err = json.Unmarshal(data, &varAgentToAgentInstantTest)

	if err != nil {
		return err
	}

	*o = AgentToAgentInstantTest(varAgentToAgentInstantTest)

	return err
}

type NullableAgentToAgentInstantTest struct {
	value *AgentToAgentInstantTest
	isSet bool
}

func (v NullableAgentToAgentInstantTest) Get() *AgentToAgentInstantTest {
	return v.value
}

func (v *NullableAgentToAgentInstantTest) Set(val *AgentToAgentInstantTest) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentToAgentInstantTest) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentToAgentInstantTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentToAgentInstantTest(val *AgentToAgentInstantTest) *NullableAgentToAgentInstantTest {
	return &NullableAgentToAgentInstantTest{value: val, isSet: true}
}

func (v NullableAgentToAgentInstantTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentToAgentInstantTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


