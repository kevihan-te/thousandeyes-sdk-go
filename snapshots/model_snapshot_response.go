/*
Test Snapshots API

Creates a new test snapshot in ThousandEyes.

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package snapshots

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the SnapshotResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SnapshotResponse{}

// SnapshotResponse struct for SnapshotResponse
type SnapshotResponse struct {
	// Snapshot ID.
	Id *string `json:"id,omitempty"`
	// The start time of the test snapshot, represented in epoch time format (in seconds).
	StartRoundId *int32 `json:"startRoundId,omitempty"`
	// The end time of the test snapshot, represented in epoch time format (in seconds).
	EndRoundId *int32 `json:"endRoundId,omitempty"`
	// The selected time of the test snapshot, represented in epoch time format (in seconds).
	RoundId *int32 `json:"roundId,omitempty"`
	// The date when the test snapshot was created in UTC time, formatted in ISO date-time.
	ShareDate *time.Time `json:"shareDate,omitempty"`
	// Snapshot test ID.
	SourceTestId *string `json:"sourceTestId,omitempty"`
	// Snapshot test ID.
	TestId *string `json:"testId,omitempty"`
	// User ID.
	Uid *string `json:"uid,omitempty"`
	// Snapshot title.
	DisplayName *string `json:"displayName,omitempty"`
	// Extra parameters.
	ExtraParams *string `json:"extraParams,omitempty"`
	Test *SnapshotTest `json:"test,omitempty"`
	Links *SnapshotLinks `json:"_links,omitempty"`
}

// NewSnapshotResponse instantiates a new SnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotResponse() *SnapshotResponse {
	this := SnapshotResponse{}
	return &this
}

// NewSnapshotResponseWithDefaults instantiates a new SnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotResponseWithDefaults() *SnapshotResponse {
	this := SnapshotResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SnapshotResponse) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SnapshotResponse) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SnapshotResponse) SetId(v string) {
	o.Id = &v
}

// GetStartRoundId returns the StartRoundId field value if set, zero value otherwise.
func (o *SnapshotResponse) GetStartRoundId() int32 {
	if o == nil || utils.IsNil(o.StartRoundId) {
		var ret int32
		return ret
	}
	return *o.StartRoundId
}

// GetStartRoundIdOk returns a tuple with the StartRoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetStartRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.StartRoundId) {
		return nil, false
	}
	return o.StartRoundId, true
}

// HasStartRoundId returns a boolean if a field has been set.
func (o *SnapshotResponse) HasStartRoundId() bool {
	if o != nil && !utils.IsNil(o.StartRoundId) {
		return true
	}

	return false
}

// SetStartRoundId gets a reference to the given int32 and assigns it to the StartRoundId field.
func (o *SnapshotResponse) SetStartRoundId(v int32) {
	o.StartRoundId = &v
}

// GetEndRoundId returns the EndRoundId field value if set, zero value otherwise.
func (o *SnapshotResponse) GetEndRoundId() int32 {
	if o == nil || utils.IsNil(o.EndRoundId) {
		var ret int32
		return ret
	}
	return *o.EndRoundId
}

// GetEndRoundIdOk returns a tuple with the EndRoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetEndRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.EndRoundId) {
		return nil, false
	}
	return o.EndRoundId, true
}

// HasEndRoundId returns a boolean if a field has been set.
func (o *SnapshotResponse) HasEndRoundId() bool {
	if o != nil && !utils.IsNil(o.EndRoundId) {
		return true
	}

	return false
}

// SetEndRoundId gets a reference to the given int32 and assigns it to the EndRoundId field.
func (o *SnapshotResponse) SetEndRoundId(v int32) {
	o.EndRoundId = &v
}

// GetRoundId returns the RoundId field value if set, zero value otherwise.
func (o *SnapshotResponse) GetRoundId() int32 {
	if o == nil || utils.IsNil(o.RoundId) {
		var ret int32
		return ret
	}
	return *o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetRoundIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.RoundId) {
		return nil, false
	}
	return o.RoundId, true
}

// HasRoundId returns a boolean if a field has been set.
func (o *SnapshotResponse) HasRoundId() bool {
	if o != nil && !utils.IsNil(o.RoundId) {
		return true
	}

	return false
}

// SetRoundId gets a reference to the given int32 and assigns it to the RoundId field.
func (o *SnapshotResponse) SetRoundId(v int32) {
	o.RoundId = &v
}

// GetShareDate returns the ShareDate field value if set, zero value otherwise.
func (o *SnapshotResponse) GetShareDate() time.Time {
	if o == nil || utils.IsNil(o.ShareDate) {
		var ret time.Time
		return ret
	}
	return *o.ShareDate
}

// GetShareDateOk returns a tuple with the ShareDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetShareDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ShareDate) {
		return nil, false
	}
	return o.ShareDate, true
}

// HasShareDate returns a boolean if a field has been set.
func (o *SnapshotResponse) HasShareDate() bool {
	if o != nil && !utils.IsNil(o.ShareDate) {
		return true
	}

	return false
}

// SetShareDate gets a reference to the given time.Time and assigns it to the ShareDate field.
func (o *SnapshotResponse) SetShareDate(v time.Time) {
	o.ShareDate = &v
}

// GetSourceTestId returns the SourceTestId field value if set, zero value otherwise.
func (o *SnapshotResponse) GetSourceTestId() string {
	if o == nil || utils.IsNil(o.SourceTestId) {
		var ret string
		return ret
	}
	return *o.SourceTestId
}

// GetSourceTestIdOk returns a tuple with the SourceTestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetSourceTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SourceTestId) {
		return nil, false
	}
	return o.SourceTestId, true
}

// HasSourceTestId returns a boolean if a field has been set.
func (o *SnapshotResponse) HasSourceTestId() bool {
	if o != nil && !utils.IsNil(o.SourceTestId) {
		return true
	}

	return false
}

// SetSourceTestId gets a reference to the given string and assigns it to the SourceTestId field.
func (o *SnapshotResponse) SetSourceTestId(v string) {
	o.SourceTestId = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *SnapshotResponse) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *SnapshotResponse) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *SnapshotResponse) SetTestId(v string) {
	o.TestId = &v
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *SnapshotResponse) GetUid() string {
	if o == nil || utils.IsNil(o.Uid) {
		var ret string
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetUidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *SnapshotResponse) HasUid() bool {
	if o != nil && !utils.IsNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given string and assigns it to the Uid field.
func (o *SnapshotResponse) SetUid(v string) {
	o.Uid = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SnapshotResponse) GetDisplayName() string {
	if o == nil || utils.IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetDisplayNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SnapshotResponse) HasDisplayName() bool {
	if o != nil && !utils.IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SnapshotResponse) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtraParams returns the ExtraParams field value if set, zero value otherwise.
func (o *SnapshotResponse) GetExtraParams() string {
	if o == nil || utils.IsNil(o.ExtraParams) {
		var ret string
		return ret
	}
	return *o.ExtraParams
}

// GetExtraParamsOk returns a tuple with the ExtraParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetExtraParamsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ExtraParams) {
		return nil, false
	}
	return o.ExtraParams, true
}

// HasExtraParams returns a boolean if a field has been set.
func (o *SnapshotResponse) HasExtraParams() bool {
	if o != nil && !utils.IsNil(o.ExtraParams) {
		return true
	}

	return false
}

// SetExtraParams gets a reference to the given string and assigns it to the ExtraParams field.
func (o *SnapshotResponse) SetExtraParams(v string) {
	o.ExtraParams = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *SnapshotResponse) GetTest() SnapshotTest {
	if o == nil || utils.IsNil(o.Test) {
		var ret SnapshotTest
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetTestOk() (*SnapshotTest, bool) {
	if o == nil || utils.IsNil(o.Test) {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *SnapshotResponse) HasTest() bool {
	if o != nil && !utils.IsNil(o.Test) {
		return true
	}

	return false
}

// SetTest gets a reference to the given SnapshotTest and assigns it to the Test field.
func (o *SnapshotResponse) SetTest(v SnapshotTest) {
	o.Test = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SnapshotResponse) GetLinks() SnapshotLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SnapshotLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotResponse) GetLinksOk() (*SnapshotLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SnapshotResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SnapshotLinks and assigns it to the Links field.
func (o *SnapshotResponse) SetLinks(v SnapshotLinks) {
	o.Links = &v
}

func (o SnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.StartRoundId) {
		toSerialize["startRoundId"] = o.StartRoundId
	}
	if !utils.IsNil(o.EndRoundId) {
		toSerialize["endRoundId"] = o.EndRoundId
	}
	if !utils.IsNil(o.RoundId) {
		toSerialize["roundId"] = o.RoundId
	}
	if !utils.IsNil(o.ShareDate) {
		toSerialize["shareDate"] = o.ShareDate
	}
	if !utils.IsNil(o.SourceTestId) {
		toSerialize["sourceTestId"] = o.SourceTestId
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !utils.IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !utils.IsNil(o.ExtraParams) {
		toSerialize["extraParams"] = o.ExtraParams
	}
	if !utils.IsNil(o.Test) {
		toSerialize["test"] = o.Test
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableSnapshotResponse struct {
	value *SnapshotResponse
	isSet bool
}

func (v NullableSnapshotResponse) Get() *SnapshotResponse {
	return v.value
}

func (v *NullableSnapshotResponse) Set(val *SnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotResponse(val *SnapshotResponse) *NullableSnapshotResponse {
	return &NullableSnapshotResponse{value: val, isSet: true}
}

func (v NullableSnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


