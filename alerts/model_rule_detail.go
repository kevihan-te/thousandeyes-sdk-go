/*
Alerts API

You can manage the following alert functionalities on the ThousandEyes platform using the Alerts API:  * **Alerts**: Retrieve alert details. Alerts are assigned to tests through alert rules.  * **Alert Rules**: Conditions that you configure in order to highlight or be notified of events of interest in your ThousandEyes tests. When an alert rule’s conditions are met, the associated alert is triggered and the alert becomes active. It remains active until the alert is cleared. Alert rules are reusable across multiple tests..  * **Alert Suppression Windows**: Suppress alerts for tests during periods such as planned maintenance. Windows can be one-time events or recurring events to handle periodic occurrences such as monthly downtime for maintenance.  For more information about the alerts, see [Alerts](https://docs.thousandeyes.com/product-documentation/alerts). 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerts

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"fmt"
)

// checks if the RuleDetail type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RuleDetail{}

// RuleDetail struct for RuleDetail
type RuleDetail struct {
	// Unique ID of the rule.
	RuleId *string `json:"ruleId,omitempty"`
	// Name of the alert rule.
	RuleName string `json:"ruleName"`
	// The expression of the alert rule.
	Expression string `json:"expression"`
	Direction *AlertDirection `json:"direction,omitempty"`
	// Send notification when alert clears.
	NotifyOnClear *bool `json:"notifyOnClear,omitempty"`
	// If set to `true`, this alert rule becomes the default for its test type and is automatically applied to newly created tests with relevant metrics. Only one default alert rule is allowed per test type.
	IsDefault *bool `json:"isDefault,omitempty"`
	AlertType AlertType `json:"alertType"`
	AlertGroupType *AlertGroupType `json:"alertGroupType,omitempty"`
	// The minimum number of agents or monitors that must meet the specified criteria to trigger the alert.
	MinimumSources *int32 `json:"minimumSources,omitempty"`
	// The minimum percentage of all assigned agents or monitors that must meet the specified criteria to trigger the alert.
	MinimumSourcesPct *int32 `json:"minimumSourcesPct,omitempty"`
	RoundsViolatingMode *AlertRoundsViolationMode `json:"roundsViolatingMode,omitempty"`
	// Specifies the divisor (y value) in the “X of Y times” condition.
	RoundsViolatingOutOf int32 `json:"roundsViolatingOutOf"`
	// Specifies the numerator (x value) in the “X of Y times” condition.
	RoundsViolatingRequired int32 `json:"roundsViolatingRequired"`
	// Set true to include covered prefixes in the BGP alert rule. Only applicable to BGP alert rules.
	IncludeCoveredPrefixes *bool `json:"includeCoveredPrefixes,omitempty"`
	SensitivityLevel *SensitivityLevel `json:"sensitivityLevel,omitempty"`
	Severity *Severity `json:"severity,omitempty"`
	// An array of endpoint agent IDs associated with the rule (get `id` from `/endpoint/agents` API). This is applicable when `alertGroupType` is `browser-session`.
	EndpointAgentIds []string `json:"endpointAgentIds,omitempty"`
	// An array of label IDs used to assign specific Endpoint Agents to the test (get `id` from `/endpoint/labels`). This is applicable when `alertGroupType` is `browser-session`.
	EndpointLabelIds []string `json:"endpointLabelIds,omitempty"`
	// A list of website domains visited during the session. This is applicable when `alertGroupType` is `browser-session`.
	VisitedSitesFilter []string `json:"visitedSitesFilter,omitempty"`
	Notifications *AlertNotification `json:"notifications,omitempty"`
	Tests []AlertSimpleTest `json:"tests,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

type _RuleDetail RuleDetail

// NewRuleDetail instantiates a new RuleDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleDetail(ruleName string, expression string, alertType AlertType, roundsViolatingOutOf int32, roundsViolatingRequired int32) *RuleDetail {
	this := RuleDetail{}
	this.RuleName = ruleName
	this.Expression = expression
	this.AlertType = alertType
	this.RoundsViolatingOutOf = roundsViolatingOutOf
	this.RoundsViolatingRequired = roundsViolatingRequired
	return &this
}

// NewRuleDetailWithDefaults instantiates a new RuleDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleDetailWithDefaults() *RuleDetail {
	this := RuleDetail{}
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *RuleDetail) GetRuleId() string {
	if o == nil || utils.IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetRuleIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *RuleDetail) HasRuleId() bool {
	if o != nil && !utils.IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *RuleDetail) SetRuleId(v string) {
	o.RuleId = &v
}

// GetRuleName returns the RuleName field value
func (o *RuleDetail) GetRuleName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetRuleNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RuleName, true
}

// SetRuleName sets field value
func (o *RuleDetail) SetRuleName(v string) {
	o.RuleName = v
}

// GetExpression returns the Expression field value
func (o *RuleDetail) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *RuleDetail) SetExpression(v string) {
	o.Expression = v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *RuleDetail) GetDirection() AlertDirection {
	if o == nil || utils.IsNil(o.Direction) {
		var ret AlertDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetDirectionOk() (*AlertDirection, bool) {
	if o == nil || utils.IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *RuleDetail) HasDirection() bool {
	if o != nil && !utils.IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given AlertDirection and assigns it to the Direction field.
func (o *RuleDetail) SetDirection(v AlertDirection) {
	o.Direction = &v
}

// GetNotifyOnClear returns the NotifyOnClear field value if set, zero value otherwise.
func (o *RuleDetail) GetNotifyOnClear() bool {
	if o == nil || utils.IsNil(o.NotifyOnClear) {
		var ret bool
		return ret
	}
	return *o.NotifyOnClear
}

// GetNotifyOnClearOk returns a tuple with the NotifyOnClear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetNotifyOnClearOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NotifyOnClear) {
		return nil, false
	}
	return o.NotifyOnClear, true
}

// HasNotifyOnClear returns a boolean if a field has been set.
func (o *RuleDetail) HasNotifyOnClear() bool {
	if o != nil && !utils.IsNil(o.NotifyOnClear) {
		return true
	}

	return false
}

// SetNotifyOnClear gets a reference to the given bool and assigns it to the NotifyOnClear field.
func (o *RuleDetail) SetNotifyOnClear(v bool) {
	o.NotifyOnClear = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *RuleDetail) GetIsDefault() bool {
	if o == nil || utils.IsNil(o.IsDefault) {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetIsDefaultOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsDefault) {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *RuleDetail) HasIsDefault() bool {
	if o != nil && !utils.IsNil(o.IsDefault) {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *RuleDetail) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetAlertType returns the AlertType field value
func (o *RuleDetail) GetAlertType() AlertType {
	if o == nil {
		var ret AlertType
		return ret
	}

	return o.AlertType
}

// GetAlertTypeOk returns a tuple with the AlertType field value
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetAlertTypeOk() (*AlertType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlertType, true
}

// SetAlertType sets field value
func (o *RuleDetail) SetAlertType(v AlertType) {
	o.AlertType = v
}

// GetAlertGroupType returns the AlertGroupType field value if set, zero value otherwise.
func (o *RuleDetail) GetAlertGroupType() AlertGroupType {
	if o == nil || utils.IsNil(o.AlertGroupType) {
		var ret AlertGroupType
		return ret
	}
	return *o.AlertGroupType
}

// GetAlertGroupTypeOk returns a tuple with the AlertGroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetAlertGroupTypeOk() (*AlertGroupType, bool) {
	if o == nil || utils.IsNil(o.AlertGroupType) {
		return nil, false
	}
	return o.AlertGroupType, true
}

// HasAlertGroupType returns a boolean if a field has been set.
func (o *RuleDetail) HasAlertGroupType() bool {
	if o != nil && !utils.IsNil(o.AlertGroupType) {
		return true
	}

	return false
}

// SetAlertGroupType gets a reference to the given AlertGroupType and assigns it to the AlertGroupType field.
func (o *RuleDetail) SetAlertGroupType(v AlertGroupType) {
	o.AlertGroupType = &v
}

// GetMinimumSources returns the MinimumSources field value if set, zero value otherwise.
func (o *RuleDetail) GetMinimumSources() int32 {
	if o == nil || utils.IsNil(o.MinimumSources) {
		var ret int32
		return ret
	}
	return *o.MinimumSources
}

// GetMinimumSourcesOk returns a tuple with the MinimumSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetMinimumSourcesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MinimumSources) {
		return nil, false
	}
	return o.MinimumSources, true
}

// HasMinimumSources returns a boolean if a field has been set.
func (o *RuleDetail) HasMinimumSources() bool {
	if o != nil && !utils.IsNil(o.MinimumSources) {
		return true
	}

	return false
}

// SetMinimumSources gets a reference to the given int32 and assigns it to the MinimumSources field.
func (o *RuleDetail) SetMinimumSources(v int32) {
	o.MinimumSources = &v
}

// GetMinimumSourcesPct returns the MinimumSourcesPct field value if set, zero value otherwise.
func (o *RuleDetail) GetMinimumSourcesPct() int32 {
	if o == nil || utils.IsNil(o.MinimumSourcesPct) {
		var ret int32
		return ret
	}
	return *o.MinimumSourcesPct
}

// GetMinimumSourcesPctOk returns a tuple with the MinimumSourcesPct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetMinimumSourcesPctOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.MinimumSourcesPct) {
		return nil, false
	}
	return o.MinimumSourcesPct, true
}

// HasMinimumSourcesPct returns a boolean if a field has been set.
func (o *RuleDetail) HasMinimumSourcesPct() bool {
	if o != nil && !utils.IsNil(o.MinimumSourcesPct) {
		return true
	}

	return false
}

// SetMinimumSourcesPct gets a reference to the given int32 and assigns it to the MinimumSourcesPct field.
func (o *RuleDetail) SetMinimumSourcesPct(v int32) {
	o.MinimumSourcesPct = &v
}

// GetRoundsViolatingMode returns the RoundsViolatingMode field value if set, zero value otherwise.
func (o *RuleDetail) GetRoundsViolatingMode() AlertRoundsViolationMode {
	if o == nil || utils.IsNil(o.RoundsViolatingMode) {
		var ret AlertRoundsViolationMode
		return ret
	}
	return *o.RoundsViolatingMode
}

// GetRoundsViolatingModeOk returns a tuple with the RoundsViolatingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetRoundsViolatingModeOk() (*AlertRoundsViolationMode, bool) {
	if o == nil || utils.IsNil(o.RoundsViolatingMode) {
		return nil, false
	}
	return o.RoundsViolatingMode, true
}

// HasRoundsViolatingMode returns a boolean if a field has been set.
func (o *RuleDetail) HasRoundsViolatingMode() bool {
	if o != nil && !utils.IsNil(o.RoundsViolatingMode) {
		return true
	}

	return false
}

// SetRoundsViolatingMode gets a reference to the given AlertRoundsViolationMode and assigns it to the RoundsViolatingMode field.
func (o *RuleDetail) SetRoundsViolatingMode(v AlertRoundsViolationMode) {
	o.RoundsViolatingMode = &v
}

// GetRoundsViolatingOutOf returns the RoundsViolatingOutOf field value
func (o *RuleDetail) GetRoundsViolatingOutOf() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RoundsViolatingOutOf
}

// GetRoundsViolatingOutOfOk returns a tuple with the RoundsViolatingOutOf field value
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetRoundsViolatingOutOfOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoundsViolatingOutOf, true
}

// SetRoundsViolatingOutOf sets field value
func (o *RuleDetail) SetRoundsViolatingOutOf(v int32) {
	o.RoundsViolatingOutOf = v
}

// GetRoundsViolatingRequired returns the RoundsViolatingRequired field value
func (o *RuleDetail) GetRoundsViolatingRequired() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RoundsViolatingRequired
}

// GetRoundsViolatingRequiredOk returns a tuple with the RoundsViolatingRequired field value
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetRoundsViolatingRequiredOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoundsViolatingRequired, true
}

// SetRoundsViolatingRequired sets field value
func (o *RuleDetail) SetRoundsViolatingRequired(v int32) {
	o.RoundsViolatingRequired = v
}

// GetIncludeCoveredPrefixes returns the IncludeCoveredPrefixes field value if set, zero value otherwise.
func (o *RuleDetail) GetIncludeCoveredPrefixes() bool {
	if o == nil || utils.IsNil(o.IncludeCoveredPrefixes) {
		var ret bool
		return ret
	}
	return *o.IncludeCoveredPrefixes
}

// GetIncludeCoveredPrefixesOk returns a tuple with the IncludeCoveredPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetIncludeCoveredPrefixesOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IncludeCoveredPrefixes) {
		return nil, false
	}
	return o.IncludeCoveredPrefixes, true
}

// HasIncludeCoveredPrefixes returns a boolean if a field has been set.
func (o *RuleDetail) HasIncludeCoveredPrefixes() bool {
	if o != nil && !utils.IsNil(o.IncludeCoveredPrefixes) {
		return true
	}

	return false
}

// SetIncludeCoveredPrefixes gets a reference to the given bool and assigns it to the IncludeCoveredPrefixes field.
func (o *RuleDetail) SetIncludeCoveredPrefixes(v bool) {
	o.IncludeCoveredPrefixes = &v
}

// GetSensitivityLevel returns the SensitivityLevel field value if set, zero value otherwise.
func (o *RuleDetail) GetSensitivityLevel() SensitivityLevel {
	if o == nil || utils.IsNil(o.SensitivityLevel) {
		var ret SensitivityLevel
		return ret
	}
	return *o.SensitivityLevel
}

// GetSensitivityLevelOk returns a tuple with the SensitivityLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetSensitivityLevelOk() (*SensitivityLevel, bool) {
	if o == nil || utils.IsNil(o.SensitivityLevel) {
		return nil, false
	}
	return o.SensitivityLevel, true
}

// HasSensitivityLevel returns a boolean if a field has been set.
func (o *RuleDetail) HasSensitivityLevel() bool {
	if o != nil && !utils.IsNil(o.SensitivityLevel) {
		return true
	}

	return false
}

// SetSensitivityLevel gets a reference to the given SensitivityLevel and assigns it to the SensitivityLevel field.
func (o *RuleDetail) SetSensitivityLevel(v SensitivityLevel) {
	o.SensitivityLevel = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *RuleDetail) GetSeverity() Severity {
	if o == nil || utils.IsNil(o.Severity) {
		var ret Severity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetSeverityOk() (*Severity, bool) {
	if o == nil || utils.IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *RuleDetail) HasSeverity() bool {
	if o != nil && !utils.IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given Severity and assigns it to the Severity field.
func (o *RuleDetail) SetSeverity(v Severity) {
	o.Severity = &v
}

// GetEndpointAgentIds returns the EndpointAgentIds field value if set, zero value otherwise.
func (o *RuleDetail) GetEndpointAgentIds() []string {
	if o == nil || utils.IsNil(o.EndpointAgentIds) {
		var ret []string
		return ret
	}
	return o.EndpointAgentIds
}

// GetEndpointAgentIdsOk returns a tuple with the EndpointAgentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetEndpointAgentIdsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EndpointAgentIds) {
		return nil, false
	}
	return o.EndpointAgentIds, true
}

// HasEndpointAgentIds returns a boolean if a field has been set.
func (o *RuleDetail) HasEndpointAgentIds() bool {
	if o != nil && !utils.IsNil(o.EndpointAgentIds) {
		return true
	}

	return false
}

// SetEndpointAgentIds gets a reference to the given []string and assigns it to the EndpointAgentIds field.
func (o *RuleDetail) SetEndpointAgentIds(v []string) {
	o.EndpointAgentIds = v
}

// GetEndpointLabelIds returns the EndpointLabelIds field value if set, zero value otherwise.
func (o *RuleDetail) GetEndpointLabelIds() []string {
	if o == nil || utils.IsNil(o.EndpointLabelIds) {
		var ret []string
		return ret
	}
	return o.EndpointLabelIds
}

// GetEndpointLabelIdsOk returns a tuple with the EndpointLabelIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetEndpointLabelIdsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.EndpointLabelIds) {
		return nil, false
	}
	return o.EndpointLabelIds, true
}

// HasEndpointLabelIds returns a boolean if a field has been set.
func (o *RuleDetail) HasEndpointLabelIds() bool {
	if o != nil && !utils.IsNil(o.EndpointLabelIds) {
		return true
	}

	return false
}

// SetEndpointLabelIds gets a reference to the given []string and assigns it to the EndpointLabelIds field.
func (o *RuleDetail) SetEndpointLabelIds(v []string) {
	o.EndpointLabelIds = v
}

// GetVisitedSitesFilter returns the VisitedSitesFilter field value if set, zero value otherwise.
func (o *RuleDetail) GetVisitedSitesFilter() []string {
	if o == nil || utils.IsNil(o.VisitedSitesFilter) {
		var ret []string
		return ret
	}
	return o.VisitedSitesFilter
}

// GetVisitedSitesFilterOk returns a tuple with the VisitedSitesFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetVisitedSitesFilterOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.VisitedSitesFilter) {
		return nil, false
	}
	return o.VisitedSitesFilter, true
}

// HasVisitedSitesFilter returns a boolean if a field has been set.
func (o *RuleDetail) HasVisitedSitesFilter() bool {
	if o != nil && !utils.IsNil(o.VisitedSitesFilter) {
		return true
	}

	return false
}

// SetVisitedSitesFilter gets a reference to the given []string and assigns it to the VisitedSitesFilter field.
func (o *RuleDetail) SetVisitedSitesFilter(v []string) {
	o.VisitedSitesFilter = v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *RuleDetail) GetNotifications() AlertNotification {
	if o == nil || utils.IsNil(o.Notifications) {
		var ret AlertNotification
		return ret
	}
	return *o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetNotificationsOk() (*AlertNotification, bool) {
	if o == nil || utils.IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *RuleDetail) HasNotifications() bool {
	if o != nil && !utils.IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given AlertNotification and assigns it to the Notifications field.
func (o *RuleDetail) SetNotifications(v AlertNotification) {
	o.Notifications = &v
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *RuleDetail) GetTests() []AlertSimpleTest {
	if o == nil || utils.IsNil(o.Tests) {
		var ret []AlertSimpleTest
		return ret
	}
	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetTestsOk() ([]AlertSimpleTest, bool) {
	if o == nil || utils.IsNil(o.Tests) {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *RuleDetail) HasTests() bool {
	if o != nil && !utils.IsNil(o.Tests) {
		return true
	}

	return false
}

// SetTests gets a reference to the given []AlertSimpleTest and assigns it to the Tests field.
func (o *RuleDetail) SetTests(v []AlertSimpleTest) {
	o.Tests = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RuleDetail) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleDetail) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RuleDetail) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *RuleDetail) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o RuleDetail) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RuleDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.RuleId) {
		toSerialize["ruleId"] = o.RuleId
	}
	toSerialize["ruleName"] = o.RuleName
	toSerialize["expression"] = o.Expression
	if !utils.IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !utils.IsNil(o.NotifyOnClear) {
		toSerialize["notifyOnClear"] = o.NotifyOnClear
	}
	if !utils.IsNil(o.IsDefault) {
		toSerialize["isDefault"] = o.IsDefault
	}
	toSerialize["alertType"] = o.AlertType
	if !utils.IsNil(o.AlertGroupType) {
		toSerialize["alertGroupType"] = o.AlertGroupType
	}
	if !utils.IsNil(o.MinimumSources) {
		toSerialize["minimumSources"] = o.MinimumSources
	}
	if !utils.IsNil(o.MinimumSourcesPct) {
		toSerialize["minimumSourcesPct"] = o.MinimumSourcesPct
	}
	if !utils.IsNil(o.RoundsViolatingMode) {
		toSerialize["roundsViolatingMode"] = o.RoundsViolatingMode
	}
	toSerialize["roundsViolatingOutOf"] = o.RoundsViolatingOutOf
	toSerialize["roundsViolatingRequired"] = o.RoundsViolatingRequired
	if !utils.IsNil(o.IncludeCoveredPrefixes) {
		toSerialize["includeCoveredPrefixes"] = o.IncludeCoveredPrefixes
	}
	if !utils.IsNil(o.SensitivityLevel) {
		toSerialize["sensitivityLevel"] = o.SensitivityLevel
	}
	if !utils.IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !utils.IsNil(o.EndpointAgentIds) {
		toSerialize["endpointAgentIds"] = o.EndpointAgentIds
	}
	if !utils.IsNil(o.EndpointLabelIds) {
		toSerialize["endpointLabelIds"] = o.EndpointLabelIds
	}
	if !utils.IsNil(o.VisitedSitesFilter) {
		toSerialize["visitedSitesFilter"] = o.VisitedSitesFilter
	}
	if !utils.IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !utils.IsNil(o.Tests) {
		toSerialize["tests"] = o.Tests
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

func (o *RuleDetail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ruleName",
		"expression",
		"alertType",
		"roundsViolatingOutOf",
		"roundsViolatingRequired",
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

	varRuleDetail := _RuleDetail{}

    err = json.Unmarshal(data, &varRuleDetail)

	if err != nil {
		return err
	}

	*o = RuleDetail(varRuleDetail)

	return err
}

type NullableRuleDetail struct {
	value *RuleDetail
	isSet bool
}

func (v NullableRuleDetail) Get() *RuleDetail {
	return v.value
}

func (v *NullableRuleDetail) Set(val *RuleDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableRuleDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableRuleDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRuleDetail(val *RuleDetail) *NullableRuleDetail {
	return &NullableRuleDetail{value: val, isSet: true}
}

func (v NullableRuleDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRuleDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


