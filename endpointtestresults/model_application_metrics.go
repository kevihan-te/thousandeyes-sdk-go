/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ApplicationMetrics type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplicationMetrics{}

// ApplicationMetrics struct for ApplicationMetrics
type ApplicationMetrics struct {
	// The name of the application.
	Name *string `json:"name,omitempty"`
	// The total CPU usage by all application processes.
	TotalCpu *float64 `json:"totalCpu,omitempty"`
	// The total percentage of memory used by all application processes.
	TotalMemoryPercentage *float64 `json:"totalMemoryPercentage,omitempty"`
	// The total memory in bytes used by all application processes.
	TotalMemoryBytes *int64 `json:"totalMemoryBytes,omitempty"`
	// A list of application processes.
	Processes []ProcessMetrics `json:"processes,omitempty"`
}

// NewApplicationMetrics instantiates a new ApplicationMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationMetrics() *ApplicationMetrics {
	this := ApplicationMetrics{}
	return &this
}

// NewApplicationMetricsWithDefaults instantiates a new ApplicationMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationMetricsWithDefaults() *ApplicationMetrics {
	this := ApplicationMetrics{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationMetrics) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationMetrics) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationMetrics) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationMetrics) SetName(v string) {
	o.Name = &v
}

// GetTotalCpu returns the TotalCpu field value if set, zero value otherwise.
func (o *ApplicationMetrics) GetTotalCpu() float64 {
	if o == nil || utils.IsNil(o.TotalCpu) {
		var ret float64
		return ret
	}
	return *o.TotalCpu
}

// GetTotalCpuOk returns a tuple with the TotalCpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationMetrics) GetTotalCpuOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.TotalCpu) {
		return nil, false
	}
	return o.TotalCpu, true
}

// HasTotalCpu returns a boolean if a field has been set.
func (o *ApplicationMetrics) HasTotalCpu() bool {
	if o != nil && !utils.IsNil(o.TotalCpu) {
		return true
	}

	return false
}

// SetTotalCpu gets a reference to the given float64 and assigns it to the TotalCpu field.
func (o *ApplicationMetrics) SetTotalCpu(v float64) {
	o.TotalCpu = &v
}

// GetTotalMemoryPercentage returns the TotalMemoryPercentage field value if set, zero value otherwise.
func (o *ApplicationMetrics) GetTotalMemoryPercentage() float64 {
	if o == nil || utils.IsNil(o.TotalMemoryPercentage) {
		var ret float64
		return ret
	}
	return *o.TotalMemoryPercentage
}

// GetTotalMemoryPercentageOk returns a tuple with the TotalMemoryPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationMetrics) GetTotalMemoryPercentageOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.TotalMemoryPercentage) {
		return nil, false
	}
	return o.TotalMemoryPercentage, true
}

// HasTotalMemoryPercentage returns a boolean if a field has been set.
func (o *ApplicationMetrics) HasTotalMemoryPercentage() bool {
	if o != nil && !utils.IsNil(o.TotalMemoryPercentage) {
		return true
	}

	return false
}

// SetTotalMemoryPercentage gets a reference to the given float64 and assigns it to the TotalMemoryPercentage field.
func (o *ApplicationMetrics) SetTotalMemoryPercentage(v float64) {
	o.TotalMemoryPercentage = &v
}

// GetTotalMemoryBytes returns the TotalMemoryBytes field value if set, zero value otherwise.
func (o *ApplicationMetrics) GetTotalMemoryBytes() int64 {
	if o == nil || utils.IsNil(o.TotalMemoryBytes) {
		var ret int64
		return ret
	}
	return *o.TotalMemoryBytes
}

// GetTotalMemoryBytesOk returns a tuple with the TotalMemoryBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationMetrics) GetTotalMemoryBytesOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.TotalMemoryBytes) {
		return nil, false
	}
	return o.TotalMemoryBytes, true
}

// HasTotalMemoryBytes returns a boolean if a field has been set.
func (o *ApplicationMetrics) HasTotalMemoryBytes() bool {
	if o != nil && !utils.IsNil(o.TotalMemoryBytes) {
		return true
	}

	return false
}

// SetTotalMemoryBytes gets a reference to the given int64 and assigns it to the TotalMemoryBytes field.
func (o *ApplicationMetrics) SetTotalMemoryBytes(v int64) {
	o.TotalMemoryBytes = &v
}

// GetProcesses returns the Processes field value if set, zero value otherwise.
func (o *ApplicationMetrics) GetProcesses() []ProcessMetrics {
	if o == nil || utils.IsNil(o.Processes) {
		var ret []ProcessMetrics
		return ret
	}
	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationMetrics) GetProcessesOk() ([]ProcessMetrics, bool) {
	if o == nil || utils.IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// HasProcesses returns a boolean if a field has been set.
func (o *ApplicationMetrics) HasProcesses() bool {
	if o != nil && !utils.IsNil(o.Processes) {
		return true
	}

	return false
}

// SetProcesses gets a reference to the given []ProcessMetrics and assigns it to the Processes field.
func (o *ApplicationMetrics) SetProcesses(v []ProcessMetrics) {
	o.Processes = v
}

func (o ApplicationMetrics) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.TotalCpu) {
		toSerialize["totalCpu"] = o.TotalCpu
	}
	if !utils.IsNil(o.TotalMemoryPercentage) {
		toSerialize["totalMemoryPercentage"] = o.TotalMemoryPercentage
	}
	if !utils.IsNil(o.TotalMemoryBytes) {
		toSerialize["totalMemoryBytes"] = o.TotalMemoryBytes
	}
	if !utils.IsNil(o.Processes) {
		toSerialize["processes"] = o.Processes
	}
	return toSerialize, nil
}

type NullableApplicationMetrics struct {
	value *ApplicationMetrics
	isSet bool
}

func (v NullableApplicationMetrics) Get() *ApplicationMetrics {
	return v.value
}

func (v *NullableApplicationMetrics) Set(val *ApplicationMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationMetrics(val *ApplicationMetrics) *NullableApplicationMetrics {
	return &NullableApplicationMetrics{value: val, isSet: true}
}

func (v NullableApplicationMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


