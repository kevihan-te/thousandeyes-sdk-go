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

// RequestMethod HTTP request method.
type RequestMethod string

// List of RequestMethod
const (
	REQUESTMETHOD_GET RequestMethod = "get"
	REQUESTMETHOD_POST RequestMethod = "post"
	REQUESTMETHOD_PUT RequestMethod = "put"
	REQUESTMETHOD_DELETE RequestMethod = "delete"
	REQUESTMETHOD_PATCH RequestMethod = "patch"
	REQUESTMETHOD_OPTIONS RequestMethod = "options"
	REQUESTMETHOD_TRACE RequestMethod = "trace"
)

// All allowed values of RequestMethod enum
var AllowedRequestMethodEnumValues = []RequestMethod{
	"get",
	"post",
	"put",
	"delete",
	"patch",
	"options",
	"trace",
}

func (v *RequestMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestMethod(value)
	for _, existing := range AllowedRequestMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestMethod", value)
}

// NewRequestMethodFromValue returns a pointer to a valid RequestMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestMethodFromValue(v string) (*RequestMethod, error) {
	ev := RequestMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestMethod: valid values are %v", v, AllowedRequestMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestMethod) IsValid() bool {
	for _, existing := range AllowedRequestMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestMethod value
func (v RequestMethod) Ptr() *RequestMethod {
	return &v
}

type NullableRequestMethod struct {
	value *RequestMethod
	isSet bool
}

func (v NullableRequestMethod) Get() *RequestMethod {
	return v.value
}

func (v *NullableRequestMethod) Set(val *RequestMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestMethod(val *RequestMethod) *NullableRequestMethod {
	return &NullableRequestMethod{value: val, isSet: true}
}

func (v NullableRequestMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

