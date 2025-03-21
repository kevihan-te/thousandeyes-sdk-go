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

// checks if the LocalNetworkResults type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LocalNetworkResults{}

// LocalNetworkResults struct for LocalNetworkResults
type LocalNetworkResults struct {
	LocalNetworks []LocalNetworkResult `json:"localNetworks,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewLocalNetworkResults instantiates a new LocalNetworkResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalNetworkResults() *LocalNetworkResults {
	this := LocalNetworkResults{}
	return &this
}

// NewLocalNetworkResultsWithDefaults instantiates a new LocalNetworkResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalNetworkResultsWithDefaults() *LocalNetworkResults {
	this := LocalNetworkResults{}
	return &this
}

// GetLocalNetworks returns the LocalNetworks field value if set, zero value otherwise.
func (o *LocalNetworkResults) GetLocalNetworks() []LocalNetworkResult {
	if o == nil || utils.IsNil(o.LocalNetworks) {
		var ret []LocalNetworkResult
		return ret
	}
	return o.LocalNetworks
}

// GetLocalNetworksOk returns a tuple with the LocalNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalNetworkResults) GetLocalNetworksOk() ([]LocalNetworkResult, bool) {
	if o == nil || utils.IsNil(o.LocalNetworks) {
		return nil, false
	}
	return o.LocalNetworks, true
}

// HasLocalNetworks returns a boolean if a field has been set.
func (o *LocalNetworkResults) HasLocalNetworks() bool {
	if o != nil && !utils.IsNil(o.LocalNetworks) {
		return true
	}

	return false
}

// SetLocalNetworks gets a reference to the given []LocalNetworkResult and assigns it to the LocalNetworks field.
func (o *LocalNetworkResults) SetLocalNetworks(v []LocalNetworkResult) {
	o.LocalNetworks = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LocalNetworkResults) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalNetworkResults) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LocalNetworkResults) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *LocalNetworkResults) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o LocalNetworkResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalNetworkResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.LocalNetworks) {
		toSerialize["localNetworks"] = o.LocalNetworks
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableLocalNetworkResults struct {
	value *LocalNetworkResults
	isSet bool
}

func (v NullableLocalNetworkResults) Get() *LocalNetworkResults {
	return v.value
}

func (v *NullableLocalNetworkResults) Set(val *LocalNetworkResults) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalNetworkResults) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalNetworkResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalNetworkResults(val *LocalNetworkResults) *NullableLocalNetworkResults {
	return &NullableLocalNetworkResults{value: val, isSet: true}
}

func (v NullableLocalNetworkResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalNetworkResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


