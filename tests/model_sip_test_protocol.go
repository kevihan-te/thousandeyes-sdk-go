/*
Tests API

This API supports listing, creating, editing, and deleting Cloud and Enterprise Agent (CEA) based tests. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tests

import (
	"encoding/json"
	"fmt"
)

// SipTestProtocol Transport layer for SIP communication.
type SipTestProtocol string

// List of SipTestProtocol
const (
	SIPTESTPROTOCOL_TCP SipTestProtocol = "tcp"
	SIPTESTPROTOCOL_TLS SipTestProtocol = "tls"
	SIPTESTPROTOCOL_UDP SipTestProtocol = "udp"
)

// All allowed values of SipTestProtocol enum
var AllowedSipTestProtocolEnumValues = []SipTestProtocol{
	"tcp",
	"tls",
	"udp",
}

func (v *SipTestProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SipTestProtocol(value)
	for _, existing := range AllowedSipTestProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SipTestProtocol", value)
}

// NewSipTestProtocolFromValue returns a pointer to a valid SipTestProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSipTestProtocolFromValue(v string) (*SipTestProtocol, error) {
	ev := SipTestProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SipTestProtocol: valid values are %v", v, AllowedSipTestProtocolEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SipTestProtocol) IsValid() bool {
	for _, existing := range AllowedSipTestProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SipTestProtocol value
func (v SipTestProtocol) Ptr() *SipTestProtocol {
	return &v
}

type NullableSipTestProtocol struct {
	value *SipTestProtocol
	isSet bool
}

func (v NullableSipTestProtocol) Get() *SipTestProtocol {
	return v.value
}

func (v *NullableSipTestProtocol) Set(val *SipTestProtocol) {
	v.value = val
	v.isSet = true
}

func (v NullableSipTestProtocol) IsSet() bool {
	return v.isSet
}

func (v *NullableSipTestProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipTestProtocol(val *SipTestProtocol) *NullableSipTestProtocol {
	return &NullableSipTestProtocol{value: val, isSet: true}
}

func (v NullableSipTestProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipTestProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

