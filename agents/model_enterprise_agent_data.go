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
	"time"
)

// checks if the EnterpriseAgentData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EnterpriseAgentData{}

// EnterpriseAgentData struct for EnterpriseAgentData
type EnterpriseAgentData struct {
	// If an enterprise agent is clustered, detailed information about each cluster member will be shown as array entries in the clusterMembers field. This field is not shown for Enterprise Agents in standalone mode, or for Cloud Agents.
	ClusterMembers []ClusterMember `json:"clusterMembers,omitempty"`
	// Shows overall utilization percentage (online Enterprise Agents and Enterprise Clusters only).
	Utilization *int32 `json:"utilization,omitempty"`
	// List of account groups. See /accounts-groups to pull a list of account IDs
	AccountGroups []AccountGroup `json:"accountGroups,omitempty"`
	Ipv6Policy *EnterpriseAgentIpv6Policy `json:"ipv6Policy,omitempty"`
	// If an enterprise agent or a cluster member presents at least one error, the errors will be shown as an array of entries in the errorDetails field (Enterprise Agents and Enterprise Cluster members only)
	ErrorDetails []ErrorDetail `json:"errorDetails,omitempty"`
	// Fully qualified domain name of the agent (Enterprise Agents only)
	Hostname *string `json:"hostname,omitempty"`
	// UTC last seen date (ISO date-time format).
	LastSeen *time.Time `json:"lastSeen,omitempty"`
	AgentState *EnterpriseAgentState `json:"agentState,omitempty"`
	// Flag indicating if the agent retains cache.
	KeepBrowserCache *bool `json:"keepBrowserCache,omitempty"`
	// UTC Agent creation date (ISO date-time format).
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// Test target IP address.
	TargetForTests *string `json:"targetForTests,omitempty"`
	// To perform rDNS lookups for public IP ranges, this field represents the public IP ranges. The range must be in CIDR notation; for example, 10.1.1.0/24. Maximum of 5 prefixes allowed (Enterprise Agents and Enterprise Agent clusters only).
	LocalResolutionPrefixes []string `json:"localResolutionPrefixes,omitempty"`
	InterfaceIpMappings []InterfaceIpMapping `json:"interfaceIpMappings,omitempty"`
}

// NewEnterpriseAgentData instantiates a new EnterpriseAgentData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAgentData() *EnterpriseAgentData {
	this := EnterpriseAgentData{}
	return &this
}

// NewEnterpriseAgentDataWithDefaults instantiates a new EnterpriseAgentData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAgentDataWithDefaults() *EnterpriseAgentData {
	this := EnterpriseAgentData{}
	return &this
}

// GetClusterMembers returns the ClusterMembers field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetClusterMembers() []ClusterMember {
	if o == nil || utils.IsNil(o.ClusterMembers) {
		var ret []ClusterMember
		return ret
	}
	return o.ClusterMembers
}

// GetClusterMembersOk returns a tuple with the ClusterMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetClusterMembersOk() ([]ClusterMember, bool) {
	if o == nil || utils.IsNil(o.ClusterMembers) {
		return nil, false
	}
	return o.ClusterMembers, true
}

// HasClusterMembers returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasClusterMembers() bool {
	if o != nil && !utils.IsNil(o.ClusterMembers) {
		return true
	}

	return false
}

// SetClusterMembers gets a reference to the given []ClusterMember and assigns it to the ClusterMembers field.
func (o *EnterpriseAgentData) SetClusterMembers(v []ClusterMember) {
	o.ClusterMembers = v
}

// GetUtilization returns the Utilization field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetUtilization() int32 {
	if o == nil || utils.IsNil(o.Utilization) {
		var ret int32
		return ret
	}
	return *o.Utilization
}

// GetUtilizationOk returns a tuple with the Utilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetUtilizationOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Utilization) {
		return nil, false
	}
	return o.Utilization, true
}

// HasUtilization returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasUtilization() bool {
	if o != nil && !utils.IsNil(o.Utilization) {
		return true
	}

	return false
}

