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

// AlertRoundsViolationMode `exact` requires the same agents to meet the threshold in consecutive rounds. `auto` is only enabled for CEA and Endpoint Scheduled test rules. The default is `any`.
type AlertRoundsViolationMode string

// List of AlertRoundsViolationMode
const (
	ALERTROUNDSVIOLATIONMODE_EXACT AlertRoundsViolationMode = "exact"
	ALERTROUNDSVIOLATIONMODE_ANY AlertRoundsViolationMode = "any"
	ALERTROUNDSVIOLATIONMODE_AUTO AlertRoundsViolationMode = "auto"
)

// All allowed values of AlertRoundsViolationMode enum
var AllowedAlertRoundsViolationModeEnumValues = []AlertRoundsViolationMode{
	"exact",
	"any",
	"auto",
}

func (v *AlertRoundsViolationMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlertRoundsViolationMode(value)
	for _, existing := range AllowedAlertRoundsViolationModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlertRoundsViolationMode", value)
}

// NewAlertRoundsViolationModeFromValue returns a pointer to a valid AlertRoundsViolationMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlertRoundsViolationModeFromValue(v string) (*AlertRoundsViolationMode, error) {
	ev := AlertRoundsViolationMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlertRoundsViolationMode: valid values are %v", v, AllowedAlertRoundsViolationModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlertRoundsViolationMode) IsValid() bool {
	for _, existing := range AllowedAlertRoundsViolationModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlertRoundsViolationMode value
func (v AlertRoundsViolationMode) Ptr() *AlertRoundsViolationMode {
	return &v
}

type NullableAlertRoundsViolationMode struct {
	value *AlertRoundsViolationMode
	isSet bool
}

func (v NullableAlertRoundsViolationMode) Get() *AlertRoundsViolationMode {
	return v.value
}

func (v *NullableAlertRoundsViolationMode) Set(val *AlertRoundsViolationMode) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertRoundsViolationMode) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertRoundsViolationMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertRoundsViolationMode(val *AlertRoundsViolationMode) *NullableAlertRoundsViolationMode {
	return &NullableAlertRoundsViolationMode{value: val, isSet: true}
}

func (v NullableAlertRoundsViolationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertRoundsViolationMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

