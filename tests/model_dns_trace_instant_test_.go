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

// checks if the DnsTraceInstantTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DnsTraceInstantTest{}

// DnsTraceInstantTest struct for DnsTraceInstantTest
type DnsTraceInstantTest struct {
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
	DnsTransportProtocol *TestDnsTransportProtocol `json:"dnsTransportProtocol,omitempty"`
	// The target record for the test, with the record type suffixed. If no record type is specified, the test defaults to an ANY record.
	Domain string `json:"domain"`
	DnsQueryClass *DnsQueryClass `json:"dnsQueryClass,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
}

type _DnsTraceInstantTest DnsTraceInstantTest

// NewDnsTraceInstantTest instantiates a new DnsTraceInstantTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsTraceInstantTest(domain string) *DnsTraceInstantTest {
	this := DnsTraceInstantTest{}
	var dnsTransportProtocol TestDnsTransportProtocol = TESTDNSTRANSPORTPROTOCOL_UDP
	this.DnsTransportProtocol = &dnsTransportProtocol
	this.Domain = domain
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	return &this
}

// NewDnsTraceInstantTestWithDefaults instantiates a new DnsTraceInstantTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsTraceInstantTestWithDefaults() *DnsTraceInstantTest {
	this := DnsTraceInstantTest{}
	var dnsTransportProtocol TestDnsTransportProtocol = TESTDNSTRANSPORTPROTOCOL_UDP
	this.DnsTransportProtocol = &dnsTransportProtocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DnsTraceInstantTest) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *DnsTraceInstantTest) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DnsTraceInstantTest) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *DnsTraceInstantTest) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *DnsTraceInstantTest) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *DnsTraceInstantTest) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *DnsTraceInstantTest) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *DnsTraceInstantTest) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *DnsTraceInstantTest) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DnsTraceInstantTest) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *DnsTraceInstantTest) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *DnsTraceInstantTest) SetLabels(v []TestLabel) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetSharedWithAccounts() []SharedWithAccount {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []SharedWithAccount
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetSharedWithAccountsOk() ([]SharedWithAccount, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []SharedWithAccount and assigns it to the SharedWithAccounts field.
func (o *DnsTraceInstantTest) SetSharedWithAccounts(v []SharedWithAccount) {
	o.SharedWithAccounts = v
}

// GetDnsTransportProtocol returns the DnsTransportProtocol field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetDnsTransportProtocol() TestDnsTransportProtocol {
	if o == nil || utils.IsNil(o.DnsTransportProtocol) {
		var ret TestDnsTransportProtocol
		return ret
	}
	return *o.DnsTransportProtocol
}

// GetDnsTransportProtocolOk returns a tuple with the DnsTransportProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetDnsTransportProtocolOk() (*TestDnsTransportProtocol, bool) {
	if o == nil || utils.IsNil(o.DnsTransportProtocol) {
		return nil, false
	}
	return o.DnsTransportProtocol, true
}

// HasDnsTransportProtocol returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasDnsTransportProtocol() bool {
	if o != nil && !utils.IsNil(o.DnsTransportProtocol) {
		return true
	}

	return false
}

// SetDnsTransportProtocol gets a reference to the given TestDnsTransportProtocol and assigns it to the DnsTransportProtocol field.
func (o *DnsTraceInstantTest) SetDnsTransportProtocol(v TestDnsTransportProtocol) {
	o.DnsTransportProtocol = &v
}

// GetDomain returns the Domain field value
func (o *DnsTraceInstantTest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DnsTraceInstantTest) SetDomain(v string) {
	o.Domain = v
}

// GetDnsQueryClass returns the DnsQueryClass field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetDnsQueryClass() DnsQueryClass {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		var ret DnsQueryClass
		return ret
	}
	return *o.DnsQueryClass
}

// GetDnsQueryClassOk returns a tuple with the DnsQueryClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetDnsQueryClassOk() (*DnsQueryClass, bool) {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		return nil, false
	}
	return o.DnsQueryClass, true
}

// HasDnsQueryClass returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasDnsQueryClass() bool {
	if o != nil && !utils.IsNil(o.DnsQueryClass) {
		return true
	}

	return false
}

// SetDnsQueryClass gets a reference to the given DnsQueryClass and assigns it to the DnsQueryClass field.
func (o *DnsTraceInstantTest) SetDnsQueryClass(v DnsQueryClass) {
	o.DnsQueryClass = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *DnsTraceInstantTest) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsTraceInstantTest) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *DnsTraceInstantTest) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *DnsTraceInstantTest) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

func (o DnsTraceInstantTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsTraceInstantTest) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

func (o *DnsTraceInstantTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
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

	varDnsTraceInstantTest := _DnsTraceInstantTest{}

    err = json.Unmarshal(data, &varDnsTraceInstantTest)

	if err != nil {
		return err
	}

	*o = DnsTraceInstantTest(varDnsTraceInstantTest)

	return err
}

type NullableDnsTraceInstantTest struct {
	value *DnsTraceInstantTest
	isSet bool
}

func (v NullableDnsTraceInstantTest) Get() *DnsTraceInstantTest {
	return v.value
}

func (v *NullableDnsTraceInstantTest) Set(val *DnsTraceInstantTest) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsTraceInstantTest) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsTraceInstantTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsTraceInstantTest(val *DnsTraceInstantTest) *NullableDnsTraceInstantTest {
	return &NullableDnsTraceInstantTest{value: val, isSet: true}
}

func (v NullableDnsTraceInstantTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsTraceInstantTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


