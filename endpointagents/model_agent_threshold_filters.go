/*
Endpoint Agents API

Manage ThousandEyes Endpoint Agents using this API.   For more information about Endpoint Agents, see [Endpoint Agents](https://docs.thousandeyes.com/product-documentation/global-vantage-points/endpoint-agents).

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointagents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AgentThresholdFilters type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentThresholdFilters{}

// AgentThresholdFilters All filters are applied based on the conditional operator (and/or).
type AgentThresholdFilters struct {
	Filters []AgentThresholdFilter `json:"filters,omitempty"`
	ConditionalOperator *ConditionalOperator `json:"conditionalOperator,omitempty"`
}

// NewAgentThresholdFilters instantiates a new AgentThresholdFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentThresholdFilters() *AgentThresholdFilters {
	this := AgentThresholdFilters{}
	return &this
}

// NewAgentThresholdFiltersWithDefaults instantiates a new AgentThresholdFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentThresholdFiltersWithDefaults() *AgentThresholdFilters {
	this := AgentThresholdFilters{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *AgentThresholdFilters) GetFilters() []AgentThresholdFilter {
	if o == nil || utils.IsNil(o.Filters) {
		var ret []AgentThresholdFilter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentThresholdFilters) GetFiltersOk() ([]AgentThresholdFilter, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *AgentThresholdFilters) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []AgentThresholdFilter and assigns it to the Filters field.
func (o *AgentThresholdFilters) SetFilters(v []AgentThresholdFilter) {
	o.Filters = v
}

// GetConditionalOperator returns the ConditionalOperator field value if set, zero value otherwise.
func (o *AgentThresholdFilters) GetConditionalOperator() ConditionalOperator {
	if o == nil || utils.IsNil(o.ConditionalOperator) {
		var ret ConditionalOperator
		return ret
	}
	return *o.ConditionalOperator
}

// GetConditionalOperatorOk returns a tuple with the ConditionalOperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentThresholdFilters) GetConditionalOperatorOk() (*ConditionalOperator, bool) {
	if o == nil || utils.IsNil(o.ConditionalOperator) {
		return nil, false
	}
	return o.ConditionalOperator, true
}

// HasConditionalOperator returns a boolean if a field has been set.
func (o *AgentThresholdFilters) HasConditionalOperator() bool {
	if o != nil && !utils.IsNil(o.ConditionalOperator) {
		return true
	}

	return false
}

// SetConditionalOperator gets a reference to the given ConditionalOperator and assigns it to the ConditionalOperator field.
func (o *AgentThresholdFilters) SetConditionalOperator(v ConditionalOperator) {
	o.ConditionalOperator = &v
}

func (o AgentThresholdFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentThresholdFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !utils.IsNil(o.ConditionalOperator) {
		toSerialize["conditionalOperator"] = o.ConditionalOperator
	}
	return toSerialize, nil
}

type NullableAgentThresholdFilters struct {
	value *AgentThresholdFilters
	isSet bool
}

func (v NullableAgentThresholdFilters) Get() *AgentThresholdFilters {
	return v.value
}

func (v *NullableAgentThresholdFilters) Set(val *AgentThresholdFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentThresholdFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentThresholdFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentThresholdFilters(val *AgentThresholdFilters) *NullableAgentThresholdFilters {
	return &NullableAgentThresholdFilters{value: val, isSet: true}
}

func (v NullableAgentThresholdFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentThresholdFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


