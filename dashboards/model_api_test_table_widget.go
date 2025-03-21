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

// checks if the ApiTestTableWidget type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiTestTableWidget{}

// ApiTestTableWidget Displays tests and statuses with options to sort and filter. It can be set to show only certain tests, like those with specific labels or failing tests.
type ApiTestTableWidget struct {
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
	// Test Table widget type,
	Type string `json:"type"`
	Filter *ApiWidgetFilterApiTestTableFilterKey `json:"filter,omitempty"`
	Exclude *ApiWidgetFilterApiTestTableFilterKey `json:"exclude,omitempty"`
	DataSource *TestTableDatasource `json:"dataSource,omitempty"`
}

type _ApiTestTableWidget ApiTestTableWidget

// NewApiTestTableWidget instantiates a new ApiTestTableWidget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTestTableWidget(type_ string) *ApiTestTableWidget {
	this := ApiTestTableWidget{}
	var visualMode VisualMode = VISUALMODE_FULL
	this.VisualMode = &visualMode
	this.Type = type_
	return &this
}

// NewApiTestTableWidgetWithDefaults instantiates a new ApiTestTableWidget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTestTableWidgetWithDefaults() *ApiTestTableWidget {
	this := ApiTestTableWidget{}
	var visualMode VisualMode = VISUALMODE_FULL
	this.VisualMode = &visualMode
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiTestTableWidget) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiTestTableWidget) SetTitle(v string) {
	o.Title = &v
}

// GetVisualMode returns the VisualMode field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetVisualMode() VisualMode {
	if o == nil || utils.IsNil(o.VisualMode) {
		var ret VisualMode
		return ret
	}
	return *o.VisualMode
}

// GetVisualModeOk returns a tuple with the VisualMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetVisualModeOk() (*VisualMode, bool) {
	if o == nil || utils.IsNil(o.VisualMode) {
		return nil, false
	}
	return o.VisualMode, true
}

// HasVisualMode returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasVisualMode() bool {
	if o != nil && !utils.IsNil(o.VisualMode) {
		return true
	}

	return false
}

// SetVisualMode gets a reference to the given VisualMode and assigns it to the VisualMode field.
func (o *ApiTestTableWidget) SetVisualMode(v VisualMode) {
	o.VisualMode = &v
}

// GetEmbedUrl returns the EmbedUrl field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetEmbedUrl() string {
	if o == nil || utils.IsNil(o.EmbedUrl) {
		var ret string
		return ret
	}
	return *o.EmbedUrl
}

// GetEmbedUrlOk returns a tuple with the EmbedUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetEmbedUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EmbedUrl) {
		return nil, false
	}
	return o.EmbedUrl, true
}

// HasEmbedUrl returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasEmbedUrl() bool {
	if o != nil && !utils.IsNil(o.EmbedUrl) {
		return true
	}

	return false
}

// SetEmbedUrl gets a reference to the given string and assigns it to the EmbedUrl field.
func (o *ApiTestTableWidget) SetEmbedUrl(v string) {
	o.EmbedUrl = &v
}

// GetIsEmbedded returns the IsEmbedded field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetIsEmbedded() bool {
	if o == nil || utils.IsNil(o.IsEmbedded) {
		var ret bool
		return ret
	}
	return *o.IsEmbedded
}

// GetIsEmbeddedOk returns a tuple with the IsEmbedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetIsEmbeddedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsEmbedded) {
		return nil, false
	}
	return o.IsEmbedded, true
}

// HasIsEmbedded returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasIsEmbedded() bool {
	if o != nil && !utils.IsNil(o.IsEmbedded) {
		return true
	}

	return false
}

// SetIsEmbedded gets a reference to the given bool and assigns it to the IsEmbedded field.
func (o *ApiTestTableWidget) SetIsEmbedded(v bool) {
	o.IsEmbedded = &v
}

// GetMetricGroup returns the MetricGroup field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetMetricGroup() MetricGroup {
	if o == nil || utils.IsNil(o.MetricGroup) {
		var ret MetricGroup
		return ret
	}
	return *o.MetricGroup
}

// GetMetricGroupOk returns a tuple with the MetricGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetMetricGroupOk() (*MetricGroup, bool) {
	if o == nil || utils.IsNil(o.MetricGroup) {
		return nil, false
	}
	return o.MetricGroup, true
}

