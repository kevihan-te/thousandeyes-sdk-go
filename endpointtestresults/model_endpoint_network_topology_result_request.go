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

// checks if the EndpointNetworkTopologyResultRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointNetworkTopologyResultRequest{}

// EndpointNetworkTopologyResultRequest struct for EndpointNetworkTopologyResultRequest
type EndpointNetworkTopologyResultRequest struct {
	SearchFilters *EndpointNetworkTopologyResultRequestFilter `json:"searchFilters,omitempty"`
}

// NewEndpointNetworkTopologyResultRequest instantiates a new EndpointNetworkTopologyResultRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointNetworkTopologyResultRequest() *EndpointNetworkTopologyResultRequest {
	this := EndpointNetworkTopologyResultRequest{}
	return &this
}

// NewEndpointNetworkTopologyResultRequestWithDefaults instantiates a new EndpointNetworkTopologyResultRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointNetworkTopologyResultRequestWithDefaults() *EndpointNetworkTopologyResultRequest {
	this := EndpointNetworkTopologyResultRequest{}
	return &this
}

// GetSearchFilters returns the SearchFilters field value if set, zero value otherwise.
func (o *EndpointNetworkTopologyResultRequest) GetSearchFilters() EndpointNetworkTopologyResultRequestFilter {
	if o == nil || utils.IsNil(o.SearchFilters) {
		var ret EndpointNetworkTopologyResultRequestFilter
		return ret
	}
	return *o.SearchFilters
}

// GetSearchFiltersOk returns a tuple with the SearchFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointNetworkTopologyResultRequest) GetSearchFiltersOk() (*EndpointNetworkTopologyResultRequestFilter, bool) {
	if o == nil || utils.IsNil(o.SearchFilters) {
		return nil, false
	}
	return o.SearchFilters, true
}

// HasSearchFilters returns a boolean if a field has been set.
func (o *EndpointNetworkTopologyResultRequest) HasSearchFilters() bool {
	if o != nil && !utils.IsNil(o.SearchFilters) {
		return true
	}

	return false
}

// SetSearchFilters gets a reference to the given EndpointNetworkTopologyResultRequestFilter and assigns it to the SearchFilters field.
func (o *EndpointNetworkTopologyResultRequest) SetSearchFilters(v EndpointNetworkTopologyResultRequestFilter) {
	o.SearchFilters = &v
}

func (o EndpointNetworkTopologyResultRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointNetworkTopologyResultRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.SearchFilters) {
		toSerialize["searchFilters"] = o.SearchFilters
	}
	return toSerialize, nil
}

type NullableEndpointNetworkTopologyResultRequest struct {
	value *EndpointNetworkTopologyResultRequest
	isSet bool
}

func (v NullableEndpointNetworkTopologyResultRequest) Get() *EndpointNetworkTopologyResultRequest {
	return v.value
}

func (v *NullableEndpointNetworkTopologyResultRequest) Set(val *EndpointNetworkTopologyResultRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointNetworkTopologyResultRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointNetworkTopologyResultRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointNetworkTopologyResultRequest(val *EndpointNetworkTopologyResultRequest) *NullableEndpointNetworkTopologyResultRequest {
	return &NullableEndpointNetworkTopologyResultRequest{value: val, isSet: true}
}

func (v NullableEndpointNetworkTopologyResultRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointNetworkTopologyResultRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


