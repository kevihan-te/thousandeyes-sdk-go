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

// checks if the BulkTagAssignment type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &BulkTagAssignment{}

// BulkTagAssignment struct for BulkTagAssignment
type BulkTagAssignment struct {
	Assignments []Assignment `json:"assignments,omitempty"`
	// The ID of the tag to assign
	TagId *string `json:"tagId,omitempty"`
}

// NewBulkTagAssignment instantiates a new BulkTagAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkTagAssignment() *BulkTagAssignment {
	this := BulkTagAssignment{}
	return &this
}

// NewBulkTagAssignmentWithDefaults instantiates a new BulkTagAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkTagAssignmentWithDefaults() *BulkTagAssignment {
	this := BulkTagAssignment{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *BulkTagAssignment) GetAssignments() []Assignment {
	if o == nil || utils.IsNil(o.Assignments) {
		var ret []Assignment
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTagAssignment) GetAssignmentsOk() ([]Assignment, bool) {
	if o == nil || utils.IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *BulkTagAssignment) HasAssignments() bool {
	if o != nil && !utils.IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []Assignment and assigns it to the Assignments field.
func (o *BulkTagAssignment) SetAssignments(v []Assignment) {
	o.Assignments = v
}

// GetTagId returns the TagId field value if set, zero value otherwise.
func (o *BulkTagAssignment) GetTagId() string {
	if o == nil || utils.IsNil(o.TagId) {
		var ret string
		return ret
	}
	return *o.TagId
}

// GetTagIdOk returns a tuple with the TagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkTagAssignment) GetTagIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TagId) {
		return nil, false
	}
	return o.TagId, true
}

// HasTagId returns a boolean if a field has been set.
func (o *BulkTagAssignment) HasTagId() bool {
	if o != nil && !utils.IsNil(o.TagId) {
		return true
	}

	return false
}

// SetTagId gets a reference to the given string and assigns it to the TagId field.
func (o *BulkTagAssignment) SetTagId(v string) {
	o.TagId = &v
}

func (o BulkTagAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkTagAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	if !utils.IsNil(o.TagId) {
		toSerialize["tagId"] = o.TagId
	}
	return toSerialize, nil
}

type NullableBulkTagAssignment struct {
	value *BulkTagAssignment
	isSet bool
}

func (v NullableBulkTagAssignment) Get() *BulkTagAssignment {
	return v.value
}

func (v *NullableBulkTagAssignment) Set(val *BulkTagAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkTagAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkTagAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkTagAssignment(val *BulkTagAssignment) *NullableBulkTagAssignment {
	return &NullableBulkTagAssignment{value: val, isSet: true}
}

func (v NullableBulkTagAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkTagAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


