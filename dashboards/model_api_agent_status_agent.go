/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ApiAgentStatusAgent type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiAgentStatusAgent{}

// ApiAgentStatusAgent Agent shown in agent status widget.
type ApiAgentStatusAgent struct {
	// Identifier of the agent.
	AgentId *string `json:"agentId,omitempty"`
	Status *EnterpriseAgentState `json:"status,omitempty"`
	IpInfo *ApiAgentStatusIpInfo `json:"ipInfo,omitempty"`
	// Name of the agent
	AgentName *string `json:"agentName,omitempty"`
	Location *ApiAgentLocation `json:"location,omitempty"`
}

// NewApiAgentStatusAgent instantiates a new ApiAgentStatusAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAgentStatusAgent() *ApiAgentStatusAgent {
	this := ApiAgentStatusAgent{}
	return &this
}

// NewApiAgentStatusAgentWithDefaults instantiates a new ApiAgentStatusAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAgentStatusAgentWithDefaults() *ApiAgentStatusAgent {
	this := ApiAgentStatusAgent{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *ApiAgentStatusAgent) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentStatusAgent) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *ApiAgentStatusAgent) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *ApiAgentStatusAgent) SetAgentId(v string) {
	o.AgentId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiAgentStatusAgent) GetStatus() EnterpriseAgentState {
	if o == nil || utils.IsNil(o.Status) {
		var ret EnterpriseAgentState
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentStatusAgent) GetStatusOk() (*EnterpriseAgentState, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiAgentStatusAgent) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EnterpriseAgentState and assigns it to the Status field.
func (o *ApiAgentStatusAgent) SetStatus(v EnterpriseAgentState) {
	o.Status = &v
}

// GetIpInfo returns the IpInfo field value if set, zero value otherwise.
func (o *ApiAgentStatusAgent) GetIpInfo() ApiAgentStatusIpInfo {
	if o == nil || utils.IsNil(o.IpInfo) {
		var ret ApiAgentStatusIpInfo
		return ret
	}
	return *o.IpInfo
}

// GetIpInfoOk returns a tuple with the IpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentStatusAgent) GetIpInfoOk() (*ApiAgentStatusIpInfo, bool) {
	if o == nil || utils.IsNil(o.IpInfo) {
		return nil, false
	}
	return o.IpInfo, true
}

// HasIpInfo returns a boolean if a field has been set.
func (o *ApiAgentStatusAgent) HasIpInfo() bool {
	if o != nil && !utils.IsNil(o.IpInfo) {
		return true
	}

	return false
}

// SetIpInfo gets a reference to the given ApiAgentStatusIpInfo and assigns it to the IpInfo field.
func (o *ApiAgentStatusAgent) SetIpInfo(v ApiAgentStatusIpInfo) {
	o.IpInfo = &v
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *ApiAgentStatusAgent) GetAgentName() string {
	if o == nil || utils.IsNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentStatusAgent) GetAgentNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *ApiAgentStatusAgent) HasAgentName() bool {
	if o != nil && !utils.IsNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *ApiAgentStatusAgent) SetAgentName(v string) {
	o.AgentName = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *ApiAgentStatusAgent) GetLocation() ApiAgentLocation {
	if o == nil || utils.IsNil(o.Location) {
		var ret ApiAgentLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAgentStatusAgent) GetLocationOk() (*ApiAgentLocation, bool) {
	if o == nil || utils.IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ApiAgentStatusAgent) HasLocation() bool {
	if o != nil && !utils.IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given ApiAgentLocation and assigns it to the Location field.
func (o *ApiAgentStatusAgent) SetLocation(v ApiAgentLocation) {
	o.Location = &v
}

func (o ApiAgentStatusAgent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAgentStatusAgent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.IpInfo) {
		toSerialize["ipInfo"] = o.IpInfo
	}
	if !utils.IsNil(o.AgentName) {
		toSerialize["agentName"] = o.AgentName
	}
	if !utils.IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	return toSerialize, nil
}

type NullableApiAgentStatusAgent struct {
	value *ApiAgentStatusAgent
	isSet bool
}

func (v NullableApiAgentStatusAgent) Get() *ApiAgentStatusAgent {
	return v.value
}

func (v *NullableApiAgentStatusAgent) Set(val *ApiAgentStatusAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAgentStatusAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAgentStatusAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAgentStatusAgent(val *ApiAgentStatusAgent) *NullableApiAgentStatusAgent {
	return &NullableApiAgentStatusAgent{value: val, isSet: true}
}

func (v NullableApiAgentStatusAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAgentStatusAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


