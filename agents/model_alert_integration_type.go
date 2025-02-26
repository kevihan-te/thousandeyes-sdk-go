/*
Agents API

 ## Overview Manage Cloud and Enterprise Agents available to your account in ThousandEyes.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agents

import (
	"encoding/json"
	"fmt"
)

// AlertIntegrationType Type of the alert integration
type AlertIntegrationType string

// List of AlertIntegrationType
const (
	ALERTINTEGRATIONTYPE_PAGER_DUTY AlertIntegrationType = "pager-duty"
	ALERTINTEGRATIONTYPE_SLACK AlertIntegrationType = "slack"
)

// All allowed values of AlertIntegrationType enum
var AllowedAlertIntegrationTypeEnumValues = []AlertIntegrationType{
	"pager-duty",
	"slack",
}

func (v *AlertIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertIntegrationType(value)
	for _, existing := range AllowedAlertIntegrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertIntegrationType", value)
}

// NewAlertIntegrationTypeFromValue returns a pointer to a valid AlertIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertIntegrationTypeFromValue(v string) (*AlertIntegrationType, error) {
	ev := AlertIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertIntegrationType: valid values are %v", v, AllowedAlertIntegrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertIntegrationType) IsValid() bool {
	for _, existing := range AllowedAlertIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertIntegrationType value
func (v AlertIntegrationType) Ptr() *AlertIntegrationType {
	return &v
}

type NullableAlertIntegrationType struct {
	value *AlertIntegrationType
	isSet bool
}

func (v NullableAlertIntegrationType) Get() *AlertIntegrationType {
	return v.value
}

func (v *NullableAlertIntegrationType) Set(val *AlertIntegrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertIntegrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertIntegrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertIntegrationType(val *AlertIntegrationType) *NullableAlertIntegrationType {
	return &NullableAlertIntegrationType{value: val, isSet: true}
}

func (v NullableAlertIntegrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertIntegrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

