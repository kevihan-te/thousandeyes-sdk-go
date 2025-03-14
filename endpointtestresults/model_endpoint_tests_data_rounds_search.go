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

// checks if the EndpointTestsDataRoundsSearch type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointTestsDataRoundsSearch{}

// EndpointTestsDataRoundsSearch struct for EndpointTestsDataRoundsSearch
type EndpointTestsDataRoundsSearch struct {
	SearchSort []EndpointTestsDataSearchSort `json:"searchSort,omitempty"`
	ThresholdFilter *EndpointTestsDataThresholdFilters `json:"thresholdFilter,omitempty"`
	SearchFilters *EndpointTestsDataSearchFilter `json:"searchFilters,omitempty"`
}

// NewEndpointTestsDataRoundsSearch instantiates a new EndpointTestsDataRoundsSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTestsDataRoundsSearch() *EndpointTestsDataRoundsSearch {
	this := EndpointTestsDataRoundsSearch{}
	return &this
}

// NewEndpointTestsDataRoundsSearchWithDefaults instantiates a new EndpointTestsDataRoundsSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTestsDataRoundsSearchWithDefaults() *EndpointTestsDataRoundsSearch {
	this := EndpointTestsDataRoundsSearch{}
	return &this
}

// GetSearchSort returns the SearchSort field value if set, zero value otherwise.
func (o *EndpointTestsDataRoundsSearch) GetSearchSort() []EndpointTestsDataSearchSort {
	if o == nil || utils.IsNil(o.SearchSort) {
		var ret []EndpointTestsDataSearchSort
		return ret
	}
	return o.SearchSort
}

// GetSearchSortOk returns a tuple with the SearchSort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTestsDataRoundsSearch) GetSearchSortOk() ([]EndpointTestsDataSearchSort, bool) {
	if o == nil || utils.IsNil(o.SearchSort) {
		return nil, false
	}
	return o.SearchSort, true
}

// HasSearchSort returns a boolean if a field has been set.
func (o *EndpointTestsDataRoundsSearch) HasSearchSort() bool {
	if o != nil && !utils.IsNil(o.SearchSort) {
		return true
	}

	return false
}

// SetSearchSort gets a reference to the given []EndpointTestsDataSearchSort and assigns it to the SearchSort field.
func (o *EndpointTestsDataRoundsSearch) SetSearchSort(v []EndpointTestsDataSearchSort) {
	o.SearchSort = v
}

// GetThresholdFilter returns the ThresholdFilter field value if set, zero value otherwise.
func (o *EndpointTestsDataRoundsSearch) GetThresholdFilter() EndpointTestsDataThresholdFilters {
	if o == nil || utils.IsNil(o.ThresholdFilter) {
		var ret EndpointTestsDataThresholdFilters
		return ret
	}
	return *o.ThresholdFilter
}

// GetThresholdFilterOk returns a tuple with the ThresholdFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTestsDataRoundsSearch) GetThresholdFilterOk() (*EndpointTestsDataThresholdFilters, bool) {
	if o == nil || utils.IsNil(o.ThresholdFilter) {
		return nil, false
	}
	return o.ThresholdFilter, true
}

// HasThresholdFilter returns a boolean if a field has been set.
func (o *EndpointTestsDataRoundsSearch) HasThresholdFilter() bool {
	if o != nil && !utils.IsNil(o.ThresholdFilter) {
		return true
	}

	return false
}

// SetThresholdFilter gets a reference to the given EndpointTestsDataThresholdFilters and assigns it to the ThresholdFilter field.
func (o *EndpointTestsDataRoundsSearch) SetThresholdFilter(v EndpointTestsDataThresholdFilters) {
	o.ThresholdFilter = &v
}

// GetSearchFilters returns the SearchFilters field value if set, zero value otherwise.
func (o *EndpointTestsDataRoundsSearch) GetSearchFilters() EndpointTestsDataSearchFilter {
	if o == nil || utils.IsNil(o.SearchFilters) {
		var ret EndpointTestsDataSearchFilter
		return ret
	}
	return *o.SearchFilters
}

// GetSearchFiltersOk returns a tuple with the SearchFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTestsDataRoundsSearch) GetSearchFiltersOk() (*EndpointTestsDataSearchFilter, bool) {
	if o == nil || utils.IsNil(o.SearchFilters) {
		return nil, false
	}
	return o.SearchFilters, true
}

// HasSearchFilters returns a boolean if a field has been set.
func (o *EndpointTestsDataRoundsSearch) HasSearchFilters() bool {
	if o != nil && !utils.IsNil(o.SearchFilters) {
		return true
	}

	return false
}

// SetSearchFilters gets a reference to the given EndpointTestsDataSearchFilter and assigns it to the SearchFilters field.
func (o *EndpointTestsDataRoundsSearch) SetSearchFilters(v EndpointTestsDataSearchFilter) {
	o.SearchFilters = &v
}

func (o EndpointTestsDataRoundsSearch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTestsDataRoundsSearch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SearchSort) {
		toSerialize["searchSort"] = o.SearchSort
	}
	if !utils.IsNil(o.ThresholdFilter) {
		toSerialize["thresholdFilter"] = o.ThresholdFilter
	}
	if !utils.IsNil(o.SearchFilters) {
		toSerialize["searchFilters"] = o.SearchFilters
	}
	return toSerialize, nil
}

type NullableEndpointTestsDataRoundsSearch struct {
	value *EndpointTestsDataRoundsSearch
	isSet bool
}

func (v NullableEndpointTestsDataRoundsSearch) Get() *EndpointTestsDataRoundsSearch {
	return v.value
}

func (v *NullableEndpointTestsDataRoundsSearch) Set(val *EndpointTestsDataRoundsSearch) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTestsDataRoundsSearch) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTestsDataRoundsSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTestsDataRoundsSearch(val *EndpointTestsDataRoundsSearch) *NullableEndpointTestsDataRoundsSearch {
	return &NullableEndpointTestsDataRoundsSearch{value: val, isSet: true}
}

func (v NullableEndpointTestsDataRoundsSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTestsDataRoundsSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


