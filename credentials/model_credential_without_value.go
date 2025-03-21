/*
Credentials API

Manage credentials for transaction tests using the Credentials API.  The following permissions are required to access Credentials API operations:  * `Settings Tests Read` for read operations. * `Settings Tests Update` for write operations. * `View sensitive data in web transaction scripts` to view the encrypted value property of credentials. * `Settings Tests Create Transaction (Tx) Tests` to create credentials.  For more information about credentials, see [Working With Secure Credentials](https://docs.thousandeyes.com/product-documentation/browser-synthetics/transaction-tests/getting-started/working-with-secure-credentials). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the CredentialWithoutValue type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CredentialWithoutValue{}

// CredentialWithoutValue struct for CredentialWithoutValue
type CredentialWithoutValue struct {
	// Unique ID of the credential.
	Id *string `json:"id,omitempty"`
	// The name of the credential.
	Name *string `json:"name,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewCredentialWithoutValue instantiates a new CredentialWithoutValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialWithoutValue() *CredentialWithoutValue {
	this := CredentialWithoutValue{}
	return &this
}

// NewCredentialWithoutValueWithDefaults instantiates a new CredentialWithoutValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialWithoutValueWithDefaults() *CredentialWithoutValue {
	this := CredentialWithoutValue{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CredentialWithoutValue) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialWithoutValue) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CredentialWithoutValue) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CredentialWithoutValue) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CredentialWithoutValue) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialWithoutValue) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CredentialWithoutValue) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CredentialWithoutValue) SetName(v string) {
	o.Name = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *CredentialWithoutValue) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialWithoutValue) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *CredentialWithoutValue) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *CredentialWithoutValue) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o CredentialWithoutValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialWithoutValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableCredentialWithoutValue struct {
	value *CredentialWithoutValue
	isSet bool
}

func (v NullableCredentialWithoutValue) Get() *CredentialWithoutValue {
	return v.value
}

func (v *NullableCredentialWithoutValue) Set(val *CredentialWithoutValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialWithoutValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialWithoutValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialWithoutValue(val *CredentialWithoutValue) *NullableCredentialWithoutValue {
	return &NullableCredentialWithoutValue{value: val, isSet: true}
}

func (v NullableCredentialWithoutValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialWithoutValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


