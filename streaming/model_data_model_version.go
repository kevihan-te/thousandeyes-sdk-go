/*
ThousandEyes for OpenTelemetry API

ThousandEyes for OpenTelemetry provides machine-to-machine integration between ThousandEyes and its customers. It allows you to export ThousandEyes telemetry data in OTel format, which is widely used in the industry. With ThousandEyes for OTel, you can leverage frameworks widely used in the observability domain - such as Splunk, Grafana, and Honeycomb - to capture and analyze ThousandEyes data. Any client that supports OTel can use ThousandEyes for OpenTelemetry.  ThousandEyes for OTel is made up of the following components:  * Data streaming APIs that you can use to configure and enable your ThousandEyes tests with OTel-compatible streams, in particular to configure how ThousandEyes telemetry data is exported to client integrations. * A set of streaming pipelines called _collectors_ that actively fetch ThousandEyes network test data, enrich the data with some additional detail, filter, and push the data to the customer-configured endpoints, depending on what you configure via the public APIs. * Third-party OTel collectors that receive, transform, filter, and export different metrics to client applications such as AppD, or any other OTel-capable client configuration.  For more information about ThousandEyes for OpenTelemetry, see the [documentation](https://docs.thousandeyes.com/product-documentation/api/opentelemetry). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
	"fmt"
)

// DataModelVersion The version of the data model used in the data stream. When using `v1`:   - The `signal` cannot be `trace`.
type DataModelVersion string

// List of DataModelVersion
const (
	DATAMODELVERSION_V1 DataModelVersion = "v1"
	DATAMODELVERSION_V2 DataModelVersion = "v2"
)

// All allowed values of DataModelVersion enum
var AllowedDataModelVersionEnumValues = []DataModelVersion{
	"v1",
	"v2",
}

func (v *DataModelVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataModelVersion(value)
	for _, existing := range AllowedDataModelVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataModelVersion", value)
}

// NewDataModelVersionFromValue returns a pointer to a valid DataModelVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataModelVersionFromValue(v string) (*DataModelVersion, error) {
	ev := DataModelVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataModelVersion: valid values are %v", v, AllowedDataModelVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataModelVersion) IsValid() bool {
	for _, existing := range AllowedDataModelVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataModelVersion value
func (v DataModelVersion) Ptr() *DataModelVersion {
	return &v
}

type NullableDataModelVersion struct {
	value *DataModelVersion
	isSet bool
}

func (v NullableDataModelVersion) Get() *DataModelVersion {
	return v.value
}

func (v *NullableDataModelVersion) Set(val *DataModelVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableDataModelVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableDataModelVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataModelVersion(val *DataModelVersion) *NullableDataModelVersion {
	return &NullableDataModelVersion{value: val, isSet: true}
}

func (v NullableDataModelVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataModelVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

