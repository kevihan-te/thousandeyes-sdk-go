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

// checks if the EnterpriseAgentsUsage type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EnterpriseAgentsUsage{}

// EnterpriseAgentsUsage struct for EnterpriseAgentsUsage
type EnterpriseAgentsUsage struct {
	Breakdowns []EnterpriseAgentUnitsByTestOwnerAccountGroup `json:"breakdowns,omitempty"`
	Links *PaginationLinks `json:"_links,omitempty"`
}

// NewEnterpriseAgentsUsage instantiates a new EnterpriseAgentsUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnterpriseAgentsUsage() *EnterpriseAgentsUsage {
	this := EnterpriseAgentsUsage{}
	return &this
}

// NewEnterpriseAgentsUsageWithDefaults instantiates a new EnterpriseAgentsUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnterpriseAgentsUsageWithDefaults() *EnterpriseAgentsUsage {
	this := EnterpriseAgentsUsage{}
	return &this
}

// GetBreakdowns returns the Breakdowns field value if set, zero value otherwise.
func (o *EnterpriseAgentsUsage) GetBreakdowns() []EnterpriseAgentUnitsByTestOwnerAccountGroup {
	if o == nil || utils.IsNil(o.Breakdowns) {
		var ret []EnterpriseAgentUnitsByTestOwnerAccountGroup
		return ret
	}
	return o.Breakdowns
}

// GetBreakdownsOk returns a tuple with the Breakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentsUsage) GetBreakdownsOk() ([]EnterpriseAgentUnitsByTestOwnerAccountGroup, bool) {
	if o == nil || utils.IsNil(o.Breakdowns) {
		return nil, false
	}
	return o.Breakdowns, true
}

// HasBreakdowns returns a boolean if a field has been set.
func (o *EnterpriseAgentsUsage) HasBreakdowns() bool {
	if o != nil && !utils.IsNil(o.Breakdowns) {
		return true
	}

	return false
}

// SetBreakdowns gets a reference to the given []EnterpriseAgentUnitsByTestOwnerAccountGroup and assigns it to the Breakdowns field.
func (o *EnterpriseAgentsUsage) SetBreakdowns(v []EnterpriseAgentUnitsByTestOwnerAccountGroup) {
	o.Breakdowns = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EnterpriseAgentsUsage) GetLinks() PaginationLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnterpriseAgentsUsage) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EnterpriseAgentsUsage) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *EnterpriseAgentsUsage) SetLinks(v PaginationLinks) {
	o.Links = &v
}

func (o EnterpriseAgentsUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnterpriseAgentsUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Breakdowns) {
		toSerialize["breakdowns"] = o.Breakdowns
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableEnterpriseAgentsUsage struct {
	value *EnterpriseAgentsUsage
	isSet bool
}

func (v NullableEnterpriseAgentsUsage) Get() *EnterpriseAgentsUsage {
	return v.value
}

func (v *NullableEnterpriseAgentsUsage) Set(val *EnterpriseAgentsUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableEnterpriseAgentsUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableEnterpriseAgentsUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnterpriseAgentsUsage(val *EnterpriseAgentsUsage) *NullableEnterpriseAgentsUsage {
	return &NullableEnterpriseAgentsUsage{value: val, isSet: true}
}

func (v NullableEnterpriseAgentsUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnterpriseAgentsUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


