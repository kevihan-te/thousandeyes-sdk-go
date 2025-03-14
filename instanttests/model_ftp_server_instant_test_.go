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
	"time"
	"fmt"
)

// checks if the FtpServerInstantTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &FtpServerInstantTest{}

// FtpServerInstantTest struct for FtpServerInstantTest
type FtpServerInstantTest struct {
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
}

type _FtpServerInstantTest FtpServerInstantTest

// NewFtpServerInstantTest instantiates a new FtpServerInstantTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFtpServerInstantTest(password string, requestType FtpServerRequestType, url string, username string) *FtpServerInstantTest {
	this := FtpServerInstantTest{}
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

// NewFtpServerInstantTestWithDefaults instantiates a new FtpServerInstantTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFtpServerInstantTestWithDefaults() *FtpServerInstantTest {
	this := FtpServerInstantTest{}
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

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *FtpServerInstantTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *FtpServerInstantTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FtpServerInstantTest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *FtpServerInstantTest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *FtpServerInstantTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *FtpServerInstantTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *FtpServerInstantTest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *FtpServerInstantTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *FtpServerInstantTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FtpServerInstantTest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *FtpServerInstantTest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *FtpServerInstantTest) SetLabels(v []TestLabel) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetSharedWithAccounts() []SharedWithAccount {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []SharedWithAccount
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetSharedWithAccountsOk() ([]SharedWithAccount, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []SharedWithAccount and assigns it to the SharedWithAccounts field.
func (o *FtpServerInstantTest) SetSharedWithAccounts(v []SharedWithAccount) {
	o.SharedWithAccounts = v
}

// GetBandwidthMeasurements returns the BandwidthMeasurements field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetBandwidthMeasurements() bool {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		var ret bool
		return ret
	}
	return *o.BandwidthMeasurements
}

// GetBandwidthMeasurementsOk returns a tuple with the BandwidthMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetBandwidthMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		return nil, false
	}
	return o.BandwidthMeasurements, true
}

// HasBandwidthMeasurements returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasBandwidthMeasurements() bool {
	if o != nil && !utils.IsNil(o.BandwidthMeasurements) {
		return true
	}

	return false
}

// SetBandwidthMeasurements gets a reference to the given bool and assigns it to the BandwidthMeasurements field.
func (o *FtpServerInstantTest) SetBandwidthMeasurements(v bool) {
	o.BandwidthMeasurements = &v
}

// GetDownloadLimit returns the DownloadLimit field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetDownloadLimit() int32 {
	if o == nil || utils.IsNil(o.DownloadLimit) {
		var ret int32
		return ret
	}
	return *o.DownloadLimit
}

// GetDownloadLimitOk returns a tuple with the DownloadLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetDownloadLimitOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.DownloadLimit) {
		return nil, false
	}
	return o.DownloadLimit, true
}

// HasDownloadLimit returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasDownloadLimit() bool {
	if o != nil && !utils.IsNil(o.DownloadLimit) {
		return true
	}

	return false
}

// SetDownloadLimit gets a reference to the given int32 and assigns it to the DownloadLimit field.
func (o *FtpServerInstantTest) SetDownloadLimit(v int32) {
	o.DownloadLimit = &v
}

// GetFtpTargetTime returns the FtpTargetTime field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetFtpTargetTime() int32 {
	if o == nil || utils.IsNil(o.FtpTargetTime) {
		var ret int32
		return ret
	}
	return *o.FtpTargetTime
}

// GetFtpTargetTimeOk returns a tuple with the FtpTargetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetFtpTargetTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FtpTargetTime) {
		return nil, false
	}
	return o.FtpTargetTime, true
}

// HasFtpTargetTime returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasFtpTargetTime() bool {
	if o != nil && !utils.IsNil(o.FtpTargetTime) {
		return true
	}

	return false
}

// SetFtpTargetTime gets a reference to the given int32 and assigns it to the FtpTargetTime field.
func (o *FtpServerInstantTest) SetFtpTargetTime(v int32) {
	o.FtpTargetTime = &v
}

// GetFtpTimeLimit returns the FtpTimeLimit field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetFtpTimeLimit() int32 {
	if o == nil || utils.IsNil(o.FtpTimeLimit) {
		var ret int32
		return ret
	}
	return *o.FtpTimeLimit
}

// GetFtpTimeLimitOk returns a tuple with the FtpTimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetFtpTimeLimitOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FtpTimeLimit) {
		return nil, false
	}
	return o.FtpTimeLimit, true
}

