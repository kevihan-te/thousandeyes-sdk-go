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

// checks if the VoiceTestRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VoiceTestRequest{}

// VoiceTestRequest struct for VoiceTestRequest
type VoiceTestRequest struct {
	Interval TestInterval `json:"interval"`
	// Indicates if alerts are enabled.
	AlertsEnabled *bool `json:"alertsEnabled,omitempty"`
	// Test is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// List of alert rules IDs to apply to the test (get `ruleId` from `/alerts/rules` endpoint. If `alertsEnabled` is set to `true` and `alertRules` is not included on test creation or update, applicable user default alert rules will be used)
	AlertRules []string `json:"alertRules,omitempty"`
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
	// Contains list of test label IDs (get `labelId` from `/labels` endpoint)
	Labels []string `json:"labels,omitempty"`
	// Contains list of account group IDs. Test is shared with the listed account groups (get `aid` from `/account-groups` endpoint)
	SharedWithAccounts []string `json:"sharedWithAccounts,omitempty"`
	// Codec label
	Codec *string `json:"codec,omitempty"`
	// Coded ID, [see the list of acceptable values](https://docs.thousandeyes.com/product-documentation/internet-and-wan-monitoring/tests/working-with-test-settings#rtp-stream-advanced-settings-tab)
	CodecId *string `json:"codecId,omitempty"`
	// DSCP label.
	Dscp *string `json:"dscp,omitempty"`
	DscpId *TestDscpId `json:"dscpId,omitempty"`
	// Duration of the test in seconds.
	Duration *int32 `json:"duration,omitempty"`
	// De-jitter buffer size in seconds.
	JitterBuffer *int32 `json:"jitterBuffer,omitempty"`
	// Number of path traces executed by the agent.
	NumPathTraces *int32 `json:"numPathTraces,omitempty"`
	// Port number for the chosen protocol.
	Port *int32 `json:"port,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	// Agent ID of the target agent for the test.
	TargetAgentId string `json:"targetAgentId"`
	// Set to `true` to enable bgp measurements.
	BgpMeasurements *bool `json:"bgpMeasurements,omitempty"`
	// Indicate if all available public BGP monitors should be used, when ommited defaults to `bgpMeasurements` value.
	UsePublicBgp *bool `json:"usePublicBgp,omitempty"`
	// Contains list of enabled BGP monitors.
	Monitors []Monitor `json:"monitors,omitempty"`
	// Contains list of Agent IDs (get `agentId` from `/agents` endpoint).
	Agents []TestAgentRequest `json:"agents"`
}

type _VoiceTestRequest VoiceTestRequest

// NewVoiceTestRequest instantiates a new VoiceTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceTestRequest(interval TestInterval, targetAgentId string, agents []TestAgentRequest) *VoiceTestRequest {
	this := VoiceTestRequest{}
	this.Interval = interval
	var enabled bool = true
	this.Enabled = &enabled
	var dscpId TestDscpId = TESTDSCPID__0
	this.DscpId = &dscpId
	var duration int32 = 5
	this.Duration = &duration
	var jitterBuffer int32 = 40
	this.JitterBuffer = &jitterBuffer
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	this.TargetAgentId = targetAgentId
	var bgpMeasurements bool = true
	this.BgpMeasurements = &bgpMeasurements
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	this.Agents = agents
	return &this
}

// NewVoiceTestRequestWithDefaults instantiates a new VoiceTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceTestRequestWithDefaults() *VoiceTestRequest {
	this := VoiceTestRequest{}
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = interval
	var enabled bool = true
	this.Enabled = &enabled
	var dscpId TestDscpId = TESTDSCPID__0
	this.DscpId = &dscpId
	var duration int32 = 5
	this.Duration = &duration
	var jitterBuffer int32 = 40
	this.JitterBuffer = &jitterBuffer
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	var bgpMeasurements bool = true
	this.BgpMeasurements = &bgpMeasurements
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	return &this
}

// GetInterval returns the Interval field value
func (o *VoiceTestRequest) GetInterval() TestInterval {
	if o == nil {
		var ret TestInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetIntervalOk() (*TestInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *VoiceTestRequest) SetInterval(v TestInterval) {
	o.Interval = v
}

// GetAlertsEnabled returns the AlertsEnabled field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetAlertsEnabled() bool {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		var ret bool
		return ret
	}
	return *o.AlertsEnabled
}

// GetAlertsEnabledOk returns a tuple with the AlertsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetAlertsEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		return nil, false
	}
	return o.AlertsEnabled, true
}

// HasAlertsEnabled returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasAlertsEnabled() bool {
	if o != nil && !utils.IsNil(o.AlertsEnabled) {
		return true
	}

	return false
}

// SetAlertsEnabled gets a reference to the given bool and assigns it to the AlertsEnabled field.
func (o *VoiceTestRequest) SetAlertsEnabled(v bool) {
	o.AlertsEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VoiceTestRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAlertRules returns the AlertRules field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetAlertRules() []string {
	if o == nil || utils.IsNil(o.AlertRules) {
		var ret []string
		return ret
	}
	return o.AlertRules
}

// GetAlertRulesOk returns a tuple with the AlertRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetAlertRulesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlertRules) {
		return nil, false
	}
	return o.AlertRules, true
}

// HasAlertRules returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasAlertRules() bool {
	if o != nil && !utils.IsNil(o.AlertRules) {
		return true
	}

	return false
}

// SetAlertRules gets a reference to the given []string and assigns it to the AlertRules field.
func (o *VoiceTestRequest) SetAlertRules(v []string) {
	o.AlertRules = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *VoiceTestRequest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *VoiceTestRequest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VoiceTestRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *VoiceTestRequest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *VoiceTestRequest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *VoiceTestRequest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *VoiceTestRequest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *VoiceTestRequest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *VoiceTestRequest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VoiceTestRequest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *VoiceTestRequest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetLabels() []string {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetLabelsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *VoiceTestRequest) SetLabels(v []string) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetSharedWithAccounts() []string {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []string
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetSharedWithAccountsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []string and assigns it to the SharedWithAccounts field.
func (o *VoiceTestRequest) SetSharedWithAccounts(v []string) {
	o.SharedWithAccounts = v
}

// GetCodec returns the Codec field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetCodec() string {
	if o == nil || utils.IsNil(o.Codec) {
		var ret string
		return ret
	}
	return *o.Codec
}

// GetCodecOk returns a tuple with the Codec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetCodecOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Codec) {
		return nil, false
	}
	return o.Codec, true
}

// HasCodec returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasCodec() bool {
	if o != nil && !utils.IsNil(o.Codec) {
		return true
	}

	return false
}

// SetCodec gets a reference to the given string and assigns it to the Codec field.
func (o *VoiceTestRequest) SetCodec(v string) {
	o.Codec = &v
}

// GetCodecId returns the CodecId field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetCodecId() string {
	if o == nil || utils.IsNil(o.CodecId) {
		var ret string
		return ret
	}
	return *o.CodecId
}

// GetCodecIdOk returns a tuple with the CodecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetCodecIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CodecId) {
		return nil, false
	}
	return o.CodecId, true
}

// HasCodecId returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasCodecId() bool {
	if o != nil && !utils.IsNil(o.CodecId) {
		return true
	}

	return false
}

// SetCodecId gets a reference to the given string and assigns it to the CodecId field.
func (o *VoiceTestRequest) SetCodecId(v string) {
	o.CodecId = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetDscp() string {
	if o == nil || utils.IsNil(o.Dscp) {
		var ret string
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetDscpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasDscp() bool {
	if o != nil && !utils.IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given string and assigns it to the Dscp field.
func (o *VoiceTestRequest) SetDscp(v string) {
	o.Dscp = &v
}

// GetDscpId returns the DscpId field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetDscpId() TestDscpId {
	if o == nil || utils.IsNil(o.DscpId) {
		var ret TestDscpId
		return ret
	}
	return *o.DscpId
}

// GetDscpIdOk returns a tuple with the DscpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetDscpIdOk() (*TestDscpId, bool) {
	if o == nil || utils.IsNil(o.DscpId) {
		return nil, false
	}
	return o.DscpId, true
}

// HasDscpId returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasDscpId() bool {
	if o != nil && !utils.IsNil(o.DscpId) {
		return true
	}

	return false
}

// SetDscpId gets a reference to the given TestDscpId and assigns it to the DscpId field.
func (o *VoiceTestRequest) SetDscpId(v TestDscpId) {
	o.DscpId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetDuration() int32 {
	if o == nil || utils.IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetDurationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *VoiceTestRequest) SetDuration(v int32) {
	o.Duration = &v
}

// GetJitterBuffer returns the JitterBuffer field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetJitterBuffer() int32 {
	if o == nil || utils.IsNil(o.JitterBuffer) {
		var ret int32
		return ret
	}
	return *o.JitterBuffer
}

// GetJitterBufferOk returns a tuple with the JitterBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetJitterBufferOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.JitterBuffer) {
		return nil, false
	}
	return o.JitterBuffer, true
}

// HasJitterBuffer returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasJitterBuffer() bool {
	if o != nil && !utils.IsNil(o.JitterBuffer) {
		return true
	}

	return false
}

// SetJitterBuffer gets a reference to the given int32 and assigns it to the JitterBuffer field.
func (o *VoiceTestRequest) SetJitterBuffer(v int32) {
	o.JitterBuffer = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *VoiceTestRequest) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *VoiceTestRequest) SetPort(v int32) {
	o.Port = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *VoiceTestRequest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetTargetAgentId returns the TargetAgentId field value
func (o *VoiceTestRequest) GetTargetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAgentId
}

// GetTargetAgentIdOk returns a tuple with the TargetAgentId field value
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetTargetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAgentId, true
}

// SetTargetAgentId sets field value
func (o *VoiceTestRequest) SetTargetAgentId(v string) {
	o.TargetAgentId = v
}

// GetBgpMeasurements returns the BgpMeasurements field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetBgpMeasurements() bool {
	if o == nil || utils.IsNil(o.BgpMeasurements) {
		var ret bool
		return ret
	}
	return *o.BgpMeasurements
}

// GetBgpMeasurementsOk returns a tuple with the BgpMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetBgpMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BgpMeasurements) {
		return nil, false
	}
	return o.BgpMeasurements, true
}

// HasBgpMeasurements returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasBgpMeasurements() bool {
	if o != nil && !utils.IsNil(o.BgpMeasurements) {
		return true
	}

	return false
}

// SetBgpMeasurements gets a reference to the given bool and assigns it to the BgpMeasurements field.
func (o *VoiceTestRequest) SetBgpMeasurements(v bool) {
	o.BgpMeasurements = &v
}

// GetUsePublicBgp returns the UsePublicBgp field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetUsePublicBgp() bool {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		var ret bool
		return ret
	}
	return *o.UsePublicBgp
}

// GetUsePublicBgpOk returns a tuple with the UsePublicBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetUsePublicBgpOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		return nil, false
	}
	return o.UsePublicBgp, true
}

// HasUsePublicBgp returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasUsePublicBgp() bool {
	if o != nil && !utils.IsNil(o.UsePublicBgp) {
		return true
	}

	return false
}

// SetUsePublicBgp gets a reference to the given bool and assigns it to the UsePublicBgp field.
func (o *VoiceTestRequest) SetUsePublicBgp(v bool) {
	o.UsePublicBgp = &v
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *VoiceTestRequest) GetMonitors() []Monitor {
	if o == nil || utils.IsNil(o.Monitors) {
		var ret []Monitor
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetMonitorsOk() ([]Monitor, bool) {
	if o == nil || utils.IsNil(o.Monitors) {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *VoiceTestRequest) HasMonitors() bool {
	if o != nil && !utils.IsNil(o.Monitors) {
		return true
	}

	return false
}

// SetMonitors gets a reference to the given []Monitor and assigns it to the Monitors field.
func (o *VoiceTestRequest) SetMonitors(v []Monitor) {
	o.Monitors = v
}

// GetAgents returns the Agents field value
func (o *VoiceTestRequest) GetAgents() []TestAgentRequest {
	if o == nil {
		var ret []TestAgentRequest
		return ret
	}

	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value
// and a boolean to check if the value has been set.
func (o *VoiceTestRequest) GetAgentsOk() ([]TestAgentRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Agents, true
}

// SetAgents sets field value
func (o *VoiceTestRequest) SetAgents(v []TestAgentRequest) {
	o.Agents = v
}

func (o VoiceTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interval"] = o.Interval
	if !utils.IsNil(o.AlertsEnabled) {
		toSerialize["alertsEnabled"] = o.AlertsEnabled
	}
	if !utils.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !utils.IsNil(o.AlertRules) {
		toSerialize["alertRules"] = o.AlertRules
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
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !utils.IsNil(o.SharedWithAccounts) {
		toSerialize["sharedWithAccounts"] = o.SharedWithAccounts
	}
	if !utils.IsNil(o.Codec) {
		toSerialize["codec"] = o.Codec
	}
	if !utils.IsNil(o.CodecId) {
		toSerialize["codecId"] = o.CodecId
	}
	if !utils.IsNil(o.Dscp) {
		toSerialize["dscp"] = o.Dscp
	}
	if !utils.IsNil(o.DscpId) {
		toSerialize["dscpId"] = o.DscpId
	}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !utils.IsNil(o.JitterBuffer) {
		toSerialize["jitterBuffer"] = o.JitterBuffer
	}
	if !utils.IsNil(o.NumPathTraces) {
		toSerialize["numPathTraces"] = o.NumPathTraces
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.RandomizedStartTime) {
		toSerialize["randomizedStartTime"] = o.RandomizedStartTime
	}
	toSerialize["targetAgentId"] = o.TargetAgentId
	if !utils.IsNil(o.BgpMeasurements) {
		toSerialize["bgpMeasurements"] = o.BgpMeasurements
	}
	if !utils.IsNil(o.UsePublicBgp) {
		toSerialize["usePublicBgp"] = o.UsePublicBgp
	}
	if !utils.IsNil(o.Monitors) {
		toSerialize["monitors"] = o.Monitors
	}
	toSerialize["agents"] = o.Agents
	return toSerialize, nil
}

func (o *VoiceTestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
		"targetAgentId",
		"agents",
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

	varVoiceTestRequest := _VoiceTestRequest{}

    err = json.Unmarshal(data, &varVoiceTestRequest)

	if err != nil {
		return err
	}

	*o = VoiceTestRequest(varVoiceTestRequest)

	return err
}

type NullableVoiceTestRequest struct {
	value *VoiceTestRequest
	isSet bool
}

func (v NullableVoiceTestRequest) Get() *VoiceTestRequest {
	return v.value
}

func (v *NullableVoiceTestRequest) Set(val *VoiceTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceTestRequest(val *VoiceTestRequest) *NullableVoiceTestRequest {
	return &NullableVoiceTestRequest{value: val, isSet: true}
}

func (v NullableVoiceTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


