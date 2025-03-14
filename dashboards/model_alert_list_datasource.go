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

// AlertListDatasource Datasource of the alert list.
type AlertListDatasource string

// List of AlertListDatasource
const (
	ALERTLISTDATASOURCE_ALERTS AlertListDatasource = "ALERTS"
	ALERTLISTDATASOURCE_DEVICES AlertListDatasource = "DEVICES"
	ALERTLISTDATASOURCE_DNSP AlertListDatasource = "DNSP"
	ALERTLISTDATASOURCE_ENDPOINT_AGENTS AlertListDatasource = "ENDPOINT_AGENTS"
	ALERTLISTDATASOURCE_ENDPOINT_SCHEDULED_TEST AlertListDatasource = "ENDPOINT_SCHEDULED_TEST"
	ALERTLISTDATASOURCE_ENDPOINT_AST_TEST AlertListDatasource = "ENDPOINT_AST_TEST"
	ALERTLISTDATASOURCE_ENDPOINT_BROWSER_SESSION AlertListDatasource = "ENDPOINT_BROWSER_SESSION"
	ALERTLISTDATASOURCE_ENDPOINT_LOCAL_NETWORK AlertListDatasource = "ENDPOINT_LOCAL_NETWORK"
	ALERTLISTDATASOURCE_ENDPOINT_LOCAL_NETWORK_WIRELESS AlertListDatasource = "ENDPOINT_LOCAL_NETWORK_WIRELESS"
	ALERTLISTDATASOURCE_ROUTING AlertListDatasource = "ROUTING"
	ALERTLISTDATASOURCE_CLOUD_AND_ENTERPRISE_AGENTS AlertListDatasource = "CLOUD_AND_ENTERPRISE_AGENTS"
	ALERTLISTDATASOURCE_INTERNET_INSIGHTS AlertListDatasource = "INTERNET_INSIGHTS"
	ALERTLISTDATASOURCE_APPDYNAMICS_SERVICE_HEALTH AlertListDatasource = "APPDYNAMICS_SERVICE_HEALTH"
	ALERTLISTDATASOURCE_CLOUD_NATIVE_MONITORING AlertListDatasource = "CLOUD_NATIVE_MONITORING"
)

// All allowed values of AlertListDatasource enum
var AllowedAlertListDatasourceEnumValues = []AlertListDatasource{
	"ALERTS",
	"DEVICES",
	"DNSP",
	"ENDPOINT_AGENTS",
	"ENDPOINT_SCHEDULED_TEST",
	"ENDPOINT_AST_TEST",
	"ENDPOINT_BROWSER_SESSION",
	"ENDPOINT_LOCAL_NETWORK",
	"ENDPOINT_LOCAL_NETWORK_WIRELESS",
	"ROUTING",
	"CLOUD_AND_ENTERPRISE_AGENTS",
	"INTERNET_INSIGHTS",
	"APPDYNAMICS_SERVICE_HEALTH",
	"CLOUD_NATIVE_MONITORING",
}

func (v *AlertListDatasource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertListDatasource(value)
	for _, existing := range AllowedAlertListDatasourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertListDatasource", value)
}

// NewAlertListDatasourceFromValue returns a pointer to a valid AlertListDatasource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertListDatasourceFromValue(v string) (*AlertListDatasource, error) {
	ev := AlertListDatasource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertListDatasource: valid values are %v", v, AllowedAlertListDatasourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertListDatasource) IsValid() bool {
	for _, existing := range AllowedAlertListDatasourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertListDatasource value
func (v AlertListDatasource) Ptr() *AlertListDatasource {
	return &v
}

type NullableAlertListDatasource struct {
	value *AlertListDatasource
	isSet bool
}

func (v NullableAlertListDatasource) Get() *AlertListDatasource {
	return v.value
}

func (v *NullableAlertListDatasource) Set(val *AlertListDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertListDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertListDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertListDatasource(val *AlertListDatasource) *NullableAlertListDatasource {
	return &NullableAlertListDatasource{value: val, isSet: true}
}

func (v NullableAlertListDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertListDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

