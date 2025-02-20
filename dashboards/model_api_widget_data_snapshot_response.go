/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the ApiWidgetDataSnapshotResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiWidgetDataSnapshotResponse{}

// ApiWidgetDataSnapshotResponse struct for ApiWidgetDataSnapshotResponse
type ApiWidgetDataSnapshotResponse struct {
	GroupLabels []ApiReportDataComponentLabelMap `json:"groupLabels,omitempty"`
	// Duration of each bin.
	BinSize *int64 `json:"binSize,omitempty"`
	Data *ApiWidgetsDataV2 `json:"data,omitempty"`
	// (Optional) When passing `window` or `startDate` parameter,  the client will also receive the `startDate` field indicating the UTC start date of the data's time range being retrieved  (ISO date-time format).
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window` or `endDate` parameter,  the client will also receive the `endDate` field indicating the UTC end date of the data's time range being retrieved  (ISO date-time format).
	EndDate *time.Time `json:"endDate,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewApiWidgetDataSnapshotResponse instantiates a new ApiWidgetDataSnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWidgetDataSnapshotResponse() *ApiWidgetDataSnapshotResponse {
	this := ApiWidgetDataSnapshotResponse{}
	return &this
}

// NewApiWidgetDataSnapshotResponseWithDefaults instantiates a new ApiWidgetDataSnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWidgetDataSnapshotResponseWithDefaults() *ApiWidgetDataSnapshotResponse {
	this := ApiWidgetDataSnapshotResponse{}
	return &this
}

// GetGroupLabels returns the GroupLabels field value if set, zero value otherwise.
func (o *ApiWidgetDataSnapshotResponse) GetGroupLabels() []ApiReportDataComponentLabelMap {
	if o == nil || utils.IsNil(o.GroupLabels) {
		var ret []ApiReportDataComponentLabelMap
		return ret
	}
	return o.GroupLabels
}

// GetGroupLabelsOk returns a tuple with the GroupLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWidgetDataSnapshotResponse) GetGroupLabelsOk() ([]ApiReportDataComponentLabelMap, bool) {
	if o == nil || utils.IsNil(o.GroupLabels) {
		return nil, false
	}
	return o.GroupLabels, true
}

// HasGroupLabels returns a boolean if a field has been set.
func (o *ApiWidgetDataSnapshotResponse) HasGroupLabels() bool {
	if o != nil && !utils.IsNil(o.GroupLabels) {
		return true
	}

	return false
}

// SetGroupLabels gets a reference to the given []ApiReportDataComponentLabelMap and assigns it to the GroupLabels field.
func (o *ApiWidgetDataSnapshotResponse) SetGroupLabels(v []ApiReportDataComponentLabelMap) {
	o.GroupLabels = v
}

// GetBinSize returns the BinSize field value if set, zero value otherwise.
func (o *ApiWidgetDataSnapshotResponse) GetBinSize() int64 {
	if o == nil || utils.IsNil(o.BinSize) {
		var ret int64
		return ret
	}
	return *o.BinSize
}

// GetBinSizeOk returns a tuple with the BinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWidgetDataSnapshotResponse) GetBinSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.BinSize) {
		return nil, false
	}
	return o.BinSize, true
}

// HasBinSize returns a boolean if a field has been set.
func (o *ApiWidgetDataSnapshotResponse) HasBinSize() bool {
	if o != nil && !utils.IsNil(o.BinSize) {
		return true
	}

	return false
}

// SetBinSize gets a reference to the given int64 and assigns it to the BinSize field.
func (o *ApiWidgetDataSnapshotResponse) SetBinSize(v int64) {
	o.BinSize = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiWidgetDataSnapshotResponse) GetData() ApiWidgetsDataV2 {
	if o == nil || utils.IsNil(o.Data) {
		var ret ApiWidgetsDataV2
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWidgetDataSnapshotResponse) GetDataOk() (*ApiWidgetsDataV2, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiWidgetDataSnapshotResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ApiWidgetsDataV2 and assigns it to the Data field.
func (o *ApiWidgetDataSnapshotResponse) SetData(v ApiWidgetsDataV2) {
	o.Data = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ApiWidgetDataSnapshotResponse) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWidgetDataSnapshotResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ApiWidgetDataSnapshotResponse) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ApiWidgetDataSnapshotResponse) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ApiWidgetDataSnapshotResponse) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWidgetDataSnapshotResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ApiWidgetDataSnapshotResponse) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ApiWidgetDataSnapshotResponse) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiWidgetDataSnapshotResponse) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWidgetDataSnapshotResponse) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiWidgetDataSnapshotResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *ApiWidgetDataSnapshotResponse) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o ApiWidgetDataSnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiWidgetDataSnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.GroupLabels) {
		toSerialize["groupLabels"] = o.GroupLabels
	}
	if !utils.IsNil(o.BinSize) {
		toSerialize["binSize"] = o.BinSize
	}
	if !utils.IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !utils.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !utils.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableApiWidgetDataSnapshotResponse struct {
	value *ApiWidgetDataSnapshotResponse
	isSet bool
}

func (v NullableApiWidgetDataSnapshotResponse) Get() *ApiWidgetDataSnapshotResponse {
	return v.value
}

func (v *NullableApiWidgetDataSnapshotResponse) Set(val *ApiWidgetDataSnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWidgetDataSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWidgetDataSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWidgetDataSnapshotResponse(val *ApiWidgetDataSnapshotResponse) *NullableApiWidgetDataSnapshotResponse {
	return &NullableApiWidgetDataSnapshotResponse{value: val, isSet: true}
}

func (v NullableApiWidgetDataSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWidgetDataSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


