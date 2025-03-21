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
	"time"
)

// checks if the ApiDefaultTimespan type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiDefaultTimespan{}

// ApiDefaultTimespan struct for ApiDefaultTimespan
type ApiDefaultTimespan struct {
	// Relative timespan in seconds.
	Duration *int64 `json:"duration,omitempty"`
	// UTC start date of the timespan range (ISO date-time format).
	Start *time.Time `json:"start,omitempty"`
	// UTC end date of the timespan range (ISO date-time format).
	End *time.Time `json:"end,omitempty"`
}

// NewApiDefaultTimespan instantiates a new ApiDefaultTimespan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDefaultTimespan() *ApiDefaultTimespan {
	this := ApiDefaultTimespan{}
	return &this
}

// NewApiDefaultTimespanWithDefaults instantiates a new ApiDefaultTimespan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDefaultTimespanWithDefaults() *ApiDefaultTimespan {
	this := ApiDefaultTimespan{}
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *ApiDefaultTimespan) GetDuration() int64 {
	if o == nil || utils.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDefaultTimespan) GetDurationOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *ApiDefaultTimespan) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *ApiDefaultTimespan) SetDuration(v int64) {
	o.Duration = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *ApiDefaultTimespan) GetStart() time.Time {
	if o == nil || utils.IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDefaultTimespan) GetStartOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *ApiDefaultTimespan) HasStart() bool {
	if o != nil && !utils.IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *ApiDefaultTimespan) SetStart(v time.Time) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *ApiDefaultTimespan) GetEnd() time.Time {
	if o == nil || utils.IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDefaultTimespan) GetEndOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *ApiDefaultTimespan) HasEnd() bool {
	if o != nil && !utils.IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *ApiDefaultTimespan) SetEnd(v time.Time) {
	o.End = &v
}

func (o ApiDefaultTimespan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDefaultTimespan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !utils.IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !utils.IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	return toSerialize, nil
}

type NullableApiDefaultTimespan struct {
	value *ApiDefaultTimespan
	isSet bool
}

func (v NullableApiDefaultTimespan) Get() *ApiDefaultTimespan {
	return v.value
}

func (v *NullableApiDefaultTimespan) Set(val *ApiDefaultTimespan) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDefaultTimespan) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDefaultTimespan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDefaultTimespan(val *ApiDefaultTimespan) *NullableApiDefaultTimespan {
	return &NullableApiDefaultTimespan{value: val, isSet: true}
}

func (v NullableApiDefaultTimespan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDefaultTimespan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


