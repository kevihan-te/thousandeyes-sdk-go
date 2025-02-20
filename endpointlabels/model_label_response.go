/*
Endpoint Agent Labels API

Manage labels applied to endpoint agents using this API. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointlabels

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the LabelResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LabelResponse{}

// LabelResponse struct for LabelResponse
type LabelResponse struct {
	// Label identifier.
	Id *string `json:"id,omitempty"`
	// The label name.
	Name *string `json:"name,omitempty"`
	// UI color
	Color *string `json:"color,omitempty"`
	MatchType *MatchType `json:"matchType,omitempty"`
	// The filters combined using the matchType to determine the label's match.
	Filters []Filter `json:"filters,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewLabelResponse instantiates a new LabelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabelResponse() *LabelResponse {
	this := LabelResponse{}
	return &this
}

// NewLabelResponseWithDefaults instantiates a new LabelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabelResponseWithDefaults() *LabelResponse {
	this := LabelResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LabelResponse) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResponse) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LabelResponse) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LabelResponse) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LabelResponse) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResponse) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LabelResponse) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LabelResponse) SetName(v string) {
	o.Name = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *LabelResponse) GetColor() string {
	if o == nil || utils.IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResponse) GetColorOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *LabelResponse) HasColor() bool {
	if o != nil && !utils.IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *LabelResponse) SetColor(v string) {
	o.Color = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *LabelResponse) GetMatchType() MatchType {
	if o == nil || utils.IsNil(o.MatchType) {
		var ret MatchType
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResponse) GetMatchTypeOk() (*MatchType, bool) {
	if o == nil || utils.IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *LabelResponse) HasMatchType() bool {
	if o != nil && !utils.IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given MatchType and assigns it to the MatchType field.
func (o *LabelResponse) SetMatchType(v MatchType) {
	o.MatchType = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *LabelResponse) GetFilters() []Filter {
	if o == nil || utils.IsNil(o.Filters) {
		var ret []Filter
		return ret
	}
	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResponse) GetFiltersOk() ([]Filter, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *LabelResponse) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given []Filter and assigns it to the Filters field.
func (o *LabelResponse) SetFilters(v []Filter) {
	o.Filters = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *LabelResponse) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LabelResponse) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *LabelResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *LabelResponse) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o LabelResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LabelResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !utils.IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !utils.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableLabelResponse struct {
	value *LabelResponse
	isSet bool
}

func (v NullableLabelResponse) Get() *LabelResponse {
	return v.value
}

func (v *NullableLabelResponse) Set(val *LabelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLabelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLabelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabelResponse(val *LabelResponse) *NullableLabelResponse {
	return &NullableLabelResponse{value: val, isSet: true}
}

func (v NullableLabelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


