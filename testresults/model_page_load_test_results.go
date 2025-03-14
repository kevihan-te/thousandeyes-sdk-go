/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the PageLoadTestResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PageLoadTestResults{}

// PageLoadTestResults struct for PageLoadTestResults
type PageLoadTestResults struct {
	Results []PageLoadTestResult `json:"results,omitempty"`
	Test *SimpleTest `json:"test,omitempty"`
	// (Optional) When passing `window` or `startDate` parameter,  the client will also receive the `startDate` field indicating the UTC start date of the data's time range being retrieved  (ISO date-time format).
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window` or `endDate` parameter,  the client will also receive the `endDate` field indicating the UTC end date of the data's time range being retrieved  (ISO date-time format).
	EndDate *time.Time `json:"endDate,omitempty"`
	Links *PaginationLinks `json:"_links,omitempty"`
}

// NewPageLoadTestResults instantiates a new PageLoadTestResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageLoadTestResults() *PageLoadTestResults {
	this := PageLoadTestResults{}
	return &this
}

// NewPageLoadTestResultsWithDefaults instantiates a new PageLoadTestResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageLoadTestResultsWithDefaults() *PageLoadTestResults {
	this := PageLoadTestResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *PageLoadTestResults) GetResults() []PageLoadTestResult {
	if o == nil || utils.IsNil(o.Results) {
		var ret []PageLoadTestResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageLoadTestResults) GetResultsOk() ([]PageLoadTestResult, bool) {
	if o == nil || utils.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *PageLoadTestResults) HasResults() bool {
	if o != nil && !utils.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []PageLoadTestResult and assigns it to the Results field.
func (o *PageLoadTestResults) SetResults(v []PageLoadTestResult) {
	o.Results = v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *PageLoadTestResults) GetTest() SimpleTest {
	if o == nil || utils.IsNil(o.Test) {
		var ret SimpleTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageLoadTestResults) GetTestOk() (*SimpleTest, bool) {
	if o == nil || utils.IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *PageLoadTestResults) HasTest() bool {
	if o != nil && !utils.IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given SimpleTest and assigns it to the Test field.
func (o *PageLoadTestResults) SetTest(v SimpleTest) {
	o.Test = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *PageLoadTestResults) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageLoadTestResults) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *PageLoadTestResults) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *PageLoadTestResults) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *PageLoadTestResults) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageLoadTestResults) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *PageLoadTestResults) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *PageLoadTestResults) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PageLoadTestResults) GetLinks() PaginationLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageLoadTestResults) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PageLoadTestResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *PageLoadTestResults) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o PageLoadTestResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageLoadTestResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !utils.IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}
	if !utils.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !utils.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullablePageLoadTestResults struct {
	value *PageLoadTestResults
	isSet bool
}

func (v NullablePageLoadTestResults) Get() *PageLoadTestResults {
	return v.value
}

func (v *NullablePageLoadTestResults) Set(val *PageLoadTestResults) {
	v.value = val
	v.isSet = true
}

func (v NullablePageLoadTestResults) IsSet() bool {
	return v.isSet
}

func (v *NullablePageLoadTestResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageLoadTestResults(val *PageLoadTestResults) *NullablePageLoadTestResults {
	return &NullablePageLoadTestResults{value: val, isSet: true}
}

func (v NullablePageLoadTestResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageLoadTestResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


