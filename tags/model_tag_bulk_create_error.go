/*
Tags API

The ThousandEyes Tags API provides a tagging system with key/value pairs. It allows you to tag assets within the ThousandEyes platform (such as agents, tests, or alert rules) with meaningful metadata. For example: `branch:sfo`, `branch:nyc`, and `team:netops`.  This feature provides:  * Support for automation. * Powerful and flexible reports/dashboards. * Support for third-party integrations.  Things to note with the ThousandEyes Tags API:  * Tags are backwards-compatible with existing labels. * Tags are separated by Tests (CEA), Agents (CEA), Endpoint Agents, Scheduled Endpoint Tests, and Reports. A single tag can only apply to one type of target object, so each tag must specify the target type of object via a `type` field. * Tags are defined in a single table so that they can be represented using a single model - `Tag`. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the TagBulkCreateError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TagBulkCreateError{}

// TagBulkCreateError struct for TagBulkCreateError
type TagBulkCreateError struct {
	Tag *map[string]TagInfo `json:"tag,omitempty"`
	// HTTP response code
	ResponseCode *int32 `json:"responseCode,omitempty"`
	// Status / error message
	Message *string `json:"message,omitempty"`
}

// NewTagBulkCreateError instantiates a new TagBulkCreateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagBulkCreateError() *TagBulkCreateError {
	this := TagBulkCreateError{}
	return &this
}

// NewTagBulkCreateErrorWithDefaults instantiates a new TagBulkCreateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagBulkCreateErrorWithDefaults() *TagBulkCreateError {
	this := TagBulkCreateError{}
	return &this
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TagBulkCreateError) GetTag() map[string]TagInfo {
	if o == nil || utils.IsNil(o.Tag) {
		var ret map[string]TagInfo
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagBulkCreateError) GetTagOk() (*map[string]TagInfo, bool) {
	if o == nil || utils.IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TagBulkCreateError) HasTag() bool {
	if o != nil && !utils.IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given map[string]TagInfo and assigns it to the Tag field.
func (o *TagBulkCreateError) SetTag(v map[string]TagInfo) {
	o.Tag = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *TagBulkCreateError) GetResponseCode() int32 {
	if o == nil || utils.IsNil(o.ResponseCode) {
		var ret int32
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagBulkCreateError) GetResponseCodeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ResponseCode) {
		return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *TagBulkCreateError) HasResponseCode() bool {
	if o != nil && !utils.IsNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given int32 and assigns it to the ResponseCode field.
func (o *TagBulkCreateError) SetResponseCode(v int32) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TagBulkCreateError) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagBulkCreateError) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TagBulkCreateError) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TagBulkCreateError) SetMessage(v string) {
	o.Message = &v
}

func (o TagBulkCreateError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagBulkCreateError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !utils.IsNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableTagBulkCreateError struct {
	value *TagBulkCreateError
	isSet bool
}

func (v NullableTagBulkCreateError) Get() *TagBulkCreateError {
	return v.value
}

func (v *NullableTagBulkCreateError) Set(val *TagBulkCreateError) {
	v.value = val
	v.isSet = true
}

func (v NullableTagBulkCreateError) IsSet() bool {
	return v.isSet
}

func (v *NullableTagBulkCreateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagBulkCreateError(val *TagBulkCreateError) *NullableTagBulkCreateError {
	return &NullableTagBulkCreateError{value: val, isSet: true}
}

func (v NullableTagBulkCreateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagBulkCreateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


