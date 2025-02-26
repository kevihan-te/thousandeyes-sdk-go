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

// LegacyWidgetSortDirection Specifies the order in which cards are sorted.
type LegacyWidgetSortDirection string

// List of LegacyWidgetSortDirection
const (
	LEGACYWIDGETSORTDIRECTION_ASCENDING LegacyWidgetSortDirection = "Ascending"
	LEGACYWIDGETSORTDIRECTION_DESCENDING LegacyWidgetSortDirection = "Descending"
)

// All allowed values of LegacyWidgetSortDirection enum
var AllowedLegacyWidgetSortDirectionEnumValues = []LegacyWidgetSortDirection{
	"Ascending",
	"Descending",
}

func (v *LegacyWidgetSortDirection) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LegacyWidgetSortDirection(value)
	for _, existing := range AllowedLegacyWidgetSortDirectionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LegacyWidgetSortDirection", value)
}

// NewLegacyWidgetSortDirectionFromValue returns a pointer to a valid LegacyWidgetSortDirection
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLegacyWidgetSortDirectionFromValue(v string) (*LegacyWidgetSortDirection, error) {
	ev := LegacyWidgetSortDirection(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LegacyWidgetSortDirection: valid values are %v", v, AllowedLegacyWidgetSortDirectionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LegacyWidgetSortDirection) IsValid() bool {
	for _, existing := range AllowedLegacyWidgetSortDirectionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LegacyWidgetSortDirection value
func (v LegacyWidgetSortDirection) Ptr() *LegacyWidgetSortDirection {
	return &v
}

type NullableLegacyWidgetSortDirection struct {
	value *LegacyWidgetSortDirection
	isSet bool
}

func (v NullableLegacyWidgetSortDirection) Get() *LegacyWidgetSortDirection {
	return v.value
}

func (v *NullableLegacyWidgetSortDirection) Set(val *LegacyWidgetSortDirection) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyWidgetSortDirection) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyWidgetSortDirection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyWidgetSortDirection(val *LegacyWidgetSortDirection) *NullableLegacyWidgetSortDirection {
	return &NullableLegacyWidgetSortDirection{value: val, isSet: true}
}

func (v NullableLegacyWidgetSortDirection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyWidgetSortDirection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

