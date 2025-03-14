/*
Agents API

 ## Overview Manage Cloud and Enterprise Agents available to your account in ThousandEyes.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agents

import (
	"encoding/json"
	"fmt"
)

// ProxyAuthType The type of authentication the proxy requires
type ProxyAuthType string

// List of ProxyAuthType
const (
	PROXYAUTHTYPE_BASIC ProxyAuthType = "basic"
	PROXYAUTHTYPE_NTLM ProxyAuthType = "ntlm"
	PROXYAUTHTYPE_KERBEROS ProxyAuthType = "kerberos"
	PROXYAUTHTYPE_UNKNOWN ProxyAuthType = "unknown"
)

// All allowed values of ProxyAuthType enum
var AllowedProxyAuthTypeEnumValues = []ProxyAuthType{
	"basic",
	"ntlm",
	"kerberos",
	"unknown",
}

func (v *ProxyAuthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProxyAuthType(value)
	for _, existing := range AllowedProxyAuthTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProxyAuthType", value)
}

// NewProxyAuthTypeFromValue returns a pointer to a valid ProxyAuthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProxyAuthTypeFromValue(v string) (*ProxyAuthType, error) {
	ev := ProxyAuthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProxyAuthType: valid values are %v", v, AllowedProxyAuthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProxyAuthType) IsValid() bool {
	for _, existing := range AllowedProxyAuthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProxyAuthType value
func (v ProxyAuthType) Ptr() *ProxyAuthType {
	return &v
}

type NullableProxyAuthType struct {
	value *ProxyAuthType
	isSet bool
}

func (v NullableProxyAuthType) Get() *ProxyAuthType {
	return v.value
}

func (v *NullableProxyAuthType) Set(val *ProxyAuthType) {
	v.value = val
	v.isSet = true
}

func (v NullableProxyAuthType) IsSet() bool {
	return v.isSet
}

func (v *NullableProxyAuthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProxyAuthType(val *ProxyAuthType) *NullableProxyAuthType {
	return &NullableProxyAuthType{value: val, isSet: true}
}

func (v NullableProxyAuthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProxyAuthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

