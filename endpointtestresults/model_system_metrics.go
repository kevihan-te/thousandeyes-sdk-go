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

// checks if the SystemMetrics type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SystemMetrics{}

// SystemMetrics struct for SystemMetrics
type SystemMetrics struct {
	// The start time of metrics collection, expressed in milliseconds since the Epoch.
	StartTimeMs *int64 `json:"startTimeMs,omitempty"`
	// The end time of metrics collection, expressed in milliseconds since the Epoch.
	EndTimeMs *int64 `json:"endTimeMs,omitempty"`
	CpuUtilization *CpuUtilization `json:"cpuUtilization,omitempty"`
	PhysicalMemoryUsedBytes *PhysicalMemoryUsedBytes `json:"physicalMemoryUsedBytes,omitempty"`
	// Total physical memory of the system.
	PhysicalMemoryTotalBytes *int64 `json:"physicalMemoryTotalBytes,omitempty"`
}

// NewSystemMetrics instantiates a new SystemMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSystemMetrics() *SystemMetrics {
	this := SystemMetrics{}
	return &this
}

// NewSystemMetricsWithDefaults instantiates a new SystemMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSystemMetricsWithDefaults() *SystemMetrics {
	this := SystemMetrics{}
	return &this
}

// GetStartTimeMs returns the StartTimeMs field value if set, zero value otherwise.
func (o *SystemMetrics) GetStartTimeMs() int64 {
	if o == nil || utils.IsNil(o.StartTimeMs) {
		var ret int64
		return ret
	}
	return *o.StartTimeMs
}

// GetStartTimeMsOk returns a tuple with the StartTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMetrics) GetStartTimeMsOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.StartTimeMs) {
		return nil, false
	}
	return o.StartTimeMs, true
}

// HasStartTimeMs returns a boolean if a field has been set.
func (o *SystemMetrics) HasStartTimeMs() bool {
	if o != nil && !utils.IsNil(o.StartTimeMs) {
		return true
	}

	return false
}

// SetStartTimeMs gets a reference to the given int64 and assigns it to the StartTimeMs field.
func (o *SystemMetrics) SetStartTimeMs(v int64) {
	o.StartTimeMs = &v
}

// GetEndTimeMs returns the EndTimeMs field value if set, zero value otherwise.
func (o *SystemMetrics) GetEndTimeMs() int64 {
	if o == nil || utils.IsNil(o.EndTimeMs) {
		var ret int64
		return ret
	}
	return *o.EndTimeMs
}

// GetEndTimeMsOk returns a tuple with the EndTimeMs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMetrics) GetEndTimeMsOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.EndTimeMs) {
		return nil, false
	}
	return o.EndTimeMs, true
}

// HasEndTimeMs returns a boolean if a field has been set.
func (o *SystemMetrics) HasEndTimeMs() bool {
	if o != nil && !utils.IsNil(o.EndTimeMs) {
		return true
	}

	return false
}

// SetEndTimeMs gets a reference to the given int64 and assigns it to the EndTimeMs field.
func (o *SystemMetrics) SetEndTimeMs(v int64) {
	o.EndTimeMs = &v
}

// GetCpuUtilization returns the CpuUtilization field value if set, zero value otherwise.
func (o *SystemMetrics) GetCpuUtilization() CpuUtilization {
	if o == nil || utils.IsNil(o.CpuUtilization) {
		var ret CpuUtilization
		return ret
	}
	return *o.CpuUtilization
}

// GetCpuUtilizationOk returns a tuple with the CpuUtilization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMetrics) GetCpuUtilizationOk() (*CpuUtilization, bool) {
	if o == nil || utils.IsNil(o.CpuUtilization) {
		return nil, false
	}
	return o.CpuUtilization, true
}

// HasCpuUtilization returns a boolean if a field has been set.
func (o *SystemMetrics) HasCpuUtilization() bool {
	if o != nil && !utils.IsNil(o.CpuUtilization) {
		return true
	}

	return false
}

// SetCpuUtilization gets a reference to the given CpuUtilization and assigns it to the CpuUtilization field.
func (o *SystemMetrics) SetCpuUtilization(v CpuUtilization) {
	o.CpuUtilization = &v
}

