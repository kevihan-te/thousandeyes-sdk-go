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

// checks if the AsnDetails type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &AsnDetails{}

// AsnDetails struct for AsnDetails
type AsnDetails struct {
	// Name of the provider.
	AsName *string `json:"asName,omitempty"`
	// Unique number assigned to an organization (also referred to as service provider).
	AsNumber *int32 `json:"asNumber,omitempty"`
}

// NewAsnDetails instantiates a new AsnDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsnDetails() *AsnDetails {
	this := AsnDetails{}
	return &this
}

// NewAsnDetailsWithDefaults instantiates a new AsnDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsnDetailsWithDefaults() *AsnDetails {
	this := AsnDetails{}
	return &this
}

// GetAsName returns the AsName field value if set, zero value otherwise.
func (o *AsnDetails) GetAsName() string {
	if o == nil || utils.IsNil(o.AsName) {
		var ret string
		return ret
	}
	return *o.AsName
}

// GetAsNameOk returns a tuple with the AsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsnDetails) GetAsNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AsName) {
		return nil, false
	}
	return o.AsName, true
}

// HasAsName returns a boolean if a field has been set.
func (o *AsnDetails) HasAsName() bool {
	if o != nil && !utils.IsNil(o.AsName) {
		return true
	}

	return false
}

// SetAsName gets a reference to the given string and assigns it to the AsName field.
func (o *AsnDetails) SetAsName(v string) {
	o.AsName = &v
}

// GetAsNumber returns the AsNumber field value if set, zero value otherwise.
func (o *AsnDetails) GetAsNumber() int32 {
	if o == nil || utils.IsNil(o.AsNumber) {
		var ret int32
		return ret
	}
	return *o.AsNumber
}

// GetAsNumberOk returns a tuple with the AsNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsnDetails) GetAsNumberOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.AsNumber) {
		return nil, false
	}
	return o.AsNumber, true
}

// HasAsNumber returns a boolean if a field has been set.
func (o *AsnDetails) HasAsNumber() bool {
	if o != nil && !utils.IsNil(o.AsNumber) {
		return true
	}

	return false
}

// SetAsNumber gets a reference to the given int32 and assigns it to the AsNumber field.
func (o *AsnDetails) SetAsNumber(v int32) {
	o.AsNumber = &v
}

func (o AsnDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AsnDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AsName) {
		toSerialize["asName"] = o.AsName
	}
	if !utils.IsNil(o.AsNumber) {
		toSerialize["asNumber"] = o.AsNumber
	}
	return toSerialize, nil
}

type NullableAsnDetails struct {
	value *AsnDetails
	isSet bool
}

func (v NullableAsnDetails) Get() *AsnDetails {
	return v.value
}

func (v *NullableAsnDetails) Set(val *AsnDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableAsnDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableAsnDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsnDetails(val *AsnDetails) *NullableAsnDetails {
	return &NullableAsnDetails{value: val, isSet: true}
}

func (v NullableAsnDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsnDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


