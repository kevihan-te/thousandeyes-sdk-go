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

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Role{}

// Role struct for Role
type Role struct {
	// Name of the role.
	Name *string `json:"name,omitempty"`
	// Unique ID representing the role.
	RoleId *string `json:"roleId,omitempty"`
	// Flag indicating if the role is built-in (Account Admin, Organization Admin, Regular User).
	IsBuiltin *bool `json:"isBuiltin,omitempty"`
	// Flag indicating whether the user has management permissions.
	HasManagementPermissions *bool `json:"hasManagementPermissions,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Role) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Role) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Role) SetName(v string) {
	o.Name = &v
}

// GetRoleId returns the RoleId field value if set, zero value otherwise.
func (o *Role) GetRoleId() string {
	if o == nil || utils.IsNil(o.RoleId) {
		var ret string
		return ret
	}
	return *o.RoleId
}

// GetRoleIdOk returns a tuple with the RoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetRoleIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RoleId) {
		return nil, false
	}
	return o.RoleId, true
}

// HasRoleId returns a boolean if a field has been set.
func (o *Role) HasRoleId() bool {
	if o != nil && !utils.IsNil(o.RoleId) {
		return true
	}

	return false
}

// SetRoleId gets a reference to the given string and assigns it to the RoleId field.
func (o *Role) SetRoleId(v string) {
	o.RoleId = &v
}

// GetIsBuiltin returns the IsBuiltin field value if set, zero value otherwise.
func (o *Role) GetIsBuiltin() bool {
	if o == nil || utils.IsNil(o.IsBuiltin) {
		var ret bool
		return ret
	}
	return *o.IsBuiltin
}

// GetIsBuiltinOk returns a tuple with the IsBuiltin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIsBuiltinOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsBuiltin) {
		return nil, false
	}
	return o.IsBuiltin, true
}

// HasIsBuiltin returns a boolean if a field has been set.
func (o *Role) HasIsBuiltin() bool {
	if o != nil && !utils.IsNil(o.IsBuiltin) {
		return true
	}

	return false
}

// SetIsBuiltin gets a reference to the given bool and assigns it to the IsBuiltin field.
func (o *Role) SetIsBuiltin(v bool) {
	o.IsBuiltin = &v
}

// GetHasManagementPermissions returns the HasManagementPermissions field value if set, zero value otherwise.
func (o *Role) GetHasManagementPermissions() bool {
	if o == nil || utils.IsNil(o.HasManagementPermissions) {
		var ret bool
		return ret
	}
	return *o.HasManagementPermissions
}

// GetHasManagementPermissionsOk returns a tuple with the HasManagementPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetHasManagementPermissionsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.HasManagementPermissions) {
		return nil, false
	}
	return o.HasManagementPermissions, true
}

// HasHasManagementPermissions returns a boolean if a field has been set.
func (o *Role) HasHasManagementPermissions() bool {
	if o != nil && !utils.IsNil(o.HasManagementPermissions) {
		return true
	}

	return false
}

// SetHasManagementPermissions gets a reference to the given bool and assigns it to the HasManagementPermissions field.
func (o *Role) SetHasManagementPermissions(v bool) {
	o.HasManagementPermissions = &v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.RoleId) {
		toSerialize["roleId"] = o.RoleId
	}
	if !utils.IsNil(o.IsBuiltin) {
		toSerialize["isBuiltin"] = o.IsBuiltin
	}
	if !utils.IsNil(o.HasManagementPermissions) {
		toSerialize["hasManagementPermissions"] = o.HasManagementPermissions
	}
	return toSerialize, nil
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


