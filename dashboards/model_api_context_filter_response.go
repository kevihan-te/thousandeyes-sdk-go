/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
	"fmt"
)

// checks if the ApiContextFilterResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiContextFilterResponse{}

// ApiContextFilterResponse Response containing dashboard filter settings and context details.
type ApiContextFilterResponse struct {
	// List of filters to be applied to a dashboard.
	Context []ApiDataSourceFilters `json:"context"`
	// Account group ID of the filter.
	Aid string `json:"aid"`
	// Unique ID of the dashboard filter.
	Id string `json:"id"`
	// The name of the dashboard filter, this must be unique.
	Name string `json:"name"`
	// An optional description of the filter.
	Description *string `json:"description,omitempty"`
	CreatedBy *ApiDashboardFilterUserDetails `json:"createdBy,omitempty"`
	// Timestamp when the filter was last modified.
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Timestamp when the filter was created.
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	ModifiedBy *ApiDashboardFilterUserDetails `json:"modifiedBy,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

type _ApiContextFilterResponse ApiContextFilterResponse

// NewApiContextFilterResponse instantiates a new ApiContextFilterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiContextFilterResponse(context []ApiDataSourceFilters, aid string, id string, name string) *ApiContextFilterResponse {
	this := ApiContextFilterResponse{}
	this.Context = context
	this.Aid = aid
	this.Id = id
	this.Name = name
	return &this
}

// NewApiContextFilterResponseWithDefaults instantiates a new ApiContextFilterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiContextFilterResponseWithDefaults() *ApiContextFilterResponse {
	this := ApiContextFilterResponse{}
	return &this
}

// GetContext returns the Context field value
func (o *ApiContextFilterResponse) GetContext() []ApiDataSourceFilters {
	if o == nil {
		var ret []ApiDataSourceFilters
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetContextOk() ([]ApiDataSourceFilters, bool) {
	if o == nil {
		return nil, false
	}
	return o.Context, true
}

// SetContext sets field value
func (o *ApiContextFilterResponse) SetContext(v []ApiDataSourceFilters) {
	o.Context = v
}

// GetAid returns the Aid field value
func (o *ApiContextFilterResponse) GetAid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aid
}

// GetAidOk returns a tuple with the Aid field value
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetAidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aid, true
}

// SetAid sets field value
func (o *ApiContextFilterResponse) SetAid(v string) {
	o.Aid = v
}

// GetId returns the Id field value
func (o *ApiContextFilterResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiContextFilterResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApiContextFilterResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiContextFilterResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiContextFilterResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiContextFilterResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiContextFilterResponse) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ApiContextFilterResponse) GetCreatedBy() ApiDashboardFilterUserDetails {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret ApiDashboardFilterUserDetails
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetCreatedByOk() (*ApiDashboardFilterUserDetails, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ApiContextFilterResponse) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given ApiDashboardFilterUserDetails and assigns it to the CreatedBy field.
func (o *ApiContextFilterResponse) SetCreatedBy(v ApiDashboardFilterUserDetails) {
	o.CreatedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *ApiContextFilterResponse) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *ApiContextFilterResponse) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *ApiContextFilterResponse) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ApiContextFilterResponse) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiContextFilterResponse) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *ApiContextFilterResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *ApiContextFilterResponse) GetModifiedBy() ApiDashboardFilterUserDetails {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret ApiDashboardFilterUserDetails
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetModifiedByOk() (*ApiDashboardFilterUserDetails, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *ApiContextFilterResponse) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given ApiDashboardFilterUserDetails and assigns it to the ModifiedBy field.
func (o *ApiContextFilterResponse) SetModifiedBy(v ApiDashboardFilterUserDetails) {
	o.ModifiedBy = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiContextFilterResponse) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiContextFilterResponse) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiContextFilterResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *ApiContextFilterResponse) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o ApiContextFilterResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiContextFilterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["aid"] = o.Aid
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !utils.IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !utils.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !utils.IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

func (o *ApiContextFilterResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"context",
		"aid",
		"id",
		"name",
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

	varApiContextFilterResponse := _ApiContextFilterResponse{}

    err = json.Unmarshal(data, &varApiContextFilterResponse)

	if err != nil {
		return err
	}

	*o = ApiContextFilterResponse(varApiContextFilterResponse)

	return err
}

type NullableApiContextFilterResponse struct {
	value *ApiContextFilterResponse
	isSet bool
}

func (v NullableApiContextFilterResponse) Get() *ApiContextFilterResponse {
	return v.value
}

func (v *NullableApiContextFilterResponse) Set(val *ApiContextFilterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiContextFilterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiContextFilterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiContextFilterResponse(val *ApiContextFilterResponse) *NullableApiContextFilterResponse {
	return &NullableApiContextFilterResponse{value: val, isSet: true}
}

func (v NullableApiContextFilterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiContextFilterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


