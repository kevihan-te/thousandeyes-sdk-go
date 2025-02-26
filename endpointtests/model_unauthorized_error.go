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
)

// checks if the UnauthorizedError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UnauthorizedError{}

// UnauthorizedError struct for UnauthorizedError
type UnauthorizedError struct {
	Error *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}

// NewUnauthorizedError instantiates a new UnauthorizedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnauthorizedError() *UnauthorizedError {
	this := UnauthorizedError{}
	return &this
}

// NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnauthorizedErrorWithDefaults() *UnauthorizedError {
	this := UnauthorizedError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *UnauthorizedError) GetError() string {
	if o == nil || utils.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetErrorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *UnauthorizedError) HasError() bool {
	if o != nil && !utils.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *UnauthorizedError) SetError(v string) {
	o.Error = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *UnauthorizedError) GetErrorDescription() string {
	if o == nil || utils.IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnauthorizedError) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *UnauthorizedError) HasErrorDescription() bool {
	if o != nil && !utils.IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *UnauthorizedError) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

func (o UnauthorizedError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnauthorizedError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !utils.IsNil(o.ErrorDescription) {
		toSerialize["error_description"] = o.ErrorDescription
	}
	return toSerialize, nil
}

type NullableUnauthorizedError struct {
	value *UnauthorizedError
	isSet bool
}

func (v NullableUnauthorizedError) Get() *UnauthorizedError {
	return v.value
}

func (v *NullableUnauthorizedError) Set(val *UnauthorizedError) {
	v.value = val
	v.isSet = true
}

func (v NullableUnauthorizedError) IsSet() bool {
	return v.isSet
}

func (v *NullableUnauthorizedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnauthorizedError(val *UnauthorizedError) *NullableUnauthorizedError {
	return &NullableUnauthorizedError{value: val, isSet: true}
}

func (v NullableUnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnauthorizedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


