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
	"fmt"
)

// checks if the TestSelfLink type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TestSelfLink{}

// TestSelfLink struct for TestSelfLink
type TestSelfLink struct {
	// Its value is either a URI [RFC3986] or a URI template [RFC6570].
	Href string `json:"href"`
	// Should be true when the link object's \"href\" property is a URI template.
	Templated *bool `json:"templated,omitempty"`
	// Used as a hint to indicate the media type expected when dereferencing the target resource.
	Type *string `json:"type,omitempty"`
	// Its presence indicates that the link is to be deprecated at a future date. Its value is a URL that should provide further information about the deprecation.
	Deprecation *string `json:"deprecation,omitempty"`
	// Its value may be used as a secondary key for selecting link objects that share the same relation type.
	Name *string `json:"name,omitempty"`
	// A URI that hints about the profile of the target resource.
	Profile *string `json:"profile,omitempty"`
	// Intended for labelling the link with a human-readable identifier
	Title *string `json:"title,omitempty"`
	// Indicates the language of the target resource
	Hreflang *string `json:"hreflang,omitempty"`
}

type _TestSelfLink TestSelfLink

// NewTestSelfLink instantiates a new TestSelfLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestSelfLink(href string) *TestSelfLink {
	this := TestSelfLink{}
	this.Href = href
	return &this
}

// NewTestSelfLinkWithDefaults instantiates a new TestSelfLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestSelfLinkWithDefaults() *TestSelfLink {
	this := TestSelfLink{}
	return &this
}

// GetHref returns the Href field value
func (o *TestSelfLink) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *TestSelfLink) SetHref(v string) {
	o.Href = v
}

// GetTemplated returns the Templated field value if set, zero value otherwise.
func (o *TestSelfLink) GetTemplated() bool {
	if o == nil || utils.IsNil(o.Templated) {
		var ret bool
		return ret
	}
	return *o.Templated
}

// GetTemplatedOk returns a tuple with the Templated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetTemplatedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Templated) {
		return nil, false
	}
	return o.Templated, true
}

// HasTemplated returns a boolean if a field has been set.
func (o *TestSelfLink) HasTemplated() bool {
	if o != nil && !utils.IsNil(o.Templated) {
		return true
	}

	return false
}

// SetTemplated gets a reference to the given bool and assigns it to the Templated field.
func (o *TestSelfLink) SetTemplated(v bool) {
	o.Templated = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TestSelfLink) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TestSelfLink) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *TestSelfLink) SetType(v string) {
	o.Type = &v
}

// GetDeprecation returns the Deprecation field value if set, zero value otherwise.
func (o *TestSelfLink) GetDeprecation() string {
	if o == nil || utils.IsNil(o.Deprecation) {
		var ret string
		return ret
	}
	return *o.Deprecation
}

// GetDeprecationOk returns a tuple with the Deprecation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetDeprecationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Deprecation) {
		return nil, false
	}
	return o.Deprecation, true
}

// HasDeprecation returns a boolean if a field has been set.
func (o *TestSelfLink) HasDeprecation() bool {
	if o != nil && !utils.IsNil(o.Deprecation) {
		return true
	}

	return false
}

// SetDeprecation gets a reference to the given string and assigns it to the Deprecation field.
func (o *TestSelfLink) SetDeprecation(v string) {
	o.Deprecation = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TestSelfLink) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TestSelfLink) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TestSelfLink) SetName(v string) {
	o.Name = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *TestSelfLink) GetProfile() string {
	if o == nil || utils.IsNil(o.Profile) {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetProfileOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *TestSelfLink) HasProfile() bool {
	if o != nil && !utils.IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *TestSelfLink) SetProfile(v string) {
	o.Profile = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TestSelfLink) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TestSelfLink) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TestSelfLink) SetTitle(v string) {
	o.Title = &v
}

// GetHreflang returns the Hreflang field value if set, zero value otherwise.
func (o *TestSelfLink) GetHreflang() string {
	if o == nil || utils.IsNil(o.Hreflang) {
		var ret string
		return ret
	}
	return *o.Hreflang
}

// GetHreflangOk returns a tuple with the Hreflang field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestSelfLink) GetHreflangOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Hreflang) {
		return nil, false
	}
	return o.Hreflang, true
}

// HasHreflang returns a boolean if a field has been set.
func (o *TestSelfLink) HasHreflang() bool {
	if o != nil && !utils.IsNil(o.Hreflang) {
		return true
	}

	return false
}

// SetHreflang gets a reference to the given string and assigns it to the Hreflang field.
func (o *TestSelfLink) SetHreflang(v string) {
	o.Hreflang = &v
}

func (o TestSelfLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TestSelfLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	if !utils.IsNil(o.Templated) {
		toSerialize["templated"] = o.Templated
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Deprecation) {
		toSerialize["deprecation"] = o.Deprecation
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !utils.IsNil(o.Hreflang) {
		toSerialize["hreflang"] = o.Hreflang
	}
	return toSerialize, nil
}

func (o *TestSelfLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTestSelfLink := _TestSelfLink{}

    err = json.Unmarshal(data, &varTestSelfLink)

	if err != nil {
		return err
	}

	*o = TestSelfLink(varTestSelfLink)

	return err
}

type NullableTestSelfLink struct {
	value *TestSelfLink
	isSet bool
}

func (v NullableTestSelfLink) Get() *TestSelfLink {
	return v.value
}

func (v *NullableTestSelfLink) Set(val *TestSelfLink) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSelfLink) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSelfLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSelfLink(val *TestSelfLink) *NullableTestSelfLink {
	return &NullableTestSelfLink{value: val, isSet: true}
}

func (v NullableTestSelfLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSelfLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


