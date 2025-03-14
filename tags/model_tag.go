/*
Tags API

The ThousandEyes Tags API provides a tagging system with key/value pairs. It allows you to tag assets within the ThousandEyes platform (such as agents, tests, or alert rules) with meaningful metadata. For example: `branch:sfo`, `branch:nyc`, and `team:netops`.  This feature provides:  * Support for automation. * Powerful and flexible reports/dashboards. * Support for third-party integrations.  Things to note with the ThousandEyes Tags API:  * Tags are backwards-compatible with existing labels. * Tags are separated by Tests (CEA), Agents (CEA), Endpoint Agents, Scheduled Endpoint Tests, and Reports. A single tag can only apply to one type of target object, so each tag must specify the target type of object via a `type` field. * Tags are defined in a single table so that they can be represented using a single model - `Tag`. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the Tag type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Tag{}

// Tag struct for Tag
type Tag struct {
	Assignments []Assignment `json:"assignments,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
	// The account group ID
	Aid *int32 `json:"aid,omitempty"`
	// Tag color
	Color *string `json:"color,omitempty"`
	// Tag creation date
	CreateDate *string `json:"createDate,omitempty"`
	Icon utils.NullableString `json:"icon,omitempty"`
	// The tag's description.
	Description utils.NullableString `json:"description,omitempty"`
	// The tag ID
	Id *string `json:"id,omitempty"`
	// The tags's key
	Key *string `json:"key,omitempty"`
	LegacyId utils.NullableFloat32 `json:"legacyId,omitempty"`
	ObjectType *ObjectType `json:"objectType,omitempty"`
	// The tag's value
	Value *string `json:"value,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag() *Tag {
	this := Tag{}
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *Tag) GetAssignments() []Assignment {
	if o == nil || utils.IsNil(o.Assignments) {
		var ret []Assignment
		return ret
	}
	return o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetAssignmentsOk() ([]Assignment, bool) {
	if o == nil || utils.IsNil(o.Assignments) {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *Tag) HasAssignments() bool {
	if o != nil && !utils.IsNil(o.Assignments) {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []Assignment and assigns it to the Assignments field.
func (o *Tag) SetAssignments(v []Assignment) {
	o.Assignments = v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *Tag) GetAccessType() AccessType {
	if o == nil || utils.IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || utils.IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *Tag) HasAccessType() bool {
	if o != nil && !utils.IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *Tag) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *Tag) GetAid() int32 {
	if o == nil || utils.IsNil(o.Aid) {
		var ret int32
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetAidOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *Tag) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given int32 and assigns it to the Aid field.
func (o *Tag) SetAid(v int32) {
	o.Aid = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *Tag) GetColor() string {
	if o == nil || utils.IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetColorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *Tag) HasColor() bool {
	if o != nil && !utils.IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *Tag) SetColor(v string) {
	o.Color = &v
}

// GetCreateDate returns the CreateDate field value if set, zero value otherwise.
func (o *Tag) GetCreateDate() string {
	if o == nil || utils.IsNil(o.CreateDate) {
		var ret string
		return ret
	}
	return *o.CreateDate
}

// GetCreateDateOk returns a tuple with the CreateDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetCreateDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreateDate) {
		return nil, false
	}
	return o.CreateDate, true
}

// HasCreateDate returns a boolean if a field has been set.
func (o *Tag) HasCreateDate() bool {
	if o != nil && !utils.IsNil(o.CreateDate) {
		return true
	}

	return false
}

// SetCreateDate gets a reference to the given string and assigns it to the CreateDate field.
func (o *Tag) SetCreateDate(v string) {
	o.CreateDate = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetIcon() string {
	if o == nil || utils.IsNil(o.Icon.Get()) {
		var ret string
		return ret
	}
	return *o.Icon.Get()
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetIconOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Icon.Get(), o.Icon.IsSet()
}

// HasIcon returns a boolean if a field has been set.
func (o *Tag) HasIcon() bool {
	if o != nil && o.Icon.IsSet() {
		return true
	}

	return false
}

// SetIcon gets a reference to the given NullableString and assigns it to the Icon field.
func (o *Tag) SetIcon(v string) {
	o.Icon.Set(&v)
}
// SetIconNil sets the value for Icon to be an explicit nil
func (o *Tag) SetIconNil() {
	o.Icon.Set(nil)
}

// UnsetIcon ensures that no value is present for Icon, not even an explicit nil
func (o *Tag) UnsetIcon() {
	o.Icon.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetDescription() string {
	if o == nil || utils.IsNil(o.Description.Get()) {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Tag) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Tag) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Tag) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Tag) UnsetDescription() {
	o.Description.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tag) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Tag) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Tag) SetId(v string) {
	o.Id = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *Tag) GetKey() string {
	if o == nil || utils.IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *Tag) HasKey() bool {
	if o != nil && !utils.IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *Tag) SetKey(v string) {
	o.Key = &v
}

// GetLegacyId returns the LegacyId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Tag) GetLegacyId() float32 {
	if o == nil || utils.IsNil(o.LegacyId.Get()) {
		var ret float32
		return ret
	}
	return *o.LegacyId.Get()
}

// GetLegacyIdOk returns a tuple with the LegacyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetLegacyIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LegacyId.Get(), o.LegacyId.IsSet()
}

// HasLegacyId returns a boolean if a field has been set.
func (o *Tag) HasLegacyId() bool {
	if o != nil && o.LegacyId.IsSet() {
		return true
	}

	return false
}

// SetLegacyId gets a reference to the given NullableFloat32 and assigns it to the LegacyId field.
func (o *Tag) SetLegacyId(v float32) {
	o.LegacyId.Set(&v)
}
// SetLegacyIdNil sets the value for LegacyId to be an explicit nil
func (o *Tag) SetLegacyIdNil() {
	o.LegacyId.Set(nil)
}

// UnsetLegacyId ensures that no value is present for LegacyId, not even an explicit nil
func (o *Tag) UnsetLegacyId() {
	o.LegacyId.Unset()
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *Tag) GetObjectType() ObjectType {
	if o == nil || utils.IsNil(o.ObjectType) {
		var ret ObjectType
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetObjectTypeOk() (*ObjectType, bool) {
	if o == nil || utils.IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// HasObjectType returns a boolean if a field has been set.
func (o *Tag) HasObjectType() bool {
	if o != nil && !utils.IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given ObjectType and assigns it to the ObjectType field.
func (o *Tag) SetObjectType(v ObjectType) {
	o.ObjectType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Tag) GetValue() string {
	if o == nil || utils.IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Tag) HasValue() bool {
	if o != nil && !utils.IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Tag) SetValue(v string) {
	o.Value = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Tag) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tag) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Tag) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *Tag) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Assignments) {
		toSerialize["assignments"] = o.Assignments
	}
	if !utils.IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !utils.IsNil(o.CreateDate) {
		toSerialize["createDate"] = o.CreateDate
	}
	if o.Icon.IsSet() {
		toSerialize["icon"] = o.Icon.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if o.LegacyId.IsSet() {
		toSerialize["legacyId"] = o.LegacyId.Get()
	}
	if !utils.IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !utils.IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


