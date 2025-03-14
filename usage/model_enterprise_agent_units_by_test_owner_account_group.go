/*
Usage API

 These usage endpoints define the following operations:  * **Usage**: Retrieve usage data for the specified time period (default is one month).          * Users must have the `View organization usage` permission to access this endpoint.     * This operation offers visibility across all account groups within the organization.     * Users with `View organization usage` permission in multiple organizations should query the operation with the `aid` query string parameter (see optional parameters) for each organization.  * **Quotas**: Obtain organization and account usage quotas. Additionally, users with the appropriate permissions can create, update, or delete these quotas.          * Users must have the necessary permissions to perform quota-related actions.  Refer to the Usage API operations for detailed usage instructions and optional parameters. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package usage

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the EnterpriseAgentUnitsByTestOwnerAccountGroup type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EnterpriseAgentUnitsByTestOwnerAccountGroup{}

// EnterpriseAgentUnitsByTestOwnerAccountGroup struct for EnterpriseAgentUnitsByTestOwnerAccountGroup
type EnterpriseAgentUnitsByTestOwnerAccountGroup struct {
	// Unique identifier of the account group where some tests are incurring the enterprise agent units.
	Aid *string `json:"aid,omitempty"`
	// Name of the account group which owns the tests that are incurring enterprise agent units.
	AccountGroupName *string `json:"accountGroupName,omitempty"`
	// Unique identifier of the enterprise agent generating usage.
	AgentId *string `json:"agentId,omitempty"`
	// Name of the enterprise agent generating usage.
	AgentName *string `json:"agentName,omitempty"`
	// Number of enterprise agent units owned by the specific account group in the usage period.
	EnterpriseUnitsUsed *int64 `json:"enterpriseUnitsUsed,omitempty"`
	// Number of enterprise units projected in the current usage period, based on units consumed to date and configuration of enabled tests. This value is updated hourly. Returns non-zero value only for organizations with metered billing.
	EnterpriseUnitsProjected *int64 `json:"enterpriseUnitsProjected,omitempty"`
	// Unique identifier of the virtual agent generating usage
	VagentId *string `json:"vagentId,omitempty"`
}

// NewEnterpriseAgentUnitsByTestOwnerAccountGroup instantiates a new EnterpriseAgentUnitsByTestOwnerAccountGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAgentUnitsByTestOwnerAccountGroup() *EnterpriseAgentUnitsByTestOwnerAccountGroup {
	this := EnterpriseAgentUnitsByTestOwnerAccountGroup{}
	return &this
}

// NewEnterpriseAgentUnitsByTestOwnerAccountGroupWithDefaults instantiates a new EnterpriseAgentUnitsByTestOwnerAccountGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAgentUnitsByTestOwnerAccountGroupWithDefaults() *EnterpriseAgentUnitsByTestOwnerAccountGroup {
	this := EnterpriseAgentUnitsByTestOwnerAccountGroup{}
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) SetAid(v string) {
	o.Aid = &v
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAccountGroupName() string {
	if o == nil || utils.IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) HasAccountGroupName() bool {
	if o != nil && !utils.IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) SetAgentId(v string) {
	o.AgentId = &v
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAgentName() string {
	if o == nil || utils.IsNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetAgentNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) HasAgentName() bool {
	if o != nil && !utils.IsNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) SetAgentName(v string) {
	o.AgentName = &v
}

// GetEnterpriseUnitsUsed returns the EnterpriseUnitsUsed field value if set, zero value otherwise.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetEnterpriseUnitsUsed() int64 {
	if o == nil || utils.IsNil(o.EnterpriseUnitsUsed) {
		var ret int64
		return ret
	}
	return *o.EnterpriseUnitsUsed
}

// GetEnterpriseUnitsUsedOk returns a tuple with the EnterpriseUnitsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetEnterpriseUnitsUsedOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.EnterpriseUnitsUsed) {
		return nil, false
	}
	return o.EnterpriseUnitsUsed, true
}

// HasEnterpriseUnitsUsed returns a boolean if a field has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) HasEnterpriseUnitsUsed() bool {
	if o != nil && !utils.IsNil(o.EnterpriseUnitsUsed) {
		return true
	}

	return false
}

// SetEnterpriseUnitsUsed gets a reference to the given int64 and assigns it to the EnterpriseUnitsUsed field.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) SetEnterpriseUnitsUsed(v int64) {
	o.EnterpriseUnitsUsed = &v
}

// GetEnterpriseUnitsProjected returns the EnterpriseUnitsProjected field value if set, zero value otherwise.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetEnterpriseUnitsProjected() int64 {
	if o == nil || utils.IsNil(o.EnterpriseUnitsProjected) {
		var ret int64
		return ret
	}
	return *o.EnterpriseUnitsProjected
}

// GetEnterpriseUnitsProjectedOk returns a tuple with the EnterpriseUnitsProjected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetEnterpriseUnitsProjectedOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.EnterpriseUnitsProjected) {
		return nil, false
	}
	return o.EnterpriseUnitsProjected, true
}

// HasEnterpriseUnitsProjected returns a boolean if a field has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) HasEnterpriseUnitsProjected() bool {
	if o != nil && !utils.IsNil(o.EnterpriseUnitsProjected) {
		return true
	}

	return false
}

// SetEnterpriseUnitsProjected gets a reference to the given int64 and assigns it to the EnterpriseUnitsProjected field.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) SetEnterpriseUnitsProjected(v int64) {
	o.EnterpriseUnitsProjected = &v
}

// GetVagentId returns the VagentId field value if set, zero value otherwise.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetVagentId() string {
	if o == nil || utils.IsNil(o.VagentId) {
		var ret string
		return ret
	}
	return *o.VagentId
}

// GetVagentIdOk returns a tuple with the VagentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) GetVagentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VagentId) {
		return nil, false
	}
	return o.VagentId, true
}

// HasVagentId returns a boolean if a field has been set.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) HasVagentId() bool {
	if o != nil && !utils.IsNil(o.VagentId) {
		return true
	}

	return false
}

// SetVagentId gets a reference to the given string and assigns it to the VagentId field.
func (o *EnterpriseAgentUnitsByTestOwnerAccountGroup) SetVagentId(v string) {
	o.VagentId = &v
}

func (o EnterpriseAgentUnitsByTestOwnerAccountGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseAgentUnitsByTestOwnerAccountGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.AccountGroupName) {
		toSerialize["accountGroupName"] = o.AccountGroupName
	}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.AgentName) {
		toSerialize["agentName"] = o.AgentName
	}
	if !utils.IsNil(o.EnterpriseUnitsUsed) {
		toSerialize["enterpriseUnitsUsed"] = o.EnterpriseUnitsUsed
	}
	if !utils.IsNil(o.EnterpriseUnitsProjected) {
		toSerialize["enterpriseUnitsProjected"] = o.EnterpriseUnitsProjected
	}
	if !utils.IsNil(o.VagentId) {
		toSerialize["vagentId"] = o.VagentId
	}
	return toSerialize, nil
}

type NullableEnterpriseAgentUnitsByTestOwnerAccountGroup struct {
	value *EnterpriseAgentUnitsByTestOwnerAccountGroup
	isSet bool
}

func (v NullableEnterpriseAgentUnitsByTestOwnerAccountGroup) Get() *EnterpriseAgentUnitsByTestOwnerAccountGroup {
	return v.value
}

func (v *NullableEnterpriseAgentUnitsByTestOwnerAccountGroup) Set(val *EnterpriseAgentUnitsByTestOwnerAccountGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAgentUnitsByTestOwnerAccountGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAgentUnitsByTestOwnerAccountGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAgentUnitsByTestOwnerAccountGroup(val *EnterpriseAgentUnitsByTestOwnerAccountGroup) *NullableEnterpriseAgentUnitsByTestOwnerAccountGroup {
	return &NullableEnterpriseAgentUnitsByTestOwnerAccountGroup{value: val, isSet: true}
}

func (v NullableEnterpriseAgentUnitsByTestOwnerAccountGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAgentUnitsByTestOwnerAccountGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


