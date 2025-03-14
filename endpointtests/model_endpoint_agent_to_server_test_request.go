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

// checks if the EndpointAgentToServerTestRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointAgentToServerTestRequest{}

// EndpointAgentToServerTestRequest struct for EndpointAgentToServerTestRequest
type EndpointAgentToServerTestRequest struct {
	AgentSelectorType *EndpointTestAgentSelectorType `json:"agentSelectorType,omitempty"`
	// List of endpoint agent IDs (obtained from `/endpoint/agents` endpoint). Required when `agentSelectorType` is set to `specific-agent`.
	Agents []string `json:"agents,omitempty"`
	// List of endpoint agent label IDs (obtained from `/endpoint/labels` endpoint), required when `agentSelectorType` is set to `agent-labels`.
	EndpointAgentLabels []string `json:"endpointAgentLabels,omitempty"`
	// Maximum number of agents which can execute the test.
	MaxMachines *int32 `json:"maxMachines,omitempty"`
	// Name of the test.
	TestName string `json:"testName"`
	// A server address without a protocol or IP address.
	ServerName string `json:"serverName"`
	// Port number.
	Port *int32 `json:"port,omitempty"`
	// Indicates whether the test should be prioritized when the number of tests assigned to an agent exceeds the license limit.
	IsPrioritized *bool `json:"isPrioritized,omitempty"`
	Interval *TestInterval `json:"interval,omitempty"`
	Protocol *EndpointTestProtocol `json:"protocol,omitempty"`
}

type _EndpointAgentToServerTestRequest EndpointAgentToServerTestRequest

// NewEndpointAgentToServerTestRequest instantiates a new EndpointAgentToServerTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAgentToServerTestRequest(testName string, serverName string) *EndpointAgentToServerTestRequest {
	this := EndpointAgentToServerTestRequest{}
	var agentSelectorType EndpointTestAgentSelectorType = ENDPOINTTESTAGENTSELECTORTYPE_ALL_AGENTS
	this.AgentSelectorType = &agentSelectorType
	var maxMachines int32 = 25
	this.MaxMachines = &maxMachines
	this.TestName = testName
	this.ServerName = serverName
	var port int32 = 443
	this.Port = &port
	var isPrioritized bool = false
	this.IsPrioritized = &isPrioritized
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = &interval
	var protocol EndpointTestProtocol = ENDPOINTTESTPROTOCOL_ICMP
	this.Protocol = &protocol
	return &this
}

// NewEndpointAgentToServerTestRequestWithDefaults instantiates a new EndpointAgentToServerTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAgentToServerTestRequestWithDefaults() *EndpointAgentToServerTestRequest {
	this := EndpointAgentToServerTestRequest{}
	var agentSelectorType EndpointTestAgentSelectorType = ENDPOINTTESTAGENTSELECTORTYPE_ALL_AGENTS
	this.AgentSelectorType = &agentSelectorType
	var maxMachines int32 = 25
	this.MaxMachines = &maxMachines
	var port int32 = 443
	this.Port = &port
	var isPrioritized bool = false
	this.IsPrioritized = &isPrioritized
	var interval TestInterval = TESTINTERVAL__60
	this.Interval = &interval
	var protocol EndpointTestProtocol = ENDPOINTTESTPROTOCOL_ICMP
	this.Protocol = &protocol
	return &this
}

// GetAgentSelectorType returns the AgentSelectorType field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetAgentSelectorType() EndpointTestAgentSelectorType {
	if o == nil || utils.IsNil(o.AgentSelectorType) {
		var ret EndpointTestAgentSelectorType
		return ret
	}
	return *o.AgentSelectorType
}

// GetAgentSelectorTypeOk returns a tuple with the AgentSelectorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetAgentSelectorTypeOk() (*EndpointTestAgentSelectorType, bool) {
	if o == nil || utils.IsNil(o.AgentSelectorType) {
		return nil, false
	}
	return o.AgentSelectorType, true
}

// HasAgentSelectorType returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasAgentSelectorType() bool {
	if o != nil && !utils.IsNil(o.AgentSelectorType) {
		return true
	}

	return false
}

// SetAgentSelectorType gets a reference to the given EndpointTestAgentSelectorType and assigns it to the AgentSelectorType field.
func (o *EndpointAgentToServerTestRequest) SetAgentSelectorType(v EndpointTestAgentSelectorType) {
	o.AgentSelectorType = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetAgents() []string {
	if o == nil || utils.IsNil(o.Agents) {
		var ret []string
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetAgentsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasAgents() bool {
	if o != nil && !utils.IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []string and assigns it to the Agents field.
func (o *EndpointAgentToServerTestRequest) SetAgents(v []string) {
	o.Agents = v
}

// GetEndpointAgentLabels returns the EndpointAgentLabels field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetEndpointAgentLabels() []string {
	if o == nil || utils.IsNil(o.EndpointAgentLabels) {
		var ret []string
		return ret
	}
	return o.EndpointAgentLabels
}

// GetEndpointAgentLabelsOk returns a tuple with the EndpointAgentLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetEndpointAgentLabelsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EndpointAgentLabels) {
		return nil, false
	}
	return o.EndpointAgentLabels, true
}

// HasEndpointAgentLabels returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasEndpointAgentLabels() bool {
	if o != nil && !utils.IsNil(o.EndpointAgentLabels) {
		return true
	}

	return false
}