// HasFtpTimeLimit returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasFtpTimeLimit() bool {
	if o != nil && !utils.IsNil(o.FtpTimeLimit) {
		return true
	}

	return false
}

// SetFtpTimeLimit gets a reference to the given int32 and assigns it to the FtpTimeLimit field.
func (o *FtpServerInstantTest) SetFtpTimeLimit(v int32) {
	o.FtpTimeLimit = &v
}

// GetMtuMeasurements returns the MtuMeasurements field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetMtuMeasurements() bool {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		var ret bool
		return ret
	}
	return *o.MtuMeasurements
}

// GetMtuMeasurementsOk returns a tuple with the MtuMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetMtuMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		return nil, false
	}
	return o.MtuMeasurements, true
}

// HasMtuMeasurements returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasMtuMeasurements() bool {
	if o != nil && !utils.IsNil(o.MtuMeasurements) {
		return true
	}

	return false
}

// SetMtuMeasurements gets a reference to the given bool and assigns it to the MtuMeasurements field.
func (o *FtpServerInstantTest) SetMtuMeasurements(v bool) {
	o.MtuMeasurements = &v
}

// GetNetworkMeasurements returns the NetworkMeasurements field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetNetworkMeasurements() bool {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		var ret bool
		return ret
	}
	return *o.NetworkMeasurements
}

// GetNetworkMeasurementsOk returns a tuple with the NetworkMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetNetworkMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		return nil, false
	}
	return o.NetworkMeasurements, true
}

// HasNetworkMeasurements returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasNetworkMeasurements() bool {
	if o != nil && !utils.IsNil(o.NetworkMeasurements) {
		return true
	}

	return false
}

// SetNetworkMeasurements gets a reference to the given bool and assigns it to the NetworkMeasurements field.
func (o *FtpServerInstantTest) SetNetworkMeasurements(v bool) {
	o.NetworkMeasurements = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *FtpServerInstantTest) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPassword returns the Password field value
func (o *FtpServerInstantTest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *FtpServerInstantTest) SetPassword(v string) {
	o.Password = v
}

// GetPathTraceMode returns the PathTraceMode field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetPathTraceMode() TestPathTraceMode {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		var ret TestPathTraceMode
		return ret
	}
	return *o.PathTraceMode
}

// GetPathTraceModeOk returns a tuple with the PathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetPathTraceModeOk() (*TestPathTraceMode, bool) {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		return nil, false
	}
	return o.PathTraceMode, true
}

// HasPathTraceMode returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.PathTraceMode) {
		return true
	}

	return false
}

// SetPathTraceMode gets a reference to the given TestPathTraceMode and assigns it to the PathTraceMode field.
func (o *FtpServerInstantTest) SetPathTraceMode(v TestPathTraceMode) {
	o.PathTraceMode = &v
}

// GetProbeMode returns the ProbeMode field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetProbeMode() TestProbeMode {
	if o == nil || utils.IsNil(o.ProbeMode) {
		var ret TestProbeMode
		return ret
	}
	return *o.ProbeMode
}

// GetProbeModeOk returns a tuple with the ProbeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetProbeModeOk() (*TestProbeMode, bool) {
	if o == nil || utils.IsNil(o.ProbeMode) {
		return nil, false
	}
	return o.ProbeMode, true
}

// HasProbeMode returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasProbeMode() bool {
	if o != nil && !utils.IsNil(o.ProbeMode) {
		return true
	}

	return false
}

// SetProbeMode gets a reference to the given TestProbeMode and assigns it to the ProbeMode field.
func (o *FtpServerInstantTest) SetProbeMode(v TestProbeMode) {
	o.ProbeMode = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetProtocol() TestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret TestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetProtocolOk() (*TestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given TestProtocol and assigns it to the Protocol field.
func (o *FtpServerInstantTest) SetProtocol(v TestProtocol) {
	o.Protocol = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *FtpServerInstantTest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetRequestType returns the RequestType field value
func (o *FtpServerInstantTest) GetRequestType() FtpServerRequestType {
	if o == nil {
		var ret FtpServerRequestType
		return ret
	}

	return o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetRequestTypeOk() (*FtpServerRequestType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestType, true
}

// SetRequestType sets field value
func (o *FtpServerInstantTest) SetRequestType(v FtpServerRequestType) {
	o.RequestType = v
}

// GetUrl returns the Url field value
func (o *FtpServerInstantTest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *FtpServerInstantTest) SetUrl(v string) {
	o.Url = v
}

// GetUseActiveFtp returns the UseActiveFtp field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetUseActiveFtp() bool {
	if o == nil || utils.IsNil(o.UseActiveFtp) {
		var ret bool
		return ret
	}
	return *o.UseActiveFtp
}

// GetUseActiveFtpOk returns a tuple with the UseActiveFtp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetUseActiveFtpOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UseActiveFtp) {
		return nil, false
	}
	return o.UseActiveFtp, true
}

// HasUseActiveFtp returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasUseActiveFtp() bool {
	if o != nil && !utils.IsNil(o.UseActiveFtp) {
		return true
	}

	return false
}

