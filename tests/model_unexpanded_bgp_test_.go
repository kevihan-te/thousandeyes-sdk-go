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

// checks if the UnexpandedBgpTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UnexpandedBgpTest{}

// UnexpandedBgpTest struct for UnexpandedBgpTest
type UnexpandedBgpTest struct {
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
	// Test is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Indicate if queries for subprefixes detected under this prefix should included.
	IncludeCoveredPrefixes *bool `json:"includeCoveredPrefixes,omitempty"`
	// Indicate if all available public BGP monitors should be used, when ommited defaults to `bgpMeasurements` value.
	UsePublicBgp *bool `json:"usePublicBgp,omitempty"`
	// Indicates if alerts are enabled.
	AlertsEnabled *bool `json:"alertsEnabled,omitempty"`
	// a.b.c.d is a network address, with the prefix length defined as e. Prefixes can be any length from 8 to 24.
	Prefix string `json:"prefix"`
}

type _UnexpandedBgpTest UnexpandedBgpTest

// NewUnexpandedBgpTest instantiates a new UnexpandedBgpTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnexpandedBgpTest(prefix string) *UnexpandedBgpTest {
	this := UnexpandedBgpTest{}
	var enabled bool = true
	this.Enabled = &enabled
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	this.Prefix = prefix
	return &this
}

// NewUnexpandedBgpTestWithDefaults instantiates a new UnexpandedBgpTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnexpandedBgpTestWithDefaults() *UnexpandedBgpTest {
	this := UnexpandedBgpTest{}
	var enabled bool = true
	this.Enabled = &enabled
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *UnexpandedBgpTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *UnexpandedBgpTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UnexpandedBgpTest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *UnexpandedBgpTest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *UnexpandedBgpTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *UnexpandedBgpTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *UnexpandedBgpTest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *UnexpandedBgpTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *UnexpandedBgpTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UnexpandedBgpTest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *UnexpandedBgpTest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UnexpandedBgpTest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetIncludeCoveredPrefixes returns the IncludeCoveredPrefixes field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetIncludeCoveredPrefixes() bool {
	if o == nil || utils.IsNil(o.IncludeCoveredPrefixes) {
		var ret bool
		return ret
	}
	return *o.IncludeCoveredPrefixes
}

// GetIncludeCoveredPrefixesOk returns a tuple with the IncludeCoveredPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetIncludeCoveredPrefixesOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IncludeCoveredPrefixes) {
		return nil, false
	}
	return o.IncludeCoveredPrefixes, true
}

// HasIncludeCoveredPrefixes returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasIncludeCoveredPrefixes() bool {
	if o != nil && !utils.IsNil(o.IncludeCoveredPrefixes) {
		return true
	}

	return false
}

// SetIncludeCoveredPrefixes gets a reference to the given bool and assigns it to the IncludeCoveredPrefixes field.
func (o *UnexpandedBgpTest) SetIncludeCoveredPrefixes(v bool) {
	o.IncludeCoveredPrefixes = &v
}

// GetUsePublicBgp returns the UsePublicBgp field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetUsePublicBgp() bool {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		var ret bool
		return ret
	}
	return *o.UsePublicBgp
}

// GetUsePublicBgpOk returns a tuple with the UsePublicBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetUsePublicBgpOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		return nil, false
	}
	return o.UsePublicBgp, true
}

// HasUsePublicBgp returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasUsePublicBgp() bool {
	if o != nil && !utils.IsNil(o.UsePublicBgp) {
		return true
	}

	return false
}

// SetUsePublicBgp gets a reference to the given bool and assigns it to the UsePublicBgp field.
func (o *UnexpandedBgpTest) SetUsePublicBgp(v bool) {
	o.UsePublicBgp = &v
}

// GetAlertsEnabled returns the AlertsEnabled field value if set, zero value otherwise.
func (o *UnexpandedBgpTest) GetAlertsEnabled() bool {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		var ret bool
		return ret
	}
	return *o.AlertsEnabled
}

// GetAlertsEnabledOk returns a tuple with the AlertsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetAlertsEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.AlertsEnabled) {
		return nil, false
	}
	return o.AlertsEnabled, true
}

// HasAlertsEnabled returns a boolean if a field has been set.
func (o *UnexpandedBgpTest) HasAlertsEnabled() bool {
	if o != nil && !utils.IsNil(o.AlertsEnabled) {
		return true
	}

	return false
}

// SetAlertsEnabled gets a reference to the given bool and assigns it to the AlertsEnabled field.
func (o *UnexpandedBgpTest) SetAlertsEnabled(v bool) {
	o.AlertsEnabled = &v
}

// GetPrefix returns the Prefix field value
func (o *UnexpandedBgpTest) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *UnexpandedBgpTest) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *UnexpandedBgpTest) SetPrefix(v string) {
	o.Prefix = v
}

func (o UnexpandedBgpTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnexpandedBgpTest) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
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
	toSerialize["prefix"] = o.Prefix
	return toSerialize, nil
}

func (o *UnexpandedBgpTest) UnmarshalJSON(data []byte) (err error) {
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

	varUnexpandedBgpTest := _UnexpandedBgpTest{}

    err = json.Unmarshal(data, &varUnexpandedBgpTest)

	if err != nil {
		return err
	}

	*o = UnexpandedBgpTest(varUnexpandedBgpTest)

	return err
}

type NullableUnexpandedBgpTest struct {
	value *UnexpandedBgpTest
	isSet bool
}

func (v NullableUnexpandedBgpTest) Get() *UnexpandedBgpTest {
	return v.value
}

func (v *NullableUnexpandedBgpTest) Set(val *UnexpandedBgpTest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnexpandedBgpTest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnexpandedBgpTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnexpandedBgpTest(val *UnexpandedBgpTest) *NullableUnexpandedBgpTest {
	return &NullableUnexpandedBgpTest{value: val, isSet: true}
}

func (v NullableUnexpandedBgpTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnexpandedBgpTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


