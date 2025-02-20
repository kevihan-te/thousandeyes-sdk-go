/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
	"fmt"
)

// checks if the VoiceInstantTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VoiceInstantTest{}

// VoiceInstantTest struct for VoiceInstantTest
type VoiceInstantTest struct {
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
}

type _VoiceInstantTest VoiceInstantTest

// NewVoiceInstantTest instantiates a new VoiceInstantTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoiceInstantTest(targetAgentId string) *VoiceInstantTest {
	this := VoiceInstantTest{}
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
	return &this
}

// NewVoiceInstantTestWithDefaults instantiates a new VoiceInstantTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoiceInstantTestWithDefaults() *VoiceInstantTest {
	this := VoiceInstantTest{}
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
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *VoiceInstantTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *VoiceInstantTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VoiceInstantTest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *VoiceInstantTest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *VoiceInstantTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *VoiceInstantTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *VoiceInstantTest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *VoiceInstantTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *VoiceInstantTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VoiceInstantTest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *VoiceInstantTest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *VoiceInstantTest) SetLabels(v []TestLabel) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetSharedWithAccounts() []SharedWithAccount {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []SharedWithAccount
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetSharedWithAccountsOk() ([]SharedWithAccount, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []SharedWithAccount and assigns it to the SharedWithAccounts field.
func (o *VoiceInstantTest) SetSharedWithAccounts(v []SharedWithAccount) {
	o.SharedWithAccounts = v
}

// GetCodec returns the Codec field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetCodec() string {
	if o == nil || utils.IsNil(o.Codec) {
		var ret string
		return ret
	}
	return *o.Codec
}

// GetCodecOk returns a tuple with the Codec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetCodecOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Codec) {
		return nil, false
	}
	return o.Codec, true
}

// HasCodec returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasCodec() bool {
	if o != nil && !utils.IsNil(o.Codec) {
		return true
	}

	return false
}

// SetCodec gets a reference to the given string and assigns it to the Codec field.
func (o *VoiceInstantTest) SetCodec(v string) {
	o.Codec = &v
}

// GetCodecId returns the CodecId field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetCodecId() string {
	if o == nil || utils.IsNil(o.CodecId) {
		var ret string
		return ret
	}
	return *o.CodecId
}

// GetCodecIdOk returns a tuple with the CodecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetCodecIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CodecId) {
		return nil, false
	}
	return o.CodecId, true
}

// HasCodecId returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasCodecId() bool {
	if o != nil && !utils.IsNil(o.CodecId) {
		return true
	}

	return false
}

// SetCodecId gets a reference to the given string and assigns it to the CodecId field.
func (o *VoiceInstantTest) SetCodecId(v string) {
	o.CodecId = &v
}

// GetDscp returns the Dscp field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetDscp() string {
	if o == nil || utils.IsNil(o.Dscp) {
		var ret string
		return ret
	}
	return *o.Dscp
}

// GetDscpOk returns a tuple with the Dscp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetDscpOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Dscp) {
		return nil, false
	}
	return o.Dscp, true
}

// HasDscp returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasDscp() bool {
	if o != nil && !utils.IsNil(o.Dscp) {
		return true
	}

	return false
}

// SetDscp gets a reference to the given string and assigns it to the Dscp field.
func (o *VoiceInstantTest) SetDscp(v string) {
	o.Dscp = &v
}

// GetDscpId returns the DscpId field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetDscpId() TestDscpId {
	if o == nil || utils.IsNil(o.DscpId) {
		var ret TestDscpId
		return ret
	}
	return *o.DscpId
}

// GetDscpIdOk returns a tuple with the DscpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetDscpIdOk() (*TestDscpId, bool) {
	if o == nil || utils.IsNil(o.DscpId) {
		return nil, false
	}
	return o.DscpId, true
}

// HasDscpId returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasDscpId() bool {
	if o != nil && !utils.IsNil(o.DscpId) {
		return true
	}

	return false
}

// SetDscpId gets a reference to the given TestDscpId and assigns it to the DscpId field.
func (o *VoiceInstantTest) SetDscpId(v TestDscpId) {
	o.DscpId = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetDuration() int32 {
	if o == nil || utils.IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetDurationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *VoiceInstantTest) SetDuration(v int32) {
	o.Duration = &v
}

// GetJitterBuffer returns the JitterBuffer field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetJitterBuffer() int32 {
	if o == nil || utils.IsNil(o.JitterBuffer) {
		var ret int32
		return ret
	}
	return *o.JitterBuffer
}

// GetJitterBufferOk returns a tuple with the JitterBuffer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetJitterBufferOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.JitterBuffer) {
		return nil, false
	}
	return o.JitterBuffer, true
}

// HasJitterBuffer returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasJitterBuffer() bool {
	if o != nil && !utils.IsNil(o.JitterBuffer) {
		return true
	}

	return false
}

// SetJitterBuffer gets a reference to the given int32 and assigns it to the JitterBuffer field.
func (o *VoiceInstantTest) SetJitterBuffer(v int32) {
	o.JitterBuffer = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *VoiceInstantTest) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *VoiceInstantTest) SetPort(v int32) {
	o.Port = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *VoiceInstantTest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *VoiceInstantTest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *VoiceInstantTest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetTargetAgentId returns the TargetAgentId field value
func (o *VoiceInstantTest) GetTargetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetAgentId
}

// GetTargetAgentIdOk returns a tuple with the TargetAgentId field value
// and a boolean to check if the value has been set.
func (o *VoiceInstantTest) GetTargetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetAgentId, true
}

// SetTargetAgentId sets field value
func (o *VoiceInstantTest) SetTargetAgentId(v string) {
	o.TargetAgentId = v
}

func (o VoiceInstantTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoiceInstantTest) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *VoiceInstantTest) UnmarshalJSON(data []byte) (err error) {
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

	varVoiceInstantTest := _VoiceInstantTest{}

    err = json.Unmarshal(data, &varVoiceInstantTest)

	if err != nil {
		return err
	}

	*o = VoiceInstantTest(varVoiceInstantTest)

	return err
}

type NullableVoiceInstantTest struct {
	value *VoiceInstantTest
	isSet bool
}

func (v NullableVoiceInstantTest) Get() *VoiceInstantTest {
	return v.value
}

func (v *NullableVoiceInstantTest) Set(val *VoiceInstantTest) {
	v.value = val
	v.isSet = true
}

func (v NullableVoiceInstantTest) IsSet() bool {
	return v.isSet
}

func (v *NullableVoiceInstantTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoiceInstantTest(val *VoiceInstantTest) *NullableVoiceInstantTest {
	return &NullableVoiceInstantTest{value: val, isSet: true}
}

func (v NullableVoiceInstantTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoiceInstantTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


