/*
Endpoint Tests API

 Manage endpoint agent dynamic and scheduled tests using the Endpoint Tests API. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// EndpointAgentSelectorConfig - Agents selection object based on agentSelectorType.
type EndpointAgentSelectorConfig struct {
	EndpointAgentLabelsSelectorConfig *EndpointAgentLabelsSelectorConfig
	EndpointAllAgentsSelectorConfig *EndpointAllAgentsSelectorConfig
	EndpointSpecificAgentsSelectorConfig *EndpointSpecificAgentsSelectorConfig
}

// EndpointAgentLabelsSelectorConfigAsEndpointAgentSelectorConfig is a convenience function that returns EndpointAgentLabelsSelectorConfig wrapped in EndpointAgentSelectorConfig
func EndpointAgentLabelsSelectorConfigAsEndpointAgentSelectorConfig(v *EndpointAgentLabelsSelectorConfig) EndpointAgentSelectorConfig {
	return EndpointAgentSelectorConfig{
		EndpointAgentLabelsSelectorConfig: v,
	}
}

// EndpointAllAgentsSelectorConfigAsEndpointAgentSelectorConfig is a convenience function that returns EndpointAllAgentsSelectorConfig wrapped in EndpointAgentSelectorConfig
func EndpointAllAgentsSelectorConfigAsEndpointAgentSelectorConfig(v *EndpointAllAgentsSelectorConfig) EndpointAgentSelectorConfig {
	return EndpointAgentSelectorConfig{
		EndpointAllAgentsSelectorConfig: v,
	}
}

// EndpointSpecificAgentsSelectorConfigAsEndpointAgentSelectorConfig is a convenience function that returns EndpointSpecificAgentsSelectorConfig wrapped in EndpointAgentSelectorConfig
func EndpointSpecificAgentsSelectorConfigAsEndpointAgentSelectorConfig(v *EndpointSpecificAgentsSelectorConfig) EndpointAgentSelectorConfig {
	return EndpointAgentSelectorConfig{
		EndpointSpecificAgentsSelectorConfig: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EndpointAgentSelectorConfig) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = utils.NewStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'agent-labels'
	if jsonDict["agentSelectorType"] == "agent-labels" {
		// try to unmarshal JSON data into EndpointAgentLabelsSelectorConfig
		err = json.Unmarshal(data, &dst.EndpointAgentLabelsSelectorConfig)
		if err == nil {
			return nil // data stored in dst.EndpointAgentLabelsSelectorConfig, return on the first match
		} else {
			dst.EndpointAgentLabelsSelectorConfig = nil
			return fmt.Errorf("failed to unmarshal EndpointAgentSelectorConfig as EndpointAgentLabelsSelectorConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'all-agents'
	if jsonDict["agentSelectorType"] == "all-agents" {
		// try to unmarshal JSON data into EndpointAllAgentsSelectorConfig
		err = json.Unmarshal(data, &dst.EndpointAllAgentsSelectorConfig)
		if err == nil {
			return nil // data stored in dst.EndpointAllAgentsSelectorConfig, return on the first match
		} else {
			dst.EndpointAllAgentsSelectorConfig = nil
			return fmt.Errorf("failed to unmarshal EndpointAgentSelectorConfig as EndpointAllAgentsSelectorConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'specific-agents'
	if jsonDict["agentSelectorType"] == "specific-agents" {
		// try to unmarshal JSON data into EndpointSpecificAgentsSelectorConfig
		err = json.Unmarshal(data, &dst.EndpointSpecificAgentsSelectorConfig)
		if err == nil {
			return nil // data stored in dst.EndpointSpecificAgentsSelectorConfig, return on the first match
		} else {
			dst.EndpointSpecificAgentsSelectorConfig = nil
			return fmt.Errorf("failed to unmarshal EndpointAgentSelectorConfig as EndpointSpecificAgentsSelectorConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EndpointAgentLabelsSelectorConfig'
	if jsonDict["agentSelectorType"] == "EndpointAgentLabelsSelectorConfig" {
		// try to unmarshal JSON data into EndpointAgentLabelsSelectorConfig
		err = json.Unmarshal(data, &dst.EndpointAgentLabelsSelectorConfig)
		if err == nil {
			return nil // data stored in dst.EndpointAgentLabelsSelectorConfig, return on the first match
		} else {
			dst.EndpointAgentLabelsSelectorConfig = nil
			return fmt.Errorf("failed to unmarshal EndpointAgentSelectorConfig as EndpointAgentLabelsSelectorConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EndpointAllAgentsSelectorConfig'
	if jsonDict["agentSelectorType"] == "EndpointAllAgentsSelectorConfig" {
		// try to unmarshal JSON data into EndpointAllAgentsSelectorConfig
		err = json.Unmarshal(data, &dst.EndpointAllAgentsSelectorConfig)
		if err == nil {
			return nil // data stored in dst.EndpointAllAgentsSelectorConfig, return on the first match
		} else {
			dst.EndpointAllAgentsSelectorConfig = nil
			return fmt.Errorf("failed to unmarshal EndpointAgentSelectorConfig as EndpointAllAgentsSelectorConfig: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EndpointSpecificAgentsSelectorConfig'
	if jsonDict["agentSelectorType"] == "EndpointSpecificAgentsSelectorConfig" {
		// try to unmarshal JSON data into EndpointSpecificAgentsSelectorConfig
		err = json.Unmarshal(data, &dst.EndpointSpecificAgentsSelectorConfig)
		if err == nil {
			return nil // data stored in dst.EndpointSpecificAgentsSelectorConfig, return on the first match
		} else {
			dst.EndpointSpecificAgentsSelectorConfig = nil
			return fmt.Errorf("failed to unmarshal EndpointAgentSelectorConfig as EndpointSpecificAgentsSelectorConfig: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EndpointAgentSelectorConfig) MarshalJSON() ([]byte, error) {
	if src.EndpointAgentLabelsSelectorConfig != nil {
		return json.Marshal(&src.EndpointAgentLabelsSelectorConfig)
	}

	if src.EndpointAllAgentsSelectorConfig != nil {
		return json.Marshal(&src.EndpointAllAgentsSelectorConfig)
	}

	if src.EndpointSpecificAgentsSelectorConfig != nil {
		return json.Marshal(&src.EndpointSpecificAgentsSelectorConfig)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EndpointAgentSelectorConfig) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.EndpointAgentLabelsSelectorConfig != nil {
		return obj.EndpointAgentLabelsSelectorConfig
	}

	if obj.EndpointAllAgentsSelectorConfig != nil {
		return obj.EndpointAllAgentsSelectorConfig
	}

	if obj.EndpointSpecificAgentsSelectorConfig != nil {
		return obj.EndpointSpecificAgentsSelectorConfig
	}

	// all schemas are nil
	return nil
}

type NullableEndpointAgentSelectorConfig struct {
	value *EndpointAgentSelectorConfig
	isSet bool
}

func (v NullableEndpointAgentSelectorConfig) Get() *EndpointAgentSelectorConfig {
	return v.value
}

func (v *NullableEndpointAgentSelectorConfig) Set(val *EndpointAgentSelectorConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAgentSelectorConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAgentSelectorConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAgentSelectorConfig(val *EndpointAgentSelectorConfig) *NullableEndpointAgentSelectorConfig {
	return &NullableEndpointAgentSelectorConfig{value: val, isSet: true}
}

func (v NullableEndpointAgentSelectorConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAgentSelectorConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


