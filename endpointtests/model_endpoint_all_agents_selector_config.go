/*
Endpoint Tests API

 Manage endpoint agent dynamic and scheduled tests using the Endpoint Tests API. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the EndpointAllAgentsSelectorConfig type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointAllAgentsSelectorConfig{}

// EndpointAllAgentsSelectorConfig Any agent selection object.
type EndpointAllAgentsSelectorConfig struct {
	AgentSelectorType string `json:"agentSelectorType"`
	// Maximum number of agents which can execute the test.
	MaxMachines *int32 `json:"maxMachines,omitempty"`
}

type _EndpointAllAgentsSelectorConfig EndpointAllAgentsSelectorConfig

// NewEndpointAllAgentsSelectorConfig instantiates a new EndpointAllAgentsSelectorConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAllAgentsSelectorConfig(agentSelectorType string) *EndpointAllAgentsSelectorConfig {
	this := EndpointAllAgentsSelectorConfig{}
	this.AgentSelectorType = agentSelectorType
	var maxMachines int32 = 25
	this.MaxMachines = &maxMachines
	return &this
}

// NewEndpointAllAgentsSelectorConfigWithDefaults instantiates a new EndpointAllAgentsSelectorConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAllAgentsSelectorConfigWithDefaults() *EndpointAllAgentsSelectorConfig {
	this := EndpointAllAgentsSelectorConfig{}
	var maxMachines int32 = 25
	this.MaxMachines = &maxMachines
	return &this
}

// GetAgentSelectorType returns the AgentSelectorType field value
func (o *EndpointAllAgentsSelectorConfig) GetAgentSelectorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentSelectorType
}

// GetAgentSelectorTypeOk returns a tuple with the AgentSelectorType field value
// and a boolean to check if the value has been set.
func (o *EndpointAllAgentsSelectorConfig) GetAgentSelectorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentSelectorType, true
}

// SetAgentSelectorType sets field value
func (o *EndpointAllAgentsSelectorConfig) SetAgentSelectorType(v string) {
	o.AgentSelectorType = v
}

// GetMaxMachines returns the MaxMachines field value if set, zero value otherwise.
func (o *EndpointAllAgentsSelectorConfig) GetMaxMachines() int32 {
	if o == nil || utils.IsNil(o.MaxMachines) {
		var ret int32
		return ret
	}
	return *o.MaxMachines
}

// GetMaxMachinesOk returns a tuple with the MaxMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAllAgentsSelectorConfig) GetMaxMachinesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxMachines) {
		return nil, false
	}
	return o.MaxMachines, true
}

// HasMaxMachines returns a boolean if a field has been set.
func (o *EndpointAllAgentsSelectorConfig) HasMaxMachines() bool {
	if o != nil && !utils.IsNil(o.MaxMachines) {
		return true
	}

	return false
}

// SetMaxMachines gets a reference to the given int32 and assigns it to the MaxMachines field.
func (o *EndpointAllAgentsSelectorConfig) SetMaxMachines(v int32) {
	o.MaxMachines = &v
}

func (o EndpointAllAgentsSelectorConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAllAgentsSelectorConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agentSelectorType"] = o.AgentSelectorType
	if !utils.IsNil(o.MaxMachines) {
		toSerialize["maxMachines"] = o.MaxMachines
	}
	return toSerialize, nil
}

func (o *EndpointAllAgentsSelectorConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agentSelectorType",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEndpointAllAgentsSelectorConfig := _EndpointAllAgentsSelectorConfig{}

    err = json.Unmarshal(data, &varEndpointAllAgentsSelectorConfig)

	if err != nil {
		return err
	}

	*o = EndpointAllAgentsSelectorConfig(varEndpointAllAgentsSelectorConfig)

	return err
}

type NullableEndpointAllAgentsSelectorConfig struct {
	value *EndpointAllAgentsSelectorConfig
	isSet bool
}

func (v NullableEndpointAllAgentsSelectorConfig) Get() *EndpointAllAgentsSelectorConfig {
	return v.value
}

func (v *NullableEndpointAllAgentsSelectorConfig) Set(val *EndpointAllAgentsSelectorConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAllAgentsSelectorConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAllAgentsSelectorConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAllAgentsSelectorConfig(val *EndpointAllAgentsSelectorConfig) *NullableEndpointAllAgentsSelectorConfig {
	return &NullableEndpointAllAgentsSelectorConfig{value: val, isSet: true}
}

func (v NullableEndpointAllAgentsSelectorConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAllAgentsSelectorConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


