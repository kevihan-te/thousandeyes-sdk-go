/*
Alerts API

You can manage the following alert functionalities on the ThousandEyes platform using the Alerts API:  * **Alerts**: Retrieve alert details. Alerts are assigned to tests through alert rules.  * **Alert Rules**: Conditions that you configure in order to highlight or be notified of events of interest in your ThousandEyes tests. When an alert rule’s conditions are met, the associated alert is triggered and the alert becomes active. It remains active until the alert is cleared. Alert rules are reusable across multiple tests..  * **Alert Suppression Windows**: Suppress alerts for tests during periods such as planned maintenance. Windows can be one-time events or recurring events to handle periodic occurrences such as monthly downtime for maintenance.  For more information about the alerts, see [Alerts](https://docs.thousandeyes.com/product-documentation/alerts). 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerts

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the Alert type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Alert{}

// Alert struct for Alert
type Alert struct {
	// A unique ID for each individual alert occurrence.
	Id *string `json:"id,omitempty"`
	AlertType *AlertType `json:"alertType,omitempty"`
	// (Optional) When passing `window` or `startDate` parameter,  the client will also receive the `startDate` field indicating the UTC start date of the data's time range being retrieved  (ISO date-time format).
	StartDate *time.Time `json:"startDate,omitempty"`
	// (Optional) When passing `window` or `endDate` parameter,  the client will also receive the `endDate` field indicating the UTC end date of the data's time range being retrieved  (ISO date-time format).
	EndDate *time.Time `json:"endDate,omitempty"`
	// Number of sources that meet the alert criteria.
	ViolationCount *int32 `json:"violationCount,omitempty"`
	// Duration in seconds the alert was active
	Duration *int64 `json:"duration,omitempty"`
	// Indicates whether the alert is currently suppressed by a real-time ASW.
	Suppressed *bool `json:"suppressed,omitempty"`
	Meta *AlertMeta `json:"meta,omitempty"`
	Links *AlertLinks `json:"_links,omitempty"`
	// A unique ID for each individual alert occurrence.
	AlertId *string `json:"alertId,omitempty"`
	// The start date and time for querying alerts.
	DateStart *string `json:"dateStart,omitempty"`
	// The end date and time for querying alerts.
	DateEnd *string `json:"dateEnd,omitempty"`
	// Unique ID of the rule.
	RuleId *int32 `json:"ruleId,omitempty"`
	// Current state of the alert. Possible values: clear or trigger.
	State *string `json:"state,omitempty"`
	// The severity of the alert.
	Severity *string `json:"severity,omitempty"`
	// Hyperlink to alerts list, with row expanded
	Permalink *string `json:"permalink,omitempty"`
	// List of hyperlinks to other areas of the API
	ApiLinks []map[string]interface{} `json:"apiLinks,omitempty"`
	// Unique ID of the rule.
	AlertRuleId *string `json:"alertRuleId,omitempty"`
	AlertState *State `json:"alertState,omitempty"`
	AlertSeverity *Severity `json:"alertSeverity,omitempty"`
}

// NewAlert instantiates a new Alert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlert() *Alert {
	this := Alert{}
	return &this
}

// NewAlertWithDefaults instantiates a new Alert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertWithDefaults() *Alert {
	this := Alert{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Alert) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Alert) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Alert) SetId(v string) {
	o.Id = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *Alert) GetAlertType() AlertType {
	if o == nil || utils.IsNil(o.AlertType) {
		var ret AlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetAlertTypeOk() (*AlertType, bool) {
	if o == nil || utils.IsNil(o.AlertType) {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *Alert) HasAlertType() bool {
	if o != nil && !utils.IsNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given AlertType and assigns it to the AlertType field.
func (o *Alert) SetAlertType(v AlertType) {
	o.AlertType = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Alert) GetStartDate() time.Time {
	if o == nil || utils.IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetStartDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Alert) HasStartDate() bool {
	if o != nil && !utils.IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Alert) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Alert) GetEndDate() time.Time {
	if o == nil || utils.IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetEndDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *Alert) HasEndDate() bool {
	if o != nil && !utils.IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *Alert) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetViolationCount returns the ViolationCount field value if set, zero value otherwise.
func (o *Alert) GetViolationCount() int32 {
	if o == nil || utils.IsNil(o.ViolationCount) {
		var ret int32
		return ret
	}
	return *o.ViolationCount
}

// GetViolationCountOk returns a tuple with the ViolationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetViolationCountOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ViolationCount) {
		return nil, false
	}
	return o.ViolationCount, true
}

// HasViolationCount returns a boolean if a field has been set.
func (o *Alert) HasViolationCount() bool {
	if o != nil && !utils.IsNil(o.ViolationCount) {
		return true
	}

	return false
}

// SetViolationCount gets a reference to the given int32 and assigns it to the ViolationCount field.
func (o *Alert) SetViolationCount(v int32) {
	o.ViolationCount = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Alert) GetDuration() int64 {
	if o == nil || utils.IsNil(o.Duration) {
		var ret int64
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetDurationOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Alert) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int64 and assigns it to the Duration field.
func (o *Alert) SetDuration(v int64) {
	o.Duration = &v
}

// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *Alert) GetSuppressed() bool {
	if o == nil || utils.IsNil(o.Suppressed) {
		var ret bool
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetSuppressedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Suppressed) {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *Alert) HasSuppressed() bool {
	if o != nil && !utils.IsNil(o.Suppressed) {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given bool and assigns it to the Suppressed field.
func (o *Alert) SetSuppressed(v bool) {
	o.Suppressed = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Alert) GetMeta() AlertMeta {
	if o == nil || utils.IsNil(o.Meta) {
		var ret AlertMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetMetaOk() (*AlertMeta, bool) {
	if o == nil || utils.IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Alert) HasMeta() bool {
	if o != nil && !utils.IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given AlertMeta and assigns it to the Meta field.
func (o *Alert) SetMeta(v AlertMeta) {
	o.Meta = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Alert) GetLinks() AlertLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret AlertLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetLinksOk() (*AlertLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Alert) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AlertLinks and assigns it to the Links field.
func (o *Alert) SetLinks(v AlertLinks) {
	o.Links = &v
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *Alert) GetAlertId() string {
	if o == nil || utils.IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetAlertIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *Alert) HasAlertId() bool {
	if o != nil && !utils.IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *Alert) SetAlertId(v string) {
	o.AlertId = &v
}

// GetDateStart returns the DateStart field value if set, zero value otherwise.
func (o *Alert) GetDateStart() string {
	if o == nil || utils.IsNil(o.DateStart) {
		var ret string
		return ret
	}
	return *o.DateStart
}

// GetDateStartOk returns a tuple with the DateStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetDateStartOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DateStart) {
		return nil, false
	}
	return o.DateStart, true
}

// HasDateStart returns a boolean if a field has been set.
func (o *Alert) HasDateStart() bool {
	if o != nil && !utils.IsNil(o.DateStart) {
		return true
	}

	return false
}

// SetDateStart gets a reference to the given string and assigns it to the DateStart field.
func (o *Alert) SetDateStart(v string) {
	o.DateStart = &v
}

// GetDateEnd returns the DateEnd field value if set, zero value otherwise.
func (o *Alert) GetDateEnd() string {
	if o == nil || utils.IsNil(o.DateEnd) {
		var ret string
		return ret
	}
	return *o.DateEnd
}

// GetDateEndOk returns a tuple with the DateEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetDateEndOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DateEnd) {
		return nil, false
	}
	return o.DateEnd, true
}

// HasDateEnd returns a boolean if a field has been set.
func (o *Alert) HasDateEnd() bool {
	if o != nil && !utils.IsNil(o.DateEnd) {
		return true
	}

	return false
}

// SetDateEnd gets a reference to the given string and assigns it to the DateEnd field.
func (o *Alert) SetDateEnd(v string) {
	o.DateEnd = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *Alert) GetRuleId() int32 {
	if o == nil || utils.IsNil(o.RuleId) {
		var ret int32
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetRuleIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *Alert) HasRuleId() bool {
	if o != nil && !utils.IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given int32 and assigns it to the RuleId field.
func (o *Alert) SetRuleId(v int32) {
	o.RuleId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Alert) GetState() string {
	if o == nil || utils.IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetStateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Alert) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Alert) SetState(v string) {
	o.State = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *Alert) GetSeverity() string {
	if o == nil || utils.IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetSeverityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *Alert) HasSeverity() bool {
	if o != nil && !utils.IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *Alert) SetSeverity(v string) {
	o.Severity = &v
}

// GetPermalink returns the Permalink field value if set, zero value otherwise.
func (o *Alert) GetPermalink() string {
	if o == nil || utils.IsNil(o.Permalink) {
		var ret string
		return ret
	}
	return *o.Permalink
}

// GetPermalinkOk returns a tuple with the Permalink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetPermalinkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Permalink) {
		return nil, false
	}
	return o.Permalink, true
}

// HasPermalink returns a boolean if a field has been set.
func (o *Alert) HasPermalink() bool {
	if o != nil && !utils.IsNil(o.Permalink) {
		return true
	}

	return false
}

// SetPermalink gets a reference to the given string and assigns it to the Permalink field.
func (o *Alert) SetPermalink(v string) {
	o.Permalink = &v
}

// GetApiLinks returns the ApiLinks field value if set, zero value otherwise.
func (o *Alert) GetApiLinks() []map[string]interface{} {
	if o == nil || utils.IsNil(o.ApiLinks) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ApiLinks
}

// GetApiLinksOk returns a tuple with the ApiLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetApiLinksOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.ApiLinks) {
		return nil, false
	}
	return o.ApiLinks, true
}

// HasApiLinks returns a boolean if a field has been set.
func (o *Alert) HasApiLinks() bool {
	if o != nil && !utils.IsNil(o.ApiLinks) {
		return true
	}

	return false
}

// SetApiLinks gets a reference to the given []map[string]interface{} and assigns it to the ApiLinks field.
func (o *Alert) SetApiLinks(v []map[string]interface{}) {
	o.ApiLinks = v
}

// GetAlertRuleId returns the AlertRuleId field value if set, zero value otherwise.
func (o *Alert) GetAlertRuleId() string {
	if o == nil || utils.IsNil(o.AlertRuleId) {
		var ret string
		return ret
	}
	return *o.AlertRuleId
}

// GetAlertRuleIdOk returns a tuple with the AlertRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetAlertRuleIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AlertRuleId) {
		return nil, false
	}
	return o.AlertRuleId, true
}

// HasAlertRuleId returns a boolean if a field has been set.
func (o *Alert) HasAlertRuleId() bool {
	if o != nil && !utils.IsNil(o.AlertRuleId) {
		return true
	}

	return false
}

// SetAlertRuleId gets a reference to the given string and assigns it to the AlertRuleId field.
func (o *Alert) SetAlertRuleId(v string) {
	o.AlertRuleId = &v
}

// GetAlertState returns the AlertState field value if set, zero value otherwise.
func (o *Alert) GetAlertState() State {
	if o == nil || utils.IsNil(o.AlertState) {
		var ret State
		return ret
	}
	return *o.AlertState
}

// GetAlertStateOk returns a tuple with the AlertState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetAlertStateOk() (*State, bool) {
	if o == nil || utils.IsNil(o.AlertState) {
		return nil, false
	}
	return o.AlertState, true
}

// HasAlertState returns a boolean if a field has been set.
func (o *Alert) HasAlertState() bool {
	if o != nil && !utils.IsNil(o.AlertState) {
		return true
	}

	return false
}

// SetAlertState gets a reference to the given State and assigns it to the AlertState field.
func (o *Alert) SetAlertState(v State) {
	o.AlertState = &v
}

// GetAlertSeverity returns the AlertSeverity field value if set, zero value otherwise.
func (o *Alert) GetAlertSeverity() Severity {
	if o == nil || utils.IsNil(o.AlertSeverity) {
		var ret Severity
		return ret
	}
	return *o.AlertSeverity
}

// GetAlertSeverityOk returns a tuple with the AlertSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alert) GetAlertSeverityOk() (*Severity, bool) {
	if o == nil || utils.IsNil(o.AlertSeverity) {
		return nil, false
	}
	return o.AlertSeverity, true
}

// HasAlertSeverity returns a boolean if a field has been set.
func (o *Alert) HasAlertSeverity() bool {
	if o != nil && !utils.IsNil(o.AlertSeverity) {
		return true
	}

	return false
}

// SetAlertSeverity gets a reference to the given Severity and assigns it to the AlertSeverity field.
func (o *Alert) SetAlertSeverity(v Severity) {
	o.AlertSeverity = &v
}

func (o Alert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Alert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !utils.IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !utils.IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !utils.IsNil(o.ViolationCount) {
		toSerialize["violationCount"] = o.ViolationCount
	}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !utils.IsNil(o.Suppressed) {
		toSerialize["suppressed"] = o.Suppressed
	}
	if !utils.IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !utils.IsNil(o.AlertId) {
		toSerialize["alertId"] = o.AlertId
	}
	if !utils.IsNil(o.DateStart) {
		toSerialize["dateStart"] = o.DateStart
	}
	if !utils.IsNil(o.DateEnd) {
		toSerialize["dateEnd"] = o.DateEnd
	}
	if !utils.IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	if !utils.IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !utils.IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !utils.IsNil(o.Permalink) {
		toSerialize["permalink"] = o.Permalink
	}
	if !utils.IsNil(o.ApiLinks) {
		toSerialize["apiLinks"] = o.ApiLinks
	}
	if !utils.IsNil(o.AlertRuleId) {
		toSerialize["alertRuleId"] = o.AlertRuleId
	}
	if !utils.IsNil(o.AlertState) {
		toSerialize["alertState"] = o.AlertState
	}
	if !utils.IsNil(o.AlertSeverity) {
		toSerialize["alertSeverity"] = o.AlertSeverity
	}
	return toSerialize, nil
}

type NullableAlert struct {
	value *Alert
	isSet bool
}

func (v NullableAlert) Get() *Alert {
	return v.value
}

func (v *NullableAlert) Set(val *Alert) {
	v.value = val
	v.isSet = true
}

func (v NullableAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlert(val *Alert) *NullableAlert {
	return &NullableAlert{value: val, isSet: true}
}

func (v NullableAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


