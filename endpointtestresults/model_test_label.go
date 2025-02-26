/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the TestLabel type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestLabel{}

// TestLabel struct for TestLabel
type TestLabel struct {
	// Label ID.
	LabelId *string `json:"labelId,omitempty"`
	// Name of the label.
	Name *string `json:"name,omitempty"`
	// Value indicating if the label in question is BuiltIn (Account Admin, Organization Admin, Regular User).
	IsBuiltin *bool `json:"isBuiltin,omitempty"`
}

// NewTestLabel instantiates a new TestLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestLabel() *TestLabel {
	this := TestLabel{}
	return &this
}

// NewTestLabelWithDefaults instantiates a new TestLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestLabelWithDefaults() *TestLabel {
	this := TestLabel{}
	return &this
}

// GetLabelId returns the LabelId field value if set, zero value otherwise.
func (o *TestLabel) GetLabelId() string {
	if o == nil || utils.IsNil(o.LabelId) {
		var ret string
		return ret
	}
	return *o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestLabel) GetLabelIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.LabelId) {
		return nil, false
	}
	return o.LabelId, true
}

// HasLabelId returns a boolean if a field has been set.
func (o *TestLabel) HasLabelId() bool {
	if o != nil && !utils.IsNil(o.LabelId) {
		return true
	}

	return false
}

// SetLabelId gets a reference to the given string and assigns it to the LabelId field.
func (o *TestLabel) SetLabelId(v string) {
	o.LabelId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TestLabel) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestLabel) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TestLabel) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TestLabel) SetName(v string) {
	o.Name = &v
}

// GetIsBuiltin returns the IsBuiltin field value if set, zero value otherwise.
func (o *TestLabel) GetIsBuiltin() bool {
	if o == nil || utils.IsNil(o.IsBuiltin) {
		var ret bool
		return ret
	}
	return *o.IsBuiltin
}

// GetIsBuiltinOk returns a tuple with the IsBuiltin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestLabel) GetIsBuiltinOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsBuiltin) {
		return nil, false
	}
	return o.IsBuiltin, true
}

// HasIsBuiltin returns a boolean if a field has been set.
func (o *TestLabel) HasIsBuiltin() bool {
	if o != nil && !utils.IsNil(o.IsBuiltin) {
		return true
	}

	return false
}

// SetIsBuiltin gets a reference to the given bool and assigns it to the IsBuiltin field.
func (o *TestLabel) SetIsBuiltin(v bool) {
	o.IsBuiltin = &v
}

func (o TestLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.LabelId) {
		toSerialize["labelId"] = o.LabelId
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.IsBuiltin) {
		toSerialize["isBuiltin"] = o.IsBuiltin
	}
	return toSerialize, nil
}

type NullableTestLabel struct {
	value *TestLabel
	isSet bool
}

func (v NullableTestLabel) Get() *TestLabel {
	return v.value
}

func (v *NullableTestLabel) Set(val *TestLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableTestLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableTestLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestLabel(val *TestLabel) *NullableTestLabel {
	return &NullableTestLabel{value: val, isSet: true}
}

func (v NullableTestLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


