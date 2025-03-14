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

// checks if the NetworkMetrics type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &NetworkMetrics{}

// NetworkMetrics struct for NetworkMetrics
type NetworkMetrics struct {
	// Network jitter.
	Jitter *int32 `json:"jitter,omitempty"`
	// Network latency.
	Latency *int32 `json:"latency,omitempty"`
	// Network loss.
	Loss *float64 `json:"loss,omitempty"`
	// Network target IP address.
	Target *string `json:"target,omitempty"`
}

// NewNetworkMetrics instantiates a new NetworkMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkMetrics() *NetworkMetrics {
	this := NetworkMetrics{}
	return &this
}

// NewNetworkMetricsWithDefaults instantiates a new NetworkMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkMetricsWithDefaults() *NetworkMetrics {
	this := NetworkMetrics{}
	return &this
}

// GetJitter returns the Jitter field value if set, zero value otherwise.
func (o *NetworkMetrics) GetJitter() int32 {
	if o == nil || utils.IsNil(o.Jitter) {
		var ret int32
		return ret
	}
	return *o.Jitter
}

// GetJitterOk returns a tuple with the Jitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetrics) GetJitterOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Jitter) {
		return nil, false
	}
	return o.Jitter, true
}

// HasJitter returns a boolean if a field has been set.
func (o *NetworkMetrics) HasJitter() bool {
	if o != nil && !utils.IsNil(o.Jitter) {
		return true
	}

	return false
}

// SetJitter gets a reference to the given int32 and assigns it to the Jitter field.
func (o *NetworkMetrics) SetJitter(v int32) {
	o.Jitter = &v
}

// GetLatency returns the Latency field value if set, zero value otherwise.
func (o *NetworkMetrics) GetLatency() int32 {
	if o == nil || utils.IsNil(o.Latency) {
		var ret int32
		return ret
	}
	return *o.Latency
}

// GetLatencyOk returns a tuple with the Latency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetrics) GetLatencyOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Latency) {
		return nil, false
	}
	return o.Latency, true
}

// HasLatency returns a boolean if a field has been set.
func (o *NetworkMetrics) HasLatency() bool {
	if o != nil && !utils.IsNil(o.Latency) {
		return true
	}

	return false
}

// SetLatency gets a reference to the given int32 and assigns it to the Latency field.
func (o *NetworkMetrics) SetLatency(v int32) {
	o.Latency = &v
}

// GetLoss returns the Loss field value if set, zero value otherwise.
func (o *NetworkMetrics) GetLoss() float64 {
	if o == nil || utils.IsNil(o.Loss) {
		var ret float64
		return ret
	}
	return *o.Loss
}

// GetLossOk returns a tuple with the Loss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetrics) GetLossOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.Loss) {
		return nil, false
	}
	return o.Loss, true
}

// HasLoss returns a boolean if a field has been set.
func (o *NetworkMetrics) HasLoss() bool {
	if o != nil && !utils.IsNil(o.Loss) {
		return true
	}

	return false
}

// SetLoss gets a reference to the given float64 and assigns it to the Loss field.
func (o *NetworkMetrics) SetLoss(v float64) {
	o.Loss = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *NetworkMetrics) GetTarget() string {
	if o == nil || utils.IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkMetrics) GetTargetOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *NetworkMetrics) HasTarget() bool {
	if o != nil && !utils.IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *NetworkMetrics) SetTarget(v string) {
	o.Target = &v
}

func (o NetworkMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Jitter) {
		toSerialize["jitter"] = o.Jitter
	}
	if !utils.IsNil(o.Latency) {
		toSerialize["latency"] = o.Latency
	}
	if !utils.IsNil(o.Loss) {
		toSerialize["loss"] = o.Loss
	}
	if !utils.IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableNetworkMetrics struct {
	value *NetworkMetrics
	isSet bool
}

func (v NullableNetworkMetrics) Get() *NetworkMetrics {
	return v.value
}

func (v *NullableNetworkMetrics) Set(val *NetworkMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkMetrics(val *NetworkMetrics) *NullableNetworkMetrics {
	return &NullableNetworkMetrics{value: val, isSet: true}
}

func (v NullableNetworkMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


