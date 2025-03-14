/*
ThousandEyes for OpenTelemetry API

ThousandEyes for OpenTelemetry provides machine-to-machine integration between ThousandEyes and its customers. It allows you to export ThousandEyes telemetry data in OTel format, which is widely used in the industry. With ThousandEyes for OTel, you can leverage frameworks widely used in the observability domain - such as Splunk, Grafana, and Honeycomb - to capture and analyze ThousandEyes data. Any client that supports OTel can use ThousandEyes for OpenTelemetry.  ThousandEyes for OTel is made up of the following components:  * Data streaming APIs that you can use to configure and enable your ThousandEyes tests with OTel-compatible streams, in particular to configure how ThousandEyes telemetry data is exported to client integrations. * A set of streaming pipelines called _collectors_ that actively fetch ThousandEyes network test data, enrich the data with some additional detail, filter, and push the data to the customer-configured endpoints, depending on what you configure via the public APIs. * Third-party OTel collectors that receive, transform, filter, and export different metrics to client applications such as AppD, or any other OTel-capable client configuration.  For more information about ThousandEyes for OpenTelemetry, see the [documentation](https://docs.thousandeyes.com/product-documentation/api/opentelemetry). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the Filters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Filters{}

// Filters Provides the ability to filter data points based on the specified test types.
type Filters struct {
	TestTypes *FiltersTestTypes `json:"testTypes,omitempty"`
}

// NewFilters instantiates a new Filters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilters() *Filters {
	this := Filters{}
	return &this
}

// NewFiltersWithDefaults instantiates a new Filters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiltersWithDefaults() *Filters {
	this := Filters{}
	return &this
}

// GetTestTypes returns the TestTypes field value if set, zero value otherwise.
func (o *Filters) GetTestTypes() FiltersTestTypes {
	if o == nil || utils.IsNil(o.TestTypes) {
		var ret FiltersTestTypes
		return ret
	}
	return *o.TestTypes
}

// GetTestTypesOk returns a tuple with the TestTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Filters) GetTestTypesOk() (*FiltersTestTypes, bool) {
	if o == nil || utils.IsNil(o.TestTypes) {
		return nil, false
	}
	return o.TestTypes, true
}

// HasTestTypes returns a boolean if a field has been set.
func (o *Filters) HasTestTypes() bool {
	if o != nil && !utils.IsNil(o.TestTypes) {
		return true
	}

	return false
}

// SetTestTypes gets a reference to the given FiltersTestTypes and assigns it to the TestTypes field.
func (o *Filters) SetTestTypes(v FiltersTestTypes) {
	o.TestTypes = &v
}

func (o Filters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Filters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TestTypes) {
		toSerialize["testTypes"] = o.TestTypes
	}
	return toSerialize, nil
}

type NullableFilters struct {
	value *Filters
	isSet bool
}

func (v NullableFilters) Get() *Filters {
	return v.value
}

func (v *NullableFilters) Set(val *Filters) {
	v.value = val
	v.isSet = true
}

func (v NullableFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilters(val *Filters) *NullableFilters {
	return &NullableFilters{value: val, isSet: true}
}

func (v NullableFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


