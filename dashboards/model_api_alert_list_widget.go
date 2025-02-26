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
	"fmt"
)

// checks if the ApiAlertListWidget type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiAlertListWidget{}

// ApiAlertListWidget struct for ApiAlertListWidget
type ApiAlertListWidget struct {
	// Identifier of the widget.
	Id *string `json:"id,omitempty"`
	// Title of the widget
	Title *string `json:"title,omitempty"`
	VisualMode *VisualMode `json:"visualMode,omitempty"`
	// When `isEmbedded` is set to `true`, an `embedUrl` is provided.
	EmbedUrl *string `json:"embedUrl,omitempty"`
	// Set to `true` if widget is marked as embedded; otherwise, set to `false`.
	IsEmbedded *bool `json:"isEmbedded,omitempty"`
	MetricGroup *MetricGroup `json:"metricGroup,omitempty"`
	Direction *DashboardMetricDirection `json:"direction,omitempty"`
	Metric *DashboardMetric `json:"metric,omitempty"`
	// (Optional) Specifies the filters applied to the widget. When present, the `filters` property displays. Each filter object has two properties: `filterProperty` and `filterValue`. The `filterProperty` can be values like `AGENT`, `ENDPOINT_MACHINE_ID`, `TEST`, `MONITOR`, etc.  The `filterValue` represents an identifier array of the selected property.
	Filters *map[string][]interface{} `json:"filters,omitempty"`
	Measure *ApiWidgetMeasure `json:"measure,omitempty"`
	FixedTimespan *ApiDuration `json:"fixedTimespan,omitempty"`
	// Deprecated
	ApiLink *string `json:"apiLink,omitempty"`
	// Excludes alert suppression window data if set to `true`.
	ShouldExcludeAlertSuppressionWindows *bool `json:"shouldExcludeAlertSuppressionWindows,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
	// Alert List widget type.
	Type string `json:"type"`
	// List of alert types configured in the widget, an empty list means all alert types.
	AlertTypes []LegacyAlertListAlertType `json:"alertTypes,omitempty"`
	// Limit the number of alerts displayed in the widget.
	LimitTo *int32 `json:"limitTo,omitempty"`
	ActiveWithin *ActiveWithin `json:"activeWithin,omitempty"`
	DataSource *AlertListDatasource `json:"dataSource,omitempty"`
}

type _ApiAlertListWidget ApiAlertListWidget

// NewApiAlertListWidget instantiates a new ApiAlertListWidget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAlertListWidget(type_ string) *ApiAlertListWidget {
	this := ApiAlertListWidget{}
	var visualMode VisualMode = VISUALMODE_FULL
	this.VisualMode = &visualMode
	this.Type = type_
	return &this
}

// NewApiAlertListWidgetWithDefaults instantiates a new ApiAlertListWidget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAlertListWidgetWithDefaults() *ApiAlertListWidget {
	this := ApiAlertListWidget{}
	var visualMode VisualMode = VISUALMODE_FULL
	this.VisualMode = &visualMode
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAlertListWidget) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiAlertListWidget) SetTitle(v string) {
	o.Title = &v
}

// GetVisualMode returns the VisualMode field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetVisualMode() VisualMode {
	if o == nil || utils.IsNil(o.VisualMode) {
		var ret VisualMode
		return ret
	}
	return *o.VisualMode
}

// GetVisualModeOk returns a tuple with the VisualMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetVisualModeOk() (*VisualMode, bool) {
	if o == nil || utils.IsNil(o.VisualMode) {
		return nil, false
	}
	return o.VisualMode, true
}

// HasVisualMode returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasVisualMode() bool {
	if o != nil && !utils.IsNil(o.VisualMode) {
		return true
	}

	return false
}

// SetVisualMode gets a reference to the given VisualMode and assigns it to the VisualMode field.
func (o *ApiAlertListWidget) SetVisualMode(v VisualMode) {
	o.VisualMode = &v
}

// GetEmbedUrl returns the EmbedUrl field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetEmbedUrl() string {
	if o == nil || utils.IsNil(o.EmbedUrl) {
		var ret string
		return ret
	}
	return *o.EmbedUrl
}

// GetEmbedUrlOk returns a tuple with the EmbedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetEmbedUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EmbedUrl) {
		return nil, false
	}
	return o.EmbedUrl, true
}

// HasEmbedUrl returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasEmbedUrl() bool {
	if o != nil && !utils.IsNil(o.EmbedUrl) {
		return true
	}

	return false
}

// SetEmbedUrl gets a reference to the given string and assigns it to the EmbedUrl field.
func (o *ApiAlertListWidget) SetEmbedUrl(v string) {
	o.EmbedUrl = &v
}

// GetIsEmbedded returns the IsEmbedded field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetIsEmbedded() bool {
	if o == nil || utils.IsNil(o.IsEmbedded) {
		var ret bool
		return ret
	}
	return *o.IsEmbedded
}

// GetIsEmbeddedOk returns a tuple with the IsEmbedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetIsEmbeddedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsEmbedded) {
		return nil, false
	}
	return o.IsEmbedded, true
}

// HasIsEmbedded returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasIsEmbedded() bool {
	if o != nil && !utils.IsNil(o.IsEmbedded) {
		return true
	}

	return false
}

// SetIsEmbedded gets a reference to the given bool and assigns it to the IsEmbedded field.
func (o *ApiAlertListWidget) SetIsEmbedded(v bool) {
	o.IsEmbedded = &v
}

// GetMetricGroup returns the MetricGroup field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetMetricGroup() MetricGroup {
	if o == nil || utils.IsNil(o.MetricGroup) {
		var ret MetricGroup
		return ret
	}
	return *o.MetricGroup
}

// GetMetricGroupOk returns a tuple with the MetricGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetMetricGroupOk() (*MetricGroup, bool) {
	if o == nil || utils.IsNil(o.MetricGroup) {
		return nil, false
	}
	return o.MetricGroup, true
}

// HasMetricGroup returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasMetricGroup() bool {
	if o != nil && !utils.IsNil(o.MetricGroup) {
		return true
	}

	return false
}

// SetMetricGroup gets a reference to the given MetricGroup and assigns it to the MetricGroup field.
func (o *ApiAlertListWidget) SetMetricGroup(v MetricGroup) {
	o.MetricGroup = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetDirection() DashboardMetricDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret DashboardMetricDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetDirectionOk() (*DashboardMetricDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given DashboardMetricDirection and assigns it to the Direction field.
func (o *ApiAlertListWidget) SetDirection(v DashboardMetricDirection) {
	o.Direction = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetMetric() DashboardMetric {
	if o == nil || utils.IsNil(o.Metric) {
		var ret DashboardMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetMetricOk() (*DashboardMetric, bool) {
	if o == nil || utils.IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasMetric() bool {
	if o != nil && !utils.IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given DashboardMetric and assigns it to the Metric field.
func (o *ApiAlertListWidget) SetMetric(v DashboardMetric) {
	o.Metric = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetFilters() map[string][]interface{} {
	if o == nil || utils.IsNil(o.Filters) {
		var ret map[string][]interface{}
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetFiltersOk() (*map[string][]interface{}, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string][]interface{} and assigns it to the Filters field.
func (o *ApiAlertListWidget) SetFilters(v map[string][]interface{}) {
	o.Filters = &v
}

// GetMeasure returns the Measure field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetMeasure() ApiWidgetMeasure {
	if o == nil || utils.IsNil(o.Measure) {
		var ret ApiWidgetMeasure
		return ret
	}
	return *o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetMeasureOk() (*ApiWidgetMeasure, bool) {
	if o == nil || utils.IsNil(o.Measure) {
		return nil, false
	}
	return o.Measure, true
}

// HasMeasure returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasMeasure() bool {
	if o != nil && !utils.IsNil(o.Measure) {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given ApiWidgetMeasure and assigns it to the Measure field.
func (o *ApiAlertListWidget) SetMeasure(v ApiWidgetMeasure) {
	o.Measure = &v
}

// GetFixedTimespan returns the FixedTimespan field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetFixedTimespan() ApiDuration {
	if o == nil || utils.IsNil(o.FixedTimespan) {
		var ret ApiDuration
		return ret
	}
	return *o.FixedTimespan
}

// GetFixedTimespanOk returns a tuple with the FixedTimespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetFixedTimespanOk() (*ApiDuration, bool) {
	if o == nil || utils.IsNil(o.FixedTimespan) {
		return nil, false
	}
	return o.FixedTimespan, true
}

// HasFixedTimespan returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasFixedTimespan() bool {
	if o != nil && !utils.IsNil(o.FixedTimespan) {
		return true
	}

	return false
}

// SetFixedTimespan gets a reference to the given ApiDuration and assigns it to the FixedTimespan field.
func (o *ApiAlertListWidget) SetFixedTimespan(v ApiDuration) {
	o.FixedTimespan = &v
}

// GetApiLink returns the ApiLink field value if set, zero value otherwise.
// Deprecated
func (o *ApiAlertListWidget) GetApiLink() string {
	if o == nil || utils.IsNil(o.ApiLink) {
		var ret string
		return ret
	}
	return *o.ApiLink
}

// GetApiLinkOk returns a tuple with the ApiLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApiAlertListWidget) GetApiLinkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ApiLink) {
		return nil, false
	}
	return o.ApiLink, true
}

// HasApiLink returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasApiLink() bool {
	if o != nil && !utils.IsNil(o.ApiLink) {
		return true
	}

	return false
}

// SetApiLink gets a reference to the given string and assigns it to the ApiLink field.
// Deprecated
func (o *ApiAlertListWidget) SetApiLink(v string) {
	o.ApiLink = &v
}

// GetShouldExcludeAlertSuppressionWindows returns the ShouldExcludeAlertSuppressionWindows field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetShouldExcludeAlertSuppressionWindows() bool {
	if o == nil || utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeAlertSuppressionWindows
}

// GetShouldExcludeAlertSuppressionWindowsOk returns a tuple with the ShouldExcludeAlertSuppressionWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetShouldExcludeAlertSuppressionWindowsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		return nil, false
	}
	return o.ShouldExcludeAlertSuppressionWindows, true
}

// HasShouldExcludeAlertSuppressionWindows returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasShouldExcludeAlertSuppressionWindows() bool {
	if o != nil && !utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		return true
	}

	return false
}

// SetShouldExcludeAlertSuppressionWindows gets a reference to the given bool and assigns it to the ShouldExcludeAlertSuppressionWindows field.
func (o *ApiAlertListWidget) SetShouldExcludeAlertSuppressionWindows(v bool) {
	o.ShouldExcludeAlertSuppressionWindows = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *ApiAlertListWidget) SetLinks(v SelfLinks) {
	o.Links = &v
}

// GetType returns the Type field value
func (o *ApiAlertListWidget) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiAlertListWidget) SetType(v string) {
	o.Type = v
}

// GetAlertTypes returns the AlertTypes field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetAlertTypes() []LegacyAlertListAlertType {
	if o == nil || utils.IsNil(o.AlertTypes) {
		var ret []LegacyAlertListAlertType
		return ret
	}
	return o.AlertTypes
}

// GetAlertTypesOk returns a tuple with the AlertTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetAlertTypesOk() ([]LegacyAlertListAlertType, bool) {
	if o == nil || utils.IsNil(o.AlertTypes) {
		return nil, false
	}
	return o.AlertTypes, true
}

// HasAlertTypes returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasAlertTypes() bool {
	if o != nil && !utils.IsNil(o.AlertTypes) {
		return true
	}

	return false
}

// SetAlertTypes gets a reference to the given []LegacyAlertListAlertType and assigns it to the AlertTypes field.
func (o *ApiAlertListWidget) SetAlertTypes(v []LegacyAlertListAlertType) {
	o.AlertTypes = v
}

// GetLimitTo returns the LimitTo field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetLimitTo() int32 {
	if o == nil || utils.IsNil(o.LimitTo) {
		var ret int32
		return ret
	}
	return *o.LimitTo
}

// GetLimitToOk returns a tuple with the LimitTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetLimitToOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.LimitTo) {
		return nil, false
	}
	return o.LimitTo, true
}

// HasLimitTo returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasLimitTo() bool {
	if o != nil && !utils.IsNil(o.LimitTo) {
		return true
	}

	return false
}

// SetLimitTo gets a reference to the given int32 and assigns it to the LimitTo field.
func (o *ApiAlertListWidget) SetLimitTo(v int32) {
	o.LimitTo = &v
}

// GetActiveWithin returns the ActiveWithin field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetActiveWithin() ActiveWithin {
	if o == nil || utils.IsNil(o.ActiveWithin) {
		var ret ActiveWithin
		return ret
	}
	return *o.ActiveWithin
}

// GetActiveWithinOk returns a tuple with the ActiveWithin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetActiveWithinOk() (*ActiveWithin, bool) {
	if o == nil || utils.IsNil(o.ActiveWithin) {
		return nil, false
	}
	return o.ActiveWithin, true
}

// HasActiveWithin returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasActiveWithin() bool {
	if o != nil && !utils.IsNil(o.ActiveWithin) {
		return true
	}

	return false
}

// SetActiveWithin gets a reference to the given ActiveWithin and assigns it to the ActiveWithin field.
func (o *ApiAlertListWidget) SetActiveWithin(v ActiveWithin) {
	o.ActiveWithin = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ApiAlertListWidget) GetDataSource() AlertListDatasource {
	if o == nil || utils.IsNil(o.DataSource) {
		var ret AlertListDatasource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListWidget) GetDataSourceOk() (*AlertListDatasource, bool) {
	if o == nil || utils.IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ApiAlertListWidget) HasDataSource() bool {
	if o != nil && !utils.IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given AlertListDatasource and assigns it to the DataSource field.
func (o *ApiAlertListWidget) SetDataSource(v AlertListDatasource) {
	o.DataSource = &v
}

func (o ApiAlertListWidget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAlertListWidget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !utils.IsNil(o.VisualMode) {
		toSerialize["visualMode"] = o.VisualMode
	}
	if !utils.IsNil(o.EmbedUrl) {
		toSerialize["embedUrl"] = o.EmbedUrl
	}
	if !utils.IsNil(o.IsEmbedded) {
		toSerialize["isEmbedded"] = o.IsEmbedded
	}
	if !utils.IsNil(o.MetricGroup) {
		toSerialize["metricGroup"] = o.MetricGroup
	}
	if !utils.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !utils.IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !utils.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !utils.IsNil(o.Measure) {
		toSerialize["measure"] = o.Measure
	}
	if !utils.IsNil(o.FixedTimespan) {
		toSerialize["fixedTimespan"] = o.FixedTimespan
	}
	if !utils.IsNil(o.ApiLink) {
		toSerialize["apiLink"] = o.ApiLink
	}
	if !utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		toSerialize["shouldExcludeAlertSuppressionWindows"] = o.ShouldExcludeAlertSuppressionWindows
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.AlertTypes) {
		toSerialize["alertTypes"] = o.AlertTypes
	}
	if !utils.IsNil(o.LimitTo) {
		toSerialize["limitTo"] = o.LimitTo
	}
	if !utils.IsNil(o.ActiveWithin) {
		toSerialize["activeWithin"] = o.ActiveWithin
	}
	if !utils.IsNil(o.DataSource) {
		toSerialize["dataSource"] = o.DataSource
	}
	return toSerialize, nil
}

func (o *ApiAlertListWidget) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varApiAlertListWidget := _ApiAlertListWidget{}

    err = json.Unmarshal(data, &varApiAlertListWidget)

	if err != nil {
		return err
	}

	*o = ApiAlertListWidget(varApiAlertListWidget)

	return err
}

type NullableApiAlertListWidget struct {
	value *ApiAlertListWidget
	isSet bool
}

func (v NullableApiAlertListWidget) Get() *ApiAlertListWidget {
	return v.value
}

func (v *NullableApiAlertListWidget) Set(val *ApiAlertListWidget) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAlertListWidget) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAlertListWidget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAlertListWidget(val *ApiAlertListWidget) *NullableApiAlertListWidget {
	return &NullableApiAlertListWidget{value: val, isSet: true}
}

func (v NullableApiAlertListWidget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAlertListWidget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


