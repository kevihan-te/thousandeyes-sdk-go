/*
Tags API

The ThousandEyes Tags API provides a tagging system with key/value pairs. It allows you to tag assets within the ThousandEyes platform (such as agents, tests, or alert rules) with meaningful metadata. For example: `branch:sfo`, `branch:nyc`, and `team:netops`.  This feature provides:  * Support for automation. * Powerful and flexible reports/dashboards. * Support for third-party integrations.  Things to note with the ThousandEyes Tags API:  * Tags are backwards-compatible with existing labels. * Tags are separated by Tests (CEA), Agents (CEA), Endpoint Agents, Scheduled Endpoint Tests, and Reports. A single tag can only apply to one type of target object, so each tag must specify the target type of object via a `type` field. * Tags are defined in a single table so that they can be represented using a single model - `Tag`. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the BulkTagAssignments type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BulkTagAssignments{}

// BulkTagAssignments struct for BulkTagAssignments
type BulkTagAssignments struct {
	Tags []BulkTagAssignment `json:"tags,omitempty"`
}

// NewBulkTagAssignments instantiates a new BulkTagAssignments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkTagAssignments() *BulkTagAssignments {
	this := BulkTagAssignments{}
	return &this
}

// NewBulkTagAssignmentsWithDefaults instantiates a new BulkTagAssignments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkTagAssignmentsWithDefaults() *BulkTagAssignments {
	this := BulkTagAssignments{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BulkTagAssignments) GetTags() []BulkTagAssignment {
	if o == nil || utils.IsNil(o.Tags) {
		var ret []BulkTagAssignment
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTagAssignments) GetTagsOk() ([]BulkTagAssignment, bool) {
	if o == nil || utils.IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BulkTagAssignments) HasTags() bool {
	if o != nil && !utils.IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []BulkTagAssignment and assigns it to the Tags field.
func (o *BulkTagAssignments) SetTags(v []BulkTagAssignment) {
	o.Tags = v
}

func (o BulkTagAssignments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkTagAssignments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableBulkTagAssignments struct {
	value *BulkTagAssignments
	isSet bool
}

func (v NullableBulkTagAssignments) Get() *BulkTagAssignments {
	return v.value
}

func (v *NullableBulkTagAssignments) Set(val *BulkTagAssignments) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkTagAssignments) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkTagAssignments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkTagAssignments(val *BulkTagAssignments) *NullableBulkTagAssignments {
	return &NullableBulkTagAssignments{value: val, isSet: true}
}

func (v NullableBulkTagAssignments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkTagAssignments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


