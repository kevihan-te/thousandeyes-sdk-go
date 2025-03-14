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

// checks if the EndpointAgents type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointAgents{}

// EndpointAgents A list of `EndpointAgents`.
type EndpointAgents struct {
	// The total number of agents.
	TotalAgents *int32 `json:"totalAgents,omitempty"`
	Agents []EndpointAgent `json:"agents,omitempty"`
}

// NewEndpointAgents instantiates a new EndpointAgents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAgents() *EndpointAgents {
	this := EndpointAgents{}
	return &this
}

// NewEndpointAgentsWithDefaults instantiates a new EndpointAgents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAgentsWithDefaults() *EndpointAgents {
	this := EndpointAgents{}
	return &this
}

// GetTotalAgents returns the TotalAgents field value if set, zero value otherwise.
func (o *EndpointAgents) GetTotalAgents() int32 {
	if o == nil || utils.IsNil(o.TotalAgents) {
		var ret int32
		return ret
	}
	return *o.TotalAgents
}

// GetTotalAgentsOk returns a tuple with the TotalAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgents) GetTotalAgentsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.TotalAgents) {
		return nil, false
	}
	return o.TotalAgents, true
}

// HasTotalAgents returns a boolean if a field has been set.
func (o *EndpointAgents) HasTotalAgents() bool {
	if o != nil && !utils.IsNil(o.TotalAgents) {
		return true
	}

	return false
}

// SetTotalAgents gets a reference to the given int32 and assigns it to the TotalAgents field.
func (o *EndpointAgents) SetTotalAgents(v int32) {
	o.TotalAgents = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *EndpointAgents) GetAgents() []EndpointAgent {
	if o == nil || utils.IsNil(o.Agents) {
		var ret []EndpointAgent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgents) GetAgentsOk() ([]EndpointAgent, bool) {
	if o == nil || utils.IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *EndpointAgents) HasAgents() bool {
	if o != nil && !utils.IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []EndpointAgent and assigns it to the Agents field.
func (o *EndpointAgents) SetAgents(v []EndpointAgent) {
	o.Agents = v
}

func (o EndpointAgents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAgents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TotalAgents) {
		toSerialize["totalAgents"] = o.TotalAgents
	}
	if !utils.IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	return toSerialize, nil
}

type NullableEndpointAgents struct {
	value *EndpointAgents
	isSet bool
}

func (v NullableEndpointAgents) Get() *EndpointAgents {
	return v.value
}

func (v *NullableEndpointAgents) Set(val *EndpointAgents) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAgents) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAgents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAgents(val *EndpointAgents) *NullableEndpointAgents {
	return &NullableEndpointAgents{value: val, isSet: true}
}

func (v NullableEndpointAgents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAgents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


