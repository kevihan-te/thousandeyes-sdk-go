/*
Test Results API

Get test result metrics for Cloud and Enterprise Agent tests.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package testresults

import (
	"encoding/json"
	"fmt"
)

// SipServerErrorType Error type, none if there is no error
type SipServerErrorType string

// List of SipServerErrorType
const (
	SIPSERVERERRORTYPE_NONE SipServerErrorType = "none"
	SIPSERVERERRORTYPE_DNS SipServerErrorType = "dns"
	SIPSERVERERRORTYPE_CONNECT SipServerErrorType = "connect"
	SIPSERVERERRORTYPE_REGISTER SipServerErrorType = "register"
	SIPSERVERERRORTYPE_INVITE SipServerErrorType = "invite"
	SIPSERVERERRORTYPE_OPTION SipServerErrorType = "option"
	SIPSERVERERRORTYPE_SERVER SipServerErrorType = "server"
	SIPSERVERERRORTYPE_SSL SipServerErrorType = "ssl"
)

// All allowed values of SipServerErrorType enum
var AllowedSipServerErrorTypeEnumValues = []SipServerErrorType{
	"none",
	"dns",
	"connect",
	"register",
	"invite",
	"option",
	"server",
	"ssl",
}

func (v *SipServerErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SipServerErrorType(value)
	for _, existing := range AllowedSipServerErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SipServerErrorType", value)
}

// NewSipServerErrorTypeFromValue returns a pointer to a valid SipServerErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSipServerErrorTypeFromValue(v string) (*SipServerErrorType, error) {
	ev := SipServerErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SipServerErrorType: valid values are %v", v, AllowedSipServerErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SipServerErrorType) IsValid() bool {
	for _, existing := range AllowedSipServerErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SipServerErrorType value
func (v SipServerErrorType) Ptr() *SipServerErrorType {
	return &v
}

type NullableSipServerErrorType struct {
	value *SipServerErrorType
	isSet bool
}

func (v NullableSipServerErrorType) Get() *SipServerErrorType {
	return v.value
}

func (v *NullableSipServerErrorType) Set(val *SipServerErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableSipServerErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableSipServerErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipServerErrorType(val *SipServerErrorType) *NullableSipServerErrorType {
	return &NullableSipServerErrorType{value: val, isSet: true}
}

func (v NullableSipServerErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipServerErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

