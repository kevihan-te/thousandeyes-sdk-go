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

// checks if the ApiPredefinedVariable type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiPredefinedVariable{}

// ApiPredefinedVariable struct for ApiPredefinedVariable
type ApiPredefinedVariable struct {
	// Variable name. Must be unique.
	Name *string `json:"name,omitempty"`
	// Variable value, will be treated as string.
	Value *string `json:"value,omitempty"`
}

// NewApiPredefinedVariable instantiates a new ApiPredefinedVariable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPredefinedVariable() *ApiPredefinedVariable {
	this := ApiPredefinedVariable{}
	return &this
}

// NewApiPredefinedVariableWithDefaults instantiates a new ApiPredefinedVariable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPredefinedVariableWithDefaults() *ApiPredefinedVariable {
	this := ApiPredefinedVariable{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiPredefinedVariable) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPredefinedVariable) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiPredefinedVariable) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiPredefinedVariable) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiPredefinedVariable) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPredefinedVariable) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiPredefinedVariable) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiPredefinedVariable) SetValue(v string) {
	o.Value = &v
}

func (o ApiPredefinedVariable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPredefinedVariable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiPredefinedVariable struct {
	value *ApiPredefinedVariable
	isSet bool
}

func (v NullableApiPredefinedVariable) Get() *ApiPredefinedVariable {
	return v.value
}

func (v *NullableApiPredefinedVariable) Set(val *ApiPredefinedVariable) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPredefinedVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPredefinedVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPredefinedVariable(val *ApiPredefinedVariable) *NullableApiPredefinedVariable {
	return &NullableApiPredefinedVariable{value: val, isSet: true}
}

func (v NullableApiPredefinedVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPredefinedVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


