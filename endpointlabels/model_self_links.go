/*
Endpoint Agent Labels API

Manage labels applied to endpoint agents using this API. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointlabels

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the SelfLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SelfLinks{}

// SelfLinks A links object containing the self link.
type SelfLinks struct {
	Self *Link `json:"self,omitempty"`
}

// NewSelfLinks instantiates a new SelfLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSelfLinks() *SelfLinks {
	this := SelfLinks{}
	return &this
}

// NewSelfLinksWithDefaults instantiates a new SelfLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSelfLinksWithDefaults() *SelfLinks {
	this := SelfLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SelfLinks) GetSelf() Link {
	if o == nil || utils.IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SelfLinks) GetSelfOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SelfLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *SelfLinks) SetSelf(v Link) {
	o.Self = &v
}

func (o SelfLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SelfLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableSelfLinks struct {
	value *SelfLinks
	isSet bool
}

func (v NullableSelfLinks) Get() *SelfLinks {
	return v.value
}

func (v *NullableSelfLinks) Set(val *SelfLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfLinks(val *SelfLinks) *NullableSelfLinks {
	return &NullableSelfLinks{value: val, isSet: true}
}

func (v NullableSelfLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