// HasMetricGroup returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasMetricGroup() bool {
	if o != nil && !utils.IsNil(o.MetricGroup) {
		return true
	}

	return false
}

// SetMetricGroup gets a reference to the given MetricGroup and assigns it to the MetricGroup field.
func (o *ApiTestTableWidget) SetMetricGroup(v MetricGroup) {
	o.MetricGroup = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetDirection() DashboardMetricDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret DashboardMetricDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetDirectionOk() (*DashboardMetricDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given DashboardMetricDirection and assigns it to the Direction field.
func (o *ApiTestTableWidget) SetDirection(v DashboardMetricDirection) {
	o.Direction = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetMetric() DashboardMetric {
	if o == nil || utils.IsNil(o.Metric) {
		var ret DashboardMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetMetricOk() (*DashboardMetric, bool) {
	if o == nil || utils.IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasMetric() bool {
	if o != nil && !utils.IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given DashboardMetric and assigns it to the Metric field.
func (o *ApiTestTableWidget) SetMetric(v DashboardMetric) {
	o.Metric = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetFilters() map[string][]interface{} {
	if o == nil || utils.IsNil(o.Filters) {
		var ret map[string][]interface{}
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetFiltersOk() (*map[string][]interface{}, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string][]interface{} and assigns it to the Filters field.
func (o *ApiTestTableWidget) SetFilters(v map[string][]interface{}) {
	o.Filters = &v
}

// GetMeasure returns the Measure field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetMeasure() ApiWidgetMeasure {
	if o == nil || utils.IsNil(o.Measure) {
		var ret ApiWidgetMeasure
		return ret
	}
	return *o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetMeasureOk() (*ApiWidgetMeasure, bool) {
	if o == nil || utils.IsNil(o.Measure) {
		return nil, false
	}
	return o.Measure, true
}

// HasMeasure returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasMeasure() bool {
	if o != nil && !utils.IsNil(o.Measure) {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given ApiWidgetMeasure and assigns it to the Measure field.
func (o *ApiTestTableWidget) SetMeasure(v ApiWidgetMeasure) {
	o.Measure = &v
}

// GetFixedTimespan returns the FixedTimespan field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetFixedTimespan() ApiDuration {
	if o == nil || utils.IsNil(o.FixedTimespan) {
		var ret ApiDuration
		return ret
	}
	return *o.FixedTimespan
}

// GetFixedTimespanOk returns a tuple with the FixedTimespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetFixedTimespanOk() (*ApiDuration, bool) {
	if o == nil || utils.IsNil(o.FixedTimespan) {
		return nil, false
	}
	return o.FixedTimespan, true
}

// HasFixedTimespan returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasFixedTimespan() bool {
	if o != nil && !utils.IsNil(o.FixedTimespan) {
		return true
	}

	return false
}

// SetFixedTimespan gets a reference to the given ApiDuration and assigns it to the FixedTimespan field.
func (o *ApiTestTableWidget) SetFixedTimespan(v ApiDuration) {
	o.FixedTimespan = &v
}

// GetApiLink returns the ApiLink field value if set, zero value otherwise.
// Deprecated
func (o *ApiTestTableWidget) GetApiLink() string {
	if o == nil || utils.IsNil(o.ApiLink) {
		var ret string
		return ret
	}
	return *o.ApiLink
}

// GetApiLinkOk returns a tuple with the ApiLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApiTestTableWidget) GetApiLinkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ApiLink) {
		return nil, false
	}
	return o.ApiLink, true
}

// HasApiLink returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasApiLink() bool {
	if o != nil && !utils.IsNil(o.ApiLink) {
		return true
	}

	return false
}

// SetApiLink gets a reference to the given string and assigns it to the ApiLink field.
// Deprecated
func (o *ApiTestTableWidget) SetApiLink(v string) {
	o.ApiLink = &v
}

// GetShouldExcludeAlertSuppressionWindows returns the ShouldExcludeAlertSuppressionWindows field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetShouldExcludeAlertSuppressionWindows() bool {
	if o == nil || utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeAlertSuppressionWindows
}

// GetShouldExcludeAlertSuppressionWindowsOk returns a tuple with the ShouldExcludeAlertSuppressionWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetShouldExcludeAlertSuppressionWindowsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		return nil, false
	}
	return o.ShouldExcludeAlertSuppressionWindows, true
}

// HasShouldExcludeAlertSuppressionWindows returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasShouldExcludeAlertSuppressionWindows() bool {
	if o != nil && !utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		return true
	}

	return false
}

