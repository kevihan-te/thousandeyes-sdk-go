/*
Event Detection API

 Event detection occurs when ThousandEyes identifies that error signals related to a component (proxy, network node, AS, server etc) have deviated from the baselines established by events. * To determine this, ThousandEyes takes the test results from all accounts groups within an organization, and analyzes that data. * Noisy test results (those that have too many errors in a short window) are removed until they stabilize, and the rest of the results are tagged with the components associated with that test result (for example, proxy, network, or server). * Next, any increase in failures from the test results and each component helps in determining the problem domain and which component may be at fault. * When this failure rate increases beyond a pre-defined threshold (set by the algorithm), an event is triggered and an email notification is sent to the user (if they've enabled email alerts).  With the Events API, you can perform the following tasks on the ThousandEyes platform: * **Retrieve Events**: Obtain a list of events and detailed information for each event. For more information about events, see [Event Detection](https://docs.thousandeyes.com/product-documentation/event-detection). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package eventdetection

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the Event type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Event{}

// Event struct for Event
type Event struct {
	// A unique ID for each event.
	Id *string `json:"id,omitempty"`
	// Event type name. Examples include, Local Agent Issue, Network Path Issue, Network Outage, DNS Issue, Server Issue, and Proxy Issue.
	TypeName *string `json:"typeName,omitempty"`
	State *EventState `json:"state,omitempty"`
	// The start date and time (in UTC, ISO 8601 format) when the event was first detected.
	StartDate *time.Time `json:"startDate,omitempty"`
	// The end date and time (in UTC, ISO 8601 format) when the event was resolved (due to timeout). This value is populated for \"ongoing\" events.
	EndDate *time.Time `json:"endDate,omitempty"`
	Severity *EventAlertSeverity `json:"severity,omitempty"`
	// Event title
	Title *string `json:"title,omitempty"`
	Type *EventType `json:"type,omitempty"`
	AffectedTests *AffectedCount `json:"affectedTests,omitempty"`
	AffectedTargets *AffectedCount `json:"affectedTargets,omitempty"`
	AffectedAgents *AffectedCount `json:"affectedAgents,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewEvent instantiates a new Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEvent() *Event {
	this := Event{}
	return &this
}

// NewEventWithDefaults instantiates a new Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventWithDefaults() *Event {
	this := Event{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Event) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Event) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Event) SetId(v string) {
	o.Id = &v
}

// GetTypeName returns the TypeName field value if set, zero value otherwise.
func (o *Event) GetTypeName() string {
	if o == nil || utils.IsNil(o.TypeName) {
		var ret string
		return ret
	}
	return *o.TypeName
}

// GetTypeNameOk returns a tuple with the TypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTypeNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TypeName) {
		return nil, false
	}
	return o.TypeName, true
}

// HasTypeName returns a boolean if a field has been set.
func (o *Event) HasTypeName() bool {
	if o != nil && !utils.IsNil(o.TypeName) {
		return true
	}

	return false
}

// SetTypeName gets a reference to the given string and assigns it to the TypeName field.
func (o *Event) SetTypeName(v string) {
	o.TypeName = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Event) GetState() EventState {
	if o == nil || utils.IsNil(o.State) {
		var ret EventState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStateOk() (*EventState, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Event) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given EventState and assigns it to the State field.
func (o *Event) SetState(v EventState) {
	o.State = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Event) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Event) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Event) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Event) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Event) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *Event) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Event) GetSeverity() EventAlertSeverity {
	if o == nil || utils.IsNil(o.Severity) {
		var ret EventAlertSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetSeverityOk() (*EventAlertSeverity, bool) {
	if o == nil || utils.IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Event) HasSeverity() bool {
	if o != nil && !utils.IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given EventAlertSeverity and assigns it to the Severity field.
func (o *Event) SetSeverity(v EventAlertSeverity) {
	o.Severity = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Event) GetTitle() string {
	if o == nil || utils.IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTitleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Event) HasTitle() bool {
	if o != nil && !utils.IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Event) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Event) GetType() EventType {
	if o == nil || utils.IsNil(o.Type) {
		var ret EventType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetTypeOk() (*EventType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Event) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EventType and assigns it to the Type field.
func (o *Event) SetType(v EventType) {
	o.Type = &v
}

// GetAffectedTests returns the AffectedTests field value if set, zero value otherwise.
func (o *Event) GetAffectedTests() AffectedCount {
	if o == nil || utils.IsNil(o.AffectedTests) {
		var ret AffectedCount
		return ret
	}
	return *o.AffectedTests
}

// GetAffectedTestsOk returns a tuple with the AffectedTests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAffectedTestsOk() (*AffectedCount, bool) {
	if o == nil || utils.IsNil(o.AffectedTests) {
		return nil, false
	}
	return o.AffectedTests, true
}

// HasAffectedTests returns a boolean if a field has been set.
func (o *Event) HasAffectedTests() bool {
	if o != nil && !utils.IsNil(o.AffectedTests) {
		return true
	}

	return false
}

// SetAffectedTests gets a reference to the given AffectedCount and assigns it to the AffectedTests field.
func (o *Event) SetAffectedTests(v AffectedCount) {
	o.AffectedTests = &v
}

// GetAffectedTargets returns the AffectedTargets field value if set, zero value otherwise.
func (o *Event) GetAffectedTargets() AffectedCount {
	if o == nil || utils.IsNil(o.AffectedTargets) {
		var ret AffectedCount
		return ret
	}
	return *o.AffectedTargets
}

// GetAffectedTargetsOk returns a tuple with the AffectedTargets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAffectedTargetsOk() (*AffectedCount, bool) {
	if o == nil || utils.IsNil(o.AffectedTargets) {
		return nil, false
	}
	return o.AffectedTargets, true
}

// HasAffectedTargets returns a boolean if a field has been set.
func (o *Event) HasAffectedTargets() bool {
	if o != nil && !utils.IsNil(o.AffectedTargets) {
		return true
	}

	return false
}

// SetAffectedTargets gets a reference to the given AffectedCount and assigns it to the AffectedTargets field.
func (o *Event) SetAffectedTargets(v AffectedCount) {
	o.AffectedTargets = &v
}

// GetAffectedAgents returns the AffectedAgents field value if set, zero value otherwise.
func (o *Event) GetAffectedAgents() AffectedCount {
	if o == nil || utils.IsNil(o.AffectedAgents) {
		var ret AffectedCount
		return ret
	}
	return *o.AffectedAgents
}

// GetAffectedAgentsOk returns a tuple with the AffectedAgents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetAffectedAgentsOk() (*AffectedCount, bool) {
	if o == nil || utils.IsNil(o.AffectedAgents) {
		return nil, false
	}
	return o.AffectedAgents, true
}

// HasAffectedAgents returns a boolean if a field has been set.
func (o *Event) HasAffectedAgents() bool {
	if o != nil && !utils.IsNil(o.AffectedAgents) {
		return true
	}

	return false
}

// SetAffectedAgents gets a reference to the given AffectedCount and assigns it to the AffectedAgents field.
func (o *Event) SetAffectedAgents(v AffectedCount) {
	o.AffectedAgents = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Event) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Event) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Event) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *Event) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.TypeName) {
		toSerialize["typeName"] = o.TypeName
	}
	if !utils.IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !utils.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !utils.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !utils.IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !utils.IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.AffectedTests) {
		toSerialize["affectedTests"] = o.AffectedTests
	}
	if !utils.IsNil(o.AffectedTargets) {
		toSerialize["affectedTargets"] = o.AffectedTargets
	}
	if !utils.IsNil(o.AffectedAgents) {
		toSerialize["affectedAgents"] = o.AffectedAgents
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


