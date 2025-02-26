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

// NumbersCardDatasource Datasource of the numbers card widget.
type NumbersCardDatasource string

// List of NumbersCardDatasource
const (
	NUMBERSCARDDATASOURCE_ALERTS NumbersCardDatasource = "ALERTS"
	NUMBERSCARDDATASOURCE_CLOUD_AND_ENTERPRISE_AGENTS NumbersCardDatasource = "CLOUD_AND_ENTERPRISE_AGENTS"
	NUMBERSCARDDATASOURCE_DEVICES NumbersCardDatasource = "DEVICES"
	NUMBERSCARDDATASOURCE_ENDPOINT_AGENTS NumbersCardDatasource = "ENDPOINT_AGENTS"
	NUMBERSCARDDATASOURCE_ENDPOINT_AST_TEST NumbersCardDatasource = "ENDPOINT_AST_TEST"
	NUMBERSCARDDATASOURCE_ENDPOINT_BROWSER_SESSION NumbersCardDatasource = "ENDPOINT_BROWSER_SESSION"
	NUMBERSCARDDATASOURCE_ENDPOINT_LOCAL_NETWORK NumbersCardDatasource = "ENDPOINT_LOCAL_NETWORK"
	NUMBERSCARDDATASOURCE_ENDPOINT_LOCAL_NETWORK_WIRELESS NumbersCardDatasource = "ENDPOINT_LOCAL_NETWORK_WIRELESS"
	NUMBERSCARDDATASOURCE_ENDPOINT_SCHEDULED_TEST NumbersCardDatasource = "ENDPOINT_SCHEDULED_TEST"
	NUMBERSCARDDATASOURCE_INTERNET_INSIGHTS NumbersCardDatasource = "INTERNET_INSIGHTS"
	NUMBERSCARDDATASOURCE_ROUTING NumbersCardDatasource = "ROUTING"
	NUMBERSCARDDATASOURCE_CLOUD_NATIVE_MONITORING NumbersCardDatasource = "CLOUD_NATIVE_MONITORING"
)

// All allowed values of NumbersCardDatasource enum
var AllowedNumbersCardDatasourceEnumValues = []NumbersCardDatasource{
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

func (v *NumbersCardDatasource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NumbersCardDatasource(value)
	for _, existing := range AllowedNumbersCardDatasourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NumbersCardDatasource", value)
}

// NewNumbersCardDatasourceFromValue returns a pointer to a valid NumbersCardDatasource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNumbersCardDatasourceFromValue(v string) (*NumbersCardDatasource, error) {
	ev := NumbersCardDatasource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NumbersCardDatasource: valid values are %v", v, AllowedNumbersCardDatasourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NumbersCardDatasource) IsValid() bool {
	for _, existing := range AllowedNumbersCardDatasourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NumbersCardDatasource value
func (v NumbersCardDatasource) Ptr() *NumbersCardDatasource {
	return &v
}

type NullableNumbersCardDatasource struct {
	value *NumbersCardDatasource
	isSet bool
}

func (v NullableNumbersCardDatasource) Get() *NumbersCardDatasource {
	return v.value
}

func (v *NullableNumbersCardDatasource) Set(val *NumbersCardDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableNumbersCardDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableNumbersCardDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNumbersCardDatasource(val *NumbersCardDatasource) *NullableNumbersCardDatasource {
	return &NullableNumbersCardDatasource{value: val, isSet: true}
}

func (v NullableNumbersCardDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNumbersCardDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

