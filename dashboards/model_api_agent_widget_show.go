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

// ApiAgentWidgetShow Ownership of the agent.
type ApiAgentWidgetShow string

// List of ApiAgentWidgetShow
const (
	APIAGENTWIDGETSHOW_OWNED ApiAgentWidgetShow = "owned"
	APIAGENTWIDGETSHOW_ALL ApiAgentWidgetShow = "all"
)

// All allowed values of ApiAgentWidgetShow enum
var AllowedApiAgentWidgetShowEnumValues = []ApiAgentWidgetShow{
	"owned",
	"all",
}

func (v *ApiAgentWidgetShow) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiAgentWidgetShow(value)
	for _, existing := range AllowedApiAgentWidgetShowEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiAgentWidgetShow", value)
}

// NewApiAgentWidgetShowFromValue returns a pointer to a valid ApiAgentWidgetShow
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiAgentWidgetShowFromValue(v string) (*ApiAgentWidgetShow, error) {
	ev := ApiAgentWidgetShow(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiAgentWidgetShow: valid values are %v", v, AllowedApiAgentWidgetShowEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiAgentWidgetShow) IsValid() bool {
	for _, existing := range AllowedApiAgentWidgetShowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiAgentWidgetShow value
func (v ApiAgentWidgetShow) Ptr() *ApiAgentWidgetShow {
	return &v
}

type NullableApiAgentWidgetShow struct {
	value *ApiAgentWidgetShow
	isSet bool
}

func (v NullableApiAgentWidgetShow) Get() *ApiAgentWidgetShow {
	return v.value
}

func (v *NullableApiAgentWidgetShow) Set(val *ApiAgentWidgetShow) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAgentWidgetShow) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAgentWidgetShow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAgentWidgetShow(val *ApiAgentWidgetShow) *NullableApiAgentWidgetShow {
	return &NullableApiAgentWidgetShow{value: val, isSet: true}
}

func (v NullableApiAgentWidgetShow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAgentWidgetShow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

