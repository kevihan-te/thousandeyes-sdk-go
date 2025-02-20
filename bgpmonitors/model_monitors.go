/*
BGP Monitors API

 Retrieve information about BGP monitors available to your ThousandEyes account. ThousandEyes ingests BGP routing data from dozens of global BGP collectors and automatically integrates that visibility as a configurable layer under service, network, and path visualization layers.  When you specify a service URL in a test, layered BGP tests automatically track reachability and path changes for any relevant prefix. When you use an IP address as the target for a test, the ThousandEyes platform monitors the relevant internet-routed prefix. You can also create specific BGP monitoring for a prefix, and can alert on hijacks and leaks.  For more information about monitors, see [Inside-Out BGP Visibility](https://docs.thousandeyes.com/product-documentation/internet-and-wan-monitoring/tests/bgp-tests/inside-out-bgp-visibility). 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bgpmonitors

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the Monitors type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Monitors{}

// Monitors struct for Monitors
type Monitors struct {
	Monitors []Monitor `json:"monitors,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewMonitors instantiates a new Monitors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitors() *Monitors {
	this := Monitors{}
	return &this
}

// NewMonitorsWithDefaults instantiates a new Monitors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorsWithDefaults() *Monitors {
	this := Monitors{}
	return &this
}

// GetMonitors returns the Monitors field value if set, zero value otherwise.
func (o *Monitors) GetMonitors() []Monitor {
	if o == nil || utils.IsNil(o.Monitors) {
		var ret []Monitor
		return ret
	}
	return o.Monitors
}

// GetMonitorsOk returns a tuple with the Monitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitors) GetMonitorsOk() ([]Monitor, bool) {
	if o == nil || utils.IsNil(o.Monitors) {
		return nil, false
	}
	return o.Monitors, true
}

// HasMonitors returns a boolean if a field has been set.
func (o *Monitors) HasMonitors() bool {
	if o != nil && !utils.IsNil(o.Monitors) {
		return true
	}

	return false
}

// SetMonitors gets a reference to the given []Monitor and assigns it to the Monitors field.
func (o *Monitors) SetMonitors(v []Monitor) {
	o.Monitors = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Monitors) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Monitors) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Monitors) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *Monitors) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o Monitors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Monitors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Monitors) {
		toSerialize["monitors"] = o.Monitors
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableMonitors struct {
	value *Monitors
	isSet bool
}

func (v NullableMonitors) Get() *Monitors {
	return v.value
}

func (v *NullableMonitors) Set(val *Monitors) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitors) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitors(val *Monitors) *NullableMonitors {
	return &NullableMonitors{value: val, isSet: true}
}

func (v NullableMonitors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


