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

// MultiMetricsTableDatasource Datasource of the Multi-Metrics table widget.
type MultiMetricsTableDatasource string

// List of MultiMetricsTableDatasource
const (
	MULTIMETRICSTABLEDATASOURCE_ALERTS MultiMetricsTableDatasource = "ALERTS"
	MULTIMETRICSTABLEDATASOURCE_CLOUD_AND_ENTERPRISE_AGENTS MultiMetricsTableDatasource = "CLOUD_AND_ENTERPRISE_AGENTS"
	MULTIMETRICSTABLEDATASOURCE_DEVICES MultiMetricsTableDatasource = "DEVICES"
	MULTIMETRICSTABLEDATASOURCE_ENDPOINT_AGENTS MultiMetricsTableDatasource = "ENDPOINT_AGENTS"
	MULTIMETRICSTABLEDATASOURCE_ENDPOINT_AST_TEST MultiMetricsTableDatasource = "ENDPOINT_AST_TEST"
	MULTIMETRICSTABLEDATASOURCE_ENDPOINT_BROWSER_SESSION MultiMetricsTableDatasource = "ENDPOINT_BROWSER_SESSION"
	MULTIMETRICSTABLEDATASOURCE_ENDPOINT_LOCAL_NETWORK MultiMetricsTableDatasource = "ENDPOINT_LOCAL_NETWORK"
	MULTIMETRICSTABLEDATASOURCE_ENDPOINT_LOCAL_NETWORK_WIRELESS MultiMetricsTableDatasource = "ENDPOINT_LOCAL_NETWORK_WIRELESS"
	MULTIMETRICSTABLEDATASOURCE_ENDPOINT_SCHEDULED_TEST MultiMetricsTableDatasource = "ENDPOINT_SCHEDULED_TEST"
	MULTIMETRICSTABLEDATASOURCE_INTERNET_INSIGHTS MultiMetricsTableDatasource = "INTERNET_INSIGHTS"
	MULTIMETRICSTABLEDATASOURCE_ROUTING MultiMetricsTableDatasource = "ROUTING"
	MULTIMETRICSTABLEDATASOURCE_CLOUD_NATIVE_MONITORING MultiMetricsTableDatasource = "CLOUD_NATIVE_MONITORING"
)

// All allowed values of MultiMetricsTableDatasource enum
var AllowedMultiMetricsTableDatasourceEnumValues = []MultiMetricsTableDatasource{
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

func (v *MultiMetricsTableDatasource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MultiMetricsTableDatasource(value)
	for _, existing := range AllowedMultiMetricsTableDatasourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MultiMetricsTableDatasource", value)
}

// NewMultiMetricsTableDatasourceFromValue returns a pointer to a valid MultiMetricsTableDatasource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMultiMetricsTableDatasourceFromValue(v string) (*MultiMetricsTableDatasource, error) {
	ev := MultiMetricsTableDatasource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MultiMetricsTableDatasource: valid values are %v", v, AllowedMultiMetricsTableDatasourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MultiMetricsTableDatasource) IsValid() bool {
	for _, existing := range AllowedMultiMetricsTableDatasourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MultiMetricsTableDatasource value
func (v MultiMetricsTableDatasource) Ptr() *MultiMetricsTableDatasource {
	return &v
}

type NullableMultiMetricsTableDatasource struct {
	value *MultiMetricsTableDatasource
	isSet bool
}

func (v NullableMultiMetricsTableDatasource) Get() *MultiMetricsTableDatasource {
	return v.value
}

func (v *NullableMultiMetricsTableDatasource) Set(val *MultiMetricsTableDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiMetricsTableDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiMetricsTableDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiMetricsTableDatasource(val *MultiMetricsTableDatasource) *NullableMultiMetricsTableDatasource {
	return &NullableMultiMetricsTableDatasource{value: val, isSet: true}
}

func (v NullableMultiMetricsTableDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiMetricsTableDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

