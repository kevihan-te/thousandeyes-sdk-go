/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the EpochTimeWindow type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EpochTimeWindow{}

// EpochTimeWindow struct for EpochTimeWindow
type EpochTimeWindow struct {
	// Epoch time (seconds) indicating the start time of the round
	StartTime *int32 `json:"startTime,omitempty"`
	// Epoch time (seconds) indicating the end time of the round
	EndTime *int32 `json:"endTime,omitempty"`
}

// NewEpochTimeWindow instantiates a new EpochTimeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEpochTimeWindow() *EpochTimeWindow {
	this := EpochTimeWindow{}
	return &this
}

// NewEpochTimeWindowWithDefaults instantiates a new EpochTimeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEpochTimeWindowWithDefaults() *EpochTimeWindow {
	this := EpochTimeWindow{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *EpochTimeWindow) GetStartTime() int32 {
	if o == nil || utils.IsNil(o.StartTime) {
		var ret int32
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpochTimeWindow) GetStartTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *EpochTimeWindow) HasStartTime() bool {
	if o != nil && !utils.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int32 and assigns it to the StartTime field.
func (o *EpochTimeWindow) SetStartTime(v int32) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *EpochTimeWindow) GetEndTime() int32 {
	if o == nil || utils.IsNil(o.EndTime) {
		var ret int32
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EpochTimeWindow) GetEndTimeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *EpochTimeWindow) HasEndTime() bool {
	if o != nil && !utils.IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given int32 and assigns it to the EndTime field.
func (o *EpochTimeWindow) SetEndTime(v int32) {
	o.EndTime = &v
}

func (o EpochTimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EpochTimeWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !utils.IsNil(o.EndTime) {
		toSerialize["endTime"] = o.EndTime
	}
	return toSerialize, nil
}

type NullableEpochTimeWindow struct {
	value *EpochTimeWindow
	isSet bool
}

func (v NullableEpochTimeWindow) Get() *EpochTimeWindow {
	return v.value
}

func (v *NullableEpochTimeWindow) Set(val *EpochTimeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableEpochTimeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableEpochTimeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpochTimeWindow(val *EpochTimeWindow) *NullableEpochTimeWindow {
	return &NullableEpochTimeWindow{value: val, isSet: true}
}

func (v NullableEpochTimeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpochTimeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


