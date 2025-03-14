/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ApiRequestVariable type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiRequestVariable{}

// ApiRequestVariable struct for ApiRequestVariable
type ApiRequestVariable struct {
	// Variable name
	Name *string `json:"name,omitempty"`
	// The JSON path of data within the Response Body to assign to this variable.
	Value *string `json:"value,omitempty"`
}

// NewApiRequestVariable instantiates a new ApiRequestVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiRequestVariable() *ApiRequestVariable {
	this := ApiRequestVariable{}
	return &this
}

// NewApiRequestVariableWithDefaults instantiates a new ApiRequestVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiRequestVariableWithDefaults() *ApiRequestVariable {
	this := ApiRequestVariable{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiRequestVariable) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequestVariable) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiRequestVariable) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiRequestVariable) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiRequestVariable) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiRequestVariable) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiRequestVariable) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiRequestVariable) SetValue(v string) {
	o.Value = &v
}

func (o ApiRequestVariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiRequestVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiRequestVariable struct {
	value *ApiRequestVariable
	isSet bool
}

func (v NullableApiRequestVariable) Get() *ApiRequestVariable {
	return v.value
}

func (v *NullableApiRequestVariable) Set(val *ApiRequestVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRequestVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRequestVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRequestVariable(val *ApiRequestVariable) *NullableApiRequestVariable {
	return &NullableApiRequestVariable{value: val, isSet: true}
}

func (v NullableApiRequestVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRequestVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


