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

// checks if the TestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestResult{}

// TestResult struct for TestResult
type TestResult struct {
	// Data point date UTC (ISO date-time format).
	Date *time.Time `json:"date,omitempty"`
	// Epoch time (seconds) indicating the start time of the round
	RoundId *int32 `json:"roundId,omitempty"`
	Links *TestResultAppLinks `json:"_links,omitempty"`
}

// NewTestResult instantiates a new TestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestResult() *TestResult {
	this := TestResult{}
	return &this
}

// NewTestResultWithDefaults instantiates a new TestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestResultWithDefaults() *TestResult {
	this := TestResult{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *TestResult) GetDate() time.Time {
	if o == nil || utils.IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResult) GetDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *TestResult) HasDate() bool {
	if o != nil && !utils.IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *TestResult) SetDate(v time.Time) {
	o.Date = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *TestResult) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResult) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *TestResult) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *TestResult) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *TestResult) GetLinks() TestResultAppLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestResultAppLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestResult) GetLinksOk() (*TestResultAppLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *TestResult) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestResultAppLinks and assigns it to the Links field.
func (o *TestResult) SetLinks(v TestResultAppLinks) {
	o.Links = &v
}

func (o TestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !utils.IsNil(o.RoundId) {
		toSerialize["roundId"] = o.RoundId
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTestResult struct {
	value *TestResult
	isSet bool
}

func (v NullableTestResult) Get() *TestResult {
	return v.value
}

func (v *NullableTestResult) Set(val *TestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestResult(val *TestResult) *NullableTestResult {
	return &NullableTestResult{value: val, isSet: true}
}

func (v NullableTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


