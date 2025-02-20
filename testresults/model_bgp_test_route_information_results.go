/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the BgpTestRouteInformationResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BgpTestRouteInformationResults{}

// BgpTestRouteInformationResults struct for BgpTestRouteInformationResults
type BgpTestRouteInformationResults struct {
	Results []BgpTestRouteInformationResult `json:"results,omitempty"`
	Test *SimpleTest `json:"test,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewBgpTestRouteInformationResults instantiates a new BgpTestRouteInformationResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpTestRouteInformationResults() *BgpTestRouteInformationResults {
	this := BgpTestRouteInformationResults{}
	return &this
}

// NewBgpTestRouteInformationResultsWithDefaults instantiates a new BgpTestRouteInformationResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpTestRouteInformationResultsWithDefaults() *BgpTestRouteInformationResults {
	this := BgpTestRouteInformationResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *BgpTestRouteInformationResults) GetResults() []BgpTestRouteInformationResult {
	if o == nil || utils.IsNil(o.Results) {
		var ret []BgpTestRouteInformationResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestRouteInformationResults) GetResultsOk() ([]BgpTestRouteInformationResult, bool) {
	if o == nil || utils.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *BgpTestRouteInformationResults) HasResults() bool {
	if o != nil && !utils.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []BgpTestRouteInformationResult and assigns it to the Results field.
func (o *BgpTestRouteInformationResults) SetResults(v []BgpTestRouteInformationResult) {
	o.Results = v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *BgpTestRouteInformationResults) GetTest() SimpleTest {
	if o == nil || utils.IsNil(o.Test) {
		var ret SimpleTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestRouteInformationResults) GetTestOk() (*SimpleTest, bool) {
	if o == nil || utils.IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *BgpTestRouteInformationResults) HasTest() bool {
	if o != nil && !utils.IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given SimpleTest and assigns it to the Test field.
func (o *BgpTestRouteInformationResults) SetTest(v SimpleTest) {
	o.Test = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BgpTestRouteInformationResults) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpTestRouteInformationResults) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BgpTestRouteInformationResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *BgpTestRouteInformationResults) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o BgpTestRouteInformationResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpTestRouteInformationResults) ToMap() (map[string]interface{}, error) {
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

type NullableBgpTestRouteInformationResults struct {
	value *BgpTestRouteInformationResults
	isSet bool
}

func (v NullableBgpTestRouteInformationResults) Get() *BgpTestRouteInformationResults {
	return v.value
}

func (v *NullableBgpTestRouteInformationResults) Set(val *BgpTestRouteInformationResults) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpTestRouteInformationResults) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpTestRouteInformationResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpTestRouteInformationResults(val *BgpTestRouteInformationResults) *NullableBgpTestRouteInformationResults {
	return &NullableBgpTestRouteInformationResults{value: val, isSet: true}
}

func (v NullableBgpTestRouteInformationResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpTestRouteInformationResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


