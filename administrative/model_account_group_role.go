/*
Administrative API

Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API. This API provides the following operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These operations can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.    For more information about the administrative models, see [Account Management](https://docs.thousandeyes.com/product-documentation/user-management).

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package administrative

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AccountGroupRole type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AccountGroupRole{}

// AccountGroupRole struct for AccountGroupRole
type AccountGroupRole struct {
	AccountGroup *AccountGroup `json:"accountGroup,omitempty"`
	Roles []Role `json:"roles,omitempty"`
}

// NewAccountGroupRole instantiates a new AccountGroupRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountGroupRole() *AccountGroupRole {
	this := AccountGroupRole{}
	return &this
}

// NewAccountGroupRoleWithDefaults instantiates a new AccountGroupRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountGroupRoleWithDefaults() *AccountGroupRole {
	this := AccountGroupRole{}
	return &this
}

// GetAccountGroup returns the AccountGroup field value if set, zero value otherwise.
func (o *AccountGroupRole) GetAccountGroup() AccountGroup {
	if o == nil || utils.IsNil(o.AccountGroup) {
		var ret AccountGroup
		return ret
	}
	return *o.AccountGroup
}

// GetAccountGroupOk returns a tuple with the AccountGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroupRole) GetAccountGroupOk() (*AccountGroup, bool) {
	if o == nil || utils.IsNil(o.AccountGroup) {
		return nil, false
	}
	return o.AccountGroup, true
}

// HasAccountGroup returns a boolean if a field has been set.
func (o *AccountGroupRole) HasAccountGroup() bool {
	if o != nil && !utils.IsNil(o.AccountGroup) {
		return true
	}

	return false
}

// SetAccountGroup gets a reference to the given AccountGroup and assigns it to the AccountGroup field.
func (o *AccountGroupRole) SetAccountGroup(v AccountGroup) {
	o.AccountGroup = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AccountGroupRole) GetRoles() []Role {
	if o == nil || utils.IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountGroupRole) GetRolesOk() ([]Role, bool) {
	if o == nil || utils.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AccountGroupRole) HasRoles() bool {
	if o != nil && !utils.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *AccountGroupRole) SetRoles(v []Role) {
	o.Roles = v
}

func (o AccountGroupRole) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountGroupRole) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccountGroup) {
		toSerialize["accountGroup"] = o.AccountGroup
	}
	if !utils.IsNil(o.Roles) {
		toSerialize["roles"] = o.Roles
	}
	return toSerialize, nil
}

type NullableAccountGroupRole struct {
	value *AccountGroupRole
	isSet bool
}

func (v NullableAccountGroupRole) Get() *AccountGroupRole {
	return v.value
}

func (v *NullableAccountGroupRole) Set(val *AccountGroupRole) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountGroupRole) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountGroupRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountGroupRole(val *AccountGroupRole) *NullableAccountGroupRole {
	return &NullableAccountGroupRole{value: val, isSet: true}
}

func (v NullableAccountGroupRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountGroupRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


