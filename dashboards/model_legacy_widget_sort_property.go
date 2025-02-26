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

// LegacyWidgetSortProperty Determines the card sorting criterion.
type LegacyWidgetSortProperty string

// List of LegacyWidgetSortProperty
const (
	LEGACYWIDGETSORTPROPERTY_ALPHABETICAL LegacyWidgetSortProperty = "Alphabetical"
	LEGACYWIDGETSORTPROPERTY_VALUE LegacyWidgetSortProperty = "Value"
)

// All allowed values of LegacyWidgetSortProperty enum
var AllowedLegacyWidgetSortPropertyEnumValues = []LegacyWidgetSortProperty{
	"Alphabetical",
	"Value",
}

func (v *LegacyWidgetSortProperty) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LegacyWidgetSortProperty(value)
	for _, existing := range AllowedLegacyWidgetSortPropertyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LegacyWidgetSortProperty", value)
}

// NewLegacyWidgetSortPropertyFromValue returns a pointer to a valid LegacyWidgetSortProperty
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLegacyWidgetSortPropertyFromValue(v string) (*LegacyWidgetSortProperty, error) {
	ev := LegacyWidgetSortProperty(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LegacyWidgetSortProperty: valid values are %v", v, AllowedLegacyWidgetSortPropertyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LegacyWidgetSortProperty) IsValid() bool {
	for _, existing := range AllowedLegacyWidgetSortPropertyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LegacyWidgetSortProperty value
func (v LegacyWidgetSortProperty) Ptr() *LegacyWidgetSortProperty {
	return &v
}

type NullableLegacyWidgetSortProperty struct {
	value *LegacyWidgetSortProperty
	isSet bool
}

func (v NullableLegacyWidgetSortProperty) Get() *LegacyWidgetSortProperty {
	return v.value
}

func (v *NullableLegacyWidgetSortProperty) Set(val *LegacyWidgetSortProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyWidgetSortProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyWidgetSortProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyWidgetSortProperty(val *LegacyWidgetSortProperty) *NullableLegacyWidgetSortProperty {
	return &NullableLegacyWidgetSortProperty{value: val, isSet: true}
}

func (v NullableLegacyWidgetSortProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyWidgetSortProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

