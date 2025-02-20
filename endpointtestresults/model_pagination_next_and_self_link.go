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

// checks if the PaginationNextAndSelfLink type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PaginationNextAndSelfLink{}

// PaginationNextAndSelfLink A links object containing a related link for forward pagination.
type PaginationNextAndSelfLink struct {
	Next *Link `json:"next,omitempty"`
	Self *Link `json:"self,omitempty"`
}

// NewPaginationNextAndSelfLink instantiates a new PaginationNextAndSelfLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationNextAndSelfLink() *PaginationNextAndSelfLink {
	this := PaginationNextAndSelfLink{}
	return &this
}

// NewPaginationNextAndSelfLinkWithDefaults instantiates a new PaginationNextAndSelfLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationNextAndSelfLinkWithDefaults() *PaginationNextAndSelfLink {
	this := PaginationNextAndSelfLink{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *PaginationNextAndSelfLink) GetNext() Link {
	if o == nil || utils.IsNil(o.Next) {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationNextAndSelfLink) GetNextOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *PaginationNextAndSelfLink) HasNext() bool {
	if o != nil && !utils.IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *PaginationNextAndSelfLink) SetNext(v Link) {
	o.Next = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PaginationNextAndSelfLink) GetSelf() Link {
	if o == nil || utils.IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationNextAndSelfLink) GetSelfOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PaginationNextAndSelfLink) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *PaginationNextAndSelfLink) SetSelf(v Link) {
	o.Self = &v
}

func (o PaginationNextAndSelfLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationNextAndSelfLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullablePaginationNextAndSelfLink struct {
	value *PaginationNextAndSelfLink
	isSet bool
}

func (v NullablePaginationNextAndSelfLink) Get() *PaginationNextAndSelfLink {
	return v.value
}

func (v *NullablePaginationNextAndSelfLink) Set(val *PaginationNextAndSelfLink) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationNextAndSelfLink) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationNextAndSelfLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationNextAndSelfLink(val *PaginationNextAndSelfLink) *NullablePaginationNextAndSelfLink {
	return &NullablePaginationNextAndSelfLink{value: val, isSet: true}
}

func (v NullablePaginationNextAndSelfLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationNextAndSelfLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


