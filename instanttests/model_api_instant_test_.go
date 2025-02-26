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

// checks if the ApiInstantTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiInstantTest{}

// ApiInstantTest struct for ApiInstantTest
type ApiInstantTest struct {
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
	// Indicates whether network data to the proxy should be collected.
	CollectProxyNetworkData *bool `json:"collectProxyNetworkData,omitempty"`
	// To disable following HTTP/301 or HTTP/302 redirect directives, set this parameter to `false`.
	FollowRedirects *bool `json:"followRedirects,omitempty"`
	// Set `true` to measure MTU sizes on network from agents to the target.
	MtuMeasurements *bool `json:"mtuMeasurements,omitempty"`
	// Enable or disable network measurements. Set to true to enable or false to disable network measurements.
	NetworkMeasurements *bool `json:"networkMeasurements,omitempty"`
	// Number of path traces executed by the agent.
	NumPathTraces *int32 `json:"numPathTraces,omitempty"`
	// Flag indicating if a proxy other than the default should be used. To override the default proxy for agents, set to `true` and specify a value for `overrideProxyId`.
	OverrideAgentProxy *bool `json:"overrideAgentProxy,omitempty"`
	// ID of the proxy to be used if the default proxy is overridden.
	OverrideProxyId *string `json:"overrideProxyId,omitempty"`
	PathTraceMode *TestPathTraceMode `json:"pathTraceMode,omitempty"`
	PredefinedVariables []ApiPredefinedVariable `json:"predefinedVariables,omitempty"`
	ProbeMode *TestProbeMode `json:"probeMode,omitempty"`
	Protocol *TestProtocol `json:"protocol,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	Requests []ApiRequest `json:"requests"`
	SslVersionId *TestSslVersionId `json:"sslVersionId,omitempty"`
	// Target time for completion metric, defaults to 50% of time limit specified in seconds. (0 means default behavior)
	TargetTime *int32 `json:"targetTime,omitempty"`
	// Time limit for transaction in seconds. Exceeding this limit will result in a Timeout error.
	TimeLimit *int32 `json:"timeLimit,omitempty"`
	// Target for the test.
	Url string `json:"url"`
	// Contains a list of credential IDs (get `credentialId` from `/credentials` endpoint).
	Credentials []string `json:"credentials,omitempty"`
}

type _ApiInstantTest ApiInstantTest

// NewApiInstantTest instantiates a new ApiInstantTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInstantTest(requests []ApiRequest, url string) *ApiInstantTest {
	this := ApiInstantTest{}
	var collectProxyNetworkData bool = false
	this.CollectProxyNetworkData = &collectProxyNetworkData
	var followRedirects bool = true
	this.FollowRedirects = &followRedirects
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var overrideAgentProxy bool = false
	this.OverrideAgentProxy = &overrideAgentProxy
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	this.Requests = requests
	var sslVersionId TestSslVersionId = TESTSSLVERSIONID__0
	this.SslVersionId = &sslVersionId
	var timeLimit int32 = 30
	this.TimeLimit = &timeLimit
	this.Url = url
	return &this
}

// NewApiInstantTestWithDefaults instantiates a new ApiInstantTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInstantTestWithDefaults() *ApiInstantTest {
	this := ApiInstantTest{}
	var collectProxyNetworkData bool = false
	this.CollectProxyNetworkData = &collectProxyNetworkData
	var followRedirects bool = true
	this.FollowRedirects = &followRedirects
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var overrideAgentProxy bool = false
	this.OverrideAgentProxy = &overrideAgentProxy
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	var sslVersionId TestSslVersionId = TESTSSLVERSIONID__0
	this.SslVersionId = &sslVersionId
	var timeLimit int32 = 30
	this.TimeLimit = &timeLimit
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ApiInstantTest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ApiInstantTest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ApiInstantTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ApiInstantTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiInstantTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *ApiInstantTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiInstantTest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiInstantTest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiInstantTest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *ApiInstantTest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *ApiInstantTest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *ApiInstantTest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ApiInstantTest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ApiInstantTest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *ApiInstantTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *ApiInstantTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ApiInstantTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *ApiInstantTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *ApiInstantTest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *ApiInstantTest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *ApiInstantTest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *ApiInstantTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *ApiInstantTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *ApiInstantTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *ApiInstantTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *ApiInstantTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *ApiInstantTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiInstantTest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiInstantTest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiInstantTest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiInstantTest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiInstantTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *ApiInstantTest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ApiInstantTest) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ApiInstantTest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *ApiInstantTest) SetLabels(v []TestLabel) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *ApiInstantTest) GetSharedWithAccounts() []SharedWithAccount {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []SharedWithAccount
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetSharedWithAccountsOk() ([]SharedWithAccount, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *ApiInstantTest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []SharedWithAccount and assigns it to the SharedWithAccounts field.
func (o *ApiInstantTest) SetSharedWithAccounts(v []SharedWithAccount) {
	o.SharedWithAccounts = v
}

// GetCollectProxyNetworkData returns the CollectProxyNetworkData field value if set, zero value otherwise.
func (o *ApiInstantTest) GetCollectProxyNetworkData() bool {
	if o == nil || utils.IsNil(o.CollectProxyNetworkData) {
		var ret bool
		return ret
	}
	return *o.CollectProxyNetworkData
}

// GetCollectProxyNetworkDataOk returns a tuple with the CollectProxyNetworkData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetCollectProxyNetworkDataOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.CollectProxyNetworkData) {
		return nil, false
	}
	return o.CollectProxyNetworkData, true
}

// HasCollectProxyNetworkData returns a boolean if a field has been set.
func (o *ApiInstantTest) HasCollectProxyNetworkData() bool {
	if o != nil && !utils.IsNil(o.CollectProxyNetworkData) {
		return true
	}

	return false
}

// SetCollectProxyNetworkData gets a reference to the given bool and assigns it to the CollectProxyNetworkData field.
func (o *ApiInstantTest) SetCollectProxyNetworkData(v bool) {
	o.CollectProxyNetworkData = &v
}

// GetFollowRedirects returns the FollowRedirects field value if set, zero value otherwise.
func (o *ApiInstantTest) GetFollowRedirects() bool {
	if o == nil || utils.IsNil(o.FollowRedirects) {
		var ret bool
		return ret
	}
	return *o.FollowRedirects
}

// GetFollowRedirectsOk returns a tuple with the FollowRedirects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetFollowRedirectsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.FollowRedirects) {
		return nil, false
	}
	return o.FollowRedirects, true
}

// HasFollowRedirects returns a boolean if a field has been set.
func (o *ApiInstantTest) HasFollowRedirects() bool {
	if o != nil && !utils.IsNil(o.FollowRedirects) {
		return true
	}

	return false
}

// SetFollowRedirects gets a reference to the given bool and assigns it to the FollowRedirects field.
func (o *ApiInstantTest) SetFollowRedirects(v bool) {
	o.FollowRedirects = &v
}

// GetMtuMeasurements returns the MtuMeasurements field value if set, zero value otherwise.
func (o *ApiInstantTest) GetMtuMeasurements() bool {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		var ret bool
		return ret
	}
	return *o.MtuMeasurements
}

// GetMtuMeasurementsOk returns a tuple with the MtuMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetMtuMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		return nil, false
	}
	return o.MtuMeasurements, true
}

// HasMtuMeasurements returns a boolean if a field has been set.
func (o *ApiInstantTest) HasMtuMeasurements() bool {
	if o != nil && !utils.IsNil(o.MtuMeasurements) {
		return true
	}

	return false
}

// SetMtuMeasurements gets a reference to the given bool and assigns it to the MtuMeasurements field.
func (o *ApiInstantTest) SetMtuMeasurements(v bool) {
	o.MtuMeasurements = &v
}

// GetNetworkMeasurements returns the NetworkMeasurements field value if set, zero value otherwise.
func (o *ApiInstantTest) GetNetworkMeasurements() bool {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		var ret bool
		return ret
	}
	return *o.NetworkMeasurements
}

// GetNetworkMeasurementsOk returns a tuple with the NetworkMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetNetworkMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		return nil, false
	}
	return o.NetworkMeasurements, true
}

// HasNetworkMeasurements returns a boolean if a field has been set.
func (o *ApiInstantTest) HasNetworkMeasurements() bool {
	if o != nil && !utils.IsNil(o.NetworkMeasurements) {
		return true
	}

	return false
}

// SetNetworkMeasurements gets a reference to the given bool and assigns it to the NetworkMeasurements field.
func (o *ApiInstantTest) SetNetworkMeasurements(v bool) {
	o.NetworkMeasurements = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *ApiInstantTest) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *ApiInstantTest) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *ApiInstantTest) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetOverrideAgentProxy returns the OverrideAgentProxy field value if set, zero value otherwise.
func (o *ApiInstantTest) GetOverrideAgentProxy() bool {
	if o == nil || utils.IsNil(o.OverrideAgentProxy) {
		var ret bool
		return ret
	}
	return *o.OverrideAgentProxy
}

// GetOverrideAgentProxyOk returns a tuple with the OverrideAgentProxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetOverrideAgentProxyOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.OverrideAgentProxy) {
		return nil, false
	}
	return o.OverrideAgentProxy, true
}

// HasOverrideAgentProxy returns a boolean if a field has been set.
func (o *ApiInstantTest) HasOverrideAgentProxy() bool {
	if o != nil && !utils.IsNil(o.OverrideAgentProxy) {
		return true
	}

	return false
}

// SetOverrideAgentProxy gets a reference to the given bool and assigns it to the OverrideAgentProxy field.
func (o *ApiInstantTest) SetOverrideAgentProxy(v bool) {
	o.OverrideAgentProxy = &v
}

// GetOverrideProxyId returns the OverrideProxyId field value if set, zero value otherwise.
func (o *ApiInstantTest) GetOverrideProxyId() string {
	if o == nil || utils.IsNil(o.OverrideProxyId) {
		var ret string
		return ret
	}
	return *o.OverrideProxyId
}

// GetOverrideProxyIdOk returns a tuple with the OverrideProxyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetOverrideProxyIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OverrideProxyId) {
		return nil, false
	}
	return o.OverrideProxyId, true
}

// HasOverrideProxyId returns a boolean if a field has been set.
func (o *ApiInstantTest) HasOverrideProxyId() bool {
	if o != nil && !utils.IsNil(o.OverrideProxyId) {
		return true
	}

	return false
}

// SetOverrideProxyId gets a reference to the given string and assigns it to the OverrideProxyId field.
func (o *ApiInstantTest) SetOverrideProxyId(v string) {
	o.OverrideProxyId = &v
}

// GetPathTraceMode returns the PathTraceMode field value if set, zero value otherwise.
func (o *ApiInstantTest) GetPathTraceMode() TestPathTraceMode {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		var ret TestPathTraceMode
		return ret
	}
	return *o.PathTraceMode
}

// GetPathTraceModeOk returns a tuple with the PathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetPathTraceModeOk() (*TestPathTraceMode, bool) {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		return nil, false
	}
	return o.PathTraceMode, true
}

// HasPathTraceMode returns a boolean if a field has been set.
func (o *ApiInstantTest) HasPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.PathTraceMode) {
		return true
	}

	return false
}

// SetPathTraceMode gets a reference to the given TestPathTraceMode and assigns it to the PathTraceMode field.
func (o *ApiInstantTest) SetPathTraceMode(v TestPathTraceMode) {
	o.PathTraceMode = &v
}

// GetPredefinedVariables returns the PredefinedVariables field value if set, zero value otherwise.
func (o *ApiInstantTest) GetPredefinedVariables() []ApiPredefinedVariable {
	if o == nil || utils.IsNil(o.PredefinedVariables) {
		var ret []ApiPredefinedVariable
		return ret
	}
	return o.PredefinedVariables
}

// GetPredefinedVariablesOk returns a tuple with the PredefinedVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetPredefinedVariablesOk() ([]ApiPredefinedVariable, bool) {
	if o == nil || utils.IsNil(o.PredefinedVariables) {
		return nil, false
	}
	return o.PredefinedVariables, true
}

// HasPredefinedVariables returns a boolean if a field has been set.
func (o *ApiInstantTest) HasPredefinedVariables() bool {
	if o != nil && !utils.IsNil(o.PredefinedVariables) {
		return true
	}

	return false
}

// SetPredefinedVariables gets a reference to the given []ApiPredefinedVariable and assigns it to the PredefinedVariables field.
func (o *ApiInstantTest) SetPredefinedVariables(v []ApiPredefinedVariable) {
	o.PredefinedVariables = v
}

// GetProbeMode returns the ProbeMode field value if set, zero value otherwise.
func (o *ApiInstantTest) GetProbeMode() TestProbeMode {
	if o == nil || utils.IsNil(o.ProbeMode) {
		var ret TestProbeMode
		return ret
	}
	return *o.ProbeMode
}

// GetProbeModeOk returns a tuple with the ProbeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetProbeModeOk() (*TestProbeMode, bool) {
	if o == nil || utils.IsNil(o.ProbeMode) {
		return nil, false
	}
	return o.ProbeMode, true
}

// HasProbeMode returns a boolean if a field has been set.
func (o *ApiInstantTest) HasProbeMode() bool {
	if o != nil && !utils.IsNil(o.ProbeMode) {
		return true
	}

	return false
}

// SetProbeMode gets a reference to the given TestProbeMode and assigns it to the ProbeMode field.
func (o *ApiInstantTest) SetProbeMode(v TestProbeMode) {
	o.ProbeMode = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApiInstantTest) GetProtocol() TestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret TestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetProtocolOk() (*TestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApiInstantTest) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given TestProtocol and assigns it to the Protocol field.
func (o *ApiInstantTest) SetProtocol(v TestProtocol) {
	o.Protocol = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *ApiInstantTest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *ApiInstantTest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *ApiInstantTest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetRequests returns the Requests field value
func (o *ApiInstantTest) GetRequests() []ApiRequest {
	if o == nil {
		var ret []ApiRequest
		return ret
	}

	return o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetRequestsOk() ([]ApiRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *ApiInstantTest) SetRequests(v []ApiRequest) {
	o.Requests = v
}

// GetSslVersionId returns the SslVersionId field value if set, zero value otherwise.
func (o *ApiInstantTest) GetSslVersionId() TestSslVersionId {
	if o == nil || utils.IsNil(o.SslVersionId) {
		var ret TestSslVersionId
		return ret
	}
	return *o.SslVersionId
}

// GetSslVersionIdOk returns a tuple with the SslVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetSslVersionIdOk() (*TestSslVersionId, bool) {
	if o == nil || utils.IsNil(o.SslVersionId) {
		return nil, false
	}
	return o.SslVersionId, true
}

// HasSslVersionId returns a boolean if a field has been set.
func (o *ApiInstantTest) HasSslVersionId() bool {
	if o != nil && !utils.IsNil(o.SslVersionId) {
		return true
	}

	return false
}

// SetSslVersionId gets a reference to the given TestSslVersionId and assigns it to the SslVersionId field.
func (o *ApiInstantTest) SetSslVersionId(v TestSslVersionId) {
	o.SslVersionId = &v
}

// GetTargetTime returns the TargetTime field value if set, zero value otherwise.
func (o *ApiInstantTest) GetTargetTime() int32 {
	if o == nil || utils.IsNil(o.TargetTime) {
		var ret int32
		return ret
	}
	return *o.TargetTime
}

// GetTargetTimeOk returns a tuple with the TargetTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetTargetTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TargetTime) {
		return nil, false
	}
	return o.TargetTime, true
}

// HasTargetTime returns a boolean if a field has been set.
func (o *ApiInstantTest) HasTargetTime() bool {
	if o != nil && !utils.IsNil(o.TargetTime) {
		return true
	}

	return false
}

// SetTargetTime gets a reference to the given int32 and assigns it to the TargetTime field.
func (o *ApiInstantTest) SetTargetTime(v int32) {
	o.TargetTime = &v
}

// GetTimeLimit returns the TimeLimit field value if set, zero value otherwise.
func (o *ApiInstantTest) GetTimeLimit() int32 {
	if o == nil || utils.IsNil(o.TimeLimit) {
		var ret int32
		return ret
	}
	return *o.TimeLimit
}

// GetTimeLimitOk returns a tuple with the TimeLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetTimeLimitOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TimeLimit) {
		return nil, false
	}
	return o.TimeLimit, true
}

// HasTimeLimit returns a boolean if a field has been set.
func (o *ApiInstantTest) HasTimeLimit() bool {
	if o != nil && !utils.IsNil(o.TimeLimit) {
		return true
	}

	return false
}

// SetTimeLimit gets a reference to the given int32 and assigns it to the TimeLimit field.
func (o *ApiInstantTest) SetTimeLimit(v int32) {
	o.TimeLimit = &v
}

// GetUrl returns the Url field value
func (o *ApiInstantTest) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ApiInstantTest) SetUrl(v string) {
	o.Url = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ApiInstantTest) GetCredentials() []string {
	if o == nil || utils.IsNil(o.Credentials) {
		var ret []string
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInstantTest) GetCredentialsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ApiInstantTest) HasCredentials() bool {
	if o != nil && !utils.IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []string and assigns it to the Credentials field.
func (o *ApiInstantTest) SetCredentials(v []string) {
	o.Credentials = v
}

func (o ApiInstantTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInstantTest) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.CollectProxyNetworkData) {
		toSerialize["collectProxyNetworkData"] = o.CollectProxyNetworkData
	}
	if !utils.IsNil(o.FollowRedirects) {
		toSerialize["followRedirects"] = o.FollowRedirects
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
	if !utils.IsNil(o.OverrideAgentProxy) {
		toSerialize["overrideAgentProxy"] = o.OverrideAgentProxy
	}
	if !utils.IsNil(o.OverrideProxyId) {
		toSerialize["overrideProxyId"] = o.OverrideProxyId
	}
	if !utils.IsNil(o.PathTraceMode) {
		toSerialize["pathTraceMode"] = o.PathTraceMode
	}
	if !utils.IsNil(o.PredefinedVariables) {
		toSerialize["predefinedVariables"] = o.PredefinedVariables
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
	toSerialize["requests"] = o.Requests
	if !utils.IsNil(o.SslVersionId) {
		toSerialize["sslVersionId"] = o.SslVersionId
	}
	if !utils.IsNil(o.TargetTime) {
		toSerialize["targetTime"] = o.TargetTime
	}
	if !utils.IsNil(o.TimeLimit) {
		toSerialize["timeLimit"] = o.TimeLimit
	}
	toSerialize["url"] = o.Url
	if !utils.IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	return toSerialize, nil
}

func (o *ApiInstantTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"requests",
		"url",
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

	varApiInstantTest := _ApiInstantTest{}

    err = json.Unmarshal(data, &varApiInstantTest)

	if err != nil {
		return err
	}

	*o = ApiInstantTest(varApiInstantTest)

	return err
}

type NullableApiInstantTest struct {
	value *ApiInstantTest
	isSet bool
}

func (v NullableApiInstantTest) Get() *ApiInstantTest {
	return v.value
}

func (v *NullableApiInstantTest) Set(val *ApiInstantTest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInstantTest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInstantTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInstantTest(val *ApiInstantTest) *NullableApiInstantTest {
	return &NullableApiInstantTest{value: val, isSet: true}
}

func (v NullableApiInstantTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInstantTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


