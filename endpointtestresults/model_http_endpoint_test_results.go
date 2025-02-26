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

// checks if the HttpEndpointTestResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpEndpointTestResults{}

// HttpEndpointTestResults struct for HttpEndpointTestResults
type HttpEndpointTestResults struct {
	Results []HttpEndpointTestResult `json:"results,omitempty"`
	Test *EndpointHttpServerTest `json:"test,omitempty"`
	// (Optional) When passing `window` or `startDate` parameter,  the client will also receive the `startDate` field indicating the UTC start date of the data's time range being retrieved  (ISO date-time format).
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window` or `endDate` parameter,  the client will also receive the `endDate` field indicating the UTC end date of the data's time range being retrieved  (ISO date-time format).
	EndDate *time.Time `json:"endDate,omitempty"`
	Links *PaginationNextAndSelfLink `json:"_links,omitempty"`
}

// NewHttpEndpointTestResults instantiates a new HttpEndpointTestResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpEndpointTestResults() *HttpEndpointTestResults {
	this := HttpEndpointTestResults{}
	return &this
}

// NewHttpEndpointTestResultsWithDefaults instantiates a new HttpEndpointTestResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpEndpointTestResultsWithDefaults() *HttpEndpointTestResults {
	this := HttpEndpointTestResults{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *HttpEndpointTestResults) GetResults() []HttpEndpointTestResult {
	if o == nil || utils.IsNil(o.Results) {
		var ret []HttpEndpointTestResult
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResults) GetResultsOk() ([]HttpEndpointTestResult, bool) {
	if o == nil || utils.IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *HttpEndpointTestResults) HasResults() bool {
	if o != nil && !utils.IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []HttpEndpointTestResult and assigns it to the Results field.
func (o *HttpEndpointTestResults) SetResults(v []HttpEndpointTestResult) {
	o.Results = v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *HttpEndpointTestResults) GetTest() EndpointHttpServerTest {
	if o == nil || utils.IsNil(o.Test) {
		var ret EndpointHttpServerTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResults) GetTestOk() (*EndpointHttpServerTest, bool) {
	if o == nil || utils.IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *HttpEndpointTestResults) HasTest() bool {
	if o != nil && !utils.IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given EndpointHttpServerTest and assigns it to the Test field.
func (o *HttpEndpointTestResults) SetTest(v EndpointHttpServerTest) {
	o.Test = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *HttpEndpointTestResults) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResults) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *HttpEndpointTestResults) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *HttpEndpointTestResults) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *HttpEndpointTestResults) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResults) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *HttpEndpointTestResults) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *HttpEndpointTestResults) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *HttpEndpointTestResults) GetLinks() PaginationNextAndSelfLink {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationNextAndSelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestResults) GetLinksOk() (*PaginationNextAndSelfLink, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *HttpEndpointTestResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationNextAndSelfLink and assigns it to the Links field.
func (o *HttpEndpointTestResults) SetLinks(v PaginationNextAndSelfLink) {
	o.Links = &v
}

func (o HttpEndpointTestResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpEndpointTestResults) ToMap() (map[string]interface{}, error) {
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

type NullableHttpEndpointTestResults struct {
	value *HttpEndpointTestResults
	isSet bool
}

func (v NullableHttpEndpointTestResults) Get() *HttpEndpointTestResults {
	return v.value
}

func (v *NullableHttpEndpointTestResults) Set(val *HttpEndpointTestResults) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpEndpointTestResults) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpEndpointTestResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpEndpointTestResults(val *HttpEndpointTestResults) *NullableHttpEndpointTestResults {
	return &NullableHttpEndpointTestResults{value: val, isSet: true}
}

func (v NullableHttpEndpointTestResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpEndpointTestResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


