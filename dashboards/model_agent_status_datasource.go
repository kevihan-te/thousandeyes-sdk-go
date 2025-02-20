/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
	"fmt"
)

// AgentStatusDatasource Datasource of the agent to retrieve status.
type AgentStatusDatasource string

// List of AgentStatusDatasource
const (
	AGENTSTATUSDATASOURCE_ALERTS AgentStatusDatasource = "ALERTS"
	AGENTSTATUSDATASOURCE_DEVICES AgentStatusDatasource = "DEVICES"
	AGENTSTATUSDATASOURCE_DNSP AgentStatusDatasource = "DNSP"
	AGENTSTATUSDATASOURCE_ENDPOINT_AGENTS AgentStatusDatasource = "ENDPOINT_AGENTS"
	AGENTSTATUSDATASOURCE_ENDPOINT_SCHEDULED_TEST AgentStatusDatasource = "ENDPOINT_SCHEDULED_TEST"
	AGENTSTATUSDATASOURCE_ENDPOINT_AST_TEST AgentStatusDatasource = "ENDPOINT_AST_TEST"
	AGENTSTATUSDATASOURCE_ENDPOINT_BROWSER_SESSION AgentStatusDatasource = "ENDPOINT_BROWSER_SESSION"
	AGENTSTATUSDATASOURCE_ENDPOINT_LOCAL_NETWORK AgentStatusDatasource = "ENDPOINT_LOCAL_NETWORK"
	AGENTSTATUSDATASOURCE_ENDPOINT_LOCAL_NETWORK_WIRELESS AgentStatusDatasource = "ENDPOINT_LOCAL_NETWORK_WIRELESS"
	AGENTSTATUSDATASOURCE_ROUTING AgentStatusDatasource = "ROUTING"
	AGENTSTATUSDATASOURCE_CLOUD_AND_ENTERPRISE_AGENTS AgentStatusDatasource = "CLOUD_AND_ENTERPRISE_AGENTS"
	AGENTSTATUSDATASOURCE_INTERNET_INSIGHTS AgentStatusDatasource = "INTERNET_INSIGHTS"
	AGENTSTATUSDATASOURCE_APPDYNAMICS_SERVICE_HEALTH AgentStatusDatasource = "APPDYNAMICS_SERVICE_HEALTH"
	AGENTSTATUSDATASOURCE_CLOUD_NATIVE_MONITORING AgentStatusDatasource = "CLOUD_NATIVE_MONITORING"
)

// All allowed values of AgentStatusDatasource enum
var AllowedAgentStatusDatasourceEnumValues = []AgentStatusDatasource{
	"ALERTS",
	"DEVICES",
	"DNSP",
	"ENDPOINT_AGENTS",
	"ENDPOINT_SCHEDULED_TEST",
	"ENDPOINT_AST_TEST",
	"ENDPOINT_BROWSER_SESSION",
	"ENDPOINT_LOCAL_NETWORK",
	"ENDPOINT_LOCAL_NETWORK_WIRELESS",
	"ROUTING",
	"CLOUD_AND_ENTERPRISE_AGENTS",
	"INTERNET_INSIGHTS",
	"APPDYNAMICS_SERVICE_HEALTH",
	"CLOUD_NATIVE_MONITORING",
}

func (v *AgentStatusDatasource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AgentStatusDatasource(value)
	for _, existing := range AllowedAgentStatusDatasourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AgentStatusDatasource", value)
}

// NewAgentStatusDatasourceFromValue returns a pointer to a valid AgentStatusDatasource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAgentStatusDatasourceFromValue(v string) (*AgentStatusDatasource, error) {
	ev := AgentStatusDatasource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AgentStatusDatasource: valid values are %v", v, AllowedAgentStatusDatasourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AgentStatusDatasource) IsValid() bool {
	for _, existing := range AllowedAgentStatusDatasourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AgentStatusDatasource value
func (v AgentStatusDatasource) Ptr() *AgentStatusDatasource {
	return &v
}

type NullableAgentStatusDatasource struct {
	value *AgentStatusDatasource
	isSet bool
}

func (v NullableAgentStatusDatasource) Get() *AgentStatusDatasource {
	return v.value
}

func (v *NullableAgentStatusDatasource) Set(val *AgentStatusDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentStatusDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentStatusDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentStatusDatasource(val *AgentStatusDatasource) *NullableAgentStatusDatasource {
	return &NullableAgentStatusDatasource{value: val, isSet: true}
}

func (v NullableAgentStatusDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentStatusDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

