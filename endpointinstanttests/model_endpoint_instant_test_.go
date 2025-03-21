/*
Endpoint Instant Scheduled Tests API

 You can create and execute a new endpoint instant scheduled test within ThousandEyes using this API. The test parameters are specified in the `POST` data.  The following applies to the Endpoint Instant Scheduled Tests API:  * To initiate the creation and execution of an instant scheduled test, the user must possess the `Edit endpoint tests` permission.  * Upon successful creation of an instant scheduled test, the API responds with an HTTP/201 CREATED status code and return the test definition. * It's important to note that the response does not include the results of the instant scheduled test. To retrieve test results, users can utilize the Endpoint Test Data endpoints. The URLs for these API test data endpoints are provided within the test definition output when an instant scheduled test is created. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointinstanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the EndpointInstantTest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointInstantTest{}

// EndpointInstantTest struct for EndpointInstantTest
type EndpointInstantTest struct {
	AgentSelectorType *EndpointTestAgentSelectorType `json:"agentSelectorType,omitempty"`
	// List of endpoint agent IDs (obtained from `/endpoint/agents` endpoint). Required when `agentSelectorType` is set to `specific-agent`.
	Agents []string `json:"agents,omitempty"`
	// List of endpoint agent label IDs (obtained from `/endpoint/labels` endpoint), required when `agentSelectorType` is set to `agent-labels`.
	EndpointAgentLabels []string `json:"endpointAgentLabels,omitempty"`
	// Maximum number of agents which can execute the test.
	MaxMachines *int32 `json:"maxMachines,omitempty"`
	// Name of the test.
	TestName string `json:"testName"`
}

type _EndpointInstantTest EndpointInstantTest

// NewEndpointInstantTest instantiates a new EndpointInstantTest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointInstantTest(testName string) *EndpointInstantTest {
	this := EndpointInstantTest{}
	var agentSelectorType EndpointTestAgentSelectorType = ENDPOINTTESTAGENTSELECTORTYPE_ALL_AGENTS
	this.AgentSelectorType = &agentSelectorType
	var maxMachines int32 = 25
	this.MaxMachines = &maxMachines
	this.TestName = testName
	return &this
}

// NewEndpointInstantTestWithDefaults instantiates a new EndpointInstantTest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointInstantTestWithDefaults() *EndpointInstantTest {
	this := EndpointInstantTest{}
	var agentSelectorType EndpointTestAgentSelectorType = ENDPOINTTESTAGENTSELECTORTYPE_ALL_AGENTS
	this.AgentSelectorType = &agentSelectorType
	var maxMachines int32 = 25
	this.MaxMachines = &maxMachines
	return &this
}

// GetAgentSelectorType returns the AgentSelectorType field value if set, zero value otherwise.
func (o *EndpointInstantTest) GetAgentSelectorType() EndpointTestAgentSelectorType {
	if o == nil || utils.IsNil(o.AgentSelectorType) {
		var ret EndpointTestAgentSelectorType
		return ret
	}
	return *o.AgentSelectorType
}

// GetAgentSelectorTypeOk returns a tuple with the AgentSelectorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointInstantTest) GetAgentSelectorTypeOk() (*EndpointTestAgentSelectorType, bool) {
	if o == nil || utils.IsNil(o.AgentSelectorType) {
		return nil, false
	}
	return o.AgentSelectorType, true
}

// HasAgentSelectorType returns a boolean if a field has been set.
func (o *EndpointInstantTest) HasAgentSelectorType() bool {
	if o != nil && !utils.IsNil(o.AgentSelectorType) {
		return true
	}

	return false
}

// SetAgentSelectorType gets a reference to the given EndpointTestAgentSelectorType and assigns it to the AgentSelectorType field.
func (o *EndpointInstantTest) SetAgentSelectorType(v EndpointTestAgentSelectorType) {
	o.AgentSelectorType = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *EndpointInstantTest) GetAgents() []string {
	if o == nil || utils.IsNil(o.Agents) {
		var ret []string
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointInstantTest) GetAgentsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *EndpointInstantTest) HasAgents() bool {
	if o != nil && !utils.IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []string and assigns it to the Agents field.
func (o *EndpointInstantTest) SetAgents(v []string) {
	o.Agents = v
}

// GetEndpointAgentLabels returns the EndpointAgentLabels field value if set, zero value otherwise.
func (o *EndpointInstantTest) GetEndpointAgentLabels() []string {
	if o == nil || utils.IsNil(o.EndpointAgentLabels) {
		var ret []string
		return ret
	}
	return o.EndpointAgentLabels
}

// GetEndpointAgentLabelsOk returns a tuple with the EndpointAgentLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointInstantTest) GetEndpointAgentLabelsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EndpointAgentLabels) {
		return nil, false
	}
	return o.EndpointAgentLabels, true
}

// HasEndpointAgentLabels returns a boolean if a field has been set.
func (o *EndpointInstantTest) HasEndpointAgentLabels() bool {
	if o != nil && !utils.IsNil(o.EndpointAgentLabels) {
		return true
	}

	return false
}

// SetEndpointAgentLabels gets a reference to the given []string and assigns it to the EndpointAgentLabels field.
func (o *EndpointInstantTest) SetEndpointAgentLabels(v []string) {
	o.EndpointAgentLabels = v
}

// GetMaxMachines returns the MaxMachines field value if set, zero value otherwise.
func (o *EndpointInstantTest) GetMaxMachines() int32 {
	if o == nil || utils.IsNil(o.MaxMachines) {
		var ret int32
		return ret
	}
	return *o.MaxMachines
}

// GetMaxMachinesOk returns a tuple with the MaxMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointInstantTest) GetMaxMachinesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MaxMachines) {
		return nil, false
	}
	return o.MaxMachines, true
}

// HasMaxMachines returns a boolean if a field has been set.
func (o *EndpointInstantTest) HasMaxMachines() bool {
	if o != nil && !utils.IsNil(o.MaxMachines) {
		return true
	}

	return false
}

// SetMaxMachines gets a reference to the given int32 and assigns it to the MaxMachines field.
func (o *EndpointInstantTest) SetMaxMachines(v int32) {
	o.MaxMachines = &v
}

// GetTestName returns the TestName field value
func (o *EndpointInstantTest) GetTestName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value
// and a boolean to check if the value has been set.
func (o *EndpointInstantTest) GetTestNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestName, true
}

// SetTestName sets field value
func (o *EndpointInstantTest) SetTestName(v string) {
	o.TestName = v
}

func (o EndpointInstantTest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointInstantTest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AgentSelectorType) {
		toSerialize["agentSelectorType"] = o.AgentSelectorType
	}
	if !utils.IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !utils.IsNil(o.EndpointAgentLabels) {
		toSerialize["endpointAgentLabels"] = o.EndpointAgentLabels
	}
	if !utils.IsNil(o.MaxMachines) {
		toSerialize["maxMachines"] = o.MaxMachines
	}
	toSerialize["testName"] = o.TestName
	return toSerialize, nil
}

func (o *EndpointInstantTest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"testName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEndpointInstantTest := _EndpointInstantTest{}

    err = json.Unmarshal(data, &varEndpointInstantTest)

	if err != nil {
		return err
	}

	*o = EndpointInstantTest(varEndpointInstantTest)

	return err
}

type NullableEndpointInstantTest struct {
	value *EndpointInstantTest
	isSet bool
}

func (v NullableEndpointInstantTest) Get() *EndpointInstantTest {
	return v.value
}

func (v *NullableEndpointInstantTest) Set(val *EndpointInstantTest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointInstantTest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointInstantTest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointInstantTest(val *EndpointInstantTest) *NullableEndpointInstantTest {
	return &NullableEndpointInstantTest{value: val, isSet: true}
}

func (v NullableEndpointInstantTest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointInstantTest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