// GetPhysicalMemoryUsedBytes returns the PhysicalMemoryUsedBytes field value if set, zero value otherwise.
func (o *SystemMetrics) GetPhysicalMemoryUsedBytes() PhysicalMemoryUsedBytes {
	if o == nil || utils.IsNil(o.PhysicalMemoryUsedBytes) {
		var ret PhysicalMemoryUsedBytes
		return ret
	}
	return *o.PhysicalMemoryUsedBytes
}

// GetPhysicalMemoryUsedBytesOk returns a tuple with the PhysicalMemoryUsedBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMetrics) GetPhysicalMemoryUsedBytesOk() (*PhysicalMemoryUsedBytes, bool) {
	if o == nil || utils.IsNil(o.PhysicalMemoryUsedBytes) {
		return nil, false
	}
	return o.PhysicalMemoryUsedBytes, true
}

// HasPhysicalMemoryUsedBytes returns a boolean if a field has been set.
func (o *SystemMetrics) HasPhysicalMemoryUsedBytes() bool {
	if o != nil && !utils.IsNil(o.PhysicalMemoryUsedBytes) {
		return true
	}

	return false
}

// SetPhysicalMemoryUsedBytes gets a reference to the given PhysicalMemoryUsedBytes and assigns it to the PhysicalMemoryUsedBytes field.
func (o *SystemMetrics) SetPhysicalMemoryUsedBytes(v PhysicalMemoryUsedBytes) {
	o.PhysicalMemoryUsedBytes = &v
}

// GetPhysicalMemoryTotalBytes returns the PhysicalMemoryTotalBytes field value if set, zero value otherwise.
func (o *SystemMetrics) GetPhysicalMemoryTotalBytes() int64 {
	if o == nil || utils.IsNil(o.PhysicalMemoryTotalBytes) {
		var ret int64
		return ret
	}
	return *o.PhysicalMemoryTotalBytes
}

// GetPhysicalMemoryTotalBytesOk returns a tuple with the PhysicalMemoryTotalBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SystemMetrics) GetPhysicalMemoryTotalBytesOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PhysicalMemoryTotalBytes) {
		return nil, false
	}
	return o.PhysicalMemoryTotalBytes, true
}

// HasPhysicalMemoryTotalBytes returns a boolean if a field has been set.
func (o *SystemMetrics) HasPhysicalMemoryTotalBytes() bool {
	if o != nil && !utils.IsNil(o.PhysicalMemoryTotalBytes) {
		return true
	}

	return false
}

// SetPhysicalMemoryTotalBytes gets a reference to the given int64 and assigns it to the PhysicalMemoryTotalBytes field.
func (o *SystemMetrics) SetPhysicalMemoryTotalBytes(v int64) {
	o.PhysicalMemoryTotalBytes = &v
}

func (o SystemMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SystemMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.StartTimeMs) {
		toSerialize["startTimeMs"] = o.StartTimeMs
	}
	if !utils.IsNil(o.EndTimeMs) {
		toSerialize["endTimeMs"] = o.EndTimeMs
	}
	if !utils.IsNil(o.CpuUtilization) {
		toSerialize["cpuUtilization"] = o.CpuUtilization
	}
	if !utils.IsNil(o.PhysicalMemoryUsedBytes) {
		toSerialize["physicalMemoryUsedBytes"] = o.PhysicalMemoryUsedBytes
	}
	if !utils.IsNil(o.PhysicalMemoryTotalBytes) {
		toSerialize["physicalMemoryTotalBytes"] = o.PhysicalMemoryTotalBytes
	}
	return toSerialize, nil
}

type NullableSystemMetrics struct {
	value *SystemMetrics
	isSet bool
}

func (v NullableSystemMetrics) Get() *SystemMetrics {
	return v.value
}

func (v *NullableSystemMetrics) Set(val *SystemMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemMetrics(val *SystemMetrics) *NullableSystemMetrics {
	return &NullableSystemMetrics{value: val, isSet: true}
}

func (v NullableSystemMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


