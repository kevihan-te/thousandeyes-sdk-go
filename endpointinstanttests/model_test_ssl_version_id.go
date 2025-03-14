/*
Endpoint Instant Scheduled Tests API

 You can create and execute a new endpoint instant scheduled test within ThousandEyes using this API. The test parameters are specified in the `POST` data.  The following applies to the Endpoint Instant Scheduled Tests API:  * To initiate the creation and execution of an instant scheduled test, the user must possess the `Edit endpoint tests` permission.  * Upon successful creation of an instant scheduled test, the API responds with an HTTP/201 CREATED status code and return the test definition. * It's important to note that the response does not include the results of the instant scheduled test. To retrieve test results, users can utilize the Endpoint Test Data endpoints. The URLs for these API test data endpoints are provided within the test definition output when an instant scheduled test is created. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointinstanttests

import (
	"encoding/json"
	"fmt"
)

// TestSslVersionId SSL version options:  * Use '0' for automatic selection. * Use '3' for SSLv3. * Use '4' for TLS v1.0. * Use '5' for TLS v1.1. * Use '6' for TLS v1.2. 
type TestSslVersionId string

// List of TestSslVersionId
const (
	TESTSSLVERSIONID__0 TestSslVersionId = "0"
	TESTSSLVERSIONID__3 TestSslVersionId = "3"
	TESTSSLVERSIONID__4 TestSslVersionId = "4"
	TESTSSLVERSIONID__5 TestSslVersionId = "5"
	TESTSSLVERSIONID__6 TestSslVersionId = "6"
)

// All allowed values of TestSslVersionId enum
var AllowedTestSslVersionIdEnumValues = []TestSslVersionId{
	"0",
	"3",
	"4",
	"5",
	"6",
}

func (v *TestSslVersionId) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TestSslVersionId(value)
	for _, existing := range AllowedTestSslVersionIdEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TestSslVersionId", value)
}

// NewTestSslVersionIdFromValue returns a pointer to a valid TestSslVersionId
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTestSslVersionIdFromValue(v string) (*TestSslVersionId, error) {
	ev := TestSslVersionId(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TestSslVersionId: valid values are %v", v, AllowedTestSslVersionIdEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TestSslVersionId) IsValid() bool {
	for _, existing := range AllowedTestSslVersionIdEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TestSslVersionId value
func (v TestSslVersionId) Ptr() *TestSslVersionId {
	return &v
}

type NullableTestSslVersionId struct {
	value *TestSslVersionId
	isSet bool
}

func (v NullableTestSslVersionId) Get() *TestSslVersionId {
	return v.value
}

func (v *NullableTestSslVersionId) Set(val *TestSslVersionId) {
	v.value = val
	v.isSet = true
}

func (v NullableTestSslVersionId) IsSet() bool {
	return v.isSet
}

func (v *NullableTestSslVersionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestSslVersionId(val *TestSslVersionId) *NullableTestSslVersionId {
	return &NullableTestSslVersionId{value: val, isSet: true}
}

func (v NullableTestSslVersionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestSslVersionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

