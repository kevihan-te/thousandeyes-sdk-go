/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the ApiContextFiltersResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiContextFiltersResponse{}

// ApiContextFiltersResponse All global dashboard filters saved.
type ApiContextFiltersResponse struct {
	DashboardFilters []ApiContextFilterResponse `json:"dashboardFilters,omitempty"`
}

// NewApiContextFiltersResponse instantiates a new ApiContextFiltersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiContextFiltersResponse() *ApiContextFiltersResponse {
	this := ApiContextFiltersResponse{}
	return &this
}

// NewApiContextFiltersResponseWithDefaults instantiates a new ApiContextFiltersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiContextFiltersResponseWithDefaults() *ApiContextFiltersResponse {
	this := ApiContextFiltersResponse{}
	return &this
}

// GetDashboardFilters returns the DashboardFilters field value if set, zero value otherwise.
func (o *ApiContextFiltersResponse) GetDashboardFilters() []ApiContextFilterResponse {
	if o == nil || utils.IsNil(o.DashboardFilters) {
		var ret []ApiContextFilterResponse
		return ret
	}
	return o.DashboardFilters
}

// GetDashboardFiltersOk returns a tuple with the DashboardFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContextFiltersResponse) GetDashboardFiltersOk() ([]ApiContextFilterResponse, bool) {
	if o == nil || utils.IsNil(o.DashboardFilters) {
		return nil, false
	}
	return o.DashboardFilters, true
}

// HasDashboardFilters returns a boolean if a field has been set.
func (o *ApiContextFiltersResponse) HasDashboardFilters() bool {
	if o != nil && !utils.IsNil(o.DashboardFilters) {
		return true
	}

	return false
}

// SetDashboardFilters gets a reference to the given []ApiContextFilterResponse and assigns it to the DashboardFilters field.
func (o *ApiContextFiltersResponse) SetDashboardFilters(v []ApiContextFilterResponse) {
	o.DashboardFilters = v
}

func (o ApiContextFiltersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiContextFiltersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DashboardFilters) {
		toSerialize["dashboardFilters"] = o.DashboardFilters
	}
	return toSerialize, nil
}

type NullableApiContextFiltersResponse struct {
	value *ApiContextFiltersResponse
	isSet bool
}

func (v NullableApiContextFiltersResponse) Get() *ApiContextFiltersResponse {
	return v.value
}

func (v *NullableApiContextFiltersResponse) Set(val *ApiContextFiltersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiContextFiltersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiContextFiltersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiContextFiltersResponse(val *ApiContextFiltersResponse) *NullableApiContextFiltersResponse {
	return &NullableApiContextFiltersResponse{value: val, isSet: true}
}

func (v NullableApiContextFiltersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiContextFiltersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