// SetUseActiveFtp gets a reference to the given bool and assigns it to the UseActiveFtp field.
func (o *FtpServerInstantTest) SetUseActiveFtp(v bool) {
	o.UseActiveFtp = &v
}

// GetUseExplicitFtps returns the UseExplicitFtps field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetUseExplicitFtps() bool {
	if o == nil || utils.IsNil(o.UseExplicitFtps) {
		var ret bool
		return ret
	}
	return *o.UseExplicitFtps
}

// GetUseExplicitFtpsOk returns a tuple with the UseExplicitFtps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetUseExplicitFtpsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UseExplicitFtps) {
		return nil, false
	}
	return o.UseExplicitFtps, true
}

// HasUseExplicitFtps returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasUseExplicitFtps() bool {
	if o != nil && !utils.IsNil(o.UseExplicitFtps) {
		return true
	}

	return false
}

// SetUseExplicitFtps gets a reference to the given bool and assigns it to the UseExplicitFtps field.
func (o *FtpServerInstantTest) SetUseExplicitFtps(v bool) {
	o.UseExplicitFtps = &v
}

// GetUsername returns the Username field value
func (o *FtpServerInstantTest) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *FtpServerInstantTest) SetUsername(v string) {
	o.Username = v
}

// GetFixedPacketRate returns the FixedPacketRate field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetFixedPacketRate() int32 {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		var ret int32
		return ret
	}
	return *o.FixedPacketRate
}

// GetFixedPacketRateOk returns a tuple with the FixedPacketRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetFixedPacketRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		return nil, false
	}
	return o.FixedPacketRate, true
}

// HasFixedPacketRate returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasFixedPacketRate() bool {
	if o != nil && !utils.IsNil(o.FixedPacketRate) {
		return true
	}

	return false
}

// SetFixedPacketRate gets a reference to the given int32 and assigns it to the FixedPacketRate field.
func (o *FtpServerInstantTest) SetFixedPacketRate(v int32) {
	o.FixedPacketRate = &v
}

// GetIpv6Policy returns the Ipv6Policy field value if set, zero value otherwise.
func (o *FtpServerInstantTest) GetIpv6Policy() TestIpv6Policy {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		var ret TestIpv6Policy
		return ret
	}
	return *o.Ipv6Policy
}

// GetIpv6PolicyOk returns a tuple with the Ipv6Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FtpServerInstantTest) GetIpv6PolicyOk() (*TestIpv6Policy, bool) {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		return nil, false
	}
	return o.Ipv6Policy, true
}

// HasIpv6Policy returns a boolean if a field has been set.
func (o *FtpServerInstantTest) HasIpv6Policy() bool {
	if o != nil && !utils.IsNil(o.Ipv6Policy) {
		return true
	}

	return false
}

// SetIpv6Policy gets a reference to the given TestIpv6Policy and assigns it to the Ipv6Policy field.
func (o *FtpServerInstantTest) SetIpv6Policy(v TestIpv6Policy) {
	o.Ipv6Policy = &v
}

func (o FtpServerInstantTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FtpServerInstantTest) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *FtpServerInstantTest) UnmarshalJSON(data []byte) (err error) {
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

	varFtpServerInstantTest := _FtpServerInstantTest{}

    err = json.Unmarshal(data, &varFtpServerInstantTest)

	if err != nil {
		return err
	}

	*o = FtpServerInstantTest(varFtpServerInstantTest)

	return err
}

type NullableFtpServerInstantTest struct {
	value *FtpServerInstantTest
	isSet bool
}

func (v NullableFtpServerInstantTest) Get() *FtpServerInstantTest {
	return v.value
}

func (v *NullableFtpServerInstantTest) Set(val *FtpServerInstantTest) {
	v.value = val
	v.isSet = true
}

func (v NullableFtpServerInstantTest) IsSet() bool {
	return v.isSet
}

func (v *NullableFtpServerInstantTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFtpServerInstantTest(val *FtpServerInstantTest) *NullableFtpServerInstantTest {
	return &NullableFtpServerInstantTest{value: val, isSet: true}
}

func (v NullableFtpServerInstantTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFtpServerInstantTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


