/*
Instant Tests API

The Instant Tests API operations lets you create and run new instant tests. You will need to be a regular user or have the following permissions:   * `API Access`   * `View tests`  The response does not include the immediate test results. Use the Test Results endpoints to get test results after creating and executing an instant test. You can find the URLs for these endpoints in the _links section of the test definition that is returned when you create the instant test. 

API version: 7.0.37
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package instanttests

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
	"fmt"
)

// checks if the DnsServerInstantTestResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DnsServerInstantTestResponse{}

// DnsServerInstantTestResponse struct for DnsServerInstantTestResponse
type DnsServerInstantTestResponse struct {
	// User that created the test.
	CreatedBy *string `json:"createdBy,omitempty"`
	// UTC created date (ISO date-time format).
	CreatedDate *time.Time `json:"createdDate,omitempty"`
	// A description of the test.
	Description *string `json:"description,omitempty"`
	// Indicates if the test is shared with the account group.
	LiveShare *bool `json:"liveShare,omitempty"`
	// User that modified the test.
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// UTC last modification date (ISO date-time format).
	ModifiedDate *time.Time `json:"modifiedDate,omitempty"`
	// Indicates if the test is a saved event.
	SavedEvent *bool `json:"savedEvent,omitempty"`
	// Each test is assigned an unique ID; this is used to access test information and results from other endpoints.
	TestId *string `json:"testId,omitempty"`
	// The name of the test. Test name must be unique.
	TestName *string `json:"testName,omitempty"`
	Type *string `json:"type,omitempty"`
	Links *TestLinks `json:"_links,omitempty"`
	Labels []TestLabel `json:"labels,omitempty"`
	SharedWithAccounts []SharedWithAccount `json:"sharedWithAccounts,omitempty"`
	// Set to `true` to enable bandwidth measurements, only applies to Enterprise agents assigned to the test.
	BandwidthMeasurements *bool `json:"bandwidthMeasurements,omitempty"`
	DnsServers []TestDnsServer `json:"dnsServers"`
	DnsTransportProtocol *TestDnsTransportProtocol `json:"dnsTransportProtocol,omitempty"`
	// The target record for the test, with the record type suffixed. If no record type is specified, the test defaults to an ANY record.
	Domain string `json:"domain"`
	// Set `true` to measure MTU sizes on network from agents to the target.
	MtuMeasurements *bool `json:"mtuMeasurements,omitempty"`
	// Enable or disable network measurements. Set to true to enable or false to disable network measurements.
	NetworkMeasurements *bool `json:"networkMeasurements,omitempty"`
	// Number of path traces executed by the agent.
	NumPathTraces *int32 `json:"numPathTraces,omitempty"`
	PathTraceMode *TestPathTraceMode `json:"pathTraceMode,omitempty"`
	ProbeMode *TestProbeMode `json:"probeMode,omitempty"`
	Protocol *TestProtocol `json:"protocol,omitempty"`
	// Indicates whether agents should randomize the start time in each test round.
	RandomizedStartTime *bool `json:"randomizedStartTime,omitempty"`
	// Set true to run query with RD (recursion desired) flag enabled.
	RecursiveQueries *bool `json:"recursiveQueries,omitempty"`
	Ipv6Policy *TestIpv6Policy `json:"ipv6Policy,omitempty"`
	// Sets packets rate sent to measure the network in packets per second.
	FixedPacketRate *int32 `json:"fixedPacketRate,omitempty"`
	DnsQueryClass *DnsQueryClass `json:"dnsQueryClass,omitempty"`
	// Contains list of agents.
	Agents []AgentResponse `json:"agents,omitempty"`
}

type _DnsServerInstantTestResponse DnsServerInstantTestResponse

// NewDnsServerInstantTestResponse instantiates a new DnsServerInstantTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsServerInstantTestResponse(dnsServers []TestDnsServer, domain string) *DnsServerInstantTestResponse {
	this := DnsServerInstantTestResponse{}
	this.DnsServers = dnsServers
	var dnsTransportProtocol TestDnsTransportProtocol = TESTDNSTRANSPORTPROTOCOL_UDP
	this.DnsTransportProtocol = &dnsTransportProtocol
	this.Domain = domain
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	var ipv6Policy TestIpv6Policy = TESTIPV6POLICY_USE_AGENT_POLICY
	this.Ipv6Policy = &ipv6Policy
	return &this
}

// NewDnsServerInstantTestResponseWithDefaults instantiates a new DnsServerInstantTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsServerInstantTestResponseWithDefaults() *DnsServerInstantTestResponse {
	this := DnsServerInstantTestResponse{}
	var dnsTransportProtocol TestDnsTransportProtocol = TESTDNSTRANSPORTPROTOCOL_UDP
	this.DnsTransportProtocol = &dnsTransportProtocol
	var networkMeasurements bool = true
	this.NetworkMeasurements = &networkMeasurements
	var numPathTraces int32 = 3
	this.NumPathTraces = &numPathTraces
	var pathTraceMode TestPathTraceMode = TESTPATHTRACEMODE_CLASSIC
	this.PathTraceMode = &pathTraceMode
	var probeMode TestProbeMode = TESTPROBEMODE_AUTO
	this.ProbeMode = &probeMode
	var protocol TestProtocol = TESTPROTOCOL_TCP
	this.Protocol = &protocol
	var randomizedStartTime bool = false
	this.RandomizedStartTime = &randomizedStartTime
	var ipv6Policy TestIpv6Policy = TESTIPV6POLICY_USE_AGENT_POLICY
	this.Ipv6Policy = &ipv6Policy
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetCreatedBy() string {
	if o == nil || utils.IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasCreatedBy() bool {
	if o != nil && !utils.IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *DnsServerInstantTestResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetCreatedDate() time.Time {
	if o == nil || utils.IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasCreatedDate() bool {
	if o != nil && !utils.IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *DnsServerInstantTestResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetDescription() string {
	if o == nil || utils.IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasDescription() bool {
	if o != nil && !utils.IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DnsServerInstantTestResponse) SetDescription(v string) {
	o.Description = &v
}

// GetLiveShare returns the LiveShare field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetLiveShare() bool {
	if o == nil || utils.IsNil(o.LiveShare) {
		var ret bool
		return ret
	}
	return *o.LiveShare
}

// GetLiveShareOk returns a tuple with the LiveShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetLiveShareOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.LiveShare) {
		return nil, false
	}
	return o.LiveShare, true
}

// HasLiveShare returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasLiveShare() bool {
	if o != nil && !utils.IsNil(o.LiveShare) {
		return true
	}

	return false
}

// SetLiveShare gets a reference to the given bool and assigns it to the LiveShare field.
func (o *DnsServerInstantTestResponse) SetLiveShare(v bool) {
	o.LiveShare = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetModifiedBy() string {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		var ret string
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetModifiedByOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ModifiedBy) {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasModifiedBy() bool {
	if o != nil && !utils.IsNil(o.ModifiedBy) {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given string and assigns it to the ModifiedBy field.
func (o *DnsServerInstantTestResponse) SetModifiedBy(v string) {
	o.ModifiedBy = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetModifiedDate() time.Time {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.ModifiedDate) {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasModifiedDate() bool {
	if o != nil && !utils.IsNil(o.ModifiedDate) {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *DnsServerInstantTestResponse) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetSavedEvent returns the SavedEvent field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetSavedEvent() bool {
	if o == nil || utils.IsNil(o.SavedEvent) {
		var ret bool
		return ret
	}
	return *o.SavedEvent
}

// GetSavedEventOk returns a tuple with the SavedEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetSavedEventOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.SavedEvent) {
		return nil, false
	}
	return o.SavedEvent, true
}

// HasSavedEvent returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasSavedEvent() bool {
	if o != nil && !utils.IsNil(o.SavedEvent) {
		return true
	}

	return false
}

// SetSavedEvent gets a reference to the given bool and assigns it to the SavedEvent field.
func (o *DnsServerInstantTestResponse) SetSavedEvent(v bool) {
	o.SavedEvent = &v
}

// GetTestId returns the TestId field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetTestId() string {
	if o == nil || utils.IsNil(o.TestId) {
		var ret string
		return ret
	}
	return *o.TestId
}

// GetTestIdOk returns a tuple with the TestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetTestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestId) {
		return nil, false
	}
	return o.TestId, true
}

// HasTestId returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasTestId() bool {
	if o != nil && !utils.IsNil(o.TestId) {
		return true
	}

	return false
}

// SetTestId gets a reference to the given string and assigns it to the TestId field.
func (o *DnsServerInstantTestResponse) SetTestId(v string) {
	o.TestId = &v
}

// GetTestName returns the TestName field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetTestName() string {
	if o == nil || utils.IsNil(o.TestName) {
		var ret string
		return ret
	}
	return *o.TestName
}

// GetTestNameOk returns a tuple with the TestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetTestNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TestName) {
		return nil, false
	}
	return o.TestName, true
}

// HasTestName returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasTestName() bool {
	if o != nil && !utils.IsNil(o.TestName) {
		return true
	}

	return false
}

// SetTestName gets a reference to the given string and assigns it to the TestName field.
func (o *DnsServerInstantTestResponse) SetTestName(v string) {
	o.TestName = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetType() string {
	if o == nil || utils.IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DnsServerInstantTestResponse) SetType(v string) {
	o.Type = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetLinks() TestLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret TestLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetLinksOk() (*TestLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given TestLinks and assigns it to the Links field.
func (o *DnsServerInstantTestResponse) SetLinks(v TestLinks) {
	o.Links = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetLabels() []TestLabel {
	if o == nil || utils.IsNil(o.Labels) {
		var ret []TestLabel
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetLabelsOk() ([]TestLabel, bool) {
	if o == nil || utils.IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []TestLabel and assigns it to the Labels field.
func (o *DnsServerInstantTestResponse) SetLabels(v []TestLabel) {
	o.Labels = v
}

// GetSharedWithAccounts returns the SharedWithAccounts field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetSharedWithAccounts() []SharedWithAccount {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		var ret []SharedWithAccount
		return ret
	}
	return o.SharedWithAccounts
}

// GetSharedWithAccountsOk returns a tuple with the SharedWithAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetSharedWithAccountsOk() ([]SharedWithAccount, bool) {
	if o == nil || utils.IsNil(o.SharedWithAccounts) {
		return nil, false
	}
	return o.SharedWithAccounts, true
}

// HasSharedWithAccounts returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasSharedWithAccounts() bool {
	if o != nil && !utils.IsNil(o.SharedWithAccounts) {
		return true
	}

	return false
}

// SetSharedWithAccounts gets a reference to the given []SharedWithAccount and assigns it to the SharedWithAccounts field.
func (o *DnsServerInstantTestResponse) SetSharedWithAccounts(v []SharedWithAccount) {
	o.SharedWithAccounts = v
}

// GetBandwidthMeasurements returns the BandwidthMeasurements field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetBandwidthMeasurements() bool {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		var ret bool
		return ret
	}
	return *o.BandwidthMeasurements
}

// GetBandwidthMeasurementsOk returns a tuple with the BandwidthMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetBandwidthMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.BandwidthMeasurements) {
		return nil, false
	}
	return o.BandwidthMeasurements, true
}

// HasBandwidthMeasurements returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasBandwidthMeasurements() bool {
	if o != nil && !utils.IsNil(o.BandwidthMeasurements) {
		return true
	}

	return false
}

// SetBandwidthMeasurements gets a reference to the given bool and assigns it to the BandwidthMeasurements field.
func (o *DnsServerInstantTestResponse) SetBandwidthMeasurements(v bool) {
	o.BandwidthMeasurements = &v
}

// GetDnsServers returns the DnsServers field value
func (o *DnsServerInstantTestResponse) GetDnsServers() []TestDnsServer {
	if o == nil {
		var ret []TestDnsServer
		return ret
	}

	return o.DnsServers
}

// GetDnsServersOk returns a tuple with the DnsServers field value
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetDnsServersOk() ([]TestDnsServer, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnsServers, true
}

// SetDnsServers sets field value
func (o *DnsServerInstantTestResponse) SetDnsServers(v []TestDnsServer) {
	o.DnsServers = v
}

// GetDnsTransportProtocol returns the DnsTransportProtocol field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetDnsTransportProtocol() TestDnsTransportProtocol {
	if o == nil || utils.IsNil(o.DnsTransportProtocol) {
		var ret TestDnsTransportProtocol
		return ret
	}
	return *o.DnsTransportProtocol
}

// GetDnsTransportProtocolOk returns a tuple with the DnsTransportProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetDnsTransportProtocolOk() (*TestDnsTransportProtocol, bool) {
	if o == nil || utils.IsNil(o.DnsTransportProtocol) {
		return nil, false
	}
	return o.DnsTransportProtocol, true
}

// HasDnsTransportProtocol returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasDnsTransportProtocol() bool {
	if o != nil && !utils.IsNil(o.DnsTransportProtocol) {
		return true
	}

	return false
}

// SetDnsTransportProtocol gets a reference to the given TestDnsTransportProtocol and assigns it to the DnsTransportProtocol field.
func (o *DnsServerInstantTestResponse) SetDnsTransportProtocol(v TestDnsTransportProtocol) {
	o.DnsTransportProtocol = &v
}

// GetDomain returns the Domain field value
func (o *DnsServerInstantTestResponse) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DnsServerInstantTestResponse) SetDomain(v string) {
	o.Domain = v
}

// GetMtuMeasurements returns the MtuMeasurements field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetMtuMeasurements() bool {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		var ret bool
		return ret
	}
	return *o.MtuMeasurements
}

// GetMtuMeasurementsOk returns a tuple with the MtuMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetMtuMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.MtuMeasurements) {
		return nil, false
	}
	return o.MtuMeasurements, true
}

// HasMtuMeasurements returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasMtuMeasurements() bool {
	if o != nil && !utils.IsNil(o.MtuMeasurements) {
		return true
	}

	return false
}

// SetMtuMeasurements gets a reference to the given bool and assigns it to the MtuMeasurements field.
func (o *DnsServerInstantTestResponse) SetMtuMeasurements(v bool) {
	o.MtuMeasurements = &v
}

// GetNetworkMeasurements returns the NetworkMeasurements field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetNetworkMeasurements() bool {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		var ret bool
		return ret
	}
	return *o.NetworkMeasurements
}

// GetNetworkMeasurementsOk returns a tuple with the NetworkMeasurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetNetworkMeasurementsOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.NetworkMeasurements) {
		return nil, false
	}
	return o.NetworkMeasurements, true
}

// HasNetworkMeasurements returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasNetworkMeasurements() bool {
	if o != nil && !utils.IsNil(o.NetworkMeasurements) {
		return true
	}

	return false
}

// SetNetworkMeasurements gets a reference to the given bool and assigns it to the NetworkMeasurements field.
func (o *DnsServerInstantTestResponse) SetNetworkMeasurements(v bool) {
	o.NetworkMeasurements = &v
}

// GetNumPathTraces returns the NumPathTraces field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetNumPathTraces() int32 {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		var ret int32
		return ret
	}
	return *o.NumPathTraces
}

// GetNumPathTracesOk returns a tuple with the NumPathTraces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetNumPathTracesOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.NumPathTraces) {
		return nil, false
	}
	return o.NumPathTraces, true
}

// HasNumPathTraces returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasNumPathTraces() bool {
	if o != nil && !utils.IsNil(o.NumPathTraces) {
		return true
	}

	return false
}

// SetNumPathTraces gets a reference to the given int32 and assigns it to the NumPathTraces field.
func (o *DnsServerInstantTestResponse) SetNumPathTraces(v int32) {
	o.NumPathTraces = &v
}

// GetPathTraceMode returns the PathTraceMode field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetPathTraceMode() TestPathTraceMode {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		var ret TestPathTraceMode
		return ret
	}
	return *o.PathTraceMode
}

// GetPathTraceModeOk returns a tuple with the PathTraceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetPathTraceModeOk() (*TestPathTraceMode, bool) {
	if o == nil || utils.IsNil(o.PathTraceMode) {
		return nil, false
	}
	return o.PathTraceMode, true
}

// HasPathTraceMode returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasPathTraceMode() bool {
	if o != nil && !utils.IsNil(o.PathTraceMode) {
		return true
	}

	return false
}

// SetPathTraceMode gets a reference to the given TestPathTraceMode and assigns it to the PathTraceMode field.
func (o *DnsServerInstantTestResponse) SetPathTraceMode(v TestPathTraceMode) {
	o.PathTraceMode = &v
}

// GetProbeMode returns the ProbeMode field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetProbeMode() TestProbeMode {
	if o == nil || utils.IsNil(o.ProbeMode) {
		var ret TestProbeMode
		return ret
	}
	return *o.ProbeMode
}

// GetProbeModeOk returns a tuple with the ProbeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetProbeModeOk() (*TestProbeMode, bool) {
	if o == nil || utils.IsNil(o.ProbeMode) {
		return nil, false
	}
	return o.ProbeMode, true
}

// HasProbeMode returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasProbeMode() bool {
	if o != nil && !utils.IsNil(o.ProbeMode) {
		return true
	}

	return false
}

// SetProbeMode gets a reference to the given TestProbeMode and assigns it to the ProbeMode field.
func (o *DnsServerInstantTestResponse) SetProbeMode(v TestProbeMode) {
	o.ProbeMode = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetProtocol() TestProtocol {
	if o == nil || utils.IsNil(o.Protocol) {
		var ret TestProtocol
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetProtocolOk() (*TestProtocol, bool) {
	if o == nil || utils.IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasProtocol() bool {
	if o != nil && !utils.IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given TestProtocol and assigns it to the Protocol field.
func (o *DnsServerInstantTestResponse) SetProtocol(v TestProtocol) {
	o.Protocol = &v
}

// GetRandomizedStartTime returns the RandomizedStartTime field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetRandomizedStartTime() bool {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		var ret bool
		return ret
	}
	return *o.RandomizedStartTime
}

// GetRandomizedStartTimeOk returns a tuple with the RandomizedStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetRandomizedStartTimeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RandomizedStartTime) {
		return nil, false
	}
	return o.RandomizedStartTime, true
}

// HasRandomizedStartTime returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasRandomizedStartTime() bool {
	if o != nil && !utils.IsNil(o.RandomizedStartTime) {
		return true
	}

	return false
}

// SetRandomizedStartTime gets a reference to the given bool and assigns it to the RandomizedStartTime field.
func (o *DnsServerInstantTestResponse) SetRandomizedStartTime(v bool) {
	o.RandomizedStartTime = &v
}

// GetRecursiveQueries returns the RecursiveQueries field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetRecursiveQueries() bool {
	if o == nil || utils.IsNil(o.RecursiveQueries) {
		var ret bool
		return ret
	}
	return *o.RecursiveQueries
}

// GetRecursiveQueriesOk returns a tuple with the RecursiveQueries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetRecursiveQueriesOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.RecursiveQueries) {
		return nil, false
	}
	return o.RecursiveQueries, true
}

// HasRecursiveQueries returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasRecursiveQueries() bool {
	if o != nil && !utils.IsNil(o.RecursiveQueries) {
		return true
	}

	return false
}

// SetRecursiveQueries gets a reference to the given bool and assigns it to the RecursiveQueries field.
func (o *DnsServerInstantTestResponse) SetRecursiveQueries(v bool) {
	o.RecursiveQueries = &v
}

// GetIpv6Policy returns the Ipv6Policy field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetIpv6Policy() TestIpv6Policy {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		var ret TestIpv6Policy
		return ret
	}
	return *o.Ipv6Policy
}

// GetIpv6PolicyOk returns a tuple with the Ipv6Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetIpv6PolicyOk() (*TestIpv6Policy, bool) {
	if o == nil || utils.IsNil(o.Ipv6Policy) {
		return nil, false
	}
	return o.Ipv6Policy, true
}

// HasIpv6Policy returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasIpv6Policy() bool {
	if o != nil && !utils.IsNil(o.Ipv6Policy) {
		return true
	}

	return false
}

// SetIpv6Policy gets a reference to the given TestIpv6Policy and assigns it to the Ipv6Policy field.
func (o *DnsServerInstantTestResponse) SetIpv6Policy(v TestIpv6Policy) {
	o.Ipv6Policy = &v
}

// GetFixedPacketRate returns the FixedPacketRate field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetFixedPacketRate() int32 {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		var ret int32
		return ret
	}
	return *o.FixedPacketRate
}

// GetFixedPacketRateOk returns a tuple with the FixedPacketRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetFixedPacketRateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.FixedPacketRate) {
		return nil, false
	}
	return o.FixedPacketRate, true
}

// HasFixedPacketRate returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasFixedPacketRate() bool {
	if o != nil && !utils.IsNil(o.FixedPacketRate) {
		return true
	}

	return false
}

// SetFixedPacketRate gets a reference to the given int32 and assigns it to the FixedPacketRate field.
func (o *DnsServerInstantTestResponse) SetFixedPacketRate(v int32) {
	o.FixedPacketRate = &v
}

// GetDnsQueryClass returns the DnsQueryClass field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetDnsQueryClass() DnsQueryClass {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		var ret DnsQueryClass
		return ret
	}
	return *o.DnsQueryClass
}

// GetDnsQueryClassOk returns a tuple with the DnsQueryClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetDnsQueryClassOk() (*DnsQueryClass, bool) {
	if o == nil || utils.IsNil(o.DnsQueryClass) {
		return nil, false
	}
	return o.DnsQueryClass, true
}

// HasDnsQueryClass returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasDnsQueryClass() bool {
	if o != nil && !utils.IsNil(o.DnsQueryClass) {
		return true
	}

	return false
}

// SetDnsQueryClass gets a reference to the given DnsQueryClass and assigns it to the DnsQueryClass field.
func (o *DnsServerInstantTestResponse) SetDnsQueryClass(v DnsQueryClass) {
	o.DnsQueryClass = &v
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *DnsServerInstantTestResponse) GetAgents() []AgentResponse {
	if o == nil || utils.IsNil(o.Agents) {
		var ret []AgentResponse
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsServerInstantTestResponse) GetAgentsOk() ([]AgentResponse, bool) {
	if o == nil || utils.IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *DnsServerInstantTestResponse) HasAgents() bool {
	if o != nil && !utils.IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []AgentResponse and assigns it to the Agents field.
func (o *DnsServerInstantTestResponse) SetAgents(v []AgentResponse) {
	o.Agents = v
}

func (o DnsServerInstantTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsServerInstantTestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !utils.IsNil(o.CreatedDate) {
		toSerialize["createdDate"] = o.CreatedDate
	}
	if !utils.IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !utils.IsNil(o.LiveShare) {
		toSerialize["liveShare"] = o.LiveShare
	}
	if !utils.IsNil(o.ModifiedBy) {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if !utils.IsNil(o.ModifiedDate) {
		toSerialize["modifiedDate"] = o.ModifiedDate
	}
	if !utils.IsNil(o.SavedEvent) {
		toSerialize["savedEvent"] = o.SavedEvent
	}
	if !utils.IsNil(o.TestId) {
		toSerialize["testId"] = o.TestId
	}
	if !utils.IsNil(o.TestName) {
		toSerialize["testName"] = o.TestName
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !utils.IsNil(o.SharedWithAccounts) {
		toSerialize["sharedWithAccounts"] = o.SharedWithAccounts
	}
	if !utils.IsNil(o.BandwidthMeasurements) {
		toSerialize["bandwidthMeasurements"] = o.BandwidthMeasurements
	}
	toSerialize["dnsServers"] = o.DnsServers
	if !utils.IsNil(o.DnsTransportProtocol) {
		toSerialize["dnsTransportProtocol"] = o.DnsTransportProtocol
	}
	toSerialize["domain"] = o.Domain
	if !utils.IsNil(o.MtuMeasurements) {
		toSerialize["mtuMeasurements"] = o.MtuMeasurements
	}
	if !utils.IsNil(o.NetworkMeasurements) {
		toSerialize["networkMeasurements"] = o.NetworkMeasurements
	}
	if !utils.IsNil(o.NumPathTraces) {
		toSerialize["numPathTraces"] = o.NumPathTraces
	}
	if !utils.IsNil(o.PathTraceMode) {
		toSerialize["pathTraceMode"] = o.PathTraceMode
	}
	if !utils.IsNil(o.ProbeMode) {
		toSerialize["probeMode"] = o.ProbeMode
	}
	if !utils.IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !utils.IsNil(o.RandomizedStartTime) {
		toSerialize["randomizedStartTime"] = o.RandomizedStartTime
	}
	if !utils.IsNil(o.RecursiveQueries) {
		toSerialize["recursiveQueries"] = o.RecursiveQueries
	}
	if !utils.IsNil(o.Ipv6Policy) {
		toSerialize["ipv6Policy"] = o.Ipv6Policy
	}
	if !utils.IsNil(o.FixedPacketRate) {
		toSerialize["fixedPacketRate"] = o.FixedPacketRate
	}
	if !utils.IsNil(o.DnsQueryClass) {
		toSerialize["dnsQueryClass"] = o.DnsQueryClass
	}
	if !utils.IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	return toSerialize, nil
}

func (o *DnsServerInstantTestResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dnsServers",
		"domain",
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

	varDnsServerInstantTestResponse := _DnsServerInstantTestResponse{}

    err = json.Unmarshal(data, &varDnsServerInstantTestResponse)

	if err != nil {
		return err
	}

	*o = DnsServerInstantTestResponse(varDnsServerInstantTestResponse)

	return err
}

type NullableDnsServerInstantTestResponse struct {
	value *DnsServerInstantTestResponse
	isSet bool
}

func (v NullableDnsServerInstantTestResponse) Get() *DnsServerInstantTestResponse {
	return v.value
}

func (v *NullableDnsServerInstantTestResponse) Set(val *DnsServerInstantTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsServerInstantTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsServerInstantTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsServerInstantTestResponse(val *DnsServerInstantTestResponse) *NullableDnsServerInstantTestResponse {
	return &NullableDnsServerInstantTestResponse{value: val, isSet: true}
}

func (v NullableDnsServerInstantTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsServerInstantTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


