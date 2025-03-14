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

// checks if the MultiTestIdEndpointTestsDataSearchFilter type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MultiTestIdEndpointTestsDataSearchFilter{}

// MultiTestIdEndpointTestsDataSearchFilter struct for MultiTestIdEndpointTestsDataSearchFilter
type MultiTestIdEndpointTestsDataSearchFilter struct {
	// Filter using the `agent-id`.
	AgentId []string `json:"agentId,omitempty"`
	TestId []string `json:"testId,omitempty"`
}

// NewMultiTestIdEndpointTestsDataSearchFilter instantiates a new MultiTestIdEndpointTestsDataSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiTestIdEndpointTestsDataSearchFilter() *MultiTestIdEndpointTestsDataSearchFilter {
	this := MultiTestIdEndpointTestsDataSearchFilter{}
	return &this
}

// NewMultiTestIdEndpointTestsDataSearchFilterWithDefaults instantiates a new MultiTestIdEndpointTestsDataSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiTestIdEndpointTestsDataSearchFilterWithDefaults() *MultiTestIdEndpointTestsDataSearchFilter {
	this := MultiTestIdEndpointTestsDataSearchFilter{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *MultiTestIdEndpointTestsDataSearchFilter) GetAgentId() []string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret []string
		return ret
	}
	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiTestIdEndpointTestsDataSearchFilter) GetAgentIdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *MultiTestIdEndpointTestsDataSearchFilter) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given []string and assigns it to the AgentId field.
func (o *MultiTestIdEndpointTestsDataSearchFilter) SetAgentId(v []string) {
	o.AgentId = v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *MultiTestIdEndpointTestsDataSearchFilter) GetTestId() []string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret []string
		return ret
	}
	return o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiTestIdEndpointTestsDataSearchFilter) GetTestIdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *MultiTestIdEndpointTestsDataSearchFilter) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given []string and assigns it to the TestId field.
func (o *MultiTestIdEndpointTestsDataSearchFilter) SetTestId(v []string) {
	o.TestId = v
}

func (o MultiTestIdEndpointTestsDataSearchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultiTestIdEndpointTestsDataSearchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	return toSerialize, nil
}

type NullableMultiTestIdEndpointTestsDataSearchFilter struct {
	value *MultiTestIdEndpointTestsDataSearchFilter
	isSet bool
}

func (v NullableMultiTestIdEndpointTestsDataSearchFilter) Get() *MultiTestIdEndpointTestsDataSearchFilter {
	return v.value
}

func (v *NullableMultiTestIdEndpointTestsDataSearchFilter) Set(val *MultiTestIdEndpointTestsDataSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiTestIdEndpointTestsDataSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiTestIdEndpointTestsDataSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiTestIdEndpointTestsDataSearchFilter(val *MultiTestIdEndpointTestsDataSearchFilter) *NullableMultiTestIdEndpointTestsDataSearchFilter {
	return &NullableMultiTestIdEndpointTestsDataSearchFilter{value: val, isSet: true}
}

func (v NullableMultiTestIdEndpointTestsDataSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiTestIdEndpointTestsDataSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


