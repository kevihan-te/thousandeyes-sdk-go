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
)

// checks if the HttpEndpointTestsDataThresholdFilters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpEndpointTestsDataThresholdFilters{}

// HttpEndpointTestsDataThresholdFilters All filters are applied based on the conditional operator (and/or).
type HttpEndpointTestsDataThresholdFilters struct {
	Filters []HttpEndpointTestsDataThresholdFilter `json:"filters,omitempty"`
	ConditionalOperator *ConditionalOperator `json:"conditionalOperator,omitempty"`
}

// NewHttpEndpointTestsDataThresholdFilters instantiates a new HttpEndpointTestsDataThresholdFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpEndpointTestsDataThresholdFilters() *HttpEndpointTestsDataThresholdFilters {
	this := HttpEndpointTestsDataThresholdFilters{}
	return &this
}

// NewHttpEndpointTestsDataThresholdFiltersWithDefaults instantiates a new HttpEndpointTestsDataThresholdFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpEndpointTestsDataThresholdFiltersWithDefaults() *HttpEndpointTestsDataThresholdFilters {
	this := HttpEndpointTestsDataThresholdFilters{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *HttpEndpointTestsDataThresholdFilters) GetFilters() []HttpEndpointTestsDataThresholdFilter {
	if o == nil || utils.IsNil(o.Filters) {
		var ret []HttpEndpointTestsDataThresholdFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestsDataThresholdFilters) GetFiltersOk() ([]HttpEndpointTestsDataThresholdFilter, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *HttpEndpointTestsDataThresholdFilters) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []HttpEndpointTestsDataThresholdFilter and assigns it to the Filters field.
func (o *HttpEndpointTestsDataThresholdFilters) SetFilters(v []HttpEndpointTestsDataThresholdFilter) {
	o.Filters = v
}

// GetConditionalOperator returns the ConditionalOperator field value if set, zero value otherwise.
func (o *HttpEndpointTestsDataThresholdFilters) GetConditionalOperator() ConditionalOperator {
	if o == nil || utils.IsNil(o.ConditionalOperator) {
		var ret ConditionalOperator
		return ret
	}
	return *o.ConditionalOperator
}

// GetConditionalOperatorOk returns a tuple with the ConditionalOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpEndpointTestsDataThresholdFilters) GetConditionalOperatorOk() (*ConditionalOperator, bool) {
	if o == nil || utils.IsNil(o.ConditionalOperator) {
		return nil, false
	}
	return o.ConditionalOperator, true
}

// HasConditionalOperator returns a boolean if a field has been set.
func (o *HttpEndpointTestsDataThresholdFilters) HasConditionalOperator() bool {
	if o != nil && !utils.IsNil(o.ConditionalOperator) {
		return true
	}

	return false
}

// SetConditionalOperator gets a reference to the given ConditionalOperator and assigns it to the ConditionalOperator field.
func (o *HttpEndpointTestsDataThresholdFilters) SetConditionalOperator(v ConditionalOperator) {
	o.ConditionalOperator = &v
}

func (o HttpEndpointTestsDataThresholdFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpEndpointTestsDataThresholdFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !utils.IsNil(o.ConditionalOperator) {
		toSerialize["conditionalOperator"] = o.ConditionalOperator
	}
	return toSerialize, nil
}

type NullableHttpEndpointTestsDataThresholdFilters struct {
	value *HttpEndpointTestsDataThresholdFilters
	isSet bool
}

func (v NullableHttpEndpointTestsDataThresholdFilters) Get() *HttpEndpointTestsDataThresholdFilters {
	return v.value
}

func (v *NullableHttpEndpointTestsDataThresholdFilters) Set(val *HttpEndpointTestsDataThresholdFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpEndpointTestsDataThresholdFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpEndpointTestsDataThresholdFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpEndpointTestsDataThresholdFilters(val *HttpEndpointTestsDataThresholdFilters) *NullableHttpEndpointTestsDataThresholdFilters {
	return &NullableHttpEndpointTestsDataThresholdFilters{value: val, isSet: true}
}

func (v NullableHttpEndpointTestsDataThresholdFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpEndpointTestsDataThresholdFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


