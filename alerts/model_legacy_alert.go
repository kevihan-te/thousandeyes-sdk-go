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
)

// checks if the LegacyAlert type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LegacyAlert{}

// LegacyAlert struct for LegacyAlert
type LegacyAlert struct {
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
}

// NewLegacyAlert instantiates a new LegacyAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLegacyAlert() *LegacyAlert {
	this := LegacyAlert{}
	return &this
}

// NewLegacyAlertWithDefaults instantiates a new LegacyAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLegacyAlertWithDefaults() *LegacyAlert {
	this := LegacyAlert{}
	return &this
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *LegacyAlert) GetAlertId() string {
	if o == nil || utils.IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetAlertIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *LegacyAlert) HasAlertId() bool {
	if o != nil && !utils.IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *LegacyAlert) SetAlertId(v string) {
	o.AlertId = &v
}

// GetDateStart returns the DateStart field value if set, zero value otherwise.
func (o *LegacyAlert) GetDateStart() string {
	if o == nil || utils.IsNil(o.DateStart) {
		var ret string
		return ret
	}
	return *o.DateStart
}

// GetDateStartOk returns a tuple with the DateStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetDateStartOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DateStart) {
		return nil, false
	}
	return o.DateStart, true
}

// HasDateStart returns a boolean if a field has been set.
func (o *LegacyAlert) HasDateStart() bool {
	if o != nil && !utils.IsNil(o.DateStart) {
		return true
	}

	return false
}

// SetDateStart gets a reference to the given string and assigns it to the DateStart field.
func (o *LegacyAlert) SetDateStart(v string) {
	o.DateStart = &v
}

// GetDateEnd returns the DateEnd field value if set, zero value otherwise.
func (o *LegacyAlert) GetDateEnd() string {
	if o == nil || utils.IsNil(o.DateEnd) {
		var ret string
		return ret
	}
	return *o.DateEnd
}

// GetDateEndOk returns a tuple with the DateEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetDateEndOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DateEnd) {
		return nil, false
	}
	return o.DateEnd, true
}

// HasDateEnd returns a boolean if a field has been set.
func (o *LegacyAlert) HasDateEnd() bool {
	if o != nil && !utils.IsNil(o.DateEnd) {
		return true
	}

	return false
}

// SetDateEnd gets a reference to the given string and assigns it to the DateEnd field.
func (o *LegacyAlert) SetDateEnd(v string) {
	o.DateEnd = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *LegacyAlert) GetRuleId() int32 {
	if o == nil || utils.IsNil(o.RuleId) {
		var ret int32
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetRuleIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *LegacyAlert) HasRuleId() bool {
	if o != nil && !utils.IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given int32 and assigns it to the RuleId field.
func (o *LegacyAlert) SetRuleId(v int32) {
	o.RuleId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LegacyAlert) GetState() string {
	if o == nil || utils.IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetStateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LegacyAlert) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LegacyAlert) SetState(v string) {
	o.State = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *LegacyAlert) GetSeverity() string {
	if o == nil || utils.IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetSeverityOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *LegacyAlert) HasSeverity() bool {
	if o != nil && !utils.IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *LegacyAlert) SetSeverity(v string) {
	o.Severity = &v
}

// GetPermalink returns the Permalink field value if set, zero value otherwise.
func (o *LegacyAlert) GetPermalink() string {
	if o == nil || utils.IsNil(o.Permalink) {
		var ret string
		return ret
	}
	return *o.Permalink
}

// GetPermalinkOk returns a tuple with the Permalink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetPermalinkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Permalink) {
		return nil, false
	}
	return o.Permalink, true
}

// HasPermalink returns a boolean if a field has been set.
func (o *LegacyAlert) HasPermalink() bool {
	if o != nil && !utils.IsNil(o.Permalink) {
		return true
	}

	return false
}

// SetPermalink gets a reference to the given string and assigns it to the Permalink field.
func (o *LegacyAlert) SetPermalink(v string) {
	o.Permalink = &v
}

// GetApiLinks returns the ApiLinks field value if set, zero value otherwise.
func (o *LegacyAlert) GetApiLinks() []map[string]interface{} {
	if o == nil || utils.IsNil(o.ApiLinks) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ApiLinks
}

// GetApiLinksOk returns a tuple with the ApiLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LegacyAlert) GetApiLinksOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.ApiLinks) {
		return nil, false
	}
	return o.ApiLinks, true
}

// HasApiLinks returns a boolean if a field has been set.
func (o *LegacyAlert) HasApiLinks() bool {
	if o != nil && !utils.IsNil(o.ApiLinks) {
		return true
	}

	return false
}

// SetApiLinks gets a reference to the given []map[string]interface{} and assigns it to the ApiLinks field.
func (o *LegacyAlert) SetApiLinks(v []map[string]interface{}) {
	o.ApiLinks = v
}

func (o LegacyAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LegacyAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableLegacyAlert struct {
	value *LegacyAlert
	isSet bool
}

func (v NullableLegacyAlert) Get() *LegacyAlert {
	return v.value
}

func (v *NullableLegacyAlert) Set(val *LegacyAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableLegacyAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableLegacyAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLegacyAlert(val *LegacyAlert) *NullableLegacyAlert {
	return &NullableLegacyAlert{value: val, isSet: true}
}

func (v NullableLegacyAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLegacyAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


