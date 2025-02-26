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

// checks if the UserRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UserRequest{}

// UserRequest struct for UserRequest
type UserRequest struct {
	// User's display name.
	Name *string `json:"name,omitempty"`
	// User's email address.
	Email *string `json:"email,omitempty"`
	// Unique ID of the login account group.
	LoginAccountGroupId *string `json:"loginAccountGroupId,omitempty"`
	AccountGroupRoles []UserAccountGroupRole `json:"accountGroupRoles,omitempty"`
	// Unique IDs representing the roles.
	AllAccountGroupRoleIds []string `json:"allAccountGroupRoleIds,omitempty"`
}

// NewUserRequest instantiates a new UserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRequest() *UserRequest {
	this := UserRequest{}
	return &this
}

// NewUserRequestWithDefaults instantiates a new UserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRequestWithDefaults() *UserRequest {
	this := UserRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserRequest) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserRequest) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserRequest) SetName(v string) {
	o.Name = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserRequest) GetEmail() string {
	if o == nil || utils.IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetEmailOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserRequest) HasEmail() bool {
	if o != nil && !utils.IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetLoginAccountGroupId returns the LoginAccountGroupId field value if set, zero value otherwise.
func (o *UserRequest) GetLoginAccountGroupId() string {
	if o == nil || utils.IsNil(o.LoginAccountGroupId) {
		var ret string
		return ret
	}
	return *o.LoginAccountGroupId
}

// GetLoginAccountGroupIdOk returns a tuple with the LoginAccountGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetLoginAccountGroupIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.LoginAccountGroupId) {
		return nil, false
	}
	return o.LoginAccountGroupId, true
}

// HasLoginAccountGroupId returns a boolean if a field has been set.
func (o *UserRequest) HasLoginAccountGroupId() bool {
	if o != nil && !utils.IsNil(o.LoginAccountGroupId) {
		return true
	}

	return false
}

// SetLoginAccountGroupId gets a reference to the given string and assigns it to the LoginAccountGroupId field.
func (o *UserRequest) SetLoginAccountGroupId(v string) {
	o.LoginAccountGroupId = &v
}

// GetAccountGroupRoles returns the AccountGroupRoles field value if set, zero value otherwise.
func (o *UserRequest) GetAccountGroupRoles() []UserAccountGroupRole {
	if o == nil || utils.IsNil(o.AccountGroupRoles) {
		var ret []UserAccountGroupRole
		return ret
	}
	return o.AccountGroupRoles
}

// GetAccountGroupRolesOk returns a tuple with the AccountGroupRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetAccountGroupRolesOk() ([]UserAccountGroupRole, bool) {
	if o == nil || utils.IsNil(o.AccountGroupRoles) {
		return nil, false
	}
	return o.AccountGroupRoles, true
}

// HasAccountGroupRoles returns a boolean if a field has been set.
func (o *UserRequest) HasAccountGroupRoles() bool {
	if o != nil && !utils.IsNil(o.AccountGroupRoles) {
		return true
	}

	return false
}

// SetAccountGroupRoles gets a reference to the given []UserAccountGroupRole and assigns it to the AccountGroupRoles field.
func (o *UserRequest) SetAccountGroupRoles(v []UserAccountGroupRole) {
	o.AccountGroupRoles = v
}

// GetAllAccountGroupRoleIds returns the AllAccountGroupRoleIds field value if set, zero value otherwise.
func (o *UserRequest) GetAllAccountGroupRoleIds() []string {
	if o == nil || utils.IsNil(o.AllAccountGroupRoleIds) {
		var ret []string
		return ret
	}
	return o.AllAccountGroupRoleIds
}

// GetAllAccountGroupRoleIdsOk returns a tuple with the AllAccountGroupRoleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRequest) GetAllAccountGroupRoleIdsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AllAccountGroupRoleIds) {
		return nil, false
	}
	return o.AllAccountGroupRoleIds, true
}

// HasAllAccountGroupRoleIds returns a boolean if a field has been set.
func (o *UserRequest) HasAllAccountGroupRoleIds() bool {
	if o != nil && !utils.IsNil(o.AllAccountGroupRoleIds) {
		return true
	}

	return false
}

// SetAllAccountGroupRoleIds gets a reference to the given []string and assigns it to the AllAccountGroupRoleIds field.
func (o *UserRequest) SetAllAccountGroupRoleIds(v []string) {
	o.AllAccountGroupRoleIds = v
}

func (o UserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !utils.IsNil(o.LoginAccountGroupId) {
		toSerialize["loginAccountGroupId"] = o.LoginAccountGroupId
	}
	if !utils.IsNil(o.AccountGroupRoles) {
		toSerialize["accountGroupRoles"] = o.AccountGroupRoles
	}
	if !utils.IsNil(o.AllAccountGroupRoleIds) {
		toSerialize["allAccountGroupRoleIds"] = o.AllAccountGroupRoleIds
	}
	return toSerialize, nil
}

type NullableUserRequest struct {
	value *UserRequest
	isSet bool
}

func (v NullableUserRequest) Get() *UserRequest {
	return v.value
}

func (v *NullableUserRequest) Set(val *UserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRequest(val *UserRequest) *NullableUserRequest {
	return &NullableUserRequest{value: val, isSet: true}
}

func (v NullableUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