// SetUtilization gets a reference to the given int32 and assigns it to the Utilization field.
func (o *EnterpriseAgentData) SetUtilization(v int32) {
	o.Utilization = &v
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetAccountGroups() []AccountGroup {
	if o == nil || utils.IsNil(o.AccountGroups) {
		var ret []AccountGroup
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetAccountGroupsOk() ([]AccountGroup, bool) {
	if o == nil || utils.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasAccountGroups() bool {
	if o != nil && !utils.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []AccountGroup and assigns it to the AccountGroups field.
func (o *EnterpriseAgentData) SetAccountGroups(v []AccountGroup) {
	o.AccountGroups = v
}

// GetIpv6Policy returns the Ipv6Policy field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetIpv6Policy() EnterpriseAgentIpv6Policy {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		var ret EnterpriseAgentIpv6Policy
		return ret
	}
	return *o.Ipv6Policy
}

// GetIpv6PolicyOk returns a tuple with the Ipv6Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetIpv6PolicyOk() (*EnterpriseAgentIpv6Policy, bool) {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		return nil, false
	}
	return o.Ipv6Policy, true
}

// HasIpv6Policy returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasIpv6Policy() bool {
	if o != nil && !utils.IsNil(o.Ipv6Policy) {
		return true
	}

	return false
}

// SetIpv6Policy gets a reference to the given EnterpriseAgentIpv6Policy and assigns it to the Ipv6Policy field.
func (o *EnterpriseAgentData) SetIpv6Policy(v EnterpriseAgentIpv6Policy) {
	o.Ipv6Policy = &v
}

// GetErrorDetails returns the ErrorDetails field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetErrorDetails() []ErrorDetail {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		var ret []ErrorDetail
		return ret
	}
	return o.ErrorDetails
}

// GetErrorDetailsOk returns a tuple with the ErrorDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetErrorDetailsOk() ([]ErrorDetail, bool) {
	if o == nil || utils.IsNil(o.ErrorDetails) {
		return nil, false
	}
	return o.ErrorDetails, true
}

// HasErrorDetails returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasErrorDetails() bool {
	if o != nil && !utils.IsNil(o.ErrorDetails) {
		return true
	}

	return false
}

// SetErrorDetails gets a reference to the given []ErrorDetail and assigns it to the ErrorDetails field.
func (o *EnterpriseAgentData) SetErrorDetails(v []ErrorDetail) {
	o.ErrorDetails = v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetHostname() string {
	if o == nil || utils.IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetHostnameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasHostname() bool {
	if o != nil && !utils.IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *EnterpriseAgentData) SetHostname(v string) {
	o.Hostname = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetLastSeen() time.Time {
	if o == nil || utils.IsNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasLastSeen() bool {
	if o != nil && !utils.IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *EnterpriseAgentData) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetAgentState returns the AgentState field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetAgentState() EnterpriseAgentState {
	if o == nil || utils.IsNil(o.AgentState) {
		var ret EnterpriseAgentState
		return ret
	}
	return *o.AgentState
}

// GetAgentStateOk returns a tuple with the AgentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetAgentStateOk() (*EnterpriseAgentState, bool) {
	if o == nil || utils.IsNil(o.AgentState) {
		return nil, false
	}
	return o.AgentState, true
}

// HasAgentState returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasAgentState() bool {
	if o != nil && !utils.IsNil(o.AgentState) {
		return true
	}

	return false
}

// SetAgentState gets a reference to the given EnterpriseAgentState and assigns it to the AgentState field.
func (o *EnterpriseAgentData) SetAgentState(v EnterpriseAgentState) {
	o.AgentState = &v
}

// GetKeepBrowserCache returns the KeepBrowserCache field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetKeepBrowserCache() bool {
	if o == nil || utils.IsNil(o.KeepBrowserCache) {
		var ret bool
		return ret
	}
	return *o.KeepBrowserCache
}

// GetKeepBrowserCacheOk returns a tuple with the KeepBrowserCache field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetKeepBrowserCacheOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.KeepBrowserCache) {
		return nil, false
	}
	return o.KeepBrowserCache, true
}

// HasKeepBrowserCache returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasKeepBrowserCache() bool {
	if o != nil && !utils.IsNil(o.KeepBrowserCache) {
		return true
	}

	return false
}

// SetKeepBrowserCache gets a reference to the given bool and assigns it to the KeepBrowserCache field.
func (o *EnterpriseAgentData) SetKeepBrowserCache(v bool) {
	o.KeepBrowserCache = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *EnterpriseAgentData) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetTargetForTests returns the TargetForTests field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetTargetForTests() string {
	if o == nil || utils.IsNil(o.TargetForTests) {
		var ret string
		return ret
	}
	return *o.TargetForTests
}

// GetTargetForTestsOk returns a tuple with the TargetForTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetTargetForTestsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TargetForTests) {
		return nil, false
	}
	return o.TargetForTests, true
}

