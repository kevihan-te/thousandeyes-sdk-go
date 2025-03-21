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

// checks if the ExporterConfig type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExporterConfig{}

// ExporterConfig Capability to set exporter configuration.
type ExporterConfig struct {
	SplunkHec *ExporterConfigSplunkHec `json:"splunkHec,omitempty"`
}

// NewExporterConfig instantiates a new ExporterConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExporterConfig() *ExporterConfig {
	this := ExporterConfig{}
	return &this
}

// NewExporterConfigWithDefaults instantiates a new ExporterConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExporterConfigWithDefaults() *ExporterConfig {
	this := ExporterConfig{}
	return &this
}

// GetSplunkHec returns the SplunkHec field value if set, zero value otherwise.
func (o *ExporterConfig) GetSplunkHec() ExporterConfigSplunkHec {
	if o == nil || utils.IsNil(o.SplunkHec) {
		var ret ExporterConfigSplunkHec
		return ret
	}
	return *o.SplunkHec
}

// GetSplunkHecOk returns a tuple with the SplunkHec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterConfig) GetSplunkHecOk() (*ExporterConfigSplunkHec, bool) {
	if o == nil || utils.IsNil(o.SplunkHec) {
		return nil, false
	}
	return o.SplunkHec, true
}

// HasSplunkHec returns a boolean if a field has been set.
func (o *ExporterConfig) HasSplunkHec() bool {
	if o != nil && !utils.IsNil(o.SplunkHec) {
		return true
	}

	return false
}

// SetSplunkHec gets a reference to the given ExporterConfigSplunkHec and assigns it to the SplunkHec field.
func (o *ExporterConfig) SetSplunkHec(v ExporterConfigSplunkHec) {
	o.SplunkHec = &v
}

func (o ExporterConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExporterConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SplunkHec) {
		toSerialize["splunkHec"] = o.SplunkHec
	}
	return toSerialize, nil
}

type NullableExporterConfig struct {
	value *ExporterConfig
	isSet bool
}

func (v NullableExporterConfig) Get() *ExporterConfig {
	return v.value
}

func (v *NullableExporterConfig) Set(val *ExporterConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableExporterConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableExporterConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExporterConfig(val *ExporterConfig) *NullableExporterConfig {
	return &NullableExporterConfig{value: val, isSet: true}
}

func (v NullableExporterConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExporterConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


