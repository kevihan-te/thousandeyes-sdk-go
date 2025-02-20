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
)

// checks if the ApiMultiMetricColumn type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiMultiMetricColumn{}

// ApiMultiMetricColumn Defines a column within a table that aggregates and displays various metrics (Multi-Metric table).
type ApiMultiMetricColumn struct {
	Id *string `json:"id,omitempty"`
	DataSource *MultiMetricsTableDatasource `json:"dataSource,omitempty"`
	MetricGroup *MetricGroup `json:"metricGroup,omitempty"`
	Direction *DashboardMetricDirection `json:"direction,omitempty"`
	Metric *DashboardMetric `json:"metric,omitempty"`
	// (Optional) Specifies the filters applied to the widget. When present, the `filters` property displays. Each filter object has two properties: `filterProperty` and `filterValue`. The `filterProperty` can be values like `AGENT`, `ENDPOINT_MACHINE_ID`, `TEST`, `MONITOR`, etc.  The `filterValue` represents an identifier array of the selected property.
	Filters *map[string][]interface{} `json:"filters,omitempty"`
	Measure *ApiWidgetMeasure `json:"measure,omitempty"`
}

// NewApiMultiMetricColumn instantiates a new ApiMultiMetricColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMultiMetricColumn() *ApiMultiMetricColumn {
	this := ApiMultiMetricColumn{}
	return &this
}

// NewApiMultiMetricColumnWithDefaults instantiates a new ApiMultiMetricColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMultiMetricColumnWithDefaults() *ApiMultiMetricColumn {
	this := ApiMultiMetricColumn{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiMultiMetricColumn) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiMetricColumn) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiMultiMetricColumn) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiMultiMetricColumn) SetId(v string) {
	o.Id = &v
}

// GetDataSource returns the DataSource field value if set, zero value otherwise.
func (o *ApiMultiMetricColumn) GetDataSource() MultiMetricsTableDatasource {
	if o == nil || utils.IsNil(o.DataSource) {
		var ret MultiMetricsTableDatasource
		return ret
	}
	return *o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiMetricColumn) GetDataSourceOk() (*MultiMetricsTableDatasource, bool) {
	if o == nil || utils.IsNil(o.DataSource) {
		return nil, false
	}
	return o.DataSource, true
}

// HasDataSource returns a boolean if a field has been set.
func (o *ApiMultiMetricColumn) HasDataSource() bool {
	if o != nil && !utils.IsNil(o.DataSource) {
		return true
	}

	return false
}

// SetDataSource gets a reference to the given MultiMetricsTableDatasource and assigns it to the DataSource field.
func (o *ApiMultiMetricColumn) SetDataSource(v MultiMetricsTableDatasource) {
	o.DataSource = &v
}

// GetMetricGroup returns the MetricGroup field value if set, zero value otherwise.
func (o *ApiMultiMetricColumn) GetMetricGroup() MetricGroup {
	if o == nil || utils.IsNil(o.MetricGroup) {
		var ret MetricGroup
		return ret
	}
	return *o.MetricGroup
}

// GetMetricGroupOk returns a tuple with the MetricGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiMetricColumn) GetMetricGroupOk() (*MetricGroup, bool) {
	if o == nil || utils.IsNil(o.MetricGroup) {
		return nil, false
	}
	return o.MetricGroup, true
}

// HasMetricGroup returns a boolean if a field has been set.
func (o *ApiMultiMetricColumn) HasMetricGroup() bool {
	if o != nil && !utils.IsNil(o.MetricGroup) {
		return true
	}

	return false
}

// SetMetricGroup gets a reference to the given MetricGroup and assigns it to the MetricGroup field.
func (o *ApiMultiMetricColumn) SetMetricGroup(v MetricGroup) {
	o.MetricGroup = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *ApiMultiMetricColumn) GetDirection() DashboardMetricDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret DashboardMetricDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiMetricColumn) GetDirectionOk() (*DashboardMetricDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *ApiMultiMetricColumn) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given DashboardMetricDirection and assigns it to the Direction field.
func (o *ApiMultiMetricColumn) SetDirection(v DashboardMetricDirection) {
	o.Direction = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *ApiMultiMetricColumn) GetMetric() DashboardMetric {
	if o == nil || utils.IsNil(o.Metric) {
		var ret DashboardMetric
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiMetricColumn) GetMetricOk() (*DashboardMetric, bool) {
	if o == nil || utils.IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *ApiMultiMetricColumn) HasMetric() bool {
	if o != nil && !utils.IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given DashboardMetric and assigns it to the Metric field.
func (o *ApiMultiMetricColumn) SetMetric(v DashboardMetric) {
	o.Metric = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ApiMultiMetricColumn) GetFilters() map[string][]interface{} {
	if o == nil || utils.IsNil(o.Filters) {
		var ret map[string][]interface{}
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiMetricColumn) GetFiltersOk() (*map[string][]interface{}, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ApiMultiMetricColumn) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given map[string][]interface{} and assigns it to the Filters field.
func (o *ApiMultiMetricColumn) SetFilters(v map[string][]interface{}) {
	o.Filters = &v
}

// GetMeasure returns the Measure field value if set, zero value otherwise.
func (o *ApiMultiMetricColumn) GetMeasure() ApiWidgetMeasure {
	if o == nil || utils.IsNil(o.Measure) {
		var ret ApiWidgetMeasure
		return ret
	}
	return *o.Measure
}

// GetMeasureOk returns a tuple with the Measure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMultiMetricColumn) GetMeasureOk() (*ApiWidgetMeasure, bool) {
	if o == nil || utils.IsNil(o.Measure) {
		return nil, false
	}
	return o.Measure, true
}

// HasMeasure returns a boolean if a field has been set.
func (o *ApiMultiMetricColumn) HasMeasure() bool {
	if o != nil && !utils.IsNil(o.Measure) {
		return true
	}

	return false
}

// SetMeasure gets a reference to the given ApiWidgetMeasure and assigns it to the Measure field.
func (o *ApiMultiMetricColumn) SetMeasure(v ApiWidgetMeasure) {
	o.Measure = &v
}

func (o ApiMultiMetricColumn) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMultiMetricColumn) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	if !utils.IsNil(o.Measure) {
		toSerialize["measure"] = o.Measure
	}
	return toSerialize, nil
}

type NullableApiMultiMetricColumn struct {
	value *ApiMultiMetricColumn
	isSet bool
}

func (v NullableApiMultiMetricColumn) Get() *ApiMultiMetricColumn {
	return v.value
}

func (v *NullableApiMultiMetricColumn) Set(val *ApiMultiMetricColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMultiMetricColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMultiMetricColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMultiMetricColumn(val *ApiMultiMetricColumn) *NullableApiMultiMetricColumn {
	return &NullableApiMultiMetricColumn{value: val, isSet: true}
}

func (v NullableApiMultiMetricColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMultiMetricColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


