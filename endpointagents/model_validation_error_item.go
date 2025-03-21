/*
Endpoint Agents API

Manage ThousandEyes Endpoint Agents using this API.   For more information about Endpoint Agents, see [Endpoint Agents](https://docs.thousandeyes.com/product-documentation/global-vantage-points/endpoint-agents).

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointagents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ValidationErrorItem type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ValidationErrorItem{}

// ValidationErrorItem struct for ValidationErrorItem
type ValidationErrorItem struct {
	// (Optional) A unique error type/code that can be referenced in the documentation for further details.
	Code *string `json:"code,omitempty"`
	// Identifies the field that triggered this particular error.
	Field *string `json:"field,omitempty"`
	// A short, human-readable summary of the error.
	Message *string `json:"message,omitempty"`
}

// NewValidationErrorItem instantiates a new ValidationErrorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationErrorItem() *ValidationErrorItem {
	this := ValidationErrorItem{}
	return &this
}

// NewValidationErrorItemWithDefaults instantiates a new ValidationErrorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorItemWithDefaults() *ValidationErrorItem {
	this := ValidationErrorItem{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ValidationErrorItem) GetCode() string {
	if o == nil || utils.IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorItem) GetCodeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ValidationErrorItem) HasCode() bool {
	if o != nil && !utils.IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ValidationErrorItem) SetCode(v string) {
	o.Code = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ValidationErrorItem) GetField() string {
	if o == nil || utils.IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorItem) GetFieldOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ValidationErrorItem) HasField() bool {
	if o != nil && !utils.IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ValidationErrorItem) SetField(v string) {
	o.Field = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ValidationErrorItem) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidationErrorItem) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ValidationErrorItem) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ValidationErrorItem) SetMessage(v string) {
	o.Message = &v
}

func (o ValidationErrorItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationErrorItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !utils.IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableValidationErrorItem struct {
	value *ValidationErrorItem
	isSet bool
}

func (v NullableValidationErrorItem) Get() *ValidationErrorItem {
	return v.value
}

func (v *NullableValidationErrorItem) Set(val *ValidationErrorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorItem(val *ValidationErrorItem) *NullableValidationErrorItem {
	return &NullableValidationErrorItem{value: val, isSet: true}
}

func (v NullableValidationErrorItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


