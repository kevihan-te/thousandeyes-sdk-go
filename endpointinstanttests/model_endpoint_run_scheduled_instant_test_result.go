/*
Endpoint Instant Scheduled Tests API

 You can create and execute a new endpoint instant scheduled test within ThousandEyes using this API. The test parameters are specified in the `POST` data.  The following applies to the Endpoint Instant Scheduled Tests API:  * To initiate the creation and execution of an instant scheduled test, the user must possess the `Edit endpoint tests` permission.  * Upon successful creation of an instant scheduled test, the API responds with an HTTP/201 CREATED status code and return the test definition. * It's important to note that the response does not include the results of the instant scheduled test. To retrieve test results, users can utilize the Endpoint Test Data endpoints. The URLs for these API test data endpoints are provided within the test definition output when an instant scheduled test is created. 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointinstanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the EndpointRunScheduledInstantTestResult type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointRunScheduledInstantTestResult{}

// EndpointRunScheduledInstantTestResult struct for EndpointRunScheduledInstantTestResult
type EndpointRunScheduledInstantTestResult struct {
	Message *string `json:"message,omitempty"`
}

// NewEndpointRunScheduledInstantTestResult instantiates a new EndpointRunScheduledInstantTestResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointRunScheduledInstantTestResult() *EndpointRunScheduledInstantTestResult {
	this := EndpointRunScheduledInstantTestResult{}
	return &this
}

// NewEndpointRunScheduledInstantTestResultWithDefaults instantiates a new EndpointRunScheduledInstantTestResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointRunScheduledInstantTestResultWithDefaults() *EndpointRunScheduledInstantTestResult {
	this := EndpointRunScheduledInstantTestResult{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *EndpointRunScheduledInstantTestResult) GetMessage() string {
	if o == nil || utils.IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointRunScheduledInstantTestResult) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *EndpointRunScheduledInstantTestResult) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *EndpointRunScheduledInstantTestResult) SetMessage(v string) {
	o.Message = &v
}

func (o EndpointRunScheduledInstantTestResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointRunScheduledInstantTestResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableEndpointRunScheduledInstantTestResult struct {
	value *EndpointRunScheduledInstantTestResult
	isSet bool
}

func (v NullableEndpointRunScheduledInstantTestResult) Get() *EndpointRunScheduledInstantTestResult {
	return v.value
}

func (v *NullableEndpointRunScheduledInstantTestResult) Set(val *EndpointRunScheduledInstantTestResult) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointRunScheduledInstantTestResult) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointRunScheduledInstantTestResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointRunScheduledInstantTestResult(val *EndpointRunScheduledInstantTestResult) *NullableEndpointRunScheduledInstantTestResult {
	return &NullableEndpointRunScheduledInstantTestResult{value: val, isSet: true}
}

func (v NullableEndpointRunScheduledInstantTestResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointRunScheduledInstantTestResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


