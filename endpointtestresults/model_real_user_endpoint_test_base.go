/*
Endpoint Test Results API

Retrieve results for scheduled and dynamic tests on endpoint agents.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointtestresults

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the RealUserEndpointTestBase type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RealUserEndpointTestBase{}

// RealUserEndpointTestBase struct for RealUserEndpointTestBase
type RealUserEndpointTestBase struct {
	// Unique ID of endpoint agent, from `/endpoint/agents` endpoint.
	AgentId *string `json:"agentId,omitempty"`
	// UTC date when endpoint real user test was committed to the controller (ISO date-time format).
	Committed *time.Time `json:"committed,omitempty"`
	// UTC date when endpoint real user test took place (ISO date-time format).
	Date *time.Time `json:"date,omitempty"`
	// Score rating a user’s experience when loading a particular page, from 0 (the worst) to 1 (the best). [More details](https://docs.thousandeyes.com/product-documentation/end-user-monitoring/viewing-data/endpoint-agent-views-reference#user-experience-score).
	ExperienceScore *float64 `json:"experienceScore,omitempty"`
	// Number of web pages visited on target website.
	NumberOfPages *int32 `json:"numberOfPages,omitempty"`
	// Name of the AS organization `sourceAddress` belongs to.
	OrganizationName *string `json:"organizationName,omitempty"`
	// Port used to visit target website.
	Port *int32 `json:"port,omitempty"`
	// Protocol used to visit target website.
	Protocol *string `json:"protocol,omitempty"`
	// Epoch time (seconds) indicating the start time of the round.
	RoundId *int32 `json:"roundId,omitempty"`
	// Public IP address of the endpoint agent during the session.
	SourceAddress *string `json:"sourceAddress,omitempty"`
	// Endpoint real user test ID. Each endpoint real user test occurrence has a unique ID.
	Id *string `json:"id,omitempty"`
	// Domain used to visit target website.
	VisitedSite *string `json:"visitedSite,omitempty"`
}

// NewRealUserEndpointTestBase instantiates a new RealUserEndpointTestBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealUserEndpointTestBase() *RealUserEndpointTestBase {
	this := RealUserEndpointTestBase{}
	return &this
}

// NewRealUserEndpointTestBaseWithDefaults instantiates a new RealUserEndpointTestBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealUserEndpointTestBaseWithDefaults() *RealUserEndpointTestBase {
	this := RealUserEndpointTestBase{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetAgentId() string {
	if o == nil || utils.IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetAgentIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasAgentId() bool {
	if o != nil && !utils.IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *RealUserEndpointTestBase) SetAgentId(v string) {
	o.AgentId = &v
}

// GetCommitted returns the Committed field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetCommitted() time.Time {
	if o == nil || utils.IsNil(o.Committed) {
		var ret time.Time
		return ret
	}
	return *o.Committed
}

// GetCommittedOk returns a tuple with the Committed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetCommittedOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Committed) {
		return nil, false
	}
	return o.Committed, true
}

// HasCommitted returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasCommitted() bool {
	if o != nil && !utils.IsNil(o.Committed) {
		return true
	}

	return false
}

// SetCommitted gets a reference to the given time.Time and assigns it to the Committed field.
func (o *RealUserEndpointTestBase) SetCommitted(v time.Time) {
	o.Committed = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetDate() time.Time {
	if o == nil || utils.IsNil(o.Date) {
		var ret time.Time
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasDate() bool {
	if o != nil && !utils.IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given time.Time and assigns it to the Date field.
func (o *RealUserEndpointTestBase) SetDate(v time.Time) {
	o.Date = &v
}

// GetExperienceScore returns the ExperienceScore field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetExperienceScore() float64 {
	if o == nil || utils.IsNil(o.ExperienceScore) {
		var ret float64
		return ret
	}
	return *o.ExperienceScore
}

// GetExperienceScoreOk returns a tuple with the ExperienceScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetExperienceScoreOk() (*float64, bool) {
	if o == nil || utils.IsNil(o.ExperienceScore) {
		return nil, false
	}
	return o.ExperienceScore, true
}

// HasExperienceScore returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasExperienceScore() bool {
	if o != nil && !utils.IsNil(o.ExperienceScore) {
		return true
	}

	return false
}

// SetExperienceScore gets a reference to the given float64 and assigns it to the ExperienceScore field.
func (o *RealUserEndpointTestBase) SetExperienceScore(v float64) {
	o.ExperienceScore = &v
}

// GetNumberOfPages returns the NumberOfPages field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetNumberOfPages() int32 {
	if o == nil || utils.IsNil(o.NumberOfPages) {
		var ret int32
		return ret
	}
	return *o.NumberOfPages
}

// GetNumberOfPagesOk returns a tuple with the NumberOfPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetNumberOfPagesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumberOfPages) {
		return nil, false
	}
	return o.NumberOfPages, true
}

// HasNumberOfPages returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasNumberOfPages() bool {
	if o != nil && !utils.IsNil(o.NumberOfPages) {
		return true
	}

	return false
}

// SetNumberOfPages gets a reference to the given int32 and assigns it to the NumberOfPages field.
func (o *RealUserEndpointTestBase) SetNumberOfPages(v int32) {
	o.NumberOfPages = &v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetOrganizationName() string {
	if o == nil || utils.IsNil(o.OrganizationName) {
		var ret string
		return ret
	}
	return *o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetOrganizationNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OrganizationName) {
		return nil, false
	}
	return o.OrganizationName, true
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasOrganizationName() bool {
	if o != nil && !utils.IsNil(o.OrganizationName) {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given string and assigns it to the OrganizationName field.
func (o *RealUserEndpointTestBase) SetOrganizationName(v string) {
	o.OrganizationName = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetPort() int32 {
	if o == nil || utils.IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetPortOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasPort() bool {
	if o != nil && !utils.IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *RealUserEndpointTestBase) SetPort(v int32) {
	o.Port = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetProtocol() string {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetProtocolOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *RealUserEndpointTestBase) SetProtocol(v string) {
	o.Protocol = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *RealUserEndpointTestBase) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetSourceAddress returns the SourceAddress field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetSourceAddress() string {
	if o == nil || utils.IsNil(o.SourceAddress) {
		var ret string
		return ret
	}
	return *o.SourceAddress
}

// GetSourceAddressOk returns a tuple with the SourceAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetSourceAddressOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceAddress) {
		return nil, false
	}
	return o.SourceAddress, true
}

// HasSourceAddress returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasSourceAddress() bool {
	if o != nil && !utils.IsNil(o.SourceAddress) {
		return true
	}

	return false
}

// SetSourceAddress gets a reference to the given string and assigns it to the SourceAddress field.
func (o *RealUserEndpointTestBase) SetSourceAddress(v string) {
	o.SourceAddress = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RealUserEndpointTestBase) SetId(v string) {
	o.Id = &v
}

// GetVisitedSite returns the VisitedSite field value if set, zero value otherwise.
func (o *RealUserEndpointTestBase) GetVisitedSite() string {
	if o == nil || utils.IsNil(o.VisitedSite) {
		var ret string
		return ret
	}
	return *o.VisitedSite
}

// GetVisitedSiteOk returns a tuple with the VisitedSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RealUserEndpointTestBase) GetVisitedSiteOk() (*string, bool) {
	if o == nil || utils.IsNil(o.VisitedSite) {
		return nil, false
	}
	return o.VisitedSite, true
}

// HasVisitedSite returns a boolean if a field has been set.
func (o *RealUserEndpointTestBase) HasVisitedSite() bool {
	if o != nil && !utils.IsNil(o.VisitedSite) {
		return true
	}

	return false
}

// SetVisitedSite gets a reference to the given string and assigns it to the VisitedSite field.
func (o *RealUserEndpointTestBase) SetVisitedSite(v string) {
	o.VisitedSite = &v
}

func (o RealUserEndpointTestBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RealUserEndpointTestBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !utils.IsNil(o.Committed) {
		toSerialize["committed"] = o.Committed
	}
	if !utils.IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !utils.IsNil(o.ExperienceScore) {
		toSerialize["experienceScore"] = o.ExperienceScore
	}
	if !utils.IsNil(o.NumberOfPages) {
		toSerialize["numberOfPages"] = o.NumberOfPages
	}
	if !utils.IsNil(o.OrganizationName) {
		toSerialize["organizationName"] = o.OrganizationName
	}
	if !utils.IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !utils.IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !utils.IsNil(o.RoundId) {
		toSerialize["roundId"] = o.RoundId
	}
	if !utils.IsNil(o.SourceAddress) {
		toSerialize["sourceAddress"] = o.SourceAddress
	}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.VisitedSite) {
		toSerialize["visitedSite"] = o.VisitedSite
	}
	return toSerialize, nil
}

type NullableRealUserEndpointTestBase struct {
	value *RealUserEndpointTestBase
	isSet bool
}

func (v NullableRealUserEndpointTestBase) Get() *RealUserEndpointTestBase {
	return v.value
}

func (v *NullableRealUserEndpointTestBase) Set(val *RealUserEndpointTestBase) {
	v.value = val
	v.isSet = true
}

func (v NullableRealUserEndpointTestBase) IsSet() bool {
	return v.isSet
}

func (v *NullableRealUserEndpointTestBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRealUserEndpointTestBase(val *RealUserEndpointTestBase) *NullableRealUserEndpointTestBase {
	return &NullableRealUserEndpointTestBase{value: val, isSet: true}
}

func (v NullableRealUserEndpointTestBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRealUserEndpointTestBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


