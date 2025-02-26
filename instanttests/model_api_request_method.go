/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
	"fmt"
)

// ApiRequestMethod HTTP request method.
type ApiRequestMethod string

// List of ApiRequestMethod
const (
	APIREQUESTMETHOD_GET ApiRequestMethod = "get"
	APIREQUESTMETHOD_POST ApiRequestMethod = "post"
	APIREQUESTMETHOD_PUT ApiRequestMethod = "put"
	APIREQUESTMETHOD_DELETE ApiRequestMethod = "delete"
	APIREQUESTMETHOD_PATCH ApiRequestMethod = "patch"
)

// All allowed values of ApiRequestMethod enum
var AllowedApiRequestMethodEnumValues = []ApiRequestMethod{
	"get",
	"post",
	"put",
	"delete",
	"patch",
}

func (v *ApiRequestMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiRequestMethod(value)
	for _, existing := range AllowedApiRequestMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiRequestMethod", value)
}

// NewApiRequestMethodFromValue returns a pointer to a valid ApiRequestMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiRequestMethodFromValue(v string) (*ApiRequestMethod, error) {
	ev := ApiRequestMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiRequestMethod: valid values are %v", v, AllowedApiRequestMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiRequestMethod) IsValid() bool {
	for _, existing := range AllowedApiRequestMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ApiRequestMethod value
func (v ApiRequestMethod) Ptr() *ApiRequestMethod {
	return &v
}

type NullableApiRequestMethod struct {
	value *ApiRequestMethod
	isSet bool
}

func (v NullableApiRequestMethod) Get() *ApiRequestMethod {
	return v.value
}

func (v *NullableApiRequestMethod) Set(val *ApiRequestMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableApiRequestMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableApiRequestMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiRequestMethod(val *ApiRequestMethod) *NullableApiRequestMethod {
	return &NullableApiRequestMethod{value: val, isSet: true}
}

func (v NullableApiRequestMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiRequestMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

