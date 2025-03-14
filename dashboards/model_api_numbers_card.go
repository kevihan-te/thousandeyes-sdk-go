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
)

// checks if the ApiNumbersCard type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiNumbersCard{}

// ApiNumbersCard An individual number card within the numbers card widget.
type ApiNumbersCard struct {
	// Minimum scale configured in the widget.
	MinScale *float32 `json:"minScale,omitempty"`
	// Maximum scale configured in the widget.
	MaxScale *float32 `json:"maxScale,omitempty"`
	Unit *ApiWidgetFixedYScalePrefix `json:"unit,omitempty"`
	// Identifier of the widget.
	Id *string `json:"id,omitempty"`
	// Description of the number card.
	Description *string `json:"description,omitempty"`
	Measure *ApiWidgetMeasure `json:"measure,omitempty"`
	CompareToPreviousValue *bool `json:"compareToPreviousValue,omitempty"`
	FixedTimespan *ApiDuration `json:"fixedTimespan,omitempty"`
	// Excludes alert suppression windows from the widget when set to `true`.
	ShouldExcludeAlertSuppressionWindows *bool `json:"shouldExcludeAlertSuppressionWindows,omitempty"`
	DataSource *NumbersCardDatasource `json:"dataSource,omitempty"`
	MetricGroup *MetricGroup `json:"metricGroup,omitempty"`
	Direction *DashboardMetricDirection `json:"direction,omitempty"`
	Metric *DashboardMetric `json:"metric,omitempty"`
	// (Optional) Specifies the filters applied to the widget. When present, the `filters` property displays. Each filter object has two properties: `filterProperty` and `filterValue`. The `filterProperty` can be values like `AGENT`, `ENDPOINT_MACHINE_ID`, `TEST`, `MONITOR`, etc.  The `filterValue` represents an identifier array of the selected property.
	Filters *map[string][]interface{} `json:"filters,omitempty"`
}

// NewApiNumbersCard instantiates a new ApiNumbersCard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiNumbersCard() *ApiNumbersCard {
	this := ApiNumbersCard{}
	return &this
}

// NewApiNumbersCardWithDefaults instantiates a new ApiNumbersCard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiNumbersCardWithDefaults() *ApiNumbersCard {
	this := ApiNumbersCard{}
	return &this
}

// GetMinScale returns the MinScale field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetMinScale() float32 {
	if o == nil || utils.IsNil(o.MinScale) {
		var ret float32
		return ret
	}
	return *o.MinScale
}

// GetMinScaleOk returns a tuple with the MinScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetMinScaleOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.MinScale) {
		return nil, false
	}
	return o.MinScale, true
}

// HasMinScale returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasMinScale() bool {
	if o != nil && !utils.IsNil(o.MinScale) {
		return true
	}

	return false
}

// SetMinScale gets a reference to the given float32 and assigns it to the MinScale field.
func (o *ApiNumbersCard) SetMinScale(v float32) {
	o.MinScale = &v
}

// GetMaxScale returns the MaxScale field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetMaxScale() float32 {
	if o == nil || utils.IsNil(o.MaxScale) {
		var ret float32
		return ret
	}
	return *o.MaxScale
}

// GetMaxScaleOk returns a tuple with the MaxScale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetMaxScaleOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.MaxScale) {
		return nil, false
	}
	return o.MaxScale, true
}

// HasMaxScale returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasMaxScale() bool {
	if o != nil && !utils.IsNil(o.MaxScale) {
		return true
	}

	return false
}

