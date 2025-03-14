/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the ApiAlertListAlert type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiAlertListAlert{}

// ApiAlertListAlert Alert shown in the alert list widget.
type ApiAlertListAlert struct {
	// Identifier of the alert.
	AlertId *string `json:"alertId,omitempty"`
	// Identifier of the test.
	TestId *string `json:"testId,omitempty"`
	// Identifier of the rule.
	RuleId *string `json:"ruleId,omitempty"`
	// Name of the agent, monitor or device producing the alert.
	AlertSource *string `json:"alertSource,omitempty"`
	// Name of the alert rule that this alert belongs to.
	AlertRule *string `json:"alertRule,omitempty"`
	AlertType *AlertListAlertType `json:"alertType,omitempty"`
	// UTC date when the alert was first active.
	StartTime *time.Time `json:"startTime,omitempty"`
	// Number of seconds the alert was active. If it’s still active, this number will increase every second.
	DurationInSeconds *int64 `json:"durationInSeconds,omitempty"`
	// Set to `true` if alert is active, `false` otherwise.
	Active *bool `json:"active,omitempty"`
}

// NewApiAlertListAlert instantiates a new ApiAlertListAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAlertListAlert() *ApiAlertListAlert {
	this := ApiAlertListAlert{}
	return &this
}

// NewApiAlertListAlertWithDefaults instantiates a new ApiAlertListAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAlertListAlertWithDefaults() *ApiAlertListAlert {
	this := ApiAlertListAlert{}
	return &this
}

// GetAlertId returns the AlertId field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetAlertId() string {
	if o == nil || utils.IsNil(o.AlertId) {
		var ret string
		return ret
	}
	return *o.AlertId
}

// GetAlertIdOk returns a tuple with the AlertId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetAlertIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AlertId) {
		return nil, false
	}
	return o.AlertId, true
}

// HasAlertId returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasAlertId() bool {
	if o != nil && !utils.IsNil(o.AlertId) {
		return true
	}

	return false
}

// SetAlertId gets a reference to the given string and assigns it to the AlertId field.
func (o *ApiAlertListAlert) SetAlertId(v string) {
	o.AlertId = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *ApiAlertListAlert) SetTestId(v string) {
	o.TestId = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetRuleId() string {
	if o == nil || utils.IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetRuleIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasRuleId() bool {
	if o != nil && !utils.IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *ApiAlertListAlert) SetRuleId(v string) {
	o.RuleId = &v
}

// GetAlertSource returns the AlertSource field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetAlertSource() string {
	if o == nil || utils.IsNil(o.AlertSource) {
		var ret string
		return ret
	}
	return *o.AlertSource
}

// GetAlertSourceOk returns a tuple with the AlertSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetAlertSourceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AlertSource) {
		return nil, false
	}
	return o.AlertSource, true
}

// HasAlertSource returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasAlertSource() bool {
	if o != nil && !utils.IsNil(o.AlertSource) {
		return true
	}

	return false
}

// SetAlertSource gets a reference to the given string and assigns it to the AlertSource field.
func (o *ApiAlertListAlert) SetAlertSource(v string) {
	o.AlertSource = &v
}

// GetAlertRule returns the AlertRule field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetAlertRule() string {
	if o == nil || utils.IsNil(o.AlertRule) {
		var ret string
		return ret
	}
	return *o.AlertRule
}

// GetAlertRuleOk returns a tuple with the AlertRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetAlertRuleOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AlertRule) {
		return nil, false
	}
	return o.AlertRule, true
}

// HasAlertRule returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasAlertRule() bool {
	if o != nil && !utils.IsNil(o.AlertRule) {
		return true
	}

	return false
}

// SetAlertRule gets a reference to the given string and assigns it to the AlertRule field.
func (o *ApiAlertListAlert) SetAlertRule(v string) {
	o.AlertRule = &v
}

// GetAlertType returns the AlertType field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetAlertType() AlertListAlertType {
	if o == nil || utils.IsNil(o.AlertType) {
		var ret AlertListAlertType
		return ret
	}
	return *o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetAlertTypeOk() (*AlertListAlertType, bool) {
	if o == nil || utils.IsNil(o.AlertType) {
		return nil, false
	}
	return o.AlertType, true
}

// HasAlertType returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasAlertType() bool {
	if o != nil && !utils.IsNil(o.AlertType) {
		return true
	}

	return false
}

// SetAlertType gets a reference to the given AlertListAlertType and assigns it to the AlertType field.
func (o *ApiAlertListAlert) SetAlertType(v AlertListAlertType) {
	o.AlertType = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetStartTime() time.Time {
	if o == nil || utils.IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasStartTime() bool {
	if o != nil && !utils.IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *ApiAlertListAlert) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetDurationInSeconds returns the DurationInSeconds field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetDurationInSeconds() int64 {
	if o == nil || utils.IsNil(o.DurationInSeconds) {
		var ret int64
		return ret
	}
	return *o.DurationInSeconds
}

// GetDurationInSecondsOk returns a tuple with the DurationInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetDurationInSecondsOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.DurationInSeconds) {
		return nil, false
	}
	return o.DurationInSeconds, true
}

// HasDurationInSeconds returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasDurationInSeconds() bool {
	if o != nil && !utils.IsNil(o.DurationInSeconds) {
		return true
	}

	return false
}

// SetDurationInSeconds gets a reference to the given int64 and assigns it to the DurationInSeconds field.
func (o *ApiAlertListAlert) SetDurationInSeconds(v int64) {
	o.DurationInSeconds = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ApiAlertListAlert) GetActive() bool {
	if o == nil || utils.IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAlertListAlert) GetActiveOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ApiAlertListAlert) HasActive() bool {
	if o != nil && !utils.IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ApiAlertListAlert) SetActive(v bool) {
	o.Active = &v
}

func (o ApiAlertListAlert) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAlertListAlert) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AlertId) {
		toSerialize["alertId"] = o.AlertId
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	if !utils.IsNil(o.AlertSource) {
		toSerialize["alertSource"] = o.AlertSource
	}
	if !utils.IsNil(o.AlertRule) {
		toSerialize["alertRule"] = o.AlertRule
	}
	if !utils.IsNil(o.AlertType) {
		toSerialize["alertType"] = o.AlertType
	}
	if !utils.IsNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !utils.IsNil(o.DurationInSeconds) {
		toSerialize["durationInSeconds"] = o.DurationInSeconds
	}
	if !utils.IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableApiAlertListAlert struct {
	value *ApiAlertListAlert
	isSet bool
}

func (v NullableApiAlertListAlert) Get() *ApiAlertListAlert {
	return v.value
}

func (v *NullableApiAlertListAlert) Set(val *ApiAlertListAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAlertListAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAlertListAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAlertListAlert(val *ApiAlertListAlert) *NullableApiAlertListAlert {
	return &NullableApiAlertListAlert{value: val, isSet: true}
}

func (v NullableApiAlertListAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAlertListAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


