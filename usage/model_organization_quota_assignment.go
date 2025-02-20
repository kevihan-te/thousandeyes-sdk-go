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

// checks if the OrganizationQuotaAssignment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &OrganizationQuotaAssignment{}

// OrganizationQuotaAssignment struct for OrganizationQuotaAssignment
type OrganizationQuotaAssignment struct {
	// Unique identifier of the organization.
	OrgId *string `json:"orgId,omitempty"`
	// List of account groups quotas.
	AccountGroups []AccountGroupQuota `json:"accountGroups,omitempty"`
}

// NewOrganizationQuotaAssignment instantiates a new OrganizationQuotaAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationQuotaAssignment() *OrganizationQuotaAssignment {
	this := OrganizationQuotaAssignment{}
	return &this
}

// NewOrganizationQuotaAssignmentWithDefaults instantiates a new OrganizationQuotaAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationQuotaAssignmentWithDefaults() *OrganizationQuotaAssignment {
	this := OrganizationQuotaAssignment{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrganizationQuotaAssignment) GetOrgId() string {
	if o == nil || utils.IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationQuotaAssignment) GetOrgIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrganizationQuotaAssignment) HasOrgId() bool {
	if o != nil && !utils.IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *OrganizationQuotaAssignment) SetOrgId(v string) {
	o.OrgId = &v
}

// GetAccountGroups returns the AccountGroups field value if set, zero value otherwise.
func (o *OrganizationQuotaAssignment) GetAccountGroups() []AccountGroupQuota {
	if o == nil || utils.IsNil(o.AccountGroups) {
		var ret []AccountGroupQuota
		return ret
	}
	return o.AccountGroups
}

// GetAccountGroupsOk returns a tuple with the AccountGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationQuotaAssignment) GetAccountGroupsOk() ([]AccountGroupQuota, bool) {
	if o == nil || utils.IsNil(o.AccountGroups) {
		return nil, false
	}
	return o.AccountGroups, true
}

// HasAccountGroups returns a boolean if a field has been set.
func (o *OrganizationQuotaAssignment) HasAccountGroups() bool {
	if o != nil && !utils.IsNil(o.AccountGroups) {
		return true
	}

	return false
}

// SetAccountGroups gets a reference to the given []AccountGroupQuota and assigns it to the AccountGroups field.
func (o *OrganizationQuotaAssignment) SetAccountGroups(v []AccountGroupQuota) {
	o.AccountGroups = v
}

func (o OrganizationQuotaAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationQuotaAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.OrgId) {
		toSerialize["orgId"] = o.OrgId
	}
	if !utils.IsNil(o.AccountGroups) {
		toSerialize["accountGroups"] = o.AccountGroups
	}
	return toSerialize, nil
}

type NullableOrganizationQuotaAssignment struct {
	value *OrganizationQuotaAssignment
	isSet bool
}

func (v NullableOrganizationQuotaAssignment) Get() *OrganizationQuotaAssignment {
	return v.value
}

func (v *NullableOrganizationQuotaAssignment) Set(val *OrganizationQuotaAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationQuotaAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationQuotaAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationQuotaAssignment(val *OrganizationQuotaAssignment) *NullableOrganizationQuotaAssignment {
	return &NullableOrganizationQuotaAssignment{value: val, isSet: true}
}

func (v NullableOrganizationQuotaAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationQuotaAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


