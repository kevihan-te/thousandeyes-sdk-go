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

// checks if the BgpTestResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BgpTestResponse{}

// BgpTestResponse struct for BgpTestResponse
type BgpTestResponse struct {
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
	// Test is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Contains list of enabled BGP monitors.
	Monitors []Monitor `json:"monitors,omitempty"`
	// Indicate if queries for subprefixes detected under this prefix should included.
	IncludeCoveredPrefixes *bool `json:"includeCoveredPrefixes,omitempty"`
	// Indicate if all available public BGP monitors should be used, when ommited defaults to `bgpMeasurements` value.
	UsePublicBgp *bool `json:"usePublicBgp,omitempty"`
	// Indicates if alerts are enabled.
	AlertsEnabled *bool `json:"alertsEnabled,omitempty"`
	// A list of enabled alert rule objects.
	AlertRules []AlertRule `json:"alertRules,omitempty"`
	// a.b.c.d is a network address, with the prefix length defined as e. Prefixes can be any length from 8 to 24.
	Prefix string `json:"prefix"`
}

type _BgpTestResponse BgpTestResponse

// NewBgpTestResponse instantiates a new BgpTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpTestResponse(prefix string) *BgpTestResponse {
	this := BgpTestResponse{}
	var enabled bool = true
	this.Enabled = &enabled
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	this.Prefix = prefix
	return &this
}

// NewBgpTestResponseWithDefaults instantiates a new BgpTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpTestResponseWithDefaults() *BgpTestResponse {
	this := BgpTestResponse{}
	var enabled bool = true
	this.Enabled = &enabled
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BgpTestResponse) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BgpTestResponse) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *BgpTestResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *BgpTestResponse) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *BgpTestResponse) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *BgpTestResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BgpTestResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BgpTestResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BgpTestResponse) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *BgpTestResponse) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *BgpTestResponse) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *BgpTestResponse) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BgpTestResponse) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BgpTestResponse) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *BgpTestResponse) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *BgpTestResponse) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *BgpTestResponse) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *BgpTestResponse) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *BgpTestResponse) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *BgpTestResponse) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *BgpTestResponse) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *BgpTestResponse) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *BgpTestResponse) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *BgpTestResponse) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *BgpTestResponse) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *BgpTestResponse) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *BgpTestResponse) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BgpTestResponse) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BgpTestResponse) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BgpTestResponse) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BgpTestResponse) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BgpTestResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *BgpTestResponse) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *BgpTestResponse) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *BgpTestResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *BgpTestResponse) SetLabels(v []TestLabel) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *BgpTestResponse) GetSharedWithAccounts() []SharedWithAccount {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []SharedWithAccount
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetSharedWithAccountsOk() ([]SharedWithAccount, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *BgpTestResponse) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []SharedWithAccount and assigns it to the SharedWithAccounts field.
func (o *BgpTestResponse) SetSharedWithAccounts(v []SharedWithAccount) {
	o.SharedWithAccounts = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *BgpTestResponse) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *BgpTestResponse) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *BgpTestResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *BgpTestResponse) GetMonitors() []Monitor {
	if o == nil || utils.IsNil(o.Monitors) {
		var ret []Monitor
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetMonitorsOk() ([]Monitor, bool) {
	if o == nil || utils.IsNil(o.Monitors) {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *BgpTestResponse) HasMonitors() bool {
	if o != nil && !utils.IsNil(o.Monitors) {
		return true
	}

	return false
}

// SetMonitors gets a reference to the given []Monitor and assigns it to the Monitors field.
func (o *BgpTestResponse) SetMonitors(v []Monitor) {
	o.Monitors = v
}

// GetIncludeCoveredPrefixes returns the IncludeCoveredPrefixes field value if set, zero value otherwise.
func (o *BgpTestResponse) GetIncludeCoveredPrefixes() bool {
	if o == nil || utils.IsNil(o.IncludeCoveredPrefixes) {
		var ret bool
		return ret
	}
	return *o.IncludeCoveredPrefixes
}

// GetIncludeCoveredPrefixesOk returns a tuple with the IncludeCoveredPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetIncludeCoveredPrefixesOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IncludeCoveredPrefixes) {
		return nil, false
	}
	return o.IncludeCoveredPrefixes, true
}

// HasIncludeCoveredPrefixes returns a boolean if a field has been set.
func (o *BgpTestResponse) HasIncludeCoveredPrefixes() bool {
	if o != nil && !utils.IsNil(o.IncludeCoveredPrefixes) {
		return true
	}

	return false
}

// SetIncludeCoveredPrefixes gets a reference to the given bool and assigns it to the IncludeCoveredPrefixes field.
func (o *BgpTestResponse) SetIncludeCoveredPrefixes(v bool) {
	o.IncludeCoveredPrefixes = &v
}

// GetUsePublicBgp returns the UsePublicBgp field value if set, zero value otherwise.
func (o *BgpTestResponse) GetUsePublicBgp() bool {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		var ret bool
		return ret
	}
	return *o.UsePublicBgp
}

