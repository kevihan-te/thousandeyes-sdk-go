/*
Endpoint Agents API

Manage ThousandEyes Endpoint Agents using this API.   For more information about Endpoint Agents, see [Endpoint Agents](https://docs.thousandeyes.com/product-documentation/global-vantage-points/endpoint-agents).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package endpointagents

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
	"time"
)

// checks if the EndpointAgent type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EndpointAgent{}

// EndpointAgent The `EndpointAgent` object, which may include multiple clients.
type EndpointAgent struct {
	// Unique ID of endpoint agent, from `/endpoint/agents` endpoint.
	Id *string `json:"id,omitempty"`
	Aid *string `json:"aid,omitempty"`
	// The name of the agent.
	Name *string `json:"name,omitempty"`
	ComputerName *string `json:"computerName,omitempty"`
	OsVersion *string `json:"osVersion,omitempty"`
	Platform *Platform `json:"platform,omitempty"`
	KernelVersion *string `json:"kernelVersion,omitempty"`
	Manufacturer *string `json:"manufacturer,omitempty"`
	Model *string `json:"model,omitempty"`
	// The last time the agent checked-in.
	LastSeen *time.Time `json:"lastSeen,omitempty"`
	Status *Status `json:"status,omitempty"`
	Deleted *bool `json:"deleted,omitempty"`
	// Version of the agent software running.
	Version *string `json:"version,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	NumberOfClients *int64 `json:"numberOfClients,omitempty"`
	PublicIP *string `json:"publicIP,omitempty"`
	Location *EndpointAgentLocation `json:"location,omitempty"`
	// List of clients (user accounts) that the agent works with. Not populated by default. 
	Clients []EndpointClient `json:"clients,omitempty"`
	TotalMemory *string `json:"totalMemory,omitempty"`
	AgentType *string `json:"agentType,omitempty"`
	// List of VPN connections on the agent. Not populated by default. 
	VpnProfiles []EndpointVpnProfile `json:"vpnProfiles,omitempty"`
	// List of network interfaces on the agent. Not populated by default. 
	NetworkInterfaceProfiles []InterfaceProfile `json:"networkInterfaceProfiles,omitempty"`
	AsnDetails *EndpointAsnDetails `json:"asnDetails,omitempty"`
	LicenseType *AgentLicenseType `json:"licenseType,omitempty"`
	// Status of TCP test support on the agent.
	TcpDriverAvailable *bool `json:"tcpDriverAvailable,omitempty"`
	// For Windows agents, the version of the NPCAP driver that the agent has loaded.
	NpcapVersion *string `json:"npcapVersion,omitempty"`
	Links *SelfLinks `json:"_links,omitempty"`
}

// NewEndpointAgent instantiates a new EndpointAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointAgent() *EndpointAgent {
	this := EndpointAgent{}
	return &this
}

// NewEndpointAgentWithDefaults instantiates a new EndpointAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointAgentWithDefaults() *EndpointAgent {
	this := EndpointAgent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EndpointAgent) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EndpointAgent) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EndpointAgent) SetId(v string) {
	o.Id = &v
}

// GetAid returns the Aid field value if set, zero value otherwise.
func (o *EndpointAgent) GetAid() string {
	if o == nil || utils.IsNil(o.Aid) {
		var ret string
		return ret
	}
	return *o.Aid
}

// GetAidOk returns a tuple with the Aid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetAidOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Aid) {
		return nil, false
	}
	return o.Aid, true
}

// HasAid returns a boolean if a field has been set.
func (o *EndpointAgent) HasAid() bool {
	if o != nil && !utils.IsNil(o.Aid) {
		return true
	}

	return false
}

// SetAid gets a reference to the given string and assigns it to the Aid field.
func (o *EndpointAgent) SetAid(v string) {
	o.Aid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EndpointAgent) GetName() string {
	if o == nil || utils.IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EndpointAgent) HasName() bool {
	if o != nil && !utils.IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EndpointAgent) SetName(v string) {
	o.Name = &v
}

// GetComputerName returns the ComputerName field value if set, zero value otherwise.
func (o *EndpointAgent) GetComputerName() string {
	if o == nil || utils.IsNil(o.ComputerName) {
		var ret string
		return ret
	}
	return *o.ComputerName
}

// GetComputerNameOk returns a tuple with the ComputerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetComputerNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ComputerName) {
		return nil, false
	}
	return o.ComputerName, true
}

// HasComputerName returns a boolean if a field has been set.
func (o *EndpointAgent) HasComputerName() bool {
	if o != nil && !utils.IsNil(o.ComputerName) {
		return true
	}

	return false
}

// SetComputerName gets a reference to the given string and assigns it to the ComputerName field.
func (o *EndpointAgent) SetComputerName(v string) {
	o.ComputerName = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *EndpointAgent) GetOsVersion() string {
	if o == nil || utils.IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetOsVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *EndpointAgent) HasOsVersion() bool {
	if o != nil && !utils.IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *EndpointAgent) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *EndpointAgent) GetPlatform() Platform {
	if o == nil || utils.IsNil(o.Platform) {
		var ret Platform
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetPlatformOk() (*Platform, bool) {
	if o == nil || utils.IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *EndpointAgent) HasPlatform() bool {
	if o != nil && !utils.IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given Platform and assigns it to the Platform field.
func (o *EndpointAgent) SetPlatform(v Platform) {
	o.Platform = &v
}

// GetKernelVersion returns the KernelVersion field value if set, zero value otherwise.
func (o *EndpointAgent) GetKernelVersion() string {
	if o == nil || utils.IsNil(o.KernelVersion) {
		var ret string
		return ret
	}
	return *o.KernelVersion
}

// GetKernelVersionOk returns a tuple with the KernelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetKernelVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.KernelVersion) {
		return nil, false
	}
	return o.KernelVersion, true
}

// HasKernelVersion returns a boolean if a field has been set.
func (o *EndpointAgent) HasKernelVersion() bool {
	if o != nil && !utils.IsNil(o.KernelVersion) {
		return true
	}

	return false
}

// SetKernelVersion gets a reference to the given string and assigns it to the KernelVersion field.
func (o *EndpointAgent) SetKernelVersion(v string) {
	o.KernelVersion = &v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise.
func (o *EndpointAgent) GetManufacturer() string {
	if o == nil || utils.IsNil(o.Manufacturer) {
		var ret string
		return ret
	}
	return *o.Manufacturer
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetManufacturerOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Manufacturer) {
		return nil, false
	}
	return o.Manufacturer, true
}

// HasManufacturer returns a boolean if a field has been set.
func (o *EndpointAgent) HasManufacturer() bool {
	if o != nil && !utils.IsNil(o.Manufacturer) {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given string and assigns it to the Manufacturer field.
func (o *EndpointAgent) SetManufacturer(v string) {
	o.Manufacturer = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *EndpointAgent) GetModel() string {
	if o == nil || utils.IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetModelOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *EndpointAgent) HasModel() bool {
	if o != nil && !utils.IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *EndpointAgent) SetModel(v string) {
	o.Model = &v
}

// GetLastSeen returns the LastSeen field value if set, zero value otherwise.
func (o *EndpointAgent) GetLastSeen() time.Time {
	if o == nil || utils.IsNil(o.LastSeen) {
		var ret time.Time
		return ret
	}
	return *o.LastSeen
}

// GetLastSeenOk returns a tuple with the LastSeen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetLastSeenOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.LastSeen) {
		return nil, false
	}
	return o.LastSeen, true
}

// HasLastSeen returns a boolean if a field has been set.
func (o *EndpointAgent) HasLastSeen() bool {
	if o != nil && !utils.IsNil(o.LastSeen) {
		return true
	}

	return false
}

// SetLastSeen gets a reference to the given time.Time and assigns it to the LastSeen field.
func (o *EndpointAgent) SetLastSeen(v time.Time) {
	o.LastSeen = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EndpointAgent) GetStatus() Status {
	if o == nil || utils.IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetStatusOk() (*Status, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EndpointAgent) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *EndpointAgent) SetStatus(v Status) {
	o.Status = &v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *EndpointAgent) GetDeleted() bool {
	if o == nil || utils.IsNil(o.Deleted) {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetDeletedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Deleted) {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *EndpointAgent) HasDeleted() bool {
	if o != nil && !utils.IsNil(o.Deleted) {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *EndpointAgent) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *EndpointAgent) GetVersion() string {
	if o == nil || utils.IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *EndpointAgent) HasVersion() bool {
	if o != nil && !utils.IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *EndpointAgent) SetVersion(v string) {
	o.Version = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EndpointAgent) GetCreatedAt() time.Time {
	if o == nil || utils.IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || utils.IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EndpointAgent) HasCreatedAt() bool {
	if o != nil && !utils.IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EndpointAgent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetNumberOfClients returns the NumberOfClients field value if set, zero value otherwise.
func (o *EndpointAgent) GetNumberOfClients() int64 {
	if o == nil || utils.IsNil(o.NumberOfClients) {
		var ret int64
		return ret
	}
	return *o.NumberOfClients
}

// GetNumberOfClientsOk returns a tuple with the NumberOfClients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetNumberOfClientsOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.NumberOfClients) {
		return nil, false
	}
	return o.NumberOfClients, true
}

// HasNumberOfClients returns a boolean if a field has been set.
func (o *EndpointAgent) HasNumberOfClients() bool {
	if o != nil && !utils.IsNil(o.NumberOfClients) {
		return true
	}

	return false
}

// SetNumberOfClients gets a reference to the given int64 and assigns it to the NumberOfClients field.
func (o *EndpointAgent) SetNumberOfClients(v int64) {
	o.NumberOfClients = &v
}

// GetPublicIP returns the PublicIP field value if set, zero value otherwise.
func (o *EndpointAgent) GetPublicIP() string {
	if o == nil || utils.IsNil(o.PublicIP) {
		var ret string
		return ret
	}
	return *o.PublicIP
}

// GetPublicIPOk returns a tuple with the PublicIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetPublicIPOk() (*string, bool) {
	if o == nil || utils.IsNil(o.PublicIP) {
		return nil, false
	}
	return o.PublicIP, true
}

// HasPublicIP returns a boolean if a field has been set.
func (o *EndpointAgent) HasPublicIP() bool {
	if o != nil && !utils.IsNil(o.PublicIP) {
		return true
	}

	return false
}

// SetPublicIP gets a reference to the given string and assigns it to the PublicIP field.
func (o *EndpointAgent) SetPublicIP(v string) {
	o.PublicIP = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *EndpointAgent) GetLocation() EndpointAgentLocation {
	if o == nil || utils.IsNil(o.Location) {
		var ret EndpointAgentLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetLocationOk() (*EndpointAgentLocation, bool) {
	if o == nil || utils.IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *EndpointAgent) HasLocation() bool {
	if o != nil && !utils.IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given EndpointAgentLocation and assigns it to the Location field.
func (o *EndpointAgent) SetLocation(v EndpointAgentLocation) {
	o.Location = &v
}

// GetClients returns the Clients field value if set, zero value otherwise.
func (o *EndpointAgent) GetClients() []EndpointClient {
	if o == nil || utils.IsNil(o.Clients) {
		var ret []EndpointClient
		return ret
	}
	return o.Clients
}

// GetClientsOk returns a tuple with the Clients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetClientsOk() ([]EndpointClient, bool) {
	if o == nil || utils.IsNil(o.Clients) {
		return nil, false
	}
	return o.Clients, true
}

// HasClients returns a boolean if a field has been set.
func (o *EndpointAgent) HasClients() bool {
	if o != nil && !utils.IsNil(o.Clients) {
		return true
	}

	return false
}

// SetClients gets a reference to the given []EndpointClient and assigns it to the Clients field.
func (o *EndpointAgent) SetClients(v []EndpointClient) {
	o.Clients = v
}

// GetTotalMemory returns the TotalMemory field value if set, zero value otherwise.
func (o *EndpointAgent) GetTotalMemory() string {
	if o == nil || utils.IsNil(o.TotalMemory) {
		var ret string
		return ret
	}
	return *o.TotalMemory
}

// GetTotalMemoryOk returns a tuple with the TotalMemory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetTotalMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TotalMemory) {
		return nil, false
	}
	return o.TotalMemory, true
}

// HasTotalMemory returns a boolean if a field has been set.
func (o *EndpointAgent) HasTotalMemory() bool {
	if o != nil && !utils.IsNil(o.TotalMemory) {
		return true
	}

	return false
}

// SetTotalMemory gets a reference to the given string and assigns it to the TotalMemory field.
func (o *EndpointAgent) SetTotalMemory(v string) {
	o.TotalMemory = &v
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *EndpointAgent) GetAgentType() string {
	if o == nil || utils.IsNil(o.AgentType) {
		var ret string
		return ret
	}
	return *o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetAgentTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.AgentType) {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *EndpointAgent) HasAgentType() bool {
	if o != nil && !utils.IsNil(o.AgentType) {
		return true
	}

	return false
}

// SetAgentType gets a reference to the given string and assigns it to the AgentType field.
func (o *EndpointAgent) SetAgentType(v string) {
	o.AgentType = &v
}

// GetVpnProfiles returns the VpnProfiles field value if set, zero value otherwise.
func (o *EndpointAgent) GetVpnProfiles() []EndpointVpnProfile {
	if o == nil || utils.IsNil(o.VpnProfiles) {
		var ret []EndpointVpnProfile
		return ret
	}
	return o.VpnProfiles
}

// GetVpnProfilesOk returns a tuple with the VpnProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetVpnProfilesOk() ([]EndpointVpnProfile, bool) {
	if o == nil || utils.IsNil(o.VpnProfiles) {
		return nil, false
	}
	return o.VpnProfiles, true
}

// HasVpnProfiles returns a boolean if a field has been set.
func (o *EndpointAgent) HasVpnProfiles() bool {
	if o != nil && !utils.IsNil(o.VpnProfiles) {
		return true
	}

	return false
}

// SetVpnProfiles gets a reference to the given []EndpointVpnProfile and assigns it to the VpnProfiles field.
func (o *EndpointAgent) SetVpnProfiles(v []EndpointVpnProfile) {
	o.VpnProfiles = v
}

// GetNetworkInterfaceProfiles returns the NetworkInterfaceProfiles field value if set, zero value otherwise.
func (o *EndpointAgent) GetNetworkInterfaceProfiles() []InterfaceProfile {
	if o == nil || utils.IsNil(o.NetworkInterfaceProfiles) {
		var ret []InterfaceProfile
		return ret
	}
	return o.NetworkInterfaceProfiles
}

// GetNetworkInterfaceProfilesOk returns a tuple with the NetworkInterfaceProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetNetworkInterfaceProfilesOk() ([]InterfaceProfile, bool) {
	if o == nil || utils.IsNil(o.NetworkInterfaceProfiles) {
		return nil, false
	}
	return o.NetworkInterfaceProfiles, true
}

// HasNetworkInterfaceProfiles returns a boolean if a field has been set.
func (o *EndpointAgent) HasNetworkInterfaceProfiles() bool {
	if o != nil && !utils.IsNil(o.NetworkInterfaceProfiles) {
		return true
	}

	return false
}

// SetNetworkInterfaceProfiles gets a reference to the given []InterfaceProfile and assigns it to the NetworkInterfaceProfiles field.
func (o *EndpointAgent) SetNetworkInterfaceProfiles(v []InterfaceProfile) {
	o.NetworkInterfaceProfiles = v
}

// GetAsnDetails returns the AsnDetails field value if set, zero value otherwise.
func (o *EndpointAgent) GetAsnDetails() EndpointAsnDetails {
	if o == nil || utils.IsNil(o.AsnDetails) {
		var ret EndpointAsnDetails
		return ret
	}
	return *o.AsnDetails
}

// GetAsnDetailsOk returns a tuple with the AsnDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetAsnDetailsOk() (*EndpointAsnDetails, bool) {
	if o == nil || utils.IsNil(o.AsnDetails) {
		return nil, false
	}
	return o.AsnDetails, true
}

// HasAsnDetails returns a boolean if a field has been set.
func (o *EndpointAgent) HasAsnDetails() bool {
	if o != nil && !utils.IsNil(o.AsnDetails) {
		return true
	}

	return false
}

// SetAsnDetails gets a reference to the given EndpointAsnDetails and assigns it to the AsnDetails field.
func (o *EndpointAgent) SetAsnDetails(v EndpointAsnDetails) {
	o.AsnDetails = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *EndpointAgent) GetLicenseType() AgentLicenseType {
	if o == nil || utils.IsNil(o.LicenseType) {
		var ret AgentLicenseType
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetLicenseTypeOk() (*AgentLicenseType, bool) {
	if o == nil || utils.IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *EndpointAgent) HasLicenseType() bool {
	if o != nil && !utils.IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given AgentLicenseType and assigns it to the LicenseType field.
func (o *EndpointAgent) SetLicenseType(v AgentLicenseType) {
	o.LicenseType = &v
}

// GetTcpDriverAvailable returns the TcpDriverAvailable field value if set, zero value otherwise.
func (o *EndpointAgent) GetTcpDriverAvailable() bool {
	if o == nil || utils.IsNil(o.TcpDriverAvailable) {
		var ret bool
		return ret
	}
	return *o.TcpDriverAvailable
}

// GetTcpDriverAvailableOk returns a tuple with the TcpDriverAvailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetTcpDriverAvailableOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.TcpDriverAvailable) {
		return nil, false
	}
	return o.TcpDriverAvailable, true
}

// HasTcpDriverAvailable returns a boolean if a field has been set.
func (o *EndpointAgent) HasTcpDriverAvailable() bool {
	if o != nil && !utils.IsNil(o.TcpDriverAvailable) {
		return true
	}

	return false
}

// SetTcpDriverAvailable gets a reference to the given bool and assigns it to the TcpDriverAvailable field.
func (o *EndpointAgent) SetTcpDriverAvailable(v bool) {
	o.TcpDriverAvailable = &v
}

// GetNpcapVersion returns the NpcapVersion field value if set, zero value otherwise.
func (o *EndpointAgent) GetNpcapVersion() string {
	if o == nil || utils.IsNil(o.NpcapVersion) {
		var ret string
		return ret
	}
	return *o.NpcapVersion
}

// GetNpcapVersionOk returns a tuple with the NpcapVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetNpcapVersionOk() (*string, bool) {
	if o == nil || utils.IsNil(o.NpcapVersion) {
		return nil, false
	}
	return o.NpcapVersion, true
}

// HasNpcapVersion returns a boolean if a field has been set.
func (o *EndpointAgent) HasNpcapVersion() bool {
	if o != nil && !utils.IsNil(o.NpcapVersion) {
		return true
	}

	return false
}

// SetNpcapVersion gets a reference to the given string and assigns it to the NpcapVersion field.
func (o *EndpointAgent) SetNpcapVersion(v string) {
	o.NpcapVersion = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *EndpointAgent) GetLinks() SelfLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret SelfLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointAgent) GetLinksOk() (*SelfLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *EndpointAgent) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLinks and assigns it to the Links field.
func (o *EndpointAgent) SetLinks(v SelfLinks) {
	o.Links = &v
}

func (o EndpointAgent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointAgent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Aid) {
		toSerialize["aid"] = o.Aid
	}
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.ComputerName) {
		toSerialize["computerName"] = o.ComputerName
	}
	if !utils.IsNil(o.OsVersion) {
		toSerialize["osVersion"] = o.OsVersion
	}
	if !utils.IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !utils.IsNil(o.KernelVersion) {
		toSerialize["kernelVersion"] = o.KernelVersion
	}
	if !utils.IsNil(o.Manufacturer) {
		toSerialize["manufacturer"] = o.Manufacturer
	}
	if !utils.IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !utils.IsNil(o.LastSeen) {
		toSerialize["lastSeen"] = o.LastSeen
	}
	if !utils.IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !utils.IsNil(o.Deleted) {
		toSerialize["deleted"] = o.Deleted
	}
	if !utils.IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !utils.IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !utils.IsNil(o.NumberOfClients) {
		toSerialize["numberOfClients"] = o.NumberOfClients
	}
	if !utils.IsNil(o.PublicIP) {
		toSerialize["publicIP"] = o.PublicIP
	}
	if !utils.IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !utils.IsNil(o.Clients) {
		toSerialize["clients"] = o.Clients
	}
	if !utils.IsNil(o.TotalMemory) {
		toSerialize["totalMemory"] = o.TotalMemory
	}
	if !utils.IsNil(o.AgentType) {
		toSerialize["agentType"] = o.AgentType
	}
	if !utils.IsNil(o.VpnProfiles) {
		toSerialize["vpnProfiles"] = o.VpnProfiles
	}
	if !utils.IsNil(o.NetworkInterfaceProfiles) {
		toSerialize["networkInterfaceProfiles"] = o.NetworkInterfaceProfiles
	}
	if !utils.IsNil(o.AsnDetails) {
		toSerialize["asnDetails"] = o.AsnDetails
	}
	if !utils.IsNil(o.LicenseType) {
		toSerialize["licenseType"] = o.LicenseType
	}
	if !utils.IsNil(o.TcpDriverAvailable) {
		toSerialize["tcpDriverAvailable"] = o.TcpDriverAvailable
	}
	if !utils.IsNil(o.NpcapVersion) {
		toSerialize["npcapVersion"] = o.NpcapVersion
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	return toSerialize, nil
}

type NullableEndpointAgent struct {
	value *EndpointAgent
	isSet bool
}

func (v NullableEndpointAgent) Get() *EndpointAgent {
	return v.value
}

func (v *NullableEndpointAgent) Set(val *EndpointAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointAgent(val *EndpointAgent) *NullableEndpointAgent {
	return &NullableEndpointAgent{value: val, isSet: true}
}

func (v NullableEndpointAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


