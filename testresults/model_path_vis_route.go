/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the PathVisRoute type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PathVisRoute{}

// PathVisRoute struct for PathVisRoute
type PathVisRoute struct {
	// Unique ID of path trace
	PathId *string `json:"pathId,omitempty"`
	// Array of hop objects indicating each step in the traceroute
	Hops []PathVisHop `json:"hops,omitempty"`
}

// NewPathVisRoute instantiates a new PathVisRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathVisRoute() *PathVisRoute {
	this := PathVisRoute{}
	return &this
}

// NewPathVisRouteWithDefaults instantiates a new PathVisRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathVisRouteWithDefaults() *PathVisRoute {
	this := PathVisRoute{}
	return &this
}

// GetPathId returns the PathId field value if set, zero value otherwise.
func (o *PathVisRoute) GetPathId() string {
	if o == nil || utils.IsNil(o.PathId) {
		var ret string
		return ret
	}
	return *o.PathId
}

// GetPathIdOk returns a tuple with the PathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisRoute) GetPathIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PathId) {
		return nil, false
	}
	return o.PathId, true
}

// HasPathId returns a boolean if a field has been set.
func (o *PathVisRoute) HasPathId() bool {
	if o != nil && !utils.IsNil(o.PathId) {
		return true
	}

	return false
}

// SetPathId gets a reference to the given string and assigns it to the PathId field.
func (o *PathVisRoute) SetPathId(v string) {
	o.PathId = &v
}

// GetHops returns the Hops field value if set, zero value otherwise.
func (o *PathVisRoute) GetHops() []PathVisHop {
	if o == nil || utils.IsNil(o.Hops) {
		var ret []PathVisHop
		return ret
	}
	return o.Hops
}

// GetHopsOk returns a tuple with the Hops field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathVisRoute) GetHopsOk() ([]PathVisHop, bool) {
	if o == nil || utils.IsNil(o.Hops) {
		return nil, false
	}
	return o.Hops, true
}

// HasHops returns a boolean if a field has been set.
func (o *PathVisRoute) HasHops() bool {
	if o != nil && !utils.IsNil(o.Hops) {
		return true
	}

	return false
}

// SetHops gets a reference to the given []PathVisHop and assigns it to the Hops field.
func (o *PathVisRoute) SetHops(v []PathVisHop) {
	o.Hops = v
}

func (o PathVisRoute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PathVisRoute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PathId) {
		toSerialize["pathId"] = o.PathId
	}
	if !utils.IsNil(o.Hops) {
		toSerialize["hops"] = o.Hops
	}
	return toSerialize, nil
}

type NullablePathVisRoute struct {
	value *PathVisRoute
	isSet bool
}

func (v NullablePathVisRoute) Get() *PathVisRoute {
	return v.value
}

func (v *NullablePathVisRoute) Set(val *PathVisRoute) {
	v.value = val
	v.isSet = true
}

func (v NullablePathVisRoute) IsSet() bool {
	return v.isSet
}

func (v *NullablePathVisRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathVisRoute(val *PathVisRoute) *NullablePathVisRoute {
	return &NullablePathVisRoute{value: val, isSet: true}
}

func (v NullablePathVisRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathVisRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