// GetUsePublicBgpOk returns a tuple with the UsePublicBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetUsePublicBgpOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		return nil, false
	}
	return o.UsePublicBgp, true
}

// HasUsePublicBgp returns a boolean if a field has been set.
func (o *BgpTestResponse) HasUsePublicBgp() bool {
	if o != nil && !utils.IsNil(o.UsePublicBgp) {
		return true
	}

	return false
}

// SetUsePublicBgp gets a reference to the given bool and assigns it to the UsePublicBgp field.
func (o *BgpTestResponse) SetUsePublicBgp(v bool) {
	o.UsePublicBgp = &v
}

// GetAlertsEnabled returns the AlertsEnabled field value if set, zero value otherwise.
func (o *BgpTestResponse) GetAlertsEnabled() bool {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		var ret bool
		return ret
	}
	return *o.AlertsEnabled
}

// GetAlertsEnabledOk returns a tuple with the AlertsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetAlertsEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		return nil, false
	}
	return o.AlertsEnabled, true
}

// HasAlertsEnabled returns a boolean if a field has been set.
func (o *BgpTestResponse) HasAlertsEnabled() bool {
	if o != nil && !utils.IsNil(o.AlertsEnabled) {
		return true
	}

	return false
}

// SetAlertsEnabled gets a reference to the given bool and assigns it to the AlertsEnabled field.
func (o *BgpTestResponse) SetAlertsEnabled(v bool) {
	o.AlertsEnabled = &v
}

// GetAlertRules returns the AlertRules field value if set, zero value otherwise.
func (o *BgpTestResponse) GetAlertRules() []AlertRule {
	if o == nil || utils.IsNil(o.AlertRules) {
		var ret []AlertRule
		return ret
	}
	return o.AlertRules
}

// GetAlertRulesOk returns a tuple with the AlertRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetAlertRulesOk() ([]AlertRule, bool) {
	if o == nil || utils.IsNil(o.AlertRules) {
		return nil, false
	}
	return o.AlertRules, true
}

// HasAlertRules returns a boolean if a field has been set.
func (o *BgpTestResponse) HasAlertRules() bool {
	if o != nil && !utils.IsNil(o.AlertRules) {
		return true
	}

	return false
}

// SetAlertRules gets a reference to the given []AlertRule and assigns it to the AlertRules field.
func (o *BgpTestResponse) SetAlertRules(v []AlertRule) {
	o.AlertRules = v
}

// GetPrefix returns the Prefix field value
func (o *BgpTestResponse) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *BgpTestResponse) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *BgpTestResponse) SetPrefix(v string) {
	o.Prefix = v
}

func (o BgpTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpTestResponse) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !utils.IsNil(o.Monitors) {
		toSerialize["monitors"] = o.Monitors
	}
	if !utils.IsNil(o.IncludeCoveredPrefixes) {
		toSerialize["includeCoveredPrefixes"] = o.IncludeCoveredPrefixes
	}
	if !utils.IsNil(o.UsePublicBgp) {
		toSerialize["usePublicBgp"] = o.UsePublicBgp
	}
	if !utils.IsNil(o.AlertsEnabled) {
		toSerialize["alertsEnabled"] = o.AlertsEnabled
	}
	if !utils.IsNil(o.AlertRules) {
		toSerialize["alertRules"] = o.AlertRules
	}
	toSerialize["prefix"] = o.Prefix
	return toSerialize, nil
}

func (o *BgpTestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prefix",
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

	varBgpTestResponse := _BgpTestResponse{}

    err = json.Unmarshal(data, &varBgpTestResponse)

	if err != nil {
		return err
	}

	*o = BgpTestResponse(varBgpTestResponse)

	return err
}

type NullableBgpTestResponse struct {
	value *BgpTestResponse
	isSet bool
}

func (v NullableBgpTestResponse) Get() *BgpTestResponse {
	return v.value
}

func (v *NullableBgpTestResponse) Set(val *BgpTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpTestResponse(val *BgpTestResponse) *NullableBgpTestResponse {
	return &NullableBgpTestResponse{value: val, isSet: true}
}

func (v NullableBgpTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


