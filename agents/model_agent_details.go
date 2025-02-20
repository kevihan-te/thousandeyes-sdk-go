/*
Agents API

 ## Overview Manage Cloud and Enterprise Agents available to your account in ThousandEyes.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// AgentDetails - struct for AgentDetails
type AgentDetails struct {
	CloudAgentDetail *CloudAgentDetail
	EnterpriseAgentClusterDetail *EnterpriseAgentClusterDetail
	EnterpriseAgentDetail *EnterpriseAgentDetail
}

// CloudAgentDetailAsAgentDetails is a convenience function that returns CloudAgentDetail wrapped in AgentDetails
func CloudAgentDetailAsAgentDetails(v *CloudAgentDetail) AgentDetails {
	return AgentDetails{
		CloudAgentDetail: v,
	}
}

// EnterpriseAgentClusterDetailAsAgentDetails is a convenience function that returns EnterpriseAgentClusterDetail wrapped in AgentDetails
func EnterpriseAgentClusterDetailAsAgentDetails(v *EnterpriseAgentClusterDetail) AgentDetails {
	return AgentDetails{
		EnterpriseAgentClusterDetail: v,
	}
}

// EnterpriseAgentDetailAsAgentDetails is a convenience function that returns EnterpriseAgentDetail wrapped in AgentDetails
func EnterpriseAgentDetailAsAgentDetails(v *EnterpriseAgentDetail) AgentDetails {
	return AgentDetails{
		EnterpriseAgentDetail: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AgentDetails) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = utils.NewStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'cloud'
	if jsonDict["agentType"] == "cloud" {
		// try to unmarshal JSON data into CloudAgentDetail
		err = json.Unmarshal(data, &dst.CloudAgentDetail)
		if err == nil {
			return nil // data stored in dst.CloudAgentDetail, return on the first match
		} else {
			dst.CloudAgentDetail = nil
			return fmt.Errorf("failed to unmarshal AgentDetails as CloudAgentDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'enterprise'
	if jsonDict["agentType"] == "enterprise" {
		// try to unmarshal JSON data into EnterpriseAgentDetail
		err = json.Unmarshal(data, &dst.EnterpriseAgentDetail)
		if err == nil {
			return nil // data stored in dst.EnterpriseAgentDetail, return on the first match
		} else {
			dst.EnterpriseAgentDetail = nil
			return fmt.Errorf("failed to unmarshal AgentDetails as EnterpriseAgentDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'enterprise-cluster'
	if jsonDict["agentType"] == "enterprise-cluster" {
		// try to unmarshal JSON data into EnterpriseAgentClusterDetail
		err = json.Unmarshal(data, &dst.EnterpriseAgentClusterDetail)
		if err == nil {
			return nil // data stored in dst.EnterpriseAgentClusterDetail, return on the first match
		} else {
			dst.EnterpriseAgentClusterDetail = nil
			return fmt.Errorf("failed to unmarshal AgentDetails as EnterpriseAgentClusterDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'CloudAgentDetail'
	if jsonDict["agentType"] == "CloudAgentDetail" {
		// try to unmarshal JSON data into CloudAgentDetail
		err = json.Unmarshal(data, &dst.CloudAgentDetail)
		if err == nil {
			return nil // data stored in dst.CloudAgentDetail, return on the first match
		} else {
			dst.CloudAgentDetail = nil
			return fmt.Errorf("failed to unmarshal AgentDetails as CloudAgentDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EnterpriseAgentClusterDetail'
	if jsonDict["agentType"] == "EnterpriseAgentClusterDetail" {
		// try to unmarshal JSON data into EnterpriseAgentClusterDetail
		err = json.Unmarshal(data, &dst.EnterpriseAgentClusterDetail)
		if err == nil {
			return nil // data stored in dst.EnterpriseAgentClusterDetail, return on the first match
		} else {
			dst.EnterpriseAgentClusterDetail = nil
			return fmt.Errorf("failed to unmarshal AgentDetails as EnterpriseAgentClusterDetail: %s", err.Error())
		}
	}

	// check if the discriminator value is 'EnterpriseAgentDetail'
	if jsonDict["agentType"] == "EnterpriseAgentDetail" {
		// try to unmarshal JSON data into EnterpriseAgentDetail
		err = json.Unmarshal(data, &dst.EnterpriseAgentDetail)
		if err == nil {
			return nil // data stored in dst.EnterpriseAgentDetail, return on the first match
		} else {
			dst.EnterpriseAgentDetail = nil
			return fmt.Errorf("failed to unmarshal AgentDetails as EnterpriseAgentDetail: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AgentDetails) MarshalJSON() ([]byte, error) {
	if src.CloudAgentDetail != nil {
		return json.Marshal(&src.CloudAgentDetail)
	}

	if src.EnterpriseAgentClusterDetail != nil {
		return json.Marshal(&src.EnterpriseAgentClusterDetail)
	}

	if src.EnterpriseAgentDetail != nil {
		return json.Marshal(&src.EnterpriseAgentDetail)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AgentDetails) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CloudAgentDetail != nil {
		return obj.CloudAgentDetail
	}

	if obj.EnterpriseAgentClusterDetail != nil {
		return obj.EnterpriseAgentClusterDetail
	}

	if obj.EnterpriseAgentDetail != nil {
		return obj.EnterpriseAgentDetail
	}

	// all schemas are nil
	return nil
}

type NullableAgentDetails struct {
	value *AgentDetails
	isSet bool
}

func (v NullableAgentDetails) Get() *AgentDetails {
	return v.value
}

func (v *NullableAgentDetails) Set(val *AgentDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentDetails(val *AgentDetails) *NullableAgentDetails {
	return &NullableAgentDetails{value: val, isSet: true}
}

func (v NullableAgentDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


