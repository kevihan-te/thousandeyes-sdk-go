/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the RealUserEndpointTestResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RealUserEndpointTestResults{}

// RealUserEndpointTestResults struct for RealUserEndpointTestResults
type RealUserEndpointTestResults struct {
	Results []RealUserEndpointTest `json:"results,omitempty"`
	// (Optional) When passing `window` or `startDate` parameter,  the client will also receive the `startDate` field indicating the UTC start date of the data's time range being retrieved  (ISO date-time format).
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window` or `endDate` parameter,  the client will also receive the `endDate` field indicating the UTC end date of the data's time range being retrieved  (ISO date-time format).
	EndDate *time.Time `json:"endDate,omitempty"`
	Links *PaginationNextLink `json:"_links,omitempty"`
}

// NewRealUserEndpointTestResults instantiates a new RealUserEndpointTestResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealUserEndpointTestResults() *RealUserEndpointTestResults {
	this := RealUserEndpointTestResults{}
	return &this
}

// NewRealUserEndpointTestResultsWithDefaults instantiates a new RealUserEndpointTestResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealUserEndpointTestResultsWithDefaults() *RealUserEndpointTestResults {
	this := RealUserEndpointTestResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *RealUserEndpointTestResults) GetResults() []RealUserEndpointTest {
	if o == nil || utils.IsNil(o.Results) {
		var ret []RealUserEndpointTest
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestResults) GetResultsOk() ([]RealUserEndpointTest, bool) {
	if o == nil || utils.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *RealUserEndpointTestResults) HasResults() bool {
	if o != nil && !utils.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []RealUserEndpointTest and assigns it to the Results field.
func (o *RealUserEndpointTestResults) SetResults(v []RealUserEndpointTest) {
	o.Results = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *RealUserEndpointTestResults) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestResults) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *RealUserEndpointTestResults) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *RealUserEndpointTestResults) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *RealUserEndpointTestResults) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestResults) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *RealUserEndpointTestResults) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *RealUserEndpointTestResults) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RealUserEndpointTestResults) GetLinks() PaginationNextLink {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationNextLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestResults) GetLinksOk() (*PaginationNextLink, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RealUserEndpointTestResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationNextLink and assigns it to the Links field.
func (o *RealUserEndpointTestResults) SetLinks(v PaginationNextLink) {
	o.Links = &v
}

func (o RealUserEndpointTestResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealUserEndpointTestResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Results) {
		toSerialize["results"] = o.Results
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

type NullableRealUserEndpointTestResults struct {
	value *RealUserEndpointTestResults
	isSet bool
}

func (v NullableRealUserEndpointTestResults) Get() *RealUserEndpointTestResults {
	return v.value
}

func (v *NullableRealUserEndpointTestResults) Set(val *RealUserEndpointTestResults) {
	v.value = val
	v.isSet = true
}

func (v NullableRealUserEndpointTestResults) IsSet() bool {
	return v.isSet
}

func (v *NullableRealUserEndpointTestResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealUserEndpointTestResults(val *RealUserEndpointTestResults) *NullableRealUserEndpointTestResults {
	return &NullableRealUserEndpointTestResults{value: val, isSet: true}
}

func (v NullableRealUserEndpointTestResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealUserEndpointTestResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


