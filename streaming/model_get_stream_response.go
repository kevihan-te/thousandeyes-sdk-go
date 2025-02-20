/*
ThousandEyes for OpenTelemetry API

ThousandEyes for OpenTelemetry provides machine-to-machine integration between ThousandEyes and its customers. It allows you to export ThousandEyes telemetry data in OTel format, which is widely used in the industry. With ThousandEyes for OTel, you can leverage frameworks widely used in the observability domain - such as Splunk, Grafana, and Honeycomb - to capture and analyze ThousandEyes data. Any client that supports OTel can use ThousandEyes for OpenTelemetry.  ThousandEyes for OTel is made up of the following components:  * Data streaming APIs that you can use to configure and enable your ThousandEyes tests with OTel-compatible streams, in particular to configure how ThousandEyes telemetry data is exported to client integrations. * A set of streaming pipelines called _collectors_ that actively fetch ThousandEyes network test data, enrich the data with some additional detail, filter, and push the data to the customer-configured endpoints, depending on what you configure via the public APIs. * Third-party OTel collectors that receive, transform, filter, and export different metrics to client applications such as AppD, or any other OTel-capable client configuration.  For more information about ThousandEyes for OpenTelemetry, see the [documentation](https://docs.thousandeyes.com/product-documentation/api/opentelemetry). 

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package streaming

import (
	"encoding/json"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/internal/utils"
)

// checks if the GetStreamResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &GetStreamResponse{}

// GetStreamResponse struct for GetStreamResponse
type GetStreamResponse struct {
	// The data stream ID
	Id *string `json:"id,omitempty"`
	// Flag to enable or disable the stream integration.
	Enabled *bool `json:"enabled,omitempty"`
	Links *StreamLinks `json:"_links,omitempty"`
	Type *StreamType `json:"type,omitempty"`
	Signal *Signal `json:"signal,omitempty"`
	EndpointType *EndpointType `json:"endpointType,omitempty"`
	// The URL ThousandEyes sends data stream to. For a URL to be valid, it needs to: - Be syntactically correct. - Be reachable. - Use the HTTPS protocol. - When using the `grpc` endpointType, streamEndpointUrl cannot contain paths:     - Valid . `grpc` - `https://example.com`     - Invalid . `grpc` - `https://example.com/collector`.     - Valid . `http` - `https://example.com/collector`.  - When using the `http` endpointType, the operation must match the exact final full URL (including the path if there is one) to which the data will be sent. Examples below:     - `https://api.honeycomb.io:443/v1/metrics`     - `https://ingest.eu0.signalfx.com/v2/datapoint/otlp`
	StreamEndpointUrl *string `json:"streamEndpointUrl,omitempty"`
	DataModelVersion *DataModelVersion `json:"dataModelVersion,omitempty"`
	// Custom headers. **Note**: When using the `splunk-hec` `type`, the `customHeaders` must contain just one element with the key `token` and the value of the *Splunk HEC Token*.
	CustomHeaders *map[string]string `json:"customHeaders,omitempty"`
	// A collection of tags that determine what tests are included in the data stream. These tag values are also included as attributes in the data stream metrics.
	TagMatch []TagMatch `json:"tagMatch,omitempty"`
	// A collection of tests to be included in the data stream.
	TestMatch []TestMatch `json:"testMatch,omitempty"`
	Filters *Filters `json:"filters,omitempty"`
	ExporterConfig *ExporterConfig `json:"exporterConfig,omitempty"`
	AuditOperation *AuditOperationWithUpdate `json:"auditOperation,omitempty"`
}

// NewGetStreamResponse instantiates a new GetStreamResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStreamResponse() *GetStreamResponse {
	this := GetStreamResponse{}
	var signal Signal = SIGNAL_METRIC
	this.Signal = &signal
	var endpointType EndpointType = ENDPOINTTYPE_GRPC
	this.EndpointType = &endpointType
	var dataModelVersion DataModelVersion = DATAMODELVERSION_V2
	this.DataModelVersion = &dataModelVersion
	return &this
}

// NewGetStreamResponseWithDefaults instantiates a new GetStreamResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStreamResponseWithDefaults() *GetStreamResponse {
	this := GetStreamResponse{}
	var signal Signal = SIGNAL_METRIC
	this.Signal = &signal
	var endpointType EndpointType = ENDPOINTTYPE_GRPC
	this.EndpointType = &endpointType
	var dataModelVersion DataModelVersion = DATAMODELVERSION_V2
	this.DataModelVersion = &dataModelVersion
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetStreamResponse) GetId() string {
	if o == nil || utils.IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetStreamResponse) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetStreamResponse) SetId(v string) {
	o.Id = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GetStreamResponse) GetEnabled() bool {
	if o == nil || utils.IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetEnabledOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GetStreamResponse) HasEnabled() bool {
	if o != nil && !utils.IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GetStreamResponse) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GetStreamResponse) GetLinks() StreamLinks {
	if o == nil || utils.IsNil(o.Links) {
		var ret StreamLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetLinksOk() (*StreamLinks, bool) {
	if o == nil || utils.IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GetStreamResponse) HasLinks() bool {
	if o != nil && !utils.IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given StreamLinks and assigns it to the Links field.
func (o *GetStreamResponse) SetLinks(v StreamLinks) {
	o.Links = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetStreamResponse) GetType() StreamType {
	if o == nil || utils.IsNil(o.Type) {
		var ret StreamType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetTypeOk() (*StreamType, bool) {
	if o == nil || utils.IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetStreamResponse) HasType() bool {
	if o != nil && !utils.IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given StreamType and assigns it to the Type field.
func (o *GetStreamResponse) SetType(v StreamType) {
	o.Type = &v
}

// GetSignal returns the Signal field value if set, zero value otherwise.
func (o *GetStreamResponse) GetSignal() Signal {
	if o == nil || utils.IsNil(o.Signal) {
		var ret Signal
		return ret
	}
	return *o.Signal
}

// GetSignalOk returns a tuple with the Signal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetSignalOk() (*Signal, bool) {
	if o == nil || utils.IsNil(o.Signal) {
		return nil, false
	}
	return o.Signal, true
}

// HasSignal returns a boolean if a field has been set.
func (o *GetStreamResponse) HasSignal() bool {
	if o != nil && !utils.IsNil(o.Signal) {
		return true
	}

	return false
}

// SetSignal gets a reference to the given Signal and assigns it to the Signal field.
func (o *GetStreamResponse) SetSignal(v Signal) {
	o.Signal = &v
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise.
func (o *GetStreamResponse) GetEndpointType() EndpointType {
	if o == nil || utils.IsNil(o.EndpointType) {
		var ret EndpointType
		return ret
	}
	return *o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetEndpointTypeOk() (*EndpointType, bool) {
	if o == nil || utils.IsNil(o.EndpointType) {
		return nil, false
	}
	return o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *GetStreamResponse) HasEndpointType() bool {
	if o != nil && !utils.IsNil(o.EndpointType) {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given EndpointType and assigns it to the EndpointType field.
func (o *GetStreamResponse) SetEndpointType(v EndpointType) {
	o.EndpointType = &v
}

// GetStreamEndpointUrl returns the StreamEndpointUrl field value if set, zero value otherwise.
func (o *GetStreamResponse) GetStreamEndpointUrl() string {
	if o == nil || utils.IsNil(o.StreamEndpointUrl) {
		var ret string
		return ret
	}
	return *o.StreamEndpointUrl
}

// GetStreamEndpointUrlOk returns a tuple with the StreamEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetStreamEndpointUrlOk() (*string, bool) {
	if o == nil || utils.IsNil(o.StreamEndpointUrl) {
		return nil, false
	}
	return o.StreamEndpointUrl, true
}

// HasStreamEndpointUrl returns a boolean if a field has been set.
func (o *GetStreamResponse) HasStreamEndpointUrl() bool {
	if o != nil && !utils.IsNil(o.StreamEndpointUrl) {
		return true
	}

	return false
}

// SetStreamEndpointUrl gets a reference to the given string and assigns it to the StreamEndpointUrl field.
func (o *GetStreamResponse) SetStreamEndpointUrl(v string) {
	o.StreamEndpointUrl = &v
}

// GetDataModelVersion returns the DataModelVersion field value if set, zero value otherwise.
func (o *GetStreamResponse) GetDataModelVersion() DataModelVersion {
	if o == nil || utils.IsNil(o.DataModelVersion) {
		var ret DataModelVersion
		return ret
	}
	return *o.DataModelVersion
}

// GetDataModelVersionOk returns a tuple with the DataModelVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetDataModelVersionOk() (*DataModelVersion, bool) {
	if o == nil || utils.IsNil(o.DataModelVersion) {
		return nil, false
	}
	return o.DataModelVersion, true
}

// HasDataModelVersion returns a boolean if a field has been set.
func (o *GetStreamResponse) HasDataModelVersion() bool {
	if o != nil && !utils.IsNil(o.DataModelVersion) {
		return true
	}

	return false
}

// SetDataModelVersion gets a reference to the given DataModelVersion and assigns it to the DataModelVersion field.
func (o *GetStreamResponse) SetDataModelVersion(v DataModelVersion) {
	o.DataModelVersion = &v
}

// GetCustomHeaders returns the CustomHeaders field value if set, zero value otherwise.
func (o *GetStreamResponse) GetCustomHeaders() map[string]string {
	if o == nil || utils.IsNil(o.CustomHeaders) {
		var ret map[string]string
		return ret
	}
	return *o.CustomHeaders
}

// GetCustomHeadersOk returns a tuple with the CustomHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetCustomHeadersOk() (*map[string]string, bool) {
	if o == nil || utils.IsNil(o.CustomHeaders) {
		return nil, false
	}
	return o.CustomHeaders, true
}

// HasCustomHeaders returns a boolean if a field has been set.
func (o *GetStreamResponse) HasCustomHeaders() bool {
	if o != nil && !utils.IsNil(o.CustomHeaders) {
		return true
	}

	return false
}

// SetCustomHeaders gets a reference to the given map[string]string and assigns it to the CustomHeaders field.
func (o *GetStreamResponse) SetCustomHeaders(v map[string]string) {
	o.CustomHeaders = &v
}

// GetTagMatch returns the TagMatch field value if set, zero value otherwise.
func (o *GetStreamResponse) GetTagMatch() []TagMatch {
	if o == nil || utils.IsNil(o.TagMatch) {
		var ret []TagMatch
		return ret
	}
	return o.TagMatch
}

// GetTagMatchOk returns a tuple with the TagMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetTagMatchOk() ([]TagMatch, bool) {
	if o == nil || utils.IsNil(o.TagMatch) {
		return nil, false
	}
	return o.TagMatch, true
}

// HasTagMatch returns a boolean if a field has been set.
func (o *GetStreamResponse) HasTagMatch() bool {
	if o != nil && !utils.IsNil(o.TagMatch) {
		return true
	}

	return false
}

// SetTagMatch gets a reference to the given []TagMatch and assigns it to the TagMatch field.
func (o *GetStreamResponse) SetTagMatch(v []TagMatch) {
	o.TagMatch = v
}

// GetTestMatch returns the TestMatch field value if set, zero value otherwise.
func (o *GetStreamResponse) GetTestMatch() []TestMatch {
	if o == nil || utils.IsNil(o.TestMatch) {
		var ret []TestMatch
		return ret
	}
	return o.TestMatch
}

// GetTestMatchOk returns a tuple with the TestMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetTestMatchOk() ([]TestMatch, bool) {
	if o == nil || utils.IsNil(o.TestMatch) {
		return nil, false
	}
	return o.TestMatch, true
}

// HasTestMatch returns a boolean if a field has been set.
func (o *GetStreamResponse) HasTestMatch() bool {
	if o != nil && !utils.IsNil(o.TestMatch) {
		return true
	}

	return false
}

// SetTestMatch gets a reference to the given []TestMatch and assigns it to the TestMatch field.
func (o *GetStreamResponse) SetTestMatch(v []TestMatch) {
	o.TestMatch = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GetStreamResponse) GetFilters() Filters {
	if o == nil || utils.IsNil(o.Filters) {
		var ret Filters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetFiltersOk() (*Filters, bool) {
	if o == nil || utils.IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GetStreamResponse) HasFilters() bool {
	if o != nil && !utils.IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given Filters and assigns it to the Filters field.
func (o *GetStreamResponse) SetFilters(v Filters) {
	o.Filters = &v
}

// GetExporterConfig returns the ExporterConfig field value if set, zero value otherwise.
func (o *GetStreamResponse) GetExporterConfig() ExporterConfig {
	if o == nil || utils.IsNil(o.ExporterConfig) {
		var ret ExporterConfig
		return ret
	}
	return *o.ExporterConfig
}

// GetExporterConfigOk returns a tuple with the ExporterConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetExporterConfigOk() (*ExporterConfig, bool) {
	if o == nil || utils.IsNil(o.ExporterConfig) {
		return nil, false
	}
	return o.ExporterConfig, true
}

// HasExporterConfig returns a boolean if a field has been set.
func (o *GetStreamResponse) HasExporterConfig() bool {
	if o != nil && !utils.IsNil(o.ExporterConfig) {
		return true
	}

	return false
}

// SetExporterConfig gets a reference to the given ExporterConfig and assigns it to the ExporterConfig field.
func (o *GetStreamResponse) SetExporterConfig(v ExporterConfig) {
	o.ExporterConfig = &v
}

// GetAuditOperation returns the AuditOperation field value if set, zero value otherwise.
func (o *GetStreamResponse) GetAuditOperation() AuditOperationWithUpdate {
	if o == nil || utils.IsNil(o.AuditOperation) {
		var ret AuditOperationWithUpdate
		return ret
	}
	return *o.AuditOperation
}

// GetAuditOperationOk returns a tuple with the AuditOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStreamResponse) GetAuditOperationOk() (*AuditOperationWithUpdate, bool) {
	if o == nil || utils.IsNil(o.AuditOperation) {
		return nil, false
	}
	return o.AuditOperation, true
}

// HasAuditOperation returns a boolean if a field has been set.
func (o *GetStreamResponse) HasAuditOperation() bool {
	if o != nil && !utils.IsNil(o.AuditOperation) {
		return true
	}

	return false
}

// SetAuditOperation gets a reference to the given AuditOperationWithUpdate and assigns it to the AuditOperation field.
func (o *GetStreamResponse) SetAuditOperation(v AuditOperationWithUpdate) {
	o.AuditOperation = &v
}

func (o GetStreamResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStreamResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !utils.IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !utils.IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !utils.IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !utils.IsNil(o.Signal) {
		toSerialize["signal"] = o.Signal
	}
	if !utils.IsNil(o.EndpointType) {
		toSerialize["endpointType"] = o.EndpointType
	}
	if !utils.IsNil(o.StreamEndpointUrl) {
		toSerialize["streamEndpointUrl"] = o.StreamEndpointUrl
	}
	if !utils.IsNil(o.DataModelVersion) {
		toSerialize["dataModelVersion"] = o.DataModelVersion
	}
	if !utils.IsNil(o.CustomHeaders) {
		toSerialize["customHeaders"] = o.CustomHeaders
	}
	if !utils.IsNil(o.TagMatch) {
		toSerialize["tagMatch"] = o.TagMatch
	}
	if !utils.IsNil(o.TestMatch) {
		toSerialize["testMatch"] = o.TestMatch
	}
	if !utils.IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !utils.IsNil(o.ExporterConfig) {
		toSerialize["exporterConfig"] = o.ExporterConfig
	}
	if !utils.IsNil(o.AuditOperation) {
		toSerialize["auditOperation"] = o.AuditOperation
	}
	return toSerialize, nil
}

type NullableGetStreamResponse struct {
	value *GetStreamResponse
	isSet bool
}

func (v NullableGetStreamResponse) Get() *GetStreamResponse {
	return v.value
}

func (v *NullableGetStreamResponse) Set(val *GetStreamResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStreamResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStreamResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStreamResponse(val *GetStreamResponse) *NullableGetStreamResponse {
	return &NullableGetStreamResponse{value: val, isSet: true}
}

func (v NullableGetStreamResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStreamResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


