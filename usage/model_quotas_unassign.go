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

// checks if the QuotasUnassign type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &QuotasUnassign{}

// QuotasUnassign struct for QuotasUnassign
type QuotasUnassign struct {
	Organizations []string `json:"organizations,omitempty"`
}

// NewQuotasUnassign instantiates a new QuotasUnassign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuotasUnassign() *QuotasUnassign {
	this := QuotasUnassign{}
	return &this
}

// NewQuotasUnassignWithDefaults instantiates a new QuotasUnassign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotasUnassignWithDefaults() *QuotasUnassign {
	this := QuotasUnassign{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *QuotasUnassign) GetOrganizations() []string {
	if o == nil || utils.IsNil(o.Organizations) {
		var ret []string
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QuotasUnassign) GetOrganizationsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Organizations) {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *QuotasUnassign) HasOrganizations() bool {
	if o != nil && !utils.IsNil(o.Organizations) {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []string and assigns it to the Organizations field.
func (o *QuotasUnassign) SetOrganizations(v []string) {
	o.Organizations = v
}

func (o QuotasUnassign) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QuotasUnassign) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Organizations) {
		toSerialize["organizations"] = o.Organizations
	}
	return toSerialize, nil
}

type NullableQuotasUnassign struct {
	value *QuotasUnassign
	isSet bool
}

func (v NullableQuotasUnassign) Get() *QuotasUnassign {
	return v.value
}

func (v *NullableQuotasUnassign) Set(val *QuotasUnassign) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotasUnassign) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotasUnassign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotasUnassign(val *QuotasUnassign) *NullableQuotasUnassign {
	return &NullableQuotasUnassign{value: val, isSet: true}
}

func (v NullableQuotasUnassign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotasUnassign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


