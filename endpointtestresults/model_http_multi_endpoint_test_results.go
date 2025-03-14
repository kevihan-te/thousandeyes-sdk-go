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

// checks if the HttpMultiEndpointTestResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpMultiEndpointTestResults{}

// HttpMultiEndpointTestResults struct for HttpMultiEndpointTestResults
type HttpMultiEndpointTestResults struct {
	Results []HttpEndpointTestResult `json:"results,omitempty"`
	// Total number of measurements that match the search criteria.
	TotalHits *int32 `json:"totalHits,omitempty"`
	// (Optional) When passing `window` or `startDate` parameter,  the client will also receive the `startDate` field indicating the UTC start date of the data's time range being retrieved  (ISO date-time format).
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window` or `endDate` parameter,  the client will also receive the `endDate` field indicating the UTC end date of the data's time range being retrieved  (ISO date-time format).
	EndDate *time.Time `json:"endDate,omitempty"`
	Links *PaginationNextAndSelfLink `json:"_links,omitempty"`
}

// NewHttpMultiEndpointTestResults instantiates a new HttpMultiEndpointTestResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpMultiEndpointTestResults() *HttpMultiEndpointTestResults {
	this := HttpMultiEndpointTestResults{}
	return &this
}

// NewHttpMultiEndpointTestResultsWithDefaults instantiates a new HttpMultiEndpointTestResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpMultiEndpointTestResultsWithDefaults() *HttpMultiEndpointTestResults {
	this := HttpMultiEndpointTestResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *HttpMultiEndpointTestResults) GetResults() []HttpEndpointTestResult {
	if o == nil || utils.IsNil(o.Results) {
		var ret []HttpEndpointTestResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpMultiEndpointTestResults) GetResultsOk() ([]HttpEndpointTestResult, bool) {
	if o == nil || utils.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *HttpMultiEndpointTestResults) HasResults() bool {
	if o != nil && !utils.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []HttpEndpointTestResult and assigns it to the Results field.
func (o *HttpMultiEndpointTestResults) SetResults(v []HttpEndpointTestResult) {
	o.Results = v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *HttpMultiEndpointTestResults) GetTotalHits() int32 {
	if o == nil || utils.IsNil(o.TotalHits) {
		var ret int32
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpMultiEndpointTestResults) GetTotalHitsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalHits) {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *HttpMultiEndpointTestResults) HasTotalHits() bool {
	if o != nil && !utils.IsNil(o.TotalHits) {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int32 and assigns it to the TotalHits field.
func (o *HttpMultiEndpointTestResults) SetTotalHits(v int32) {
	o.TotalHits = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *HttpMultiEndpointTestResults) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpMultiEndpointTestResults) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *HttpMultiEndpointTestResults) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *HttpMultiEndpointTestResults) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *HttpMultiEndpointTestResults) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpMultiEndpointTestResults) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *HttpMultiEndpointTestResults) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *HttpMultiEndpointTestResults) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *HttpMultiEndpointTestResults) GetLinks() PaginationNextAndSelfLink {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationNextAndSelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpMultiEndpointTestResults) GetLinksOk() (*PaginationNextAndSelfLink, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HttpMultiEndpointTestResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationNextAndSelfLink and assigns it to the Links field.
func (o *HttpMultiEndpointTestResults) SetLinks(v PaginationNextAndSelfLink) {
	o.Links = &v
}

func (o HttpMultiEndpointTestResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpMultiEndpointTestResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !utils.IsNil(o.TotalHits) {
		toSerialize["totalHits"] = o.TotalHits
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

type NullableHttpMultiEndpointTestResults struct {
	value *HttpMultiEndpointTestResults
	isSet bool
}

func (v NullableHttpMultiEndpointTestResults) Get() *HttpMultiEndpointTestResults {
	return v.value
}

func (v *NullableHttpMultiEndpointTestResults) Set(val *HttpMultiEndpointTestResults) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpMultiEndpointTestResults) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpMultiEndpointTestResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpMultiEndpointTestResults(val *HttpMultiEndpointTestResults) *NullableHttpMultiEndpointTestResults {
	return &NullableHttpMultiEndpointTestResults{value: val, isSet: true}
}

func (v NullableHttpMultiEndpointTestResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpMultiEndpointTestResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


