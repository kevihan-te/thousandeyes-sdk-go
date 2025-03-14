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

// GeoMapDatasource Datasource of the geo map widget.
type GeoMapDatasource string

// List of GeoMapDatasource
const (
	GEOMAPDATASOURCE_ALERTS GeoMapDatasource = "ALERTS"
	GEOMAPDATASOURCE_CLOUD_AND_ENTERPRISE_AGENTS GeoMapDatasource = "CLOUD_AND_ENTERPRISE_AGENTS"
	GEOMAPDATASOURCE_ENDPOINT_AGENTS GeoMapDatasource = "ENDPOINT_AGENTS"
	GEOMAPDATASOURCE_ENDPOINT_AST_TEST GeoMapDatasource = "ENDPOINT_AST_TEST"
	GEOMAPDATASOURCE_ENDPOINT_BROWSER_SESSION GeoMapDatasource = "ENDPOINT_BROWSER_SESSION"
	GEOMAPDATASOURCE_ENDPOINT_LOCAL_NETWORK GeoMapDatasource = "ENDPOINT_LOCAL_NETWORK"
	GEOMAPDATASOURCE_ENDPOINT_LOCAL_NETWORK_WIRELESS GeoMapDatasource = "ENDPOINT_LOCAL_NETWORK_WIRELESS"
	GEOMAPDATASOURCE_ENDPOINT_SCHEDULED_TEST GeoMapDatasource = "ENDPOINT_SCHEDULED_TEST"
	GEOMAPDATASOURCE_INTERNET_INSIGHTS GeoMapDatasource = "INTERNET_INSIGHTS"
	GEOMAPDATASOURCE_ROUTING GeoMapDatasource = "ROUTING"
	GEOMAPDATASOURCE_CLOUD_NATIVE_MONITORING GeoMapDatasource = "CLOUD_NATIVE_MONITORING"
)

// All allowed values of GeoMapDatasource enum
var AllowedGeoMapDatasourceEnumValues = []GeoMapDatasource{
	"ALERTS",
	"CLOUD_AND_ENTERPRISE_AGENTS",
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

func (v *GeoMapDatasource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GeoMapDatasource(value)
	for _, existing := range AllowedGeoMapDatasourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GeoMapDatasource", value)
}

// NewGeoMapDatasourceFromValue returns a pointer to a valid GeoMapDatasource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGeoMapDatasourceFromValue(v string) (*GeoMapDatasource, error) {
	ev := GeoMapDatasource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GeoMapDatasource: valid values are %v", v, AllowedGeoMapDatasourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GeoMapDatasource) IsValid() bool {
	for _, existing := range AllowedGeoMapDatasourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GeoMapDatasource value
func (v GeoMapDatasource) Ptr() *GeoMapDatasource {
	return &v
}

type NullableGeoMapDatasource struct {
	value *GeoMapDatasource
	isSet bool
}

func (v NullableGeoMapDatasource) Get() *GeoMapDatasource {
	return v.value
}

func (v *NullableGeoMapDatasource) Set(val *GeoMapDatasource) {
	v.value = val
	v.isSet = true
}

func (v NullableGeoMapDatasource) IsSet() bool {
	return v.isSet
}

func (v *NullableGeoMapDatasource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeoMapDatasource(val *GeoMapDatasource) *NullableGeoMapDatasource {
	return &NullableGeoMapDatasource{value: val, isSet: true}
}

func (v NullableGeoMapDatasource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeoMapDatasource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

