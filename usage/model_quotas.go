/*
Usage API

 These usage endpoints define the following operations:  * **Usage**: Retrieve usage data for the specified time period (default is one month).          * Users must have the `View organization usage` permission to access this endpoint.     * This operation offers visibility across all account groups within the organization.     * Users with `View organization usage` permission in multiple organizations should query the operation with the `aid` query string parameter (see optional parameters) for each organization.  * **Quotas**: Obtain organization and account usage quotas. Additionally, users with the appropriate permissions can create, update, or delete these quotas.          * Users must have the necessary permissions to perform quota-related actions.  Refer to the Usage API operations for detailed usage instructions and optional parameters. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package usage

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the Quotas type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Quotas{}

// Quotas struct for Quotas
type Quotas struct {
	Quotas []Quota `json:"quotas,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewQuotas instantiates a new Quotas object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotas() *Quotas {
	this := Quotas{}
	return &this
}

// NewQuotasWithDefaults instantiates a new Quotas object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotasWithDefaults() *Quotas {
	this := Quotas{}
	return &this
}

// GetQuotas returns the Quotas field value if set, zero value otherwise.
func (o *Quotas) GetQuotas() []Quota {
	if o == nil || utils.IsNil(o.Quotas) {
		var ret []Quota
		return ret
	}
	return o.Quotas
}

// GetQuotasOk returns a tuple with the Quotas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quotas) GetQuotasOk() ([]Quota, bool) {
	if o == nil || utils.IsNil(o.Quotas) {
		return nil, false
	}
	return o.Quotas, true
}

// HasQuotas returns a boolean if a field has been set.
func (o *Quotas) HasQuotas() bool {
	if o != nil && !utils.IsNil(o.Quotas) {
		return true
	}

	return false
}

// SetQuotas gets a reference to the given []Quota and assigns it to the Quotas field.
func (o *Quotas) SetQuotas(v []Quota) {
	o.Quotas = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Quotas) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quotas) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Quotas) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *Quotas) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o Quotas) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Quotas) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Quotas) {
		toSerialize["quotas"] = o.Quotas
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableQuotas struct {
	value *Quotas
	isSet bool
}

func (v NullableQuotas) Get() *Quotas {
	return v.value
}

func (v *NullableQuotas) Set(val *Quotas) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotas) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotas) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotas(val *Quotas) *NullableQuotas {
	return &NullableQuotas{value: val, isSet: true}
}

func (v NullableQuotas) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotas) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


