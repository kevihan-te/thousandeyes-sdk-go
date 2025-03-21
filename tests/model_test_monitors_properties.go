/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the TestMonitorsProperties type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestMonitorsProperties{}

// TestMonitorsProperties struct for TestMonitorsProperties
type TestMonitorsProperties struct {
	// Set to `true` to enable bgp measurements.
	BgpMeasurements *bool `json:"bgpMeasurements,omitempty"`
	// Indicate if all available public BGP monitors should be used, when ommited defaults to `bgpMeasurements` value.
	UsePublicBgp *bool `json:"usePublicBgp,omitempty"`
	// Contains list of enabled BGP monitors.
	Monitors []Monitor `json:"monitors,omitempty"`
}

// NewTestMonitorsProperties instantiates a new TestMonitorsProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestMonitorsProperties() *TestMonitorsProperties {
	this := TestMonitorsProperties{}
	var bgpMeasurements bool = true
	this.BgpMeasurements = &bgpMeasurements
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	return &this
}

// NewTestMonitorsPropertiesWithDefaults instantiates a new TestMonitorsProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestMonitorsPropertiesWithDefaults() *TestMonitorsProperties {
	this := TestMonitorsProperties{}
	var bgpMeasurements bool = true
	this.BgpMeasurements = &bgpMeasurements
	var usePublicBgp bool = true
	this.UsePublicBgp = &usePublicBgp
	return &this
}

// GetBgpMeasurements returns the BgpMeasurements field value if set, zero value otherwise.
func (o *TestMonitorsProperties) GetBgpMeasurements() bool {
	if o == nil || utils.IsNil(o.BgpMeasurements) {
		var ret bool
		return ret
	}
	return *o.BgpMeasurements
}

// GetBgpMeasurementsOk returns a tuple with the BgpMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestMonitorsProperties) GetBgpMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BgpMeasurements) {
		return nil, false
	}
	return o.BgpMeasurements, true
}

// HasBgpMeasurements returns a boolean if a field has been set.
func (o *TestMonitorsProperties) HasBgpMeasurements() bool {
	if o != nil && !utils.IsNil(o.BgpMeasurements) {
		return true
	}

	return false
}

// SetBgpMeasurements gets a reference to the given bool and assigns it to the BgpMeasurements field.
func (o *TestMonitorsProperties) SetBgpMeasurements(v bool) {
	o.BgpMeasurements = &v
}

// GetUsePublicBgp returns the UsePublicBgp field value if set, zero value otherwise.
func (o *TestMonitorsProperties) GetUsePublicBgp() bool {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		var ret bool
		return ret
	}
	return *o.UsePublicBgp
}

// GetUsePublicBgpOk returns a tuple with the UsePublicBgp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestMonitorsProperties) GetUsePublicBgpOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.UsePublicBgp) {
		return nil, false
	}
	return o.UsePublicBgp, true
}

// HasUsePublicBgp returns a boolean if a field has been set.
func (o *TestMonitorsProperties) HasUsePublicBgp() bool {
	if o != nil && !utils.IsNil(o.UsePublicBgp) {
		return true
	}

	return false
}

// SetUsePublicBgp gets a reference to the given bool and assigns it to the UsePublicBgp field.
func (o *TestMonitorsProperties) SetUsePublicBgp(v bool) {
	o.UsePublicBgp = &v
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *TestMonitorsProperties) GetMonitors() []Monitor {
	if o == nil || utils.IsNil(o.Monitors) {
		var ret []Monitor
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestMonitorsProperties) GetMonitorsOk() ([]Monitor, bool) {
	if o == nil || utils.IsNil(o.Monitors) {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *TestMonitorsProperties) HasMonitors() bool {
	if o != nil && !utils.IsNil(o.Monitors) {
		return true
	}

	return false
}

// SetMonitors gets a reference to the given []Monitor and assigns it to the Monitors field.
func (o *TestMonitorsProperties) SetMonitors(v []Monitor) {
	o.Monitors = v
}

func (o TestMonitorsProperties) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestMonitorsProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.BgpMeasurements) {
		toSerialize["bgpMeasurements"] = o.BgpMeasurements
	}
	if !utils.IsNil(o.UsePublicBgp) {
		toSerialize["usePublicBgp"] = o.UsePublicBgp
	}
	if !utils.IsNil(o.Monitors) {
		toSerialize["monitors"] = o.Monitors
	}
	return toSerialize, nil
}

type NullableTestMonitorsProperties struct {
	value *TestMonitorsProperties
	isSet bool
}

func (v NullableTestMonitorsProperties) Get() *TestMonitorsProperties {
	return v.value
}

func (v *NullableTestMonitorsProperties) Set(val *TestMonitorsProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableTestMonitorsProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableTestMonitorsProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestMonitorsProperties(val *TestMonitorsProperties) *NullableTestMonitorsProperties {
	return &NullableTestMonitorsProperties{value: val, isSet: true}
}

func (v NullableTestMonitorsProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestMonitorsProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


