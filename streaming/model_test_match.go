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

// checks if the TestMatch type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestMatch{}

// TestMatch struct for TestMatch
type TestMatch struct {
	// The ID of the test to match.
	Id *string `json:"id,omitempty"`
	Domain *TestMatchDomain `json:"domain,omitempty"`
}

// NewTestMatch instantiates a new TestMatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestMatch() *TestMatch {
	this := TestMatch{}
	return &this
}

// NewTestMatchWithDefaults instantiates a new TestMatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestMatchWithDefaults() *TestMatch {
	this := TestMatch{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TestMatch) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestMatch) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TestMatch) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TestMatch) SetId(v string) {
	o.Id = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *TestMatch) GetDomain() TestMatchDomain {
	if o == nil || utils.IsNil(o.Domain) {
		var ret TestMatchDomain
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestMatch) GetDomainOk() (*TestMatchDomain, bool) {
	if o == nil || utils.IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *TestMatch) HasDomain() bool {
	if o != nil && !utils.IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given TestMatchDomain and assigns it to the Domain field.
func (o *TestMatch) SetDomain(v TestMatchDomain) {
	o.Domain = &v
}

func (o TestMatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestMatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	return toSerialize, nil
}

type NullableTestMatch struct {
	value *TestMatch
	isSet bool
}

func (v NullableTestMatch) Get() *TestMatch {
	return v.value
}

func (v *NullableTestMatch) Set(val *TestMatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTestMatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTestMatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestMatch(val *TestMatch) *NullableTestMatch {
	return &NullableTestMatch{value: val, isSet: true}
}

func (v NullableTestMatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestMatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


