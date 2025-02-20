/*
Dashboards API

Manage ThousandEyes Dashboards.

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dashboards

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the ApiDashboardSnapshot type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApiDashboardSnapshot{}

// ApiDashboardSnapshot struct for ApiDashboardSnapshot
type ApiDashboardSnapshot struct {
	// Identifier of the account group that the snapshot belongs to.
	AccountId *int64 `json:"accountId,omitempty"`
	// UTC date when dashboard snapshot was created.
	CreatedDate *string `json:"createdDate,omitempty"`
	// Expiration date of the snapshot. If unspecified, the snapshot expires 1 year from its creation date. The expiration date must be set within 5 years from the current date.
	ExpirationDate *string `json:"expirationDate,omitempty"`
	// Hyperlink to dashboard snapshot in ThousandEyes Application
	Permalink *string `json:"permalink,omitempty"`
	// A links array containing the self link.
	ApiLinks []map[string]interface{} `json:"apiLinks,omitempty"`
	// Identifier of the dashboard snapshot.
	SnapshotId *string `json:"snapshotId,omitempty"`
	// Name of the dashboard snapshot.
	SnapshotName *string `json:"snapshotName,omitempty"`
	// Identifier of the account group that the snapshot belongs to.
	Aid *string `json:"aid,omitempty"`
	// Set `true` if dashboard snapshot is shared, `false` otherwise.
	IsShared *bool `json:"isShared,omitempty"`
	// UTC date when dashboard snapshot was created (ISO date-time format).
	SnapshotCreatedDate *time.Time `json:"snapshotCreatedDate,omitempty"`
	Dashboard *ApiDashboard `json:"dashboard,omitempty"`
	Widgets []ApiWidget `json:"widgets,omitempty"`
	// Set `true` if dashboard snapshot was generated from a schedule, `false` otherwise.
	IsScheduled *bool `json:"isScheduled,omitempty"`
	TimeSpan *ApiReportSnapshotTimeSpan `json:"timeSpan,omitempty"`
	// Expiration date of the snapshot. If unspecified, the snapshot expires 1 year from its creation date. The expiration date must be set within 5 years from the current date and adhere to the ISO date-time format.
	SnapshotExpirationDate *time.Time `json:"snapshotExpirationDate,omitempty"`
	Links *AppAndSelfLinks `json:"_links,omitempty"`
}

// NewApiDashboardSnapshot instantiates a new ApiDashboardSnapshot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDashboardSnapshot() *ApiDashboardSnapshot {
	this := ApiDashboardSnapshot{}
	return &this
}

// NewApiDashboardSnapshotWithDefaults instantiates a new ApiDashboardSnapshot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDashboardSnapshotWithDefaults() *ApiDashboardSnapshot {
	this := ApiDashboardSnapshot{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetAccountId() int64 {
	if o == nil || utils.IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetAccountIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasAccountId() bool {
	if o != nil && !utils.IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *ApiDashboardSnapshot) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetCreatedDate() string {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret string
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetCreatedDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given string and assigns it to the CreatedDate field.
func (o *ApiDashboardSnapshot) SetCreatedDate(v string) {
	o.CreatedDate = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetExpirationDate() string {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		var ret string
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetExpirationDateOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasExpirationDate() bool {
	if o != nil && !utils.IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given string and assigns it to the ExpirationDate field.
func (o *ApiDashboardSnapshot) SetExpirationDate(v string) {
	o.ExpirationDate = &v
}

// GetPermalink returns the Permalink field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetPermalink() string {
	if o == nil || utils.IsNil(o.Permalink) {
		var ret string
		return ret
	}
	return *o.Permalink
}

// GetPermalinkOk returns a tuple with the Permalink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetPermalinkOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Permalink) {
		return nil, false
	}
	return o.Permalink, true
}

// HasPermalink returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasPermalink() bool {
	if o != nil && !utils.IsNil(o.Permalink) {
		return true
	}

	return false
}

// SetPermalink gets a reference to the given string and assigns it to the Permalink field.
func (o *ApiDashboardSnapshot) SetPermalink(v string) {
	o.Permalink = &v
}

// GetApiLinks returns the ApiLinks field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetApiLinks() []map[string]interface{} {
	if o == nil || utils.IsNil(o.ApiLinks) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ApiLinks
}

// GetApiLinksOk returns a tuple with the ApiLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetApiLinksOk() ([]map[string]interface{}, bool) {
	if o == nil || utils.IsNil(o.ApiLinks) {
		return nil, false
	}
	return o.ApiLinks, true
}

// HasApiLinks returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasApiLinks() bool {
	if o != nil && !utils.IsNil(o.ApiLinks) {
		return true
	}

	return false
}

// SetApiLinks gets a reference to the given []map[string]interface{} and assigns it to the ApiLinks field.
func (o *ApiDashboardSnapshot) SetApiLinks(v []map[string]interface{}) {
	o.ApiLinks = v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetSnapshotId() string {
	if o == nil || utils.IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetSnapshotIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasSnapshotId() bool {
	if o != nil && !utils.IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ApiDashboardSnapshot) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetSnapshotName returns the SnapshotName field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetSnapshotName() string {
	if o == nil || utils.IsNil(o.SnapshotName) {
		var ret string
		return ret
	}
	return *o.SnapshotName
}

// GetSnapshotNameOk returns a tuple with the SnapshotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetSnapshotNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SnapshotName) {
		return nil, false
	}
	return o.SnapshotName, true
}

// HasSnapshotName returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasSnapshotName() bool {
	if o != nil && !utils.IsNil(o.SnapshotName) {
		return true
	}

	return false
}

// SetSnapshotName gets a reference to the given string and assigns it to the SnapshotName field.
func (o *ApiDashboardSnapshot) SetSnapshotName(v string) {
	o.SnapshotName = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *ApiDashboardSnapshot) SetAid(v string) {
	o.Aid = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetIsShared() bool {
	if o == nil || utils.IsNil(o.IsShared) {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetIsSharedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsShared) {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasIsShared() bool {
	if o != nil && !utils.IsNil(o.IsShared) {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *ApiDashboardSnapshot) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetSnapshotCreatedDate returns the SnapshotCreatedDate field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetSnapshotCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.SnapshotCreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.SnapshotCreatedDate
}

// GetSnapshotCreatedDateOk returns a tuple with the SnapshotCreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetSnapshotCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.SnapshotCreatedDate) {
		return nil, false
	}
	return o.SnapshotCreatedDate, true
}

// HasSnapshotCreatedDate returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasSnapshotCreatedDate() bool {
	if o != nil && !utils.IsNil(o.SnapshotCreatedDate) {
		return true
	}

	return false
}

// SetSnapshotCreatedDate gets a reference to the given time.Time and assigns it to the SnapshotCreatedDate field.
func (o *ApiDashboardSnapshot) SetSnapshotCreatedDate(v time.Time) {
	o.SnapshotCreatedDate = &v
}

// GetDashboard returns the Dashboard field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetDashboard() ApiDashboard {
	if o == nil || utils.IsNil(o.Dashboard) {
		var ret ApiDashboard
		return ret
	}
	return *o.Dashboard
}

// GetDashboardOk returns a tuple with the Dashboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetDashboardOk() (*ApiDashboard, bool) {
	if o == nil || utils.IsNil(o.Dashboard) {
		return nil, false
	}
	return o.Dashboard, true
}

// HasDashboard returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasDashboard() bool {
	if o != nil && !utils.IsNil(o.Dashboard) {
		return true
	}

	return false
}

// SetDashboard gets a reference to the given ApiDashboard and assigns it to the Dashboard field.
func (o *ApiDashboardSnapshot) SetDashboard(v ApiDashboard) {
	o.Dashboard = &v
}

// GetWidgets returns the Widgets field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetWidgets() []ApiWidget {
	if o == nil || utils.IsNil(o.Widgets) {
		var ret []ApiWidget
		return ret
	}
	return o.Widgets
}

// GetWidgetsOk returns a tuple with the Widgets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetWidgetsOk() ([]ApiWidget, bool) {
	if o == nil || utils.IsNil(o.Widgets) {
		return nil, false
	}
	return o.Widgets, true
}

// HasWidgets returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasWidgets() bool {
	if o != nil && !utils.IsNil(o.Widgets) {
		return true
	}

	return false
}

// SetWidgets gets a reference to the given []ApiWidget and assigns it to the Widgets field.
func (o *ApiDashboardSnapshot) SetWidgets(v []ApiWidget) {
	o.Widgets = v
}

// GetIsScheduled returns the IsScheduled field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetIsScheduled() bool {
	if o == nil || utils.IsNil(o.IsScheduled) {
		var ret bool
		return ret
	}
	return *o.IsScheduled
}

// GetIsScheduledOk returns a tuple with the IsScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetIsScheduledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsScheduled) {
		return nil, false
	}
	return o.IsScheduled, true
}

// HasIsScheduled returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasIsScheduled() bool {
	if o != nil && !utils.IsNil(o.IsScheduled) {
		return true
	}

	return false
}

// SetIsScheduled gets a reference to the given bool and assigns it to the IsScheduled field.
func (o *ApiDashboardSnapshot) SetIsScheduled(v bool) {
	o.IsScheduled = &v
}

// GetTimeSpan returns the TimeSpan field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetTimeSpan() ApiReportSnapshotTimeSpan {
	if o == nil || utils.IsNil(o.TimeSpan) {
		var ret ApiReportSnapshotTimeSpan
		return ret
	}
	return *o.TimeSpan
}

// GetTimeSpanOk returns a tuple with the TimeSpan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetTimeSpanOk() (*ApiReportSnapshotTimeSpan, bool) {
	if o == nil || utils.IsNil(o.TimeSpan) {
		return nil, false
	}
	return o.TimeSpan, true
}

// HasTimeSpan returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasTimeSpan() bool {
	if o != nil && !utils.IsNil(o.TimeSpan) {
		return true
	}

	return false
}

// SetTimeSpan gets a reference to the given ApiReportSnapshotTimeSpan and assigns it to the TimeSpan field.
func (o *ApiDashboardSnapshot) SetTimeSpan(v ApiReportSnapshotTimeSpan) {
	o.TimeSpan = &v
}

// GetSnapshotExpirationDate returns the SnapshotExpirationDate field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetSnapshotExpirationDate() time.Time {
	if o == nil || utils.IsNil(o.SnapshotExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.SnapshotExpirationDate
}

// GetSnapshotExpirationDateOk returns a tuple with the SnapshotExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetSnapshotExpirationDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.SnapshotExpirationDate) {
		return nil, false
	}
	return o.SnapshotExpirationDate, true
}

// HasSnapshotExpirationDate returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasSnapshotExpirationDate() bool {
	if o != nil && !utils.IsNil(o.SnapshotExpirationDate) {
		return true
	}

	return false
}

// SetSnapshotExpirationDate gets a reference to the given time.Time and assigns it to the SnapshotExpirationDate field.
func (o *ApiDashboardSnapshot) SetSnapshotExpirationDate(v time.Time) {
	o.SnapshotExpirationDate = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiDashboardSnapshot) GetLinks() AppAndSelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret AppAndSelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDashboardSnapshot) GetLinksOk() (*AppAndSelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiDashboardSnapshot) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAndSelfLinks and assigns it to the Links field.
func (o *ApiDashboardSnapshot) SetLinks(v AppAndSelfLinks) {
	o.Links = &v
}

func (o ApiDashboardSnapshot) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDashboardSnapshot) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !utils.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !utils.IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
	if !utils.IsNil(o.Permalink) {
		toSerialize["permalink"] = o.Permalink
	}
	if !utils.IsNil(o.ApiLinks) {
		toSerialize["apiLinks"] = o.ApiLinks
	}
	if !utils.IsNil(o.SnapshotId) {
		toSerialize["snapshotId"] = o.SnapshotId
	}
	if !utils.IsNil(o.SnapshotName) {
		toSerialize["snapshotName"] = o.SnapshotName
	}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.IsShared) {
		toSerialize["isShared"] = o.IsShared
	}
	if !utils.IsNil(o.SnapshotCreatedDate) {
		toSerialize["snapshotCreatedDate"] = o.SnapshotCreatedDate
	}
	if !utils.IsNil(o.Dashboard) {
		toSerialize["dashboard"] = o.Dashboard
	}
	if !utils.IsNil(o.Widgets) {
		toSerialize["widgets"] = o.Widgets
	}
	if !utils.IsNil(o.IsScheduled) {
		toSerialize["isScheduled"] = o.IsScheduled
	}
	if !utils.IsNil(o.TimeSpan) {
		toSerialize["timeSpan"] = o.TimeSpan
	}
	if !utils.IsNil(o.SnapshotExpirationDate) {
		toSerialize["snapshotExpirationDate"] = o.SnapshotExpirationDate
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableApiDashboardSnapshot struct {
	value *ApiDashboardSnapshot
	isSet bool
}

func (v NullableApiDashboardSnapshot) Get() *ApiDashboardSnapshot {
	return v.value
}

func (v *NullableApiDashboardSnapshot) Set(val *ApiDashboardSnapshot) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDashboardSnapshot) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDashboardSnapshot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDashboardSnapshot(val *ApiDashboardSnapshot) *NullableApiDashboardSnapshot {
	return &NullableApiDashboardSnapshot{value: val, isSet: true}
}

func (v NullableApiDashboardSnapshot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDashboardSnapshot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


