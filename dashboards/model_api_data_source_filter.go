/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the ApiDataSourceFilter type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiDataSourceFilter{}

// ApiDataSourceFilter List of different filter properties for a single datasource.
type ApiDataSourceFilter struct {
	// Data source property to filter by.
	FilterId string `json:"filterId"`
	// Values to filter by based on the specified `filterId`.
	Values []string `json:"values"`
	// Dashboard metric associated with the filter property.
	MetricIds []string `json:"metricIds"`
}

type _ApiDataSourceFilter ApiDataSourceFilter

// NewApiDataSourceFilter instantiates a new ApiDataSourceFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDataSourceFilter(filterId string, values []string, metricIds []string) *ApiDataSourceFilter {
	this := ApiDataSourceFilter{}
	this.FilterId = filterId
	this.Values = values
	this.MetricIds = metricIds
	return &this
}

// NewApiDataSourceFilterWithDefaults instantiates a new ApiDataSourceFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDataSourceFilterWithDefaults() *ApiDataSourceFilter {
	this := ApiDataSourceFilter{}
	return &this
}

// GetFilterId returns the FilterId field value
func (o *ApiDataSourceFilter) GetFilterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FilterId
}

// GetFilterIdOk returns a tuple with the FilterId field value
// and a boolean to check if the value has been set.
func (o *ApiDataSourceFilter) GetFilterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterId, true
}

// SetFilterId sets field value
func (o *ApiDataSourceFilter) SetFilterId(v string) {
	o.FilterId = v
}

// GetValues returns the Values field value
func (o *ApiDataSourceFilter) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *ApiDataSourceFilter) GetValuesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ApiDataSourceFilter) SetValues(v []string) {
	o.Values = v
}

// GetMetricIds returns the MetricIds field value
func (o *ApiDataSourceFilter) GetMetricIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.MetricIds
}

// GetMetricIdsOk returns a tuple with the MetricIds field value
// and a boolean to check if the value has been set.
func (o *ApiDataSourceFilter) GetMetricIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MetricIds, true
}

// SetMetricIds sets field value
func (o *ApiDataSourceFilter) SetMetricIds(v []string) {
	o.MetricIds = v
}

func (o ApiDataSourceFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDataSourceFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filterId"] = o.FilterId
	toSerialize["values"] = o.Values
	toSerialize["metricIds"] = o.MetricIds
	return toSerialize, nil
}

func (o *ApiDataSourceFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filterId",
		"values",
		"metricIds",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiDataSourceFilter := _ApiDataSourceFilter{}

    err = json.Unmarshal(data, &varApiDataSourceFilter)

	if err != nil {
		return err
	}

	*o = ApiDataSourceFilter(varApiDataSourceFilter)

	return err
}

type NullableApiDataSourceFilter struct {
	value *ApiDataSourceFilter
	isSet bool
}

func (v NullableApiDataSourceFilter) Get() *ApiDataSourceFilter {
	return v.value
}

func (v *NullableApiDataSourceFilter) Set(val *ApiDataSourceFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDataSourceFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDataSourceFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDataSourceFilter(val *ApiDataSourceFilter) *NullableApiDataSourceFilter {
	return &NullableApiDataSourceFilter{value: val, isSet: true}
}

func (v NullableApiDataSourceFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDataSourceFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


