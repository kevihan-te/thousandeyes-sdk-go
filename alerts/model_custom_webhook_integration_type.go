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

// CustomWebhookIntegrationType Integration type.
type CustomWebhookIntegrationType string

// List of CustomWebhookIntegrationType
const (
	CUSTOMWEBHOOKINTEGRATIONTYPE_CUSTOM_WEBHOOK CustomWebhookIntegrationType = "custom-webhook"
)

// All allowed values of CustomWebhookIntegrationType enum
var AllowedCustomWebhookIntegrationTypeEnumValues = []CustomWebhookIntegrationType{
	"custom-webhook",
}

func (v *CustomWebhookIntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomWebhookIntegrationType(value)
	for _, existing := range AllowedCustomWebhookIntegrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomWebhookIntegrationType", value)
}

// NewCustomWebhookIntegrationTypeFromValue returns a pointer to a valid CustomWebhookIntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomWebhookIntegrationTypeFromValue(v string) (*CustomWebhookIntegrationType, error) {
	ev := CustomWebhookIntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomWebhookIntegrationType: valid values are %v", v, AllowedCustomWebhookIntegrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomWebhookIntegrationType) IsValid() bool {
	for _, existing := range AllowedCustomWebhookIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomWebhookIntegrationType value
func (v CustomWebhookIntegrationType) Ptr() *CustomWebhookIntegrationType {
	return &v
}

type NullableCustomWebhookIntegrationType struct {
	value *CustomWebhookIntegrationType
	isSet bool
}

func (v NullableCustomWebhookIntegrationType) Get() *CustomWebhookIntegrationType {
	return v.value
}

func (v *NullableCustomWebhookIntegrationType) Set(val *CustomWebhookIntegrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomWebhookIntegrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomWebhookIntegrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomWebhookIntegrationType(val *CustomWebhookIntegrationType) *NullableCustomWebhookIntegrationType {
	return &NullableCustomWebhookIntegrationType{value: val, isSet: true}
}

func (v NullableCustomWebhookIntegrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomWebhookIntegrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

