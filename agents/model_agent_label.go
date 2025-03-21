/*
Agents API

 ## Overview Manage Cloud and Enterprise Agents available to your account in ThousandEyes.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the AgentLabel type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AgentLabel{}

// AgentLabel struct for AgentLabel
type AgentLabel struct {
	// Label Id.
	LabelId *string `json:"labelId,omitempty"`
	// Name of the label.
	Name *string `json:"name,omitempty"`
}

// NewAgentLabel instantiates a new AgentLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentLabel() *AgentLabel {
	this := AgentLabel{}
	return &this
}

// NewAgentLabelWithDefaults instantiates a new AgentLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentLabelWithDefaults() *AgentLabel {
	this := AgentLabel{}
	return &this
}

// GetLabelId returns the LabelId field value if set, zero value otherwise.
func (o *AgentLabel) GetLabelId() string {
	if o == nil || utils.IsNil(o.LabelId) {
		var ret string
		return ret
	}
	return *o.LabelId
}

// GetLabelIdOk returns a tuple with the LabelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentLabel) GetLabelIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.LabelId) {
		return nil, false
	}
	return o.LabelId, true
}

// HasLabelId returns a boolean if a field has been set.
func (o *AgentLabel) HasLabelId() bool {
	if o != nil && !utils.IsNil(o.LabelId) {
		return true
	}

	return false
}

// SetLabelId gets a reference to the given string and assigns it to the LabelId field.
func (o *AgentLabel) SetLabelId(v string) {
	o.LabelId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentLabel) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentLabel) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentLabel) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentLabel) SetName(v string) {
	o.Name = &v
}

func (o AgentLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.LabelId) {
		toSerialize["labelId"] = o.LabelId
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableAgentLabel struct {
	value *AgentLabel
	isSet bool
}

func (v NullableAgentLabel) Get() *AgentLabel {
	return v.value
}

func (v *NullableAgentLabel) Set(val *AgentLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentLabel(val *AgentLabel) *NullableAgentLabel {
	return &NullableAgentLabel{value: val, isSet: true}
}

func (v NullableAgentLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


