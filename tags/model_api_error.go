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

// checks if the ApiError type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiError{}

// ApiError struct for ApiError
type ApiError struct {
	Timestamp *int64 `json:"timestamp,omitempty"`
	Status *int32 `json:"status,omitempty"`
	Errors *string `json:"errors,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewApiError instantiates a new ApiError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiError() *ApiError {
	this := ApiError{}
	return &this
}

// NewApiErrorWithDefaults instantiates a new ApiError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorWithDefaults() *ApiError {
	this := ApiError{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ApiError) GetTimestamp() int64 {
	if o == nil || utils.IsNil(o.Timestamp) {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetTimestampOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ApiError) HasTimestamp() bool {
	if o != nil && !utils.IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *ApiError) SetTimestamp(v int64) {
	o.Timestamp = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiError) GetStatus() int32 {
	if o == nil || utils.IsNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetStatusOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiError) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *ApiError) SetStatus(v int32) {
	o.Status = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ApiError) GetErrors() string {
	if o == nil || utils.IsNil(o.Errors) {
		var ret string
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetErrorsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ApiError) HasErrors() bool {
	if o != nil && !utils.IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given string and assigns it to the Errors field.
func (o *ApiError) SetErrors(v string) {
	o.Errors = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ApiError) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiError) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ApiError) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ApiError) SetPath(v string) {
	o.Path = &v
}

func (o ApiError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !utils.IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	return toSerialize, nil
}

type NullableApiError struct {
	value *ApiError
	isSet bool
}

func (v NullableApiError) Get() *ApiError {
	return v.value
}

func (v *NullableApiError) Set(val *ApiError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiError(val *ApiError) *NullableApiError {
	return &NullableApiError{value: val, isSet: true}
}

func (v NullableApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