// SetShouldExcludeAlertSuppressionWindows gets a reference to the given bool and assigns it to the ShouldExcludeAlertSuppressionWindows field.
func (o *ApiTestTableWidget) SetShouldExcludeAlertSuppressionWindows(v bool) {
	o.ShouldExcludeAlertSuppressionWindows = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *ApiTestTableWidget) SetLinks(v SelfLinks) {
	o.Links = &v
}

// GetType returns the Type field value
func (o *ApiTestTableWidget) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApiTestTableWidget) SetType(v string) {
	o.Type = v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetFilter() ApiWidgetFilterApiTestTableFilterKey {
	if o == nil || utils.IsNil(o.Filter) {
		var ret ApiWidgetFilterApiTestTableFilterKey
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetFilterOk() (*ApiWidgetFilterApiTestTableFilterKey, bool) {
	if o == nil || utils.IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasFilter() bool {
	if o != nil && !utils.IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given ApiWidgetFilterApiTestTableFilterKey and assigns it to the Filter field.
func (o *ApiTestTableWidget) SetFilter(v ApiWidgetFilterApiTestTableFilterKey) {
	o.Filter = &v
}

// GetExclude returns the Exclude field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetExclude() ApiWidgetFilterApiTestTableFilterKey {
	if o == nil || utils.IsNil(o.Exclude) {
		var ret ApiWidgetFilterApiTestTableFilterKey
		return ret
	}
	return *o.Exclude
}

// GetExcludeOk returns a tuple with the Exclude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetExcludeOk() (*ApiWidgetFilterApiTestTableFilterKey, bool) {
	if o == nil || utils.IsNil(o.Exclude) {
		return nil, false
	}
	return o.Exclude, true
}

// HasExclude returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasExclude() bool {
	if o != nil && !utils.IsNil(o.Exclude) {
		return true
	}

	return false
}

// SetExclude gets a reference to the given ApiWidgetFilterApiTestTableFilterKey and assigns it to the Exclude field.
func (o *ApiTestTableWidget) SetExclude(v ApiWidgetFilterApiTestTableFilterKey) {
	o.Exclude = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ApiTestTableWidget) GetDataSource() TestTableDatasource {
	if o == nil || utils.IsNil(o.DataSource) {
		var ret TestTableDatasource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTestTableWidget) GetDataSourceOk() (*TestTableDatasource, bool) {
	if o == nil || utils.IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ApiTestTableWidget) HasDataSource() bool {
	if o != nil && !utils.IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given TestTableDatasource and assigns it to the DataSource field.
func (o *ApiTestTableWidget) SetDataSource(v TestTableDatasource) {
	o.DataSource = &v
}

func (o ApiTestTableWidget) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTestTableWidget) ToMap() (map[string]interface{}, error) {
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
	if !utils.IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !utils.IsNil(o.Exclude) {
		toSerialize["exclude"] = o.Exclude
	}
	if !utils.IsNil(o.DataSource) {
		toSerialize["dataSource"] = o.DataSource
	}
	return toSerialize, nil
}

func (o *ApiTestTableWidget) UnmarshalJSON(data []byte) (err error) {
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

	varApiTestTableWidget := _ApiTestTableWidget{}

    err = json.Unmarshal(data, &varApiTestTableWidget)

	if err != nil {
		return err
	}

	*o = ApiTestTableWidget(varApiTestTableWidget)

	return err
}

type NullableApiTestTableWidget struct {
	value *ApiTestTableWidget
	isSet bool
}

func (v NullableApiTestTableWidget) Get() *ApiTestTableWidget {
	return v.value
}

func (v *NullableApiTestTableWidget) Set(val *ApiTestTableWidget) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTestTableWidget) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTestTableWidget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTestTableWidget(val *ApiTestTableWidget) *NullableApiTestTableWidget {
	return &NullableApiTestTableWidget{value: val, isSet: true}
}

func (v NullableApiTestTableWidget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTestTableWidget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


