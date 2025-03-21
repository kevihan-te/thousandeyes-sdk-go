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

// checks if the TagMatch type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TagMatch{}

// TagMatch struct for TagMatch
type TagMatch struct {
	// The name of the tag key to match
	Key *string `json:"key,omitempty"`
	// The value of the tag to match
	Value *string `json:"value,omitempty"`
}

// NewTagMatch instantiates a new TagMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagMatch() *TagMatch {
	this := TagMatch{}
	return &this
}

// NewTagMatchWithDefaults instantiates a new TagMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagMatchWithDefaults() *TagMatch {
	this := TagMatch{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TagMatch) GetKey() string {
	if o == nil || utils.IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagMatch) GetKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TagMatch) HasKey() bool {
	if o != nil && !utils.IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TagMatch) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TagMatch) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagMatch) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TagMatch) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TagMatch) SetValue(v string) {
	o.Value = &v
}

func (o TagMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTagMatch struct {
	value *TagMatch
	isSet bool
}

func (v NullableTagMatch) Get() *TagMatch {
	return v.value
}

func (v *NullableTagMatch) Set(val *TagMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTagMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTagMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagMatch(val *TagMatch) *NullableTagMatch {
	return &NullableTagMatch{value: val, isSet: true}
}

func (v NullableTagMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