// SetEndpointAgentLabels gets a reference to the given []string and assigns it to the EndpointAgentLabels field.
func (o *EndpointAgentToServerTestRequest) SetEndpointAgentLabels(v []string) {
	o.EndpointAgentLabels = v
}

// GetMaxMachines returns the MaxMachines field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetMaxMachines() int32 {
	if o == nil || utils.IsNil(o.MaxMachines) {
		var ret int32
		return ret
	}
	return *o.MaxMachines
}

// GetMaxMachinesOk returns a tuple with the MaxMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetMaxMachinesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxMachines) {
		return nil, false
	}
	return o.MaxMachines, true
}

// HasMaxMachines returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasMaxMachines() bool {
	if o != nil && !utils.IsNil(o.MaxMachines) {
		return true
	}

	return false
}

// SetMaxMachines gets a reference to the given int32 and assigns it to the MaxMachines field.
func (o *EndpointAgentToServerTestRequest) SetMaxMachines(v int32) {
	o.MaxMachines = &v
}

// GetTestName returns the TestName field value
func (o *EndpointAgentToServerTestRequest) GetTestName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetTestNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestName, true
}

// SetTestName sets field value
func (o *EndpointAgentToServerTestRequest) SetTestName(v string) {
	o.TestName = v
}

// GetServerName returns the ServerName field value
func (o *EndpointAgentToServerTestRequest) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *EndpointAgentToServerTestRequest) SetServerName(v string) {
	o.ServerName = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *EndpointAgentToServerTestRequest) SetPort(v int32) {
	o.Port = &v
}

// GetIsPrioritized returns the IsPrioritized field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetIsPrioritized() bool {
	if o == nil || utils.IsNil(o.IsPrioritized) {
		var ret bool
		return ret
	}
	return *o.IsPrioritized
}

// GetIsPrioritizedOk returns a tuple with the IsPrioritized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetIsPrioritizedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsPrioritized) {
		return nil, false
	}
	return o.IsPrioritized, true
}

// HasIsPrioritized returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasIsPrioritized() bool {
	if o != nil && !utils.IsNil(o.IsPrioritized) {
		return true
	}

	return false
}

// SetIsPrioritized gets a reference to the given bool and assigns it to the IsPrioritized field.
func (o *EndpointAgentToServerTestRequest) SetIsPrioritized(v bool) {
	o.IsPrioritized = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetInterval() TestInterval {
	if o == nil || utils.IsNil(o.Interval) {
		var ret TestInterval
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetIntervalOk() (*TestInterval, bool) {
	if o == nil || utils.IsNil(o.Interval) {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasInterval() bool {
	if o != nil && !utils.IsNil(o.Interval) {
		return true
	}

	return false
}

// SetInterval gets a reference to the given TestInterval and assigns it to the Interval field.
func (o *EndpointAgentToServerTestRequest) SetInterval(v TestInterval) {
	o.Interval = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *EndpointAgentToServerTestRequest) GetProtocol() EndpointTestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret EndpointTestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgentToServerTestRequest) GetProtocolOk() (*EndpointTestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *EndpointAgentToServerTestRequest) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given EndpointTestProtocol and assigns it to the Protocol field.
func (o *EndpointAgentToServerTestRequest) SetProtocol(v EndpointTestProtocol) {
	o.Protocol = &v
}

func (o EndpointAgentToServerTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAgentToServerTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AgentSelectorType) {
		toSerialize["agentSelectorType"] = o.AgentSelectorType
	}
	if !utils.IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !utils.IsNil(o.EndpointAgentLabels) {
		toSerialize["endpointAgentLabels"] = o.EndpointAgentLabels
	}
	if !utils.IsNil(o.MaxMachines) {
		toSerialize["maxMachines"] = o.MaxMachines
	}
	toSerialize["testName"] = o.TestName
	toSerialize["serverName"] = o.ServerName
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.IsPrioritized) {
		toSerialize["isPrioritized"] = o.IsPrioritized
	}
	if !utils.IsNil(o.Interval) {
		toSerialize["interval"] = o.Interval
	}
	if !utils.IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	return toSerialize, nil
}

func (o *EndpointAgentToServerTestRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"testName",
		"serverName",
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

	varEndpointAgentToServerTestRequest := _EndpointAgentToServerTestRequest{}

    err = json.Unmarshal(data, &varEndpointAgentToServerTestRequest)

	if err != nil {
		return err
	}

	*o = EndpointAgentToServerTestRequest(varEndpointAgentToServerTestRequest)

	return err
}

type NullableEndpointAgentToServerTestRequest struct {
	value *EndpointAgentToServerTestRequest
	isSet bool
}

func (v NullableEndpointAgentToServerTestRequest) Get() *EndpointAgentToServerTestRequest {
	return v.value
}

func (v *NullableEndpointAgentToServerTestRequest) Set(val *EndpointAgentToServerTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAgentToServerTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAgentToServerTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAgentToServerTestRequest(val *EndpointAgentToServerTestRequest) *NullableEndpointAgentToServerTestRequest {
	return &NullableEndpointAgentToServerTestRequest{value: val, isSet: true}
}

func (v NullableEndpointAgentToServerTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAgentToServerTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


