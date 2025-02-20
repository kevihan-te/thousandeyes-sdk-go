/*
Test Snapshots API

Creates a new test snapshot in ThousandEyes.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package snapshots

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the SnapshotLinks type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SnapshotLinks{}

// SnapshotLinks struct for SnapshotLinks
type SnapshotLinks struct {
	AppLink *Link `json:"appLink,omitempty"`
	Self *Link `json:"self,omitempty"`
}

// NewSnapshotLinks instantiates a new SnapshotLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotLinks() *SnapshotLinks {
	this := SnapshotLinks{}
	return &this
}

// NewSnapshotLinksWithDefaults instantiates a new SnapshotLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotLinksWithDefaults() *SnapshotLinks {
	this := SnapshotLinks{}
	return &this
}

// GetAppLink returns the AppLink field value if set, zero value otherwise.
func (o *SnapshotLinks) GetAppLink() Link {
	if o == nil || utils.IsNil(o.AppLink) {
		var ret Link
		return ret
	}
	return *o.AppLink
}

// GetAppLinkOk returns a tuple with the AppLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotLinks) GetAppLinkOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.AppLink) {
		return nil, false
	}
	return o.AppLink, true
}

// HasAppLink returns a boolean if a field has been set.
func (o *SnapshotLinks) HasAppLink() bool {
	if o != nil && !utils.IsNil(o.AppLink) {
		return true
	}

	return false
}

// SetAppLink gets a reference to the given Link and assigns it to the AppLink field.
func (o *SnapshotLinks) SetAppLink(v Link) {
	o.AppLink = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *SnapshotLinks) GetSelf() Link {
	if o == nil || utils.IsNil(o.Self) {
		var ret Link
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotLinks) GetSelfOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *SnapshotLinks) HasSelf() bool {
	if o != nil && !utils.IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given Link and assigns it to the Self field.
func (o *SnapshotLinks) SetSelf(v Link) {
	o.Self = &v
}

func (o SnapshotLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AppLink) {
		toSerialize["appLink"] = o.AppLink
	}
	if !utils.IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableSnapshotLinks struct {
	value *SnapshotLinks
	isSet bool
}

func (v NullableSnapshotLinks) Get() *SnapshotLinks {
	return v.value
}

func (v *NullableSnapshotLinks) Set(val *SnapshotLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotLinks(val *SnapshotLinks) *NullableSnapshotLinks {
	return &NullableSnapshotLinks{value: val, isSet: true}
}

func (v NullableSnapshotLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


