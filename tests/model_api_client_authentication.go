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

// ApiClientAuthentication The OAuth2 client authentication location type.
type ApiClientAuthentication string

// List of ApiClientAuthentication
const (
	APICLIENTAUTHENTICATION_BASIC_AUTH_HEADER ApiClientAuthentication = "basic-auth-header"
	APICLIENTAUTHENTICATION_IN_BODY ApiClientAuthentication = "in-body"
)

// All allowed values of ApiClientAuthentication enum
var AllowedApiClientAuthenticationEnumValues = []ApiClientAuthentication{
	"basic-auth-header",
	"in-body",
}

func (v *ApiClientAuthentication) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiClientAuthentication(value)
	for _, existing := range AllowedApiClientAuthenticationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiClientAuthentication", value)
}

// NewApiClientAuthenticationFromValue returns a pointer to a valid ApiClientAuthentication
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiClientAuthenticationFromValue(v string) (*ApiClientAuthentication, error) {
	ev := ApiClientAuthentication(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiClientAuthentication: valid values are %v", v, AllowedApiClientAuthenticationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiClientAuthentication) IsValid() bool {
	for _, existing := range AllowedApiClientAuthenticationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiClientAuthentication value
func (v ApiClientAuthentication) Ptr() *ApiClientAuthentication {
	return &v
}

type NullableApiClientAuthentication struct {
	value *ApiClientAuthentication
	isSet bool
}

func (v NullableApiClientAuthentication) Get() *ApiClientAuthentication {
	return v.value
}

func (v *NullableApiClientAuthentication) Set(val *ApiClientAuthentication) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClientAuthentication) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClientAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClientAuthentication(val *ApiClientAuthentication) *NullableApiClientAuthentication {
	return &NullableApiClientAuthentication{value: val, isSet: true}
}

func (v NullableApiClientAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClientAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

