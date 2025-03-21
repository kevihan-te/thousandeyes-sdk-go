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

// checks if the TargetTraceroute type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TargetTraceroute{}

// TargetTraceroute struct for TargetTraceroute
type TargetTraceroute struct {
	// The target IP address.
	Destination *string `json:"destination,omitempty"`
	// Only present when there is an error
	Error *string `json:"error,omitempty"`
	InfoFlags []string `json:"infoFlags,omitempty"`
	InternalErrors []string `json:"internalErrors,omitempty"`
	Hops []TracerouteHop `json:"hops,omitempty"`
}

// NewTargetTraceroute instantiates a new TargetTraceroute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTargetTraceroute() *TargetTraceroute {
	this := TargetTraceroute{}
	return &this
}

// NewTargetTracerouteWithDefaults instantiates a new TargetTraceroute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTargetTracerouteWithDefaults() *TargetTraceroute {
	this := TargetTraceroute{}
	return &this
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *TargetTraceroute) GetDestination() string {
	if o == nil || utils.IsNil(o.Destination) {
		var ret string
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTraceroute) GetDestinationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *TargetTraceroute) HasDestination() bool {
	if o != nil && !utils.IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given string and assigns it to the Destination field.
func (o *TargetTraceroute) SetDestination(v string) {
	o.Destination = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *TargetTraceroute) GetError() string {
	if o == nil || utils.IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTraceroute) GetErrorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *TargetTraceroute) HasError() bool {
	if o != nil && !utils.IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *TargetTraceroute) SetError(v string) {
	o.Error = &v
}

// GetInfoFlags returns the InfoFlags field value if set, zero value otherwise.
func (o *TargetTraceroute) GetInfoFlags() []string {
	if o == nil || utils.IsNil(o.InfoFlags) {
		var ret []string
		return ret
	}
	return o.InfoFlags
}

// GetInfoFlagsOk returns a tuple with the InfoFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTraceroute) GetInfoFlagsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.InfoFlags) {
		return nil, false
	}
	return o.InfoFlags, true
}

// HasInfoFlags returns a boolean if a field has been set.
func (o *TargetTraceroute) HasInfoFlags() bool {
	if o != nil && !utils.IsNil(o.InfoFlags) {
		return true
	}

	return false
}

// SetInfoFlags gets a reference to the given []string and assigns it to the InfoFlags field.
func (o *TargetTraceroute) SetInfoFlags(v []string) {
	o.InfoFlags = v
}

// GetInternalErrors returns the InternalErrors field value if set, zero value otherwise.
func (o *TargetTraceroute) GetInternalErrors() []string {
	if o == nil || utils.IsNil(o.InternalErrors) {
		var ret []string
		return ret
	}
	return o.InternalErrors
}

// GetInternalErrorsOk returns a tuple with the InternalErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTraceroute) GetInternalErrorsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.InternalErrors) {
		return nil, false
	}
	return o.InternalErrors, true
}

// HasInternalErrors returns a boolean if a field has been set.
func (o *TargetTraceroute) HasInternalErrors() bool {
	if o != nil && !utils.IsNil(o.InternalErrors) {
		return true
	}

	return false
}

// SetInternalErrors gets a reference to the given []string and assigns it to the InternalErrors field.
func (o *TargetTraceroute) SetInternalErrors(v []string) {
	o.InternalErrors = v
}

// GetHops returns the Hops field value if set, zero value otherwise.
func (o *TargetTraceroute) GetHops() []TracerouteHop {
	if o == nil || utils.IsNil(o.Hops) {
		var ret []TracerouteHop
		return ret
	}
	return o.Hops
}

// GetHopsOk returns a tuple with the Hops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TargetTraceroute) GetHopsOk() ([]TracerouteHop, bool) {
	if o == nil || utils.IsNil(o.Hops) {
		return nil, false
	}
	return o.Hops, true
}

// HasHops returns a boolean if a field has been set.
func (o *TargetTraceroute) HasHops() bool {
	if o != nil && !utils.IsNil(o.Hops) {
		return true
	}

	return false
}

// SetHops gets a reference to the given []TracerouteHop and assigns it to the Hops field.
func (o *TargetTraceroute) SetHops(v []TracerouteHop) {
	o.Hops = v
}

func (o TargetTraceroute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TargetTraceroute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !utils.IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !utils.IsNil(o.InfoFlags) {
		toSerialize["infoFlags"] = o.InfoFlags
	}
	if !utils.IsNil(o.InternalErrors) {
		toSerialize["internalErrors"] = o.InternalErrors
	}
	if !utils.IsNil(o.Hops) {
		toSerialize["hops"] = o.Hops
	}
	return toSerialize, nil
}

type NullableTargetTraceroute struct {
	value *TargetTraceroute
	isSet bool
}

func (v NullableTargetTraceroute) Get() *TargetTraceroute {
	return v.value
}

func (v *NullableTargetTraceroute) Set(val *TargetTraceroute) {
	v.value = val
	v.isSet = true
}

func (v NullableTargetTraceroute) IsSet() bool {
	return v.isSet
}

func (v *NullableTargetTraceroute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTargetTraceroute(val *TargetTraceroute) *NullableTargetTraceroute {
	return &NullableTargetTraceroute{value: val, isSet: true}
}

func (v NullableTargetTraceroute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTargetTraceroute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


