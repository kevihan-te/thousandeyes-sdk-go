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

// checks if the EndpointTestsDataThresholdFilters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointTestsDataThresholdFilters{}

// EndpointTestsDataThresholdFilters All filters are applied based on the conditional operator (and/or).
type EndpointTestsDataThresholdFilters struct {
	Filters []EndpointTestsDataThresholdFilter `json:"filters,omitempty"`
	ConditionalOperator *ConditionalOperator `json:"conditionalOperator,omitempty"`
}

// NewEndpointTestsDataThresholdFilters instantiates a new EndpointTestsDataThresholdFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTestsDataThresholdFilters() *EndpointTestsDataThresholdFilters {
	this := EndpointTestsDataThresholdFilters{}
	return &this
}

// NewEndpointTestsDataThresholdFiltersWithDefaults instantiates a new EndpointTestsDataThresholdFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTestsDataThresholdFiltersWithDefaults() *EndpointTestsDataThresholdFilters {
	this := EndpointTestsDataThresholdFilters{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *EndpointTestsDataThresholdFilters) GetFilters() []EndpointTestsDataThresholdFilter {
	if o == nil || utils.IsNil(o.Filters) {
		var ret []EndpointTestsDataThresholdFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTestsDataThresholdFilters) GetFiltersOk() ([]EndpointTestsDataThresholdFilter, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *EndpointTestsDataThresholdFilters) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []EndpointTestsDataThresholdFilter and assigns it to the Filters field.
func (o *EndpointTestsDataThresholdFilters) SetFilters(v []EndpointTestsDataThresholdFilter) {
	o.Filters = v
}

// GetConditionalOperator returns the ConditionalOperator field value if set, zero value otherwise.
func (o *EndpointTestsDataThresholdFilters) GetConditionalOperator() ConditionalOperator {
	if o == nil || utils.IsNil(o.ConditionalOperator) {
		var ret ConditionalOperator
		return ret
	}
	return *o.ConditionalOperator
}

// GetConditionalOperatorOk returns a tuple with the ConditionalOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTestsDataThresholdFilters) GetConditionalOperatorOk() (*ConditionalOperator, bool) {
	if o == nil || utils.IsNil(o.ConditionalOperator) {
		return nil, false
	}
	return o.ConditionalOperator, true
}

// HasConditionalOperator returns a boolean if a field has been set.
func (o *EndpointTestsDataThresholdFilters) HasConditionalOperator() bool {
	if o != nil && !utils.IsNil(o.ConditionalOperator) {
		return true
	}

	return false
}

// SetConditionalOperator gets a reference to the given ConditionalOperator and assigns it to the ConditionalOperator field.
func (o *EndpointTestsDataThresholdFilters) SetConditionalOperator(v ConditionalOperator) {
	o.ConditionalOperator = &v
}

func (o EndpointTestsDataThresholdFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTestsDataThresholdFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !utils.IsNil(o.ConditionalOperator) {
		toSerialize["conditionalOperator"] = o.ConditionalOperator
	}
	return toSerialize, nil
}

type NullableEndpointTestsDataThresholdFilters struct {
	value *EndpointTestsDataThresholdFilters
	isSet bool
}

func (v NullableEndpointTestsDataThresholdFilters) Get() *EndpointTestsDataThresholdFilters {
	return v.value
}

func (v *NullableEndpointTestsDataThresholdFilters) Set(val *EndpointTestsDataThresholdFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTestsDataThresholdFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTestsDataThresholdFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTestsDataThresholdFilters(val *EndpointTestsDataThresholdFilters) *NullableEndpointTestsDataThresholdFilters {
	return &NullableEndpointTestsDataThresholdFilters{value: val, isSet: true}
}

func (v NullableEndpointTestsDataThresholdFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTestsDataThresholdFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