// HasTargetForTests returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasTargetForTests() bool {
	if o != nil && !utils.IsNil(o.TargetForTests) {
		return true
	}

	return false
}

// SetTargetForTests gets a reference to the given string and assigns it to the TargetForTests field.
func (o *EnterpriseAgentData) SetTargetForTests(v string) {
	o.TargetForTests = &v
}

// GetLocalResolutionPrefixes returns the LocalResolutionPrefixes field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetLocalResolutionPrefixes() []string {
	if o == nil || utils.IsNil(o.LocalResolutionPrefixes) {
		var ret []string
		return ret
	}
	return o.LocalResolutionPrefixes
}

// GetLocalResolutionPrefixesOk returns a tuple with the LocalResolutionPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetLocalResolutionPrefixesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.LocalResolutionPrefixes) {
		return nil, false
	}
	return o.LocalResolutionPrefixes, true
}

// HasLocalResolutionPrefixes returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasLocalResolutionPrefixes() bool {
	if o != nil && !utils.IsNil(o.LocalResolutionPrefixes) {
		return true
	}

	return false
}

// SetLocalResolutionPrefixes gets a reference to the given []string and assigns it to the LocalResolutionPrefixes field.
func (o *EnterpriseAgentData) SetLocalResolutionPrefixes(v []string) {
	o.LocalResolutionPrefixes = v
}

// GetInterfaceIpMappings returns the InterfaceIpMappings field value if set, zero value otherwise.
func (o *EnterpriseAgentData) GetInterfaceIpMappings() []InterfaceIpMapping {
	if o == nil || utils.IsNil(o.InterfaceIpMappings) {
		var ret []InterfaceIpMapping
		return ret
	}
	return o.InterfaceIpMappings
}

// GetInterfaceIpMappingsOk returns a tuple with the InterfaceIpMappings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentData) GetInterfaceIpMappingsOk() ([]InterfaceIpMapping, bool) {
	if o == nil || utils.IsNil(o.InterfaceIpMappings) {
		return nil, false
	}
	return o.InterfaceIpMappings, true
}

// HasInterfaceIpMappings returns a boolean if a field has been set.
func (o *EnterpriseAgentData) HasInterfaceIpMappings() bool {
	if o != nil && !utils.IsNil(o.InterfaceIpMappings) {
		return true
	}

	return false
}

// SetInterfaceIpMappings gets a reference to the given []InterfaceIpMapping and assigns it to the InterfaceIpMappings field.
func (o *EnterpriseAgentData) SetInterfaceIpMappings(v []InterfaceIpMapping) {
	o.InterfaceIpMappings = v
}

func (o EnterpriseAgentData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseAgentData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ClusterMembers) {
		toSerialize["clusterMembers"] = o.ClusterMembers
	}
	if !utils.IsNil(o.Utilization) {
		toSerialize["utilization"] = o.Utilization
	}
	if !utils.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	if !utils.IsNil(o.Ipv6Policy) {
		toSerialize["ipv6Policy"] = o.Ipv6Policy
	}
	if !utils.IsNil(o.ErrorDetails) {
		toSerialize["errorDetails"] = o.ErrorDetails
	}
	if !utils.IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !utils.IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !utils.IsNil(o.AgentState) {
		toSerialize["agentState"] = o.AgentState
	}
	if !utils.IsNil(o.KeepBrowserCache) {
		toSerialize["keepBrowserCache"] = o.KeepBrowserCache
	}
	if !utils.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !utils.IsNil(o.TargetForTests) {
		toSerialize["targetForTests"] = o.TargetForTests
	}
	if !utils.IsNil(o.LocalResolutionPrefixes) {
		toSerialize["localResolutionPrefixes"] = o.LocalResolutionPrefixes
	}
	if !utils.IsNil(o.InterfaceIpMappings) {
		toSerialize["interfaceIpMappings"] = o.InterfaceIpMappings
	}
	return toSerialize, nil
}

type NullableEnterpriseAgentData struct {
	value *EnterpriseAgentData
	isSet bool
}

func (v NullableEnterpriseAgentData) Get() *EnterpriseAgentData {
	return v.value
}

func (v *NullableEnterpriseAgentData) Set(val *EnterpriseAgentData) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAgentData) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAgentData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAgentData(val *EnterpriseAgentData) *NullableEnterpriseAgentData {
	return &NullableEnterpriseAgentData{value: val, isSet: true}
}

func (v NullableEnterpriseAgentData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAgentData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


