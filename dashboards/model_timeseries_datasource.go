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

// TimeseriesDatasource Datasource of the Timeseries widget.
type TimeseriesDatasource string

// List of TimeseriesDatasource
const (
	TIMESERIESDATASOURCE_ALERTS TimeseriesDatasource = "ALERTS"
	TIMESERIESDATASOURCE_CLOUD_AND_ENTERPRISE_AGENTS TimeseriesDatasource = "CLOUD_AND_ENTERPRISE_AGENTS"
	TIMESERIESDATASOURCE_DEVICES TimeseriesDatasource = "DEVICES"
	TIMESERIESDATASOURCE_ENDPOINT_AGENTS TimeseriesDatasource = "ENDPOINT_AGENTS"
	TIMESERIESDATASOURCE_ENDPOINT_AST_TEST TimeseriesDatasource = "ENDPOINT_AST_TEST"
	TIMESERIESDATASOURCE_ENDPOINT_BROWSER_SESSION TimeseriesDatasource = "ENDPOINT_BROWSER_SESSION"
	TIMESERIESDATASOURCE_ENDPOINT_LOCAL_NETWORK TimeseriesDatasource = "ENDPOINT_LOCAL_NETWORK"
	TIMESERIESDATASOURCE_ENDPOINT_LOCAL_NETWORK_WIRELESS TimeseriesDatasource = "ENDPOINT_LOCAL_NETWORK_WIRELESS"
	TIMESERIESDATASOURCE_ENDPOINT_SCHEDULED_TEST TimeseriesDatasource = "ENDPOINT_SCHEDULED_TEST"
	TIMESERIESDATASOURCE_INTERNET_INSIGHTS TimeseriesDatasource = "INTERNET_INSIGHTS"
	TIMESERIESDATASOURCE_ROUTING TimeseriesDatasource = "ROUTING"
	TIMESERIESDATASOURCE_CLOUD_NATIVE_MONITORING TimeseriesDatasource = "CLOUD_NATIVE_MONITORING"
)

// All allowed values of TimeseriesDatasource enum
var AllowedTimeseriesDatasourceEnumValues = []TimeseriesDatasource{
	"ALERTS",
	"CLOUD_AND_ENTERPRISE_AGENTS",
	"DEVICES",
	"ENDPOINT_AGENTS",
	"ENDPOINT_AST_TEST",
	"ENDPOINT_BROWSER_SESSION",
	"ENDPOINT_LOCAL_NETWORK",
	"ENDPOINT_LOCAL_NETWORK_WIRELESS",
	"ENDPOINT_SCHEDULED_TEST",
	"INTERNET_INSIGHTS",
	"ROUTING",
	"CLOUD_NATIVE_MONITORING",
}

func (v *TimeseriesDatasource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeseriesDatasource(value)
	for _, existing := range AllowedTimeseriesDatasourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeseriesDatasource", value)
}

// NewTimeseriesDatasourceFromValue returns a pointer to a valid TimeseriesDatasource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeseriesDatasourceFromValue(v string) (*TimeseriesDatasource, error) {
	ev := TimeseriesDatasource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeseriesDatasource: valid values are %v", v, AllowedTimeseriesDatasourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeseriesDatasource) IsValid() bool {
	for _, existing := range AllowedTimeseriesDatasourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TimeseriesDatasource value
func (v TimeseriesDatasource) Ptr() *TimeseriesDatasource {
	return &v
}

type NullableTimeseriesDatasource struct {
	value *TimeseriesDatasource
	isSet bool
}

func (v NullableTimeseriesDatasource) Get() *TimeseriesDatasource {
	return v.value
}

func (v *NullableTimeseriesDatasource) Set(val *TimeseriesDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesDatasource(val *TimeseriesDatasource) *NullableTimeseriesDatasource {
	return &NullableTimeseriesDatasource{value: val, isSet: true}
}

func (v NullableTimeseriesDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

