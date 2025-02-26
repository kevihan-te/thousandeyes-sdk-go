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

// checks if the DnsTraceTestRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DnsTraceTestRequest{}

// DnsTraceTestRequest struct for DnsTraceTestRequest
type DnsTraceTestRequest struct {
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
	DnsTransportProtocol *TestDnsTransportProtocol `json:"dnsTransportProtocol,omitempty"`
	// The target record for the test, with the record type suffixed. If no record type is specified, the test defaults to an ANY record.
	Domain string `json:"domain"`
	DnsQueryClass *DnsQueryClass `json:"dnsQueryClass,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	// Contains list of Agent IDs (get `agentId` from `/agents` endpoint).
	Agents []TestAgentRequest `json:"agents"`
}

type _DnsTraceTestRequest DnsTraceTestRequest

// NewDnsTraceTestRequest instantiates a new DnsTraceTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsTraceTestRequest(interval TestInterval, domain string, agents []TestAgentRequest) *DnsTraceTestRequest {
	this := DnsTraceTestRequest{}
	this.Interval = interval
	var enabled bool = true
	this.Enabled = &enabled
	var dnsTransportProtocol TestDnsTransportProtocol = TESTDNSTRANSPORTPROTOCOL_UDP
	this.DnsTransportProtocol = &dnsTransportProtocol
	this.Domain = domain
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	this.Agents = agents
	return &this
}

// NewDnsTraceTestRequestWithDefaults instantiates a new DnsTraceTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsTraceTestRequestWithDefaults() *DnsTraceTestRequest {
	this := DnsTraceTestRequest{}
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = interval
	var enabled bool = true
	this.Enabled = &enabled
	var dnsTransportProtocol TestDnsTransportProtocol = TESTDNSTRANSPORTPROTOCOL_UDP
	this.DnsTransportProtocol = &dnsTransportProtocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	return &this
}

// GetInterval returns the Interval field value
func (o *DnsTraceTestRequest) GetInterval() TestInterval {
	if o == nil {
		var ret TestInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetIntervalOk() (*TestInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *DnsTraceTestRequest) SetInterval(v TestInterval) {
	o.Interval = v
}

// GetAlertsEnabled returns the AlertsEnabled field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetAlertsEnabled() bool {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		var ret bool
		return ret
	}
	return *o.AlertsEnabled
}

// GetAlertsEnabledOk returns a tuple with the AlertsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetAlertsEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		return nil, false
	}
	return o.AlertsEnabled, true
}

// HasAlertsEnabled returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasAlertsEnabled() bool {
	if o != nil && !utils.IsNil(o.AlertsEnabled) {
		return true
	}

	return false
}

// SetAlertsEnabled gets a reference to the given bool and assigns it to the AlertsEnabled field.
func (o *DnsTraceTestRequest) SetAlertsEnabled(v bool) {
	o.AlertsEnabled = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DnsTraceTestRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetAlertRules returns the AlertRules field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetAlertRules() []string {
	if o == nil || utils.IsNil(o.AlertRules) {
		var ret []string
		return ret
	}
	return o.AlertRules
}

// GetAlertRulesOk returns a tuple with the AlertRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetAlertRulesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AlertRules) {
		return nil, false
	}
	return o.AlertRules, true
}

// HasAlertRules returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasAlertRules() bool {
	if o != nil && !utils.IsNil(o.AlertRules) {
		return true
	}

	return false
}

// SetAlertRules gets a reference to the given []string and assigns it to the AlertRules field.
func (o *DnsTraceTestRequest) SetAlertRules(v []string) {
	o.AlertRules = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DnsTraceTestRequest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *DnsTraceTestRequest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DnsTraceTestRequest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *DnsTraceTestRequest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *DnsTraceTestRequest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *DnsTraceTestRequest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *DnsTraceTestRequest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *DnsTraceTestRequest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *DnsTraceTestRequest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DnsTraceTestRequest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *DnsTraceTestRequest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetLabels() []string {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetLabelsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *DnsTraceTestRequest) SetLabels(v []string) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetSharedWithAccounts() []string {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []string
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetSharedWithAccountsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []string and assigns it to the SharedWithAccounts field.
func (o *DnsTraceTestRequest) SetSharedWithAccounts(v []string) {
	o.SharedWithAccounts = v
}

// GetDnsTransportProtocol returns the DnsTransportProtocol field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetDnsTransportProtocol() TestDnsTransportProtocol {
	if o == nil || utils.IsNil(o.DnsTransportProtocol) {
		var ret TestDnsTransportProtocol
		return ret
	}
	return *o.DnsTransportProtocol
}

// GetDnsTransportProtocolOk returns a tuple with the DnsTransportProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetDnsTransportProtocolOk() (*TestDnsTransportProtocol, bool) {
	if o == nil || utils.IsNil(o.DnsTransportProtocol) {
		return nil, false
	}
	return o.DnsTransportProtocol, true
}

// HasDnsTransportProtocol returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasDnsTransportProtocol() bool {
	if o != nil && !utils.IsNil(o.DnsTransportProtocol) {
		return true
	}

	return false
}

// SetDnsTransportProtocol gets a reference to the given TestDnsTransportProtocol and assigns it to the DnsTransportProtocol field.
func (o *DnsTraceTestRequest) SetDnsTransportProtocol(v TestDnsTransportProtocol) {
	o.DnsTransportProtocol = &v
}

// GetDomain returns the Domain field value
func (o *DnsTraceTestRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DnsTraceTestRequest) SetDomain(v string) {
	o.Domain = v
}

// GetDnsQueryClass returns the DnsQueryClass field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetDnsQueryClass() DnsQueryClass {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		var ret DnsQueryClass
		return ret
	}
	return *o.DnsQueryClass
}

// GetDnsQueryClassOk returns a tuple with the DnsQueryClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetDnsQueryClassOk() (*DnsQueryClass, bool) {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		return nil, false
	}
	return o.DnsQueryClass, true
}

// HasDnsQueryClass returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasDnsQueryClass() bool {
	if o != nil && !utils.IsNil(o.DnsQueryClass) {
		return true
	}

	return false
}

// SetDnsQueryClass gets a reference to the given DnsQueryClass and assigns it to the DnsQueryClass field.
func (o *DnsTraceTestRequest) SetDnsQueryClass(v DnsQueryClass) {
	o.DnsQueryClass = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *DnsTraceTestRequest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *DnsTraceTestRequest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *DnsTraceTestRequest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetAgents returns the Agents field value
func (o *DnsTraceTestRequest) GetAgents() []TestAgentRequest {
	if o == nil {
		var ret []TestAgentRequest
		return ret
	}

	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value
// and a boolean to check if the value has been set.
func (o *DnsTraceTestRequest) GetAgentsOk() ([]TestAgentRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Agents, true
}

// SetAgents sets field value
func (o *DnsTraceTestRequest) SetAgents(v []TestAgentRequest) {
	o.Agents = v
}

func (o DnsTraceTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsTraceTestRequest) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.DnsTransportProtocol) {
		toSerialize["dnsTransportProtocol"] = o.DnsTransportProtocol
	}
	toSerialize["domain"] = o.Domain
	if !utils.IsNil(o.DnsQueryClass) {
		toSerialize["dnsQueryClass"] = o.DnsQueryClass
	}
	if !utils.IsNil(o.RandomizedStartTime) {
		toSerialize["randomizedStartTime"] = o.RandomizedStartTime
	}
	toSerialize["agents"] = o.Agents
	return toSerialize, nil
}

func (o *DnsTraceTestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
		"domain",
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

	varDnsTraceTestRequest := _DnsTraceTestRequest{}

    err = json.Unmarshal(data, &varDnsTraceTestRequest)

	if err != nil {
		return err
	}

	*o = DnsTraceTestRequest(varDnsTraceTestRequest)

	return err
}

type NullableDnsTraceTestRequest struct {
	value *DnsTraceTestRequest
	isSet bool
}

func (v NullableDnsTraceTestRequest) Get() *DnsTraceTestRequest {
	return v.value
}

func (v *NullableDnsTraceTestRequest) Set(val *DnsTraceTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsTraceTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsTraceTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsTraceTestRequest(val *DnsTraceTestRequest) *NullableDnsTraceTestRequest {
	return &NullableDnsTraceTestRequest{value: val, isSet: true}
}

func (v NullableDnsTraceTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsTraceTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


