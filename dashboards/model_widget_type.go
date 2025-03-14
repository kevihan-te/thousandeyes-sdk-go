/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
	"fmt"
)

// WidgetType Type of the Widget
type WidgetType string

// List of WidgetType
const (
	WIDGETTYPE_BAR_CHART_STACKED WidgetType = "Bar Chart: Stacked"
	WIDGETTYPE_BAR_CHART_GROUPED WidgetType = "Bar Chart: Grouped"
	WIDGETTYPE_TIME_SERIES_LINE WidgetType = "Time Series: Line"
	WIDGETTYPE_TIME_SERIES_STACKED_AREA WidgetType = "Time Series: Stacked Area"
	WIDGETTYPE_PIE_CHART WidgetType = "Pie Chart"
	WIDGETTYPE_TABLE WidgetType = "Table"
	WIDGETTYPE_MULTI_METRIC_TABLE WidgetType = "Multi Metric Table"
	WIDGETTYPE_NUMBER WidgetType = "Number"
	WIDGETTYPE_AGENT_STATUS WidgetType = "Agent Status"
	WIDGETTYPE_COLOR_GRID WidgetType = "Color Grid"
	WIDGETTYPE_ALERT_LIST WidgetType = "Alert List"
	WIDGETTYPE_TEST_TABLE WidgetType = "Test Table"
	WIDGETTYPE_MAP WidgetType = "Map"
	WIDGETTYPE_BOX_AND_WHISKERS WidgetType = "Box and Whiskers"
)

// All allowed values of WidgetType enum
var AllowedWidgetTypeEnumValues = []WidgetType{
	"Bar Chart: Stacked",
	"Bar Chart: Grouped",
	"Time Series: Line",
	"Time Series: Stacked Area",
	"Pie Chart",
	"Table",
	"Multi Metric Table",
	"Number",
	"Agent Status",
	"Color Grid",
	"Alert List",
	"Test Table",
	"Map",
	"Box and Whiskers",
}

func (v *WidgetType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetType(value)
	for _, existing := range AllowedWidgetTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetType", value)
}

// NewWidgetTypeFromValue returns a pointer to a valid WidgetType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWidgetTypeFromValue(v string) (*WidgetType, error) {
	ev := WidgetType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WidgetType: valid values are %v", v, AllowedWidgetTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WidgetType) IsValid() bool {
	for _, existing := range AllowedWidgetTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WidgetType value
func (v WidgetType) Ptr() *WidgetType {
	return &v
}

type NullableWidgetType struct {
	value *WidgetType
	isSet bool
}

func (v NullableWidgetType) Get() *WidgetType {
	return v.value
}

func (v *NullableWidgetType) Set(val *WidgetType) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetType) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetType(val *WidgetType) *NullableWidgetType {
	return &NullableWidgetType{value: val, isSet: true}
}

func (v NullableWidgetType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