// SetMaxScale gets a reference to the given float32 and assigns it to the MaxScale field.
func (o *ApiNumbersCard) SetMaxScale(v float32) {
	o.MaxScale = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetUnit() ApiWidgetFixedYScalePrefix {
	if o == nil || utils.IsNil(o.Unit) {
		var ret ApiWidgetFixedYScalePrefix
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetUnitOk() (*ApiWidgetFixedYScalePrefix, bool) {
	if o == nil || utils.IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasUnit() bool {
	if o != nil && !utils.IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given ApiWidgetFixedYScalePrefix and assigns it to the Unit field.
func (o *ApiNumbersCard) SetUnit(v ApiWidgetFixedYScalePrefix) {
	o.Unit = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiNumbersCard) SetId(v string) {
	o.Id = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiNumbersCard) SetDescription(v string) {
	o.Description = &v
}

// GetMeasure returns the Measure field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetMeasure() ApiWidgetMeasure {
	if o == nil || utils.IsNil(o.Measure) {
		var ret ApiWidgetMeasure
		return ret
	}
	return *o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetMeasureOk() (*ApiWidgetMeasure, bool) {
	if o == nil || utils.IsNil(o.Measure) {
		return nil, false
	}
	return o.Measure, true
}

// HasMeasure returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasMeasure() bool {
	if o != nil && !utils.IsNil(o.Measure) {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given ApiWidgetMeasure and assigns it to the Measure field.
func (o *ApiNumbersCard) SetMeasure(v ApiWidgetMeasure) {
	o.Measure = &v
}

// GetCompareToPreviousValue returns the CompareToPreviousValue field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetCompareToPreviousValue() bool {
	if o == nil || utils.IsNil(o.CompareToPreviousValue) {
		var ret bool
		return ret
	}
	return *o.CompareToPreviousValue
}

// GetCompareToPreviousValueOk returns a tuple with the CompareToPreviousValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetCompareToPreviousValueOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.CompareToPreviousValue) {
		return nil, false
	}
	return o.CompareToPreviousValue, true
}

// HasCompareToPreviousValue returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasCompareToPreviousValue() bool {
	if o != nil && !utils.IsNil(o.CompareToPreviousValue) {
		return true
	}

	return false
}

// SetCompareToPreviousValue gets a reference to the given bool and assigns it to the CompareToPreviousValue field.
func (o *ApiNumbersCard) SetCompareToPreviousValue(v bool) {
	o.CompareToPreviousValue = &v
}

// GetFixedTimespan returns the FixedTimespan field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetFixedTimespan() ApiDuration {
	if o == nil || utils.IsNil(o.FixedTimespan) {
		var ret ApiDuration
		return ret
	}
	return *o.FixedTimespan
}

// GetFixedTimespanOk returns a tuple with the FixedTimespan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetFixedTimespanOk() (*ApiDuration, bool) {
	if o == nil || utils.IsNil(o.FixedTimespan) {
		return nil, false
	}
	return o.FixedTimespan, true
}

// HasFixedTimespan returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasFixedTimespan() bool {
	if o != nil && !utils.IsNil(o.FixedTimespan) {
		return true
	}

	return false
}

// SetFixedTimespan gets a reference to the given ApiDuration and assigns it to the FixedTimespan field.
func (o *ApiNumbersCard) SetFixedTimespan(v ApiDuration) {
	o.FixedTimespan = &v
}

// GetShouldExcludeAlertSuppressionWindows returns the ShouldExcludeAlertSuppressionWindows field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetShouldExcludeAlertSuppressionWindows() bool {
	if o == nil || utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		var ret bool
		return ret
	}
	return *o.ShouldExcludeAlertSuppressionWindows
}

// GetShouldExcludeAlertSuppressionWindowsOk returns a tuple with the ShouldExcludeAlertSuppressionWindows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetShouldExcludeAlertSuppressionWindowsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		return nil, false
	}
	return o.ShouldExcludeAlertSuppressionWindows, true
}

// HasShouldExcludeAlertSuppressionWindows returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasShouldExcludeAlertSuppressionWindows() bool {
	if o != nil && !utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		return true
	}

	return false
}

