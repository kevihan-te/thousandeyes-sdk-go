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

// checks if the CpuUtilization type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CpuUtilization{}

// CpuUtilization struct for CpuUtilization
type CpuUtilization struct {
	// The minimum sampled usage value recorded during the monitored period.
	Min *float64 `json:"min,omitempty"`
	// The maximum sampled usage value recorded during the monitored period.
	Max *float64 `json:"max,omitempty"`
	// The mean (average) sampled usage value recorded during the monitored period.
	Mean *float64 `json:"mean,omitempty"`
	// The median sampled usage value recorded during the monitored period.
	Median *float64 `json:"median,omitempty"`
	// The standard deviation of sampled usage values recorded during the monitored period.
	StdDev *float64 `json:"stdDev,omitempty"`
	// The total number of samples collected during the monitored period.
	Count *int32 `json:"count,omitempty"`
}

// NewCpuUtilization instantiates a new CpuUtilization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCpuUtilization() *CpuUtilization {
	this := CpuUtilization{}
	return &this
}

// NewCpuUtilizationWithDefaults instantiates a new CpuUtilization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuUtilizationWithDefaults() *CpuUtilization {
	this := CpuUtilization{}
	return &this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *CpuUtilization) GetMin() float64 {
	if o == nil || utils.IsNil(o.Min) {
		var ret float64
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuUtilization) GetMinOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Min) {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *CpuUtilization) HasMin() bool {
	if o != nil && !utils.IsNil(o.Min) {
		return true
	}

	return false
}

// SetMin gets a reference to the given float64 and assigns it to the Min field.
func (o *CpuUtilization) SetMin(v float64) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *CpuUtilization) GetMax() float64 {
	if o == nil || utils.IsNil(o.Max) {
		var ret float64
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuUtilization) GetMaxOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Max) {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *CpuUtilization) HasMax() bool {
	if o != nil && !utils.IsNil(o.Max) {
		return true
	}

	return false
}

// SetMax gets a reference to the given float64 and assigns it to the Max field.
func (o *CpuUtilization) SetMax(v float64) {
	o.Max = &v
}

// GetMean returns the Mean field value if set, zero value otherwise.
func (o *CpuUtilization) GetMean() float64 {
	if o == nil || utils.IsNil(o.Mean) {
		var ret float64
		return ret
	}
	return *o.Mean
}

// GetMeanOk returns a tuple with the Mean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuUtilization) GetMeanOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Mean) {
		return nil, false
	}
	return o.Mean, true
}

// HasMean returns a boolean if a field has been set.
func (o *CpuUtilization) HasMean() bool {
	if o != nil && !utils.IsNil(o.Mean) {
		return true
	}

	return false
}

// SetMean gets a reference to the given float64 and assigns it to the Mean field.
func (o *CpuUtilization) SetMean(v float64) {
	o.Mean = &v
}

// GetMedian returns the Median field value if set, zero value otherwise.
func (o *CpuUtilization) GetMedian() float64 {
	if o == nil || utils.IsNil(o.Median) {
		var ret float64
		return ret
	}
	return *o.Median
}

// GetMedianOk returns a tuple with the Median field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuUtilization) GetMedianOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Median) {
		return nil, false
	}
	return o.Median, true
}

// HasMedian returns a boolean if a field has been set.
func (o *CpuUtilization) HasMedian() bool {
	if o != nil && !utils.IsNil(o.Median) {
		return true
	}

	return false
}

// SetMedian gets a reference to the given float64 and assigns it to the Median field.
func (o *CpuUtilization) SetMedian(v float64) {
	o.Median = &v
}

// GetStdDev returns the StdDev field value if set, zero value otherwise.
func (o *CpuUtilization) GetStdDev() float64 {
	if o == nil || utils.IsNil(o.StdDev) {
		var ret float64
		return ret
	}
	return *o.StdDev
}

// GetStdDevOk returns a tuple with the StdDev field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuUtilization) GetStdDevOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.StdDev) {
		return nil, false
	}
	return o.StdDev, true
}

// HasStdDev returns a boolean if a field has been set.
func (o *CpuUtilization) HasStdDev() bool {
	if o != nil && !utils.IsNil(o.StdDev) {
		return true
	}

	return false
}

// SetStdDev gets a reference to the given float64 and assigns it to the StdDev field.
func (o *CpuUtilization) SetStdDev(v float64) {
	o.StdDev = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CpuUtilization) GetCount() int32 {
	if o == nil || utils.IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CpuUtilization) GetCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CpuUtilization) HasCount() bool {
	if o != nil && !utils.IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CpuUtilization) SetCount(v int32) {
	o.Count = &v
}

func (o CpuUtilization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuUtilization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Min) {
		toSerialize["min"] = o.Min
	}
	if !utils.IsNil(o.Max) {
		toSerialize["max"] = o.Max
	}
	if !utils.IsNil(o.Mean) {
		toSerialize["mean"] = o.Mean
	}
	if !utils.IsNil(o.Median) {
		toSerialize["median"] = o.Median
	}
	if !utils.IsNil(o.StdDev) {
		toSerialize["stdDev"] = o.StdDev
	}
	if !utils.IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableCpuUtilization struct {
	value *CpuUtilization
	isSet bool
}

func (v NullableCpuUtilization) Get() *CpuUtilization {
	return v.value
}

func (v *NullableCpuUtilization) Set(val *CpuUtilization) {
	v.value = val
	v.isSet = true
}

func (v NullableCpuUtilization) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuUtilization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuUtilization(val *CpuUtilization) *NullableCpuUtilization {
	return &NullableCpuUtilization{value: val, isSet: true}
}

func (v NullableCpuUtilization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuUtilization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


