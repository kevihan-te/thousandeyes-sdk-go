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

// checks if the StreamLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &StreamLinks{}

// StreamLinks struct for StreamLinks
type StreamLinks struct {
	Self *StreamSelfLink `json:"self,omitempty"`
}

// NewStreamLinks instantiates a new StreamLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamLinks() *StreamLinks {
	this := StreamLinks{}
	return &this
}

// NewStreamLinksWithDefaults instantiates a new StreamLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamLinksWithDefaults() *StreamLinks {
	this := StreamLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *StreamLinks) GetSelf() StreamSelfLink {
	if o == nil || utils.IsNil(o.Self) {
		var ret StreamSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamLinks) GetSelfOk() (*StreamSelfLink, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *StreamLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given StreamSelfLink and assigns it to the Self field.
func (o *StreamLinks) SetSelf(v StreamSelfLink) {
	o.Self = &v
}

func (o StreamLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableStreamLinks struct {
	value *StreamLinks
	isSet bool
}

func (v NullableStreamLinks) Get() *StreamLinks {
	return v.value
}

func (v *NullableStreamLinks) Set(val *StreamLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamLinks(val *StreamLinks) *NullableStreamLinks {
	return &NullableStreamLinks{value: val, isSet: true}
}

func (v NullableStreamLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


