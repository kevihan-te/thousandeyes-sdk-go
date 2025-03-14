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

// checks if the EndpointTestLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointTestLinks{}

// EndpointTestLinks A list of links that can be accessed to get more information.
type EndpointTestLinks struct {
	Self *EndpointTestSelfLink `json:"self,omitempty"`
	// Reference to the test results.
	TestResults []Link `json:"testResults,omitempty"`
}

// NewEndpointTestLinks instantiates a new EndpointTestLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTestLinks() *EndpointTestLinks {
	this := EndpointTestLinks{}
	return &this
}

// NewEndpointTestLinksWithDefaults instantiates a new EndpointTestLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTestLinksWithDefaults() *EndpointTestLinks {
	this := EndpointTestLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EndpointTestLinks) GetSelf() EndpointTestSelfLink {
	if o == nil || utils.IsNil(o.Self) {
		var ret EndpointTestSelfLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTestLinks) GetSelfOk() (*EndpointTestSelfLink, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EndpointTestLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given EndpointTestSelfLink and assigns it to the Self field.
func (o *EndpointTestLinks) SetSelf(v EndpointTestSelfLink) {
	o.Self = &v
}

// GetTestResults returns the TestResults field value if set, zero value otherwise.
func (o *EndpointTestLinks) GetTestResults() []Link {
	if o == nil || utils.IsNil(o.TestResults) {
		var ret []Link
		return ret
	}
	return o.TestResults
}

// GetTestResultsOk returns a tuple with the TestResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTestLinks) GetTestResultsOk() ([]Link, bool) {
	if o == nil || utils.IsNil(o.TestResults) {
		return nil, false
	}
	return o.TestResults, true
}

// HasTestResults returns a boolean if a field has been set.
func (o *EndpointTestLinks) HasTestResults() bool {
	if o != nil && !utils.IsNil(o.TestResults) {
		return true
	}

	return false
}

// SetTestResults gets a reference to the given []Link and assigns it to the TestResults field.
func (o *EndpointTestLinks) SetTestResults(v []Link) {
	o.TestResults = v
}

func (o EndpointTestLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTestLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !utils.IsNil(o.TestResults) {
		toSerialize["testResults"] = o.TestResults
	}
	return toSerialize, nil
}

type NullableEndpointTestLinks struct {
	value *EndpointTestLinks
	isSet bool
}

func (v NullableEndpointTestLinks) Get() *EndpointTestLinks {
	return v.value
}

func (v *NullableEndpointTestLinks) Set(val *EndpointTestLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTestLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTestLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTestLinks(val *EndpointTestLinks) *NullableEndpointTestLinks {
	return &NullableEndpointTestLinks{value: val, isSet: true}
}

func (v NullableEndpointTestLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTestLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


