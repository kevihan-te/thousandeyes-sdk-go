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

// checks if the NetworkDynamicEndpointTestResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkDynamicEndpointTestResults{}

// NetworkDynamicEndpointTestResults struct for NetworkDynamicEndpointTestResults
type NetworkDynamicEndpointTestResults struct {
	Results []NetworkDynamicEndpointTestResult `json:"results,omitempty"`
	Test *DynamicTest `json:"test,omitempty"`
	// Total number of measurements that match the search criteria.
	TotalHits *int32 `json:"totalHits,omitempty"`
	// (Optional) When passing `window` or `startDate` parameter,  the client will also receive the `startDate` field indicating the UTC start date of the data's time range being retrieved  (ISO date-time format).
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window` or `endDate` parameter,  the client will also receive the `endDate` field indicating the UTC end date of the data's time range being retrieved  (ISO date-time format).
	EndDate *time.Time `json:"endDate,omitempty"`
	Links *PaginationNextLink `json:"_links,omitempty"`
}

// NewNetworkDynamicEndpointTestResults instantiates a new NetworkDynamicEndpointTestResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDynamicEndpointTestResults() *NetworkDynamicEndpointTestResults {
	this := NetworkDynamicEndpointTestResults{}
	return &this
}

// NewNetworkDynamicEndpointTestResultsWithDefaults instantiates a new NetworkDynamicEndpointTestResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDynamicEndpointTestResultsWithDefaults() *NetworkDynamicEndpointTestResults {
	this := NetworkDynamicEndpointTestResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *NetworkDynamicEndpointTestResults) GetResults() []NetworkDynamicEndpointTestResult {
	if o == nil || utils.IsNil(o.Results) {
		var ret []NetworkDynamicEndpointTestResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDynamicEndpointTestResults) GetResultsOk() ([]NetworkDynamicEndpointTestResult, bool) {
	if o == nil || utils.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *NetworkDynamicEndpointTestResults) HasResults() bool {
	if o != nil && !utils.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []NetworkDynamicEndpointTestResult and assigns it to the Results field.
func (o *NetworkDynamicEndpointTestResults) SetResults(v []NetworkDynamicEndpointTestResult) {
	o.Results = v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *NetworkDynamicEndpointTestResults) GetTest() DynamicTest {
	if o == nil || utils.IsNil(o.Test) {
		var ret DynamicTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDynamicEndpointTestResults) GetTestOk() (*DynamicTest, bool) {
	if o == nil || utils.IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *NetworkDynamicEndpointTestResults) HasTest() bool {
	if o != nil && !utils.IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given DynamicTest and assigns it to the Test field.
func (o *NetworkDynamicEndpointTestResults) SetTest(v DynamicTest) {
	o.Test = &v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *NetworkDynamicEndpointTestResults) GetTotalHits() int32 {
	if o == nil || utils.IsNil(o.TotalHits) {
		var ret int32
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDynamicEndpointTestResults) GetTotalHitsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalHits) {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *NetworkDynamicEndpointTestResults) HasTotalHits() bool {
	if o != nil && !utils.IsNil(o.TotalHits) {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int32 and assigns it to the TotalHits field.
func (o *NetworkDynamicEndpointTestResults) SetTotalHits(v int32) {
	o.TotalHits = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *NetworkDynamicEndpointTestResults) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDynamicEndpointTestResults) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *NetworkDynamicEndpointTestResults) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *NetworkDynamicEndpointTestResults) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *NetworkDynamicEndpointTestResults) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDynamicEndpointTestResults) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *NetworkDynamicEndpointTestResults) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *NetworkDynamicEndpointTestResults) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *NetworkDynamicEndpointTestResults) GetLinks() PaginationNextLink {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationNextLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDynamicEndpointTestResults) GetLinksOk() (*PaginationNextLink, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *NetworkDynamicEndpointTestResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationNextLink and assigns it to the Links field.
func (o *NetworkDynamicEndpointTestResults) SetLinks(v PaginationNextLink) {
	o.Links = &v
}

func (o NetworkDynamicEndpointTestResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDynamicEndpointTestResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !utils.IsNil(o.Test) {
		toSerialize["test"] = o.Test
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

type NullableNetworkDynamicEndpointTestResults struct {
	value *NetworkDynamicEndpointTestResults
	isSet bool
}

func (v NullableNetworkDynamicEndpointTestResults) Get() *NetworkDynamicEndpointTestResults {
	return v.value
}

func (v *NullableNetworkDynamicEndpointTestResults) Set(val *NetworkDynamicEndpointTestResults) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkDynamicEndpointTestResults) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkDynamicEndpointTestResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkDynamicEndpointTestResults(val *NetworkDynamicEndpointTestResults) *NullableNetworkDynamicEndpointTestResults {
	return &NullableNetworkDynamicEndpointTestResults{value: val, isSet: true}
}

func (v NullableNetworkDynamicEndpointTestResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkDynamicEndpointTestResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


