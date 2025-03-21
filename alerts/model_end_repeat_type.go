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

// EndRepeatType End repeat options type.
type EndRepeatType string

// List of EndRepeatType
const (
	ENDREPEATTYPE_COUNT EndRepeatType = "count"
	ENDREPEATTYPE_NEVER EndRepeatType = "never"
	ENDREPEATTYPE_DATE EndRepeatType = "date"
)

// All allowed values of EndRepeatType enum
var AllowedEndRepeatTypeEnumValues = []EndRepeatType{
	"count",
	"never",
	"date",
}

func (v *EndRepeatType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EndRepeatType(value)
	for _, existing := range AllowedEndRepeatTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EndRepeatType", value)
}

// NewEndRepeatTypeFromValue returns a pointer to a valid EndRepeatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEndRepeatTypeFromValue(v string) (*EndRepeatType, error) {
	ev := EndRepeatType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EndRepeatType: valid values are %v", v, AllowedEndRepeatTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EndRepeatType) IsValid() bool {
	for _, existing := range AllowedEndRepeatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EndRepeatType value
func (v EndRepeatType) Ptr() *EndRepeatType {
	return &v
}

type NullableEndRepeatType struct {
	value *EndRepeatType
	isSet bool
}

func (v NullableEndRepeatType) Get() *EndRepeatType {
	return v.value
}

func (v *NullableEndRepeatType) Set(val *EndRepeatType) {
	v.value = val
	v.isSet = true
}

func (v NullableEndRepeatType) IsSet() bool {
	return v.isSet
}

func (v *NullableEndRepeatType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndRepeatType(val *EndRepeatType) *NullableEndRepeatType {
	return &NullableEndRepeatType{value: val, isSet: true}
}

func (v NullableEndRepeatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndRepeatType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

