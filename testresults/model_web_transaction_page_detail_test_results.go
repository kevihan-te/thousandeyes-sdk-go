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
)

// checks if the WebTransactionPageDetailTestResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WebTransactionPageDetailTestResults{}

// WebTransactionPageDetailTestResults struct for WebTransactionPageDetailTestResults
type WebTransactionPageDetailTestResults struct {
	Results []WebTransactionPageDetailTestResult `json:"results,omitempty"`
	Test *SimpleTest `json:"test,omitempty"`
	Links *PaginationLinks `json:"_links,omitempty"`
}

// NewWebTransactionPageDetailTestResults instantiates a new WebTransactionPageDetailTestResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebTransactionPageDetailTestResults() *WebTransactionPageDetailTestResults {
	this := WebTransactionPageDetailTestResults{}
	return &this
}

// NewWebTransactionPageDetailTestResultsWithDefaults instantiates a new WebTransactionPageDetailTestResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebTransactionPageDetailTestResultsWithDefaults() *WebTransactionPageDetailTestResults {
	this := WebTransactionPageDetailTestResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *WebTransactionPageDetailTestResults) GetResults() []WebTransactionPageDetailTestResult {
	if o == nil || utils.IsNil(o.Results) {
		var ret []WebTransactionPageDetailTestResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionPageDetailTestResults) GetResultsOk() ([]WebTransactionPageDetailTestResult, bool) {
	if o == nil || utils.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *WebTransactionPageDetailTestResults) HasResults() bool {
	if o != nil && !utils.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []WebTransactionPageDetailTestResult and assigns it to the Results field.
func (o *WebTransactionPageDetailTestResults) SetResults(v []WebTransactionPageDetailTestResult) {
	o.Results = v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *WebTransactionPageDetailTestResults) GetTest() SimpleTest {
	if o == nil || utils.IsNil(o.Test) {
		var ret SimpleTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionPageDetailTestResults) GetTestOk() (*SimpleTest, bool) {
	if o == nil || utils.IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *WebTransactionPageDetailTestResults) HasTest() bool {
	if o != nil && !utils.IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given SimpleTest and assigns it to the Test field.
func (o *WebTransactionPageDetailTestResults) SetTest(v SimpleTest) {
	o.Test = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *WebTransactionPageDetailTestResults) GetLinks() PaginationLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebTransactionPageDetailTestResults) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *WebTransactionPageDetailTestResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *WebTransactionPageDetailTestResults) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o WebTransactionPageDetailTestResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebTransactionPageDetailTestResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !utils.IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableWebTransactionPageDetailTestResults struct {
	value *WebTransactionPageDetailTestResults
	isSet bool
}

func (v NullableWebTransactionPageDetailTestResults) Get() *WebTransactionPageDetailTestResults {
	return v.value
}

func (v *NullableWebTransactionPageDetailTestResults) Set(val *WebTransactionPageDetailTestResults) {
	v.value = val
	v.isSet = true
}

func (v NullableWebTransactionPageDetailTestResults) IsSet() bool {
	return v.isSet
}

func (v *NullableWebTransactionPageDetailTestResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebTransactionPageDetailTestResults(val *WebTransactionPageDetailTestResults) *NullableWebTransactionPageDetailTestResults {
	return &NullableWebTransactionPageDetailTestResults{value: val, isSet: true}
}

func (v NullableWebTransactionPageDetailTestResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebTransactionPageDetailTestResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


