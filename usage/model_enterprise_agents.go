/*
Usage API

 These usage endpoints define the following operations:  * **Usage**: Retrieve usage data for the specified time period (default is one month).          * Users must have the `View organization usage` permission to access this endpoint.     * This operation offers visibility across all account groups within the organization.     * Users with `View organization usage` permission in multiple organizations should query the operation with the `aid` query string parameter (see optional parameters) for each organization.  * **Quotas**: Obtain organization and account usage quotas. Additionally, users with the appropriate permissions can create, update, or delete these quotas.          * Users must have the necessary permissions to perform quota-related actions.  Refer to the Usage API operations for detailed usage instructions and optional parameters. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package usage

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the EnterpriseAgents type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EnterpriseAgents{}

// EnterpriseAgents struct for EnterpriseAgents
type EnterpriseAgents struct {
	// A unique identifier that specifies the account group that owns the enterprise agents.
	Aid *string `json:"aid,omitempty"`
	// Name of the account group which owns the enterprise agents.
	AccountGroupName *string `json:"accountGroupName,omitempty"`
	// Number of enterprise agents owned by the specific account group in the usage period.
	EnterpriseAgentsUsed *int64 `json:"enterpriseAgentsUsed,omitempty"`
}

// NewEnterpriseAgents instantiates a new EnterpriseAgents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAgents() *EnterpriseAgents {
	this := EnterpriseAgents{}
	return &this
}

// NewEnterpriseAgentsWithDefaults instantiates a new EnterpriseAgents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAgentsWithDefaults() *EnterpriseAgents {
	this := EnterpriseAgents{}
	return &this
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *EnterpriseAgents) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgents) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *EnterpriseAgents) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *EnterpriseAgents) SetAid(v string) {
	o.Aid = &v
}

// GetAccountGroupName returns the AccountGroupName field value if set, zero value otherwise.
func (o *EnterpriseAgents) GetAccountGroupName() string {
	if o == nil || utils.IsNil(o.AccountGroupName) {
		var ret string
		return ret
	}
	return *o.AccountGroupName
}

// GetAccountGroupNameOk returns a tuple with the AccountGroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgents) GetAccountGroupNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AccountGroupName) {
		return nil, false
	}
	return o.AccountGroupName, true
}

// HasAccountGroupName returns a boolean if a field has been set.
func (o *EnterpriseAgents) HasAccountGroupName() bool {
	if o != nil && !utils.IsNil(o.AccountGroupName) {
		return true
	}

	return false
}

// SetAccountGroupName gets a reference to the given string and assigns it to the AccountGroupName field.
func (o *EnterpriseAgents) SetAccountGroupName(v string) {
	o.AccountGroupName = &v
}

// GetEnterpriseAgentsUsed returns the EnterpriseAgentsUsed field value if set, zero value otherwise.
func (o *EnterpriseAgents) GetEnterpriseAgentsUsed() int64 {
	if o == nil || utils.IsNil(o.EnterpriseAgentsUsed) {
		var ret int64
		return ret
	}
	return *o.EnterpriseAgentsUsed
}

// GetEnterpriseAgentsUsedOk returns a tuple with the EnterpriseAgentsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgents) GetEnterpriseAgentsUsedOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.EnterpriseAgentsUsed) {
		return nil, false
	}
	return o.EnterpriseAgentsUsed, true
}

// HasEnterpriseAgentsUsed returns a boolean if a field has been set.
func (o *EnterpriseAgents) HasEnterpriseAgentsUsed() bool {
	if o != nil && !utils.IsNil(o.EnterpriseAgentsUsed) {
		return true
	}

	return false
}

// SetEnterpriseAgentsUsed gets a reference to the given int64 and assigns it to the EnterpriseAgentsUsed field.
func (o *EnterpriseAgents) SetEnterpriseAgentsUsed(v int64) {
	o.EnterpriseAgentsUsed = &v
}

func (o EnterpriseAgents) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseAgents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.AccountGroupName) {
		toSerialize["accountGroupName"] = o.AccountGroupName
	}
	if !utils.IsNil(o.EnterpriseAgentsUsed) {
		toSerialize["enterpriseAgentsUsed"] = o.EnterpriseAgentsUsed
	}
	return toSerialize, nil
}

type NullableEnterpriseAgents struct {
	value *EnterpriseAgents
	isSet bool
}

func (v NullableEnterpriseAgents) Get() *EnterpriseAgents {
	return v.value
}

func (v *NullableEnterpriseAgents) Set(val *EnterpriseAgents) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAgents) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAgents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAgents(val *EnterpriseAgents) *NullableEnterpriseAgents {
	return &NullableEnterpriseAgents{value: val, isSet: true}
}

func (v NullableEnterpriseAgents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAgents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


