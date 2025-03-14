/*
Alerts API

You can manage the following alert functionalities on the ThousandEyes platform using the Alerts API:  * **Alerts**: Retrieve alert details. Alerts are assigned to tests through alert rules.  * **Alert Rules**: Conditions that you configure in order to highlight or be notified of events of interest in your ThousandEyes tests. When an alert rule’s conditions are met, the associated alert is triggered and the alert becomes active. It remains active until the alert is cleared. Alert rules are reusable across multiple tests..  * **Alert Suppression Windows**: Suppress alerts for tests during periods such as planned maintenance. Windows can be one-time events or recurring events to handle periodic occurrences such as monthly downtime for maintenance.  For more information about the alerts, see [Alerts](https://docs.thousandeyes.com/product-documentation/alerts). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerts

import (
	"encoding/json"
	"fmt"
)

// DaysOfWeek Specifies the day to activate the alert suppression window. Applicable only when `intervalType` is set to `week`.
type DaysOfWeek string

// List of DaysOfWeek
const (
	DAYSOFWEEK_SUN DaysOfWeek = "sun"
	DAYSOFWEEK_MON DaysOfWeek = "mon"
	DAYSOFWEEK_TUE DaysOfWeek = "tue"
	DAYSOFWEEK_WED DaysOfWeek = "wed"
	DAYSOFWEEK_THU DaysOfWeek = "thu"
	DAYSOFWEEK_FRI DaysOfWeek = "fri"
	DAYSOFWEEK_SAT DaysOfWeek = "sat"
)

// All allowed values of DaysOfWeek enum
var AllowedDaysOfWeekEnumValues = []DaysOfWeek{
	"sun",
	"mon",
	"tue",
	"wed",
	"thu",
	"fri",
	"sat",
}

func (v *DaysOfWeek) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DaysOfWeek(value)
	for _, existing := range AllowedDaysOfWeekEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DaysOfWeek", value)
}

// NewDaysOfWeekFromValue returns a pointer to a valid DaysOfWeek
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDaysOfWeekFromValue(v string) (*DaysOfWeek, error) {
	ev := DaysOfWeek(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DaysOfWeek: valid values are %v", v, AllowedDaysOfWeekEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DaysOfWeek) IsValid() bool {
	for _, existing := range AllowedDaysOfWeekEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DaysOfWeek value
func (v DaysOfWeek) Ptr() *DaysOfWeek {
	return &v
}

type NullableDaysOfWeek struct {
	value *DaysOfWeek
	isSet bool
}

func (v NullableDaysOfWeek) Get() *DaysOfWeek {
	return v.value
}

func (v *NullableDaysOfWeek) Set(val *DaysOfWeek) {
	v.value = val
	v.isSet = true
}

func (v NullableDaysOfWeek) IsSet() bool {
	return v.isSet
}

func (v *NullableDaysOfWeek) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaysOfWeek(val *DaysOfWeek) *NullableDaysOfWeek {
	return &NullableDaysOfWeek{value: val, isSet: true}
}

func (v NullableDaysOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaysOfWeek) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

