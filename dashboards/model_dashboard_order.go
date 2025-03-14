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

// DashboardOrder the model 'DashboardOrder'
type DashboardOrder string

// List of DashboardOrder
const (
	DASHBOARDORDER_ASC DashboardOrder = "asc"
	DASHBOARDORDER_DESC DashboardOrder = "desc"
)

// All allowed values of DashboardOrder enum
var AllowedDashboardOrderEnumValues = []DashboardOrder{
	"asc",
	"desc",
}

func (v *DashboardOrder) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DashboardOrder(value)
	for _, existing := range AllowedDashboardOrderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DashboardOrder", value)
}

// NewDashboardOrderFromValue returns a pointer to a valid DashboardOrder
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDashboardOrderFromValue(v string) (*DashboardOrder, error) {
	ev := DashboardOrder(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DashboardOrder: valid values are %v", v, AllowedDashboardOrderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DashboardOrder) IsValid() bool {
	for _, existing := range AllowedDashboardOrderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DashboardOrder value
func (v DashboardOrder) Ptr() *DashboardOrder {
	return &v
}

type NullableDashboardOrder struct {
	value *DashboardOrder
	isSet bool
}

func (v NullableDashboardOrder) Get() *DashboardOrder {
	return v.value
}

func (v *NullableDashboardOrder) Set(val *DashboardOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardOrder(val *DashboardOrder) *NullableDashboardOrder {
	return &NullableDashboardOrder{value: val, isSet: true}
}

func (v NullableDashboardOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

