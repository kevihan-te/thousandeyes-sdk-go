/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ApiReportDataComponentLabelMapEntry type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiReportDataComponentLabelMapEntry{}

// ApiReportDataComponentLabelMapEntry Represents a mapping entry of a group label.
type ApiReportDataComponentLabelMapEntry struct {
	// Identifier of the group.
	GroupId *string `json:"groupId,omitempty"`
	// Label of the group.
	GroupLabel *string `json:"groupLabel,omitempty"`
}

// NewApiReportDataComponentLabelMapEntry instantiates a new ApiReportDataComponentLabelMapEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReportDataComponentLabelMapEntry() *ApiReportDataComponentLabelMapEntry {
	this := ApiReportDataComponentLabelMapEntry{}
	return &this
}

// NewApiReportDataComponentLabelMapEntryWithDefaults instantiates a new ApiReportDataComponentLabelMapEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReportDataComponentLabelMapEntryWithDefaults() *ApiReportDataComponentLabelMapEntry {
	this := ApiReportDataComponentLabelMapEntry{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *ApiReportDataComponentLabelMapEntry) GetGroupId() string {
	if o == nil || utils.IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportDataComponentLabelMapEntry) GetGroupIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *ApiReportDataComponentLabelMapEntry) HasGroupId() bool {
	if o != nil && !utils.IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *ApiReportDataComponentLabelMapEntry) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupLabel returns the GroupLabel field value if set, zero value otherwise.
func (o *ApiReportDataComponentLabelMapEntry) GetGroupLabel() string {
	if o == nil || utils.IsNil(o.GroupLabel) {
		var ret string
		return ret
	}
	return *o.GroupLabel
}

// GetGroupLabelOk returns a tuple with the GroupLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReportDataComponentLabelMapEntry) GetGroupLabelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.GroupLabel) {
		return nil, false
	}
	return o.GroupLabel, true
}

// HasGroupLabel returns a boolean if a field has been set.
func (o *ApiReportDataComponentLabelMapEntry) HasGroupLabel() bool {
	if o != nil && !utils.IsNil(o.GroupLabel) {
		return true
	}

	return false
}

// SetGroupLabel gets a reference to the given string and assigns it to the GroupLabel field.
func (o *ApiReportDataComponentLabelMapEntry) SetGroupLabel(v string) {
	o.GroupLabel = &v
}

func (o ApiReportDataComponentLabelMapEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReportDataComponentLabelMapEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !utils.IsNil(o.GroupLabel) {
		toSerialize["groupLabel"] = o.GroupLabel
	}
	return toSerialize, nil
}

type NullableApiReportDataComponentLabelMapEntry struct {
	value *ApiReportDataComponentLabelMapEntry
	isSet bool
}

func (v NullableApiReportDataComponentLabelMapEntry) Get() *ApiReportDataComponentLabelMapEntry {
	return v.value
}

func (v *NullableApiReportDataComponentLabelMapEntry) Set(val *ApiReportDataComponentLabelMapEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReportDataComponentLabelMapEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReportDataComponentLabelMapEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReportDataComponentLabelMapEntry(val *ApiReportDataComponentLabelMapEntry) *NullableApiReportDataComponentLabelMapEntry {
	return &NullableApiReportDataComponentLabelMapEntry{value: val, isSet: true}
}

func (v NullableApiReportDataComponentLabelMapEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReportDataComponentLabelMapEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