// SetShouldExcludeAlertSuppressionWindows gets a reference to the given bool and assigns it to the ShouldExcludeAlertSuppressionWindows field.
func (o *ApiNumbersCard) SetShouldExcludeAlertSuppressionWindows(v bool) {
	o.ShouldExcludeAlertSuppressionWindows = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetDataSource() NumbersCardDatasource {
	if o == nil || utils.IsNil(o.DataSource) {
		var ret NumbersCardDatasource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetDataSourceOk() (*NumbersCardDatasource, bool) {
	if o == nil || utils.IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasDataSource() bool {
	if o != nil && !utils.IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given NumbersCardDatasource and assigns it to the DataSource field.
func (o *ApiNumbersCard) SetDataSource(v NumbersCardDatasource) {
	o.DataSource = &v
}

// GetMetricGroup returns the MetricGroup field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetMetricGroup() MetricGroup {
	if o == nil || utils.IsNil(o.MetricGroup) {
		var ret MetricGroup
		return ret
	}
	return *o.MetricGroup
}

// GetMetricGroupOk returns a tuple with the MetricGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetMetricGroupOk() (*MetricGroup, bool) {
	if o == nil || utils.IsNil(o.MetricGroup) {
		return nil, false
	}
	return o.MetricGroup, true
}

// HasMetricGroup returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasMetricGroup() bool {
	if o != nil && !utils.IsNil(o.MetricGroup) {
		return true
	}

	return false
}

// SetMetricGroup gets a reference to the given MetricGroup and assigns it to the MetricGroup field.
func (o *ApiNumbersCard) SetMetricGroup(v MetricGroup) {
	o.MetricGroup = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetDirection() DashboardMetricDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret DashboardMetricDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetDirectionOk() (*DashboardMetricDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given DashboardMetricDirection and assigns it to the Direction field.
func (o *ApiNumbersCard) SetDirection(v DashboardMetricDirection) {
	o.Direction = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetMetric() DashboardMetric {
	if o == nil || utils.IsNil(o.Metric) {
		var ret DashboardMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetMetricOk() (*DashboardMetric, bool) {
	if o == nil || utils.IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasMetric() bool {
	if o != nil && !utils.IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given DashboardMetric and assigns it to the Metric field.
func (o *ApiNumbersCard) SetMetric(v DashboardMetric) {
	o.Metric = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ApiNumbersCard) GetFilters() map[string][]interface{} {
	if o == nil || utils.IsNil(o.Filters) {
		var ret map[string][]interface{}
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiNumbersCard) GetFiltersOk() (*map[string][]interface{}, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ApiNumbersCard) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string][]interface{} and assigns it to the Filters field.
func (o *ApiNumbersCard) SetFilters(v map[string][]interface{}) {
	o.Filters = &v
}

func (o ApiNumbersCard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiNumbersCard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.MinScale) {
		toSerialize["minScale"] = o.MinScale
	}
	if !utils.IsNil(o.MaxScale) {
		toSerialize["maxScale"] = o.MaxScale
	}
	if !utils.IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.Measure) {
		toSerialize["measure"] = o.Measure
	}
	if !utils.IsNil(o.CompareToPreviousValue) {
		toSerialize["compareToPreviousValue"] = o.CompareToPreviousValue
	}
	if !utils.IsNil(o.FixedTimespan) {
		toSerialize["fixedTimespan"] = o.FixedTimespan
	}
	if !utils.IsNil(o.ShouldExcludeAlertSuppressionWindows) {
		toSerialize["shouldExcludeAlertSuppressionWindows"] = o.ShouldExcludeAlertSuppressionWindows
	}
	if !utils.IsNil(o.DataSource) {
		toSerialize["dataSource"] = o.DataSource
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
	return toSerialize, nil
}

type NullableApiNumbersCard struct {
	value *ApiNumbersCard
	isSet bool
}

func (v NullableApiNumbersCard) Get() *ApiNumbersCard {
	return v.value
}

func (v *NullableApiNumbersCard) Set(val *ApiNumbersCard) {
	v.value = val
	v.isSet = true
}

func (v NullableApiNumbersCard) IsSet() bool {
	return v.isSet
}

func (v *NullableApiNumbersCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiNumbersCard(val *ApiNumbersCard) *NullableApiNumbersCard {
	return &NullableApiNumbersCard{value: val, isSet: true}
}

func (v NullableApiNumbersCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiNumbersCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


