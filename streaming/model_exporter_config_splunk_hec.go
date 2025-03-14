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

// checks if the ExporterConfigSplunkHec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExporterConfigSplunkHec{}

// ExporterConfigSplunkHec Splunk HEC configuration. This can only be configured when the `type` is `splunk-hec`.
type ExporterConfigSplunkHec struct {
	// The Splunk HEC `token`. This is a required field.
	Token *string `json:"token,omitempty"`
	// The Splunk HEC `source`. This field is optional.
	Source *string `json:"source,omitempty"`
	// The Splunk HEC `sourceType`. This field is optional.
	SourceType *string `json:"sourceType,omitempty"`
	// The name of the Splunk HEC index where the event data will be stored.  This field is optional.
	Index *string `json:"index,omitempty"`
}

// NewExporterConfigSplunkHec instantiates a new ExporterConfigSplunkHec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExporterConfigSplunkHec() *ExporterConfigSplunkHec {
	this := ExporterConfigSplunkHec{}
	var source string = "ThousandEyesOTel"
	this.Source = &source
	var sourceType string = "ThousandEyesOTel"
	this.SourceType = &sourceType
	return &this
}

// NewExporterConfigSplunkHecWithDefaults instantiates a new ExporterConfigSplunkHec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExporterConfigSplunkHecWithDefaults() *ExporterConfigSplunkHec {
	this := ExporterConfigSplunkHec{}
	var source string = "ThousandEyesOTel"
	this.Source = &source
	var sourceType string = "ThousandEyesOTel"
	this.SourceType = &sourceType
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ExporterConfigSplunkHec) GetToken() string {
	if o == nil || utils.IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterConfigSplunkHec) GetTokenOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ExporterConfigSplunkHec) HasToken() bool {
	if o != nil && !utils.IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ExporterConfigSplunkHec) SetToken(v string) {
	o.Token = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ExporterConfigSplunkHec) GetSource() string {
	if o == nil || utils.IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterConfigSplunkHec) GetSourceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ExporterConfigSplunkHec) HasSource() bool {
	if o != nil && !utils.IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ExporterConfigSplunkHec) SetSource(v string) {
	o.Source = &v
}

// GetSourceType returns the SourceType field value if set, zero value otherwise.
func (o *ExporterConfigSplunkHec) GetSourceType() string {
	if o == nil || utils.IsNil(o.SourceType) {
		var ret string
		return ret
	}
	return *o.SourceType
}

// GetSourceTypeOk returns a tuple with the SourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterConfigSplunkHec) GetSourceTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceType) {
		return nil, false
	}
	return o.SourceType, true
}

// HasSourceType returns a boolean if a field has been set.
func (o *ExporterConfigSplunkHec) HasSourceType() bool {
	if o != nil && !utils.IsNil(o.SourceType) {
		return true
	}

	return false
}

// SetSourceType gets a reference to the given string and assigns it to the SourceType field.
func (o *ExporterConfigSplunkHec) SetSourceType(v string) {
	o.SourceType = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ExporterConfigSplunkHec) GetIndex() string {
	if o == nil || utils.IsNil(o.Index) {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExporterConfigSplunkHec) GetIndexOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *ExporterConfigSplunkHec) HasIndex() bool {
	if o != nil && !utils.IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *ExporterConfigSplunkHec) SetIndex(v string) {
	o.Index = &v
}

func (o ExporterConfigSplunkHec) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExporterConfigSplunkHec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !utils.IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !utils.IsNil(o.SourceType) {
		toSerialize["sourceType"] = o.SourceType
	}
	if !utils.IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	return toSerialize, nil
}

type NullableExporterConfigSplunkHec struct {
	value *ExporterConfigSplunkHec
	isSet bool
}

func (v NullableExporterConfigSplunkHec) Get() *ExporterConfigSplunkHec {
	return v.value
}

func (v *NullableExporterConfigSplunkHec) Set(val *ExporterConfigSplunkHec) {
	v.value = val
	v.isSet = true
}

func (v NullableExporterConfigSplunkHec) IsSet() bool {
	return v.isSet
}

func (v *NullableExporterConfigSplunkHec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExporterConfigSplunkHec(val *ExporterConfigSplunkHec) *NullableExporterConfigSplunkHec {
	return &NullableExporterConfigSplunkHec{value: val, isSet: true}
}

func (v NullableExporterConfigSplunkHec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExporterConfigSplunkHec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


